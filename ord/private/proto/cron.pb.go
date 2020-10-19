// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cron.proto

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

type CronServiceResponse struct {
	Pager                *Pager   `protobuf:"bytes,1,opt,name=pager,proto3" json:"pager,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CronServiceResponse) Reset()         { *m = CronServiceResponse{} }
func (m *CronServiceResponse) String() string { return proto.CompactTextString(m) }
func (*CronServiceResponse) ProtoMessage()    {}
func (*CronServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21379dc2cd81a5b8, []int{0}
}

func (m *CronServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CronServiceResponse.Unmarshal(m, b)
}
func (m *CronServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CronServiceResponse.Marshal(b, m, deterministic)
}
func (m *CronServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronServiceResponse.Merge(m, src)
}
func (m *CronServiceResponse) XXX_Size() int {
	return xxx_messageInfo_CronServiceResponse.Size(m)
}
func (m *CronServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CronServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CronServiceResponse proto.InternalMessageInfo

func (m *CronServiceResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *CronServiceResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CronServiceResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*CronServiceResponse)(nil), "geiqin.srv.ord.private.CronServiceResponse")
}

func init() {
	proto.RegisterFile("cron.proto", fileDescriptor_21379dc2cd81a5b8)
}

var fileDescriptor_21379dc2cd81a5b8 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xbd, 0x4a, 0xc0, 0x30,
	0x10, 0x80, 0x8d, 0x7f, 0xc3, 0xb5, 0x53, 0x04, 0x29, 0x45, 0x41, 0x3a, 0x09, 0x42, 0x90, 0xf6,
	0x11, 0x8a, 0x83, 0x93, 0x12, 0x37, 0x71, 0xa9, 0xe9, 0xb5, 0x04, 0x6c, 0x2e, 0x5e, 0x42, 0xc1,
	0xc7, 0xf2, 0x29, 0x7c, 0x2d, 0x69, 0x8a, 0xe0, 0xd0, 0x3a, 0x39, 0xe7, 0xfb, 0xbe, 0x1c, 0x77,
	0x00, 0x86, 0xc9, 0x29, 0xcf, 0x14, 0x49, 0x9e, 0x8f, 0x68, 0xdf, 0xad, 0x53, 0x81, 0x67, 0x45,
	0xdc, 0x2b, 0xcf, 0x76, 0xee, 0x22, 0x96, 0xb9, 0xa1, 0x69, 0xfa, 0xa1, 0xaa, 0x4f, 0x01, 0x67,
	0x2d, 0x93, 0x7b, 0x42, 0x9e, 0xad, 0x41, 0x8d, 0xc1, 0x93, 0x0b, 0x28, 0x1b, 0x38, 0xf1, 0xdd,
	0x88, 0x5c, 0x88, 0x2b, 0x71, 0x9d, 0xd5, 0x97, 0x6a, 0xbb, 0xa6, 0x1e, 0x17, 0x48, 0xaf, 0xec,
	0x22, 0x21, 0x33, 0x71, 0x71, 0xf8, 0xb7, 0x74, 0xb7, 0x40, 0x7a, 0x65, 0xe5, 0x2d, 0x1c, 0x5b,
	0x37, 0x50, 0x71, 0x94, 0x9c, 0x8b, 0x3d, 0xe7, 0xde, 0x0d, 0xa4, 0x13, 0x59, 0x7f, 0x09, 0xc8,
	0x7e, 0xcd, 0x2c, 0x9f, 0x01, 0x1e, 0xb8, 0x47, 0x6e, 0xdf, 0x28, 0xa0, 0xdc, 0xff, 0x75, 0xf2,
	0xf1, 0xa3, 0xbc, 0xd9, 0x7b, 0xde, 0xd8, 0x42, 0x75, 0x20, 0x5f, 0x20, 0x4f, 0x6d, 0x8d, 0x06,
	0xad, 0x8f, 0xff, 0x5b, 0x7f, 0x3d, 0x4d, 0x47, 0x68, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x0a, 0x03, 0xcf, 0xb8, 0x01, 0x00, 0x00,
}