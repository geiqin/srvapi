// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store.proto

package geiqin_srv_mts

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

type Store struct {
	Id                   int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LogoId               int64        `protobuf:"varint,3,opt,name=logo_id,json=logoId,proto3" json:"logo_id,omitempty"`
	LogoUrl              string       `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	CreatorId            int64        `protobuf:"varint,6,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	ApplicationId        int32        `protobuf:"varint,7,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Version              string       `protobuf:"bytes,8,opt,name=version,proto3" json:"version,omitempty"`
	Industry             int64        `protobuf:"varint,9,opt,name=industry,proto3" json:"industry,omitempty"`
	AreaId               int64        `protobuf:"varint,10,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Address              string       `protobuf:"bytes,11,opt,name=address,proto3" json:"address,omitempty"`
	Lng                  string       `protobuf:"bytes,12,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  string       `protobuf:"bytes,13,opt,name=lat,proto3" json:"lat,omitempty"`
	CustomDomain         string       `protobuf:"bytes,14,opt,name=custom_domain,json=customDomain,proto3" json:"custom_domain,omitempty"`
	Status               string       `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	Enabled              bool         `protobuf:"varint,16,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Trusted              bool         `protobuf:"varint,17,opt,name=trusted,proto3" json:"trusted,omitempty"`
	StartedAt            string       `protobuf:"bytes,18,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	ExpiredAt            string       `protobuf:"bytes,19,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	CreatedAt            string       `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string       `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Application          *Application `protobuf:"bytes,22,opt,name=application,proto3" json:"application,omitempty"`
	Ids                  []int64      `protobuf:"varint,23,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Store) Reset()         { *m = Store{} }
func (m *Store) String() string { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()    {}
func (*Store) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{0}
}

func (m *Store) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Store.Unmarshal(m, b)
}
func (m *Store) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Store.Marshal(b, m, deterministic)
}
func (m *Store) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Store.Merge(m, src)
}
func (m *Store) XXX_Size() int {
	return xxx_messageInfo_Store.Size(m)
}
func (m *Store) XXX_DiscardUnknown() {
	xxx_messageInfo_Store.DiscardUnknown(m)
}

var xxx_messageInfo_Store proto.InternalMessageInfo

func (m *Store) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Store) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Store) GetLogoId() int64 {
	if m != nil {
		return m.LogoId
	}
	return 0
}

func (m *Store) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *Store) GetCreatorId() int64 {
	if m != nil {
		return m.CreatorId
	}
	return 0
}

func (m *Store) GetApplicationId() int32 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *Store) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Store) GetIndustry() int64 {
	if m != nil {
		return m.Industry
	}
	return 0
}

func (m *Store) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Store) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Store) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *Store) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *Store) GetCustomDomain() string {
	if m != nil {
		return m.CustomDomain
	}
	return ""
}

func (m *Store) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Store) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Store) GetTrusted() bool {
	if m != nil {
		return m.Trusted
	}
	return false
}

func (m *Store) GetStartedAt() string {
	if m != nil {
		return m.StartedAt
	}
	return ""
}

func (m *Store) GetExpiredAt() string {
	if m != nil {
		return m.ExpiredAt
	}
	return ""
}

func (m *Store) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Store) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Store) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *Store) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type StoreSecret struct {
	StoreKey             string   `protobuf:"bytes,1,opt,name=store_key,json=storeKey,proto3" json:"store_key,omitempty"`
	StoreSecret          string   `protobuf:"bytes,2,opt,name=store_secret,json=storeSecret,proto3" json:"store_secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreSecret) Reset()         { *m = StoreSecret{} }
func (m *StoreSecret) String() string { return proto.CompactTextString(m) }
func (*StoreSecret) ProtoMessage()    {}
func (*StoreSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{1}
}

func (m *StoreSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreSecret.Unmarshal(m, b)
}
func (m *StoreSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreSecret.Marshal(b, m, deterministic)
}
func (m *StoreSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreSecret.Merge(m, src)
}
func (m *StoreSecret) XXX_Size() int {
	return xxx_messageInfo_StoreSecret.Size(m)
}
func (m *StoreSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreSecret.DiscardUnknown(m)
}

var xxx_messageInfo_StoreSecret proto.InternalMessageInfo

func (m *StoreSecret) GetStoreKey() string {
	if m != nil {
		return m.StoreKey
	}
	return ""
}

func (m *StoreSecret) GetStoreSecret() string {
	if m != nil {
		return m.StoreSecret
	}
	return ""
}

//
type StoreResponse struct {
	Entity               *Store   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Store `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{2}
}

func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (m *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(m, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetEntity() *Store {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *StoreResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *StoreResponse) GetItems() []*Store {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *StoreResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *StoreResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Store)(nil), "geiqin.srv.mts.Store")
	proto.RegisterType((*StoreSecret)(nil), "geiqin.srv.mts.StoreSecret")
	proto.RegisterType((*StoreResponse)(nil), "geiqin.srv.mts.StoreResponse")
}

func init() {
	proto.RegisterFile("store.proto", fileDescriptor_98bbca36ef968dfc)
}

var fileDescriptor_98bbca36ef968dfc = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdf, 0x6e, 0x1a, 0x39,
	0x14, 0xc6, 0x97, 0x0c, 0x10, 0x38, 0xfc, 0xd9, 0xc4, 0x9b, 0x3f, 0x0e, 0x51, 0x24, 0x96, 0xd5,
	0x4a, 0x48, 0x55, 0xb9, 0x48, 0xaf, 0x13, 0x95, 0x96, 0xa8, 0x42, 0x6d, 0xa5, 0x6a, 0x68, 0xd4,
	0x4b, 0xe4, 0x8c, 0x4f, 0xc0, 0xca, 0xcc, 0x78, 0x6a, 0x7b, 0xd2, 0xf2, 0x3c, 0x7d, 0xc0, 0xf6,
	0x11, 0x2a, 0xdb, 0x93, 0x96, 0x10, 0xaa, 0x46, 0xe5, 0xce, 0xe7, 0xfb, 0x7e, 0xfe, 0x38, 0x1c,
	0x1f, 0x04, 0x34, 0xb4, 0x91, 0x0a, 0x07, 0x99, 0x92, 0x46, 0x92, 0xf6, 0x0c, 0xc5, 0x47, 0x91,
	0x0e, 0xb4, 0xba, 0x1d, 0x24, 0x46, 0x77, 0x9a, 0x91, 0x4c, 0x12, 0x99, 0x7a, 0xb7, 0xb3, 0xcb,
	0xb2, 0x2c, 0x16, 0x11, 0x33, 0xe2, 0x4e, 0xea, 0x7d, 0x2b, 0x43, 0x65, 0x62, 0x03, 0x48, 0x1b,
	0xb6, 0x04, 0xa7, 0xa5, 0x6e, 0xa9, 0x1f, 0x84, 0x5b, 0x82, 0x13, 0x02, 0xe5, 0x94, 0x25, 0x48,
	0xb7, 0xba, 0xa5, 0x7e, 0x3d, 0x74, 0x67, 0x72, 0x08, 0xdb, 0xb1, 0x9c, 0xc9, 0xa9, 0xe0, 0x34,
	0x70, 0x60, 0xd5, 0x96, 0x63, 0x4e, 0x8e, 0xa0, 0xe6, 0x8c, 0x5c, 0xc5, 0xb4, 0xec, 0x2e, 0x38,
	0xf0, 0x52, 0xc5, 0xe4, 0x04, 0x20, 0x52, 0xc8, 0x8c, 0x54, 0xf6, 0x5a, 0xd5, 0x5d, 0xab, 0x17,
	0xca, 0x98, 0x93, 0xff, 0xa1, 0xbd, 0xd4, 0x95, 0x45, 0xb6, 0xbb, 0xa5, 0x7e, 0x25, 0x6c, 0x2d,
	0xa9, 0x63, 0x4e, 0x28, 0x6c, 0xdf, 0xa2, 0xd2, 0x42, 0xa6, 0xb4, 0xe6, 0xf3, 0x8b, 0x92, 0x74,
	0xa0, 0x26, 0x52, 0x9e, 0x6b, 0xa3, 0x16, 0xb4, 0xee, 0xd2, 0x7f, 0xd4, 0xb6, 0x5f, 0xa6, 0x90,
	0xd9, 0x54, 0xf0, 0xfd, 0xda, 0xd2, 0xc7, 0x31, 0xce, 0x15, 0x6a, 0x4d, 0x1b, 0x3e, 0xae, 0x28,
	0xc9, 0x0e, 0x04, 0x71, 0x3a, 0xa3, 0x4d, 0xa7, 0xda, 0xa3, 0x53, 0x98, 0xa1, 0xad, 0x42, 0x61,
	0x86, 0xfc, 0x07, 0xad, 0x28, 0xd7, 0x46, 0x26, 0x53, 0x2e, 0x13, 0x26, 0x52, 0xda, 0x76, 0x5e,
	0xd3, 0x8b, 0x23, 0xa7, 0x91, 0x03, 0xa8, 0x6a, 0xc3, 0x4c, 0xae, 0xe9, 0xdf, 0xce, 0x2d, 0x2a,
	0xfb, 0xd1, 0x98, 0xb2, 0xab, 0x18, 0x39, 0xdd, 0xe9, 0x96, 0xfa, 0xb5, 0xf0, 0xae, 0xb4, 0x8e,
	0x51, 0xb9, 0x36, 0xc8, 0xe9, 0xae, 0x77, 0x8a, 0xd2, 0xce, 0x50, 0x1b, 0xa6, 0x0c, 0xf2, 0x29,
	0x33, 0x94, 0xb8, 0xbc, 0x7a, 0xa1, 0x0c, 0x8d, 0xb5, 0xf1, 0x73, 0x26, 0x94, 0xb7, 0xff, 0xf1,
	0x76, 0xa1, 0x78, 0xdb, 0xcd, 0xdb, 0xdb, 0x7b, 0xde, 0x2e, 0x14, 0x6f, 0xe7, 0x19, 0xbf, 0xb3,
	0xf7, 0xbd, 0x5d, 0x28, 0x43, 0x43, 0xce, 0xa0, 0xb1, 0xf4, 0x14, 0xf4, 0xa0, 0x5b, 0xea, 0x37,
	0x4e, 0x8f, 0x07, 0xf7, 0x17, 0x6d, 0x30, 0xfc, 0x89, 0x84, 0xcb, 0xbc, 0x9d, 0x9e, 0xe0, 0x9a,
	0x1e, 0x76, 0x83, 0x7e, 0x10, 0xda, 0x63, 0xef, 0x2d, 0x34, 0xdc, 0xc6, 0x4d, 0x30, 0x52, 0x68,
	0xc8, 0x31, 0xd4, 0xdd, 0x06, 0x4f, 0x6f, 0x70, 0xe1, 0xd6, 0xaf, 0x1e, 0xd6, 0x9c, 0xf0, 0x1a,
	0x17, 0xe4, 0x5f, 0x68, 0x7a, 0x53, 0x3b, 0xb8, 0x58, 0x46, 0xbf, 0xf2, 0xfe, 0x7e, 0xef, 0x6b,
	0x09, 0x5a, 0x2e, 0x2f, 0x44, 0x9d, 0xc9, 0x54, 0x23, 0x79, 0x0a, 0x55, 0x4c, 0x8d, 0x30, 0x3e,
	0xae, 0x71, 0xba, 0xbf, 0xda, 0xac, 0xc7, 0x0b, 0x88, 0x3c, 0x81, 0x4a, 0xc6, 0x66, 0xa8, 0x5c,
	0xf8, 0x1a, 0xfa, 0x9d, 0x35, 0x43, 0xcf, 0x58, 0x58, 0x18, 0x4c, 0x34, 0x0d, 0xba, 0xc1, 0xaf,
	0xa3, 0x3d, 0x63, 0x61, 0x54, 0x4a, 0x2a, 0xf7, 0x93, 0x58, 0x03, 0x5f, 0x58, 0x33, 0xf4, 0x0c,
	0xe9, 0x43, 0x59, 0xa4, 0xd7, 0x92, 0x56, 0x1c, 0xbb, 0xb7, 0xca, 0x8e, 0xd3, 0x6b, 0x19, 0x3a,
	0xe2, 0xf4, 0x4b, 0x05, 0x9a, 0xc5, 0x04, 0xd5, 0xad, 0x88, 0x90, 0x3c, 0x87, 0xea, 0x4b, 0xf7,
	0x9c, 0x64, 0x7d, 0x3f, 0x9d, 0x93, 0xf5, 0x6d, 0x16, 0x03, 0xeb, 0xfd, 0x65, 0x13, 0x2e, 0xdd,
	0x8b, 0x6f, 0x92, 0x30, 0xc2, 0x18, 0x37, 0x48, 0x18, 0x41, 0x79, 0x98, 0x9b, 0x39, 0x39, 0x5e,
	0x0b, 0xfa, 0xd7, 0x7e, 0x98, 0xf2, 0x5e, 0xde, 0x60, 0x7a, 0xbf, 0x8f, 0xc9, 0x27, 0x61, 0xa2,
	0xf9, 0xa3, 0xfb, 0x58, 0x4d, 0x38, 0x87, 0xb2, 0x1d, 0xf6, 0xc3, 0xfb, 0x17, 0x49, 0x66, 0x16,
	0xbf, 0xff, 0x1e, 0x67, 0x10, 0xbc, 0x42, 0xf3, 0xc7, 0x63, 0x38, 0x87, 0xf2, 0x1b, 0x19, 0xdd,
	0x6c, 0xf4, 0x94, 0x69, 0xbc, 0x49, 0xc2, 0x08, 0xaa, 0x13, 0x64, 0x2a, 0x9a, 0x93, 0xa3, 0x55,
	0xf4, 0x05, 0xd3, 0xf8, 0x61, 0x8e, 0x8f, 0x48, 0xb9, 0xaa, 0xba, 0x3f, 0x98, 0x67, 0xdf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x5b, 0x51, 0x49, 0x0f, 0xa0, 0x06, 0x00, 0x00,
}
