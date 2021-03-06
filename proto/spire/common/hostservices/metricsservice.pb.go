// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metricsservice.proto

package hostservices

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Label struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{0}
}

func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (m *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(m, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetGaugeRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	Val                  float32  `protobuf:"fixed32,2,opt,name=val,proto3" json:"val,omitempty"`
	Labels               []*Label `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGaugeRequest) Reset()         { *m = SetGaugeRequest{} }
func (m *SetGaugeRequest) String() string { return proto.CompactTextString(m) }
func (*SetGaugeRequest) ProtoMessage()    {}
func (*SetGaugeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{1}
}

func (m *SetGaugeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGaugeRequest.Unmarshal(m, b)
}
func (m *SetGaugeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGaugeRequest.Marshal(b, m, deterministic)
}
func (m *SetGaugeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGaugeRequest.Merge(m, src)
}
func (m *SetGaugeRequest) XXX_Size() int {
	return xxx_messageInfo_SetGaugeRequest.Size(m)
}
func (m *SetGaugeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGaugeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetGaugeRequest proto.InternalMessageInfo

func (m *SetGaugeRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SetGaugeRequest) GetVal() float32 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *SetGaugeRequest) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type SetGaugeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGaugeResponse) Reset()         { *m = SetGaugeResponse{} }
func (m *SetGaugeResponse) String() string { return proto.CompactTextString(m) }
func (*SetGaugeResponse) ProtoMessage()    {}
func (*SetGaugeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{2}
}

func (m *SetGaugeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGaugeResponse.Unmarshal(m, b)
}
func (m *SetGaugeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGaugeResponse.Marshal(b, m, deterministic)
}
func (m *SetGaugeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGaugeResponse.Merge(m, src)
}
func (m *SetGaugeResponse) XXX_Size() int {
	return xxx_messageInfo_SetGaugeResponse.Size(m)
}
func (m *SetGaugeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGaugeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetGaugeResponse proto.InternalMessageInfo

type EmitKeyRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	Val                  float32  `protobuf:"fixed32,2,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmitKeyRequest) Reset()         { *m = EmitKeyRequest{} }
func (m *EmitKeyRequest) String() string { return proto.CompactTextString(m) }
func (*EmitKeyRequest) ProtoMessage()    {}
func (*EmitKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{3}
}

func (m *EmitKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmitKeyRequest.Unmarshal(m, b)
}
func (m *EmitKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmitKeyRequest.Marshal(b, m, deterministic)
}
func (m *EmitKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmitKeyRequest.Merge(m, src)
}
func (m *EmitKeyRequest) XXX_Size() int {
	return xxx_messageInfo_EmitKeyRequest.Size(m)
}
func (m *EmitKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmitKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmitKeyRequest proto.InternalMessageInfo

func (m *EmitKeyRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *EmitKeyRequest) GetVal() float32 {
	if m != nil {
		return m.Val
	}
	return 0
}

type EmitKeyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmitKeyResponse) Reset()         { *m = EmitKeyResponse{} }
func (m *EmitKeyResponse) String() string { return proto.CompactTextString(m) }
func (*EmitKeyResponse) ProtoMessage()    {}
func (*EmitKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{4}
}

func (m *EmitKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmitKeyResponse.Unmarshal(m, b)
}
func (m *EmitKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmitKeyResponse.Marshal(b, m, deterministic)
}
func (m *EmitKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmitKeyResponse.Merge(m, src)
}
func (m *EmitKeyResponse) XXX_Size() int {
	return xxx_messageInfo_EmitKeyResponse.Size(m)
}
func (m *EmitKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmitKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmitKeyResponse proto.InternalMessageInfo

type IncrCounterRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	Val                  float32  `protobuf:"fixed32,2,opt,name=val,proto3" json:"val,omitempty"`
	Labels               []*Label `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncrCounterRequest) Reset()         { *m = IncrCounterRequest{} }
func (m *IncrCounterRequest) String() string { return proto.CompactTextString(m) }
func (*IncrCounterRequest) ProtoMessage()    {}
func (*IncrCounterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{5}
}

func (m *IncrCounterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrCounterRequest.Unmarshal(m, b)
}
func (m *IncrCounterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrCounterRequest.Marshal(b, m, deterministic)
}
func (m *IncrCounterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrCounterRequest.Merge(m, src)
}
func (m *IncrCounterRequest) XXX_Size() int {
	return xxx_messageInfo_IncrCounterRequest.Size(m)
}
func (m *IncrCounterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrCounterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncrCounterRequest proto.InternalMessageInfo

func (m *IncrCounterRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *IncrCounterRequest) GetVal() float32 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *IncrCounterRequest) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type IncrCounterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncrCounterResponse) Reset()         { *m = IncrCounterResponse{} }
func (m *IncrCounterResponse) String() string { return proto.CompactTextString(m) }
func (*IncrCounterResponse) ProtoMessage()    {}
func (*IncrCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{6}
}

func (m *IncrCounterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrCounterResponse.Unmarshal(m, b)
}
func (m *IncrCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrCounterResponse.Marshal(b, m, deterministic)
}
func (m *IncrCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrCounterResponse.Merge(m, src)
}
func (m *IncrCounterResponse) XXX_Size() int {
	return xxx_messageInfo_IncrCounterResponse.Size(m)
}
func (m *IncrCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IncrCounterResponse proto.InternalMessageInfo

type AddSampleRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	Val                  float32  `protobuf:"fixed32,2,opt,name=val,proto3" json:"val,omitempty"`
	Labels               []*Label `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSampleRequest) Reset()         { *m = AddSampleRequest{} }
func (m *AddSampleRequest) String() string { return proto.CompactTextString(m) }
func (*AddSampleRequest) ProtoMessage()    {}
func (*AddSampleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{7}
}

func (m *AddSampleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSampleRequest.Unmarshal(m, b)
}
func (m *AddSampleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSampleRequest.Marshal(b, m, deterministic)
}
func (m *AddSampleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSampleRequest.Merge(m, src)
}
func (m *AddSampleRequest) XXX_Size() int {
	return xxx_messageInfo_AddSampleRequest.Size(m)
}
func (m *AddSampleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSampleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddSampleRequest proto.InternalMessageInfo

func (m *AddSampleRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *AddSampleRequest) GetVal() float32 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *AddSampleRequest) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type AddSampleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSampleResponse) Reset()         { *m = AddSampleResponse{} }
func (m *AddSampleResponse) String() string { return proto.CompactTextString(m) }
func (*AddSampleResponse) ProtoMessage()    {}
func (*AddSampleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{8}
}

func (m *AddSampleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSampleResponse.Unmarshal(m, b)
}
func (m *AddSampleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSampleResponse.Marshal(b, m, deterministic)
}
func (m *AddSampleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSampleResponse.Merge(m, src)
}
func (m *AddSampleResponse) XXX_Size() int {
	return xxx_messageInfo_AddSampleResponse.Size(m)
}
func (m *AddSampleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSampleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddSampleResponse proto.InternalMessageInfo

type MeasureSinceRequest struct {
	Key []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	// Unix time in nanoseconds
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Labels               []*Label `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeasureSinceRequest) Reset()         { *m = MeasureSinceRequest{} }
func (m *MeasureSinceRequest) String() string { return proto.CompactTextString(m) }
func (*MeasureSinceRequest) ProtoMessage()    {}
func (*MeasureSinceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{9}
}

func (m *MeasureSinceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeasureSinceRequest.Unmarshal(m, b)
}
func (m *MeasureSinceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeasureSinceRequest.Marshal(b, m, deterministic)
}
func (m *MeasureSinceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeasureSinceRequest.Merge(m, src)
}
func (m *MeasureSinceRequest) XXX_Size() int {
	return xxx_messageInfo_MeasureSinceRequest.Size(m)
}
func (m *MeasureSinceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MeasureSinceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MeasureSinceRequest proto.InternalMessageInfo

func (m *MeasureSinceRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MeasureSinceRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *MeasureSinceRequest) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type MeasureSinceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeasureSinceResponse) Reset()         { *m = MeasureSinceResponse{} }
func (m *MeasureSinceResponse) String() string { return proto.CompactTextString(m) }
func (*MeasureSinceResponse) ProtoMessage()    {}
func (*MeasureSinceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_654b7be56ee6e138, []int{10}
}

func (m *MeasureSinceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeasureSinceResponse.Unmarshal(m, b)
}
func (m *MeasureSinceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeasureSinceResponse.Marshal(b, m, deterministic)
}
func (m *MeasureSinceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeasureSinceResponse.Merge(m, src)
}
func (m *MeasureSinceResponse) XXX_Size() int {
	return xxx_messageInfo_MeasureSinceResponse.Size(m)
}
func (m *MeasureSinceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MeasureSinceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MeasureSinceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Label)(nil), "spire.common.hostservices.Label")
	proto.RegisterType((*SetGaugeRequest)(nil), "spire.common.hostservices.SetGaugeRequest")
	proto.RegisterType((*SetGaugeResponse)(nil), "spire.common.hostservices.SetGaugeResponse")
	proto.RegisterType((*EmitKeyRequest)(nil), "spire.common.hostservices.EmitKeyRequest")
	proto.RegisterType((*EmitKeyResponse)(nil), "spire.common.hostservices.EmitKeyResponse")
	proto.RegisterType((*IncrCounterRequest)(nil), "spire.common.hostservices.IncrCounterRequest")
	proto.RegisterType((*IncrCounterResponse)(nil), "spire.common.hostservices.IncrCounterResponse")
	proto.RegisterType((*AddSampleRequest)(nil), "spire.common.hostservices.AddSampleRequest")
	proto.RegisterType((*AddSampleResponse)(nil), "spire.common.hostservices.AddSampleResponse")
	proto.RegisterType((*MeasureSinceRequest)(nil), "spire.common.hostservices.MeasureSinceRequest")
	proto.RegisterType((*MeasureSinceResponse)(nil), "spire.common.hostservices.MeasureSinceResponse")
}

func init() { proto.RegisterFile("metricsservice.proto", fileDescriptor_654b7be56ee6e138) }

var fileDescriptor_654b7be56ee6e138 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0x95, 0x3a, 0x2d, 0x64, 0x8a, 0xda, 0x74, 0x12, 0x90, 0xf1, 0xc9, 0xf2, 0x29, 0x14,
	0xb0, 0x45, 0x41, 0x02, 0x8e, 0x80, 0x10, 0x42, 0xd0, 0x8b, 0x7d, 0xe3, 0x84, 0xe3, 0x4e, 0x92,
	0x15, 0xb6, 0xd7, 0xec, 0xae, 0x23, 0xe5, 0xcf, 0xf0, 0x5b, 0x91, 0x77, 0x97, 0xe0, 0x04, 0xb0,
	0x8c, 0x2a, 0xe5, 0x36, 0xd9, 0xbc, 0x37, 0xdf, 0xcb, 0xee, 0x53, 0x60, 0x5a, 0x90, 0x12, 0x2c,
	0x93, 0x92, 0xc4, 0x9a, 0x65, 0x14, 0x56, 0x82, 0x2b, 0x8e, 0x0f, 0x65, 0xc5, 0x04, 0x85, 0x19,
	0x2f, 0x0a, 0x5e, 0x86, 0x2b, 0x2e, 0x95, 0xfd, 0x5e, 0x06, 0xcf, 0xe0, 0xf8, 0x73, 0x3a, 0xa7,
	0x1c, 0x11, 0x86, 0x65, 0x5a, 0x90, 0x3b, 0xf0, 0x07, 0xb3, 0x51, 0xac, 0x67, 0x9c, 0xc2, 0xf1,
	0x3a, 0xcd, 0x6b, 0x72, 0x8f, 0xf4, 0xa1, 0xf9, 0x10, 0x70, 0x38, 0x4f, 0x48, 0x7d, 0x48, 0xeb,
	0x25, 0xc5, 0xf4, 0xbd, 0x26, 0xa9, 0x70, 0x0c, 0xce, 0x37, 0xda, 0xb8, 0x03, 0xdf, 0x99, 0x8d,
	0xe2, 0x66, 0x6c, 0x4e, 0xd6, 0x69, 0xae, 0x8d, 0x47, 0x71, 0x33, 0xe2, 0x2b, 0x38, 0xc9, 0x1b,
	0x92, 0x74, 0x1d, 0xdf, 0x99, 0x9d, 0x5e, 0xf9, 0xe1, 0x3f, 0x53, 0x85, 0x3a, 0x52, 0x6c, 0xf5,
	0x01, 0xc2, 0xf8, 0x37, 0x50, 0x56, 0xbc, 0x94, 0x14, 0xbc, 0x80, 0xb3, 0xf7, 0x05, 0x53, 0x9f,
	0x68, 0xf3, 0x1f, 0x19, 0x82, 0x0b, 0x38, 0xdf, 0xba, 0xec, 0x22, 0x01, 0xf8, 0xb1, 0xcc, 0xc4,
	0x3b, 0x5e, 0x97, 0x8a, 0xc4, 0x61, 0x7e, 0xd0, 0x7d, 0x98, 0xec, 0x30, 0x6d, 0x94, 0x0a, 0xc6,
	0x6f, 0x6e, 0x6e, 0x92, 0xb4, 0xa8, 0xf2, 0x03, 0xdd, 0xec, 0x04, 0x2e, 0x5a, 0x44, 0x1b, 0xa3,
	0x86, 0xc9, 0x35, 0xa5, 0xb2, 0x16, 0x94, 0xb0, 0x32, 0xeb, 0x48, 0x82, 0x30, 0x54, 0xac, 0x30,
	0xed, 0x70, 0x62, 0x3d, 0xdf, 0x22, 0xcb, 0x03, 0x98, 0xee, 0x62, 0x4d, 0x9c, 0xab, 0x1f, 0x43,
	0x38, 0xbb, 0x36, 0xad, 0x4e, 0x8c, 0x13, 0x33, 0xb8, 0xfb, 0xab, 0x10, 0x78, 0xd9, 0x01, 0xd8,
	0xab, 0xa9, 0xf7, 0xb8, 0x97, 0xd6, 0x70, 0xf1, 0x2b, 0xdc, 0xb1, 0x5d, 0xc1, 0x47, 0x1d, 0xbe,
	0xdd, 0x16, 0x7a, 0x97, 0x7d, 0xa4, 0x96, 0x90, 0xc3, 0x69, 0xab, 0x06, 0xf8, 0xb4, 0xc3, 0xfa,
	0x67, 0x45, 0xbd, 0xb0, 0xaf, 0xdc, 0xd2, 0x16, 0x30, 0xda, 0xbe, 0x35, 0x76, 0xdd, 0xc4, 0x7e,
	0x07, 0xbd, 0x27, 0xfd, 0xc4, 0x96, 0xc3, 0xe1, 0x5e, 0xfb, 0x1d, 0xb1, 0x2b, 0xe7, 0x5f, 0x7a,
	0xe6, 0x45, 0xbd, 0xf5, 0x06, 0xf8, 0xf6, 0xf5, 0x97, 0x97, 0x4b, 0xa6, 0x56, 0xf5, 0xbc, 0xb1,
	0x44, 0xb2, 0x62, 0x8b, 0x05, 0x45, 0x7a, 0x47, 0xa4, 0xff, 0xfe, 0xec, 0x6c, 0xf6, 0x45, 0xed,
	0x7d, 0xf3, 0x13, 0x2d, 0x78, 0xfe, 0x33, 0x00, 0x00, 0xff, 0xff, 0x38, 0x05, 0x4a, 0x83, 0x37,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	SetGauge(ctx context.Context, in *SetGaugeRequest, opts ...grpc.CallOption) (*SetGaugeResponse, error)
	EmitKey(ctx context.Context, in *EmitKeyRequest, opts ...grpc.CallOption) (*EmitKeyResponse, error)
	IncrCounter(ctx context.Context, in *IncrCounterRequest, opts ...grpc.CallOption) (*IncrCounterResponse, error)
	AddSample(ctx context.Context, in *AddSampleRequest, opts ...grpc.CallOption) (*AddSampleResponse, error)
	MeasureSince(ctx context.Context, in *MeasureSinceRequest, opts ...grpc.CallOption) (*MeasureSinceResponse, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) SetGauge(ctx context.Context, in *SetGaugeRequest, opts ...grpc.CallOption) (*SetGaugeResponse, error) {
	out := new(SetGaugeResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/SetGauge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) EmitKey(ctx context.Context, in *EmitKeyRequest, opts ...grpc.CallOption) (*EmitKeyResponse, error) {
	out := new(EmitKeyResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/EmitKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) IncrCounter(ctx context.Context, in *IncrCounterRequest, opts ...grpc.CallOption) (*IncrCounterResponse, error) {
	out := new(IncrCounterResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/IncrCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) AddSample(ctx context.Context, in *AddSampleRequest, opts ...grpc.CallOption) (*AddSampleResponse, error) {
	out := new(AddSampleResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/AddSample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) MeasureSince(ctx context.Context, in *MeasureSinceRequest, opts ...grpc.CallOption) (*MeasureSinceResponse, error) {
	out := new(MeasureSinceResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/MeasureSince", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	SetGauge(context.Context, *SetGaugeRequest) (*SetGaugeResponse, error)
	EmitKey(context.Context, *EmitKeyRequest) (*EmitKeyResponse, error)
	IncrCounter(context.Context, *IncrCounterRequest) (*IncrCounterResponse, error)
	AddSample(context.Context, *AddSampleRequest) (*AddSampleResponse, error)
	MeasureSince(context.Context, *MeasureSinceRequest) (*MeasureSinceResponse, error)
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_SetGauge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGaugeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).SetGauge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/SetGauge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).SetGauge(ctx, req.(*SetGaugeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_EmitKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).EmitKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/EmitKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).EmitKey(ctx, req.(*EmitKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_IncrCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).IncrCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/IncrCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).IncrCounter(ctx, req.(*IncrCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_AddSample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).AddSample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/AddSample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).AddSample(ctx, req.(*AddSampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_MeasureSince_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasureSinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).MeasureSince(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/MeasureSince",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).MeasureSince(ctx, req.(*MeasureSinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.common.hostservices.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetGauge",
			Handler:    _MetricsService_SetGauge_Handler,
		},
		{
			MethodName: "EmitKey",
			Handler:    _MetricsService_EmitKey_Handler,
		},
		{
			MethodName: "IncrCounter",
			Handler:    _MetricsService_IncrCounter_Handler,
		},
		{
			MethodName: "AddSample",
			Handler:    _MetricsService_AddSample_Handler,
		},
		{
			MethodName: "MeasureSince",
			Handler:    _MetricsService_MeasureSince_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metricsservice.proto",
}
