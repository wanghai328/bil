// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb.proto

/*
Package render is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	PB
*/
package render

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PB struct {
	Code    int64                `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	TTL     uint64               `protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	Data    *google_protobuf.Any `protobuf:"bytes,4,opt,name=Data" json:"Data,omitempty"`
}

func (m *PB) Reset()                    { *m = PB{} }
func (m *PB) String() string            { return proto.CompactTextString(m) }
func (*PB) ProtoMessage()               {}
func (*PB) Descriptor() ([]byte, []int) { return fileDescriptorPb, []int{0} }

func (m *PB) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PB) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PB) GetTTL() uint64 {
	if m != nil {
		return m.TTL
	}
	return 0
}

func (m *PB) GetData() *google_protobuf.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*PB)(nil), "render.PB")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptorPb) }

var fileDescriptorPb = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x48, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x4a, 0xcd, 0x4b, 0x49, 0x2d, 0x92, 0x92, 0x4c, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x26, 0x95, 0xa6, 0xe9, 0x27, 0xe6, 0x55, 0x42, 0x94,
	0x28, 0xe5, 0x71, 0x31, 0x05, 0x38, 0x09, 0x09, 0x71, 0xb1, 0x38, 0xe7, 0xa7, 0xa4, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x30, 0x07, 0x81, 0xd9, 0x42, 0x12, 0x5c, 0xec, 0xbe, 0xa9, 0xc5, 0xc5, 0x89,
	0xe9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x00, 0x17, 0x73, 0x48,
	0x88, 0x8f, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x88, 0x29, 0xa4, 0xc1, 0xc5, 0xe2, 0x92,
	0x58, 0x92, 0x28, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa2, 0x07, 0xb1, 0x4f, 0x0f, 0x66,
	0x9f, 0x9e, 0x63, 0x5e, 0x65, 0x10, 0x58, 0x45, 0x12, 0x1b, 0x58, 0xcc, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x7a, 0x92, 0x16, 0x71, 0xa5, 0x00, 0x00, 0x00,
}
