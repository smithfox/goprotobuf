// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 The Go Authors.  All rights reserved.
// http://github.com/smithfox/goprotobuf/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package proto

/*
 * Types and routines for supporting protocol buffer extensions.
 */

import (
	"errors"
	"reflect"
	"strconv"
	"sync"
)

// ErrMissingExtension is the error returned by GetExtension if the named extension is not in the message.
var ErrMissingExtension = errors.New("proto: missing extension")

// ExtensionRange represents a range of message extensions for a protocol buffer.
// Used in code generated by the protocol compiler.
type ExtensionRange struct {
	Start, End int32 // both inclusive
}

// extendableProto is an interface implemented by any protocol buffer that may be extended.
type extendableProto interface {
	Message
	ExtensionRangeArray() []ExtensionRange
	ExtensionMap() map[int32]Extension
}

var extendableProtoType = reflect.TypeOf((*extendableProto)(nil)).Elem()

// ExtensionDesc represents an extension specification.
// Used in generated code from the protocol compiler.
type ExtensionDesc struct {
	ExtendedType  Message     // nil pointer to the type that is being extended
	ExtensionType interface{} // nil pointer to the extension type
	Field         int32       // field number
	Name          string      // fully-qualified name of extension, for text formatting
	Tag           string      // protobuf tag style
}

func (ed *ExtensionDesc) repeated() bool {
	t := reflect.TypeOf(ed.ExtensionType)
	return t.Kind() == reflect.Slice && t.Elem().Kind() != reflect.Uint8
}

// Extension represents an extension in a message.
type Extension struct {
	// When an extension is stored in a message using SetExtension
	// only desc and value are set. When the message is marshaled
	// enc will be set to the encoded form of the message.
	//
	// When a message is unmarshaled and contains extensions, each
	// extension will have only enc set. When such an extension is
	// accessed using GetExtension (or GetExtensions) desc and value
	// will be set.
	desc  *ExtensionDesc
	value interface{}
	enc   []byte
}

// SetRawExtension is for testing only.
func SetRawExtension(base extendableProto, id int32, b []byte) {
	base.ExtensionMap()[id] = Extension{enc: b}
}

// isExtensionField returns true iff the given field number is in an extension range.
func isExtensionField(pb extendableProto, field int32) bool {
	for _, er := range pb.ExtensionRangeArray() {
		if er.Start <= field && field <= er.End {
			return true
		}
	}
	return false
}

// checkExtensionTypes checks that the given extension is valid for pb.
func checkExtensionTypes(pb extendableProto, extension *ExtensionDesc) error {
	// Check the extended type.
	if a, b := reflect.TypeOf(pb), reflect.TypeOf(extension.ExtendedType); a != b {
		return errors.New("proto: bad extended type; " + b.String() + " does not extend " + a.String())
	}
	// Check the range.
	if !isExtensionField(pb, extension.Field) {
		return errors.New("proto: bad extension number; not in declared ranges")
	}
	return nil
}

// extPropKey is sufficient to uniquely identify an extension.
type extPropKey struct {
	base  reflect.Type
	field int32
}

var extProp = struct {
	sync.RWMutex
	m map[extPropKey]*Properties
}{
	m: make(map[extPropKey]*Properties),
}

func extensionProperties(ed *ExtensionDesc) *Properties {
	key := extPropKey{base: reflect.TypeOf(ed.ExtendedType), field: ed.Field}

	extProp.RLock()
	if prop, ok := extProp.m[key]; ok {
		extProp.RUnlock()
		return prop
	}
	extProp.RUnlock()

	extProp.Lock()
	defer extProp.Unlock()
	// Check again.
	if prop, ok := extProp.m[key]; ok {
		return prop
	}

	prop := new(Properties)
	prop.Init(reflect.TypeOf(ed.ExtensionType), "unknown_name", ed.Tag, nil)
	extProp.m[key] = prop
	return prop
}

// encodeExtensionMap encodes any unmarshaled (unencoded) extensions in m.
func encodeExtensionMap(m map[int32]Extension) error {
	for k, e := range m {
		if e.value == nil || e.desc == nil {
			// Extension is only in its encoded form.
			continue
		}

		// We don't skip extensions that have an encoded form set,
		// because the extension value may have been mutated after
		// the last time this function was called.

		et := reflect.TypeOf(e.desc.ExtensionType)
		props := extensionProperties(e.desc)

		p := NewBuffer(nil)
		// If e.value has type T, the encoder expects a *struct{ X T }.
		// Pass a *T with a zero field and hope it all works out.
		x := reflect.New(et)
		x.Elem().Set(reflect.ValueOf(e.value))
		if err := props.enc(p, props, toStructPointer(x)); err != nil {
			return err
		}
		e.enc = p.buf
		m[k] = e
	}
	return nil
}

func sizeExtensionMap(m map[int32]Extension) (n int) {
	for _, e := range m {
		if e.value == nil || e.desc == nil {
			// Extension is only in its encoded form.
			n += len(e.enc)
			continue
		}

		// We don't skip extensions that have an encoded form set,
		// because the extension value may have been mutated after
		// the last time this function was called.

		et := reflect.TypeOf(e.desc.ExtensionType)
		props := extensionProperties(e.desc)

		// If e.value has type T, the encoder expects a *struct{ X T }.
		// Pass a *T with a zero field and hope it all works out.
		x := reflect.New(et)
		x.Elem().Set(reflect.ValueOf(e.value))
		n += props.size(props, toStructPointer(x))
	}
	return
}

// HasExtension returns whether the given extension is present in pb.
func HasExtension(pb extendableProto, extension *ExtensionDesc) bool {
	// TODO: Check types, field numbers, etc.?
	_, ok := pb.ExtensionMap()[extension.Field]
	return ok
}

// ClearExtension removes the given extension from pb.
func ClearExtension(pb extendableProto, extension *ExtensionDesc) {
	// TODO: Check types, field numbers, etc.?
	delete(pb.ExtensionMap(), extension.Field)
}

// GetExtension parses and returns the given extension of pb.
// If the extension is not present it returns ErrMissingExtension.
func GetExtension(pb extendableProto, extension *ExtensionDesc) (interface{}, error) {
	if err := checkExtensionTypes(pb, extension); err != nil {
		return nil, err
	}

	e, ok := pb.ExtensionMap()[extension.Field]
	if !ok {
		return nil, ErrMissingExtension
	}
	if e.value != nil {
		// Already decoded. Check the descriptor, though.
		if e.desc != extension {
			// This shouldn't happen. If it does, it means that
			// GetExtension was called twice with two different
			// descriptors with the same field number.
			return nil, errors.New("proto: descriptor conflict")
		}
		return e.value, nil
	}

	v, err := decodeExtension(e.enc, extension)
	if err != nil {
		return nil, err
	}

	// Remember the decoded version and drop the encoded version.
	// That way it is safe to mutate what we return.
	e.value = v
	e.desc = extension
	e.enc = nil
	return e.value, nil
}

// decodeExtension decodes an extension encoded in b.
func decodeExtension(b []byte, extension *ExtensionDesc) (interface{}, error) {
	o := NewBuffer(b)

	t := reflect.TypeOf(extension.ExtensionType)
	rep := extension.repeated()

	props := extensionProperties(extension)

	// t is a pointer to a struct, pointer to basic type or a slice.
	// Allocate a "field" to store the pointer/slice itself; the
	// pointer/slice will be stored here. We pass
	// the address of this field to props.dec.
	// This passes a zero field and a *t and lets props.dec
	// interpret it as a *struct{ x t }.
	value := reflect.New(t).Elem()

	for {
		// Discard wire type and field number varint. It isn't needed.
		if _, err := o.DecodeVarint(); err != nil {
			return nil, err
		}

		if err := props.dec(o, props, toStructPointer(value.Addr())); err != nil {
			return nil, err
		}

		if !rep || o.index >= len(o.buf) {
			break
		}
	}
	return value.Interface(), nil
}

// GetExtensions returns a slice of the extensions present in pb that are also listed in es.
// The returned slice has the same length as es; missing extensions will appear as nil elements.
func GetExtensions(pb Message, es []*ExtensionDesc) (extensions []interface{}, err error) {
	epb, ok := pb.(extendableProto)
	if !ok {
		err = errors.New("proto: not an extendable proto")
		return
	}
	extensions = make([]interface{}, len(es))
	for i, e := range es {
		extensions[i], err = GetExtension(epb, e)
		if err != nil {
			return
		}
	}
	return
}

// SetExtension sets the specified extension of pb to the specified value.
func SetExtension(pb extendableProto, extension *ExtensionDesc, value interface{}) error {
	if err := checkExtensionTypes(pb, extension); err != nil {
		return err
	}
	typ := reflect.TypeOf(extension.ExtensionType)
	if typ != reflect.TypeOf(value) {
		return errors.New("proto: bad extension value type")
	}

	pb.ExtensionMap()[extension.Field] = Extension{desc: extension, value: value}
	return nil
}

// A global registry of extensions.
// The generated code will register the generated descriptors by calling RegisterExtension.

var extensionMaps = make(map[reflect.Type]map[int32]*ExtensionDesc)

// RegisterExtension is called from the generated code.
func RegisterExtension(desc *ExtensionDesc) {
	st := reflect.TypeOf(desc.ExtendedType).Elem()
	m := extensionMaps[st]
	if m == nil {
		m = make(map[int32]*ExtensionDesc)
		extensionMaps[st] = m
	}
	if _, ok := m[desc.Field]; ok {
		panic("proto: duplicate extension registered: " + st.String() + " " + strconv.Itoa(int(desc.Field)))
	}
	m[desc.Field] = desc
}

// RegisteredExtensions returns a map of the registered extensions of a
// protocol buffer struct, indexed by the extension number.
// The argument pb should be a nil pointer to the struct type.
func RegisteredExtensions(pb Message) map[int32]*ExtensionDesc {
	return extensionMaps[reflect.TypeOf(pb).Elem()]
}
