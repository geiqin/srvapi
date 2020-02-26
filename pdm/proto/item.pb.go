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
	Cats                 []*Cat         `protobuf:"bytes,47,rep,name=cats,proto3" json:"cats,omitempty"`
	Tags                 []*Tag         `protobuf:"bytes,48,rep,name=tags,proto3" json:"tags,omitempty"`
	Skus                 []*Sku         `protobuf:"bytes,49,rep,name=skus,proto3" json:"skus,omitempty"`
	Galleries            []*ItemGallery `protobuf:"bytes,50,rep,name=galleries,proto3" json:"galleries,omitempty"`
	Prices               []*ItemPrice   `protobuf:"bytes,51,rep,name=prices,proto3" json:"prices,omitempty"`
	Rights               []*ItemRight   `protobuf:"bytes,52,rep,name=rights,proto3" json:"rights,omitempty"`
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
	// 1159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5b, 0x73, 0x13, 0x37,
	0x14, 0xae, 0x73, 0x59, 0xdb, 0x72, 0x2e, 0x8d, 0xc2, 0x45, 0x49, 0xb8, 0x98, 0xb4, 0x14, 0x17,
	0xa8, 0x5b, 0x42, 0x5f, 0x98, 0xb6, 0xb4, 0x90, 0xb4, 0x4c, 0x3a, 0x0c, 0x43, 0xd7, 0x30, 0x3c,
	0xee, 0xc8, 0xbb, 0x27, 0x8e, 0x9a, 0x5d, 0x69, 0x91, 0xb4, 0x01, 0xf7, 0xb9, 0xbf, 0xa1, 0xbf,
	0xb2, 0x3f, 0xa2, 0x73, 0x8e, 0x64, 0x68, 0xc1, 0x1d, 0x66, 0xc2, 0x9b, 0xce, 0xf7, 0x7d, 0xe7,
	0xac, 0x74, 0xf4, 0x49, 0x5a, 0xc6, 0x94, 0x87, 0x6a, 0x58, 0x5b, 0xe3, 0x0d, 0x5f, 0x9b, 0x80,
	0x7a, 0xa9, 0xf4, 0xd0, 0xd9, 0xd3, 0x61, 0x5d, 0x54, 0xdb, 0x2b, 0xb9, 0xa9, 0x2a, 0xa3, 0x03,
	0xbb, 0xbd, 0x81, 0xca, 0x47, 0xb2, 0x2c, 0xc1, 0x4e, 0x23, 0xb4, 0x8e, 0xd0, 0x53, 0xab, 0x72,
	0xf8, 0x37, 0x90, 0xaa, 0xc9, 0xb1, 0x8f, 0x40, 0xd7, 0xcb, 0x49, 0x1c, 0xf6, 0xc6, 0x56, 0xea,
	0x62, 0x86, 0xbb, 0x93, 0x66, 0x36, 0xcc, 0x65, 0x54, 0xef, 0xfe, 0xb9, 0xca, 0x96, 0x0e, 0x3d,
	0x54, 0x7c, 0x8d, 0x2d, 0xa8, 0x42, 0xb4, 0xfa, 0xad, 0xc1, 0x62, 0xba, 0xa0, 0x0a, 0xce, 0xd9,
	0x92, 0x96, 0x15, 0x88, 0x85, 0x7e, 0x6b, 0xd0, 0x4d, 0x69, 0x8c, 0x58, 0xa3, 0x95, 0x17, 0x8b,
	0x01, 0xc3, 0x31, 0xbf, 0xc8, 0xda, 0x38, 0x83, 0xcc, 0x69, 0xb1, 0x44, 0x70, 0x82, 0xe1, 0x48,
	0xf3, 0x2d, 0xd6, 0xa1, 0xcf, 0x67, 0xaa, 0x10, 0xcb, 0x54, 0xb6, 0x4d, 0xf1, 0x61, 0xc1, 0xaf,
	0xb2, 0x9e, 0x97, 0xaf, 0x8d, 0x36, 0xd5, 0x14, 0xd9, 0x84, 0x58, 0x36, 0x83, 0x0e, 0x0b, 0xbe,
	0xc3, 0xba, 0xe3, 0x66, 0x9a, 0xbd, 0x6c, 0x8c, 0x97, 0xa2, 0xdd, 0x6f, 0x0d, 0x96, 0xd3, 0xce,
	0xb8, 0x99, 0xfe, 0x86, 0x31, 0xdf, 0x66, 0x9d, 0x97, 0x8d, 0xd4, 0x5e, 0xf9, 0xa9, 0xe8, 0x04,
	0x6e, 0x16, 0xe3, 0x47, 0x9d, 0x29, 0x8b, 0x4c, 0x37, 0x95, 0xe8, 0x12, 0xd7, 0xc6, 0xf8, 0x49,
	0x53, 0xf1, 0x5d, 0xb6, 0xaa, 0xb4, 0xf2, 0xd9, 0x1b, 0x9e, 0x11, 0xdf, 0x43, 0x70, 0x14, 0x35,
	0x97, 0x19, 0x3b, 0xb2, 0xe6, 0x0f, 0xd0, 0x24, 0xe8, 0x91, 0xa0, 0x1b, 0x90, 0x48, 0x57, 0xa6,
	0x80, 0x32, 0xf3, 0xd3, 0x1a, 0xc4, 0x0a, 0x2d, 0xb7, 0x4b, 0xc8, 0xb3, 0x69, 0x0d, 0xfc, 0x02,
	0x4b, 0x4a, 0xe5, 0x3c, 0x14, 0x62, 0xb5, 0xdf, 0x1a, 0x74, 0xd2, 0x18, 0x11, 0x6e, 0xf2, 0x13,
	0x28, 0xc4, 0x5a, 0xc4, 0x29, 0xc2, 0x72, 0xc7, 0xaa, 0x80, 0xcc, 0x79, 0x93, 0x9f, 0x88, 0x75,
	0xe2, 0xba, 0x88, 0x8c, 0x10, 0xe0, 0x43, 0xb6, 0xf9, 0xbb, 0x51, 0x3a, 0x2b, 0xe1, 0x14, 0xca,
	0xac, 0x50, 0x2e, 0x37, 0x8d, 0xf6, 0xe2, 0x53, 0xd2, 0x6d, 0x20, 0xf5, 0x18, 0x99, 0x83, 0x48,
	0xf0, 0xeb, 0x6c, 0xad, 0x6e, 0x6c, 0x7e, 0x2c, 0x1d, 0x64, 0x16, 0x0d, 0x21, 0x36, 0x48, 0xba,
	0x3a, 0x43, 0xc9, 0x25, 0xfc, 0x26, 0xdb, 0x90, 0x8d, 0x37, 0x19, 0x4e, 0x4e, 0xe9, 0x49, 0xe6,
	0x55, 0x05, 0x82, 0xd3, 0x5a, 0xd6, 0x91, 0x78, 0x1c, 0xf0, 0x67, 0xaa, 0x02, 0xdc, 0x87, 0xda,
	0x38, 0x1f, 0xd6, 0xbb, 0x49, 0x9a, 0x0e, 0x02, 0xb4, 0xdc, 0x2d, 0x46, 0xe3, 0xec, 0x08, 0x40,
	0x9c, 0xeb, 0xb7, 0x06, 0x0b, 0x69, 0x1b, 0xe3, 0x5f, 0x00, 0xf8, 0x15, 0xd6, 0x33, 0x8d, 0xcf,
	0xc8, 0x18, 0xda, 0x88, 0xf3, 0xa1, 0x53, 0xa6, 0xf1, 0x68, 0xb5, 0x27, 0x86, 0x0b, 0xd6, 0xce,
	0x8d, 0xf6, 0xa0, 0xbd, 0xb8, 0x40, 0xdc, 0x2c, 0x44, 0xc6, 0x35, 0x55, 0x25, 0xed, 0x54, 0x5c,
	0x0c, 0x4c, 0x0c, 0xf9, 0x0d, 0xb6, 0xee, 0xa0, 0x2c, 0x71, 0xca, 0xb5, 0x51, 0xda, 0xdf, 0xdd,
	0x13, 0x82, 0x14, 0x6b, 0x11, 0x7e, 0x1a, 0x50, 0x74, 0x64, 0xa5, 0x74, 0x36, 0x6e, 0xa6, 0x62,
	0x8b, 0x76, 0x30, 0xa9, 0x94, 0x7e, 0xd8, 0x4c, 0xf9, 0x35, 0xb6, 0xa2, 0x8d, 0xf3, 0xc6, 0x42,
	0x86, 0x29, 0x62, 0x9b, 0xda, 0xd3, 0x8b, 0xd8, 0x08, 0xca, 0x12, 0xd7, 0xe4, 0x8f, 0x9b, 0x6a,
	0x8c, 0xb6, 0xdc, 0x09, 0xa6, 0xa5, 0x38, 0x78, 0x32, 0x50, 0x8d, 0x2d, 0xc5, 0xa5, 0xd0, 0x0b,
	0x02, 0x9e, 0xdb, 0x92, 0x1c, 0x0d, 0x55, 0x5d, 0x4a, 0x0f, 0x98, 0x7a, 0x99, 0xbe, 0xcb, 0x66,
	0xd0, 0x61, 0x81, 0xeb, 0x1a, 0x4b, 0x9b, 0x9b, 0x02, 0xc4, 0x95, 0xb0, 0xae, 0x18, 0xf2, 0x73,
	0x6c, 0xb9, 0xc6, 0xf3, 0x2c, 0xae, 0x52, 0x0f, 0x43, 0x80, 0x73, 0x35, 0x56, 0x4d, 0x94, 0xce,
	0x02, 0xd9, 0x27, 0xb2, 0x17, 0x30, 0x3a, 0xff, 0x68, 0x9f, 0x1c, 0xfb, 0x1f, 0x04, 0xd7, 0x48,
	0xd0, 0x45, 0x24, 0xd0, 0x3b, 0xac, 0x5b, 0xbd, 0x49, 0xdf, 0x25, 0xb6, 0x53, 0xcd, 0x72, 0x91,
	0x94, 0xaf, 0x23, 0xf9, 0x59, 0x24, 0xe5, 0xeb, 0x40, 0x5e, 0x60, 0xc9, 0x2b, 0x20, 0x03, 0x7d,
	0x4e, 0x4c, 0x8c, 0x30, 0xc9, 0xd5, 0x90, 0x67, 0x85, 0xf4, 0x52, 0x5c, 0x0f, 0x1d, 0x40, 0xe0,
	0x40, 0x7a, 0x89, 0xb3, 0xb1, 0x70, 0xaa, 0xe0, 0x15, 0x1d, 0x9d, 0x2f, 0xc2, 0xd1, 0x09, 0x48,
	0x3c, 0x3a, 0x44, 0x06, 0x0f, 0xdf, 0x08, 0x34, 0x22, 0xfb, 0xe4, 0xdd, 0x3e, 0x5b, 0x21, 0x7a,
	0x26, 0x18, 0x84, 0x06, 0x22, 0xf6, 0x22, 0x28, 0xe2, 0x95, 0x10, 0xe8, 0x2f, 0xdf, 0x5c, 0x09,
	0x81, 0xbc, 0xc2, 0x7a, 0x48, 0xce, 0xb2, 0x6f, 0x86, 0xf2, 0xe3, 0x66, 0x1a, 0x93, 0xd1, 0x55,
	0xc6, 0xa2, 0xad, 0xc5, 0xad, 0xd9, 0xad, 0x40, 0x21, 0x35, 0xd1, 0x82, 0xf4, 0x50, 0x64, 0xd2,
	0x8b, 0xdb, 0xc1, 0xa8, 0x11, 0x79, 0xe0, 0x91, 0x6e, 0xea, 0x62, 0x46, 0x7f, 0x15, 0xe8, 0x88,
	0x3c, 0xf0, 0xfc, 0x16, 0x5b, 0xa6, 0x3b, 0x4d, 0x0c, 0xfb, 0xad, 0x41, 0x6f, 0xef, 0xfc, 0xf0,
	0xbf, 0xd7, 0xf9, 0xf0, 0x21, 0x92, 0x69, 0xd0, 0xf0, 0x1b, 0x6c, 0x29, 0x97, 0xde, 0x89, 0xaf,
	0xfb, 0x8b, 0x83, 0xde, 0xde, 0xe6, 0xbb, 0xda, 0x7d, 0xe9, 0x53, 0x12, 0xa0, 0xd0, 0xcb, 0x89,
	0x13, 0xdf, 0xcc, 0x17, 0x3e, 0x93, 0x93, 0x94, 0x04, 0x28, 0x74, 0x27, 0x8d, 0x13, 0x77, 0xe6,
	0x0b, 0x47, 0x27, 0x4d, 0x4a, 0x02, 0x7e, 0x8f, 0x75, 0x27, 0xf4, 0x8c, 0x28, 0x70, 0x62, 0x8f,
	0xd4, 0x3b, 0xef, 0xaa, 0x0f, 0xdf, 0xbe, 0x35, 0xe9, 0x5b, 0x35, 0xbf, 0xc3, 0x12, 0x72, 0x89,
	0x13, 0x77, 0x29, 0x6f, 0x6b, 0x5e, 0x1e, 0xf9, 0x26, 0x8d, 0x42, 0x4c, 0xa1, 0xfb, 0xc7, 0x89,
	0x6f, 0xff, 0x3f, 0x85, 0x2e, 0xa3, 0x34, 0x0a, 0x77, 0xff, 0x6e, 0xb1, 0x15, 0x42, 0xc1, 0xd5,
	0x46, 0x3b, 0xe0, 0xb7, 0x59, 0x02, 0xe1, 0x8a, 0x6f, 0x51, 0x6b, 0xcf, 0xcd, 0xad, 0x11, 0x35,
	0xb8, 0x0f, 0xb5, 0x9c, 0x80, 0xa5, 0xd7, 0x6a, 0xce, 0x3e, 0x3c, 0x45, 0x32, 0x0d, 0x1a, 0x7e,
	0x93, 0x2d, 0xe3, 0xc5, 0xe4, 0xc4, 0x22, 0xcd, 0x6e, 0x7e, 0xe5, 0x20, 0xc1, 0xc2, 0x60, 0xad,
	0xb1, 0xf4, 0xb6, 0xcd, 0x29, 0xfc, 0x33, 0x92, 0x69, 0xd0, 0xf0, 0x01, 0x5b, 0x52, 0xfa, 0xc8,
	0xd0, 0x6b, 0x37, 0xaf, 0xae, 0x3e, 0x32, 0x29, 0x29, 0xf6, 0xfe, 0x4a, 0x58, 0x0f, 0x3f, 0x33,
	0x02, 0x7b, 0x8a, 0x27, 0xee, 0x3e, 0x4b, 0xf6, 0xc9, 0x73, 0x7c, 0xee, 0x6c, 0xb6, 0x2f, 0xcd,
	0x9d, 0x63, 0xec, 0xd5, 0xee, 0x27, 0x98, 0xff, 0x9c, 0x4c, 0x79, 0xc6, 0xfc, 0xef, 0x59, 0x72,
	0x00, 0x25, 0x78, 0xe0, 0xfc, 0x3d, 0x65, 0xf1, 0xc1, 0xec, 0xef, 0xd8, 0xd2, 0x63, 0x7c, 0xb0,
	0x36, 0xdf, 0xcf, 0x75, 0x1f, 0x4c, 0xfe, 0x81, 0x25, 0xcf, 0x75, 0x79, 0xe6, 0xf4, 0xfb, 0xac,
	0x1d, 0x1f, 0xac, 0xb3, 0xe5, 0xff, 0xc4, 0xba, 0x07, 0x50, 0x7e, 0x4c, 0x85, 0x1f, 0x59, 0x27,
	0x85, 0xdc, 0x9c, 0x82, 0x9d, 0x9e, 0x79, 0x09, 0x07, 0xe0, 0xbc, 0x35, 0x67, 0xcc, 0xbf, 0xc7,
	0x16, 0x1f, 0x81, 0x3f, 0xd3, 0xce, 0xed, 0xb3, 0x64, 0x04, 0xd2, 0xe6, 0xc7, 0xfc, 0xbd, 0x33,
	0xfa, 0x50, 0x3a, 0x78, 0x71, 0x0c, 0x16, 0x3e, 0x58, 0xe4, 0x57, 0xb6, 0x1a, 0x8a, 0x04, 0x0b,
	0x15, 0x1f, 0x51, 0x6b, 0x9c, 0xd0, 0x5f, 0xe9, 0xdd, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0x1d, 0x98, 0xf1, 0x24, 0x0b, 0x00, 0x00,
}
