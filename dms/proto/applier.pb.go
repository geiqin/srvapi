// Code generated by protoc-gen-go. DO NOT EDIT.
// source: applier.proto

package geiqin_srv_dms

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

type ApplierWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top                  int32    `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Mobile               string   `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Keywords             string   `protobuf:"bytes,6,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Status               string   `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Ids                  []int64  `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	RankType             string   `protobuf:"bytes,9,opt,name=rank_type,json=rankType,proto3" json:"rank_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplierWhere) Reset()         { *m = ApplierWhere{} }
func (m *ApplierWhere) String() string { return proto.CompactTextString(m) }
func (*ApplierWhere) ProtoMessage()    {}
func (*ApplierWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_83e53247d573818c, []int{0}
}

func (m *ApplierWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplierWhere.Unmarshal(m, b)
}
func (m *ApplierWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplierWhere.Marshal(b, m, deterministic)
}
func (m *ApplierWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplierWhere.Merge(m, src)
}
func (m *ApplierWhere) XXX_Size() int {
	return xxx_messageInfo_ApplierWhere.Size(m)
}
func (m *ApplierWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplierWhere.DiscardUnknown(m)
}

var xxx_messageInfo_ApplierWhere proto.InternalMessageInfo

func (m *ApplierWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *ApplierWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ApplierWhere) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *ApplierWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ApplierWhere) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *ApplierWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *ApplierWhere) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ApplierWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ApplierWhere) GetRankType() string {
	if m != nil {
		return m.RankType
	}
	return ""
}

type Applier struct {
	Id                   int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DistributorId        int64                 `protobuf:"varint,2,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	RankId               int32                 `protobuf:"varint,3,opt,name=rank_id,json=rankId,proto3" json:"rank_id,omitempty"`
	DisplayName          string                `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Autoed               bool                  `protobuf:"varint,5,opt,name=autoed,proto3" json:"autoed,omitempty"`
	CustomerId           int64                 `protobuf:"varint,6,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Status               string                `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Reason               string                `protobuf:"bytes,8,opt,name=reason,proto3" json:"reason,omitempty"`
	Remarks              string                `protobuf:"bytes,9,opt,name=remarks,proto3" json:"remarks,omitempty"`
	Addons               string                `protobuf:"bytes,10,opt,name=addons,proto3" json:"addons,omitempty"`
	CreatedAt            string                `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string                `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Distributor          *Distributor          `protobuf:"bytes,13,opt,name=distributor,proto3" json:"distributor,omitempty"`
	Rank                 *Rank                 `protobuf:"bytes,14,opt,name=rank,proto3" json:"rank,omitempty"`
	InviteSn             string                `protobuf:"bytes,15,opt,name=invite_sn,json=inviteSn,proto3" json:"invite_sn,omitempty"`
	Condition            *ApplierCondition     `protobuf:"bytes,16,opt,name=condition,proto3" json:"condition,omitempty"`
	MeetCondition        *ApplierMeetCondition `protobuf:"bytes,17,opt,name=meet_condition,json=meetCondition,proto3" json:"meet_condition,omitempty"`
	RankType             string                `protobuf:"bytes,18,opt,name=rank_type,json=rankType,proto3" json:"rank_type,omitempty"`
	Mobile               string                `protobuf:"bytes,19,opt,name=mobile,proto3" json:"mobile,omitempty"`
	InviteName           string                `protobuf:"bytes,20,opt,name=invite_name,json=inviteName,proto3" json:"invite_name,omitempty"`
	RankName             string                `protobuf:"bytes,21,opt,name=rank_name,json=rankName,proto3" json:"rank_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Applier) Reset()         { *m = Applier{} }
func (m *Applier) String() string { return proto.CompactTextString(m) }
func (*Applier) ProtoMessage()    {}
func (*Applier) Descriptor() ([]byte, []int) {
	return fileDescriptor_83e53247d573818c, []int{1}
}

func (m *Applier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Applier.Unmarshal(m, b)
}
func (m *Applier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Applier.Marshal(b, m, deterministic)
}
func (m *Applier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Applier.Merge(m, src)
}
func (m *Applier) XXX_Size() int {
	return xxx_messageInfo_Applier.Size(m)
}
func (m *Applier) XXX_DiscardUnknown() {
	xxx_messageInfo_Applier.DiscardUnknown(m)
}

var xxx_messageInfo_Applier proto.InternalMessageInfo

func (m *Applier) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Applier) GetDistributorId() int64 {
	if m != nil {
		return m.DistributorId
	}
	return 0
}

func (m *Applier) GetRankId() int32 {
	if m != nil {
		return m.RankId
	}
	return 0
}

func (m *Applier) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Applier) GetAutoed() bool {
	if m != nil {
		return m.Autoed
	}
	return false
}

func (m *Applier) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Applier) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Applier) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Applier) GetRemarks() string {
	if m != nil {
		return m.Remarks
	}
	return ""
}

func (m *Applier) GetAddons() string {
	if m != nil {
		return m.Addons
	}
	return ""
}

func (m *Applier) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Applier) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Applier) GetDistributor() *Distributor {
	if m != nil {
		return m.Distributor
	}
	return nil
}

func (m *Applier) GetRank() *Rank {
	if m != nil {
		return m.Rank
	}
	return nil
}

func (m *Applier) GetInviteSn() string {
	if m != nil {
		return m.InviteSn
	}
	return ""
}

func (m *Applier) GetCondition() *ApplierCondition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *Applier) GetMeetCondition() *ApplierMeetCondition {
	if m != nil {
		return m.MeetCondition
	}
	return nil
}

func (m *Applier) GetRankType() string {
	if m != nil {
		return m.RankType
	}
	return ""
}

func (m *Applier) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Applier) GetInviteName() string {
	if m != nil {
		return m.InviteName
	}
	return ""
}

func (m *Applier) GetRankName() string {
	if m != nil {
		return m.RankName
	}
	return ""
}

type ApplierResponse struct {
	Entity               *Applier   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Applier `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ApplierResponse) Reset()         { *m = ApplierResponse{} }
func (m *ApplierResponse) String() string { return proto.CompactTextString(m) }
func (*ApplierResponse) ProtoMessage()    {}
func (*ApplierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83e53247d573818c, []int{2}
}

func (m *ApplierResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplierResponse.Unmarshal(m, b)
}
func (m *ApplierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplierResponse.Marshal(b, m, deterministic)
}
func (m *ApplierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplierResponse.Merge(m, src)
}
func (m *ApplierResponse) XXX_Size() int {
	return xxx_messageInfo_ApplierResponse.Size(m)
}
func (m *ApplierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplierResponse proto.InternalMessageInfo

func (m *ApplierResponse) GetEntity() *Applier {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ApplierResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ApplierResponse) GetItems() []*Applier {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ApplierResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ApplierResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

type ApplyInfo struct {
	IsApply              bool                    `protobuf:"varint,1,opt,name=is_apply,json=isApply,proto3" json:"is_apply,omitempty"`
	Applier              *Applier                `protobuf:"bytes,2,opt,name=applier,proto3" json:"applier,omitempty"`
	Ranks                []*Rank                 `protobuf:"bytes,3,rep,name=ranks,proto3" json:"ranks,omitempty"`
	Condition            []*ApplierCondition     `protobuf:"bytes,4,rep,name=condition,proto3" json:"condition,omitempty"`
	MeetCondition        []*ApplierMeetCondition `protobuf:"bytes,5,rep,name=meet_condition,json=meetCondition,proto3" json:"meet_condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ApplyInfo) Reset()         { *m = ApplyInfo{} }
func (m *ApplyInfo) String() string { return proto.CompactTextString(m) }
func (*ApplyInfo) ProtoMessage()    {}
func (*ApplyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_83e53247d573818c, []int{3}
}

func (m *ApplyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyInfo.Unmarshal(m, b)
}
func (m *ApplyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyInfo.Marshal(b, m, deterministic)
}
func (m *ApplyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyInfo.Merge(m, src)
}
func (m *ApplyInfo) XXX_Size() int {
	return xxx_messageInfo_ApplyInfo.Size(m)
}
func (m *ApplyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyInfo proto.InternalMessageInfo

func (m *ApplyInfo) GetIsApply() bool {
	if m != nil {
		return m.IsApply
	}
	return false
}

func (m *ApplyInfo) GetApplier() *Applier {
	if m != nil {
		return m.Applier
	}
	return nil
}

func (m *ApplyInfo) GetRanks() []*Rank {
	if m != nil {
		return m.Ranks
	}
	return nil
}

func (m *ApplyInfo) GetCondition() []*ApplierCondition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *ApplyInfo) GetMeetCondition() []*ApplierMeetCondition {
	if m != nil {
		return m.MeetCondition
	}
	return nil
}

type ApplyInfoResponse struct {
	Entity               *ApplyInfo `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Error                *Error     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info      `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ApplyInfoResponse) Reset()         { *m = ApplyInfoResponse{} }
func (m *ApplyInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ApplyInfoResponse) ProtoMessage()    {}
func (*ApplyInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83e53247d573818c, []int{4}
}

func (m *ApplyInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyInfoResponse.Unmarshal(m, b)
}
func (m *ApplyInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyInfoResponse.Marshal(b, m, deterministic)
}
func (m *ApplyInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyInfoResponse.Merge(m, src)
}
func (m *ApplyInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ApplyInfoResponse.Size(m)
}
func (m *ApplyInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyInfoResponse proto.InternalMessageInfo

func (m *ApplyInfoResponse) GetEntity() *ApplyInfo {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ApplyInfoResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ApplyInfoResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

// 加入销售员的配置条件
type ApplierCondition struct {
	RankId               int32             `protobuf:"varint,1,opt,name=rank_id,json=rankId,proto3" json:"rank_id,omitempty"`
	IsBought             bool              `protobuf:"varint,2,opt,name=is_bought,json=isBought,proto3" json:"is_bought,omitempty"`
	ConsumptionAmount    float32           `protobuf:"fixed32,3,opt,name=consumption_amount,json=consumptionAmount,proto3" json:"consumption_amount,omitempty"`
	ConsumptionNum       int32             `protobuf:"varint,4,opt,name=consumption_num,json=consumptionNum,proto3" json:"consumption_num,omitempty"`
	IsBindIdcard         bool              `protobuf:"varint,5,opt,name=is_bind_idcard,json=isBindIdcard,proto3" json:"is_bind_idcard,omitempty"`
	IsBindMobile         bool              `protobuf:"varint,6,opt,name=is_bind_mobile,json=isBindMobile,proto3" json:"is_bind_mobile,omitempty"`
	IsInformation        bool              `protobuf:"varint,7,opt,name=is_information,json=isInformation,proto3" json:"is_information,omitempty"`
	InformationFields    string            `protobuf:"bytes,8,opt,name=information_fields,json=informationFields,proto3" json:"information_fields,omitempty"`
	Joined               bool              `protobuf:"varint,9,opt,name=joined,proto3" json:"joined,omitempty"`
	JoinAmount           float32           `protobuf:"fixed32,10,opt,name=join_amount,json=joinAmount,proto3" json:"join_amount,omitempty"`
	PromotionAmount      float32           `protobuf:"fixed32,11,opt,name=promotion_amount,json=promotionAmount,proto3" json:"promotion_amount,omitempty"`
	PrimaryNum           int32             `protobuf:"varint,12,opt,name=primary_num,json=primaryNum,proto3" json:"primary_num,omitempty"`
	SecondNum            int32             `protobuf:"varint,13,opt,name=second_num,json=secondNum,proto3" json:"second_num,omitempty"`
	GoodsCondition       []*GoodsCondition `protobuf:"bytes,14,rep,name=goods_condition,json=goodsCondition,proto3" json:"goods_condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ApplierCondition) Reset()         { *m = ApplierCondition{} }
func (m *ApplierCondition) String() string { return proto.CompactTextString(m) }
func (*ApplierCondition) ProtoMessage()    {}
func (*ApplierCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_83e53247d573818c, []int{5}
}

func (m *ApplierCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplierCondition.Unmarshal(m, b)
}
func (m *ApplierCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplierCondition.Marshal(b, m, deterministic)
}
func (m *ApplierCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplierCondition.Merge(m, src)
}
func (m *ApplierCondition) XXX_Size() int {
	return xxx_messageInfo_ApplierCondition.Size(m)
}
func (m *ApplierCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplierCondition.DiscardUnknown(m)
}

var xxx_messageInfo_ApplierCondition proto.InternalMessageInfo

func (m *ApplierCondition) GetRankId() int32 {
	if m != nil {
		return m.RankId
	}
	return 0
}

func (m *ApplierCondition) GetIsBought() bool {
	if m != nil {
		return m.IsBought
	}
	return false
}

func (m *ApplierCondition) GetConsumptionAmount() float32 {
	if m != nil {
		return m.ConsumptionAmount
	}
	return 0
}

func (m *ApplierCondition) GetConsumptionNum() int32 {
	if m != nil {
		return m.ConsumptionNum
	}
	return 0
}

func (m *ApplierCondition) GetIsBindIdcard() bool {
	if m != nil {
		return m.IsBindIdcard
	}
	return false
}

func (m *ApplierCondition) GetIsBindMobile() bool {
	if m != nil {
		return m.IsBindMobile
	}
	return false
}

func (m *ApplierCondition) GetIsInformation() bool {
	if m != nil {
		return m.IsInformation
	}
	return false
}

func (m *ApplierCondition) GetInformationFields() string {
	if m != nil {
		return m.InformationFields
	}
	return ""
}

func (m *ApplierCondition) GetJoined() bool {
	if m != nil {
		return m.Joined
	}
	return false
}

func (m *ApplierCondition) GetJoinAmount() float32 {
	if m != nil {
		return m.JoinAmount
	}
	return 0
}

func (m *ApplierCondition) GetPromotionAmount() float32 {
	if m != nil {
		return m.PromotionAmount
	}
	return 0
}

func (m *ApplierCondition) GetPrimaryNum() int32 {
	if m != nil {
		return m.PrimaryNum
	}
	return 0
}

func (m *ApplierCondition) GetSecondNum() int32 {
	if m != nil {
		return m.SecondNum
	}
	return 0
}

func (m *ApplierCondition) GetGoodsCondition() []*GoodsCondition {
	if m != nil {
		return m.GoodsCondition
	}
	return nil
}

// 满足加入销售员配置条件项
type ApplierMeetCondition struct {
	RankId               int32    `protobuf:"varint,1,opt,name=rank_id,json=rankId,proto3" json:"rank_id,omitempty"`
	HasBought            bool     `protobuf:"varint,2,opt,name=has_bought,json=hasBought,proto3" json:"has_bought,omitempty"`
	HasConsumptionAmount bool     `protobuf:"varint,3,opt,name=has_consumption_amount,json=hasConsumptionAmount,proto3" json:"has_consumption_amount,omitempty"`
	HasConsumptionNum    bool     `protobuf:"varint,4,opt,name=has_consumption_num,json=hasConsumptionNum,proto3" json:"has_consumption_num,omitempty"`
	HasBindIdcard        bool     `protobuf:"varint,5,opt,name=has_bind_idcard,json=hasBindIdcard,proto3" json:"has_bind_idcard,omitempty"`
	HasBindMobile        bool     `protobuf:"varint,6,opt,name=has_bind_mobile,json=hasBindMobile,proto3" json:"has_bind_mobile,omitempty"`
	HasInformation       bool     `protobuf:"varint,7,opt,name=has_information,json=hasInformation,proto3" json:"has_information,omitempty"`
	HasJoined            bool     `protobuf:"varint,8,opt,name=has_joined,json=hasJoined,proto3" json:"has_joined,omitempty"`
	HasPromotionAmount   bool     `protobuf:"varint,9,opt,name=has_promotion_amount,json=hasPromotionAmount,proto3" json:"has_promotion_amount,omitempty"`
	HasPrimaryNum        bool     `protobuf:"varint,10,opt,name=has_primary_num,json=hasPrimaryNum,proto3" json:"has_primary_num,omitempty"`
	HasSecondNum         bool     `protobuf:"varint,11,opt,name=has_second_num,json=hasSecondNum,proto3" json:"has_second_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplierMeetCondition) Reset()         { *m = ApplierMeetCondition{} }
func (m *ApplierMeetCondition) String() string { return proto.CompactTextString(m) }
func (*ApplierMeetCondition) ProtoMessage()    {}
func (*ApplierMeetCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_83e53247d573818c, []int{6}
}

func (m *ApplierMeetCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplierMeetCondition.Unmarshal(m, b)
}
func (m *ApplierMeetCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplierMeetCondition.Marshal(b, m, deterministic)
}
func (m *ApplierMeetCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplierMeetCondition.Merge(m, src)
}
func (m *ApplierMeetCondition) XXX_Size() int {
	return xxx_messageInfo_ApplierMeetCondition.Size(m)
}
func (m *ApplierMeetCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplierMeetCondition.DiscardUnknown(m)
}

var xxx_messageInfo_ApplierMeetCondition proto.InternalMessageInfo

func (m *ApplierMeetCondition) GetRankId() int32 {
	if m != nil {
		return m.RankId
	}
	return 0
}

func (m *ApplierMeetCondition) GetHasBought() bool {
	if m != nil {
		return m.HasBought
	}
	return false
}

func (m *ApplierMeetCondition) GetHasConsumptionAmount() bool {
	if m != nil {
		return m.HasConsumptionAmount
	}
	return false
}

func (m *ApplierMeetCondition) GetHasConsumptionNum() bool {
	if m != nil {
		return m.HasConsumptionNum
	}
	return false
}

func (m *ApplierMeetCondition) GetHasBindIdcard() bool {
	if m != nil {
		return m.HasBindIdcard
	}
	return false
}

func (m *ApplierMeetCondition) GetHasBindMobile() bool {
	if m != nil {
		return m.HasBindMobile
	}
	return false
}

func (m *ApplierMeetCondition) GetHasInformation() bool {
	if m != nil {
		return m.HasInformation
	}
	return false
}

func (m *ApplierMeetCondition) GetHasJoined() bool {
	if m != nil {
		return m.HasJoined
	}
	return false
}

func (m *ApplierMeetCondition) GetHasPromotionAmount() bool {
	if m != nil {
		return m.HasPromotionAmount
	}
	return false
}

func (m *ApplierMeetCondition) GetHasPrimaryNum() bool {
	if m != nil {
		return m.HasPrimaryNum
	}
	return false
}

func (m *ApplierMeetCondition) GetHasSecondNum() bool {
	if m != nil {
		return m.HasSecondNum
	}
	return false
}

func init() {
	proto.RegisterType((*ApplierWhere)(nil), "geiqin.srv.dms.ApplierWhere")
	proto.RegisterType((*Applier)(nil), "geiqin.srv.dms.Applier")
	proto.RegisterType((*ApplierResponse)(nil), "geiqin.srv.dms.ApplierResponse")
	proto.RegisterType((*ApplyInfo)(nil), "geiqin.srv.dms.ApplyInfo")
	proto.RegisterType((*ApplyInfoResponse)(nil), "geiqin.srv.dms.ApplyInfoResponse")
	proto.RegisterType((*ApplierCondition)(nil), "geiqin.srv.dms.ApplierCondition")
	proto.RegisterType((*ApplierMeetCondition)(nil), "geiqin.srv.dms.ApplierMeetCondition")
}

func init() {
	proto.RegisterFile("applier.proto", fileDescriptor_83e53247d573818c)
}

var fileDescriptor_83e53247d573818c = []byte{
	// 1210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x97, 0xdd, 0x72, 0x1b, 0xb5,
	0x17, 0xc0, 0xff, 0xb6, 0x63, 0x7b, 0x7d, 0xfc, 0x15, 0xab, 0x6e, 0xbb, 0x4d, 0xff, 0x25, 0xae,
	0xa7, 0xa5, 0x06, 0xa6, 0x86, 0x06, 0x6e, 0x61, 0xc6, 0x4d, 0x4b, 0xc7, 0x30, 0xe9, 0x64, 0xd6,
	0xcc, 0x70, 0xb9, 0xa3, 0x58, 0x4a, 0x2c, 0x92, 0x5d, 0x2d, 0x2b, 0x39, 0xe0, 0x3e, 0x06, 0xf7,
	0x3c, 0x03, 0x2f, 0xc2, 0x1b, 0x30, 0xc3, 0x0d, 0x57, 0x3c, 0x00, 0xf7, 0x8c, 0x8e, 0xb4, 0xf6,
	0xda, 0x89, 0xa1, 0x90, 0x0b, 0xee, 0x74, 0x3e, 0x74, 0x74, 0xce, 0xd1, 0x4f, 0xc7, 0x6b, 0x68,
	0xd2, 0x24, 0xb9, 0x10, 0x3c, 0x1d, 0x26, 0xa9, 0xd4, 0x92, 0xb4, 0xce, 0xb8, 0xf8, 0x56, 0xc4,
	0x43, 0x95, 0x5e, 0x0e, 0x59, 0xa4, 0xf6, 0x1a, 0x53, 0x19, 0x45, 0x32, 0xb6, 0xd6, 0xbd, 0x0e,
	0x13, 0x4a, 0xa7, 0xe2, 0x64, 0xae, 0xa5, 0xdb, 0xb0, 0x07, 0x29, 0x8d, 0xcf, 0xdd, 0xba, 0x7b,
	0x26, 0x25, 0x53, 0x87, 0x32, 0x66, 0x42, 0x8b, 0x6c, 0x53, 0xff, 0xd7, 0x02, 0x34, 0x46, 0xf6,
	0x90, 0xaf, 0x67, 0x3c, 0xe5, 0xa4, 0x0b, 0xe5, 0x84, 0x9e, 0x71, 0xe6, 0x17, 0x7a, 0x85, 0x41,
	0x39, 0xb0, 0x02, 0xb9, 0x0f, 0x35, 0xb3, 0x08, 0x95, 0x78, 0xc3, 0xfd, 0x22, 0x5a, 0x3c, 0xa3,
	0x98, 0x88, 0x37, 0x9c, 0xec, 0x42, 0x49, 0xcb, 0xc4, 0x2f, 0xa1, 0xda, 0x2c, 0x49, 0x0b, 0x8a,
	0x82, 0xf9, 0x3b, 0xbd, 0xc2, 0xa0, 0x14, 0x14, 0x05, 0x23, 0x77, 0xa0, 0x12, 0xc9, 0x13, 0x71,
	0xc1, 0xfd, 0x72, 0xaf, 0x30, 0xa8, 0x05, 0x4e, 0x22, 0x7b, 0xe0, 0x9d, 0xf3, 0xc5, 0x77, 0x32,
	0x65, 0xca, 0xaf, 0xa0, 0x65, 0x29, 0x9b, 0x3d, 0x4a, 0x53, 0x3d, 0x57, 0x7e, 0xd5, 0xee, 0xb1,
	0x92, 0x39, 0x4d, 0x30, 0xe5, 0x7b, 0xbd, 0xd2, 0xa0, 0x14, 0x98, 0xa5, 0x49, 0xce, 0xd4, 0x19,
	0xea, 0x45, 0xc2, 0xfd, 0x9a, 0x0d, 0x63, 0x14, 0x5f, 0x2d, 0x12, 0xde, 0xff, 0xa5, 0x0c, 0x55,
	0x57, 0xa0, 0x4b, 0xab, 0xb0, 0x4c, 0xeb, 0x31, 0xb4, 0x72, 0x3d, 0x0b, 0x05, 0xc3, 0xd2, 0x4a,
	0x41, 0x33, 0xa7, 0x1d, 0x33, 0x72, 0x17, 0xaa, 0x18, 0x5f, 0x30, 0x57, 0x63, 0xc5, 0x88, 0x63,
	0x46, 0x1e, 0x42, 0x83, 0x09, 0x95, 0x5c, 0xd0, 0x45, 0x18, 0xd3, 0x88, 0x63, 0xc1, 0xb5, 0xa0,
	0xee, 0x74, 0xaf, 0x69, 0xc4, 0x4d, 0x15, 0x74, 0xae, 0x25, 0x67, 0x58, 0xb9, 0x17, 0x38, 0x89,
	0xec, 0x43, 0x7d, 0x3a, 0x57, 0x5a, 0x46, 0x1c, 0xcf, 0xad, 0xe0, 0xb9, 0x90, 0xa9, 0xc6, 0x6c,
	0x6b, 0xf9, 0x77, 0xa0, 0x92, 0x72, 0xaa, 0x64, 0xec, 0x7b, 0x56, 0x6f, 0x25, 0xe2, 0x43, 0x35,
	0xe5, 0x11, 0x4d, 0xcf, 0x95, 0x6b, 0x41, 0x26, 0x62, 0x0a, 0x8c, 0xc9, 0x58, 0xf9, 0x60, 0x77,
	0x58, 0x89, 0x3c, 0x00, 0x98, 0xa6, 0x9c, 0x6a, 0xce, 0x42, 0xaa, 0xfd, 0x3a, 0xda, 0x6a, 0x4e,
	0x33, 0xd2, 0xc6, 0x3c, 0x4f, 0x58, 0x66, 0x6e, 0x58, 0xb3, 0xd3, 0x8c, 0x34, 0xf9, 0x14, 0xea,
	0xb9, 0x2e, 0xf9, 0xcd, 0x5e, 0x61, 0x50, 0x3f, 0xb8, 0x3f, 0x5c, 0x27, 0x74, 0xf8, 0x62, 0xe5,
	0x12, 0xe4, 0xfd, 0xc9, 0x00, 0x76, 0x4c, 0x13, 0xfd, 0x16, 0xee, 0xeb, 0x6e, 0xee, 0x0b, 0x68,
	0x7c, 0x1e, 0xa0, 0x87, 0xb9, 0x5d, 0x11, 0x5f, 0x0a, 0xcd, 0x43, 0x15, 0xfb, 0x6d, 0x7b, 0xbb,
	0x56, 0x31, 0x89, 0xc9, 0x67, 0x50, 0x9b, 0x66, 0x44, 0xfb, 0xbb, 0x18, 0xab, 0xb7, 0x19, 0xcb,
	0xdd, 0xfe, 0x92, 0xfc, 0x60, 0xb5, 0x85, 0x7c, 0x09, 0xad, 0x88, 0x73, 0x1d, 0xae, 0x82, 0x74,
	0x30, 0xc8, 0xa3, 0x2d, 0x41, 0x8e, 0x38, 0xd7, 0xab, 0x40, 0xcd, 0x28, 0x2f, 0xae, 0x73, 0x48,
	0xd6, 0x39, 0xcc, 0x3d, 0x81, 0x5b, 0x6b, 0x4f, 0x60, 0x1f, 0xea, 0xae, 0x3c, 0x44, 0xa8, 0x8b,
	0x46, 0xb0, 0x2a, 0x24, 0x28, 0x8b, 0x8a, 0xe6, 0xdb, 0xab, 0xa8, 0xc6, 0xd8, 0xff, 0xa3, 0x00,
	0x6d, 0x97, 0x5a, 0xc0, 0x55, 0x22, 0x63, 0xc5, 0xc9, 0x87, 0x50, 0xe1, 0xb1, 0x16, 0x7a, 0x81,
	0xa4, 0xd7, 0x0f, 0xee, 0x6e, 0xa9, 0x25, 0x70, 0x6e, 0xe4, 0x03, 0xfb, 0xe4, 0x53, 0xa4, 0xbf,
	0x7e, 0x70, 0x7b, 0xd3, 0xff, 0xd8, 0x18, 0xed, 0x24, 0x48, 0xc9, 0x53, 0x28, 0x0b, 0xcd, 0x23,
	0xe5, 0x97, 0x7a, 0xa5, 0xbf, 0x0a, 0x6e, 0xbd, 0x4c, 0x6c, 0x9e, 0xa6, 0x32, 0xc5, 0xb7, 0x71,
	0x4d, 0xec, 0x97, 0xc6, 0x18, 0x58, 0x1f, 0x03, 0x85, 0x88, 0x4f, 0x25, 0x3e, 0x95, 0x6b, 0xa0,
	0x18, 0xc7, 0xa7, 0x32, 0x40, 0x8f, 0xfe, 0x0f, 0x45, 0xa8, 0x99, 0x93, 0x16, 0x46, 0x47, 0xee,
	0x81, 0x27, 0x54, 0x68, 0x66, 0xa5, 0xad, 0xd9, 0x0b, 0xaa, 0x42, 0xa1, 0x99, 0x3c, 0x83, 0xaa,
	0x9b, 0xa1, 0xae, 0xba, 0xad, 0x09, 0x67, 0x7e, 0xe4, 0x7d, 0x28, 0x9b, 0xfe, 0x66, 0x15, 0x5e,
	0xcf, 0xa6, 0x75, 0x59, 0xe7, 0x6f, 0x07, 0xfd, 0x6f, 0xc8, 0x5f, 0x19, 0x83, 0xfc, 0x1b, 0xfe,
	0xfa, 0x3f, 0x16, 0xa0, 0xb3, 0x6c, 0xca, 0x12, 0x87, 0x67, 0x1b, 0x38, 0xdc, 0xbb, 0x2e, 0xb4,
	0xdd, 0x92, 0x03, 0xc2, 0x5e, 0x5a, 0xf1, 0x1f, 0x5c, 0x5a, 0xe9, 0x6f, 0x2f, 0xed, 0xa7, 0x1d,
	0xd8, 0xdd, 0x6c, 0x46, 0x7e, 0xb8, 0x16, 0xd6, 0x86, 0xab, 0x79, 0xf7, 0x2a, 0x3c, 0x91, 0xf3,
	0xb3, 0x99, 0xc6, 0x44, 0xbc, 0xc0, 0x13, 0xea, 0x39, 0xca, 0xe4, 0x29, 0x90, 0xa9, 0x8c, 0xd5,
	0x3c, 0x4a, 0x4c, 0x90, 0x90, 0x46, 0x72, 0x1e, 0x6b, 0x4c, 0xa1, 0x18, 0x74, 0x72, 0x96, 0x11,
	0x1a, 0xc8, 0x13, 0x68, 0xe7, 0xdd, 0xe3, 0x79, 0x84, 0x3c, 0x96, 0x83, 0x56, 0x4e, 0xfd, 0x7a,
	0x1e, 0x91, 0x47, 0xd0, 0x32, 0x87, 0x8a, 0x98, 0x85, 0x82, 0x4d, 0x69, 0x9a, 0x8d, 0xed, 0x86,
	0x50, 0xcf, 0x45, 0xcc, 0xc6, 0xa8, 0xcb, 0x7b, 0xb9, 0x37, 0x5d, 0xc9, 0x7b, 0x1d, 0xd9, 0x97,
	0xfd, 0x18, 0xbd, 0x4c, 0xe5, 0x69, 0x44, 0xf1, 0x6e, 0xab, 0xe8, 0xd5, 0x14, 0x6a, 0xbc, 0x52,
	0x9a, 0x52, 0x72, 0x3e, 0xe1, 0xa9, 0xe0, 0x17, 0xf8, 0xf3, 0x66, 0x1e, 0x7a, 0x27, 0x67, 0xf9,
	0x1c, 0x0d, 0x66, 0x8e, 0x7c, 0x23, 0x45, 0xcc, 0x19, 0x8e, 0x79, 0x2f, 0x70, 0x92, 0x99, 0x23,
	0x66, 0x95, 0xb5, 0x02, 0xb0, 0x15, 0x60, 0x54, 0xae, 0x07, 0xef, 0xc1, 0x6e, 0x92, 0xca, 0x48,
	0xe6, 0x1b, 0x56, 0x47, 0xaf, 0xf6, 0x52, 0xef, 0x5c, 0xf7, 0xa1, 0x9e, 0xa4, 0x22, 0xa2, 0xe9,
	0x02, 0x5b, 0xd5, 0xc0, 0x56, 0x81, 0x53, 0x99, 0x36, 0x3d, 0x00, 0x50, 0xdc, 0x30, 0x8b, 0xf6,
	0x26, 0xda, 0x6b, 0x56, 0x63, 0xcc, 0xaf, 0xa0, 0x8d, 0x1f, 0x1b, 0x39, 0xac, 0x5b, 0x88, 0xf5,
	0x3b, 0x9b, 0x74, 0xbc, 0x5a, 0xfb, 0x26, 0x09, 0x5a, 0xeb, 0xdf, 0x28, 0xfd, 0x9f, 0x4b, 0xd0,
	0xbd, 0x8e, 0xfc, 0xed, 0xd4, 0x3c, 0x00, 0x98, 0xd1, 0x0d, 0x6c, 0x6a, 0x33, 0x9a, 0x71, 0xf3,
	0x09, 0xdc, 0x31, 0xe6, 0x2d, 0xec, 0x78, 0x41, 0x77, 0x46, 0xcd, 0xf1, 0x1b, 0xf8, 0x0c, 0xe1,
	0xd6, 0xe6, 0xae, 0x0c, 0x21, 0x2f, 0xe8, 0xac, 0x6f, 0x31, 0xf5, 0xbf, 0x0b, 0x6d, 0x4c, 0xe2,
	0x0a, 0x46, 0x4d, 0x93, 0xc9, 0x8a, 0xa3, 0xbc, 0xdf, 0x1a, 0x48, 0x99, 0x9f, 0x23, 0xe9, 0x89,
	0xf5, 0xbb, 0x8a, 0x52, 0x6b, 0x46, 0xd7, 0x58, 0x72, 0xd5, 0x3b, 0x40, 0xbc, 0x65, 0xf5, 0x5f,
	0x58, 0x46, 0x3e, 0x02, 0x53, 0x5f, 0x78, 0x05, 0x03, 0x4b, 0x12, 0x99, 0x51, 0x75, 0xbc, 0x41,
	0x82, 0xcb, 0x30, 0x4f, 0x03, 0x2c, 0x33, 0x3c, 0x5e, 0x01, 0xf1, 0x08, 0x4c, 0x2a, 0x61, 0x0e,
	0x8a, 0xba, 0x7d, 0x11, 0x33, 0xaa, 0x26, 0x19, 0x17, 0x07, 0xbf, 0x17, 0x61, 0xf7, 0x68, 0xe1,
	0x2e, 0x74, 0xc2, 0xd3, 0x4b, 0x31, 0xe5, 0x64, 0x04, 0xe5, 0xc3, 0x19, 0x9f, 0x9e, 0x93, 0xab,
	0x63, 0x26, 0x4a, 0xf4, 0x62, 0x6f, 0x7f, 0xdb, 0xc0, 0x76, 0x03, 0xae, 0xff, 0x3f, 0x72, 0x08,
	0x65, 0x3b, 0xed, 0xb7, 0x84, 0x78, 0xb8, 0x7d, 0xe4, 0xad, 0x82, 0xbc, 0x80, 0xca, 0x64, 0x7e,
	0x12, 0x09, 0x4d, 0xb6, 0xfd, 0x44, 0xbc, 0x4d, 0x2a, 0x23, 0x28, 0xbd, 0xe2, 0x37, 0x0b, 0xf1,
	0x02, 0x2a, 0x2f, 0xbf, 0x17, 0x4a, 0xab, 0x9b, 0x44, 0x39, 0xf8, 0xad, 0x00, 0xad, 0x8d, 0x4e,
	0x1f, 0x66, 0x9d, 0xfe, 0x8f, 0x0b, 0x1c, 0x43, 0x65, 0xc2, 0x69, 0x3a, 0x9d, 0x91, 0xff, 0x6f,
	0x71, 0xc6, 0xbf, 0x22, 0x6f, 0x11, 0xea, 0xa4, 0x82, 0xff, 0x62, 0x3e, 0xfe, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x14, 0xfa, 0x6a, 0x33, 0x29, 0x0d, 0x00, 0x00,
}
