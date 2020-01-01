// Code generated by protoc-gen-go. DO NOT EDIT.
// source: failures.proto

package util

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Failure is a message intended for notification of a system failure
type Failure struct {
	ErrorCode            uint32   `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}
func (*Failure) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc815f87e23048c4, []int{0}
}

func (m *Failure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Failure.Unmarshal(m, b)
}
func (m *Failure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Failure.Marshal(b, m, deterministic)
}
func (m *Failure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failure.Merge(m, src)
}
func (m *Failure) XXX_Size() int {
	return xxx_messageInfo_Failure.Size(m)
}
func (m *Failure) XXX_DiscardUnknown() {
	xxx_messageInfo_Failure.DiscardUnknown(m)
}

var xxx_messageInfo_Failure proto.InternalMessageInfo

func (m *Failure) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *Failure) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Failure)(nil), "util.Failure")
}

func init() { proto.RegisterFile("failures.proto", fileDescriptor_cc815f87e23048c4) }

var fileDescriptor_cc815f87e23048c4 = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4b, 0xcc, 0xcc,
	0x29, 0x2d, 0x4a, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2d, 0xc9, 0xcc,
	0x51, 0x72, 0xe4, 0x62, 0x77, 0x83, 0x88, 0x0b, 0xc9, 0x70, 0x71, 0xa6, 0x16, 0x15, 0xe5, 0x17,
	0x39, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0x21, 0x04, 0x84, 0x24, 0xb8,
	0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x60,
	0xdc, 0x24, 0x36, 0xb0, 0x79, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x58, 0x83, 0x1d,
	0x61, 0x00, 0x00, 0x00,
}