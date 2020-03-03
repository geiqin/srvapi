// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fanWeixin.proto

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

type FanWeixin struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FanWeixin) Reset()         { *m = FanWeixin{} }
func (m *FanWeixin) String() string { return proto.CompactTextString(m) }
func (*FanWeixin) ProtoMessage()    {}
func (*FanWeixin) Descriptor() ([]byte, []int) {
	return fileDescriptor_750d834c580d7f43, []int{0}
}

func (m *FanWeixin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FanWeixin.Unmarshal(m, b)
}
func (m *FanWeixin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FanWeixin.Marshal(b, m, deterministic)
}
func (m *FanWeixin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FanWeixin.Merge(m, src)
}
func (m *FanWeixin) XXX_Size() int {
	return xxx_messageInfo_FanWeixin.Size(m)
}
func (m *FanWeixin) XXX_DiscardUnknown() {
	xxx_messageInfo_FanWeixin.DiscardUnknown(m)
}

var xxx_messageInfo_FanWeixin proto.InternalMessageInfo

func (m *FanWeixin) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

//
type FanWeixinResponse struct {
	Entity               *FanWeixin   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*FanWeixin `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FanWeixinResponse) Reset()         { *m = FanWeixinResponse{} }
func (m *FanWeixinResponse) String() string { return proto.CompactTextString(m) }
func (*FanWeixinResponse) ProtoMessage()    {}
func (*FanWeixinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_750d834c580d7f43, []int{1}
}

func (m *FanWeixinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FanWeixinResponse.Unmarshal(m, b)
}
func (m *FanWeixinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FanWeixinResponse.Marshal(b, m, deterministic)
}
func (m *FanWeixinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FanWeixinResponse.Merge(m, src)
}
func (m *FanWeixinResponse) XXX_Size() int {
	return xxx_messageInfo_FanWeixinResponse.Size(m)
}
func (m *FanWeixinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FanWeixinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FanWeixinResponse proto.InternalMessageInfo

func (m *FanWeixinResponse) GetEntity() *FanWeixin {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *FanWeixinResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *FanWeixinResponse) GetItems() []*FanWeixin {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *FanWeixinResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *FanWeixinResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*FanWeixin)(nil), "geiqin.srv.crm.FanWeixin")
	proto.RegisterType((*FanWeixinResponse)(nil), "geiqin.srv.crm.FanWeixinResponse")
}

func init() { proto.RegisterFile("fanWeixin.proto", fileDescriptor_750d834c580d7f43) }

var fileDescriptor_750d834c580d7f43 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0xd2, 0x04, 0x1c, 0xa5, 0xea, 0xa2, 0x10, 0xeb, 0xa5, 0xe6, 0x54, 0x10, 0x22,
	0xd6, 0x5f, 0xe0, 0x47, 0xfd, 0xb8, 0x49, 0x8a, 0xf4, 0x1c, 0x93, 0x49, 0x3b, 0x60, 0x76, 0xe3,
	0x64, 0x29, 0xfa, 0xd7, 0xf5, 0x22, 0xd9, 0xc4, 0x8a, 0xad, 0x2d, 0x3d, 0xf4, 0xfc, 0x3e, 0xef,
	0x33, 0x33, 0xcb, 0xc2, 0x5e, 0x16, 0xcb, 0x11, 0xd2, 0x3b, 0xc9, 0xb0, 0x60, 0xa5, 0x95, 0x68,
	0x8f, 0x91, 0xde, 0x48, 0x86, 0x25, 0x4f, 0xc3, 0x84, 0xf3, 0xce, 0x6e, 0xa2, 0xf2, 0x5c, 0x35,
	0x69, 0x70, 0x02, 0xdb, 0x77, 0x3f, 0x05, 0xd1, 0x06, 0x9b, 0x52, 0xdf, 0xea, 0x5a, 0x3d, 0x37,
	0xb2, 0x29, 0x0d, 0xbe, 0x2c, 0x38, 0x98, 0xa5, 0x11, 0x96, 0x85, 0x92, 0x25, 0x8a, 0x0b, 0xf0,
	0x50, 0x6a, 0xd2, 0x1f, 0x86, 0xdc, 0xe9, 0x1f, 0x87, 0x7f, 0x27, 0x84, 0xbf, 0x95, 0x06, 0x14,
	0x67, 0xe0, 0x16, 0xf1, 0x18, 0xd9, 0xb7, 0x4d, 0xe3, 0x68, 0xbe, 0xf1, 0x54, 0x85, 0x51, 0xcd,
	0x88, 0x73, 0x70, 0x49, 0x63, 0x5e, 0xfa, 0x4e, 0xd7, 0x59, 0xad, 0xaf, 0xb9, 0xca, 0x8e, 0xcc,
	0x8a, 0xfd, 0xd6, 0xff, 0xf6, 0x41, 0x15, 0x46, 0x35, 0x23, 0x7a, 0xd0, 0x22, 0x99, 0x29, 0xdf,
	0x35, 0xec, 0xe1, 0x3c, 0xfb, 0x28, 0x33, 0x15, 0x19, 0xa2, 0xff, 0x69, 0xc3, 0xfe, 0x6c, 0xd6,
	0x10, 0x79, 0x4a, 0x09, 0x8a, 0x07, 0xf0, 0x6e, 0x18, 0x63, 0x8d, 0x62, 0xf9, 0x5e, 0x9d, 0xd3,
	0xe5, 0x2b, 0x37, 0x8f, 0x18, 0x6c, 0x55, 0xa6, 0xe7, 0x22, 0xdd, 0x84, 0xe9, 0x0a, 0xbc, 0x5b,
	0x7c, 0x45, 0x8d, 0x42, 0x2c, 0x9c, 0x93, 0xae, 0xa7, 0x18, 0x80, 0x73, 0x8f, 0x7a, 0x13, 0x37,
	0x0d, 0x31, 0xe6, 0x64, 0xb2, 0x68, 0xba, 0x8e, 0x4b, 0x1c, 0x4d, 0x90, 0x71, 0x2d, 0xd3, 0x8b,
	0x67, 0xbe, 0xe7, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x8d, 0x99, 0x1d, 0xcf, 0x02,
	0x00, 0x00,
}