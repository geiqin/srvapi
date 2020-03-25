// Code generated by protoc-gen-go. DO NOT EDIT.
// source: itemDisplay.proto

package geiqin_srv_pdm

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

type ItemDisplay struct {
	Id                    int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name                  string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Unit                  string         `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit"`
	ItemSn                string         `protobuf:"bytes,4,opt,name=item_sn,json=itemSn,proto3" json:"item_sn"`
	TaxonomyId            int64          `protobuf:"varint,5,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id"`
	Quantity              int32          `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity"`
	SoldNum               int32          `protobuf:"varint,7,opt,name=sold_num,json=soldNum,proto3" json:"sold_num"`
	ModelType             string         `protobuf:"bytes,8,opt,name=model_type,json=modelType,proto3" json:"model_type"`
	HideStock             bool           `protobuf:"varint,9,opt,name=hide_stock,json=hideStock,proto3" json:"hide_stock"`
	Content               string         `protobuf:"bytes,10,opt,name=content,proto3" json:"content"`
	Summary               string         `protobuf:"bytes,11,opt,name=summary,proto3" json:"summary"`
	SellingPoint          string         `protobuf:"bytes,12,opt,name=selling_point,json=sellingPoint,proto3" json:"selling_point"`
	MinBuy                int32          `protobuf:"varint,13,opt,name=min_buy,json=minBuy,proto3" json:"min_buy"`
	NostoreSell           bool           `protobuf:"varint,14,opt,name=nostore_sell,json=nostoreSell,proto3" json:"nostore_sell"`
	ThumbUrl              string         `protobuf:"bytes,15,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url"`
	TemplateId            int32          `protobuf:"varint,16,opt,name=template_id,json=templateId,proto3" json:"template_id"`
	Barcode               string         `protobuf:"bytes,17,opt,name=barcode,proto3" json:"barcode"`
	Price                 float32        `protobuf:"fixed32,18,opt,name=price,proto3" json:"price"`
	OriginPrice           float32        `protobuf:"fixed32,19,opt,name=origin_price,json=originPrice,proto3" json:"origin_price"`
	MinPrice              float32        `protobuf:"fixed32,20,opt,name=min_price,json=minPrice,proto3" json:"min_price"`
	MaxPrice              float32        `protobuf:"fixed32,21,opt,name=max_price,json=maxPrice,proto3" json:"max_price"`
	Weight                float32        `protobuf:"fixed32,22,opt,name=weight,proto3" json:"weight"`
	ReviewNum             int32          `protobuf:"varint,23,opt,name=review_num,json=reviewNum,proto3" json:"review_num"`
	ViewCount             int32          `protobuf:"varint,24,opt,name=view_count,json=viewCount,proto3" json:"view_count"`
	Brand                 *Brand         `protobuf:"bytes,25,opt,name=brand,proto3" json:"brand,omitempty"`
	Cats                  []*Cat         `protobuf:"bytes,26,rep,name=cats,proto3" json:"cats,omitempty" gorm:"many2many:item_cats;`
	Tags                  []*Tag         `protobuf:"bytes,27,rep,name=tags,proto3" json:"tags,omitempty" gorm:"many2many:item_tags;"`
	Skus                  []*Sku         `protobuf:"bytes,28,rep,name=skus,proto3" json:"skus,omitempty"`
	Galleries             []*ItemGallery `protobuf:"bytes,29,rep,name=galleries,proto3" json:"galleries,omitempty"`
	Specs                 []*Spec        `protobuf:"bytes,30,rep,name=specs,proto3" json:"specs,omitempty" gorm:"many2many:spec_item_indices;"`
	IsMemberGoods         bool           `protobuf:"varint,31,opt,name=is_member_goods,json=isMemberGoods,proto3" json:"is_member_goods"`
	IsBuy                 bool           `protobuf:"varint,32,opt,name=is_buy,json=isBuy,proto3" json:"is_buy"`
	IsMember              bool           `protobuf:"varint,33,opt,name=is_member,json=isMember,proto3" json:"is_member"`
	MemberPrice           float32        `protobuf:"fixed32,34,opt,name=member_price,json=memberPrice,proto3" json:"member_price"`
	IsDistribution        bool           `protobuf:"varint,35,opt,name=is_distribution,json=isDistribution,proto3" json:"is_distribution"`
	DistributionShareIcon string         `protobuf:"bytes,36,opt,name=distribution_share_icon,json=distributionShareIcon,proto3" json:"distribution_share_icon"`
	XXX_NoUnkeyedLiteral  struct{}       `json:"-"`
	XXX_unrecognized      []byte         `json:"-"`
	XXX_sizecache         int32          `json:"-"`
}

func (m *ItemDisplay) Reset()         { *m = ItemDisplay{} }
func (m *ItemDisplay) String() string { return proto.CompactTextString(m) }
func (*ItemDisplay) ProtoMessage()    {}
func (*ItemDisplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_efb47c94d29c250f, []int{0}
}

func (m *ItemDisplay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemDisplay.Unmarshal(m, b)
}
func (m *ItemDisplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemDisplay.Marshal(b, m, deterministic)
}
func (m *ItemDisplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemDisplay.Merge(m, src)
}
func (m *ItemDisplay) XXX_Size() int {
	return xxx_messageInfo_ItemDisplay.Size(m)
}
func (m *ItemDisplay) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemDisplay.DiscardUnknown(m)
}

var xxx_messageInfo_ItemDisplay proto.InternalMessageInfo

func (m *ItemDisplay) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemDisplay) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ItemDisplay) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *ItemDisplay) GetItemSn() string {
	if m != nil {
		return m.ItemSn
	}
	return ""
}

func (m *ItemDisplay) GetTaxonomyId() int64 {
	if m != nil {
		return m.TaxonomyId
	}
	return 0
}

func (m *ItemDisplay) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *ItemDisplay) GetSoldNum() int32 {
	if m != nil {
		return m.SoldNum
	}
	return 0
}

func (m *ItemDisplay) GetModelType() string {
	if m != nil {
		return m.ModelType
	}
	return ""
}

func (m *ItemDisplay) GetHideStock() bool {
	if m != nil {
		return m.HideStock
	}
	return false
}

func (m *ItemDisplay) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ItemDisplay) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *ItemDisplay) GetSellingPoint() string {
	if m != nil {
		return m.SellingPoint
	}
	return ""
}

func (m *ItemDisplay) GetMinBuy() int32 {
	if m != nil {
		return m.MinBuy
	}
	return 0
}

func (m *ItemDisplay) GetNostoreSell() bool {
	if m != nil {
		return m.NostoreSell
	}
	return false
}

func (m *ItemDisplay) GetThumbUrl() string {
	if m != nil {
		return m.ThumbUrl
	}
	return ""
}

func (m *ItemDisplay) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func (m *ItemDisplay) GetBarcode() string {
	if m != nil {
		return m.Barcode
	}
	return ""
}

func (m *ItemDisplay) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ItemDisplay) GetOriginPrice() float32 {
	if m != nil {
		return m.OriginPrice
	}
	return 0
}

func (m *ItemDisplay) GetMinPrice() float32 {
	if m != nil {
		return m.MinPrice
	}
	return 0
}

func (m *ItemDisplay) GetMaxPrice() float32 {
	if m != nil {
		return m.MaxPrice
	}
	return 0
}

func (m *ItemDisplay) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *ItemDisplay) GetReviewNum() int32 {
	if m != nil {
		return m.ReviewNum
	}
	return 0
}

func (m *ItemDisplay) GetViewCount() int32 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *ItemDisplay) GetBrand() *Brand {
	if m != nil {
		return m.Brand
	}
	return nil
}

func (m *ItemDisplay) GetCats() []*Cat {
	if m != nil {
		return m.Cats
	}
	return nil
}

func (m *ItemDisplay) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ItemDisplay) GetSkus() []*Sku {
	if m != nil {
		return m.Skus
	}
	return nil
}

func (m *ItemDisplay) GetGalleries() []*ItemGallery {
	if m != nil {
		return m.Galleries
	}
	return nil
}

func (m *ItemDisplay) GetSpecs() []*Spec {
	if m != nil {
		return m.Specs
	}
	return nil
}

func (m *ItemDisplay) GetIsMemberGoods() bool {
	if m != nil {
		return m.IsMemberGoods
	}
	return false
}

func (m *ItemDisplay) GetIsBuy() bool {
	if m != nil {
		return m.IsBuy
	}
	return false
}

func (m *ItemDisplay) GetIsMember() bool {
	if m != nil {
		return m.IsMember
	}
	return false
}

func (m *ItemDisplay) GetMemberPrice() float32 {
	if m != nil {
		return m.MemberPrice
	}
	return 0
}

func (m *ItemDisplay) GetIsDistribution() bool {
	if m != nil {
		return m.IsDistribution
	}
	return false
}

func (m *ItemDisplay) GetDistributionShareIcon() string {
	if m != nil {
		return m.DistributionShareIcon
	}
	return ""
}

type ItemDisplayWhere struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	SoldNum              int32    `protobuf:"varint,3,opt,name=sold_num,json=soldNum,proto3" json:"sold_num,omitempty"`
	SkuId                int64    `protobuf:"varint,4,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemDisplayWhere) Reset()         { *m = ItemDisplayWhere{} }
func (m *ItemDisplayWhere) String() string { return proto.CompactTextString(m) }
func (*ItemDisplayWhere) ProtoMessage()    {}
func (*ItemDisplayWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_efb47c94d29c250f, []int{1}
}

func (m *ItemDisplayWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemDisplayWhere.Unmarshal(m, b)
}
func (m *ItemDisplayWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemDisplayWhere.Marshal(b, m, deterministic)
}
func (m *ItemDisplayWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemDisplayWhere.Merge(m, src)
}
func (m *ItemDisplayWhere) XXX_Size() int {
	return xxx_messageInfo_ItemDisplayWhere.Size(m)
}
func (m *ItemDisplayWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemDisplayWhere.DiscardUnknown(m)
}

var xxx_messageInfo_ItemDisplayWhere proto.InternalMessageInfo

func (m *ItemDisplayWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemDisplayWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ItemDisplayWhere) GetSoldNum() int32 {
	if m != nil {
		return m.SoldNum
	}
	return 0
}

func (m *ItemDisplayWhere) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

type ItemDisplayResponse struct {
	Entity               *ItemDisplay   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager         `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*ItemDisplay `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error         `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info          `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ItemDisplayResponse) Reset()         { *m = ItemDisplayResponse{} }
func (m *ItemDisplayResponse) String() string { return proto.CompactTextString(m) }
func (*ItemDisplayResponse) ProtoMessage()    {}
func (*ItemDisplayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_efb47c94d29c250f, []int{2}
}

func (m *ItemDisplayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemDisplayResponse.Unmarshal(m, b)
}
func (m *ItemDisplayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemDisplayResponse.Marshal(b, m, deterministic)
}
func (m *ItemDisplayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemDisplayResponse.Merge(m, src)
}
func (m *ItemDisplayResponse) XXX_Size() int {
	return xxx_messageInfo_ItemDisplayResponse.Size(m)
}
func (m *ItemDisplayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemDisplayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemDisplayResponse proto.InternalMessageInfo

func (m *ItemDisplayResponse) GetEntity() *ItemDisplay {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ItemDisplayResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ItemDisplayResponse) GetItems() []*ItemDisplay {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ItemDisplayResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ItemDisplayResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemDisplay)(nil), "geiqin.srv.pdm.ItemDisplay")
	proto.RegisterType((*ItemDisplayWhere)(nil), "geiqin.srv.pdm.ItemDisplayWhere")
	proto.RegisterType((*ItemDisplayResponse)(nil), "geiqin.srv.pdm.ItemDisplayResponse")
}

func init() { proto.RegisterFile("itemDisplay.proto", fileDescriptor_efb47c94d29c250f) }

var fileDescriptor_efb47c94d29c250f = []byte{
	// 913 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdf, 0x6f, 0xdb, 0x36,
	0x10, 0x9e, 0x7f, 0xc6, 0x3a, 0x27, 0x69, 0xca, 0xc4, 0x0d, 0x63, 0xaf, 0xab, 0xea, 0x0c, 0xab,
	0xb1, 0x01, 0x06, 0x96, 0x02, 0x03, 0xf6, 0xb8, 0xb6, 0x43, 0xe0, 0x01, 0x1d, 0x02, 0xb9, 0xc5,
	0x1e, 0x05, 0x5a, 0x62, 0x64, 0xc2, 0x12, 0xa9, 0x92, 0x54, 0x1a, 0x3f, 0xee, 0xcf, 0xd8, 0xdf,
	0xba, 0x97, 0x81, 0x47, 0x39, 0x71, 0x5a, 0xaf, 0xc0, 0x80, 0xbd, 0xdd, 0x7d, 0xdf, 0x77, 0xc7,
	0x13, 0xef, 0x78, 0x82, 0xc7, 0xc2, 0xf2, 0xe2, 0x8d, 0x30, 0x65, 0xce, 0xd6, 0xd3, 0x52, 0x2b,
	0xab, 0xc8, 0x61, 0xc6, 0xc5, 0x07, 0x21, 0xa7, 0x46, 0xdf, 0x4c, 0xcb, 0xb4, 0x18, 0xee, 0x27,
	0xaa, 0x28, 0x94, 0xf4, 0xec, 0x10, 0x5c, 0x40, 0x6d, 0x63, 0xf0, 0x25, 0xcb, 0x73, 0xae, 0xeb,
	0xe0, 0x61, 0x60, 0x59, 0x56, 0x9b, 0xfd, 0x85, 0x66, 0x32, 0xdd, 0xe0, 0x66, 0x55, 0x6d, 0xcc,
	0x84, 0xd9, 0x4d, 0x32, 0x53, 0xf2, 0xc4, 0xdb, 0xe3, 0xbf, 0x02, 0xe8, 0xcf, 0xee, 0x8b, 0x21,
	0x87, 0xd0, 0x14, 0x29, 0x6d, 0x84, 0x8d, 0x49, 0x2b, 0x6a, 0x8a, 0x94, 0x10, 0x68, 0x4b, 0x56,
	0x70, 0xda, 0x0c, 0x1b, 0x93, 0x20, 0x42, 0xdb, 0x61, 0x95, 0x14, 0x96, 0xb6, 0x3c, 0xe6, 0x6c,
	0x72, 0x0a, 0x7b, 0xae, 0xac, 0xd8, 0x48, 0xda, 0x46, 0xb8, 0xeb, 0xdc, 0xb9, 0x24, 0xcf, 0xa0,
	0x6f, 0xd9, 0xad, 0x92, 0xaa, 0x58, 0xc7, 0x22, 0xa5, 0x1d, 0xcc, 0x0c, 0x1b, 0x68, 0x96, 0x92,
	0x21, 0xf4, 0x3e, 0x54, 0x4c, 0x5a, 0x61, 0xd7, 0xb4, 0x1b, 0x36, 0x26, 0x9d, 0xe8, 0xce, 0x27,
	0x67, 0xd0, 0x33, 0x2a, 0x4f, 0x63, 0x59, 0x15, 0x74, 0x0f, 0xb9, 0x3d, 0xe7, 0xff, 0x5e, 0x15,
	0xe4, 0x29, 0x40, 0xa1, 0x52, 0x9e, 0xc7, 0x76, 0x5d, 0x72, 0xda, 0xc3, 0x33, 0x03, 0x44, 0xde,
	0xad, 0x4b, 0xee, 0xe8, 0xa5, 0x48, 0x79, 0x6c, 0xac, 0x4a, 0x56, 0x34, 0x08, 0x1b, 0x93, 0x5e,
	0x14, 0x38, 0x64, 0xee, 0x00, 0x42, 0x61, 0x2f, 0x51, 0xd2, 0x72, 0x69, 0x29, 0x60, 0xe8, 0xc6,
	0x75, 0x8c, 0xa9, 0x8a, 0x82, 0xe9, 0x35, 0xed, 0x7b, 0xa6, 0x76, 0xc9, 0x39, 0x1c, 0x18, 0x9e,
	0xe7, 0x42, 0x66, 0x71, 0xa9, 0x84, 0xb4, 0x74, 0x1f, 0xf9, 0xfd, 0x1a, 0xbc, 0x72, 0x98, 0xbb,
	0x87, 0x42, 0xc8, 0x78, 0x51, 0xad, 0xe9, 0x01, 0x16, 0xdc, 0x2d, 0x84, 0x7c, 0x55, 0xad, 0xc9,
	0x73, 0xd8, 0x97, 0xca, 0x58, 0xa5, 0x79, 0xec, 0x02, 0xe8, 0x21, 0x96, 0xd4, 0xaf, 0xb1, 0x39,
	0xcf, 0x73, 0x32, 0x82, 0xc0, 0x2e, 0xab, 0x62, 0x11, 0x57, 0x3a, 0xa7, 0x8f, 0x30, 0x79, 0x0f,
	0x81, 0xf7, 0x3a, 0xc7, 0x7b, 0xe4, 0x45, 0x99, 0x33, 0xcb, 0xdd, 0x3d, 0x1e, 0x61, 0x72, 0xd8,
	0x40, 0xb3, 0xd4, 0x15, 0xbe, 0x60, 0x3a, 0x51, 0x29, 0xa7, 0x8f, 0x7d, 0xe1, 0xb5, 0x4b, 0x4e,
	0xa0, 0x53, 0x6a, 0x91, 0x70, 0x4a, 0xc2, 0xc6, 0xa4, 0x19, 0x79, 0xc7, 0x15, 0xa4, 0xb4, 0xc8,
	0x84, 0x8c, 0x3d, 0x79, 0x8c, 0x64, 0xdf, 0x63, 0x57, 0x28, 0x19, 0x41, 0x50, 0xdc, 0xf1, 0x27,
	0xc8, 0xf7, 0x8a, 0x6d, 0x92, 0xdd, 0xd6, 0xe4, 0xa0, 0x26, 0xd9, 0xad, 0x27, 0x9f, 0x40, 0xf7,
	0x23, 0x17, 0xd9, 0xd2, 0xd2, 0x27, 0xc8, 0xd4, 0x9e, 0x6b, 0x8b, 0xe6, 0x37, 0x82, 0x7f, 0xc4,
	0x96, 0x9e, 0xe2, 0x47, 0x04, 0x1e, 0xa9, 0x9b, 0x8a, 0x64, 0xa2, 0x2a, 0x69, 0x29, 0xf5, 0xb4,
	0x43, 0x5e, 0x3b, 0x80, 0xfc, 0x00, 0x1d, 0x9c, 0x6e, 0x7a, 0x16, 0x36, 0x26, 0xfd, 0x8b, 0xc1,
	0xf4, 0xe1, 0x9b, 0x99, 0xbe, 0x72, 0x64, 0xe4, 0x35, 0xe4, 0x05, 0xb4, 0x13, 0x66, 0x0d, 0x1d,
	0x86, 0xad, 0x49, 0xff, 0xe2, 0xf8, 0x53, 0xed, 0x6b, 0x66, 0x23, 0x14, 0x38, 0xa1, 0x65, 0x99,
	0xa1, 0xa3, 0xdd, 0xc2, 0x77, 0x2c, 0x8b, 0x50, 0xe0, 0x84, 0x66, 0x55, 0x19, 0xfa, 0xf5, 0x6e,
	0xe1, 0x7c, 0x55, 0x45, 0x28, 0x20, 0x3f, 0x43, 0x90, 0xe1, 0xfb, 0x14, 0xdc, 0xd0, 0xa7, 0xa8,
	0x1e, 0x7d, 0xaa, 0x9e, 0xdd, 0x3f, 0xe2, 0xe8, 0x5e, 0x4d, 0xbe, 0x87, 0x8e, 0x7b, 0x9d, 0x86,
	0x7e, 0x83, 0x61, 0x27, 0x9f, 0x1d, 0x52, 0xf2, 0x24, 0xf2, 0x12, 0xf2, 0x1d, 0x3c, 0x12, 0x26,
	0x2e, 0x78, 0xb1, 0xe0, 0x3a, 0xce, 0x94, 0x4a, 0x0d, 0x7d, 0x86, 0x53, 0x75, 0x20, 0xcc, 0x5b,
	0x44, 0x2f, 0x1d, 0x48, 0x06, 0xd0, 0x15, 0x06, 0x47, 0x32, 0x44, 0xba, 0x23, 0x8c, 0x9b, 0xc8,
	0x11, 0x04, 0x77, 0xe1, 0xf4, 0x39, 0x32, 0xbd, 0x4d, 0xa0, 0x9b, 0x8e, 0x3a, 0xb1, 0x6f, 0xf0,
	0xd8, 0x4f, 0x87, 0xc7, 0x7c, 0x8f, 0x5f, 0xe0, 0xf1, 0xa9, 0x30, 0x56, 0x8b, 0x45, 0x65, 0x85,
	0x92, 0xf4, 0x1c, 0xb3, 0x1c, 0x0a, 0xf3, 0x66, 0x0b, 0x25, 0x3f, 0xc1, 0xe9, 0xb6, 0x2a, 0x36,
	0x4b, 0xa6, 0x79, 0x2c, 0x12, 0x25, 0xe9, 0xb7, 0x38, 0xa9, 0x83, 0x6d, 0x7a, 0xee, 0xd8, 0x59,
	0xa2, 0xe4, 0xf8, 0x1a, 0x8e, 0xb6, 0x56, 0xd3, 0x1f, 0x4b, 0xae, 0xf9, 0x67, 0xfb, 0xe9, 0x08,
	0x5a, 0x22, 0x35, 0xb4, 0x19, 0xb6, 0x26, 0xad, 0xc8, 0x99, 0x0f, 0x76, 0x46, 0xeb, 0xe1, 0xce,
	0x18, 0x40, 0xd7, 0xac, 0x2a, 0xf7, 0x7c, 0xda, 0x98, 0xa0, 0x63, 0x56, 0xd5, 0x2c, 0x1d, 0xff,
	0xd9, 0x84, 0xe3, 0xad, 0x83, 0x22, 0x6e, 0x4a, 0x25, 0x0d, 0x27, 0x2f, 0xa1, 0xcb, 0xfd, 0x5e,
	0x6a, 0xe0, 0xbc, 0xed, 0xec, 0xe1, 0x26, 0xa8, 0x96, 0xba, 0x19, 0x2d, 0x59, 0xc6, 0x35, 0x6e,
	0xcc, 0x1d, 0x33, 0x7a, 0xe5, 0xc8, 0xc8, 0x6b, 0xc8, 0x8f, 0xd0, 0x71, 0x6b, 0xd2, 0xd0, 0xd6,
	0xbf, 0x0f, 0xc9, 0xe6, 0x00, 0xaf, 0x74, 0xf9, 0xb9, 0xd6, 0x4a, 0xe3, 0x27, 0xec, 0xc8, 0xff,
	0xab, 0x23, 0x23, 0xaf, 0x21, 0x13, 0x68, 0x0b, 0x79, 0xad, 0x70, 0xeb, 0xee, 0x18, 0xa6, 0x99,
	0xbc, 0x56, 0x11, 0x2a, 0x2e, 0xfe, 0x6e, 0xc0, 0xc1, 0xdb, 0xb5, 0x3b, 0x6f, 0xce, 0xf5, 0x8d,
	0x6b, 0xef, 0x6f, 0xd0, 0x9d, 0x73, 0xa6, 0x93, 0x25, 0x39, 0xdb, 0x55, 0x16, 0xb6, 0x63, 0x78,
	0xfe, 0xa5, 0x8a, 0xeb, 0x7b, 0x1c, 0x7f, 0x45, 0x2e, 0xa1, 0x75, 0xc9, 0xed, 0xff, 0x90, 0xe8,
	0x3d, 0xc0, 0x2f, 0x69, 0x3a, 0xaf, 0xfb, 0x19, 0x7e, 0x21, 0xe8, 0xbf, 0xa4, 0x5d, 0x74, 0xf1,
	0x67, 0xf8, 0xf2, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x19, 0xe5, 0xe0, 0x98, 0x07, 0x00,
	0x00,
}