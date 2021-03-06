// Code generated by protoc-gen-go. DO NOT EDIT.
// source: integral.proto

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

type Integral struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Integral) Reset()         { *m = Integral{} }
func (m *Integral) String() string { return proto.CompactTextString(m) }
func (*Integral) ProtoMessage()    {}
func (*Integral) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d836c8d70e1b53, []int{0}
}

func (m *Integral) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Integral.Unmarshal(m, b)
}
func (m *Integral) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Integral.Marshal(b, m, deterministic)
}
func (m *Integral) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Integral.Merge(m, src)
}
func (m *Integral) XXX_Size() int {
	return xxx_messageInfo_Integral.Size(m)
}
func (m *Integral) XXX_DiscardUnknown() {
	xxx_messageInfo_Integral.DiscardUnknown(m)
}

var xxx_messageInfo_Integral proto.InternalMessageInfo

func (m *Integral) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

//
type IntegralResponse struct {
	Entity               *Integral   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Integral `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *IntegralResponse) Reset()         { *m = IntegralResponse{} }
func (m *IntegralResponse) String() string { return proto.CompactTextString(m) }
func (*IntegralResponse) ProtoMessage()    {}
func (*IntegralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d836c8d70e1b53, []int{1}
}

func (m *IntegralResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegralResponse.Unmarshal(m, b)
}
func (m *IntegralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegralResponse.Marshal(b, m, deterministic)
}
func (m *IntegralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegralResponse.Merge(m, src)
}
func (m *IntegralResponse) XXX_Size() int {
	return xxx_messageInfo_IntegralResponse.Size(m)
}
func (m *IntegralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntegralResponse proto.InternalMessageInfo

func (m *IntegralResponse) GetEntity() *Integral {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *IntegralResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *IntegralResponse) GetItems() []*Integral {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *IntegralResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *IntegralResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Integral)(nil), "geiqin.srv.crm.Integral")
	proto.RegisterType((*IntegralResponse)(nil), "geiqin.srv.crm.IntegralResponse")
}

func init() {
	proto.RegisterFile("integral.proto", fileDescriptor_d3d836c8d70e1b53)
}

var fileDescriptor_d3d836c8d70e1b53 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xdd, 0x4a, 0xfb, 0x30,
	0x18, 0x87, 0xff, 0x6d, 0xd7, 0xf2, 0xe7, 0x55, 0xaa, 0x04, 0x85, 0xd8, 0xa3, 0xd2, 0xa3, 0x81,
	0x10, 0x64, 0xde, 0x81, 0x1f, 0x0c, 0x11, 0x44, 0xba, 0x03, 0x8f, 0x63, 0xf7, 0xae, 0x0b, 0xac,
	0x49, 0x7d, 0x13, 0x06, 0x5e, 0xa1, 0xb7, 0xe4, 0xa1, 0x34, 0xed, 0x04, 0xe7, 0xe7, 0xc1, 0x8e,
	0x9f, 0xe7, 0xf7, 0x24, 0x84, 0x40, 0xaa, 0xb4, 0xc3, 0x9a, 0xe4, 0x4a, 0xb4, 0x64, 0x9c, 0x61,
	0x69, 0x8d, 0xea, 0x49, 0x69, 0x61, 0x69, 0x2d, 0x2a, 0x6a, 0xb2, 0xfd, 0xca, 0x34, 0x8d, 0xd1,
	0x3d, 0x2d, 0x32, 0xf8, 0x7f, 0x33, 0xf8, 0x2c, 0x85, 0x50, 0xcd, 0x79, 0x90, 0x07, 0xe3, 0xb8,
	0x0c, 0xd5, 0xbc, 0x78, 0x0d, 0xe0, 0x70, 0x03, 0x4b, 0xb4, 0xad, 0xd1, 0x16, 0xd9, 0x19, 0x24,
	0xa8, 0x9d, 0x72, 0xcf, 0x5e, 0xdc, 0x9b, 0x70, 0xf1, 0xb1, 0x2f, 0xde, 0x17, 0x83, 0xc7, 0x4e,
	0x21, 0x6e, 0x65, 0x8d, 0xc4, 0x43, 0x3f, 0x38, 0xde, 0x1e, 0xdc, 0x77, 0xb0, 0xec, 0x1d, 0x26,
	0x20, 0x56, 0x0e, 0x1b, 0xcb, 0xa3, 0x3c, 0xfa, 0xb1, 0xde, 0x6b, 0x5d, 0x1c, 0x89, 0x0c, 0xf1,
	0xd1, 0xd7, 0xf1, 0xeb, 0x0e, 0x96, 0xbd, 0xc3, 0xc6, 0x30, 0x52, 0x7a, 0x61, 0x78, 0xec, 0xdd,
	0xa3, 0xcf, 0xed, 0x85, 0x29, 0xbd, 0x31, 0x79, 0x09, 0xe1, 0x60, 0x73, 0xd4, 0x0c, 0x69, 0xad,
	0x2a, 0x64, 0x77, 0x90, 0x5c, 0x12, 0x4a, 0x87, 0x2c, 0xff, 0xf6, 0x56, 0xc3, 0x2b, 0x65, 0xbf,
	0x1a, 0xc5, 0xbf, 0xae, 0x77, 0x85, 0x2b, 0xdc, 0x59, 0xef, 0x16, 0xa2, 0x29, 0xba, 0x1d, 0xc5,
	0xa6, 0x90, 0xcc, 0x50, 0x52, 0xb5, 0x64, 0x27, 0xdb, 0xf6, 0x85, 0xb4, 0xf8, 0xb0, 0x44, 0xfa,
	0x53, 0xe8, 0x31, 0xf1, 0xff, 0xec, 0xfc, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x44, 0x84, 0xbe,
	0x97, 0x02, 0x00, 0x00,
}
