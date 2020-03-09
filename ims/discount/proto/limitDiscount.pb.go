// Code generated by protoc-gen-go. DO NOT EDIT.
// source: limitDiscount.proto

package geiqin_srv_ims_discount

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

type LimitDiscount struct {
	Id                   int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StartAt              string              `protobuf:"bytes,3,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                string              `protobuf:"bytes,4,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	PeriodType           int32               `protobuf:"varint,5,opt,name=period_type,json=periodType,proto3" json:"period_type,omitempty"`
	PeriodStartTime      string              `protobuf:"bytes,6,opt,name=period_start_time,json=periodStartTime,proto3" json:"period_start_time,omitempty"`
	PeriodEndTime        string              `protobuf:"bytes,7,opt,name=period_end_time,json=periodEndTime,proto3" json:"period_end_time,omitempty"`
	PeriodDays           string              `protobuf:"bytes,8,opt,name=period_days,json=periodDays,proto3" json:"period_days,omitempty"`
	LimitType            int32               `protobuf:"varint,9,opt,name=limit_type,json=limitType,proto3" json:"limit_type,omitempty"`
	LimitNum             int32               `protobuf:"varint,10,opt,name=limit_num,json=limitNum,proto3" json:"limit_num,omitempty"`
	Lable                string              `protobuf:"bytes,11,opt,name=lable,proto3" json:"lable,omitempty"`
	Erased               string              `protobuf:"bytes,12,opt,name=erased,proto3" json:"erased,omitempty"`
	Status               int32               `protobuf:"varint,13,opt,name=Status,proto3" json:"Status,omitempty"`
	CreatedAt            string              `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string              `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Goodses              *LimitDiscountGoods `protobuf:"bytes,16,opt,name=goodses,proto3" json:"goodses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LimitDiscount) Reset()         { *m = LimitDiscount{} }
func (m *LimitDiscount) String() string { return proto.CompactTextString(m) }
func (*LimitDiscount) ProtoMessage()    {}
func (*LimitDiscount) Descriptor() ([]byte, []int) {
	return fileDescriptor_3163fe83b6bcf4d5, []int{0}
}

func (m *LimitDiscount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LimitDiscount.Unmarshal(m, b)
}
func (m *LimitDiscount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LimitDiscount.Marshal(b, m, deterministic)
}
func (m *LimitDiscount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitDiscount.Merge(m, src)
}
func (m *LimitDiscount) XXX_Size() int {
	return xxx_messageInfo_LimitDiscount.Size(m)
}
func (m *LimitDiscount) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitDiscount.DiscardUnknown(m)
}

var xxx_messageInfo_LimitDiscount proto.InternalMessageInfo

func (m *LimitDiscount) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LimitDiscount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LimitDiscount) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *LimitDiscount) GetEndAt() string {
	if m != nil {
		return m.EndAt
	}
	return ""
}

func (m *LimitDiscount) GetPeriodType() int32 {
	if m != nil {
		return m.PeriodType
	}
	return 0
}

func (m *LimitDiscount) GetPeriodStartTime() string {
	if m != nil {
		return m.PeriodStartTime
	}
	return ""
}

func (m *LimitDiscount) GetPeriodEndTime() string {
	if m != nil {
		return m.PeriodEndTime
	}
	return ""
}

func (m *LimitDiscount) GetPeriodDays() string {
	if m != nil {
		return m.PeriodDays
	}
	return ""
}

func (m *LimitDiscount) GetLimitType() int32 {
	if m != nil {
		return m.LimitType
	}
	return 0
}

func (m *LimitDiscount) GetLimitNum() int32 {
	if m != nil {
		return m.LimitNum
	}
	return 0
}

func (m *LimitDiscount) GetLable() string {
	if m != nil {
		return m.Lable
	}
	return ""
}

func (m *LimitDiscount) GetErased() string {
	if m != nil {
		return m.Erased
	}
	return ""
}

func (m *LimitDiscount) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *LimitDiscount) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *LimitDiscount) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *LimitDiscount) GetGoodses() *LimitDiscountGoods {
	if m != nil {
		return m.Goodses
	}
	return nil
}

//
type LimitDiscountResponse struct {
	Entity               *LimitDiscount   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager           `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*LimitDiscount `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error           `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info            `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LimitDiscountResponse) Reset()         { *m = LimitDiscountResponse{} }
func (m *LimitDiscountResponse) String() string { return proto.CompactTextString(m) }
func (*LimitDiscountResponse) ProtoMessage()    {}
func (*LimitDiscountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3163fe83b6bcf4d5, []int{1}
}

func (m *LimitDiscountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LimitDiscountResponse.Unmarshal(m, b)
}
func (m *LimitDiscountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LimitDiscountResponse.Marshal(b, m, deterministic)
}
func (m *LimitDiscountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitDiscountResponse.Merge(m, src)
}
func (m *LimitDiscountResponse) XXX_Size() int {
	return xxx_messageInfo_LimitDiscountResponse.Size(m)
}
func (m *LimitDiscountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitDiscountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LimitDiscountResponse proto.InternalMessageInfo

func (m *LimitDiscountResponse) GetEntity() *LimitDiscount {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *LimitDiscountResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *LimitDiscountResponse) GetItems() []*LimitDiscount {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *LimitDiscountResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *LimitDiscountResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*LimitDiscount)(nil), "geiqin.srv.ims.discount.LimitDiscount")
	proto.RegisterType((*LimitDiscountResponse)(nil), "geiqin.srv.ims.discount.LimitDiscountResponse")
}

func init() { proto.RegisterFile("limitDiscount.proto", fileDescriptor_3163fe83b6bcf4d5) }

var fileDescriptor_3163fe83b6bcf4d5 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x6b, 0x13, 0x4d,
	0x14, 0x7e, 0xf3, 0xb1, 0xdb, 0xe6, 0x6c, 0x3f, 0x5e, 0xc7, 0x56, 0xc7, 0x4a, 0x35, 0xe4, 0xa2,
	0x04, 0x85, 0x05, 0xa3, 0x97, 0x22, 0xc4, 0xb6, 0x14, 0xa1, 0x88, 0x6c, 0x2a, 0xde, 0x19, 0x26,
	0xbb, 0xa7, 0xe9, 0x40, 0x76, 0x66, 0x9d, 0x99, 0x14, 0x72, 0xeb, 0x6f, 0xf2, 0x3f, 0xf9, 0x37,
	0x64, 0xce, 0x6c, 0xd1, 0x20, 0x91, 0xf4, 0x22, 0x77, 0x39, 0xcf, 0xc7, 0x79, 0xce, 0xce, 0x39,
	0x04, 0x1e, 0xce, 0x64, 0x29, 0xdd, 0x99, 0xb4, 0xb9, 0x9e, 0x2b, 0x97, 0x56, 0x46, 0x3b, 0xcd,
	0x1e, 0x4f, 0x51, 0x7e, 0x93, 0x2a, 0xb5, 0xe6, 0x36, 0x95, 0xa5, 0x4d, 0x8b, 0x9a, 0x3e, 0xda,
	0xc9, 0x75, 0x59, 0x6a, 0x15, 0x64, 0x47, 0x7c, 0xc9, 0x7b, 0xa1, 0x75, 0x61, 0x03, 0xd3, 0xfb,
	0xde, 0x86, 0xdd, 0xcb, 0x3f, 0x49, 0xb6, 0x07, 0x4d, 0x59, 0xf0, 0x46, 0xb7, 0xd1, 0x6f, 0x65,
	0x4d, 0x59, 0x30, 0x06, 0x6d, 0x25, 0x4a, 0xe4, 0xcd, 0x6e, 0xa3, 0xdf, 0xc9, 0xe8, 0x37, 0x7b,
	0x02, 0xdb, 0xd6, 0x09, 0xe3, 0xc6, 0xc2, 0xf1, 0x16, 0xe1, 0x5b, 0x54, 0x0f, 0x1d, 0x3b, 0x84,
	0x18, 0x55, 0xe1, 0x89, 0x36, 0x11, 0x11, 0xaa, 0x62, 0xe8, 0xd8, 0x73, 0x48, 0x2a, 0x34, 0x52,
	0x17, 0x63, 0xb7, 0xa8, 0x90, 0x47, 0xdd, 0x46, 0x3f, 0xca, 0x20, 0x40, 0x57, 0x8b, 0x0a, 0xd9,
	0x0b, 0x78, 0x50, 0x0b, 0x42, 0x67, 0x27, 0x4b, 0xe4, 0x31, 0xb5, 0xd8, 0x0f, 0xc4, 0xc8, 0xe3,
	0x57, 0xb2, 0x44, 0x76, 0x02, 0x35, 0x34, 0xf6, 0x51, 0xa4, 0xdc, 0x22, 0xe5, 0x6e, 0x80, 0xcf,
	0x55, 0x41, 0xba, 0xdf, 0xa1, 0x85, 0x58, 0x58, 0xbe, 0x4d, 0x9a, 0x3a, 0xf4, 0x4c, 0x2c, 0x2c,
	0x3b, 0x06, 0xa0, 0x97, 0x09, 0x43, 0x75, 0x68, 0xa8, 0x0e, 0x21, 0x34, 0xd3, 0x53, 0x08, 0xc5,
	0x58, 0xcd, 0x4b, 0x0e, 0xc4, 0x6e, 0x13, 0xf0, 0x71, 0x5e, 0xb2, 0x03, 0x88, 0x66, 0x62, 0x32,
	0x43, 0x9e, 0x84, 0xef, 0xa4, 0x82, 0x3d, 0x82, 0x18, 0x8d, 0xb0, 0x58, 0xf0, 0x1d, 0x82, 0xeb,
	0xca, 0xe3, 0x23, 0x27, 0xdc, 0xdc, 0xf2, 0x5d, 0xea, 0x53, 0x57, 0x7e, 0x82, 0xdc, 0xa0, 0x70,
	0x48, 0x4f, 0xb6, 0x47, 0x9e, 0x4e, 0x8d, 0x0c, 0x9d, 0xa7, 0xe7, 0x55, 0x71, 0x47, 0xef, 0x07,
	0xba, 0x46, 0x86, 0x8e, 0x9d, 0xc3, 0xd6, 0xd4, 0x2f, 0x13, 0x2d, 0xff, 0xbf, 0xdb, 0xe8, 0x27,
	0x83, 0x97, 0xe9, 0x8a, 0x83, 0x48, 0x2f, 0xff, 0xba, 0x80, 0xec, 0xce, 0xdb, 0xfb, 0xd1, 0x84,
	0xc3, 0x25, 0x3e, 0x43, 0x5b, 0x69, 0x65, 0x91, 0xbd, 0xf3, 0xdb, 0x74, 0xd2, 0x2d, 0xe8, 0x20,
	0x92, 0xc1, 0xc9, 0x7a, 0xfd, 0xb3, 0xda, 0xc5, 0xde, 0x40, 0x54, 0x89, 0x29, 0x1a, 0xba, 0x9e,
	0x64, 0xf0, 0x6c, 0xa5, 0xfd, 0x93, 0x57, 0x65, 0x41, 0xcc, 0xde, 0x42, 0x24, 0x1d, 0x96, 0x96,
	0xb7, 0xba, 0xad, 0x7b, 0x84, 0x06, 0x93, 0xcf, 0x44, 0x63, 0xb4, 0xa1, 0x03, 0xfc, 0x57, 0xe6,
	0xb9, 0x57, 0x65, 0x41, 0xcc, 0x5e, 0x41, 0x5b, 0xaa, 0x6b, 0x4d, 0x97, 0x99, 0x0c, 0x8e, 0x57,
	0x9a, 0x3e, 0xa8, 0x6b, 0x9d, 0x91, 0x74, 0xf0, 0xb3, 0x0d, 0x07, 0x4b, 0x13, 0x8c, 0xd0, 0xdc,
	0xca, 0x1c, 0xd9, 0x04, 0xe2, 0x53, 0x5a, 0x21, 0x5b, 0x73, 0xf4, 0xa3, 0x74, 0xcd, 0x4f, 0xac,
	0xf7, 0xd2, 0xfb, 0xcf, 0x67, 0x7c, 0xa6, 0x3b, 0xd8, 0x6c, 0xc6, 0x19, 0xce, 0x70, 0xd3, 0x19,
	0xa7, 0x42, 0xe5, 0x38, 0xdb, 0x60, 0xc6, 0x18, 0x5a, 0x17, 0xe8, 0x36, 0x18, 0xf0, 0x15, 0xe2,
	0x11, 0x0a, 0x93, 0xdf, 0xb0, 0xde, 0x4a, 0xef, 0x7b, 0x61, 0xf1, 0xcb, 0x0d, 0x1a, 0xbc, 0x7f,
	0xff, 0x49, 0x4c, 0x7f, 0xd6, 0xaf, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x43, 0x59, 0xe5, 0x0b,
	0x04, 0x06, 0x00, 0x00,
}