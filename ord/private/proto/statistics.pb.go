// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statistics.proto

package geiqin_srv_ord_private

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

type StatsRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Days                 int32    `protobuf:"varint,2,opt,name=days,proto3" json:"days,omitempty"`
	Paged                int32    `protobuf:"varint,3,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsRequest) Reset()         { *m = StatsRequest{} }
func (m *StatsRequest) String() string { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()    {}
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4ebcdede33dbb6, []int{0}
}

func (m *StatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsRequest.Unmarshal(m, b)
}
func (m *StatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsRequest.Marshal(b, m, deterministic)
}
func (m *StatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsRequest.Merge(m, src)
}
func (m *StatsRequest) XXX_Size() int {
	return xxx_messageInfo_StatsRequest.Size(m)
}
func (m *StatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsRequest proto.InternalMessageInfo

func (m *StatsRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StatsRequest) GetDays() int32 {
	if m != nil {
		return m.Days
	}
	return 0
}

func (m *StatsRequest) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *StatsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func init() {
	proto.RegisterType((*StatsRequest)(nil), "geiqin.srv.ord.private.StatsRequest")
}

func init() {
	proto.RegisterFile("statistics.proto", fileDescriptor_fe4ebcdede33dbb6)
}

var fileDescriptor_fe4ebcdede33dbb6 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xad, 0xa6, 0x62, 0x06, 0x85, 0xb2, 0x88, 0xc4, 0x78, 0x29, 0xc5, 0x43, 0x11, 0x5c,
	0x50, 0x1f, 0x41, 0x41, 0x6f, 0x85, 0xa4, 0x9e, 0x75, 0xdb, 0x1d, 0xeb, 0x42, 0xcd, 0x6e, 0x77,
	0xa6, 0x85, 0xf4, 0x8d, 0x7d, 0x0b, 0xc9, 0xaa, 0x21, 0x8a, 0x16, 0x3d, 0xe4, 0x36, 0xfb, 0xff,
	0x03, 0xdf, 0xfc, 0x3b, 0x03, 0x3d, 0x62, 0xc5, 0x86, 0xd8, 0x4c, 0x49, 0x3a, 0x6f, 0xd9, 0x8a,
	0xa3, 0x19, 0x9a, 0x85, 0x29, 0x24, 0xf9, 0x95, 0xb4, 0x5e, 0x4b, 0xe7, 0xcd, 0x4a, 0x31, 0xa6,
	0x3d, 0xeb, 0x35, 0xfa, 0x9c, 0x15, 0x7f, 0x74, 0xa6, 0xc7, 0x41, 0xb9, 0x2f, 0x9e, 0x55, 0xa1,
	0xe7, 0xa8, 0x1b, 0xd6, 0xc0, 0xc0, 0x7e, 0x78, 0x66, 0xb8, 0x58, 0x22, 0xb1, 0x10, 0x10, 0x71,
	0xe9, 0x30, 0xe9, 0xf4, 0x3b, 0xc3, 0x38, 0x0b, 0x75, 0xa5, 0x69, 0x55, 0x52, 0xb2, 0xdd, 0xef,
	0x0c, 0xbb, 0x59, 0xa8, 0xc5, 0x21, 0x74, 0x9d, 0x9a, 0xa1, 0x4e, 0x76, 0x82, 0xf8, 0xfe, 0x10,
	0x27, 0x10, 0x57, 0xc5, 0x03, 0x99, 0x35, 0x26, 0x51, 0x70, 0xf6, 0x2a, 0x21, 0x37, 0x6b, 0xbc,
	0x7c, 0x8d, 0x00, 0xf2, 0x3a, 0x84, 0x78, 0x04, 0x18, 0x55, 0x63, 0x8d, 0x2d, 0xab, 0xb9, 0x38,
	0x95, 0x3f, 0xa7, 0x91, 0xcd, 0xe9, 0xd2, 0xb3, 0xdf, 0xba, 0x46, 0x75, 0xe4, 0x0c, 0xc9, 0xd9,
	0x82, 0x70, 0xb0, 0x25, 0xa6, 0x00, 0x63, 0xab, 0x55, 0xf9, 0x1f, 0xc2, 0xf9, 0x46, 0xc2, 0x8d,
	0x2a, 0xbf, 0x43, 0x26, 0x10, 0x7f, 0x5a, 0xd4, 0x16, 0xe3, 0x09, 0xe2, 0x7a, 0x79, 0x7f, 0x64,
	0x5c, 0x6c, 0x64, 0xdc, 0x35, 0x0e, 0xe1, 0xcb, 0x87, 0x1d, 0x5c, 0x2f, 0x89, 0xed, 0x0b, 0xfa,
	0x5b, 0x6b, 0x35, 0xb5, 0xb4, 0x95, 0x1a, 0x12, 0xfc, 0x36, 0x20, 0x93, 0xdd, 0x70, 0xdd, 0x57,
	0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39, 0x8b, 0xb0, 0xc9, 0x36, 0x03, 0x00, 0x00,
}
