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
	ShipmentMethod       string         `protobuf:"bytes,57,opt,name=shipment_method,json=shipmentMethod,proto3" json:"shipment_method,omitempty"`
	CashExchangeMoney    float32        `protobuf:"fixed32,58,opt,name=cash_exchange_money,json=cashExchangeMoney,proto3" json:"cash_exchange_money,omitempty"`
	CashExchangePoints   int32          `protobuf:"varint,59,opt,name=cash_exchange_points,json=cashExchangePoints,proto3" json:"cash_exchange_points,omitempty"`
	IsDistribution       bool           `protobuf:"varint,60,opt,name=is_distribution,json=isDistribution,proto3" json:"is_distribution,omitempty"`
	IsConfirmed          bool           `protobuf:"varint,61,opt,name=is_confirmed,json=isConfirmed,proto3" json:"is_confirmed,omitempty"`
	Food                 *OrderFood     `protobuf:"bytes,62,opt,name=food,proto3" json:"food,omitempty"`
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

func (m *Order) GetShipmentMethod() string {
	if m != nil {
		return m.ShipmentMethod
	}
	return ""
}

func (m *Order) GetCashExchangeMoney() float32 {
	if m != nil {
		return m.CashExchangeMoney
	}
	return 0
}

func (m *Order) GetCashExchangePoints() int32 {
	if m != nil {
		return m.CashExchangePoints
	}
	return 0
}

func (m *Order) GetIsDistribution() bool {
	if m != nil {
		return m.IsDistribution
	}
	return false
}

func (m *Order) GetIsConfirmed() bool {
	if m != nil {
		return m.IsConfirmed
	}
	return false
}

func (m *Order) GetFood() *OrderFood {
	if m != nil {
		return m.Food
	}
	return nil
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
	Code                 string   `protobuf:"bytes,25,opt,name=code,proto3" json:"code,omitempty"`
	FetchStatus          string   `protobuf:"bytes,26,opt,name=fetch_status,json=fetchStatus,proto3" json:"fetch_status,omitempty"`
	IsDistribution       bool     `protobuf:"varint,27,opt,name=is_distribution,json=isDistribution,proto3" json:"is_distribution,omitempty"`
	IsConfirmed          bool     `protobuf:"varint,28,opt,name=is_confirmed,json=isConfirmed,proto3" json:"is_confirmed,omitempty"`
	Id                   int64    `protobuf:"varint,29,opt,name=id,proto3" json:"id,omitempty"`
	TableId              int64    `protobuf:"varint,62,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
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

func (m *OrderWhere) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *OrderWhere) GetFetchStatus() string {
	if m != nil {
		return m.FetchStatus
	}
	return ""
}

func (m *OrderWhere) GetIsDistribution() bool {
	if m != nil {
		return m.IsDistribution
	}
	return false
}

func (m *OrderWhere) GetIsConfirmed() bool {
	if m != nil {
		return m.IsConfirmed
	}
	return false
}

func (m *OrderWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderWhere) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
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
	// 1765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x58, 0xdb, 0x72, 0x1b, 0xc7,
	0x11, 0x0d, 0x2f, 0xb8, 0x35, 0x48, 0xd0, 0x1a, 0xdd, 0x46, 0xa4, 0x64, 0x41, 0x94, 0x65, 0x21,
	0x71, 0xc2, 0x72, 0xd1, 0xe5, 0x5c, 0x6d, 0x55, 0x31, 0xa0, 0x95, 0x62, 0x55, 0x68, 0x31, 0x0b,
	0x3b, 0x8a, 0x92, 0x07, 0x64, 0xb0, 0xd3, 0x00, 0x26, 0xc2, 0xee, 0xac, 0x67, 0x67, 0x69, 0xad,
	0x7f, 0x22, 0x8f, 0xf9, 0x85, 0x7c, 0x41, 0x7e, 0x27, 0x7f, 0x92, 0x4a, 0x4d, 0xcf, 0x2c, 0x04,
	0x4a, 0x14, 0xa4, 0x54, 0x41, 0x4f, 0x79, 0xdb, 0x3e, 0x7d, 0xa6, 0x7b, 0xd0, 0xdb, 0xb7, 0x05,
	0xb4, 0xb5, 0x91, 0x68, 0x0e, 0x32, 0xa3, 0xad, 0x66, 0x37, 0x26, 0xa8, 0xbe, 0x53, 0xe9, 0x41,
	0x6e, 0xce, 0x0f, 0xb4, 0x91, 0x07, 0x99, 0x51, 0xe7, 0xc2, 0xe2, 0xee, 0x56, 0xac, 0x93, 0x44,
	0xa7, 0x9e, 0xb5, 0xbb, 0x35, 0x2a, 0x4a, 0x95, 0x4e, 0x82, 0x74, 0x85, 0x0c, 0x1c, 0xa3, 0x15,
	0x6a, 0x16, 0x20, 0x46, 0xd0, 0x91, 0x94, 0x06, 0xf3, 0x3c, 0x60, 0x9d, 0x7c, 0xaa, 0xb2, 0x04,
	0x53, 0x7b, 0xe1, 0x58, 0x5f, 0x17, 0xd9, 0xdc, 0xee, 0x0e, 0x41, 0x8f, 0xb5, 0x96, 0x1e, 0xd8,
	0xff, 0xcf, 0x36, 0xd4, 0x9e, 0x38, 0x8c, 0x75, 0x60, 0x5d, 0x49, 0xbe, 0xd6, 0x5d, 0xeb, 0x6d,
	0x44, 0xeb, 0x4a, 0xb2, 0x5b, 0xd0, 0x24, 0xf2, 0x30, 0x4f, 0xf9, 0x7a, 0x77, 0xad, 0xd7, 0x8a,
	0x1a, 0x24, 0x0f, 0x52, 0x76, 0x07, 0xc0, 0xab, 0x6c, 0x99, 0x21, 0xdf, 0x20, 0x65, 0x8b, 0x90,
	0x6f, 0xca, 0x0c, 0xd9, 0x5d, 0x68, 0xc7, 0x45, 0x6e, 0x75, 0x82, 0x66, 0xa8, 0x24, 0xdf, 0x24,
	0x93, 0x50, 0x41, 0x27, 0x92, 0xed, 0x41, 0xcb, 0x6a, 0x2b, 0x66, 0xc3, 0xb4, 0x48, 0x78, 0xbd,
	0xbb, 0xd6, 0xab, 0x45, 0x4d, 0x02, 0xbe, 0x2e, 0x12, 0x76, 0x0f, 0xb6, 0xbc, 0xf2, 0x7b, 0x54,
	0x93, 0xa9, 0xe5, 0x8d, 0xee, 0x5a, 0x6f, 0x3d, 0x6a, 0x13, 0xf6, 0x94, 0x20, 0xb6, 0x0b, 0xcd,
	0xb8, 0x30, 0x06, 0xd3, 0xb8, 0xe4, 0x4d, 0xf2, 0x3e, 0x97, 0xd9, 0x0d, 0xa8, 0x8b, 0x44, 0x17,
	0xa9, 0xe5, 0x2d, 0x3a, 0x18, 0x24, 0x76, 0x1f, 0xb6, 0x47, 0x38, 0xd6, 0x06, 0x87, 0x41, 0x0d,
	0xa4, 0xde, 0xf2, 0xe0, 0xd1, 0x9c, 0x24, 0x62, 0x5b, 0x88, 0x59, 0x45, 0x6a, 0x7b, 0x92, 0x07,
	0x03, 0xe9, 0x2e, 0xb4, 0xbd, 0x76, 0x98, 0x09, 0x25, 0xf9, 0x16, 0x51, 0xc0, 0x43, 0x67, 0x42,
	0x49, 0xf6, 0x10, 0x76, 0x02, 0xc1, 0xe0, 0xb8, 0x48, 0x25, 0x4a, 0xbe, 0x4d, 0xa4, 0x8e, 0x87,
	0xa3, 0x80, 0xb2, 0x07, 0xd0, 0xc1, 0x17, 0x99, 0x7b, 0x83, 0x95, 0xbf, 0x0e, 0xf1, 0xb6, 0x03,
	0x1a, 0x1c, 0x3e, 0x84, 0x1d, 0xa9, 0xf2, 0x98, 0x2c, 0x06, 0xde, 0x8e, 0xb7, 0x57, 0xc1, 0x81,
	0xc8, 0x60, 0x93, 0xae, 0xf4, 0x41, 0x77, 0xad, 0xd7, 0x8c, 0xe8, 0xd9, 0xc5, 0x6a, 0x7e, 0x8b,
	0x2b, 0x84, 0xcf, 0x65, 0x17, 0xea, 0x58, 0xa4, 0x43, 0x89, 0x33, 0x75, 0x8e, 0xa6, 0xe4, 0x8c,
	0xf4, 0xed, 0x58, 0xa4, 0xc7, 0x01, 0x62, 0x37, 0xa1, 0xe1, 0x28, 0xb1, 0x96, 0xfc, 0x2a, 0x69,
	0xeb, 0xb1, 0x48, 0xfb, 0x9a, 0xec, 0xaa, 0xf4, 0x5c, 0xab, 0x18, 0x25, 0xbf, 0xe6, 0xed, 0x56,
	0xb2, 0xd3, 0x25, 0x5a, 0xaa, 0xb1, 0x42, 0xc9, 0xaf, 0x7b, 0x5d, 0x25, 0xb3, 0x2e, 0xb4, 0x73,
	0x31, 0xc6, 0x49, 0x21, 0x8c, 0xbb, 0xd2, 0x0d, 0xef, 0x72, 0x01, 0x62, 0x1f, 0x02, 0x18, 0xb4,
	0x85, 0x49, 0xc5, 0x68, 0x86, 0xfc, 0x26, 0x11, 0x16, 0x10, 0xc6, 0xa1, 0x91, 0x17, 0xa3, 0xbf,
	0x61, 0x6c, 0x39, 0xf7, 0x79, 0x19, 0x44, 0xf7, 0xfb, 0x47, 0x5a, 0x96, 0xfc, 0x16, 0xc1, 0xf4,
	0xec, 0xd2, 0x38, 0x13, 0xa5, 0xcf, 0xd4, 0x5d, 0x4f, 0xcf, 0x44, 0x49, 0x79, 0xba, 0x07, 0xad,
	0x91, 0x11, 0x69, 0x3c, 0x75, 0x59, 0xba, 0x47, 0x59, 0xda, 0xf4, 0x80, 0xcf, 0xd1, 0x78, 0xa6,
	0x30, 0xb5, 0x43, 0x95, 0xf1, 0xdb, 0x21, 0xc9, 0x08, 0x38, 0xc9, 0x9c, 0xa3, 0xf1, 0x4c, 0x4c,
	0xf8, 0x1d, 0xca, 0x5d, 0x7a, 0x76, 0x58, 0x82, 0x89, 0xe6, 0x1f, 0x7a, 0xe7, 0xee, 0xd9, 0x15,
	0xca, 0xa8, 0x28, 0xd1, 0x0c, 0x49, 0x73, 0xd7, 0x17, 0x0a, 0x21, 0xa7, 0x4e, 0xed, 0x7c, 0x88,
	0x34, 0xc6, 0x99, 0xbb, 0x40, 0xd7, 0xd7, 0x81, 0x07, 0x4e, 0xa4, 0xcb, 0xc5, 0xa0, 0x34, 0x28,
	0x72, 0x9d, 0xf2, 0x7b, 0x74, 0x7c, 0xcb, 0x83, 0x11, 0x61, 0x2e, 0xdb, 0x73, 0x2b, 0x6c, 0x91,
	0xf3, 0x7d, 0xd2, 0x06, 0xc9, 0x39, 0x76, 0xbf, 0x3a, 0xe8, 0xee, 0x7b, 0xc7, 0x99, 0x28, 0x07,
	0x5e, 0x7d, 0x1f, 0xb6, 0x7d, 0x12, 0x54, 0x8c, 0x8f, 0xbc, 0x6d, 0x0f, 0x06, 0xd2, 0x1e, 0xb4,
	0xf0, 0x45, 0xa6, 0x5c, 0xc5, 0x58, 0xfe, 0xb1, 0x8f, 0x80, 0x07, 0x8e, 0xac, 0xcb, 0x0b, 0x97,
	0x5e, 0x4e, 0xf5, 0xd0, 0x7b, 0x76, 0xe2, 0x91, 0x75, 0x9e, 0x5d, 0x1b, 0xca, 0x90, 0x74, 0x3d,
	0xef, 0x39, 0x20, 0x47, 0xd6, 0x19, 0xcd, 0xd5, 0x24, 0xf5, 0xda, 0x1f, 0x7b, 0xa3, 0x1e, 0x38,
	0xa2, 0xca, 0xaa, 0x72, 0xd3, 0xa9, 0x7f, 0x42, 0x6a, 0xa8, 0x20, 0x4f, 0x18, 0xab, 0x54, 0xe5,
	0x53, 0x4f, 0xf8, 0xc4, 0x13, 0x2a, 0xc8, 0x7b, 0x8f, 0x0d, 0x0a, 0xeb, 0xf5, 0x3f, 0xf5, 0xde,
	0x03, 0xe2, 0xd5, 0x45, 0x26, 0x2b, 0xf5, 0xcf, 0xbc, 0x3a, 0x20, 0x47, 0x54, 0x68, 0xd9, 0x4c,
	0xd8, 0xb1, 0x36, 0xc9, 0x30, 0xd7, 0x85, 0x89, 0x91, 0x1f, 0x10, 0xa7, 0x53, 0xc1, 0x03, 0x42,
	0xd9, 0x23, 0x68, 0x08, 0xdf, 0x7a, 0xf9, 0xcf, 0xbb, 0x6b, 0xbd, 0xf6, 0xe1, 0x47, 0x07, 0x97,
	0xb7, 0xf5, 0x83, 0x27, 0x0b, 0x6d, 0x3a, 0xaa, 0x0e, 0xb1, 0x2f, 0xa1, 0x21, 0xa9, 0x9b, 0xe7,
	0xfc, 0x17, 0xdd, 0x8d, 0x5e, 0xfb, 0xf0, 0xfe, 0xd2, 0xf3, 0xbe, 0xf3, 0x47, 0xd5, 0x19, 0xf6,
	0x85, 0xeb, 0x7f, 0xbe, 0x9b, 0xf2, 0x5f, 0x92, 0xff, 0xee, 0x9b, 0xce, 0xf7, 0x03, 0x2f, 0x9a,
	0x9f, 0x70, 0xbf, 0xb2, 0x1a, 0x14, 0xc3, 0x04, 0xed, 0x54, 0x4b, 0xfe, 0x2b, 0xff, 0x2b, 0x2b,
	0xf8, 0x94, 0x50, 0x76, 0x00, 0x57, 0x63, 0x91, 0x4f, 0x87, 0xf8, 0x22, 0x9e, 0x8a, 0x74, 0x82,
	0xc3, 0x44, 0xa7, 0x58, 0xf2, 0x5f, 0x53, 0xef, 0xb9, 0xe2, 0x54, 0x5f, 0x05, 0xcd, 0xa9, 0x53,
	0xb0, 0x4f, 0xe1, 0xda, 0x45, 0x7e, 0xa6, 0x55, 0x6a, 0x73, 0xfe, 0x1b, 0xca, 0x6c, 0xb6, 0x78,
	0xe0, 0x8c, 0x34, 0xee, 0x2a, 0x2a, 0x1f, 0x4a, 0x95, 0x5b, 0xa3, 0x46, 0x85, 0x55, 0x3a, 0xe5,
	0x5f, 0x50, 0xbd, 0x77, 0x54, 0x7e, 0xbc, 0x80, 0xba, 0x4e, 0xa5, 0xf2, 0x61, 0xac, 0xd3, 0xb1,
	0x32, 0x09, 0x4a, 0xfe, 0xa5, 0x6f, 0x1b, 0x2a, 0xef, 0x57, 0x10, 0xfb, 0x1c, 0x36, 0xc7, 0x5a,
	0x4b, 0xfe, 0x88, 0x02, 0x72, 0x6f, 0x69, 0x40, 0xdd, 0x00, 0x8c, 0x88, 0xbe, 0xff, 0xcf, 0x75,
	0x68, 0x56, 0x41, 0x7a, 0x6d, 0x06, 0x2e, 0x4e, 0xb2, 0xf9, 0x18, 0x9c, 0x4f, 0xb2, 0x41, 0xea,
	0x8a, 0x3e, 0x15, 0x49, 0x35, 0x03, 0xe9, 0xd9, 0xdd, 0x55, 0xaa, 0x3c, 0x9b, 0x89, 0x72, 0x48,
	0xba, 0x4d, 0xd2, 0xb5, 0x03, 0xf6, 0xb5, 0xa3, 0x50, 0x53, 0x16, 0x33, 0x52, 0xd7, 0x7c, 0x11,
	0x54, 0xb2, 0x2b, 0x69, 0x25, 0x63, 0x61, 0x24, 0x4d, 0xc6, 0x56, 0x14, 0x24, 0x87, 0x4f, 0x30,
	0x95, 0x68, 0x68, 0x22, 0xd6, 0xa2, 0x20, 0x39, 0x3c, 0xd1, 0x23, 0x35, 0xc3, 0x30, 0x0a, 0x83,
	0xc4, 0xae, 0x41, 0x0d, 0x13, 0xa1, 0x66, 0x34, 0x07, 0x5b, 0x91, 0x17, 0x5c, 0xfd, 0x89, 0x73,
	0x61, 0x05, 0x4d, 0x66, 0xf0, 0x3d, 0xcf, 0x03, 0x27, 0xd2, 0x95, 0x47, 0x50, 0x16, 0x66, 0x46,
	0xb3, 0xaf, 0x15, 0x05, 0xfa, 0xb7, 0x66, 0xb6, 0xff, 0xef, 0x3a, 0x00, 0x85, 0xef, 0xe9, 0x14,
	0x0d, 0x39, 0xc8, 0xc4, 0x04, 0x7d, 0xbc, 0x6a, 0x91, 0x17, 0x9c, 0x03, 0xf7, 0x30, 0xcc, 0xd5,
	0x0f, 0x48, 0x01, 0xab, 0x45, 0x4d, 0x07, 0x0c, 0xd4, 0x0f, 0xbe, 0x75, 0x6b, 0x63, 0x55, 0x3a,
	0x09, 0x11, 0xab, 0x44, 0x17, 0x91, 0xe7, 0x58, 0x7e, 0xaf, 0x8d, 0xcc, 0x43, 0xc0, 0xe6, 0xf2,
	0x85, 0x4d, 0xa4, 0xb6, 0x6c, 0x13, 0xa9, 0xbf, 0xba, 0x89, 0x2c, 0x36, 0xff, 0xc6, 0xc5, 0xe6,
	0xff, 0xea, 0xec, 0x6b, 0x2e, 0x9d, 0x7d, 0xad, 0x0b, 0xb3, 0xef, 0x95, 0x05, 0x07, 0x5e, 0x5b,
	0x70, 0xb8, 0xeb, 0x8e, 0xa5, 0x2b, 0xa5, 0xb0, 0x41, 0x54, 0xa2, 0xb3, 0xa9, 0xf2, 0xe1, 0xd8,
	0x20, 0xd2, 0xe2, 0xd0, 0x8c, 0xea, 0x2a, 0x7f, 0x6c, 0x10, 0x17, 0x3a, 0xf9, 0xf6, 0x92, 0x4e,
	0xde, 0x79, 0x6b, 0x27, 0xdf, 0xb9, 0xa4, 0x93, 0x3f, 0x80, 0x4e, 0xf8, 0x9d, 0x15, 0xeb, 0x03,
	0x62, 0x6d, 0x07, 0xf4, 0x65, 0xc3, 0x7f, 0x39, 0x0f, 0xaf, 0xf8, 0x57, 0x37, 0x9f, 0x87, 0xd5,
	0xc8, 0x63, 0x0b, 0x23, 0xef, 0x95, 0x59, 0x7e, 0xf5, 0xf5, 0x59, 0x7e, 0x0b, 0x9a, 0xb9, 0x15,
	0xc6, 0xba, 0x76, 0x7b, 0x2d, 0xbc, 0x71, 0x27, 0x1f, 0x59, 0x76, 0x1d, 0xea, 0x98, 0x52, 0x1f,
	0xbe, 0x1e, 0x12, 0x34, 0x75, 0x3d, 0xb8, 0xda, 0x61, 0x6e, 0x2c, 0xec, 0x30, 0x1f, 0xc0, 0x86,
	0x92, 0x39, 0xbf, 0xd9, 0xdd, 0xe8, 0x6d, 0x44, 0xee, 0xf1, 0xb2, 0x1e, 0xc6, 0x2f, 0xed, 0x61,
	0x0c, 0x36, 0x63, 0x2d, 0xb1, 0x5a, 0x09, 0xdc, 0xb3, 0x7b, 0xf5, 0x63, 0xb4, 0xf1, 0xb4, 0x0a,
	0x86, 0x5f, 0x0b, 0xda, 0x84, 0x85, 0x50, 0x5c, 0xd2, 0x98, 0xf6, 0xde, 0xa9, 0x31, 0xdd, 0x7e,
	0xbd, 0x31, 0xf9, 0xa6, 0x72, 0x67, 0x71, 0xb1, 0xb6, 0x6e, 0x91, 0x71, 0x51, 0x7e, 0x44, 0x68,
	0x83, 0xe4, 0x13, 0xb9, 0xff, 0x8f, 0x75, 0xd8, 0xa6, 0x0a, 0x8b, 0x30, 0xcf, 0x74, 0x9a, 0x23,
	0xfb, 0xdc, 0x45, 0xc9, 0x2a, 0x5b, 0x52, 0x95, 0xb5, 0x0f, 0xef, 0x2c, 0xed, 0x6b, 0x51, 0x20,
	0xb3, 0xcf, 0x7c, 0x6d, 0x1a, 0xaa, 0xc0, 0x25, 0xa7, 0xce, 0x1c, 0xc9, 0x97, 0xae, 0x71, 0x87,
	0x94, 0xc5, 0x24, 0xe7, 0x1b, 0x34, 0x93, 0xde, 0xe2, 0xca, 0x73, 0xdd, 0x21, 0x34, 0x46, 0x1b,
	0xaa, 0xda, 0x25, 0x87, 0xbe, 0x72, 0xa4, 0xc8, 0x73, 0xd9, 0xa7, 0xb0, 0xa9, 0xd2, 0xb1, 0xa6,
	0x6a, 0x6e, 0x1f, 0xde, 0x7e, 0xd3, 0x99, 0x93, 0x74, 0xac, 0x23, 0x62, 0x1e, 0xfe, 0xab, 0x06,
	0x9d, 0xd3, 0x92, 0x3c, 0x0f, 0xd0, 0x9c, 0xab, 0x18, 0xd9, 0x9f, 0xa1, 0x11, 0x82, 0xcc, 0x1e,
	0xbc, 0xc9, 0xc2, 0x6f, 0xe9, 0x33, 0x2a, 0xc2, 0xef, 0x0a, 0xcc, 0xed, 0xee, 0xc7, 0x6f, 0xa3,
	0xf9, 0xa0, 0xef, 0xff, 0x88, 0xfd, 0x09, 0xea, 0x83, 0x62, 0x94, 0x28, 0xfb, 0xae, 0xa6, 0x1f,
	0x2c, 0x0f, 0xd6, 0x4b, 0xcb, 0x67, 0x50, 0x3f, 0xc6, 0x19, 0x5a, 0x64, 0xcb, 0xe3, 0xfb, 0x3f,
	0x59, 0xec, 0xd3, 0x4e, 0xb8, 0x32, 0x8b, 0x7f, 0x80, 0x46, 0x84, 0x31, 0xaa, 0xcc, 0xae, 0xcc,
	0xe4, 0x29, 0x6c, 0xfc, 0x0e, 0xed, 0x2a, 0x6f, 0x78, 0xec, 0xe7, 0xe9, 0xca, 0x4c, 0x7e, 0x0b,
	0xf5, 0x01, 0x0a, 0x13, 0x4f, 0xd9, 0xfe, 0xd2, 0x23, 0x34, 0xfc, 0xde, 0xd9, 0xec, 0xe1, 0xdf,
	0x3b, 0xb0, 0x75, 0x21, 0x6d, 0x9f, 0x41, 0xbd, 0x4f, 0x0b, 0xe9, 0xea, 0xb3, 0xf6, 0x09, 0xd4,
	0xfa, 0x53, 0x8c, 0x9f, 0xaf, 0x30, 0x26, 0xed, 0x53, 0xf7, 0xe1, 0x56, 0x9e, 0x19, 0x77, 0xf5,
	0x55, 0x99, 0x7d, 0x0a, 0xdb, 0xde, 0x6c, 0x58, 0x8c, 0x57, 0x59, 0x0a, 0x2b, 0x2e, 0xae, 0xff,
	0xd7, 0xbc, 0x65, 0x03, 0xd8, 0xfc, 0xbd, 0xca, 0xed, 0x6a, 0x8d, 0x52, 0x63, 0x49, 0x84, 0x79,
	0xbe, 0xba, 0x57, 0xfe, 0x17, 0xd8, 0x1c, 0x4c, 0x55, 0xc6, 0x1e, 0xbe, 0xe9, 0xc0, 0x20, 0x2c,
	0x04, 0x55, 0x39, 0xf5, 0xde, 0x4e, 0x9c, 0x1b, 0x8f, 0x61, 0xab, 0xda, 0x06, 0xdf, 0x9f, 0x93,
	0x67, 0x00, 0x7d, 0x2d, 0xf1, 0x8f, 0x68, 0xd4, 0xb8, 0x5c, 0x6d, 0xbc, 0xff, 0x0a, 0xad, 0xc7,
	0xb4, 0xd5, 0xbc, 0xb7, 0xcb, 0x0b, 0x80, 0x08, 0x33, 0xa1, 0xcc, 0x7b, 0x7d, 0x09, 0xd5, 0x67,
	0xe3, 0xfb, 0x73, 0x82, 0xb0, 0xf3, 0xe4, 0xe5, 0xbf, 0xa5, 0x94, 0xf9, 0xcb, 0xbf, 0xc9, 0x3d,
	0x71, 0xf7, 0x93, 0x77, 0x20, 0x2d, 0xb8, 0xf9, 0x06, 0x6a, 0x7d, 0xfa, 0xab, 0x6e, 0xd5, 0x65,
	0x55, 0x6d, 0x42, 0xab, 0x2a, 0xab, 0x67, 0x00, 0x03, 0xb4, 0x76, 0x86, 0xf4, 0xd5, 0xb2, 0xca,
	0xdb, 0x8e, 0xea, 0xf4, 0xcf, 0xf3, 0x67, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xdf, 0x22,
	0x9d, 0x17, 0x17, 0x00, 0x00,
}
