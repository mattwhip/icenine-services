// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login_facebook.proto

package useraccounts

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

type LoginFacebookRequest struct {
	// The Facebook access token
	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	// The unique ID of the device being used to access the system
	DeviceID string `protobuf:"bytes,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	// The model information of the device being used to access the system
	DeviceModelInfo string `protobuf:"bytes,3,opt,name=deviceModelInfo,proto3" json:"deviceModelInfo,omitempty"`
	// The Facebook graph API user ID
	FacebookID string `protobuf:"bytes,4,opt,name=facebookID,proto3" json:"facebookID,omitempty"`
	// The operating system of the device being used to access the system
	OperatingSystem      string   `protobuf:"bytes,5,opt,name=operatingSystem,proto3" json:"operatingSystem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginFacebookRequest) Reset()         { *m = LoginFacebookRequest{} }
func (m *LoginFacebookRequest) String() string { return proto.CompactTextString(m) }
func (*LoginFacebookRequest) ProtoMessage()    {}
func (*LoginFacebookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8917dc018e20751, []int{0}
}

func (m *LoginFacebookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginFacebookRequest.Unmarshal(m, b)
}
func (m *LoginFacebookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginFacebookRequest.Marshal(b, m, deterministic)
}
func (m *LoginFacebookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginFacebookRequest.Merge(m, src)
}
func (m *LoginFacebookRequest) XXX_Size() int {
	return xxx_messageInfo_LoginFacebookRequest.Size(m)
}
func (m *LoginFacebookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginFacebookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginFacebookRequest proto.InternalMessageInfo

func (m *LoginFacebookRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LoginFacebookRequest) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *LoginFacebookRequest) GetDeviceModelInfo() string {
	if m != nil {
		return m.DeviceModelInfo
	}
	return ""
}

func (m *LoginFacebookRequest) GetFacebookID() string {
	if m != nil {
		return m.FacebookID
	}
	return ""
}

func (m *LoginFacebookRequest) GetOperatingSystem() string {
	if m != nil {
		return m.OperatingSystem
	}
	return ""
}

type LoginFacebookResponse struct {
	// The json web token used to authorize the user in other
	// parts of the system
	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	// The coins in the user's wallet
	Coins int64 `protobuf:"varint,2,opt,name=coins,proto3" json:"coins,omitempty"`
	// The skill level of the user
	SkillLevel           float32  `protobuf:"fixed32,3,opt,name=skillLevel,proto3" json:"skillLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginFacebookResponse) Reset()         { *m = LoginFacebookResponse{} }
func (m *LoginFacebookResponse) String() string { return proto.CompactTextString(m) }
func (*LoginFacebookResponse) ProtoMessage()    {}
func (*LoginFacebookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8917dc018e20751, []int{1}
}

func (m *LoginFacebookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginFacebookResponse.Unmarshal(m, b)
}
func (m *LoginFacebookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginFacebookResponse.Marshal(b, m, deterministic)
}
func (m *LoginFacebookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginFacebookResponse.Merge(m, src)
}
func (m *LoginFacebookResponse) XXX_Size() int {
	return xxx_messageInfo_LoginFacebookResponse.Size(m)
}
func (m *LoginFacebookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginFacebookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginFacebookResponse proto.InternalMessageInfo

func (m *LoginFacebookResponse) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func (m *LoginFacebookResponse) GetCoins() int64 {
	if m != nil {
		return m.Coins
	}
	return 0
}

func (m *LoginFacebookResponse) GetSkillLevel() float32 {
	if m != nil {
		return m.SkillLevel
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginFacebookRequest)(nil), "useraccounts.LoginFacebookRequest")
	proto.RegisterType((*LoginFacebookResponse)(nil), "useraccounts.LoginFacebookResponse")
}

func init() { proto.RegisterFile("login_facebook.proto", fileDescriptor_d8917dc018e20751) }

var fileDescriptor_d8917dc018e20751 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0xd6, 0x15, 0x1d, 0x05, 0x25, 0x54, 0x08, 0x1e, 0x64, 0xd9, 0xd3, 0x9e, 0xbc,
	0xf8, 0x0a, 0x8b, 0x50, 0x58, 0x2f, 0xd1, 0xfb, 0x92, 0xcd, 0xce, 0x96, 0xd8, 0x98, 0xa9, 0x9d,
	0xb4, 0xe2, 0x1b, 0xfa, 0x58, 0x92, 0xb4, 0x4a, 0xe9, 0x6d, 0xe6, 0xcb, 0xcf, 0xf0, 0xfd, 0x81,
	0xc2, 0x51, 0x65, 0xfd, 0xfe, 0xa4, 0x0d, 0x1e, 0x88, 0xea, 0xc7, 0xa6, 0xa5, 0x40, 0xe2, 0xba,
	0x63, 0x6c, 0xb5, 0x31, 0xd4, 0xf9, 0xc0, 0xeb, 0x9f, 0x0c, 0x8a, 0x5d, 0x8c, 0x3d, 0x8f, 0x29,
	0x85, 0x9f, 0x1d, 0x72, 0x10, 0x2b, 0xb8, 0xd2, 0xc6, 0x20, 0xf3, 0x1b, 0xd5, 0xe8, 0x65, 0xb6,
	0xca, 0x36, 0x97, 0x6a, 0x8a, 0xc4, 0x3d, 0x5c, 0x1c, 0xb1, 0xb7, 0x06, 0xcb, 0xad, 0x5c, 0xa4,
	0xe7, 0xff, 0x5d, 0x6c, 0xe0, 0x66, 0x98, 0x5f, 0xe8, 0x88, 0xae, 0xf4, 0x27, 0x92, 0x79, 0x8a,
	0xcc, 0xb1, 0x78, 0x00, 0xf8, 0x13, 0x2c, 0xb7, 0xf2, 0x2c, 0x85, 0x26, 0x24, 0x5e, 0xa2, 0x06,
	0x5b, 0x1d, 0xac, 0xaf, 0x5e, 0xbf, 0x39, 0xe0, 0x87, 0x5c, 0x0e, 0x97, 0x66, 0x78, 0xbd, 0x87,
	0xbb, 0x59, 0x13, 0x6e, 0xc8, 0x33, 0x8a, 0x5b, 0xc8, 0xdf, 0xbf, 0xc2, 0x58, 0x21, 0x8e, 0xa2,
	0x80, 0xa5, 0x21, 0xeb, 0x39, 0x79, 0xe7, 0x6a, 0x58, 0xa2, 0x0a, 0xd7, 0xd6, 0xb9, 0x1d, 0xf6,
	0xe8, 0x92, 0xef, 0x42, 0x4d, 0xc8, 0xe1, 0x3c, 0x7d, 0xe0, 0xd3, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x13, 0x10, 0xf7, 0x9e, 0x58, 0x01, 0x00, 0x00,
}
