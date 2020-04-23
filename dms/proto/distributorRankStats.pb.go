// Code generated by protoc-gen-go. DO NOT EDIT.
// source: distributorRankStats.proto

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

type DistributorRankStats struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DistributorId        int64    `protobuf:"varint,2,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	RankId               int32    `protobuf:"varint,3,opt,name=rank_id,json=rankId,proto3" json:"rank_id,omitempty"`
	RecommendNum         int32    `protobuf:"varint,4,opt,name=recommend_num,json=recommendNum,proto3" json:"recommend_num,omitempty"`
	CustomerNum          int32    `protobuf:"varint,5,opt,name=customer_num,json=customerNum,proto3" json:"customer_num,omitempty"`
	PromotionNum         int32    `protobuf:"varint,6,opt,name=promotion_num,json=promotionNum,proto3" json:"promotion_num,omitempty"`
	ConsumptionNum       int32    `protobuf:"varint,7,opt,name=consumption_num,json=consumptionNum,proto3" json:"consumption_num,omitempty"`
	PromotionAmount      float32  `protobuf:"fixed32,8,opt,name=promotion_amount,json=promotionAmount,proto3" json:"promotion_amount,omitempty"`
	ConsumptionAmount    float32  `protobuf:"fixed32,9,opt,name=consumption_amount,json=consumptionAmount,proto3" json:"consumption_amount,omitempty"`
	CreatedAt            string   `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributorRankStats) Reset()         { *m = DistributorRankStats{} }
func (m *DistributorRankStats) String() string { return proto.CompactTextString(m) }
func (*DistributorRankStats) ProtoMessage()    {}
func (*DistributorRankStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b48ec36e672b10c, []int{0}
}

func (m *DistributorRankStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributorRankStats.Unmarshal(m, b)
}
func (m *DistributorRankStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributorRankStats.Marshal(b, m, deterministic)
}
func (m *DistributorRankStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributorRankStats.Merge(m, src)
}
func (m *DistributorRankStats) XXX_Size() int {
	return xxx_messageInfo_DistributorRankStats.Size(m)
}
func (m *DistributorRankStats) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributorRankStats.DiscardUnknown(m)
}

var xxx_messageInfo_DistributorRankStats proto.InternalMessageInfo

func (m *DistributorRankStats) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DistributorRankStats) GetDistributorId() int64 {
	if m != nil {
		return m.DistributorId
	}
	return 0
}

func (m *DistributorRankStats) GetRankId() int32 {
	if m != nil {
		return m.RankId
	}
	return 0
}

func (m *DistributorRankStats) GetRecommendNum() int32 {
	if m != nil {
		return m.RecommendNum
	}
	return 0
}

func (m *DistributorRankStats) GetCustomerNum() int32 {
	if m != nil {
		return m.CustomerNum
	}
	return 0
}

func (m *DistributorRankStats) GetPromotionNum() int32 {
	if m != nil {
		return m.PromotionNum
	}
	return 0
}

func (m *DistributorRankStats) GetConsumptionNum() int32 {
	if m != nil {
		return m.ConsumptionNum
	}
	return 0
}

func (m *DistributorRankStats) GetPromotionAmount() float32 {
	if m != nil {
		return m.PromotionAmount
	}
	return 0
}

func (m *DistributorRankStats) GetConsumptionAmount() float32 {
	if m != nil {
		return m.ConsumptionAmount
	}
	return 0
}

func (m *DistributorRankStats) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *DistributorRankStats) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*DistributorRankStats)(nil), "geiqin.srv.dms.DistributorRankStats")
}

func init() {
	proto.RegisterFile("distributorRankStats.proto", fileDescriptor_7b48ec36e672b10c)
}

var fileDescriptor_7b48ec36e672b10c = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4e, 0x83, 0x30,
	0x18, 0xc7, 0x03, 0x38, 0x26, 0xdf, 0x06, 0xd3, 0xc6, 0x44, 0x62, 0x62, 0x82, 0x2e, 0x46, 0x3c,
	0xc8, 0xc5, 0x27, 0x20, 0xf1, 0xb2, 0x8b, 0x07, 0x7c, 0x00, 0xd2, 0xd1, 0xc6, 0x34, 0x4b, 0x5b,
	0x6c, 0xbf, 0xfa, 0x24, 0x3e, 0xb0, 0xa1, 0x63, 0xc8, 0x61, 0xc7, 0xef, 0xf7, 0xfb, 0xe5, 0x7f,
	0xf9, 0xe0, 0x8e, 0x09, 0x8b, 0x46, 0xec, 0x1d, 0x6a, 0xd3, 0x50, 0x75, 0xf8, 0x44, 0x8a, 0xb6,
	0xea, 0x8d, 0x46, 0x4d, 0xb2, 0x2f, 0x2e, 0xbe, 0x85, 0xaa, 0xac, 0xf9, 0xa9, 0x98, 0xb4, 0x8f,
	0xbf, 0x11, 0xdc, 0xbc, 0x9f, 0xc9, 0x49, 0x06, 0xa1, 0x60, 0x79, 0x50, 0x04, 0x65, 0xd4, 0x84,
	0x82, 0x91, 0x27, 0xc8, 0x66, 0xb3, 0xad, 0x60, 0x79, 0xe8, 0x5d, 0x3a, 0xa3, 0x3b, 0x46, 0x6e,
	0x61, 0x69, 0xa8, 0x3a, 0x0c, 0x3e, 0x2a, 0x82, 0x72, 0xd1, 0xc4, 0xc3, 0xb9, 0x63, 0x64, 0x0b,
	0xa9, 0xe1, 0x9d, 0x96, 0x92, 0x2b, 0xd6, 0x2a, 0x27, 0xf3, 0x0b, 0xaf, 0xd7, 0x13, 0xfc, 0x70,
	0x92, 0x3c, 0xc0, 0xba, 0x73, 0x16, 0xb5, 0xe4, 0xc6, 0x37, 0x0b, 0xdf, 0xac, 0x4e, 0x6c, 0x48,
	0xb6, 0x90, 0xf6, 0x46, 0x4b, 0x8d, 0x42, 0x2b, 0xdf, 0xc4, 0xc7, 0x9d, 0x09, 0x0e, 0xd1, 0x33,
	0x6c, 0x3a, 0xad, 0xac, 0x93, 0xfd, 0x94, 0x2d, 0x7d, 0x96, 0xcd, 0xf0, 0x10, 0xbe, 0xc0, 0xd5,
	0xff, 0x1a, 0x95, 0xda, 0x29, 0xcc, 0x2f, 0x8b, 0xa0, 0x0c, 0x9b, 0xcd, 0xc4, 0x6b, 0x8f, 0xc9,
	0x2b, 0x90, 0xf9, 0xe6, 0x18, 0x27, 0x3e, 0xbe, 0x9e, 0x99, 0x31, 0xbf, 0x07, 0xe8, 0x0c, 0xa7,
	0xc8, 0x59, 0x4b, 0x31, 0x87, 0x22, 0x28, 0x93, 0x26, 0x19, 0x49, 0xed, 0xb5, 0xeb, 0xd9, 0x49,
	0xaf, 0x8e, 0x7a, 0x24, 0x35, 0xee, 0x63, 0xff, 0xad, 0xb7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x60, 0x86, 0xb9, 0x78, 0xcb, 0x01, 0x00, 0x00,
}
