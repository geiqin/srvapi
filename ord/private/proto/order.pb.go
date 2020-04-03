// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

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

type Order struct {
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	OrderSn              string         `protobuf:"bytes,2,opt,name=order_sn,json=orderSn,proto3" json:"order_sn"`
	OrderType            string         `protobuf:"bytes,3,opt,name=order_type,json=orderType,proto3" json:"order_type"`
	CustomerId           int64          `protobuf:"varint,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	TotalNum             int32          `protobuf:"varint,6,opt,name=total_num,json=totalNum,proto3" json:"total_num"`
	TotalWeight          float32        `protobuf:"fixed32,7,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight"`
	Currency             string         `protobuf:"bytes,8,opt,name=currency,proto3" json:"currency"`
	Amount               float32        `protobuf:"fixed32,9,opt,name=amount,proto3" json:"amount"`
	BeforeAmount         float32        `protobuf:"fixed32,10,opt,name=before_amount,json=beforeAmount,proto3" json:"before_amount"`
	ActualAmount         float32        `protobuf:"fixed32,11,opt,name=actual_amount,json=actualAmount,proto3" json:"actual_amount"`
	AmountPaid           float32        `protobuf:"fixed32,12,opt,name=amount_paid,json=amountPaid,proto3" json:"amount_paid"`
	AmountRefunded       float32        `protobuf:"fixed32,13,opt,name=amount_refunded,json=amountRefunded,proto3" json:"amount_refunded"`
	ExpressAmount        float32        `protobuf:"fixed32,14,opt,name=express_amount,json=expressAmount,proto3" json:"express_amount"`
	DiscountAmount       float32        `protobuf:"fixed32,15,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount"`
	Paid                 bool           `protobuf:"varint,16,opt,name=paid,proto3" json:"paid"`
	Refunded             bool           `protobuf:"varint,17,opt,name=refunded,proto3" json:"refunded"`
	CanDelivery          bool           `protobuf:"varint,18,opt,name=can_delivery,json=canDelivery,proto3" json:"can_delivery"`
	CanCod               bool           `protobuf:"varint,19,opt,name=can_cod,json=canCod,proto3" json:"can_cod"`
	Invoiced             bool           `protobuf:"varint,20,opt,name=invoiced,proto3" json:"invoiced"`
	Modified             bool           `protobuf:"varint,21,opt,name=modified,proto3" json:"modified"`
	Safeguarded          bool           `protobuf:"varint,22,opt,name=safeguarded,proto3" json:"safeguarded"`
	Returnable           bool           `protobuf:"varint,23,opt,name=returnable,proto3" json:"returnable"`
	Subject              string         `protobuf:"bytes,24,opt,name=subject,proto3" json:"subject"`
	Body                 string         `protobuf:"bytes,25,opt,name=body,proto3" json:"body"`
	PayType              string         `protobuf:"bytes,26,opt,name=pay_type,json=payType,proto3" json:"pay_type"`
	BranchId             int64          `protobuf:"varint,27,opt,name=branch_id,json=branchId,proto3" json:"branch_id"`
	ClientIp             string         `protobuf:"bytes,28,opt,name=client_ip,json=clientIp,proto3" json:"client_ip"`
	Flag                 int32          `protobuf:"varint,29,opt,name=flag,proto3" json:"flag"`
	Memo                 string         `protobuf:"bytes,30,opt,name=memo,proto3" json:"memo"`
	BuyerMemo            string         `protobuf:"bytes,31,opt,name=buyer_memo,json=buyerMemo,proto3" json:"buyer_memo"`
	CancelId             int32          `protobuf:"varint,32,opt,name=cancel_id,json=cancelId,proto3" json:"cancel_id"`
	CancelReason         string         `protobuf:"bytes,33,opt,name=cancel_reason,json=cancelReason,proto3" json:"cancel_reason"`
	Status               string         `protobuf:"bytes,34,opt,name=status,proto3" json:"status"`
	PayStatus            string         `protobuf:"bytes,35,opt,name=pay_status,json=payStatus,proto3" json:"pay_status"`
	RefundStatus         string         `protobuf:"bytes,36,opt,name=refund_status,json=refundStatus,proto3" json:"refund_status"`
	ExpireAt             string         `protobuf:"bytes,38,opt,name=expire_at,json=expireAt,proto3" json:"expire_at"`
	PaidAt               string         `protobuf:"bytes,39,opt,name=paid_at,json=paidAt,proto3" json:"paid_at"`
	ShippedAt            string         `protobuf:"bytes,40,opt,name=shipped_at,json=shippedAt,proto3" json:"shipped_at"`
	SignedAt             string         `protobuf:"bytes,41,opt,name=signed_at,json=signedAt,proto3" json:"signed_at"`
	RefundedAt           string         `protobuf:"bytes,42,opt,name=refunded_at,json=refundedAt,proto3" json:"refunded_at"`
	FinishedAt           string         `protobuf:"bytes,43,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at"`
	CreatedAt            string         `protobuf:"bytes,44,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string         `protobuf:"bytes,45,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PlatformSource       string         `protobuf:"bytes,46,opt,name=platform_source,json=platformSource,proto3" json:"platform_source"`
	Address              *OrderAddress  `protobuf:"bytes,54,opt,name=address,proto3" json:"address,omitempty"`
	Details              []*OrderDetail `protobuf:"bytes,55,rep,name=details,proto3" json:"details,omitempty"`
	Customer             *Customer      `protobuf:"bytes,56,opt,name=customer,proto3" json:"customer,omitempty"`
	Delivery             []*Delivery    `protobuf:"bytes,57,rep,name=delivery,proto3" json:"delivery,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetOrderSn() string {
	if m != nil {
		return m.OrderSn
	}
	return ""
}

func (m *Order) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *Order) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Order) GetTotalNum() int32 {
	if m != nil {
		return m.TotalNum
	}
	return 0
}

func (m *Order) GetTotalWeight() float32 {
	if m != nil {
		return m.TotalWeight
	}
	return 0
}

func (m *Order) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Order) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Order) GetBeforeAmount() float32 {
	if m != nil {
		return m.BeforeAmount
	}
	return 0
}

func (m *Order) GetActualAmount() float32 {
	if m != nil {
		return m.ActualAmount
	}
	return 0
}

func (m *Order) GetAmountPaid() float32 {
	if m != nil {
		return m.AmountPaid
	}
	return 0
}

func (m *Order) GetAmountRefunded() float32 {
	if m != nil {
		return m.AmountRefunded
	}
	return 0
}

func (m *Order) GetExpressAmount() float32 {
	if m != nil {
		return m.ExpressAmount
	}
	return 0
}

func (m *Order) GetDiscountAmount() float32 {
	if m != nil {
		return m.DiscountAmount
	}
	return 0
}

func (m *Order) GetPaid() bool {
	if m != nil {
		return m.Paid
	}
	return false
}

func (m *Order) GetRefunded() bool {
	if m != nil {
		return m.Refunded
	}
	return false
}

func (m *Order) GetCanDelivery() bool {
	if m != nil {
		return m.CanDelivery
	}
	return false
}

func (m *Order) GetCanCod() bool {
	if m != nil {
		return m.CanCod
	}
	return false
}

func (m *Order) GetInvoiced() bool {
	if m != nil {
		return m.Invoiced
	}
	return false
}

func (m *Order) GetModified() bool {
	if m != nil {
		return m.Modified
	}
	return false
}

func (m *Order) GetSafeguarded() bool {
	if m != nil {
		return m.Safeguarded
	}
	return false
}

func (m *Order) GetReturnable() bool {
	if m != nil {
		return m.Returnable
	}
	return false
}

func (m *Order) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Order) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Order) GetPayType() string {
	if m != nil {
		return m.PayType
	}
	return ""
}

func (m *Order) GetBranchId() int64 {
	if m != nil {
		return m.BranchId
	}
	return 0
}

func (m *Order) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *Order) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *Order) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Order) GetBuyerMemo() string {
	if m != nil {
		return m.BuyerMemo
	}
	return ""
}

func (m *Order) GetCancelId() int32 {
	if m != nil {
		return m.CancelId
	}
	return 0
}

func (m *Order) GetCancelReason() string {
	if m != nil {
		return m.CancelReason
	}
	return ""
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Order) GetPayStatus() string {
	if m != nil {
		return m.PayStatus
	}
	return ""
}

func (m *Order) GetRefundStatus() string {
	if m != nil {
		return m.RefundStatus
	}
	return ""
}

func (m *Order) GetExpireAt() string {
	if m != nil {
		return m.ExpireAt
	}
	return ""
}

func (m *Order) GetPaidAt() string {
	if m != nil {
		return m.PaidAt
	}
	return ""
}

func (m *Order) GetShippedAt() string {
	if m != nil {
		return m.ShippedAt
	}
	return ""
}

func (m *Order) GetSignedAt() string {
	if m != nil {
		return m.SignedAt
	}
	return ""
}

func (m *Order) GetRefundedAt() string {
	if m != nil {
		return m.RefundedAt
	}
	return ""
}

func (m *Order) GetFinishedAt() string {
	if m != nil {
		return m.FinishedAt
	}
	return ""
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Order) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Order) GetPlatformSource() string {
	if m != nil {
		return m.PlatformSource
	}
	return ""
}

func (m *Order) GetAddress() *OrderAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Order) GetDetails() []*OrderDetail {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Order) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

func (m *Order) GetDelivery() []*Delivery {
	if m != nil {
		return m.Delivery
	}
	return nil
}

type Customer struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CustomerSn           string   `protobuf:"bytes,2,opt,name=customer_sn,json=customerSn,proto3" json:"customer_sn"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	DisplayName          string   `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name"`
	Realname             string   `protobuf:"bytes,5,opt,name=realname,proto3" json:"realname"`
	Idcard               string   `protobuf:"bytes,6,opt,name=idcard,proto3" json:"idcard"`
	Gender               int32    `protobuf:"varint,7,opt,name=gender,proto3" json:"gender"`
	Mobile               string   `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile"`
	Email                string   `protobuf:"bytes,9,opt,name=email,proto3" json:"email"`
	AvatarId             int64    `protobuf:"varint,10,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id"`
	AvatarUrl            string   `protobuf:"bytes,11,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer) GetCustomerSn() string {
	if m != nil {
		return m.CustomerSn
	}
	return ""
}

func (m *Customer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Customer) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Customer) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *Customer) GetIdcard() string {
	if m != nil {
		return m.Idcard
	}
	return ""
}

func (m *Customer) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *Customer) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Customer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Customer) GetAvatarId() int64 {
	if m != nil {
		return m.AvatarId
	}
	return 0
}

func (m *Customer) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type OrderWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting              string   `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Keywords             string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	OrderSn              string   `protobuf:"bytes,5,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
	OrderType            string   `protobuf:"bytes,6,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
	PayType              string   `protobuf:"bytes,7,opt,name=pay_type,json=payType,proto3" json:"pay_type,omitempty"`
	CanDelivery          bool     `protobuf:"varint,8,opt,name=can_delivery,json=canDelivery,proto3" json:"can_delivery,omitempty"`
	CanCod               bool     `protobuf:"varint,9,opt,name=can_cod,json=canCod,proto3" json:"can_cod,omitempty"`
	CustomerId           int64    `protobuf:"varint,10,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Payment              float32  `protobuf:"fixed32,11,opt,name=payment,proto3" json:"payment,omitempty"`
	IsFree               bool     `protobuf:"varint,13,opt,name=is_free,json=isFree,proto3" json:"is_free,omitempty"`
	Status               string   `protobuf:"bytes,26,opt,name=status,proto3" json:"status,omitempty"`
	PayStatus            string   `protobuf:"bytes,27,opt,name=pay_status,json=payStatus,proto3" json:"pay_status,omitempty"`
	RefundStatus         string   `protobuf:"bytes,28,opt,name=refund_status,json=refundStatus,proto3" json:"refund_status,omitempty"`
	DeliverStatus        string   `protobuf:"bytes,29,opt,name=deliver_status,json=deliverStatus,proto3" json:"deliver_status,omitempty"`
	BranchId             int32    `protobuf:"varint,32,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Flag                 int32    `protobuf:"varint,41,opt,name=flag,proto3" json:"flag,omitempty"`
	Safeguarded          bool     `protobuf:"varint,43,opt,name=safeguarded,proto3" json:"safeguarded,omitempty"`
	StartAt              string   `protobuf:"bytes,51,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                string   `protobuf:"bytes,52,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderWhere) Reset()         { *m = OrderWhere{} }
func (m *OrderWhere) String() string { return proto.CompactTextString(m) }
func (*OrderWhere) ProtoMessage()    {}
func (*OrderWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *OrderWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderWhere.Unmarshal(m, b)
}
func (m *OrderWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderWhere.Marshal(b, m, deterministic)
}
func (m *OrderWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderWhere.Merge(m, src)
}
func (m *OrderWhere) XXX_Size() int {
	return xxx_messageInfo_OrderWhere.Size(m)
}
func (m *OrderWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderWhere.DiscardUnknown(m)
}

var xxx_messageInfo_OrderWhere proto.InternalMessageInfo

func (m *OrderWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *OrderWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *OrderWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *OrderWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *OrderWhere) GetOrderSn() string {
	if m != nil {
		return m.OrderSn
	}
	return ""
}

func (m *OrderWhere) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *OrderWhere) GetPayType() string {
	if m != nil {
		return m.PayType
	}
	return ""
}

func (m *OrderWhere) GetCanDelivery() bool {
	if m != nil {
		return m.CanDelivery
	}
	return false
}

func (m *OrderWhere) GetCanCod() bool {
	if m != nil {
		return m.CanCod
	}
	return false
}

func (m *OrderWhere) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *OrderWhere) GetPayment() float32 {
	if m != nil {
		return m.Payment
	}
	return 0
}

func (m *OrderWhere) GetIsFree() bool {
	if m != nil {
		return m.IsFree
	}
	return false
}

func (m *OrderWhere) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *OrderWhere) GetPayStatus() string {
	if m != nil {
		return m.PayStatus
	}
	return ""
}

func (m *OrderWhere) GetRefundStatus() string {
	if m != nil {
		return m.RefundStatus
	}
	return ""
}

func (m *OrderWhere) GetDeliverStatus() string {
	if m != nil {
		return m.DeliverStatus
	}
	return ""
}

func (m *OrderWhere) GetBranchId() int32 {
	if m != nil {
		return m.BranchId
	}
	return 0
}

func (m *OrderWhere) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *OrderWhere) GetSafeguarded() bool {
	if m != nil {
		return m.Safeguarded
	}
	return false
}

func (m *OrderWhere) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *OrderWhere) GetEndAt() string {
	if m != nil {
		return m.EndAt
	}
	return ""
}

type OrderResponse struct {
	Entity               *Order   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Order `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderResponse) Reset()         { *m = OrderResponse{} }
func (m *OrderResponse) String() string { return proto.CompactTextString(m) }
func (*OrderResponse) ProtoMessage()    {}
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

func (m *OrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderResponse.Unmarshal(m, b)
}
func (m *OrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderResponse.Marshal(b, m, deterministic)
}
func (m *OrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResponse.Merge(m, src)
}
func (m *OrderResponse) XXX_Size() int {
	return xxx_messageInfo_OrderResponse.Size(m)
}
func (m *OrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResponse proto.InternalMessageInfo

func (m *OrderResponse) GetEntity() *Order {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *OrderResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *OrderResponse) GetItems() []*Order {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *OrderResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *OrderResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "geiqin.srv.ord.private.Order")
	proto.RegisterType((*Customer)(nil), "geiqin.srv.ord.private.Customer")
	proto.RegisterType((*OrderWhere)(nil), "geiqin.srv.ord.private.OrderWhere")
	proto.RegisterType((*OrderResponse)(nil), "geiqin.srv.ord.private.OrderResponse")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 1478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x58, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0xfe, 0x7d, 0xd0, 0x69, 0x25, 0x3b, 0x7f, 0xf6, 0xcf, 0x81, 0xb1, 0xe3, 0x44, 0xb1, 0xe3,
	0x44, 0xf9, 0xd3, 0x1a, 0x85, 0xd3, 0x23, 0xd0, 0x16, 0x50, 0xed, 0xb4, 0xf0, 0x85, 0x13, 0x97,
	0x4a, 0x90, 0x36, 0x28, 0x20, 0xac, 0xc8, 0x91, 0xb4, 0x0d, 0xb9, 0x64, 0x96, 0x4b, 0x27, 0xcc,
	0x8b, 0xb4, 0x8f, 0xd0, 0xab, 0x5e, 0xf6, 0x81, 0xfa, 0x24, 0xc5, 0xce, 0x2c, 0xe5, 0x43, 0x62,
	0x39, 0x05, 0xd4, 0xab, 0xde, 0x71, 0xbe, 0xf9, 0x66, 0x86, 0x5e, 0x7d, 0xfb, 0x0d, 0x61, 0xd6,
	0x4c, 0x74, 0x08, 0x7a, 0x2b, 0xd5, 0x89, 0x49, 0xf8, 0x95, 0x11, 0xc8, 0x97, 0x52, 0x6d, 0x65,
	0xfa, 0x70, 0x2b, 0xd1, 0xe1, 0x56, 0xaa, 0xe5, 0xa1, 0x30, 0xb0, 0xd2, 0x0a, 0x92, 0x38, 0x4e,
	0x14, 0xb1, 0x56, 0x5a, 0x83, 0xbc, 0x90, 0x6a, 0xe4, 0xa2, 0x8b, 0xd8, 0x60, 0x17, 0x8c, 0x90,
	0x91, 0x83, 0x38, 0x42, 0xdd, 0x30, 0xd4, 0x90, 0x65, 0x0e, 0x5b, 0x0e, 0x21, 0x92, 0x87, 0xa0,
	0x0b, 0x8a, 0xd7, 0x7f, 0x6f, 0xb1, 0xca, 0x63, 0x4b, 0xe3, 0xcb, 0x6c, 0x5e, 0x86, 0xde, 0x5c,
	0x7b, 0xae, 0xb3, 0xe0, 0xcf, 0xcb, 0x90, 0x5f, 0x63, 0x75, 0xac, 0xef, 0x67, 0xca, 0x9b, 0x6f,
	0xcf, 0x75, 0x1a, 0x7e, 0x0d, 0xe3, 0x9e, 0xe2, 0x6b, 0x8c, 0x51, 0xca, 0x14, 0x29, 0x78, 0x0b,
	0x98, 0x6c, 0x20, 0xf2, 0xa4, 0x48, 0x81, 0xdf, 0x64, 0xcd, 0x20, 0xcf, 0x4c, 0x12, 0x83, 0xee,
	0xcb, 0xd0, 0x5b, 0xc4, 0x96, 0xac, 0x84, 0xf6, 0x42, 0xbe, 0xca, 0x1a, 0x26, 0x31, 0x22, 0xea,
	0xab, 0x3c, 0xf6, 0xaa, 0xed, 0xb9, 0x4e, 0xc5, 0xaf, 0x23, 0xf0, 0x28, 0x8f, 0xf9, 0x2d, 0xd6,
	0xa2, 0xe4, 0x2b, 0x90, 0xa3, 0xb1, 0xf1, 0x6a, 0xed, 0xb9, 0xce, 0xbc, 0xdf, 0x44, 0xec, 0x19,
	0x42, 0x7c, 0x85, 0xd5, 0x83, 0x5c, 0x6b, 0x50, 0x41, 0xe1, 0xd5, 0x71, 0xfa, 0x24, 0xe6, 0x57,
	0x58, 0x55, 0xc4, 0x49, 0xae, 0x8c, 0xd7, 0xc0, 0x42, 0x17, 0xf1, 0x0d, 0xb6, 0x34, 0x80, 0x61,
	0xa2, 0xa1, 0xef, 0xd2, 0x0c, 0xd3, 0x2d, 0x02, 0xbb, 0x13, 0x92, 0x08, 0x4c, 0x2e, 0xa2, 0x92,
	0xd4, 0x24, 0x12, 0x81, 0x8e, 0x74, 0x93, 0x35, 0x29, 0xdb, 0x4f, 0x85, 0x0c, 0xbd, 0x16, 0x52,
	0x18, 0x41, 0x07, 0x42, 0x86, 0xfc, 0x2e, 0xbb, 0xe0, 0x08, 0x1a, 0x86, 0xb9, 0x0a, 0x21, 0xf4,
	0x96, 0x90, 0xb4, 0x4c, 0xb0, 0xef, 0x50, 0xbe, 0xc9, 0x96, 0xe1, 0x75, 0x6a, 0x7f, 0x9d, 0x72,
	0xde, 0x32, 0xf2, 0x96, 0x1c, 0xea, 0x06, 0xde, 0x65, 0x17, 0x42, 0x99, 0x05, 0xd8, 0xd1, 0xf1,
	0x2e, 0x50, 0xbf, 0x12, 0x76, 0x44, 0xce, 0x16, 0xf1, 0x95, 0xfe, 0xdb, 0x9e, 0xeb, 0xd4, 0x7d,
	0x7c, 0xb6, 0x67, 0x35, 0x79, 0x8b, 0x8b, 0x88, 0x4f, 0x62, 0x7b, 0xd4, 0x81, 0x50, 0xfd, 0x52,
	0x12, 0x1e, 0xc7, 0x7c, 0x33, 0x10, 0x6a, 0xd7, 0x41, 0xfc, 0x2a, 0xab, 0x59, 0x4a, 0x90, 0x84,
	0xde, 0xff, 0x30, 0x5b, 0x0d, 0x84, 0xda, 0x49, 0xb0, 0xaf, 0x54, 0x87, 0x89, 0x0c, 0x20, 0xf4,
	0x2e, 0x51, 0xdf, 0x32, 0xb6, 0xb9, 0x38, 0x09, 0xe5, 0x50, 0x42, 0xe8, 0x5d, 0xa6, 0x5c, 0x19,
	0xf3, 0x36, 0x6b, 0x66, 0x62, 0x08, 0xa3, 0x5c, 0x68, 0xfb, 0x4a, 0x57, 0x68, 0xe4, 0x31, 0x88,
	0xdf, 0x60, 0x4c, 0x83, 0xc9, 0xb5, 0x12, 0x83, 0x08, 0xbc, 0xab, 0x48, 0x38, 0x86, 0x70, 0x8f,
	0xd5, 0xb2, 0x7c, 0xf0, 0x33, 0x04, 0xc6, 0xf3, 0x48, 0x97, 0x2e, 0xb4, 0x7f, 0xff, 0x20, 0x09,
	0x0b, 0xef, 0x1a, 0xc2, 0xf8, 0x6c, 0x65, 0x9c, 0x8a, 0x82, 0x94, 0xba, 0x42, 0xf4, 0x54, 0x14,
	0xa8, 0xd3, 0x55, 0xd6, 0x18, 0x68, 0xa1, 0x82, 0xb1, 0x55, 0xe9, 0x2a, 0xaa, 0xb4, 0x4e, 0x00,
	0x69, 0x34, 0x88, 0x24, 0x28, 0xd3, 0x97, 0xa9, 0x77, 0xdd, 0x89, 0x0c, 0x81, 0xbd, 0xd4, 0x0e,
	0x1a, 0x46, 0x62, 0xe4, 0xad, 0xa1, 0x76, 0xf1, 0xd9, 0x62, 0x31, 0xc4, 0x89, 0x77, 0x83, 0x86,
	0xdb, 0x67, 0x7b, 0x51, 0x06, 0x79, 0x01, 0xba, 0x8f, 0x99, 0x9b, 0x74, 0x51, 0x10, 0xd9, 0xb7,
	0x69, 0x3b, 0x43, 0xa8, 0x00, 0x22, 0xfb, 0x02, 0x6d, 0xba, 0x07, 0x04, 0xec, 0x85, 0x56, 0x8b,
	0x2e, 0xa9, 0x41, 0x64, 0x89, 0xf2, 0x6e, 0x61, 0x79, 0x8b, 0x40, 0x1f, 0x31, 0xab, 0xf6, 0xcc,
	0x08, 0x93, 0x67, 0xde, 0x3a, 0x66, 0x5d, 0x64, 0x07, 0xdb, 0xbf, 0xda, 0xe5, 0x36, 0x68, 0x70,
	0x2a, 0x8a, 0x1e, 0xa5, 0x37, 0xd8, 0x12, 0x89, 0xa0, 0x64, 0xdc, 0xa6, 0xde, 0x04, 0x3a, 0xd2,
	0x2a, 0x6b, 0xc0, 0xeb, 0x54, 0xda, 0x1b, 0x63, 0xbc, 0x3b, 0x74, 0x02, 0x04, 0x74, 0x8d, 0xd5,
	0x85, 0x95, 0x97, 0x4d, 0xdd, 0xa5, 0xc9, 0x36, 0xec, 0x1a, 0x3b, 0x39, 0x1b, 0xcb, 0x34, 0x05,
	0xcc, 0x75, 0x68, 0xb2, 0x43, 0xba, 0xc6, 0x36, 0xcd, 0xe4, 0x48, 0x51, 0xf6, 0x1e, 0x35, 0x25,
	0xa0, 0x8b, 0x37, 0xab, 0xd4, 0xa6, 0x4d, 0xff, 0x1f, 0xd3, 0xac, 0x84, 0x88, 0x30, 0x94, 0x4a,
	0x66, 0x63, 0x22, 0xdc, 0x27, 0x42, 0x09, 0xd1, 0xf4, 0x40, 0x83, 0x30, 0x94, 0xff, 0x80, 0xa6,
	0x3b, 0x84, 0xd2, 0x79, 0x1a, 0x96, 0xe9, 0x0f, 0x29, 0xed, 0x90, 0x2e, 0x5e, 0xb4, 0x34, 0x12,
	0x66, 0x98, 0xe8, 0xb8, 0x9f, 0x25, 0xb9, 0x0e, 0xc0, 0xdb, 0x42, 0xce, 0x72, 0x09, 0xf7, 0x10,
	0xe5, 0x5f, 0xb3, 0x9a, 0x20, 0x5b, 0xf5, 0x3e, 0x6d, 0xcf, 0x75, 0x9a, 0xdb, 0xb7, 0xb7, 0xde,
	0x6d, 0xd9, 0x5b, 0x8f, 0x8f, 0x59, 0xb0, 0x5f, 0x16, 0xf1, 0xaf, 0x58, 0x2d, 0x44, 0xa7, 0xce,
	0xbc, 0xcf, 0xda, 0x0b, 0x9d, 0xe6, 0xf6, 0xc6, 0xd4, 0x7a, 0x72, 0x75, 0xbf, 0xac, 0xe1, 0x5f,
	0x5a, 0xff, 0x23, 0x37, 0xf5, 0x3e, 0xc7, 0xf9, 0xed, 0xb3, 0xea, 0x77, 0x1c, 0xcf, 0x9f, 0x54,
	0xd8, 0xea, 0xc9, 0x8d, 0xff, 0x02, 0xa7, 0x9f, 0x59, 0x5d, 0xda, 0x80, 0x3f, 0xa9, 0x58, 0xff,
	0x6d, 0x9e, 0xd5, 0xcb, 0xa6, 0x6f, 0xed, 0x8c, 0xe3, 0xce, 0x3f, 0x59, 0x1b, 0x13, 0xe7, 0xef,
	0x29, 0x7b, 0x49, 0x94, 0x88, 0xcb, 0x9d, 0x81, 0xcf, 0xd6, 0x85, 0x42, 0x99, 0xa5, 0x91, 0x28,
	0xfa, 0x98, 0x5b, 0xc4, 0x5c, 0xd3, 0x61, 0x8f, 0x2c, 0x05, 0x4d, 0x4c, 0x44, 0x98, 0xae, 0x90,
	0x68, 0xca, 0xd8, 0x5e, 0x01, 0x19, 0x06, 0x42, 0x87, 0xb8, 0x49, 0x1a, 0xbe, 0x8b, 0x2c, 0x3e,
	0x02, 0x15, 0x82, 0xc6, 0x0d, 0x52, 0xf1, 0x5d, 0x64, 0xf1, 0x38, 0x19, 0xc8, 0x08, 0xdc, 0xea,
	0x70, 0x11, 0xbf, 0xc4, 0x2a, 0x10, 0x0b, 0x19, 0xe1, 0xde, 0x68, 0xf8, 0x14, 0x58, 0xbd, 0x8a,
	0x43, 0x61, 0x04, 0x6e, 0x32, 0x46, 0x1e, 0x41, 0xc0, 0x5e, 0x68, 0xe5, 0xe4, 0x92, 0xb9, 0x8e,
	0x70, 0x57, 0x34, 0x7c, 0x47, 0x7f, 0xaa, 0xa3, 0xf5, 0x3f, 0x17, 0x19, 0xc3, 0xdf, 0xef, 0xd9,
	0x18, 0x34, 0x0e, 0x48, 0xc5, 0x08, 0xe8, 0xbc, 0x2a, 0x3e, 0x05, 0x76, 0x80, 0x7d, 0xe8, 0x67,
	0xf2, 0x0d, 0xe0, 0x81, 0x55, 0xfc, 0xba, 0x05, 0x7a, 0xf2, 0x0d, 0x59, 0x5d, 0xa2, 0x8d, 0x54,
	0x23, 0x77, 0x62, 0x65, 0x68, 0x4f, 0xe4, 0x05, 0x14, 0xaf, 0x12, 0x1d, 0x66, 0xee, 0xc0, 0x26,
	0xf1, 0x89, 0xcd, 0x5d, 0x99, 0xb6, 0xb9, 0xab, 0xa7, 0x37, 0xf7, 0x71, 0xb3, 0xac, 0x9d, 0x34,
	0xcb, 0xd3, 0xbb, 0xa2, 0x3e, 0x75, 0x57, 0x34, 0x4e, 0xec, 0x8a, 0x53, 0x1f, 0x04, 0xec, 0xad,
	0x0f, 0x02, 0xcf, 0xba, 0x49, 0x11, 0xc3, 0x64, 0xe3, 0x96, 0xa1, 0xed, 0x29, 0xb3, 0xfe, 0x50,
	0x03, 0xe0, 0x0e, 0xad, 0xfb, 0x55, 0x99, 0x7d, 0xab, 0x01, 0x8e, 0x39, 0xdf, 0xca, 0x14, 0xe7,
	0x5b, 0x3d, 0xd7, 0xf9, 0xae, 0xbf, 0xc3, 0xf9, 0x36, 0x59, 0xf9, 0x99, 0x54, 0xb2, 0xd6, 0x90,
	0xb5, 0xe4, 0xd0, 0x23, 0x83, 0x3c, 0xda, 0x1f, 0xce, 0xbe, 0x27, 0xfb, 0xa3, 0x5c, 0x11, 0xf7,
	0x8e, 0xad, 0x88, 0x53, 0xbb, 0xef, 0xfe, 0xdb, 0xbb, 0xef, 0x1a, 0xab, 0x67, 0x46, 0x68, 0x63,
	0xed, 0xe9, 0x81, 0xfb, 0xc5, 0x6d, 0xdc, 0x35, 0xfc, 0x32, 0xab, 0x82, 0x42, 0xdf, 0xfa, 0xd8,
	0x09, 0x54, 0x85, 0x5d, 0xb3, 0xfe, 0xcb, 0x3c, 0x5b, 0x42, 0x91, 0xf9, 0x90, 0xa5, 0x89, 0xca,
	0x80, 0x7f, 0x62, 0x89, 0x46, 0x9a, 0x02, 0x85, 0xd6, 0xdc, 0x5e, 0x9b, 0xea, 0x2d, 0xbe, 0x23,
	0xf3, 0x07, 0x24, 0x4f, 0x8d, 0x22, 0x9c, 0x52, 0x75, 0x60, 0x49, 0xa4, 0x5e, 0x6d, 0x8b, 0xa4,
	0x81, 0x38, 0xf3, 0x16, 0xd0, 0x48, 0xce, 0x19, 0x45, 0x5c, 0x5b, 0x04, 0x5a, 0x27, 0x1a, 0x85,
	0x3b, 0xa5, 0xe8, 0xa1, 0x25, 0xf9, 0xc4, 0xe5, 0x1f, 0xb1, 0x45, 0xa9, 0x86, 0x09, 0x0a, 0xba,
	0xb9, 0x7d, 0xfd, 0xac, 0x9a, 0x3d, 0x35, 0x4c, 0x7c, 0x64, 0x6e, 0xff, 0x51, 0x61, 0xcb, 0xfb,
	0x05, 0x4e, 0xee, 0x81, 0x3e, 0x94, 0x01, 0xf0, 0xe7, 0xac, 0xb6, 0x93, 0xa8, 0xa1, 0xd4, 0x31,
	0xdf, 0x3c, 0xab, 0xc3, 0x37, 0xf8, 0x55, 0xed, 0xc3, 0xcb, 0x1c, 0x32, 0xb3, 0x72, 0xe7, 0x3c,
	0x1a, 0x1d, 0xfa, 0xfa, 0x7f, 0xf8, 0x0f, 0xac, 0xda, 0xcb, 0x07, 0xb1, 0x34, 0xef, 0xdb, 0x7a,
	0x73, 0xfa, 0x61, 0x1d, 0x75, 0x3e, 0x60, 0xd5, 0x5d, 0x88, 0xc0, 0x00, 0x9f, 0x7e, 0xbe, 0x7f,
	0xab, 0xe3, 0x0e, 0x7e, 0x46, 0xcc, 0xac, 0xe3, 0xf7, 0xac, 0xe6, 0x43, 0x00, 0x32, 0x35, 0x33,
	0x6b, 0xb9, 0xcf, 0x16, 0xbe, 0x03, 0x33, 0xcb, 0x37, 0xdc, 0xa5, 0x95, 0x32, 0xb3, 0x96, 0x4f,
	0x59, 0xb5, 0x07, 0x42, 0x07, 0x63, 0xbe, 0x3e, 0xb5, 0x04, 0xfd, 0xff, 0xbd, 0xdb, 0x6e, 0xff,
	0x5a, 0x67, 0xad, 0x13, 0xb2, 0xfd, 0x91, 0x55, 0x77, 0xf0, 0x1b, 0x66, 0xf6, 0xaa, 0x7d, 0xcc,
	0x2a, 0x3b, 0x63, 0x08, 0x5e, 0xcc, 0xf0, 0x4c, 0x9a, 0xfb, 0xf6, 0x5b, 0xbf, 0x38, 0xd0, 0xf6,
	0xd5, 0x67, 0xd5, 0xf6, 0x19, 0x5b, 0xa2, 0xb6, 0xee, 0x5b, 0x6a, 0x96, 0x57, 0x61, 0xc6, 0x97,
	0xeb, 0xdf, 0xaa, 0x5b, 0xf2, 0x80, 0x58, 0xe8, 0x17, 0xb3, 0xfb, 0x75, 0x9e, 0xb0, 0xc5, 0xde,
	0x58, 0xa6, 0xfc, 0xdc, 0x2f, 0xd4, 0x95, 0xce, 0xb9, 0xdf, 0xb0, 0x47, 0x5d, 0x9f, 0x33, 0xe6,
	0x43, 0x2a, 0xa4, 0xfe, 0x07, 0x7a, 0xff, 0xc4, 0x5a, 0x0f, 0x5f, 0x07, 0x63, 0xa1, 0x46, 0x30,
	0xfb, 0xee, 0x83, 0x2a, 0xfe, 0xd7, 0xe6, 0xc1, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0x97,
	0x27, 0x18, 0x2f, 0x12, 0x00, 0x00,
}
