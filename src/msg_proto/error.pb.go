// Code generated by protoc-gen-go.
// source: error.proto
// DO NOT EDIT!

/*
Package msg_proto is a generated protocol buffer package.

It is generated from these files:
	error.proto
	frontend.proto
	gate.proto
	msgcmd.proto
	result.proto

It has these top-level messages:
	Token
	RequestLogin
	CmdResult
*/
package msg_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Error int32

const (
	Error_Init_None         Error = -1
	Error_Sucess            Error = 0
	Error_UnKnow_User       Error = 1
	Error_Mismatch_Password Error = 2
)

var Error_name = map[int32]string{
	-1: "Init_None",
	0:  "Sucess",
	1:  "UnKnow_User",
	2:  "Mismatch_Password",
}
var Error_value = map[string]int32{
	"Init_None":         -1,
	"Sucess":            0,
	"UnKnow_User":       1,
	"Mismatch_Password": 2,
}

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}
func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}
func (x *Error) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Error_value, data, "Error")
	if err != nil {
		return err
	}
	*x = Error(value)
	return nil
}
func (Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("msg_proto.Error", Error_name, Error_value)
}

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x2a, 0xca,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0x2d, 0x4e, 0x8f, 0x07, 0x33, 0xb5,
	0x82, 0xb9, 0x58, 0x5d, 0x41, 0x32, 0x42, 0x62, 0x5c, 0x9c, 0x9e, 0x79, 0x99, 0x25, 0xf1, 0x7e,
	0xf9, 0x79, 0xa9, 0x02, 0xff, 0x61, 0x80, 0x51, 0x88, 0x8b, 0x8b, 0x2d, 0xb8, 0x34, 0x39, 0xb5,
	0xb8, 0x58, 0x80, 0x41, 0x88, 0x9f, 0x8b, 0x3b, 0x34, 0xcf, 0x3b, 0x2f, 0xbf, 0x3c, 0x3e, 0xb4,
	0x38, 0xb5, 0x48, 0x80, 0x51, 0x48, 0x94, 0x4b, 0xd0, 0x37, 0xb3, 0x38, 0x37, 0xb1, 0x24, 0x39,
	0x23, 0x3e, 0x20, 0xb1, 0xb8, 0xb8, 0x3c, 0xbf, 0x28, 0x45, 0x80, 0x09, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x46, 0x65, 0x90, 0x6e, 0x6d, 0x00, 0x00, 0x00,
}
