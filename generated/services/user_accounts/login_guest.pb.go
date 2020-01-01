// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login_guest.proto

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

type LoginGuestRequest struct {
	// The unique ID of the device being used to access the system
	DeviceID string `protobuf:"bytes,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	// The model information of the device being used to access the system
	DeviceModelInfo string `protobuf:"bytes,2,opt,name=deviceModelInfo,proto3" json:"deviceModelInfo,omitempty"`
	// The operating system of the device being used to access the system
	OperatingSystem      string   `protobuf:"bytes,3,opt,name=operatingSystem,proto3" json:"operatingSystem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginGuestRequest) Reset()         { *m = LoginGuestRequest{} }
func (m *LoginGuestRequest) String() string { return proto.CompactTextString(m) }
func (*LoginGuestRequest) ProtoMessage()    {}
func (*LoginGuestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1da7b0e7f80f2ea0, []int{0}
}

func (m *LoginGuestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginGuestRequest.Unmarshal(m, b)
}
func (m *LoginGuestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginGuestRequest.Marshal(b, m, deterministic)
}
func (m *LoginGuestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginGuestRequest.Merge(m, src)
}
func (m *LoginGuestRequest) XXX_Size() int {
	return xxx_messageInfo_LoginGuestRequest.Size(m)
}
func (m *LoginGuestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginGuestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginGuestRequest proto.InternalMessageInfo

func (m *LoginGuestRequest) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *LoginGuestRequest) GetDeviceModelInfo() string {
	if m != nil {
		return m.DeviceModelInfo
	}
	return ""
}

func (m *LoginGuestRequest) GetOperatingSystem() string {
	if m != nil {
		return m.OperatingSystem
	}
	return ""
}

type LoginGuestResponse struct {
	// The json web token used to authorize the user in HTTP APIs
	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	// The ID used to identify this user when Kafka messages are sent from
	// the player to the server
	UserID []byte `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	// The server-to-user kafka topic used to send messages from the server
	// to the player
	ServerToUserTopic string `protobuf:"bytes,3,opt,name=serverToUserTopic,proto3" json:"serverToUserTopic,omitempty"`
	// The coins in the guest user's wallet
	Coins int64 `protobuf:"varint,4,opt,name=coins,proto3" json:"coins,omitempty"`
	// The skill level of the guest user
	SkillLevel float32 `protobuf:"fixed32,5,opt,name=skillLevel,proto3" json:"skillLevel,omitempty"`
	// The daily bonus availability
	IsDailyBonusAvailable bool `protobuf:"varint,6,opt,name=isDailyBonusAvailable,proto3" json:"isDailyBonusAvailable,omitempty"`
	// The number of seconds until the next daily bonus is available
	SecondsUntilDailyBonusAvailable int32 `protobuf:"varint,7,opt,name=secondsUntilDailyBonusAvailable,proto3" json:"secondsUntilDailyBonusAvailable,omitempty"`
	// The current daily bonus streak
	DailyBonusStreak int32 `protobuf:"varint,8,opt,name=dailyBonusStreak,proto3" json:"dailyBonusStreak,omitempty"`
	// The wheel values of the next daily bonus to play
	DailyBonusWheelValues []int64 `protobuf:"varint,9,rep,packed,name=dailyBonusWheelValues,proto3" json:"dailyBonusWheelValues,omitempty"`
	// The topic used to send game messages from user to server if game is
	// in progress
	GameInProgressTopic string `protobuf:"bytes,10,opt,name=gameInProgressTopic,proto3" json:"gameInProgressTopic,omitempty"`
	// The server-specified offset intended for the client to use on WebSocket
	// topic subscription
	ServerToUserTopicOffset int64    `protobuf:"varint,11,opt,name=serverToUserTopicOffset,proto3" json:"serverToUserTopicOffset,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *LoginGuestResponse) Reset()         { *m = LoginGuestResponse{} }
func (m *LoginGuestResponse) String() string { return proto.CompactTextString(m) }
func (*LoginGuestResponse) ProtoMessage()    {}
func (*LoginGuestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1da7b0e7f80f2ea0, []int{1}
}

func (m *LoginGuestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginGuestResponse.Unmarshal(m, b)
}
func (m *LoginGuestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginGuestResponse.Marshal(b, m, deterministic)
}
func (m *LoginGuestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginGuestResponse.Merge(m, src)
}
func (m *LoginGuestResponse) XXX_Size() int {
	return xxx_messageInfo_LoginGuestResponse.Size(m)
}
func (m *LoginGuestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginGuestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginGuestResponse proto.InternalMessageInfo

func (m *LoginGuestResponse) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func (m *LoginGuestResponse) GetUserID() []byte {
	if m != nil {
		return m.UserID
	}
	return nil
}

func (m *LoginGuestResponse) GetServerToUserTopic() string {
	if m != nil {
		return m.ServerToUserTopic
	}
	return ""
}

func (m *LoginGuestResponse) GetCoins() int64 {
	if m != nil {
		return m.Coins
	}
	return 0
}

func (m *LoginGuestResponse) GetSkillLevel() float32 {
	if m != nil {
		return m.SkillLevel
	}
	return 0
}

func (m *LoginGuestResponse) GetIsDailyBonusAvailable() bool {
	if m != nil {
		return m.IsDailyBonusAvailable
	}
	return false
}

func (m *LoginGuestResponse) GetSecondsUntilDailyBonusAvailable() int32 {
	if m != nil {
		return m.SecondsUntilDailyBonusAvailable
	}
	return 0
}

func (m *LoginGuestResponse) GetDailyBonusStreak() int32 {
	if m != nil {
		return m.DailyBonusStreak
	}
	return 0
}

func (m *LoginGuestResponse) GetDailyBonusWheelValues() []int64 {
	if m != nil {
		return m.DailyBonusWheelValues
	}
	return nil
}

func (m *LoginGuestResponse) GetGameInProgressTopic() string {
	if m != nil {
		return m.GameInProgressTopic
	}
	return ""
}

func (m *LoginGuestResponse) GetServerToUserTopicOffset() int64 {
	if m != nil {
		return m.ServerToUserTopicOffset
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginGuestRequest)(nil), "useraccounts.LoginGuestRequest")
	proto.RegisterType((*LoginGuestResponse)(nil), "useraccounts.LoginGuestResponse")
}

func init() { proto.RegisterFile("login_guest.proto", fileDescriptor_1da7b0e7f80f2ea0) }

var fileDescriptor_1da7b0e7f80f2ea0 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0xcb, 0xd3, 0x40,
	0x10, 0x26, 0x6f, 0x6c, 0xed, 0x3b, 0x16, 0x6c, 0xd7, 0xaf, 0xc5, 0x83, 0x86, 0x9e, 0x82, 0x88,
	0x08, 0x7a, 0xf0, 0xaa, 0x14, 0x34, 0x50, 0x51, 0xb6, 0x56, 0x8f, 0xb2, 0x4d, 0xa6, 0x71, 0xed,
	0x76, 0x37, 0xee, 0x6c, 0x22, 0x3d, 0xfb, 0x87, 0xfd, 0x09, 0xb2, 0x49, 0xa9, 0xa5, 0x8d, 0x78,
	0x49, 0xe6, 0xf9, 0x18, 0xe6, 0xe1, 0x61, 0x61, 0xaa, 0x6d, 0xa9, 0xcc, 0xd7, 0xb2, 0x46, 0xf2,
	0xcf, 0x2a, 0x67, 0xbd, 0x65, 0xe3, 0x9a, 0xd0, 0xc9, 0x3c, 0xb7, 0xb5, 0xf1, 0x34, 0xfb, 0x15,
	0xc1, 0x74, 0x11, 0x3c, 0x6f, 0x83, 0x45, 0xe0, 0x8f, 0xf0, 0x63, 0x0f, 0x61, 0x54, 0x60, 0xa3,
	0x72, 0xcc, 0xe6, 0x3c, 0x4a, 0xa2, 0xf4, 0x5a, 0x1c, 0x31, 0x4b, 0xe1, 0x76, 0x37, 0xbf, 0xb7,
	0x05, 0xea, 0xcc, 0x6c, 0x2c, 0xbf, 0x6a, 0x2d, 0xe7, 0x74, 0x70, 0xda, 0x0a, 0x9d, 0xf4, 0xca,
	0x94, 0xcb, 0x3d, 0x79, 0xdc, 0xf1, 0xb8, 0x73, 0x9e, 0xd1, 0xb3, 0xdf, 0x31, 0xb0, 0xd3, 0x14,
	0x54, 0x59, 0x43, 0xc8, 0x26, 0x10, 0x7f, 0xff, 0xe9, 0x0f, 0x09, 0xc2, 0xc8, 0xee, 0xc3, 0x30,
	0xc4, 0xcf, 0xe6, 0xed, 0xcd, 0xb1, 0x38, 0x20, 0xf6, 0x14, 0xa6, 0x84, 0xae, 0x41, 0xf7, 0xc9,
	0xae, 0x28, 0x7c, 0x2b, 0x95, 0x1f, 0x8e, 0x5d, 0x0a, 0xec, 0x2e, 0x0c, 0x72, 0xab, 0x0c, 0xf1,
	0x1b, 0x49, 0x94, 0xc6, 0xa2, 0x03, 0xec, 0x11, 0x00, 0x6d, 0x95, 0xd6, 0x0b, 0x6c, 0x50, 0xf3,
	0x41, 0x12, 0xa5, 0x57, 0xe2, 0x84, 0x61, 0x2f, 0xe1, 0x9e, 0xa2, 0xb9, 0x54, 0x7a, 0xff, 0xc6,
	0x9a, 0x9a, 0x5e, 0x37, 0x52, 0x69, 0xb9, 0xd6, 0xc8, 0x87, 0x49, 0x94, 0x8e, 0x44, 0xbf, 0xc8,
	0xde, 0xc1, 0x63, 0xc2, 0xdc, 0x9a, 0x82, 0x56, 0xc6, 0x2b, 0xdd, 0xb7, 0x7f, 0x33, 0x89, 0xd2,
	0x81, 0xf8, 0x9f, 0x8d, 0x3d, 0x81, 0x49, 0x71, 0xa4, 0x97, 0xde, 0xa1, 0xdc, 0xf2, 0x51, 0xbb,
	0x7a, 0xc1, 0x87, 0xac, 0x7f, 0xb9, 0x2f, 0xdf, 0x10, 0xf5, 0x67, 0xa9, 0x6b, 0x24, 0x7e, 0x9d,
	0xc4, 0x69, 0x2c, 0xfa, 0x45, 0xf6, 0x1c, 0xee, 0x94, 0x72, 0x87, 0x99, 0xf9, 0xe8, 0x6c, 0xe9,
	0x90, 0xa8, 0xeb, 0x11, 0xda, 0x1e, 0xfb, 0x24, 0xf6, 0x0a, 0x1e, 0x5c, 0xd4, 0xfb, 0x61, 0xb3,
	0x21, 0xf4, 0xfc, 0x56, 0xdb, 0xed, 0xbf, 0xe4, 0xf5, 0xb0, 0x7d, 0x8d, 0x2f, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x39, 0xa1, 0xb9, 0x02, 0xa2, 0x02, 0x00, 0x00,
}
