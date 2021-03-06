// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package geiqin_srv_uim

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
	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
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

func (m *BaseWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
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
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
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

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

//令牌应答
type TokenResponse struct {
	Entity               *Token   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenResponse) Reset()         { *m = TokenResponse{} }
func (m *TokenResponse) String() string { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()    {}
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

func (m *TokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenResponse.Unmarshal(m, b)
}
func (m *TokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenResponse.Marshal(b, m, deterministic)
}
func (m *TokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenResponse.Merge(m, src)
}
func (m *TokenResponse) XXX_Size() int {
	return xxx_messageInfo_TokenResponse.Size(m)
}
func (m *TokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenResponse proto.InternalMessageInfo

func (m *TokenResponse) GetEntity() *Token {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *TokenResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *TokenResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "geiqin.srv.uim.Empty")
	proto.RegisterType((*Info)(nil), "geiqin.srv.uim.Info")
	proto.RegisterType((*Error)(nil), "geiqin.srv.uim.Error")
	proto.RegisterType((*Pager)(nil), "geiqin.srv.uim.Pager")
	proto.RegisterType((*BaseWhere)(nil), "geiqin.srv.uim.BaseWhere")
	proto.RegisterType((*Token)(nil), "geiqin.srv.uim.Token")
	proto.RegisterType((*TokenResponse)(nil), "geiqin.srv.uim.TokenResponse")
}

func init() {
	proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206)
}

var fileDescriptor_555bd8c177793206 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcd, 0x6e, 0x13, 0x3f,
	0x10, 0x57, 0x3e, 0x36, 0xc9, 0x4e, 0xd2, 0xfe, 0xff, 0xb2, 0x5a, 0x58, 0x40, 0x48, 0x61, 0xb9,
	0x44, 0x42, 0xe4, 0x50, 0xc4, 0x0b, 0x00, 0x3d, 0xe4, 0x04, 0x32, 0x48, 0x1c, 0x23, 0x77, 0x3d,
	0x4d, 0xad, 0x66, 0xed, 0xc5, 0x76, 0x8a, 0xd2, 0xf7, 0xe0, 0x21, 0x78, 0x3e, 0x5e, 0x00, 0xcd,
	0xd8, 0x41, 0x6d, 0x05, 0x37, 0xff, 0x3e, 0xec, 0xf9, 0x69, 0xc6, 0x03, 0xb3, 0xc6, 0xb5, 0xad,
	0xb3, 0xcb, 0xce, 0xbb, 0xe8, 0xc4, 0xf1, 0x06, 0xcd, 0x37, 0x63, 0x97, 0xc1, 0xdf, 0x2c, 0x77,
	0xa6, 0xad, 0xc7, 0x50, 0x9c, 0xb7, 0x5d, 0xdc, 0xd7, 0x57, 0x30, 0x5c, 0xd9, 0x4b, 0x27, 0x1e,
	0xc1, 0xc8, 0xee, 0xda, 0x0b, 0xf4, 0x55, 0x6f, 0xde, 0x5b, 0x14, 0x32, 0x23, 0xe2, 0x3b, 0x15,
	0x02, 0xea, 0xaa, 0x3f, 0xef, 0x2d, 0x26, 0x32, 0x23, 0x51, 0xc1, 0xb8, 0x71, 0x36, 0xa2, 0x8d,
	0xd5, 0x60, 0xde, 0x5b, 0x94, 0xf2, 0x00, 0xe9, 0x46, 0x88, 0x2a, 0xee, 0x42, 0x35, 0x64, 0x21,
	0xa3, 0xfa, 0x2d, 0x14, 0xe7, 0xde, 0x3b, 0x2f, 0x04, 0x0c, 0x1b, 0xa7, 0x31, 0x17, 0xe2, 0x33,
	0x3d, 0xd7, 0x62, 0x08, 0x6a, 0x83, 0x5c, 0xa7, 0x94, 0x07, 0x58, 0xff, 0xec, 0x41, 0xf1, 0x49,
	0x6d, 0xd0, 0x8b, 0x13, 0x28, 0x3a, 0xb5, 0x41, 0x9d, 0x2f, 0x26, 0x40, 0x6c, 0x74, 0x51, 0x6d,
	0xf9, 0x5e, 0x21, 0x13, 0x10, 0xcf, 0x01, 0x48, 0x5e, 0x37, 0x6e, 0x97, 0x13, 0x16, 0xb2, 0x24,
	0xe6, 0x3d, 0x11, 0xe2, 0x19, 0x30, 0x58, 0x07, 0x73, 0x8b, 0x1c, 0xb3, 0x90, 0x13, 0x22, 0x3e,
	0x9b, 0x5b, 0x64, 0xd1, 0xe3, 0xcd, 0x9a, 0x88, 0xaa, 0xc8, 0xa2, 0xc7, 0x1b, 0x4a, 0x41, 0xe2,
	0x56, 0x85, 0x98, 0xc4, 0x51, 0x12, 0x89, 0x20, 0xb1, 0xfe, 0xd5, 0x87, 0xf2, 0x9d, 0x0a, 0xf8,
	0xf5, 0x0a, 0x3d, 0xfe, 0x23, 0xef, 0xbd, 0xd2, 0xfd, 0x07, 0xa5, 0x2b, 0x18, 0x07, 0xe7, 0xa3,
	0xb1, 0x9b, 0x43, 0x57, 0x33, 0x14, 0xc7, 0xd0, 0x37, 0x9a, 0xa3, 0x0e, 0x64, 0xdf, 0x68, 0xea,
	0x72, 0xeb, 0x2e, 0xcc, 0x36, 0x25, 0x2c, 0x65, 0x46, 0xd4, 0x5c, 0xab, 0xda, 0x14, 0xad, 0x94,
	0x7c, 0xe6, 0x16, 0x99, 0xb8, 0xc5, 0x6a, 0xcc, 0x64, 0x02, 0xe2, 0x29, 0x4c, 0xae, 0x71, 0xff,
	0xdd, 0x79, 0x1d, 0xaa, 0x09, 0x0b, 0x7f, 0xf0, 0x9d, 0x19, 0x96, 0x77, 0x67, 0x28, 0x1e, 0xc3,
	0x58, 0x79, 0x54, 0x6b, 0xa3, 0x2b, 0xe0, 0x28, 0x23, 0x82, 0x2b, 0x4d, 0x42, 0xdc, 0x77, 0x48,
	0xc2, 0x34, 0xfd, 0x1f, 0x82, 0x2b, 0x2d, 0x4e, 0x61, 0xd4, 0xa8, 0x48, 0xfc, 0x2c, 0x75, 0xa1,
	0x51, 0x71, 0xa5, 0x69, 0x3e, 0x21, 0x2a, 0x1f, 0xd7, 0x5a, 0x45, 0xac, 0x8e, 0xb8, 0x48, 0xc9,
	0xcc, 0x07, 0x15, 0x51, 0x3c, 0x81, 0x09, 0x5a, 0x9d, 0xc4, 0xe3, 0xd4, 0x08, 0xb4, 0x9a, 0xa5,
	0xff, 0x61, 0x60, 0x74, 0xa8, 0xfe, 0x9b, 0x0f, 0x16, 0x03, 0x49, 0xc7, 0xfa, 0x23, 0x14, 0x5f,
	0xdc, 0x35, 0x5a, 0xf1, 0x02, 0x66, 0xaa, 0x69, 0x30, 0x84, 0x75, 0x24, 0xcc, 0x7d, 0x2f, 0xe5,
	0x34, 0x71, 0xc9, 0xf2, 0x12, 0x8e, 0x3c, 0x5e, 0x7a, 0x0c, 0x57, 0xd9, 0x93, 0x7e, 0xdb, 0x2c,
	0x93, 0x6c, 0xaa, 0x7f, 0xf4, 0xe0, 0x88, 0x4f, 0x12, 0x43, 0xe7, 0x6c, 0x40, 0xf1, 0x1a, 0x46,
	0x68, 0xa3, 0x89, 0x7b, 0x7e, 0x73, 0x7a, 0x76, 0xba, 0xbc, 0xbf, 0x4f, 0xcb, 0x64, 0xcf, 0x26,
	0xf1, 0x0a, 0x0a, 0xa4, 0xaf, 0xce, 0xaf, 0xff, 0xc5, 0xcd, 0x7b, 0x20, 0x93, 0x47, 0x2c, 0x60,
	0x68, 0xec, 0xa5, 0xe3, 0x81, 0x4f, 0xcf, 0x4e, 0x1e, 0x7a, 0x69, 0x3b, 0x25, 0x3b, 0x2e, 0x46,
	0xbc, 0xcb, 0x6f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xf0, 0x31, 0x89, 0xdb, 0x03, 0x00,
	0x00,
}
