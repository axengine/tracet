// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet.proto

package wallet

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

type GetWalletRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWalletRequest) Reset()         { *m = GetWalletRequest{} }
func (m *GetWalletRequest) String() string { return proto.CompactTextString(m) }
func (*GetWalletRequest) ProtoMessage()    {}
func (*GetWalletRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{0}
}

func (m *GetWalletRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWalletRequest.Unmarshal(m, b)
}
func (m *GetWalletRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWalletRequest.Marshal(b, m, deterministic)
}
func (m *GetWalletRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWalletRequest.Merge(m, src)
}
func (m *GetWalletRequest) XXX_Size() int {
	return xxx_messageInfo_GetWalletRequest.Size(m)
}
func (m *GetWalletRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWalletRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWalletRequest proto.InternalMessageInfo

func (m *GetWalletRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetWalletResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol"`
	Balance              float64  `protobuf:"fixed64,3,opt,name=balance,proto3" json:"balance"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWalletResponse) Reset()         { *m = GetWalletResponse{} }
func (m *GetWalletResponse) String() string { return proto.CompactTextString(m) }
func (*GetWalletResponse) ProtoMessage()    {}
func (*GetWalletResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{1}
}

func (m *GetWalletResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWalletResponse.Unmarshal(m, b)
}
func (m *GetWalletResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWalletResponse.Marshal(b, m, deterministic)
}
func (m *GetWalletResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWalletResponse.Merge(m, src)
}
func (m *GetWalletResponse) XXX_Size() int {
	return xxx_messageInfo_GetWalletResponse.Size(m)
}
func (m *GetWalletResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWalletResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWalletResponse proto.InternalMessageInfo

func (m *GetWalletResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetWalletResponse) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *GetWalletResponse) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*GetWalletRequest)(nil), "wallet.GetWalletRequest")
	proto.RegisterType((*GetWalletResponse)(nil), "wallet.GetWalletResponse")
}

func init() { proto.RegisterFile("wallet.proto", fileDescriptor_b88fd140af4deb6f) }

var fileDescriptor_b88fd140af4deb6f = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4f, 0xcc, 0xc9,
	0x49, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x94, 0xb8, 0x04,
	0xdc, 0x53, 0x4b, 0xc2, 0xc1, 0x9c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x3e, 0x2e,
	0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0xa6, 0xcc, 0x14, 0xa5, 0x50, 0x2e,
	0x41, 0x24, 0x35, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0xe8, 0x8a, 0x84, 0xc4, 0xb8, 0xd8, 0x8a,
	0x2b, 0x73, 0x93, 0xf2, 0x73, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x09,
	0x2e, 0xf6, 0xa4, 0xc4, 0x9c, 0xc4, 0xbc, 0xe4, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xc6, 0x20,
	0x18, 0xd7, 0xc8, 0x87, 0x8b, 0x0d, 0x62, 0xa6, 0x90, 0x13, 0x17, 0x27, 0xdc, 0x02, 0x21, 0x09,
	0x3d, 0xa8, 0x43, 0xd1, 0xdd, 0x25, 0x25, 0x89, 0x45, 0x06, 0xe2, 0x1a, 0x25, 0x86, 0x24, 0x36,
	0xb0, 0xbf, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x8c, 0x2b, 0x5b, 0xe7, 0x00, 0x00,
	0x00,
}
