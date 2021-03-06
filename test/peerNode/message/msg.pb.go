// Code generated by protoc-gen-go.
// source: msg.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	CG_Game_Start_CG
	CG_Game_Start_GC
*/
package message

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CG_Game_Start_GC_GameStartType int32

const (
	CG_Game_Start_GC_eGS_Success CG_Game_Start_GC_GameStartType = 0
	CG_Game_Start_GC_eGS_Failure CG_Game_Start_GC_GameStartType = 1
)

var CG_Game_Start_GC_GameStartType_name = map[int32]string{
	0: "eGS_Success",
	1: "eGS_Failure",
}
var CG_Game_Start_GC_GameStartType_value = map[string]int32{
	"eGS_Success": 0,
	"eGS_Failure": 1,
}

func (x CG_Game_Start_GC_GameStartType) Enum() *CG_Game_Start_GC_GameStartType {
	p := new(CG_Game_Start_GC_GameStartType)
	*p = x
	return p
}
func (x CG_Game_Start_GC_GameStartType) String() string {
	return proto.EnumName(CG_Game_Start_GC_GameStartType_name, int32(x))
}
func (x *CG_Game_Start_GC_GameStartType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CG_Game_Start_GC_GameStartType_value, data, "CG_Game_Start_GC_GameStartType")
	if err != nil {
		return err
	}
	*x = CG_Game_Start_GC_GameStartType(value)
	return nil
}

type CG_Game_Start_CG struct {
	ValidateNum      *int32  `protobuf:"varint,1,opt,name=validateNum" json:"validateNum,omitempty"`
	AccID            *uint64 `protobuf:"varint,2,opt" json:"AccID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CG_Game_Start_CG) Reset()         { *m = CG_Game_Start_CG{} }
func (m *CG_Game_Start_CG) String() string { return proto.CompactTextString(m) }
func (*CG_Game_Start_CG) ProtoMessage()    {}

func (m *CG_Game_Start_CG) GetValidateNum() int32 {
	if m != nil && m.ValidateNum != nil {
		return *m.ValidateNum
	}
	return 0
}

func (m *CG_Game_Start_CG) GetAccID() uint64 {
	if m != nil && m.AccID != nil {
		return *m.AccID
	}
	return 0
}

type CG_Game_Start_GC struct {
	GS_Type          *CG_Game_Start_GC_GameStartType `protobuf:"varint,2,opt,enum=message.CG_Game_Start_GC_GameStartType,def=0" json:"GS_Type,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *CG_Game_Start_GC) Reset()         { *m = CG_Game_Start_GC{} }
func (m *CG_Game_Start_GC) String() string { return proto.CompactTextString(m) }
func (*CG_Game_Start_GC) ProtoMessage()    {}

const Default_CG_Game_Start_GC_GS_Type CG_Game_Start_GC_GameStartType = CG_Game_Start_GC_eGS_Success

func (m *CG_Game_Start_GC) GetGS_Type() CG_Game_Start_GC_GameStartType {
	if m != nil && m.GS_Type != nil {
		return *m.GS_Type
	}
	return Default_CG_Game_Start_GC_GS_Type
}

func init() {
	proto.RegisterEnum("message.CG_Game_Start_GC_GameStartType", CG_Game_Start_GC_GameStartType_name, CG_Game_Start_GC_GameStartType_value)
}
