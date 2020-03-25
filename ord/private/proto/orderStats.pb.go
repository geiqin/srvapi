// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderStats.proto

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

type OrderStats struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PayAmount            float32  `protobuf:"fixed32,2,opt,name=pay_amount,json=payAmount,proto3" json:"pay_amount,omitempty"`
	CouponAmount         float32  `protobuf:"fixed32,3,opt,name=coupon_amount,json=couponAmount,proto3" json:"coupon_amount,omitempty"`
	ExpressAmount        float32  `protobuf:"fixed32,4,opt,name=express_amount,json=expressAmount,proto3" json:"express_amount,omitempty"`
	RefundAmount         float32  `protobuf:"fixed32,5,opt,name=refund_amount,json=refundAmount,proto3" json:"refund_amount,omitempty"`
	PayNum               int32    `protobuf:"varint,6,opt,name=pay_num,json=payNum,proto3" json:"pay_num,omitempty"`
	BuyerNum             int32    `protobuf:"varint,7,opt,name=buyer_num,json=buyerNum,proto3" json:"buyer_num,omitempty"`
	GoodsNum             int32    `protobuf:"varint,8,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num,omitempty"`
	RefundNum            int32    `protobuf:"varint,9,opt,name=refund_num,json=refundNum,proto3" json:"refund_num,omitempty"`
	NewCustomerNum       int32    `protobuf:"varint,10,opt,name=new_customer_num,json=newCustomerNum,proto3" json:"new_customer_num,omitempty"`
	OldCustomerNum       int32    `protobuf:"varint,11,opt,name=old_customer_num,json=oldCustomerNum,proto3" json:"old_customer_num,omitempty"`
	StatisticDate        string   `protobuf:"bytes,12,opt,name=statistic_date,json=statisticDate,proto3" json:"statistic_date,omitempty"`
	CreatedAt            string   `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderStats) Reset()         { *m = OrderStats{} }
func (m *OrderStats) String() string { return proto.CompactTextString(m) }
func (*OrderStats) ProtoMessage()    {}
func (*OrderStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_a14a673301f078e7, []int{0}
}

func (m *OrderStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderStats.Unmarshal(m, b)
}
func (m *OrderStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderStats.Marshal(b, m, deterministic)
}
func (m *OrderStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderStats.Merge(m, src)
}
func (m *OrderStats) XXX_Size() int {
	return xxx_messageInfo_OrderStats.Size(m)
}
func (m *OrderStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderStats.DiscardUnknown(m)
}

var xxx_messageInfo_OrderStats proto.InternalMessageInfo

func (m *OrderStats) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderStats) GetPayAmount() float32 {
	if m != nil {
		return m.PayAmount
	}
	return 0
}

func (m *OrderStats) GetCouponAmount() float32 {
	if m != nil {
		return m.CouponAmount
	}
	return 0
}

func (m *OrderStats) GetExpressAmount() float32 {
	if m != nil {
		return m.ExpressAmount
	}
	return 0
}

func (m *OrderStats) GetRefundAmount() float32 {
	if m != nil {
		return m.RefundAmount
	}
	return 0
}

func (m *OrderStats) GetPayNum() int32 {
	if m != nil {
		return m.PayNum
	}
	return 0
}

func (m *OrderStats) GetBuyerNum() int32 {
	if m != nil {
		return m.BuyerNum
	}
	return 0
}

func (m *OrderStats) GetGoodsNum() int32 {
	if m != nil {
		return m.GoodsNum
	}
	return 0
}

func (m *OrderStats) GetRefundNum() int32 {
	if m != nil {
		return m.RefundNum
	}
	return 0
}

func (m *OrderStats) GetNewCustomerNum() int32 {
	if m != nil {
		return m.NewCustomerNum
	}
	return 0
}

func (m *OrderStats) GetOldCustomerNum() int32 {
	if m != nil {
		return m.OldCustomerNum
	}
	return 0
}

func (m *OrderStats) GetStatisticDate() string {
	if m != nil {
		return m.StatisticDate
	}
	return ""
}

func (m *OrderStats) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *OrderStats) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type OrderDayStats struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PayAmount            float32  `protobuf:"fixed32,2,opt,name=pay_amount,json=payAmount,proto3" json:"pay_amount,omitempty"`
	CouponAmount         float32  `protobuf:"fixed32,3,opt,name=coupon_amount,json=couponAmount,proto3" json:"coupon_amount,omitempty"`
	ExpressAmount        float32  `protobuf:"fixed32,4,opt,name=express_amount,json=expressAmount,proto3" json:"express_amount,omitempty"`
	RefundAmount         float32  `protobuf:"fixed32,5,opt,name=refund_amount,json=refundAmount,proto3" json:"refund_amount,omitempty"`
	PayNum               int32    `protobuf:"varint,6,opt,name=pay_num,json=payNum,proto3" json:"pay_num,omitempty"`
	BuyerNum             int32    `protobuf:"varint,7,opt,name=buyer_num,json=buyerNum,proto3" json:"buyer_num,omitempty"`
	GoodsNum             int32    `protobuf:"varint,8,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num,omitempty"`
	RefundNum            int32    `protobuf:"varint,9,opt,name=refund_num,json=refundNum,proto3" json:"refund_num,omitempty"`
	NewCustomerNum       int32    `protobuf:"varint,10,opt,name=new_customer_num,json=newCustomerNum,proto3" json:"new_customer_num,omitempty"`
	OldCustomerNum       int32    `protobuf:"varint,11,opt,name=old_customer_num,json=oldCustomerNum,proto3" json:"old_customer_num,omitempty"`
	CreatedAt            string   `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderDayStats) Reset()         { *m = OrderDayStats{} }
func (m *OrderDayStats) String() string { return proto.CompactTextString(m) }
func (*OrderDayStats) ProtoMessage()    {}
func (*OrderDayStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_a14a673301f078e7, []int{1}
}

func (m *OrderDayStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderDayStats.Unmarshal(m, b)
}
func (m *OrderDayStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderDayStats.Marshal(b, m, deterministic)
}
func (m *OrderDayStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderDayStats.Merge(m, src)
}
func (m *OrderDayStats) XXX_Size() int {
	return xxx_messageInfo_OrderDayStats.Size(m)
}
func (m *OrderDayStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderDayStats.DiscardUnknown(m)
}

var xxx_messageInfo_OrderDayStats proto.InternalMessageInfo

func (m *OrderDayStats) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderDayStats) GetPayAmount() float32 {
	if m != nil {
		return m.PayAmount
	}
	return 0
}

func (m *OrderDayStats) GetCouponAmount() float32 {
	if m != nil {
		return m.CouponAmount
	}
	return 0
}

func (m *OrderDayStats) GetExpressAmount() float32 {
	if m != nil {
		return m.ExpressAmount
	}
	return 0
}

func (m *OrderDayStats) GetRefundAmount() float32 {
	if m != nil {
		return m.RefundAmount
	}
	return 0
}

func (m *OrderDayStats) GetPayNum() int32 {
	if m != nil {
		return m.PayNum
	}
	return 0
}

func (m *OrderDayStats) GetBuyerNum() int32 {
	if m != nil {
		return m.BuyerNum
	}
	return 0
}

func (m *OrderDayStats) GetGoodsNum() int32 {
	if m != nil {
		return m.GoodsNum
	}
	return 0
}

func (m *OrderDayStats) GetRefundNum() int32 {
	if m != nil {
		return m.RefundNum
	}
	return 0
}

func (m *OrderDayStats) GetNewCustomerNum() int32 {
	if m != nil {
		return m.NewCustomerNum
	}
	return 0
}

func (m *OrderDayStats) GetOldCustomerNum() int32 {
	if m != nil {
		return m.OldCustomerNum
	}
	return 0
}

func (m *OrderDayStats) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *OrderDayStats) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type OrderStatsResponse struct {
	Entity               *OrderStats   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager        `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*OrderStats `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error        `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info         `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *OrderStatsResponse) Reset()         { *m = OrderStatsResponse{} }
func (m *OrderStatsResponse) String() string { return proto.CompactTextString(m) }
func (*OrderStatsResponse) ProtoMessage()    {}
func (*OrderStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a14a673301f078e7, []int{2}
}

func (m *OrderStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderStatsResponse.Unmarshal(m, b)
}
func (m *OrderStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderStatsResponse.Marshal(b, m, deterministic)
}
func (m *OrderStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderStatsResponse.Merge(m, src)
}
func (m *OrderStatsResponse) XXX_Size() int {
	return xxx_messageInfo_OrderStatsResponse.Size(m)
}
func (m *OrderStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderStatsResponse proto.InternalMessageInfo

func (m *OrderStatsResponse) GetEntity() *OrderStats {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *OrderStatsResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *OrderStatsResponse) GetItems() []*OrderStats {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *OrderStatsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *OrderStatsResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

type OrderDayStatsResponse struct {
	Entity               *OrderDayStats   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager           `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*OrderDayStats `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error           `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info            `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *OrderDayStatsResponse) Reset()         { *m = OrderDayStatsResponse{} }
func (m *OrderDayStatsResponse) String() string { return proto.CompactTextString(m) }
func (*OrderDayStatsResponse) ProtoMessage()    {}
func (*OrderDayStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a14a673301f078e7, []int{3}
}

func (m *OrderDayStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderDayStatsResponse.Unmarshal(m, b)
}
func (m *OrderDayStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderDayStatsResponse.Marshal(b, m, deterministic)
}
func (m *OrderDayStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderDayStatsResponse.Merge(m, src)
}
func (m *OrderDayStatsResponse) XXX_Size() int {
	return xxx_messageInfo_OrderDayStatsResponse.Size(m)
}
func (m *OrderDayStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderDayStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderDayStatsResponse proto.InternalMessageInfo

func (m *OrderDayStatsResponse) GetEntity() *OrderDayStats {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *OrderDayStatsResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *OrderDayStatsResponse) GetItems() []*OrderDayStats {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *OrderDayStatsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *OrderDayStatsResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderStats)(nil), "geiqin.srv.ord.private.OrderStats")
	proto.RegisterType((*OrderDayStats)(nil), "geiqin.srv.ord.private.OrderDayStats")
	proto.RegisterType((*OrderStatsResponse)(nil), "geiqin.srv.ord.private.OrderStatsResponse")
	proto.RegisterType((*OrderDayStatsResponse)(nil), "geiqin.srv.ord.private.OrderDayStatsResponse")
}

func init() { proto.RegisterFile("orderStats.proto", fileDescriptor_a14a673301f078e7) }

var fileDescriptor_a14a673301f078e7 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0x41, 0x6b, 0xdb, 0x3e,
	0x18, 0xc6, 0x89, 0xdd, 0xa4, 0xf5, 0x9b, 0xd8, 0x14, 0xc1, 0xff, 0x3f, 0xb3, 0x2d, 0x10, 0x32,
	0x02, 0x39, 0x99, 0x91, 0x5e, 0xc6, 0xc6, 0x0e, 0x61, 0xdd, 0x61, 0x97, 0x6d, 0x78, 0x1f, 0x20,
	0xa8, 0x96, 0x12, 0x04, 0xb5, 0xa4, 0x49, 0x72, 0x3b, 0x7f, 0xa1, 0x5d, 0x77, 0xd8, 0x87, 0xd9,
	0xd7, 0x19, 0x7e, 0xa5, 0xa4, 0x4d, 0x59, 0x43, 0x61, 0xd0, 0xd3, 0xae, 0xcf, 0xf3, 0x7b, 0xf4,
	0x8a, 0xf7, 0x91, 0x0d, 0xa7, 0xca, 0x30, 0x6e, 0xbe, 0x38, 0xea, 0x6c, 0xa1, 0x8d, 0x72, 0x8a,
	0xfc, 0xbf, 0xe1, 0xe2, 0xab, 0x90, 0x85, 0x35, 0x57, 0x85, 0x32, 0xac, 0xd0, 0x46, 0x5c, 0x51,
	0xc7, 0x9f, 0x8e, 0x2a, 0x55, 0xd7, 0x4a, 0x7a, 0x6a, 0xfa, 0x2b, 0x06, 0xf8, 0xb4, 0x8b, 0x92,
	0x0c, 0x22, 0xc1, 0xf2, 0xde, 0xa4, 0x37, 0xef, 0x97, 0x91, 0x60, 0x64, 0x0c, 0xa0, 0x69, 0xbb,
	0xa2, 0xb5, 0x6a, 0xa4, 0xcb, 0xa3, 0x49, 0x6f, 0x1e, 0x95, 0x89, 0xa6, 0xed, 0x12, 0x05, 0xf2,
	0x02, 0xd2, 0x4a, 0x35, 0x5a, 0xc9, 0x2d, 0x11, 0x23, 0x31, 0xf2, 0x62, 0x80, 0x66, 0x90, 0xf1,
	0x6f, 0xda, 0x70, 0x6b, 0xb7, 0xd4, 0x11, 0x52, 0x69, 0x50, 0x6f, 0xce, 0x32, 0x7c, 0xdd, 0x48,
	0xb6, 0xa5, 0xfa, 0xfe, 0x2c, 0x2f, 0x06, 0xe8, 0x09, 0x1c, 0x77, 0xf7, 0x91, 0x4d, 0x9d, 0x0f,
	0xf0, 0x92, 0x03, 0x4d, 0xdb, 0x8f, 0x4d, 0x4d, 0x9e, 0x41, 0x72, 0xd1, 0xb4, 0xdc, 0xa0, 0x75,
	0x8c, 0xd6, 0x09, 0x0a, 0xc1, 0xdc, 0x28, 0xc5, 0x2c, 0x9a, 0x27, 0xde, 0x44, 0xa1, 0x33, 0xc7,
	0x00, 0x61, 0x6e, 0xe7, 0x26, 0xe8, 0x26, 0x5e, 0xe9, 0xec, 0x39, 0x9c, 0x4a, 0x7e, 0xbd, 0xaa,
	0x1a, 0xeb, 0x54, 0x1d, 0xce, 0x07, 0x84, 0x32, 0xc9, 0xaf, 0xdf, 0x05, 0x39, 0x90, 0xea, 0x92,
	0xed, 0x93, 0x43, 0x4f, 0xaa, 0x4b, 0x76, 0x9b, 0x9c, 0x41, 0x66, 0x1d, 0x75, 0xc2, 0x3a, 0x51,
	0xad, 0x18, 0x75, 0x3c, 0x1f, 0x4d, 0x7a, 0xf3, 0xa4, 0x4c, 0x77, 0xea, 0x39, 0x75, 0xbc, 0xbb,
	0x59, 0x65, 0x38, 0x75, 0x9c, 0xad, 0xa8, 0xcb, 0x53, 0x44, 0x92, 0xa0, 0x2c, 0x5d, 0x67, 0x37,
	0x9a, 0x6d, 0xed, 0xcc, 0xdb, 0x41, 0x59, 0xba, 0xe9, 0xcf, 0x18, 0x52, 0x6c, 0xf6, 0x9c, 0xb6,
	0xff, 0xca, 0x7d, 0xbc, 0x72, 0xf7, 0x5b, 0x1b, 0x1d, 0x6e, 0x2d, 0xbd, 0xdb, 0xda, 0xf7, 0x08,
	0xc8, 0xcd, 0xf7, 0x58, 0x72, 0xab, 0x95, 0xb4, 0x9c, 0xbc, 0x86, 0x01, 0x97, 0x4e, 0xb8, 0x16,
	0xeb, 0x1b, 0x2e, 0xa6, 0xc5, 0x9f, 0xbf, 0xee, 0xe2, 0x56, 0x36, 0x24, 0xc8, 0x19, 0xf4, 0x35,
	0xdd, 0x70, 0x83, 0x0d, 0x0f, 0x17, 0xe3, 0xfb, 0xa2, 0x9f, 0x3b, 0xa8, 0xf4, 0x2c, 0x79, 0x05,
	0x7d, 0xe1, 0x78, 0x6d, 0xf3, 0x78, 0x12, 0x3f, 0x70, 0x9e, 0x0f, 0x74, 0xe3, 0xb8, 0x31, 0xca,
	0xe0, 0x43, 0x38, 0x30, 0xee, 0x7d, 0x07, 0x95, 0x9e, 0x25, 0x2f, 0xe1, 0x48, 0xc8, 0xb5, 0xc2,
	0x67, 0x31, 0x5c, 0x3c, 0xbf, 0x2f, 0xf3, 0x41, 0xae, 0x55, 0x89, 0xe4, 0xf4, 0x47, 0x04, 0xff,
	0xed, 0x3d, 0xef, 0xdd, 0xae, 0xde, 0xde, 0xd9, 0xd5, 0xec, 0xe0, 0xdd, 0x77, 0xf1, 0xbf, 0x5a,
	0xd7, 0x9b, 0xfd, 0x75, 0x3d, 0x70, 0xe4, 0xa3, 0x6e, 0xec, 0x62, 0x80, 0x7f, 0xfc, 0xb3, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0xe6, 0x43, 0x08, 0x2b, 0x06, 0x00, 0x00,
}