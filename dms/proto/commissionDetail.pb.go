// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commissionDetail.proto

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

type CommissionDetail struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CommissionId         int64    `protobuf:"varint,2,opt,name=commission_id,json=commissionId,proto3" json:"commission_id,omitempty"`
	OrderDetailId        int64    `protobuf:"varint,3,opt,name=order_detail_id,json=orderDetailId,proto3" json:"order_detail_id,omitempty"`
	ItemId               int64    `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,5,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	ThumbUrl             string   `protobuf:"bytes,7,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	Num                  int32    `protobuf:"varint,8,opt,name=num,proto3" json:"num,omitempty"`
	Price                float32  `protobuf:"fixed32,9,opt,name=price,proto3" json:"price,omitempty"`
	CommissionRate       float32  `protobuf:"fixed32,10,opt,name=commission_rate,json=commissionRate,proto3" json:"commission_rate,omitempty"`
	CommissionMoney      float32  `protobuf:"fixed32,11,opt,name=commission_money,json=commissionMoney,proto3" json:"commission_money,omitempty"`
	CreatedAt            string   `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommissionDetail) Reset()         { *m = CommissionDetail{} }
func (m *CommissionDetail) String() string { return proto.CompactTextString(m) }
func (*CommissionDetail) ProtoMessage()    {}
func (*CommissionDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4da00a9d4c67a27, []int{0}
}

func (m *CommissionDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommissionDetail.Unmarshal(m, b)
}
func (m *CommissionDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommissionDetail.Marshal(b, m, deterministic)
}
func (m *CommissionDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommissionDetail.Merge(m, src)
}
func (m *CommissionDetail) XXX_Size() int {
	return xxx_messageInfo_CommissionDetail.Size(m)
}
func (m *CommissionDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_CommissionDetail.DiscardUnknown(m)
}

var xxx_messageInfo_CommissionDetail proto.InternalMessageInfo

func (m *CommissionDetail) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CommissionDetail) GetCommissionId() int64 {
	if m != nil {
		return m.CommissionId
	}
	return 0
}

func (m *CommissionDetail) GetOrderDetailId() int64 {
	if m != nil {
		return m.OrderDetailId
	}
	return 0
}

func (m *CommissionDetail) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *CommissionDetail) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *CommissionDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CommissionDetail) GetThumbUrl() string {
	if m != nil {
		return m.ThumbUrl
	}
	return ""
}

func (m *CommissionDetail) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *CommissionDetail) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CommissionDetail) GetCommissionRate() float32 {
	if m != nil {
		return m.CommissionRate
	}
	return 0
}

func (m *CommissionDetail) GetCommissionMoney() float32 {
	if m != nil {
		return m.CommissionMoney
	}
	return 0
}

func (m *CommissionDetail) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *CommissionDetail) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CommissionDetailResponse struct {
	Entity               *CommissionDetail   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager              `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*CommissionDetail `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error              `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info               `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CommissionDetailResponse) Reset()         { *m = CommissionDetailResponse{} }
func (m *CommissionDetailResponse) String() string { return proto.CompactTextString(m) }
func (*CommissionDetailResponse) ProtoMessage()    {}
func (*CommissionDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4da00a9d4c67a27, []int{1}
}

func (m *CommissionDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommissionDetailResponse.Unmarshal(m, b)
}
func (m *CommissionDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommissionDetailResponse.Marshal(b, m, deterministic)
}
func (m *CommissionDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommissionDetailResponse.Merge(m, src)
}
func (m *CommissionDetailResponse) XXX_Size() int {
	return xxx_messageInfo_CommissionDetailResponse.Size(m)
}
func (m *CommissionDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommissionDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommissionDetailResponse proto.InternalMessageInfo

func (m *CommissionDetailResponse) GetEntity() *CommissionDetail {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *CommissionDetailResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *CommissionDetailResponse) GetItems() []*CommissionDetail {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CommissionDetailResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CommissionDetailResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*CommissionDetail)(nil), "geiqin.srv.dms.CommissionDetail")
	proto.RegisterType((*CommissionDetailResponse)(nil), "geiqin.srv.dms.CommissionDetailResponse")
}

func init() { proto.RegisterFile("commissionDetail.proto", fileDescriptor_d4da00a9d4c67a27) }

var fileDescriptor_d4da00a9d4c67a27 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x5f, 0x6b, 0x13, 0x41,
	0x14, 0xc5, 0x4d, 0x36, 0xbb, 0x6d, 0x6e, 0xfe, 0x34, 0x5c, 0x5a, 0x3b, 0x56, 0x84, 0x25, 0x82,
	0xae, 0x08, 0x79, 0x88, 0x20, 0xbe, 0xd6, 0x56, 0x24, 0x0f, 0xa2, 0x6c, 0xa8, 0x3e, 0x86, 0x6d,
	0xe6, 0xb6, 0x19, 0x9a, 0xdd, 0x89, 0x33, 0xb3, 0x85, 0x7e, 0x05, 0x9f, 0xfc, 0x56, 0x7e, 0x2d,
	0x99, 0xbb, 0x91, 0xc4, 0x55, 0xa8, 0x42, 0xde, 0x76, 0xcf, 0xf9, 0xdd, 0x73, 0x77, 0xe6, 0xc0,
	0xc2, 0xc3, 0xb9, 0xce, 0x73, 0x65, 0xad, 0xd2, 0xc5, 0x39, 0xb9, 0x4c, 0x2d, 0x47, 0x2b, 0xa3,
	0x9d, 0xc6, 0xfe, 0x35, 0xa9, 0xaf, 0xaa, 0x18, 0x59, 0x73, 0x3b, 0x92, 0xb9, 0x3d, 0xe9, 0x7a,
	0x4e, 0x17, 0x95, 0x3b, 0xfc, 0x16, 0xc0, 0xe0, 0xac, 0x36, 0x88, 0x7d, 0x68, 0x2a, 0x29, 0x1a,
	0x71, 0x23, 0x09, 0xd2, 0xa6, 0x92, 0xf8, 0x14, 0x7a, 0x9b, 0xf0, 0x99, 0x92, 0xa2, 0xc9, 0x56,
	0x77, 0x23, 0x4e, 0x24, 0x3e, 0x83, 0x03, 0x6d, 0x24, 0x99, 0x99, 0xe4, 0x10, 0x8f, 0x05, 0x8c,
	0xf5, 0x58, 0xae, 0xa2, 0x27, 0x12, 0x8f, 0x61, 0x4f, 0x39, 0xca, 0xbd, 0xdf, 0x62, 0x3f, 0xf2,
	0xaf, 0x13, 0x89, 0x47, 0x10, 0xd9, 0x9b, 0xd2, 0xeb, 0x21, 0xeb, 0xa1, 0xbd, 0x29, 0x27, 0x12,
	0x11, 0x5a, 0x45, 0x96, 0x93, 0x88, 0xe2, 0x46, 0xd2, 0x4e, 0xf9, 0x19, 0x1f, 0x43, 0xdb, 0x2d,
	0xca, 0xfc, 0x72, 0x56, 0x9a, 0xa5, 0xd8, 0x63, 0x63, 0x9f, 0x85, 0x0b, 0xb3, 0xc4, 0x01, 0x04,
	0x45, 0x99, 0x8b, 0xfd, 0xb8, 0x91, 0x84, 0xa9, 0x7f, 0xc4, 0x43, 0x08, 0x57, 0x46, 0xcd, 0x49,
	0xb4, 0xe3, 0x46, 0xd2, 0x4c, 0xab, 0x17, 0x7c, 0x0e, 0x07, 0x5b, 0xa7, 0x32, 0x99, 0x23, 0x01,
	0xec, 0xf7, 0x37, 0x72, 0x9a, 0x39, 0xc2, 0x17, 0x30, 0xd8, 0x02, 0x73, 0x5d, 0xd0, 0x9d, 0xe8,
	0x30, 0xb9, 0x15, 0xf0, 0xc1, 0xcb, 0xf8, 0x04, 0x60, 0x6e, 0x28, 0x73, 0x24, 0x67, 0x99, 0x13,
	0x5d, 0xfe, 0xb2, 0xf6, 0x5a, 0x39, 0x75, 0xde, 0x2e, 0x57, 0xf2, 0x97, 0xdd, 0xab, 0xec, 0xb5,
	0x72, 0xea, 0x86, 0xdf, 0x9b, 0x20, 0xea, 0x65, 0xa4, 0x64, 0x57, 0xba, 0xb0, 0x84, 0x6f, 0x20,
	0xa2, 0xc2, 0x29, 0x77, 0xc7, 0xc5, 0x74, 0xc6, 0xf1, 0xe8, 0xf7, 0x62, 0x47, 0x7f, 0x4c, 0xae,
	0x79, 0x7c, 0x09, 0xe1, 0x2a, 0xbb, 0x26, 0xc3, 0xb5, 0x75, 0xc6, 0x47, 0xf5, 0xc1, 0x4f, 0xde,
	0x4c, 0x2b, 0x06, 0x5f, 0x43, 0xe8, 0xfb, 0xb0, 0x22, 0x88, 0x83, 0x7f, 0xda, 0x52, 0xe1, 0x7e,
	0x09, 0x19, 0xa3, 0x0d, 0x97, 0xfa, 0x97, 0x25, 0xef, 0xbc, 0x99, 0x56, 0x0c, 0x26, 0xd0, 0x52,
	0xc5, 0x95, 0xe6, 0xa2, 0x3b, 0xe3, 0xc3, 0x3a, 0x3b, 0x29, 0xae, 0x74, 0xca, 0xc4, 0xf8, 0x47,
	0x00, 0xc7, 0xf5, 0x95, 0x53, 0x32, 0xb7, 0xbe, 0xc0, 0xcf, 0x10, 0x9d, 0xf1, 0xd5, 0xe2, 0xbd,
	0x5f, 0x79, 0x92, 0xdc, 0x7b, 0x8e, 0xf5, 0x3d, 0x0f, 0x1f, 0xf8, 0xdc, 0x0b, 0xee, 0x64, 0xf7,
	0xb9, 0xe7, 0xb4, 0xa4, 0x9d, 0xe7, 0x4e, 0x21, 0x78, 0x4f, 0x6e, 0xc7, 0xa1, 0x1f, 0x21, 0x9a,
	0x52, 0x66, 0xe6, 0x0b, 0x7c, 0x54, 0x9f, 0x7a, 0x9b, 0x59, 0xfa, 0xb2, 0x20, 0x43, 0xff, 0x13,
	0x78, 0x19, 0xf1, 0x0f, 0xe7, 0xd5, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xb0, 0xf0, 0x18,
	0xa8, 0x04, 0x00, 0x00,
}