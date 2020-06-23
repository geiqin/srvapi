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
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderSn              string         `protobuf:"bytes,2,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
	OrderType            string         `protobuf:"bytes,3,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
	CustomerId           int64          `protobuf:"varint,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	TotalNum             int32          `protobuf:"varint,6,opt,name=total_num,json=totalNum,proto3" json:"total_num,omitempty"`
	TotalWeight          float32        `protobuf:"fixed32,7,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty"`
	Currency             string         `protobuf:"bytes,8,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float32        `protobuf:"fixed32,9,opt,name=amount,proto3" json:"amount,omitempty"`
	BeforeAmount         float32        `protobuf:"fixed32,10,opt,name=before_amount,json=beforeAmount,proto3" json:"before_amount,omitempty"`
	ActualAmount         float32        `protobuf:"fixed32,11,opt,name=actual_amount,json=actualAmount,proto3" json:"actual_amount,omitempty"`
	AmountPaid           float32        `protobuf:"fixed32,12,opt,name=amount_paid,json=amountPaid,proto3" json:"amount_paid,omitempty"`
	AmountRefunded       float32        `protobuf:"fixed32,13,opt,name=amount_refunded,json=amountRefunded,proto3" json:"amount_refunded,omitempty"`
	ExpressAmount        float32        `protobuf:"fixed32,14,opt,name=express_amount,json=expressAmount,proto3" json:"express_amount,omitempty"`
	DiscountAmount       float32        `protobuf:"fixed32,15,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty"`
	Paid                 bool           `protobuf:"varint,16,opt,name=paid,proto3" json:"paid,omitempty"`
	Refunded             bool           `protobuf:"varint,17,opt,name=refunded,proto3" json:"refunded,omitempty"`
	CanDelivery          bool           `protobuf:"varint,18,opt,name=can_delivery,json=canDelivery,proto3" json:"can_delivery,omitempty"`
	CanCod               bool           `protobuf:"varint,19,opt,name=can_cod,json=canCod,proto3" json:"can_cod,omitempty"`
	Invoiced             bool           `protobuf:"varint,20,opt,name=invoiced,proto3" json:"invoiced,omitempty"`
	Modified             bool           `protobuf:"varint,21,opt,name=modified,proto3" json:"modified,omitempty"`
	Safeguarded          bool           `protobuf:"varint,22,opt,name=safeguarded,proto3" json:"safeguarded,omitempty"`
	Returnable           bool           `protobuf:"varint,23,opt,name=returnable,proto3" json:"returnable,omitempty"`
	Subject              string         `protobuf:"bytes,24,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string         `protobuf:"bytes,25,opt,name=body,proto3" json:"body,omitempty"`
	PayType              string         `protobuf:"bytes,26,opt,name=pay_type,json=payType,proto3" json:"pay_type,omitempty"`
	BranchId             int64          `protobuf:"varint,27,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	ClientIp             string         `protobuf:"bytes,28,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	Flag                 int32          `protobuf:"varint,29,opt,name=flag,proto3" json:"flag,omitempty"`
	Memo                 string         `protobuf:"bytes,30,opt,name=memo,proto3" json:"memo,omitempty"`
	BuyerMemo            string         `protobuf:"bytes,31,opt,name=buyer_memo,json=buyerMemo,proto3" json:"buyer_memo,omitempty"`
	CancelId             int32          `protobuf:"varint,32,opt,name=cancel_id,json=cancelId,proto3" json:"cancel_id,omitempty"`
	CancelReason         string         `protobuf:"bytes,33,opt,name=cancel_reason,json=cancelReason,proto3" json:"cancel_reason,omitempty"`
	Status               string         `protobuf:"bytes,34,opt,name=status,proto3" json:"status,omitempty"`
	PayStatus            string         `protobuf:"bytes,35,opt,name=pay_status,json=payStatus,proto3" json:"pay_status,omitempty"`
	RefundStatus         string         `protobuf:"bytes,36,opt,name=refund_status,json=refundStatus,proto3" json:"refund_status,omitempty"`
	ExpireAt             string         `protobuf:"bytes,38,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	PaidAt               string         `protobuf:"bytes,39,opt,name=paid_at,json=paidAt,proto3" json:"paid_at,omitempty"`
	ShippedAt            string         `protobuf:"bytes,40,opt,name=shipped_at,json=shippedAt,proto3" json:"shipped_at,omitempty"`
	SignedAt             string         `protobuf:"bytes,41,opt,name=signed_at,json=signedAt,proto3" json:"signed_at,omitempty"`
	RefundedAt           string         `protobuf:"bytes,42,opt,name=refunded_at,json=refundedAt,proto3" json:"refunded_at,omitempty"`
	FinishedAt           string         `protobuf:"bytes,43,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	CreatedAt            string         `protobuf:"bytes,44,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string         `protobuf:"bytes,45,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PlatformSource       string         `protobuf:"bytes,46,opt,name=platform_source,json=platformSource,proto3" json:"platform_source,omitempty"`
	Address              *OrderAddress  `protobuf:"bytes,54,opt,name=address,proto3" json:"address,omitempty"`
	Details              []*OrderDetail `protobuf:"bytes,55,rep,name=details,proto3" json:"details,omitempty"`
	Customer             *Customer      `protobuf:"bytes,56,opt,name=customer,proto3" json:"customer,omitempty"`
	Delivery             []*Delivery    `protobuf:"bytes,57,rep,name=delivery,proto3" json:"delivery,omitempty"`
	ShipmentMethod       string         `protobuf:"bytes,58,opt,name=shipment_method,json=shipmentMethod,proto3" json:"shipment_method,omitempty"`
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

func (m *Order) GetShipmentMethod() string {
	if m != nil {
		return m.ShipmentMethod
	}
	return ""
}

type Customer struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerSn           string   `protobuf:"bytes,2,opt,name=customer_sn,json=customerSn,proto3" json:"customer_sn,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string   `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Realname             string   `protobuf:"bytes,5,opt,name=realname,proto3" json:"realname,omitempty"`
	Idcard               string   `protobuf:"bytes,6,opt,name=idcard,proto3" json:"idcard,omitempty"`
	Gender               int32    `protobuf:"varint,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Mobile               string   `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email                string   `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	AvatarId             int64    `protobuf:"varint,10,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,11,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
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
	IsFree               bool     `protobuf:"varint,12,opt,name=is_free,json=isFree,proto3" json:"is_free,omitempty"`
	Status               string   `protobuf:"bytes,13,opt,name=status,proto3" json:"status,omitempty"`
	PayStatus            string   `protobuf:"bytes,14,opt,name=pay_status,json=payStatus,proto3" json:"pay_status,omitempty"`
	RefundStatus         string   `protobuf:"bytes,15,opt,name=refund_status,json=refundStatus,proto3" json:"refund_status,omitempty"`
	DeliverStatus        string   `protobuf:"bytes,16,opt,name=deliver_status,json=deliverStatus,proto3" json:"deliver_status,omitempty"`
	BranchId             int32    `protobuf:"varint,17,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Flag                 int32    `protobuf:"varint,18,opt,name=flag,proto3" json:"flag,omitempty"`
	Safeguarded          bool     `protobuf:"varint,19,opt,name=safeguarded,proto3" json:"safeguarded,omitempty"`
	StartAt              string   `protobuf:"bytes,20,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                string   `protobuf:"bytes,21,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	Paid                 bool     `protobuf:"varint,22,opt,name=paid,proto3" json:"paid,omitempty"`
	Ids                  []int64  `protobuf:"varint,23,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	ShipmentMethod       string   `protobuf:"bytes,24,opt,name=shipment_method,json=shipmentMethod,proto3" json:"shipment_method,omitempty"`
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

func (m *OrderWhere) GetPaid() bool {
	if m != nil {
		return m.Paid
	}
	return false
}

func (m *OrderWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *OrderWhere) GetShipmentMethod() string {
	if m != nil {
		return m.ShipmentMethod
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

func init() {
	proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077)
}

var fileDescriptor_cd01338c35d87077 = []byte{
	// 1560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x58, 0xdb, 0x6e, 0x1b, 0xc9,
	0x11, 0x8d, 0x44, 0xf1, 0xd6, 0xbc, 0xc8, 0xee, 0xb5, 0xe5, 0xb6, 0xb4, 0x5e, 0x73, 0xa5, 0xd5,
	0x9a, 0xc9, 0x26, 0x42, 0xa0, 0x45, 0xae, 0x48, 0x02, 0x30, 0xd2, 0x26, 0x10, 0x10, 0xad, 0x94,
	0xe1, 0x1a, 0x4e, 0x8c, 0x00, 0x44, 0x73, 0xa6, 0x48, 0x76, 0x3c, 0x37, 0xf7, 0xf4, 0xc8, 0x1e,
	0xff, 0x48, 0x7e, 0x21, 0xaf, 0x79, 0xc9, 0x97, 0x04, 0xc8, 0xef, 0x04, 0x5d, 0xd5, 0x43, 0x51,
	0xb2, 0x44, 0x39, 0x00, 0xfd, 0xb4, 0x6f, 0x53, 0xe7, 0x9c, 0xaa, 0x1a, 0xf5, 0xd4, 0xa5, 0x45,
	0xd6, 0x4a, 0x74, 0x00, 0xfa, 0x20, 0xd5, 0x89, 0x49, 0xf8, 0xd6, 0x14, 0xd4, 0x6b, 0x15, 0x1f,
	0x64, 0xfa, 0xe2, 0x20, 0xd1, 0xc1, 0x41, 0xaa, 0xd5, 0x85, 0x34, 0xb0, 0xdd, 0xf6, 0x93, 0x28,
	0x4a, 0x62, 0x52, 0x6d, 0xb7, 0xc7, 0x79, 0xa1, 0xe2, 0xa9, 0xb3, 0xee, 0x63, 0x80, 0x63, 0x30,
	0x52, 0x85, 0x0e, 0xe2, 0x08, 0x0d, 0x82, 0x40, 0x43, 0x96, 0x39, 0xac, 0x1b, 0x40, 0xa8, 0x2e,
	0x40, 0x17, 0x57, 0xdc, 0x8e, 0x92, 0x3c, 0x2d, 0xe3, 0xee, 0xfe, 0xb7, 0xcd, 0xaa, 0x67, 0x16,
	0xe5, 0x5d, 0xb6, 0xae, 0x02, 0xb1, 0xd6, 0x5b, 0xeb, 0x57, 0xbc, 0x75, 0x15, 0xf0, 0xc7, 0xac,
	0x81, 0xf2, 0x51, 0x16, 0x8b, 0xf5, 0xde, 0x5a, 0xbf, 0xe9, 0xd5, 0xd1, 0x1e, 0xc6, 0xfc, 0x09,
	0x63, 0x44, 0x99, 0x22, 0x05, 0x51, 0x41, 0xb2, 0x89, 0xc8, 0x77, 0x45, 0x0a, 0xfc, 0x29, 0x6b,
	0xf9, 0x79, 0x66, 0x92, 0x08, 0xf4, 0x48, 0x05, 0x62, 0x03, 0x43, 0xb2, 0x12, 0x3a, 0x09, 0xf8,
	0x0e, 0x6b, 0x9a, 0xc4, 0xc8, 0x70, 0x14, 0xe7, 0x91, 0xa8, 0xf5, 0xd6, 0xfa, 0x55, 0xaf, 0x81,
	0xc0, 0xb7, 0x79, 0xc4, 0x3f, 0x67, 0x6d, 0x22, 0xdf, 0x80, 0x9a, 0xce, 0x8c, 0xa8, 0xf7, 0xd6,
	0xfa, 0xeb, 0x5e, 0x0b, 0xb1, 0x17, 0x08, 0xf1, 0x6d, 0xd6, 0xf0, 0x73, 0xad, 0x21, 0xf6, 0x0b,
	0xd1, 0xc0, 0xec, 0x73, 0x9b, 0x6f, 0xb1, 0x9a, 0x8c, 0x92, 0x3c, 0x36, 0xa2, 0x89, 0x8e, 0xce,
	0xe2, 0x7b, 0xac, 0x33, 0x86, 0x49, 0xa2, 0x61, 0xe4, 0x68, 0x86, 0x74, 0x9b, 0xc0, 0xc1, 0x5c,
	0x24, 0x7d, 0x93, 0xcb, 0xb0, 0x14, 0xb5, 0x48, 0x44, 0xa0, 0x13, 0x3d, 0x65, 0x2d, 0x62, 0x47,
	0xa9, 0x54, 0x81, 0x68, 0xa3, 0x84, 0x11, 0x74, 0x2e, 0x55, 0xc0, 0x9f, 0xb1, 0x4d, 0x27, 0xd0,
	0x30, 0xc9, 0xe3, 0x00, 0x02, 0xd1, 0x41, 0x51, 0x97, 0x60, 0xcf, 0xa1, 0x7c, 0x9f, 0x75, 0xe1,
	0x6d, 0x6a, 0x3f, 0x58, 0x99, 0xaf, 0x8b, 0xba, 0x8e, 0x43, 0x5d, 0xc2, 0x67, 0x6c, 0x33, 0x50,
	0x99, 0x8f, 0x11, 0x9d, 0x6e, 0x93, 0xe2, 0x95, 0xb0, 0x13, 0x72, 0xb6, 0x81, 0xaf, 0x74, 0xaf,
	0xb7, 0xd6, 0x6f, 0x78, 0xf8, 0x6c, 0xcf, 0x6a, 0xfe, 0x16, 0xf7, 0x11, 0x9f, 0xdb, 0xf6, 0xa8,
	0x7d, 0x19, 0x8f, 0xca, 0x2a, 0x11, 0x1c, 0xf9, 0x96, 0x2f, 0xe3, 0x63, 0x07, 0xf1, 0x47, 0xac,
	0x6e, 0x25, 0x7e, 0x12, 0x88, 0x4f, 0x90, 0xad, 0xf9, 0x32, 0x3e, 0x4a, 0x30, 0xae, 0x8a, 0x2f,
	0x12, 0xe5, 0x43, 0x20, 0x1e, 0x50, 0xdc, 0xd2, 0xb6, 0x5c, 0x94, 0x04, 0x6a, 0xa2, 0x20, 0x10,
	0x0f, 0x89, 0x2b, 0x6d, 0xde, 0x63, 0xad, 0x4c, 0x4e, 0x60, 0x9a, 0x4b, 0x6d, 0x5f, 0x69, 0x8b,
	0x52, 0x2e, 0x40, 0xfc, 0x33, 0xc6, 0x34, 0x98, 0x5c, 0xc7, 0x72, 0x1c, 0x82, 0x78, 0x84, 0x82,
	0x05, 0x84, 0x0b, 0x56, 0xcf, 0xf2, 0xf1, 0xdf, 0xc1, 0x37, 0x42, 0x50, 0x5d, 0x3a, 0xd3, 0xfe,
	0xfd, 0xe3, 0x24, 0x28, 0xc4, 0x63, 0x84, 0xf1, 0xd9, 0x96, 0x71, 0x2a, 0x0b, 0xaa, 0xd4, 0x6d,
	0x92, 0xa7, 0xb2, 0xc0, 0x3a, 0xdd, 0x61, 0xcd, 0xb1, 0x96, 0xb1, 0x3f, 0xb3, 0x55, 0xba, 0x83,
	0x55, 0xda, 0x20, 0x80, 0x6a, 0xd4, 0x0f, 0x15, 0xc4, 0x66, 0xa4, 0x52, 0xf1, 0xa9, 0x2b, 0x32,
	0x04, 0x4e, 0x52, 0x9b, 0x68, 0x12, 0xca, 0xa9, 0x78, 0x82, 0xb5, 0x8b, 0xcf, 0x16, 0x8b, 0x20,
	0x4a, 0xc4, 0x67, 0x94, 0xdc, 0x3e, 0xdb, 0x46, 0x19, 0xe7, 0x05, 0xe8, 0x11, 0x32, 0x4f, 0xa9,
	0x51, 0x10, 0x39, 0xb5, 0xb4, 0xcd, 0x21, 0x63, 0x1f, 0x42, 0xfb, 0x02, 0x3d, 0xea, 0x03, 0x02,
	0x4e, 0x02, 0x5b, 0x8b, 0x8e, 0xd4, 0x20, 0xb3, 0x24, 0x16, 0x9f, 0xa3, 0x7b, 0x9b, 0x40, 0x0f,
	0x31, 0x5b, 0xed, 0x99, 0x91, 0x26, 0xcf, 0xc4, 0x2e, 0xb2, 0xce, 0xb2, 0x89, 0xed, 0x5f, 0xed,
	0xb8, 0x3d, 0x4a, 0x9c, 0xca, 0x62, 0x48, 0xf4, 0x1e, 0xeb, 0x50, 0x11, 0x94, 0x8a, 0x2f, 0x28,
	0x36, 0x81, 0x4e, 0xb4, 0xc3, 0x9a, 0xf0, 0x36, 0x55, 0xb6, 0x63, 0x8c, 0xf8, 0x92, 0x4e, 0x80,
	0x80, 0x81, 0xb1, 0x75, 0x61, 0xcb, 0xcb, 0x52, 0xcf, 0x28, 0xb3, 0x35, 0x07, 0xc6, 0x66, 0xce,
	0x66, 0x2a, 0x4d, 0x01, 0xb9, 0x3e, 0x65, 0x76, 0xc8, 0xc0, 0xd8, 0xa0, 0x99, 0x9a, 0xc6, 0xc4,
	0xfe, 0x90, 0x82, 0x12, 0x30, 0xc0, 0xce, 0x2a, 0x6b, 0xd3, 0xd2, 0x3f, 0x42, 0x9a, 0x95, 0x10,
	0x09, 0x26, 0x2a, 0x56, 0xd9, 0x8c, 0x04, 0x5f, 0x91, 0xa0, 0x84, 0x28, 0xbb, 0xaf, 0x41, 0x1a,
	0xe2, 0x7f, 0x4c, 0xd9, 0x1d, 0x42, 0x74, 0x9e, 0x06, 0x25, 0xfd, 0x13, 0xa2, 0x1d, 0x32, 0xc0,
	0x46, 0x4b, 0x43, 0x69, 0x26, 0x89, 0x8e, 0x46, 0x59, 0x92, 0x6b, 0x1f, 0xc4, 0x01, 0x6a, 0xba,
	0x25, 0x3c, 0x44, 0x94, 0xff, 0x8e, 0xd5, 0x25, 0x4d, 0x5a, 0xf1, 0xf3, 0xde, 0x5a, 0xbf, 0x75,
	0xf8, 0xc5, 0xc1, 0xcd, 0x53, 0xfc, 0xe0, 0x6c, 0x61, 0x2a, 0x7b, 0xa5, 0x13, 0xff, 0x2d, 0xab,
	0x07, 0x38, 0xbc, 0x33, 0xf1, 0x8b, 0x5e, 0xa5, 0xdf, 0x3a, 0xdc, 0x5b, 0xea, 0x4f, 0x83, 0xde,
	0x2b, 0x7d, 0xf8, 0x6f, 0xec, 0xfc, 0xa3, 0x69, 0x2a, 0x7e, 0x89, 0xf9, 0x7b, 0xb7, 0xf9, 0x1f,
	0x39, 0x9d, 0x37, 0xf7, 0xb0, 0xde, 0xf3, 0x8e, 0xff, 0x15, 0x66, 0xbf, 0xd5, 0xbb, 0x1c, 0x03,
	0xde, 0xdc, 0xc3, 0x9e, 0x91, 0xfd, 0x9a, 0x91, 0xed, 0x8c, 0x08, 0xcc, 0x2c, 0x09, 0xc4, 0xaf,
	0xe9, 0x8c, 0x4a, 0xf8, 0x14, 0xd1, 0xdd, 0x7f, 0xae, 0xb3, 0x46, 0x99, 0xfd, 0xbd, 0xe5, 0xb2,
	0xb8, 0x22, 0xe6, 0xfb, 0x65, 0xbe, 0x22, 0x86, 0xb1, 0xed, 0xa6, 0x58, 0x46, 0xe5, 0x72, 0xc1,
	0x67, 0x3b, 0xae, 0x02, 0x95, 0xa5, 0xa1, 0x2c, 0x46, 0xc8, 0x6d, 0x20, 0xd7, 0x72, 0xd8, 0xb7,
	0x56, 0x82, 0xd3, 0x4e, 0x86, 0x48, 0x57, 0xa9, 0xba, 0x4a, 0xdb, 0xf6, 0x8a, 0x0a, 0x7c, 0xa9,
	0x03, 0x5c, 0x39, 0x4d, 0xcf, 0x59, 0x16, 0x9f, 0x42, 0x1c, 0x80, 0xc6, 0x55, 0x53, 0xf5, 0x9c,
	0x65, 0xf1, 0x28, 0x19, 0xab, 0x10, 0xdc, 0x8e, 0x71, 0x16, 0x7f, 0xc0, 0xaa, 0x10, 0x49, 0x15,
	0xe2, 0x82, 0x69, 0x7a, 0x64, 0xd8, 0xc2, 0x96, 0x17, 0xd2, 0x48, 0x5c, 0x79, 0x8c, 0x86, 0x09,
	0x01, 0x27, 0x81, 0xad, 0x3b, 0x47, 0xe6, 0x3a, 0xc4, 0xa5, 0xd2, 0xf4, 0x9c, 0xfc, 0xb9, 0x0e,
	0x77, 0xff, 0x55, 0x65, 0x0c, 0x3f, 0xf4, 0x8b, 0x19, 0x68, 0x4c, 0x90, 0xca, 0x29, 0xd0, 0x79,
	0x55, 0x3d, 0x32, 0x6c, 0x02, 0xfb, 0x30, 0xca, 0xd4, 0x3b, 0xc0, 0x03, 0xab, 0x7a, 0x0d, 0x0b,
	0x0c, 0xd5, 0x3b, 0x9a, 0x89, 0x89, 0x36, 0x2a, 0x9e, 0xba, 0x13, 0x2b, 0x4d, 0x7b, 0x22, 0xaf,
	0xa0, 0x78, 0x93, 0xe8, 0x20, 0x73, 0x07, 0x36, 0xb7, 0xaf, 0xac, 0xf8, 0xea, 0xb2, 0x15, 0x5f,
	0xbb, 0xbe, 0xe2, 0x17, 0xa7, 0x6a, 0xfd, 0xea, 0x54, 0xbd, 0xbe, 0x54, 0x1a, 0x4b, 0x97, 0x4a,
	0xf3, 0xca, 0x52, 0xb9, 0x76, 0x73, 0x60, 0xef, 0xdd, 0x1c, 0x84, 0x1d, 0x3b, 0x85, 0xad, 0x32,
	0xb7, 0x9a, 0x4b, 0xd3, 0xc6, 0x54, 0xd9, 0x68, 0xa2, 0x01, 0x70, 0x23, 0x37, 0xbc, 0x9a, 0xca,
	0xfe, 0xa0, 0x01, 0x16, 0x46, 0x64, 0x67, 0xc9, 0x88, 0xec, 0xde, 0x39, 0x22, 0x37, 0x6f, 0x18,
	0x91, 0xfb, 0xac, 0xbc, 0x62, 0x95, 0xaa, 0x7b, 0xa8, 0xea, 0x38, 0xf4, 0x72, 0x92, 0x5e, 0x2e,
	0x9a, 0xfb, 0xf4, 0xe9, 0xe6, 0x8b, 0xa6, 0xdc, 0x25, 0x7c, 0x61, 0x97, 0x5c, 0x5b, 0x92, 0x9f,
	0xbc, 0xbf, 0x24, 0x1f, 0xb3, 0x46, 0x66, 0xa4, 0x36, 0x76, 0x8e, 0x3d, 0x70, 0x5f, 0xdc, 0xda,
	0x03, 0xc3, 0x1f, 0xb2, 0x1a, 0xc4, 0x38, 0xe0, 0x1e, 0xba, 0x02, 0x8d, 0xed, 0x70, 0x2b, 0x2f,
	0x07, 0x5b, 0x0b, 0x97, 0x83, 0x7b, 0xac, 0xa2, 0x82, 0x4c, 0x3c, 0xea, 0x55, 0xfa, 0x15, 0xcf,
	0x3e, 0xde, 0xd4, 0xde, 0xe2, 0xc6, 0xf6, 0xfe, 0xc7, 0x3a, 0xeb, 0x60, 0xcd, 0x7a, 0x90, 0xa5,
	0x49, 0x9c, 0x01, 0xff, 0x99, 0xcd, 0x6b, 0x94, 0x29, 0xb0, 0x6e, 0x5b, 0x87, 0x4f, 0x96, 0xce,
	0x34, 0xcf, 0x89, 0xf9, 0xd7, 0x54, 0xed, 0x1a, 0x6b, 0x7a, 0x89, 0xd7, 0xb9, 0x15, 0x51, 0x33,
	0x68, 0xeb, 0xa4, 0x0c, 0x44, 0x99, 0xa8, 0xe0, 0x00, 0xbb, 0x23, 0x15, 0x69, 0xad, 0x13, 0x68,
	0x9d, 0x68, 0xec, 0x83, 0x25, 0x4e, 0xdf, 0x58, 0x91, 0x47, 0x5a, 0xfe, 0x53, 0xb6, 0xa1, 0xe2,
	0x49, 0x82, 0xfd, 0xd1, 0x3a, 0xfc, 0xf4, 0x36, 0x9f, 0x93, 0x78, 0x92, 0x78, 0xa8, 0x3c, 0xfc,
	0x77, 0x95, 0x75, 0x4f, 0x0b, 0xcc, 0x3c, 0x04, 0x7d, 0xa1, 0x7c, 0xe0, 0x2f, 0x59, 0xfd, 0x28,
	0x89, 0x27, 0x4a, 0x47, 0x7c, 0xff, 0xb6, 0x08, 0xbf, 0xc7, 0x0b, 0xbe, 0x07, 0xaf, 0x73, 0xc8,
	0xcc, 0xf6, 0x97, 0x77, 0xc9, 0xe8, 0xd0, 0x77, 0x7f, 0xc0, 0xff, 0xc2, 0x6a, 0xc3, 0x7c, 0x1c,
	0x29, 0xf3, 0xa1, 0xa1, 0xf7, 0x97, 0x1f, 0xd6, 0x65, 0xe4, 0x73, 0x56, 0x3b, 0x86, 0x10, 0x0c,
	0xf0, 0xe5, 0xe7, 0xfb, 0x7f, 0x45, 0x3c, 0xc2, 0xeb, 0xcb, 0xca, 0x22, 0xfe, 0x99, 0xd5, 0x3d,
	0xf0, 0x41, 0xa5, 0x66, 0x65, 0x21, 0x4f, 0x59, 0xe5, 0x8f, 0x60, 0x56, 0xf9, 0x86, 0xc7, 0xb4,
	0xa1, 0x56, 0x16, 0xf2, 0x39, 0xab, 0x0d, 0x41, 0x6a, 0x7f, 0xc6, 0x77, 0x97, 0xba, 0xe0, 0x3a,
	0xf9, 0xe0, 0xb0, 0x87, 0xff, 0x69, 0xb2, 0xf6, 0x95, 0xb2, 0xfd, 0x2b, 0xab, 0x1d, 0xe1, 0xdd,
	0x69, 0xf5, 0x55, 0x7b, 0xc6, 0xaa, 0x47, 0x33, 0xf0, 0x5f, 0xad, 0xf0, 0x4c, 0x5a, 0xa7, 0xf6,
	0x7f, 0x8c, 0xe2, 0x5c, 0xdb, 0x57, 0x5f, 0x55, 0xd8, 0x17, 0xac, 0x43, 0x61, 0xdd, 0x1d, 0x6e,
	0x95, 0xad, 0xb0, 0xe2, 0xe6, 0xfa, 0xbe, 0xd6, 0x2d, 0x1f, 0xb2, 0x8d, 0x3f, 0xa9, 0xcc, 0xac,
	0x36, 0x28, 0x0e, 0x96, 0x48, 0xea, 0x57, 0xab, 0xfb, 0xe4, 0xdf, 0xb1, 0x8d, 0xe1, 0x4c, 0xa5,
	0xfc, 0xce, 0xeb, 0xf6, 0x76, 0xff, 0xce, 0x0b, 0xf9, 0x65, 0xd4, 0x97, 0x8c, 0x79, 0x90, 0x4a,
	0xa5, 0x3f, 0x42, 0xec, 0xbf, 0xb1, 0xf6, 0x37, 0x6f, 0xfd, 0x99, 0x8c, 0xa7, 0xf0, 0x11, 0xa2,
	0x03, 0xdb, 0x3c, 0xbb, 0xfc, 0x41, 0x0a, 0x3f, 0xe1, 0xf2, 0xff, 0x83, 0x48, 0xb8, 0xfd, 0xd5,
	0x07, 0x88, 0x2e, 0xd3, 0x8c, 0x6b, 0xf8, 0x4b, 0xd7, 0xd7, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0xfd, 0x91, 0x0e, 0xba, 0x76, 0x13, 0x00, 0x00,
}
