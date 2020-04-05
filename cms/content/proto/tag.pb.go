// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tag.proto

package geiqin_srv_cms_content

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

// 字典信息
type Tag struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Slug                 string   `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug"`
	ParentId             int32    `protobuf:"varint,4,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	Memo                 string   `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	ArticleCount         int32    `protobuf:"varint,6,opt,name=article_count,json=articleCount,proto3" json:"article_count"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Children             []*Tag   `protobuf:"bytes,9,rep,name=children,proto3" json:"children,omitempty"`
	Ids                  []int32  `protobuf:"varint,10,rep,packed,name=ids,proto3" json:"ids,omitempty"`
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

func (m *Tag) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Tag) GetParentId() int32 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Tag) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Tag) GetArticleCount() int32 {
	if m != nil {
		return m.ArticleCount
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

func (m *Tag) GetChildren() []*Tag {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Tag) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
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
	proto.RegisterType((*Tag)(nil), "geiqin.srv.cms.content.Tag")
	proto.RegisterType((*TagResponse)(nil), "geiqin.srv.cms.content.TagResponse")
}

func init() {
	proto.RegisterFile("tag.proto", fileDescriptor_27f545bcde37ecb5)
}

var fileDescriptor_27f545bcde37ecb5 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x69, 0xd2, 0x84, 0x66, 0xba, 0x20, 0x34, 0x07, 0x64, 0xed, 0xb2, 0x52, 0xe9, 0x5e,
	0x7a, 0x8a, 0xa0, 0x3d, 0x70, 0x5e, 0x16, 0x84, 0x2a, 0xed, 0x61, 0x95, 0x2d, 0xe2, 0x58, 0x19,
	0x67, 0x36, 0xb5, 0xd4, 0xd8, 0xc1, 0x76, 0x57, 0xe2, 0x31, 0xb8, 0xf3, 0x7e, 0xbc, 0x06, 0xb2,
	0x1d, 0x38, 0xd1, 0x96, 0x43, 0x6f, 0xa3, 0x7f, 0xbe, 0xff, 0x77, 0x66, 0x46, 0x0a, 0x14, 0x8e,
	0x37, 0x65, 0x67, 0xb4, 0xd3, 0xf8, 0xb2, 0x21, 0xf9, 0x4d, 0xaa, 0xd2, 0x9a, 0xc7, 0x52, 0xb4,
	0xb6, 0x14, 0x5a, 0x39, 0x52, 0xee, 0xfc, 0x4c, 0xe8, 0xb6, 0xd5, 0x2a, 0x52, 0xd3, 0x9f, 0x09,
	0xa4, 0x2b, 0xde, 0xe0, 0x73, 0x48, 0x64, 0xcd, 0x06, 0x93, 0xc1, 0x2c, 0xab, 0x12, 0x59, 0x23,
	0xc2, 0x50, 0xf1, 0x96, 0x58, 0x32, 0x19, 0xcc, 0x8a, 0x2a, 0xd4, 0x5e, 0xb3, 0xdb, 0x5d, 0xc3,
	0xd2, 0xa8, 0xf9, 0x1a, 0x2f, 0xa0, 0xe8, 0xb8, 0x21, 0xe5, 0xd6, 0xb2, 0x66, 0xc3, 0x60, 0x1f,
	0x45, 0x61, 0x19, 0x42, 0x5a, 0x6a, 0x35, 0xcb, 0xa2, 0xc1, 0xd7, 0x78, 0x05, 0xcf, 0xb8, 0x71,
	0x52, 0x6c, 0x69, 0x2d, 0xf4, 0x4e, 0x39, 0x96, 0x07, 0xd3, 0x59, 0x2f, 0xde, 0x78, 0x0d, 0x2f,
	0x01, 0x84, 0x21, 0xee, 0xa8, 0x5e, 0x73, 0xc7, 0x9e, 0x06, 0x7b, 0xd1, 0x2b, 0xd7, 0xa1, 0xbd,
	0xeb, 0xea, 0x3f, 0xed, 0x51, 0x6c, 0xf7, 0xca, 0xb5, 0xc3, 0x77, 0x30, 0x12, 0x1b, 0xb9, 0xad,
	0x0d, 0x29, 0x56, 0x4c, 0xd2, 0xd9, 0x78, 0x7e, 0x51, 0xfe, 0x7b, 0x19, 0xe5, 0x8a, 0x37, 0xd5,
	0x5f, 0x18, 0x5f, 0x40, 0x2a, 0x6b, 0xcb, 0x60, 0x92, 0xce, 0xb2, 0xca, 0x97, 0xd3, 0x1f, 0x09,
	0x8c, 0x3d, 0x43, 0xb6, 0xd3, 0xca, 0x12, 0x2e, 0x20, 0x27, 0xe5, 0xa4, 0xfb, 0x1e, 0x56, 0x75,
	0x24, 0xb8, 0x47, 0x71, 0x01, 0x59, 0xc7, 0x1b, 0x32, 0x61, 0x99, 0xe3, 0xf9, 0xe5, 0x3e, 0xcf,
	0x9d, 0x87, 0xaa, 0xc8, 0xe2, 0x5b, 0xc8, 0xa4, 0xa3, 0xd6, 0xb2, 0xf4, 0xf8, 0x04, 0x91, 0xf4,
	0xef, 0x90, 0x31, 0xda, 0x84, 0x3b, 0x1c, 0x78, 0xe7, 0xa3, 0x87, 0xaa, 0xc8, 0xe2, 0x1b, 0x18,
	0x4a, 0xf5, 0x10, 0x6f, 0x34, 0x9e, 0xbf, 0xda, 0xe7, 0x59, 0xaa, 0x07, 0x5d, 0x05, 0x72, 0xfe,
	0x2b, 0x05, 0x58, 0xf1, 0xe6, 0x9e, 0xcc, 0xa3, 0x14, 0x84, 0xb7, 0x90, 0xdf, 0x84, 0xcb, 0xe0,
	0xa1, 0x6f, 0x3c, 0xbf, 0x3a, 0x34, 0x40, 0xbf, 0xde, 0xe9, 0x13, 0x9f, 0xf6, 0x39, 0x1c, 0xf2,
	0x54, 0x69, 0x1f, 0x68, 0x4b, 0x27, 0x4a, 0x5b, 0x42, 0xfa, 0x89, 0xdc, 0x49, 0xa2, 0x2a, 0xc8,
	0xef, 0x89, 0x1b, 0xb1, 0xc1, 0xd7, 0xfb, 0x0c, 0xef, 0xb9, 0xa5, 0x2f, 0x1b, 0x32, 0xf4, 0xbf,
	0x99, 0x77, 0x30, 0xbc, 0x95, 0xd6, 0x9d, 0x2e, 0xf1, 0x6b, 0x1e, 0xfe, 0x11, 0x8b, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0xd2, 0x9f, 0xae, 0x56, 0x04, 0x00, 0x00,
}
