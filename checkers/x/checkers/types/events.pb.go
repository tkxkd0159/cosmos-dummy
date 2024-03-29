// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/checkers/events.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventCreateGame struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	GameIndex string `protobuf:"bytes,2,opt,name=game_index,json=gameIndex,proto3" json:"game_index,omitempty"`
	Black     string `protobuf:"bytes,3,opt,name=black,proto3" json:"black,omitempty"`
	Red       string `protobuf:"bytes,4,opt,name=red,proto3" json:"red,omitempty"`
	Wager     uint64 `protobuf:"varint,5,opt,name=wager,proto3" json:"wager,omitempty"`
}

func (m *EventCreateGame) Reset()         { *m = EventCreateGame{} }
func (m *EventCreateGame) String() string { return proto.CompactTextString(m) }
func (*EventCreateGame) ProtoMessage()    {}
func (*EventCreateGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6dd934f465c975c, []int{0}
}
func (m *EventCreateGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateGame.Merge(m, src)
}
func (m *EventCreateGame) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateGame) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateGame.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateGame proto.InternalMessageInfo

func (m *EventCreateGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EventCreateGame) GetGameIndex() string {
	if m != nil {
		return m.GameIndex
	}
	return ""
}

func (m *EventCreateGame) GetBlack() string {
	if m != nil {
		return m.Black
	}
	return ""
}

func (m *EventCreateGame) GetRed() string {
	if m != nil {
		return m.Red
	}
	return ""
}

func (m *EventCreateGame) GetWager() uint64 {
	if m != nil {
		return m.Wager
	}
	return 0
}

type EventMove struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	GameIndex string `protobuf:"bytes,2,opt,name=game_index,json=gameIndex,proto3" json:"game_index,omitempty"`
	CapturedX int64  `protobuf:"varint,3,opt,name=captured_x,json=capturedX,proto3" json:"captured_x,omitempty"`
	CapturedY int64  `protobuf:"varint,4,opt,name=captured_y,json=capturedY,proto3" json:"captured_y,omitempty"`
	Winner    string `protobuf:"bytes,5,opt,name=winner,proto3" json:"winner,omitempty"`
	Board     string `protobuf:"bytes,6,opt,name=board,proto3" json:"board,omitempty"`
}

func (m *EventMove) Reset()         { *m = EventMove{} }
func (m *EventMove) String() string { return proto.CompactTextString(m) }
func (*EventMove) ProtoMessage()    {}
func (*EventMove) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6dd934f465c975c, []int{1}
}
func (m *EventMove) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMove.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMove.Merge(m, src)
}
func (m *EventMove) XXX_Size() int {
	return m.Size()
}
func (m *EventMove) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMove.DiscardUnknown(m)
}

var xxx_messageInfo_EventMove proto.InternalMessageInfo

func (m *EventMove) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EventMove) GetGameIndex() string {
	if m != nil {
		return m.GameIndex
	}
	return ""
}

func (m *EventMove) GetCapturedX() int64 {
	if m != nil {
		return m.CapturedX
	}
	return 0
}

func (m *EventMove) GetCapturedY() int64 {
	if m != nil {
		return m.CapturedY
	}
	return 0
}

func (m *EventMove) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *EventMove) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

type EventRejectGame struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	GameIndex string `protobuf:"bytes,2,opt,name=game_index,json=gameIndex,proto3" json:"game_index,omitempty"`
}

func (m *EventRejectGame) Reset()         { *m = EventRejectGame{} }
func (m *EventRejectGame) String() string { return proto.CompactTextString(m) }
func (*EventRejectGame) ProtoMessage()    {}
func (*EventRejectGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6dd934f465c975c, []int{2}
}
func (m *EventRejectGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRejectGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRejectGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRejectGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRejectGame.Merge(m, src)
}
func (m *EventRejectGame) XXX_Size() int {
	return m.Size()
}
func (m *EventRejectGame) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRejectGame.DiscardUnknown(m)
}

var xxx_messageInfo_EventRejectGame proto.InternalMessageInfo

func (m *EventRejectGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EventRejectGame) GetGameIndex() string {
	if m != nil {
		return m.GameIndex
	}
	return ""
}

type EventForfeitGame struct {
	GameIndex string `protobuf:"bytes,1,opt,name=game_index,json=gameIndex,proto3" json:"game_index,omitempty"`
	Winner    string `protobuf:"bytes,2,opt,name=winner,proto3" json:"winner,omitempty"`
	Board     string `protobuf:"bytes,3,opt,name=board,proto3" json:"board,omitempty"`
}

func (m *EventForfeitGame) Reset()         { *m = EventForfeitGame{} }
func (m *EventForfeitGame) String() string { return proto.CompactTextString(m) }
func (*EventForfeitGame) ProtoMessage()    {}
func (*EventForfeitGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6dd934f465c975c, []int{3}
}
func (m *EventForfeitGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventForfeitGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventForfeitGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventForfeitGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventForfeitGame.Merge(m, src)
}
func (m *EventForfeitGame) XXX_Size() int {
	return m.Size()
}
func (m *EventForfeitGame) XXX_DiscardUnknown() {
	xxx_messageInfo_EventForfeitGame.DiscardUnknown(m)
}

var xxx_messageInfo_EventForfeitGame proto.InternalMessageInfo

func (m *EventForfeitGame) GetGameIndex() string {
	if m != nil {
		return m.GameIndex
	}
	return ""
}

func (m *EventForfeitGame) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *EventForfeitGame) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func init() {
	proto.RegisterType((*EventCreateGame)(nil), "checkers.checkers.EventCreateGame")
	proto.RegisterType((*EventMove)(nil), "checkers.checkers.EventMove")
	proto.RegisterType((*EventRejectGame)(nil), "checkers.checkers.EventRejectGame")
	proto.RegisterType((*EventForfeitGame)(nil), "checkers.checkers.EventForfeitGame")
}

func init() { proto.RegisterFile("checkers/checkers/events.proto", fileDescriptor_c6dd934f465c975c) }

var fileDescriptor_c6dd934f465c975c = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x4e, 0x02, 0x41,
	0x10, 0xc6, 0x59, 0x0e, 0x30, 0x37, 0x8d, 0xb8, 0x31, 0x66, 0x2d, 0xd8, 0x90, 0xab, 0xa8, 0xb0,
	0xe0, 0x0d, 0x34, 0x6a, 0x34, 0xb1, 0xb9, 0x4a, 0x6d, 0xc8, 0xb2, 0x37, 0x22, 0x22, 0xb7, 0x64,
	0x59, 0xe1, 0x78, 0x03, 0x4b, 0x1f, 0xc5, 0xc7, 0xb0, 0xa4, 0xb4, 0x34, 0x77, 0x2f, 0x62, 0x76,
	0xef, 0x8e, 0x3f, 0xc6, 0x8e, 0x6e, 0xbe, 0xef, 0xdb, 0x9d, 0xfc, 0x66, 0x32, 0xc0, 0xe5, 0x33,
	0xca, 0x31, 0xea, 0xd9, 0xd9, 0xba, 0xc0, 0x39, 0xc6, 0x66, 0xd6, 0x9d, 0x6a, 0x65, 0x14, 0x3d,
	0x2a, 0xed, 0x6e, 0x59, 0x04, 0xef, 0x04, 0x0e, 0x2f, 0xed, 0x9b, 0x0b, 0x8d, 0xc2, 0xe0, 0xb5,
	0x98, 0x20, 0x65, 0x70, 0x20, 0xad, 0x52, 0x9a, 0x91, 0x36, 0xe9, 0xf8, 0x61, 0x29, 0x69, 0x0b,
	0x60, 0x28, 0x26, 0xd8, 0x1f, 0xc5, 0x11, 0x26, 0xac, 0xea, 0x42, 0xdf, 0x3a, 0x37, 0xd6, 0xa0,
	0xc7, 0x50, 0x1f, 0xbc, 0x0a, 0x39, 0x66, 0x9e, 0x4b, 0x72, 0x41, 0x9b, 0xe0, 0x69, 0x8c, 0x58,
	0xcd, 0x79, 0xb6, 0xb4, 0xef, 0x16, 0x62, 0x88, 0x9a, 0xd5, 0xdb, 0xa4, 0x53, 0x0b, 0x73, 0x11,
	0x7c, 0x12, 0xf0, 0x1d, 0xca, 0x9d, 0x9a, 0xef, 0x01, 0xd1, 0x02, 0x90, 0x62, 0x6a, 0xde, 0x34,
	0x46, 0xfd, 0xc4, 0x91, 0x78, 0xa1, 0x5f, 0x3a, 0xf7, 0x3b, 0xf1, 0xd2, 0x41, 0x6d, 0xc5, 0x0f,
	0xf4, 0x04, 0x1a, 0x8b, 0x51, 0x1c, 0x17, 0x6c, 0x7e, 0x58, 0x28, 0x37, 0x9a, 0x12, 0x3a, 0x62,
	0x8d, 0x62, 0x34, 0x2b, 0x82, 0xdb, 0x62, 0x79, 0x21, 0xbe, 0xa0, 0x34, 0x7b, 0x2d, 0x2f, 0xe8,
	0x43, 0xd3, 0xf5, 0xba, 0x52, 0xfa, 0x09, 0x47, 0x79, 0xb3, 0xdd, 0x2f, 0xe4, 0xef, 0xa8, 0x1b,
	0xd8, 0xea, 0xff, 0xb0, 0xde, 0x16, 0xec, 0x79, 0xef, 0x2b, 0xe5, 0x64, 0x95, 0x72, 0xf2, 0x93,
	0x72, 0xf2, 0x91, 0xf1, 0xca, 0x2a, 0xe3, 0x95, 0xef, 0x8c, 0x57, 0x1e, 0x4f, 0xd7, 0xe7, 0x92,
	0x6c, 0x2e, 0xc7, 0x2c, 0xa7, 0x38, 0x1b, 0x34, 0xdc, 0xe5, 0xf4, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0xc3, 0x10, 0x71, 0x5b, 0x02, 0x00, 0x00,
}

func (m *EventCreateGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Wager != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Wager))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Red) > 0 {
		i -= len(m.Red)
		copy(dAtA[i:], m.Red)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Red)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Black) > 0 {
		i -= len(m.Black)
		copy(dAtA[i:], m.Black)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Black)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.GameIndex) > 0 {
		i -= len(m.GameIndex)
		copy(dAtA[i:], m.GameIndex)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.GameIndex)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventMove) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMove) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMove) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Board) > 0 {
		i -= len(m.Board)
		copy(dAtA[i:], m.Board)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Board)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0x2a
	}
	if m.CapturedY != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.CapturedY))
		i--
		dAtA[i] = 0x20
	}
	if m.CapturedX != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.CapturedX))
		i--
		dAtA[i] = 0x18
	}
	if len(m.GameIndex) > 0 {
		i -= len(m.GameIndex)
		copy(dAtA[i:], m.GameIndex)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.GameIndex)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventRejectGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRejectGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRejectGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GameIndex) > 0 {
		i -= len(m.GameIndex)
		copy(dAtA[i:], m.GameIndex)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.GameIndex)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventForfeitGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventForfeitGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventForfeitGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Board) > 0 {
		i -= len(m.Board)
		copy(dAtA[i:], m.Board)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Board)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GameIndex) > 0 {
		i -= len(m.GameIndex)
		copy(dAtA[i:], m.GameIndex)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.GameIndex)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreateGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.GameIndex)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Black)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Red)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Wager != 0 {
		n += 1 + sovEvents(uint64(m.Wager))
	}
	return n
}

func (m *EventMove) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.GameIndex)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.CapturedX != 0 {
		n += 1 + sovEvents(uint64(m.CapturedX))
	}
	if m.CapturedY != 0 {
		n += 1 + sovEvents(uint64(m.CapturedY))
	}
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Board)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventRejectGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.GameIndex)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventForfeitGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GameIndex)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Board)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreateGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GameIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Black", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Black = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Red", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Red = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wager", wireType)
			}
			m.Wager = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Wager |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventMove) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventMove: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMove: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GameIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapturedX", wireType)
			}
			m.CapturedX = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CapturedX |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapturedY", wireType)
			}
			m.CapturedY = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CapturedY |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Board", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Board = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventRejectGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRejectGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRejectGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GameIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventForfeitGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventForfeitGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventForfeitGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GameIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Board", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Board = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
