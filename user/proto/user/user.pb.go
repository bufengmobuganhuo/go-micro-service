// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

/*
Package go_micro_service_user is a generated protocol buffer package.

It is generated from these files:
	proto/user/user.proto

It has these top-level messages:
	UserRegisterRequest
	UserRegisterResponse
	UserLoginRequest
	UserLoginResponse
	UserInfoRequest
	UserInfoResponse
*/
package go_micro_service_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserRegisterRequest struct {
	UserName  string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	Pwd       string `protobuf:"bytes,3,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *UserRegisterRequest) Reset()                    { *m = UserRegisterRequest{} }
func (m *UserRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRegisterRequest) ProtoMessage()               {}
func (*UserRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRegisterRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserRegisterRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserRegisterRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserRegisterResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UserRegisterResponse) Reset()                    { *m = UserRegisterResponse{} }
func (m *UserRegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*UserRegisterResponse) ProtoMessage()               {}
func (*UserRegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserRegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserLoginRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Pwd      string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *UserLoginRequest) Reset()                    { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()               {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserLoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserLoginRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserLoginResponse struct {
	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess" json:"is_success,omitempty"`
}

func (m *UserLoginResponse) Reset()                    { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string            { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()               {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserLoginResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

type UserInfoRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *UserInfoRequest) Reset()                    { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()               {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserInfoRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type UserInfoResponse struct {
	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
}

func (m *UserInfoResponse) Reset()                    { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()               {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserInfoResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfoResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfoResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRegisterRequest)(nil), "go.micro.service.user.UserRegisterRequest")
	proto.RegisterType((*UserRegisterResponse)(nil), "go.micro.service.user.UserRegisterResponse")
	proto.RegisterType((*UserLoginRequest)(nil), "go.micro.service.user.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "go.micro.service.user.UserLoginResponse")
	proto.RegisterType((*UserInfoRequest)(nil), "go.micro.service.user.UserInfoRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "go.micro.service.user.UserInfoResponse")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0xb5, 0xa0, 0x6d, 0x19, 0x0f, 0xd6, 0xd5, 0xc6, 0xa6, 0xa6, 0x89, 0xd9, 0x83, 0x6d, 0x34,
	0x59, 0x4d, 0xfd, 0x02, 0x4f, 0xa6, 0x89, 0xf1, 0x80, 0xf1, 0x66, 0x52, 0x2b, 0x4c, 0xc9, 0x1e,
	0x60, 0x71, 0x07, 0xf4, 0xab, 0xfc, 0x47, 0xb3, 0x0b, 0x28, 0x35, 0x0d, 0xe5, 0x02, 0xcc, 0xbc,
	0xf7, 0xe6, 0x2d, 0x6f, 0x16, 0x86, 0xa9, 0x56, 0x99, 0xba, 0xc9, 0x09, 0xb5, 0x7d, 0x08, 0x5b,
	0xb3, 0x61, 0xa4, 0x44, 0x2c, 0x03, 0xad, 0x04, 0xa1, 0xfe, 0x94, 0x01, 0x0a, 0x03, 0xf2, 0x00,
	0x4e, 0x5e, 0x08, 0xb5, 0x8f, 0x91, 0xa4, 0xcc, 0xbc, 0x3f, 0x72, 0xa4, 0x8c, 0x9d, 0x83, 0x67,
	0xe0, 0x65, 0xb2, 0x8a, 0x71, 0xd4, 0xb9, 0xe8, 0xcc, 0x3c, 0xbf, 0x6f, 0x1a, 0x4f, 0xab, 0x18,
	0xd9, 0x04, 0x60, 0x2d, 0x35, 0x65, 0x05, 0xea, 0x58, 0xd4, 0xb3, 0x1d, 0x0b, 0x0f, 0xc0, 0x4d,
	0xbf, 0xc2, 0x91, 0x6b, 0xfb, 0xe6, 0x93, 0xdf, 0xc2, 0xe9, 0xa6, 0x09, 0xa5, 0x2a, 0x21, 0x64,
	0x23, 0xe8, 0xc5, 0x48, 0xb4, 0x8a, 0x2a, 0x8f, 0xaa, 0xe4, 0xf7, 0x30, 0x30, 0x8a, 0x47, 0x15,
	0xc9, 0xa4, 0xd5, 0x99, 0x4a, 0x53, 0xe7, 0xcf, 0x74, 0x0e, 0xc7, 0xb5, 0x11, 0xa5, 0xe3, 0x04,
	0x40, 0xd2, 0x92, 0xf2, 0x20, 0x40, 0x22, 0x3b, 0xa4, 0xef, 0x7b, 0x92, 0x9e, 0x8b, 0x06, 0x17,
	0x70, 0x64, 0x34, 0x8b, 0x64, 0xad, 0xda, 0xb8, 0xf2, 0xa8, 0x38, 0x66, 0xc1, 0x2f, 0x2d, 0xce,
	0xa0, 0x67, 0x05, 0x32, 0xb4, 0x74, 0xd7, 0xef, 0x9a, 0x72, 0x11, 0x6e, 0x4e, 0x72, 0x1a, 0x33,
	0x75, 0xff, 0x65, 0x3a, 0xff, 0x76, 0x60, 0xdf, 0x38, 0x31, 0x84, 0x7e, 0x15, 0x23, 0xbb, 0x12,
	0x5b, 0x77, 0x2a, 0xb6, 0x2c, 0x74, 0x7c, 0xdd, 0x8a, 0x5b, 0xfc, 0x02, 0xdf, 0x63, 0xaf, 0x70,
	0x60, 0x83, 0x63, 0xd3, 0x06, 0x5d, 0x7d, 0x3b, 0xe3, 0xd9, 0x6e, 0xe2, 0xef, 0xf4, 0x37, 0x38,
	0x7c, 0xc0, 0xac, 0x4a, 0x8e, 0x5d, 0x36, 0x48, 0x6b, 0xab, 0x18, 0x4f, 0x77, 0xf2, 0x2a, 0x87,
	0xf7, 0xae, 0xbd, 0xf4, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x77, 0x6b, 0x81, 0x0d,
	0x03, 0x00, 0x00,
}