// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

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

// 用户信息
type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string   `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	AvatarId             string   `protobuf:"bytes,4,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Mobile               string   `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Gender               string   `protobuf:"bytes,9,opt,name=gender,proto3" json:"gender,omitempty"`
	CreatedAt            string   `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Profile              *Profile `protobuf:"bytes,12,opt,name=profile,proto3" json:"profile,omitempty"`
	Roles                []*Role  `protobuf:"bytes,13,rep,name=roles,proto3" json:"roles,omitempty" gorm:"many2many:user_roles;"`
	Ids                  []int64  `protobuf:"varint,14,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetAvatarId() string {
	if m != nil {
		return m.AvatarId
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *User) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *User) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *User) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *User) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

// 用户信息
type UserWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting              string   `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string   `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Mobile               string   `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Gender               string   `protobuf:"bytes,8,opt,name=gender,proto3" json:"gender,omitempty"`
	Password             string   `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserWhere) Reset()         { *m = UserWhere{} }
func (m *UserWhere) String() string { return proto.CompactTextString(m) }
func (*UserWhere) ProtoMessage()    {}
func (*UserWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserWhere.Unmarshal(m, b)
}
func (m *UserWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserWhere.Marshal(b, m, deterministic)
}
func (m *UserWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserWhere.Merge(m, src)
}
func (m *UserWhere) XXX_Size() int {
	return xxx_messageInfo_UserWhere.Size(m)
}
func (m *UserWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_UserWhere.DiscardUnknown(m)
}

var xxx_messageInfo_UserWhere proto.InternalMessageInfo

func (m *UserWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *UserWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *UserWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *UserWhere) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserWhere) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *UserWhere) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserWhere) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserWhere) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserWhere) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserResponse struct {
	Entity               *User    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*User  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetEntity() *User {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *UserResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *UserResponse) GetItems() []*User {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *UserResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *UserResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "geiqin.srv.uim.User")
	proto.RegisterType((*UserWhere)(nil), "geiqin.srv.uim.UserWhere")
	proto.RegisterType((*UserResponse)(nil), "geiqin.srv.uim.UserResponse")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0xdf, 0x2c, 0x7f, 0xda, 0x9e, 0x74, 0xd3, 0x2b, 0x6b, 0x80, 0x29, 0x4c, 0x2a, 0xbd,
	0x8a, 0x00, 0x55, 0xa2, 0x5c, 0x22, 0x90, 0xa6, 0x81, 0xa6, 0x49, 0x08, 0xa1, 0x8c, 0xc1, 0x65,
	0xe5, 0x35, 0x67, 0x99, 0x45, 0x12, 0x07, 0xdb, 0x1d, 0xea, 0x3e, 0x0a, 0xe2, 0x23, 0x72, 0xc3,
	0x37, 0x40, 0xb6, 0xd3, 0x69, 0x5d, 0xc3, 0x86, 0x7a, 0x97, 0x73, 0x9e, 0xc7, 0xc7, 0x27, 0xcf,
	0x2f, 0x2d, 0xc0, 0x5c, 0xa1, 0x1c, 0xd7, 0x52, 0x68, 0x41, 0x76, 0x72, 0xe4, 0xdf, 0x78, 0x35,
	0x56, 0xf2, 0x62, 0x3c, 0xe7, 0xe5, 0x60, 0xbb, 0x96, 0xe2, 0x8c, 0x17, 0xe8, 0xe4, 0x41, 0x7f,
	0x26, 0xca, 0x52, 0x54, 0x4d, 0x05, 0x52, 0x2c, 0x95, 0xd1, 0x0f, 0x1f, 0x82, 0x13, 0x85, 0x92,
	0xec, 0xc0, 0x16, 0xcf, 0xa8, 0x37, 0xf4, 0x12, 0x3f, 0xdd, 0xe2, 0x19, 0x21, 0x10, 0x54, 0xac,
	0x44, 0xba, 0x35, 0xf4, 0x92, 0x5e, 0x6a, 0x9f, 0xc9, 0x13, 0xe8, 0x67, 0x5c, 0xd5, 0x05, 0x5b,
	0x4c, 0xad, 0xe6, 0x5b, 0x2d, 0x6e, 0x7a, 0x1f, 0x8c, 0xe5, 0x11, 0xf4, 0xd8, 0x05, 0xd3, 0x4c,
	0x4e, 0x79, 0x46, 0x03, 0xab, 0x77, 0x5d, 0xe3, 0x28, 0x23, 0x7b, 0x00, 0x8d, 0x38, 0x97, 0x05,
	0x0d, 0xad, 0xda, 0xd8, 0x4f, 0x64, 0x41, 0x06, 0xd0, 0xad, 0x99, 0x52, 0xdf, 0x85, 0xcc, 0x68,
	0xe4, 0x8e, 0x2e, 0x6b, 0xb2, 0x0b, 0x21, 0x96, 0x8c, 0x17, 0xb4, 0x63, 0x05, 0x57, 0x90, 0xfb,
	0x10, 0x95, 0xe2, 0x94, 0x17, 0x48, 0xbb, 0xb6, 0xdd, 0x54, 0xa6, 0x9f, 0x63, 0x95, 0xa1, 0xa4,
	0x3d, 0xd7, 0x77, 0x95, 0x59, 0x60, 0x26, 0x91, 0x69, 0xcc, 0xa6, 0x4c, 0x53, 0x70, 0x0b, 0x34,
	0x9d, 0x7d, 0x6d, 0xe4, 0x79, 0x9d, 0x2d, 0xe5, 0xd8, 0xc9, 0x4d, 0x67, 0x5f, 0x93, 0x17, 0xd0,
	0x69, 0x62, 0xa5, 0xfd, 0xa1, 0x97, 0xc4, 0x93, 0x07, 0xe3, 0xd5, 0xd8, 0xc7, 0x1f, 0x9d, 0x9c,
	0x2e, 0x7d, 0xe4, 0x29, 0x84, 0x26, 0x6c, 0x45, 0xb7, 0x87, 0x7e, 0x12, 0x4f, 0x76, 0x6f, 0x1e,
	0x48, 0x45, 0x81, 0xa9, 0xb3, 0x90, 0xff, 0xc1, 0xe7, 0x99, 0xa2, 0x3b, 0x43, 0x3f, 0xf1, 0x53,
	0xf3, 0x38, 0xfa, 0xed, 0x41, 0xcf, 0xc0, 0xf9, 0x72, 0x8e, 0x12, 0x4d, 0x04, 0x35, 0xcb, 0xd1,
	0x41, 0x0a, 0x53, 0x57, 0x98, 0xc0, 0xcd, 0xc3, 0x54, 0xf1, 0x4b, 0x07, 0x2b, 0x34, 0xa9, 0xe5,
	0x78, 0xcc, 0x2f, 0x91, 0x50, 0xe8, 0x28, 0x21, 0x35, 0xaf, 0xf2, 0x86, 0xd5, 0xb2, 0xbc, 0xc2,
	0x1b, 0xdc, 0x82, 0x37, 0x5c, 0xc7, 0x7b, 0x85, 0x21, 0x6a, 0xc7, 0xd0, 0xf9, 0x0b, 0x86, 0xee,
	0x0a, 0x86, 0xeb, 0xa0, 0x7b, 0xab, 0xa0, 0x47, 0xbf, 0x3c, 0xe8, 0x9b, 0x77, 0x4e, 0x51, 0xd5,
	0xa2, 0x52, 0x48, 0x9e, 0x43, 0x84, 0x95, 0xe6, 0x7a, 0x61, 0xdf, 0xbb, 0x25, 0x43, 0xeb, 0x6e,
	0x3c, 0xe4, 0x99, 0x0b, 0x49, 0xda, 0x28, 0xe2, 0xc9, 0xbd, 0x35, 0x42, 0x46, 0x74, 0xd9, 0x49,
	0x43, 0x87, 0x6b, 0x2c, 0x15, 0xf5, 0xdb, 0xe9, 0xd8, 0xc9, 0xce, 0x62, 0x06, 0xa3, 0x94, 0x42,
	0xda, 0xc4, 0x5a, 0x06, 0xbf, 0x33, 0x62, 0xea, 0x3c, 0x24, 0x81, 0x80, 0x57, 0x67, 0xc2, 0x26,
	0xd8, 0x32, 0xf7, 0xa8, 0x3a, 0x13, 0xa9, 0x75, 0x4c, 0x7e, 0x06, 0x10, 0x9b, 0x6b, 0x8e, 0x51,
	0x5e, 0xf0, 0x19, 0x92, 0x37, 0x10, 0x1d, 0xd8, 0xef, 0x91, 0xb4, 0x6e, 0x33, 0x78, 0xdc, 0xba,
	0x63, 0x93, 0xd5, 0xe8, 0x3f, 0x73, 0xfe, 0xc4, 0x7e, 0xb0, 0x1b, 0x9e, 0x7f, 0x05, 0xfe, 0x21,
	0xea, 0xcd, 0x2f, 0x7f, 0x8b, 0x05, 0x6e, 0x7c, 0xf9, 0x01, 0x44, 0xc7, 0xc8, 0xe4, 0xec, 0x9c,
	0x3c, 0x6c, 0x73, 0xda, 0x9f, 0xc1, 0x3f, 0x0c, 0x09, 0xdf, 0x8b, 0x9c, 0x57, 0xb7, 0xcd, 0xd8,
	0xbb, 0x29, 0x7d, 0x12, 0x5f, 0xb1, 0xba, 0x36, 0xe4, 0x35, 0x04, 0x06, 0x12, 0x59, 0xc7, 0x5c,
	0xd6, 0x7a, 0x71, 0xe7, 0x0e, 0x87, 0xb0, 0xfd, 0x99, 0x15, 0xdc, 0x70, 0xb0, 0x93, 0xd7, 0xe7,
	0xd8, 0xf6, 0x9d, 0x7b, 0x9c, 0x46, 0xf6, 0x5f, 0xfa, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0x95, 0x14, 0xa1, 0xec, 0x05, 0x00, 0x00,
}