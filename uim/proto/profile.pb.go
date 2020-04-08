// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile.proto

package geiqin_srv_uim

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

// 用户信息
type Profile struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName          string   `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{0}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Profile) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *Profile) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Profile) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ProfileResponse struct {
	Entity               *Profile   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Profile `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProfileResponse) Reset()         { *m = ProfileResponse{} }
func (m *ProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ProfileResponse) ProtoMessage()    {}
func (*ProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{1}
}

func (m *ProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileResponse.Unmarshal(m, b)
}
func (m *ProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileResponse.Marshal(b, m, deterministic)
}
func (m *ProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileResponse.Merge(m, src)
}
func (m *ProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ProfileResponse.Size(m)
}
func (m *ProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileResponse proto.InternalMessageInfo

func (m *ProfileResponse) GetEntity() *Profile {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ProfileResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ProfileResponse) GetItems() []*Profile {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ProfileResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ProfileResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "geiqin.srv.uim.Profile")
	proto.RegisterType((*ProfileResponse)(nil), "geiqin.srv.uim.ProfileResponse")
}

func init() {
	proto.RegisterFile("profile.proto", fileDescriptor_744bf7a47b381504)
}

var fileDescriptor_744bf7a47b381504 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0x02, 0x41,
	0x18, 0xc6, 0x5b, 0x57, 0x0d, 0x5f, 0xcb, 0xe0, 0xa5, 0x68, 0x13, 0x22, 0xf3, 0x24, 0x44, 0x1b,
	0xd8, 0x27, 0x30, 0xa8, 0xe8, 0x12, 0xb1, 0x1e, 0x3a, 0xca, 0xb4, 0xfb, 0xaa, 0x03, 0xcd, 0x9f,
	0x66, 0x47, 0xc1, 0xbe, 0x4a, 0x9f, 0xb3, 0x7b, 0xec, 0xcc, 0x78, 0xc8, 0x32, 0xba, 0x3e, 0xcf,
	0x6f, 0x9f, 0xf7, 0xc7, 0x32, 0xb0, 0xaf, 0x8d, 0x9a, 0xf2, 0x57, 0x4a, 0xb5, 0x51, 0x56, 0x61,
	0x67, 0x46, 0xfc, 0x8d, 0xcb, 0xb4, 0x34, 0xcb, 0x74, 0xc1, 0x45, 0x77, 0x2f, 0x57, 0x42, 0x28,
	0xe9, 0xdb, 0xfe, 0x3b, 0xec, 0x3e, 0x79, 0x1c, 0x3b, 0x50, 0xe3, 0x45, 0x12, 0xf5, 0xa2, 0x41,
	0x9c, 0xd5, 0x78, 0x81, 0xe7, 0x50, 0xa1, 0x9a, 0xc9, 0xd5, 0x44, 0x32, 0x41, 0x49, 0xad, 0x17,
	0x0d, 0x5a, 0x59, 0x3b, 0x64, 0x8f, 0x4c, 0x10, 0x9e, 0x02, 0xe4, 0x86, 0x98, 0xa5, 0x62, 0xc2,
	0x6c, 0x12, 0x3b, 0xa0, 0x15, 0x92, 0x91, 0xad, 0xea, 0x85, 0x2e, 0xd6, 0x75, 0xdd, 0xd7, 0x21,
	0x19, 0xd9, 0xfe, 0x67, 0x04, 0x07, 0xe1, 0x78, 0x46, 0xa5, 0x56, 0xb2, 0x24, 0xbc, 0x82, 0x26,
	0x49, 0xcb, 0xed, 0xca, 0x89, 0xb4, 0x87, 0xc7, 0xe9, 0x77, 0xfd, 0x74, 0xfd, 0x41, 0xc0, 0xf0,
	0x02, 0x1a, 0x9a, 0xcd, 0xc8, 0x38, 0xbd, 0xf6, 0xf0, 0xe8, 0x07, 0x5f, 0x95, 0x99, 0x67, 0xf0,
	0x12, 0x1a, 0xdc, 0x92, 0x28, 0x93, 0xb8, 0x17, 0xff, 0x35, 0xee, 0xa9, 0x6a, 0x9b, 0x8c, 0x51,
	0xc6, 0xa9, 0xff, 0xb2, 0x7d, 0x5b, 0x95, 0x99, 0x67, 0x70, 0x00, 0x75, 0x2e, 0xa7, 0x2a, 0x69,
	0x38, 0xf6, 0x70, 0x93, 0x7d, 0x90, 0x53, 0x95, 0x39, 0x62, 0xf8, 0x11, 0x41, 0x27, 0x5c, 0x1a,
	0x93, 0x59, 0xf2, 0x9c, 0x70, 0x04, 0xf1, 0x3d, 0x59, 0xdc, 0x26, 0xd4, 0x3d, 0xdb, 0x66, 0x1a,
	0xfe, 0x5b, 0x7f, 0x07, 0xef, 0xa0, 0x39, 0x26, 0x66, 0xf2, 0x39, 0x9e, 0x6c, 0xc2, 0x37, 0xac,
	0xa4, 0xe7, 0x39, 0x99, 0xff, 0xec, 0xbc, 0x34, 0xdd, 0xc3, 0xb8, 0xfe, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xc6, 0x79, 0x23, 0x9d, 0x47, 0x02, 0x00, 0x00,
}
