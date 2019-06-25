// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calcpb/calc.proto

package calcpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Numbers struct {
	Numbers              []int32  `protobuf:"varint,1,rep,packed,name=numbers,proto3" json:"numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Numbers) Reset()         { *m = Numbers{} }
func (m *Numbers) String() string { return proto.CompactTextString(m) }
func (*Numbers) ProtoMessage()    {}
func (*Numbers) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{0}
}

func (m *Numbers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Numbers.Unmarshal(m, b)
}
func (m *Numbers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Numbers.Marshal(b, m, deterministic)
}
func (m *Numbers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Numbers.Merge(m, src)
}
func (m *Numbers) XXX_Size() int {
	return xxx_messageInfo_Numbers.Size(m)
}
func (m *Numbers) XXX_DiscardUnknown() {
	xxx_messageInfo_Numbers.DiscardUnknown(m)
}

var xxx_messageInfo_Numbers proto.InternalMessageInfo

func (m *Numbers) GetNumbers() []int32 {
	if m != nil {
		return m.Numbers
	}
	return nil
}

type SumRequest struct {
	Numbers              *Numbers `protobuf:"bytes,1,opt,name=numbers,proto3" json:"numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetNumbers() *Numbers {
	if m != nil {
		return m.Numbers
	}
	return nil
}

type SumResponse struct {
	Result               string   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{2}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type PrimeNumberDecompositionRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionRequest) Reset()         { *m = PrimeNumberDecompositionRequest{} }
func (m *PrimeNumberDecompositionRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionRequest) ProtoMessage()    {}
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{3}
}

func (m *PrimeNumberDecompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionRequest.Merge(m, src)
}
func (m *PrimeNumberDecompositionRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Size(m)
}
func (m *PrimeNumberDecompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionRequest proto.InternalMessageInfo

func (m *PrimeNumberDecompositionRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionResponse) Reset()         { *m = PrimeNumberDecompositionResponse{} }
func (m *PrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{4}
}

func (m *PrimeNumberDecompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionResponse.Merge(m, src)
}
func (m *PrimeNumberDecompositionResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Size(m)
}
func (m *PrimeNumberDecompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionResponse proto.InternalMessageInfo

func (m *PrimeNumberDecompositionResponse) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAverageRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{5}
}

func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(m, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	Number               float32  `protobuf:"fixed32,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{6}
}

func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(m, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetNumber() float32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximumRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumRequest) Reset()         { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()    {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{7}
}

func (m *FindMaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumRequest.Unmarshal(m, b)
}
func (m *FindMaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumRequest.Merge(m, src)
}
func (m *FindMaximumRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximumRequest.Size(m)
}
func (m *FindMaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumRequest proto.InternalMessageInfo

func (m *FindMaximumRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximumResponse struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumResponse) Reset()         { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()    {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{8}
}

func (m *FindMaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumResponse.Unmarshal(m, b)
}
func (m *FindMaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumResponse.Merge(m, src)
}
func (m *FindMaximumResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximumResponse.Size(m)
}
func (m *FindMaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumResponse proto.InternalMessageInfo

func (m *FindMaximumResponse) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{9}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	NumberRoot           float32  `protobuf:"fixed32,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootResponse) Reset()         { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()    {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{10}
}

func (m *SquareRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootResponse.Unmarshal(m, b)
}
func (m *SquareRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootResponse.Marshal(b, m, deterministic)
}
func (m *SquareRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootResponse.Merge(m, src)
}
func (m *SquareRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquareRootResponse.Size(m)
}
func (m *SquareRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootResponse proto.InternalMessageInfo

func (m *SquareRootResponse) GetNumberRoot() float32 {
	if m != nil {
		return m.NumberRoot
	}
	return 0
}

func init() {
	proto.RegisterType((*Numbers)(nil), "calc.Numbers")
	proto.RegisterType((*SumRequest)(nil), "calc.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calc.SumResponse")
	proto.RegisterType((*PrimeNumberDecompositionRequest)(nil), "calc.PrimeNumberDecompositionRequest")
	proto.RegisterType((*PrimeNumberDecompositionResponse)(nil), "calc.PrimeNumberDecompositionResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calc.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calc.ComputeAverageResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "calc.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calc.FindMaximumResponse")
	proto.RegisterType((*SquareRootRequest)(nil), "calc.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "calc.SquareRootResponse")
}

func init() { proto.RegisterFile("calculator/calcpb/calc.proto", fileDescriptor_cfc74e58bc0fa04b) }

var fileDescriptor_cfc74e58bc0fa04b = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0xae, 0xa1, 0x98, 0x76, 0xac, 0x56, 0x65, 0xaa, 0x52, 0xd7, 0x45, 0x02, 0xb9, 0xa2, 0xb5,
	0x54, 0x0a, 0x88, 0x8a, 0x43, 0x72, 0x23, 0x89, 0xa2, 0x5c, 0xf2, 0x23, 0xfb, 0x96, 0x4b, 0x64,
	0x9c, 0x55, 0x64, 0x89, 0xf5, 0x9a, 0xf5, 0x2e, 0xca, 0x1b, 0xe4, 0xb5, 0x23, 0xbc, 0x76, 0x6c,
	0xc0, 0xc8, 0x27, 0xcf, 0xec, 0x7c, 0x3f, 0xab, 0xf9, 0xbc, 0xd0, 0x0b, 0xfc, 0x55, 0x20, 0x57,
	0xbe, 0x60, 0x7c, 0xb2, 0x2d, 0xe3, 0x65, 0xfa, 0x19, 0xc7, 0x9c, 0x09, 0x86, 0xef, 0xb7, 0xb5,
	0xfd, 0x0b, 0xda, 0x37, 0x92, 0x2e, 0x09, 0x4f, 0xd0, 0x84, 0x76, 0xa4, 0x4a, 0x53, 0x1b, 0x34,
	0x9d, 0x96, 0x9b, 0xb7, 0xf6, 0x1c, 0xc0, 0x93, 0xd4, 0x25, 0x6b, 0x49, 0x12, 0x81, 0x7f, 0xca,
	0x38, 0xcd, 0x31, 0x66, 0x9f, 0xc6, 0xa9, 0x6c, 0xa6, 0x53, 0xd0, 0x86, 0x60, 0xa4, 0xb4, 0x24,
	0x66, 0x51, 0x42, 0xb0, 0x0b, 0x3a, 0x27, 0x89, 0x5c, 0x09, 0xb3, 0x31, 0xd0, 0x9c, 0x8f, 0x6e,
	0xd6, 0xd9, 0x27, 0xd0, 0xbf, 0xe3, 0x21, 0x25, 0x8a, 0x7f, 0x41, 0x02, 0x46, 0x63, 0x96, 0x84,
	0x22, 0x64, 0x51, 0x6e, 0xd9, 0x05, 0x5d, 0x89, 0xa6, 0x8e, 0x2d, 0x37, 0xeb, 0xec, 0x53, 0x18,
	0x1c, 0xa7, 0x16, 0xb6, 0x95, 0xdc, 0x09, 0x7c, 0x3b, 0x67, 0x34, 0x96, 0x82, 0x2c, 0x36, 0x84,
	0xfb, 0x4f, 0xa4, 0xce, 0x6c, 0x0a, 0xdd, 0x7d, 0x42, 0xa5, 0x45, 0xe3, 0x8d, 0x31, 0x02, 0xbc,
	0x0c, 0xa3, 0xc7, 0x6b, 0xff, 0x39, 0xa4, 0xc5, 0xfe, 0x8e, 0xe9, 0xff, 0x83, 0xaf, 0x3b, 0xe8,
	0x9a, 0xfb, 0xff, 0x85, 0x8e, 0xb7, 0x96, 0x3e, 0x27, 0x2e, 0x63, 0xa2, 0x4e, 0x7b, 0x0e, 0x58,
	0x06, 0x67, 0xd2, 0x7d, 0x30, 0xd4, 0xfc, 0x81, 0x33, 0x26, 0xb2, 0xcb, 0x83, 0x3a, 0xda, 0x02,
	0x67, 0x2f, 0xcd, 0x34, 0x79, 0x8f, 0xf0, 0x4d, 0x18, 0x10, 0x1c, 0x41, 0xd3, 0x93, 0x14, 0xbf,
	0xa8, 0xbc, 0x8b, 0x5f, 0xc2, 0xea, 0x94, 0x4e, 0x94, 0xb6, 0xfd, 0x0e, 0x29, 0x98, 0xc7, 0xc2,
	0xc1, 0xa1, 0x22, 0xd4, 0xe4, 0x6e, 0xfd, 0xae, 0x83, 0xe5, 0x66, 0x53, 0x0d, 0x6f, 0xe1, 0xf3,
	0x6e, 0x3c, 0xf8, 0x53, 0xb1, 0x2b, 0x53, 0xb6, 0x7a, 0xd5, 0xc3, 0x5c, 0xd0, 0xd1, 0xf0, 0x0a,
	0x8c, 0x52, 0x1e, 0x68, 0x2a, 0xc2, 0x61, 0xa0, 0xd6, 0x8f, 0x8a, 0x49, 0xa1, 0x33, 0xd5, 0x70,
	0x01, 0x50, 0x6c, 0x1f, 0xbf, 0x67, 0xcb, 0xda, 0x0f, 0xcf, 0x32, 0x0f, 0x07, 0xb9, 0xcc, 0xd9,
	0x87, 0x7b, 0x5d, 0x3d, 0xe1, 0xa5, 0x9e, 0x3e, 0xdf, 0xff, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x59, 0x7d, 0xc8, 0x75, 0xde, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (SumService_PrimeNumberDecompositionClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (SumService_ComputeAverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (SumService_FindMaximumClient, error)
	// This rpc will throw exception if number is negative; error -- INVALID_ARGUMENT
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calc.SumService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (SumService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[0], "/calc.SumService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SumService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type sumServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *sumServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (SumService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[1], "/calc.SumService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceComputeAverageClient{stream}
	return x, nil
}

type SumService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type sumServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *sumServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (SumService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[2], "/calc.SumService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceFindMaximumClient{stream}
	return x, nil
}

type SumService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type sumServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *sumServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calc.SumService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, SumService_PrimeNumberDecompositionServer) error
	ComputeAverage(SumService_ComputeAverageServer) error
	FindMaximum(SumService_FindMaximumServer) error
	// This rpc will throw exception if number is negative; error -- INVALID_ARGUMENT
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedSumServiceServer) PrimeNumberDecomposition(req *PrimeNumberDecompositionRequest, srv SumService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}
func (*UnimplementedSumServiceServer) ComputeAverage(srv SumService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (*UnimplementedSumServiceServer) FindMaximum(srv SumService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}
func (*UnimplementedSumServiceServer) SquareRoot(ctx context.Context, req *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SumServiceServer).PrimeNumberDecomposition(m, &sumServicePrimeNumberDecompositionServer{stream})
}

type SumService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type sumServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *sumServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SumService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServiceServer).ComputeAverage(&sumServiceComputeAverageServer{stream})
}

type SumService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type sumServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *sumServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SumService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServiceServer).FindMaximum(&sumServiceFindMaximumServer{stream})
}

type SumService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type sumServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *sumServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SumService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.SumService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumService_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _SumService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _SumService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _SumService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _SumService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}