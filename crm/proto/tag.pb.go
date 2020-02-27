// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tag.proto

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

type Tag struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 int32    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	ConditionType        int32    `protobuf:"varint,4,opt,name=condition_type,json=conditionType,proto3" json:"condition_type,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f545bcde37ecb5, []int{0}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Tag) GetConditionType() int32 {
	if m != nil {
		return m.ConditionType
	}
	return 0
}

func (m *Tag) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Tag) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

//
type TagResponse struct {
	Entity               *Tag     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Tag   `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagResponse) Reset()         { *m = TagResponse{} }
func (m *TagResponse) String() string { return proto.CompactTextString(m) }
func (*TagResponse) ProtoMessage()    {}
func (*TagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f545bcde37ecb5, []int{1}
}

func (m *TagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagResponse.Unmarshal(m, b)
}
func (m *TagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagResponse.Marshal(b, m, deterministic)
}
func (m *TagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagResponse.Merge(m, src)
}
func (m *TagResponse) XXX_Size() int {
	return xxx_messageInfo_TagResponse.Size(m)
}
func (m *TagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TagResponse proto.InternalMessageInfo

func (m *TagResponse) GetEntity() *Tag {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *TagResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *TagResponse) GetItems() []*Tag {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *TagResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *TagResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Tag)(nil), "geiqin.srv.crm.Tag")
	proto.RegisterType((*TagResponse)(nil), "geiqin.srv.crm.TagResponse")
}

func init() { proto.RegisterFile("tag.proto", fileDescriptor_27f545bcde37ecb5) }

var fileDescriptor_27f545bcde37ecb5 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0x24, 0x4d, 0xa0, 0x37, 0xbf, 0x5d, 0x8c, 0x0a, 0xb1, 0x22, 0x94, 0x82, 0x50,
	0x29, 0x64, 0x11, 0x97, 0x8a, 0xd0, 0xaa, 0x48, 0x77, 0x92, 0x46, 0x5c, 0x96, 0x31, 0xb9, 0x4d,
	0x07, 0xcc, 0x4c, 0x9c, 0x8c, 0x85, 0xbe, 0x8e, 0x6f, 0xe3, 0xa3, 0xf8, 0x16, 0x32, 0x93, 0x28,
	0x58, 0x2b, 0x85, 0xee, 0x86, 0x7b, 0xbe, 0x73, 0xb8, 0xe7, 0x32, 0xd0, 0x56, 0x34, 0x0f, 0x4b,
	0x29, 0x94, 0x20, 0x9d, 0x1c, 0xd9, 0x0b, 0xe3, 0x61, 0x25, 0x97, 0x61, 0x2a, 0x8b, 0xee, 0xff,
	0x54, 0x14, 0x85, 0xe0, 0xb5, 0xda, 0x7f, 0xb3, 0xc0, 0x49, 0x68, 0x4e, 0x3a, 0x60, 0xb3, 0x2c,
	0xb0, 0x7a, 0xd6, 0xc0, 0x8d, 0x6d, 0x96, 0x11, 0x02, 0x2d, 0x4e, 0x0b, 0x0c, 0xec, 0x9e, 0x35,
	0x68, 0xc7, 0xe6, 0xad, 0x67, 0x6a, 0x55, 0x62, 0xe0, 0x18, 0xca, 0xbc, 0xc9, 0x29, 0x74, 0x52,
	0xc1, 0x33, 0xa6, 0x98, 0xe0, 0x33, 0xa3, 0xb6, 0x8c, 0xba, 0xf7, 0x3d, 0x4d, 0x34, 0x76, 0x02,
	0x90, 0x4a, 0xa4, 0x0a, 0xb3, 0x19, 0x55, 0x81, 0x6b, 0x42, 0xdb, 0xcd, 0x64, 0xa4, 0xb4, 0xfc,
	0x5a, 0x66, 0x5f, 0xb2, 0x57, 0xcb, 0xcd, 0x64, 0xa4, 0xfa, 0x1f, 0x16, 0xf8, 0x09, 0xcd, 0x63,
	0xac, 0x4a, 0xc1, 0x2b, 0x24, 0x43, 0xf0, 0x90, 0x2b, 0xa6, 0x56, 0x66, 0x61, 0x3f, 0xda, 0x0f,
	0x7f, 0x76, 0x0c, 0x35, 0xdc, 0x20, 0x64, 0x08, 0x6e, 0x49, 0x73, 0x94, 0xa6, 0x8a, 0x1f, 0x1d,
	0xae, 0xb3, 0xf7, 0x5a, 0x8c, 0x6b, 0x86, 0x9c, 0x81, 0xcb, 0x14, 0x16, 0x55, 0xe0, 0xf4, 0x9c,
	0xbf, 0x82, 0x6b, 0x42, 0xe7, 0xa2, 0x94, 0x42, 0x9a, 0xc2, 0x1b, 0x72, 0x6f, 0xb5, 0x18, 0xd7,
	0x0c, 0x19, 0x40, 0x8b, 0xf1, 0xb9, 0x30, 0xcd, 0xfd, 0xe8, 0x60, 0x9d, 0x9d, 0xf0, 0xb9, 0x88,
	0x0d, 0x11, 0xbd, 0xdb, 0x00, 0x09, 0xcd, 0xa7, 0x28, 0x97, 0x2c, 0x45, 0x72, 0x09, 0xde, 0xb5,
	0x39, 0x13, 0xd9, 0xb4, 0x4b, 0xf7, 0x78, 0xd3, 0x82, 0xcd, 0x99, 0xfa, 0xff, 0xb4, 0xfb, 0xc1,
	0x5c, 0x71, 0x27, 0xf7, 0x15, 0x78, 0x37, 0xf8, 0x8c, 0x0a, 0xc9, 0xaf, 0x72, 0x93, 0x6c, 0xc2,
	0xd5, 0x36, 0xff, 0x05, 0x38, 0x77, 0xa8, 0x76, 0x34, 0x8f, 0xc1, 0x9b, 0x22, 0x95, 0xe9, 0x82,
	0x1c, 0xad, 0x83, 0x63, 0x5a, 0xe1, 0xe3, 0x02, 0x25, 0x6e, 0xc9, 0x78, 0xf2, 0xcc, 0x1f, 0x3f,
	0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xed, 0x17, 0x91, 0x0e, 0x03, 0x00, 0x00,
}
