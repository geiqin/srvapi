// Code generated by protoc-gen-go. DO NOT EDIT.
// source: withDrawCfg.proto

package geiqin_srv_dms

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

type WithdrawCfg struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MinMoney             float32  `protobuf:"fixed32,2,opt,name=min_money,json=minMoney,proto3" json:"min_money,omitempty"`
	MaxMoney             float32  `protobuf:"fixed32,3,opt,name=max_money,json=maxMoney,proto3" json:"max_money,omitempty"`
	ServiceMoneyRate     float32  `protobuf:"fixed32,4,opt,name=service_money_rate,json=serviceMoneyRate,proto3" json:"service_money_rate,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawCfg) Reset()         { *m = WithdrawCfg{} }
func (m *WithdrawCfg) String() string { return proto.CompactTextString(m) }
func (*WithdrawCfg) ProtoMessage()    {}
func (*WithdrawCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e1aa2ad12578cb, []int{0}
}

func (m *WithdrawCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawCfg.Unmarshal(m, b)
}
func (m *WithdrawCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawCfg.Marshal(b, m, deterministic)
}
func (m *WithdrawCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawCfg.Merge(m, src)
}
func (m *WithdrawCfg) XXX_Size() int {
	return xxx_messageInfo_WithdrawCfg.Size(m)
}
func (m *WithdrawCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawCfg.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawCfg proto.InternalMessageInfo

func (m *WithdrawCfg) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WithdrawCfg) GetMinMoney() float32 {
	if m != nil {
		return m.MinMoney
	}
	return 0
}

func (m *WithdrawCfg) GetMaxMoney() float32 {
	if m != nil {
		return m.MaxMoney
	}
	return 0
}

func (m *WithdrawCfg) GetServiceMoneyRate() float32 {
	if m != nil {
		return m.ServiceMoneyRate
	}
	return 0
}

func (m *WithdrawCfg) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *WithdrawCfg) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type WithdrawCfgResponse struct {
	Entity               *WithdrawCfg   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager         `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*WithdrawCfg `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error         `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info          `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WithdrawCfgResponse) Reset()         { *m = WithdrawCfgResponse{} }
func (m *WithdrawCfgResponse) String() string { return proto.CompactTextString(m) }
func (*WithdrawCfgResponse) ProtoMessage()    {}
func (*WithdrawCfgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e1aa2ad12578cb, []int{1}
}

func (m *WithdrawCfgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawCfgResponse.Unmarshal(m, b)
}
func (m *WithdrawCfgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawCfgResponse.Marshal(b, m, deterministic)
}
func (m *WithdrawCfgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawCfgResponse.Merge(m, src)
}
func (m *WithdrawCfgResponse) XXX_Size() int {
	return xxx_messageInfo_WithdrawCfgResponse.Size(m)
}
func (m *WithdrawCfgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawCfgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawCfgResponse proto.InternalMessageInfo

func (m *WithdrawCfgResponse) GetEntity() *WithdrawCfg {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *WithdrawCfgResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *WithdrawCfgResponse) GetItems() []*WithdrawCfg {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *WithdrawCfgResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *WithdrawCfgResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*WithdrawCfg)(nil), "geiqin.srv.dms.WithdrawCfg")
	proto.RegisterType((*WithdrawCfgResponse)(nil), "geiqin.srv.dms.WithdrawCfgResponse")
}

func init() {
	proto.RegisterFile("withDrawCfg.proto", fileDescriptor_46e1aa2ad12578cb)
}

var fileDescriptor_46e1aa2ad12578cb = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0x87, 0x6f, 0x92, 0xb6, 0xdc, 0x4e, 0x2e, 0xe5, 0x3a, 0x2a, 0x84, 0x16, 0x21, 0xd4, 0x4d,
	0x40, 0x09, 0x98, 0x3e, 0x41, 0xa9, 0x22, 0x5d, 0x08, 0x32, 0x5d, 0xb8, 0x2c, 0x63, 0x73, 0xda,
	0xce, 0x62, 0x66, 0xe2, 0xe4, 0xd8, 0x3f, 0x4b, 0xdf, 0xc5, 0x37, 0xf1, 0xc5, 0x24, 0x33, 0x11,
	0x6a, 0x95, 0xea, 0x32, 0xe7, 0xfb, 0xce, 0x39, 0x39, 0x3f, 0x86, 0x1c, 0xad, 0x05, 0x2e, 0xaf,
	0x0d, 0x5f, 0x8f, 0xe6, 0x8b, 0xb4, 0x30, 0x1a, 0x35, 0xed, 0x2c, 0x40, 0x3c, 0x09, 0x95, 0x96,
	0x66, 0x95, 0xe6, 0xb2, 0xec, 0xfe, 0x9b, 0x69, 0x29, 0xb5, 0x72, 0xb4, 0xff, 0xe6, 0x91, 0xf0,
	0x41, 0xe0, 0x32, 0x77, 0x3d, 0xb4, 0x43, 0x7c, 0x91, 0x47, 0x5e, 0xec, 0x25, 0x01, 0xf3, 0x45,
	0x4e, 0x7b, 0xa4, 0x2d, 0x85, 0x9a, 0x4a, 0xad, 0x60, 0x1b, 0xf9, 0xb1, 0x97, 0xf8, 0xec, 0xaf,
	0x14, 0xea, 0xae, 0xfa, 0xb6, 0x90, 0x6f, 0x6a, 0x18, 0xd4, 0x90, 0x6f, 0x1c, 0xbc, 0x24, 0xb4,
	0x04, 0xb3, 0x12, 0x33, 0x70, 0xc2, 0xd4, 0x70, 0x84, 0xa8, 0x61, 0xad, 0xff, 0x35, 0xb1, 0x26,
	0xe3, 0x08, 0xf4, 0x8c, 0x90, 0x99, 0x01, 0x8e, 0x90, 0x4f, 0x39, 0x46, 0xcd, 0xd8, 0x4b, 0xda,
	0xac, 0x5d, 0x57, 0x86, 0x58, 0xe1, 0xe7, 0x22, 0xff, 0xc0, 0x2d, 0x87, 0xeb, 0xca, 0x10, 0xfb,
	0x2f, 0x3e, 0x39, 0xde, 0xb9, 0x82, 0x41, 0x59, 0x68, 0x55, 0x02, 0x1d, 0x90, 0x16, 0x28, 0x14,
	0xb8, 0xb5, 0x17, 0x85, 0x59, 0x2f, 0xfd, 0x1c, 0x46, 0xba, 0xdb, 0x54, 0xab, 0xf4, 0x82, 0x34,
	0x0b, 0xbe, 0x00, 0x63, 0xcf, 0x0d, 0xb3, 0xd3, 0xfd, 0x9e, 0xfb, 0x0a, 0x32, 0xe7, 0xd0, 0x2b,
	0xd2, 0x14, 0x08, 0xb2, 0x8c, 0x82, 0x38, 0xf8, 0x69, 0x81, 0x33, 0xab, 0xf9, 0x60, 0x8c, 0x36,
	0x36, 0x8b, 0x6f, 0xe6, 0xdf, 0x54, 0x90, 0x39, 0x87, 0x26, 0xa4, 0x21, 0xd4, 0x5c, 0xdb, 0x44,
	0xc2, 0xec, 0x64, 0xdf, 0x1d, 0xab, 0xb9, 0x66, 0xd6, 0xc8, 0x5e, 0x3d, 0x42, 0x77, 0xb6, 0x4d,
	0x5c, 0xc2, 0x74, 0x44, 0x82, 0x5b, 0x40, 0xfa, 0x75, 0x8b, 0x2c, 0x70, 0xdb, 0x3d, 0x3f, 0xf4,
	0xbf, 0x75, 0x8a, 0xfd, 0x3f, 0x74, 0x4c, 0x82, 0x09, 0x20, 0x3d, 0x74, 0xdd, 0x2f, 0x47, 0x3d,
	0xb6, 0xec, 0xbb, 0x1b, 0xbc, 0x07, 0x00, 0x00, 0xff, 0xff, 0x40, 0xd9, 0xf2, 0x82, 0xaa, 0x02,
	0x00, 0x00,
}
