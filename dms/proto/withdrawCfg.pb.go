// Code generated by protoc-gen-go. DO NOT EDIT.
// source: withdrawCfg.proto

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
	DailyLimit           float32  `protobuf:"fixed32,3,opt,name=daily_limit,json=dailyLimit,proto3" json:"daily_limit,omitempty"`
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
	return fileDescriptor_9bdc854c8b4e0934, []int{0}
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

func (m *WithdrawCfg) GetDailyLimit() float32 {
	if m != nil {
		return m.DailyLimit
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
	return fileDescriptor_9bdc854c8b4e0934, []int{1}
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
	proto.RegisterFile("withdrawCfg.proto", fileDescriptor_9bdc854c8b4e0934)
}

var fileDescriptor_9bdc854c8b4e0934 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0x2b, 0xc9, 0x36, 0xf5, 0xa8, 0x98, 0x76, 0xda, 0x82, 0xb0, 0x29, 0x15, 0xee, 0x45,
	0xd0, 0x22, 0xa8, 0xfc, 0x04, 0xc6, 0x84, 0x60, 0x48, 0x20, 0xac, 0x0f, 0x39, 0x0a, 0xc5, 0x1a,
	0x3b, 0x0b, 0xde, 0x5d, 0x65, 0xb5, 0xb1, 0xd1, 0x31, 0xef, 0x92, 0x77, 0xc9, 0x6b, 0x05, 0xad,
	0x14, 0xe2, 0x38, 0xc1, 0xc9, 0x51, 0xf3, 0xfd, 0xf3, 0x8f, 0xfe, 0x9f, 0x85, 0x6f, 0x3b, 0x6e,
	0xae, 0x73, 0x9d, 0xed, 0x66, 0xab, 0x75, 0x5c, 0x68, 0x65, 0x14, 0x0e, 0xd6, 0xc4, 0x6f, 0xb8,
	0x8c, 0x4b, 0xbd, 0x8d, 0x73, 0x51, 0x0e, 0xbf, 0x2c, 0x95, 0x10, 0x4a, 0x36, 0x74, 0xfc, 0xe0,
	0x80, 0x7f, 0xf9, 0xbc, 0x83, 0x03, 0x70, 0x79, 0x1e, 0x38, 0xa1, 0x13, 0x79, 0xcc, 0xe5, 0x39,
	0x8e, 0xa0, 0x2f, 0xb8, 0x4c, 0x85, 0x92, 0x54, 0x05, 0x6e, 0xe8, 0x44, 0x2e, 0xfb, 0x2c, 0xb8,
	0x3c, 0xaf, 0xbf, 0xf1, 0x37, 0xf8, 0x79, 0xc6, 0x37, 0x55, 0xba, 0xe1, 0x82, 0x9b, 0xc0, 0xb3,
	0x18, 0xec, 0xe8, 0xac, 0x9e, 0xe0, 0x3f, 0xc0, 0x92, 0xf4, 0x96, 0x2f, 0xa9, 0x71, 0x48, 0x75,
	0x66, 0x28, 0xe8, 0x58, 0xdd, 0xd7, 0x96, 0x58, 0x2b, 0x96, 0x19, 0xc2, 0x5f, 0x00, 0x4b, 0x4d,
	0x99, 0xa1, 0x3c, 0xcd, 0x4c, 0xd0, 0x0d, 0x9d, 0xa8, 0xcf, 0xfa, 0xed, 0x64, 0x6a, 0x6a, 0x7c,
	0x5b, 0xe4, 0x4f, 0xb8, 0xd7, 0xe0, 0x76, 0x32, 0x35, 0xe3, 0x3b, 0x17, 0xbe, 0xef, 0x25, 0x61,
	0x54, 0x16, 0x4a, 0x96, 0x84, 0x13, 0xe8, 0x91, 0x34, 0xdc, 0x54, 0x36, 0x95, 0x9f, 0x8c, 0xe2,
	0x97, 0x85, 0xc4, 0xfb, 0x4b, 0xad, 0x14, 0xff, 0x42, 0xb7, 0xc8, 0xd6, 0xa4, 0x6d, 0x64, 0x3f,
	0xf9, 0x79, 0xb8, 0x73, 0x51, 0x43, 0xd6, 0x68, 0xf0, 0x3f, 0x74, 0xb9, 0x21, 0x51, 0x06, 0x5e,
	0xe8, 0xbd, 0x77, 0xa0, 0x51, 0xd6, 0xfe, 0xa4, 0xb5, 0xd2, 0xb6, 0x8b, 0x37, 0xfc, 0x4f, 0x6a,
	0xc8, 0x1a, 0x0d, 0x46, 0xd0, 0xe1, 0x72, 0xa5, 0x6c, 0x23, 0x7e, 0xf2, 0xe3, 0x50, 0x3b, 0x97,
	0x2b, 0xc5, 0xac, 0x22, 0xb9, 0x77, 0x00, 0xf7, 0xae, 0x2d, 0x9a, 0x86, 0x71, 0x06, 0xde, 0x29,
	0x19, 0x7c, 0x7d, 0x45, 0x14, 0xa6, 0x1a, 0xfe, 0x39, 0xf6, 0xbf, 0x6d, 0x8b, 0xe3, 0x4f, 0x38,
	0x07, 0x6f, 0x41, 0x06, 0x8f, 0xa5, 0xfb, 0xa0, 0xd5, 0x55, 0xcf, 0xbe, 0xbd, 0xc9, 0x63, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8f, 0xb1, 0x0d, 0x61, 0xae, 0x02, 0x00, 0x00,
}