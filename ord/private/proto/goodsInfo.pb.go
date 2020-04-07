// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goodsInfo.proto

package geiqin_srv_ord_private

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

type GoodsInfo struct {
	ItemId               int64    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	ItemSn               string   `protobuf:"bytes,2,opt,name=item_sn,json=itemSn,proto3" json:"item_sn"`
	ModelType            string   `protobuf:"bytes,3,opt,name=model_type,json=modelType,proto3" json:"model_type"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	Unit                 string   `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit"`
	BrandId              int32    `protobuf:"varint,6,opt,name=brand_id,json=brandId,proto3" json:"brand_id"`
	TaxonomyId           int64    `protobuf:"varint,7,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id"`
	Quantity             int32    `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity"`
	ThumbId              int64    `protobuf:"varint,9,opt,name=thumb_id,json=thumbId,proto3" json:"thumb_id"`
	ThumbUrl             string   `protobuf:"bytes,10,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url"`
	Barcode              string   `protobuf:"bytes,11,opt,name=barcode,proto3" json:"barcode"`
	Price                float32  `protobuf:"fixed32,12,opt,name=price,proto3" json:"price"`
	OriginPrice          float32  `protobuf:"fixed32,13,opt,name=origin_price,json=originPrice,proto3" json:"origin_price"`
	CostPrice            float32  `protobuf:"fixed32,14,opt,name=cost_price,json=costPrice,proto3" json:"cost_price"`
	Weight               float32  `protobuf:"fixed32,15,opt,name=weight,proto3" json:"weight"`
	SkuId                int64    `protobuf:"varint,16,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	SkuSn                string   `protobuf:"bytes,17,opt,name=sku_sn,json=skuSn,proto3" json:"sku_sn"`
	SkuName              string   `protobuf:"bytes,18,opt,name=sku_name,json=skuName,proto3" json:"sku_name"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsInfo) Reset()         { *m = GoodsInfo{} }
func (m *GoodsInfo) String() string { return proto.CompactTextString(m) }
func (*GoodsInfo) ProtoMessage()    {}
func (*GoodsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b7d92f8597b270c, []int{0}
}

func (m *GoodsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsInfo.Unmarshal(m, b)
}
func (m *GoodsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsInfo.Marshal(b, m, deterministic)
}
func (m *GoodsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsInfo.Merge(m, src)
}
func (m *GoodsInfo) XXX_Size() int {
	return xxx_messageInfo_GoodsInfo.Size(m)
}
func (m *GoodsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsInfo proto.InternalMessageInfo

func (m *GoodsInfo) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *GoodsInfo) GetItemSn() string {
	if m != nil {
		return m.ItemSn
	}
	return ""
}

func (m *GoodsInfo) GetModelType() string {
	if m != nil {
		return m.ModelType
	}
	return ""
}

func (m *GoodsInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GoodsInfo) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *GoodsInfo) GetBrandId() int32 {
	if m != nil {
		return m.BrandId
	}
	return 0
}

func (m *GoodsInfo) GetTaxonomyId() int64 {
	if m != nil {
		return m.TaxonomyId
	}
	return 0
}

func (m *GoodsInfo) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *GoodsInfo) GetThumbId() int64 {
	if m != nil {
		return m.ThumbId
	}
	return 0
}

func (m *GoodsInfo) GetThumbUrl() string {
	if m != nil {
		return m.ThumbUrl
	}
	return ""
}

func (m *GoodsInfo) GetBarcode() string {
	if m != nil {
		return m.Barcode
	}
	return ""
}

func (m *GoodsInfo) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *GoodsInfo) GetOriginPrice() float32 {
	if m != nil {
		return m.OriginPrice
	}
	return 0
}

func (m *GoodsInfo) GetCostPrice() float32 {
	if m != nil {
		return m.CostPrice
	}
	return 0
}

func (m *GoodsInfo) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *GoodsInfo) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *GoodsInfo) GetSkuSn() string {
	if m != nil {
		return m.SkuSn
	}
	return ""
}

func (m *GoodsInfo) GetSkuName() string {
	if m != nil {
		return m.SkuName
	}
	return ""
}

func init() {
	proto.RegisterType((*GoodsInfo)(nil), "geiqin.srv.ord.private.GoodsInfo")
}

func init() {
	proto.RegisterFile("goodsInfo.proto", fileDescriptor_0b7d92f8597b270c)
}

var fileDescriptor_0b7d92f8597b270c = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x92, 0xc1, 0x6e, 0xa3, 0x30,
	0x10, 0x86, 0x45, 0x12, 0x20, 0x4c, 0xb2, 0x9b, 0x5d, 0x6b, 0x37, 0xeb, 0x6d, 0x15, 0x95, 0xf6,
	0xc4, 0x89, 0x4b, 0x1f, 0xa2, 0xe2, 0x52, 0x55, 0x49, 0x7b, 0x8e, 0x20, 0x76, 0x89, 0x95, 0x60,
	0x13, 0xdb, 0xa4, 0xe5, 0x89, 0xfb, 0x1a, 0x95, 0xc7, 0x21, 0x37, 0xff, 0xdf, 0x37, 0x03, 0x3f,
	0xc8, 0xb0, 0xa8, 0x95, 0x62, 0xa6, 0x90, 0xef, 0x2a, 0x6f, 0xb5, 0xb2, 0x8a, 0x2c, 0x6b, 0x2e,
	0x4e, 0x42, 0xe6, 0x46, 0x9f, 0x73, 0xa5, 0x59, 0xde, 0x6a, 0x71, 0x2e, 0x2d, 0x7f, 0xf8, 0x1a,
	0x43, 0xf2, 0x34, 0xcc, 0x92, 0x7f, 0x10, 0x0b, 0xcb, 0x9b, 0xad, 0x60, 0x34, 0x48, 0x83, 0x6c,
	0xbc, 0x8e, 0x5c, 0x2c, 0xd8, 0x55, 0x18, 0x49, 0x47, 0x69, 0x90, 0x25, 0x5e, 0x6c, 0x24, 0x59,
	0x01, 0x34, 0x8a, 0xf1, 0xe3, 0xd6, 0xf6, 0x2d, 0xa7, 0x63, 0x74, 0x09, 0x92, 0xd7, 0xbe, 0xe5,
	0x84, 0xc0, 0x44, 0x96, 0x0d, 0xa7, 0x13, 0x14, 0x78, 0x76, 0xac, 0x93, 0xc2, 0xd2, 0xd0, 0x33,
	0x77, 0x26, 0xff, 0x61, 0x5a, 0xe9, 0x52, 0x32, 0xf7, 0xe6, 0x28, 0x0d, 0xb2, 0x70, 0x1d, 0x63,
	0x2e, 0x18, 0xb9, 0x83, 0x99, 0x2d, 0x3f, 0x95, 0x54, 0x4d, 0xef, 0x6c, 0x8c, 0xbd, 0x60, 0x40,
	0x05, 0x23, 0x37, 0x30, 0x3d, 0x75, 0xa5, 0xb4, 0xc2, 0xf6, 0x74, 0x8a, 0xbb, 0xd7, 0xec, 0x9e,
	0x6b, 0xf7, 0x5d, 0x53, 0xb9, 0xcd, 0x04, 0x37, 0x63, 0xcc, 0x05, 0x23, 0xb7, 0x90, 0x78, 0xd5,
	0xe9, 0x23, 0x05, 0xec, 0xe2, 0x67, 0xdf, 0xf4, 0x91, 0x50, 0x88, 0xab, 0x52, 0xef, 0x14, 0xe3,
	0x74, 0x86, 0x6a, 0x88, 0xe4, 0x0f, 0x84, 0xad, 0x16, 0x3b, 0x4e, 0xe7, 0x69, 0x90, 0x8d, 0xd6,
	0x3e, 0x90, 0x7b, 0x98, 0x2b, 0x2d, 0x6a, 0x21, 0xb7, 0x5e, 0xfe, 0x40, 0x39, 0xf3, 0xec, 0x05,
	0x47, 0x56, 0x00, 0x3b, 0x65, 0xec, 0x65, 0xe0, 0x27, 0x0e, 0x24, 0x8e, 0x78, 0xbd, 0x84, 0xe8,
	0x83, 0x8b, 0x7a, 0x6f, 0xe9, 0x02, 0xd5, 0x25, 0x91, 0xbf, 0x10, 0x99, 0x43, 0xe7, 0xfa, 0xff,
	0xc2, 0xfe, 0xa1, 0x39, 0x74, 0x05, 0x1b, 0xb0, 0x91, 0xf4, 0x37, 0xf6, 0x73, 0x78, 0x23, 0xdd,
	0xf7, 0x3a, 0x8c, 0xff, 0x9c, 0xf8, 0xe2, 0xe6, 0xd0, 0x3d, 0x97, 0x0d, 0xaf, 0x22, 0xbc, 0x08,
	0x8f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xac, 0xb4, 0x08, 0x67, 0x1b, 0x02, 0x00, 0x00,
}
