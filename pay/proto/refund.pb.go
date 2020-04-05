// Code generated by protoc-gen-go. DO NOT EDIT.
// source: refund.proto

package geiqin_srv_pay

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

//退款
type Refund struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RefundSn             string   `protobuf:"bytes,2,opt,name=refund_sn,json=refundSn,proto3" json:"refund_sn,omitempty"`
	TotalFee             float32  `protobuf:"fixed32,3,opt,name=total_fee,json=totalFee,proto3" json:"total_fee,omitempty"`
	RefundFee            float32  `protobuf:"fixed32,4,opt,name=refund_fee,json=refundFee,proto3" json:"refund_fee,omitempty"`
	Currency             string   `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	OutTradeId           int64    `protobuf:"varint,7,opt,name=out_trade_id,json=outTradeId,proto3" json:"out_trade_id,omitempty"`
	ChargeId             int64    `protobuf:"varint,8,opt,name=charge_id,json=chargeId,proto3" json:"charge_id,omitempty"`
	ChargeSn             string   `protobuf:"bytes,9,opt,name=charge_sn,json=chargeSn,proto3" json:"charge_sn,omitempty"`
	ChargeTransactionNo  string   `protobuf:"bytes,10,opt,name=charge_transaction_no,json=chargeTransactionNo,proto3" json:"charge_transaction_no,omitempty"`
	TargetUserType       string   `protobuf:"bytes,11,opt,name=target_user_type,json=targetUserType,proto3" json:"target_user_type,omitempty"`
	TargetUserId         int64    `protobuf:"varint,12,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	RefundDesc           string   `protobuf:"bytes,13,opt,name=refund_desc,json=refundDesc,proto3" json:"refund_desc,omitempty"`
	RefundFunder         string   `protobuf:"bytes,14,opt,name=refund_funder,json=refundFunder,proto3" json:"refund_funder,omitempty"`
	Metadata             string   `protobuf:"bytes,15,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ReturnExtra          string   `protobuf:"bytes,16,opt,name=return_extra,json=returnExtra,proto3" json:"return_extra,omitempty"`
	TransactionNo        string   `protobuf:"bytes,17,opt,name=transaction_no,json=transactionNo,proto3" json:"transaction_no,omitempty"`
	Status               string   `protobuf:"bytes,18,opt,name=status,proto3" json:"status,omitempty"`
	HandledAt            string   `protobuf:"bytes,19,opt,name=handled_at,json=handledAt,proto3" json:"handled_at,omitempty"`
	CreatedAt            string   `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Refund) Reset()         { *m = Refund{} }
func (m *Refund) String() string { return proto.CompactTextString(m) }
func (*Refund) ProtoMessage()    {}
func (*Refund) Descriptor() ([]byte, []int) {
	return fileDescriptor_27efa219a51643e4, []int{0}
}

func (m *Refund) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Refund.Unmarshal(m, b)
}
func (m *Refund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Refund.Marshal(b, m, deterministic)
}
func (m *Refund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Refund.Merge(m, src)
}
func (m *Refund) XXX_Size() int {
	return xxx_messageInfo_Refund.Size(m)
}
func (m *Refund) XXX_DiscardUnknown() {
	xxx_messageInfo_Refund.DiscardUnknown(m)
}

var xxx_messageInfo_Refund proto.InternalMessageInfo

func (m *Refund) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Refund) GetRefundSn() string {
	if m != nil {
		return m.RefundSn
	}
	return ""
}

func (m *Refund) GetTotalFee() float32 {
	if m != nil {
		return m.TotalFee
	}
	return 0
}

func (m *Refund) GetRefundFee() float32 {
	if m != nil {
		return m.RefundFee
	}
	return 0
}

func (m *Refund) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Refund) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Refund) GetOutTradeId() int64 {
	if m != nil {
		return m.OutTradeId
	}
	return 0
}

func (m *Refund) GetChargeId() int64 {
	if m != nil {
		return m.ChargeId
	}
	return 0
}

func (m *Refund) GetChargeSn() string {
	if m != nil {
		return m.ChargeSn
	}
	return ""
}

func (m *Refund) GetChargeTransactionNo() string {
	if m != nil {
		return m.ChargeTransactionNo
	}
	return ""
}

func (m *Refund) GetTargetUserType() string {
	if m != nil {
		return m.TargetUserType
	}
	return ""
}

func (m *Refund) GetTargetUserId() int64 {
	if m != nil {
		return m.TargetUserId
	}
	return 0
}

func (m *Refund) GetRefundDesc() string {
	if m != nil {
		return m.RefundDesc
	}
	return ""
}

func (m *Refund) GetRefundFunder() string {
	if m != nil {
		return m.RefundFunder
	}
	return ""
}

func (m *Refund) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Refund) GetReturnExtra() string {
	if m != nil {
		return m.ReturnExtra
	}
	return ""
}

func (m *Refund) GetTransactionNo() string {
	if m != nil {
		return m.TransactionNo
	}
	return ""
}

func (m *Refund) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Refund) GetHandledAt() string {
	if m != nil {
		return m.HandledAt
	}
	return ""
}

func (m *Refund) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Refund) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

//
type RefundResponse struct {
	Entity               *Refund   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Refund `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error    `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info     `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RefundResponse) Reset()         { *m = RefundResponse{} }
func (m *RefundResponse) String() string { return proto.CompactTextString(m) }
func (*RefundResponse) ProtoMessage()    {}
func (*RefundResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27efa219a51643e4, []int{1}
}

func (m *RefundResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefundResponse.Unmarshal(m, b)
}
func (m *RefundResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefundResponse.Marshal(b, m, deterministic)
}
func (m *RefundResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefundResponse.Merge(m, src)
}
func (m *RefundResponse) XXX_Size() int {
	return xxx_messageInfo_RefundResponse.Size(m)
}
func (m *RefundResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefundResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefundResponse proto.InternalMessageInfo

func (m *RefundResponse) GetEntity() *Refund {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *RefundResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *RefundResponse) GetItems() []*Refund {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *RefundResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *RefundResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Refund)(nil), "geiqin.srv.pay.Refund")
	proto.RegisterType((*RefundResponse)(nil), "geiqin.srv.pay.RefundResponse")
}

func init() {
	proto.RegisterFile("refund.proto", fileDescriptor_27efa219a51643e4)
}

var fileDescriptor_27efa219a51643e4 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdf, 0x6e, 0x13, 0x39,
	0x14, 0xc6, 0x37, 0x49, 0x33, 0x9b, 0x9c, 0x24, 0xb3, 0x5d, 0xf7, 0x8f, 0xbc, 0x5d, 0x01, 0xa1,
	0x80, 0x14, 0x09, 0x94, 0x8b, 0xf0, 0x00, 0xa8, 0xa5, 0x2d, 0xea, 0x0d, 0x42, 0x93, 0x56, 0x5c,
	0x8e, 0xcc, 0xcc, 0x49, 0x33, 0x52, 0x63, 0x07, 0xfb, 0x4c, 0x45, 0x5e, 0x87, 0x57, 0xe3, 0x0d,
	0x78, 0x02, 0xe4, 0x63, 0x47, 0xa5, 0x11, 0x85, 0x8b, 0xde, 0xc5, 0xdf, 0xef, 0xcb, 0x37, 0x67,
	0x3e, 0x1f, 0x0d, 0xf4, 0x2d, 0xce, 0x6a, 0x5d, 0x8e, 0x97, 0xd6, 0x90, 0x11, 0xe9, 0x15, 0x56,
	0x9f, 0x2b, 0x3d, 0x76, 0xf6, 0x66, 0xbc, 0x54, 0xab, 0x83, 0x7e, 0x61, 0x16, 0x0b, 0xa3, 0x03,
	0x3d, 0xfc, 0xda, 0x86, 0x24, 0x63, 0xbb, 0x48, 0xa1, 0x59, 0x95, 0xb2, 0x31, 0x6c, 0x8c, 0xda,
	0x59, 0xb3, 0x2a, 0xc5, 0xff, 0xd0, 0x0d, 0x41, 0xb9, 0xd3, 0xb2, 0x39, 0x6c, 0x8c, 0xba, 0x59,
	0x27, 0x08, 0x53, 0xed, 0x21, 0x19, 0x52, 0xd7, 0xf9, 0x0c, 0x51, 0xb6, 0x86, 0x8d, 0x51, 0x33,
	0xeb, 0xb0, 0x70, 0x86, 0x28, 0x1e, 0x01, 0xc4, 0x7f, 0x7a, 0xba, 0xc5, 0x34, 0x66, 0x79, 0x7c,
	0x00, 0x9d, 0xa2, 0xb6, 0x16, 0x75, 0xb1, 0x92, 0xed, 0x90, 0xbb, 0x3e, 0x0b, 0x01, 0x5b, 0xb4,
	0x5a, 0xa2, 0x4c, 0x58, 0xe7, 0xdf, 0x62, 0x08, 0x7d, 0x53, 0x53, 0x4e, 0x56, 0x95, 0x98, 0x57,
	0xa5, 0xfc, 0x7b, 0xd8, 0x18, 0xb5, 0x32, 0x30, 0x35, 0x5d, 0x78, 0xe9, 0x9c, 0x47, 0x2d, 0xe6,
	0xca, 0x5e, 0x31, 0xee, 0x30, 0xee, 0x04, 0xe1, 0x0e, 0x74, 0x5a, 0x76, 0xe3, 0xf3, 0x58, 0x98,
	0x6a, 0x31, 0x81, 0xbd, 0x08, 0xc9, 0x2a, 0xed, 0x54, 0x41, 0x95, 0xd1, 0xb9, 0x36, 0x12, 0xd8,
	0xb8, 0x13, 0xe0, 0xc5, 0x2d, 0x7b, 0x6f, 0xc4, 0x08, 0xb6, 0xc9, 0xab, 0x94, 0xd7, 0x0e, 0x6d,
	0xce, 0xf3, 0xf6, 0xd8, 0x9e, 0x06, 0xfd, 0xd2, 0xa1, 0xbd, 0xf0, 0x93, 0x3f, 0x87, 0xf4, 0x67,
	0x67, 0x55, 0xca, 0x3e, 0x0f, 0xd7, 0xbf, 0xf5, 0x9d, 0x97, 0xe2, 0x09, 0xf4, 0x62, 0x5d, 0x25,
	0xba, 0x42, 0x0e, 0x38, 0x2a, 0x36, 0x78, 0x82, 0xae, 0x10, 0xcf, 0x60, 0xb0, 0xee, 0xb3, 0xd6,
	0x25, 0x5a, 0x99, 0xb2, 0x25, 0xde, 0xf3, 0x19, 0x6b, 0xbe, 0xd5, 0x05, 0x92, 0x2a, 0x15, 0x29,
	0xf9, 0x4f, 0x78, 0xcb, 0xf5, 0x59, 0x3c, 0xf5, 0x3b, 0x41, 0xb5, 0xd5, 0x39, 0x7e, 0x21, 0xab,
	0xe4, 0x36, 0xf3, 0x5e, 0xd0, 0x4e, 0xbd, 0x24, 0x5e, 0x40, 0xba, 0xd1, 0xc0, 0xbf, 0x6c, 0x1a,
	0xd0, 0x9d, 0x77, 0xdf, 0x87, 0xc4, 0x91, 0xa2, 0xda, 0x49, 0xc1, 0x38, 0x9e, 0xfc, 0x95, 0xcf,
	0x95, 0x2e, 0xaf, 0xb1, 0xcc, 0x15, 0xc9, 0x1d, 0x66, 0xdd, 0xa8, 0x1c, 0x91, 0xc7, 0x85, 0x45,
	0x45, 0x01, 0xef, 0x06, 0x1c, 0x95, 0x80, 0xeb, 0x65, 0xb9, 0xc6, 0x7b, 0x01, 0x47, 0xe5, 0x88,
	0x0e, 0xbf, 0x37, 0x20, 0x0d, 0x4b, 0x9a, 0xa1, 0x5b, 0x1a, 0xed, 0x50, 0x8c, 0x21, 0x41, 0x4d,
	0x15, 0xad, 0x78, 0x61, 0x7b, 0x93, 0xfd, 0xf1, 0xdd, 0x35, 0x1f, 0x47, 0x7f, 0x74, 0x89, 0x97,
	0xd0, 0x5e, 0xaa, 0x2b, 0xb4, 0xbc, 0xc8, 0xbd, 0xc9, 0xde, 0xa6, 0xfd, 0x83, 0x87, 0x59, 0xf0,
	0x88, 0x57, 0xd0, 0xae, 0x08, 0x17, 0x4e, 0xb6, 0x86, 0xad, 0xdf, 0x64, 0x07, 0x93, 0x8f, 0x46,
	0x6b, 0x8d, 0xe5, 0x45, 0xff, 0x45, 0xf4, 0xa9, 0x87, 0x59, 0xf0, 0x88, 0x11, 0x6c, 0x55, 0x7a,
	0x66, 0x78, 0xef, 0x7b, 0x93, 0xdd, 0x4d, 0xef, 0xb9, 0x9e, 0x99, 0x8c, 0x1d, 0x93, 0x6f, 0x4d,
	0x18, 0x84, 0x07, 0x4d, 0xd1, 0xde, 0x54, 0x05, 0x8a, 0x63, 0x48, 0xde, 0x72, 0x65, 0xe2, 0x9e,
	0x89, 0x0e, 0x1e, 0xdf, 0x33, 0x69, 0x6c, 0xed, 0xf0, 0x2f, 0x9f, 0x71, 0xc9, 0xbd, 0x3e, 0x2c,
	0xe3, 0x04, 0xaf, 0xf1, 0x41, 0x19, 0x6f, 0xa0, 0xf5, 0x0e, 0xe9, 0x01, 0x01, 0xa7, 0x90, 0x4c,
	0x51, 0xd9, 0x62, 0x2e, 0xfe, 0xdb, 0xf4, 0x1e, 0x2b, 0x87, 0x1f, 0xe7, 0x68, 0xf1, 0xcf, 0x31,
	0x9f, 0x12, 0xfe, 0x0c, 0xbe, 0xfe, 0x11, 0x00, 0x00, 0xff, 0xff, 0x86, 0x3d, 0xd2, 0x22, 0x34,
	0x05, 0x00, 0x00,
}
