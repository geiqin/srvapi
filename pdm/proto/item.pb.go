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
	Id                int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Unit              string  `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	ItemSn            string  `protobuf:"bytes,4,opt,name=item_sn,json=itemSn,proto3" json:"item_sn,omitempty"`
	BrandId           int32   `protobuf:"varint,5,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	TaxonomyId        int64   `protobuf:"varint,6,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id,omitempty"`
	BuyQuota          int32   `protobuf:"varint,7,opt,name=buy_quota,json=buyQuota,proto3" json:"buy_quota,omitempty"`
	Quantity          int32   `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
	SoldNum           int32   `protobuf:"varint,9,opt,name=sold_num,json=soldNum,proto3" json:"sold_num,omitempty"`
	InitSoldNum       int32   `protobuf:"varint,10,opt,name=init_sold_num,json=initSoldNum,proto3" json:"init_sold_num,omitempty"`
	ModelType         string  `protobuf:"bytes,11,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	Listed            bool    `protobuf:"varint,12,opt,name=listed,proto3" json:"listed,omitempty"`
	Locked            bool    `protobuf:"varint,13,opt,name=locked,proto3" json:"locked,omitempty"`
	HideStock         bool    `protobuf:"varint,14,opt,name=hide_stock,json=hideStock,proto3" json:"hide_stock,omitempty"`
	JoinLevelDiscount bool    `protobuf:"varint,15,opt,name=join_level_discount,json=joinLevelDiscount,proto3" json:"join_level_discount,omitempty"`
	PurchaseRight     bool    `protobuf:"varint,16,opt,name=purchase_right,json=purchaseRight,proto3" json:"purchase_right,omitempty"`
	AutoListingTime   string  `protobuf:"bytes,17,opt,name=auto_listing_time,json=autoListingTime,proto3" json:"auto_listing_time,omitempty"`
	OutItemNo         string  `protobuf:"bytes,18,opt,name=out_item_no,json=outItemNo,proto3" json:"out_item_no,omitempty"`
	Content           string  `protobuf:"bytes,19,opt,name=content,proto3" json:"content,omitempty"`
	Summary           string  `protobuf:"bytes,20,opt,name=summary,proto3" json:"summary,omitempty"`
	SellingPoint      string  `protobuf:"bytes,21,opt,name=selling_point,json=sellingPoint,proto3" json:"selling_point,omitempty"`
	MinBuy            int32   `protobuf:"varint,22,opt,name=min_buy,json=minBuy,proto3" json:"min_buy,omitempty"`
	NostoreSell       bool    `protobuf:"varint,23,opt,name=nostore_sell,json=nostoreSell,proto3" json:"nostore_sell,omitempty"`
	ThumbId           int64   `protobuf:"varint,24,opt,name=thumb_id,json=thumbId,proto3" json:"thumb_id,omitempty"`
	ThumbUrl          string  `protobuf:"bytes,25,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	TemplateId        int32   `protobuf:"varint,26,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Barcode           string  `protobuf:"bytes,27,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Price             float32 `protobuf:"fixed32,28,opt,name=price,proto3" json:"price,omitempty"`
	OriginPrice       float32 `protobuf:"fixed32,29,opt,name=origin_price,json=originPrice,proto3" json:"origin_price,omitempty"`
	CostPrice         float32 `protobuf:"fixed32,30,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	MinPrice          float32 `protobuf:"fixed32,31,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice          float32 `protobuf:"fixed32,32,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	Weight            float32 `protobuf:"fixed32,33,opt,name=weight,proto3" json:"weight,omitempty"`
	ReviewNum         int32   `protobuf:"varint,34,opt,name=review_num,json=reviewNum,proto3" json:"review_num,omitempty"`
	ViewCount         int32   `protobuf:"varint,35,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	BuyCount          int32   `protobuf:"varint,36,opt,name=buy_count,json=buyCount,proto3" json:"buy_count,omitempty"`
	Sorting           int32   `protobuf:"varint,37,opt,name=sorting,proto3" json:"sorting,omitempty"`
	ListedAt          string  `protobuf:"bytes,38,opt,name=listed_at,json=listedAt,proto3" json:"listed_at,omitempty"`
	CreatedAt         string  `protobuf:"bytes,39,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt         string  `protobuf:"bytes,40,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Brand             *Brand  `protobuf:"bytes,41,opt,name=brand,proto3" json:"brand,omitempty"`
	// gorm:"many2many:item_cats;"
	Cats []*Cat `protobuf:"bytes,42,rep,name=cats,proto3" json:"cats,omitempty" gorm:"many2many:item_cats;"`
	// gorm:"many2many:item_tags;"
	Tags        []*Tag           `protobuf:"bytes,43,rep,name=tags,proto3" json:"tags,omitempty" gorm:"many2many:item_tags;"`
	Skus        []*Sku           `protobuf:"bytes,44,rep,name=skus,proto3" json:"skus,omitempty"`
	Galleries   []*ItemGallery   `protobuf:"bytes,45,rep,name=galleries,proto3" json:"galleries,omitempty"`
	Prices      []*ItemPrice     `protobuf:"bytes,46,rep,name=prices,proto3" json:"prices,omitempty"`
	Rights      []*ItemRight     `protobuf:"bytes,47,rep,name=rights,proto3" json:"rights,omitempty"`
	SpecIndexes []*SpecItemIndex `protobuf:"bytes,48,rep,name=spec_indexes,json=specIndexes,proto3" json:"spec_indexes,omitempty"`
	Ids         []int64          `protobuf:"varint,49,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	// gorm:"many2many:spec_item_indices;"
	Specs                []*Spec      `protobuf:"bytes,50,rep,name=specs,proto3" json:"specs,omitempty" gorm:"many2many:spec_item_indices;"`
	Skuitem              *Sku         `protobuf:"bytes,51,opt,name=skuitem,proto3" json:"skuitem,omitempty"`
	SkuId                int64        `protobuf:"varint,52,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	IsPresale            bool         `protobuf:"varint,53,opt,name=is_presale,json=isPresale,proto3" json:"is_presale,omitempty"`
	ItemPresale          *ItemPresale `protobuf:"bytes,54,opt,name=item_presale,json=itemPresale,proto3" json:"item_presale,omitempty"`
	IsSku                bool         `protobuf:"varint,55,opt,name=is_sku,json=isSku,proto3" json:"is_sku,omitempty"`
	Food                 *Food        `protobuf:"bytes,56,opt,name=food,proto3" json:"food,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
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

func (m *Item) GetBrandId() int32 {
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

func (m *Item) GetSellingPoint() string {
	if m != nil {
		return m.SellingPoint
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

func (m *Item) GetBuyCount() int32 {
	if m != nil {
		return m.BuyCount
	}
	return 0
}

func (m *Item) GetSorting() int32 {
	if m != nil {
		return m.Sorting
	}
	return 0
}

func (m *Item) GetListedAt() string {
	if m != nil {
		return m.ListedAt
	}
	return ""
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

func (m *Item) GetSpecIndexes() []*SpecItemIndex {
	if m != nil {
		return m.SpecIndexes
	}
	return nil
}

func (m *Item) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Item) GetSpecs() []*Spec {
	if m != nil {
		return m.Specs
	}
	return nil
}

func (m *Item) GetSkuitem() *Sku {
	if m != nil {
		return m.Skuitem
	}
	return nil
}

func (m *Item) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *Item) GetIsPresale() bool {
	if m != nil {
		return m.IsPresale
	}
	return false
}

func (m *Item) GetItemPresale() *ItemPresale {
	if m != nil {
		return m.ItemPresale
	}
	return nil
}

func (m *Item) GetIsSku() bool {
	if m != nil {
		return m.IsSku
	}
	return false
}

func (m *Item) GetFood() *Food {
	if m != nil {
		return m.Food
	}
	return nil
}

type GoodsWhere struct {
	ItemIds              []int64  `protobuf:"varint,1,rep,packed,name=item_ids,json=itemIds,proto3" json:"item_ids,omitempty"`
	SkuIds               []int64  `protobuf:"varint,2,rep,packed,name=sku_ids,json=skuIds,proto3" json:"sku_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsWhere) Reset()         { *m = GoodsWhere{} }
func (m *GoodsWhere) String() string { return proto.CompactTextString(m) }
func (*GoodsWhere) ProtoMessage()    {}
func (*GoodsWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{1}
}

func (m *GoodsWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsWhere.Unmarshal(m, b)
}
func (m *GoodsWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsWhere.Marshal(b, m, deterministic)
}
func (m *GoodsWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsWhere.Merge(m, src)
}
func (m *GoodsWhere) XXX_Size() int {
	return xxx_messageInfo_GoodsWhere.Size(m)
}
func (m *GoodsWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsWhere.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsWhere proto.InternalMessageInfo

func (m *GoodsWhere) GetItemIds() []int64 {
	if m != nil {
		return m.ItemIds
	}
	return nil
}

func (m *GoodsWhere) GetSkuIds() []int64 {
	if m != nil {
		return m.SkuIds
	}
	return nil
}

type ItemWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting              string   `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Keywords             string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	ItemSn               string   `protobuf:"bytes,6,opt,name=item_sn,json=itemSn,proto3" json:"item_sn,omitempty"`
	BrandId              int32    `protobuf:"varint,7,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	TagId                int32    `protobuf:"varint,8,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	CatId                int32    `protobuf:"varint,9,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"`
	TaxonomyId           int64    `protobuf:"varint,10,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id,omitempty"`
	MinPrice             float32  `protobuf:"fixed32,11,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice             float32  `protobuf:"fixed32,12,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	ModelType            string   `protobuf:"bytes,13,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	Status               string   `protobuf:"bytes,14,opt,name=status,proto3" json:"status,omitempty"`
	CouponId             int64    `protobuf:"varint,15,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	Valid                bool     `protobuf:"varint,16,opt,name=valid,proto3" json:"valid,omitempty"`
	Invite               string   `protobuf:"bytes,17,opt,name=invite,proto3" json:"invite,omitempty"`
	Id                   int64    `protobuf:"varint,18,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,19,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	SkuId                int64    `protobuf:"varint,20,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	CatSlug              string   `protobuf:"bytes,21,opt,name=cat_slug,json=catSlug,proto3" json:"cat_slug,omitempty"`
	IsSku                bool     `protobuf:"varint,22,opt,name=is_sku,json=isSku,proto3" json:"is_sku,omitempty"`
	KindType             string   `protobuf:"bytes,23,opt,name=kind_type,json=kindType,proto3" json:"kind_type,omitempty"`
	KindId               int64    `protobuf:"varint,24,opt,name=kind_id,json=kindId,proto3" json:"kind_id,omitempty"`
	KindIds              []int64  `protobuf:"varint,25,rep,packed,name=kind_ids,json=kindIds,proto3" json:"kind_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemWhere) Reset()         { *m = ItemWhere{} }
func (m *ItemWhere) String() string { return proto.CompactTextString(m) }
func (*ItemWhere) ProtoMessage()    {}
func (*ItemWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{2}
}

func (m *ItemWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemWhere.Unmarshal(m, b)
}
func (m *ItemWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemWhere.Marshal(b, m, deterministic)
}
func (m *ItemWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemWhere.Merge(m, src)
}
func (m *ItemWhere) XXX_Size() int {
	return xxx_messageInfo_ItemWhere.Size(m)
}
func (m *ItemWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemWhere.DiscardUnknown(m)
}

var xxx_messageInfo_ItemWhere proto.InternalMessageInfo

func (m *ItemWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *ItemWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ItemWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *ItemWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *ItemWhere) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ItemWhere) GetItemSn() string {
	if m != nil {
		return m.ItemSn
	}
	return ""
}

func (m *ItemWhere) GetBrandId() int32 {
	if m != nil {
		return m.BrandId
	}
	return 0
}

func (m *ItemWhere) GetTagId() int32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *ItemWhere) GetCatId() int32 {
	if m != nil {
		return m.CatId
	}
	return 0
}

func (m *ItemWhere) GetTaxonomyId() int64 {
	if m != nil {
		return m.TaxonomyId
	}
	return 0
}

func (m *ItemWhere) GetMinPrice() float32 {
	if m != nil {
		return m.MinPrice
	}
	return 0
}

func (m *ItemWhere) GetMaxPrice() float32 {
	if m != nil {
		return m.MaxPrice
	}
	return 0
}

func (m *ItemWhere) GetModelType() string {
	if m != nil {
		return m.ModelType
	}
	return ""
}

func (m *ItemWhere) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ItemWhere) GetCouponId() int64 {
	if m != nil {
		return m.CouponId
	}
	return 0
}

func (m *ItemWhere) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *ItemWhere) GetInvite() string {
	if m != nil {
		return m.Invite
	}
	return ""
}

func (m *ItemWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ItemWhere) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *ItemWhere) GetCatSlug() string {
	if m != nil {
		return m.CatSlug
	}
	return ""
}

func (m *ItemWhere) GetIsSku() bool {
	if m != nil {
		return m.IsSku
	}
	return false
}

func (m *ItemWhere) GetKindType() string {
	if m != nil {
		return m.KindType
	}
	return ""
}

func (m *ItemWhere) GetKindId() int64 {
	if m != nil {
		return m.KindId
	}
	return 0
}

func (m *ItemWhere) GetKindIds() []int64 {
	if m != nil {
		return m.KindIds
	}
	return nil
}

type MemberGoods struct {
	ItemId               int64    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId                int64    `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	MinMemberPrice       float32  `protobuf:"fixed32,3,opt,name=min_member_price,json=minMemberPrice,proto3" json:"min_member_price,omitempty"`
	MaxMemberPrice       float32  `protobuf:"fixed32,4,opt,name=max_member_price,json=maxMemberPrice,proto3" json:"max_member_price,omitempty"`
	Discount             float32  `protobuf:"fixed32,5,opt,name=discount,proto3" json:"discount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberGoods) Reset()         { *m = MemberGoods{} }
func (m *MemberGoods) String() string { return proto.CompactTextString(m) }
func (*MemberGoods) ProtoMessage()    {}
func (*MemberGoods) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{3}
}

func (m *MemberGoods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberGoods.Unmarshal(m, b)
}
func (m *MemberGoods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberGoods.Marshal(b, m, deterministic)
}
func (m *MemberGoods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberGoods.Merge(m, src)
}
func (m *MemberGoods) XXX_Size() int {
	return xxx_messageInfo_MemberGoods.Size(m)
}
func (m *MemberGoods) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberGoods.DiscardUnknown(m)
}

var xxx_messageInfo_MemberGoods proto.InternalMessageInfo

func (m *MemberGoods) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *MemberGoods) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *MemberGoods) GetMinMemberPrice() float32 {
	if m != nil {
		return m.MinMemberPrice
	}
	return 0
}

func (m *MemberGoods) GetMaxMemberPrice() float32 {
	if m != nil {
		return m.MaxMemberPrice
	}
	return 0
}

func (m *MemberGoods) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
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
	return fileDescriptor_6007f868cf6553df, []int{4}
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
	proto.RegisterType((*GoodsWhere)(nil), "geiqin.srv.pdm.GoodsWhere")
	proto.RegisterType((*ItemWhere)(nil), "geiqin.srv.pdm.ItemWhere")
	proto.RegisterType((*MemberGoods)(nil), "geiqin.srv.pdm.MemberGoods")
	proto.RegisterType((*ItemResponse)(nil), "geiqin.srv.pdm.ItemResponse")
}

func init() {
	proto.RegisterFile("item.proto", fileDescriptor_6007f868cf6553df)
}

var fileDescriptor_6007f868cf6553df = []byte{
	// 1624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x58, 0x6d, 0x73, 0x1b, 0xb7,
	0x11, 0x2e, 0x45, 0xf1, 0xc8, 0x5b, 0x4a, 0xb2, 0x0d, 0xbf, 0x41, 0x52, 0x9c, 0x30, 0x4a, 0xd3,
	0xb0, 0x4e, 0xa2, 0x36, 0x4e, 0x5f, 0xa7, 0x9d, 0xd4, 0xb6, 0xdc, 0x7a, 0xd8, 0x71, 0x33, 0xee,
	0xd1, 0x99, 0x7e, 0xbc, 0x81, 0xee, 0x60, 0x0a, 0xe5, 0x1d, 0x40, 0x1f, 0x70, 0xb2, 0x98, 0xaf,
	0xfd, 0x3d, 0x9d, 0xe9, 0x1f, 0xe9, 0x3f, 0xe9, 0x8f, 0xe8, 0xec, 0x02, 0x47, 0xd1, 0xf6, 0x25,
	0xe3, 0xa1, 0xbe, 0xdd, 0xee, 0xf3, 0xec, 0x02, 0x0b, 0xec, 0x0b, 0xe6, 0x00, 0x94, 0x93, 0xe5,
	0xf1, 0xa2, 0x32, 0xce, 0xb0, 0xbd, 0x99, 0x54, 0xaf, 0x94, 0x3e, 0xb6, 0xd5, 0xf9, 0xf1, 0x22,
	0x2f, 0x0f, 0x76, 0x32, 0x53, 0x96, 0x46, 0x7b, 0xf4, 0xe0, 0x06, 0x32, 0x9f, 0x8a, 0xa2, 0x90,
	0xd5, 0x32, 0xa8, 0xae, 0xa1, 0xea, 0x79, 0xa5, 0x32, 0xb9, 0xae, 0x48, 0xd4, 0xec, 0xcc, 0x05,
	0x45, 0xec, 0xc4, 0x2c, 0x7c, 0x0e, 0x4f, 0x2b, 0xa1, 0xf3, 0x46, 0x6f, 0xe7, 0x75, 0xf3, 0x99,
	0x89, 0x86, 0x0d, 0x76, 0x21, 0xb3, 0xf5, 0xe5, 0x9e, 0x57, 0xd2, 0x8a, 0xa2, 0xf1, 0x0e, 0x2f,
	0x8d, 0x09, 0x0e, 0x8e, 0xfe, 0xbb, 0x07, 0xdb, 0x13, 0x27, 0x4b, 0xb6, 0x07, 0x5b, 0x2a, 0xe7,
	0x9d, 0x51, 0x67, 0xdc, 0x4d, 0xb6, 0x54, 0xce, 0x18, 0x6c, 0x6b, 0x51, 0x4a, 0xbe, 0x35, 0xea,
	0x8c, 0xe3, 0x84, 0xbe, 0x51, 0x57, 0x6b, 0xe5, 0x78, 0xd7, 0xeb, 0xf0, 0x9b, 0xdd, 0x85, 0x3e,
	0xae, 0x90, 0x5a, 0xcd, 0xb7, 0x49, 0x1d, 0xa1, 0x38, 0xd5, 0x6c, 0x1f, 0x06, 0xb4, 0xd3, 0x54,
	0xe5, 0xbc, 0x37, 0xea, 0x8c, 0x7b, 0x49, 0x9f, 0xe4, 0x49, 0xce, 0x3e, 0x82, 0xa1, 0x13, 0x17,
	0x46, 0x9b, 0x72, 0x89, 0x68, 0x44, 0x8b, 0x42, 0xa3, 0x9a, 0xe4, 0xec, 0x10, 0xe2, 0xd3, 0x7a,
	0x99, 0xbe, 0xaa, 0x8d, 0x13, 0xbc, 0x4f, 0xc6, 0x83, 0xd3, 0x7a, 0xf9, 0x77, 0x94, 0xd9, 0x01,
	0x0c, 0x5e, 0xd5, 0x42, 0x3b, 0xe5, 0x96, 0x7c, 0xe0, 0xb1, 0x46, 0xc6, 0x45, 0xad, 0x29, 0xf2,
	0x54, 0xd7, 0x25, 0x8f, 0xfd, 0xa2, 0x28, 0x7f, 0x5b, 0x97, 0xec, 0x08, 0x76, 0x95, 0x56, 0x2e,
	0x5d, 0xe1, 0x40, 0xf8, 0x10, 0x95, 0xd3, 0xc0, 0xb9, 0x07, 0x50, 0x9a, 0x5c, 0x16, 0xa9, 0x5b,
	0x2e, 0x24, 0x1f, 0x52, 0x3c, 0x31, 0x69, 0x5e, 0x2c, 0x17, 0x92, 0xdd, 0x81, 0xa8, 0x50, 0xd6,
	0xc9, 0x9c, 0xef, 0x8c, 0x3a, 0xe3, 0x41, 0x12, 0x24, 0xd2, 0x9b, 0x6c, 0x2e, 0x73, 0xbe, 0x1b,
	0xf4, 0x24, 0xa1, 0xbb, 0x33, 0x95, 0xcb, 0xd4, 0x3a, 0x93, 0xcd, 0xf9, 0x1e, 0x61, 0x31, 0x6a,
	0xa6, 0xa8, 0x60, 0xc7, 0x70, 0xf3, 0x9f, 0x46, 0xe9, 0xb4, 0x90, 0xe7, 0xb2, 0x48, 0x73, 0x65,
	0x33, 0x53, 0x6b, 0xc7, 0xaf, 0x11, 0xef, 0x06, 0x42, 0xcf, 0x10, 0x79, 0x12, 0x00, 0xf6, 0x29,
	0xec, 0x2d, 0xea, 0x2a, 0x3b, 0x13, 0x56, 0xa6, 0x15, 0x26, 0x07, 0xbf, 0x4e, 0xd4, 0xdd, 0x46,
	0x4b, 0x19, 0xc3, 0xee, 0xc3, 0x0d, 0x51, 0x3b, 0x93, 0xe2, 0xe6, 0x94, 0x9e, 0xa5, 0x4e, 0x95,
	0x92, 0xdf, 0xa0, 0x58, 0xae, 0x21, 0xf0, 0xcc, 0xeb, 0x5f, 0xa8, 0x52, 0xb2, 0x0f, 0x61, 0x68,
	0x6a, 0x97, 0xd2, 0x0d, 0x6a, 0xc3, 0x99, 0x8f, 0xd8, 0xd4, 0x0e, 0x73, 0xe2, 0x5b, 0xc3, 0x38,
	0xf4, 0x33, 0xa3, 0x9d, 0xd4, 0x8e, 0xdf, 0x24, 0xac, 0x11, 0x11, 0xb1, 0x75, 0x59, 0x8a, 0x6a,
	0xc9, 0x6f, 0x79, 0x24, 0x88, 0xec, 0x13, 0xd8, 0xb5, 0xb2, 0x28, 0x70, 0xe9, 0x85, 0x51, 0xda,
	0xf1, 0xdb, 0x84, 0xef, 0x04, 0xe5, 0x73, 0xd4, 0x61, 0xda, 0x94, 0x4a, 0xa7, 0xa7, 0xf5, 0x92,
	0xdf, 0xa1, 0x7b, 0x88, 0x4a, 0xa5, 0x1f, 0xd7, 0x4b, 0xf6, 0x31, 0xec, 0x68, 0x63, 0x9d, 0xa9,
	0x64, 0x8a, 0x06, 0xfc, 0x2e, 0x85, 0x38, 0x0c, 0xba, 0xa9, 0x2c, 0x0a, 0xbc, 0x64, 0x77, 0x56,
	0x97, 0xa7, 0x98, 0x3b, 0x9c, 0x72, 0xa7, 0x4f, 0xb2, 0x4f, 0x1c, 0x0f, 0xd5, 0x55, 0xc1, 0xf7,
	0x69, 0x5d, 0xcf, 0xfd, 0xae, 0x2a, 0x28, 0xed, 0x64, 0xb9, 0x28, 0x84, 0x93, 0x68, 0x7a, 0x40,
	0xeb, 0x42, 0xa3, 0x9a, 0xe4, 0x18, 0xd3, 0xa9, 0xa8, 0x32, 0x93, 0x4b, 0x7e, 0xe8, 0x63, 0x0a,
	0x22, 0xbb, 0x05, 0xbd, 0x05, 0xd6, 0x27, 0xff, 0x60, 0xd4, 0x19, 0x6f, 0x25, 0x5e, 0xc0, 0xbd,
	0x9a, 0x4a, 0xcd, 0x94, 0x4e, 0x3d, 0x78, 0x8f, 0xc0, 0xa1, 0xd7, 0x51, 0x3d, 0x63, 0x0a, 0x64,
	0xc6, 0xba, 0x40, 0xf8, 0x90, 0x08, 0x31, 0x6a, 0x3c, 0x7c, 0x08, 0x71, 0xb9, 0x32, 0xff, 0x88,
	0xd0, 0x41, 0xd9, 0xd8, 0x22, 0x28, 0x2e, 0x02, 0x38, 0x0a, 0xa0, 0xb8, 0xf0, 0xe0, 0x1d, 0x88,
	0x5e, 0x4b, 0x4a, 0x82, 0x8f, 0x09, 0x09, 0x12, 0x2e, 0x58, 0xc9, 0x73, 0x25, 0x5f, 0x53, 0x8e,
	0x1f, 0x51, 0x8c, 0xb1, 0xd7, 0x84, 0x0c, 0x27, 0xd0, 0xa7, 0xda, 0x27, 0x1e, 0x46, 0xcd, 0x09,
	0xa5, 0x58, 0x28, 0x3c, 0x8f, 0xfe, 0x74, 0x55, 0x78, 0x1e, 0xc4, 0x2b, 0x37, 0x15, 0xe6, 0x0e,
	0xff, 0xb4, 0xa9, 0x2d, 0x12, 0xd1, 0xcc, 0x97, 0x42, 0x2a, 0x1c, 0xff, 0x99, 0x3f, 0x76, 0xaf,
	0x78, 0x44, 0x3b, 0xca, 0x2a, 0x29, 0x02, 0xfa, 0x99, 0x4f, 0xb1, 0xa0, 0xf1, 0x70, 0xbd, 0xc8,
	0x1b, 0x78, 0xec, 0xe1, 0xa0, 0x79, 0xe4, 0xd8, 0xe7, 0xd0, 0xa3, 0xb6, 0xc1, 0x7f, 0x3e, 0xea,
	0x8c, 0x87, 0x0f, 0x6e, 0x1f, 0xbf, 0xd9, 0x5c, 0x8f, 0x1f, 0x23, 0x98, 0x78, 0x0e, 0xfb, 0x0c,
	0xb6, 0x33, 0xe1, 0x2c, 0xbf, 0x3f, 0xea, 0x8e, 0x87, 0x0f, 0x6e, 0xbe, 0xcd, 0x3d, 0x11, 0x2e,
	0x21, 0x02, 0x12, 0x9d, 0x98, 0x59, 0xfe, 0x79, 0x3b, 0xf1, 0x85, 0x98, 0x25, 0x44, 0x40, 0xa2,
	0x9d, 0xd7, 0x96, 0x7f, 0xd1, 0x4e, 0x9c, 0xce, 0xeb, 0x84, 0x08, 0xec, 0xf7, 0x10, 0xcf, 0xa8,
	0xa9, 0x2b, 0x69, 0xf9, 0x97, 0xc4, 0x3e, 0x7c, 0x9b, 0x3d, 0xb9, 0xec, 0xfc, 0xc9, 0x25, 0x9b,
	0x7d, 0x05, 0x11, 0xdd, 0xb1, 0xe5, 0xc7, 0x64, 0xb7, 0xdf, 0x66, 0x47, 0xb7, 0x9e, 0x04, 0x22,
	0x9a, 0x50, 0x07, 0xb0, 0xfc, 0x17, 0x3f, 0x6c, 0x42, 0xed, 0x20, 0x09, 0x44, 0xf6, 0x10, 0x76,
	0x70, 0x2c, 0xa4, 0x4a, 0xe7, 0xf2, 0x42, 0x5a, 0xfe, 0x4b, 0x32, 0xbc, 0xf7, 0x4e, 0x44, 0x0b,
	0x99, 0xa1, 0xf1, 0x04, 0x69, 0xc9, 0x10, 0x4d, 0x26, 0xde, 0x82, 0x5d, 0x87, 0xae, 0xca, 0x2d,
	0xff, 0x6a, 0xd4, 0x1d, 0x77, 0x13, 0xfc, 0x64, 0xf7, 0xa1, 0x87, 0x04, 0xcb, 0x1f, 0x90, 0xb3,
	0x5b, 0x6d, 0xce, 0x12, 0x4f, 0x61, 0x5f, 0x42, 0xdf, 0xce, 0x6b, 0xec, 0x34, 0xfc, 0x6b, 0xba,
	0xca, 0xd6, 0xc3, 0x6c, 0x38, 0xec, 0x36, 0x44, 0x76, 0x5e, 0x63, 0x9d, 0xfe, 0x8a, 0x4a, 0xbc,
	0x67, 0xe7, 0xf5, 0x84, 0x5a, 0xaa, 0xb2, 0xe9, 0xc2, 0xcf, 0x33, 0xfe, 0x6b, 0xdf, 0x52, 0x95,
	0x0d, 0x03, 0x8e, 0x7d, 0x03, 0x3b, 0xd4, 0xcb, 0x1a, 0xc2, 0x6f, 0x68, 0xa5, 0xc3, 0xf6, 0x03,
	0x25, 0x4a, 0x32, 0x5c, 0x1b, 0x90, 0xb8, 0xaa, 0xb2, 0xa9, 0x9d, 0xd7, 0xfc, 0xb7, 0xe4, 0xba,
	0xa7, 0xec, 0x74, 0x5e, 0xb3, 0x31, 0x6c, 0xe3, 0xcc, 0xe4, 0xbf, 0x23, 0x77, 0xef, 0x84, 0xf9,
	0x17, 0x63, 0xf2, 0x84, 0x18, 0x47, 0x0f, 0x01, 0x9e, 0x1a, 0x93, 0xdb, 0x7f, 0x9c, 0xc9, 0x4a,
	0x62, 0xa7, 0xa2, 0xed, 0xe0, 0xb1, 0x75, 0xe8, 0xd8, 0x68, 0x58, 0x4e, 0x72, 0x8b, 0x0d, 0xd0,
	0xc7, 0x67, 0xf9, 0x16, 0x21, 0x11, 0x05, 0x68, 0x8f, 0xfe, 0xd5, 0x83, 0x18, 0xf7, 0xe7, 0x3d,
	0x60, 0xe3, 0x11, 0x33, 0xe9, 0x27, 0x73, 0x2f, 0xf1, 0x02, 0xd6, 0x1b, 0x7e, 0xa4, 0x56, 0x7d,
	0xef, 0x27, 0x74, 0x2f, 0x19, 0xa0, 0x62, 0xaa, 0xbe, 0x97, 0xeb, 0x65, 0xda, 0x0d, 0x9d, 0x39,
	0x94, 0xe9, 0x01, 0x0c, 0xe6, 0x72, 0xf9, 0xda, 0x54, 0xb9, 0x0d, 0xc3, 0x7a, 0x25, 0xaf, 0xe6,
	0x7d, 0x6f, 0x6d, 0xde, 0xaf, 0xcd, 0xf6, 0xe8, 0x07, 0x67, 0x7b, 0xff, 0xcd, 0xd9, 0x7e, 0x1b,
	0x22, 0x27, 0x66, 0x08, 0xf8, 0xd9, 0xdc, 0x73, 0x62, 0xe6, 0xd5, 0x99, 0x70, 0xa8, 0xf6, 0x63,
	0xb9, 0x97, 0x09, 0xf7, 0xee, 0x4b, 0x00, 0xda, 0x5e, 0x02, 0x97, 0x0d, 0x72, 0xf8, 0x63, 0x0d,
	0x72, 0xe7, 0xad, 0x06, 0xf9, 0xe6, 0x2c, 0xdf, 0x6d, 0x99, 0xe5, 0xd6, 0x09, 0x57, 0x5b, 0x9a,
	0xcb, 0x71, 0x12, 0x24, 0xf4, 0x99, 0x99, 0x7a, 0x61, 0x34, 0xee, 0xe7, 0x1a, 0xed, 0x67, 0xe0,
	0x15, 0x93, 0x1c, 0x6f, 0xe3, 0x5c, 0x14, 0x2a, 0x0f, 0x83, 0xd7, 0x0b, 0xe8, 0x4a, 0xe9, 0x73,
	0xe5, 0x9a, 0x29, 0x1b, 0xa4, 0xf0, 0xa4, 0x62, 0xab, 0x27, 0x55, 0xa8, 0x9f, 0x9b, 0x97, 0xf5,
	0x73, 0x99, 0xe4, 0xb7, 0xd6, 0x93, 0x7c, 0x1f, 0x06, 0x78, 0x58, 0xb6, 0xa8, 0x67, 0x61, 0x78,
	0xf6, 0x33, 0xe1, 0xa6, 0x45, 0x3d, 0x5b, 0x4b, 0xd0, 0x3b, 0xeb, 0x09, 0x7a, 0x08, 0xf1, 0x5c,
	0xe9, 0xdc, 0xc7, 0x7a, 0x37, 0x5c, 0xad, 0xd2, 0x39, 0x85, 0x7a, 0x17, 0xfa, 0x04, 0xae, 0xc6,
	0x65, 0x84, 0xa2, 0x5f, 0x27, 0x00, 0x96, 0xef, 0xfb, 0xf4, 0xf4, 0x88, 0x3d, 0xfa, 0x77, 0x07,
	0x86, 0x7f, 0x93, 0xe5, 0xa9, 0xac, 0x28, 0x9d, 0x57, 0xa9, 0xb0, 0x7a, 0x23, 0x46, 0x3e, 0x91,
	0xd7, 0x42, 0xd8, 0x5a, 0x0f, 0x61, 0x0c, 0xd7, 0xf1, 0xde, 0x4a, 0x72, 0x11, 0x6e, 0xa8, 0x4b,
	0x37, 0xb4, 0x57, 0x2a, 0xed, 0x3d, 0xfb, 0x7b, 0x42, 0xa6, 0xb8, 0x78, 0x93, 0xb9, 0x1d, 0x98,
	0xe2, 0x62, 0x9d, 0x79, 0x00, 0x83, 0xd5, 0x23, 0xa9, 0xe7, 0x6f, 0xbb, 0x91, 0x8f, 0xfe, 0xd7,
	0x81, 0x1d, 0xea, 0x79, 0xd2, 0x2e, 0x8c, 0xb6, 0x92, 0x7d, 0x01, 0x91, 0xf4, 0x6f, 0xc4, 0x4e,
	0x7b, 0xd1, 0x12, 0x3b, 0x70, 0x70, 0xca, 0x60, 0xfd, 0x54, 0x14, 0x44, 0xcb, 0x94, 0x79, 0x8e,
	0xa0, 0xaf, 0xbe, 0x0a, 0xbb, 0x1e, 0x06, 0x6f, 0x79, 0xb7, 0xbd, 0xeb, 0x91, 0x67, 0x4f, 0x41,
	0xc7, 0xb2, 0xaa, 0x4c, 0x45, 0x21, 0xb5, 0x38, 0xfe, 0x33, 0x82, 0x89, 0xe7, 0x60, 0x9b, 0x51,
	0xfa, 0xa5, 0xa1, 0xe0, 0xda, 0xfc, 0xea, 0x97, 0x26, 0x21, 0xc6, 0x83, 0xff, 0xc4, 0x30, 0xc4,
	0x65, 0xa6, 0xb2, 0x3a, 0xc7, 0xa3, 0xf9, 0x06, 0xa2, 0x13, 0x9a, 0xa8, 0xac, 0x75, 0x37, 0x07,
	0x1f, 0xb4, 0xee, 0x31, 0x9c, 0xd5, 0xd1, 0x4f, 0xd0, 0xfe, 0x3b, 0x1a, 0xb9, 0x9b, 0xdb, 0x3f,
	0x91, 0x85, 0xdc, 0xd8, 0xfe, 0x8f, 0xb0, 0xfd, 0x0c, 0x9f, 0xc4, 0x9b, 0xef, 0x5e, 0x17, 0x9b,
	0xdb, 0xff, 0x09, 0xfa, 0xe1, 0x51, 0xbc, 0xa1, 0x83, 0x47, 0x10, 0x3f, 0x91, 0xc5, 0x95, 0x5c,
	0x3c, 0x84, 0x41, 0x22, 0x33, 0x73, 0x2e, 0xab, 0xe5, 0xe6, 0x51, 0x3c, 0x91, 0xd6, 0x55, 0x66,
	0x53, 0x07, 0x7f, 0x80, 0xee, 0x53, 0xe9, 0x36, 0x34, 0x7e, 0x0a, 0x31, 0x75, 0x0a, 0x3c, 0x48,
	0x76, 0xf0, 0x36, 0xf9, 0x72, 0x26, 0xbe, 0x5f, 0x2a, 0x39, 0xa1, 0x8a, 0x2b, 0x1c, 0x83, 0xb2,
	0x8b, 0x42, 0x2c, 0x37, 0xbf, 0xcc, 0xe9, 0xbc, 0xbe, 0xd2, 0x1e, 0x30, 0x9d, 0xf1, 0x1c, 0x36,
	0xb3, 0x3e, 0x81, 0x68, 0x2a, 0x45, 0x95, 0x9d, 0xb1, 0xd6, 0x67, 0xdd, 0xfb, 0x1d, 0xe3, 0x5f,
	0x61, 0xd7, 0x3b, 0xf1, 0x75, 0x99, 0x5f, 0xc5, 0xd7, 0x09, 0x0c, 0xbd, 0x97, 0xc7, 0x4b, 0x7c,
	0xba, 0x6c, 0x14, 0xd5, 0x69, 0x44, 0x3f, 0x1c, 0xbe, 0xfe, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x74, 0x49, 0x71, 0x66, 0x2a, 0x11, 0x00, 0x00,
}
