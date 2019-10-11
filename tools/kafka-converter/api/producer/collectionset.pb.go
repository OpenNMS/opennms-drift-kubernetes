// Code generated by protoc-gen-go. DO NOT EDIT.
// source: collectionset.proto

package producer

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

type NumericAttribute_Type int32

const (
	NumericAttribute_GAUGE   NumericAttribute_Type = 0
	NumericAttribute_COUNTER NumericAttribute_Type = 1
)

var NumericAttribute_Type_name = map[int32]string{
	0: "GAUGE",
	1: "COUNTER",
}

var NumericAttribute_Type_value = map[string]int32{
	"GAUGE":   0,
	"COUNTER": 1,
}

func (x NumericAttribute_Type) String() string {
	return proto.EnumName(NumericAttribute_Type_name, int32(x))
}

func (NumericAttribute_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a82f6b1e217a62f, []int{1, 0}
}

type StringAttribute struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringAttribute) Reset()         { *m = StringAttribute{} }
func (m *StringAttribute) String() string { return proto.CompactTextString(m) }
func (*StringAttribute) ProtoMessage()    {}
func (*StringAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a82f6b1e217a62f, []int{0}
}

func (m *StringAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringAttribute.Unmarshal(m, b)
}
func (m *StringAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringAttribute.Marshal(b, m, deterministic)
}
func (m *StringAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringAttribute.Merge(m, src)
}
func (m *StringAttribute) XXX_Size() int {
	return xxx_messageInfo_StringAttribute.Size(m)
}
func (m *StringAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_StringAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_StringAttribute proto.InternalMessageInfo

func (m *StringAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StringAttribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type NumericAttribute struct {
	Group                string                `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value                float64               `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Type                 NumericAttribute_Type `protobuf:"varint,4,opt,name=type,proto3,enum=NumericAttribute_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NumericAttribute) Reset()         { *m = NumericAttribute{} }
func (m *NumericAttribute) String() string { return proto.CompactTextString(m) }
func (*NumericAttribute) ProtoMessage()    {}
func (*NumericAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a82f6b1e217a62f, []int{1}
}

func (m *NumericAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumericAttribute.Unmarshal(m, b)
}
func (m *NumericAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumericAttribute.Marshal(b, m, deterministic)
}
func (m *NumericAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumericAttribute.Merge(m, src)
}
func (m *NumericAttribute) XXX_Size() int {
	return xxx_messageInfo_NumericAttribute.Size(m)
}
func (m *NumericAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_NumericAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_NumericAttribute proto.InternalMessageInfo

func (m *NumericAttribute) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *NumericAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NumericAttribute) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *NumericAttribute) GetType() NumericAttribute_Type {
	if m != nil {
		return m.Type
	}
	return NumericAttribute_GAUGE
}

type NodeLevelResource struct {
	NodeId               int64    `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ForeignSource        string   `protobuf:"bytes,2,opt,name=foreign_source,json=foreignSource,proto3" json:"foreign_source,omitempty"`
	ForeignId            string   `protobuf:"bytes,3,opt,name=foreign_id,json=foreignId,proto3" json:"foreign_id,omitempty"`
	NodeLabel            string   `protobuf:"bytes,4,opt,name=node_label,json=nodeLabel,proto3" json:"node_label,omitempty"`
	Location             string   `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeLevelResource) Reset()         { *m = NodeLevelResource{} }
func (m *NodeLevelResource) String() string { return proto.CompactTextString(m) }
func (*NodeLevelResource) ProtoMessage()    {}
func (*NodeLevelResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a82f6b1e217a62f, []int{2}
}

func (m *NodeLevelResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeLevelResource.Unmarshal(m, b)
}
func (m *NodeLevelResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeLevelResource.Marshal(b, m, deterministic)
}
func (m *NodeLevelResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeLevelResource.Merge(m, src)
}
func (m *NodeLevelResource) XXX_Size() int {
	return xxx_messageInfo_NodeLevelResource.Size(m)
}
func (m *NodeLevelResource) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeLevelResource.DiscardUnknown(m)
}

var xxx_messageInfo_NodeLevelResource proto.InternalMessageInfo

func (m *NodeLevelResource) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeLevelResource) GetForeignSource() string {
	if m != nil {
		return m.ForeignSource
	}
	return ""
}

func (m *NodeLevelResource) GetForeignId() string {
	if m != nil {
		return m.ForeignId
	}
	return ""
}

func (m *NodeLevelResource) GetNodeLabel() string {
	if m != nil {
		return m.NodeLabel
	}
	return ""
}

func (m *NodeLevelResource) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type InterfaceLevelResource struct {
	Node                 *NodeLevelResource `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Instance             string             `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *InterfaceLevelResource) Reset()         { *m = InterfaceLevelResource{} }
func (m *InterfaceLevelResource) String() string { return proto.CompactTextString(m) }
func (*InterfaceLevelResource) ProtoMessage()    {}
func (*InterfaceLevelResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a82f6b1e217a62f, []int{3}
}

func (m *InterfaceLevelResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceLevelResource.Unmarshal(m, b)
}
func (m *InterfaceLevelResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceLevelResource.Marshal(b, m, deterministic)
}
func (m *InterfaceLevelResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceLevelResource.Merge(m, src)
}
func (m *InterfaceLevelResource) XXX_Size() int {
	return xxx_messageInfo_InterfaceLevelResource.Size(m)
}
func (m *InterfaceLevelResource) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceLevelResource.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceLevelResource proto.InternalMessageInfo

func (m *InterfaceLevelResource) GetNode() *NodeLevelResource {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *InterfaceLevelResource) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

type GenericTypeResource struct {
	Node                 *NodeLevelResource `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Type                 string             `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Instance             string             `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GenericTypeResource) Reset()         { *m = GenericTypeResource{} }
func (m *GenericTypeResource) String() string { return proto.CompactTextString(m) }
func (*GenericTypeResource) ProtoMessage()    {}
func (*GenericTypeResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a82f6b1e217a62f, []int{4}
}

func (m *GenericTypeResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericTypeResource.Unmarshal(m, b)
}
func (m *GenericTypeResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericTypeResource.Marshal(b, m, deterministic)
}
func (m *GenericTypeResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericTypeResource.Merge(m, src)
}
func (m *GenericTypeResource) XXX_Size() int {
	return xxx_messageInfo_GenericTypeResource.Size(m)
}
func (m *GenericTypeResource) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericTypeResource.DiscardUnknown(m)
}

var xxx_messageInfo_GenericTypeResource proto.InternalMessageInfo

func (m *GenericTypeResource) GetNode() *NodeLevelResource {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *GenericTypeResource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GenericTypeResource) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

type ResponseTimeResource struct {
	Instance             string   `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseTimeResource) Reset()         { *m = ResponseTimeResource{} }
func (m *ResponseTimeResource) String() string { return proto.CompactTextString(m) }
func (*ResponseTimeResource) ProtoMessage()    {}
func (*ResponseTimeResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a82f6b1e217a62f, []int{5}
}

func (m *ResponseTimeResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseTimeResource.Unmarshal(m, b)
}
func (m *ResponseTimeResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseTimeResource.Marshal(b, m, deterministic)
}
func (m *ResponseTimeResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseTimeResource.Merge(m, src)
}
func (m *ResponseTimeResource) XXX_Size() int {
	return xxx_messageInfo_ResponseTimeResource.Size(m)
}
func (m *ResponseTimeResource) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseTimeResource.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseTimeResource proto.InternalMessageInfo

func (m *ResponseTimeResource) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *ResponseTimeResource) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type CollectionSetResource struct {
	// Types that are valid to be assigned to Resource:
	//	*CollectionSetResource_Node
	//	*CollectionSetResource_Interface
	//	*CollectionSetResource_Generic
	//	*CollectionSetResource_Response
	Resource             isCollectionSetResource_Resource `protobuf_oneof:"resource"`
	String_              []*StringAttribute               `protobuf:"bytes,10,rep,name=string,proto3" json:"string,omitempty"`
	Numeric              []*NumericAttribute              `protobuf:"bytes,11,rep,name=numeric,proto3" json:"numeric,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CollectionSetResource) Reset()         { *m = CollectionSetResource{} }
func (m *CollectionSetResource) String() string { return proto.CompactTextString(m) }
func (*CollectionSetResource) ProtoMessage()    {}
func (*CollectionSetResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a82f6b1e217a62f, []int{6}
}

func (m *CollectionSetResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionSetResource.Unmarshal(m, b)
}
func (m *CollectionSetResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionSetResource.Marshal(b, m, deterministic)
}
func (m *CollectionSetResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionSetResource.Merge(m, src)
}
func (m *CollectionSetResource) XXX_Size() int {
	return xxx_messageInfo_CollectionSetResource.Size(m)
}
func (m *CollectionSetResource) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionSetResource.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionSetResource proto.InternalMessageInfo

type isCollectionSetResource_Resource interface {
	isCollectionSetResource_Resource()
}

type CollectionSetResource_Node struct {
	Node *NodeLevelResource `protobuf:"bytes,1,opt,name=node,proto3,oneof"`
}

type CollectionSetResource_Interface struct {
	Interface *InterfaceLevelResource `protobuf:"bytes,2,opt,name=interface,proto3,oneof"`
}

type CollectionSetResource_Generic struct {
	Generic *GenericTypeResource `protobuf:"bytes,3,opt,name=generic,proto3,oneof"`
}

type CollectionSetResource_Response struct {
	Response *ResponseTimeResource `protobuf:"bytes,4,opt,name=response,proto3,oneof"`
}

func (*CollectionSetResource_Node) isCollectionSetResource_Resource() {}

func (*CollectionSetResource_Interface) isCollectionSetResource_Resource() {}

func (*CollectionSetResource_Generic) isCollectionSetResource_Resource() {}

func (*CollectionSetResource_Response) isCollectionSetResource_Resource() {}

func (m *CollectionSetResource) GetResource() isCollectionSetResource_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *CollectionSetResource) GetNode() *NodeLevelResource {
	if x, ok := m.GetResource().(*CollectionSetResource_Node); ok {
		return x.Node
	}
	return nil
}

func (m *CollectionSetResource) GetInterface() *InterfaceLevelResource {
	if x, ok := m.GetResource().(*CollectionSetResource_Interface); ok {
		return x.Interface
	}
	return nil
}

func (m *CollectionSetResource) GetGeneric() *GenericTypeResource {
	if x, ok := m.GetResource().(*CollectionSetResource_Generic); ok {
		return x.Generic
	}
	return nil
}

func (m *CollectionSetResource) GetResponse() *ResponseTimeResource {
	if x, ok := m.GetResource().(*CollectionSetResource_Response); ok {
		return x.Response
	}
	return nil
}

func (m *CollectionSetResource) GetString_() []*StringAttribute {
	if m != nil {
		return m.String_
	}
	return nil
}

func (m *CollectionSetResource) GetNumeric() []*NumericAttribute {
	if m != nil {
		return m.Numeric
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CollectionSetResource) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CollectionSetResource_Node)(nil),
		(*CollectionSetResource_Interface)(nil),
		(*CollectionSetResource_Generic)(nil),
		(*CollectionSetResource_Response)(nil),
	}
}

type CollectionSet struct {
	Timestamp            int64                    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Resource             []*CollectionSetResource `protobuf:"bytes,2,rep,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CollectionSet) Reset()         { *m = CollectionSet{} }
func (m *CollectionSet) String() string { return proto.CompactTextString(m) }
func (*CollectionSet) ProtoMessage()    {}
func (*CollectionSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a82f6b1e217a62f, []int{7}
}

func (m *CollectionSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionSet.Unmarshal(m, b)
}
func (m *CollectionSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionSet.Marshal(b, m, deterministic)
}
func (m *CollectionSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionSet.Merge(m, src)
}
func (m *CollectionSet) XXX_Size() int {
	return xxx_messageInfo_CollectionSet.Size(m)
}
func (m *CollectionSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionSet.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionSet proto.InternalMessageInfo

func (m *CollectionSet) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *CollectionSet) GetResource() []*CollectionSetResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterEnum("NumericAttribute_Type", NumericAttribute_Type_name, NumericAttribute_Type_value)
	proto.RegisterType((*StringAttribute)(nil), "StringAttribute")
	proto.RegisterType((*NumericAttribute)(nil), "NumericAttribute")
	proto.RegisterType((*NodeLevelResource)(nil), "NodeLevelResource")
	proto.RegisterType((*InterfaceLevelResource)(nil), "InterfaceLevelResource")
	proto.RegisterType((*GenericTypeResource)(nil), "GenericTypeResource")
	proto.RegisterType((*ResponseTimeResource)(nil), "ResponseTimeResource")
	proto.RegisterType((*CollectionSetResource)(nil), "CollectionSetResource")
	proto.RegisterType((*CollectionSet)(nil), "CollectionSet")
}

func init() { proto.RegisterFile("collectionset.proto", fileDescriptor_8a82f6b1e217a62f) }

var fileDescriptor_8a82f6b1e217a62f = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xb6, 0xfc, 0x13, 0x47, 0x23, 0x92, 0x3a, 0x1b, 0x27, 0x11, 0xa5, 0x2d, 0x46, 0xd0, 0x22,
	0x5a, 0x10, 0x45, 0x39, 0xf4, 0xd0, 0x53, 0x12, 0x82, 0x6d, 0x08, 0x2e, 0xac, 0x9d, 0x4b, 0x29,
	0x04, 0x59, 0x9a, 0x98, 0x05, 0x49, 0x2b, 0x56, 0xab, 0x40, 0x5e, 0xa6, 0x6f, 0xd0, 0xb7, 0xe8,
	0x83, 0x95, 0x5d, 0xad, 0xe5, 0xda, 0x31, 0x94, 0xde, 0x34, 0x33, 0xdf, 0x7c, 0x33, 0xf3, 0xcd,
	0x8e, 0xe0, 0x34, 0xe6, 0x69, 0x8a, 0xb1, 0x64, 0x3c, 0x2f, 0x51, 0x06, 0x85, 0xe0, 0x92, 0x7b,
	0x5f, 0xe1, 0xd5, 0x5c, 0x0a, 0x96, 0xaf, 0xae, 0xa4, 0x14, 0x6c, 0x59, 0x49, 0x24, 0x04, 0xba,
	0x79, 0x94, 0xa1, 0x6b, 0x8d, 0x2c, 0xdf, 0xa6, 0xfa, 0x9b, 0x0c, 0xa1, 0xf7, 0x14, 0xa5, 0x15,
	0xba, 0x6d, 0xed, 0xac, 0x0d, 0xef, 0xa7, 0x05, 0x83, 0x59, 0x95, 0xa1, 0x60, 0xf1, 0x26, 0x7d,
	0x08, 0xbd, 0x95, 0xe0, 0x55, 0x61, 0xf2, 0x6b, 0xa3, 0x21, 0x6d, 0xef, 0x23, 0xed, 0x8c, 0x2c,
	0xdf, 0x32, 0xa4, 0xe4, 0x23, 0x74, 0xe5, 0x73, 0x81, 0x6e, 0x77, 0x64, 0xf9, 0xc7, 0xe1, 0x79,
	0xb0, 0x5b, 0x20, 0x58, 0x3c, 0x17, 0x48, 0x35, 0xc6, 0x7b, 0x07, 0x5d, 0x65, 0x11, 0x1b, 0x7a,
	0xe3, 0xab, 0xfb, 0xf1, 0xed, 0xa0, 0x45, 0x1c, 0xe8, 0xdf, 0x7c, 0xbb, 0x9f, 0x2d, 0x6e, 0xe9,
	0xc0, 0xf2, 0x7e, 0x59, 0x70, 0x32, 0xe3, 0x09, 0xde, 0xe1, 0x13, 0xa6, 0x14, 0x4b, 0x5e, 0x89,
	0x18, 0xc9, 0x05, 0xf4, 0x73, 0x9e, 0xe0, 0x03, 0x4b, 0x74, 0x8f, 0x1d, 0x7a, 0xa0, 0xcc, 0x69,
	0x42, 0xde, 0xc3, 0xf1, 0x23, 0x17, 0xc8, 0x56, 0xf9, 0x43, 0x0d, 0x35, 0xed, 0x1e, 0x19, 0xef,
	0xbc, 0xce, 0x7f, 0x0b, 0xb0, 0x86, 0xb1, 0x44, 0x37, 0x6f, 0x53, 0xdb, 0x78, 0xa6, 0x89, 0x0a,
	0x6b, 0xfa, 0x34, 0x5a, 0x62, 0xaa, 0xc7, 0xb0, 0xa9, 0xad, 0x3c, 0x77, 0xca, 0x41, 0x5e, 0xc3,
	0x61, 0xca, 0xe3, 0x48, 0xad, 0xc1, 0xed, 0xe9, 0x60, 0x63, 0x7b, 0x3f, 0xe0, 0x7c, 0x9a, 0x4b,
	0x14, 0x8f, 0x51, 0xbc, 0xd3, 0xf3, 0x07, 0xe8, 0x2a, 0x0a, 0xdd, 0xb0, 0x13, 0x92, 0xe0, 0xc5,
	0x54, 0x54, 0xc7, 0x15, 0x3b, 0xcb, 0x4b, 0x19, 0xe5, 0x4d, 0xf3, 0x8d, 0xed, 0x65, 0x70, 0x3a,
	0xc6, 0x5c, 0x89, 0xa9, 0x25, 0xfc, 0x5f, 0x6a, 0x62, 0x16, 0x63, 0x56, 0xa8, 0xbe, 0xb7, 0xca,
	0x75, 0x76, 0xca, 0xcd, 0x60, 0x48, 0xb1, 0x2c, 0xd4, 0x6b, 0x5b, 0xb0, 0x6c, 0x53, 0xef, 0xef,
	0x1c, 0x6b, 0x3b, 0x67, 0x4b, 0x9c, 0xf6, 0x8e, 0x38, 0xbf, 0xdb, 0x70, 0x76, 0xd3, 0x3c, 0xe1,
	0x39, 0xca, 0x86, 0xd1, 0xff, 0xd7, 0x04, 0x93, 0x96, 0x99, 0xe1, 0x0b, 0xd8, 0x6c, 0x2d, 0xb0,
	0x2e, 0xe0, 0x84, 0x17, 0xc1, 0x7e, 0xc9, 0x27, 0x2d, 0xba, 0xc1, 0x92, 0xcf, 0xd0, 0x5f, 0xd5,
	0xda, 0xe9, 0x39, 0x9d, 0x70, 0x18, 0xec, 0xd1, 0x72, 0xd2, 0xa2, 0x6b, 0x18, 0xb9, 0x84, 0x43,
	0x61, 0xc6, 0xd7, 0x8f, 0xc0, 0x09, 0xcf, 0x82, 0x7d, 0x7a, 0x4c, 0x5a, 0xb4, 0x01, 0x12, 0x1f,
	0x0e, 0x4a, 0x7d, 0x8e, 0x2e, 0x8c, 0x3a, 0xbe, 0x13, 0x0e, 0x82, 0x9d, 0xeb, 0xa4, 0x26, 0x4e,
	0x3e, 0x41, 0x3f, 0xaf, 0x2f, 0xc3, 0x75, 0x34, 0xf4, 0xe4, 0xc5, 0xa5, 0xd0, 0x35, 0xe2, 0x1a,
	0x74, 0x2f, 0xba, 0x9c, 0x17, 0xc1, 0xd1, 0x96, 0x8a, 0xe4, 0x0d, 0xd8, 0x92, 0x65, 0x58, 0xca,
	0x28, 0x2b, 0xcc, 0x41, 0x6c, 0x1c, 0x24, 0xdc, 0xa4, 0xba, 0x6d, 0x5d, 0xe8, 0x3c, 0xd8, 0xbb,
	0x05, 0xda, 0xe0, 0xae, 0xe1, 0xfb, 0x61, 0x21, 0x78, 0x52, 0xc5, 0x28, 0x96, 0x07, 0xfa, 0x3f,
	0x73, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x72, 0xfd, 0x4c, 0x1f, 0x7e, 0x04, 0x00, 0x00,
}