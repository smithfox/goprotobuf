// Code generated by protoc-gen-go.
// source: my_test/test.proto
// DO NOT EDIT!

package my_test

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import multi2 "multi/multi1.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type HatType int32

const (
	HatType_FEDORA HatType = 1
	HatType_FEZ    HatType = 2
)

var HatType_name = map[int32]string{
	1: "FEDORA",
	2: "FEZ",
}
var HatType_value = map[string]int32{
	"FEDORA": 1,
	"FEZ":    2,
}

func (x HatType) Enum() *HatType {
	p := new(HatType)
	*p = x
	return p
}
func (x HatType) String() string {
	return proto.EnumName(HatType_name, int32(x))
}
func (x HatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *HatType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HatType_value, data, "HatType")
	if err != nil {
		return err
	}
	*x = HatType(value)
	return nil
}

type Days int32

const (
	Days_MONDAY  Days = 1
	Days_TUESDAY Days = 2
	Days_LUNDI   Days = 1
)

var Days_name = map[int32]string{
	1: "MONDAY",
	2: "TUESDAY",
	// Duplicate value: 1: "LUNDI",
}
var Days_value = map[string]int32{
	"MONDAY":  1,
	"TUESDAY": 2,
	"LUNDI":   1,
}

func (x Days) Enum() *Days {
	p := new(Days)
	*p = x
	return p
}
func (x Days) String() string {
	return proto.EnumName(Days_name, int32(x))
}
func (x Days) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Days) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Days_value, data, "Days")
	if err != nil {
		return err
	}
	*x = Days(value)
	return nil
}

type Request_Color int32

const (
	Request_RED   Request_Color = 0
	Request_GREEN Request_Color = 1
	Request_BLUE  Request_Color = 2
)

var Request_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var Request_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Request_Color) Enum() *Request_Color {
	p := new(Request_Color)
	*p = x
	return p
}
func (x Request_Color) String() string {
	return proto.EnumName(Request_Color_name, int32(x))
}
func (x Request_Color) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Request_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Request_Color_value, data, "Request_Color")
	if err != nil {
		return err
	}
	*x = Request_Color(value)
	return nil
}

type Reply_Entry_Game int32

const (
	Reply_Entry_FOOTBALL Reply_Entry_Game = 1
	Reply_Entry_TENNIS   Reply_Entry_Game = 2
)

var Reply_Entry_Game_name = map[int32]string{
	1: "FOOTBALL",
	2: "TENNIS",
}
var Reply_Entry_Game_value = map[string]int32{
	"FOOTBALL": 1,
	"TENNIS":   2,
}

func (x Reply_Entry_Game) Enum() *Reply_Entry_Game {
	p := new(Reply_Entry_Game)
	*p = x
	return p
}
func (x Reply_Entry_Game) String() string {
	return proto.EnumName(Reply_Entry_Game_name, int32(x))
}
func (x Reply_Entry_Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Reply_Entry_Game) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Reply_Entry_Game_value, data, "Reply_Entry_Game")
	if err != nil {
		return err
	}
	*x = Reply_Entry_Game(value)
	return nil
}

type Request struct {
	Key              []int64            `protobuf:"varint,1,rep,name=key" json:"key,omitempty"`
	Hue              *Request_Color     `protobuf:"varint,3,opt,name=hue,enum=my.test.Request_Color" json:"hue,omitempty"`
	Hat              *HatType           `protobuf:"varint,4,opt,name=hat,enum=my.test.HatType,def=1" json:"hat,omitempty"`
	Deadline         *float32           `protobuf:"fixed32,7,opt,name=deadline,def=inf" json:"deadline,omitempty"`
	Somegroup        *Request_SomeGroup `protobuf:"group,8,opt,name=SomeGroup" json:"somegroup,omitempty"`
	Reset_           *int32             `protobuf:"varint,12,opt,name=reset" json:"reset,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (this *Request) Reset()         { *this = Request{} }
func (this *Request) String() string { return proto.CompactTextString(this) }
func (*Request) ProtoMessage()       {}

const Default_Request_Hat HatType = HatType_FEDORA

var Default_Request_Deadline float32 = float32(math.Inf(1))

func (this *Request) GetKey() []int64 {
	if this != nil {
		return this.Key
	}
	return nil
}

func (this *Request) GetHue() Request_Color {
	if this != nil && this.Hue != nil {
		return *this.Hue
	}
	return 0
}

func (this *Request) GetHat() HatType {
	if this != nil && this.Hat != nil {
		return *this.Hat
	}
	return Default_Request_Hat
}

func (this *Request) GetDeadline() float32 {
	if this != nil && this.Deadline != nil {
		return *this.Deadline
	}
	return Default_Request_Deadline
}

func (this *Request) GetSomegroup() *Request_SomeGroup {
	if this != nil {
		return this.Somegroup
	}
	return nil
}

func (this *Request) GetReset_() int32 {
	if this != nil && this.Reset_ != nil {
		return *this.Reset_
	}
	return 0
}

type Request_SomeGroup struct {
	GroupField *int32 `protobuf:"varint,9,opt,name=group_field" json:"group_field,omitempty"`
}

func (this *Request_SomeGroup) Reset() { *this = Request_SomeGroup{} }

func (this *Request_SomeGroup) GetGroupField() int32 {
	if this != nil && this.GroupField != nil {
		return *this.GroupField
	}
	return 0
}

type Reply struct {
	Found            []*Reply_Entry            `protobuf:"bytes,1,rep,name=found" json:"found,omitempty"`
	CompactKeys      []int32                   `protobuf:"varint,2,rep,packed,name=compact_keys" json:"compact_keys,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (this *Reply) Reset()         { *this = Reply{} }
func (this *Reply) String() string { return proto.CompactTextString(this) }
func (*Reply) ProtoMessage()       {}

var extRange_Reply = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Reply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Reply
}
func (this *Reply) ExtensionMap() map[int32]proto.Extension {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32]proto.Extension)
	}
	return this.XXX_extensions
}

func (this *Reply) GetFound() []*Reply_Entry {
	if this != nil {
		return this.Found
	}
	return nil
}

func (this *Reply) GetCompactKeys() []int32 {
	if this != nil {
		return this.CompactKeys
	}
	return nil
}

type Reply_Entry struct {
	KeyThatNeeds_1234Camel_CasIng *int64 `protobuf:"varint,1,req,name=key_that_needs_1234camel_CasIng" json:"key_that_needs_1234camel_CasIng,omitempty"`
	Value                         *int64 `protobuf:"varint,2,opt,name=value,def=7" json:"value,omitempty"`
	XMyFieldName_2                *int64 `protobuf:"varint,3,opt,name=_my_field_name_2" json:"_my_field_name_2,omitempty"`
	XXX_unrecognized              []byte `json:"-"`
}

func (this *Reply_Entry) Reset()         { *this = Reply_Entry{} }
func (this *Reply_Entry) String() string { return proto.CompactTextString(this) }
func (*Reply_Entry) ProtoMessage()       {}

const Default_Reply_Entry_Value int64 = 7

func (this *Reply_Entry) GetKeyThatNeeds_1234Camel_CasIng() int64 {
	if this != nil && this.KeyThatNeeds_1234Camel_CasIng != nil {
		return *this.KeyThatNeeds_1234Camel_CasIng
	}
	return 0
}

func (this *Reply_Entry) GetValue() int64 {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return Default_Reply_Entry_Value
}

func (this *Reply_Entry) GetXMyFieldName_2() int64 {
	if this != nil && this.XMyFieldName_2 != nil {
		return *this.XMyFieldName_2
	}
	return 0
}

type ReplyExtensions struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *ReplyExtensions) Reset()         { *this = ReplyExtensions{} }
func (this *ReplyExtensions) String() string { return proto.CompactTextString(this) }
func (*ReplyExtensions) ProtoMessage()       {}

var E_ReplyExtensions_Time = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*float64)(nil),
	Field:         101,
	Name:          "my.test.ReplyExtensions.time",
	Tag:           "fixed64,101,opt,name=time",
}

type OldReply struct {
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (this *OldReply) Reset()         { *this = OldReply{} }
func (this *OldReply) String() string { return proto.CompactTextString(this) }
func (*OldReply) ProtoMessage()       {}

func (this *OldReply) Marshal() ([]byte, error) {
	return proto.MarshalMessageSet(this.ExtensionMap())
}
func (this *OldReply) Unmarshal(buf []byte) error {
	return proto.UnmarshalMessageSet(buf, this.ExtensionMap())
}

// ensure OldReply satisfies proto.Marshaler and proto.Unmarshaler
var _ proto.Marshaler = (*OldReply)(nil)
var _ proto.Unmarshaler = (*OldReply)(nil)

var extRange_OldReply = []proto.ExtensionRange{
	{100, 536870911},
}

func (*OldReply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OldReply
}
func (this *OldReply) ExtensionMap() map[int32]proto.Extension {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32]proto.Extension)
	}
	return this.XXX_extensions
}

var E_Tag = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*string)(nil),
	Field:         103,
	Name:          "my.test.tag",
	Tag:           "bytes,103,opt,name=tag",
}

func init() {
	proto.RegisterEnum("my.test.HatType", HatType_name, HatType_value)
	proto.RegisterEnum("my.test.Days", Days_name, Days_value)
	proto.RegisterEnum("my.test.Request_Color", Request_Color_name, Request_Color_value)
	proto.RegisterEnum("my.test.Reply_Entry_Game", Reply_Entry_Game_name, Reply_Entry_Game_value)
	proto.RegisterExtension(E_ReplyExtensions_Time)
	proto.RegisterExtension(E_Tag)
}
