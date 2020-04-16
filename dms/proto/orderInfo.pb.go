// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderInfo.proto

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

type OrderInfo struct {
	Id                   int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId              int64              `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderSn              string             `protobuf:"bytes,3,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
	PlatformSource       string             `protobuf:"bytes,4,opt,name=platform_source,json=platformSource,proto3" json:"platform_source,omitempty"`
	OrderAmount          float32            `protobuf:"fixed32,5,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	Money                float32            `protobuf:"fixed32,7,opt,name=money,proto3" json:"money,omitempty"`
	IncomeType           string             `protobuf:"bytes,8,opt,name=income_type,json=incomeType,proto3" json:"income_type,omitempty"`
	IncomeRate           float32            `protobuf:"fixed32,9,opt,name=income_rate,json=incomeRate,proto3" json:"income_rate,omitempty"`
	Status               int32              `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	BuyerName            string             `protobuf:"bytes,12,opt,name=buyer_name,json=buyerName,proto3" json:"buyer_name,omitempty"`
	Mobile               string             `protobuf:"bytes,13,opt,name=mobile,proto3" json:"mobile,omitempty"`
	OrderStatus          string             `protobuf:"bytes,14,opt,name=order_status,json=orderStatus,proto3" json:"order_status,omitempty"`
	TotalNum             int32              `protobuf:"varint,15,opt,name=total_num,json=totalNum,proto3" json:"total_num,omitempty"`
	CanDelivery          bool               `protobuf:"varint,16,opt,name=can_delivery,json=canDelivery,proto3" json:"can_delivery,omitempty"`
	CanLogistics         bool               `protobuf:"varint,17,opt,name=can_logistics,json=canLogistics,proto3" json:"can_logistics,omitempty"`
	OrderCreatedAt       string             `protobuf:"bytes,18,opt,name=order_created_at,json=orderCreatedAt,proto3" json:"order_created_at,omitempty"`
	OrderDetails         []*OrderDetailInfo `protobuf:"bytes,19,rep,name=order_details,json=orderDetails,proto3" json:"order_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OrderInfo) Reset()         { *m = OrderInfo{} }
func (m *OrderInfo) String() string { return proto.CompactTextString(m) }
func (*OrderInfo) ProtoMessage()    {}
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e838811beff491e, []int{0}
}

func (m *OrderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderInfo.Unmarshal(m, b)
}
func (m *OrderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderInfo.Marshal(b, m, deterministic)
}
func (m *OrderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderInfo.Merge(m, src)
}
func (m *OrderInfo) XXX_Size() int {
	return xxx_messageInfo_OrderInfo.Size(m)
}
func (m *OrderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderInfo proto.InternalMessageInfo

func (m *OrderInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderInfo) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderInfo) GetOrderSn() string {
	if m != nil {
		return m.OrderSn
	}
	return ""
}

func (m *OrderInfo) GetPlatformSource() string {
	if m != nil {
		return m.PlatformSource
	}
	return ""
}

func (m *OrderInfo) GetOrderAmount() float32 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

func (m *OrderInfo) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *OrderInfo) GetIncomeType() string {
	if m != nil {
		return m.IncomeType
	}
	return ""
}

func (m *OrderInfo) GetIncomeRate() float32 {
	if m != nil {
		return m.IncomeRate
	}
	return 0
}

func (m *OrderInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *OrderInfo) GetBuyerName() string {
	if m != nil {
		return m.BuyerName
	}
	return ""
}

func (m *OrderInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *OrderInfo) GetOrderStatus() string {
	if m != nil {
		return m.OrderStatus
	}
	return ""
}

func (m *OrderInfo) GetTotalNum() int32 {
	if m != nil {
		return m.TotalNum
	}
	return 0
}

func (m *OrderInfo) GetCanDelivery() bool {
	if m != nil {
		return m.CanDelivery
	}
	return false
}

func (m *OrderInfo) GetCanLogistics() bool {
	if m != nil {
		return m.CanLogistics
	}
	return false
}

func (m *OrderInfo) GetOrderCreatedAt() string {
	if m != nil {
		return m.OrderCreatedAt
	}
	return ""
}

func (m *OrderInfo) GetOrderDetails() []*OrderDetailInfo {
	if m != nil {
		return m.OrderDetails
	}
	return nil
}

type OrderDetailInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId              int64    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ItemId               int64    `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ThumbUrl             string   `protobuf:"bytes,5,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	SkuId                int64    `protobuf:"varint,6,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	ItemSn               string   `protobuf:"bytes,7,opt,name=item_sn,json=itemSn,proto3" json:"item_sn,omitempty"`
	SkuSn                string   `protobuf:"bytes,8,opt,name=sku_sn,json=skuSn,proto3" json:"sku_sn,omitempty"`
	ModelType            string   `protobuf:"bytes,9,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	Name                 string   `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Num                  int32    `protobuf:"varint,12,opt,name=num,proto3" json:"num,omitempty"`
	Price                float32  `protobuf:"fixed32,13,opt,name=price,proto3" json:"price,omitempty"`
	TotalPrice           float32  `protobuf:"fixed32,14,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderDetailInfo) Reset()         { *m = OrderDetailInfo{} }
func (m *OrderDetailInfo) String() string { return proto.CompactTextString(m) }
func (*OrderDetailInfo) ProtoMessage()    {}
func (*OrderDetailInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e838811beff491e, []int{1}
}

func (m *OrderDetailInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderDetailInfo.Unmarshal(m, b)
}
func (m *OrderDetailInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderDetailInfo.Marshal(b, m, deterministic)
}
func (m *OrderDetailInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderDetailInfo.Merge(m, src)
}
func (m *OrderDetailInfo) XXX_Size() int {
	return xxx_messageInfo_OrderDetailInfo.Size(m)
}
func (m *OrderDetailInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderDetailInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderDetailInfo proto.InternalMessageInfo

func (m *OrderDetailInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderDetailInfo) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderDetailInfo) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *OrderDetailInfo) GetThumbUrl() string {
	if m != nil {
		return m.ThumbUrl
	}
	return ""
}

func (m *OrderDetailInfo) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *OrderDetailInfo) GetItemSn() string {
	if m != nil {
		return m.ItemSn
	}
	return ""
}

func (m *OrderDetailInfo) GetSkuSn() string {
	if m != nil {
		return m.SkuSn
	}
	return ""
}

func (m *OrderDetailInfo) GetModelType() string {
	if m != nil {
		return m.ModelType
	}
	return ""
}

func (m *OrderDetailInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrderDetailInfo) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *OrderDetailInfo) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderDetailInfo) GetTotalPrice() float32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

type OrderInfoWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Mobile               string   `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	OrderId              int64    `protobuf:"varint,5,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderSn              string   `protobuf:"bytes,6,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
	Keywords             string   `protobuf:"bytes,7,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Ids                  []int64  `protobuf:"varint,9,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	DistributorId        int64    `protobuf:"varint,10,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	CustomerId           int64    `protobuf:"varint,11,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderInfoWhere) Reset()         { *m = OrderInfoWhere{} }
func (m *OrderInfoWhere) String() string { return proto.CompactTextString(m) }
func (*OrderInfoWhere) ProtoMessage()    {}
func (*OrderInfoWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e838811beff491e, []int{2}
}

func (m *OrderInfoWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderInfoWhere.Unmarshal(m, b)
}
func (m *OrderInfoWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderInfoWhere.Marshal(b, m, deterministic)
}
func (m *OrderInfoWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderInfoWhere.Merge(m, src)
}
func (m *OrderInfoWhere) XXX_Size() int {
	return xxx_messageInfo_OrderInfoWhere.Size(m)
}
func (m *OrderInfoWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderInfoWhere.DiscardUnknown(m)
}

var xxx_messageInfo_OrderInfoWhere proto.InternalMessageInfo

func (m *OrderInfoWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *OrderInfoWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *OrderInfoWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderInfoWhere) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *OrderInfoWhere) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderInfoWhere) GetOrderSn() string {
	if m != nil {
		return m.OrderSn
	}
	return ""
}

func (m *OrderInfoWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *OrderInfoWhere) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *OrderInfoWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *OrderInfoWhere) GetDistributorId() int64 {
	if m != nil {
		return m.DistributorId
	}
	return 0
}

func (m *OrderInfoWhere) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

type OrderInfoResponse struct {
	Entity               *OrderInfo   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*OrderInfo `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OrderInfoResponse) Reset()         { *m = OrderInfoResponse{} }
func (m *OrderInfoResponse) String() string { return proto.CompactTextString(m) }
func (*OrderInfoResponse) ProtoMessage()    {}
func (*OrderInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e838811beff491e, []int{3}
}

func (m *OrderInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderInfoResponse.Unmarshal(m, b)
}
func (m *OrderInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderInfoResponse.Marshal(b, m, deterministic)
}
func (m *OrderInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderInfoResponse.Merge(m, src)
}
func (m *OrderInfoResponse) XXX_Size() int {
	return xxx_messageInfo_OrderInfoResponse.Size(m)
}
func (m *OrderInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderInfoResponse proto.InternalMessageInfo

func (m *OrderInfoResponse) GetEntity() *OrderInfo {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *OrderInfoResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *OrderInfoResponse) GetItems() []*OrderInfo {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *OrderInfoResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *OrderInfoResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderInfo)(nil), "geiqin.srv.dms.OrderInfo")
	proto.RegisterType((*OrderDetailInfo)(nil), "geiqin.srv.dms.OrderDetailInfo")
	proto.RegisterType((*OrderInfoWhere)(nil), "geiqin.srv.dms.OrderInfoWhere")
	proto.RegisterType((*OrderInfoResponse)(nil), "geiqin.srv.dms.OrderInfoResponse")
}

func init() {
	proto.RegisterFile("orderInfo.proto", fileDescriptor_2e838811beff491e)
}

var fileDescriptor_2e838811beff491e = []byte{
	// 798 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xd1, 0x6e, 0xe3, 0x44,
	0x14, 0xc5, 0x71, 0x92, 0xc6, 0x37, 0x69, 0xda, 0x1d, 0x76, 0x61, 0xb6, 0x08, 0x36, 0x1b, 0x84,
	0xb0, 0x84, 0x14, 0x44, 0xf8, 0x82, 0x15, 0x45, 0xa8, 0x88, 0x5d, 0x56, 0x36, 0x88, 0x47, 0xcb,
	0xb1, 0x6f, 0xdb, 0x51, 0x3d, 0x33, 0x61, 0x66, 0x5c, 0xe4, 0xfd, 0x23, 0xc4, 0x03, 0x4f, 0xbc,
	0xf1, 0x67, 0xbc, 0xa0, 0xb9, 0xe3, 0x74, 0xd3, 0xaa, 0x45, 0x42, 0xe2, 0x85, 0xb7, 0x99, 0x73,
	0xcf, 0x3d, 0xb9, 0x3e, 0xf7, 0x8c, 0x02, 0x47, 0xda, 0xd4, 0x68, 0xce, 0xd4, 0xb9, 0x5e, 0x6d,
	0x8d, 0x76, 0x9a, 0xcd, 0x2f, 0x50, 0xfc, 0x2c, 0xd4, 0xca, 0x9a, 0xeb, 0x55, 0x2d, 0xed, 0xc9,
	0xac, 0xd2, 0x52, 0x6a, 0x15, 0xaa, 0xcb, 0x3f, 0x87, 0x90, 0x7c, 0xbf, 0xeb, 0x60, 0x73, 0x18,
	0x88, 0x9a, 0x47, 0x8b, 0x28, 0x8d, 0xb3, 0x81, 0xa8, 0xd9, 0x53, 0x98, 0x90, 0x5c, 0x21, 0x6a,
	0x3e, 0x20, 0xf4, 0x20, 0xc8, 0xef, 0x95, 0xac, 0xe2, 0xf1, 0x22, 0x4a, 0x93, 0xbe, 0x94, 0x2b,
	0xf6, 0x29, 0x1c, 0x6d, 0x9b, 0xd2, 0x9d, 0x6b, 0x23, 0x0b, 0xab, 0x5b, 0x53, 0x21, 0x1f, 0x12,
	0x63, 0xbe, 0x83, 0x73, 0x42, 0xd9, 0x73, 0x98, 0x05, 0x8d, 0x52, 0xea, 0x56, 0x39, 0x3e, 0x5a,
	0x44, 0xe9, 0x20, 0x9b, 0x12, 0xf6, 0x82, 0x20, 0xf6, 0x18, 0x46, 0x52, 0x2b, 0xec, 0xf8, 0x01,
	0xd5, 0xc2, 0x85, 0x3d, 0x83, 0xa9, 0x50, 0x95, 0x96, 0x58, 0xb8, 0x6e, 0x8b, 0x7c, 0x42, 0xea,
	0x10, 0xa0, 0x1f, 0xba, 0x2d, 0xee, 0x11, 0x4c, 0xe9, 0x90, 0x27, 0xd4, 0xdc, 0x13, 0xb2, 0xd2,
	0x21, 0x7b, 0x0f, 0xc6, 0xd6, 0x95, 0xae, 0xb5, 0x1c, 0x16, 0x51, 0x3a, 0xca, 0xfa, 0x1b, 0xfb,
	0x10, 0x60, 0xd3, 0x76, 0x68, 0x0a, 0x55, 0x4a, 0xe4, 0x33, 0x12, 0x4e, 0x08, 0x79, 0x55, 0x4a,
	0x6a, 0x93, 0x7a, 0x23, 0x1a, 0xe4, 0x87, 0x54, 0xea, 0x6f, 0x6f, 0xbf, 0xa4, 0x17, 0x9d, 0x53,
	0x35, 0x7c, 0x49, 0x1e, 0x94, 0x3f, 0x80, 0xc4, 0x69, 0x57, 0x36, 0x85, 0x6a, 0x25, 0x3f, 0xa2,
	0x1f, 0x9d, 0x10, 0xf0, 0xaa, 0x95, 0xbe, 0xbf, 0x2a, 0x55, 0x51, 0x63, 0x23, 0xae, 0xd1, 0x74,
	0xfc, 0x78, 0x11, 0xa5, 0x93, 0x6c, 0x5a, 0x95, 0xea, 0xb4, 0x87, 0xd8, 0xc7, 0x70, 0xe8, 0x29,
	0x8d, 0xbe, 0x10, 0xd6, 0x89, 0xca, 0xf2, 0x47, 0xc4, 0xf1, 0x7d, 0xdf, 0xed, 0x30, 0x96, 0xc2,
	0x71, 0x98, 0xa3, 0x32, 0x58, 0x3a, 0xac, 0x8b, 0xd2, 0x71, 0x16, 0xbc, 0x27, 0xfc, 0xab, 0x00,
	0xbf, 0x70, 0xec, 0x14, 0x0e, 0x03, 0xb3, 0x46, 0x57, 0x8a, 0xc6, 0xf2, 0x77, 0x17, 0x71, 0x3a,
	0x5d, 0x3f, 0x5b, 0xdd, 0x8e, 0xcb, 0x8a, 0xc2, 0x71, 0x4a, 0x1c, 0x1f, 0x91, 0x2c, 0x7c, 0x67,
	0x00, 0xec, 0xf2, 0x8f, 0x01, 0x1c, 0xdd, 0x61, 0xfc, 0x9b, 0x10, 0xbd, 0x0f, 0x07, 0xc2, 0xa1,
	0xf4, 0x95, 0x98, 0x2a, 0x63, 0x7f, 0x3d, 0xab, 0xc9, 0xac, 0xcb, 0x56, 0x6e, 0x8a, 0xd6, 0x34,
	0x14, 0x8b, 0x24, 0x9b, 0x10, 0xf0, 0xa3, 0x69, 0xd8, 0x13, 0x18, 0xdb, 0xab, 0xd6, 0x37, 0x8d,
	0xa9, 0x69, 0x64, 0xaf, 0xda, 0x3d, 0x31, 0xab, 0x28, 0x2c, 0x49, 0x10, 0xcb, 0xd5, 0x8e, 0x6f,
	0x55, 0x1f, 0x14, 0xcf, 0xcf, 0x95, 0x5f, 0xb5, 0xd4, 0x35, 0x36, 0x21, 0x43, 0x49, 0x58, 0x35,
	0x21, 0x14, 0x21, 0x06, 0x43, 0xca, 0x00, 0x50, 0x81, 0xce, 0xec, 0x18, 0x62, 0xbf, 0xbd, 0x19,
	0x6d, 0xcf, 0x1f, 0x7d, 0x3e, 0xb7, 0x46, 0x54, 0x21, 0x0f, 0x83, 0x2c, 0x5c, 0x7c, 0xfc, 0xc2,
	0xae, 0x43, 0x6d, 0x1e, 0xe2, 0x47, 0xd0, 0x6b, 0x8f, 0x2c, 0x7f, 0x1f, 0xc0, 0xfc, 0xe6, 0xd9,
	0xfd, 0x74, 0x89, 0x06, 0x49, 0xa9, 0xbc, 0xc0, 0xe0, 0xdc, 0x28, 0x0b, 0x17, 0x6f, 0x84, 0x3f,
	0x14, 0x56, 0xbc, 0x41, 0x72, 0x6f, 0x94, 0x4d, 0x3c, 0x90, 0x8b, 0x37, 0xd8, 0x3b, 0x1d, 0xdf,
	0x38, 0xfd, 0x36, 0x9d, 0xc3, 0x5b, 0xe9, 0xdc, 0xdf, 0xc0, 0xe8, 0xe1, 0x67, 0x3c, 0xbe, 0xfd,
	0x8c, 0x4f, 0x60, 0x72, 0x85, 0xdd, 0x2f, 0xda, 0xd4, 0xb6, 0x37, 0xf4, 0xe6, 0xbe, 0xf7, 0x7c,
	0x82, 0xa5, 0xbb, 0xe7, 0x73, 0x0c, 0xb1, 0xa8, 0x2d, 0x4f, 0x16, 0x71, 0x1a, 0x67, 0xfe, 0xc8,
	0x3e, 0x81, 0x79, 0x2d, 0xac, 0x33, 0x62, 0xd3, 0x3a, 0x4d, 0x13, 0x00, 0x4d, 0x70, 0xb8, 0x87,
	0x9e, 0xd5, 0xde, 0xb1, 0xaa, 0xb5, 0x4e, 0xcb, 0x30, 0xe5, 0x94, 0x38, 0xb0, 0x83, 0xce, 0xea,
	0xe5, 0x5f, 0x11, 0x3c, 0xba, 0x71, 0x2c, 0x43, 0xbb, 0xd5, 0xca, 0x22, 0xfb, 0x02, 0xc6, 0xa8,
	0x9c, 0x70, 0x1d, 0xb9, 0x36, 0x5d, 0x3f, 0xbd, 0x37, 0xbe, 0xd4, 0xd2, 0x13, 0xd9, 0x67, 0xc1,
	0x67, 0x43, 0x6e, 0x4e, 0xd7, 0x4f, 0xee, 0x76, 0xbc, 0xf6, 0xc5, 0x60, 0xbf, 0x61, 0x9f, 0xc3,
	0xc8, 0x87, 0xc8, 0xf2, 0x98, 0x5e, 0xc7, 0x3f, 0xc8, 0x07, 0x9e, 0x57, 0x47, 0x63, 0xb4, 0xa1,
	0x0d, 0xdc, 0xa3, 0xfe, 0xb5, 0x2f, 0x66, 0x81, 0xc3, 0x52, 0x18, 0x0a, 0x75, 0xae, 0x69, 0x27,
	0xd3, 0xf5, 0xe3, 0xbb, 0x5c, 0xd2, 0x25, 0xc6, 0xfa, 0xb7, 0x08, 0xe6, 0x2f, 0x3b, 0xfa, 0xb5,
	0x1c, 0xcd, 0xb5, 0xcf, 0xd8, 0xb7, 0x10, 0x7f, 0x83, 0x8e, 0x7d, 0xf4, 0xe0, 0x48, 0x14, 0xab,
	0x93, 0xe7, 0x0f, 0x8f, 0xdc, 0x9b, 0xb8, 0x7c, 0x87, 0xbd, 0x84, 0x71, 0x8e, 0xa5, 0xa9, 0x2e,
	0xff, 0x13, 0xb9, 0xf5, 0xaf, 0x11, 0xcc, 0xfe, 0x27, 0xb3, 0x6e, 0xc6, 0xf4, 0x3f, 0xf8, 0xe5,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x7d, 0x5f, 0x27, 0x38, 0x07, 0x00, 0x00,
}