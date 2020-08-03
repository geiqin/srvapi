// Code generated by protoc-gen-go. DO NOT EDIT.
// source: download.proto

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

// 下载信息
type Download struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	FileName             string   `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	RawName              string   `protobuf:"bytes,6,opt,name=raw_name,json=rawName,proto3" json:"raw_name,omitempty"`
	Path                 string   `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
	Url                  string   `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	MimeType             string   `protobuf:"bytes,9,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Size                 int32    `protobuf:"varint,10,opt,name=size,proto3" json:"size,omitempty"`
	Memo                 string   `protobuf:"bytes,11,opt,name=memo,proto3" json:"memo,omitempty"`
	UserId               int64    `protobuf:"varint,12,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status               int32    `protobuf:"varint,13,opt,name=status,proto3" json:"status,omitempty"`
	BuiltAt              string   `protobuf:"bytes,14,opt,name=built_at,json=builtAt,proto3" json:"built_at,omitempty"`
	CreatedAt            string   `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ids                  []int64  `protobuf:"varint,17,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Download) Reset()         { *m = Download{} }
func (m *Download) String() string { return proto.CompactTextString(m) }
func (*Download) ProtoMessage()    {}
func (*Download) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ce52b48c9eea83, []int{0}
}

func (m *Download) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Download.Unmarshal(m, b)
}
func (m *Download) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Download.Marshal(b, m, deterministic)
}
func (m *Download) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Download.Merge(m, src)
}
func (m *Download) XXX_Size() int {
	return xxx_messageInfo_Download.Size(m)
}
func (m *Download) XXX_DiscardUnknown() {
	xxx_messageInfo_Download.DiscardUnknown(m)
}

var xxx_messageInfo_Download proto.InternalMessageInfo

func (m *Download) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Download) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Download) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Download) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Download) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Download) GetRawName() string {
	if m != nil {
		return m.RawName
	}
	return ""
}

func (m *Download) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Download) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Download) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *Download) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Download) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Download) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Download) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Download) GetBuiltAt() string {
	if m != nil {
		return m.BuiltAt
	}
	return ""
}

func (m *Download) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Download) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Download) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type DownloadWhere struct {
	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	//base params
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Title                string   `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Hash                 string   `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
	Status               string   `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	BuiltAt              string   `protobuf:"bytes,10,opt,name=built_at,json=builtAt,proto3" json:"built_at,omitempty"`
	CreatedAt            string   `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ids                  []int64  `protobuf:"varint,13,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadWhere) Reset()         { *m = DownloadWhere{} }
func (m *DownloadWhere) String() string { return proto.CompactTextString(m) }
func (*DownloadWhere) ProtoMessage()    {}
func (*DownloadWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ce52b48c9eea83, []int{1}
}

func (m *DownloadWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadWhere.Unmarshal(m, b)
}
func (m *DownloadWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadWhere.Marshal(b, m, deterministic)
}
func (m *DownloadWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadWhere.Merge(m, src)
}
func (m *DownloadWhere) XXX_Size() int {
	return xxx_messageInfo_DownloadWhere.Size(m)
}
func (m *DownloadWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadWhere.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadWhere proto.InternalMessageInfo

func (m *DownloadWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *DownloadWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *DownloadWhere) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *DownloadWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DownloadWhere) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DownloadWhere) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DownloadWhere) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *DownloadWhere) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DownloadWhere) GetBuiltAt() string {
	if m != nil {
		return m.BuiltAt
	}
	return ""
}

func (m *DownloadWhere) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *DownloadWhere) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *DownloadWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

//
type DownloadResponse struct {
	Entity               *Download   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Download `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DownloadResponse) Reset()         { *m = DownloadResponse{} }
func (m *DownloadResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadResponse) ProtoMessage()    {}
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ce52b48c9eea83, []int{2}
}

func (m *DownloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadResponse.Unmarshal(m, b)
}
func (m *DownloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadResponse.Marshal(b, m, deterministic)
}
func (m *DownloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadResponse.Merge(m, src)
}
func (m *DownloadResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadResponse.Size(m)
}
func (m *DownloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadResponse proto.InternalMessageInfo

func (m *DownloadResponse) GetEntity() *Download {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *DownloadResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *DownloadResponse) GetItems() []*Download {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *DownloadResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DownloadResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Download)(nil), "geiqin.srv.cms.media.Download")
	proto.RegisterType((*DownloadWhere)(nil), "geiqin.srv.cms.media.DownloadWhere")
	proto.RegisterType((*DownloadResponse)(nil), "geiqin.srv.cms.media.DownloadResponse")
}

func init() {
	proto.RegisterFile("download.proto", fileDescriptor_f7ce52b48c9eea83)
}

var fileDescriptor_f7ce52b48c9eea83 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x25, 0xb1, 0xe3, 0x38, 0x93, 0x7e, 0x84, 0x55, 0x05, 0x4b, 0x2b, 0x50, 0x54, 0x24, 0x94,
	0x93, 0x25, 0x02, 0xe2, 0x5e, 0x68, 0x85, 0x7a, 0x01, 0xe4, 0x14, 0x7a, 0x8c, 0xb6, 0xf1, 0xb4,
	0x59, 0xc9, 0x5f, 0xac, 0xd7, 0xad, 0xca, 0x0f, 0x81, 0x5f, 0xc0, 0x4f, 0xe3, 0x7f, 0xa0, 0x99,
	0x75, 0x68, 0x85, 0xd2, 0xb4, 0x87, 0xdc, 0xde, 0xce, 0x9b, 0x7d, 0xb3, 0xfb, 0xde, 0x3a, 0x81,
	0xad, 0xa4, 0xb8, 0xca, 0xd3, 0x42, 0x25, 0x51, 0x69, 0x0a, 0x5b, 0x88, 0x9d, 0x0b, 0xd4, 0xdf,
	0x75, 0x1e, 0x55, 0xe6, 0x32, 0x9a, 0x65, 0x55, 0x94, 0x61, 0xa2, 0xd5, 0xee, 0xc6, 0xac, 0xc8,
	0xb2, 0x22, 0x77, 0x3d, 0xfb, 0xbf, 0x3c, 0x08, 0x0f, 0x9b, 0x6d, 0x62, 0x0b, 0xda, 0x3a, 0x91,
	0xad, 0x61, 0x6b, 0xe4, 0xc5, 0x6d, 0x9d, 0x08, 0x01, 0xfe, 0x5c, 0x55, 0x73, 0xd9, 0x1e, 0xb6,
	0x46, 0xbd, 0x98, 0xb1, 0xd8, 0x81, 0x8e, 0xd5, 0x36, 0x45, 0xe9, 0x71, 0xd1, 0x2d, 0xa8, 0xd3,
	0x5e, 0x97, 0x28, 0x7d, 0xd7, 0x49, 0x58, 0xec, 0x41, 0xef, 0x5c, 0xa7, 0x38, 0xcd, 0x55, 0x86,
	0xb2, 0xc3, 0x44, 0x48, 0x85, 0x4f, 0x2a, 0x43, 0xf1, 0x0c, 0x42, 0xa3, 0xae, 0x1c, 0x17, 0x30,
	0xd7, 0x35, 0xea, 0x8a, 0x29, 0x01, 0x7e, 0xa9, 0xec, 0x5c, 0x76, 0x9d, 0x16, 0x61, 0x31, 0x00,
	0xaf, 0x36, 0xa9, 0x0c, 0xb9, 0x44, 0x90, 0xd4, 0x33, 0x9d, 0xe1, 0x94, 0xc7, 0xf6, 0x9c, 0x3a,
	0x15, 0x4e, 0x68, 0xb4, 0x00, 0xbf, 0xd2, 0x3f, 0x50, 0xc2, 0xb0, 0x35, 0xea, 0xc4, 0x8c, 0xa9,
	0x96, 0x61, 0x56, 0xc8, 0xbe, 0x93, 0x25, 0x2c, 0x9e, 0x42, 0xb7, 0xae, 0xd0, 0x4c, 0x75, 0x22,
	0x37, 0xf8, 0xd6, 0x01, 0x2d, 0x8f, 0x13, 0xf1, 0x04, 0x82, 0xca, 0x2a, 0x5b, 0x57, 0x72, 0x93,
	0x25, 0x9a, 0x15, 0x1d, 0xfb, 0xac, 0xd6, 0xa9, 0x9d, 0x2a, 0x2b, 0xb7, 0xdc, 0xb1, 0x79, 0x7d,
	0x60, 0xc5, 0x73, 0x80, 0x99, 0x41, 0x65, 0x31, 0x21, 0x72, 0x9b, 0xc9, 0x5e, 0x53, 0x71, 0x74,
	0x5d, 0x26, 0x0b, 0x7a, 0xe0, 0xe8, 0xa6, 0x72, 0x60, 0xe9, 0x82, 0x3a, 0xa9, 0xe4, 0xe3, 0xa1,
	0x37, 0xf2, 0x62, 0x82, 0xfb, 0xbf, 0xdb, 0xb0, 0xb9, 0x48, 0xe6, 0x74, 0x8e, 0x06, 0xc9, 0xfa,
	0x52, 0x5d, 0xa0, 0x4b, 0xa8, 0x13, 0xbb, 0x05, 0x19, 0x41, 0x60, 0xca, 0x17, 0x6e, 0x33, 0x13,
	0x52, 0x61, 0x42, 0x97, 0x1e, 0x80, 0x67, 0x8b, 0x92, 0xb3, 0xea, 0xc4, 0x04, 0x9b, 0x8c, 0xfd,
	0xdb, 0x19, 0xdf, 0x0a, 0x81, 0xf1, 0x4d, 0xc6, 0xdd, 0xff, 0x32, 0xe6, 0xd7, 0x10, 0xde, 0x7a,
	0x0d, 0x37, 0x3e, 0xb9, 0x08, 0x96, 0xf9, 0x04, 0xab, 0x7c, 0xea, 0xaf, 0xf6, 0x69, 0xe3, 0x0e,
	0x9f, 0x36, 0x6f, 0x7c, 0xfa, 0xd9, 0x86, 0xc1, 0xc2, 0xa7, 0x18, 0xab, 0xb2, 0xc8, 0x2b, 0x14,
	0xef, 0x20, 0xc0, 0xdc, 0x6a, 0x7b, 0xcd, 0x5e, 0xf5, 0xc7, 0x2f, 0xa2, 0x65, 0xdf, 0x42, 0xf4,
	0x6f, 0x5f, 0xd3, 0x2d, 0x5e, 0x3b, 0x8b, 0x0d, 0x1b, 0xd9, 0x1f, 0xef, 0x2d, 0xdf, 0xf6, 0x85,
	0x5a, 0x9c, 0xff, 0x46, 0xbc, 0x85, 0x8e, 0xb6, 0x98, 0x55, 0xd2, 0x1b, 0x7a, 0x0f, 0x98, 0xe4,
	0x9a, 0x69, 0x10, 0x1a, 0x53, 0x18, 0x4e, 0xe2, 0xce, 0x41, 0x47, 0xd4, 0x12, 0xbb, 0x4e, 0x11,
	0x81, 0xaf, 0xf3, 0xf3, 0x82, 0x3f, 0xa5, 0xfe, 0x78, 0x77, 0xf9, 0x8e, 0xe3, 0xfc, 0xbc, 0x88,
	0xb9, 0x6f, 0xfc, 0xc7, 0x87, 0xed, 0xc5, 0xd8, 0x09, 0x9a, 0x4b, 0x3d, 0x43, 0xf1, 0x0d, 0xe0,
	0x7d, 0xad, 0xd3, 0x64, 0x62, 0x95, 0xb1, 0xe2, 0x9e, 0xb3, 0xee, 0xbe, 0xba, 0xe7, 0x2e, 0x8d,
	0xdb, 0xfb, 0x8f, 0xc4, 0x09, 0x84, 0xac, 0x7b, 0x94, 0x27, 0x6b, 0x54, 0x8d, 0x21, 0xf8, 0xc0,
	0x0f, 0x63, 0xbd, 0x9a, 0x5f, 0xf9, 0x35, 0xad, 0x57, 0xf3, 0x10, 0x53, 0x5c, 0xab, 0xe6, 0x67,
	0xf0, 0x3e, 0xe2, 0x3a, 0x23, 0x3a, 0x85, 0x60, 0x82, 0xca, 0xcc, 0xe6, 0xe2, 0xe5, 0xea, 0x3d,
	0xfc, 0x63, 0xf3, 0x70, 0xe1, 0xb3, 0x80, 0xff, 0x49, 0xde, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x3e, 0xab, 0xc9, 0x90, 0x7f, 0x06, 0x00, 0x00,
}
