// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customerSecret.proto

package geiqin_srv_crm

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

type CustomerSecret struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           int64    `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Locked               bool     `protobuf:"varint,4,opt,name=locked,proto3" json:"locked,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerSecret) Reset()         { *m = CustomerSecret{} }
func (m *CustomerSecret) String() string { return proto.CompactTextString(m) }
func (*CustomerSecret) ProtoMessage()    {}
func (*CustomerSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_26c6da4b2a1e067e, []int{0}
}

func (m *CustomerSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerSecret.Unmarshal(m, b)
}
func (m *CustomerSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerSecret.Marshal(b, m, deterministic)
}
func (m *CustomerSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerSecret.Merge(m, src)
}
func (m *CustomerSecret) XXX_Size() int {
	return xxx_messageInfo_CustomerSecret.Size(m)
}
func (m *CustomerSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerSecret.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerSecret proto.InternalMessageInfo

func (m *CustomerSecret) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerSecret) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *CustomerSecret) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CustomerSecret) GetLocked() bool {
	if m != nil {
		return m.Locked
	}
	return false
}

//
type CustomerSecretResponse struct {
	Entity               *CustomerSecret   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager            `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*CustomerSecret `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error            `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info             `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CustomerSecretResponse) Reset()         { *m = CustomerSecretResponse{} }
func (m *CustomerSecretResponse) String() string { return proto.CompactTextString(m) }
func (*CustomerSecretResponse) ProtoMessage()    {}
func (*CustomerSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_26c6da4b2a1e067e, []int{1}
}

func (m *CustomerSecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerSecretResponse.Unmarshal(m, b)
}
func (m *CustomerSecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerSecretResponse.Marshal(b, m, deterministic)
}
func (m *CustomerSecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerSecretResponse.Merge(m, src)
}
func (m *CustomerSecretResponse) XXX_Size() int {
	return xxx_messageInfo_CustomerSecretResponse.Size(m)
}
func (m *CustomerSecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerSecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerSecretResponse proto.InternalMessageInfo

func (m *CustomerSecretResponse) GetEntity() *CustomerSecret {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *CustomerSecretResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *CustomerSecretResponse) GetItems() []*CustomerSecret {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CustomerSecretResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CustomerSecretResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomerSecret)(nil), "geiqin.srv.crm.CustomerSecret")
	proto.RegisterType((*CustomerSecretResponse)(nil), "geiqin.srv.crm.CustomerSecretResponse")
}

func init() {
	proto.RegisterFile("customerSecret.proto", fileDescriptor_26c6da4b2a1e067e)
}

var fileDescriptor_26c6da4b2a1e067e = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0xd2, 0x86, 0x3a, 0x91, 0x1c, 0x96, 0x5a, 0x96, 0x1e, 0x34, 0xf4, 0x14, 0x10,
	0x72, 0xa8, 0xe2, 0x0b, 0x88, 0x87, 0xde, 0x64, 0x7d, 0x00, 0xa9, 0xd9, 0x69, 0x59, 0x34, 0x3b,
	0x71, 0x76, 0xab, 0xf8, 0x0a, 0x3e, 0xb5, 0x64, 0xd3, 0x0a, 0x09, 0x1e, 0x3c, 0xce, 0xfe, 0xdf,
	0xbf, 0xff, 0xfc, 0x0c, 0xcc, 0xeb, 0x83, 0xf3, 0xd4, 0x20, 0x3f, 0x61, 0xcd, 0xe8, 0xab, 0x96,
	0xc9, 0x93, 0xc8, 0xf7, 0x68, 0xde, 0x8d, 0xad, 0x1c, 0x7f, 0x54, 0x35, 0x37, 0xcb, 0xf3, 0x9a,
	0x9a, 0x86, 0x6c, 0xaf, 0xae, 0x0e, 0x90, 0xdf, 0x0f, 0x5c, 0x22, 0x87, 0xd8, 0x68, 0x19, 0x15,
	0x51, 0x99, 0xa8, 0xd8, 0x68, 0x71, 0x05, 0xd9, 0xe9, 0xdf, 0x67, 0xa3, 0x65, 0x1c, 0x04, 0x38,
	0x3d, 0x6d, 0xb4, 0x58, 0xc2, 0xac, 0xdd, 0x3a, 0xf7, 0x49, 0xac, 0x65, 0x52, 0x44, 0xe5, 0x99,
	0xfa, 0x9d, 0xc5, 0x02, 0xd2, 0x37, 0xaa, 0x5f, 0x51, 0xcb, 0x49, 0x11, 0x95, 0x33, 0x75, 0x9c,
	0x56, 0xdf, 0x31, 0x2c, 0x86, 0xb9, 0x0a, 0x5d, 0x4b, 0xd6, 0xa1, 0xb8, 0x83, 0x14, 0xad, 0x37,
	0xfe, 0x2b, 0xec, 0x90, 0xad, 0x2f, 0xab, 0x61, 0x81, 0x6a, 0xe4, 0x3b, 0xd2, 0xe2, 0x1a, 0xa6,
	0xed, 0x76, 0x8f, 0x1c, 0x36, 0xcc, 0xd6, 0x17, 0x63, 0xdb, 0x63, 0x27, 0xaa, 0x9e, 0x11, 0xb7,
	0x30, 0x35, 0x1e, 0x1b, 0x27, 0x93, 0x22, 0xf9, 0x47, 0x46, 0x0f, 0x77, 0x11, 0xc8, 0x4c, 0x1c,
	0xca, 0xfc, 0x11, 0xf1, 0xd0, 0x89, 0xaa, 0x67, 0x44, 0x09, 0x13, 0x63, 0x77, 0x24, 0xa7, 0x81,
	0x9d, 0x8f, 0xd9, 0x8d, 0xdd, 0x91, 0x0a, 0xc4, 0x4b, 0x1a, 0x4e, 0x71, 0xf3, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xa0, 0x4f, 0x31, 0x12, 0xc0, 0x01, 0x00, 0x00,
}
