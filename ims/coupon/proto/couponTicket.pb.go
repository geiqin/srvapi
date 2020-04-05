// Code generated by protoc-gen-go. DO NOT EDIT.
// source: couponTicket.proto

package geiqin_srv_ims_coupon

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

type CouponTicketWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keywords             string   `protobuf:"bytes,3,opt,name=keywords,proto3" json:"keywords,omitempty"`
	CustomerId           int64    `protobuf:"varint,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CouponId             int64    `protobuf:"varint,5,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	Id                   int64    `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Valid                bool     `protobuf:"varint,8,opt,name=valid,proto3" json:"valid,omitempty"`
	Status               string   `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CouponTicketWhere) Reset()         { *m = CouponTicketWhere{} }
func (m *CouponTicketWhere) String() string { return proto.CompactTextString(m) }
func (*CouponTicketWhere) ProtoMessage()    {}
func (*CouponTicketWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_b05ea0f1ee352a83, []int{0}
}

func (m *CouponTicketWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CouponTicketWhere.Unmarshal(m, b)
}
func (m *CouponTicketWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CouponTicketWhere.Marshal(b, m, deterministic)
}
func (m *CouponTicketWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CouponTicketWhere.Merge(m, src)
}
func (m *CouponTicketWhere) XXX_Size() int {
	return xxx_messageInfo_CouponTicketWhere.Size(m)
}
func (m *CouponTicketWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_CouponTicketWhere.DiscardUnknown(m)
}

var xxx_messageInfo_CouponTicketWhere proto.InternalMessageInfo

func (m *CouponTicketWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *CouponTicketWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *CouponTicketWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *CouponTicketWhere) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *CouponTicketWhere) GetCouponId() int64 {
	if m != nil {
		return m.CouponId
	}
	return 0
}

func (m *CouponTicketWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CouponTicketWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *CouponTicketWhere) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *CouponTicketWhere) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

// 优惠劵凭证
type CouponTicket struct {
	Id                   int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TicketSn             string      `protobuf:"bytes,2,opt,name=ticket_sn,json=ticketSn,proto3" json:"ticket_sn,omitempty"`
	StartAt              string      `protobuf:"bytes,3,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                string      `protobuf:"bytes,4,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	CouponId             int64       `protobuf:"varint,5,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	CustomerId           int64       `protobuf:"varint,6,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	OrderId              int64       `protobuf:"varint,7,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	IsUsed               bool        `protobuf:"varint,8,opt,name=is_used,json=isUsed,proto3" json:"is_used,omitempty"`
	UseAt                string      `protobuf:"bytes,9,opt,name=use_at,json=useAt,proto3" json:"use_at,omitempty"`
	Status               int32       `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string      `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string      `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ids                  []int64     `protobuf:"varint,13,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Coupon               *CouponInfo `protobuf:"bytes,14,opt,name=coupon,proto3" json:"coupon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CouponTicket) Reset()         { *m = CouponTicket{} }
func (m *CouponTicket) String() string { return proto.CompactTextString(m) }
func (*CouponTicket) ProtoMessage()    {}
func (*CouponTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor_b05ea0f1ee352a83, []int{1}
}

func (m *CouponTicket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CouponTicket.Unmarshal(m, b)
}
func (m *CouponTicket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CouponTicket.Marshal(b, m, deterministic)
}
func (m *CouponTicket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CouponTicket.Merge(m, src)
}
func (m *CouponTicket) XXX_Size() int {
	return xxx_messageInfo_CouponTicket.Size(m)
}
func (m *CouponTicket) XXX_DiscardUnknown() {
	xxx_messageInfo_CouponTicket.DiscardUnknown(m)
}

var xxx_messageInfo_CouponTicket proto.InternalMessageInfo

func (m *CouponTicket) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CouponTicket) GetTicketSn() string {
	if m != nil {
		return m.TicketSn
	}
	return ""
}

func (m *CouponTicket) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *CouponTicket) GetEndAt() string {
	if m != nil {
		return m.EndAt
	}
	return ""
}

func (m *CouponTicket) GetCouponId() int64 {
	if m != nil {
		return m.CouponId
	}
	return 0
}

func (m *CouponTicket) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *CouponTicket) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *CouponTicket) GetIsUsed() bool {
	if m != nil {
		return m.IsUsed
	}
	return false
}

func (m *CouponTicket) GetUseAt() string {
	if m != nil {
		return m.UseAt
	}
	return ""
}

func (m *CouponTicket) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CouponTicket) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *CouponTicket) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *CouponTicket) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *CouponTicket) GetCoupon() *CouponInfo {
	if m != nil {
		return m.Coupon
	}
	return nil
}

// 优惠券
type CouponInfo struct {
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CouponSn             string         `protobuf:"bytes,2,opt,name=coupon_sn,json=couponSn,proto3" json:"coupon_sn,omitempty"`
	Title                string         `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	IsAtLeast            bool           `protobuf:"varint,4,opt,name=is_at_least,json=isAtLeast,proto3" json:"is_at_least,omitempty"`
	AtLeast              float32        `protobuf:"fixed32,5,opt,name=at_least,json=atLeast,proto3" json:"at_least,omitempty"`
	PreferentialType     int32          `protobuf:"varint,6,opt,name=preferential_type,json=preferentialType,proto3" json:"preferential_type,omitempty"`
	PreferentialMoney    float32        `protobuf:"fixed32,7,opt,name=preferential_money,json=preferentialMoney,proto3" json:"preferential_money,omitempty"`
	PreferentialDiscount float32        `protobuf:"fixed32,8,opt,name=preferential_discount,json=preferentialDiscount,proto3" json:"preferential_discount,omitempty"`
	ExchangeNum          int32          `protobuf:"varint,9,opt,name=exchange_num,json=exchangeNum,proto3" json:"exchange_num,omitempty"`
	RangeType            string         `protobuf:"bytes,10,opt,name=range_type,json=rangeType,proto3" json:"range_type,omitempty"`
	Description          string         `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	Goodses              []*CouponGoods `protobuf:"bytes,12,rep,name=goodses,proto3" json:"goodses,omitempty"`
	Ids                  []int64        `protobuf:"varint,13,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CouponInfo) Reset()         { *m = CouponInfo{} }
func (m *CouponInfo) String() string { return proto.CompactTextString(m) }
func (*CouponInfo) ProtoMessage()    {}
func (*CouponInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b05ea0f1ee352a83, []int{2}
}

func (m *CouponInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CouponInfo.Unmarshal(m, b)
}
func (m *CouponInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CouponInfo.Marshal(b, m, deterministic)
}
func (m *CouponInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CouponInfo.Merge(m, src)
}
func (m *CouponInfo) XXX_Size() int {
	return xxx_messageInfo_CouponInfo.Size(m)
}
func (m *CouponInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CouponInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CouponInfo proto.InternalMessageInfo

func (m *CouponInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CouponInfo) GetCouponSn() string {
	if m != nil {
		return m.CouponSn
	}
	return ""
}

func (m *CouponInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CouponInfo) GetIsAtLeast() bool {
	if m != nil {
		return m.IsAtLeast
	}
	return false
}

func (m *CouponInfo) GetAtLeast() float32 {
	if m != nil {
		return m.AtLeast
	}
	return 0
}

func (m *CouponInfo) GetPreferentialType() int32 {
	if m != nil {
		return m.PreferentialType
	}
	return 0
}

func (m *CouponInfo) GetPreferentialMoney() float32 {
	if m != nil {
		return m.PreferentialMoney
	}
	return 0
}

func (m *CouponInfo) GetPreferentialDiscount() float32 {
	if m != nil {
		return m.PreferentialDiscount
	}
	return 0
}

func (m *CouponInfo) GetExchangeNum() int32 {
	if m != nil {
		return m.ExchangeNum
	}
	return 0
}

func (m *CouponInfo) GetRangeType() string {
	if m != nil {
		return m.RangeType
	}
	return ""
}

func (m *CouponInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CouponInfo) GetGoodses() []*CouponGoods {
	if m != nil {
		return m.Goodses
	}
	return nil
}

func (m *CouponInfo) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

//
type CouponTicketResponse struct {
	Entity               *CouponTicket   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager          `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*CouponTicket `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error          `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info           `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	Ids                  []int64         `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CouponTicketResponse) Reset()         { *m = CouponTicketResponse{} }
func (m *CouponTicketResponse) String() string { return proto.CompactTextString(m) }
func (*CouponTicketResponse) ProtoMessage()    {}
func (*CouponTicketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b05ea0f1ee352a83, []int{3}
}

func (m *CouponTicketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CouponTicketResponse.Unmarshal(m, b)
}
func (m *CouponTicketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CouponTicketResponse.Marshal(b, m, deterministic)
}
func (m *CouponTicketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CouponTicketResponse.Merge(m, src)
}
func (m *CouponTicketResponse) XXX_Size() int {
	return xxx_messageInfo_CouponTicketResponse.Size(m)
}
func (m *CouponTicketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CouponTicketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CouponTicketResponse proto.InternalMessageInfo

func (m *CouponTicketResponse) GetEntity() *CouponTicket {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *CouponTicketResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *CouponTicketResponse) GetItems() []*CouponTicket {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CouponTicketResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CouponTicketResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *CouponTicketResponse) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*CouponTicketWhere)(nil), "geiqin.srv.ims.coupon.CouponTicketWhere")
	proto.RegisterType((*CouponTicket)(nil), "geiqin.srv.ims.coupon.CouponTicket")
	proto.RegisterType((*CouponInfo)(nil), "geiqin.srv.ims.coupon.CouponInfo")
	proto.RegisterType((*CouponTicketResponse)(nil), "geiqin.srv.ims.coupon.CouponTicketResponse")
}

func init() {
	proto.RegisterFile("couponTicket.proto", fileDescriptor_b05ea0f1ee352a83)
}

var fileDescriptor_b05ea0f1ee352a83 = []byte{
	// 831 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xc6, 0xde, 0x78, 0xbd, 0x3e, 0x6b, 0x4a, 0x33, 0x24, 0xb0, 0x4d, 0x28, 0xb8, 0xe6, 0xc6,
	0x52, 0x85, 0x91, 0xdc, 0xab, 0x0a, 0x6e, 0xcc, 0x8f, 0x2a, 0x4b, 0x2d, 0x42, 0xeb, 0x46, 0x88,
	0x1b, 0x56, 0x9b, 0x9d, 0x13, 0x67, 0x54, 0x7b, 0x66, 0x99, 0x99, 0x0d, 0xb8, 0xcf, 0xc1, 0x0d,
	0x8f, 0xc2, 0x2d, 0x2f, 0xc0, 0xab, 0xf0, 0x08, 0x68, 0xce, 0x8c, 0x13, 0x87, 0x26, 0x51, 0x7a,
	0x91, 0xde, 0xed, 0xf9, 0xbe, 0x73, 0xe6, 0xfc, 0x7c, 0x67, 0x67, 0x17, 0x58, 0xa5, 0x9a, 0x5a,
	0xc9, 0x97, 0xa2, 0x7a, 0x85, 0x76, 0x5c, 0x6b, 0x65, 0x15, 0xdb, 0x5f, 0xa0, 0xf8, 0x55, 0xc8,
	0xb1, 0xd1, 0x67, 0x63, 0xb1, 0x32, 0x63, 0xef, 0x72, 0xd0, 0xaf, 0xd4, 0x6a, 0xa5, 0xa4, 0x77,
	0x3a, 0xe8, 0x1f, 0x37, 0x6b, 0x21, 0x17, 0xc1, 0xda, 0xf5, 0x3e, 0xcf, 0x94, 0xe2, 0xc6, 0x43,
	0xc3, 0x7f, 0x5b, 0xb0, 0xfb, 0xed, 0xd6, 0xe1, 0x3f, 0x9d, 0xa2, 0x46, 0xb6, 0x07, 0x9d, 0xba,
	0x5c, 0x20, 0xcf, 0x5a, 0x83, 0xd6, 0xa8, 0x93, 0x7b, 0x83, 0x1d, 0x42, 0xcf, 0x3d, 0x14, 0x46,
	0xbc, 0xc6, 0xac, 0x4d, 0x4c, 0xe2, 0x80, 0xb9, 0x78, 0x8d, 0xec, 0x00, 0x92, 0x57, 0xb8, 0xfe,
	0x4d, 0x69, 0x6e, 0xb2, 0x68, 0xd0, 0x1a, 0xf5, 0xf2, 0x73, 0x9b, 0x7d, 0x06, 0x69, 0xd5, 0x18,
	0xab, 0x56, 0xa8, 0x0b, 0xc1, 0xb3, 0x9d, 0x41, 0x6b, 0x14, 0xe5, 0xb0, 0x81, 0x66, 0x74, 0xb2,
	0x2f, 0xcd, 0xd1, 0x1d, 0xa2, 0x13, 0x0f, 0xcc, 0x38, 0xbb, 0x07, 0x6d, 0xc1, 0xb3, 0x98, 0xd0,
	0xb6, 0xe0, 0xec, 0x3e, 0x44, 0x82, 0x9b, 0xac, 0x3b, 0x88, 0x46, 0x51, 0xee, 0x1e, 0x5d, 0xb9,
	0x67, 0xe5, 0x52, 0xf0, 0x2c, 0x19, 0xb4, 0x46, 0x49, 0xee, 0x0d, 0xf6, 0x11, 0xc4, 0xc6, 0x96,
	0xb6, 0x31, 0x59, 0x8f, 0xea, 0x09, 0xd6, 0xf0, 0xcf, 0x08, 0xfa, 0xdb, 0x2d, 0x87, 0x04, 0xad,
	0xf3, 0x04, 0x87, 0xd0, 0xb3, 0xc4, 0x14, 0x46, 0x52, 0x9f, 0xbd, 0x3c, 0xf1, 0xc0, 0x5c, 0xb2,
	0x07, 0x90, 0x18, 0x5b, 0x6a, 0x5b, 0x94, 0x36, 0xf4, 0xd9, 0x25, 0x7b, 0x6a, 0xd9, 0x3e, 0xc4,
	0x28, 0xb9, 0x23, 0x76, 0x88, 0xe8, 0xa0, 0xe4, 0x53, 0x7b, 0x73, 0x73, 0xff, 0x1b, 0x4d, 0xfc,
	0xc6, 0x68, 0x1e, 0x40, 0xa2, 0x34, 0xf7, 0x6c, 0x97, 0xd8, 0x2e, 0xd9, 0x33, 0xce, 0x3e, 0x86,
	0xae, 0x30, 0x45, 0x63, 0x70, 0xd3, 0x78, 0x2c, 0xcc, 0x91, 0x41, 0xee, 0x0a, 0x69, 0x0c, 0xba,
	0x42, 0x7c, 0xe7, 0x9d, 0xc6, 0xe0, 0xd4, 0x6e, 0x0d, 0x04, 0x48, 0xbc, 0x60, 0xb1, 0x87, 0x00,
	0x95, 0xc6, 0xd2, 0x22, 0xd5, 0x9e, 0x52, 0x48, 0x2f, 0x20, 0x53, 0xeb, 0xe8, 0xa6, 0xe6, 0x1b,
	0xba, 0xef, 0xe9, 0x80, 0x4c, 0xed, 0x46, 0x8e, 0xf7, 0x2f, 0xe4, 0x78, 0x0a, 0xb1, 0xef, 0x2f,
	0xbb, 0x37, 0x68, 0x8d, 0xd2, 0xc9, 0xa3, 0xf1, 0x95, 0xab, 0x3a, 0xf6, 0x22, 0xcc, 0xe4, 0x89,
	0xca, 0x43, 0xc0, 0xf0, 0x9f, 0x08, 0xe0, 0x02, 0xbe, 0x4a, 0x99, 0x30, 0xca, 0x0b, 0x65, 0x3c,
	0x30, 0x97, 0x6e, 0x0b, 0xac, 0xb0, 0x4b, 0x0c, 0xb2, 0x78, 0x83, 0x7d, 0x0a, 0xa9, 0x30, 0x45,
	0x69, 0x8b, 0x25, 0x96, 0xc6, 0x2b, 0x93, 0xe4, 0x3d, 0x61, 0xa6, 0xf6, 0xb9, 0x03, 0xdc, 0x7c,
	0xcf, 0x49, 0x27, 0x4e, 0x3b, 0xef, 0x96, 0x81, 0x7a, 0x0c, 0xbb, 0xb5, 0xc6, 0x13, 0xd4, 0x28,
	0xad, 0x28, 0x97, 0x85, 0x5d, 0xd7, 0x48, 0x0a, 0x75, 0xf2, 0xfb, 0xdb, 0xc4, 0xcb, 0x75, 0x8d,
	0xec, 0x0b, 0x60, 0x97, 0x9c, 0x57, 0x4a, 0xe2, 0x9a, 0x14, 0x6b, 0xe7, 0x97, 0x8e, 0x79, 0xe1,
	0x08, 0xf6, 0x04, 0xf6, 0x2f, 0xb9, 0x73, 0x61, 0x2a, 0xd5, 0x48, 0x4b, 0x4a, 0xb6, 0xf3, 0xbd,
	0x6d, 0xf2, 0xbb, 0xc0, 0xb1, 0x47, 0xd0, 0xc7, 0xdf, 0xab, 0xd3, 0x52, 0x2e, 0xb0, 0x90, 0xcd,
	0x8a, 0xd4, 0xed, 0xe4, 0xe9, 0x06, 0xfb, 0xa1, 0x59, 0x39, 0xb1, 0x34, 0xf1, 0x54, 0x2c, 0x78,
	0xb1, 0x08, 0xa1, 0x2a, 0x07, 0x90, 0x72, 0x34, 0x95, 0x16, 0xb5, 0x15, 0x4a, 0x06, 0xad, 0xb7,
	0x21, 0xf6, 0x35, 0x74, 0x17, 0xee, 0x7e, 0x40, 0x93, 0xf5, 0x07, 0xd1, 0x28, 0x9d, 0x0c, 0x6f,
	0x54, 0x8f, 0xee, 0x92, 0x7c, 0x13, 0xf2, 0xe6, 0x32, 0x0c, 0xff, 0x6a, 0xc3, 0xde, 0xf6, 0xdb,
	0x96, 0xa3, 0xa9, 0x95, 0x34, 0xc8, 0xbe, 0x72, 0x6f, 0x8b, 0x15, 0x76, 0x4d, 0xfa, 0xa6, 0x93,
	0xcf, 0x6f, 0xcc, 0x13, 0x82, 0x43, 0x08, 0x9b, 0xf8, 0x0b, 0x4a, 0xd3, 0x12, 0xa4, 0x93, 0x4f,
	0xae, 0x89, 0xfd, 0xd1, 0xf9, 0xf8, 0xeb, 0x4b, 0xb3, 0xa7, 0xd0, 0x11, 0x16, 0x57, 0xee, 0x7a,
	0x8a, 0x6e, 0x9b, 0xcf, 0x47, 0xb8, 0x74, 0xa8, 0xb5, 0xd2, 0xb4, 0x3e, 0xd7, 0xa7, 0xfb, 0xde,
	0xf9, 0xe4, 0xde, 0x95, 0x7d, 0x09, 0x3b, 0x42, 0x9e, 0x28, 0x5a, 0xaa, 0x74, 0x72, 0x78, 0x4d,
	0x08, 0x6d, 0x3f, 0x39, 0x6e, 0x66, 0x17, 0x9f, 0xcf, 0x6e, 0xf2, 0x47, 0x04, 0x1f, 0x6e, 0x97,
	0x33, 0x47, 0x7d, 0x26, 0x2a, 0x64, 0x3f, 0x43, 0x74, 0x64, 0x90, 0xdd, 0xa6, 0x83, 0x83, 0xc7,
	0xb7, 0x69, 0x33, 0x68, 0x32, 0x7c, 0x8f, 0xfd, 0x02, 0xd1, 0x33, 0xb4, 0x6c, 0x74, 0x8b, 0x28,
	0xfa, 0x54, 0xbc, 0xed, 0xf9, 0x25, 0xc4, 0x73, 0x2c, 0x75, 0x75, 0x7a, 0x77, 0x29, 0x0a, 0xd8,
	0x79, 0x2e, 0xcc, 0xdd, 0xf5, 0x30, 0xf9, 0xbb, 0x0d, 0xfb, 0x2f, 0xd6, 0x57, 0x09, 0xf3, 0x4e,
	0xba, 0xfb, 0xe0, 0xc8, 0x94, 0xc7, 0x4b, 0xf4, 0xcc, 0x8c, 0x1b, 0xf6, 0xf0, 0x9a, 0x13, 0xbe,
	0xa1, 0x6f, 0xff, 0xdb, 0x26, 0x40, 0x00, 0x9f, 0xe0, 0x4e, 0x87, 0x78, 0x1c, 0xd3, 0xff, 0xc7,
	0x93, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xad, 0x4c, 0x0a, 0x2b, 0xdb, 0x08, 0x00, 0x00,
}
