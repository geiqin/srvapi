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
	ItemId               int64    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemSn               string   `protobuf:"bytes,2,opt,name=item_sn,json=itemSn,proto3" json:"item_sn,omitempty"`
	ModelType            string   `protobuf:"bytes,3,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Unit                 string   `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	BrandId              int32    `protobuf:"varint,6,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	TaxonomyId           int64    `protobuf:"varint,7,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id,omitempty"`
	Quantity             int32    `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ThumbId              int64    `protobuf:"varint,9,opt,name=thumb_id,json=thumbId,proto3" json:"thumb_id,omitempty"`
	ThumbUrl             string   `protobuf:"bytes,10,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	Barcode              string   `protobuf:"bytes,11,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Price                float32  `protobuf:"fixed32,12,opt,name=price,proto3" json:"price,omitempty"`
	OriginPrice          float32  `protobuf:"fixed32,13,opt,name=origin_price,json=originPrice,proto3" json:"origin_price,omitempty"`
	CostPrice            float32  `protobuf:"fixed32,14,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	Weight               float32  `protobuf:"fixed32,15,opt,name=weight,proto3" json:"weight,omitempty"`
	SkuId                int64    `protobuf:"varint,16,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	SkuSn                string   `protobuf:"bytes,17,opt,name=sku_sn,json=skuSn,proto3" json:"sku_sn,omitempty"`
	SkuName              string   `protobuf:"bytes,18,opt,name=sku_name,json=skuName,proto3" json:"sku_name,omitempty"`
	SkuSpecDesc          string   `protobuf:"bytes,19,opt,name=sku_spec_desc,json=skuSpecDesc,proto3" json:"sku_spec_desc,omitempty"`
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

func (m *GoodsInfo) GetSkuSpecDesc() string {
	if m != nil {
		return m.SkuSpecDesc
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
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x92, 0xd1, 0xce, 0x93, 0x30,
	0x1c, 0xc5, 0xc3, 0xf7, 0x0d, 0x18, 0x7f, 0x36, 0xa7, 0x55, 0x67, 0xd5, 0x2c, 0xe2, 0xae, 0xb8,
	0xe2, 0xc6, 0x57, 0x30, 0x31, 0xdc, 0x18, 0xb3, 0xe9, 0x35, 0x01, 0x5a, 0x59, 0xb3, 0xd1, 0xb2,
	0xb6, 0x4c, 0x79, 0x06, 0x5f, 0xda, 0xf4, 0xdf, 0xb1, 0xbb, 0x9e, 0xf3, 0x3b, 0x07, 0x0e, 0x0d,
	0xb0, 0xe9, 0x94, 0x62, 0xa6, 0x94, 0xbf, 0x55, 0x31, 0x68, 0x65, 0x15, 0xd9, 0x76, 0x5c, 0x5c,
	0x85, 0x2c, 0x8c, 0xbe, 0x15, 0x4a, 0xb3, 0x62, 0xd0, 0xe2, 0x56, 0x5b, 0xbe, 0xff, 0xb7, 0x80,
	0xe4, 0xdb, 0x9c, 0x25, 0xef, 0x20, 0x16, 0x96, 0xf7, 0x95, 0x60, 0x34, 0xc8, 0x82, 0xfc, 0xf9,
	0x10, 0x39, 0x59, 0xb2, 0x07, 0x30, 0x92, 0x3e, 0x65, 0x41, 0x9e, 0x78, 0x70, 0x94, 0x64, 0x07,
	0xd0, 0x2b, 0xc6, 0x2f, 0x95, 0x9d, 0x06, 0x4e, 0x9f, 0x91, 0x25, 0xe8, 0xfc, 0x9c, 0x06, 0x4e,
	0x08, 0x2c, 0x64, 0xdd, 0x73, 0xba, 0x40, 0x80, 0x67, 0xe7, 0x8d, 0x52, 0x58, 0x1a, 0x7a, 0xcf,
	0x9d, 0xc9, 0x7b, 0x58, 0x36, 0xba, 0x96, 0xcc, 0xbd, 0x39, 0xca, 0x82, 0x3c, 0x3c, 0xc4, 0xa8,
	0x4b, 0x46, 0x3e, 0x41, 0x6a, 0xeb, 0xbf, 0x4a, 0xaa, 0x7e, 0x72, 0x34, 0xc6, 0x5d, 0x30, 0x5b,
	0x25, 0x23, 0x1f, 0x60, 0x79, 0x1d, 0x6b, 0x69, 0x85, 0x9d, 0xe8, 0x12, 0xbb, 0x0f, 0xed, 0x9e,
	0x6b, 0x4f, 0x63, 0xdf, 0xb8, 0x66, 0x82, 0xcd, 0x18, 0x75, 0xc9, 0xc8, 0x47, 0x48, 0x3c, 0x1a,
	0xf5, 0x85, 0x02, 0x6e, 0xf1, 0xd9, 0x5f, 0xfa, 0x42, 0x28, 0xc4, 0x4d, 0xad, 0x5b, 0xc5, 0x38,
	0x4d, 0x11, 0xcd, 0x92, 0xbc, 0x81, 0x70, 0xd0, 0xa2, 0xe5, 0x74, 0x95, 0x05, 0xf9, 0xd3, 0xc1,
	0x0b, 0xf2, 0x19, 0x56, 0x4a, 0x8b, 0x4e, 0xc8, 0xca, 0xc3, 0x35, 0xc2, 0xd4, 0x7b, 0x3f, 0x30,
	0xb2, 0x03, 0x68, 0x95, 0xb1, 0xf7, 0xc0, 0x0b, 0x0c, 0x24, 0xce, 0xf1, 0x78, 0x0b, 0xd1, 0x1f,
	0x2e, 0xba, 0x93, 0xa5, 0x1b, 0x44, 0x77, 0x45, 0xde, 0x42, 0x64, 0xce, 0xa3, 0xdb, 0xff, 0x12,
	0xf7, 0x87, 0xe6, 0x3c, 0x96, 0x6c, 0xb6, 0x8d, 0xa4, 0xaf, 0x70, 0x9f, 0xb3, 0x8f, 0xd2, 0x7d,
	0xaf, 0xb3, 0xf1, 0xce, 0x89, 0x1f, 0x6e, 0xce, 0xe3, 0x77, 0x77, 0xed, 0x7b, 0x58, 0x63, 0x63,
	0xe0, 0x6d, 0xc5, 0xb8, 0x69, 0xe9, 0x6b, 0xe4, 0xa9, 0x2b, 0x0e, 0xbc, 0xfd, 0xca, 0x4d, 0xdb,
	0x44, 0xf8, 0xb3, 0x7c, 0xf9, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x28, 0x99, 0xa3, 0x3f, 0x02,
	0x00, 0x00,
}
