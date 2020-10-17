// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kindTag.proto

package geiqin_srv_eat

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

type KindTagWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top                  int32    `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	Id                   int32    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int32  `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	IsSystem             bool     `protobuf:"varint,6,opt,name=is_system,json=isSystem,proto3" json:"is_system,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KindTagWhere) Reset()         { *m = KindTagWhere{} }
func (m *KindTagWhere) String() string { return proto.CompactTextString(m) }
func (*KindTagWhere) ProtoMessage()    {}
func (*KindTagWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_656ab9901fc954f5, []int{0}
}

func (m *KindTagWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KindTagWhere.Unmarshal(m, b)
}
func (m *KindTagWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KindTagWhere.Marshal(b, m, deterministic)
}
func (m *KindTagWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KindTagWhere.Merge(m, src)
}
func (m *KindTagWhere) XXX_Size() int {
	return xxx_messageInfo_KindTagWhere.Size(m)
}
func (m *KindTagWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_KindTagWhere.DiscardUnknown(m)
}

var xxx_messageInfo_KindTagWhere proto.InternalMessageInfo

func (m *KindTagWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *KindTagWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *KindTagWhere) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *KindTagWhere) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *KindTagWhere) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *KindTagWhere) GetIsSystem() bool {
	if m != nil {
		return m.IsSystem
	}
	return false
}

type KindTag struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: validate:"required" label:"菜单标签名称"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" validate:"required" label:"菜单标签名称"`
	// @inject_tag: validate:"required" label:"菜单图标"
	Icon                 string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty" validate:"required" label:"菜单图标"`
	IsSystem             bool     `protobuf:"varint,4,opt,name=is_system,json=isSystem,proto3" json:"is_system,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KindTag) Reset()         { *m = KindTag{} }
func (m *KindTag) String() string { return proto.CompactTextString(m) }
func (*KindTag) ProtoMessage()    {}
func (*KindTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_656ab9901fc954f5, []int{1}
}

func (m *KindTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KindTag.Unmarshal(m, b)
}
func (m *KindTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KindTag.Marshal(b, m, deterministic)
}
func (m *KindTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KindTag.Merge(m, src)
}
func (m *KindTag) XXX_Size() int {
	return xxx_messageInfo_KindTag.Size(m)
}
func (m *KindTag) XXX_DiscardUnknown() {
	xxx_messageInfo_KindTag.DiscardUnknown(m)
}

var xxx_messageInfo_KindTag proto.InternalMessageInfo

func (m *KindTag) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *KindTag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KindTag) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *KindTag) GetIsSystem() bool {
	if m != nil {
		return m.IsSystem
	}
	return false
}

func (m *KindTag) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *KindTag) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type KindTagResponse struct {
	Entity               *KindTag   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*KindTag `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *KindTagResponse) Reset()         { *m = KindTagResponse{} }
func (m *KindTagResponse) String() string { return proto.CompactTextString(m) }
func (*KindTagResponse) ProtoMessage()    {}
func (*KindTagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_656ab9901fc954f5, []int{2}
}

func (m *KindTagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KindTagResponse.Unmarshal(m, b)
}
func (m *KindTagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KindTagResponse.Marshal(b, m, deterministic)
}
func (m *KindTagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KindTagResponse.Merge(m, src)
}
func (m *KindTagResponse) XXX_Size() int {
	return xxx_messageInfo_KindTagResponse.Size(m)
}
func (m *KindTagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KindTagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KindTagResponse proto.InternalMessageInfo

func (m *KindTagResponse) GetEntity() *KindTag {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *KindTagResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *KindTagResponse) GetItems() []*KindTag {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *KindTagResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *KindTagResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*KindTagWhere)(nil), "geiqin.srv.eat.KindTagWhere")
	proto.RegisterType((*KindTag)(nil), "geiqin.srv.eat.KindTag")
	proto.RegisterType((*KindTagResponse)(nil), "geiqin.srv.eat.KindTagResponse")
}

func init() {
	proto.RegisterFile("kindTag.proto", fileDescriptor_656ab9901fc954f5)
}

var fileDescriptor_656ab9901fc954f5 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0xcd, 0x87, 0xed, 0xed, 0x5a, 0x65, 0x58, 0x31, 0xac, 0x8a, 0xa5, 0x4f, 0x05, 0x31,
	0x42, 0xfd, 0x05, 0xc5, 0x95, 0xa5, 0xe8, 0x83, 0xa4, 0x8a, 0x8f, 0x25, 0x26, 0x77, 0xbb, 0x17,
	0xcd, 0x4c, 0x9c, 0x19, 0x17, 0x76, 0x7f, 0x86, 0xcf, 0xfe, 0x16, 0xff, 0x95, 0xef, 0x32, 0x77,
	0x66, 0x85, 0x16, 0x2b, 0x0b, 0x79, 0xbb, 0x39, 0xe7, 0xcc, 0x99, 0x33, 0x87, 0x1b, 0xb8, 0xf7,
	0x85, 0x64, 0xf3, 0xa1, 0xda, 0x16, 0x9d, 0x56, 0x56, 0x89, 0xc9, 0x16, 0xe9, 0x1b, 0xc9, 0xc2,
	0xe8, 0xcb, 0x02, 0x2b, 0x7b, 0x72, 0x54, 0xab, 0xb6, 0x55, 0xd2, 0xb3, 0xb3, 0x1f, 0x11, 0x1c,
	0xbd, 0xf5, 0xfa, 0x4f, 0x17, 0xa8, 0x51, 0x1c, 0x43, 0xda, 0x55, 0x5b, 0x6c, 0xf2, 0x68, 0x1a,
	0xcd, 0xd3, 0xd2, 0x7f, 0x88, 0xc7, 0x30, 0x72, 0xc3, 0xc6, 0xd0, 0x35, 0xe6, 0x03, 0x66, 0x86,
	0x0e, 0x58, 0xd3, 0x35, 0x8a, 0x07, 0x10, 0x5b, 0xd5, 0xe5, 0x31, 0xc3, 0x6e, 0x14, 0x13, 0x18,
	0x50, 0x93, 0x27, 0x0c, 0x0c, 0xa8, 0x71, 0x0a, 0x6a, 0x4c, 0x9e, 0x4e, 0x63, 0xa7, 0xa0, 0xc6,
	0x38, 0x43, 0x32, 0x1b, 0x73, 0x65, 0x2c, 0xb6, 0x79, 0x36, 0x8d, 0xe6, 0xc3, 0x72, 0x48, 0x66,
	0xcd, 0xdf, 0xb3, 0x9f, 0x11, 0xdc, 0x0d, 0xa1, 0x82, 0x55, 0xf4, 0xd7, 0x4a, 0x40, 0x22, 0xab,
	0xd6, 0x87, 0x18, 0x95, 0x3c, 0x3b, 0x8c, 0x6a, 0x25, 0x39, 0xc1, 0xa8, 0xe4, 0x79, 0xf7, 0x82,
	0x64, 0xf7, 0x02, 0xf1, 0x14, 0xa0, 0xd6, 0x58, 0x59, 0x6c, 0x36, 0x95, 0xcd, 0x53, 0x3e, 0x36,
	0x0a, 0xc8, 0xd2, 0x3a, 0xfa, 0x7b, 0xd7, 0xdc, 0xd0, 0x99, 0xa7, 0x03, 0xb2, 0xb4, 0xb3, 0xdf,
	0x11, 0xdc, 0x0f, 0xf1, 0x4a, 0x34, 0x9d, 0x92, 0x06, 0xc5, 0x4b, 0xc8, 0x50, 0x5a, 0xb2, 0x57,
	0x1c, 0x75, 0xbc, 0x78, 0x54, 0xec, 0xd6, 0x5e, 0xdc, 0x1c, 0x08, 0x32, 0xf1, 0xdc, 0xf7, 0xac,
	0xf9, 0x21, 0xe3, 0xc5, 0xc3, 0x7d, 0xfd, 0x7b, 0x47, 0xfa, 0xfa, 0xb5, 0x78, 0x01, 0x29, 0x59,
	0x6c, 0x4d, 0x1e, 0x4f, 0xe3, 0xff, 0x99, 0x7b, 0x95, 0xf3, 0x46, 0xad, 0x95, 0xe6, 0x77, 0xff,
	0xc3, 0xfb, 0x8d, 0x23, 0x4b, 0xaf, 0x11, 0x73, 0x48, 0x48, 0x9e, 0x2b, 0x6e, 0x61, 0xbc, 0x38,
	0xde, 0xd7, 0xae, 0xe4, 0xb9, 0x2a, 0x59, 0xb1, 0xf8, 0x15, 0xc3, 0x24, 0xdc, 0xb4, 0x46, 0x7d,
	0x49, 0x35, 0x8a, 0x15, 0x64, 0x6b, 0xac, 0x74, 0x7d, 0x21, 0x9e, 0x1c, 0xc8, 0xc4, 0x5b, 0x75,
	0xf2, 0xec, 0x50, 0xe2, 0xd0, 0xdf, 0xec, 0x8e, 0x38, 0x83, 0xe4, 0x1d, 0x19, 0xdb, 0xdf, 0x68,
	0x09, 0xf1, 0x19, 0x5a, 0x71, 0xa8, 0xa4, 0xdb, 0x58, 0x9c, 0x42, 0xf6, 0x9a, 0xb7, 0xa1, 0xaf,
	0xcb, 0x47, 0x5e, 0x9a, 0x5e, 0x2e, 0x2b, 0xc8, 0x4e, 0xf1, 0x2b, 0x5a, 0xec, 0xdd, 0xcc, 0xe7,
	0x8c, 0xff, 0xf9, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x71, 0x0e, 0xb0, 0xa1, 0x22, 0x04,
	0x00, 0x00,
}
