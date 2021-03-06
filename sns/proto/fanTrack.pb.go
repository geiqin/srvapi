// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fanTrack.proto

package geiqin_srv_sns

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

type FanTrack struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	CreatedAt            string   `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ids                  []int64  `protobuf:"varint,20,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FanTrack) Reset()         { *m = FanTrack{} }
func (m *FanTrack) String() string { return proto.CompactTextString(m) }
func (*FanTrack) ProtoMessage()    {}
func (*FanTrack) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ebdb3459b3bd922, []int{0}
}

func (m *FanTrack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FanTrack.Unmarshal(m, b)
}
func (m *FanTrack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FanTrack.Marshal(b, m, deterministic)
}
func (m *FanTrack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FanTrack.Merge(m, src)
}
func (m *FanTrack) XXX_Size() int {
	return xxx_messageInfo_FanTrack.Size(m)
}
func (m *FanTrack) XXX_DiscardUnknown() {
	xxx_messageInfo_FanTrack.DiscardUnknown(m)
}

var xxx_messageInfo_FanTrack proto.InternalMessageInfo

func (m *FanTrack) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FanTrack) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *FanTrack) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *FanTrack) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *FanTrack) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type FanTrackResponse struct {
	Entity               *FanTrack   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*FanTrack `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FanTrackResponse) Reset()         { *m = FanTrackResponse{} }
func (m *FanTrackResponse) String() string { return proto.CompactTextString(m) }
func (*FanTrackResponse) ProtoMessage()    {}
func (*FanTrackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ebdb3459b3bd922, []int{1}
}

func (m *FanTrackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FanTrackResponse.Unmarshal(m, b)
}
func (m *FanTrackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FanTrackResponse.Marshal(b, m, deterministic)
}
func (m *FanTrackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FanTrackResponse.Merge(m, src)
}
func (m *FanTrackResponse) XXX_Size() int {
	return xxx_messageInfo_FanTrackResponse.Size(m)
}
func (m *FanTrackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FanTrackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FanTrackResponse proto.InternalMessageInfo

func (m *FanTrackResponse) GetEntity() *FanTrack {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *FanTrackResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *FanTrackResponse) GetItems() []*FanTrack {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *FanTrackResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *FanTrackResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*FanTrack)(nil), "geiqin.srv.sns.FanTrack")
	proto.RegisterType((*FanTrackResponse)(nil), "geiqin.srv.sns.FanTrackResponse")
}

func init() {
	proto.RegisterFile("fanTrack.proto", fileDescriptor_2ebdb3459b3bd922)
}

var fileDescriptor_2ebdb3459b3bd922 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0xb3, 0x16, 0xf7, 0x26, 0x65, 0x3e, 0xa6, 0x04, 0x41, 0x28, 0x3b, 0x15, 0x84,
	0x20, 0xf3, 0x13, 0xec, 0xa0, 0xe0, 0x4d, 0x82, 0x77, 0xa9, 0xed, 0xeb, 0x08, 0xd2, 0xa4, 0x26,
	0x99, 0xe0, 0xd5, 0x4f, 0xed, 0x51, 0x9a, 0x66, 0x82, 0x43, 0xbc, 0x25, 0xef, 0xf7, 0x7b, 0x7f,
	0xfe, 0x21, 0x50, 0x74, 0xb5, 0x7e, 0xb2, 0x75, 0xf3, 0x2a, 0x06, 0x6b, 0xbc, 0xc1, 0x62, 0x47,
	0xea, 0x4d, 0x69, 0xe1, 0xec, 0xbb, 0x70, 0xda, 0x5d, 0x9e, 0x36, 0xa6, 0xef, 0x8d, 0x9e, 0xe8,
	0xfa, 0x33, 0x81, 0x93, 0xfb, 0xb8, 0x80, 0x05, 0xa4, 0xaa, 0xe5, 0x49, 0x99, 0x54, 0x4c, 0xa6,
	0xaa, 0xc5, 0x0b, 0xc8, 0x9d, 0xd9, 0xdb, 0x86, 0x78, 0x5a, 0x26, 0xd5, 0x5c, 0xc6, 0x1b, 0x5e,
	0x01, 0x34, 0x96, 0x6a, 0x4f, 0xed, 0x73, 0xed, 0xf9, 0x59, 0x60, 0xf3, 0x38, 0xd9, 0xfa, 0x11,
	0xef, 0x87, 0xf6, 0x80, 0x71, 0xc2, 0x71, 0xb2, 0xf5, 0xb8, 0x04, 0xa6, 0x5a, 0xc7, 0x57, 0x25,
	0xab, 0x98, 0x1c, 0x8f, 0xeb, 0xaf, 0x04, 0x96, 0x87, 0x12, 0x92, 0xdc, 0x60, 0xb4, 0x23, 0xbc,
	0x81, 0x9c, 0xb4, 0x57, 0xfe, 0x23, 0x14, 0x5a, 0x6c, 0xb8, 0xf8, 0xfd, 0x10, 0xf1, 0xb3, 0x11,
	0x3d, 0xbc, 0x86, 0x6c, 0xa8, 0x77, 0x64, 0x43, 0xdb, 0xc5, 0xe6, 0xfc, 0x78, 0xe1, 0x71, 0x84,
	0x72, 0x72, 0x50, 0x40, 0xa6, 0x3c, 0xf5, 0x8e, 0xb3, 0x92, 0xfd, 0x9b, 0x3e, 0x69, 0x63, 0x38,
	0x59, 0x6b, 0x2c, 0x9f, 0xfd, 0x1d, 0x7e, 0x37, 0x42, 0x39, 0x39, 0x58, 0xc1, 0x4c, 0xe9, 0xce,
	0xf0, 0x2c, 0xb8, 0xab, 0x63, 0xf7, 0x41, 0x77, 0x46, 0x06, 0xe3, 0x25, 0x0f, 0xdf, 0x70, 0xfb,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x51, 0x08, 0xfd, 0xb6, 0x01, 0x00, 0x00,
}
