// Code generated by protoc-gen-go. DO NOT EDIT.
// source: purchaseDetail.proto

package geiqin_srv_wms

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

type PurchaseDetail struct {
	Id                   int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PurchaseId           int64      `protobuf:"varint,2,opt,name=purchase_id,json=purchaseId,proto3" json:"purchase_id,omitempty"`
	ItemId               int64      `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64      `protobuf:"varint,4,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Quantity             int32      `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	CostPrice            float32    `protobuf:"fixed32,6,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	CostTotal            float32    `protobuf:"fixed32,7,opt,name=cost_total,json=costTotal,proto3" json:"cost_total,omitempty"`
	Memo                 string     `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo,omitempty"`
	CreatedAt            string     `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string     `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Goods                *GoodsInfo `protobuf:"bytes,11,opt,name=goods,proto3" json:"goods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PurchaseDetail) Reset()         { *m = PurchaseDetail{} }
func (m *PurchaseDetail) String() string { return proto.CompactTextString(m) }
func (*PurchaseDetail) ProtoMessage()    {}
func (*PurchaseDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae65a33796795f26, []int{0}
}

func (m *PurchaseDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseDetail.Unmarshal(m, b)
}
func (m *PurchaseDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseDetail.Marshal(b, m, deterministic)
}
func (m *PurchaseDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseDetail.Merge(m, src)
}
func (m *PurchaseDetail) XXX_Size() int {
	return xxx_messageInfo_PurchaseDetail.Size(m)
}
func (m *PurchaseDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseDetail.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseDetail proto.InternalMessageInfo

func (m *PurchaseDetail) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PurchaseDetail) GetPurchaseId() int64 {
	if m != nil {
		return m.PurchaseId
	}
	return 0
}

func (m *PurchaseDetail) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *PurchaseDetail) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *PurchaseDetail) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *PurchaseDetail) GetCostPrice() float32 {
	if m != nil {
		return m.CostPrice
	}
	return 0
}

func (m *PurchaseDetail) GetCostTotal() float32 {
	if m != nil {
		return m.CostTotal
	}
	return 0
}

func (m *PurchaseDetail) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *PurchaseDetail) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *PurchaseDetail) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *PurchaseDetail) GetGoods() *GoodsInfo {
	if m != nil {
		return m.Goods
	}
	return nil
}

type PurchaseDetailResponse struct {
	Entity               *PurchaseDetail   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager            `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*PurchaseDetail `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error            `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info             `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PurchaseDetailResponse) Reset()         { *m = PurchaseDetailResponse{} }
func (m *PurchaseDetailResponse) String() string { return proto.CompactTextString(m) }
func (*PurchaseDetailResponse) ProtoMessage()    {}
func (*PurchaseDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae65a33796795f26, []int{1}
}

func (m *PurchaseDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseDetailResponse.Unmarshal(m, b)
}
func (m *PurchaseDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseDetailResponse.Marshal(b, m, deterministic)
}
func (m *PurchaseDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseDetailResponse.Merge(m, src)
}
func (m *PurchaseDetailResponse) XXX_Size() int {
	return xxx_messageInfo_PurchaseDetailResponse.Size(m)
}
func (m *PurchaseDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseDetailResponse proto.InternalMessageInfo

func (m *PurchaseDetailResponse) GetEntity() *PurchaseDetail {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *PurchaseDetailResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *PurchaseDetailResponse) GetItems() []*PurchaseDetail {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *PurchaseDetailResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *PurchaseDetailResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*PurchaseDetail)(nil), "geiqin.srv.wms.PurchaseDetail")
	proto.RegisterType((*PurchaseDetailResponse)(nil), "geiqin.srv.wms.PurchaseDetailResponse")
}

func init() { proto.RegisterFile("purchaseDetail.proto", fileDescriptor_ae65a33796795f26) }

var fileDescriptor_ae65a33796795f26 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x6a, 0xdb, 0x30,
	0x18, 0x85, 0xb1, 0x1d, 0x3b, 0xf1, 0xef, 0x91, 0x81, 0x48, 0x36, 0x2d, 0xb0, 0xcd, 0xe4, 0xca,
	0x30, 0xf0, 0x20, 0x1b, 0xbb, 0x0f, 0xac, 0x14, 0xdf, 0x05, 0xd1, 0xfb, 0xe0, 0x5a, 0x4a, 0x2a,
	0x12, 0x5b, 0x8e, 0x24, 0xb7, 0xf4, 0x15, 0xfa, 0x70, 0x7d, 0xa6, 0x22, 0x29, 0x4e, 0x49, 0xda,
	0x8b, 0xde, 0x59, 0xe7, 0x3b, 0xff, 0x7f, 0xf0, 0x91, 0x60, 0xd2, 0x76, 0xb2, 0xba, 0x2b, 0x15,
	0xfb, 0xcf, 0x74, 0xc9, 0xf7, 0x79, 0x2b, 0x85, 0x16, 0x68, 0xbc, 0x65, 0xfc, 0xc0, 0x9b, 0x5c,
	0xc9, 0xfb, 0xfc, 0xa1, 0x56, 0xb3, 0x4f, 0x95, 0xa8, 0x6b, 0xd1, 0x38, 0x3a, 0xfb, 0xbc, 0x15,
	0x82, 0xaa, 0xa2, 0xd9, 0x08, 0x27, 0xcc, 0x9f, 0x7d, 0x18, 0xaf, 0xce, 0xf6, 0xa0, 0x31, 0xf8,
	0x9c, 0x62, 0x2f, 0xf5, 0xb2, 0x80, 0xf8, 0x9c, 0xa2, 0x9f, 0x90, 0xf4, 0x49, 0x6b, 0x4e, 0xb1,
	0x6f, 0x01, 0xf4, 0x52, 0x41, 0xd1, 0x57, 0x18, 0x72, 0xcd, 0x6a, 0x03, 0x03, 0x0b, 0x23, 0x73,
	0x2c, 0x28, 0x9a, 0x42, 0xa4, 0x76, 0x9d, 0xd1, 0x07, 0x56, 0x0f, 0xd5, 0xae, 0x2b, 0x28, 0x9a,
	0xc1, 0xe8, 0xd0, 0x95, 0x8d, 0xe6, 0xfa, 0x11, 0x87, 0xa9, 0x97, 0x85, 0xe4, 0x74, 0x46, 0xdf,
	0x01, 0x2a, 0xa1, 0xf4, 0xba, 0x95, 0xbc, 0x62, 0x38, 0x4a, 0xbd, 0xcc, 0x27, 0xb1, 0x51, 0x56,
	0x46, 0x38, 0x61, 0x2d, 0x74, 0xb9, 0xc7, 0xc3, 0x57, 0x7c, 0x63, 0x04, 0x84, 0x60, 0x50, 0xb3,
	0x5a, 0xe0, 0x51, 0xea, 0x65, 0x31, 0xb1, 0xdf, 0x76, 0x44, 0xb2, 0x52, 0x33, 0xba, 0x2e, 0x35,
	0x8e, 0x2d, 0x89, 0x8f, 0xca, 0x52, 0x1b, 0xdc, 0xb5, 0xb4, 0xc7, 0xe0, 0xf0, 0x51, 0x59, 0x6a,
	0xf4, 0x1b, 0x42, 0x5b, 0x19, 0x4e, 0x52, 0x2f, 0x4b, 0x16, 0xdf, 0xf2, 0xf3, 0x7a, 0xf3, 0xeb,
	0xbe, 0x4f, 0xe2, 0x7c, 0xf3, 0x27, 0x1f, 0xbe, 0x9c, 0x17, 0x4a, 0x98, 0x6a, 0x45, 0xa3, 0x18,
	0xfa, 0x07, 0x11, 0x73, 0x7f, 0xed, 0xd9, 0x65, 0x3f, 0x2e, 0x97, 0x5d, 0xcc, 0x1d, 0xdd, 0xe8,
	0x17, 0x84, 0x6d, 0xb9, 0x65, 0xd2, 0x56, 0x9f, 0x2c, 0xa6, 0x6f, 0xc6, 0x0c, 0x24, 0xce, 0x83,
	0xfe, 0x42, 0x68, 0xda, 0x57, 0x38, 0x48, 0x83, 0x0f, 0x64, 0x38, 0xb3, 0x89, 0x60, 0x52, 0x0a,
	0x69, 0x2f, 0xea, 0x9d, 0x88, 0x2b, 0x03, 0x89, 0xf3, 0xa0, 0x0c, 0x06, 0xbc, 0xd9, 0x08, 0x7b,
	0x77, 0xc9, 0x62, 0x72, 0xe9, 0xb5, 0x6d, 0x58, 0xc7, 0x6d, 0x64, 0x1f, 0xd9, 0x9f, 0x97, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbc, 0x4a, 0xd4, 0x54, 0xab, 0x02, 0x00, 0x00,
}
