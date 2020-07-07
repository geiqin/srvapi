// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goodsInfo.proto

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
	proto.RegisterType((*GoodsInfo)(nil), "geiqin.srv.crm.integral.GoodsInfo")
}

func init() {
	proto.RegisterFile("goodsInfo.proto", fileDescriptor_0b7d92f8597b270c)
}

var fileDescriptor_0b7d92f8597b270c = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x92, 0xd1, 0x6e, 0x9b, 0x30,
	0x14, 0x86, 0x45, 0x12, 0x20, 0x9c, 0x64, 0xcb, 0x66, 0x6d, 0x8b, 0xb7, 0x29, 0x1a, 0xdb, 0x15,
	0x57, 0xdc, 0xec, 0x21, 0x2a, 0x6e, 0xaa, 0x2a, 0x69, 0xaf, 0x23, 0xc0, 0x2e, 0xb1, 0x02, 0x36,
	0xb1, 0x4d, 0x5b, 0x9e, 0xb8, 0xaf, 0x51, 0xf9, 0x38, 0xe4, 0xce, 0xff, 0xf7, 0x9d, 0x03, 0x3f,
	0xc8, 0xb0, 0x69, 0x94, 0x62, 0xa6, 0x90, 0xcf, 0x2a, 0xef, 0xb5, 0xb2, 0x8a, 0x6c, 0x1b, 0x2e,
	0x2e, 0x42, 0xe6, 0x46, 0xbf, 0xe4, 0xb5, 0xee, 0x72, 0x21, 0x2d, 0x6f, 0x74, 0xd9, 0xfe, 0x7b,
	0x9f, 0x43, 0x72, 0x37, 0x0d, 0x93, 0x2d, 0xc4, 0xc2, 0xf2, 0xee, 0x28, 0x18, 0x0d, 0xd2, 0x20,
	0x9b, 0xef, 0x23, 0x17, 0x0b, 0x76, 0x13, 0x46, 0xd2, 0x59, 0x1a, 0x64, 0x89, 0x17, 0x07, 0x49,
	0x76, 0x00, 0x9d, 0x62, 0xbc, 0x3d, 0xda, 0xb1, 0xe7, 0x74, 0x8e, 0x2e, 0x41, 0xf2, 0x38, 0xf6,
	0x9c, 0x10, 0x58, 0xc8, 0xb2, 0xe3, 0x74, 0x81, 0x02, 0xcf, 0x8e, 0x0d, 0x52, 0x58, 0x1a, 0x7a,
	0xe6, 0xce, 0xe4, 0x27, 0x2c, 0x2b, 0x5d, 0x4a, 0xe6, 0xde, 0x1c, 0xa5, 0x41, 0x16, 0xee, 0x63,
	0xcc, 0x05, 0x23, 0x7f, 0x60, 0x65, 0xcb, 0x37, 0x25, 0x55, 0x37, 0x3a, 0x1b, 0x63, 0x2f, 0x98,
	0x50, 0xc1, 0xc8, 0x2f, 0x58, 0x5e, 0x86, 0x52, 0x5a, 0x61, 0x47, 0xba, 0xc4, 0xdd, 0x5b, 0x76,
	0xcf, 0xb5, 0xa7, 0xa1, 0xab, 0xdc, 0x66, 0x82, 0x9b, 0x31, 0xe6, 0x82, 0x91, 0xdf, 0x90, 0x78,
	0x35, 0xe8, 0x96, 0x02, 0x76, 0xf1, 0xb3, 0x4f, 0xba, 0x25, 0x14, 0xe2, 0xaa, 0xd4, 0xb5, 0x62,
	0x9c, 0xae, 0x50, 0x4d, 0x91, 0x7c, 0x83, 0xb0, 0xd7, 0xa2, 0xe6, 0x74, 0x9d, 0x06, 0xd9, 0x6c,
	0xef, 0x03, 0xf9, 0x0b, 0x6b, 0xa5, 0x45, 0x23, 0xe4, 0xd1, 0xcb, 0x4f, 0x28, 0x57, 0x9e, 0x3d,
	0xe0, 0xc8, 0x0e, 0xa0, 0x56, 0xc6, 0x5e, 0x07, 0x3e, 0xe3, 0x40, 0xe2, 0x88, 0xd7, 0x3f, 0x20,
	0x7a, 0xe5, 0xa2, 0x39, 0x59, 0xba, 0x41, 0x75, 0x4d, 0xe4, 0x3b, 0x44, 0xe6, 0x3c, 0xb8, 0xfe,
	0x5f, 0xb0, 0x7f, 0x68, 0xce, 0x43, 0xc1, 0x26, 0x6c, 0x24, 0xfd, 0x8a, 0xfd, 0x1c, 0x3e, 0x48,
	0xf7, 0xbd, 0x0e, 0xe3, 0x3f, 0x27, 0xbe, 0xb8, 0x39, 0x0f, 0xf7, 0x65, 0xc7, 0xab, 0x08, 0x6f,
	0xc2, 0xff, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x39, 0x36, 0xaf, 0x1c, 0x02, 0x00, 0x00,
}