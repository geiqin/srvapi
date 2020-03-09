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
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RewardId             int64    `protobuf:"varint,2,opt,name=reward_id,json=rewardId,proto3" json:"reward_id,omitempty"`
	UnitType             int32    `protobuf:"varint,3,opt,name=unit_type,json=unitType,proto3" json:"unit_type,omitempty"`
	ConditionPrice       float32  `protobuf:"fixed32,4,opt,name=condition_price,json=conditionPrice,proto3" json:"condition_price,omitempty"`
	ConditionNum         int32    `protobuf:"varint,5,opt,name=condition_num,json=conditionNum,proto3" json:"condition_num,omitempty"`
	Preferent            int32    `protobuf:"varint,6,opt,name=preferent,proto3" json:"preferent,omitempty"`
	Money                float32  `protobuf:"fixed32,7,opt,name=money,proto3" json:"money,omitempty"`
	Discount             float32  `protobuf:"fixed32,8,opt,name=discount,proto3" json:"discount,omitempty"`
	Point                int32    `protobuf:"varint,9,opt,name=point,proto3" json:"point,omitempty"`
	PresentId            int64    `protobuf:"varint,10,opt,name=present_id,json=presentId,proto3" json:"present_id,omitempty"`
	PresentNum           int32    `protobuf:"varint,11,opt,name=present_num,json=presentNum,proto3" json:"present_num,omitempty"`
	Description          string   `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            string   `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *RewardStep) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
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
	return fileDescriptor_0aff996c4aaa35fc, []int{1}
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
	proto.RegisterType((*RewardStepResponse)(nil), "geiqin.srv.ims.reward.RewardStepResponse")
}

func init() { proto.RegisterFile("rewardStep.proto", fileDescriptor_0aff996c4aaa35fc) }

var fileDescriptor_0aff996c4aaa35fc = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x8e, 0x13, 0x31,
	0x10, 0x87, 0xb5, 0xbb, 0xb7, 0x21, 0x3b, 0x9b, 0x0b, 0xc8, 0x02, 0xc9, 0xba, 0x3b, 0xc4, 0x72,
	0x14, 0xa4, 0x5a, 0xa4, 0x50, 0x20, 0xca, 0x2b, 0x28, 0xd2, 0xa0, 0x93, 0xa1, 0x8f, 0xc2, 0x7a,
	0x72, 0x9a, 0x62, 0x6d, 0xe3, 0xf5, 0x82, 0xf2, 0x38, 0xb4, 0x3c, 0x25, 0xf2, 0x38, 0x7f, 0xae,
	0x20, 0xd2, 0x95, 0xfe, 0xe6, 0xfb, 0x79, 0xc6, 0x1a, 0xc3, 0x0b, 0x8f, 0xbf, 0x37, 0x5e, 0x7f,
	0x0b, 0xe8, 0x5a, 0xe7, 0x6d, 0xb0, 0xe2, 0xd5, 0x03, 0xd2, 0x4f, 0x32, 0xed, 0xe0, 0x7f, 0xb5,
	0xd4, 0x0f, 0x6d, 0x12, 0xae, 0x66, 0x9d, 0xed, 0x7b, 0x6b, 0x92, 0x74, 0xfb, 0xb7, 0x00, 0x50,
	0xc7, 0xa4, 0x98, 0x43, 0x4e, 0x5a, 0x66, 0x4d, 0xb6, 0x28, 0x54, 0x4e, 0x5a, 0x5c, 0x43, 0x95,
	0x62, 0x6b, 0xd2, 0x32, 0x67, 0x3c, 0x4d, 0x60, 0xc5, 0xc5, 0xd1, 0x50, 0x58, 0x87, 0x9d, 0x43,
	0x59, 0x34, 0xd9, 0xa2, 0x54, 0xd3, 0x08, 0xbe, 0xef, 0x1c, 0x8a, 0xf7, 0xf0, 0xbc, 0xb3, 0x46,
	0x53, 0x20, 0x6b, 0xd6, 0xce, 0x53, 0x87, 0xf2, 0xa2, 0xc9, 0x16, 0xb9, 0x9a, 0x1f, 0xf1, 0x7d,
	0xa4, 0xe2, 0x1d, 0x5c, 0x9e, 0x44, 0x33, 0xf6, 0xb2, 0xe4, 0x9b, 0x66, 0x47, 0xf8, 0x75, 0xec,
	0xc5, 0x0d, 0x54, 0xce, 0xe3, 0x16, 0x3d, 0x9a, 0x20, 0x27, 0x2c, 0x9c, 0x80, 0x78, 0x09, 0x65,
	0x6f, 0x0d, 0xee, 0xe4, 0x33, 0xee, 0x90, 0x0e, 0xe2, 0x0a, 0xa6, 0x9a, 0x86, 0xce, 0x8e, 0x26,
	0xc8, 0x29, 0x17, 0x8e, 0xe7, 0x98, 0x70, 0x96, 0x4c, 0x90, 0x15, 0xdf, 0x95, 0x0e, 0xe2, 0x35,
	0x80, 0xf3, 0x38, 0xa0, 0x09, 0xf1, 0xb9, 0xc0, 0xcf, 0xad, 0xf6, 0x64, 0xa5, 0xc5, 0x1b, 0xa8,
	0x0f, 0xe5, 0x38, 0x67, 0xcd, 0xd1, 0x43, 0x22, 0x4e, 0xd9, 0x40, 0xad, 0x71, 0xe8, 0x3c, 0xb9,
	0x38, 0xb7, 0x9c, 0x35, 0xd9, 0xa2, 0x52, 0x8f, 0x51, 0xec, 0xd0, 0x79, 0xdc, 0x04, 0xd4, 0xeb,
	0x4d, 0x90, 0x97, 0x2c, 0x54, 0x7b, 0x72, 0xc7, 0x03, 0x8c, 0x4e, 0x1f, 0xca, 0xf3, 0x54, 0xde,
	0x93, 0xbb, 0x70, 0xfb, 0x27, 0x07, 0x71, 0x5a, 0x96, 0xc2, 0xc1, 0x59, 0x33, 0xa0, 0xf8, 0x0c,
	0x13, 0x34, 0x81, 0xc2, 0x8e, 0x17, 0x57, 0x2f, 0xdf, 0xb6, 0xff, 0xdd, 0x7c, 0xfb, 0x28, 0xba,
	0x0f, 0x88, 0x25, 0x94, 0x6e, 0xf3, 0x80, 0x9e, 0x77, 0x5b, 0x2f, 0x6f, 0xce, 0x24, 0xef, 0xa3,
	0xa3, 0x92, 0x2a, 0x3e, 0x41, 0x49, 0x01, 0xfb, 0x41, 0x16, 0x4d, 0xf1, 0xb4, 0x6e, 0xc9, 0x8f,
	0xcd, 0xd0, 0x7b, 0xeb, 0xf9, 0x23, 0x9c, 0x6f, 0xf6, 0x25, 0x3a, 0x2a, 0xa9, 0xe2, 0x03, 0x5c,
	0x90, 0xd9, 0x5a, 0xfe, 0x14, 0xf5, 0xf2, 0xfa, 0x4c, 0x64, 0x65, 0xb6, 0x56, 0xb1, 0xf8, 0x63,
	0xc2, 0xff, 0xfa, 0xe3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x1e, 0x31, 0xa1, 0x10, 0x03,
	0x00, 0x00,
}