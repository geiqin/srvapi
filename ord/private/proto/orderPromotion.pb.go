// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderPromotion.proto

package geiqin_srv_ord_private

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

type OrderPromotionParams struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           int64    `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	PromotionId          int64    `protobuf:"varint,3,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id,omitempty"`
	OrderId              int64    `protobuf:"varint,4,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PromotionType        string   `protobuf:"bytes,5,opt,name=promotion_type,json=promotionType,proto3" json:"promotion_type,omitempty"`
	Ids                  []int64  `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	StartDate            string   `protobuf:"bytes,7,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              string   `protobuf:"bytes,8,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderPromotionParams) Reset()         { *m = OrderPromotionParams{} }
func (m *OrderPromotionParams) String() string { return proto.CompactTextString(m) }
func (*OrderPromotionParams) ProtoMessage()    {}
func (*OrderPromotionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a828aeaf0d55488, []int{0}
}

func (m *OrderPromotionParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderPromotionParams.Unmarshal(m, b)
}
func (m *OrderPromotionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderPromotionParams.Marshal(b, m, deterministic)
}
func (m *OrderPromotionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPromotionParams.Merge(m, src)
}
func (m *OrderPromotionParams) XXX_Size() int {
	return xxx_messageInfo_OrderPromotionParams.Size(m)
}
func (m *OrderPromotionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPromotionParams.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPromotionParams proto.InternalMessageInfo

func (m *OrderPromotionParams) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderPromotionParams) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *OrderPromotionParams) GetPromotionId() int64 {
	if m != nil {
		return m.PromotionId
	}
	return 0
}

func (m *OrderPromotionParams) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderPromotionParams) GetPromotionType() string {
	if m != nil {
		return m.PromotionType
	}
	return ""
}

func (m *OrderPromotionParams) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *OrderPromotionParams) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *OrderPromotionParams) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

type OrderPromotion struct {
	Id                   int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           int64                  `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	OrderId              int64                  `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PromotionId          int64                  `protobuf:"varint,4,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id,omitempty"`
	PromotionType        string                 `protobuf:"bytes,5,opt,name=promotion_type,json=promotionType,proto3" json:"promotion_type,omitempty"`
	Price                float32                `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt            string                 `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string                 `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Goodses              []*OrderPromotionGoods `protobuf:"bytes,9,rep,name=goodses,proto3" json:"goodses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *OrderPromotion) Reset()         { *m = OrderPromotion{} }
func (m *OrderPromotion) String() string { return proto.CompactTextString(m) }
func (*OrderPromotion) ProtoMessage()    {}
func (*OrderPromotion) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a828aeaf0d55488, []int{1}
}

func (m *OrderPromotion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderPromotion.Unmarshal(m, b)
}
func (m *OrderPromotion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderPromotion.Marshal(b, m, deterministic)
}
func (m *OrderPromotion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPromotion.Merge(m, src)
}
func (m *OrderPromotion) XXX_Size() int {
	return xxx_messageInfo_OrderPromotion.Size(m)
}
func (m *OrderPromotion) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPromotion.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPromotion proto.InternalMessageInfo

func (m *OrderPromotion) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderPromotion) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *OrderPromotion) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderPromotion) GetPromotionId() int64 {
	if m != nil {
		return m.PromotionId
	}
	return 0
}

func (m *OrderPromotion) GetPromotionType() string {
	if m != nil {
		return m.PromotionType
	}
	return ""
}

func (m *OrderPromotion) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderPromotion) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *OrderPromotion) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *OrderPromotion) GetGoodses() []*OrderPromotionGoods {
	if m != nil {
		return m.Goodses
	}
	return nil
}

type OrderPromotionGoods struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderPromotionId     int64    `protobuf:"varint,2,opt,name=order_promotion_id,json=orderPromotionId,proto3" json:"order_promotion_id,omitempty"`
	ItemId               int64    `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,4,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	GoodsNum             int32    `protobuf:"varint,5,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num,omitempty"`
	Price                float32  `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderPromotionGoods) Reset()         { *m = OrderPromotionGoods{} }
func (m *OrderPromotionGoods) String() string { return proto.CompactTextString(m) }
func (*OrderPromotionGoods) ProtoMessage()    {}
func (*OrderPromotionGoods) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a828aeaf0d55488, []int{2}
}

func (m *OrderPromotionGoods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderPromotionGoods.Unmarshal(m, b)
}
func (m *OrderPromotionGoods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderPromotionGoods.Marshal(b, m, deterministic)
}
func (m *OrderPromotionGoods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPromotionGoods.Merge(m, src)
}
func (m *OrderPromotionGoods) XXX_Size() int {
	return xxx_messageInfo_OrderPromotionGoods.Size(m)
}
func (m *OrderPromotionGoods) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPromotionGoods.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPromotionGoods proto.InternalMessageInfo

func (m *OrderPromotionGoods) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderPromotionGoods) GetOrderPromotionId() int64 {
	if m != nil {
		return m.OrderPromotionId
	}
	return 0
}

func (m *OrderPromotionGoods) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *OrderPromotionGoods) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *OrderPromotionGoods) GetGoodsNum() int32 {
	if m != nil {
		return m.GoodsNum
	}
	return 0
}

func (m *OrderPromotionGoods) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderPromotionGoods) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *OrderPromotionGoods) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type OrderPromotionResponse struct {
	Entity               *OrderPromotion   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager            `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*OrderPromotion `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error            `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info             `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OrderPromotionResponse) Reset()         { *m = OrderPromotionResponse{} }
func (m *OrderPromotionResponse) String() string { return proto.CompactTextString(m) }
func (*OrderPromotionResponse) ProtoMessage()    {}
func (*OrderPromotionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a828aeaf0d55488, []int{3}
}

func (m *OrderPromotionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderPromotionResponse.Unmarshal(m, b)
}
func (m *OrderPromotionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderPromotionResponse.Marshal(b, m, deterministic)
}
func (m *OrderPromotionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPromotionResponse.Merge(m, src)
}
func (m *OrderPromotionResponse) XXX_Size() int {
	return xxx_messageInfo_OrderPromotionResponse.Size(m)
}
func (m *OrderPromotionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPromotionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPromotionResponse proto.InternalMessageInfo

func (m *OrderPromotionResponse) GetEntity() *OrderPromotion {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *OrderPromotionResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *OrderPromotionResponse) GetItems() []*OrderPromotion {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *OrderPromotionResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *OrderPromotionResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderPromotionParams)(nil), "geiqin.srv.ord.private.OrderPromotionParams")
	proto.RegisterType((*OrderPromotion)(nil), "geiqin.srv.ord.private.OrderPromotion")
	proto.RegisterType((*OrderPromotionGoods)(nil), "geiqin.srv.ord.private.OrderPromotionGoods")
	proto.RegisterType((*OrderPromotionResponse)(nil), "geiqin.srv.ord.private.OrderPromotionResponse")
}

func init() {
	proto.RegisterFile("orderPromotion.proto", fileDescriptor_9a828aeaf0d55488)
}

var fileDescriptor_9a828aeaf0d55488 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xe1, 0x6b, 0xd3, 0x40,
	0x18, 0xc6, 0x4d, 0xd2, 0xa4, 0xeb, 0x9b, 0x59, 0xc6, 0xd9, 0xcd, 0x38, 0x1d, 0xd6, 0x82, 0x52,
	0x70, 0x04, 0xe9, 0xbe, 0x8a, 0x30, 0xdc, 0x18, 0x01, 0xd1, 0x92, 0xe9, 0xe7, 0x92, 0xf5, 0xde,
	0x76, 0xc7, 0x4c, 0x2e, 0xde, 0x5d, 0x0a, 0xfd, 0xb7, 0xfc, 0x0f, 0xfc, 0x9f, 0x44, 0xfc, 0x28,
	0x77, 0x69, 0x56, 0xe3, 0x36, 0x89, 0x42, 0xbf, 0x35, 0xef, 0xfb, 0x3c, 0x7d, 0xee, 0xf9, 0x5d,
	0x20, 0xd0, 0xe3, 0x82, 0xa2, 0x18, 0x0b, 0x9e, 0x72, 0xc5, 0x78, 0x16, 0xe6, 0x82, 0x2b, 0x4e,
	0xf6, 0xe6, 0xc8, 0xbe, 0xb0, 0x2c, 0x94, 0x62, 0x11, 0x72, 0x41, 0xc3, 0x5c, 0xb0, 0x45, 0xa2,
	0x70, 0x7f, 0x7b, 0xca, 0xd3, 0xb4, 0x52, 0x0d, 0x7e, 0x5a, 0xd0, 0xfb, 0x50, 0xb3, 0x8f, 0x13,
	0x91, 0xa4, 0x92, 0x74, 0xc1, 0x66, 0x34, 0xb0, 0xfa, 0xd6, 0xd0, 0x89, 0x6d, 0x46, 0xc9, 0x53,
	0xf0, 0xa7, 0x85, 0x54, 0x3c, 0x45, 0x31, 0x61, 0x34, 0xb0, 0xcd, 0x02, 0xaa, 0x51, 0x44, 0xc9,
	0x33, 0xd8, 0xce, 0xab, 0xff, 0xd0, 0x0a, 0xc7, 0x28, 0xfc, 0xeb, 0x59, 0x44, 0xc9, 0x23, 0xd8,
	0x32, 0x47, 0xd5, 0xeb, 0x96, 0x59, 0xb7, 0xcd, 0x73, 0x44, 0xc9, 0x73, 0xe8, 0xae, 0xdd, 0x6a,
	0x99, 0x63, 0xe0, 0xf6, 0xad, 0x61, 0x27, 0xbe, 0x7f, 0x3d, 0xfd, 0xb8, 0xcc, 0x91, 0xec, 0x80,
	0xc3, 0xa8, 0x0c, 0xbc, 0xbe, 0x33, 0x74, 0x62, 0xfd, 0x93, 0x1c, 0x00, 0x48, 0x95, 0x08, 0x35,
	0xa1, 0x89, 0xc2, 0xa0, 0x6d, 0x4c, 0x1d, 0x33, 0x39, 0x49, 0x14, 0xea, 0x48, 0xcc, 0x68, 0xb9,
	0xdc, 0x32, 0xcb, 0x36, 0x66, 0x54, 0xaf, 0x06, 0xdf, 0x6c, 0xe8, 0xd6, 0xab, 0xff, 0x7b, 0xe9,
	0xdf, 0x1b, 0x39, 0xf5, 0x46, 0x7f, 0xf2, 0x68, 0xdd, 0xe4, 0xd1, 0xb0, 0x74, 0x0f, 0xdc, 0x5c,
	0xb0, 0x29, 0x06, 0x5e, 0xdf, 0x1a, 0xda, 0x71, 0xf9, 0xa0, 0x8b, 0x4f, 0x05, 0x26, 0x0a, 0xe9,
	0x24, 0x51, 0x55, 0xf1, 0xd5, 0xe4, 0x58, 0xe9, 0x75, 0x91, 0xd3, 0x6a, 0x5d, 0x56, 0xef, 0xac,
	0x26, 0xc7, 0x8a, 0x9c, 0x42, 0x7b, 0xce, 0x39, 0x95, 0x28, 0x83, 0x4e, 0xdf, 0x19, 0xfa, 0xa3,
	0x97, 0xe1, 0xed, 0xef, 0x4b, 0x58, 0x47, 0x74, 0xa6, 0x4d, 0x71, 0xe5, 0x1d, 0x7c, 0xb7, 0xe0,
	0xc1, 0x2d, 0x82, 0x1b, 0x20, 0x0f, 0x81, 0x94, 0x9c, 0x6a, 0x48, 0x4a, 0x9e, 0x3b, 0xf5, 0xd7,
	0x37, 0xa2, 0xe4, 0x21, 0xb4, 0x99, 0xc2, 0x74, 0x0d, 0xd5, 0xd3, 0x8f, 0x11, 0x25, 0xbb, 0xe0,
	0xc9, 0xab, 0x62, 0x4d, 0xd3, 0x95, 0x57, 0x45, 0x44, 0xc9, 0x63, 0xe8, 0x98, 0x03, 0x4d, 0xb2,
	0x22, 0x35, 0x08, 0xdd, 0x78, 0xcb, 0x0c, 0xde, 0x17, 0xe9, 0x26, 0xe8, 0x0d, 0xbe, 0xda, 0xb0,
	0x57, 0xaf, 0x1d, 0xa3, 0xcc, 0x79, 0x26, 0x91, 0xbc, 0x01, 0x0f, 0x33, 0xc5, 0xd4, 0xd2, 0xb4,
	0xf7, 0x47, 0x2f, 0x9a, 0x71, 0x8d, 0x57, 0x2e, 0x72, 0x04, 0x6e, 0x9e, 0xcc, 0x51, 0x18, 0x38,
	0xfe, 0xe8, 0xe0, 0x2e, 0xfb, 0x58, 0x8b, 0xe2, 0x52, 0x4b, 0x5e, 0x83, 0xab, 0x09, 0xc9, 0xc0,
	0x31, 0x77, 0xd9, 0x34, 0xb3, 0x34, 0xe9, 0x48, 0x14, 0x82, 0x0b, 0x03, 0xf5, 0x2f, 0x91, 0xa7,
	0x5a, 0x14, 0x97, 0x5a, 0xf2, 0x0a, 0x5a, 0x2c, 0x9b, 0x71, 0x83, 0xdb, 0x1f, 0x3d, 0xb9, 0xcb,
	0x13, 0x65, 0x33, 0x1e, 0x1b, 0xe5, 0xe8, 0x47, 0x0b, 0x76, 0xeb, 0x07, 0x38, 0x47, 0xb1, 0xd0,
	0x97, 0x71, 0x01, 0xde, 0x5b, 0x83, 0x9e, 0x34, 0x3c, 0xf9, 0x7e, 0xd8, 0xb0, 0xe1, 0xea, 0x56,
	0x06, 0xf7, 0x74, 0xc6, 0x27, 0x73, 0x7f, 0x9b, 0xcd, 0x38, 0xc1, 0xcf, 0xb8, 0xd1, 0x8c, 0x09,
	0x38, 0x67, 0xa8, 0x36, 0x18, 0x70, 0x09, 0xde, 0x39, 0x26, 0x62, 0x7a, 0x49, 0x0e, 0x9b, 0x79,
	0xcb, 0x0f, 0xc6, 0x7f, 0x24, 0xcd, 0xa0, 0xf5, 0x8e, 0x49, 0xb5, 0xe9, 0x9c, 0x0b, 0xcf, 0x7c,
	0xea, 0x8e, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x26, 0xd3, 0x87, 0x28, 0x07, 0x00, 0x00,
}
