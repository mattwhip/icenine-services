// Code generated by protoc-gen-go. DO NOT EDIT.
// source: table.proto

package teenpatti

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

// TableStatus is the state of the table
type TableStatus struct {
	// state of the game
	State uint32 `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	// array of pots
	Pots []*Pot `protobuf:"bytes,2,rep,name=pots,proto3" json:"pots,omitempty"`
	// amount of the current stake
	CurrentStakeAmount int64 `protobuf:"varint,3,opt,name=currentStakeAmount,proto3" json:"currentStakeAmount,omitempty"`
	// userID of whose turn it is
	CurrentUser string `protobuf:"bytes,4,opt,name=currentUser,proto3" json:"currentUser,omitempty"`
	// users is a map of user info in the table
	Users                map[int32]*User `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TableStatus) Reset()         { *m = TableStatus{} }
func (m *TableStatus) String() string { return proto.CompactTextString(m) }
func (*TableStatus) ProtoMessage()    {}
func (*TableStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{0}
}

func (m *TableStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableStatus.Unmarshal(m, b)
}
func (m *TableStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableStatus.Marshal(b, m, deterministic)
}
func (m *TableStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableStatus.Merge(m, src)
}
func (m *TableStatus) XXX_Size() int {
	return xxx_messageInfo_TableStatus.Size(m)
}
func (m *TableStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TableStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TableStatus proto.InternalMessageInfo

func (m *TableStatus) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *TableStatus) GetPots() []*Pot {
	if m != nil {
		return m.Pots
	}
	return nil
}

func (m *TableStatus) GetCurrentStakeAmount() int64 {
	if m != nil {
		return m.CurrentStakeAmount
	}
	return 0
}

func (m *TableStatus) GetCurrentUser() string {
	if m != nil {
		return m.CurrentUser
	}
	return ""
}

func (m *TableStatus) GetUsers() map[int32]*User {
	if m != nil {
		return m.Users
	}
	return nil
}

// TableConfiguration contains the configuration of the table
type TableConfiguration struct {
	//time given to user to decide action
	MilliSecondsToDecide int64 `protobuf:"varint,1,opt,name=milliSecondsToDecide,proto3" json:"milliSecondsToDecide,omitempty"`
	//time between rounds
	MilliSecondsBetweenRounds int64 `protobuf:"varint,2,opt,name=milliSecondsBetweenRounds,proto3" json:"milliSecondsBetweenRounds,omitempty"`
	//time for user action animation
	MilliSecondsForUserActionAnimation int64 `protobuf:"varint,3,opt,name=milliSecondsForUserActionAnimation,proto3" json:"milliSecondsForUserActionAnimation,omitempty"`
	//time for celebration
	MilliSecondsForCelebration int64 `protobuf:"varint,4,opt,name=milliSecondsForCelebration,proto3" json:"milliSecondsForCelebration,omitempty"`
	//time before first round
	MilliSecondsBeforeFirstRound int64 `protobuf:"varint,5,opt,name=milliSecondsBeforeFirstRound,proto3" json:"milliSecondsBeforeFirstRound,omitempty"`
	//time before the last user in table is forced to exit
	MilliSecondsBeforeLastUserExit int64 `protobuf:"varint,6,opt,name=milliSecondsBeforeLastUserExit,proto3" json:"milliSecondsBeforeLastUserExit,omitempty"`
	//number of seats in table
	NumberOfSeats int32 `protobuf:"varint,7,opt,name=numberOfSeats,proto3" json:"numberOfSeats,omitempty"`
	//number of turn a user does not respond before being kicked out
	IdleTurnLimit int32 `protobuf:"varint,8,opt,name=idleTurnLimit,proto3" json:"idleTurnLimit,omitempty"`
	//starting bet amount
	StartingStakeAmount  int64    `protobuf:"varint,9,opt,name=startingStakeAmount,proto3" json:"startingStakeAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableConfiguration) Reset()         { *m = TableConfiguration{} }
func (m *TableConfiguration) String() string { return proto.CompactTextString(m) }
func (*TableConfiguration) ProtoMessage()    {}
func (*TableConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{1}
}

func (m *TableConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableConfiguration.Unmarshal(m, b)
}
func (m *TableConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableConfiguration.Marshal(b, m, deterministic)
}
func (m *TableConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableConfiguration.Merge(m, src)
}
func (m *TableConfiguration) XXX_Size() int {
	return xxx_messageInfo_TableConfiguration.Size(m)
}
func (m *TableConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_TableConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_TableConfiguration proto.InternalMessageInfo

func (m *TableConfiguration) GetMilliSecondsToDecide() int64 {
	if m != nil {
		return m.MilliSecondsToDecide
	}
	return 0
}

func (m *TableConfiguration) GetMilliSecondsBetweenRounds() int64 {
	if m != nil {
		return m.MilliSecondsBetweenRounds
	}
	return 0
}

func (m *TableConfiguration) GetMilliSecondsForUserActionAnimation() int64 {
	if m != nil {
		return m.MilliSecondsForUserActionAnimation
	}
	return 0
}

func (m *TableConfiguration) GetMilliSecondsForCelebration() int64 {
	if m != nil {
		return m.MilliSecondsForCelebration
	}
	return 0
}

func (m *TableConfiguration) GetMilliSecondsBeforeFirstRound() int64 {
	if m != nil {
		return m.MilliSecondsBeforeFirstRound
	}
	return 0
}

func (m *TableConfiguration) GetMilliSecondsBeforeLastUserExit() int64 {
	if m != nil {
		return m.MilliSecondsBeforeLastUserExit
	}
	return 0
}

func (m *TableConfiguration) GetNumberOfSeats() int32 {
	if m != nil {
		return m.NumberOfSeats
	}
	return 0
}

func (m *TableConfiguration) GetIdleTurnLimit() int32 {
	if m != nil {
		return m.IdleTurnLimit
	}
	return 0
}

func (m *TableConfiguration) GetStartingStakeAmount() int64 {
	if m != nil {
		return m.StartingStakeAmount
	}
	return 0
}

// Pot contains information of bets in a round
type Pot struct {
	// amount in a single Pot
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// userIDs of users that have a stake in the pot
	Users []string `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	// userID of winner of this pot
	Winner               string   `protobuf:"bytes,3,opt,name=winner,proto3" json:"winner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pot) Reset()         { *m = Pot{} }
func (m *Pot) String() string { return proto.CompactTextString(m) }
func (*Pot) ProtoMessage()    {}
func (*Pot) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{2}
}

func (m *Pot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pot.Unmarshal(m, b)
}
func (m *Pot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pot.Marshal(b, m, deterministic)
}
func (m *Pot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pot.Merge(m, src)
}
func (m *Pot) XXX_Size() int {
	return xxx_messageInfo_Pot.Size(m)
}
func (m *Pot) XXX_DiscardUnknown() {
	xxx_messageInfo_Pot.DiscardUnknown(m)
}

var xxx_messageInfo_Pot proto.InternalMessageInfo

func (m *Pot) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Pot) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Pot) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

// User contains information of user in a table
type User struct {
	// user id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// blind is true if user is blind in round else false
	Blind bool `protobuf:"varint,2,opt,name=blind,proto3" json:"blind,omitempty"`
	// inPlay is true when user is in current round else false
	InPlay bool `protobuf:"varint,3,opt,name=inPlay,proto3" json:"inPlay,omitempty"`
	// lastBet is the amount user bet in their last turn
	LastBet int64 `protobuf:"varint,4,opt,name=lastBet,proto3" json:"lastBet,omitempty"`
	// totalBet is the amount user bet in the current round so far
	TotalBet int64 `protobuf:"varint,5,opt,name=totalBet,proto3" json:"totalBet,omitempty"`
	// cards of user
	Cards                *Cards   `protobuf:"bytes,6,opt,name=cards,proto3" json:"cards,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_448a2743262f7a00, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetBlind() bool {
	if m != nil {
		return m.Blind
	}
	return false
}

func (m *User) GetInPlay() bool {
	if m != nil {
		return m.InPlay
	}
	return false
}

func (m *User) GetLastBet() int64 {
	if m != nil {
		return m.LastBet
	}
	return 0
}

func (m *User) GetTotalBet() int64 {
	if m != nil {
		return m.TotalBet
	}
	return 0
}

func (m *User) GetCards() *Cards {
	if m != nil {
		return m.Cards
	}
	return nil
}

func init() {
	proto.RegisterType((*TableStatus)(nil), "teenpatti.TableStatus")
	proto.RegisterMapType((map[int32]*User)(nil), "teenpatti.TableStatus.UsersEntry")
	proto.RegisterType((*TableConfiguration)(nil), "teenpatti.TableConfiguration")
	proto.RegisterType((*Pot)(nil), "teenpatti.Pot")
	proto.RegisterType((*User)(nil), "teenpatti.User")
}

func init() { proto.RegisterFile("table.proto", fileDescriptor_448a2743262f7a00) }

var fileDescriptor_448a2743262f7a00 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x61, 0x8b, 0xd3, 0x40,
	0x10, 0x25, 0x49, 0x73, 0xd7, 0x4c, 0xb8, 0xf3, 0x58, 0x0f, 0x89, 0x45, 0x24, 0x06, 0x95, 0x7e,
	0x2a, 0x52, 0x3f, 0x28, 0x22, 0xc2, 0xb5, 0x5e, 0x41, 0x3c, 0xb4, 0x6c, 0xeb, 0x0f, 0xd8, 0x36,
	0xdb, 0x63, 0xb9, 0x74, 0xb7, 0xec, 0x4e, 0x3c, 0xfb, 0x63, 0x04, 0xf1, 0x97, 0xca, 0xee, 0xc6,
	0x33, 0xd5, 0x7a, 0xfa, 0xad, 0xf3, 0xde, 0xeb, 0xdb, 0xb7, 0x3b, 0x33, 0x81, 0x14, 0xd9, 0xa2,
	0xe2, 0x83, 0x8d, 0x56, 0xa8, 0x48, 0x82, 0x9c, 0xcb, 0x0d, 0x43, 0x14, 0x3d, 0xa8, 0x0d, 0xd7,
	0x1e, 0x2e, 0xbe, 0x86, 0x90, 0xce, 0xad, 0x6c, 0x86, 0x0c, 0x6b, 0x43, 0x4e, 0x21, 0x36, 0xc8,
	0x90, 0x67, 0x41, 0x1e, 0xf4, 0x8f, 0xa8, 0x2f, 0x48, 0x01, 0x9d, 0x8d, 0x42, 0x93, 0x85, 0x79,
	0xd4, 0x4f, 0x87, 0xc7, 0x83, 0x1b, 0xaf, 0xc1, 0x54, 0x21, 0x75, 0x1c, 0x19, 0x00, 0x59, 0xd6,
	0x5a, 0x73, 0x89, 0x33, 0x64, 0x57, 0xfc, 0x6c, 0xad, 0x6a, 0x89, 0x59, 0x94, 0x07, 0xfd, 0x88,
	0xee, 0x61, 0x48, 0x0e, 0x69, 0x83, 0x7e, 0x32, 0x5c, 0x67, 0x9d, 0x3c, 0xe8, 0x27, 0xb4, 0x0d,
	0x91, 0x17, 0x10, 0xdb, 0xa4, 0x26, 0x8b, 0xdd, 0xb1, 0x8f, 0x5a, 0xc7, 0xb6, 0x22, 0x0f, 0xac,
	0xd6, 0x9c, 0x4b, 0xd4, 0x5b, 0xea, 0xf5, 0xbd, 0x77, 0x00, 0xbf, 0x40, 0x72, 0x02, 0xd1, 0x15,
	0xdf, 0xba, 0x0b, 0xc5, 0xd4, 0xfe, 0x24, 0x4f, 0x20, 0xfe, 0xcc, 0xaa, 0x9a, 0x67, 0x61, 0x1e,
	0xf4, 0xd3, 0xe1, 0x9d, 0x96, 0xb1, 0xfd, 0x1f, 0xf5, 0xec, 0xab, 0xf0, 0x65, 0x50, 0x7c, 0xeb,
	0x00, 0x71, 0x87, 0x8d, 0x95, 0x5c, 0x89, 0xcb, 0x5a, 0x33, 0x14, 0x4a, 0x92, 0x21, 0x9c, 0xae,
	0x45, 0x55, 0x89, 0x19, 0x5f, 0x2a, 0x59, 0x9a, 0xb9, 0x7a, 0xcb, 0x97, 0xa2, 0xf4, 0xaf, 0x16,
	0xd1, 0xbd, 0x1c, 0x79, 0x0d, 0xf7, 0xdb, 0xf8, 0x88, 0xe3, 0x35, 0xe7, 0x92, 0xaa, 0x5a, 0x96,
	0xc6, 0x25, 0x89, 0xe8, 0xdf, 0x05, 0xe4, 0x03, 0x14, 0x6d, 0x72, 0xa2, 0xb4, 0x8d, 0x7a, 0xb6,
	0xb4, 0x71, 0xce, 0xa4, 0x58, 0xbb, 0x5c, 0xcd, 0x73, 0xff, 0x87, 0x92, 0xbc, 0x81, 0xde, 0x6f,
	0xaa, 0x31, 0xaf, 0xf8, 0xc2, 0xdf, 0xcf, 0x75, 0x23, 0xa2, 0xb7, 0x28, 0xc8, 0x08, 0x1e, 0xec,
	0x86, 0x5d, 0x29, 0xcd, 0x27, 0x42, 0x1b, 0x74, 0x81, 0xb3, 0xd8, 0x39, 0xdc, 0xaa, 0x21, 0x13,
	0x78, 0xf8, 0x27, 0x7f, 0xc1, 0x8c, 0x6b, 0xff, 0xf9, 0x17, 0x81, 0xd9, 0x81, 0x73, 0xf9, 0x87,
	0x8a, 0x3c, 0x86, 0x23, 0x59, 0xaf, 0x17, 0x5c, 0x7f, 0x5c, 0xcd, 0x38, 0x43, 0x93, 0x1d, 0xba,
	0x5e, 0xef, 0x82, 0x56, 0x25, 0xca, 0x8a, 0xcf, 0x6b, 0x2d, 0x2f, 0xc4, 0x5a, 0x60, 0xd6, 0xf5,
	0xaa, 0x1d, 0x90, 0x3c, 0x83, 0xbb, 0x06, 0x99, 0x46, 0x21, 0x2f, 0xdb, 0x73, 0x9c, 0xb8, 0x20,
	0xfb, 0xa8, 0xe2, 0x3d, 0x44, 0x53, 0x85, 0xe4, 0x1e, 0x1c, 0x30, 0xaf, 0xf5, 0x43, 0xd0, 0x54,
	0x76, 0xa3, 0xfc, 0x14, 0xdb, 0xe5, 0x49, 0x9a, 0x11, 0xb5, 0xea, 0x6b, 0x21, 0x25, 0xd7, 0xae,
	0x65, 0x09, 0x6d, 0xaa, 0xe2, 0x7b, 0x00, 0x1d, 0x37, 0xfc, 0xc7, 0x10, 0x8a, 0xd2, 0x59, 0x25,
	0x34, 0x14, 0xa5, 0xb5, 0x59, 0x54, 0x42, 0x96, 0x6e, 0x52, 0xba, 0xd4, 0x17, 0xd6, 0x46, 0xc8,
	0x69, 0xc5, 0xb6, 0xce, 0xa6, 0x4b, 0x9b, 0x8a, 0x64, 0x70, 0x58, 0x31, 0x83, 0x23, 0x8e, 0x4d,
	0x2b, 0x7f, 0x96, 0xa4, 0x07, 0x5d, 0x54, 0xc8, 0x2a, 0x4b, 0xf9, 0x1e, 0xdd, 0xd4, 0xe4, 0x29,
	0xc4, 0x4b, 0xa6, 0x4b, 0xe3, 0x9e, 0x3d, 0x1d, 0x9e, 0xb4, 0xf6, 0x62, 0x6c, 0x71, 0xea, 0xe9,
	0xc5, 0x81, 0xfb, 0x76, 0x3c, 0xff, 0x11, 0x00, 0x00, 0xff, 0xff, 0x97, 0x79, 0x17, 0x3f, 0x61,
	0x04, 0x00, 0x00,
}
