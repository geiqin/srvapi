// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consumer.proto

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

type Consumer struct {
	Id                   int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	DistributorId        int64        `protobuf:"varint,2,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id"`
	CustomerId           int64        `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	OrderNum             int32        `protobuf:"varint,4,opt,name=order_num,json=orderNum,proto3" json:"order_num"`
	OrderAmount          float32      `protobuf:"fixed32,5,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount"`
	TotalOrderNum        int32        `protobuf:"varint,6,opt,name=total_order_num,json=totalOrderNum,proto3" json:"total_order_num"`
	TotalOrderAmount     float32      `protobuf:"fixed32,7,opt,name=total_order_amount,json=totalOrderAmount,proto3" json:"total_order_amount"`
	Status               string       `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
	ChangeNum            int32        `protobuf:"varint,9,opt,name=change_num,json=changeNum,proto3" json:"change_num"`
	LastChangeAt         string       `protobuf:"bytes,10,opt,name=last_change_at,json=lastChangeAt,proto3" json:"last_change_at"`
	JoinAt               string       `protobuf:"bytes,11,opt,name=join_at,json=joinAt,proto3" json:"join_at"`
	Mobile               string       `protobuf:"bytes,12,opt,name=mobile,proto3" json:"mobile"`
	ConsumerDisplayName  string       `protobuf:"bytes,13,opt,name=consumer_display_name,json=consumerDisplayName,proto3" json:"consumer_display_name"`
	ConsumerAvatarUrl    string       `protobuf:"bytes,14,opt,name=consumer_avatar_url,json=consumerAvatarUrl,proto3" json:"consumer_avatar_url"`
	CreatedAt            string       `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string       `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Distributor          *Distributor `protobuf:"bytes,17,opt,name=distributor,proto3" json:"distributor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consumer) Reset()         { *m = Consumer{} }
func (m *Consumer) String() string { return proto.CompactTextString(m) }
func (*Consumer) ProtoMessage()    {}
func (*Consumer) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dcb7dada8cc1b, []int{0}
}

func (m *Consumer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consumer.Unmarshal(m, b)
}
func (m *Consumer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consumer.Marshal(b, m, deterministic)
}
func (m *Consumer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consumer.Merge(m, src)
}
func (m *Consumer) XXX_Size() int {
	return xxx_messageInfo_Consumer.Size(m)
}
func (m *Consumer) XXX_DiscardUnknown() {
	xxx_messageInfo_Consumer.DiscardUnknown(m)
}

var xxx_messageInfo_Consumer proto.InternalMessageInfo

func (m *Consumer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Consumer) GetDistributorId() int64 {
	if m != nil {
		return m.DistributorId
	}
	return 0
}

func (m *Consumer) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Consumer) GetOrderNum() int32 {
	if m != nil {
		return m.OrderNum
	}
	return 0
}

func (m *Consumer) GetOrderAmount() float32 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

func (m *Consumer) GetTotalOrderNum() int32 {
	if m != nil {
		return m.TotalOrderNum
	}
	return 0
}

func (m *Consumer) GetTotalOrderAmount() float32 {
	if m != nil {
		return m.TotalOrderAmount
	}
	return 0
}

func (m *Consumer) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Consumer) GetChangeNum() int32 {
	if m != nil {
		return m.ChangeNum
	}
	return 0
}

func (m *Consumer) GetLastChangeAt() string {
	if m != nil {
		return m.LastChangeAt
	}
	return ""
}

func (m *Consumer) GetJoinAt() string {
	if m != nil {
		return m.JoinAt
	}
	return ""
}

func (m *Consumer) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Consumer) GetConsumerDisplayName() string {
	if m != nil {
		return m.ConsumerDisplayName
	}
	return ""
}

func (m *Consumer) GetConsumerAvatarUrl() string {
	if m != nil {
		return m.ConsumerAvatarUrl
	}
	return ""
}

func (m *Consumer) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Consumer) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Consumer) GetDistributor() *Distributor {
	if m != nil {
		return m.Distributor
	}
	return nil
}

type ConsumerWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	StartDate            string   `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              string   `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	CustomerId           int64    `protobuf:"varint,7,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumerWhere) Reset()         { *m = ConsumerWhere{} }
func (m *ConsumerWhere) String() string { return proto.CompactTextString(m) }
func (*ConsumerWhere) ProtoMessage()    {}
func (*ConsumerWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dcb7dada8cc1b, []int{1}
}

func (m *ConsumerWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerWhere.Unmarshal(m, b)
}
func (m *ConsumerWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerWhere.Marshal(b, m, deterministic)
}
func (m *ConsumerWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerWhere.Merge(m, src)
}
func (m *ConsumerWhere) XXX_Size() int {
	return xxx_messageInfo_ConsumerWhere.Size(m)
}
func (m *ConsumerWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerWhere.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerWhere proto.InternalMessageInfo

func (m *ConsumerWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *ConsumerWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ConsumerWhere) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *ConsumerWhere) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ConsumerWhere) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *ConsumerWhere) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *ConsumerWhere) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

type ConsumerResponse struct {
	Entity               *Consumer   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Consumer `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ConsumerResponse) Reset()         { *m = ConsumerResponse{} }
func (m *ConsumerResponse) String() string { return proto.CompactTextString(m) }
func (*ConsumerResponse) ProtoMessage()    {}
func (*ConsumerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dcb7dada8cc1b, []int{2}
}

func (m *ConsumerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerResponse.Unmarshal(m, b)
}
func (m *ConsumerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerResponse.Marshal(b, m, deterministic)
}
func (m *ConsumerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerResponse.Merge(m, src)
}
func (m *ConsumerResponse) XXX_Size() int {
	return xxx_messageInfo_ConsumerResponse.Size(m)
}
func (m *ConsumerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerResponse proto.InternalMessageInfo

func (m *ConsumerResponse) GetEntity() *Consumer {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ConsumerResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *ConsumerResponse) GetItems() []*Consumer {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ConsumerResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ConsumerResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Consumer)(nil), "geiqin.srv.dms.Consumer")
	proto.RegisterType((*ConsumerWhere)(nil), "geiqin.srv.dms.ConsumerWhere")
	proto.RegisterType((*ConsumerResponse)(nil), "geiqin.srv.dms.ConsumerResponse")
}

func init() { proto.RegisterFile("consumer.proto", fileDescriptor_376dcb7dada8cc1b) }

var fileDescriptor_376dcb7dada8cc1b = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0x39, 0xff, 0x33, 0x4e, 0xd2, 0x66, 0x7f, 0x2d, 0x98, 0x56, 0x15, 0x26, 0x02, 0x14,
	0x09, 0x64, 0x21, 0x73, 0xe6, 0x10, 0xb5, 0x1c, 0x2a, 0x44, 0x41, 0xae, 0x10, 0x37, 0xcc, 0x36,
	0x9e, 0xb6, 0x8b, 0xe2, 0xdd, 0xb0, 0xbb, 0xae, 0xd4, 0x7e, 0x47, 0x3e, 0x07, 0x5f, 0x81, 0x23,
	0xf2, 0xac, 0xdd, 0xa4, 0x11, 0xe5, 0xc6, 0xcd, 0xf3, 0xde, 0x9b, 0x37, 0x3b, 0x7f, 0x64, 0x18,
	0xcd, 0x95, 0x34, 0x45, 0x8e, 0x3a, 0x5a, 0x6a, 0x65, 0x15, 0x1b, 0x5d, 0xa0, 0xf8, 0x2e, 0x64,
	0x64, 0xf4, 0x55, 0x94, 0xe5, 0x66, 0x6f, 0x30, 0x57, 0x79, 0xae, 0xa4, 0x63, 0xf7, 0xc6, 0x99,
	0x30, 0x56, 0x8b, 0xb3, 0xc2, 0xaa, 0x2a, 0x61, 0xf2, 0xb3, 0x05, 0xbd, 0xc3, 0xca, 0x83, 0x8d,
	0xa0, 0x21, 0xb2, 0xc0, 0x0b, 0xbd, 0x69, 0x33, 0x69, 0x88, 0x8c, 0x3d, 0x83, 0xd1, 0x5a, 0x46,
	0x2a, 0xb2, 0xa0, 0x41, 0xdc, 0x70, 0x0d, 0x3d, 0xce, 0xd8, 0x63, 0xf0, 0xe7, 0x85, 0xb1, 0x2a,
	0x47, 0xd2, 0x34, 0x49, 0x03, 0x35, 0x74, 0x9c, 0xb1, 0x7d, 0xe8, 0x2b, 0x9d, 0xa1, 0x4e, 0x65,
	0x91, 0x07, 0xad, 0xd0, 0x9b, 0xb6, 0x93, 0x1e, 0x01, 0x27, 0x45, 0xce, 0x9e, 0xc0, 0xc0, 0x91,
	0x3c, 0x57, 0x85, 0xb4, 0x41, 0x3b, 0xf4, 0xa6, 0x8d, 0xc4, 0x27, 0x6c, 0x46, 0x10, 0x7b, 0x0e,
	0x5b, 0x56, 0x59, 0xbe, 0x48, 0x57, 0x2e, 0x1d, 0x72, 0x19, 0x12, 0xfc, 0xa1, 0xb6, 0x7a, 0x09,
	0x6c, 0x5d, 0x57, 0x19, 0x76, 0xc9, 0x70, 0x7b, 0x25, 0xad, 0x5c, 0x1f, 0x40, 0xc7, 0x58, 0x6e,
	0x0b, 0x13, 0xf4, 0x42, 0x6f, 0xda, 0x4f, 0xaa, 0x88, 0x1d, 0x00, 0xcc, 0x2f, 0xb9, 0xbc, 0x40,
	0x2a, 0xd4, 0xa7, 0x42, 0x7d, 0x87, 0x94, 0x45, 0x9e, 0xc2, 0x68, 0xc1, 0x8d, 0x4d, 0x2b, 0x0d,
	0xb7, 0x01, 0x50, 0xfa, 0xa0, 0x44, 0x0f, 0x09, 0x9c, 0x59, 0xf6, 0x10, 0xba, 0xdf, 0x94, 0x90,
	0x25, 0xed, 0x3b, 0xf7, 0x32, 0x9c, 0x51, 0xd5, 0x5c, 0x9d, 0x89, 0x05, 0x06, 0x03, 0x87, 0xbb,
	0x88, 0xc5, 0xb0, 0x5b, 0xef, 0x32, 0xcd, 0x84, 0x59, 0x2e, 0xf8, 0x75, 0x2a, 0x79, 0x8e, 0xc1,
	0x90, 0x64, 0xff, 0xd7, 0xe4, 0x91, 0xe3, 0x4e, 0x78, 0x8e, 0x2c, 0x82, 0x5b, 0x38, 0xe5, 0x57,
	0xdc, 0x72, 0x9d, 0x16, 0x7a, 0x11, 0x8c, 0x28, 0x63, 0x5c, 0x53, 0x33, 0x62, 0x3e, 0xe9, 0x05,
	0x75, 0xa6, 0x91, 0x5b, 0xcc, 0xca, 0x77, 0x6d, 0x91, 0xac, 0x5f, 0x21, 0x33, 0x5b, 0xd2, 0xc5,
	0x32, 0xab, 0xe9, 0x6d, 0x47, 0x57, 0xc8, 0xcc, 0xb2, 0x37, 0xe0, 0xaf, 0xed, 0x3d, 0x18, 0x87,
	0xde, 0xd4, 0x8f, 0xf7, 0xa3, 0xbb, 0x17, 0x17, 0x1d, 0xad, 0x24, 0xc9, 0xba, 0x7e, 0xf2, 0xc3,
	0x83, 0x61, 0x7d, 0x69, 0x9f, 0x2f, 0x51, 0x23, 0xdb, 0x81, 0xf6, 0x92, 0x5f, 0xa0, 0xbb, 0xb8,
	0x76, 0xe2, 0x82, 0xf2, 0x58, 0xca, 0x8f, 0xd4, 0x88, 0x1b, 0xa4, 0x7b, 0x6b, 0x27, 0xbd, 0x12,
	0x38, 0x15, 0x37, 0xb8, 0x36, 0xbd, 0xe6, 0x9d, 0xe9, 0xad, 0x76, 0xd9, 0xda, 0xdc, 0xa5, 0xb1,
	0x5c, 0xdb, 0xb4, 0x6c, 0x82, 0x4e, 0xab, 0x9f, 0xf4, 0x09, 0x39, 0xe2, 0x16, 0xd9, 0x23, 0xe8,
	0xa1, 0xcc, 0x1c, 0xd9, 0x21, 0xb2, 0x8b, 0x32, 0x23, 0x6a, 0xe3, 0xa8, 0xbb, 0x9b, 0x47, 0x3d,
	0xf9, 0xe5, 0xc1, 0x76, 0xdd, 0x4f, 0x82, 0x66, 0xa9, 0xa4, 0x41, 0xf6, 0x0a, 0x3a, 0x28, 0xad,
	0xb0, 0xd7, 0xd4, 0x93, 0x1f, 0x07, 0x9b, 0xe3, 0xb9, 0xcd, 0xa8, 0x74, 0xec, 0x85, 0x1b, 0x82,
	0xa6, 0x56, 0xfd, 0x78, 0x77, 0x33, 0xe1, 0x63, 0x49, 0xba, 0xd9, 0x68, 0x16, 0x41, 0x5b, 0x58,
	0xcc, 0x4d, 0xd0, 0x0c, 0x9b, 0x7f, 0x75, 0x77, 0xb2, 0xd2, 0x1c, 0xb5, 0x56, 0x9a, 0xa6, 0xf2,
	0x07, 0xf3, 0xb7, 0x25, 0x99, 0x38, 0x0d, 0x9b, 0x42, 0x4b, 0xc8, 0x73, 0x45, 0x53, 0xf2, 0xe3,
	0x9d, 0x4d, 0xed, 0xb1, 0x3c, 0x57, 0x09, 0x29, 0xe2, 0xaf, 0x30, 0x7e, 0x7f, 0x5d, 0xd7, 0x3a,
	0x45, 0x7d, 0x25, 0xe6, 0xc8, 0xde, 0x41, 0xe7, 0x14, 0xb9, 0x9e, 0x5f, 0xb2, 0x83, 0xfb, 0x9e,
	0x45, 0x6b, 0xdf, 0x0b, 0xef, 0x7d, 0x75, 0x35, 0xc5, 0xc9, 0x7f, 0xf1, 0x17, 0xd8, 0xfa, 0x97,
	0xfe, 0x67, 0x1d, 0xfa, 0xfb, 0xbd, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x57, 0x18, 0xa8,
	0x40, 0x05, 0x00, 0x00,
}