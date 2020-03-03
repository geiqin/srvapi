// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package geiqin_srv_crm

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
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	Total                int32    `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
	PageCount            int32    `protobuf:"varint,3,opt,name=page_count,json=pageCount,proto3" json:"page_count"`
	PageSize             int32    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	PrevPage             int32    `protobuf:"varint,5,opt,name=prev_page,json=prevPage,proto3" json:"prev_page"`
	LastPage             int32    `protobuf:"varint,6,opt,name=last_page,json=lastPage,proto3" json:"last_page"`
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

//令牌
type Token struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Error                *Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Token) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "geiqin.srv.crm.Empty")
	proto.RegisterType((*Info)(nil), "geiqin.srv.crm.Info")
	proto.RegisterType((*Error)(nil), "geiqin.srv.crm.Error")
	proto.RegisterType((*Pager)(nil), "geiqin.srv.crm.Pager")
	proto.RegisterType((*BaseWhere)(nil), "geiqin.srv.crm.BaseWhere")
	proto.RegisterType((*Token)(nil), "geiqin.srv.crm.Token")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x8e, 0xd3, 0x3e,
	0x10, 0xc6, 0x95, 0xa6, 0x6e, 0x9b, 0x69, 0xff, 0xfd, 0xaf, 0xac, 0x2e, 0x04, 0x10, 0x52, 0xc9,
	0xa9, 0x12, 0x52, 0x0f, 0x8b, 0x78, 0x01, 0x60, 0x0f, 0xbd, 0x21, 0x83, 0xc4, 0x31, 0x72, 0xe3,
	0xd9, 0xae, 0xb5, 0x8d, 0x1d, 0x6c, 0xb7, 0xa8, 0xfb, 0x22, 0x3c, 0x03, 0xcf, 0xc6, 0x4b, 0xa0,
	0xb1, 0x53, 0xb4, 0xbb, 0x12, 0x37, 0xff, 0xbe, 0xcf, 0xce, 0x7c, 0x19, 0x8f, 0x61, 0xd6, 0xd8,
	0xb6, 0xb5, 0x66, 0xdd, 0x39, 0x1b, 0x2c, 0x9f, 0xef, 0x50, 0x7f, 0xd7, 0x66, 0xed, 0xdd, 0x71,
	0xdd, 0xb8, 0xb6, 0x1a, 0x03, 0xbb, 0x6e, 0xbb, 0x70, 0xaa, 0x6e, 0x61, 0xb8, 0x31, 0x37, 0x96,
	0x3f, 0x83, 0x91, 0x39, 0xb4, 0x5b, 0x74, 0x65, 0xb6, 0xcc, 0x56, 0x4c, 0xf4, 0x44, 0x7a, 0x27,
	0xbd, 0x47, 0x55, 0x0e, 0x96, 0xd9, 0x6a, 0x22, 0x7a, 0xe2, 0x25, 0x8c, 0x1b, 0x6b, 0x02, 0x9a,
	0x50, 0xe6, 0xcb, 0x6c, 0x55, 0x88, 0x33, 0xd2, 0x09, 0x1f, 0x64, 0x38, 0xf8, 0x72, 0x18, 0x8d,
	0x9e, 0xaa, 0xf7, 0xc0, 0xae, 0x9d, 0xb3, 0x8e, 0x73, 0x18, 0x36, 0x56, 0x61, 0x5f, 0x28, 0xae,
	0xe9, 0x73, 0x2d, 0x7a, 0x2f, 0x77, 0x18, 0xeb, 0x14, 0xe2, 0x8c, 0xd5, 0xaf, 0x0c, 0xd8, 0x67,
	0xb9, 0x43, 0xc7, 0x17, 0xc0, 0x3a, 0xb9, 0x43, 0xd5, 0x1f, 0x4c, 0x40, 0x6a, 0xb0, 0x41, 0xee,
	0xe3, 0x39, 0x26, 0x12, 0xf0, 0xd7, 0x00, 0x64, 0xd7, 0x8d, 0x3d, 0xf4, 0x09, 0x99, 0x28, 0x48,
	0xf9, 0x48, 0x02, 0x7f, 0x05, 0x11, 0x6a, 0xaf, 0xef, 0x31, 0xc6, 0x64, 0x62, 0x42, 0xc2, 0x17,
	0x7d, 0x8f, 0xd1, 0x74, 0x78, 0xac, 0x49, 0x28, 0x59, 0x6f, 0x3a, 0x3c, 0x52, 0x0a, 0x32, 0xf7,
	0xd2, 0x87, 0x64, 0x8e, 0x92, 0x49, 0x02, 0x99, 0xd5, 0xef, 0x01, 0x14, 0x1f, 0xa4, 0xc7, 0x6f,
	0xb7, 0xe8, 0xf0, 0x1f, 0x79, 0x1f, 0x95, 0x1e, 0x3c, 0x29, 0x7d, 0x01, 0x79, 0xb0, 0x5d, 0x9f,
	0x97, 0x96, 0x7c, 0x0e, 0x03, 0xad, 0x62, 0xc4, 0x5c, 0x0c, 0xb4, 0xa2, 0xee, 0xb6, 0x76, 0xab,
	0xf7, 0x29, 0x59, 0x21, 0x7a, 0xa2, 0xa6, 0x1a, 0xd9, 0xa6, 0x48, 0x85, 0x88, 0xeb, 0xd8, 0x1a,
	0x1d, 0xf6, 0x58, 0x8e, 0xa3, 0x98, 0x80, 0xbf, 0x84, 0xc9, 0x1d, 0x9e, 0x7e, 0x58, 0xa7, 0x7c,
	0x39, 0x89, 0xc6, 0x5f, 0x7e, 0x70, 0x77, 0xc5, 0xc3, 0xbb, 0xe3, 0xcf, 0x61, 0x2c, 0x1d, 0xca,
	0x5a, 0xab, 0x12, 0x62, 0x94, 0x11, 0xe1, 0x46, 0x91, 0x11, 0x4e, 0x1d, 0x92, 0x31, 0x4d, 0x73,
	0x43, 0xb8, 0x51, 0xfc, 0x12, 0x46, 0x8d, 0x0c, 0xa4, 0xcf, 0xd2, 0xdf, 0x37, 0x32, 0x6c, 0x14,
	0xdd, 0x8b, 0x0f, 0xd2, 0x85, 0x5a, 0xc9, 0x80, 0xe5, 0x7f, 0xb1, 0x48, 0x11, 0x95, 0x4f, 0x32,
	0x20, 0x7f, 0x01, 0x13, 0x34, 0x2a, 0x99, 0xf3, 0x34, 0x07, 0x68, 0x54, 0xb4, 0x2e, 0x20, 0xd7,
	0xca, 0x97, 0xff, 0x2f, 0xf3, 0x55, 0x2e, 0x68, 0x59, 0xfd, 0xcc, 0x80, 0x7d, 0xb5, 0x77, 0x68,
	0xf8, 0x1b, 0x98, 0xc9, 0xa6, 0x41, 0xef, 0xeb, 0x40, 0x1c, 0x1b, 0x5e, 0x88, 0x69, 0xd2, 0xd2,
	0x96, 0x05, 0xb0, 0xa3, 0xdc, 0xeb, 0xf3, 0x18, 0x27, 0xe0, 0x6f, 0x81, 0x21, 0xcd, 0x64, 0xec,
	0xf8, 0xf4, 0xea, 0x72, 0xfd, 0xf8, 0x99, 0xac, 0xe3, 0xc0, 0x8a, 0xb4, 0x87, 0xaf, 0x60, 0xa8,
	0xcd, 0x8d, 0x8d, 0x8d, 0x9f, 0x5e, 0x2d, 0x9e, 0xee, 0xa5, 0x67, 0x24, 0xe2, 0x8e, 0xed, 0x28,
	0x3e, 0xba, 0x77, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x12, 0xa3, 0x06, 0x84, 0x03, 0x00,
	0x00,
}
