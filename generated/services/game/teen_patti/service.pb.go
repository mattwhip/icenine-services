// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package teenpatti

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// A request for table generation
type TableRequest struct {
	//id of the kafka table
	TableID string `protobuf:"bytes,1,opt,name=tableID,proto3" json:"tableID,omitempty"`
	//time given to user to decide action
	MilliSecondsToDecide int64 `protobuf:"varint,2,opt,name=milliSecondsToDecide,proto3" json:"milliSecondsToDecide,omitempty"`
	//time between rounds
	MilliSecondsBetweenRounds int64 `protobuf:"varint,3,opt,name=milliSecondsBetweenRounds,proto3" json:"milliSecondsBetweenRounds,omitempty"`
	//time for user action animation
	MilliSecondsForUserActionAnimation int64 `protobuf:"varint,4,opt,name=milliSecondsForUserActionAnimation,proto3" json:"milliSecondsForUserActionAnimation,omitempty"`
	//time for celebration
	MilliSecondsForCelebration int64 `protobuf:"varint,5,opt,name=milliSecondsForCelebration,proto3" json:"milliSecondsForCelebration,omitempty"`
	//time before starting first round
	MilliSecondsBeforeFirstRound int64 `protobuf:"varint,6,opt,name=milliSecondsBeforeFirstRound,proto3" json:"milliSecondsBeforeFirstRound,omitempty"`
	//time before the last user in table is forced to exit
	MilliSecondsBeforeLastUserExit int64 `protobuf:"varint,7,opt,name=milliSecondsBeforeLastUserExit,proto3" json:"milliSecondsBeforeLastUserExit,omitempty"`
	//number of turn a user does not respond before being kicked out
	IdleTurnLimit int32 `protobuf:"varint,8,opt,name=idleTurnLimit,proto3" json:"idleTurnLimit,omitempty"`
	//number of seats in the table
	NumberOfSeats int32 `protobuf:"varint,9,opt,name=numberOfSeats,proto3" json:"numberOfSeats,omitempty"`
	//starting stake amount in table
	StartingStakeAmount int64 `protobuf:"varint,10,opt,name=startingStakeAmount,proto3" json:"startingStakeAmount,omitempty"`
	//list of users joining the table
	Users                []string `protobuf:"bytes,11,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableRequest) Reset()         { *m = TableRequest{} }
func (m *TableRequest) String() string { return proto.CompactTextString(m) }
func (*TableRequest) ProtoMessage()    {}
func (*TableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *TableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableRequest.Unmarshal(m, b)
}
func (m *TableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableRequest.Marshal(b, m, deterministic)
}
func (m *TableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableRequest.Merge(m, src)
}
func (m *TableRequest) XXX_Size() int {
	return xxx_messageInfo_TableRequest.Size(m)
}
func (m *TableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TableRequest proto.InternalMessageInfo

func (m *TableRequest) GetTableID() string {
	if m != nil {
		return m.TableID
	}
	return ""
}

func (m *TableRequest) GetMilliSecondsToDecide() int64 {
	if m != nil {
		return m.MilliSecondsToDecide
	}
	return 0
}

func (m *TableRequest) GetMilliSecondsBetweenRounds() int64 {
	if m != nil {
		return m.MilliSecondsBetweenRounds
	}
	return 0
}

func (m *TableRequest) GetMilliSecondsForUserActionAnimation() int64 {
	if m != nil {
		return m.MilliSecondsForUserActionAnimation
	}
	return 0
}

func (m *TableRequest) GetMilliSecondsForCelebration() int64 {
	if m != nil {
		return m.MilliSecondsForCelebration
	}
	return 0
}

func (m *TableRequest) GetMilliSecondsBeforeFirstRound() int64 {
	if m != nil {
		return m.MilliSecondsBeforeFirstRound
	}
	return 0
}

func (m *TableRequest) GetMilliSecondsBeforeLastUserExit() int64 {
	if m != nil {
		return m.MilliSecondsBeforeLastUserExit
	}
	return 0
}

func (m *TableRequest) GetIdleTurnLimit() int32 {
	if m != nil {
		return m.IdleTurnLimit
	}
	return 0
}

func (m *TableRequest) GetNumberOfSeats() int32 {
	if m != nil {
		return m.NumberOfSeats
	}
	return 0
}

func (m *TableRequest) GetStartingStakeAmount() int64 {
	if m != nil {
		return m.StartingStakeAmount
	}
	return 0
}

func (m *TableRequest) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

// A response with table generation
type TableResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableResponse) Reset()         { *m = TableResponse{} }
func (m *TableResponse) String() string { return proto.CompactTextString(m) }
func (*TableResponse) ProtoMessage()    {}
func (*TableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *TableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableResponse.Unmarshal(m, b)
}
func (m *TableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableResponse.Marshal(b, m, deterministic)
}
func (m *TableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableResponse.Merge(m, src)
}
func (m *TableResponse) XXX_Size() int {
	return xxx_messageInfo_TableResponse.Size(m)
}
func (m *TableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TableResponse proto.InternalMessageInfo

// JoinTable adds a user to a table
type JoinTable struct {
	// an ID of user
	UserID               []string `protobuf:"bytes,1,rep,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinTable) Reset()         { *m = JoinTable{} }
func (m *JoinTable) String() string { return proto.CompactTextString(m) }
func (*JoinTable) ProtoMessage()    {}
func (*JoinTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *JoinTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinTable.Unmarshal(m, b)
}
func (m *JoinTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinTable.Marshal(b, m, deterministic)
}
func (m *JoinTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinTable.Merge(m, src)
}
func (m *JoinTable) XXX_Size() int {
	return xxx_messageInfo_JoinTable.Size(m)
}
func (m *JoinTable) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinTable.DiscardUnknown(m)
}

var xxx_messageInfo_JoinTable proto.InternalMessageInfo

func (m *JoinTable) GetUserID() []string {
	if m != nil {
		return m.UserID
	}
	return nil
}

func init() {
	proto.RegisterType((*TableRequest)(nil), "teenpatti.TableRequest")
	proto.RegisterType((*TableResponse)(nil), "teenpatti.TableResponse")
	proto.RegisterType((*JoinTable)(nil), "teenpatti.JoinTable")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdd, 0x8e, 0xd3, 0x30,
	0x10, 0x85, 0x09, 0xfd, 0x23, 0x53, 0x2a, 0x24, 0x53, 0x81, 0xa9, 0x10, 0x8a, 0x02, 0x17, 0xb9,
	0xaa, 0x50, 0xb9, 0x45, 0x48, 0xfd, 0xa1, 0x12, 0xa8, 0xa2, 0x28, 0x0d, 0x0f, 0xe0, 0x24, 0x53,
	0x64, 0x91, 0xd8, 0xc5, 0x9e, 0x00, 0xcf, 0xc8, 0x53, 0xad, 0xe2, 0x74, 0x57, 0xed, 0xb6, 0xdb,
	0xdd, 0x3b, 0xcf, 0x9c, 0x6f, 0x8e, 0xcf, 0xc5, 0x81, 0x81, 0x45, 0xf3, 0x47, 0x66, 0x38, 0xde,
	0x19, 0x4d, 0x9a, 0xf9, 0x84, 0xa8, 0x76, 0x82, 0x48, 0x86, 0xff, 0xdb, 0xf0, 0x34, 0x11, 0x69,
	0x81, 0x31, 0xfe, 0xae, 0xd0, 0x12, 0xe3, 0xd0, 0xa3, 0x7a, 0xfe, 0xb2, 0xe0, 0x5e, 0xe0, 0x45,
	0x7e, 0x7c, 0x3d, 0xb2, 0x09, 0x0c, 0x4b, 0x59, 0x14, 0x72, 0x83, 0x99, 0x56, 0xb9, 0x4d, 0xf4,
	0x02, 0x33, 0x99, 0x23, 0x7f, 0x1c, 0x78, 0x51, 0x2b, 0x3e, 0xab, 0xb1, 0x8f, 0xf0, 0xea, 0x70,
	0x3f, 0x43, 0xfa, 0x8b, 0xa8, 0x62, 0x5d, 0xa9, 0xdc, 0xf2, 0x96, 0x3b, 0xbc, 0x1b, 0x60, 0xdf,
	0x20, 0x3c, 0x14, 0x97, 0xda, 0xfc, 0xb0, 0x68, 0xa6, 0x19, 0x49, 0xad, 0xa6, 0x4a, 0x96, 0xa2,
	0x7e, 0xf0, 0xb6, 0xb3, 0x79, 0x00, 0xc9, 0x3e, 0xc1, 0xe8, 0x16, 0x35, 0xc7, 0x02, 0x53, 0xd3,
	0xf8, 0x74, 0x9c, 0xcf, 0x05, 0x82, 0xcd, 0xe0, 0xf5, 0x71, 0xd8, 0xad, 0x36, 0xb8, 0x94, 0xc6,
	0x92, 0x0b, 0xcc, 0xbb, 0xce, 0xe1, 0x22, 0xc3, 0x96, 0xf0, 0xe6, 0x54, 0x5f, 0x09, 0x4b, 0x75,
	0xe0, 0xcf, 0xff, 0x24, 0xf1, 0x9e, 0x73, 0xb9, 0x87, 0x62, 0xef, 0x60, 0x20, 0xf3, 0x02, 0x93,
	0xca, 0xa8, 0x95, 0x2c, 0x25, 0xf1, 0x27, 0x81, 0x17, 0x75, 0xe2, 0xe3, 0x65, 0x4d, 0xa9, 0xaa,
	0x4c, 0xd1, 0xac, 0xb7, 0x1b, 0x14, 0x64, 0xb9, 0xdf, 0x50, 0x47, 0x4b, 0xf6, 0x1e, 0x9e, 0x5b,
	0x12, 0x86, 0xa4, 0xfa, 0xb9, 0x21, 0xf1, 0x0b, 0xa7, 0xa5, 0xae, 0x14, 0x71, 0x70, 0x41, 0xce,
	0x49, 0x6c, 0x08, 0x9d, 0xca, 0xa2, 0xb1, 0xbc, 0x1f, 0xb4, 0x22, 0x3f, 0x6e, 0x86, 0xf0, 0x19,
	0x0c, 0xf6, 0x5d, 0xb2, 0x3b, 0xad, 0x2c, 0x86, 0x6f, 0xc1, 0xff, 0xaa, 0xa5, 0x72, 0x4b, 0xf6,
	0x02, 0xba, 0x35, 0xe6, 0x8a, 0x55, 0x1f, 0xed, 0xa7, 0xc9, 0x1a, 0xfc, 0x04, 0x51, 0x7d, 0xaf,
	0xfb, 0xc8, 0x66, 0xd0, 0x9f, 0x1b, 0x14, 0x84, 0xcd, 0xcd, 0xcb, 0xf1, 0x4d, 0x55, 0xc7, 0x87,
	0x35, 0x1d, 0xf1, 0x53, 0x61, 0xff, 0xe7, 0xa3, 0xb4, 0xeb, 0x5a, 0xfe, 0xe1, 0x2a, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0x37, 0x97, 0x21, 0xf6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TeenPattiClient is the client API for TeenPatti service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeenPattiClient interface {
	CreateTable(ctx context.Context, in *TableRequest, opts ...grpc.CallOption) (*TableResponse, error)
}

type teenPattiClient struct {
	cc *grpc.ClientConn
}

func NewTeenPattiClient(cc *grpc.ClientConn) TeenPattiClient {
	return &teenPattiClient{cc}
}

func (c *teenPattiClient) CreateTable(ctx context.Context, in *TableRequest, opts ...grpc.CallOption) (*TableResponse, error) {
	out := new(TableResponse)
	err := c.cc.Invoke(ctx, "/teenpatti.TeenPatti/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeenPattiServer is the server API for TeenPatti service.
type TeenPattiServer interface {
	CreateTable(context.Context, *TableRequest) (*TableResponse, error)
}

// UnimplementedTeenPattiServer can be embedded to have forward compatible implementations.
type UnimplementedTeenPattiServer struct {
}

func (*UnimplementedTeenPattiServer) CreateTable(ctx context.Context, req *TableRequest) (*TableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}

func RegisterTeenPattiServer(s *grpc.Server, srv TeenPattiServer) {
	s.RegisterService(&_TeenPatti_serviceDesc, srv)
}

func _TeenPatti_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeenPattiServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teenpatti.TeenPatti/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeenPattiServer).CreateTable(ctx, req.(*TableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TeenPatti_serviceDesc = grpc.ServiceDesc{
	ServiceName: "teenpatti.TeenPatti",
	HandlerType: (*TeenPattiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _TeenPatti_CreateTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
