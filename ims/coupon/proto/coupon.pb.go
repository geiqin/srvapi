// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coupon.proto

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

type CouponWhere struct {
	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	//base params
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Keywords             string   `protobuf:"bytes,6,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Status               int32    `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	PreferentialType     int32    `protobuf:"varint,8,opt,name=preferential_type,json=preferentialType,proto3" json:"preferential_type,omitempty"`
	CustomerId           int64    `protobuf:"varint,9,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Num                  int32    `protobuf:"varint,10,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CouponWhere) Reset()         { *m = CouponWhere{} }
func (m *CouponWhere) String() string { return proto.CompactTextString(m) }
func (*CouponWhere) ProtoMessage()    {}
func (*CouponWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{0}
}

func (m *CouponWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CouponWhere.Unmarshal(m, b)
}
func (m *CouponWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CouponWhere.Marshal(b, m, deterministic)
}
func (m *CouponWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CouponWhere.Merge(m, src)
}
func (m *CouponWhere) XXX_Size() int {
	return xxx_messageInfo_CouponWhere.Size(m)
}
func (m *CouponWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_CouponWhere.DiscardUnknown(m)
}

var xxx_messageInfo_CouponWhere proto.InternalMessageInfo

func (m *CouponWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *CouponWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *CouponWhere) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *CouponWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CouponWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *CouponWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *CouponWhere) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CouponWhere) GetPreferentialType() int32 {
	if m != nil {
		return m.PreferentialType
	}
	return 0
}

func (m *CouponWhere) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *CouponWhere) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

//优惠劵
type Coupon struct {
	Id                   int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CouponSn             string          `protobuf:"bytes,3,opt,name=coupon_sn,json=couponSn,proto3" json:"coupon_sn"`
	Title                string          `protobuf:"bytes,4,opt,name=title,proto3" json:"title"`
	Total                int32           `protobuf:"varint,5,opt,name=total,proto3" json:"total"`
	Stock                int32           `protobuf:"varint,6,opt,name=stock,proto3" json:"stock"`
	IsAtLeast            bool            `protobuf:"varint,7,opt,name=is_at_least,json=isAtLeast,proto3" json:"is_at_least"`
	AtLeast              float32         `protobuf:"fixed32,8,opt,name=at_least,json=atLeast,proto3" json:"at_least"`
	EffectiveType        int32           `protobuf:"varint,9,opt,name=effective_type,json=effectiveType,proto3" json:"effective_type"`
	EffectiveFixedTerm   int32           `protobuf:"varint,10,opt,name=effective_fixed_term,json=effectiveFixedTerm,proto3" json:"effective_fixed_term"`
	EffectiveStartAt     string          `protobuf:"bytes,11,opt,name=effective_start_at,json=effectiveStartAt,proto3" json:"effective_start_at"`
	EffectiveEndAt       string          `protobuf:"bytes,12,opt,name=effective_end_at,json=effectiveEndAt,proto3" json:"effective_end_at"`
	ExpireNotice         bool            `protobuf:"varint,13,opt,name=expire_notice,json=expireNotice,proto3" json:"expire_notice"`
	IsExpired            bool            `protobuf:"varint,14,opt,name=is_expired,json=isExpired,proto3" json:"is_expired"`
	IsOngoing            bool            `protobuf:"varint,15,opt,name=is_ongoing,json=isOngoing,proto3" json:"is_ongoing"`
	IsForbidPreference   bool            `protobuf:"varint,16,opt,name=is_forbid_preference,json=isForbidPreference,proto3" json:"is_forbid_preference"`
	IsShare              bool            `protobuf:"varint,17,opt,name=is_share,json=isShare,proto3" json:"is_share"`
	IsCancel             bool            `protobuf:"varint,18,opt,name=is_cancel,json=isCancel,proto3" json:"is_cancel"`
	IsDelete             bool            `protobuf:"varint,19,opt,name=is_delete,json=isDelete,proto3" json:"is_delete"`
	IsSyncWeixin         bool            `protobuf:"varint,20,opt,name=is_sync_weixin,json=isSyncWeixin,proto3" json:"is_sync_weixin"`
	NeedUserLevel        int32           `protobuf:"varint,21,opt,name=need_user_level,json=needUserLevel,proto3" json:"need_user_level"`
	PreferentialType     int32           `protobuf:"varint,22,opt,name=preferential_type,json=preferentialType,proto3" json:"preferential_type"`
	PreferentialMoney    float32         `protobuf:"fixed32,23,opt,name=preferential_money,json=preferentialMoney,proto3" json:"preferential_money"`
	PreferentialDiscount float32         `protobuf:"fixed32,24,opt,name=preferential_discount,json=preferentialDiscount,proto3" json:"preferential_discount"`
	Quota                int32           `protobuf:"varint,25,opt,name=quota,proto3" json:"quota"`
	ExchangeNum          int32           `protobuf:"varint,26,opt,name=exchange_num,json=exchangeNum,proto3" json:"exchange_num"`
	RangeType            string          `protobuf:"bytes,27,opt,name=range_type,json=rangeType,proto3" json:"range_type"`
	ServicePhone         string          `protobuf:"bytes,28,opt,name=service_phone,json=servicePhone,proto3" json:"service_phone"`
	TotalFansTaked       int32           `protobuf:"varint,29,opt,name=total_fans_taked,json=totalFansTaked,proto3" json:"total_fans_taked"`
	TotalFansUsed        int32           `protobuf:"varint,30,opt,name=total_fans_used,json=totalFansUsed,proto3" json:"total_fans_used"`
	TotalTaked           int32           `protobuf:"varint,31,opt,name=total_taked,json=totalTaked,proto3" json:"total_taked"`
	TotalUsed            int32           `protobuf:"varint,32,opt,name=total_used,json=totalUsed,proto3" json:"total_used"`
	CouponUrl            string          `protobuf:"bytes,33,opt,name=coupon_url,json=couponUrl,proto3" json:"coupon_url"`
	Description          string          `protobuf:"bytes,34,opt,name=description,proto3" json:"description"`
	Status               int32           `protobuf:"varint,35,opt,name=status,proto3" json:"status"`
	WeixinTitle          string          `protobuf:"bytes,36,opt,name=weixin_title,json=weixinTitle,proto3" json:"weixin_title"`
	WeixinSubTitle       string          `protobuf:"bytes,37,opt,name=weixin_sub_title,json=weixinSubTitle,proto3" json:"weixin_sub_title"`
	WeixinColor          string          `protobuf:"bytes,38,opt,name=weixin_color,json=weixinColor,proto3" json:"weixin_color"`
	ReceiveType          string          `protobuf:"bytes,39,opt,name=receive_type,json=receiveType,proto3" json:"receive_type"`
	CreatedAt            string          `protobuf:"bytes,40,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string          `protobuf:"bytes,41,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Tickets              []*CouponTicket `protobuf:"bytes,42,rep,name=tickets,proto3" json:"tickets,omitempty"`
	Goodses              []*CouponGoods  `protobuf:"bytes,43,rep,name=goodses,proto3" json:"goodses,omitempty"`
	Ids                  []int64         `protobuf:"varint,44,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	IsReceive            bool            `protobuf:"varint,45,opt,name=is_receive,json=isReceive,proto3" json:"is_receive"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Coupon) Reset()         { *m = Coupon{} }
func (m *Coupon) String() string { return proto.CompactTextString(m) }
func (*Coupon) ProtoMessage()    {}
func (*Coupon) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{1}
}

func (m *Coupon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coupon.Unmarshal(m, b)
}
func (m *Coupon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coupon.Marshal(b, m, deterministic)
}
func (m *Coupon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coupon.Merge(m, src)
}
func (m *Coupon) XXX_Size() int {
	return xxx_messageInfo_Coupon.Size(m)
}
func (m *Coupon) XXX_DiscardUnknown() {
	xxx_messageInfo_Coupon.DiscardUnknown(m)
}

var xxx_messageInfo_Coupon proto.InternalMessageInfo

func (m *Coupon) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Coupon) GetCouponSn() string {
	if m != nil {
		return m.CouponSn
	}
	return ""
}

func (m *Coupon) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Coupon) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Coupon) GetStock() int32 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func (m *Coupon) GetIsAtLeast() bool {
	if m != nil {
		return m.IsAtLeast
	}
	return false
}

func (m *Coupon) GetAtLeast() float32 {
	if m != nil {
		return m.AtLeast
	}
	return 0
}

func (m *Coupon) GetEffectiveType() int32 {
	if m != nil {
		return m.EffectiveType
	}
	return 0
}

func (m *Coupon) GetEffectiveFixedTerm() int32 {
	if m != nil {
		return m.EffectiveFixedTerm
	}
	return 0
}

func (m *Coupon) GetEffectiveStartAt() string {
	if m != nil {
		return m.EffectiveStartAt
	}
	return ""
}

func (m *Coupon) GetEffectiveEndAt() string {
	if m != nil {
		return m.EffectiveEndAt
	}
	return ""
}

func (m *Coupon) GetExpireNotice() bool {
	if m != nil {
		return m.ExpireNotice
	}
	return false
}

func (m *Coupon) GetIsExpired() bool {
	if m != nil {
		return m.IsExpired
	}
	return false
}

func (m *Coupon) GetIsOngoing() bool {
	if m != nil {
		return m.IsOngoing
	}
	return false
}

func (m *Coupon) GetIsForbidPreference() bool {
	if m != nil {
		return m.IsForbidPreference
	}
	return false
}

func (m *Coupon) GetIsShare() bool {
	if m != nil {
		return m.IsShare
	}
	return false
}

func (m *Coupon) GetIsCancel() bool {
	if m != nil {
		return m.IsCancel
	}
	return false
}

func (m *Coupon) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

func (m *Coupon) GetIsSyncWeixin() bool {
	if m != nil {
		return m.IsSyncWeixin
	}
	return false
}

func (m *Coupon) GetNeedUserLevel() int32 {
	if m != nil {
		return m.NeedUserLevel
	}
	return 0
}

func (m *Coupon) GetPreferentialType() int32 {
	if m != nil {
		return m.PreferentialType
	}
	return 0
}

func (m *Coupon) GetPreferentialMoney() float32 {
	if m != nil {
		return m.PreferentialMoney
	}
	return 0
}

func (m *Coupon) GetPreferentialDiscount() float32 {
	if m != nil {
		return m.PreferentialDiscount
	}
	return 0
}

func (m *Coupon) GetQuota() int32 {
	if m != nil {
		return m.Quota
	}
	return 0
}

func (m *Coupon) GetExchangeNum() int32 {
	if m != nil {
		return m.ExchangeNum
	}
	return 0
}

func (m *Coupon) GetRangeType() string {
	if m != nil {
		return m.RangeType
	}
	return ""
}

func (m *Coupon) GetServicePhone() string {
	if m != nil {
		return m.ServicePhone
	}
	return ""
}

func (m *Coupon) GetTotalFansTaked() int32 {
	if m != nil {
		return m.TotalFansTaked
	}
	return 0
}

func (m *Coupon) GetTotalFansUsed() int32 {
	if m != nil {
		return m.TotalFansUsed
	}
	return 0
}

func (m *Coupon) GetTotalTaked() int32 {
	if m != nil {
		return m.TotalTaked
	}
	return 0
}

func (m *Coupon) GetTotalUsed() int32 {
	if m != nil {
		return m.TotalUsed
	}
	return 0
}

func (m *Coupon) GetCouponUrl() string {
	if m != nil {
		return m.CouponUrl
	}
	return ""
}

func (m *Coupon) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Coupon) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Coupon) GetWeixinTitle() string {
	if m != nil {
		return m.WeixinTitle
	}
	return ""
}

func (m *Coupon) GetWeixinSubTitle() string {
	if m != nil {
		return m.WeixinSubTitle
	}
	return ""
}

func (m *Coupon) GetWeixinColor() string {
	if m != nil {
		return m.WeixinColor
	}
	return ""
}

func (m *Coupon) GetReceiveType() string {
	if m != nil {
		return m.ReceiveType
	}
	return ""
}

func (m *Coupon) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Coupon) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Coupon) GetTickets() []*CouponTicket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

func (m *Coupon) GetGoodses() []*CouponGoods {
	if m != nil {
		return m.Goodses
	}
	return nil
}

func (m *Coupon) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Coupon) GetIsReceive() bool {
	if m != nil {
		return m.IsReceive
	}
	return false
}

//
type CouponResponse struct {
	Entity               *Coupon   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Coupon `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error    `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info     `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CouponResponse) Reset()         { *m = CouponResponse{} }
func (m *CouponResponse) String() string { return proto.CompactTextString(m) }
func (*CouponResponse) ProtoMessage()    {}
func (*CouponResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{2}
}

func (m *CouponResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CouponResponse.Unmarshal(m, b)
}
func (m *CouponResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CouponResponse.Marshal(b, m, deterministic)
}
func (m *CouponResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CouponResponse.Merge(m, src)
}
func (m *CouponResponse) XXX_Size() int {
	return xxx_messageInfo_CouponResponse.Size(m)
}
func (m *CouponResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CouponResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CouponResponse proto.InternalMessageInfo

func (m *CouponResponse) GetEntity() *Coupon {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *CouponResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *CouponResponse) GetItems() []*Coupon {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CouponResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CouponResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*CouponWhere)(nil), "geiqin.srv.ims.coupon.CouponWhere")
	proto.RegisterType((*Coupon)(nil), "geiqin.srv.ims.coupon.Coupon")
	proto.RegisterType((*CouponResponse)(nil), "geiqin.srv.ims.coupon.CouponResponse")
}

func init() { proto.RegisterFile("coupon.proto", fileDescriptor_a727a1a30518ca78) }

var fileDescriptor_a727a1a30518ca78 = []byte{
	// 1205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0x5d, 0x73, 0x13, 0x37,
	0x14, 0x6d, 0xe2, 0xd8, 0xb1, 0xaf, 0x9d, 0x0f, 0xd4, 0x40, 0x45, 0x42, 0xc0, 0x24, 0x40, 0xdd,
	0x02, 0x6e, 0xc7, 0x4c, 0xdf, 0xda, 0x87, 0x4c, 0xf8, 0x18, 0x66, 0x80, 0x66, 0xd6, 0x31, 0xf4,
	0x6d, 0x67, 0xb3, 0x7b, 0xed, 0x68, 0xb2, 0x96, 0x16, 0x49, 0x1b, 0x62, 0x7e, 0x41, 0xfb, 0xdc,
	0xe9, 0xff, 0xe8, 0x4f, 0xec, 0xe8, 0x6a, 0xd7, 0x76, 0x66, 0x48, 0xca, 0x83, 0xdf, 0x56, 0xe7,
	0x1c, 0xdd, 0xd5, 0xb9, 0x7b, 0x75, 0xaf, 0x0d, 0xad, 0x58, 0xe5, 0x99, 0x92, 0xdd, 0x4c, 0x2b,
	0xab, 0xd8, 0xcd, 0x11, 0x8a, 0x8f, 0x42, 0x76, 0x8d, 0x3e, 0xef, 0x8a, 0xb1, 0xe9, 0x7a, 0x72,
	0xbb, 0x15, 0xab, 0xf1, 0xb8, 0x14, 0x6d, 0xdf, 0xf0, 0xe8, 0x2b, 0xa5, 0x12, 0x53, 0x40, 0xcc,
	0x43, 0xc7, 0x22, 0x3e, 0x43, 0xeb, 0xb1, 0xbd, 0x3f, 0x97, 0xa1, 0x79, 0x48, 0xf0, 0x87, 0x53,
	0xd4, 0xc8, 0xb6, 0xa0, 0x9a, 0x45, 0x23, 0x4c, 0xf8, 0x52, 0x7b, 0xa9, 0x53, 0x0d, 0xfc, 0x82,
	0xed, 0x40, 0xc3, 0x3d, 0x84, 0x46, 0x7c, 0x46, 0xbe, 0x4c, 0x4c, 0xdd, 0x01, 0x7d, 0xf1, 0x19,
	0xd9, 0x26, 0x54, 0xac, 0xca, 0x78, 0x85, 0x60, 0xf7, 0xc8, 0xd6, 0x61, 0x59, 0x24, 0x7c, 0xa5,
	0xbd, 0xd4, 0xa9, 0x04, 0xcb, 0x22, 0x71, 0x0a, 0x91, 0x18, 0x5e, 0x6d, 0x57, 0x3a, 0x95, 0xc0,
	0x3d, 0xb2, 0x6d, 0xa8, 0x9f, 0xe1, 0xe4, 0x93, 0xd2, 0x89, 0xe1, 0xb5, 0xf6, 0x52, 0xa7, 0x11,
	0x4c, 0xd7, 0xec, 0x16, 0xd4, 0x8c, 0x8d, 0x6c, 0x6e, 0xf8, 0x2a, 0x85, 0x2c, 0x56, 0xec, 0x31,
	0xdc, 0xc8, 0x34, 0x0e, 0x51, 0xa3, 0xb4, 0x22, 0x4a, 0x43, 0x3b, 0xc9, 0x90, 0xd7, 0x49, 0xb2,
	0x39, 0x4f, 0x1c, 0x4f, 0x32, 0x64, 0xf7, 0xa0, 0x19, 0xe7, 0xc6, 0xaa, 0x31, 0xea, 0x50, 0x24,
	0xbc, 0x41, 0x67, 0x81, 0x12, 0x7a, 0x4d, 0x67, 0x92, 0xf9, 0x98, 0x83, 0x3f, 0xb5, 0xcc, 0xc7,
	0x7b, 0x7f, 0xb7, 0xa0, 0xe6, 0x53, 0x51, 0x18, 0x58, 0x9a, 0x1a, 0xd8, 0x81, 0x86, 0xcf, 0x5d,
	0x68, 0x24, 0x19, 0x6d, 0x04, 0x75, 0x0f, 0xf4, 0xa5, 0x4b, 0x99, 0x15, 0x36, 0x45, 0x32, 0xdc,
	0x08, 0xfc, 0x82, 0x50, 0x65, 0xa3, 0x94, 0x57, 0x7d, 0x22, 0x69, 0xe1, 0x50, 0x63, 0x55, 0x7c,
	0x46, 0xa6, 0xab, 0x81, 0x5f, 0xb0, 0xbb, 0xd0, 0x14, 0x26, 0x8c, 0x6c, 0x98, 0x62, 0x64, 0x2c,
	0xd9, 0xae, 0x07, 0x0d, 0x61, 0x0e, 0xec, 0x1b, 0x07, 0xb0, 0xdb, 0x50, 0x9f, 0x92, 0xce, 0xf0,
	0x72, 0xb0, 0x1a, 0x15, 0xd4, 0x43, 0x58, 0xc7, 0xe1, 0x10, 0x63, 0x2b, 0xce, 0xd1, 0x67, 0xa4,
	0x41, 0x91, 0xd7, 0xa6, 0x28, 0xa5, 0xe3, 0x67, 0xd8, 0x9a, 0xc9, 0x86, 0xe2, 0x02, 0x93, 0xd0,
	0xa2, 0x2e, 0xed, 0xb3, 0x29, 0xf7, 0xd2, 0x51, 0xc7, 0xa8, 0xc7, 0xec, 0x09, 0xcc, 0xd0, 0xd0,
	0xd8, 0x48, 0xdb, 0x30, 0xb2, 0xbc, 0x49, 0x16, 0x37, 0xa7, 0x4c, 0xdf, 0x11, 0x07, 0x96, 0x75,
	0x60, 0x86, 0x85, 0x28, 0x13, 0xa7, 0x6d, 0x91, 0x76, 0x76, 0xbc, 0x17, 0x32, 0x39, 0xb0, 0x6c,
	0x1f, 0xd6, 0xf0, 0x22, 0x13, 0x1a, 0x43, 0xa9, 0xac, 0x88, 0x91, 0xaf, 0x91, 0xdb, 0x96, 0x07,
	0xdf, 0x11, 0xc6, 0x76, 0x01, 0x84, 0x09, 0x3d, 0x94, 0xf0, 0xf5, 0x32, 0x1f, 0x2f, 0x3c, 0x50,
	0xd0, 0x4a, 0x8e, 0x94, 0x90, 0x23, 0xbe, 0x51, 0xd2, 0xbf, 0x7b, 0xc0, 0x99, 0x15, 0x26, 0x1c,
	0x2a, 0x7d, 0x22, 0x92, 0xb0, 0xac, 0x8c, 0x18, 0xf9, 0x26, 0x09, 0x99, 0x30, 0x2f, 0x89, 0x3a,
	0x9a, 0x32, 0x2e, 0xc1, 0xc2, 0x84, 0xe6, 0x34, 0xd2, 0xc8, 0x6f, 0x90, 0x6a, 0x55, 0x98, 0xbe,
	0x5b, 0xba, 0x4f, 0x2f, 0x4c, 0x18, 0x47, 0x32, 0xc6, 0x94, 0x33, 0xe2, 0xea, 0xc2, 0x1c, 0xd2,
	0xba, 0x20, 0x13, 0x4c, 0xd1, 0x22, 0xff, 0xb6, 0x24, 0x9f, 0xd3, 0x9a, 0x3d, 0x80, 0x75, 0x17,
	0x74, 0x22, 0xe3, 0xf0, 0x13, 0x8a, 0x0b, 0x21, 0xf9, 0x96, 0xb7, 0x2a, 0x4c, 0x7f, 0x22, 0xe3,
	0x0f, 0x84, 0xb1, 0x47, 0xb0, 0x21, 0x11, 0x93, 0x30, 0x37, 0xa8, 0xc3, 0x14, 0xcf, 0x31, 0xe5,
	0x37, 0xfd, 0x17, 0x74, 0xf0, 0xc0, 0xa0, 0x7e, 0xe3, 0xc0, 0x2f, 0x57, 0xff, 0xad, 0x2b, 0xaa,
	0xff, 0x29, 0xb0, 0x4b, 0xe2, 0xb1, 0x92, 0x38, 0xe1, 0xdf, 0x51, 0xe9, 0x5c, 0x0a, 0xf3, 0xd6,
	0x11, 0xec, 0x19, 0xdc, 0xbc, 0x24, 0x4f, 0x84, 0x89, 0x55, 0x2e, 0x2d, 0xe7, 0xb4, 0x63, 0x6b,
	0x9e, 0x7c, 0x5e, 0x70, 0xae, 0x94, 0x3f, 0xe6, 0xca, 0x46, 0xfc, 0xb6, 0x2f, 0x65, 0x5a, 0xb0,
	0xfb, 0xd0, 0xc2, 0x8b, 0xf8, 0x34, 0x92, 0x23, 0x0c, 0xdd, 0xfd, 0xda, 0x26, 0xb2, 0x59, 0x62,
	0xef, 0xf2, 0xb1, 0xfb, 0x7a, 0x9a, 0x78, 0xb2, 0xb0, 0x43, 0x55, 0xd2, 0x20, 0x84, 0xce, 0xbe,
	0x0f, 0x6b, 0x06, 0xf5, 0xb9, 0x88, 0x31, 0xcc, 0x4e, 0x95, 0x44, 0x7e, 0x87, 0x14, 0xad, 0x02,
	0x3c, 0x72, 0x98, 0xab, 0x37, 0xba, 0x50, 0xe1, 0x30, 0x92, 0x26, 0xb4, 0xd1, 0x19, 0x26, 0x7c,
	0x97, 0x5e, 0xb5, 0x4e, 0xf8, 0xcb, 0x48, 0x9a, 0x63, 0x87, 0xba, 0xfc, 0xce, 0x29, 0x73, 0x83,
	0x09, 0xbf, 0xeb, 0xf3, 0x3b, 0x15, 0x0e, 0x0c, 0x26, 0xae, 0x61, 0x78, 0x9d, 0x0f, 0x76, 0x8f,
	0x34, 0x40, 0x90, 0x0f, 0xb4, 0x0b, 0x7e, 0xe5, 0x63, 0xb4, 0x89, 0x6f, 0x10, 0x42, 0xfb, 0x77,
	0x01, 0x8a, 0x16, 0x91, 0xeb, 0x94, 0xdf, 0xf7, 0xae, 0x3c, 0x32, 0xd0, 0x29, 0x6b, 0x43, 0x33,
	0x41, 0x13, 0x6b, 0x91, 0x59, 0xa1, 0x24, 0xdf, 0x23, 0x7e, 0x1e, 0x9a, 0x6b, 0x7b, 0xfb, 0x97,
	0xda, 0xde, 0x7d, 0x68, 0xf9, 0xf2, 0x09, 0x7d, 0x97, 0x79, 0xe0, 0xb7, 0x7a, 0xec, 0x98, 0x7a,
	0x4d, 0x07, 0x36, 0x0b, 0x89, 0xc9, 0x4f, 0x0a, 0xd9, 0x43, 0x7f, 0xfb, 0x3c, 0xde, 0xcf, 0x4f,
	0xbc, 0x72, 0x16, 0x2c, 0x56, 0xa9, 0xd2, 0xfc, 0xd1, 0x7c, 0xb0, 0x43, 0x07, 0x39, 0x89, 0xc6,
	0x18, 0xa7, 0xfd, 0xe4, 0x7b, 0x2f, 0x29, 0x30, 0xfa, 0x44, 0xce, 0xab, 0xc6, 0xc8, 0x22, 0xdd,
	0xf3, 0x4e, 0xe1, 0xd5, 0x23, 0x07, 0xd6, 0xd1, 0x79, 0x96, 0x94, 0xf4, 0x0f, 0x9e, 0x2e, 0x90,
	0x03, 0xcb, 0x7e, 0x83, 0x55, 0x4b, 0x23, 0xc8, 0xf0, 0x1f, 0xdb, 0x95, 0x4e, 0xb3, 0xb7, 0xdf,
	0xfd, 0xe2, 0x40, 0xeb, 0x1e, 0xce, 0x8d, 0xab, 0xa0, 0xdc, 0xc3, 0x7e, 0x85, 0xd5, 0x91, 0x1b,
	0x6a, 0x68, 0xf8, 0x63, 0xda, 0xbe, 0x77, 0xed, 0x76, 0x1a, 0x80, 0x41, 0xb9, 0xa5, 0x1c, 0x45,
	0x4f, 0x66, 0xa3, 0xc8, 0x37, 0x93, 0xc2, 0x1e, 0x7f, 0x5a, 0x36, 0x93, 0xc0, 0x03, 0x7b, 0xff,
	0x2c, 0xc3, 0xba, 0x8f, 0x14, 0xa0, 0xc9, 0x94, 0x34, 0xc8, 0x7e, 0x81, 0x9a, 0xbb, 0x0b, 0x76,
	0x42, 0x13, 0xa2, 0xd9, 0xdb, 0xbd, 0xf6, 0x00, 0x41, 0x21, 0x66, 0x3d, 0x3f, 0x5a, 0x35, 0x0d,
	0xd0, 0x66, 0xef, 0xce, 0x15, 0xbb, 0x8e, 0x9c, 0xc6, 0x0f, 0x5e, 0xcd, 0x9e, 0x41, 0x55, 0x58,
	0x1c, 0x1b, 0x5e, 0x21, 0xab, 0xff, 0xf3, 0x26, 0xaf, 0x75, 0x2f, 0x42, 0xad, 0x95, 0xa6, 0x81,
	0x74, 0xf5, 0x8b, 0x5e, 0x38, 0x4d, 0xe0, 0xa5, 0xec, 0x27, 0x58, 0x11, 0x72, 0xa8, 0x68, 0x5a,
	0x35, 0x7b, 0x3b, 0x57, 0x6c, 0x79, 0x2d, 0x87, 0x2a, 0x20, 0x61, 0xef, 0xaf, 0x15, 0xd8, 0x78,
	0x3b, 0xf1, 0x2f, 0xee, 0xfb, 0xab, 0xc9, 0xde, 0x42, 0xe5, 0x15, 0x5a, 0x76, 0xfd, 0x29, 0xb7,
	0x1f, 0x5e, 0x6f, 0xa2, 0xc8, 0xf2, 0xde, 0x37, 0xec, 0x08, 0x6a, 0xcf, 0xd1, 0x46, 0x22, 0x5d,
	0x58, 0xc4, 0x77, 0xb0, 0xf2, 0x46, 0x98, 0xc5, 0x9d, 0xf0, 0x0f, 0xd8, 0x18, 0xc8, 0x58, 0xa5,
	0x29, 0xc6, 0x16, 0x13, 0x0a, 0x7d, 0x65, 0xb6, 0xc7, 0x99, 0x9d, 0x7c, 0x7d, 0xe4, 0x3e, 0x34,
	0xde, 0x47, 0xa9, 0x48, 0x16, 0x7a, 0xdc, 0xf7, 0xb0, 0x5a, 0x94, 0x35, 0xbb, 0xfe, 0xd2, 0xd0,
	0x6f, 0xc1, 0xaf, 0x8e, 0xdb, 0xfb, 0xb7, 0x0a, 0x6b, 0x97, 0x2b, 0xe1, 0x08, 0x6a, 0x87, 0xd4,
	0x0f, 0x16, 0x59, 0x0c, 0x03, 0x6a, 0x21, 0x0b, 0x8b, 0x38, 0x70, 0xe5, 0x45, 0x93, 0x7a, 0x91,
	0xc9, 0x70, 0x61, 0x8b, 0x5f, 0x07, 0x0b, 0x0d, 0xbb, 0xe0, 0xbb, 0x15, 0xcc, 0xd7, 0xd7, 0x82,
	0x6a, 0x76, 0x00, 0xb5, 0x3e, 0x46, 0x3a, 0x3e, 0x5d, 0xac, 0xf3, 0x3e, 0xac, 0xbc, 0x5a, 0x74,
	0xc9, 0x9e, 0xd4, 0xe8, 0xef, 0xcf, 0xb3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xc9, 0x73,
	0x48, 0x5a, 0x0d, 0x00, 0x00,
}
