// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commission.proto

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

type Commission struct {
	Id                   int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CommissionSn         string              `protobuf:"bytes,2,opt,name=commission_sn,json=commissionSn,proto3" json:"commission_sn,omitempty"`
	OrderId              int64               `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderSn              string              `protobuf:"bytes,4,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
	PlatformSource       string              `protobuf:"bytes,5,opt,name=platform_source,json=platformSource,proto3" json:"platform_source,omitempty"`
	OrderAmount          float32             `protobuf:"fixed32,6,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	DistributorId        int64               `protobuf:"varint,7,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	InviteeId            int64               `protobuf:"varint,8,opt,name=invitee_id,json=inviteeId,proto3" json:"invitee_id,omitempty"`
	Direction            string              `protobuf:"bytes,9,opt,name=direction,proto3" json:"direction,omitempty"`
	Money                float32             `protobuf:"fixed32,10,opt,name=money,proto3" json:"money,omitempty"`
	IncomeType           string              `protobuf:"bytes,11,opt,name=income_type,json=incomeType,proto3" json:"income_type,omitempty"`
	IncomeRate           float32             `protobuf:"fixed32,12,opt,name=income_rate,json=incomeRate,proto3" json:"income_rate,omitempty"`
	Status               int32               `protobuf:"varint,15,opt,name=status,proto3" json:"status,omitempty"`
	Memo                 string              `protobuf:"bytes,16,opt,name=memo,proto3" json:"memo,omitempty"`
	OperatorId           int64               `protobuf:"varint,17,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	CreatedAt            string              `protobuf:"bytes,18,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string              `protobuf:"bytes,19,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Details              []*CommissionDetail `protobuf:"bytes,20,rep,name=details,proto3" json:"details,omitempty"`
	Invitee              *Distributor        `protobuf:"bytes,21,opt,name=invitee,proto3" json:"invitee,omitempty"`
	Distributor          *Distributor        `protobuf:"bytes,22,opt,name=distributor,proto3" json:"distributor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Commission) Reset()         { *m = Commission{} }
func (m *Commission) String() string { return proto.CompactTextString(m) }
func (*Commission) ProtoMessage()    {}
func (*Commission) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3823d0793ff9520, []int{0}
}

func (m *Commission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commission.Unmarshal(m, b)
}
func (m *Commission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commission.Marshal(b, m, deterministic)
}
func (m *Commission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commission.Merge(m, src)
}
func (m *Commission) XXX_Size() int {
	return xxx_messageInfo_Commission.Size(m)
}
func (m *Commission) XXX_DiscardUnknown() {
	xxx_messageInfo_Commission.DiscardUnknown(m)
}

var xxx_messageInfo_Commission proto.InternalMessageInfo

func (m *Commission) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Commission) GetCommissionSn() string {
	if m != nil {
		return m.CommissionSn
	}
	return ""
}

func (m *Commission) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Commission) GetOrderSn() string {
	if m != nil {
		return m.OrderSn
	}
	return ""
}

func (m *Commission) GetPlatformSource() string {
	if m != nil {
		return m.PlatformSource
	}
	return ""
}

func (m *Commission) GetOrderAmount() float32 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

func (m *Commission) GetDistributorId() int64 {
	if m != nil {
		return m.DistributorId
	}
	return 0
}

func (m *Commission) GetInviteeId() int64 {
	if m != nil {
		return m.InviteeId
	}
	return 0
}

func (m *Commission) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *Commission) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *Commission) GetIncomeType() string {
	if m != nil {
		return m.IncomeType
	}
	return ""
}

func (m *Commission) GetIncomeRate() float32 {
	if m != nil {
		return m.IncomeRate
	}
	return 0
}

func (m *Commission) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Commission) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Commission) GetOperatorId() int64 {
	if m != nil {
		return m.OperatorId
	}
	return 0
}

func (m *Commission) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Commission) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Commission) GetDetails() []*CommissionDetail {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Commission) GetInvitee() *Distributor {
	if m != nil {
		return m.Invitee
	}
	return nil
}

func (m *Commission) GetDistributor() *Distributor {
	if m != nil {
		return m.Distributor
	}
	return nil
}

type CommissionResponse struct {
	Entity               *Commission   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager        `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Commission `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error        `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info         `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CommissionResponse) Reset()         { *m = CommissionResponse{} }
func (m *CommissionResponse) String() string { return proto.CompactTextString(m) }
func (*CommissionResponse) ProtoMessage()    {}
func (*CommissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3823d0793ff9520, []int{1}
}

func (m *CommissionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommissionResponse.Unmarshal(m, b)
}
func (m *CommissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommissionResponse.Marshal(b, m, deterministic)
}
func (m *CommissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommissionResponse.Merge(m, src)
}
func (m *CommissionResponse) XXX_Size() int {
	return xxx_messageInfo_CommissionResponse.Size(m)
}
func (m *CommissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommissionResponse proto.InternalMessageInfo

func (m *CommissionResponse) GetEntity() *Commission {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *CommissionResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *CommissionResponse) GetItems() []*Commission {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CommissionResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CommissionResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Commission)(nil), "geiqin.srv.dms.Commission")
	proto.RegisterType((*CommissionResponse)(nil), "geiqin.srv.dms.CommissionResponse")
}

func init() {
	proto.RegisterFile("commission.proto", fileDescriptor_a3823d0793ff9520)
}

var fileDescriptor_a3823d0793ff9520 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xbf, 0x24, 0x4d, 0xd2, 0x8c, 0xd3, 0xb4, 0x99, 0xaf, 0xad, 0xb6, 0x01, 0x84, 0x09,
	0x42, 0x58, 0x42, 0x8a, 0x90, 0x11, 0x17, 0x24, 0x0e, 0xa5, 0x45, 0x55, 0x38, 0x55, 0x4e, 0x25,
	0x8e, 0x91, 0x6b, 0x4f, 0xdb, 0x95, 0xea, 0x5d, 0xb3, 0xbb, 0xa9, 0x94, 0x77, 0xe3, 0x05, 0x78,
	0x29, 0x84, 0x3c, 0x76, 0x70, 0x88, 0xa0, 0x20, 0xc1, 0xcd, 0xfb, 0x9f, 0xdf, 0xfc, 0x3d, 0xb3,
	0x3b, 0xbb, 0xb0, 0x97, 0xe8, 0x2c, 0x93, 0xd6, 0x4a, 0xad, 0x26, 0xb9, 0xd1, 0x4e, 0xe3, 0xe0,
	0x9a, 0xe4, 0x27, 0xa9, 0x26, 0xd6, 0xdc, 0x4d, 0xd2, 0xcc, 0x8e, 0xfa, 0x05, 0xb1, 0x8a, 0x8e,
	0x0e, 0x6b, 0xfe, 0x94, 0x5c, 0x2c, 0x6f, 0x2b, 0x7d, 0x98, 0x4a, 0xeb, 0x8c, 0xbc, 0x5c, 0x38,
	0x6d, 0x4a, 0x69, 0xfc, 0xb9, 0x0d, 0x70, 0xf2, 0x9d, 0xc6, 0x01, 0x34, 0x65, 0x2a, 0x1a, 0x7e,
	0x23, 0x68, 0x45, 0x4d, 0x99, 0xe2, 0x53, 0xd8, 0xa9, 0xbd, 0xe6, 0x56, 0x89, 0xa6, 0xdf, 0x08,
	0x7a, 0x51, 0xbf, 0x16, 0x67, 0x0a, 0x8f, 0x60, 0x5b, 0x9b, 0x94, 0xcc, 0x5c, 0xa6, 0xa2, 0xc5,
	0xa9, 0x5d, 0x5e, 0x4f, 0xd3, 0x3a, 0x64, 0x95, 0xd8, 0xe2, 0xd4, 0x32, 0x34, 0x53, 0xf8, 0x1c,
	0x76, 0xf3, 0xdb, 0xd8, 0x5d, 0x69, 0x93, 0xcd, 0xad, 0x5e, 0x98, 0x84, 0x44, 0x9b, 0x89, 0xc1,
	0x4a, 0x9e, 0xb1, 0x8a, 0x4f, 0xa0, 0x5f, 0x7a, 0xc4, 0x99, 0x5e, 0x28, 0x27, 0x3a, 0x7e, 0x23,
	0x68, 0x46, 0x1e, 0x6b, 0xc7, 0x2c, 0xe1, 0x33, 0x18, 0xac, 0xb5, 0x56, 0xd4, 0xd1, 0xe5, 0x3a,
	0x76, 0xd6, 0xd4, 0x69, 0x8a, 0x8f, 0x00, 0xa4, 0xba, 0x93, 0x8e, 0xa8, 0x40, 0xb6, 0x19, 0xe9,
	0x55, 0xca, 0x34, 0xc5, 0x87, 0xd0, 0x4b, 0xa5, 0xa1, 0xc4, 0x49, 0xad, 0x44, 0x8f, 0x6b, 0xa9,
	0x05, 0xdc, 0x87, 0x76, 0xa6, 0x15, 0x2d, 0x05, 0xf0, 0xff, 0xcb, 0x05, 0x3e, 0x06, 0x4f, 0xaa,
	0x44, 0x67, 0x34, 0x77, 0xcb, 0x9c, 0x84, 0xc7, 0x59, 0x50, 0x4a, 0x17, 0xcb, 0x9c, 0xd6, 0x00,
	0x13, 0x3b, 0x12, 0x7d, 0x4e, 0xae, 0x80, 0x28, 0x76, 0x84, 0x87, 0xd0, 0xb1, 0x2e, 0x76, 0x0b,
	0x2b, 0x76, 0xfd, 0x46, 0xd0, 0x8e, 0xaa, 0x15, 0x22, 0x6c, 0x65, 0x94, 0x69, 0xb1, 0xc7, 0x96,
	0xfc, 0x5d, 0x98, 0xe9, 0x9c, 0x4c, 0x5c, 0x35, 0x39, 0xe4, 0x0e, 0x60, 0x25, 0x95, 0x1d, 0x26,
	0x86, 0x62, 0x47, 0xe9, 0x3c, 0x76, 0x02, 0xcb, 0x1e, 0x2a, 0xe5, 0xd8, 0x15, 0xe1, 0x45, 0x9e,
	0xae, 0xc2, 0xff, 0x97, 0xe1, 0x4a, 0x39, 0x76, 0xf8, 0x06, 0xba, 0x29, 0xcf, 0x8b, 0x15, 0xfb,
	0x7e, 0x2b, 0xf0, 0x42, 0x7f, 0xf2, 0xe3, 0x9c, 0x4d, 0x4e, 0x36, 0x06, 0x2b, 0x5a, 0x25, 0xe0,
	0x6b, 0xe8, 0x56, 0x3b, 0x29, 0x0e, 0xfc, 0x46, 0xe0, 0x85, 0x0f, 0x36, 0x73, 0x4f, 0xeb, 0xb3,
	0x88, 0x56, 0x2c, 0xbe, 0x05, 0x6f, 0xed, 0x8c, 0xc4, 0xe1, 0xef, 0x53, 0xd7, 0xf9, 0xf1, 0xd7,
	0x06, 0x60, 0x5d, 0x53, 0x44, 0x36, 0xd7, 0xca, 0x12, 0x86, 0xd0, 0x21, 0xe5, 0xa4, 0x5b, 0xf2,
	0x28, 0x7b, 0xe1, 0xe8, 0xd7, 0x7d, 0x44, 0x15, 0x89, 0x2f, 0xa0, 0x9d, 0xc7, 0xd7, 0x64, 0x78,
	0xc4, 0xbd, 0xf0, 0x60, 0x33, 0xe5, 0xbc, 0x08, 0x46, 0x25, 0x83, 0x2f, 0xa1, 0x2d, 0x1d, 0x65,
	0x56, 0xb4, 0x78, 0x9f, 0xee, 0xf3, 0x2f, 0xc1, 0xc2, 0x9e, 0x8c, 0xd1, 0x86, 0xaf, 0xc1, 0x4f,
	0xec, 0xdf, 0x17, 0xc1, 0xa8, 0x64, 0x30, 0x80, 0x2d, 0xa9, 0xae, 0x34, 0x5f, 0x08, 0x2f, 0xdc,
	0xdf, 0x64, 0xa7, 0xea, 0x4a, 0x47, 0x4c, 0x84, 0x5f, 0x9a, 0x30, 0xac, 0x7f, 0x36, 0x23, 0x73,
	0x27, 0x13, 0xc2, 0x0f, 0xd0, 0x39, 0xe1, 0x43, 0xc7, 0x7b, 0x2a, 0x1b, 0x8d, 0xef, 0xa9, 0xba,
	0xda, 0xc9, 0xf1, 0x7f, 0x78, 0x06, 0xad, 0x33, 0x72, 0xff, 0xc0, 0x68, 0x0a, 0x9d, 0x19, 0xc5,
	0x26, 0xb9, 0xc1, 0xa3, 0x4d, 0xfe, 0x5d, 0x6c, 0xe9, 0xe3, 0x0d, 0x19, 0xfa, 0x43, 0xab, 0x0b,
	0x18, 0x5a, 0xb6, 0x3a, 0x27, 0x53, 0xbc, 0x14, 0xb1, 0x4a, 0xe8, 0xaf, 0x5d, 0x2f, 0x3b, 0xfc,
	0x24, 0xbe, 0xfa, 0x16, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xa2, 0x54, 0xe7, 0x6f, 0x05, 0x00, 0x00,
}
