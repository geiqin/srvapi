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
	LeaderId             int64              `protobuf:"varint,20,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	LeaderName           string             `protobuf:"bytes,21,opt,name=leader_name,json=leaderName,proto3" json:"leader_name,omitempty"`
	LeaderMobile         string             `protobuf:"bytes,22,opt,name=leader_mobile,json=leaderMobile,proto3" json:"leader_mobile,omitempty"`
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

func (m *OrderInfo) GetLeaderId() int64 {
	if m != nil {
		return m.LeaderId
	}
	return 0
}

func (m *OrderInfo) GetLeaderName() string {
	if m != nil {
		return m.LeaderName
	}
	return ""
}

func (m *OrderInfo) GetLeaderMobile() string {
	if m != nil {
		return m.LeaderMobile
	}
	return ""
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
	LeaderId             int64    `protobuf:"varint,12,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
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

func (m *OrderInfoWhere) GetLeaderId() int64 {
	if m != nil {
		return m.LeaderId
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
	// 865 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6e, 0xe4, 0x44,
	0x13, 0xfd, 0x3c, 0x8e, 0x27, 0xe3, 0x9a, 0x9f, 0x24, 0xbd, 0xc9, 0x7e, 0xbd, 0x41, 0xb0, 0xb3,
	0xb3, 0x42, 0x58, 0x42, 0x1a, 0xc4, 0xf0, 0x04, 0x0b, 0x41, 0x28, 0x88, 0x2c, 0x2b, 0x0f, 0x88,
	0x4b, 0xcb, 0x63, 0x77, 0x92, 0x56, 0xec, 0xee, 0xa1, 0xbb, 0x1d, 0xe4, 0x7d, 0x19, 0xae, 0x11,
	0x12, 0xe2, 0x86, 0xb7, 0xe0, 0x8d, 0xb8, 0x41, 0x5d, 0xe5, 0x49, 0x26, 0xd1, 0x06, 0x09, 0x89,
	0x9b, 0xdc, 0x75, 0x9f, 0x3a, 0x55, 0x53, 0x7d, 0xea, 0x94, 0x35, 0xb0, 0xa7, 0x4d, 0x29, 0xcc,
	0xa9, 0x3a, 0xd7, 0xf3, 0xb5, 0xd1, 0x4e, 0xb3, 0xc9, 0x85, 0x90, 0x3f, 0x4a, 0x35, 0xb7, 0xe6,
	0x7a, 0x5e, 0xd6, 0xf6, 0x78, 0x54, 0xe8, 0xba, 0xd6, 0x8a, 0xa2, 0xb3, 0x9f, 0x23, 0x88, 0xbf,
	0xdd, 0x64, 0xb0, 0x09, 0xf4, 0x64, 0xc9, 0x83, 0x69, 0x90, 0x84, 0x69, 0x4f, 0x96, 0xec, 0x19,
	0x0c, 0xb0, 0x5c, 0x26, 0x4b, 0xde, 0x43, 0x74, 0x97, 0xca, 0x6f, 0x85, 0xac, 0xe2, 0xe1, 0x34,
	0x48, 0xe2, 0x2e, 0xb4, 0x54, 0xec, 0x23, 0xd8, 0x5b, 0x57, 0xb9, 0x3b, 0xd7, 0xa6, 0xce, 0xac,
	0x6e, 0x4c, 0x21, 0xf8, 0x0e, 0x32, 0x26, 0x1b, 0x78, 0x89, 0x28, 0x7b, 0x01, 0x23, 0xaa, 0x91,
	0xd7, 0xba, 0x51, 0x8e, 0x47, 0xd3, 0x20, 0xe9, 0xa5, 0x43, 0xc4, 0x5e, 0x21, 0xc4, 0x0e, 0x21,
	0xaa, 0xb5, 0x12, 0x2d, 0xdf, 0xc5, 0x18, 0x5d, 0xd8, 0x73, 0x18, 0x4a, 0x55, 0xe8, 0x5a, 0x64,
	0xae, 0x5d, 0x0b, 0x3e, 0xc0, 0xea, 0x40, 0xd0, 0x77, 0xed, 0x5a, 0x6c, 0x11, 0x4c, 0xee, 0x04,
	0x8f, 0x31, 0xb9, 0x23, 0xa4, 0xb9, 0x13, 0xec, 0x29, 0xf4, 0xad, 0xcb, 0x5d, 0x63, 0x39, 0x4c,
	0x83, 0x24, 0x4a, 0xbb, 0x1b, 0x7b, 0x1f, 0x60, 0xd5, 0xb4, 0xc2, 0x64, 0x2a, 0xaf, 0x05, 0x1f,
	0x61, 0xe1, 0x18, 0x91, 0xd7, 0x79, 0x8d, 0x69, 0xb5, 0x5e, 0xc9, 0x4a, 0xf0, 0x31, 0x86, 0xba,
	0xdb, 0xed, 0x4b, 0xba, 0xa2, 0x13, 0x8c, 0xd2, 0x4b, 0x96, 0x54, 0xf9, 0x3d, 0x88, 0x9d, 0x76,
	0x79, 0x95, 0xa9, 0xa6, 0xe6, 0x7b, 0xf8, 0xa3, 0x03, 0x04, 0x5e, 0x37, 0xb5, 0xcf, 0x2f, 0x72,
	0x95, 0x95, 0xa2, 0x92, 0xd7, 0xc2, 0xb4, 0x7c, 0x7f, 0x1a, 0x24, 0x83, 0x74, 0x58, 0xe4, 0xea,
	0xa4, 0x83, 0xd8, 0x4b, 0x18, 0x7b, 0x4a, 0xa5, 0x2f, 0xa4, 0x75, 0xb2, 0xb0, 0xfc, 0x00, 0x39,
	0x3e, 0xef, 0x9b, 0x0d, 0xc6, 0x12, 0xd8, 0xa7, 0x3e, 0x0a, 0x23, 0x72, 0x27, 0xca, 0x2c, 0x77,
	0x9c, 0x91, 0xf6, 0x88, 0x7f, 0x41, 0xf0, 0x2b, 0xc7, 0x4e, 0x60, 0x4c, 0xcc, 0x52, 0xb8, 0x5c,
	0x56, 0x96, 0x3f, 0x99, 0x86, 0xc9, 0x70, 0xf1, 0x7c, 0x7e, 0xd7, 0x2e, 0x73, 0x34, 0xc7, 0x09,
	0x72, 0xbc, 0x45, 0x52, 0x7a, 0x27, 0x01, 0xf8, 0xa8, 0x4a, 0xe4, 0x9d, 0x43, 0x0e, 0xd1, 0x21,
	0x03, 0x02, 0x4e, 0x4b, 0x3f, 0x84, 0x2e, 0x88, 0x62, 0x1e, 0xd1, 0x94, 0x08, 0x42, 0x35, 0x5f,
	0xc2, 0xb8, 0x23, 0x74, 0xa2, 0x3e, 0x45, 0xca, 0x88, 0xc0, 0x33, 0xc4, 0x66, 0x7f, 0xf4, 0x60,
	0xef, 0x5e, 0x13, 0xff, 0xc6, 0xa7, 0xff, 0x87, 0x5d, 0xe9, 0x44, 0xed, 0x23, 0x21, 0x46, 0xfa,
	0xfe, 0x7a, 0x5a, 0xe2, 0x3c, 0x2e, 0x9b, 0x7a, 0x95, 0x35, 0xa6, 0x42, 0xe7, 0xc5, 0xe9, 0x00,
	0x81, 0xef, 0x4d, 0xc5, 0x8e, 0xa0, 0x6f, 0xaf, 0x1a, 0x9f, 0xd4, 0xc7, 0xa4, 0xc8, 0x5e, 0x35,
	0x5b, 0xc5, 0xac, 0x42, 0x3f, 0xc6, 0x54, 0x6c, 0xa9, 0x36, 0x7c, 0xab, 0x3a, 0x2f, 0x7a, 0xfe,
	0x52, 0x79, 0x37, 0xd5, 0xba, 0x14, 0x15, 0xd9, 0x34, 0x26, 0x37, 0x21, 0x82, 0x2e, 0x65, 0xb0,
	0x83, 0xca, 0x00, 0x06, 0xf0, 0xcc, 0xf6, 0x21, 0xf4, 0x06, 0x19, 0xa1, 0x41, 0xfc, 0xd1, 0xaf,
	0xc0, 0xda, 0xc8, 0x82, 0x2c, 0xd7, 0x4b, 0xe9, 0xe2, 0xc5, 0x25, 0x3b, 0x51, 0x6c, 0x42, 0x0e,
	0x47, 0xe8, 0x8d, 0x47, 0x66, 0x7f, 0xf6, 0x60, 0x72, 0xb3, 0xd9, 0x3f, 0x5c, 0x0a, 0x23, 0xb0,
	0x52, 0x7e, 0x21, 0x48, 0xb9, 0x28, 0xa5, 0x8b, 0x17, 0xc2, 0x1f, 0x32, 0x2b, 0xdf, 0x0a, 0x54,
	0x2f, 0x4a, 0x07, 0x1e, 0x58, 0xca, 0xb7, 0xa2, 0x53, 0x3a, 0xbc, 0x51, 0xfa, 0x76, 0x01, 0x76,
	0xee, 0x2c, 0xc0, 0xf6, 0x04, 0xa2, 0x87, 0xbf, 0x14, 0xfd, 0xbb, 0x5f, 0x8a, 0x63, 0x18, 0x5c,
	0x89, 0xf6, 0x27, 0x6d, 0x4a, 0xdb, 0x09, 0x7a, 0x73, 0xdf, 0xda, 0x50, 0x92, 0x74, 0xb3, 0xa1,
	0xfb, 0x10, 0xca, 0xd2, 0xf2, 0x78, 0x1a, 0x26, 0x61, 0xea, 0x8f, 0xec, 0x43, 0x98, 0x94, 0xd2,
	0x3a, 0x23, 0x57, 0x8d, 0xd3, 0xd8, 0x01, 0x60, 0x07, 0xe3, 0x2d, 0x94, 0xec, 0x58, 0x34, 0xd6,
	0xe9, 0x9a, 0xba, 0x1c, 0x22, 0x07, 0x36, 0x10, 0x39, 0xe2, 0xd6, 0xcc, 0xa3, 0xbb, 0x66, 0x9e,
	0xfd, 0x15, 0xc0, 0xc1, 0x8d, 0x9c, 0xa9, 0xb0, 0x6b, 0xad, 0xac, 0x60, 0x9f, 0x42, 0x5f, 0x28,
	0x27, 0x5d, 0x8b, 0x92, 0x0e, 0x17, 0xcf, 0xde, 0xb9, 0x3e, 0x98, 0xd2, 0x11, 0xd9, 0xc7, 0x34,
	0x04, 0x83, 0x52, 0x0f, 0x17, 0x47, 0xf7, 0x33, 0xde, 0xf8, 0x20, 0xcd, 0xc6, 0xb0, 0x4f, 0x20,
	0xf2, 0x0e, 0xb3, 0x3c, 0xc4, 0xed, 0xfc, 0x87, 0xf2, 0xc4, 0xf3, 0xd5, 0x85, 0x31, 0xda, 0xe0,
	0x78, 0xde, 0x51, 0xfd, 0x4b, 0x1f, 0x4c, 0x89, 0xc3, 0x12, 0xd8, 0x91, 0xea, 0x5c, 0xe3, 0xc0,
	0x86, 0x8b, 0xc3, 0xfb, 0x5c, 0xac, 0x8b, 0x8c, 0xc5, 0xaf, 0x01, 0x4c, 0xce, 0x5a, 0xfc, 0xb5,
	0xa5, 0x30, 0xd7, 0xde, 0x80, 0x5f, 0x43, 0xf8, 0x95, 0x70, 0xec, 0x83, 0x07, 0x5b, 0x42, 0xcf,
	0x1d, 0xbf, 0x78, 0xb8, 0xe5, 0x4e, 0xc4, 0xd9, 0xff, 0xd8, 0x19, 0xf4, 0x97, 0x22, 0x37, 0xc5,
	0xe5, 0x7f, 0x52, 0x6e, 0xf1, 0x4b, 0x00, 0xa3, 0xc7, 0xd2, 0xeb, 0x6f, 0x01, 0x1c, 0x7c, 0xae,
	0x55, 0x63, 0x1f, 0x4b, 0xc3, 0xbf, 0x07, 0xf0, 0xe4, 0xac, 0x7d, 0x4c, 0x2d, 0xaf, 0xfa, 0xf8,
	0x5f, 0xe7, 0xb3, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x24, 0xe1, 0xc7, 0x0c, 0x1c, 0x09, 0x00,
	0x00,
}
