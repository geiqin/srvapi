// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goods.proto

package geiqin_srv_dms

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

type Goods struct {
	Id                   int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId               int64              `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64              `protobuf:"varint,3,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Disabled             bool               `protobuf:"varint,6,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Default              bool               `protobuf:"varint,7,opt,name=default,proto3" json:"default,omitempty"`
	CreatedAt            string             `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string             `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	GoodsCommission      []*GoodsCommission `protobuf:"bytes,10,rep,name=goods_commission,json=goodsCommission,proto3" json:"goods_commission,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Goods) Reset()         { *m = Goods{} }
func (m *Goods) String() string { return proto.CompactTextString(m) }
func (*Goods) ProtoMessage()    {}
func (*Goods) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30593c5487368b0, []int{0}
}

func (m *Goods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Goods.Unmarshal(m, b)
}
func (m *Goods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Goods.Marshal(b, m, deterministic)
}
func (m *Goods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Goods.Merge(m, src)
}
func (m *Goods) XXX_Size() int {
	return xxx_messageInfo_Goods.Size(m)
}
func (m *Goods) XXX_DiscardUnknown() {
	xxx_messageInfo_Goods.DiscardUnknown(m)
}

var xxx_messageInfo_Goods proto.InternalMessageInfo

func (m *Goods) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Goods) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *Goods) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *Goods) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Goods) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

func (m *Goods) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Goods) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Goods) GetGoodsCommission() []*GoodsCommission {
	if m != nil {
		return m.GoodsCommission
	}
	return nil
}

type GoodsWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top                  int32    `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Keywords             string   `protobuf:"bytes,6,opt,name=keywords,proto3" json:"keywords,omitempty"`
	CatId                int32    `protobuf:"varint,7,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"`
	Type                 string   `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Sorting              string   `protobuf:"bytes,9,opt,name=sorting,proto3" json:"sorting,omitempty"`
	ItemId               int64    `protobuf:"varint,10,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,11,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	ItemName             string   `protobuf:"bytes,12,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	ItemPrice            string   `protobuf:"bytes,13,opt,name=item_price,json=itemPrice,proto3" json:"item_price,omitempty"`
	ItemIds              []int64  `protobuf:"varint,14,rep,packed,name=item_ids,json=itemIds,proto3" json:"item_ids,omitempty"`
	CustomerId           int64    `protobuf:"varint,15,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Client               bool     `protobuf:"varint,16,opt,name=client,proto3" json:"client,omitempty"`
	Disabled             bool     `protobuf:"varint,17,opt,name=disabled,proto3" json:"disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsWhere) Reset()         { *m = GoodsWhere{} }
func (m *GoodsWhere) String() string { return proto.CompactTextString(m) }
func (*GoodsWhere) ProtoMessage()    {}
func (*GoodsWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30593c5487368b0, []int{1}
}

func (m *GoodsWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsWhere.Unmarshal(m, b)
}
func (m *GoodsWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsWhere.Marshal(b, m, deterministic)
}
func (m *GoodsWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsWhere.Merge(m, src)
}
func (m *GoodsWhere) XXX_Size() int {
	return xxx_messageInfo_GoodsWhere.Size(m)
}
func (m *GoodsWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsWhere.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsWhere proto.InternalMessageInfo

func (m *GoodsWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *GoodsWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GoodsWhere) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *GoodsWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GoodsWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *GoodsWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *GoodsWhere) GetCatId() int32 {
	if m != nil {
		return m.CatId
	}
	return 0
}

func (m *GoodsWhere) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GoodsWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *GoodsWhere) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *GoodsWhere) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *GoodsWhere) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

func (m *GoodsWhere) GetItemPrice() string {
	if m != nil {
		return m.ItemPrice
	}
	return ""
}

func (m *GoodsWhere) GetItemIds() []int64 {
	if m != nil {
		return m.ItemIds
	}
	return nil
}

func (m *GoodsWhere) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *GoodsWhere) GetClient() bool {
	if m != nil {
		return m.Client
	}
	return false
}

func (m *GoodsWhere) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

type GoodsResponse struct {
	Entity               *Goods   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Goods `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsResponse) Reset()         { *m = GoodsResponse{} }
func (m *GoodsResponse) String() string { return proto.CompactTextString(m) }
func (*GoodsResponse) ProtoMessage()    {}
func (*GoodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a30593c5487368b0, []int{2}
}

func (m *GoodsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsResponse.Unmarshal(m, b)
}
func (m *GoodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsResponse.Marshal(b, m, deterministic)
}
func (m *GoodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsResponse.Merge(m, src)
}
func (m *GoodsResponse) XXX_Size() int {
	return xxx_messageInfo_GoodsResponse.Size(m)
}
func (m *GoodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsResponse proto.InternalMessageInfo

func (m *GoodsResponse) GetEntity() *Goods {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *GoodsResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *GoodsResponse) GetItems() []*Goods {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *GoodsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GoodsResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Goods)(nil), "geiqin.srv.dms.Goods")
	proto.RegisterType((*GoodsWhere)(nil), "geiqin.srv.dms.GoodsWhere")
	proto.RegisterType((*GoodsResponse)(nil), "geiqin.srv.dms.GoodsResponse")
}

func init() {
	proto.RegisterFile("goods.proto", fileDescriptor_a30593c5487368b0)
}

var fileDescriptor_a30593c5487368b0 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xfd, 0x92, 0xe9, 0xe4, 0xe7, 0xa6, 0x4d, 0xf3, 0x59, 0x2d, 0x98, 0x54, 0x55, 0x43, 0x56,
	0x91, 0x10, 0x59, 0x84, 0x35, 0x42, 0x55, 0x41, 0x51, 0x90, 0x40, 0x95, 0xb3, 0x60, 0xc1, 0x22,
	0x9a, 0x8e, 0x6f, 0x53, 0x2b, 0x9d, 0xf1, 0x60, 0x3b, 0x45, 0xe9, 0x0b, 0xf0, 0x3a, 0xbc, 0x12,
	0x4f, 0xc0, 0x2b, 0x20, 0xdf, 0x99, 0x84, 0xa6, 0x9a, 0x82, 0x04, 0x62, 0xe7, 0x7b, 0xce, 0xc9,
	0xb5, 0xef, 0x39, 0x37, 0x03, 0xad, 0xb9, 0xd6, 0xd2, 0x0e, 0x33, 0xa3, 0x9d, 0x66, 0xed, 0x39,
	0xaa, 0x4f, 0x2a, 0x1d, 0x5a, 0x73, 0x33, 0x94, 0x89, 0xed, 0xee, 0xc6, 0x3a, 0x49, 0x74, 0x9a,
	0xb3, 0xdd, 0x7d, 0x92, 0x4e, 0xd2, 0x4b, 0x5d, 0x00, 0x87, 0x04, 0x9c, 0xe9, 0x24, 0x51, 0xd6,
	0xaa, 0xb5, 0xae, 0xff, 0xa5, 0x0a, 0xe1, 0xd8, 0x33, 0xac, 0x0d, 0x55, 0x25, 0x79, 0xa5, 0x57,
	0x19, 0x04, 0xa2, 0xaa, 0x24, 0x7b, 0x0c, 0x75, 0xe5, 0x30, 0x99, 0x29, 0xc9, 0xab, 0x04, 0xd6,
	0x7c, 0x39, 0x91, 0xec, 0x10, 0x6a, 0x76, 0xb1, 0xf4, 0x78, 0x40, 0x78, 0x68, 0x17, 0xcb, 0x89,
	0x64, 0x5d, 0x68, 0x48, 0x65, 0xa3, 0x8b, 0x6b, 0x94, 0xbc, 0xd6, 0xab, 0x0c, 0x1a, 0x62, 0x53,
	0x33, 0x0e, 0x75, 0x89, 0x97, 0xd1, 0xf2, 0xda, 0xf1, 0x3a, 0x51, 0xeb, 0x92, 0x1d, 0x03, 0xc4,
	0x06, 0x23, 0x87, 0x72, 0x16, 0x39, 0xde, 0xe8, 0x55, 0x06, 0x4d, 0xd1, 0x2c, 0x90, 0x53, 0xa2,
	0x97, 0x99, 0x5c, 0xd3, 0xcd, 0x9c, 0x2e, 0x90, 0x53, 0xc7, 0xde, 0x42, 0x87, 0xc6, 0x9a, 0xc5,
	0x9b, 0xb9, 0x38, 0xf4, 0x82, 0x41, 0x6b, 0x74, 0x32, 0xdc, 0xb6, 0x67, 0x38, 0xde, 0x1e, 0x5f,
	0xec, 0xdf, 0xf3, 0xa3, 0xff, 0x35, 0x00, 0x20, 0xd1, 0x87, 0x2b, 0x34, 0xc8, 0x0e, 0x20, 0xcc,
	0xa2, 0x39, 0xe6, 0x8e, 0x84, 0x22, 0x2f, 0xd8, 0x11, 0x34, 0xfd, 0x61, 0x66, 0xd5, 0x2d, 0x92,
	0x2d, 0xa1, 0x68, 0x78, 0x60, 0xaa, 0x6e, 0x91, 0x75, 0x20, 0x70, 0x3a, 0x23, 0x57, 0x42, 0xe1,
	0x8f, 0x85, 0xa7, 0x3b, 0x1b, 0x4f, 0x3b, 0x10, 0x28, 0x69, 0x79, 0xd8, 0x0b, 0x06, 0x81, 0xf0,
	0x47, 0xef, 0xda, 0x02, 0x57, 0x9f, 0xb5, 0x91, 0x96, 0x5c, 0x6b, 0x8a, 0x4d, 0xed, 0x8d, 0x8e,
	0x23, 0xe7, 0x8d, 0xae, 0xe7, 0x6f, 0x88, 0x23, 0x37, 0x91, 0x8c, 0xc1, 0x8e, 0x5b, 0x65, 0x58,
	0x98, 0x45, 0x67, 0x6f, 0xb0, 0xd5, 0xc6, 0xa9, 0x74, 0x5e, 0x98, 0xb4, 0x2e, 0xef, 0xc6, 0x08,
	0x0f, 0xc4, 0xd8, 0xba, 0x1b, 0xe3, 0x11, 0x34, 0x49, 0x9f, 0x46, 0x09, 0xf2, 0xdd, 0xfc, 0x45,
	0x1e, 0x78, 0x1f, 0x25, 0xe8, 0xe3, 0x20, 0x32, 0x33, 0x2a, 0x46, 0xbe, 0x97, 0xc7, 0xe1, 0x91,
	0x73, 0x0f, 0xb0, 0x27, 0xd0, 0x28, 0xee, 0xb2, 0xbc, 0x4d, 0x33, 0xd6, 0xf3, 0xcb, 0x2c, 0x3b,
	0x81, 0x56, 0xbc, 0xb4, 0x4e, 0x27, 0x68, 0xfc, 0x95, 0xfb, 0x74, 0x25, 0xac, 0xa1, 0x89, 0x64,
	0x8f, 0xa0, 0x16, 0x5f, 0x2b, 0x4c, 0x1d, 0xef, 0xd0, 0x86, 0x14, 0xd5, 0xd6, 0x5a, 0xfd, 0xbf,
	0xbd, 0x56, 0xfd, 0xef, 0x15, 0xd8, 0xa3, 0xc8, 0x04, 0xda, 0x4c, 0xa7, 0x16, 0xd9, 0x73, 0xa8,
	0x61, 0xea, 0x94, 0x5b, 0x51, 0x6c, 0xad, 0xd1, 0x61, 0xe9, 0x1a, 0x88, 0x42, 0xc4, 0x9e, 0xe5,
	0x21, 0x1b, 0x8a, 0xb2, 0x44, 0x7d, 0xee, 0xc9, 0x3c, 0x7b, 0xe3, 0xc5, 0x7e, 0x1a, 0xcb, 0x03,
	0xda, 0xb0, 0x07, 0x5a, 0xe7, 0x1a, 0x2f, 0x46, 0x63, 0xb4, 0xa1, 0xf0, 0x4b, 0xc4, 0x6f, 0x3c,
	0x29, 0x72, 0x0d, 0x1b, 0xc0, 0x8e, 0x4a, 0x2f, 0x35, 0x0f, 0x49, 0x7b, 0x70, 0x5f, 0xeb, 0xff,
	0xc5, 0x82, 0x14, 0xa3, 0x8f, 0xd0, 0x7e, 0xb7, 0xa2, 0x8b, 0xa6, 0x68, 0x6e, 0xbc, 0xe7, 0x13,
	0xa8, 0x4d, 0x31, 0x32, 0xf1, 0x15, 0xeb, 0x96, 0x3e, 0x88, 0xb6, 0xb9, 0xfb, 0xb4, 0x94, 0xa3,
	0xc6, 0x85, 0x75, 0xfd, 0xff, 0x46, 0xdf, 0xaa, 0xb0, 0xfb, 0x8f, 0x7a, 0xb3, 0x31, 0x34, 0x5e,
	0xaf, 0xbf, 0x06, 0xbf, 0x6a, 0x76, 0x5c, 0xee, 0xea, 0xcf, 0x46, 0xaf, 0x20, 0x3c, 0xbb, 0xc2,
	0x78, 0xc1, 0xca, 0xfd, 0xff, 0x7d, 0x83, 0x97, 0x10, 0x8c, 0xd1, 0xfd, 0xcd, 0xcf, 0xa7, 0x7f,
	0xfe, 0xf3, 0x8b, 0x1a, 0x7d, 0x76, 0x5f, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xdf, 0x1f,
	0xd3, 0xcb, 0x05, 0x00, 0x00,
}
