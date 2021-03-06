// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type GetUserRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Sex                  int32    `protobuf:"varint,3,opt,name=sex,proto3" json:"sex"`
	Symbol               string   `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol"`
	Balance              float64  `protobuf:"fixed64,5,opt,name=balance,proto3" json:"balance"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetUserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUserResponse) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *GetUserResponse) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *GetUserResponse) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "user.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "user.GetUserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x31, 0x0f, 0x82, 0x30,
	0x10, 0x85, 0x2d, 0x14, 0x88, 0x37, 0xa0, 0xb9, 0xa8, 0x69, 0x9c, 0x1a, 0xa6, 0x4e, 0x0c, 0xba,
	0x38, 0xba, 0xb9, 0x37, 0xf1, 0x07, 0x80, 0xdc, 0x40, 0x02, 0x14, 0x29, 0x24, 0xf2, 0xef, 0x0d,
	0x15, 0x4d, 0x64, 0x7b, 0xef, 0xf2, 0x25, 0xf7, 0x3d, 0x80, 0xc1, 0x52, 0x97, 0xb6, 0x9d, 0xe9,
	0x0d, 0xf2, 0x29, 0x27, 0x12, 0xe2, 0x1b, 0xf5, 0x77, 0x4b, 0x9d, 0xa6, 0xe7, 0x40, 0xb6, 0xc7,
	0x18, 0xbc, 0xb2, 0x10, 0x4c, 0x32, 0xe5, 0x6b, 0xaf, 0x2c, 0x92, 0x11, 0x36, 0x3f, 0xc2, 0xb6,
	0xa6, 0xb1, 0xb4, 0x44, 0x10, 0x81, 0x37, 0x59, 0x4d, 0xc2, 0x93, 0x4c, 0xad, 0xb5, 0xcb, 0xb8,
	0x05, 0xdf, 0xd2, 0x4b, 0xf8, 0x92, 0xa9, 0x40, 0x4f, 0x11, 0x0f, 0x10, 0xda, 0xb1, 0xce, 0x4d,
	0x25, 0xb8, 0xe3, 0xe6, 0x86, 0x02, 0xa2, 0x3c, 0xab, 0xb2, 0xe6, 0x41, 0x22, 0x90, 0x4c, 0x31,
	0xfd, 0xad, 0xa7, 0x2b, 0xf0, 0xe9, 0x2f, 0x5e, 0x20, 0x9a, 0x15, 0x70, 0x97, 0xba, 0x09, 0xff,
	0xce, 0xc7, 0xfd, 0xe2, 0xfa, 0xf1, 0x4c, 0x56, 0x79, 0xe8, 0xb6, 0x9e, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x07, 0x2b, 0x37, 0x06, 0xf9, 0x00, 0x00, 0x00,
}
