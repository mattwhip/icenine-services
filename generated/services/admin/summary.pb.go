// Code generated by protoc-gen-go. DO NOT EDIT.
// source: summary.proto

package admin

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

type SummaryRequest struct {
	ForceUpdate          bool     `protobuf:"varint,1,opt,name=forceUpdate,proto3" json:"forceUpdate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SummaryRequest) Reset()         { *m = SummaryRequest{} }
func (m *SummaryRequest) String() string { return proto.CompactTextString(m) }
func (*SummaryRequest) ProtoMessage()    {}
func (*SummaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{0}
}

func (m *SummaryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryRequest.Unmarshal(m, b)
}
func (m *SummaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryRequest.Marshal(b, m, deterministic)
}
func (m *SummaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryRequest.Merge(m, src)
}
func (m *SummaryRequest) XXX_Size() int {
	return xxx_messageInfo_SummaryRequest.Size(m)
}
func (m *SummaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryRequest proto.InternalMessageInfo

func (m *SummaryRequest) GetForceUpdate() bool {
	if m != nil {
		return m.ForceUpdate
	}
	return false
}

type SummaryResponse struct {
	LastUpdatedUnixTime  int64    `protobuf:"varint,1,opt,name=lastUpdatedUnixTime,proto3" json:"lastUpdatedUnixTime,omitempty"`
	TotalUsers           int32    `protobuf:"varint,2,opt,name=totalUsers,proto3" json:"totalUsers,omitempty"`
	TotalDevices         int32    `protobuf:"varint,3,opt,name=totalDevices,proto3" json:"totalDevices,omitempty"`
	TotalAccounts        int32    `protobuf:"varint,4,opt,name=totalAccounts,proto3" json:"totalAccounts,omitempty"`
	AliveUsers           int32    `protobuf:"varint,5,opt,name=aliveUsers,proto3" json:"aliveUsers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SummaryResponse) Reset()         { *m = SummaryResponse{} }
func (m *SummaryResponse) String() string { return proto.CompactTextString(m) }
func (*SummaryResponse) ProtoMessage()    {}
func (*SummaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{1}
}

func (m *SummaryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryResponse.Unmarshal(m, b)
}
func (m *SummaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryResponse.Marshal(b, m, deterministic)
}
func (m *SummaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryResponse.Merge(m, src)
}
func (m *SummaryResponse) XXX_Size() int {
	return xxx_messageInfo_SummaryResponse.Size(m)
}
func (m *SummaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryResponse proto.InternalMessageInfo

func (m *SummaryResponse) GetLastUpdatedUnixTime() int64 {
	if m != nil {
		return m.LastUpdatedUnixTime
	}
	return 0
}

func (m *SummaryResponse) GetTotalUsers() int32 {
	if m != nil {
		return m.TotalUsers
	}
	return 0
}

func (m *SummaryResponse) GetTotalDevices() int32 {
	if m != nil {
		return m.TotalDevices
	}
	return 0
}

func (m *SummaryResponse) GetTotalAccounts() int32 {
	if m != nil {
		return m.TotalAccounts
	}
	return 0
}

func (m *SummaryResponse) GetAliveUsers() int32 {
	if m != nil {
		return m.AliveUsers
	}
	return 0
}

func init() {
	proto.RegisterType((*SummaryRequest)(nil), "admin.SummaryRequest")
	proto.RegisterType((*SummaryResponse)(nil), "admin.SummaryResponse")
}

func init() { proto.RegisterFile("summary.proto", fileDescriptor_f7168d0e3f3f5589) }

var fileDescriptor_f7168d0e3f3f5589 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xcd, 0xcd,
	0x4d, 0x2c, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x4c, 0xc9, 0xcd, 0xcc,
	0x53, 0x32, 0xe2, 0xe2, 0x0b, 0x86, 0x88, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x29,
	0x70, 0x71, 0xa7, 0xe5, 0x17, 0x25, 0xa7, 0x86, 0x16, 0xa4, 0x24, 0x96, 0xa4, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x04, 0x21, 0x0b, 0x29, 0x9d, 0x65, 0xe4, 0xe2, 0x87, 0x6b, 0x2a, 0x2e, 0xc8,
	0xcf, 0x2b, 0x4e, 0x15, 0x32, 0xe0, 0x12, 0xce, 0x49, 0x2c, 0x2e, 0x81, 0xa8, 0x48, 0x09, 0xcd,
	0xcb, 0xac, 0x08, 0xc9, 0xcc, 0x85, 0xe8, 0x66, 0x0e, 0xc2, 0x26, 0x25, 0x24, 0xc7, 0xc5, 0x55,
	0x92, 0x5f, 0x92, 0x98, 0x13, 0x5a, 0x9c, 0x5a, 0x54, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1a,
	0x84, 0x24, 0x22, 0xa4, 0xc4, 0xc5, 0x03, 0xe6, 0xb9, 0xa4, 0x96, 0x65, 0x26, 0xa7, 0x16, 0x4b,
	0x30, 0x83, 0x55, 0xa0, 0x88, 0x09, 0xa9, 0x70, 0xf1, 0x82, 0xf9, 0x8e, 0xc9, 0xc9, 0xf9, 0xa5,
	0x79, 0x25, 0xc5, 0x12, 0x2c, 0x60, 0x45, 0xa8, 0x82, 0x20, 0x9b, 0x12, 0x73, 0x32, 0xcb, 0x52,
	0x21, 0x36, 0xb1, 0x42, 0x6c, 0x42, 0x88, 0x24, 0xb1, 0x81, 0x43, 0xc4, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x0e, 0xf7, 0x5d, 0xb4, 0x22, 0x01, 0x00, 0x00,
}