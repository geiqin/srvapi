// Code generated by protoc-gen-go. DO NOT EDIT.
// source: withdrawal.proto

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

//提现
type Withdrawal struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount               float32  `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	AssetTransaction     string   `protobuf:"bytes,3,opt,name=asset_transaction,json=assetTransaction,proto3" json:"asset_transaction,omitempty"`
	BalanceTransaction   string   `protobuf:"bytes,4,opt,name=balance_transaction,json=balanceTransaction,proto3" json:"balance_transaction,omitempty"`
	Channel              string   `protobuf:"bytes,5,opt,name=channel,proto3" json:"channel,omitempty"`
	Extra                string   `protobuf:"bytes,6,opt,name=extra,proto3" json:"extra,omitempty"`
	FailureMsg           string   `protobuf:"bytes,7,opt,name=failure_msg,json=failureMsg,proto3" json:"failure_msg,omitempty"`
	Fee                  float32  `protobuf:"fixed32,8,opt,name=fee,proto3" json:"fee,omitempty"`
	Livemode             bool     `protobuf:"varint,9,opt,name=livemode,proto3" json:"livemode,omitempty"`
	OperationUrl         string   `protobuf:"bytes,10,opt,name=operation_url,json=operationUrl,proto3" json:"operation_url,omitempty"`
	OrderNo              string   `protobuf:"bytes,11,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	Source               string   `protobuf:"bytes,12,opt,name=source,proto3" json:"source,omitempty"`
	CanceledAt           string   `protobuf:"bytes,13,opt,name=canceled_at,json=canceledAt,proto3" json:"canceled_at,omitempty"`
	SucceededAt          string   `protobuf:"bytes,14,opt,name=succeeded_at,json=succeededAt,proto3" json:"succeeded_at,omitempty"`
	TargetUserType       string   `protobuf:"bytes,15,opt,name=target_user_type,json=targetUserType,proto3" json:"target_user_type,omitempty"`
	TargetUserId         int64    `protobuf:"varint,16,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	ServiceFee           float32  `protobuf:"fixed32,17,opt,name=service_fee,json=serviceFee,proto3" json:"service_fee,omitempty"`
	SettleAccountId      int64    `protobuf:"varint,18,opt,name=settle_account_id,json=settleAccountId,proto3" json:"settle_account_id,omitempty"`
	StoreId              int64    `protobuf:"varint,19,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Meno                 string   `protobuf:"bytes,20,opt,name=meno,proto3" json:"meno,omitempty"`
	Metadata             string   `protobuf:"bytes,21,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Status               string   `protobuf:"bytes,22,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string   `protobuf:"bytes,23,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,24,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Withdrawal) Reset()         { *m = Withdrawal{} }
func (m *Withdrawal) String() string { return proto.CompactTextString(m) }
func (*Withdrawal) ProtoMessage()    {}
func (*Withdrawal) Descriptor() ([]byte, []int) {
	return fileDescriptor_729e0e9b3824e84b, []int{0}
}

func (m *Withdrawal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Withdrawal.Unmarshal(m, b)
}
func (m *Withdrawal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Withdrawal.Marshal(b, m, deterministic)
}
func (m *Withdrawal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Withdrawal.Merge(m, src)
}
func (m *Withdrawal) XXX_Size() int {
	return xxx_messageInfo_Withdrawal.Size(m)
}
func (m *Withdrawal) XXX_DiscardUnknown() {
	xxx_messageInfo_Withdrawal.DiscardUnknown(m)
}

var xxx_messageInfo_Withdrawal proto.InternalMessageInfo

func (m *Withdrawal) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Withdrawal) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Withdrawal) GetAssetTransaction() string {
	if m != nil {
		return m.AssetTransaction
	}
	return ""
}

func (m *Withdrawal) GetBalanceTransaction() string {
	if m != nil {
		return m.BalanceTransaction
	}
	return ""
}

func (m *Withdrawal) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Withdrawal) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *Withdrawal) GetFailureMsg() string {
	if m != nil {
		return m.FailureMsg
	}
	return ""
}

func (m *Withdrawal) GetFee() float32 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Withdrawal) GetLivemode() bool {
	if m != nil {
		return m.Livemode
	}
	return false
}

func (m *Withdrawal) GetOperationUrl() string {
	if m != nil {
		return m.OperationUrl
	}
	return ""
}

func (m *Withdrawal) GetOrderNo() string {
	if m != nil {
		return m.OrderNo
	}
	return ""
}

func (m *Withdrawal) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Withdrawal) GetCanceledAt() string {
	if m != nil {
		return m.CanceledAt
	}
	return ""
}

func (m *Withdrawal) GetSucceededAt() string {
	if m != nil {
		return m.SucceededAt
	}
	return ""
}

func (m *Withdrawal) GetTargetUserType() string {
	if m != nil {
		return m.TargetUserType
	}
	return ""
}

func (m *Withdrawal) GetTargetUserId() int64 {
	if m != nil {
		return m.TargetUserId
	}
	return 0
}

func (m *Withdrawal) GetServiceFee() float32 {
	if m != nil {
		return m.ServiceFee
	}
	return 0
}

func (m *Withdrawal) GetSettleAccountId() int64 {
	if m != nil {
		return m.SettleAccountId
	}
	return 0
}

func (m *Withdrawal) GetStoreId() int64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Withdrawal) GetMeno() string {
	if m != nil {
		return m.Meno
	}
	return ""
}

func (m *Withdrawal) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Withdrawal) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Withdrawal) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Withdrawal) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

//
type WithdrawalResponse struct {
	Entity               *Withdrawal   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager        `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Withdrawal `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error        `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info         `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *WithdrawalResponse) Reset()         { *m = WithdrawalResponse{} }
func (m *WithdrawalResponse) String() string { return proto.CompactTextString(m) }
func (*WithdrawalResponse) ProtoMessage()    {}
func (*WithdrawalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_729e0e9b3824e84b, []int{1}
}

func (m *WithdrawalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawalResponse.Unmarshal(m, b)
}
func (m *WithdrawalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawalResponse.Marshal(b, m, deterministic)
}
func (m *WithdrawalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawalResponse.Merge(m, src)
}
func (m *WithdrawalResponse) XXX_Size() int {
	return xxx_messageInfo_WithdrawalResponse.Size(m)
}
func (m *WithdrawalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawalResponse proto.InternalMessageInfo

func (m *WithdrawalResponse) GetEntity() *Withdrawal {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *WithdrawalResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *WithdrawalResponse) GetItems() []*Withdrawal {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *WithdrawalResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *WithdrawalResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Withdrawal)(nil), "geiqin.srv.pay.Withdrawal")
	proto.RegisterType((*WithdrawalResponse)(nil), "geiqin.srv.pay.WithdrawalResponse")
}

func init() {
	proto.RegisterFile("withdrawal.proto", fileDescriptor_729e0e9b3824e84b)
}

var fileDescriptor_729e0e9b3824e84b = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xda, 0x48,
	0x14, 0x5e, 0x20, 0x10, 0x72, 0x20, 0x04, 0x4e, 0x7e, 0x76, 0x82, 0xb4, 0x5a, 0x96, 0xed, 0x05,
	0x6a, 0x24, 0x5a, 0xd1, 0x27, 0xa0, 0x7f, 0x11, 0x95, 0x5a, 0x55, 0x4e, 0xa2, 0x5c, 0x5a, 0x13,
	0xfb, 0x00, 0x96, 0x6c, 0x8f, 0x3b, 0x33, 0x4e, 0xca, 0x33, 0xf4, 0x51, 0xfb, 0x0e, 0x55, 0x35,
	0xc7, 0x0e, 0x24, 0x51, 0x95, 0xf6, 0x22, 0x77, 0xfe, 0x7e, 0xe6, 0xe3, 0xf8, 0x1b, 0x1f, 0xa0,
	0x7b, 0x13, 0xd9, 0x65, 0xa8, 0xe5, 0x8d, 0x8c, 0xc7, 0x99, 0x56, 0x56, 0x61, 0x67, 0x41, 0xd1,
	0x97, 0x28, 0x1d, 0x1b, 0x7d, 0x3d, 0xce, 0xe4, 0xaa, 0xdf, 0x0e, 0x54, 0x92, 0xa8, 0xb4, 0x50,
	0x87, 0xdf, 0xeb, 0x00, 0x97, 0xeb, 0x23, 0xd8, 0x81, 0x6a, 0x14, 0x8a, 0xca, 0xa0, 0x32, 0xaa,
	0x79, 0xd5, 0x28, 0xc4, 0x23, 0x68, 0xc8, 0x44, 0xe5, 0xa9, 0x15, 0xd5, 0x41, 0x65, 0x54, 0xf5,
	0x4a, 0x84, 0x27, 0xd0, 0x93, 0xc6, 0x90, 0xf5, 0xad, 0x96, 0xa9, 0x91, 0x81, 0x8d, 0x54, 0x2a,
	0x6a, 0x83, 0xca, 0x68, 0xc7, 0xeb, 0xb2, 0x70, 0xbe, 0xe1, 0xf1, 0x05, 0xec, 0x5f, 0xc9, 0x58,
	0xa6, 0x01, 0xdd, 0xb3, 0x6f, 0xb1, 0x1d, 0x4b, 0xe9, 0xee, 0x01, 0x01, 0xdb, 0xc1, 0x52, 0xa6,
	0x29, 0xc5, 0xa2, 0xce, 0xa6, 0x5b, 0x88, 0x07, 0x50, 0xa7, 0xaf, 0x56, 0x4b, 0xd1, 0x60, 0xbe,
	0x00, 0xf8, 0x2f, 0xb4, 0xe6, 0x32, 0x8a, 0x73, 0x4d, 0x7e, 0x62, 0x16, 0x62, 0x9b, 0x35, 0x28,
	0xa9, 0x8f, 0x66, 0x81, 0x5d, 0xa8, 0xcd, 0x89, 0x44, 0x93, 0xdf, 0xc1, 0x3d, 0x62, 0x1f, 0x9a,
	0x71, 0x74, 0x4d, 0x89, 0x0a, 0x49, 0xec, 0x0c, 0x2a, 0xa3, 0xa6, 0xb7, 0xc6, 0xf8, 0x3f, 0xec,
	0xaa, 0x8c, 0xb4, 0x74, 0xb3, 0xf8, 0xb9, 0x8e, 0x05, 0x70, 0x60, 0x7b, 0x4d, 0x5e, 0xe8, 0x18,
	0x8f, 0xa1, 0xa9, 0x74, 0x48, 0xda, 0x4f, 0x95, 0x68, 0x15, 0x43, 0x32, 0xfe, 0xa4, 0x5c, 0x69,
	0x46, 0xe5, 0x3a, 0x20, 0xd1, 0x66, 0xa1, 0x44, 0x6e, 0xcc, 0xc0, 0xbd, 0x6a, 0x4c, 0xa1, 0x2f,
	0xad, 0xd8, 0x2d, 0xc6, 0xbc, 0xa5, 0xa6, 0x16, 0xff, 0x83, 0xb6, 0xc9, 0x83, 0x80, 0x28, 0x2c,
	0x1c, 0x1d, 0x76, 0xb4, 0xd6, 0xdc, 0xd4, 0xe2, 0x08, 0xba, 0x56, 0xea, 0x05, 0x59, 0x3f, 0x37,
	0xa4, 0x7d, 0xbb, 0xca, 0x48, 0xec, 0xb1, 0xad, 0x53, 0xf0, 0x17, 0x86, 0xf4, 0xf9, 0x2a, 0x23,
	0x7c, 0x06, 0x9d, 0xbb, 0xce, 0x28, 0x14, 0x5d, 0xbe, 0xd6, 0xf6, 0xc6, 0x37, 0x0b, 0xdd, 0x4c,
	0x86, 0xf4, 0x75, 0x14, 0x90, 0xef, 0x1a, 0xea, 0x71, 0x43, 0x50, 0x52, 0xef, 0x89, 0xf0, 0x39,
	0xf4, 0x0c, 0x59, 0x1b, 0x93, 0x2f, 0x83, 0xc0, 0xdd, 0xbd, 0x4b, 0x42, 0x4e, 0xda, 0x2b, 0x84,
	0x69, 0xc1, 0xcf, 0x42, 0xd7, 0x89, 0xb1, 0x4a, 0x93, 0xb3, 0xec, 0xb3, 0x65, 0x9b, 0xf1, 0x2c,
	0x44, 0x84, 0xad, 0x84, 0x52, 0x25, 0x0e, 0x78, 0x56, 0x7e, 0x76, 0x77, 0x90, 0x90, 0x95, 0xa1,
	0xb4, 0x52, 0x1c, 0x32, 0xbf, 0xc6, 0xdc, 0xa1, 0x95, 0x36, 0x37, 0xe2, 0xa8, 0xec, 0x90, 0x11,
	0xfe, 0x03, 0x10, 0x68, 0x92, 0xb6, 0x28, 0xe8, 0x6f, 0xd6, 0x76, 0x4a, 0x66, 0x6a, 0x9d, 0x9c,
	0x67, 0xe1, 0xad, 0x2c, 0x0a, 0xb9, 0x64, 0xa6, 0x76, 0xf8, 0xa3, 0x02, 0xb8, 0xf9, 0xda, 0x3d,
	0x32, 0x99, 0x4a, 0x0d, 0xe1, 0x04, 0x1a, 0x94, 0xda, 0xc8, 0xae, 0xf8, 0xcb, 0x6f, 0x4d, 0xfa,
	0xe3, 0xfb, 0x3b, 0x33, 0xbe, 0x73, 0xa6, 0x74, 0xe2, 0x09, 0xd4, 0x33, 0xb9, 0x20, 0xcd, 0x8b,
	0xd1, 0x9a, 0x1c, 0x3e, 0x3c, 0xf2, 0xd9, 0x89, 0x5e, 0xe1, 0xc1, 0x97, 0x50, 0x8f, 0x2c, 0x25,
	0x46, 0xd4, 0x06, 0xb5, 0xdf, 0xe4, 0x17, 0x46, 0x17, 0x4f, 0x5a, 0x2b, 0xcd, 0x5b, 0xf2, 0x8b,
	0xf8, 0x77, 0x4e, 0xf4, 0x0a, 0x0f, 0x8e, 0x60, 0x2b, 0x4a, 0xe7, 0x8a, 0x97, 0xa5, 0x35, 0x39,
	0x78, 0xe8, 0x9d, 0xa5, 0x73, 0xe5, 0xb1, 0x63, 0xf2, 0xad, 0x06, 0xbd, 0xcd, 0x8f, 0x9d, 0x15,
	0xd7, 0x8c, 0x1f, 0xa0, 0xf1, 0x86, 0x2b, 0xc4, 0x47, 0x26, 0xeb, 0x0f, 0x1f, 0x99, 0xba, 0x6c,
	0x72, 0xf8, 0x97, 0xcb, 0xba, 0xe0, 0xbe, 0x9f, 0x26, 0xeb, 0x2d, 0xc5, 0xf4, 0x24, 0x59, 0xa7,
	0x50, 0x3b, 0x25, 0xfb, 0x04, 0x41, 0x33, 0x68, 0x9c, 0x91, 0xd4, 0xc1, 0x12, 0x8f, 0x1f, 0xfa,
	0x5f, 0x4b, 0x43, 0x97, 0x4b, 0xd2, 0xf4, 0x67, 0x51, 0x57, 0x0d, 0xfe, 0x0f, 0x7e, 0xf5, 0x33,
	0x00, 0x00, 0xff, 0xff, 0x43, 0x1c, 0x37, 0xb2, 0xb5, 0x05, 0x00, 0x00,
}
