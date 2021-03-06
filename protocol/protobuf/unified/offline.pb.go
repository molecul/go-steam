// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_offline.steamclient.proto

package unified

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
// is compatible with the proto package unified is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package unified to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type COffline_GetOfflineLogonTicket_Request struct {
	Priority             *uint32  `protobuf:"varint,1,opt,name=priority" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *COffline_GetOfflineLogonTicket_Request) Reset() {
	*m = COffline_GetOfflineLogonTicket_Request{}
}
func (m *COffline_GetOfflineLogonTicket_Request) String() string { return proto.CompactTextString(m) }
func (*COffline_GetOfflineLogonTicket_Request) ProtoMessage()    {}
func (*COffline_GetOfflineLogonTicket_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_86478a2284c871a9, []int{0}
}

func (m *COffline_GetOfflineLogonTicket_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Request.Unmarshal(m, b)
}
func (m *COffline_GetOfflineLogonTicket_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Request.Marshal(b, m, deterministic)
}
func (m *COffline_GetOfflineLogonTicket_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COffline_GetOfflineLogonTicket_Request.Merge(m, src)
}
func (m *COffline_GetOfflineLogonTicket_Request) XXX_Size() int {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Request.Size(m)
}
func (m *COffline_GetOfflineLogonTicket_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_COffline_GetOfflineLogonTicket_Request.DiscardUnknown(m)
}

var xxx_messageInfo_COffline_GetOfflineLogonTicket_Request proto.InternalMessageInfo

func (m *COffline_GetOfflineLogonTicket_Request) GetPriority() uint32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

type COffline_GetOfflineLogonTicket_Response struct {
	SerializedTicket     []byte   `protobuf:"bytes,1,opt,name=serialized_ticket" json:"serialized_ticket,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *COffline_GetOfflineLogonTicket_Response) Reset() {
	*m = COffline_GetOfflineLogonTicket_Response{}
}
func (m *COffline_GetOfflineLogonTicket_Response) String() string { return proto.CompactTextString(m) }
func (*COffline_GetOfflineLogonTicket_Response) ProtoMessage()    {}
func (*COffline_GetOfflineLogonTicket_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_86478a2284c871a9, []int{1}
}

func (m *COffline_GetOfflineLogonTicket_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Response.Unmarshal(m, b)
}
func (m *COffline_GetOfflineLogonTicket_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Response.Marshal(b, m, deterministic)
}
func (m *COffline_GetOfflineLogonTicket_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COffline_GetOfflineLogonTicket_Response.Merge(m, src)
}
func (m *COffline_GetOfflineLogonTicket_Response) XXX_Size() int {
	return xxx_messageInfo_COffline_GetOfflineLogonTicket_Response.Size(m)
}
func (m *COffline_GetOfflineLogonTicket_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_COffline_GetOfflineLogonTicket_Response.DiscardUnknown(m)
}

var xxx_messageInfo_COffline_GetOfflineLogonTicket_Response proto.InternalMessageInfo

func (m *COffline_GetOfflineLogonTicket_Response) GetSerializedTicket() []byte {
	if m != nil {
		return m.SerializedTicket
	}
	return nil
}

func (m *COffline_GetOfflineLogonTicket_Response) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type COffline_GetUnsignedOfflineLogonTicket_Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *COffline_GetUnsignedOfflineLogonTicket_Request) Reset() {
	*m = COffline_GetUnsignedOfflineLogonTicket_Request{}
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) String() string {
	return proto.CompactTextString(m)
}
func (*COffline_GetUnsignedOfflineLogonTicket_Request) ProtoMessage() {}
func (*COffline_GetUnsignedOfflineLogonTicket_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_86478a2284c871a9, []int{2}
}

func (m *COffline_GetUnsignedOfflineLogonTicket_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request.Unmarshal(m, b)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request.Marshal(b, m, deterministic)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request.Merge(m, src)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) XXX_Size() int {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request.Size(m)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request.DiscardUnknown(m)
}

var xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Request proto.InternalMessageInfo

type COffline_OfflineLogonTicket struct {
	Accountid            *uint32  `protobuf:"varint,1,opt,name=accountid" json:"accountid,omitempty"`
	Rtime32CreationTime  *uint32  `protobuf:"fixed32,2,opt,name=rtime32_creation_time" json:"rtime32_creation_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *COffline_OfflineLogonTicket) Reset()         { *m = COffline_OfflineLogonTicket{} }
func (m *COffline_OfflineLogonTicket) String() string { return proto.CompactTextString(m) }
func (*COffline_OfflineLogonTicket) ProtoMessage()    {}
func (*COffline_OfflineLogonTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor_86478a2284c871a9, []int{3}
}

func (m *COffline_OfflineLogonTicket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COffline_OfflineLogonTicket.Unmarshal(m, b)
}
func (m *COffline_OfflineLogonTicket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COffline_OfflineLogonTicket.Marshal(b, m, deterministic)
}
func (m *COffline_OfflineLogonTicket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COffline_OfflineLogonTicket.Merge(m, src)
}
func (m *COffline_OfflineLogonTicket) XXX_Size() int {
	return xxx_messageInfo_COffline_OfflineLogonTicket.Size(m)
}
func (m *COffline_OfflineLogonTicket) XXX_DiscardUnknown() {
	xxx_messageInfo_COffline_OfflineLogonTicket.DiscardUnknown(m)
}

var xxx_messageInfo_COffline_OfflineLogonTicket proto.InternalMessageInfo

func (m *COffline_OfflineLogonTicket) GetAccountid() uint32 {
	if m != nil && m.Accountid != nil {
		return *m.Accountid
	}
	return 0
}

func (m *COffline_OfflineLogonTicket) GetRtime32CreationTime() uint32 {
	if m != nil && m.Rtime32CreationTime != nil {
		return *m.Rtime32CreationTime
	}
	return 0
}

type COffline_GetUnsignedOfflineLogonTicket_Response struct {
	Ticket               *COffline_OfflineLogonTicket `protobuf:"bytes,1,opt,name=ticket" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *COffline_GetUnsignedOfflineLogonTicket_Response) Reset() {
	*m = COffline_GetUnsignedOfflineLogonTicket_Response{}
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) String() string {
	return proto.CompactTextString(m)
}
func (*COffline_GetUnsignedOfflineLogonTicket_Response) ProtoMessage() {}
func (*COffline_GetUnsignedOfflineLogonTicket_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_86478a2284c871a9, []int{4}
}

func (m *COffline_GetUnsignedOfflineLogonTicket_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response.Unmarshal(m, b)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response.Marshal(b, m, deterministic)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response.Merge(m, src)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) XXX_Size() int {
	return xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response.Size(m)
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response.DiscardUnknown(m)
}

var xxx_messageInfo_COffline_GetUnsignedOfflineLogonTicket_Response proto.InternalMessageInfo

func (m *COffline_GetUnsignedOfflineLogonTicket_Response) GetTicket() *COffline_OfflineLogonTicket {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func init() {
	proto.RegisterType((*COffline_GetOfflineLogonTicket_Request)(nil), "COffline_GetOfflineLogonTicket_Request")
	proto.RegisterType((*COffline_GetOfflineLogonTicket_Response)(nil), "COffline_GetOfflineLogonTicket_Response")
	proto.RegisterType((*COffline_GetUnsignedOfflineLogonTicket_Request)(nil), "COffline_GetUnsignedOfflineLogonTicket_Request")
	proto.RegisterType((*COffline_OfflineLogonTicket)(nil), "COffline_OfflineLogonTicket")
	proto.RegisterType((*COffline_GetUnsignedOfflineLogonTicket_Response)(nil), "COffline_GetUnsignedOfflineLogonTicket_Response")
}

func init() {
	proto.RegisterFile("steammessages_offline.steamclient.proto", fileDescriptor_86478a2284c871a9)
}

var fileDescriptor_86478a2284c871a9 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x4a, 0xe3, 0x40,
	0x14, 0x26, 0x5b, 0xd8, 0xee, 0xce, 0xee, 0xc2, 0x76, 0xa0, 0x90, 0xcd, 0xb6, 0x30, 0xe4, 0xc2,
	0xf6, 0xa2, 0xa4, 0xa5, 0x5e, 0x29, 0x78, 0xa3, 0x48, 0x11, 0x84, 0x82, 0x28, 0x5e, 0x0e, 0x63,
	0x72, 0x12, 0x07, 0xd3, 0x99, 0x3a, 0x73, 0x22, 0xe8, 0x95, 0xf4, 0x55, 0x7c, 0x86, 0x3c, 0x80,
	0x6f, 0x26, 0xf9, 0x41, 0x2b, 0xd6, 0xda, 0xde, 0x25, 0x67, 0xbe, 0xf3, 0xcd, 0xf7, 0x33, 0xa4,
	0x67, 0x11, 0xc4, 0x6c, 0x06, 0xd6, 0x8a, 0x04, 0x2c, 0xd7, 0x71, 0x9c, 0x4a, 0x05, 0x41, 0x39,
	0x0d, 0x53, 0x09, 0x0a, 0x83, 0xb9, 0xd1, 0xa8, 0xbd, 0xc1, 0x7b, 0x60, 0xa6, 0x64, 0x2c, 0x21,
	0xe2, 0x57, 0xc2, 0xae, 0x40, 0xfb, 0xfb, 0x64, 0xe7, 0x68, 0x5a, 0x71, 0xf1, 0x09, 0x60, 0xfd,
	0x79, 0xaa, 0x13, 0xad, 0xce, 0x65, 0x78, 0x03, 0xc8, 0xcf, 0xe0, 0x36, 0x03, 0x8b, 0xf4, 0x2f,
	0xf9, 0x31, 0x37, 0x52, 0x1b, 0x89, 0xf7, 0xae, 0xc3, 0x9c, 0xfe, 0x1f, 0xff, 0x92, 0xf4, 0xbe,
	0xdc, 0xb5, 0x73, 0xad, 0x2c, 0xd0, 0x7f, 0xa4, 0x65, 0xc1, 0x48, 0x91, 0xca, 0x07, 0x88, 0x38,
	0x96, 0xa7, 0x25, 0xcb, 0x6f, 0xda, 0x22, 0x3f, 0xad, 0x4c, 0x94, 0xc0, 0xcc, 0x80, 0xfb, 0xad,
	0x18, 0xf9, 0x23, 0x12, 0x2c, 0x13, 0x5f, 0xa8, 0x02, 0x00, 0xd1, 0xe7, 0xe2, 0xfc, 0x29, 0xf9,
	0xff, 0xba, 0xf1, 0x11, 0x56, 0xdc, 0x21, 0xc2, 0x50, 0x67, 0x0a, 0x65, 0x54, 0x89, 0xa7, 0x5d,
	0xd2, 0x36, 0x28, 0x67, 0xb0, 0x3b, 0xe6, 0xa1, 0x01, 0x81, 0x52, 0x2b, 0x5e, 0xfc, 0x97, 0x12,
	0x9a, 0x3e, 0x27, 0xc3, 0x8d, 0x25, 0xd4, 0x1e, 0x07, 0xe4, 0xfb, 0x92, 0xb1, 0x5f, 0xe3, 0x4e,
	0xb0, 0x46, 0xd2, 0xf8, 0xa9, 0x41, 0x9a, 0xf5, 0x98, 0xe6, 0x0e, 0x69, 0xaf, 0x0c, 0x90, 0xf6,
	0x82, 0xcd, 0xda, 0xf1, 0xfa, 0xc1, 0x86, 0x55, 0xf8, 0x27, 0x8b, 0xdc, 0x3d, 0x9e, 0x00, 0x32,
	0xc1, 0xde, 0x2a, 0x61, 0x42, 0x45, 0xac, 0xb2, 0xc8, 0xea, 0xa7, 0xc5, 0xd2, 0x62, 0x9b, 0x55,
	0x96, 0x58, 0xac, 0x0d, 0xc3, 0x6b, 0x60, 0x61, 0x66, 0x0c, 0x28, 0x64, 0x99, 0x05, 0x43, 0x9f,
	0x1d, 0xd2, 0x5d, 0x1b, 0x0e, 0x1d, 0x6e, 0x59, 0xa4, 0x37, 0x0a, 0xb6, 0x8c, 0xdd, 0x3f, 0x58,
	0xe4, 0xee, 0x5e, 0xe9, 0x47, 0xb1, 0x4c, 0x6d, 0xeb, 0xc1, 0xeb, 0x2c, 0x72, 0xd7, 0xad, 0xf9,
	0x99, 0x05, 0x44, 0xa9, 0x12, 0x5b, 0x24, 0x73, 0x27, 0x43, 0x38, 0x6c, 0x3c, 0x3a, 0xce, 0x4b,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0xd3, 0xb5, 0xf7, 0x7b, 0x03, 0x00, 0x00,
}
