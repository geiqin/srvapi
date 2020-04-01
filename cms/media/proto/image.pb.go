// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image.proto

package geiqin_srv_cms_media

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

// 图片信息
type Image struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	CatId                int32    `protobuf:"varint,4,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"`
	FileName             string   `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	RawName              string   `protobuf:"bytes,6,opt,name=raw_name,json=rawName,proto3" json:"raw_name,omitempty"`
	Hash                 string   `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty"`
	Path                 string   `protobuf:"bytes,8,opt,name=path,proto3" json:"path,omitempty"`
	Url                  string   `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`
	MimeType             string   `protobuf:"bytes,10,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Size                 float32  `protobuf:"fixed32,11,opt,name=size,proto3" json:"size,omitempty"`
	Width                int32    `protobuf:"varint,12,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,13,opt,name=height,proto3" json:"height,omitempty"`
	Memo                 string   `protobuf:"bytes,14,opt,name=memo,proto3" json:"memo,omitempty"`
	Taxonomy             string   `protobuf:"bytes,15,opt,name=taxonomy,proto3" json:"taxonomy,omitempty"`
	UserId               int64    `protobuf:"varint,19,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Cat                  *Cat     `protobuf:"bytes,22,opt,name=cat,proto3" json:"cat,omitempty"`
	Ids                  []int64  `protobuf:"varint,23,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{0}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Image) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Image) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Image) GetCatId() int32 {
	if m != nil {
		return m.CatId
	}
	return 0
}

func (m *Image) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Image) GetRawName() string {
	if m != nil {
		return m.RawName
	}
	return ""
}

func (m *Image) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Image) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Image) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *Image) GetSize() float32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Image) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Image) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Image) GetTaxonomy() string {
	if m != nil {
		return m.Taxonomy
	}
	return ""
}

func (m *Image) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Image) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Image) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Image) GetCat() *Cat {
	if m != nil {
		return m.Cat
	}
	return nil
}

func (m *Image) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

// 修改图片信息
type UpdateImage struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CatId                int32    `protobuf:"varint,4,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"`
	Memo                 string   `protobuf:"bytes,14,opt,name=memo,proto3" json:"memo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateImage) Reset()         { *m = UpdateImage{} }
func (m *UpdateImage) String() string { return proto.CompactTextString(m) }
func (*UpdateImage) ProtoMessage()    {}
func (*UpdateImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{1}
}

func (m *UpdateImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateImage.Unmarshal(m, b)
}
func (m *UpdateImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateImage.Marshal(b, m, deterministic)
}
func (m *UpdateImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateImage.Merge(m, src)
}
func (m *UpdateImage) XXX_Size() int {
	return xxx_messageInfo_UpdateImage.Size(m)
}
func (m *UpdateImage) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateImage.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateImage proto.InternalMessageInfo

func (m *UpdateImage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateImage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateImage) GetCatId() int32 {
	if m != nil {
		return m.CatId
	}
	return 0
}

func (m *UpdateImage) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

//
type ImageResponse struct {
	Entity               *Image   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Image `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageResponse) Reset()         { *m = ImageResponse{} }
func (m *ImageResponse) String() string { return proto.CompactTextString(m) }
func (*ImageResponse) ProtoMessage()    {}
func (*ImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9624c68e2b547544, []int{2}
}

func (m *ImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageResponse.Unmarshal(m, b)
}
func (m *ImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageResponse.Marshal(b, m, deterministic)
}
func (m *ImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageResponse.Merge(m, src)
}
func (m *ImageResponse) XXX_Size() int {
	return xxx_messageInfo_ImageResponse.Size(m)
}
func (m *ImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImageResponse proto.InternalMessageInfo

func (m *ImageResponse) GetEntity() *Image {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ImageResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ImageResponse) GetItems() []*Image {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ImageResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ImageResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Image)(nil), "geiqin.srv.cms.media.Image")
	proto.RegisterType((*UpdateImage)(nil), "geiqin.srv.cms.media.UpdateImage")
	proto.RegisterType((*ImageResponse)(nil), "geiqin.srv.cms.media.ImageResponse")
}

func init() { proto.RegisterFile("image.proto", fileDescriptor_9624c68e2b547544) }

var fileDescriptor_9624c68e2b547544 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x1c, 0xff, 0x34, 0x19, 0xb7, 0xfd, 0xd0, 0xd0, 0x9f, 0x6d, 0x2a, 0x84, 0x09, 0x37,
	0x96, 0x90, 0x2c, 0x91, 0x3e, 0x41, 0x29, 0x08, 0x05, 0x55, 0xa8, 0x72, 0x41, 0xdc, 0x11, 0x2d,
	0xf6, 0x34, 0x5e, 0x29, 0x6b, 0x9b, 0xf5, 0xb6, 0x25, 0xbc, 0x03, 0xcf, 0xc9, 0x25, 0xaf, 0x80,
	0x76, 0x37, 0x95, 0xb8, 0x48, 0x0a, 0x95, 0x72, 0x37, 0x73, 0xe6, 0x9c, 0x39, 0xb3, 0x33, 0x96,
	0x21, 0x16, 0x92, 0xcf, 0x28, 0x6b, 0x55, 0xa3, 0x1b, 0xdc, 0x9b, 0x91, 0xf8, 0x2a, 0xea, 0xac,
	0x53, 0x37, 0x59, 0x21, 0xbb, 0x4c, 0x52, 0x29, 0xf8, 0x70, 0xbb, 0x68, 0xa4, 0x6c, 0x6a, 0xc7,
	0x19, 0x0e, 0x0a, 0xae, 0x5d, 0x38, 0xfa, 0xe5, 0x43, 0x38, 0x31, 0x72, 0xdc, 0x85, 0x9e, 0x28,
	0x99, 0x97, 0x78, 0xa9, 0x9f, 0xf7, 0x44, 0x89, 0x7b, 0x10, 0x6a, 0xa1, 0xe7, 0xc4, 0x7a, 0x89,
	0x97, 0x0e, 0x72, 0x97, 0x20, 0x42, 0xa0, 0x17, 0x2d, 0x31, 0xdf, 0x82, 0x36, 0xc6, 0x7d, 0x88,
	0x0a, 0xae, 0xa7, 0xa2, 0x64, 0x41, 0xe2, 0xa5, 0x61, 0x1e, 0x16, 0x5c, 0x4f, 0x4a, 0x3c, 0x86,
	0xc1, 0x95, 0x98, 0xd3, 0xb4, 0xe6, 0x92, 0x58, 0x68, 0xf9, 0x7d, 0x03, 0xbc, 0xe7, 0x92, 0xf0,
	0x08, 0xfa, 0x8a, 0xdf, 0xba, 0x5a, 0x64, 0x6b, 0x5b, 0x8a, 0xdf, 0xda, 0x12, 0x42, 0x50, 0xf1,
	0xae, 0x62, 0x5b, 0xce, 0xc2, 0xc4, 0x06, 0x6b, 0xb9, 0xae, 0x58, 0xdf, 0x61, 0x26, 0xc6, 0x47,
	0xe0, 0x5f, 0xab, 0x39, 0x1b, 0x58, 0xc8, 0x84, 0xc6, 0x51, 0x0a, 0x49, 0x53, 0x3b, 0x21, 0x38,
	0x47, 0x03, 0x7c, 0x30, 0x53, 0x22, 0x04, 0x9d, 0xf8, 0x4e, 0x2c, 0x4e, 0xbc, 0xb4, 0x97, 0xdb,
	0xd8, 0xbc, 0xf1, 0x56, 0x94, 0xba, 0x62, 0xdb, 0x6e, 0x70, 0x9b, 0xe0, 0x01, 0x44, 0x15, 0x89,
	0x59, 0xa5, 0xd9, 0x8e, 0x85, 0x97, 0x99, 0xe9, 0x20, 0x49, 0x36, 0x6c, 0xd7, 0x0d, 0x61, 0x62,
	0x1c, 0x42, 0x5f, 0xf3, 0x6f, 0x4d, 0xdd, 0xc8, 0x05, 0xfb, 0xdf, 0x39, 0xde, 0xe5, 0x78, 0x08,
	0x5b, 0xd7, 0x1d, 0x29, 0xb3, 0x98, 0xc7, 0x76, 0xad, 0x91, 0x49, 0x27, 0x25, 0x3e, 0x01, 0x28,
	0x14, 0x71, 0x4d, 0xe5, 0x94, 0x6b, 0xb6, 0x67, 0x65, 0x83, 0x25, 0x72, 0xaa, 0x4d, 0xf9, 0xba,
	0x2d, 0xef, 0xca, 0xfb, 0xae, 0xbc, 0x44, 0x4e, 0x35, 0xbe, 0x00, 0xbf, 0xe0, 0x9a, 0x1d, 0x24,
	0x5e, 0x1a, 0x8f, 0x8f, 0xb2, 0x55, 0xf7, 0xce, 0xce, 0xb8, 0xce, 0x0d, 0xcb, 0x2c, 0x49, 0x94,
	0x1d, 0x3b, 0x4c, 0xfc, 0xd4, 0xcf, 0x4d, 0x38, 0xfa, 0x0c, 0xf1, 0x47, 0xdb, 0xeb, 0x21, 0x67,
	0x5f, 0x73, 0xe2, 0x15, 0x1b, 0x19, 0xfd, 0xe8, 0xc1, 0x8e, 0x6d, 0x9d, 0x53, 0xd7, 0x36, 0x75,
	0x47, 0x78, 0x02, 0x11, 0xd5, 0x5a, 0xe8, 0x85, 0xb5, 0x89, 0xc7, 0xc7, 0xab, 0x67, 0x76, 0xa2,
	0x25, 0x15, 0x5f, 0x42, 0xd8, 0xf2, 0x19, 0x29, 0x3b, 0xc7, 0x5a, 0xcd, 0x85, 0xa1, 0xe4, 0x8e,
	0x69, 0x24, 0x42, 0x93, 0xec, 0x98, 0x9f, 0xf8, 0x7f, 0xb3, 0x71, 0x4c, 0x23, 0x21, 0xa5, 0x1a,
	0x65, 0x9f, 0xb5, 0x56, 0xf2, 0xc6, 0x50, 0x72, 0xc7, 0xc4, 0x0c, 0x02, 0x51, 0x5f, 0x35, 0xf6,
	0x8b, 0x8e, 0xc7, 0xc3, 0x35, 0x26, 0xf5, 0x55, 0x93, 0x5b, 0xde, 0xf8, 0xa7, 0x0f, 0xdb, 0xd6,
	0xf3, 0x92, 0xd4, 0x8d, 0x28, 0x08, 0x73, 0x88, 0xdc, 0x01, 0xf0, 0xd9, 0x6a, 0xf1, 0x1f, 0xe7,
	0x19, 0x3e, 0xbf, 0xef, 0x11, 0xcb, 0x05, 0x8f, 0xfe, 0xc3, 0x73, 0x88, 0x5e, 0xd3, 0x9c, 0x34,
	0xe1, 0x7d, 0xaf, 0xfe, 0xd7, 0x6e, 0x13, 0xf0, 0xdf, 0x92, 0xde, 0x48, 0xab, 0x77, 0x10, 0x9c,
	0x8b, 0x6e, 0x33, 0xbd, 0x2e, 0x20, 0xba, 0x24, 0xae, 0x8a, 0x0a, 0x9f, 0xae, 0x16, 0xbc, 0xe2,
	0x1d, 0x7d, 0xaa, 0x48, 0x3d, 0x64, 0x6d, 0x97, 0xa4, 0xcf, 0xf8, 0x46, 0xe6, 0xfb, 0x12, 0xd9,
	0x5f, 0xea, 0xc9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0xe3, 0xb9, 0x9b, 0x90, 0x05, 0x00,
	0x00,
}
