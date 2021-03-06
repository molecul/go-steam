// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_unified_base.steamclient.proto

package unified

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type EProtoExecutionSite int32

const (
	EProtoExecutionSite_k_EProtoExecutionSiteUnknown     EProtoExecutionSite = 0
	EProtoExecutionSite_k_EProtoExecutionSiteSteamClient EProtoExecutionSite = 2
)

var EProtoExecutionSite_name = map[int32]string{
	0: "k_EProtoExecutionSiteUnknown",
	2: "k_EProtoExecutionSiteSteamClient",
}

var EProtoExecutionSite_value = map[string]int32{
	"k_EProtoExecutionSiteUnknown":     0,
	"k_EProtoExecutionSiteSteamClient": 2,
}

func (x EProtoExecutionSite) Enum() *EProtoExecutionSite {
	p := new(EProtoExecutionSite)
	*p = x
	return p
}

func (x EProtoExecutionSite) String() string {
	return proto.EnumName(EProtoExecutionSite_name, int32(x))
}

func (x *EProtoExecutionSite) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EProtoExecutionSite_value, data, "EProtoExecutionSite")
	if err != nil {
		return err
	}
	*x = EProtoExecutionSite(value)
	return nil
}

func (EProtoExecutionSite) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9054b4d44b9cef17, []int{0}
}

type NoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoResponse) Reset()         { *m = NoResponse{} }
func (m *NoResponse) String() string { return proto.CompactTextString(m) }
func (*NoResponse) ProtoMessage()    {}
func (*NoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9054b4d44b9cef17, []int{0}
}

func (m *NoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoResponse.Unmarshal(m, b)
}
func (m *NoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoResponse.Marshal(b, m, deterministic)
}
func (m *NoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoResponse.Merge(m, src)
}
func (m *NoResponse) XXX_Size() int {
	return xxx_messageInfo_NoResponse.Size(m)
}
func (m *NoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NoResponse proto.InternalMessageInfo

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "description",
	Tag:           "bytes,50000,opt,name=description",
	Filename:      "steammessages_unified_base.steamclient.proto",
}

var E_ServiceDescription = &proto.ExtensionDesc{
	ExtendedType:  (*protobuf.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "service_description",
	Tag:           "bytes,50000,opt,name=service_description",
	Filename:      "steammessages_unified_base.steamclient.proto",
}

var E_ServiceExecutionSite = &proto.ExtensionDesc{
	ExtendedType:  (*protobuf.ServiceOptions)(nil),
	ExtensionType: (*EProtoExecutionSite)(nil),
	Field:         50008,
	Name:          "service_execution_site",
	Tag:           "varint,50008,opt,name=service_execution_site,enum=EProtoExecutionSite,def=0",
	Filename:      "steammessages_unified_base.steamclient.proto",
}

var E_MethodDescription = &proto.ExtensionDesc{
	ExtendedType:  (*protobuf.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "method_description",
	Tag:           "bytes,50000,opt,name=method_description",
	Filename:      "steammessages_unified_base.steamclient.proto",
}

var E_EnumDescription = &proto.ExtensionDesc{
	ExtendedType:  (*protobuf.EnumOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "enum_description",
	Tag:           "bytes,50000,opt,name=enum_description",
	Filename:      "steammessages_unified_base.steamclient.proto",
}

var E_EnumValueDescription = &proto.ExtensionDesc{
	ExtendedType:  (*protobuf.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "enum_value_description",
	Tag:           "bytes,50000,opt,name=enum_value_description",
	Filename:      "steammessages_unified_base.steamclient.proto",
}

func init() {
	proto.RegisterEnum("EProtoExecutionSite", EProtoExecutionSite_name, EProtoExecutionSite_value)
	proto.RegisterType((*NoResponse)(nil), "NoResponse")
	proto.RegisterExtension(E_Description)
	proto.RegisterExtension(E_ServiceDescription)
	proto.RegisterExtension(E_ServiceExecutionSite)
	proto.RegisterExtension(E_MethodDescription)
	proto.RegisterExtension(E_EnumDescription)
	proto.RegisterExtension(E_EnumValueDescription)
}

func init() {
	proto.RegisterFile("steammessages_unified_base.steamclient.proto", fileDescriptor_9054b4d44b9cef17)
}

var fileDescriptor_9054b4d44b9cef17 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x5b, 0xc5, 0x83, 0xa3, 0x48, 0x49, 0xa5, 0x88, 0x54, 0x8d, 0xe2, 0x41, 0x44, 0xb6,
	0x20, 0x1e, 0x24, 0x88, 0x07, 0x4b, 0xc5, 0x8b, 0x1f, 0x18, 0xf4, 0x26, 0x21, 0x4d, 0xa6, 0x71,
	0x69, 0xb2, 0x1b, 0x32, 0xbb, 0xd5, 0xa3, 0x27, 0x7f, 0x9f, 0x47, 0x7f, 0x8e, 0x64, 0x43, 0xc0,
	0x7c, 0xa0, 0xc7, 0x64, 0xde, 0xe7, 0xdd, 0x67, 0x06, 0x4e, 0x48, 0xa1, 0x9f, 0x24, 0x48, 0xe4,
	0x47, 0x48, 0x9e, 0x16, 0x7c, 0xc6, 0x31, 0xf4, 0xa6, 0x3e, 0x21, 0x33, 0xa3, 0x20, 0xe6, 0x28,
	0x14, 0x4b, 0x33, 0xa9, 0xe4, 0xb6, 0x1d, 0x49, 0x19, 0xc5, 0x38, 0x32, 0x5f, 0x53, 0x3d, 0x1b,
	0x85, 0x48, 0x41, 0xc6, 0x53, 0x25, 0xb3, 0x22, 0x71, 0xb0, 0x0e, 0x70, 0x27, 0x1f, 0x91, 0x52,
	0x29, 0x08, 0x8f, 0x5f, 0xa0, 0x3f, 0x79, 0xc8, 0xff, 0x4f, 0xde, 0x31, 0xd0, 0x8a, 0x4b, 0xe1,
	0x72, 0x85, 0x96, 0x0d, 0xc3, 0xb9, 0xd7, 0x32, 0x78, 0x12, 0x73, 0x21, 0xdf, 0x44, 0xaf, 0x63,
	0x1d, 0x82, 0xdd, 0x9a, 0x70, 0x73, 0xa5, 0xb1, 0x51, 0xea, 0x2d, 0x39, 0x67, 0xb0, 0x56, 0x0a,
	0x70, 0x29, 0xac, 0x1d, 0x56, 0xe8, 0xb1, 0x52, 0x8f, 0x5d, 0x73, 0x8c, 0xc3, 0x7b, 0x33, 0xa5,
	0xad, 0xaf, 0xcf, 0x65, 0xbb, 0x7b, 0xb4, 0xea, 0x5c, 0x42, 0x9f, 0x30, 0x5b, 0xf0, 0x00, 0xbd,
	0xdf, 0xf4, 0x5e, 0x83, 0x76, 0x8b, 0x54, 0x9d, 0xd7, 0x30, 0x28, 0x79, 0x2c, 0xdd, 0x3c, 0xca,
	0xf7, 0xfa, 0xb7, 0xe2, 0xdb, 0x54, 0x6c, 0x9c, 0x6e, 0xb2, 0x96, 0xdd, 0x9c, 0x3f, 0x8f, 0xe2,
	0x5c, 0x80, 0x95, 0xa0, 0x7a, 0x95, 0x61, 0xc5, 0x7a, 0xb7, 0xf1, 0xe4, 0xad, 0x09, 0xd5, 0xa5,
	0xcf, 0xa1, 0x87, 0x42, 0x27, 0x15, 0x76, 0xd8, 0x60, 0x27, 0x42, 0x27, 0x75, 0x72, 0x0c, 0x03,
	0x43, 0x2e, 0xfc, 0x58, 0x57, 0x2f, 0xb6, 0xdf, 0xca, 0x3f, 0xe7, 0xb9, 0x5a, 0xc9, 0xd5, 0xca,
	0x4d, 0xf7, 0xa3, 0xdb, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xf6, 0x07, 0xbb, 0x6e, 0x02,
	0x00, 0x00,
}
