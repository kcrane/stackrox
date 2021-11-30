// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/schedule.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Schedule_IntervalType int32

const (
	Schedule_UNSET   Schedule_IntervalType = 0
	Schedule_DAILY   Schedule_IntervalType = 1
	Schedule_WEEKLY  Schedule_IntervalType = 2
	Schedule_MONTHLY Schedule_IntervalType = 3
)

var Schedule_IntervalType_name = map[int32]string{
	0: "UNSET",
	1: "DAILY",
	2: "WEEKLY",
	3: "MONTHLY",
}

var Schedule_IntervalType_value = map[string]int32{
	"UNSET":   0,
	"DAILY":   1,
	"WEEKLY":  2,
	"MONTHLY": 3,
}

func (x Schedule_IntervalType) String() string {
	return proto.EnumName(Schedule_IntervalType_name, int32(x))
}

func (Schedule_IntervalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_50417c3585ef67fa, []int{0, 0}
}

type Schedule struct {
	IntervalType Schedule_IntervalType `protobuf:"varint,1,opt,name=interval_type,json=intervalType,proto3,enum=storage.Schedule_IntervalType" json:"interval_type,omitempty"`
	Hour         int32                 `protobuf:"varint,2,opt,name=hour,proto3" json:"hour,omitempty"`
	Minute       int32                 `protobuf:"varint,3,opt,name=minute,proto3" json:"minute,omitempty"`
	// Types that are valid to be assigned to Interval:
	//	*Schedule_Weekly
	//	*Schedule_DaysOfWeek_
	//	*Schedule_DaysOfMonth_
	Interval             isSchedule_Interval `protobuf_oneof:"Interval"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_50417c3585ef67fa, []int{0}
}
func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(m, src)
}
func (m *Schedule) XXX_Size() int {
	return m.Size()
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

type isSchedule_Interval interface {
	isSchedule_Interval()
	MarshalTo([]byte) (int, error)
	Size() int
	Clone() isSchedule_Interval
}

type Schedule_Weekly struct {
	Weekly *Schedule_WeeklyInterval `protobuf:"bytes,4,opt,name=weekly,proto3,oneof" json:"weekly,omitempty"`
}
type Schedule_DaysOfWeek_ struct {
	DaysOfWeek *Schedule_DaysOfWeek `protobuf:"bytes,5,opt,name=days_of_week,json=daysOfWeek,proto3,oneof" json:"days_of_week,omitempty"`
}
type Schedule_DaysOfMonth_ struct {
	DaysOfMonth *Schedule_DaysOfMonth `protobuf:"bytes,6,opt,name=days_of_month,json=daysOfMonth,proto3,oneof" json:"days_of_month,omitempty"`
}

func (*Schedule_Weekly) isSchedule_Interval() {}
func (m *Schedule_Weekly) Clone() isSchedule_Interval {
	if m == nil {
		return nil
	}
	cloned := new(Schedule_Weekly)
	*cloned = *m

	cloned.Weekly = m.Weekly.Clone()
	return cloned
}
func (*Schedule_DaysOfWeek_) isSchedule_Interval() {}
func (m *Schedule_DaysOfWeek_) Clone() isSchedule_Interval {
	if m == nil {
		return nil
	}
	cloned := new(Schedule_DaysOfWeek_)
	*cloned = *m

	cloned.DaysOfWeek = m.DaysOfWeek.Clone()
	return cloned
}
func (*Schedule_DaysOfMonth_) isSchedule_Interval() {}
func (m *Schedule_DaysOfMonth_) Clone() isSchedule_Interval {
	if m == nil {
		return nil
	}
	cloned := new(Schedule_DaysOfMonth_)
	*cloned = *m

	cloned.DaysOfMonth = m.DaysOfMonth.Clone()
	return cloned
}

func (m *Schedule) GetInterval() isSchedule_Interval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *Schedule) GetIntervalType() Schedule_IntervalType {
	if m != nil {
		return m.IntervalType
	}
	return Schedule_UNSET
}

func (m *Schedule) GetHour() int32 {
	if m != nil {
		return m.Hour
	}
	return 0
}

func (m *Schedule) GetMinute() int32 {
	if m != nil {
		return m.Minute
	}
	return 0
}

func (m *Schedule) GetWeekly() *Schedule_WeeklyInterval {
	if x, ok := m.GetInterval().(*Schedule_Weekly); ok {
		return x.Weekly
	}
	return nil
}

func (m *Schedule) GetDaysOfWeek() *Schedule_DaysOfWeek {
	if x, ok := m.GetInterval().(*Schedule_DaysOfWeek_); ok {
		return x.DaysOfWeek
	}
	return nil
}

func (m *Schedule) GetDaysOfMonth() *Schedule_DaysOfMonth {
	if x, ok := m.GetInterval().(*Schedule_DaysOfMonth_); ok {
		return x.DaysOfMonth
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Schedule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Schedule_Weekly)(nil),
		(*Schedule_DaysOfWeek_)(nil),
		(*Schedule_DaysOfMonth_)(nil),
	}
}

func (m *Schedule) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Schedule) Clone() *Schedule {
	if m == nil {
		return nil
	}
	cloned := new(Schedule)
	*cloned = *m

	if m.Interval != nil {
		cloned.Interval = m.Interval.Clone()
	}
	return cloned
}

type Schedule_WeeklyInterval struct {
	Day                  int32    `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Schedule_WeeklyInterval) Reset()         { *m = Schedule_WeeklyInterval{} }
func (m *Schedule_WeeklyInterval) String() string { return proto.CompactTextString(m) }
func (*Schedule_WeeklyInterval) ProtoMessage()    {}
func (*Schedule_WeeklyInterval) Descriptor() ([]byte, []int) {
	return fileDescriptor_50417c3585ef67fa, []int{0, 0}
}
func (m *Schedule_WeeklyInterval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schedule_WeeklyInterval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schedule_WeeklyInterval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schedule_WeeklyInterval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule_WeeklyInterval.Merge(m, src)
}
func (m *Schedule_WeeklyInterval) XXX_Size() int {
	return m.Size()
}
func (m *Schedule_WeeklyInterval) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule_WeeklyInterval.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule_WeeklyInterval proto.InternalMessageInfo

func (m *Schedule_WeeklyInterval) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *Schedule_WeeklyInterval) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Schedule_WeeklyInterval) Clone() *Schedule_WeeklyInterval {
	if m == nil {
		return nil
	}
	cloned := new(Schedule_WeeklyInterval)
	*cloned = *m

	return cloned
}

// Sunday = 0, Monday = 1, .... Saturday =  6
type Schedule_DaysOfWeek struct {
	Days                 []int32  `protobuf:"varint,1,rep,packed,name=days,proto3" json:"days,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Schedule_DaysOfWeek) Reset()         { *m = Schedule_DaysOfWeek{} }
func (m *Schedule_DaysOfWeek) String() string { return proto.CompactTextString(m) }
func (*Schedule_DaysOfWeek) ProtoMessage()    {}
func (*Schedule_DaysOfWeek) Descriptor() ([]byte, []int) {
	return fileDescriptor_50417c3585ef67fa, []int{0, 1}
}
func (m *Schedule_DaysOfWeek) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schedule_DaysOfWeek) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schedule_DaysOfWeek.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schedule_DaysOfWeek) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule_DaysOfWeek.Merge(m, src)
}
func (m *Schedule_DaysOfWeek) XXX_Size() int {
	return m.Size()
}
func (m *Schedule_DaysOfWeek) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule_DaysOfWeek.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule_DaysOfWeek proto.InternalMessageInfo

func (m *Schedule_DaysOfWeek) GetDays() []int32 {
	if m != nil {
		return m.Days
	}
	return nil
}

func (m *Schedule_DaysOfWeek) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Schedule_DaysOfWeek) Clone() *Schedule_DaysOfWeek {
	if m == nil {
		return nil
	}
	cloned := new(Schedule_DaysOfWeek)
	*cloned = *m

	if m.Days != nil {
		cloned.Days = make([]int32, len(m.Days))
		copy(cloned.Days, m.Days)
	}
	return cloned
}

// 1 for 1st, 2 for 2nd .... 31 for 31st
type Schedule_DaysOfMonth struct {
	Days                 []int32  `protobuf:"varint,1,rep,packed,name=days,proto3" json:"days,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Schedule_DaysOfMonth) Reset()         { *m = Schedule_DaysOfMonth{} }
func (m *Schedule_DaysOfMonth) String() string { return proto.CompactTextString(m) }
func (*Schedule_DaysOfMonth) ProtoMessage()    {}
func (*Schedule_DaysOfMonth) Descriptor() ([]byte, []int) {
	return fileDescriptor_50417c3585ef67fa, []int{0, 2}
}
func (m *Schedule_DaysOfMonth) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schedule_DaysOfMonth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schedule_DaysOfMonth.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schedule_DaysOfMonth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule_DaysOfMonth.Merge(m, src)
}
func (m *Schedule_DaysOfMonth) XXX_Size() int {
	return m.Size()
}
func (m *Schedule_DaysOfMonth) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule_DaysOfMonth.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule_DaysOfMonth proto.InternalMessageInfo

func (m *Schedule_DaysOfMonth) GetDays() []int32 {
	if m != nil {
		return m.Days
	}
	return nil
}

func (m *Schedule_DaysOfMonth) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Schedule_DaysOfMonth) Clone() *Schedule_DaysOfMonth {
	if m == nil {
		return nil
	}
	cloned := new(Schedule_DaysOfMonth)
	*cloned = *m

	if m.Days != nil {
		cloned.Days = make([]int32, len(m.Days))
		copy(cloned.Days, m.Days)
	}
	return cloned
}

func init() {
	proto.RegisterEnum("storage.Schedule_IntervalType", Schedule_IntervalType_name, Schedule_IntervalType_value)
	proto.RegisterType((*Schedule)(nil), "storage.Schedule")
	proto.RegisterType((*Schedule_WeeklyInterval)(nil), "storage.Schedule.WeeklyInterval")
	proto.RegisterType((*Schedule_DaysOfWeek)(nil), "storage.Schedule.DaysOfWeek")
	proto.RegisterType((*Schedule_DaysOfMonth)(nil), "storage.Schedule.DaysOfMonth")
}

func init() { proto.RegisterFile("storage/schedule.proto", fileDescriptor_50417c3585ef67fa) }

var fileDescriptor_50417c3585ef67fa = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x4e, 0xc2, 0x40,
	0x18, 0xec, 0x52, 0x5a, 0xf0, 0x2b, 0x92, 0x66, 0x0f, 0xa4, 0x12, 0x6d, 0x2a, 0x27, 0x4e, 0x35,
	0x41, 0x4f, 0x26, 0x26, 0x8a, 0x90, 0x94, 0xc8, 0x4f, 0x52, 0x30, 0x04, 0x2f, 0xa4, 0xd2, 0x45,
	0x1a, 0x7e, 0x96, 0xb4, 0x45, 0xed, 0x9b, 0xf8, 0x48, 0x1e, 0x7d, 0x04, 0x83, 0x6f, 0xe0, 0x13,
	0x98, 0x5d, 0x5b, 0xc1, 0xa0, 0xb7, 0x99, 0xd9, 0xd9, 0xf9, 0x76, 0x27, 0x1f, 0x14, 0x82, 0x90,
	0xfa, 0xce, 0x03, 0x39, 0x09, 0x46, 0x13, 0xe2, 0xae, 0x66, 0xc4, 0x5c, 0xfa, 0x34, 0xa4, 0x38,
	0x13, 0xeb, 0xa5, 0x4f, 0x11, 0xb2, 0xdd, 0xf8, 0x0c, 0x5f, 0xc3, 0xbe, 0xb7, 0x08, 0x89, 0xff,
	0xe8, 0xcc, 0x86, 0x61, 0xb4, 0x24, 0x1a, 0x32, 0x50, 0x39, 0x5f, 0xd1, 0xcd, 0xd8, 0x6d, 0x26,
	0x4e, 0xb3, 0x11, 0xdb, 0x7a, 0xd1, 0x92, 0xd8, 0x39, 0x6f, 0x8b, 0x61, 0x0c, 0xe9, 0x09, 0x5d,
	0xf9, 0x5a, 0xca, 0x40, 0x65, 0xc9, 0xe6, 0x18, 0x17, 0x40, 0x9e, 0x7b, 0x8b, 0x55, 0x48, 0x34,
	0x91, 0xab, 0x31, 0xc3, 0xe7, 0x20, 0x3f, 0x11, 0x32, 0x9d, 0x45, 0x5a, 0xda, 0x40, 0x65, 0xa5,
	0x62, 0xec, 0x4e, 0xea, 0xf3, 0xf3, 0x64, 0x9e, 0x25, 0xd8, 0xf1, 0x0d, 0x7c, 0x09, 0x39, 0xd7,
	0x89, 0x82, 0x21, 0x1d, 0x0f, 0x99, 0xa2, 0x49, 0x3c, 0xe1, 0x70, 0x37, 0xa1, 0xe6, 0x44, 0x41,
	0x67, 0xcc, 0x72, 0x2c, 0xc1, 0x06, 0xf7, 0x87, 0xb1, 0xef, 0x26, 0x09, 0x73, 0xba, 0x08, 0x27,
	0x9a, 0xcc, 0x23, 0x8e, 0xfe, 0x8b, 0x68, 0x31, 0x93, 0x25, 0xd8, 0x8a, 0xbb, 0xa1, 0xc5, 0x12,
	0xe4, 0x7f, 0x3f, 0x11, 0xab, 0x20, 0xba, 0x4e, 0xc4, 0xbb, 0x93, 0x6c, 0x06, 0x8b, 0x06, 0xc0,
	0xe6, 0x11, 0xac, 0x20, 0x16, 0xa0, 0x21, 0x43, 0x64, 0x05, 0x31, 0x5c, 0x3c, 0x06, 0x65, 0x6b,
	0xc6, 0x5f, 0x96, 0xd2, 0x05, 0xe4, 0xb6, 0x5b, 0xc7, 0x7b, 0x20, 0xdd, 0xb6, 0xbb, 0xf5, 0x9e,
	0x2a, 0x30, 0x58, 0xbb, 0x6a, 0x34, 0x07, 0x2a, 0xc2, 0x00, 0x72, 0xbf, 0x5e, 0xbf, 0x69, 0x0e,
	0xd4, 0x14, 0x56, 0x20, 0xd3, 0xea, 0xb4, 0x7b, 0x56, 0x73, 0xa0, 0x8a, 0x55, 0x80, 0x6c, 0x72,
	0xbd, 0x7a, 0xf6, 0xba, 0xd6, 0xd1, 0xdb, 0x5a, 0x47, 0xef, 0x6b, 0x1d, 0xbd, 0x7c, 0xe8, 0x02,
	0x1c, 0x78, 0xd4, 0x0c, 0x42, 0x67, 0x34, 0xf5, 0xe9, 0xf3, 0xf7, 0x8a, 0x24, 0x25, 0xdc, 0x25,
	0xab, 0x72, 0x2f, 0x73, 0xfd, 0xf4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x35, 0x0f, 0x56, 0x54,
	0x02, 0x00, 0x00,
}

func (m *Schedule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schedule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schedule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Interval != nil {
		{
			size := m.Interval.Size()
			i -= size
			if _, err := m.Interval.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Minute != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.Minute))
		i--
		dAtA[i] = 0x18
	}
	if m.Hour != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.Hour))
		i--
		dAtA[i] = 0x10
	}
	if m.IntervalType != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.IntervalType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Schedule_Weekly) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schedule_Weekly) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Weekly != nil {
		{
			size, err := m.Weekly.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSchedule(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Schedule_DaysOfWeek_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schedule_DaysOfWeek_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DaysOfWeek != nil {
		{
			size, err := m.DaysOfWeek.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSchedule(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Schedule_DaysOfMonth_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schedule_DaysOfMonth_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DaysOfMonth != nil {
		{
			size, err := m.DaysOfMonth.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSchedule(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Schedule_WeeklyInterval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schedule_WeeklyInterval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schedule_WeeklyInterval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Day != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.Day))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Schedule_DaysOfWeek) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schedule_DaysOfWeek) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schedule_DaysOfWeek) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Days) > 0 {
		dAtA5 := make([]byte, len(m.Days)*10)
		var j4 int
		for _, num1 := range m.Days {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		i -= j4
		copy(dAtA[i:], dAtA5[:j4])
		i = encodeVarintSchedule(dAtA, i, uint64(j4))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Schedule_DaysOfMonth) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schedule_DaysOfMonth) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schedule_DaysOfMonth) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Days) > 0 {
		dAtA7 := make([]byte, len(m.Days)*10)
		var j6 int
		for _, num1 := range m.Days {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA7[j6] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j6++
			}
			dAtA7[j6] = uint8(num)
			j6++
		}
		i -= j6
		copy(dAtA[i:], dAtA7[:j6])
		i = encodeVarintSchedule(dAtA, i, uint64(j6))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchedule(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchedule(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Schedule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IntervalType != 0 {
		n += 1 + sovSchedule(uint64(m.IntervalType))
	}
	if m.Hour != 0 {
		n += 1 + sovSchedule(uint64(m.Hour))
	}
	if m.Minute != 0 {
		n += 1 + sovSchedule(uint64(m.Minute))
	}
	if m.Interval != nil {
		n += m.Interval.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Schedule_Weekly) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Weekly != nil {
		l = m.Weekly.Size()
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}
func (m *Schedule_DaysOfWeek_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DaysOfWeek != nil {
		l = m.DaysOfWeek.Size()
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}
func (m *Schedule_DaysOfMonth_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DaysOfMonth != nil {
		l = m.DaysOfMonth.Size()
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}
func (m *Schedule_WeeklyInterval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Day != 0 {
		n += 1 + sovSchedule(uint64(m.Day))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Schedule_DaysOfWeek) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Days) > 0 {
		l = 0
		for _, e := range m.Days {
			l += sovSchedule(uint64(e))
		}
		n += 1 + sovSchedule(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Schedule_DaysOfMonth) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Days) > 0 {
		l = 0
		for _, e := range m.Days {
			l += sovSchedule(uint64(e))
		}
		n += 1 + sovSchedule(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSchedule(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSchedule(x uint64) (n int) {
	return sovSchedule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Schedule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Schedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntervalType", wireType)
			}
			m.IntervalType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntervalType |= Schedule_IntervalType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hour", wireType)
			}
			m.Hour = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hour |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minute", wireType)
			}
			m.Minute = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Minute |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weekly", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Schedule_WeeklyInterval{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Interval = &Schedule_Weekly{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaysOfWeek", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Schedule_DaysOfWeek{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Interval = &Schedule_DaysOfWeek_{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaysOfMonth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Schedule_DaysOfMonth{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Interval = &Schedule_DaysOfMonth_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Schedule_WeeklyInterval) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WeeklyInterval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WeeklyInterval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Day", wireType)
			}
			m.Day = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Day |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Schedule_DaysOfWeek) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DaysOfWeek: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DaysOfWeek: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSchedule
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Days = append(m.Days, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSchedule
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthSchedule
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthSchedule
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Days) == 0 {
					m.Days = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSchedule
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Days = append(m.Days, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Days", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Schedule_DaysOfMonth) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DaysOfMonth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DaysOfMonth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSchedule
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Days = append(m.Days, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSchedule
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthSchedule
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthSchedule
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Days) == 0 {
					m.Days = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSchedule
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Days = append(m.Days, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Days", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSchedule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchedule
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSchedule
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchedule
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchedule
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchedule        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchedule          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchedule = fmt.Errorf("proto: unexpected end of group")
)
