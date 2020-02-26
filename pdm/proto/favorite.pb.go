// Code generated by protoc-gen-go. DO NOT EDIT.
// source: favorite.proto

//import "proto/item.proto";

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
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           int64    `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ItemId               int64    `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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
	return fileDescriptor_d331d96d24e0e82b, []int{1}
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
	proto.RegisterType((*FavoriteResponse)(nil), "geiqin.srv.pdm.FavoriteResponse")
}

func init() { proto.RegisterFile("favorite.proto", fileDescriptor_d331d96d24e0e82b) }

var fileDescriptor_d331d96d24e0e82b = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x4b, 0xe3, 0x50,
	0x14, 0x85, 0x27, 0x49, 0x93, 0x99, 0xde, 0x0c, 0x9d, 0xe1, 0x32, 0xc3, 0x64, 0x0a, 0x62, 0xe8,
	0xaa, 0x20, 0x04, 0x89, 0x6b, 0xd1, 0x5a, 0x6d, 0xe9, 0x4e, 0xd2, 0x85, 0xcb, 0x12, 0xf3, 0x6e,
	0xdb, 0x87, 0x26, 0x2f, 0xbe, 0xbc, 0x16, 0xfc, 0x1b, 0xfe, 0x46, 0x7f, 0x84, 0x4b, 0xc9, 0x4b,
	0x22, 0x58, 0xad, 0x94, 0x2e, 0x73, 0xbf, 0x73, 0x4e, 0xce, 0xbd, 0x3c, 0xe8, 0xcc, 0xe3, 0xb5,
	0x90, 0x5c, 0x51, 0x90, 0x4b, 0xa1, 0x04, 0x76, 0x16, 0xc4, 0x1f, 0x78, 0x16, 0x14, 0x72, 0x1d,
	0xe4, 0x2c, 0xed, 0xfe, 0x4c, 0x44, 0x9a, 0x8a, 0xac, 0xa2, 0xbd, 0x27, 0x03, 0x7e, 0x8c, 0x6a,
	0x03, 0x76, 0xc0, 0xe4, 0xcc, 0x33, 0x7c, 0xa3, 0x6f, 0x45, 0x26, 0x67, 0x78, 0x08, 0x6e, 0xb2,
	0x2a, 0x94, 0x48, 0x49, 0xce, 0x38, 0xf3, 0x4c, 0x0d, 0xa0, 0x19, 0x4d, 0x18, 0xfe, 0x83, 0xef,
	0x5c, 0x51, 0x5a, 0x42, 0x4b, 0x43, 0xa7, 0xfc, 0x9c, 0x30, 0x3c, 0x00, 0x48, 0x24, 0xc5, 0x8a,
	0xd8, 0x2c, 0x56, 0x5e, 0xcb, 0x37, 0xfa, 0xed, 0xa8, 0x5d, 0x4f, 0x06, 0xaa, 0xc4, 0xab, 0x9c,
	0x35, 0xd8, 0xae, 0x70, 0x3d, 0x19, 0xa8, 0xde, 0x8b, 0x01, 0xbf, 0x9b, 0x52, 0x11, 0x15, 0xb9,
	0xc8, 0x0a, 0xc2, 0x63, 0x70, 0x28, 0x53, 0x5c, 0x3d, 0xea, 0x82, 0x6e, 0xe8, 0x05, 0xef, 0x17,
	0x0b, 0xde, 0x1c, 0xb5, 0x0e, 0x8f, 0xc0, 0xce, 0xe3, 0x05, 0x49, 0x5d, 0xdc, 0x0d, 0xff, 0x6e,
	0x1a, 0xae, 0x4b, 0x18, 0x55, 0x1a, 0x0c, 0xc0, 0x2e, 0xbb, 0x17, 0x9e, 0xe5, 0x5b, 0x5f, 0xa6,
	0x57, 0xb2, 0x32, 0x9c, 0xa4, 0x14, 0x52, 0x2f, 0xf7, 0x49, 0xf8, 0x55, 0x09, 0xa3, 0x4a, 0x83,
	0x7d, 0x68, 0xf1, 0x6c, 0x2e, 0xf4, 0xa6, 0x6e, 0xf8, 0x67, 0x53, 0x3b, 0xc9, 0xe6, 0x22, 0xd2,
	0x8a, 0xf0, 0xd9, 0x84, 0x5f, 0xcd, 0xaf, 0xa6, 0x24, 0xd7, 0x3c, 0x21, 0x1c, 0x81, 0x33, 0xd4,
	0xa7, 0xc3, 0xad, 0xad, 0xba, 0xfe, 0xd6, 0xbe, 0xf5, 0xfd, 0x7a, 0xdf, 0xf0, 0x1c, 0x9c, 0x4b,
	0xba, 0x27, 0x45, 0x88, 0x1f, 0x1a, 0xb0, 0x9d, 0x12, 0x4e, 0xc1, 0x1a, 0x93, 0xda, 0xdb, 0x7e,
	0x06, 0xf6, 0x70, 0x49, 0xc9, 0xdd, 0xde, 0x01, 0x63, 0x70, 0xa6, 0x14, 0xcb, 0x64, 0x89, 0xff,
	0x37, 0xd5, 0x17, 0x71, 0x41, 0x37, 0x4b, 0x92, 0x3b, 0x9d, 0xe2, 0xd6, 0xd1, 0xaf, 0xff, 0xe4,
	0x35, 0x00, 0x00, 0xff, 0xff, 0x48, 0x6c, 0xc8, 0x94, 0x2d, 0x03, 0x00, 0x00,
}
