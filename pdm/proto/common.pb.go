// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Info struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Passed               bool     `protobuf:"varint,2,opt,name=passed,proto3" json:"passed,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Info) GetPassed() bool {
	if m != nil {
		return m.Passed
	}
	return false
}

func (m *Info) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Info) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Pager struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	Total                int32    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	PageCount            int32    `protobuf:"varint,3,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	PageSize             int32    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PrevPage             int32    `protobuf:"varint,5,opt,name=prev_page,json=prevPage,proto3" json:"prev_page,omitempty"`
	LastPage             int32    `protobuf:"varint,6,opt,name=last_page,json=lastPage,proto3" json:"last_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pager) Reset()         { *m = Pager{} }
func (m *Pager) String() string { return proto.CompactTextString(m) }
func (*Pager) ProtoMessage()    {}
func (*Pager) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *Pager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pager.Unmarshal(m, b)
}
func (m *Pager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pager.Marshal(b, m, deterministic)
}
func (m *Pager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pager.Merge(m, src)
}
func (m *Pager) XXX_Size() int {
	return xxx_messageInfo_Pager.Size(m)
}
func (m *Pager) XXX_DiscardUnknown() {
	xxx_messageInfo_Pager.DiscardUnknown(m)
}

var xxx_messageInfo_Pager proto.InternalMessageInfo

func (m *Pager) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *Pager) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Pager) GetPageCount() int32 {
	if m != nil {
		return m.PageCount
	}
	return 0
}

func (m *Pager) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Pager) GetPrevPage() int32 {
	if m != nil {
		return m.PrevPage
	}
	return 0
}

func (m *Pager) GetLastPage() int32 {
	if m != nil {
		return m.LastPage
	}
	return 0
}

type BaseWhere struct {
	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	//base params
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Mobile               string   `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Title                string   `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Keywords             string   `protobuf:"bytes,8,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Status               string   `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	AreaId               int64    `protobuf:"varint,10,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	TypeId               int32    `protobuf:"varint,11,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	CatId                int32    `protobuf:"varint,12,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"`
	StartDate            string   `protobuf:"bytes,13,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              string   `protobuf:"bytes,14,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Ids                  []int64  `protobuf:"varint,15,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseWhere) Reset()         { *m = BaseWhere{} }
func (m *BaseWhere) String() string { return proto.CompactTextString(m) }
func (*BaseWhere) ProtoMessage()    {}
func (*BaseWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *BaseWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseWhere.Unmarshal(m, b)
}
func (m *BaseWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseWhere.Marshal(b, m, deterministic)
}
func (m *BaseWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseWhere.Merge(m, src)
}
func (m *BaseWhere) XXX_Size() int {
	return xxx_messageInfo_BaseWhere.Size(m)
}
func (m *BaseWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseWhere.DiscardUnknown(m)
}

var xxx_messageInfo_BaseWhere proto.InternalMessageInfo

func (m *BaseWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *BaseWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *BaseWhere) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *BaseWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BaseWhere) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *BaseWhere) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BaseWhere) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BaseWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *BaseWhere) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *BaseWhere) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *BaseWhere) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *BaseWhere) GetCatId() int32 {
	if m != nil {
		return m.CatId
	}
	return 0
}

func (m *BaseWhere) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *BaseWhere) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *BaseWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "geiqin.srv.pdm.Empty")
	proto.RegisterType((*Info)(nil), "geiqin.srv.pdm.Info")
	proto.RegisterType((*Error)(nil), "geiqin.srv.pdm.Error")
	proto.RegisterType((*Pager)(nil), "geiqin.srv.pdm.Pager")
	proto.RegisterType((*BaseWhere)(nil), "geiqin.srv.pdm.BaseWhere")
}

func init() {
	proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206)
}

var fileDescriptor_555bd8c177793206 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x95, 0xa4, 0x49, 0x93, 0x61, 0x29, 0x2b, 0x8b, 0x3f, 0x06, 0x84, 0x54, 0xe5, 0xd4,
	0x53, 0x2f, 0x88, 0x17, 0x00, 0xf6, 0xd0, 0x1b, 0x32, 0x07, 0x8e, 0x91, 0x1b, 0x0f, 0x5d, 0x8b,
	0x26, 0x0e, 0xf6, 0xb4, 0xa8, 0xfb, 0x46, 0x3c, 0x1b, 0x2f, 0x81, 0xc6, 0xce, 0xa2, 0x05, 0x89,
	0x9b, 0x7f, 0xdf, 0x4f, 0xce, 0x7c, 0xb1, 0x0d, 0x57, 0xbd, 0x1b, 0x06, 0x37, 0x6e, 0x27, 0xef,
	0xc8, 0x89, 0xd5, 0x01, 0xed, 0x77, 0x3b, 0x6e, 0x83, 0x3f, 0x6f, 0x27, 0x33, 0xb4, 0x4b, 0x28,
	0x6f, 0x86, 0x89, 0x2e, 0xed, 0x2d, 0x2c, 0x76, 0xe3, 0x57, 0x27, 0x9e, 0x43, 0x35, 0x9e, 0x86,
	0x3d, 0x7a, 0x99, 0xad, 0xb3, 0x4d, 0xa9, 0x66, 0xe2, 0x7c, 0xd2, 0x21, 0xa0, 0x91, 0xf9, 0x3a,
	0xdb, 0xd4, 0x6a, 0x26, 0x21, 0x61, 0xd9, 0xbb, 0x91, 0x70, 0x24, 0x59, 0xac, 0xb3, 0x4d, 0xa3,
	0xee, 0x91, 0x77, 0x04, 0xd2, 0x74, 0x0a, 0x72, 0x11, 0xc5, 0x4c, 0xed, 0x3b, 0x28, 0x6f, 0xbc,
	0x77, 0x5e, 0x08, 0x58, 0xf4, 0xce, 0xe0, 0x3c, 0x28, 0xae, 0xf9, 0x73, 0x03, 0x86, 0xa0, 0x0f,
	0x18, 0xe7, 0x34, 0xea, 0x1e, 0xdb, 0x9f, 0x19, 0x94, 0x9f, 0xf4, 0x01, 0xbd, 0x78, 0x0a, 0xe5,
	0xa4, 0x0f, 0x68, 0xe6, 0x8d, 0x09, 0x38, 0x25, 0x47, 0xfa, 0x18, 0xf7, 0x95, 0x2a, 0x81, 0x78,
	0x03, 0xc0, 0xba, 0xeb, 0xdd, 0x69, 0x6e, 0x58, 0xaa, 0x86, 0x93, 0x0f, 0x1c, 0x88, 0xd7, 0x10,
	0xa1, 0x0b, 0xf6, 0x0e, 0x63, 0xcd, 0x52, 0xd5, 0x1c, 0x7c, 0xb6, 0x77, 0x18, 0xa5, 0xc7, 0x73,
	0xc7, 0x81, 0x2c, 0x67, 0xe9, 0xf1, 0xcc, 0x2d, 0x58, 0x1e, 0x75, 0xa0, 0x24, 0xab, 0x24, 0x39,
	0x60, 0xd9, 0xfe, 0xca, 0xa1, 0x79, 0xaf, 0x03, 0x7e, 0xb9, 0x45, 0x8f, 0xff, 0xe9, 0xfb, 0xd7,
	0xe8, 0xfc, 0x9f, 0xd1, 0xd7, 0x50, 0x90, 0x9b, 0xe6, 0xbe, 0xbc, 0x14, 0x2b, 0xc8, 0xad, 0x89,
	0x15, 0x0b, 0x95, 0x5b, 0xc3, 0xa7, 0x3b, 0xb8, 0xbd, 0x3d, 0xa6, 0x66, 0x8d, 0x9a, 0x89, 0x0f,
	0x75, 0xd4, 0x43, 0xaa, 0xd4, 0xa8, 0xb8, 0x8e, 0x47, 0x63, 0xe9, 0x88, 0x72, 0x19, 0xc3, 0x04,
	0xe2, 0x15, 0xd4, 0xdf, 0xf0, 0xf2, 0xc3, 0x79, 0x13, 0x64, 0x1d, 0xc5, 0x1f, 0x7e, 0x70, 0x77,
	0xcd, 0xc3, 0xbb, 0x13, 0x2f, 0x60, 0xa9, 0x3d, 0xea, 0xce, 0x1a, 0x09, 0xb1, 0x4a, 0xc5, 0xb8,
	0x33, 0x2c, 0xe8, 0x32, 0x21, 0x8b, 0x47, 0xe9, 0xdd, 0x30, 0xee, 0x8c, 0x78, 0x06, 0x55, 0xaf,
	0x89, 0xf3, 0xab, 0xf4, 0xf7, 0xbd, 0xa6, 0x9d, 0xe1, 0x7b, 0x09, 0xa4, 0x3d, 0x75, 0x46, 0x13,
	0xca, 0xc7, 0x71, 0x48, 0x13, 0x93, 0x8f, 0x9a, 0x50, 0xbc, 0x84, 0x1a, 0x47, 0x93, 0xe4, 0x2a,
	0xbd, 0x03, 0x1c, 0x4d, 0x54, 0xd7, 0x50, 0x58, 0x13, 0xe4, 0x93, 0x75, 0xb1, 0x29, 0x14, 0x2f,
	0xf7, 0x55, 0x7c, 0xda, 0x6f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x58, 0x95, 0xa6, 0x19, 0xea,
	0x02, 0x00, 0x00,
}
