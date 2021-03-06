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
	CustomerId           int64    `protobuf:"varint,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	StartAt              string   `protobuf:"bytes,6,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                string   `protobuf:"bytes,7,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	StatisticAt          string   `protobuf:"bytes,8,opt,name=statistic_at,json=statisticAt,proto3" json:"statistic_at,omitempty"`
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

func (m *StatsRequest) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *StatsRequest) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *StatsRequest) GetEndAt() string {
	if m != nil {
		return m.EndAt
	}
	return ""
}

func (m *StatsRequest) GetStatisticAt() string {
	if m != nil {
		return m.StatisticAt
	}
	return ""
}

func init() {
	proto.RegisterType((*StatsRequest)(nil), "geiqin.srv.ord.private.StatsRequest")
}

func init() {
	proto.RegisterFile("statistics.proto", fileDescriptor_fe4ebcdede33dbb6)
}

var fileDescriptor_fe4ebcdede33dbb6 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x31, 0xad, 0xf3, 0x67, 0x6a, 0xa4, 0xb2, 0x02, 0xe4, 0x1a, 0x21, 0x42, 0xc5, 0x21,
	0x42, 0xc2, 0x12, 0xe1, 0xc4, 0x31, 0x02, 0x04, 0x1c, 0x10, 0x92, 0x13, 0xce, 0x61, 0xeb, 0x9d,
	0x96, 0x15, 0xf1, 0xae, 0xbb, 0x3b, 0x8d, 0xe4, 0x1e, 0x78, 0x1d, 0x1e, 0x8b, 0x57, 0x41, 0xde,
	0xda, 0xc6, 0x44, 0x4d, 0x65, 0x0b, 0xc1, 0x6d, 0xf7, 0xfb, 0x26, 0xf3, 0xdb, 0xf9, 0x34, 0x31,
	0x1c, 0x5a, 0xe2, 0x24, 0x2d, 0xc9, 0xd4, 0xc6, 0xb9, 0xd1, 0xa4, 0xd9, 0x83, 0x33, 0x94, 0xe7,
	0x52, 0xc5, 0xd6, 0x6c, 0x62, 0x6d, 0x44, 0x9c, 0x1b, 0xb9, 0xe1, 0x84, 0x51, 0x90, 0xea, 0x2c,
	0xd3, 0xea, 0xaa, 0x2a, 0x3a, 0xd4, 0x46, 0xa0, 0x59, 0x10, 0xa7, 0xea, 0x77, 0xd1, 0x91, 0x53,
	0x3e, 0xab, 0xaf, 0x5c, 0x89, 0x35, 0x8a, 0x96, 0x75, 0xfc, 0xd3, 0x83, 0xc0, 0xdd, 0x13, 0x3c,
	0xbf, 0x40, 0x4b, 0x8c, 0xc1, 0x3e, 0x15, 0x39, 0x86, 0xde, 0xc4, 0x9b, 0x8e, 0x13, 0x77, 0x2e,
	0x35, 0xc1, 0x0b, 0x1b, 0xde, 0x9e, 0x78, 0x53, 0x3f, 0x71, 0x67, 0x76, 0x0f, 0xfc, 0x9c, 0x9f,
	0xa1, 0x08, 0xf7, 0x9c, 0x78, 0x75, 0x61, 0x0f, 0x61, 0x5c, 0x1e, 0x56, 0x56, 0x5e, 0x62, 0xb8,
	0xef, 0x9c, 0x51, 0x29, 0x2c, 0xe4, 0x25, 0xb2, 0xc7, 0x70, 0x90, 0x5e, 0x58, 0xd2, 0x19, 0x9a,
	0x95, 0x14, 0xa1, 0x3f, 0xf1, 0xa6, 0x7b, 0x09, 0xd4, 0xd2, 0x07, 0xc1, 0x8e, 0x60, 0x64, 0x89,
	0x1b, 0x5a, 0x71, 0x0a, 0x07, 0x8e, 0x3f, 0x74, 0xf7, 0x39, 0xb1, 0xfb, 0x30, 0x40, 0x25, 0x4a,
	0x63, 0xe8, 0x0c, 0x1f, 0x95, 0x98, 0x13, 0x7b, 0x02, 0x41, 0x93, 0x52, 0x69, 0x8e, 0x9c, 0x79,
	0xd0, 0x68, 0x73, 0x9a, 0x7d, 0x87, 0xe0, 0x63, 0xb1, 0x68, 0xa2, 0x64, 0x0a, 0xc6, 0x4d, 0x12,
	0xec, 0x69, 0x7c, 0x7d, 0xa4, 0x71, 0x3b, 0x93, 0xe8, 0xd5, 0xae, 0xaa, 0xd7, 0xd5, 0xdb, 0x3f,
	0x95, 0xf9, 0xbe, 0x6f, 0xa5, 0x9b, 0xa0, 0xcd, 0xb5, 0xb2, 0x78, 0x7c, 0x6b, 0xf6, 0x63, 0x08,
	0xd0, 0xc2, 0xa7, 0x00, 0xae, 0xfa, 0x0d, 0x97, 0xeb, 0xa2, 0x23, 0xff, 0xf9, 0xae, 0xaa, 0xaa,
	0x53, 0xb1, 0xc5, 0x64, 0x5f, 0x2a, 0xc8, 0x52, 0x13, 0x5f, 0x77, 0x84, 0x3c, 0xbb, 0x11, 0xb2,
	0x4d, 0x48, 0x01, 0x96, 0x5a, 0xf0, 0xa2, 0x0f, 0xa1, 0xf7, 0x18, 0x27, 0x30, 0xae, 0x2d, 0xfb,
	0xaf, 0x18, 0xa7, 0xfd, 0xd7, 0xe1, 0xc5, 0x8d, 0x8c, 0xeb, 0xd7, 0x80, 0x9d, 0xc2, 0x9d, 0x7a,
	0x5b, 0xde, 0x69, 0x2d, 0xfe, 0x7a, 0x9e, 0x3f, 0x9a, 0xb5, 0x38, 0xdf, 0x7e, 0x73, 0xdc, 0x73,
	0x3a, 0x72, 0x66, 0x9d, 0x56, 0x7c, 0x7b, 0xa8, 0x0d, 0xdc, 0xad, 0xfd, 0xff, 0xf9, 0x9f, 0x62,
	0x4b, 0xf0, 0x13, 0xb4, 0x48, 0xec, 0xd1, 0xae, 0x2e, 0x6f, 0xb3, 0x9c, 0x8a, 0x7e, 0x3b, 0x7d,
	0x32, 0x70, 0x9f, 0xc4, 0x97, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x75, 0xfa, 0x21, 0x79,
	0x05, 0x00, 0x00,
}
