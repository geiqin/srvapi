// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fetch.proto

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

type FetchWhere struct {
	Paged                int32    `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorting              string   `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	Keywords             string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Id                   int64    `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	Ids                  []int64  `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Lng                  string   `protobuf:"bytes,7,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                  string   `protobuf:"bytes,8,opt,name=lat,proto3" json:"lat,omitempty"`
	Name                 string   `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchWhere) Reset()         { *m = FetchWhere{} }
func (m *FetchWhere) String() string { return proto.CompactTextString(m) }
func (*FetchWhere) ProtoMessage()    {}
func (*FetchWhere) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3607b78a5221b8, []int{0}
}

func (m *FetchWhere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchWhere.Unmarshal(m, b)
}
func (m *FetchWhere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchWhere.Marshal(b, m, deterministic)
}
func (m *FetchWhere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchWhere.Merge(m, src)
}
func (m *FetchWhere) XXX_Size() int {
	return xxx_messageInfo_FetchWhere.Size(m)
}
func (m *FetchWhere) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchWhere.DiscardUnknown(m)
}

var xxx_messageInfo_FetchWhere proto.InternalMessageInfo

func (m *FetchWhere) GetPaged() int32 {
	if m != nil {
		return m.Paged
	}
	return 0
}

func (m *FetchWhere) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *FetchWhere) GetSorting() string {
	if m != nil {
		return m.Sorting
	}
	return ""
}

func (m *FetchWhere) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *FetchWhere) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FetchWhere) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FetchWhere) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *FetchWhere) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *FetchWhere) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Fetch struct {
	Id                     int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                   string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AreaId                 int64           `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Addr                   string          `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	Lng                    string          `protobuf:"bytes,5,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat                    string          `protobuf:"bytes,6,opt,name=lat,proto3" json:"lat,omitempty"`
	Tel                    string          `protobuf:"bytes,7,opt,name=tel,proto3" json:"tel,omitempty"`
	Mobile                 string          `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Reception              string          `protobuf:"bytes,9,opt,name=reception,proto3" json:"reception,omitempty"`
	ReceptionRepeatWeeks   int32           `protobuf:"varint,10,opt,name=reception_repeat_weeks,json=receptionRepeatWeeks,proto3" json:"reception_repeat_weeks,omitempty"`
	ReceptionRepeatWeekArr []int32         `protobuf:"varint,11,rep,packed,name=reception_repeat_week_arr,json=receptionRepeatWeekArr,proto3" json:"reception_repeat_week_arr,omitempty"`
	IsFetchTime            bool            `protobuf:"varint,12,opt,name=is_fetch_time,json=isFetchTime,proto3" json:"is_fetch_time,omitempty"`
	Fetch                  string          `protobuf:"bytes,13,opt,name=fetch,proto3" json:"fetch,omitempty"`
	FetchRepeatWeeks       int32           `protobuf:"varint,14,opt,name=fetch_repeat_weeks,json=fetchRepeatWeeks,proto3" json:"fetch_repeat_weeks,omitempty"`
	FetchRepeatWeekArr     []int32         `protobuf:"varint,15,rep,packed,name=fetch_repeat_week_arr,json=fetchRepeatWeekArr,proto3" json:"fetch_repeat_week_arr,omitempty"`
	SubFetchTime           string          `protobuf:"bytes,16,opt,name=sub_fetch_time,json=subFetchTime,proto3" json:"sub_fetch_time,omitempty"`
	Appointment            string          `protobuf:"bytes,17,opt,name=appointment,proto3" json:"appointment,omitempty"`
	AppointmentNum         int32           `protobuf:"varint,18,opt,name=appointment_num,json=appointmentNum,proto3" json:"appointment_num,omitempty"`
	MaxAppointmentNum      int32           `protobuf:"varint,19,opt,name=max_appointment_num,json=maxAppointmentNum,proto3" json:"max_appointment_num,omitempty"`
	Memo                   string          `protobuf:"bytes,20,opt,name=memo,proto3" json:"memo,omitempty"`
	CreatedAt              string          `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt              string          `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ReceptionTimes         []*Times        `protobuf:"bytes,23,rep,name=reception_times,json=receptionTimes,proto3" json:"reception_times,omitempty"`
	FetchTimes             []*Times        `protobuf:"bytes,24,rep,name=fetch_times,json=fetchTimes,proto3" json:"fetch_times,omitempty"`
	Galleries              []*FetchGallery `protobuf:"bytes,25,rep,name=galleries,proto3" json:"galleries,omitempty"`
	Area                   *AreaInfo       `protobuf:"bytes,26,opt,name=area,proto3" json:"area,omitempty"`
	Distance               float32         `protobuf:"fixed32,27,opt,name=distance,proto3" json:"distance,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *Fetch) Reset()         { *m = Fetch{} }
func (m *Fetch) String() string { return proto.CompactTextString(m) }
func (*Fetch) ProtoMessage()    {}
func (*Fetch) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3607b78a5221b8, []int{1}
}

func (m *Fetch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fetch.Unmarshal(m, b)
}
func (m *Fetch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fetch.Marshal(b, m, deterministic)
}
func (m *Fetch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fetch.Merge(m, src)
}
func (m *Fetch) XXX_Size() int {
	return xxx_messageInfo_Fetch.Size(m)
}
func (m *Fetch) XXX_DiscardUnknown() {
	xxx_messageInfo_Fetch.DiscardUnknown(m)
}

var xxx_messageInfo_Fetch proto.InternalMessageInfo

func (m *Fetch) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Fetch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Fetch) GetAreaId() int64 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Fetch) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Fetch) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *Fetch) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *Fetch) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *Fetch) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Fetch) GetReception() string {
	if m != nil {
		return m.Reception
	}
	return ""
}

func (m *Fetch) GetReceptionRepeatWeeks() int32 {
	if m != nil {
		return m.ReceptionRepeatWeeks
	}
	return 0
}

func (m *Fetch) GetReceptionRepeatWeekArr() []int32 {
	if m != nil {
		return m.ReceptionRepeatWeekArr
	}
	return nil
}

func (m *Fetch) GetIsFetchTime() bool {
	if m != nil {
		return m.IsFetchTime
	}
	return false
}

func (m *Fetch) GetFetch() string {
	if m != nil {
		return m.Fetch
	}
	return ""
}

func (m *Fetch) GetFetchRepeatWeeks() int32 {
	if m != nil {
		return m.FetchRepeatWeeks
	}
	return 0
}

func (m *Fetch) GetFetchRepeatWeekArr() []int32 {
	if m != nil {
		return m.FetchRepeatWeekArr
	}
	return nil
}

func (m *Fetch) GetSubFetchTime() string {
	if m != nil {
		return m.SubFetchTime
	}
	return ""
}

func (m *Fetch) GetAppointment() string {
	if m != nil {
		return m.Appointment
	}
	return ""
}

func (m *Fetch) GetAppointmentNum() int32 {
	if m != nil {
		return m.AppointmentNum
	}
	return 0
}

func (m *Fetch) GetMaxAppointmentNum() int32 {
	if m != nil {
		return m.MaxAppointmentNum
	}
	return 0
}

func (m *Fetch) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Fetch) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Fetch) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Fetch) GetReceptionTimes() []*Times {
	if m != nil {
		return m.ReceptionTimes
	}
	return nil
}

func (m *Fetch) GetFetchTimes() []*Times {
	if m != nil {
		return m.FetchTimes
	}
	return nil
}

func (m *Fetch) GetGalleries() []*FetchGallery {
	if m != nil {
		return m.Galleries
	}
	return nil
}

func (m *Fetch) GetArea() *AreaInfo {
	if m != nil {
		return m.Area
	}
	return nil
}

func (m *Fetch) GetDistance() float32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

type FetchResponse struct {
	Entity               *Fetch   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items                []*Fetch `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error                *Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info                 *Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchResponse) Reset()         { *m = FetchResponse{} }
func (m *FetchResponse) String() string { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()    {}
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3607b78a5221b8, []int{2}
}

func (m *FetchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchResponse.Unmarshal(m, b)
}
func (m *FetchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchResponse.Marshal(b, m, deterministic)
}
func (m *FetchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchResponse.Merge(m, src)
}
func (m *FetchResponse) XXX_Size() int {
	return xxx_messageInfo_FetchResponse.Size(m)
}
func (m *FetchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchResponse proto.InternalMessageInfo

func (m *FetchResponse) GetEntity() *Fetch {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *FetchResponse) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func (m *FetchResponse) GetItems() []*Fetch {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *FetchResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *FetchResponse) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*FetchWhere)(nil), "geiqin.srv.tms.FetchWhere")
	proto.RegisterType((*Fetch)(nil), "geiqin.srv.tms.Fetch")
	proto.RegisterType((*FetchResponse)(nil), "geiqin.srv.tms.FetchResponse")
}

func init() {
	proto.RegisterFile("fetch.proto", fileDescriptor_4c3607b78a5221b8)
}

var fileDescriptor_4c3607b78a5221b8 = []byte{
	// 888 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x6d, 0x6f, 0x1b, 0x45,
	0x10, 0xe6, 0x72, 0xf6, 0xc5, 0x1e, 0x27, 0x4e, 0xba, 0x4d, 0xd2, 0xad, 0xdb, 0x8a, 0x93, 0x85,
	0xc4, 0x49, 0x14, 0x4b, 0x18, 0x84, 0x04, 0x12, 0xa8, 0x26, 0x6d, 0x42, 0x51, 0x05, 0xe8, 0x02,
	0xea, 0xc7, 0xd3, 0xc6, 0x37, 0x76, 0x56, 0xf5, 0xbd, 0xb0, 0xbb, 0x6e, 0x9b, 0xfe, 0x34, 0xbe,
	0xf1, 0x0b, 0xf8, 0xce, 0x9f, 0xe0, 0x2f, 0xa0, 0x9d, 0x3d, 0xbf, 0xc4, 0x39, 0x15, 0xa9, 0xfe,
	0xb6, 0x33, 0xcf, 0x33, 0xb3, 0xcf, 0xce, 0xed, 0x3e, 0x36, 0x74, 0x26, 0x68, 0xc6, 0x57, 0x83,
	0x52, 0x15, 0xa6, 0x60, 0xdd, 0x29, 0xca, 0x3f, 0x64, 0x3e, 0xd0, 0xea, 0xf5, 0xc0, 0x64, 0xba,
	0xb7, 0x37, 0x2e, 0xb2, 0xac, 0xc8, 0x1d, 0xda, 0xeb, 0x18, 0x99, 0xa1, 0xae, 0x02, 0x46, 0x75,
	0xe7, 0x62, 0x36, 0x43, 0x75, 0x5d, 0xe5, 0xba, 0x42, 0xa1, 0x78, 0x9e, 0x4f, 0x8a, 0x2a, 0xde,
	0xd7, 0x68, 0x8c, 0xcc, 0xa7, 0x2e, 0xec, 0xff, 0xed, 0x01, 0x9c, 0xd9, 0xaa, 0x97, 0x57, 0xa8,
	0x90, 0x1d, 0x41, 0xb3, 0x14, 0x53, 0x4c, 0xb9, 0x17, 0x7a, 0x51, 0x33, 0x76, 0x01, 0x7b, 0x00,
	0x6d, 0xbb, 0x48, 0xb4, 0x7c, 0x87, 0x7c, 0x87, 0x90, 0x96, 0x4d, 0x5c, 0xc8, 0x77, 0xc8, 0x38,
	0xec, 0xea, 0x42, 0xd9, 0x96, 0xdc, 0x0f, 0xbd, 0xa8, 0x1d, 0x2f, 0x42, 0xd6, 0x83, 0xd6, 0x2b,
	0xbc, 0x7e, 0x53, 0xa8, 0x54, 0xf3, 0x06, 0x41, 0xcb, 0x98, 0x75, 0x61, 0x47, 0xa6, 0xbc, 0x19,
	0x7a, 0x91, 0x1f, 0xef, 0xc8, 0x94, 0x1d, 0x82, 0x2f, 0x53, 0xcd, 0x83, 0xd0, 0x8f, 0xfc, 0xd8,
	0x2e, 0x6d, 0x66, 0x96, 0x4f, 0xf9, 0x2e, 0x15, 0xda, 0x25, 0x65, 0x84, 0xe1, 0xad, 0x2a, 0x23,
	0x0c, 0x63, 0xd0, 0xc8, 0x45, 0x86, 0xbc, 0x4d, 0x29, 0x5a, 0xf7, 0xff, 0xda, 0x85, 0x26, 0x9d,
	0xa8, 0xda, 0xc3, 0x5b, 0xee, 0xb1, 0x60, 0xef, 0xac, 0xd8, 0xec, 0x1e, 0xec, 0xda, 0x01, 0x25,
	0x32, 0x25, 0xf5, 0x7e, 0x1c, 0xd0, 0xbc, 0x88, 0x2c, 0xd2, 0x54, 0x55, 0xc2, 0x69, 0xbd, 0x90,
	0xd4, 0xbc, 0x25, 0x29, 0x58, 0x49, 0x3a, 0x04, 0xdf, 0xe0, 0x6c, 0x21, 0xdb, 0xe0, 0x8c, 0x9d,
	0x40, 0x90, 0x15, 0x97, 0x72, 0x86, 0x95, 0xf2, 0x2a, 0x62, 0x0f, 0xa1, 0xad, 0x70, 0x8c, 0xa5,
	0x91, 0x45, 0x5e, 0x9d, 0x60, 0x95, 0x60, 0x5f, 0xc1, 0xc9, 0x32, 0x48, 0x14, 0x96, 0x28, 0x4c,
	0xf2, 0x06, 0xf1, 0x95, 0xe6, 0x40, 0x1f, 0xe0, 0x68, 0x89, 0xc6, 0x04, 0xbe, 0xb4, 0x18, 0xfb,
	0x06, 0xee, 0xd7, 0x56, 0x25, 0x42, 0x29, 0xde, 0x09, 0xfd, 0xa8, 0x19, 0x9f, 0xd4, 0x14, 0x8e,
	0x94, 0x62, 0x7d, 0xd8, 0x97, 0x3a, 0xa1, 0x1b, 0x94, 0xd8, 0x4b, 0xc5, 0xf7, 0x42, 0x2f, 0x6a,
	0xc5, 0x1d, 0xa9, 0x69, 0x9a, 0xbf, 0xc9, 0x8c, 0xae, 0x07, 0x11, 0xf8, 0x3e, 0xc9, 0x75, 0x01,
	0x7b, 0x0c, 0xee, 0xe2, 0xdd, 0x94, 0xd9, 0x25, 0x99, 0x87, 0x84, 0xac, 0x4b, 0xfc, 0x02, 0x8e,
	0x6f, 0xb1, 0x49, 0xde, 0x01, 0xc9, 0x63, 0x1b, 0x05, 0x56, 0xda, 0x27, 0xd0, 0xd5, 0xf3, 0xcb,
	0x75, 0x6d, 0x87, 0xb4, 0xff, 0x9e, 0x9e, 0x5f, 0xae, 0xc4, 0x85, 0xd0, 0x11, 0x65, 0x59, 0xc8,
	0xdc, 0x64, 0x98, 0x1b, 0x7e, 0x87, 0x28, 0xeb, 0x29, 0xf6, 0x29, 0x1c, 0xac, 0x85, 0x49, 0x3e,
	0xcf, 0x38, 0x23, 0x95, 0xdd, 0xb5, 0xf4, 0xcf, 0xf3, 0x8c, 0x0d, 0xe0, 0x6e, 0x26, 0xde, 0x26,
	0x9b, 0xe4, 0xbb, 0x44, 0xbe, 0x93, 0x89, 0xb7, 0xa3, 0x9b, 0x7c, 0x06, 0x8d, 0x0c, 0xb3, 0x82,
	0x1f, 0xb9, 0xcb, 0x62, 0xd7, 0xec, 0x11, 0xc0, 0x58, 0xa1, 0x30, 0x98, 0x26, 0xc2, 0xf0, 0x63,
	0xf7, 0x7d, 0xab, 0xcc, 0xc8, 0x58, 0x78, 0x5e, 0xa6, 0x0b, 0xf8, 0xc4, 0xc1, 0x55, 0x66, 0x64,
	0xd8, 0xf7, 0x70, 0xb0, 0xfa, 0x90, 0xf4, 0xc6, 0xf9, 0xbd, 0xd0, 0x8f, 0x3a, 0xc3, 0xe3, 0xc1,
	0x4d, 0x3f, 0x18, 0xd8, 0xb3, 0xeb, 0xb8, 0xbb, 0x64, 0x53, 0xcc, 0xbe, 0xae, 0x4c, 0xa4, 0xaa,
	0xe5, 0xef, 0xab, 0x85, 0xc9, 0x62, 0x86, 0x9a, 0x7d, 0x0b, 0xed, 0x29, 0xf9, 0x87, 0x44, 0xcd,
	0xef, 0x53, 0xd5, 0xc3, 0xcd, 0xaa, 0xb3, 0x35, 0x97, 0x89, 0x57, 0x74, 0xf6, 0x18, 0x1a, 0xf6,
	0xf1, 0xf0, 0x5e, 0xe8, 0x45, 0x9d, 0x21, 0xdf, 0x2c, 0x1b, 0x55, 0x46, 0x14, 0x13, 0xcb, 0xba,
	0x43, 0x2a, 0xb5, 0x11, 0xf9, 0x18, 0xf9, 0x83, 0xd0, 0x8b, 0x76, 0xe2, 0x65, 0xdc, 0xff, 0xd7,
	0x83, 0xfd, 0x33, 0x77, 0x0f, 0x74, 0x59, 0xe4, 0x1a, 0xd9, 0xe7, 0x10, 0x60, 0x6e, 0xa4, 0xb9,
	0xa6, 0xf7, 0x5c, 0x73, 0x14, 0x47, 0xaf, 0x48, 0xec, 0x33, 0xe7, 0x63, 0x8a, 0xde, 0x7a, 0x0d,
	0xfb, 0x57, 0x0b, 0x3a, 0x7b, 0x53, 0x96, 0x2c, 0x0d, 0x66, 0x9a, 0xfb, 0xf5, 0x53, 0x72, 0xad,
	0x1d, 0xc7, 0x92, 0x51, 0xa9, 0xc2, 0x19, 0x43, 0x0d, 0xf9, 0x99, 0x05, 0x63, 0xc7, 0x61, 0x11,
	0x34, 0x64, 0x3e, 0x29, 0xc8, 0x31, 0x3a, 0xc3, 0xa3, 0x4d, 0xae, 0x9b, 0x86, 0x65, 0x0c, 0xff,
	0xf1, 0xe0, 0x80, 0xf6, 0x39, 0x9d, 0x4c, 0x2f, 0x50, 0xbd, 0x96, 0x63, 0x64, 0x4f, 0xa0, 0xf1,
	0x4b, 0x89, 0x39, 0xbb, 0xbd, 0x47, 0x56, 0x9a, 0xeb, 0xde, 0xc7, 0x9b, 0xe9, 0x0b, 0xe7, 0xec,
	0x8b, 0x99, 0xf5, 0x3f, 0x62, 0x23, 0x68, 0x9e, 0xce, 0x0a, 0x8d, 0x5b, 0xb4, 0xf8, 0x01, 0x82,
	0xe7, 0x7a, 0x3b, 0x19, 0xc3, 0x3f, 0x1b, 0xb0, 0x47, 0x87, 0x5b, 0x9d, 0x2c, 0x38, 0xa5, 0x97,
	0xc0, 0xea, 0x87, 0xdd, 0x7b, 0x54, 0xff, 0x0d, 0x56, 0xb2, 0x9e, 0x40, 0xf0, 0x3b, 0x3d, 0x96,
	0x0f, 0xee, 0xf0, 0x0c, 0x82, 0xa7, 0x38, 0x43, 0x83, 0xac, 0x57, 0x4b, 0xa5, 0x1f, 0xc4, 0xff,
	0x6f, 0xf3, 0x1d, 0xf8, 0xe7, 0x68, 0xb6, 0x51, 0x71, 0x81, 0x42, 0x8d, 0xaf, 0xb6, 0x53, 0x71,
	0x0a, 0x8d, 0x17, 0x52, 0x9b, 0xed, 0x9a, 0xfc, 0x04, 0xad, 0x73, 0x34, 0xce, 0x07, 0xde, 0xd7,
	0x28, 0xdc, 0xc4, 0x9e, 0x0a, 0x83, 0x76, 0xfb, 0xb5, 0x5e, 0x3f, 0x42, 0xfb, 0x1c, 0xcd, 0x0b,
	0x61, 0x70, 0x4b, 0x55, 0x97, 0x01, 0xfd, 0x51, 0xf9, 0xf2, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x03, 0xe9, 0xfb, 0xd6, 0x15, 0x09, 0x00, 0x00,
}
