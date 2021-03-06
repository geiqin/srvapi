// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fetchGallery.proto

package geiqin_srv_tms

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

type FetchGallery struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FetchId              int64    `protobuf:"varint,2,opt,name=fetch_id,json=fetchId,proto3" json:"fetch_id,omitempty"`
	MediaId              int64    `protobuf:"varint,3,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	MediaUrl             string   `protobuf:"bytes,4,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	Defaulted            bool     `protobuf:"varint,5,opt,name=defaulted,proto3" json:"defaulted,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ids                  []int64  `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchGallery) Reset()         { *m = FetchGallery{} }
func (m *FetchGallery) String() string { return proto.CompactTextString(m) }
func (*FetchGallery) ProtoMessage()    {}
func (*FetchGallery) Descriptor() ([]byte, []int) {
	return fileDescriptor_866a1ec15d712e72, []int{0}
}

func (m *FetchGallery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchGallery.Unmarshal(m, b)
}
func (m *FetchGallery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchGallery.Marshal(b, m, deterministic)
}
func (m *FetchGallery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchGallery.Merge(m, src)
}
func (m *FetchGallery) XXX_Size() int {
	return xxx_messageInfo_FetchGallery.Size(m)
}
func (m *FetchGallery) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchGallery.DiscardUnknown(m)
}

var xxx_messageInfo_FetchGallery proto.InternalMessageInfo

func (m *FetchGallery) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FetchGallery) GetFetchId() int64 {
	if m != nil {
		return m.FetchId
	}
	return 0
}

func (m *FetchGallery) GetMediaId() int64 {
	if m != nil {
		return m.MediaId
	}
	return 0
}

func (m *FetchGallery) GetMediaUrl() string {
	if m != nil {
		return m.MediaUrl
	}
	return ""
}

func (m *FetchGallery) GetDefaulted() bool {
	if m != nil {
		return m.Defaulted
	}
	return false
}

func (m *FetchGallery) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *FetchGallery) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *FetchGallery) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*FetchGallery)(nil), "geiqin.srv.tms.FetchGallery")
}

func init() {
	proto.RegisterFile("fetchGallery.proto", fileDescriptor_866a1ec15d712e72)
}

var fileDescriptor_866a1ec15d712e72 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xc9, 0x46, 0xdb, 0xdd, 0x41, 0x8a, 0xcc, 0x29, 0xa2, 0x42, 0xf0, 0x94, 0xd3, 0x5e,
	0x7c, 0x82, 0x5e, 0x94, 0x5e, 0x03, 0x9e, 0x4b, 0xec, 0x4c, 0x35, 0x90, 0xda, 0x9a, 0x9d, 0x15,
	0x7c, 0x5a, 0x5f, 0x45, 0x36, 0x59, 0xb1, 0xb7, 0xf9, 0xbf, 0xef, 0x9f, 0x81, 0x01, 0xdc, 0xb3,
	0xec, 0xde, 0x9f, 0x43, 0x4a, 0x9c, 0xbf, 0xfb, 0x53, 0x3e, 0xca, 0x11, 0x57, 0x6f, 0x1c, 0x3f,
	0xe3, 0x47, 0x3f, 0xe4, 0xaf, 0x5e, 0x0e, 0xc3, 0xc3, 0x8f, 0x82, 0xab, 0xa7, 0xb3, 0x1a, 0xae,
	0xa0, 0x89, 0x64, 0x94, 0x55, 0x4e, 0xfb, 0x26, 0x12, 0xde, 0x40, 0x5b, 0xce, 0x6c, 0x23, 0x99,
	0xa6, 0xd0, 0x65, 0xc9, 0x9b, 0xa2, 0x0e, 0x4c, 0x31, 0x4c, 0x4a, 0x57, 0x55, 0xf2, 0x86, 0xf0,
	0x16, 0xba, 0xaa, 0xc6, 0x9c, 0xcc, 0x85, 0x55, 0xae, 0xf3, 0xb5, 0xfb, 0x92, 0x13, 0xde, 0x41,
	0x47, 0xbc, 0x0f, 0x63, 0x12, 0x26, 0x73, 0x69, 0x95, 0x6b, 0xfd, 0x3f, 0xc0, 0x7b, 0x80, 0x5d,
	0xe6, 0x20, 0x4c, 0xdb, 0x20, 0x66, 0x51, 0x76, 0xbb, 0x99, 0xac, 0x65, 0xd2, 0xe3, 0x89, 0xfe,
	0xf4, 0xb2, 0xea, 0x99, 0xac, 0x05, 0xaf, 0x41, 0x47, 0x1a, 0x4c, 0x6b, 0xb5, 0xd3, 0x7e, 0x1a,
	0x5f, 0x17, 0xe5, 0xf1, 0xc7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0xe7, 0x32, 0x98, 0x0e,
	0x01, 0x00, 0x00,
}
