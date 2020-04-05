// Code generated by protoc-gen-go. DO NOT EDIT.
// source: itemSaleStats.proto

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

type ItemSaleStats struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId               int64    `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SaleNum              int32    `protobuf:"varint,3,opt,name=sale_num,json=saleNum,proto3" json:"sale_num,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemSaleStats) Reset()         { *m = ItemSaleStats{} }
func (m *ItemSaleStats) String() string { return proto.CompactTextString(m) }
func (*ItemSaleStats) ProtoMessage()    {}
func (*ItemSaleStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d2e52fef479617, []int{0}
}

func (m *ItemSaleStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemSaleStats.Unmarshal(m, b)
}
func (m *ItemSaleStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemSaleStats.Marshal(b, m, deterministic)
}
func (m *ItemSaleStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemSaleStats.Merge(m, src)
}
func (m *ItemSaleStats) XXX_Size() int {
	return xxx_messageInfo_ItemSaleStats.Size(m)
}
func (m *ItemSaleStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemSaleStats.DiscardUnknown(m)
}

var xxx_messageInfo_ItemSaleStats proto.InternalMessageInfo

func (m *ItemSaleStats) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemSaleStats) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ItemSaleStats) GetSaleNum() int32 {
	if m != nil {
		return m.SaleNum
	}
	return 0
}

func (m *ItemSaleStats) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ItemSaleStats) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ItemSaleStatsResponse struct {
	Entity               *ItemSaleStats   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager           `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*ItemSaleStats `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error           `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info            `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ItemSaleStatsResponse) Reset()         { *m = ItemSaleStatsResponse{} }
func (m *ItemSaleStatsResponse) String() string { return proto.CompactTextString(m) }
func (*ItemSaleStatsResponse) ProtoMessage()    {}
func (*ItemSaleStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d2e52fef479617, []int{1}
}

func (m *ItemSaleStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemSaleStatsResponse.Unmarshal(m, b)
}
func (m *ItemSaleStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemSaleStatsResponse.Marshal(b, m, deterministic)
}
func (m *ItemSaleStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemSaleStatsResponse.Merge(m, src)
}
func (m *ItemSaleStatsResponse) XXX_Size() int {
	return xxx_messageInfo_ItemSaleStatsResponse.Size(m)
}
func (m *ItemSaleStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemSaleStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemSaleStatsResponse proto.InternalMessageInfo

func (m *ItemSaleStatsResponse) GetEntity() *ItemSaleStats {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ItemSaleStatsResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ItemSaleStatsResponse) GetItems() []*ItemSaleStats {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ItemSaleStatsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ItemSaleStatsResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemSaleStats)(nil), "geiqin.srv.pdm.ItemSaleStats")
	proto.RegisterType((*ItemSaleStatsResponse)(nil), "geiqin.srv.pdm.ItemSaleStatsResponse")
}

func init() {
	proto.RegisterFile("itemSaleStats.proto", fileDescriptor_a9d2e52fef479617)
}

var fileDescriptor_a9d2e52fef479617 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x69, 0xbb, 0x6e, 0xee, 0x4c, 0x77, 0x11, 0x1d, 0x56, 0x61, 0x50, 0x76, 0x55, 0x10,
	0x7a, 0xd1, 0xe1, 0x03, 0xec, 0xc2, 0x8b, 0xdd, 0x88, 0x64, 0x0f, 0x50, 0xe2, 0x72, 0x36, 0x02,
	0x4b, 0x52, 0x93, 0x54, 0xf0, 0x09, 0xbc, 0xf6, 0x8d, 0x25, 0x69, 0x05, 0xab, 0x82, 0x97, 0xe7,
	0x7c, 0xff, 0x9f, 0xff, 0xfc, 0x04, 0x2e, 0x85, 0x43, 0xb9, 0x63, 0x27, 0xdc, 0x39, 0xe6, 0x6c,
	0xd9, 0x18, 0xed, 0x34, 0x99, 0x1f, 0x51, 0xbc, 0x08, 0x55, 0x5a, 0xf3, 0x5a, 0x36, 0x5c, 0xde,
	0x9e, 0xef, 0xb5, 0x94, 0x5a, 0x75, 0x74, 0xf5, 0x11, 0xc1, 0xc5, 0xf6, 0xbb, 0x8b, 0xcc, 0x21,
	0x16, 0x3c, 0x8b, 0xf2, 0xa8, 0x48, 0x68, 0x2c, 0x38, 0xb9, 0x86, 0x89, 0x7f, 0xb6, 0x16, 0x3c,
	0x8b, 0xc3, 0x72, 0xec, 0xc7, 0x2d, 0x27, 0x37, 0x70, 0x66, 0xd9, 0x09, 0x6b, 0xd5, 0xca, 0x2c,
	0xc9, 0xa3, 0x22, 0xa5, 0x13, 0x3f, 0x3f, 0xb6, 0x92, 0x2c, 0x01, 0xf6, 0x06, 0x99, 0x43, 0x5e,
	0x33, 0x97, 0x8d, 0xf2, 0xa8, 0x98, 0xd2, 0x69, 0xbf, 0xd9, 0x38, 0x8f, 0xdb, 0x86, 0x7f, 0xe1,
	0xb4, 0xc3, 0xfd, 0x66, 0xe3, 0x56, 0xef, 0x31, 0x2c, 0x06, 0x37, 0x51, 0xb4, 0x8d, 0x56, 0x16,
	0xc9, 0x3d, 0x8c, 0x51, 0x39, 0xe1, 0xde, 0xc2, 0x7d, 0xb3, 0x6a, 0x59, 0x0e, 0xcb, 0x95, 0x43,
	0x5b, 0x2f, 0x26, 0x77, 0x90, 0x36, 0xec, 0x88, 0x26, 0x14, 0x98, 0x55, 0x8b, 0x9f, 0xae, 0x27,
	0x0f, 0x69, 0xa7, 0x21, 0x6b, 0x48, 0x7d, 0x41, 0x9b, 0x25, 0x79, 0xf2, 0x7f, 0x44, 0xa7, 0xf5,
	0x09, 0x68, 0x8c, 0x36, 0xa1, 0xeb, 0x1f, 0x09, 0x0f, 0x1e, 0xd2, 0x4e, 0x43, 0x0a, 0x18, 0x09,
	0x75, 0xd0, 0xa1, 0xf8, 0xac, 0xba, 0xfa, 0x15, 0xa0, 0x0e, 0x9a, 0x06, 0xc5, 0xf3, 0x38, 0x7c,
	0xd2, 0xfa, 0x33, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xda, 0x12, 0x75, 0xd9, 0x01, 0x00, 0x00,
}
