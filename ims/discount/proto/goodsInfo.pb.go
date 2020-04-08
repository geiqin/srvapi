// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goodsInfo.proto

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
	SoldNum              int32    `protobuf:"varint,19,opt,name=sold_num,json=soldNum,proto3" json:"sold_num"`
	Disabled             bool     `protobuf:"varint,20,opt,name=disabled,proto3" json:"disabled"`
	SkuNum               int32    `protobuf:"varint,21,opt,name=sku_num,json=skuNum,proto3" json:"sku_num"`
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

func (m *GoodsInfo) GetSoldNum() int32 {
	if m != nil {
		return m.SoldNum
	}
	return 0
}

func (m *GoodsInfo) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *GoodsInfo) GetSkuNum() int32 {
	if m != nil {
		return m.SkuNum
	}
	return 0
}

type GoodsInfoWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top                  int32    `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	PromotionId          int64    `protobuf:"varint,4,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id"`
	PromotionType        string   `protobuf:"bytes,5,opt,name=promotion_type,json=promotionType,proto3" json:"promotion_type"`
	Keywords             string   `protobuf:"bytes,6,opt,name=keywords,proto3" json:"keywords"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsInfoWhere) Reset()         { *m = GoodsInfoWhere{} }
func (m *GoodsInfoWhere) String() string { return proto.CompactTextString(m) }
func (*GoodsInfoWhere) ProtoMessage()    {}
func (*GoodsInfoWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b7d92f8597b270c, []int{1}
}

func (m *GoodsInfoWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsInfoWhere.Unmarshal(m, b)
}
func (m *GoodsInfoWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsInfoWhere.Marshal(b, m, deterministic)
}
func (m *GoodsInfoWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsInfoWhere.Merge(m, src)
}
func (m *GoodsInfoWhere) XXX_Size() int {
	return xxx_messageInfo_GoodsInfoWhere.Size(m)
}
func (m *GoodsInfoWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsInfoWhere.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsInfoWhere proto.InternalMessageInfo

func (m *GoodsInfoWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *GoodsInfoWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GoodsInfoWhere) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *GoodsInfoWhere) GetPromotionId() int64 {
	if m != nil {
		return m.PromotionId
	}
	return 0
}

func (m *GoodsInfoWhere) GetPromotionType() string {
	if m != nil {
		return m.PromotionType
	}
	return ""
}

func (m *GoodsInfoWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

type GoodsInfoResponse struct {
	Entity               *GoodsInfo   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager                *Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items                []*GoodsInfo `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error                *Error       `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info                 *Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GoodsInfoResponse) Reset()         { *m = GoodsInfoResponse{} }
func (m *GoodsInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GoodsInfoResponse) ProtoMessage()    {}
func (*GoodsInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b7d92f8597b270c, []int{2}
}

func (m *GoodsInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsInfoResponse.Unmarshal(m, b)
}
func (m *GoodsInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsInfoResponse.Marshal(b, m, deterministic)
}
func (m *GoodsInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsInfoResponse.Merge(m, src)
}
func (m *GoodsInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GoodsInfoResponse.Size(m)
}
func (m *GoodsInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsInfoResponse proto.InternalMessageInfo

func (m *GoodsInfoResponse) GetEntity() *GoodsInfo {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *GoodsInfoResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *GoodsInfoResponse) GetItems() []*GoodsInfo {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *GoodsInfoResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GoodsInfoResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*GoodsInfo)(nil), "geiqin.srv.ims.discount.GoodsInfo")
	proto.RegisterType((*GoodsInfoWhere)(nil), "geiqin.srv.ims.discount.GoodsInfoWhere")
	proto.RegisterType((*GoodsInfoResponse)(nil), "geiqin.srv.ims.discount.GoodsInfoResponse")
}

func init() { proto.RegisterFile("goodsInfo.proto", fileDescriptor_0b7d92f8597b270c) }

var fileDescriptor_0b7d92f8597b270c = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0x36, 0xd9, 0x24, 0x3b, 0xdb, 0x4f, 0xd3, 0x52, 0x53, 0x54, 0x08, 0x91, 0x90, 0x72,
	0x8a, 0x44, 0xe1, 0x80, 0xb8, 0x23, 0xb4, 0x97, 0xaa, 0xda, 0x82, 0x38, 0x46, 0x9b, 0xd8, 0x4d,
	0xad, 0xc4, 0xf6, 0xd6, 0xf6, 0x52, 0xd2, 0x5f, 0xc4, 0x3f, 0xe0, 0xc0, 0x9f, 0x43, 0x33, 0x6e,
	0xb6, 0xa7, 0x20, 0x6e, 0x7e, 0xef, 0xcd, 0xcb, 0xcc, 0xbc, 0xce, 0x16, 0x0e, 0x16, 0xd6, 0x0a,
	0x5f, 0x98, 0x1b, 0x3b, 0xa9, 0x9d, 0x0d, 0x96, 0x9d, 0x2e, 0xa4, 0xba, 0x53, 0x66, 0xe2, 0xdd,
	0x8f, 0x89, 0xd2, 0x7e, 0x22, 0x94, 0x9f, 0xdb, 0xc6, 0x84, 0xb3, 0xdd, 0xb9, 0xd5, 0xda, 0x9a,
	0x58, 0x36, 0xfa, 0xdd, 0x85, 0xec, 0xcb, 0xc6, 0xca, 0x4e, 0xa1, 0xaf, 0x82, 0xd4, 0x53, 0x25,
	0x78, 0x32, 0x4c, 0xc6, 0x9d, 0xb2, 0x87, 0xb0, 0x10, 0xad, 0xe0, 0x0d, 0xdf, 0x19, 0x26, 0xe3,
	0x2c, 0x0a, 0xd7, 0x86, 0x9d, 0x03, 0x68, 0x2b, 0xe4, 0x6a, 0x1a, 0xd6, 0xb5, 0xe4, 0x1d, 0xd2,
	0x32, 0x62, 0xbe, 0xae, 0x6b, 0xc9, 0x18, 0x74, 0x4d, 0xa5, 0x25, 0xef, 0x92, 0x40, 0x6f, 0xe4,
	0x1a, 0xa3, 0x02, 0x4f, 0x23, 0x87, 0x6f, 0xf6, 0x02, 0x06, 0x33, 0x57, 0x19, 0x81, 0x9d, 0x7b,
	0xc3, 0x64, 0x9c, 0x96, 0x7d, 0xc2, 0x85, 0x60, 0xaf, 0x21, 0x0f, 0xd5, 0x4f, 0x6b, 0xac, 0x5e,
	0xa3, 0xda, 0xa7, 0xb9, 0x60, 0x43, 0x15, 0x82, 0x9d, 0xc1, 0xe0, 0xae, 0xa9, 0x4c, 0x50, 0x61,
	0xcd, 0x07, 0xe4, 0x6d, 0x31, 0xfe, 0x6e, 0xb8, 0x6d, 0xf4, 0x0c, 0x9d, 0x19, 0x39, 0xfb, 0x84,
	0x0b, 0xc1, 0x5e, 0x42, 0x16, 0xa5, 0xc6, 0xad, 0x38, 0xd0, 0x2c, 0xb1, 0xf6, 0x9b, 0x5b, 0x31,
	0x0e, 0xfd, 0x59, 0xe5, 0xe6, 0x56, 0x48, 0x9e, 0x93, 0xb4, 0x81, 0xec, 0x18, 0xd2, 0xda, 0xa9,
	0xb9, 0xe4, 0xbb, 0xc3, 0x64, 0xbc, 0x53, 0x46, 0xc0, 0xde, 0xc0, 0xae, 0x75, 0x6a, 0xa1, 0xcc,
	0x34, 0x8a, 0x7b, 0x24, 0xe6, 0x91, 0xbb, 0xa2, 0x92, 0x73, 0x80, 0xb9, 0xf5, 0xe1, 0xb1, 0x60,
	0x9f, 0x0a, 0x32, 0x64, 0xa2, 0xfc, 0x1c, 0x7a, 0xf7, 0x52, 0x2d, 0x6e, 0x03, 0x3f, 0x20, 0xe9,
	0x11, 0xb1, 0x13, 0xe8, 0xf9, 0x65, 0x83, 0xf3, 0x1f, 0xd2, 0xfc, 0xa9, 0x5f, 0x36, 0x85, 0xd8,
	0xd0, 0xde, 0xf0, 0x23, 0x9a, 0x0f, 0xe9, 0x6b, 0x83, 0xfb, 0x22, 0x4d, 0x99, 0xb3, 0x38, 0xb8,
	0x5f, 0x36, 0x97, 0x18, 0x3b, 0x4a, 0x76, 0x25, 0xa6, 0xa6, 0xd1, 0xfc, 0x59, 0x8c, 0x18, 0xf1,
	0x65, 0xa3, 0x31, 0x41, 0xa1, 0x7c, 0x35, 0x5b, 0x49, 0xc1, 0x8f, 0x87, 0xc9, 0x78, 0x50, 0xb6,
	0x18, 0xff, 0xf2, 0xf4, 0x8b, 0x8d, 0xe6, 0x27, 0xe4, 0xc2, 0xbe, 0x97, 0x8d, 0x1e, 0xfd, 0x49,
	0x60, 0xbf, 0xbd, 0x9c, 0xef, 0xb7, 0xd2, 0xc5, 0x6c, 0xaa, 0x85, 0x8c, 0xc7, 0x93, 0x96, 0x11,
	0x60, 0xd0, 0xf8, 0x98, 0x7a, 0xf5, 0x20, 0xe9, 0x7a, 0xd2, 0x72, 0x80, 0xc4, 0xb5, 0x7a, 0x90,
	0xec, 0x10, 0x3a, 0xc1, 0xd6, 0x74, 0x38, 0x69, 0x89, 0x4f, 0x8c, 0xb2, 0x76, 0x56, 0xdb, 0xa0,
	0xac, 0xc1, 0xb5, 0xbb, 0xb4, 0x76, 0xde, 0x72, 0x85, 0x60, 0x6f, 0x61, 0xff, 0xa9, 0x84, 0x0e,
	0x2f, 0xde, 0xd2, 0x5e, 0xcb, 0xd2, 0xf1, 0x9d, 0xc1, 0x60, 0x29, 0xd7, 0xf7, 0xd6, 0x09, 0x4f,
	0x47, 0x95, 0x95, 0x2d, 0x1e, 0xfd, 0xda, 0x81, 0xa3, 0x76, 0xfa, 0x52, 0xfa, 0xda, 0x1a, 0x2f,
	0xd9, 0x27, 0xe8, 0xc9, 0x78, 0x48, 0xb8, 0x41, 0x7e, 0x31, 0x9a, 0x6c, 0xf9, 0x8a, 0x26, 0x4f,
	0xde, 0x47, 0x07, 0xfb, 0x10, 0x97, 0x77, 0xb4, 0x62, 0x7e, 0xf1, 0x6a, 0xab, 0xf5, 0x0a, 0xab,
	0x62, 0x38, 0x8e, 0x7d, 0x84, 0x14, 0xbf, 0x24, 0xcf, 0x3b, 0xc3, 0xce, 0x7f, 0x36, 0x8c, 0x06,
	0xec, 0x27, 0x9d, 0xb3, 0x8e, 0x02, 0xfa, 0x57, 0xbf, 0xcf, 0x58, 0x55, 0xc6, 0x62, 0xf6, 0x0e,
	0xba, 0xca, 0xdc, 0x58, 0x0a, 0x2c, 0xbf, 0x38, 0xdf, 0x6a, 0xa2, 0x4e, 0x54, 0x3a, 0xeb, 0xd1,
	0x7f, 0x8a, 0xf7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xf7, 0x74, 0x42, 0x63, 0x04, 0x00,
	0x00,
}
