// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderFood.proto

package geiqin_srv_eat

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

type OrderFoodWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top                  int32    `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	OrderId              int64    `protobuf:"varint,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	TableId              int64    `protobuf:"varint,7,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	Type                 string   `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderFoodWhere) Reset()         { *m = OrderFoodWhere{} }
func (m *OrderFoodWhere) String() string { return proto.CompactTextString(m) }
func (*OrderFoodWhere) ProtoMessage()    {}
func (*OrderFoodWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_b25400c348127c6a, []int{0}
}

func (m *OrderFoodWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderFoodWhere.Unmarshal(m, b)
}
func (m *OrderFoodWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderFoodWhere.Marshal(b, m, deterministic)
}
func (m *OrderFoodWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderFoodWhere.Merge(m, src)
}
func (m *OrderFoodWhere) XXX_Size() int {
	return xxx_messageInfo_OrderFoodWhere.Size(m)
}
func (m *OrderFoodWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderFoodWhere.DiscardUnknown(m)
}

var xxx_messageInfo_OrderFoodWhere proto.InternalMessageInfo

func (m *OrderFoodWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *OrderFoodWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *OrderFoodWhere) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *OrderFoodWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderFoodWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *OrderFoodWhere) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderFoodWhere) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *OrderFoodWhere) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type OrderFood struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	OrderId              int64    `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	TableId              int64    `protobuf:"varint,4,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	TableName            string   `protobuf:"bytes,5,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	SurchargeType        int32    `protobuf:"varint,6,opt,name=surcharge_type,json=surchargeType,proto3" json:"surcharge_type,omitempty"`
	SurchargeFee         float32  `protobuf:"fixed32,7,opt,name=surcharge_fee,json=surchargeFee,proto3" json:"surcharge_fee,omitempty"`
	SurchargeFeeName     string   `protobuf:"bytes,8,opt,name=surcharge_fee_name,json=surchargeFeeName,proto3" json:"surcharge_fee_name,omitempty"`
	TotalSurchargeFee    float32  `protobuf:"fixed32,9,opt,name=total_surcharge_fee,json=totalSurchargeFee,proto3" json:"total_surcharge_fee,omitempty"`
	BoxFee               float32  `protobuf:"fixed32,10,opt,name=box_fee,json=boxFee,proto3" json:"box_fee,omitempty"`
	GuestNum             int32    `protobuf:"varint,11,opt,name=guest_num,json=guestNum,proto3" json:"guest_num,omitempty"`
	IsSettlement         bool     `protobuf:"varint,12,opt,name=is_settlement,json=isSettlement,proto3" json:"is_settlement,omitempty"`
	CreatedAt            string   `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderFood) Reset()         { *m = OrderFood{} }
func (m *OrderFood) String() string { return proto.CompactTextString(m) }
func (*OrderFood) ProtoMessage()    {}
func (*OrderFood) Descriptor() ([]byte, []int) {
	return fileDescriptor_b25400c348127c6a, []int{1}
}

func (m *OrderFood) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderFood.Unmarshal(m, b)
}
func (m *OrderFood) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderFood.Marshal(b, m, deterministic)
}
func (m *OrderFood) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderFood.Merge(m, src)
}
func (m *OrderFood) XXX_Size() int {
	return xxx_messageInfo_OrderFood.Size(m)
}
func (m *OrderFood) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderFood.DiscardUnknown(m)
}

var xxx_messageInfo_OrderFood proto.InternalMessageInfo

func (m *OrderFood) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderFood) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OrderFood) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderFood) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *OrderFood) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *OrderFood) GetSurchargeType() int32 {
	if m != nil {
		return m.SurchargeType
	}
	return 0
}

func (m *OrderFood) GetSurchargeFee() float32 {
	if m != nil {
		return m.SurchargeFee
	}
	return 0
}

func (m *OrderFood) GetSurchargeFeeName() string {
	if m != nil {
		return m.SurchargeFeeName
	}
	return ""
}

func (m *OrderFood) GetTotalSurchargeFee() float32 {
	if m != nil {
		return m.TotalSurchargeFee
	}
	return 0
}

func (m *OrderFood) GetBoxFee() float32 {
	if m != nil {
		return m.BoxFee
	}
	return 0
}

func (m *OrderFood) GetGuestNum() int32 {
	if m != nil {
		return m.GuestNum
	}
	return 0
}

func (m *OrderFood) GetIsSettlement() bool {
	if m != nil {
		return m.IsSettlement
	}
	return false
}

func (m *OrderFood) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *OrderFood) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type OrderFoodResponse struct {
	Entity               *OrderFood   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*OrderFood `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info        `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OrderFoodResponse) Reset()         { *m = OrderFoodResponse{} }
func (m *OrderFoodResponse) String() string { return proto.CompactTextString(m) }
func (*OrderFoodResponse) ProtoMessage()    {}
func (*OrderFoodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b25400c348127c6a, []int{2}
}

func (m *OrderFoodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderFoodResponse.Unmarshal(m, b)
}
func (m *OrderFoodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderFoodResponse.Marshal(b, m, deterministic)
}
func (m *OrderFoodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderFoodResponse.Merge(m, src)
}
func (m *OrderFoodResponse) XXX_Size() int {
	return xxx_messageInfo_OrderFoodResponse.Size(m)
}
func (m *OrderFoodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderFoodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderFoodResponse proto.InternalMessageInfo

func (m *OrderFoodResponse) GetEntity() *OrderFood {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *OrderFoodResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *OrderFoodResponse) GetItems() []*OrderFood {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *OrderFoodResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *OrderFoodResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderFoodWhere)(nil), "geiqin.srv.eat.OrderFoodWhere")
	proto.RegisterType((*OrderFood)(nil), "geiqin.srv.eat.OrderFood")
	proto.RegisterType((*OrderFoodResponse)(nil), "geiqin.srv.eat.OrderFoodResponse")
}

func init() {
	proto.RegisterFile("orderFood.proto", fileDescriptor_b25400c348127c6a)
}

var fileDescriptor_b25400c348127c6a = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0xfd, 0xd2, 0xb4, 0x69, 0x73, 0xdb, 0xf5, 0xeb, 0xcc, 0x10, 0xde, 0x26, 0x50, 0x28, 0x42,
	0x8a, 0x04, 0x0a, 0xa2, 0xfc, 0x82, 0x09, 0x31, 0x28, 0x0f, 0x03, 0xa5, 0x48, 0x3c, 0x46, 0x69,
	0x73, 0xdb, 0x59, 0x6a, 0xe2, 0xe0, 0xb8, 0xd3, 0xb6, 0x5f, 0xc4, 0xdf, 0xe0, 0x89, 0x1f, 0xc5,
	0x0b, 0xf2, 0x75, 0x9b, 0xae, 0x13, 0x45, 0xbc, 0xd9, 0xe7, 0x9c, 0x7b, 0x8e, 0x7d, 0x7d, 0x13,
	0xf8, 0x5f, 0xaa, 0x0c, 0xd5, 0xb9, 0x94, 0x59, 0x54, 0x2a, 0xa9, 0x25, 0xeb, 0x2f, 0x50, 0x7c,
	0x13, 0x45, 0x54, 0xa9, 0xab, 0x08, 0x53, 0x7d, 0xd2, 0x9b, 0xc9, 0x3c, 0x97, 0x85, 0x65, 0x87,
	0x3f, 0x1c, 0xe8, 0x7f, 0xda, 0x54, 0x7c, 0xbd, 0x44, 0x85, 0xec, 0x08, 0x5a, 0x65, 0xba, 0xc0,
	0x8c, 0x3b, 0x81, 0x13, 0xb6, 0x62, 0xbb, 0x61, 0xa7, 0xe0, 0x9b, 0x45, 0x52, 0x89, 0x5b, 0xe4,
	0x0d, 0x62, 0x3a, 0x06, 0x98, 0x88, 0x5b, 0x64, 0x03, 0x70, 0xb5, 0x2c, 0xb9, 0x4b, 0xb0, 0x59,
	0xb2, 0x3e, 0x34, 0x44, 0xc6, 0x9b, 0x81, 0x13, 0xba, 0x71, 0x43, 0x64, 0x46, 0x21, 0xb2, 0x8a,
	0xb7, 0x02, 0x37, 0x74, 0x63, 0xb3, 0x64, 0xc7, 0xd0, 0xa1, 0xa3, 0x26, 0x22, 0xe3, 0x1e, 0xe9,
	0xda, 0xb4, 0x1f, 0x67, 0x86, 0xd2, 0xe9, 0x74, 0x89, 0x86, 0x6a, 0x5b, 0x8a, 0xf6, 0xe3, 0x8c,
	0x31, 0x68, 0xea, 0x9b, 0x12, 0x79, 0x27, 0x70, 0x42, 0x3f, 0xa6, 0xf5, 0xf0, 0xa7, 0x0b, 0x7e,
	0x7d, 0x87, 0x75, 0xb2, 0x53, 0x27, 0x6f, 0x2a, 0x1a, 0xdb, 0x8a, 0x9d, 0x6c, 0x77, 0x7f, 0x76,
	0x73, 0x37, 0xfb, 0x31, 0x80, 0xa5, 0x8a, 0x34, 0x47, 0xde, 0x22, 0x3f, 0x9f, 0x90, 0x8b, 0x34,
	0x47, 0xf6, 0x1c, 0xfa, 0xd5, 0x4a, 0xcd, 0x2e, 0x53, 0xb5, 0xc0, 0x84, 0x22, 0x3d, 0xea, 0xc7,
	0x41, 0x8d, 0x7e, 0x31, 0xd9, 0xcf, 0x60, 0x0b, 0x24, 0x73, 0x44, 0xba, 0x61, 0x23, 0xee, 0xd5,
	0xe0, 0x39, 0x22, 0x7b, 0x09, 0x6c, 0x47, 0x64, 0x23, 0xed, 0xa5, 0x07, 0x77, 0x95, 0x94, 0x1c,
	0xc1, 0x03, 0x2d, 0x75, 0xba, 0x4c, 0x76, 0x8d, 0x7d, 0x32, 0x3e, 0x24, 0x6a, 0x72, 0xd7, 0xfd,
	0x11, 0xb4, 0xa7, 0xf2, 0x9a, 0x34, 0x40, 0x1a, 0x6f, 0x2a, 0xaf, 0x0d, 0x71, 0x0a, 0xfe, 0x62,
	0x85, 0x95, 0x4e, 0x8a, 0x55, 0xce, 0xbb, 0xf6, 0x91, 0x09, 0xb8, 0x58, 0xe5, 0xe6, 0xe0, 0xa2,
	0x4a, 0x2a, 0xd4, 0x7a, 0x89, 0x39, 0x16, 0x9a, 0xf7, 0x02, 0x27, 0xec, 0xc4, 0x3d, 0x51, 0x4d,
	0x6a, 0xcc, 0xf4, 0x68, 0xa6, 0x30, 0xd5, 0x98, 0x25, 0xa9, 0xe6, 0x07, 0xb6, 0x47, 0x6b, 0xe4,
	0x8c, 0xe8, 0x55, 0x99, 0x6d, 0xe8, 0xbe, 0xa5, 0xd7, 0xc8, 0x99, 0x1e, 0xfe, 0x72, 0xe0, 0xb0,
	0x7e, 0xc9, 0x18, 0xab, 0x52, 0x16, 0x15, 0xb2, 0xd7, 0xe0, 0x61, 0xa1, 0x85, 0xbe, 0xa1, 0x57,
	0xed, 0x8e, 0x8e, 0xa3, 0xdd, 0x91, 0x8e, 0xb6, 0x25, 0x6b, 0x21, 0x7b, 0x61, 0x67, 0x58, 0xd1,
	0xab, 0x77, 0x47, 0x0f, 0xef, 0x57, 0x7c, 0x36, 0xa4, 0x1d, 0x6d, 0xc5, 0x5e, 0x41, 0x4b, 0x68,
	0xcc, 0x2b, 0xee, 0x06, 0xee, 0xdf, 0xed, 0xad, 0xce, 0xb8, 0xa3, 0x52, 0x52, 0xd1, 0x80, 0xfc,
	0xc1, 0xfd, 0x9d, 0x21, 0x63, 0xab, 0x61, 0x21, 0x34, 0x45, 0x31, 0x97, 0x34, 0x2f, 0xdd, 0xd1,
	0xd1, 0x7d, 0xed, 0xb8, 0x98, 0xcb, 0x98, 0x14, 0xa3, 0xef, 0x0e, 0x0c, 0xea, 0xac, 0x09, 0xaa,
	0x2b, 0x31, 0x43, 0xf6, 0x01, 0xbc, 0xb7, 0xd4, 0x3e, 0xb6, 0xff, 0x5c, 0x27, 0x4f, 0xf7, 0x1f,
	0x79, 0xdd, 0xc4, 0xe1, 0x7f, 0xec, 0x23, 0xb8, 0xef, 0x51, 0xb3, 0x27, 0x7b, 0xb5, 0xf4, 0xf9,
	0xff, 0x93, 0xd7, 0xd4, 0xa3, 0xbf, 0xc7, 0x9b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x9c,
	0x35, 0xa9, 0x6e, 0x04, 0x00, 0x00,
}
