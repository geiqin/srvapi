// Code generated by protoc-gen-go. DO NOT EDIT.
// source: area.proto

package geiqin_srv_bas

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

// 区域信息
type Area struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AreaId               int64    `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	LevelType            string   `protobuf:"bytes,3,opt,name=level_type,json=levelType,proto3" json:"level_type,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Shortname            string   `protobuf:"bytes,5,opt,name=shortname,proto3" json:"shortname,omitempty"`
	ParentPath           string   `protobuf:"bytes,6,opt,name=parent_path,json=parentPath,proto3" json:"parent_path,omitempty"`
	Province             string   `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`
	District             string   `protobuf:"bytes,9,opt,name=district,proto3" json:"district,omitempty"`
	ProvinceShortname    string   `protobuf:"bytes,10,opt,name=province_shortname,json=provinceShortname,proto3" json:"province_shortname,omitempty"`
	CityShortname        string   `protobuf:"bytes,11,opt,name=city_shortname,json=cityShortname,proto3" json:"city_shortname,omitempty"`
	DistrictShortname    string   `protobuf:"bytes,12,opt,name=district_shortname,json=districtShortname,proto3" json:"district_shortname,omitempty"`
	ProvincePinyin       string   `protobuf:"bytes,13,opt,name=province_pinyin,json=provincePinyin,proto3" json:"province_pinyin,omitempty"`
	CityPinyin           string   `protobuf:"bytes,14,opt,name=city_pinyin,json=cityPinyin,proto3" json:"city_pinyin,omitempty"`
	DistrictPinyin       string   `protobuf:"bytes,15,opt,name=district_pinyin,json=districtPinyin,proto3" json:"district_pinyin,omitempty"`
	Pinyin               string   `protobuf:"bytes,16,opt,name=pinyin,proto3" json:"pinyin,omitempty"`
	Jianpin              string   `protobuf:"bytes,17,opt,name=jianpin,proto3" json:"jianpin,omitempty"`
	FirstChar            string   `protobuf:"bytes,18,opt,name=first_char,json=firstChar,proto3" json:"first_char,omitempty"`
	CityCode             string   `protobuf:"bytes,19,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	ZipCode              string   `protobuf:"bytes,20,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	Lng                  string   `protobuf:"bytes,21,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  string   `protobuf:"bytes,22,opt,name=lat,proto3" json:"lat,omitempty"`
	Remark1              string   `protobuf:"bytes,23,opt,name=remark1,proto3" json:"remark1,omitempty"`
	Remark2              string   `protobuf:"bytes,24,opt,name=remark2,proto3" json:"remark2,omitempty"`
	Children             []*Area  `protobuf:"bytes,25,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Area) Reset()         { *m = Area{} }
func (m *Area) String() string { return proto.CompactTextString(m) }
func (*Area) ProtoMessage()    {}
func (*Area) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4d14b2d06184cbd, []int{0}
}

func (m *Area) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Area.Unmarshal(m, b)
}
func (m *Area) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Area.Marshal(b, m, deterministic)
}
func (m *Area) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Area.Merge(m, src)
}
func (m *Area) XXX_Size() int {
	return xxx_messageInfo_Area.Size(m)
}
func (m *Area) XXX_DiscardUnknown() {
	xxx_messageInfo_Area.DiscardUnknown(m)
}

var xxx_messageInfo_Area proto.InternalMessageInfo

func (m *Area) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Area) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Area) GetLevelType() string {
	if m != nil {
		return m.LevelType
	}
	return ""
}

func (m *Area) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Area) GetShortname() string {
	if m != nil {
		return m.Shortname
	}
	return ""
}

func (m *Area) GetParentPath() string {
	if m != nil {
		return m.ParentPath
	}
	return ""
}

func (m *Area) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Area) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Area) GetDistrict() string {
	if m != nil {
		return m.District
	}
	return ""
}

func (m *Area) GetProvinceShortname() string {
	if m != nil {
		return m.ProvinceShortname
	}
	return ""
}

func (m *Area) GetCityShortname() string {
	if m != nil {
		return m.CityShortname
	}
	return ""
}

func (m *Area) GetDistrictShortname() string {
	if m != nil {
		return m.DistrictShortname
	}
	return ""
}

func (m *Area) GetProvincePinyin() string {
	if m != nil {
		return m.ProvincePinyin
	}
	return ""
}

func (m *Area) GetCityPinyin() string {
	if m != nil {
		return m.CityPinyin
	}
	return ""
}

func (m *Area) GetDistrictPinyin() string {
	if m != nil {
		return m.DistrictPinyin
	}
	return ""
}

func (m *Area) GetPinyin() string {
	if m != nil {
		return m.Pinyin
	}
	return ""
}

func (m *Area) GetJianpin() string {
	if m != nil {
		return m.Jianpin
	}
	return ""
}

func (m *Area) GetFirstChar() string {
	if m != nil {
		return m.FirstChar
	}
	return ""
}

func (m *Area) GetCityCode() string {
	if m != nil {
		return m.CityCode
	}
	return ""
}

func (m *Area) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *Area) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *Area) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *Area) GetRemark1() string {
	if m != nil {
		return m.Remark1
	}
	return ""
}

func (m *Area) GetRemark2() string {
	if m != nil {
		return m.Remark2
	}
	return ""
}

func (m *Area) GetChildren() []*Area {
	if m != nil {
		return m.Children
	}
	return nil
}

type AreaResponse struct {
	Entity               *Area    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Area  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AreaResponse) Reset()         { *m = AreaResponse{} }
func (m *AreaResponse) String() string { return proto.CompactTextString(m) }
func (*AreaResponse) ProtoMessage()    {}
func (*AreaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4d14b2d06184cbd, []int{1}
}

func (m *AreaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AreaResponse.Unmarshal(m, b)
}
func (m *AreaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AreaResponse.Marshal(b, m, deterministic)
}
func (m *AreaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AreaResponse.Merge(m, src)
}
func (m *AreaResponse) XXX_Size() int {
	return xxx_messageInfo_AreaResponse.Size(m)
}
func (m *AreaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AreaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AreaResponse proto.InternalMessageInfo

func (m *AreaResponse) GetEntity() *Area {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *AreaResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *AreaResponse) GetItems() []*Area {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *AreaResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Area)(nil), "geiqin.srv.bas.Area")
	proto.RegisterType((*AreaResponse)(nil), "geiqin.srv.bas.AreaResponse")
}

func init() { proto.RegisterFile("area.proto", fileDescriptor_e4d14b2d06184cbd) }

var fileDescriptor_e4d14b2d06184cbd = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6e, 0x12, 0x41,
	0x14, 0xc6, 0x85, 0x85, 0x2d, 0x1c, 0x5a, 0xda, 0x1e, 0xfb, 0x67, 0x5a, 0x6b, 0x6c, 0x48, 0x4c,
	0x89, 0x5a, 0xa2, 0x78, 0x69, 0x6f, 0x94, 0x18, 0x63, 0xe2, 0x05, 0x81, 0x26, 0x5e, 0x92, 0xe9,
	0xee, 0x91, 0x1d, 0x85, 0xd9, 0x71, 0x76, 0x42, 0x42, 0x5f, 0xc1, 0xb7, 0xf2, 0xa9, 0xbc, 0x34,
	0x33, 0xb3, 0xbb, 0x60, 0x83, 0x9a, 0xf4, 0x6e, 0xce, 0xf7, 0xfd, 0xe6, 0x3b, 0x87, 0xd9, 0x19,
	0x00, 0xb8, 0x26, 0xde, 0x53, 0x3a, 0x35, 0x29, 0xb6, 0xa7, 0x24, 0xbe, 0x0b, 0xd9, 0xcb, 0xf4,
	0xa2, 0x77, 0xc3, 0xb3, 0xd3, 0xed, 0x28, 0x9d, 0xcf, 0x53, 0xe9, 0xdd, 0xce, 0xaf, 0x3a, 0xd4,
	0xde, 0x6a, 0xe2, 0xd8, 0x86, 0xaa, 0x88, 0x59, 0xe5, 0xbc, 0xd2, 0x0d, 0x46, 0x55, 0x11, 0xe3,
	0x31, 0x6c, 0xd9, 0x90, 0x89, 0x88, 0x59, 0xd5, 0x89, 0xa1, 0x2d, 0x3f, 0xc6, 0xf8, 0x18, 0x60,
	0x46, 0x0b, 0x9a, 0x4d, 0xcc, 0x52, 0x11, 0x0b, 0xce, 0x2b, 0xdd, 0xe6, 0xa8, 0xe9, 0x94, 0xeb,
	0xa5, 0x22, 0x44, 0xa8, 0x49, 0x3e, 0x27, 0x56, 0x73, 0x86, 0x5b, 0xe3, 0x19, 0x34, 0xb3, 0x24,
	0xd5, 0xc6, 0x19, 0x75, 0xbf, 0xa3, 0x14, 0xf0, 0x09, 0xb4, 0x14, 0xd7, 0x24, 0xcd, 0x44, 0x71,
	0x93, 0xb0, 0xd0, 0xf9, 0xe0, 0xa5, 0x21, 0x37, 0x09, 0x9e, 0x42, 0x43, 0xe9, 0x74, 0x21, 0x64,
	0x44, 0x6c, 0xcb, 0xb9, 0x65, 0x6d, 0xdb, 0x45, 0xc2, 0x2c, 0x59, 0xc3, 0xb7, 0xb3, 0x6b, 0xcb,
	0xc7, 0x22, 0x33, 0x5a, 0x44, 0x86, 0x35, 0x3d, 0x5f, 0xd4, 0x78, 0x09, 0x58, 0xec, 0x9d, 0xac,
	0x66, 0x02, 0x47, 0xed, 0x17, 0xce, 0xb8, 0x9c, 0xed, 0x29, 0xb4, 0x6d, 0xe4, 0x1a, 0xda, 0x72,
	0xe8, 0x8e, 0x55, 0x57, 0xd8, 0x25, 0x60, 0xd1, 0x61, 0x0d, 0xdd, 0xf6, 0xa9, 0x85, 0xb3, 0xc2,
	0x2f, 0x60, 0xb7, 0x1c, 0x42, 0x09, 0xb9, 0x14, 0x92, 0xed, 0x38, 0xb6, 0x5d, 0xc8, 0x43, 0xa7,
	0xda, 0xa3, 0x71, 0xed, 0x73, 0xa8, 0xed, 0x8f, 0xc6, 0x4a, 0x39, 0x70, 0x01, 0xbb, 0x65, 0xe3,
	0x1c, 0xda, 0xf5, 0x49, 0x85, 0x9c, 0x83, 0x47, 0x10, 0xe6, 0xfe, 0x9e, 0xf3, 0xf3, 0x0a, 0x19,
	0x6c, 0x7d, 0x15, 0x5c, 0x2a, 0x21, 0xd9, 0xbe, 0x33, 0x8a, 0xd2, 0x7e, 0xe7, 0x2f, 0x42, 0x67,
	0x66, 0x12, 0x25, 0x5c, 0x33, 0xf4, 0x5f, 0xcd, 0x29, 0x83, 0x84, 0x6b, 0x7c, 0x04, 0x4d, 0x37,
	0x5a, 0x94, 0xc6, 0xc4, 0x1e, 0xfa, 0x53, 0xb6, 0xc2, 0x20, 0x8d, 0x09, 0x4f, 0xa0, 0x71, 0x2b,
	0x94, 0xf7, 0x0e, 0x7c, 0xec, 0xad, 0x50, 0xce, 0xda, 0x83, 0x60, 0x26, 0xa7, 0xec, 0xd0, 0xa9,
	0x76, 0xe9, 0x14, 0x6e, 0xd8, 0x51, 0xae, 0x70, 0x63, 0x87, 0xd2, 0x34, 0xe7, 0xfa, 0xdb, 0x2b,
	0x76, 0xec, 0x77, 0xe7, 0xe5, 0xca, 0xe9, 0x33, 0xb6, 0xee, 0xf4, 0xf1, 0x25, 0x34, 0xa2, 0x44,
	0xcc, 0x62, 0x4d, 0x92, 0x9d, 0x9c, 0x07, 0xdd, 0x56, 0xff, 0xa0, 0xf7, 0xe7, 0xcd, 0xef, 0xd9,
	0x7b, 0x3e, 0x2a, 0xa9, 0xce, 0xcf, 0x0a, 0x6c, 0x3b, 0x89, 0x32, 0x95, 0xca, 0x8c, 0xf0, 0x05,
	0x84, 0x24, 0x8d, 0xbd, 0x4d, 0xf6, 0x19, 0xfc, 0x2d, 0x20, 0x67, 0xf0, 0x39, 0xd4, 0x15, 0x9f,
	0x92, 0x76, 0xcf, 0xa3, 0xd5, 0x3f, 0xbc, 0x0b, 0x0f, 0xad, 0x39, 0xf2, 0x0c, 0x3e, 0x83, 0xba,
	0x30, 0x34, 0xcf, 0x58, 0xf0, 0x8f, 0xd1, 0x3c, 0x62, 0x83, 0x49, 0xeb, 0x54, 0xbb, 0x27, 0xb4,
	0x21, 0xf8, 0xbd, 0x35, 0x47, 0x9e, 0xe9, 0xff, 0xa8, 0x42, 0xcb, 0x6e, 0x1e, 0x93, 0x5e, 0x88,
	0x88, 0xf0, 0x0d, 0x04, 0x1f, 0xc8, 0xe0, 0xc6, 0x06, 0xa7, 0x67, 0x1b, 0xdb, 0xe6, 0x3f, 0xbf,
	0xf3, 0x00, 0x07, 0x10, 0x8e, 0x89, 0xeb, 0x28, 0xc1, 0x93, 0xbb, 0xe4, 0x3b, 0x9e, 0xd1, 0xe7,
	0x84, 0x34, 0xfd, 0x37, 0xe4, 0x0a, 0x6a, 0xd7, 0x9a, 0xe8, 0x9e, 0x23, 0x5c, 0x41, 0xed, 0x93,
	0xc8, 0xee, 0xf9, 0x03, 0x6e, 0x42, 0xf7, 0xa7, 0xf6, 0xfa, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xdd, 0x11, 0x64, 0x94, 0x00, 0x05, 0x00, 0x00,
}