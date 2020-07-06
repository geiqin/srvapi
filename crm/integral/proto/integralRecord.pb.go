// Code generated by protoc-gen-go. DO NOT EDIT.
// source: integralRecord.proto

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

type IntegralRecordWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting              string   `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Keywords             string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id                   int64    `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	CustomerId           int64    `protobuf:"varint,7,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Status               int32    `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegralRecordWhere) Reset()         { *m = IntegralRecordWhere{} }
func (m *IntegralRecordWhere) String() string { return proto.CompactTextString(m) }
func (*IntegralRecordWhere) ProtoMessage()    {}
func (*IntegralRecordWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_19634b875d7080b9, []int{0}
}

func (m *IntegralRecordWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegralRecordWhere.Unmarshal(m, b)
}
func (m *IntegralRecordWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegralRecordWhere.Marshal(b, m, deterministic)
}
func (m *IntegralRecordWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegralRecordWhere.Merge(m, src)
}
func (m *IntegralRecordWhere) XXX_Size() int {
	return xxx_messageInfo_IntegralRecordWhere.Size(m)
}
func (m *IntegralRecordWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegralRecordWhere.DiscardUnknown(m)
}

var xxx_messageInfo_IntegralRecordWhere proto.InternalMessageInfo

func (m *IntegralRecordWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *IntegralRecordWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *IntegralRecordWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *IntegralRecordWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *IntegralRecordWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IntegralRecordWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *IntegralRecordWhere) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *IntegralRecordWhere) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type IntegralRecord struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           int64    `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	IntegralId           int64    `protobuf:"varint,3,opt,name=integral_id,json=integralId,proto3" json:"integral_id,omitempty"`
	Type                 int32    `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Direction            string   `protobuf:"bytes,5,opt,name=direction,proto3" json:"direction,omitempty"`
	Value                int32    `protobuf:"varint,6,opt,name=value,proto3" json:"value,omitempty"`
	Balance              int32    `protobuf:"varint,7,opt,name=balance,proto3" json:"balance,omitempty"`
	UsedValue            int32    `protobuf:"varint,8,opt,name=used_value,json=usedValue,proto3" json:"used_value,omitempty"`
	EffectiveAt          string   `protobuf:"bytes,9,opt,name=effective_at,json=effectiveAt,proto3" json:"effective_at,omitempty"`
	ExpirationAt         string   `protobuf:"bytes,10,opt,name=expiration_at,json=expirationAt,proto3" json:"expiration_at,omitempty"`
	Memo                 string   `protobuf:"bytes,11,opt,name=memo,proto3" json:"memo,omitempty"`
	Status               int32    `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string   `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegralRecord) Reset()         { *m = IntegralRecord{} }
func (m *IntegralRecord) String() string { return proto.CompactTextString(m) }
func (*IntegralRecord) ProtoMessage()    {}
func (*IntegralRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_19634b875d7080b9, []int{1}
}

func (m *IntegralRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegralRecord.Unmarshal(m, b)
}
func (m *IntegralRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegralRecord.Marshal(b, m, deterministic)
}
func (m *IntegralRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegralRecord.Merge(m, src)
}
func (m *IntegralRecord) XXX_Size() int {
	return xxx_messageInfo_IntegralRecord.Size(m)
}
func (m *IntegralRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegralRecord.DiscardUnknown(m)
}

var xxx_messageInfo_IntegralRecord proto.InternalMessageInfo

func (m *IntegralRecord) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IntegralRecord) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *IntegralRecord) GetIntegralId() int64 {
	if m != nil {
		return m.IntegralId
	}
	return 0
}

func (m *IntegralRecord) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *IntegralRecord) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *IntegralRecord) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *IntegralRecord) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *IntegralRecord) GetUsedValue() int32 {
	if m != nil {
		return m.UsedValue
	}
	return 0
}

func (m *IntegralRecord) GetEffectiveAt() string {
	if m != nil {
		return m.EffectiveAt
	}
	return ""
}

func (m *IntegralRecord) GetExpirationAt() string {
	if m != nil {
		return m.ExpirationAt
	}
	return ""
}

func (m *IntegralRecord) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *IntegralRecord) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *IntegralRecord) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *IntegralRecord) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type IntegralRecordResponse struct {
	Entity               *IntegralRecord   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager            `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*IntegralRecord `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error            `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info             `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IntegralRecordResponse) Reset()         { *m = IntegralRecordResponse{} }
func (m *IntegralRecordResponse) String() string { return proto.CompactTextString(m) }
func (*IntegralRecordResponse) ProtoMessage()    {}
func (*IntegralRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19634b875d7080b9, []int{2}
}

func (m *IntegralRecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegralRecordResponse.Unmarshal(m, b)
}
func (m *IntegralRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegralRecordResponse.Marshal(b, m, deterministic)
}
func (m *IntegralRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegralRecordResponse.Merge(m, src)
}
func (m *IntegralRecordResponse) XXX_Size() int {
	return xxx_messageInfo_IntegralRecordResponse.Size(m)
}
func (m *IntegralRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegralRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntegralRecordResponse proto.InternalMessageInfo

func (m *IntegralRecordResponse) GetEntity() *IntegralRecord {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *IntegralRecordResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *IntegralRecordResponse) GetItems() []*IntegralRecord {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *IntegralRecordResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *IntegralRecordResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*IntegralRecordWhere)(nil), "geiqin.srv.crm.integral.IntegralRecordWhere")
	proto.RegisterType((*IntegralRecord)(nil), "geiqin.srv.crm.integral.IntegralRecord")
	proto.RegisterType((*IntegralRecordResponse)(nil), "geiqin.srv.crm.integral.IntegralRecordResponse")
}

func init() {
	proto.RegisterFile("integralRecord.proto", fileDescriptor_19634b875d7080b9)
}

var fileDescriptor_19634b875d7080b9 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0x39, 0xae, 0xdd, 0x78, 0x9c, 0x46, 0x3f, 0x2d, 0xa5, 0x5d, 0x05, 0x4a, 0x43, 0x38,
	0x90, 0x03, 0x32, 0x22, 0x70, 0x45, 0x28, 0x07, 0x0e, 0x39, 0x20, 0x21, 0x47, 0x82, 0x63, 0xe4,
	0x7a, 0x27, 0xe9, 0x8a, 0xd8, 0x6b, 0x76, 0x37, 0x81, 0xf4, 0x88, 0x78, 0x05, 0x9e, 0x85, 0x67,
	0xe1, 0x6d, 0xd0, 0xee, 0xda, 0x69, 0x53, 0xd4, 0xaa, 0x9c, 0xb8, 0xed, 0x7c, 0xf3, 0xcd, 0x9f,
	0x6f, 0x66, 0x6c, 0x38, 0xe4, 0xa5, 0xc6, 0x85, 0xcc, 0x96, 0x29, 0xe6, 0x42, 0xb2, 0xa4, 0x92,
	0x42, 0x0b, 0x72, 0xbc, 0x40, 0xfe, 0x99, 0x97, 0x89, 0x92, 0xeb, 0x24, 0x97, 0x45, 0xd2, 0x90,
	0x7a, 0x9d, 0x5c, 0x14, 0x85, 0x28, 0x1d, 0x6d, 0xf0, 0xcb, 0x83, 0x7b, 0x93, 0x9d, 0xf8, 0x8f,
	0xe7, 0x28, 0x91, 0x1c, 0x42, 0x50, 0x65, 0x0b, 0x64, 0xd4, 0xeb, 0x7b, 0xc3, 0x20, 0x75, 0x06,
	0x79, 0x00, 0x91, 0x79, 0xcc, 0x14, 0xbf, 0x40, 0xda, 0xb2, 0x9e, 0xb6, 0x01, 0xa6, 0xfc, 0x02,
	0x09, 0x85, 0x7d, 0x25, 0xa4, 0xe6, 0xe5, 0x82, 0xfa, 0x7d, 0x6f, 0x18, 0xa5, 0x8d, 0x49, 0x7a,
	0xd0, 0xfe, 0x84, 0x9b, 0x2f, 0x42, 0x32, 0x45, 0xf7, 0xac, 0x6b, 0x6b, 0x93, 0x2e, 0xb4, 0x38,
	0xa3, 0x41, 0xdf, 0x1b, 0xfa, 0x69, 0x8b, 0x33, 0xf2, 0x3f, 0xf8, 0x9c, 0x29, 0x1a, 0xf6, 0xfd,
	0xa1, 0x9f, 0x9a, 0x27, 0x39, 0x85, 0x38, 0x5f, 0x29, 0x2d, 0x0a, 0x94, 0x33, 0xce, 0xe8, 0xbe,
	0xa5, 0x42, 0x03, 0x4d, 0x18, 0x39, 0x82, 0x50, 0xe9, 0x4c, 0xaf, 0x14, 0x6d, 0xdb, 0x96, 0x6a,
	0x6b, 0xf0, 0xc3, 0x87, 0xee, 0xae, 0xb6, 0xba, 0x9a, 0xb7, 0xad, 0x76, 0x2d, 0x77, 0xeb, 0x8f,
	0xdc, 0xa7, 0x10, 0x37, 0x93, 0x33, 0x04, 0xdf, 0x11, 0x1a, 0x68, 0xc2, 0x08, 0x81, 0x3d, 0xbd,
	0xa9, 0xd0, 0xea, 0x0a, 0x52, 0xfb, 0x26, 0x0f, 0x21, 0x62, 0x5c, 0x62, 0xae, 0xb9, 0x28, 0xad,
	0xb4, 0x28, 0xbd, 0x04, 0xcc, 0x68, 0xd7, 0xd9, 0x72, 0x85, 0x34, 0x74, 0xa3, 0xb5, 0x86, 0x99,
	0xde, 0x59, 0xb6, 0xcc, 0xca, 0x1c, 0xad, 0xc2, 0x20, 0x6d, 0x4c, 0x72, 0x02, 0xb0, 0x52, 0xc8,
	0x66, 0x2e, 0xc8, 0x49, 0x8c, 0x0c, 0xf2, 0xc1, 0x06, 0x3e, 0x86, 0x0e, 0xce, 0xe7, 0x26, 0xf7,
	0x1a, 0x67, 0x99, 0xa6, 0x91, 0xad, 0x17, 0x6f, 0xb1, 0xb1, 0x26, 0x4f, 0xe0, 0x00, 0xbf, 0x56,
	0x5c, 0x66, 0xa6, 0xbe, 0xe1, 0x80, 0xe5, 0x74, 0x2e, 0xc1, 0xb1, 0x36, 0x42, 0x0a, 0x2c, 0x04,
	0x8d, 0xad, 0xcf, 0xbe, 0xaf, 0x4c, 0xb6, 0x73, 0x75, 0xb2, 0xa6, 0xa5, 0x5c, 0x62, 0xa6, 0x91,
	0x99, 0x6c, 0x07, 0x4e, 0x61, 0x8d, 0x8c, 0xb5, 0xed, 0xb8, 0x62, 0x8d, 0xbb, 0xeb, 0xdc, 0x35,
	0x32, 0xd6, 0x83, 0x9f, 0x2d, 0x38, 0xda, 0xdd, 0x4b, 0x8a, 0xaa, 0x12, 0xa5, 0x42, 0xf2, 0x06,
	0x42, 0x2c, 0x35, 0xd7, 0x1b, 0xbb, 0xa3, 0x78, 0xf4, 0x34, 0xb9, 0xe1, 0x8c, 0x93, 0x6b, 0x09,
	0xea, 0x30, 0xf2, 0xca, 0xdd, 0xad, 0xb4, 0xab, 0x8c, 0x47, 0x8f, 0x6e, 0x8c, 0x7f, 0x6f, 0x58,
	0xee, 0xae, 0x25, 0x79, 0x0d, 0x01, 0xd7, 0x58, 0x28, 0xea, 0xf7, 0xfd, 0xbf, 0xa9, 0xea, 0xa2,
	0x4c, 0x51, 0x94, 0x52, 0x48, 0x7b, 0x04, 0xb7, 0x15, 0x7d, 0x6b, 0x58, 0xa9, 0x23, 0x93, 0x17,
	0xb0, 0xc7, 0xcb, 0xb9, 0xb0, 0x07, 0x12, 0x8f, 0x4e, 0x6e, 0xa9, 0x39, 0x17, 0xa9, 0xa5, 0x8e,
	0xbe, 0x79, 0x70, 0x7f, 0xb7, 0x85, 0x29, 0xca, 0x35, 0xcf, 0x91, 0x70, 0x08, 0xa7, 0x98, 0xc9,
	0xfc, 0x9c, 0x3c, 0xbb, 0x63, 0xf3, 0xf6, 0x3b, 0xef, 0x3d, 0xbf, 0xab, 0xd4, 0x7a, 0x43, 0x83,
	0xff, 0x46, 0xdf, 0x3d, 0x38, 0x7e, 0xb7, 0xf9, 0xd7, 0x6d, 0x9c, 0x85, 0xf6, 0x07, 0xf6, 0xf2,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x2d, 0xf2, 0x2e, 0xff, 0x04, 0x00, 0x00,
}