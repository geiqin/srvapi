// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

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

type Item struct {
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Unit                 string         `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	ItemSn               string         `protobuf:"bytes,4,opt,name=item_sn,json=itemSn,proto3" json:"item_sn,omitempty"`
	BrandId              int64          `protobuf:"varint,5,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	TaxonomyId           int64          `protobuf:"varint,6,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id,omitempty"`
	BuyQuota             int32          `protobuf:"varint,7,opt,name=buy_quota,json=buyQuota,proto3" json:"buy_quota,omitempty"`
	Quantity             int32          `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
	SoldNum              int32          `protobuf:"varint,9,opt,name=sold_num,json=soldNum,proto3" json:"sold_num,omitempty"`
	InitSoldNum          int32          `protobuf:"varint,10,opt,name=init_sold_num,json=initSoldNum,proto3" json:"init_sold_num,omitempty"`
	FrozenNum            int32          `protobuf:"varint,11,opt,name=frozen_num,json=frozenNum,proto3" json:"frozen_num,omitempty"`
	ModelType            string         `protobuf:"bytes,12,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	Listed               bool           `protobuf:"varint,13,opt,name=listed,proto3" json:"listed,omitempty"`
	Locked               bool           `protobuf:"varint,14,opt,name=locked,proto3" json:"locked,omitempty"`
	HideStock            bool           `protobuf:"varint,15,opt,name=hide_stock,json=hideStock,proto3" json:"hide_stock,omitempty"`
	JoinLevelDiscount    bool           `protobuf:"varint,16,opt,name=join_level_discount,json=joinLevelDiscount,proto3" json:"join_level_discount,omitempty"`
	PurchaseRight        bool           `protobuf:"varint,17,opt,name=purchase_right,json=purchaseRight,proto3" json:"purchase_right,omitempty"`
	AutoListingTime      string         `protobuf:"bytes,18,opt,name=auto_listing_time,json=autoListingTime,proto3" json:"auto_listing_time,omitempty"`
	PostType             string         `protobuf:"bytes,19,opt,name=post_type,json=postType,proto3" json:"post_type,omitempty"`
	PostFee              float32        `protobuf:"fixed32,20,opt,name=post_fee,json=postFee,proto3" json:"post_fee,omitempty"`
	OutItemNo            string         `protobuf:"bytes,21,opt,name=out_item_no,json=outItemNo,proto3" json:"out_item_no,omitempty"`
	Content              string         `protobuf:"bytes,22,opt,name=content,proto3" json:"content,omitempty"`
	Summary              string         `protobuf:"bytes,23,opt,name=summary,proto3" json:"summary,omitempty"`
	SellingPoint32       string         `protobuf:"bytes,24,opt,name=selling_point32,json=sellingPoint32,proto3" json:"selling_point32,omitempty"`
	MinBuy               int32          `protobuf:"varint,25,opt,name=min_buy,json=minBuy,proto3" json:"min_buy,omitempty"`
	NostoreSell          bool           `protobuf:"varint,26,opt,name=nostore_sell,json=nostoreSell,proto3" json:"nostore_sell,omitempty"`
	ThumbId              int64          `protobuf:"varint,27,opt,name=thumb_id,json=thumbId,proto3" json:"thumb_id,omitempty"`
	ThumbUrl             string         `protobuf:"bytes,28,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	TemplateId           int32          `protobuf:"varint,29,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Barcode              string         `protobuf:"bytes,30,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Price                float32        `protobuf:"fixed32,31,opt,name=price,proto3" json:"price,omitempty"`
	OriginPrice          float32        `protobuf:"fixed32,32,opt,name=origin_price,json=originPrice,proto3" json:"origin_price,omitempty"`
	CostPrice            float32        `protobuf:"fixed32,33,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	MinPrice             float32        `protobuf:"fixed32,34,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice             float32        `protobuf:"fixed32,35,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	Weight               float32        `protobuf:"fixed32,36,opt,name=weight,proto3" json:"weight,omitempty"`
	SpecData             string         `protobuf:"bytes,37,opt,name=spec_data,json=specData,proto3" json:"spec_data,omitempty"`
	ReviewNum            int32          `protobuf:"varint,38,opt,name=review_num,json=reviewNum,proto3" json:"review_num,omitempty"`
	ViewCount            int32          `protobuf:"varint,39,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	ViewWCount           int32          `protobuf:"varint,40,opt,name=view_w_count,json=viewWCount,proto3" json:"view_w_count,omitempty"`
	BuyCount             int32          `protobuf:"varint,41,opt,name=buy_count,json=buyCount,proto3" json:"buy_count,omitempty"`
	BuyWCount            int32          `protobuf:"varint,42,opt,name=buy_w_count,json=buyWCount,proto3" json:"buy_w_count,omitempty"`
	Sorting              int32          `protobuf:"varint,43,opt,name=sorting,proto3" json:"sorting,omitempty"`
	CreatedAt            string         `protobuf:"bytes,44,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string         `protobuf:"bytes,45,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Brand                *Brand         `protobuf:"bytes,46,opt,name=brand,proto3" json:"brand,omitempty"`
	Cats                 []*Cat         `protobuf:"bytes,47,rep,name=cats,proto3" json:"cats,omitempty" gorm:"many2many:item_cats;"`
	Tags                 []*Tag         `protobuf:"bytes,48,rep,name=tags,proto3" json:"tags,omitempty" gorm:"many2many:item_tags;"`
	Skus                 []*Sku         `protobuf:"bytes,49,rep,name=skus,proto3" json:"skus,omitempty"`
	Galleries            []*ItemGallery `protobuf:"bytes,50,rep,name=galleries,proto3" json:"galleries,omitempty"`
	Prices               []*ItemPrice   `protobuf:"bytes,51,rep,name=prices,proto3" json:"prices,omitempty"`
	Rights               []*ItemRight   `protobuf:"bytes,52,rep,name=rights,proto3" json:"rights,omitempty"`
	Ids                  []int64        `protobuf:"varint,53,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Item) GetItemSn() string {
	if m != nil {
		return m.ItemSn
	}
	return ""
}

func (m *Item) GetBrandId() int64 {
	if m != nil {
		return m.BrandId
	}
	return 0
}

func (m *Item) GetTaxonomyId() int64 {
	if m != nil {
		return m.TaxonomyId
	}
	return 0
}

func (m *Item) GetBuyQuota() int32 {
	if m != nil {
		return m.BuyQuota
	}
	return 0
}

func (m *Item) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *Item) GetSoldNum() int32 {
	if m != nil {
		return m.SoldNum
	}
	return 0
}

func (m *Item) GetInitSoldNum() int32 {
	if m != nil {
		return m.InitSoldNum
	}
	return 0
}

func (m *Item) GetFrozenNum() int32 {
	if m != nil {
		return m.FrozenNum
	}
	return 0
}

func (m *Item) GetModelType() string {
	if m != nil {
		return m.ModelType
	}
	return ""
}

func (m *Item) GetListed() bool {
	if m != nil {
		return m.Listed
	}
	return false
}

func (m *Item) GetLocked() bool {
	if m != nil {
		return m.Locked
	}
	return false
}

func (m *Item) GetHideStock() bool {
	if m != nil {
		return m.HideStock
	}
	return false
}

func (m *Item) GetJoinLevelDiscount() bool {
	if m != nil {
		return m.JoinLevelDiscount
	}
	return false
}

func (m *Item) GetPurchaseRight() bool {
	if m != nil {
		return m.PurchaseRight
	}
	return false
}

func (m *Item) GetAutoListingTime() string {
	if m != nil {
		return m.AutoListingTime
	}
	return ""
}

func (m *Item) GetPostType() string {
	if m != nil {
		return m.PostType
	}
	return ""
}

func (m *Item) GetPostFee() float32 {
	if m != nil {
		return m.PostFee
	}
	return 0
}

func (m *Item) GetOutItemNo() string {
	if m != nil {
		return m.OutItemNo
	}
	return ""
}

func (m *Item) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Item) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Item) GetSellingPoint32() string {
	if m != nil {
		return m.SellingPoint32
	}
	return ""
}

func (m *Item) GetMinBuy() int32 {
	if m != nil {
		return m.MinBuy
	}
	return 0
}

func (m *Item) GetNostoreSell() bool {
	if m != nil {
		return m.NostoreSell
	}
	return false
}

func (m *Item) GetThumbId() int64 {
	if m != nil {
		return m.ThumbId
	}
	return 0
}

func (m *Item) GetThumbUrl() string {
	if m != nil {
		return m.ThumbUrl
	}
	return ""
}

func (m *Item) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func (m *Item) GetBarcode() string {
	if m != nil {
		return m.Barcode
	}
	return ""
}

func (m *Item) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetOriginPrice() float32 {
	if m != nil {
		return m.OriginPrice
	}
	return 0
}

func (m *Item) GetCostPrice() float32 {
	if m != nil {
		return m.CostPrice
	}
	return 0
}

func (m *Item) GetMinPrice() float32 {
	if m != nil {
		return m.MinPrice
	}
	return 0
}

func (m *Item) GetMaxPrice() float32 {
	if m != nil {
		return m.MaxPrice
	}
	return 0
}

func (m *Item) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Item) GetSpecData() string {
	if m != nil {
		return m.SpecData
	}
	return ""
}

func (m *Item) GetReviewNum() int32 {
	if m != nil {
		return m.ReviewNum
	}
	return 0
}

func (m *Item) GetViewCount() int32 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *Item) GetViewWCount() int32 {
	if m != nil {
		return m.ViewWCount
	}
	return 0
}

func (m *Item) GetBuyCount() int32 {
	if m != nil {
		return m.BuyCount
	}
	return 0
}

func (m *Item) GetBuyWCount() int32 {
	if m != nil {
		return m.BuyWCount
	}
	return 0
}

func (m *Item) GetSorting() int32 {
	if m != nil {
		return m.Sorting
	}
	return 0
}

func (m *Item) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Item) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Item) GetBrand() *Brand {
	if m != nil {
		return m.Brand
	}
	return nil
}

func (m *Item) GetCats() []*Cat {
	if m != nil {
		return m.Cats
	}
	return nil
}

func (m *Item) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Item) GetSkus() []*Sku {
	if m != nil {
		return m.Skus
	}
	return nil
}

func (m *Item) GetGalleries() []*ItemGallery {
	if m != nil {
		return m.Galleries
	}
	return nil
}

func (m *Item) GetPrices() []*ItemPrice {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *Item) GetRights() []*ItemRight {
	if m != nil {
		return m.Rights
	}
	return nil
}

func (m *Item) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ItemResponse struct {
	Entity               *Item    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Item  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemResponse) Reset()         { *m = ItemResponse{} }
func (m *ItemResponse) String() string { return proto.CompactTextString(m) }
func (*ItemResponse) ProtoMessage()    {}
func (*ItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{1}
}

func (m *ItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemResponse.Unmarshal(m, b)
}
func (m *ItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemResponse.Marshal(b, m, deterministic)
}
func (m *ItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemResponse.Merge(m, src)
}
func (m *ItemResponse) XXX_Size() int {
	return xxx_messageInfo_ItemResponse.Size(m)
}
func (m *ItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemResponse proto.InternalMessageInfo

func (m *ItemResponse) GetEntity() *Item {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ItemResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ItemResponse) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ItemResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ItemResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "geiqin.srv.pdm.Item")
	proto.RegisterType((*ItemResponse)(nil), "geiqin.srv.pdm.ItemResponse")
}

func init() { proto.RegisterFile("item.proto", fileDescriptor_6007f868cf6553df) }

var fileDescriptor_6007f868cf6553df = []byte{
	// 1158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5d, 0x73, 0x1b, 0x35,
	0x17, 0x7e, 0x1d, 0x27, 0xfe, 0x90, 0xf3, 0xd1, 0xa8, 0x5f, 0x4a, 0xd2, 0x0f, 0x37, 0x2f, 0xa5,
	0xa6, 0x2d, 0x86, 0xa6, 0x70, 0xc1, 0xc0, 0x00, 0x6d, 0x02, 0x9d, 0x30, 0x9d, 0x4e, 0x59, 0xb7,
	0xd3, 0xcb, 0x1d, 0x79, 0xf7, 0xc4, 0x16, 0xd9, 0x95, 0xb6, 0x92, 0x36, 0xad, 0xf9, 0x21, 0x5c,
	0xf3, 0x03, 0xf9, 0x11, 0xcc, 0x39, 0x92, 0x53, 0x28, 0x66, 0x98, 0x71, 0xef, 0x74, 0x9e, 0xe7,
	0x39, 0xc7, 0xd2, 0xd9, 0x47, 0x47, 0x66, 0x4c, 0x79, 0x28, 0x87, 0x95, 0x35, 0xde, 0xf0, 0xcd,
	0x09, 0xa8, 0xd7, 0x4a, 0x0f, 0x9d, 0x3d, 0x1b, 0x56, 0x79, 0xb9, 0xbb, 0x9e, 0x99, 0xb2, 0x34,
	0x3a, 0xb0, 0xbb, 0xdb, 0xa8, 0x7c, 0x22, 0x8b, 0x02, 0xec, 0x2c, 0x42, 0x5b, 0x08, 0x3d, 0xb7,
	0x2a, 0x83, 0xbf, 0x02, 0x89, 0x9a, 0x4c, 0x7d, 0x04, 0xba, 0x5e, 0x4e, 0xe2, 0xb2, 0x37, 0xb6,
	0x52, 0xe7, 0x73, 0xdc, 0x9d, 0xd6, 0xf3, 0x65, 0x26, 0xa3, 0x7a, 0xff, 0xb7, 0x0d, 0xb6, 0x7a,
	0xec, 0xa1, 0xe4, 0x9b, 0x6c, 0x45, 0xe5, 0xa2, 0xd1, 0x6f, 0x0c, 0x9a, 0xc9, 0x8a, 0xca, 0x39,
	0x67, 0xab, 0x5a, 0x96, 0x20, 0x56, 0xfa, 0x8d, 0x41, 0x37, 0xa1, 0x35, 0x62, 0xb5, 0x56, 0x5e,
	0x34, 0x03, 0x86, 0x6b, 0x7e, 0x95, 0xb5, 0x71, 0x07, 0xa9, 0xd3, 0x62, 0x95, 0xe0, 0x16, 0x86,
	0x23, 0xcd, 0x77, 0x58, 0x87, 0x7e, 0x3e, 0x55, 0xb9, 0x58, 0xa3, 0xb2, 0x6d, 0x8a, 0x8f, 0x73,
	0x7e, 0x93, 0xf5, 0xbc, 0x7c, 0x6b, 0xb4, 0x29, 0x67, 0xc8, 0xb6, 0x88, 0x65, 0x73, 0xe8, 0x38,
	0xe7, 0x7b, 0xac, 0x3b, 0xae, 0x67, 0xe9, 0xeb, 0xda, 0x78, 0x29, 0xda, 0xfd, 0xc6, 0x60, 0x2d,
	0xe9, 0x8c, 0xeb, 0xd9, 0xcf, 0x18, 0xf3, 0x5d, 0xd6, 0x79, 0x5d, 0x4b, 0xed, 0x95, 0x9f, 0x89,
	0x4e, 0xe0, 0xe6, 0x31, 0xfe, 0xa8, 0x33, 0x45, 0x9e, 0xea, 0xba, 0x14, 0x5d, 0xe2, 0xda, 0x18,
	0x3f, 0xab, 0x4b, 0xbe, 0xcf, 0x36, 0x94, 0x56, 0x3e, 0x3d, 0xe7, 0x19, 0xf1, 0x3d, 0x04, 0x47,
	0x51, 0x73, 0x9d, 0xb1, 0x13, 0x6b, 0x7e, 0x05, 0x4d, 0x82, 0x1e, 0x09, 0xba, 0x01, 0x89, 0x74,
	0x69, 0x72, 0x28, 0x52, 0x3f, 0xab, 0x40, 0xac, 0xd3, 0x71, 0xbb, 0x84, 0xbc, 0x98, 0x55, 0xc0,
	0xaf, 0xb0, 0x56, 0xa1, 0x9c, 0x87, 0x5c, 0x6c, 0xf4, 0x1b, 0x83, 0x4e, 0x12, 0x23, 0xc2, 0x4d,
	0x76, 0x0a, 0xb9, 0xd8, 0x8c, 0x38, 0x45, 0x58, 0x6e, 0xaa, 0x72, 0x48, 0x9d, 0x37, 0xd9, 0xa9,
	0xd8, 0x22, 0xae, 0x8b, 0xc8, 0x08, 0x01, 0x3e, 0x64, 0x17, 0x7f, 0x31, 0x4a, 0xa7, 0x05, 0x9c,
	0x41, 0x91, 0xe6, 0xca, 0x65, 0xa6, 0xd6, 0x5e, 0x5c, 0x20, 0xdd, 0x36, 0x52, 0x4f, 0x91, 0x39,
	0x8a, 0x04, 0xbf, 0xcd, 0x36, 0xab, 0xda, 0x66, 0x53, 0xe9, 0x20, 0xb5, 0x68, 0x08, 0xb1, 0x4d,
	0xd2, 0x8d, 0x39, 0x4a, 0x2e, 0xe1, 0x77, 0xd9, 0xb6, 0xac, 0xbd, 0x49, 0x71, 0x73, 0x4a, 0x4f,
	0x52, 0xaf, 0x4a, 0x10, 0x9c, 0xce, 0xb2, 0x85, 0xc4, 0xd3, 0x80, 0xbf, 0x50, 0x25, 0xe0, 0x77,
	0xa8, 0x8c, 0xf3, 0xe1, 0xbc, 0x17, 0x49, 0xd3, 0x41, 0x80, 0x8e, 0xbb, 0xc3, 0x68, 0x9d, 0x9e,
	0x00, 0x88, 0x4b, 0xfd, 0xc6, 0x60, 0x25, 0x69, 0x63, 0xfc, 0x23, 0x00, 0xbf, 0xc1, 0x7a, 0xa6,
	0xf6, 0x29, 0x19, 0x43, 0x1b, 0x71, 0x39, 0x74, 0xca, 0xd4, 0x1e, 0xad, 0xf6, 0xcc, 0x70, 0xc1,
	0xda, 0x99, 0xd1, 0x1e, 0xb4, 0x17, 0x57, 0x88, 0x9b, 0x87, 0xc8, 0xb8, 0xba, 0x2c, 0xa5, 0x9d,
	0x89, 0xab, 0x81, 0x89, 0x21, 0xbf, 0xc3, 0xb6, 0x1c, 0x14, 0x05, 0x6e, 0xb9, 0x32, 0x4a, 0xfb,
	0x87, 0x07, 0x42, 0x90, 0x62, 0x33, 0xc2, 0xcf, 0x03, 0x8a, 0x8e, 0x2c, 0x95, 0x4e, 0xc7, 0xf5,
	0x4c, 0xec, 0xd0, 0x17, 0x6c, 0x95, 0x4a, 0x3f, 0xae, 0x67, 0xfc, 0x16, 0x5b, 0xd7, 0xc6, 0x79,
	0x63, 0x21, 0xc5, 0x14, 0xb1, 0x4b, 0xed, 0xe9, 0x45, 0x6c, 0x04, 0x45, 0x81, 0x67, 0xf2, 0xd3,
	0xba, 0x1c, 0xa3, 0x2d, 0xf7, 0x82, 0x69, 0x29, 0x0e, 0x9e, 0x0c, 0x54, 0x6d, 0x0b, 0x71, 0x2d,
	0xf4, 0x82, 0x80, 0x97, 0xb6, 0x20, 0x47, 0x43, 0x59, 0x15, 0xd2, 0x03, 0xa6, 0x5e, 0xa7, 0xdf,
	0x65, 0x73, 0xe8, 0x38, 0xc7, 0x73, 0x8d, 0xa5, 0xcd, 0x4c, 0x0e, 0xe2, 0x46, 0x38, 0x57, 0x0c,
	0xf9, 0x25, 0xb6, 0x56, 0xe1, 0x7d, 0x16, 0x37, 0xa9, 0x87, 0x21, 0xc0, 0xbd, 0x1a, 0xab, 0x26,
	0x4a, 0xa7, 0x81, 0xec, 0x13, 0xd9, 0x0b, 0x18, 0xdd, 0x7f, 0xb4, 0x4f, 0x86, 0xfd, 0x0f, 0x82,
	0x5b, 0x24, 0xe8, 0x22, 0x12, 0xe8, 0x3d, 0xd6, 0x2d, 0xcf, 0xd3, 0xf7, 0x89, 0xed, 0x94, 0xf3,
	0x5c, 0x24, 0xe5, 0xdb, 0x48, 0xfe, 0x3f, 0x92, 0xf2, 0x6d, 0x20, 0xaf, 0xb0, 0xd6, 0x1b, 0x20,
	0x03, 0x7d, 0x44, 0x4c, 0x8c, 0x30, 0xc9, 0x55, 0x90, 0xa5, 0xb9, 0xf4, 0x52, 0xdc, 0x0e, 0x1d,
	0x40, 0xe0, 0x48, 0x7a, 0x89, 0xbb, 0xb1, 0x70, 0xa6, 0xe0, 0x0d, 0x5d, 0x9d, 0x8f, 0xc3, 0xd5,
	0x09, 0x48, 0xbc, 0x3a, 0x44, 0x06, 0x0f, 0xdf, 0x09, 0x34, 0x22, 0x87, 0xe4, 0xdd, 0x3e, 0x5b,
	0x27, 0x7a, 0x2e, 0x18, 0x84, 0x06, 0x22, 0xf6, 0x2a, 0x28, 0xe2, 0x48, 0x08, 0xf4, 0x27, 0xe7,
	0x23, 0x21, 0x90, 0x37, 0x58, 0x0f, 0xc9, 0x79, 0xf6, 0xdd, 0x50, 0x7e, 0x5c, 0xcf, 0x62, 0x32,
	0xba, 0xca, 0x58, 0xb4, 0xb5, 0xb8, 0x37, 0x9f, 0x0a, 0x14, 0x52, 0x13, 0x2d, 0x48, 0x0f, 0x79,
	0x2a, 0xbd, 0xb8, 0x1f, 0x8c, 0x1a, 0x91, 0x47, 0x1e, 0xe9, 0xba, 0xca, 0xe7, 0xf4, 0xa7, 0x81,
	0x8e, 0xc8, 0x23, 0xcf, 0xef, 0xb1, 0x35, 0x9a, 0x69, 0x62, 0xd8, 0x6f, 0x0c, 0x7a, 0x07, 0x97,
	0x87, 0x7f, 0x1f, 0xe7, 0xc3, 0xc7, 0x48, 0x26, 0x41, 0xc3, 0xef, 0xb0, 0xd5, 0x4c, 0x7a, 0x27,
	0x3e, 0xeb, 0x37, 0x07, 0xbd, 0x83, 0x8b, 0xef, 0x6b, 0x0f, 0xa5, 0x4f, 0x48, 0x80, 0x42, 0x2f,
	0x27, 0x4e, 0x7c, 0xbe, 0x58, 0xf8, 0x42, 0x4e, 0x12, 0x12, 0xa0, 0xd0, 0x9d, 0xd6, 0x4e, 0x3c,
	0x58, 0x2c, 0x1c, 0x9d, 0xd6, 0x09, 0x09, 0xf8, 0x57, 0xac, 0x3b, 0xa1, 0x67, 0x44, 0x81, 0x13,
	0x07, 0xa4, 0xde, 0x7b, 0x5f, 0x7d, 0xfc, 0xee, 0xad, 0x49, 0xde, 0xa9, 0xf9, 0x03, 0xd6, 0x22,
	0x97, 0x38, 0xf1, 0x90, 0xf2, 0x76, 0x16, 0xe5, 0x91, 0x6f, 0x92, 0x28, 0xc4, 0x14, 0x9a, 0x3f,
	0x4e, 0x7c, 0xf1, 0xef, 0x29, 0x34, 0x8c, 0x92, 0x28, 0xe4, 0x17, 0x58, 0x53, 0xe5, 0x4e, 0x7c,
	0xd9, 0x6f, 0x0e, 0x9a, 0x09, 0x2e, 0xf7, 0xff, 0x68, 0xb0, 0x75, 0xd2, 0x81, 0xab, 0x8c, 0x76,
	0xc0, 0xef, 0xb3, 0x16, 0x84, 0xa1, 0xdf, 0xa0, 0x66, 0x5f, 0x5a, 0x58, 0x35, 0x6a, 0xf0, 0xcb,
	0x54, 0x72, 0x02, 0x96, 0xde, 0xaf, 0x05, 0x5f, 0xe6, 0x39, 0x92, 0x49, 0xd0, 0xf0, 0xbb, 0x6c,
	0x0d, 0x47, 0x95, 0x13, 0x4d, 0xda, 0xef, 0xe2, 0xca, 0x41, 0x82, 0x85, 0xc1, 0x5a, 0x63, 0xe9,
	0xb5, 0x5b, 0x50, 0xf8, 0x07, 0x24, 0x93, 0xa0, 0xe1, 0x03, 0xb6, 0xaa, 0xf4, 0x89, 0xa1, 0xf7,
	0x6f, 0x51, 0x5d, 0x7d, 0x62, 0x12, 0x52, 0x1c, 0xfc, 0xde, 0x62, 0x3d, 0xfc, 0x99, 0x11, 0xd8,
	0x33, 0xbc, 0x83, 0xdf, 0xb2, 0xd6, 0x21, 0xb9, 0x90, 0x2f, 0xdc, 0xcd, 0xee, 0xb5, 0x85, 0x7b,
	0x8c, 0xbd, 0xda, 0xff, 0x1f, 0xe6, 0xbf, 0x24, 0x9b, 0x2e, 0x9f, 0x7f, 0x04, 0x05, 0x2c, 0x9d,
	0xff, 0x0d, 0x5b, 0x7d, 0x8a, 0x8f, 0xd8, 0xf2, 0xbb, 0xd7, 0xc5, 0xf2, 0xf9, 0xdf, 0xb1, 0x76,
	0x7c, 0xc6, 0x96, 0x2c, 0xf0, 0x88, 0x75, 0x8f, 0xa0, 0xf8, 0xa0, 0x12, 0xdf, 0xb3, 0x4e, 0x02,
	0x99, 0x39, 0x03, 0x3b, 0x5b, 0xfe, 0x14, 0x47, 0xe0, 0xbc, 0x35, 0xcb, 0x16, 0xf8, 0x9a, 0x35,
	0x9f, 0x80, 0x5f, 0x32, 0xf9, 0x90, 0xb5, 0x46, 0x20, 0x6d, 0x36, 0xe5, 0xff, 0xb8, 0xbf, 0x8f,
	0xa5, 0x83, 0x57, 0x53, 0xb0, 0xf0, 0x9f, 0x45, 0x7e, 0x62, 0x1b, 0xa1, 0x48, 0x30, 0x53, 0xfe,
	0x01, 0xb5, 0xc6, 0x2d, 0xfa, 0xc7, 0xfa, 0xf0, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x33,
	0x33, 0x92, 0x40, 0x0b, 0x00, 0x00,
}
