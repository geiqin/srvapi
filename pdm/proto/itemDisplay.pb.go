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
	Id                    int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                  string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Unit                  string         `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	ItemSn                string         `protobuf:"bytes,4,opt,name=item_sn,json=itemSn,proto3" json:"item_sn,omitempty"`
	TaxonomyId            int64          `protobuf:"varint,5,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id,omitempty"`
	Quantity              int32          `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	SoldNum               int32          `protobuf:"varint,7,opt,name=sold_num,json=soldNum,proto3" json:"sold_num,omitempty"`
	ModelType             string         `protobuf:"bytes,8,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	HideStock             bool           `protobuf:"varint,9,opt,name=hide_stock,json=hideStock,proto3" json:"hide_stock,omitempty"`
	Content               string         `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty"`
	Summary               string         `protobuf:"bytes,11,opt,name=summary,proto3" json:"summary,omitempty"`
	SellingPoint          string         `protobuf:"bytes,12,opt,name=selling_point,json=sellingPoint,proto3" json:"selling_point,omitempty"`
	MinBuy                int32          `protobuf:"varint,13,opt,name=min_buy,json=minBuy,proto3" json:"min_buy,omitempty"`
	NostoreSell           bool           `protobuf:"varint,14,opt,name=nostore_sell,json=nostoreSell,proto3" json:"nostore_sell,omitempty"`
	ThumbUrl              string         `protobuf:"bytes,15,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	TemplateId            int32          `protobuf:"varint,16,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Barcode               string         `protobuf:"bytes,17,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Price                 float32        `protobuf:"fixed32,18,opt,name=price,proto3" json:"price,omitempty"`
	OriginPrice           float32        `protobuf:"fixed32,19,opt,name=origin_price,json=originPrice,proto3" json:"origin_price,omitempty"`
	MinPrice              float32        `protobuf:"fixed32,20,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice              float32        `protobuf:"fixed32,21,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	Weight                float32        `protobuf:"fixed32,22,opt,name=weight,proto3" json:"weight,omitempty"`
	ReviewNum             int32          `protobuf:"varint,23,opt,name=review_num,json=reviewNum,proto3" json:"review_num,omitempty"`
	ViewCount             int32          `protobuf:"varint,24,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	Brand                 *Brand         `protobuf:"bytes,25,opt,name=brand,proto3" json:"brand,omitempty"`
	Cats                  []*Cat         `protobuf:"bytes,26,rep,name=cats,proto3" json:"cats,omitempty" gorm:"many2many:item_cats;"`
	Tags                  []*Tag         `protobuf:"bytes,27,rep,name=tags,proto3" json:"tags,omitempty" gorm:"many2many:item_tags;"`
	Skus                  []*Sku         `protobuf:"bytes,28,rep,name=skus,proto3" json:"skus,omitempty"`
	Galleries             []*ItemGallery `protobuf:"bytes,29,rep,name=galleries,proto3" json:"galleries,omitempty"`
	Specs                 []*Spec        `protobuf:"bytes,30,rep,name=specs,proto3" json:"specs,omitempty" gorm:"many2many:spec_item_indices;"`
	IsMemberGoods         bool           `protobuf:"varint,31,opt,name=is_member_goods,json=isMemberGoods,proto3" json:"is_member_goods,omitempty"`
	IsBuy                 bool           `protobuf:"varint,32,opt,name=is_buy,json=isBuy,proto3" json:"is_buy,omitempty"`
	IsMember              bool           `protobuf:"varint,33,opt,name=is_member,json=isMember,proto3" json:"is_member,omitempty"`
	MemberPrice           float32        `protobuf:"fixed32,34,opt,name=member_price,json=memberPrice,proto3" json:"member_price,omitempty"`
	IsDistribution        bool           `protobuf:"varint,35,opt,name=is_distribution,json=isDistribution,proto3" json:"is_distribution,omitempty"`
	DistributionShareIcon string         `protobuf:"bytes,36,opt,name=distribution_share_icon,json=distributionShareIcon,proto3" json:"distribution_share_icon,omitempty"`
	Listed                bool           `protobuf:"varint,37,opt,name=listed,proto3" json:"listed,omitempty"`
	IsSku                 bool           `protobuf:"varint,38,opt,name=is_sku,json=isSku,proto3" json:"is_sku,omitempty"`
	Food                  *Food          `protobuf:"bytes,39,opt,name=food,proto3" json:"food,omitempty"`
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

func (m *ItemDisplay) GetListed() bool {
	if m != nil {
		return m.Listed
	}
	return false
}

func (m *ItemDisplay) GetIsSku() bool {
	if m != nil {
		return m.IsSku
	}
	return false
}

func (m *ItemDisplay) GetFood() *Food {
	if m != nil {
		return m.Food
	}
	return nil
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

func init() {
	proto.RegisterFile("itemDisplay.proto", fileDescriptor_efb47c94d29c250f)
}

var fileDescriptor_efb47c94d29c250f = []byte{
	// 955 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x51, 0x6f, 0xdb, 0x36,
	0x10, 0x9e, 0xed, 0xd8, 0xb1, 0xce, 0x49, 0x9a, 0x32, 0x49, 0xc3, 0x38, 0xeb, 0xaa, 0x26, 0x5b,
	0x63, 0x6c, 0x80, 0x81, 0xa5, 0xc0, 0x80, 0x3d, 0xae, 0xed, 0x16, 0x78, 0x40, 0x87, 0x40, 0x6e,
	0xb1, 0x47, 0x81, 0x96, 0x18, 0x99, 0xb0, 0x44, 0xaa, 0x24, 0x95, 0xc6, 0x8f, 0xfb, 0xcb, 0x7b,
	0xdd, 0xcb, 0xc0, 0xa3, 0x94, 0x38, 0x8d, 0x57, 0x60, 0x40, 0xdf, 0x78, 0xdf, 0xf7, 0xdd, 0xf1,
	0x44, 0x7e, 0x3c, 0xc1, 0x63, 0x61, 0x79, 0xf1, 0x46, 0x98, 0x32, 0x67, 0xcb, 0x71, 0xa9, 0x95,
	0x55, 0x64, 0x27, 0xe3, 0xe2, 0x83, 0x90, 0x63, 0xa3, 0xaf, 0xc7, 0x65, 0x5a, 0x0c, 0xb7, 0x12,
	0x55, 0x14, 0x4a, 0x7a, 0x76, 0x08, 0x2e, 0xa1, 0x5e, 0x63, 0xf2, 0x05, 0xcb, 0x73, 0xae, 0xeb,
	0xe4, 0x61, 0x60, 0x59, 0x56, 0x2f, 0x07, 0x33, 0xcd, 0x64, 0xda, 0xe0, 0x66, 0x51, 0x35, 0xcb,
	0x84, 0xd9, 0xa6, 0x98, 0x29, 0x79, 0xd2, 0xac, 0xaf, 0x94, 0xaa, 0xd5, 0x27, 0x7f, 0x07, 0x30,
	0x98, 0xdc, 0x35, 0x46, 0x76, 0xa0, 0x2d, 0x52, 0xda, 0x0a, 0x5b, 0xa3, 0x4e, 0xd4, 0x16, 0x29,
	0x21, 0xb0, 0x21, 0x59, 0xc1, 0x69, 0x3b, 0x6c, 0x8d, 0x82, 0x08, 0xd7, 0x0e, 0xab, 0xa4, 0xb0,
	0xb4, 0xe3, 0x31, 0xb7, 0x26, 0x87, 0xb0, 0xe9, 0x5a, 0x8c, 0x8d, 0xa4, 0x1b, 0x08, 0xf7, 0x5c,
	0x38, 0x95, 0xe4, 0x19, 0x0c, 0x2c, 0xbb, 0x51, 0x52, 0x15, 0xcb, 0x58, 0xa4, 0xb4, 0x8b, 0x95,
	0xa1, 0x81, 0x26, 0x29, 0x19, 0x42, 0xff, 0x43, 0xc5, 0xa4, 0x15, 0x76, 0x49, 0x7b, 0x61, 0x6b,
	0xd4, 0x8d, 0x6e, 0x63, 0x72, 0x04, 0x7d, 0xa3, 0xf2, 0x34, 0x96, 0x55, 0x41, 0x37, 0x91, 0xdb,
	0x74, 0xf1, 0x1f, 0x55, 0x41, 0x9e, 0x02, 0x14, 0x2a, 0xe5, 0x79, 0x6c, 0x97, 0x25, 0xa7, 0x7d,
	0xdc, 0x33, 0x40, 0xe4, 0xdd, 0xb2, 0xe4, 0x8e, 0x9e, 0x8b, 0x94, 0xc7, 0xc6, 0xaa, 0x64, 0x41,
	0x83, 0xb0, 0x35, 0xea, 0x47, 0x81, 0x43, 0xa6, 0x0e, 0x20, 0x14, 0x36, 0x13, 0x25, 0x2d, 0x97,
	0x96, 0x02, 0xa6, 0x36, 0xa1, 0x63, 0x4c, 0x55, 0x14, 0x4c, 0x2f, 0xe9, 0xc0, 0x33, 0x75, 0x48,
	0x4e, 0x61, 0xdb, 0xf0, 0x3c, 0x17, 0x32, 0x8b, 0x4b, 0x25, 0xa4, 0xa5, 0x5b, 0xc8, 0x6f, 0xd5,
	0xe0, 0xa5, 0xc3, 0xdc, 0x39, 0x14, 0x42, 0xc6, 0xb3, 0x6a, 0x49, 0xb7, 0xb1, 0xe1, 0x5e, 0x21,
	0xe4, 0xab, 0x6a, 0x49, 0x9e, 0xc3, 0x96, 0x54, 0xc6, 0x2a, 0xcd, 0x63, 0x97, 0x40, 0x77, 0xb0,
	0xa5, 0x41, 0x8d, 0x4d, 0x79, 0x9e, 0x93, 0x63, 0x08, 0xec, 0xbc, 0x2a, 0x66, 0x71, 0xa5, 0x73,
	0xfa, 0x08, 0x8b, 0xf7, 0x11, 0x78, 0xaf, 0x73, 0x3c, 0x47, 0x5e, 0x94, 0x39, 0xb3, 0xdc, 0x9d,
	0xe3, 0x2e, 0x16, 0x87, 0x06, 0x9a, 0xa4, 0xae, 0xf1, 0x19, 0xd3, 0x89, 0x4a, 0x39, 0x7d, 0xec,
	0x1b, 0xaf, 0x43, 0xb2, 0x0f, 0xdd, 0x52, 0x8b, 0x84, 0x53, 0x12, 0xb6, 0x46, 0xed, 0xc8, 0x07,
	0xae, 0x21, 0xa5, 0x45, 0x26, 0x64, 0xec, 0xc9, 0x3d, 0x24, 0x07, 0x1e, 0xbb, 0x44, 0xc9, 0x31,
	0x04, 0xc5, 0x2d, 0xbf, 0x8f, 0x7c, 0xbf, 0x58, 0x25, 0xd9, 0x4d, 0x4d, 0x1e, 0xd4, 0x24, 0xbb,
	0xf1, 0xe4, 0x13, 0xe8, 0x7d, 0xe4, 0x22, 0x9b, 0x5b, 0xfa, 0x04, 0x99, 0x3a, 0x72, 0xd7, 0xa2,
	0xf9, 0xb5, 0xe0, 0x1f, 0xf1, 0x4a, 0x0f, 0xf1, 0x23, 0x02, 0x8f, 0xd4, 0x97, 0x8a, 0x64, 0xa2,
	0x2a, 0x69, 0x29, 0xf5, 0xb4, 0x43, 0x5e, 0x3b, 0x80, 0xfc, 0x00, 0x5d, 0x74, 0x3a, 0x3d, 0x0a,
	0x5b, 0xa3, 0xc1, 0xf9, 0xc1, 0xf8, 0xfe, 0xfb, 0x19, 0xbf, 0x72, 0x64, 0xe4, 0x35, 0xe4, 0x0c,
	0x36, 0x12, 0x66, 0x0d, 0x1d, 0x86, 0x9d, 0xd1, 0xe0, 0x7c, 0xef, 0x53, 0xed, 0x6b, 0x66, 0x23,
	0x14, 0x38, 0xa1, 0x65, 0x99, 0xa1, 0xc7, 0xeb, 0x85, 0xef, 0x58, 0x16, 0xa1, 0xc0, 0x09, 0xcd,
	0xa2, 0x32, 0xf4, 0xeb, 0xf5, 0xc2, 0xe9, 0xa2, 0x8a, 0x50, 0x40, 0x7e, 0x86, 0x20, 0xc3, 0xb7,
	0x2a, 0xb8, 0xa1, 0x4f, 0x51, 0x7d, 0xfc, 0xa9, 0x7a, 0x72, 0xf7, 0xa0, 0xa3, 0x3b, 0x35, 0xf9,
	0x1e, 0xba, 0xee, 0xa5, 0x1a, 0xfa, 0x0d, 0xa6, 0xed, 0x3f, 0xd8, 0xa4, 0xe4, 0x49, 0xe4, 0x25,
	0xe4, 0x05, 0x3c, 0x12, 0x26, 0x2e, 0x78, 0x31, 0xe3, 0x3a, 0xce, 0x94, 0x4a, 0x0d, 0x7d, 0x86,
	0xae, 0xda, 0x16, 0xe6, 0x2d, 0xa2, 0x17, 0x0e, 0x24, 0x07, 0xd0, 0x13, 0x06, 0x2d, 0x19, 0x22,
	0xdd, 0x15, 0xc6, 0x39, 0xf2, 0x18, 0x82, 0xdb, 0x74, 0xfa, 0x1c, 0x99, 0x7e, 0x93, 0xe8, 0xdc,
	0x51, 0x17, 0xf6, 0x17, 0x7c, 0xe2, 0xdd, 0xe1, 0x31, 0x7f, 0xc7, 0x67, 0xb8, 0x7d, 0x2a, 0x8c,
	0xd5, 0x62, 0x56, 0x59, 0xa1, 0x24, 0x3d, 0xc5, 0x2a, 0x3b, 0xc2, 0xbc, 0x59, 0x41, 0xc9, 0x4f,
	0x70, 0xb8, 0xaa, 0x8a, 0xcd, 0x9c, 0x69, 0x1e, 0x8b, 0x44, 0x49, 0xfa, 0x2d, 0x3a, 0xf5, 0x60,
	0x95, 0x9e, 0x3a, 0x76, 0x92, 0x28, 0xe9, 0x4c, 0x94, 0x0b, 0x63, 0x79, 0x4a, 0xbf, 0xc3, 0xba,
	0x75, 0x54, 0x7f, 0x8f, 0x59, 0x54, 0xf4, 0x45, 0xf3, 0x3d, 0xd3, 0x45, 0x45, 0x46, 0xb0, 0xe1,
	0x06, 0x1b, 0x3d, 0x43, 0x73, 0x3c, 0x38, 0xb9, 0xdf, 0x94, 0x4a, 0x23, 0x54, 0x9c, 0x5c, 0xc1,
	0xee, 0xca, 0xcc, 0xfb, 0x73, 0xce, 0x35, 0x7f, 0x30, 0xf8, 0x76, 0xa1, 0x23, 0x52, 0x43, 0xdb,
	0x61, 0x67, 0xd4, 0x89, 0xdc, 0xf2, 0xde, 0x30, 0xea, 0xdc, 0x1f, 0x46, 0x07, 0xd0, 0x33, 0x8b,
	0xca, 0xbd, 0xcb, 0x0d, 0x2c, 0xd0, 0x35, 0x8b, 0x6a, 0x92, 0x9e, 0xfc, 0xd5, 0x86, 0xbd, 0x95,
	0x8d, 0x22, 0x6e, 0x4a, 0x25, 0x0d, 0x27, 0x2f, 0xa1, 0xc7, 0xfd, 0xc0, 0x6b, 0x61, 0xaf, 0x6b,
	0xcd, 0xd1, 0x24, 0xd5, 0x52, 0x67, 0xfe, 0x92, 0x65, 0x5c, 0xe3, 0x28, 0x5e, 0x63, 0xfe, 0x4b,
	0x47, 0x46, 0x5e, 0x43, 0x7e, 0x84, 0xae, 0x9b, 0xbf, 0x86, 0x76, 0xfe, 0xdb, 0x7d, 0xcd, 0x06,
	0x5e, 0xe9, 0xea, 0x73, 0xad, 0x95, 0xc6, 0x4f, 0x58, 0x53, 0xff, 0x57, 0x47, 0x46, 0x5e, 0xe3,
	0xce, 0x5a, 0xc8, 0x2b, 0x85, 0xe3, 0x7c, 0xcd, 0x59, 0x4f, 0xe4, 0x95, 0x8a, 0x50, 0x71, 0xfe,
	0x4f, 0x0b, 0xb6, 0xdf, 0x2e, 0xdd, 0x7e, 0x53, 0xae, 0xaf, 0x9d, 0x6f, 0x7e, 0x87, 0xde, 0x94,
	0x33, 0x9d, 0xcc, 0xc9, 0xd1, 0xba, 0xb6, 0xf0, 0x3a, 0x86, 0xa7, 0x9f, 0xeb, 0xb8, 0x3e, 0xc7,
	0x93, 0xaf, 0xc8, 0x05, 0x74, 0x2e, 0xb8, 0xfd, 0x02, 0x85, 0xde, 0x03, 0xfc, 0x92, 0xa6, 0xd3,
	0xfa, 0x3e, 0xc3, 0xcf, 0x24, 0xfd, 0x9f, 0xb2, 0xb3, 0x1e, 0xfe, 0x65, 0x5f, 0xfe, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xf3, 0xe3, 0xc0, 0x14, 0xfd, 0x07, 0x00, 0x00,
}
