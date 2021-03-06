// Code generated by protoc-gen-go. DO NOT EDIT.
// source: areaInfo.proto

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

type AreaInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AreaId               int64    `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	LevelType            string   `protobuf:"bytes,3,opt,name=level_type,json=levelType,proto3" json:"level_type,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Shortname            string   `protobuf:"bytes,5,opt,name=shortname,proto3" json:"shortname,omitempty"`
	ParentPath           string   `protobuf:"bytes,6,opt,name=parent_path,json=parentPath,proto3" json:"parent_path,omitempty"`
	Province             string   `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`
	District             string   `protobuf:"bytes,9,opt,name=district,proto3" json:"district,omitempty"`
	ProvinceShortname    string   `protobuf:"bytes,11,opt,name=province_shortname,json=provinceShortname,proto3" json:"province_shortname,omitempty"`
	CityShortname        string   `protobuf:"bytes,12,opt,name=city_shortname,json=cityShortname,proto3" json:"city_shortname,omitempty"`
	DistrictShortname    string   `protobuf:"bytes,13,opt,name=district_shortname,json=districtShortname,proto3" json:"district_shortname,omitempty"`
	ProvincePinyin       string   `protobuf:"bytes,14,opt,name=province_pinyin,json=provincePinyin,proto3" json:"province_pinyin,omitempty"`
	CityPinyin           string   `protobuf:"bytes,15,opt,name=city_pinyin,json=cityPinyin,proto3" json:"city_pinyin,omitempty"`
	DistrictPinyin       string   `protobuf:"bytes,16,opt,name=district_pinyin,json=districtPinyin,proto3" json:"district_pinyin,omitempty"`
	Pinyin               string   `protobuf:"bytes,17,opt,name=pinyin,proto3" json:"pinyin,omitempty"`
	Jianpin              string   `protobuf:"bytes,18,opt,name=jianpin,proto3" json:"jianpin,omitempty"`
	FirstChar            string   `protobuf:"bytes,19,opt,name=first_char,json=firstChar,proto3" json:"first_char,omitempty"`
	CityCode             string   `protobuf:"bytes,20,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	ZipCode              string   `protobuf:"bytes,21,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	Lng                  string   `protobuf:"bytes,22,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  string   `protobuf:"bytes,23,opt,name=lat,proto3" json:"lat,omitempty"`
	Remark1              string   `protobuf:"bytes,24,opt,name=remark1,proto3" json:"remark1,omitempty"`
	Remark2              string   `protobuf:"bytes,25,opt,name=remark2,proto3" json:"remark2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AreaInfo) Reset()         { *m = AreaInfo{} }
func (m *AreaInfo) String() string { return proto.CompactTextString(m) }
func (*AreaInfo) ProtoMessage()    {}
func (*AreaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4d5c69474b95158, []int{0}
}

func (m *AreaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AreaInfo.Unmarshal(m, b)
}
func (m *AreaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AreaInfo.Marshal(b, m, deterministic)
}
func (m *AreaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AreaInfo.Merge(m, src)
}
func (m *AreaInfo) XXX_Size() int {
	return xxx_messageInfo_AreaInfo.Size(m)
}
func (m *AreaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AreaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AreaInfo proto.InternalMessageInfo

func (m *AreaInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AreaInfo) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *AreaInfo) GetLevelType() string {
	if m != nil {
		return m.LevelType
	}
	return ""
}

func (m *AreaInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AreaInfo) GetShortname() string {
	if m != nil {
		return m.Shortname
	}
	return ""
}

func (m *AreaInfo) GetParentPath() string {
	if m != nil {
		return m.ParentPath
	}
	return ""
}

func (m *AreaInfo) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *AreaInfo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *AreaInfo) GetDistrict() string {
	if m != nil {
		return m.District
	}
	return ""
}

func (m *AreaInfo) GetProvinceShortname() string {
	if m != nil {
		return m.ProvinceShortname
	}
	return ""
}

func (m *AreaInfo) GetCityShortname() string {
	if m != nil {
		return m.CityShortname
	}
	return ""
}

func (m *AreaInfo) GetDistrictShortname() string {
	if m != nil {
		return m.DistrictShortname
	}
	return ""
}

func (m *AreaInfo) GetProvincePinyin() string {
	if m != nil {
		return m.ProvincePinyin
	}
	return ""
}

func (m *AreaInfo) GetCityPinyin() string {
	if m != nil {
		return m.CityPinyin
	}
	return ""
}

func (m *AreaInfo) GetDistrictPinyin() string {
	if m != nil {
		return m.DistrictPinyin
	}
	return ""
}

func (m *AreaInfo) GetPinyin() string {
	if m != nil {
		return m.Pinyin
	}
	return ""
}

func (m *AreaInfo) GetJianpin() string {
	if m != nil {
		return m.Jianpin
	}
	return ""
}

func (m *AreaInfo) GetFirstChar() string {
	if m != nil {
		return m.FirstChar
	}
	return ""
}

func (m *AreaInfo) GetCityCode() string {
	if m != nil {
		return m.CityCode
	}
	return ""
}

func (m *AreaInfo) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *AreaInfo) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *AreaInfo) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *AreaInfo) GetRemark1() string {
	if m != nil {
		return m.Remark1
	}
	return ""
}

func (m *AreaInfo) GetRemark2() string {
	if m != nil {
		return m.Remark2
	}
	return ""
}

func init() {
	proto.RegisterType((*AreaInfo)(nil), "geiqin.srv.tms.AreaInfo")
}

func init() {
	proto.RegisterFile("areaInfo.proto", fileDescriptor_e4d5c69474b95158)
}

var fileDescriptor_e4d5c69474b95158 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xcf, 0x4e, 0x1b, 0x31,
	0x10, 0x87, 0x95, 0x04, 0x92, 0xec, 0x50, 0x36, 0x30, 0x6d, 0x61, 0xe8, 0x1f, 0x15, 0x55, 0xaa,
	0xe0, 0x42, 0xa4, 0xd2, 0x27, 0xa8, 0x38, 0xf5, 0x86, 0x68, 0xef, 0x2b, 0x77, 0xd7, 0xb0, 0xd3,
	0x26, 0x5e, 0xd7, 0x6b, 0x45, 0x0a, 0x8f, 0xd8, 0xa7, 0x42, 0x1e, 0xdb, 0x9b, 0xdc, 0x3c, 0xdf,
	0xef, 0xf3, 0xcc, 0xc8, 0xbb, 0x50, 0x2a, 0xa7, 0xd5, 0x0f, 0xf3, 0xd8, 0x2d, 0xad, 0xeb, 0x7c,
	0x87, 0xe5, 0x93, 0xe6, 0x7f, 0x6c, 0x96, 0xbd, 0xdb, 0x2c, 0xfd, 0xba, 0xff, 0xfc, 0xff, 0x10,
	0xe6, 0xdf, 0x93, 0x82, 0x25, 0x8c, 0xb9, 0xa1, 0xd1, 0xe5, 0xe8, 0x7a, 0xf2, 0x30, 0xe6, 0x06,
	0xcf, 0x61, 0x16, 0xae, 0x57, 0xdc, 0xd0, 0x58, 0xe0, 0x54, 0xba, 0x35, 0xf8, 0x11, 0x60, 0xa5,
	0x37, 0x7a, 0x55, 0xf9, 0xad, 0xd5, 0x34, 0xb9, 0x1c, 0x5d, 0x17, 0x0f, 0x85, 0x90, 0x5f, 0x5b,
	0xab, 0x11, 0xe1, 0xc0, 0xa8, 0xb5, 0xa6, 0x03, 0x09, 0xe4, 0x8c, 0x1f, 0xa0, 0xe8, 0xdb, 0xce,
	0x79, 0x09, 0x0e, 0xe3, 0x8d, 0x01, 0xe0, 0x27, 0x38, 0xb2, 0xca, 0x69, 0xe3, 0x2b, 0xab, 0x7c,
	0x4b, 0x53, 0xc9, 0x21, 0xa2, 0x7b, 0xe5, 0x5b, 0x7c, 0x07, 0x73, 0xeb, 0xba, 0x0d, 0x9b, 0x5a,
	0xd3, 0x4c, 0xd2, 0xa1, 0x0e, 0xe3, 0x6a, 0xf6, 0x5b, 0x9a, 0xc7, 0x71, 0xe1, 0x1c, 0xfc, 0x86,
	0x7b, 0xef, 0xb8, 0xf6, 0x54, 0x44, 0x3f, 0xd7, 0x78, 0x03, 0x98, 0xef, 0x56, 0xbb, 0x9d, 0x8e,
	0xc4, 0x3a, 0xcd, 0xc9, 0xcf, 0x61, 0xb7, 0x2f, 0x50, 0x86, 0x96, 0x7b, 0xea, 0x2b, 0x51, 0x8f,
	0x03, 0xdd, 0x69, 0x37, 0x80, 0x79, 0xc2, 0x9e, 0x7a, 0x1c, 0xbb, 0xe6, 0x64, 0xa7, 0x5f, 0xc1,
	0x62, 0x58, 0xc2, 0xb2, 0xd9, 0xb2, 0xa1, 0x52, 0xdc, 0x32, 0xe3, 0x7b, 0xa1, 0xe1, 0x69, 0x64,
	0x7c, 0x92, 0x16, 0xf1, 0x69, 0x02, 0x4a, 0xc2, 0x15, 0x2c, 0x86, 0xc1, 0x49, 0x3a, 0x89, 0x9d,
	0x32, 0x4e, 0xe2, 0x19, 0x4c, 0x53, 0x7e, 0x2a, 0x79, 0xaa, 0x90, 0x60, 0xf6, 0x87, 0x95, 0xb1,
	0x6c, 0x08, 0x25, 0xc8, 0x65, 0xf8, 0xce, 0x8f, 0xec, 0x7a, 0x5f, 0xd5, 0xad, 0x72, 0xf4, 0x3a,
	0x7e, 0x35, 0x21, 0x77, 0xad, 0x72, 0xf8, 0x1e, 0x0a, 0x59, 0xad, 0xee, 0x1a, 0x4d, 0x6f, 0xe2,
	0x2b, 0x07, 0x70, 0xd7, 0x35, 0x1a, 0x2f, 0x60, 0xfe, 0xcc, 0x36, 0x66, 0x6f, 0x63, 0xdb, 0x67,
	0xb6, 0x12, 0x9d, 0xc0, 0x64, 0x65, 0x9e, 0xe8, 0x4c, 0x68, 0x38, 0x0a, 0x51, 0x9e, 0xce, 0x13,
	0x51, 0x3e, 0x2c, 0xe5, 0xf4, 0x5a, 0xb9, 0xbf, 0x5f, 0x89, 0xe2, 0xed, 0x54, 0xee, 0x92, 0x5b,
	0xba, 0xd8, 0x4f, 0x6e, 0x7f, 0x4f, 0xe5, 0x1f, 0xff, 0xf6, 0x12, 0x00, 0x00, 0xff, 0xff, 0x52,
	0xc0, 0xd4, 0x7e, 0xf5, 0x02, 0x00, 0x00,
}
