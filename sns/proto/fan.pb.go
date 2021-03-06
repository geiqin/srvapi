// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fan.proto

package geiqin_srv_sns

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

type Fan struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FanSn                string   `protobuf:"bytes,2,opt,name=fan_sn,json=fanSn,proto3" json:"fan_sn,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	OpenId               string   `protobuf:"bytes,4,opt,name=open_id,json=openId,proto3" json:"open_id,omitempty"`
	UnionId              string   `protobuf:"bytes,5,opt,name=union_id,json=unionId,proto3" json:"union_id,omitempty"`
	NickName             string   `protobuf:"bytes,6,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,7,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Gender               int32    `protobuf:"varint,8,opt,name=gender,proto3" json:"gender,omitempty"`
	Province             string   `protobuf:"bytes,9,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,10,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,11,opt,name=country,proto3" json:"country,omitempty"`
	Followed             bool     `protobuf:"varint,12,opt,name=followed,proto3" json:"followed,omitempty"`
	FollowedAt           string   `protobuf:"bytes,13,opt,name=followed_at,json=followedAt,proto3" json:"followed_at,omitempty"`
	CustomerId           int64    `protobuf:"varint,14,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Remark               string   `protobuf:"bytes,15,opt,name=remark,proto3" json:"remark,omitempty"`
	Mobile               string   `protobuf:"bytes,16,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email                string   `protobuf:"bytes,17,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt            string   `protobuf:"bytes,18,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,19,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ids                  []int64  `protobuf:"varint,20,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fan) Reset()         { *m = Fan{} }
func (m *Fan) String() string { return proto.CompactTextString(m) }
func (*Fan) ProtoMessage()    {}
func (*Fan) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a953796bd812057, []int{0}
}

func (m *Fan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fan.Unmarshal(m, b)
}
func (m *Fan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fan.Marshal(b, m, deterministic)
}
func (m *Fan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fan.Merge(m, src)
}
func (m *Fan) XXX_Size() int {
	return xxx_messageInfo_Fan.Size(m)
}
func (m *Fan) XXX_DiscardUnknown() {
	xxx_messageInfo_Fan.DiscardUnknown(m)
}

var xxx_messageInfo_Fan proto.InternalMessageInfo

func (m *Fan) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Fan) GetFanSn() string {
	if m != nil {
		return m.FanSn
	}
	return ""
}

func (m *Fan) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Fan) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *Fan) GetUnionId() string {
	if m != nil {
		return m.UnionId
	}
	return ""
}

func (m *Fan) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *Fan) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *Fan) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *Fan) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Fan) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Fan) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Fan) GetFollowed() bool {
	if m != nil {
		return m.Followed
	}
	return false
}

func (m *Fan) GetFollowedAt() string {
	if m != nil {
		return m.FollowedAt
	}
	return ""
}

func (m *Fan) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Fan) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *Fan) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Fan) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Fan) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Fan) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Fan) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

//粉丝查询参数
type FanWhere struct {
	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	//以下为自定义参数
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Gender               int32    `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Mobile               string   `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Source               string   `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	Status               string   `protobuf:"bytes,13,opt,name=status,proto3" json:"status,omitempty"`
	Keywords             string   `protobuf:"bytes,14,opt,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FanWhere) Reset()         { *m = FanWhere{} }
func (m *FanWhere) String() string { return proto.CompactTextString(m) }
func (*FanWhere) ProtoMessage()    {}
func (*FanWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a953796bd812057, []int{1}
}

func (m *FanWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FanWhere.Unmarshal(m, b)
}
func (m *FanWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FanWhere.Marshal(b, m, deterministic)
}
func (m *FanWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FanWhere.Merge(m, src)
}
func (m *FanWhere) XXX_Size() int {
	return xxx_messageInfo_FanWhere.Size(m)
}
func (m *FanWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_FanWhere.DiscardUnknown(m)
}

var xxx_messageInfo_FanWhere proto.InternalMessageInfo

func (m *FanWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *FanWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *FanWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *FanWhere) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FanWhere) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *FanWhere) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *FanWhere) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *FanWhere) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FanWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

//
type FanResponse struct {
	Entity               *Fan     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Fan   `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FanResponse) Reset()         { *m = FanResponse{} }
func (m *FanResponse) String() string { return proto.CompactTextString(m) }
func (*FanResponse) ProtoMessage()    {}
func (*FanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a953796bd812057, []int{2}
}

func (m *FanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FanResponse.Unmarshal(m, b)
}
func (m *FanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FanResponse.Marshal(b, m, deterministic)
}
func (m *FanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FanResponse.Merge(m, src)
}
func (m *FanResponse) XXX_Size() int {
	return xxx_messageInfo_FanResponse.Size(m)
}
func (m *FanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FanResponse proto.InternalMessageInfo

func (m *FanResponse) GetEntity() *Fan {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *FanResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *FanResponse) GetItems() []*Fan {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *FanResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *FanResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Fan)(nil), "geiqin.srv.sns.Fan")
	proto.RegisterType((*FanWhere)(nil), "geiqin.srv.sns.FanWhere")
	proto.RegisterType((*FanResponse)(nil), "geiqin.srv.sns.FanResponse")
}

func init() {
	proto.RegisterFile("fan.proto", fileDescriptor_7a953796bd812057)
}

var fileDescriptor_7a953796bd812057 = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0x13, 0x3b,
	0x14, 0x7d, 0xc9, 0x64, 0x26, 0x93, 0x9b, 0xbe, 0xbe, 0x3e, 0xb7, 0x05, 0xd3, 0x0a, 0x11, 0x65,
	0x15, 0x54, 0x29, 0x8b, 0xb0, 0x42, 0x62, 0x13, 0x09, 0x82, 0x22, 0x01, 0x42, 0x13, 0x55, 0x2c,
	0x23, 0x77, 0xe6, 0x26, 0xb5, 0x3a, 0x63, 0x0f, 0xb6, 0x27, 0x55, 0xfa, 0x01, 0x2c, 0xf8, 0x51,
	0xc4, 0x5f, 0x20, 0xdb, 0x93, 0x96, 0x94, 0xa0, 0x4a, 0xdd, 0xf9, 0x9c, 0x73, 0x7d, 0x72, 0x7d,
	0xee, 0x9d, 0x40, 0x67, 0xc1, 0xc4, 0xb0, 0x54, 0xd2, 0x48, 0xb2, 0xbf, 0x44, 0xfe, 0x95, 0x8b,
	0xa1, 0x56, 0xab, 0xa1, 0x16, 0xfa, 0x64, 0x2f, 0x95, 0x45, 0x21, 0x6b, 0xb5, 0xff, 0xbd, 0x05,
	0xc1, 0x84, 0x09, 0xb2, 0x0f, 0x4d, 0x9e, 0xd1, 0x46, 0xaf, 0x31, 0x08, 0x92, 0x26, 0xcf, 0xc8,
	0x31, 0x44, 0x0b, 0x26, 0xe6, 0x5a, 0xd0, 0x66, 0xaf, 0x31, 0xe8, 0x24, 0xe1, 0x82, 0x89, 0x99,
	0x20, 0x4f, 0x20, 0xd2, 0xb2, 0x52, 0x29, 0xd2, 0xc0, 0xd1, 0x35, 0x22, 0x4f, 0xa1, 0x2d, 0x4b,
	0x14, 0x73, 0x9e, 0xd1, 0x96, 0x17, 0x2c, 0x9c, 0x66, 0xe4, 0x19, 0xc4, 0x95, 0xe0, 0xd2, 0x29,
	0xa1, 0x53, 0xda, 0x0e, 0x4f, 0x33, 0x72, 0x0a, 0x1d, 0xc1, 0xd3, 0xab, 0xb9, 0x60, 0x05, 0xd2,
	0xc8, 0x69, 0xb1, 0x25, 0x3e, 0xb1, 0x02, 0xc9, 0x73, 0x00, 0xb6, 0x62, 0x86, 0xa9, 0x79, 0xa5,
	0x72, 0xda, 0x76, 0x6a, 0xc7, 0x33, 0xe7, 0x2a, 0xb7, 0x7d, 0x2c, 0x51, 0x64, 0xa8, 0x68, 0xdc,
	0x6b, 0x0c, 0xc2, 0xa4, 0x46, 0xe4, 0x04, 0xe2, 0x52, 0xc9, 0x15, 0x17, 0x29, 0xd2, 0x8e, 0xb7,
	0xdc, 0x60, 0x42, 0xa0, 0x95, 0x72, 0xb3, 0xa6, 0xe0, 0x78, 0x77, 0x26, 0x14, 0xda, 0xa9, 0xac,
	0x84, 0x51, 0x6b, 0xda, 0xf5, 0xdd, 0xd5, 0xd0, 0x3a, 0x2d, 0x64, 0x9e, 0xcb, 0x6b, 0xcc, 0xe8,
	0x5e, 0xaf, 0x31, 0x88, 0x93, 0x5b, 0x4c, 0x5e, 0x40, 0x77, 0x73, 0x9e, 0x33, 0x43, 0xff, 0x75,
	0x37, 0x61, 0x43, 0x8d, 0x8d, 0x2d, 0x48, 0x2b, 0x6d, 0x64, 0x81, 0xca, 0x3e, 0x7c, 0xdf, 0xc5,
	0x0a, 0x1b, 0x6a, 0x9a, 0xd9, 0xfe, 0x15, 0x16, 0x4c, 0x5d, 0xd1, 0xff, 0x7c, 0x5c, 0x1e, 0x59,
	0xbe, 0x90, 0x17, 0x3c, 0x47, 0x7a, 0xe0, 0x79, 0x8f, 0xc8, 0x11, 0x84, 0x58, 0x30, 0x9e, 0xd3,
	0xff, 0xfd, 0x34, 0x1c, 0xb0, 0x21, 0xa5, 0x0a, 0x99, 0xf1, 0x6d, 0x10, 0x1f, 0x52, 0xcd, 0x8c,
	0x8d, 0x95, 0xab, 0x32, 0xdb, 0xc8, 0x87, 0x5e, 0xae, 0x99, 0xb1, 0x21, 0x07, 0x10, 0xf0, 0x4c,
	0xd3, 0xa3, 0x5e, 0x30, 0x08, 0x12, 0x7b, 0xec, 0xff, 0x68, 0x40, 0x3c, 0x61, 0xe2, 0xcb, 0x25,
	0x2a, 0xf7, 0x93, 0x25, 0x5b, 0xa2, 0x5f, 0x8a, 0x30, 0xf1, 0xc0, 0x0e, 0xcd, 0x1e, 0xe6, 0x9a,
	0xdf, 0xa0, 0x5b, 0x8d, 0x30, 0x89, 0x2d, 0x31, 0xe3, 0x37, 0x68, 0xd3, 0xd4, 0x52, 0x19, 0x2e,
	0x96, 0xf5, 0x7a, 0x6c, 0xa0, 0xcd, 0xde, 0x8d, 0xd9, 0x2f, 0x87, 0x3b, 0xff, 0x36, 0xc3, 0x70,
	0x6b, 0x86, 0x77, 0x19, 0x44, 0x5b, 0x19, 0xdc, 0xed, 0x5e, 0x7b, 0x6b, 0xf7, 0x2c, 0x6f, 0x98,
	0xa9, 0x74, 0x3d, 0x88, 0x1a, 0xd9, 0x09, 0x5e, 0xe1, 0xfa, 0x5a, 0xaa, 0x4c, 0xbb, 0x09, 0x74,
	0x92, 0x5b, 0xdc, 0xff, 0xd9, 0x80, 0xee, 0x84, 0x89, 0x04, 0x75, 0x29, 0x85, 0x46, 0x72, 0x06,
	0x11, 0x0a, 0x63, 0xb7, 0xc3, 0xbe, 0xb6, 0x3b, 0x3a, 0x1c, 0x6e, 0x7f, 0x35, 0x43, 0x5b, 0x5c,
	0x97, 0x90, 0x33, 0x9f, 0x8c, 0x72, 0xef, 0xef, 0x8e, 0x8e, 0xef, 0xd7, 0x7e, 0xb6, 0xa2, 0x0f,
	0x4c, 0x91, 0x97, 0x10, 0x72, 0x83, 0x85, 0xa6, 0x41, 0x2f, 0xf8, 0x9b, 0xb1, 0xaf, 0xb0, 0xbe,
	0xa8, 0x94, 0x54, 0x2e, 0xa5, 0x1d, 0xbe, 0xef, 0xac, 0x98, 0xf8, 0x1a, 0x32, 0x80, 0x16, 0x17,
	0x0b, 0xe9, 0xb2, 0xeb, 0x8e, 0x8e, 0xee, 0xd7, 0x4e, 0xc5, 0x42, 0x26, 0xae, 0x62, 0xf4, 0x01,
	0xf6, 0x3e, 0xae, 0x27, 0x4c, 0xcc, 0x50, 0xad, 0x78, 0x8a, 0xe4, 0x0d, 0xb4, 0xac, 0x4a, 0xfe,
	0xf4, 0x2f, 0x4a, 0xb3, 0x3e, 0x39, 0xdd, 0xd5, 0x61, 0x9d, 0x53, 0xff, 0x9f, 0xd1, 0xb7, 0x26,
	0xc0, 0x96, 0x59, 0x74, 0xee, 0x36, 0x8a, 0xec, 0x7a, 0xd9, 0x03, 0x66, 0xf6, 0xf6, 0x5b, 0xcc,
	0xf1, 0x91, 0xb7, 0x5f, 0x43, 0xf0, 0x1e, 0xcd, 0xa3, 0xae, 0x8e, 0x21, 0x9a, 0x21, 0x53, 0xe9,
	0x25, 0xa1, 0x3b, 0x0a, 0xdd, 0x07, 0xf0, 0x80, 0xc5, 0x45, 0xe4, 0xfe, 0x40, 0x5f, 0xfd, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0xe7, 0x2e, 0xbf, 0xc6, 0x6b, 0x05, 0x00, 0x00,
}
