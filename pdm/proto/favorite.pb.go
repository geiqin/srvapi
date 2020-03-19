// Code generated by protoc-gen-go. DO NOT EDIT.
// source: favorite.proto

package geiqin_srv_pdm

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

type Favorite struct {
	Id                   int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           int64        `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ItemId               int64        `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	CreatedAt            string       `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string       `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Item                 *ItemDisplay `protobuf:"bytes,6,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Favorite) Reset()         { *m = Favorite{} }
func (m *Favorite) String() string { return proto.CompactTextString(m) }
func (*Favorite) ProtoMessage()    {}
func (*Favorite) Descriptor() ([]byte, []int) {
	return fileDescriptor_d331d96d24e0e82b, []int{0}
}

func (m *Favorite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Favorite.Unmarshal(m, b)
}
func (m *Favorite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Favorite.Marshal(b, m, deterministic)
}
func (m *Favorite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Favorite.Merge(m, src)
}
func (m *Favorite) XXX_Size() int {
	return xxx_messageInfo_Favorite.Size(m)
}
func (m *Favorite) XXX_DiscardUnknown() {
	xxx_messageInfo_Favorite.DiscardUnknown(m)
}

var xxx_messageInfo_Favorite proto.InternalMessageInfo

func (m *Favorite) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Favorite) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Favorite) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *Favorite) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Favorite) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Favorite) GetItem() *ItemDisplay {
	if m != nil {
		return m.Item
	}
	return nil
}

type FavoriteWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top                  int32    `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	CustomerId           int64    `protobuf:"varint,6,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ItemId               int64    `protobuf:"varint,7,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FavoriteWhere) Reset()         { *m = FavoriteWhere{} }
func (m *FavoriteWhere) String() string { return proto.CompactTextString(m) }
func (*FavoriteWhere) ProtoMessage()    {}
func (*FavoriteWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_d331d96d24e0e82b, []int{1}
}

func (m *FavoriteWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FavoriteWhere.Unmarshal(m, b)
}
func (m *FavoriteWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FavoriteWhere.Marshal(b, m, deterministic)
}
func (m *FavoriteWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FavoriteWhere.Merge(m, src)
}
func (m *FavoriteWhere) XXX_Size() int {
	return xxx_messageInfo_FavoriteWhere.Size(m)
}
func (m *FavoriteWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_FavoriteWhere.DiscardUnknown(m)
}

var xxx_messageInfo_FavoriteWhere proto.InternalMessageInfo

func (m *FavoriteWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *FavoriteWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *FavoriteWhere) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *FavoriteWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FavoriteWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FavoriteWhere) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *FavoriteWhere) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

type FavoriteResponse struct {
	Entity               *Favorite   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Favorite `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FavoriteResponse) Reset()         { *m = FavoriteResponse{} }
func (m *FavoriteResponse) String() string { return proto.CompactTextString(m) }
func (*FavoriteResponse) ProtoMessage()    {}
func (*FavoriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d331d96d24e0e82b, []int{2}
}

func (m *FavoriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FavoriteResponse.Unmarshal(m, b)
}
func (m *FavoriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FavoriteResponse.Marshal(b, m, deterministic)
}
func (m *FavoriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FavoriteResponse.Merge(m, src)
}
func (m *FavoriteResponse) XXX_Size() int {
	return xxx_messageInfo_FavoriteResponse.Size(m)
}
func (m *FavoriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FavoriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FavoriteResponse proto.InternalMessageInfo

func (m *FavoriteResponse) GetEntity() *Favorite {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *FavoriteResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *FavoriteResponse) GetItems() []*Favorite {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *FavoriteResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *FavoriteResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Favorite)(nil), "geiqin.srv.pdm.Favorite")
	proto.RegisterType((*FavoriteWhere)(nil), "geiqin.srv.pdm.FavoriteWhere")
	proto.RegisterType((*FavoriteResponse)(nil), "geiqin.srv.pdm.FavoriteResponse")
}

func init() { proto.RegisterFile("favorite.proto", fileDescriptor_d331d96d24e0e82b) }

var fileDescriptor_d331d96d24e0e82b = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0xec, 0x36, 0x13, 0x08, 0xe9, 0xa8, 0x08, 0x2b, 0x55, 0x85, 0x95, 0x93, 0x25,
	0x24, 0x83, 0xcc, 0x2f, 0xa8, 0xfa, 0x21, 0x45, 0x08, 0x09, 0x6d, 0x0e, 0xdc, 0x88, 0x8c, 0x77,
	0xd2, 0xac, 0xa8, 0xbd, 0x66, 0xbd, 0x8d, 0x94, 0xfe, 0x2a, 0x7e, 0x00, 0x37, 0xfe, 0x14, 0x47,
	0xb4, 0xbb, 0x76, 0x81, 0x40, 0x2a, 0xa4, 0xf6, 0xb6, 0x3b, 0x6f, 0xf6, 0xcd, 0x7b, 0x6f, 0x2c,
	0xc3, 0x68, 0x99, 0xaf, 0xa5, 0x12, 0x9a, 0xd2, 0x5a, 0x49, 0x2d, 0x71, 0x74, 0x49, 0xe2, 0x8b,
	0xa8, 0xd2, 0x46, 0xad, 0xd3, 0x9a, 0x97, 0x93, 0xc7, 0x85, 0x2c, 0x4b, 0x59, 0x39, 0x74, 0x72,
	0x20, 0x34, 0x95, 0x67, 0xa2, 0xa9, 0xaf, 0xf2, 0x8d, 0x2b, 0x4d, 0xbf, 0x7b, 0xb0, 0x7f, 0xd1,
	0x72, 0xe0, 0x08, 0x7a, 0x82, 0x47, 0x5e, 0xec, 0x25, 0x3e, 0xeb, 0x09, 0x8e, 0x2f, 0x60, 0x58,
	0x5c, 0x37, 0x5a, 0x96, 0xa4, 0x16, 0x82, 0x47, 0x3d, 0x0b, 0x40, 0x57, 0x9a, 0x71, 0x7c, 0x0e,
	0x7b, 0x86, 0xd2, 0x80, 0xbe, 0x05, 0x43, 0x73, 0x9d, 0x71, 0x3c, 0x06, 0x28, 0x14, 0xe5, 0x9a,
	0xf8, 0x22, 0xd7, 0x51, 0x3f, 0xf6, 0x92, 0x01, 0x1b, 0xb4, 0x95, 0x13, 0x6d, 0xe0, 0xeb, 0x9a,
	0x77, 0x70, 0xe0, 0xe0, 0xb6, 0x72, 0xa2, 0xf1, 0x15, 0xf4, 0x0d, 0x4f, 0x14, 0xc6, 0x5e, 0x32,
	0xcc, 0x8e, 0xd2, 0x3f, 0x4d, 0xa5, 0xb3, 0x5f, 0x2e, 0x98, 0x6d, 0x9c, 0x7e, 0xf5, 0xe0, 0x49,
	0xe7, 0xe2, 0xc3, 0x8a, 0x14, 0xe1, 0x21, 0x04, 0x75, 0x7e, 0x49, 0xce, 0x4d, 0xc0, 0xdc, 0x05,
	0x8f, 0x60, 0x60, 0x0e, 0x8b, 0x46, 0xdc, 0x90, 0xb5, 0x13, 0xb0, 0x7d, 0x53, 0x98, 0x8b, 0x1b,
	0xc2, 0x31, 0xf8, 0x5a, 0xd6, 0xd6, 0x48, 0xc0, 0xcc, 0xb1, 0xcd, 0xa3, 0x7f, 0x9b, 0xc7, 0x18,
	0x7c, 0xc1, 0x9b, 0x28, 0x88, 0xfd, 0xc4, 0x67, 0xe6, 0xb8, 0x9d, 0x50, 0x78, 0x57, 0x42, 0x7b,
	0xbf, 0x27, 0x34, 0xfd, 0xe1, 0xc1, 0xb8, 0x93, 0xcc, 0xa8, 0xa9, 0x65, 0xd5, 0x10, 0xbe, 0x86,
	0x90, 0x2a, 0x2d, 0xf4, 0xc6, 0xca, 0x1e, 0x66, 0xd1, 0xb6, 0xf5, 0xdb, 0x17, 0x6d, 0x1f, 0xbe,
	0x74, 0x3e, 0x95, 0x75, 0x33, 0xcc, 0x9e, 0x6d, 0x3f, 0x78, 0x6f, 0x40, 0x67, 0x5f, 0x61, 0x0a,
	0x81, 0x99, 0xde, 0x44, 0x7e, 0xec, 0xdf, 0xc9, 0xee, 0xda, 0x0c, 0x39, 0x29, 0x25, 0x95, 0x8d,
	0xe0, 0x1f, 0xe4, 0xe7, 0x06, 0x64, 0xae, 0x07, 0x13, 0xe8, 0x8b, 0x6a, 0x29, 0xed, 0x36, 0x87,
	0xd9, 0xe1, 0x5f, 0x4b, 0xab, 0x96, 0x92, 0xd9, 0x8e, 0xec, 0x23, 0x3c, 0xed, 0x26, 0xcd, 0x49,
	0xad, 0x45, 0x41, 0xf8, 0x16, 0xc2, 0x39, 0xe5, 0xaa, 0x58, 0xe1, 0xf1, 0x2e, 0x51, 0x76, 0xaf,
	0x93, 0x78, 0xa7, 0xe6, 0x36, 0xc3, 0xe9, 0xa3, 0xec, 0x5b, 0x0f, 0x0e, 0xde, 0x6d, 0xb6, 0x47,
	0x5c, 0x40, 0x78, 0x6a, 0x3f, 0x40, 0xdc, 0xe9, 0xfb, 0x7f, 0xd8, 0x8d, 0xd4, 0x33, 0xba, 0x22,
	0x4d, 0x0f, 0x20, 0x15, 0xcf, 0x21, 0x38, 0x5d, 0x51, 0xf1, 0xf9, 0xfe, 0x9a, 0x1e, 0x2c, 0xbe,
	0x4f, 0xa1, 0xfd, 0x33, 0xbc, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x85, 0x62, 0x00, 0xe6, 0x5c,
	0x04, 0x00, 0x00,
}
