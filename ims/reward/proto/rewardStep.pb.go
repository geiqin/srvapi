// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rewardStep.proto

package geiqin_srv_ims_reward

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

type RewardStep struct {
	Id                   int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RewardId             int64           `protobuf:"varint,2,opt,name=reward_id,json=rewardId,proto3" json:"reward_id,omitempty"`
	UnitType             int32           `protobuf:"varint,3,opt,name=unit_type,json=unitType,proto3" json:"unit_type,omitempty"`
	ConditionPrice       float32         `protobuf:"fixed32,4,opt,name=condition_price,json=conditionPrice,proto3" json:"condition_price,omitempty"`
	ConditionNum         int32           `protobuf:"varint,5,opt,name=condition_num,json=conditionNum,proto3" json:"condition_num,omitempty"`
	Preferent            int32           `protobuf:"varint,6,opt,name=preferent,proto3" json:"preferent,omitempty"`
	Money                float32         `protobuf:"fixed32,7,opt,name=money,proto3" json:"money,omitempty"`
	Discount             float32         `protobuf:"fixed32,8,opt,name=discount,proto3" json:"discount,omitempty"`
	Point                int32           `protobuf:"varint,9,opt,name=point,proto3" json:"point,omitempty"`
	PresentId            int64           `protobuf:"varint,10,opt,name=present_id,json=presentId,proto3" json:"present_id,omitempty"`
	PresentNum           int32           `protobuf:"varint,11,opt,name=present_num,json=presentNum,proto3" json:"present_num,omitempty"`
	CreatedAt            string          `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string          `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Coupons              []*RewardCoupon `protobuf:"bytes,14,rep,name=coupons,proto3" json:"coupons,omitempty"`
	FreePostage          bool            `protobuf:"varint,15,opt,name=free_postage,json=freePostage,proto3" json:"free_postage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RewardStep) Reset()         { *m = RewardStep{} }
func (m *RewardStep) String() string { return proto.CompactTextString(m) }
func (*RewardStep) ProtoMessage()    {}
func (*RewardStep) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aff996c4aaa35fc, []int{0}
}

func (m *RewardStep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewardStep.Unmarshal(m, b)
}
func (m *RewardStep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewardStep.Marshal(b, m, deterministic)
}
func (m *RewardStep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardStep.Merge(m, src)
}
func (m *RewardStep) XXX_Size() int {
	return xxx_messageInfo_RewardStep.Size(m)
}
func (m *RewardStep) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardStep.DiscardUnknown(m)
}

var xxx_messageInfo_RewardStep proto.InternalMessageInfo

func (m *RewardStep) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RewardStep) GetRewardId() int64 {
	if m != nil {
		return m.RewardId
	}
	return 0
}

func (m *RewardStep) GetUnitType() int32 {
	if m != nil {
		return m.UnitType
	}
	return 0
}

func (m *RewardStep) GetConditionPrice() float32 {
	if m != nil {
		return m.ConditionPrice
	}
	return 0
}

func (m *RewardStep) GetConditionNum() int32 {
	if m != nil {
		return m.ConditionNum
	}
	return 0
}

func (m *RewardStep) GetPreferent() int32 {
	if m != nil {
		return m.Preferent
	}
	return 0
}

func (m *RewardStep) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *RewardStep) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *RewardStep) GetPoint() int32 {
	if m != nil {
		return m.Point
	}
	return 0
}

func (m *RewardStep) GetPresentId() int64 {
	if m != nil {
		return m.PresentId
	}
	return 0
}

func (m *RewardStep) GetPresentNum() int32 {
	if m != nil {
		return m.PresentNum
	}
	return 0
}

func (m *RewardStep) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *RewardStep) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *RewardStep) GetCoupons() []*RewardCoupon {
	if m != nil {
		return m.Coupons
	}
	return nil
}

func (m *RewardStep) GetFreePostage() bool {
	if m != nil {
		return m.FreePostage
	}
	return false
}

type RewardCoupon struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RewardStepId         int64    `protobuf:"varint,2,opt,name=reward_step_id,json=rewardStepId,proto3" json:"reward_step_id,omitempty"`
	CouponId             int64    `protobuf:"varint,3,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	CouponNum            int32    `protobuf:"varint,4,opt,name=coupon_num,json=couponNum,proto3" json:"coupon_num,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RewardCoupon) Reset()         { *m = RewardCoupon{} }
func (m *RewardCoupon) String() string { return proto.CompactTextString(m) }
func (*RewardCoupon) ProtoMessage()    {}
func (*RewardCoupon) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aff996c4aaa35fc, []int{1}
}

func (m *RewardCoupon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewardCoupon.Unmarshal(m, b)
}
func (m *RewardCoupon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewardCoupon.Marshal(b, m, deterministic)
}
func (m *RewardCoupon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardCoupon.Merge(m, src)
}
func (m *RewardCoupon) XXX_Size() int {
	return xxx_messageInfo_RewardCoupon.Size(m)
}
func (m *RewardCoupon) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardCoupon.DiscardUnknown(m)
}

var xxx_messageInfo_RewardCoupon proto.InternalMessageInfo

func (m *RewardCoupon) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RewardCoupon) GetRewardStepId() int64 {
	if m != nil {
		return m.RewardStepId
	}
	return 0
}

func (m *RewardCoupon) GetCouponId() int64 {
	if m != nil {
		return m.CouponId
	}
	return 0
}

func (m *RewardCoupon) GetCouponNum() int32 {
	if m != nil {
		return m.CouponNum
	}
	return 0
}

func (m *RewardCoupon) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *RewardCoupon) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

//
type RewardStepResponse struct {
	Entity               *RewardStep   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager        `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*RewardStep `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error        `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info         `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RewardStepResponse) Reset()         { *m = RewardStepResponse{} }
func (m *RewardStepResponse) String() string { return proto.CompactTextString(m) }
func (*RewardStepResponse) ProtoMessage()    {}
func (*RewardStepResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aff996c4aaa35fc, []int{2}
}

func (m *RewardStepResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewardStepResponse.Unmarshal(m, b)
}
func (m *RewardStepResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewardStepResponse.Marshal(b, m, deterministic)
}
func (m *RewardStepResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardStepResponse.Merge(m, src)
}
func (m *RewardStepResponse) XXX_Size() int {
	return xxx_messageInfo_RewardStepResponse.Size(m)
}
func (m *RewardStepResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardStepResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RewardStepResponse proto.InternalMessageInfo

func (m *RewardStepResponse) GetEntity() *RewardStep {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *RewardStepResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *RewardStepResponse) GetItems() []*RewardStep {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *RewardStepResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *RewardStepResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*RewardStep)(nil), "geiqin.srv.ims.reward.RewardStep")
	proto.RegisterType((*RewardCoupon)(nil), "geiqin.srv.ims.reward.RewardCoupon")
	proto.RegisterType((*RewardStepResponse)(nil), "geiqin.srv.ims.reward.RewardStepResponse")
}

func init() {
	proto.RegisterFile("rewardStep.proto", fileDescriptor_0aff996c4aaa35fc)
}

var fileDescriptor_0aff996c4aaa35fc = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x95, 0xa6, 0xe9, 0x36, 0x93, 0x6c, 0x17, 0x59, 0x20, 0x59, 0xfb, 0x47, 0x64, 0xbb,
	0x48, 0xe4, 0x14, 0xa4, 0x72, 0x40, 0x1c, 0x38, 0xac, 0x10, 0x87, 0x5e, 0x50, 0x65, 0xb8, 0x57,
	0x25, 0x99, 0x56, 0x3e, 0xc4, 0x36, 0x8e, 0x03, 0xea, 0xe3, 0xf0, 0x22, 0x3c, 0x10, 0x4f, 0x81,
	0x3c, 0xee, 0x1f, 0xc4, 0xd2, 0x15, 0x47, 0xff, 0xe6, 0xfb, 0x3c, 0x99, 0xf9, 0x1c, 0x78, 0x62,
	0xf1, 0xfb, 0xca, 0x36, 0x9f, 0x1c, 0x9a, 0xca, 0x58, 0xed, 0x34, 0x7b, 0xb6, 0x41, 0xf9, 0x55,
	0xaa, 0xaa, 0xb3, 0xdf, 0x2a, 0xd9, 0x76, 0x55, 0x10, 0x5c, 0xe6, 0xb5, 0x6e, 0x5b, 0xad, 0x82,
	0x68, 0xfa, 0x2b, 0x06, 0x10, 0x07, 0x27, 0x9b, 0xc0, 0x40, 0x36, 0x3c, 0x2a, 0xa2, 0x32, 0x16,
	0x03, 0xd9, 0xb0, 0x2b, 0x48, 0x83, 0x6d, 0x29, 0x1b, 0x3e, 0x20, 0x3c, 0x0e, 0x60, 0x4e, 0xc5,
	0x5e, 0x49, 0xb7, 0x74, 0x5b, 0x83, 0x3c, 0x2e, 0xa2, 0x32, 0x11, 0x63, 0x0f, 0x3e, 0x6f, 0x0d,
	0xb2, 0x97, 0x70, 0x51, 0x6b, 0xd5, 0x48, 0x27, 0xb5, 0x5a, 0x1a, 0x2b, 0x6b, 0xe4, 0xc3, 0x22,
	0x2a, 0x07, 0x62, 0x72, 0xc0, 0x0b, 0x4f, 0xd9, 0x1d, 0x9c, 0x1f, 0x85, 0xaa, 0x6f, 0x79, 0x42,
	0x37, 0xe5, 0x07, 0xf8, 0xb1, 0x6f, 0xd9, 0x35, 0xa4, 0xc6, 0xe2, 0x1a, 0x2d, 0x2a, 0xc7, 0x47,
	0x24, 0x38, 0x02, 0xf6, 0x14, 0x92, 0x56, 0x2b, 0xdc, 0xf2, 0x33, 0xea, 0x10, 0x0e, 0xec, 0x12,
	0xc6, 0x8d, 0xec, 0x6a, 0xdd, 0x2b, 0xc7, 0xc7, 0x54, 0x38, 0x9c, 0xbd, 0xc3, 0x68, 0xa9, 0x1c,
	0x4f, 0xe9, 0xae, 0x70, 0x60, 0x37, 0x00, 0xc6, 0x62, 0x87, 0xca, 0xf9, 0x71, 0x81, 0xc6, 0x4d,
	0x77, 0x64, 0xde, 0xb0, 0xe7, 0x90, 0xed, 0xcb, 0xfe, 0x3b, 0x33, 0xb2, 0xee, 0x1d, 0xfe, 0x2b,
	0x6f, 0x00, 0x6a, 0x8b, 0x2b, 0x87, 0xcd, 0x72, 0xe5, 0x78, 0x5e, 0x44, 0x65, 0x2a, 0xd2, 0x1d,
	0xb9, 0xa7, 0xeb, 0x7b, 0xd3, 0xec, 0xcb, 0xe7, 0xa1, 0xbc, 0x23, 0xf7, 0x8e, 0xbd, 0x83, 0xb3,
	0x5a, 0xf7, 0x46, 0xab, 0x8e, 0x4f, 0x8a, 0xb8, 0xcc, 0x66, 0x77, 0xd5, 0x3f, 0x13, 0xac, 0x42,
	0x5e, 0xef, 0x49, 0x2b, 0xf6, 0x1e, 0x76, 0x0b, 0xf9, 0xda, 0x22, 0x2e, 0x8d, 0xee, 0xdc, 0x6a,
	0x83, 0xfc, 0xa2, 0x88, 0xca, 0xb1, 0xc8, 0x3c, 0x5b, 0x04, 0x34, 0xfd, 0x19, 0x41, 0xfe, 0xa7,
	0xf9, 0x41, 0xdc, 0x2f, 0x60, 0xb2, 0x8b, 0xbb, 0x73, 0x68, 0x8e, 0x99, 0xe7, 0xc7, 0xc7, 0x15,
	0x72, 0x0f, 0x4d, 0xbd, 0x20, 0x0e, 0x8f, 0x22, 0x80, 0x79, 0x43, 0x3b, 0x08, 0x45, 0xbf, 0xa3,
	0x61, 0x88, 0x2a, 0x90, 0x87, 0x2b, 0x4a, 0x1e, 0x5f, 0xd1, 0xe8, 0xaf, 0x15, 0x4d, 0x7f, 0x0c,
	0x80, 0x1d, 0x5f, 0xab, 0xc0, 0xce, 0x4f, 0x8e, 0xec, 0x2d, 0x8c, 0x50, 0x39, 0xe9, 0xb6, 0x34,
	0x4a, 0x36, 0xbb, 0x7d, 0x74, 0x71, 0x64, 0xdd, 0x19, 0xd8, 0x0c, 0x12, 0xb3, 0xda, 0xa0, 0xa5,
	0x41, 0xb3, 0xd9, 0xf5, 0x09, 0xe7, 0xc2, 0x6b, 0x44, 0x90, 0xb2, 0x37, 0x90, 0x48, 0x87, 0x6d,
	0xc7, 0x63, 0x8a, 0xe9, 0x3f, 0xba, 0x05, 0xbd, 0x6f, 0x86, 0xd6, 0x6a, 0x4b, 0x6b, 0x39, 0xdd,
	0xec, 0x83, 0xd7, 0x88, 0x20, 0x65, 0xaf, 0x60, 0x28, 0xd5, 0x5a, 0xd3, 0xaa, 0xb2, 0xd9, 0xd5,
	0x09, 0xcb, 0x5c, 0xad, 0xb5, 0x20, 0xe1, 0x97, 0x11, 0xfd, 0xd8, 0xaf, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x0c, 0x5b, 0x95, 0xe8, 0x11, 0x04, 0x00, 0x00,
}
