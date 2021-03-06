// Code generated by protoc-gen-go. DO NOT EDIT.
// source: integralRule.proto

package geiqin_srv_crm_integral

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

type IntegralRuleWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting              string   `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Keywords             string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id                   int64    `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegralRuleWhere) Reset()         { *m = IntegralRuleWhere{} }
func (m *IntegralRuleWhere) String() string { return proto.CompactTextString(m) }
func (*IntegralRuleWhere) ProtoMessage()    {}
func (*IntegralRuleWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_89448b0b7b8f9508, []int{0}
}

func (m *IntegralRuleWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegralRuleWhere.Unmarshal(m, b)
}
func (m *IntegralRuleWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegralRuleWhere.Marshal(b, m, deterministic)
}
func (m *IntegralRuleWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegralRuleWhere.Merge(m, src)
}
func (m *IntegralRuleWhere) XXX_Size() int {
	return xxx_messageInfo_IntegralRuleWhere.Size(m)
}
func (m *IntegralRuleWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegralRuleWhere.DiscardUnknown(m)
}

var xxx_messageInfo_IntegralRuleWhere proto.InternalMessageInfo

func (m *IntegralRuleWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *IntegralRuleWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *IntegralRuleWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *IntegralRuleWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *IntegralRuleWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IntegralRuleWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type IntegralRule struct {
	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Points               int32            `protobuf:"varint,2,opt,name=points,proto3" json:"points,omitempty"`
	ReceiveType          int32            `protobuf:"varint,3,opt,name=receive_type,json=receiveType,proto3" json:"receive_type,omitempty"`
	DealCount            int32            `protobuf:"varint,4,opt,name=deal_count,json=dealCount,proto3" json:"deal_count,omitempty"`
	DealAmount           float32          `protobuf:"fixed32,5,opt,name=deal_amount,json=dealAmount,proto3" json:"deal_amount,omitempty"`
	GoodsType            int32            `protobuf:"varint,6,opt,name=goods_type,json=goodsType,proto3" json:"goods_type,omitempty"`
	CreatedAt            string           `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string           `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Goods                []*IntegralGoods `protobuf:"bytes,9,rep,name=goods,proto3" json:"goods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IntegralRule) Reset()         { *m = IntegralRule{} }
func (m *IntegralRule) String() string { return proto.CompactTextString(m) }
func (*IntegralRule) ProtoMessage()    {}
func (*IntegralRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_89448b0b7b8f9508, []int{1}
}

func (m *IntegralRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegralRule.Unmarshal(m, b)
}
func (m *IntegralRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegralRule.Marshal(b, m, deterministic)
}
func (m *IntegralRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegralRule.Merge(m, src)
}
func (m *IntegralRule) XXX_Size() int {
	return xxx_messageInfo_IntegralRule.Size(m)
}
func (m *IntegralRule) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegralRule.DiscardUnknown(m)
}

var xxx_messageInfo_IntegralRule proto.InternalMessageInfo

func (m *IntegralRule) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IntegralRule) GetPoints() int32 {
	if m != nil {
		return m.Points
	}
	return 0
}

func (m *IntegralRule) GetReceiveType() int32 {
	if m != nil {
		return m.ReceiveType
	}
	return 0
}

func (m *IntegralRule) GetDealCount() int32 {
	if m != nil {
		return m.DealCount
	}
	return 0
}

func (m *IntegralRule) GetDealAmount() float32 {
	if m != nil {
		return m.DealAmount
	}
	return 0
}

func (m *IntegralRule) GetGoodsType() int32 {
	if m != nil {
		return m.GoodsType
	}
	return 0
}

func (m *IntegralRule) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *IntegralRule) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *IntegralRule) GetGoods() []*IntegralGoods {
	if m != nil {
		return m.Goods
	}
	return nil
}

type IntegralGoods struct {
	Id                   int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId               int64      `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64      `protobuf:"varint,3,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Value                int32      `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
	Ratio                int32      `protobuf:"varint,5,opt,name=ratio,proto3" json:"ratio,omitempty"`
	IntegralRuleId       int64      `protobuf:"varint,6,opt,name=integral_rule_id,json=integralRuleId,proto3" json:"integral_rule_id,omitempty"`
	CreatedAt            string     `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string     `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Goods                *GoodsInfo `protobuf:"bytes,9,opt,name=goods,proto3" json:"goods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IntegralGoods) Reset()         { *m = IntegralGoods{} }
func (m *IntegralGoods) String() string { return proto.CompactTextString(m) }
func (*IntegralGoods) ProtoMessage()    {}
func (*IntegralGoods) Descriptor() ([]byte, []int) {
	return fileDescriptor_89448b0b7b8f9508, []int{2}
}

func (m *IntegralGoods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegralGoods.Unmarshal(m, b)
}
func (m *IntegralGoods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegralGoods.Marshal(b, m, deterministic)
}
func (m *IntegralGoods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegralGoods.Merge(m, src)
}
func (m *IntegralGoods) XXX_Size() int {
	return xxx_messageInfo_IntegralGoods.Size(m)
}
func (m *IntegralGoods) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegralGoods.DiscardUnknown(m)
}

var xxx_messageInfo_IntegralGoods proto.InternalMessageInfo

func (m *IntegralGoods) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IntegralGoods) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *IntegralGoods) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *IntegralGoods) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *IntegralGoods) GetRatio() int32 {
	if m != nil {
		return m.Ratio
	}
	return 0
}

func (m *IntegralGoods) GetIntegralRuleId() int64 {
	if m != nil {
		return m.IntegralRuleId
	}
	return 0
}

func (m *IntegralGoods) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *IntegralGoods) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *IntegralGoods) GetGoods() *GoodsInfo {
	if m != nil {
		return m.Goods
	}
	return nil
}

type IntegralRuleResponse struct {
	Entity               *IntegralRule   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager          `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*IntegralRule `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error          `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info           `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IntegralRuleResponse) Reset()         { *m = IntegralRuleResponse{} }
func (m *IntegralRuleResponse) String() string { return proto.CompactTextString(m) }
func (*IntegralRuleResponse) ProtoMessage()    {}
func (*IntegralRuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89448b0b7b8f9508, []int{3}
}

func (m *IntegralRuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegralRuleResponse.Unmarshal(m, b)
}
func (m *IntegralRuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegralRuleResponse.Marshal(b, m, deterministic)
}
func (m *IntegralRuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegralRuleResponse.Merge(m, src)
}
func (m *IntegralRuleResponse) XXX_Size() int {
	return xxx_messageInfo_IntegralRuleResponse.Size(m)
}
func (m *IntegralRuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegralRuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntegralRuleResponse proto.InternalMessageInfo

func (m *IntegralRuleResponse) GetEntity() *IntegralRule {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *IntegralRuleResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *IntegralRuleResponse) GetItems() []*IntegralRule {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *IntegralRuleResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *IntegralRuleResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*IntegralRuleWhere)(nil), "geiqin.srv.crm.integral.IntegralRuleWhere")
	proto.RegisterType((*IntegralRule)(nil), "geiqin.srv.crm.integral.IntegralRule")
	proto.RegisterType((*IntegralGoods)(nil), "geiqin.srv.crm.integral.IntegralGoods")
	proto.RegisterType((*IntegralRuleResponse)(nil), "geiqin.srv.crm.integral.IntegralRuleResponse")
}

func init() {
	proto.RegisterFile("integralRule.proto", fileDescriptor_89448b0b7b8f9508)
}

var fileDescriptor_89448b0b7b8f9508 = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0xbd, 0xb6, 0x6b, 0xb7, 0x19, 0xf7, 0xf6, 0xf6, 0xee, 0xed, 0xa5, 0x56, 0x50, 0x21, 0x44,
	0x02, 0x59, 0x48, 0x44, 0x22, 0xf0, 0x80, 0x04, 0x3c, 0x54, 0x05, 0x55, 0x79, 0x43, 0x5b, 0x10,
	0x8f, 0xc1, 0xf5, 0x4e, 0xd3, 0x55, 0x13, 0xaf, 0x59, 0xaf, 0x83, 0xd2, 0x5f, 0xe0, 0x1b, 0xe0,
	0x1f, 0x90, 0xf8, 0x40, 0xb4, 0xb3, 0x76, 0x09, 0x42, 0x2d, 0x11, 0xa8, 0x6f, 0x9e, 0x33, 0x73,
	0xe6, 0x8c, 0xcf, 0x8c, 0x13, 0x60, 0xb2, 0x30, 0x38, 0xd1, 0xd9, 0x94, 0xd7, 0x53, 0x1c, 0x94,
	0x5a, 0x19, 0xc5, 0x76, 0x27, 0x28, 0xdf, 0xcb, 0x62, 0x50, 0xe9, 0xf9, 0x20, 0xd7, 0xb3, 0x41,
	0x5b, 0xd2, 0xdd, 0xcc, 0xd5, 0x6c, 0xa6, 0x0a, 0x57, 0xd6, 0xfd, 0x67, 0xa2, 0x94, 0xa8, 0x46,
	0xc5, 0x89, 0x72, 0x40, 0xff, 0xb3, 0x07, 0xff, 0x8e, 0x96, 0xda, 0xbd, 0x3d, 0x45, 0x8d, 0x6c,
	0x07, 0xc2, 0x32, 0x9b, 0xa0, 0x48, 0xbc, 0x9e, 0x97, 0x86, 0xdc, 0x05, 0xec, 0x26, 0x74, 0xec,
	0xc3, 0xb8, 0x92, 0xe7, 0x98, 0xf8, 0x94, 0xd9, 0xb0, 0xc0, 0x91, 0x3c, 0x47, 0x96, 0xc0, 0x7a,
	0xa5, 0xb4, 0x91, 0xc5, 0x24, 0x09, 0x7a, 0x5e, 0xda, 0xe1, 0x6d, 0xc8, 0xba, 0xb0, 0x71, 0x86,
	0x8b, 0x0f, 0x4a, 0x8b, 0x2a, 0x59, 0xa3, 0xd4, 0x45, 0xcc, 0xb6, 0xc0, 0x97, 0x22, 0x09, 0x7b,
	0x5e, 0x1a, 0x70, 0x5f, 0x0a, 0xb6, 0x0d, 0x81, 0x14, 0x55, 0x12, 0xf5, 0x82, 0x34, 0xe0, 0xf6,
	0xb1, 0xff, 0xd5, 0x87, 0xcd, 0xe5, 0x01, 0x1b, 0x8a, 0x77, 0x41, 0xb9, 0x01, 0x51, 0xa9, 0x64,
	0x61, 0xaa, 0x66, 0xa4, 0x26, 0x62, 0x77, 0x60, 0x53, 0x63, 0x8e, 0x72, 0x8e, 0x63, 0xb3, 0x28,
	0x91, 0xa6, 0x0a, 0x79, 0xdc, 0x60, 0xaf, 0x17, 0x25, 0xb2, 0x3d, 0x00, 0x81, 0xd9, 0x74, 0x9c,
	0xab, 0xba, 0x30, 0x34, 0x5b, 0xc8, 0x3b, 0x16, 0x39, 0xb0, 0x00, 0xbb, 0x0d, 0x31, 0xa5, 0xb3,
	0x19, 0xe5, 0xed, 0x94, 0x3e, 0x27, 0xc6, 0x3e, 0x21, 0x96, 0x4f, 0x7e, 0x3a, 0x81, 0xc8, 0xf1,
	0x09, 0x69, 0xdb, 0xe7, 0x1a, 0x33, 0x83, 0x62, 0x9c, 0x99, 0x64, 0x9d, 0x5e, 0xbd, 0xd3, 0x20,
	0xfb, 0xc4, 0xae, 0x4b, 0xd1, 0xa6, 0x37, 0x5c, 0xba, 0x41, 0xf6, 0x0d, 0x7b, 0x06, 0x21, 0xb5,
	0x4a, 0x3a, 0xbd, 0x20, 0x8d, 0x87, 0xf7, 0x06, 0x97, 0x6c, 0x78, 0xd0, 0xba, 0x73, 0x68, 0xab,
	0xb9, 0x23, 0xf5, 0x3f, 0xf9, 0xf0, 0xf7, 0x0f, 0x89, 0x9f, 0x7c, 0xdb, 0x85, 0x75, 0x69, 0x70,
	0x36, 0x96, 0x82, 0x8c, 0x0b, 0x78, 0x64, 0xc3, 0x91, 0x60, 0xff, 0x43, 0x54, 0x9d, 0xd5, 0x16,
	0x0f, 0x08, 0x0f, 0xab, 0xb3, 0x7a, 0x24, 0xec, 0x4d, 0xcc, 0xb3, 0x69, 0x8d, 0x8d, 0x4f, 0x2e,
	0xb0, 0xa8, 0xce, 0x8c, 0x54, 0xe4, 0x4e, 0xc8, 0x5d, 0xc0, 0x52, 0xd8, 0x6e, 0xc7, 0x1b, 0xeb,
	0x7a, 0x8a, 0xb6, 0x59, 0x44, 0xcd, 0xb6, 0x96, 0x6f, 0x77, 0x24, 0xfe, 0xd0, 0xa3, 0x27, 0xdf,
	0x3d, 0xf2, 0xd2, 0x78, 0xd8, 0xbf, 0xd4, 0xa3, 0xc3, 0xf6, 0xec, 0x5b, 0x7f, 0xbe, 0xf8, 0xb0,
	0xb3, 0x7c, 0x56, 0x1c, 0xab, 0x52, 0x15, 0x15, 0xb2, 0xe7, 0x10, 0x61, 0x61, 0xa4, 0x59, 0x90,
	0x55, 0xf1, 0xf0, 0xee, 0x2f, 0x7d, 0x27, 0x7a, 0x43, 0x62, 0x8f, 0xdd, 0x97, 0xa3, 0xc9, 0xd3,
	0x78, 0x78, 0xeb, 0x52, 0xf6, 0x2b, 0x5b, 0xe5, 0xbe, 0x2c, 0xcd, 0x9e, 0x42, 0x68, 0xcd, 0xaf,
	0x92, 0x80, 0x76, 0xbd, 0xa2, 0xa6, 0xe3, 0x58, 0x49, 0xd4, 0x5a, 0x69, 0x5a, 0xcc, 0x55, 0x92,
	0x2f, 0x6d, 0x15, 0x77, 0xc5, 0xec, 0x21, 0xac, 0xc9, 0xe2, 0xc4, 0xed, 0x2d, 0x1e, 0xee, 0x5d,
	0xa1, 0x78, 0xa2, 0x38, 0x95, 0x0e, 0x3f, 0xae, 0xc1, 0x7f, 0xcb, 0x03, 0x1c, 0xa1, 0x9e, 0xcb,
	0x1c, 0xd9, 0x3b, 0x88, 0x0e, 0x68, 0x63, 0x6c, 0xb5, 0xc1, 0xbb, 0x0f, 0x56, 0x7b, 0xbf, 0x66,
	0x25, 0xfd, 0xbf, 0xac, 0xc2, 0x1b, 0x5a, 0xfa, 0xb5, 0x29, 0x20, 0x44, 0x2f, 0x70, 0x8a, 0x06,
	0xd9, 0xfd, 0x95, 0xa8, 0xf4, 0x3b, 0xf9, 0x5b, 0x32, 0x47, 0x98, 0xe9, 0xfc, 0xf4, 0x7a, 0x65,
	0x8e, 0x21, 0x38, 0x44, 0x73, 0xad, 0x1a, 0xc7, 0x11, 0xfd, 0x81, 0x3c, 0xfa, 0x16, 0x00, 0x00,
	0xff, 0xff, 0xc1, 0x45, 0x4e, 0x97, 0x8e, 0x06, 0x00, 0x00,
}
