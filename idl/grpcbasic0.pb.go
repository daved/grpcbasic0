// Code generated by protoc-gen-go.
// source: grpcbasic0.proto
// DO NOT EDIT!

/*
Package idl is a generated protocol buffer package.

User Service

User Service API provides create, read, and read (many) access to a list of
users.

It is generated from these files:
	grpcbasic0.proto

It has these top-level messages:
	User
	UserRecordReq
	UserGetReq
	Users
	UsersGetReq
*/
package idl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id      int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age     int64  `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
	Fortune string `protobuf:"bytes,4,opt,name=fortune" json:"fortune,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetFortune() string {
	if m != nil {
		return m.Fortune
	}
	return ""
}

type UserRecordReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age  int64  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
}

func (m *UserRecordReq) Reset()                    { *m = UserRecordReq{} }
func (m *UserRecordReq) String() string            { return proto.CompactTextString(m) }
func (*UserRecordReq) ProtoMessage()               {}
func (*UserRecordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserRecordReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRecordReq) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

type UserGetReq struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *UserGetReq) Reset()                    { *m = UserGetReq{} }
func (m *UserGetReq) String() string            { return proto.CompactTextString(m) }
func (*UserGetReq) ProtoMessage()               {}
func (*UserGetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserGetReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Users struct {
	Users []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *Users) Reset()                    { *m = Users{} }
func (m *Users) String() string            { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()               {}
func (*Users) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UsersGetReq struct {
	Start int64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Desc  bool  `protobuf:"varint,3,opt,name=desc" json:"desc,omitempty"`
}

func (m *UsersGetReq) Reset()                    { *m = UsersGetReq{} }
func (m *UsersGetReq) String() string            { return proto.CompactTextString(m) }
func (*UsersGetReq) ProtoMessage()               {}
func (*UsersGetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UsersGetReq) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *UsersGetReq) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *UsersGetReq) GetDesc() bool {
	if m != nil {
		return m.Desc
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "idl.User")
	proto.RegisterType((*UserRecordReq)(nil), "idl.UserRecordReq")
	proto.RegisterType((*UserGetReq)(nil), "idl.UserGetReq")
	proto.RegisterType((*Users)(nil), "idl.Users")
	proto.RegisterType((*UsersGetReq)(nil), "idl.UsersGetReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	RecordUser(ctx context.Context, in *UserRecordReq, opts ...grpc.CallOption) (*User, error)
	GetUser(ctx context.Context, in *UserGetReq, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *UsersGetReq, opts ...grpc.CallOption) (*Users, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) RecordUser(ctx context.Context, in *UserRecordReq, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/idl.UserService/RecordUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *UserGetReq, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/idl.UserService/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *UsersGetReq, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := grpc.Invoke(ctx, "/idl.UserService/GetUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	RecordUser(context.Context, *UserRecordReq) (*User, error)
	GetUser(context.Context, *UserGetReq) (*User, error)
	GetUsers(context.Context, *UsersGetReq) (*Users, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_RecordUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RecordUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.UserService/RecordUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RecordUser(ctx, req.(*UserRecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*UserGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*UsersGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordUser",
			Handler:    _UserService_RecordUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcbasic0.proto",
}

func init() { proto.RegisterFile("grpcbasic0.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0xa4, 0xa5, 0xed, 0x55, 0xa5, 0xe5, 0x00, 0x29, 0xaa, 0x2a, 0x51, 0x79, 0x8a,
	0x18, 0x1a, 0x28, 0x62, 0x01, 0x06, 0xb6, 0x4e, 0x2c, 0x41, 0xb0, 0xbb, 0xb1, 0x89, 0x2c, 0x95,
	0xb8, 0xd8, 0x6e, 0x17, 0xc4, 0xc2, 0x2b, 0xf0, 0x56, 0xac, 0xbc, 0x02, 0x0f, 0x82, 0x6c, 0x27,
	0x29, 0x82, 0xed, 0xee, 0xf2, 0xff, 0xdf, 0xfd, 0x17, 0xc3, 0xa8, 0x50, 0xeb, 0x7c, 0x49, 0xb5,
	0xc8, 0xcf, 0x66, 0x6b, 0x25, 0x8d, 0xc4, 0x48, 0xb0, 0xd5, 0x78, 0x52, 0x48, 0x59, 0xac, 0x78,
	0x4a, 0xd7, 0x22, 0xa5, 0x65, 0x29, 0x0d, 0x35, 0x42, 0x96, 0xda, 0x4b, 0xc8, 0x23, 0xb4, 0x1e,
	0x34, 0x57, 0xb8, 0x0f, 0xa1, 0x60, 0x71, 0x30, 0x0d, 0x92, 0x28, 0x0b, 0x05, 0x43, 0x84, 0x56,
	0x49, 0x9f, 0x79, 0x1c, 0x4e, 0x83, 0xa4, 0x97, 0xb9, 0x1a, 0x47, 0x10, 0xd1, 0x82, 0xc7, 0x91,
	0x13, 0xd9, 0x12, 0x63, 0xe8, 0x3c, 0x49, 0x65, 0x36, 0x25, 0x8f, 0x5b, 0x4e, 0x58, 0xb7, 0xe4,
	0x12, 0x06, 0x96, 0x9b, 0xf1, 0x5c, 0x2a, 0x96, 0xf1, 0x97, 0x06, 0x18, 0xfc, 0x07, 0x86, 0x0d,
	0x90, 0x4c, 0x00, 0xac, 0x6d, 0xc1, 0x8d, 0xf5, 0xfc, 0x09, 0x45, 0x12, 0x68, 0xdb, 0xaf, 0x1a,
	0x4f, 0xa0, 0xbd, 0xb1, 0x45, 0x1c, 0x4c, 0xa3, 0xa4, 0x3f, 0xef, 0xcd, 0x04, 0x5b, 0xcd, 0xdc,
	0x3e, 0x3f, 0x27, 0x77, 0xd0, 0x77, 0xca, 0x0a, 0x74, 0x04, 0x6d, 0x6d, 0xa8, 0x32, 0x15, 0xcb,
	0x37, 0x76, 0x9a, 0xcb, 0x4d, 0x69, 0xaa, 0x00, 0xbe, 0xb1, 0x41, 0x19, 0xd7, 0xb9, 0x3b, 0xb3,
	0x9b, 0xb9, 0x7a, 0xfe, 0x19, 0x78, 0xde, 0x3d, 0x57, 0x5b, 0x91, 0x73, 0xbc, 0x05, 0xf0, 0x97,
	0xb9, 0x7f, 0x87, 0xbb, 0xf5, 0xf5, 0xb9, 0xe3, 0x5d, 0x24, 0x72, 0xf8, 0xfe, 0xf5, 0xfd, 0x11,
	0x0e, 0x48, 0x37, 0xdd, 0x9e, 0xa7, 0x36, 0xde, 0x55, 0x70, 0x8a, 0x37, 0xd0, 0x59, 0x70, 0xe3,
	0xec, 0xc3, 0x46, 0xea, 0xd3, 0xfe, 0xf6, 0x1e, 0x3b, 0xef, 0x10, 0x07, 0xb5, 0x37, 0x7d, 0x15,
	0xec, 0x0d, 0xaf, 0xa1, 0x5b, 0xb9, 0x35, 0x8e, 0x1a, 0x75, 0x75, 0xed, 0x18, 0x76, 0x13, 0x72,
	0xe0, 0x00, 0x7d, 0xec, 0xd5, 0x00, 0xbd, 0xdc, 0x73, 0x2f, 0x7f, 0xf1, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x80, 0x52, 0x6c, 0x91, 0x30, 0x02, 0x00, 0x00,
}
