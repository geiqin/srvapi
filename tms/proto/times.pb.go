// Code generated by protoc-gen-go. DO NOT EDIT.
// source: times.proto

package geiqin_srv_tms

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

type TimesWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	TargetId             int64    `protobuf:"varint,5,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimesWhere) Reset()         { *m = TimesWhere{} }
func (m *TimesWhere) String() string { return proto.CompactTextString(m) }
func (*TimesWhere) ProtoMessage()    {}
func (*TimesWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_68bf99941fb82d86, []int{0}
}

func (m *TimesWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimesWhere.Unmarshal(m, b)
}
func (m *TimesWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimesWhere.Marshal(b, m, deterministic)
}
func (m *TimesWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimesWhere.Merge(m, src)
}
func (m *TimesWhere) XXX_Size() int {
	return xxx_messageInfo_TimesWhere.Size(m)
}
func (m *TimesWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_TimesWhere.DiscardUnknown(m)
}

var xxx_messageInfo_TimesWhere proto.InternalMessageInfo

func (m *TimesWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *TimesWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *TimesWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TimesWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *TimesWhere) GetTargetId() int64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *TimesWhere) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Times struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TargetId             int64    `protobuf:"varint,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	StartTime            string   `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string   `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Times) Reset()         { *m = Times{} }
func (m *Times) String() string { return proto.CompactTextString(m) }
func (*Times) ProtoMessage()    {}
func (*Times) Descriptor() ([]byte, []int) {
	return fileDescriptor_68bf99941fb82d86, []int{1}
}

func (m *Times) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Times.Unmarshal(m, b)
}
func (m *Times) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Times.Marshal(b, m, deterministic)
}
func (m *Times) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Times.Merge(m, src)
}
func (m *Times) XXX_Size() int {
	return xxx_messageInfo_Times.Size(m)
}
func (m *Times) XXX_DiscardUnknown() {
	xxx_messageInfo_Times.DiscardUnknown(m)
}

var xxx_messageInfo_Times proto.InternalMessageInfo

func (m *Times) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Times) GetTargetId() int64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *Times) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Times) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Times) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Times) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Times) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type TimesResponse struct {
	Entity               *Times   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Times `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimesResponse) Reset()         { *m = TimesResponse{} }
func (m *TimesResponse) String() string { return proto.CompactTextString(m) }
func (*TimesResponse) ProtoMessage()    {}
func (*TimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68bf99941fb82d86, []int{2}
}

func (m *TimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimesResponse.Unmarshal(m, b)
}
func (m *TimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimesResponse.Marshal(b, m, deterministic)
}
func (m *TimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimesResponse.Merge(m, src)
}
func (m *TimesResponse) XXX_Size() int {
	return xxx_messageInfo_TimesResponse.Size(m)
}
func (m *TimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TimesResponse proto.InternalMessageInfo

func (m *TimesResponse) GetEntity() *Times {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *TimesResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *TimesResponse) GetItems() []*Times {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *TimesResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *TimesResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

type DateList struct {
	Date                 string       `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Type                 int32        `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Times                []*DateTimes `protobuf:"bytes,4,rep,name=times,proto3" json:"times,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DateList) Reset()         { *m = DateList{} }
func (m *DateList) String() string { return proto.CompactTextString(m) }
func (*DateList) ProtoMessage()    {}
func (*DateList) Descriptor() ([]byte, []int) {
	return fileDescriptor_68bf99941fb82d86, []int{3}
}

func (m *DateList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateList.Unmarshal(m, b)
}
func (m *DateList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateList.Marshal(b, m, deterministic)
}
func (m *DateList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateList.Merge(m, src)
}
func (m *DateList) XXX_Size() int {
	return xxx_messageInfo_DateList.Size(m)
}
func (m *DateList) XXX_DiscardUnknown() {
	xxx_messageInfo_DateList.DiscardUnknown(m)
}

var xxx_messageInfo_DateList proto.InternalMessageInfo

func (m *DateList) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *DateList) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *DateList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DateList) GetTimes() []*DateTimes {
	if m != nil {
		return m.Times
	}
	return nil
}

type DateTimes struct {
	StartTime            string   `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string   `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateTimes) Reset()         { *m = DateTimes{} }
func (m *DateTimes) String() string { return proto.CompactTextString(m) }
func (*DateTimes) ProtoMessage()    {}
func (*DateTimes) Descriptor() ([]byte, []int) {
	return fileDescriptor_68bf99941fb82d86, []int{4}
}

func (m *DateTimes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateTimes.Unmarshal(m, b)
}
func (m *DateTimes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateTimes.Marshal(b, m, deterministic)
}
func (m *DateTimes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateTimes.Merge(m, src)
}
func (m *DateTimes) XXX_Size() int {
	return xxx_messageInfo_DateTimes.Size(m)
}
func (m *DateTimes) XXX_DiscardUnknown() {
	xxx_messageInfo_DateTimes.DiscardUnknown(m)
}

var xxx_messageInfo_DateTimes proto.InternalMessageInfo

func (m *DateTimes) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *DateTimes) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

type DateListResponse struct {
	Entity               *DateList   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*DateList `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DateListResponse) Reset()         { *m = DateListResponse{} }
func (m *DateListResponse) String() string { return proto.CompactTextString(m) }
func (*DateListResponse) ProtoMessage()    {}
func (*DateListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68bf99941fb82d86, []int{5}
}

func (m *DateListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateListResponse.Unmarshal(m, b)
}
func (m *DateListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateListResponse.Marshal(b, m, deterministic)
}
func (m *DateListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateListResponse.Merge(m, src)
}
func (m *DateListResponse) XXX_Size() int {
	return xxx_messageInfo_DateListResponse.Size(m)
}
func (m *DateListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DateListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DateListResponse proto.InternalMessageInfo

func (m *DateListResponse) GetEntity() *DateList {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *DateListResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *DateListResponse) GetItems() []*DateList {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *DateListResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DateListResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*TimesWhere)(nil), "geiqin.srv.tms.TimesWhere")
	proto.RegisterType((*Times)(nil), "geiqin.srv.tms.Times")
	proto.RegisterType((*TimesResponse)(nil), "geiqin.srv.tms.TimesResponse")
	proto.RegisterType((*DateList)(nil), "geiqin.srv.tms.DateList")
	proto.RegisterType((*DateTimes)(nil), "geiqin.srv.tms.DateTimes")
	proto.RegisterType((*DateListResponse)(nil), "geiqin.srv.tms.DateListResponse")
}

func init() {
	proto.RegisterFile("times.proto", fileDescriptor_68bf99941fb82d86)
}

var fileDescriptor_68bf99941fb82d86 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x76, 0xf2, 0xd7, 0xcd, 0xd9, 0x5a, 0xca, 0x50, 0x21, 0xad, 0x2c, 0x84, 0x5c, 0x05, 0xc4,
	0x28, 0xeb, 0x0b, 0x58, 0xda, 0xbd, 0x28, 0x78, 0x21, 0xa9, 0xe2, 0xe5, 0x12, 0x77, 0x4e, 0xd7,
	0x01, 0xf3, 0xe3, 0xcc, 0x58, 0x69, 0x1f, 0xc3, 0x07, 0xf0, 0x59, 0x7c, 0x23, 0x6f, 0xbd, 0x94,
	0x39, 0x93, 0x46, 0x76, 0xc9, 0x2a, 0x54, 0xbc, 0x3b, 0x7b, 0xbe, 0x2f, 0xdf, 0xcc, 0x39, 0xdf,
	0xb7, 0x03, 0x53, 0x23, 0x6b, 0xd4, 0x45, 0xa7, 0x5a, 0xd3, 0xf2, 0x83, 0x35, 0xca, 0x4f, 0xb2,
	0x29, 0xb4, 0xba, 0x2e, 0x4c, 0xad, 0x4f, 0xf6, 0x57, 0x6d, 0x5d, 0xb7, 0x8d, 0x43, 0xb3, 0xaf,
	0x0c, 0xe0, 0x8d, 0x65, 0xbf, 0xfb, 0x80, 0x0a, 0xf9, 0x11, 0x84, 0x5d, 0xb5, 0x46, 0x91, 0xb0,
	0x94, 0xe5, 0x61, 0xe9, 0x7e, 0xf0, 0xc7, 0x10, 0xdb, 0x62, 0xa9, 0xe5, 0x2d, 0x26, 0x1e, 0x21,
	0x13, 0xdb, 0xb8, 0x94, 0xb7, 0xc8, 0x0f, 0xc0, 0x93, 0x22, 0xf1, 0x53, 0x96, 0xfb, 0xa5, 0x27,
	0x05, 0x3f, 0x04, 0x5f, 0x0a, 0x9d, 0x04, 0xa9, 0x9f, 0xfb, 0xa5, 0x2d, 0xed, 0xe7, 0xa6, 0x52,
	0x6b, 0x34, 0x4b, 0x29, 0x92, 0x90, 0x88, 0x13, 0xd7, 0xb8, 0x10, 0x9c, 0x43, 0x60, 0x6e, 0x3a,
	0x4c, 0xa2, 0x94, 0xe5, 0x71, 0x49, 0x75, 0xf6, 0x9d, 0x41, 0x48, 0x97, 0xea, 0xc5, 0xd9, 0x20,
	0xbe, 0x21, 0xe5, 0xed, 0x90, 0xf2, 0x7f, 0x4b, 0xf1, 0x19, 0x80, 0x36, 0x95, 0x32, 0x4b, 0xbb,
	0x92, 0x24, 0x20, 0x24, 0xa6, 0x8e, 0x3d, 0x80, 0x1f, 0xc3, 0x04, 0x1b, 0xe1, 0xc0, 0x90, 0xc0,
	0x3d, 0x6c, 0x04, 0x41, 0x33, 0x80, 0x95, 0xc2, 0xca, 0xa0, 0x58, 0x56, 0xa6, 0xbf, 0x5e, 0xdc,
	0x77, 0x4e, 0x8d, 0x85, 0x3f, 0x77, 0xe2, 0x0e, 0xde, 0x73, 0x70, 0xdf, 0x39, 0x35, 0xd9, 0x0f,
	0x06, 0x0f, 0x69, 0x84, 0x12, 0x75, 0xd7, 0x36, 0x1a, 0xf9, 0x53, 0x88, 0xb0, 0x31, 0xd2, 0xdc,
	0xd0, 0x38, 0xd3, 0xf9, 0xa3, 0x62, 0xd3, 0x98, 0xc2, 0xd1, 0x7b, 0x12, 0x7f, 0xe2, 0x9c, 0x50,
	0x34, 0xe5, 0x08, 0xfb, 0xb5, 0x05, 0x9d, 0x41, 0xca, 0x92, 0xa5, 0xc1, 0x5a, 0x27, 0x7e, 0xea,
	0xef, 0x96, 0x76, 0x1c, 0x4b, 0x46, 0xa5, 0x5a, 0x45, 0xdb, 0x18, 0x21, 0x2f, 0x2c, 0x58, 0x3a,
	0x0e, 0xcf, 0x21, 0x90, 0xcd, 0x55, 0x4b, 0xcb, 0x99, 0xce, 0x8f, 0xb6, 0xb9, 0x17, 0xcd, 0x55,
	0x5b, 0x12, 0x23, 0xfb, 0x02, 0x93, 0xf3, 0xca, 0xe0, 0x2b, 0xa9, 0x8d, 0x75, 0xc2, 0x2e, 0x82,
	0x26, 0x8d, 0x4b, 0xaa, 0x07, 0x77, 0x5c, 0x7e, 0x9c, 0x3b, 0x1c, 0x82, 0xa6, 0xaa, 0x07, 0xc7,
	0x6c, 0xcd, 0x9f, 0x41, 0x48, 0xf1, 0xa5, 0x04, 0x4d, 0xe7, 0xc7, 0xdb, 0x47, 0xda, 0x43, 0xfa,
	0x79, 0x88, 0x97, 0x2d, 0x20, 0x1e, 0x7a, 0x5b, 0x7e, 0xb3, 0x3f, 0xf9, 0xed, 0x6d, 0xf8, 0x9d,
	0xfd, 0x64, 0x70, 0x78, 0x37, 0xc0, 0x60, 0xda, 0xf3, 0x2d, 0xd3, 0x92, 0xb1, 0xdb, 0xd0, 0x17,
	0xf7, 0xf2, 0xad, 0xd8, 0xf4, 0x6d, 0xb7, 0xfa, 0x7f, 0xb5, 0x6e, 0xfe, 0xcd, 0x83, 0x7d, 0x5a,
	0xdf, 0x25, 0xaa, 0x6b, 0xb9, 0x42, 0xfe, 0x12, 0xa2, 0x33, 0x4a, 0x3a, 0x1f, 0x8f, 0xd2, 0xc9,
	0x6c, 0x3c, 0x61, 0xfd, 0xda, 0xb2, 0x07, 0x56, 0xe1, 0x2d, 0xfd, 0x19, 0xee, 0xad, 0xb0, 0x80,
	0xe8, 0x1c, 0x3f, 0xa2, 0x41, 0x7e, 0x32, 0x4a, 0xa5, 0x07, 0xeb, 0xef, 0x32, 0x67, 0x10, 0x50,
	0x24, 0xff, 0x45, 0xe4, 0x7d, 0x44, 0x8f, 0xe5, 0x8b, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a,
	0x79, 0x23, 0xe3, 0x59, 0x05, 0x00, 0x00,
}
