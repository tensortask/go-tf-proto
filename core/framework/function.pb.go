// Code generated by protoc-gen-go. DO NOT EDIT.
// source: function.proto

package tensorflow

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

// A library is a set of named functions.
type FunctionDefLibrary struct {
	Function             []*FunctionDef `protobuf:"bytes,1,rep,name=function,proto3" json:"function,omitempty"`
	Gradient             []*GradientDef `protobuf:"bytes,2,rep,name=gradient,proto3" json:"gradient,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FunctionDefLibrary) Reset()         { *m = FunctionDefLibrary{} }
func (m *FunctionDefLibrary) String() string { return proto.CompactTextString(m) }
func (*FunctionDefLibrary) ProtoMessage()    {}
func (*FunctionDefLibrary) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac74addf543d91a, []int{0}
}

func (m *FunctionDefLibrary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionDefLibrary.Unmarshal(m, b)
}
func (m *FunctionDefLibrary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionDefLibrary.Marshal(b, m, deterministic)
}
func (m *FunctionDefLibrary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionDefLibrary.Merge(m, src)
}
func (m *FunctionDefLibrary) XXX_Size() int {
	return xxx_messageInfo_FunctionDefLibrary.Size(m)
}
func (m *FunctionDefLibrary) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionDefLibrary.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionDefLibrary proto.InternalMessageInfo

func (m *FunctionDefLibrary) GetFunction() []*FunctionDef {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *FunctionDefLibrary) GetGradient() []*GradientDef {
	if m != nil {
		return m.Gradient
	}
	return nil
}

// A function can be instantiated when the runtime can bind every attr
// with a value. When a GraphDef has a call to a function, it must
// have binding for every attr defined in the signature.
//
// TODO(zhifengc):
//   * device spec, etc.
type FunctionDef struct {
	// The definition of the function's name, arguments, return values,
	// attrs etc.
	Signature *OpDef `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Attributes specific to this function definition.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// By convention, "op" in node_def is resolved by consulting with a
	// user-defined library first. If not resolved, "func" is assumed to
	// be a builtin op.
	NodeDef []*NodeDef `protobuf:"bytes,3,rep,name=node_def,json=nodeDef,proto3" json:"node_def,omitempty"`
	// A mapping from the output arg names from `signature` to the
	// outputs from `node_def` that should be returned by the function.
	Ret                  map[string]string `protobuf:"bytes,4,rep,name=ret,proto3" json:"ret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FunctionDef) Reset()         { *m = FunctionDef{} }
func (m *FunctionDef) String() string { return proto.CompactTextString(m) }
func (*FunctionDef) ProtoMessage()    {}
func (*FunctionDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac74addf543d91a, []int{1}
}

func (m *FunctionDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionDef.Unmarshal(m, b)
}
func (m *FunctionDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionDef.Marshal(b, m, deterministic)
}
func (m *FunctionDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionDef.Merge(m, src)
}
func (m *FunctionDef) XXX_Size() int {
	return xxx_messageInfo_FunctionDef.Size(m)
}
func (m *FunctionDef) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionDef.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionDef proto.InternalMessageInfo

func (m *FunctionDef) GetSignature() *OpDef {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FunctionDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *FunctionDef) GetNodeDef() []*NodeDef {
	if m != nil {
		return m.NodeDef
	}
	return nil
}

func (m *FunctionDef) GetRet() map[string]string {
	if m != nil {
		return m.Ret
	}
	return nil
}

// GradientDef defines the gradient function of a function defined in
// a function library.
//
// A gradient function g (specified by gradient_func) for a function f
// (specified by function_name) must follow the following:
//
// The function 'f' must be a numerical function which takes N inputs
// and produces M outputs. Its gradient function 'g', which is a
// function taking N + M inputs and produces N outputs.
//
// I.e. if we have
//    (y1, y2, ..., y_M) = f(x1, x2, ..., x_N),
// then, g is
//    (dL/dx1, dL/dx2, ..., dL/dx_N) = g(x1, x2, ..., x_N,
//                                      dL/dy1, dL/dy2, ..., dL/dy_M),
// where L is a scalar-value function of (x1, x2, ..., xN) (e.g., the
// loss function). dL/dx_i is the partial derivative of L with respect
// to x_i.
type GradientDef struct {
	FunctionName         string   `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	GradientFunc         string   `protobuf:"bytes,2,opt,name=gradient_func,json=gradientFunc,proto3" json:"gradient_func,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GradientDef) Reset()         { *m = GradientDef{} }
func (m *GradientDef) String() string { return proto.CompactTextString(m) }
func (*GradientDef) ProtoMessage()    {}
func (*GradientDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac74addf543d91a, []int{2}
}

func (m *GradientDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientDef.Unmarshal(m, b)
}
func (m *GradientDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientDef.Marshal(b, m, deterministic)
}
func (m *GradientDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientDef.Merge(m, src)
}
func (m *GradientDef) XXX_Size() int {
	return xxx_messageInfo_GradientDef.Size(m)
}
func (m *GradientDef) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientDef.DiscardUnknown(m)
}

var xxx_messageInfo_GradientDef proto.InternalMessageInfo

func (m *GradientDef) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *GradientDef) GetGradientFunc() string {
	if m != nil {
		return m.GradientFunc
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionDefLibrary)(nil), "tensorflow.FunctionDefLibrary")
	proto.RegisterType((*FunctionDef)(nil), "tensorflow.FunctionDef")
	proto.RegisterMapType((map[string]*AttrValue)(nil), "tensorflow.FunctionDef.AttrEntry")
	proto.RegisterMapType((map[string]string)(nil), "tensorflow.FunctionDef.RetEntry")
	proto.RegisterType((*GradientDef)(nil), "tensorflow.GradientDef")
}

func init() { proto.RegisterFile("function.proto", fileDescriptor_8ac74addf543d91a) }

var fileDescriptor_8ac74addf543d91a = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x6b, 0xea, 0x30,
	0x18, 0xc6, 0x69, 0xab, 0xe7, 0xb4, 0x6f, 0x3d, 0xe2, 0xc9, 0x39, 0x63, 0xa1, 0x57, 0xce, 0xdd,
	0x08, 0x83, 0x0a, 0xca, 0xc6, 0xd8, 0xdd, 0x64, 0x7f, 0x60, 0x0c, 0x27, 0xbd, 0xd8, 0x2e, 0x4b,
	0x5c, 0x53, 0x29, 0x6a, 0x22, 0x31, 0x4e, 0xbc, 0xd9, 0xd7, 0xdd, 0x57, 0xd8, 0xe5, 0x48, 0xda,
	0x68, 0x60, 0xf3, 0x2e, 0xc9, 0xfb, 0xfc, 0x9e, 0x27, 0x6f, 0xde, 0x40, 0x33, 0x5f, 0xb3, 0x57,
	0x59, 0x70, 0x16, 0x2f, 0x05, 0x97, 0x1c, 0x81, 0xa4, 0x6c, 0xc5, 0x45, 0x3e, 0xe7, 0x9b, 0xa8,
	0x45, 0xa4, 0x14, 0xe9, 0x1b, 0x99, 0xaf, 0x69, 0x59, 0x8d, 0x9a, 0x8c, 0x67, 0x34, 0xcd, 0x68,
	0x5e, 0xed, 0x1b, 0x7c, 0xb9, 0xdf, 0x75, 0xde, 0x01, 0xdd, 0x55, 0x6e, 0x37, 0x34, 0x7f, 0x2c,
	0x26, 0x82, 0x88, 0x2d, 0x1a, 0x80, 0x6f, 0x32, 0xb0, 0xd3, 0xf6, 0xba, 0x61, 0xff, 0x38, 0xde,
	0x87, 0xc4, 0x16, 0x91, 0xec, 0x84, 0x0a, 0x9a, 0x0a, 0x92, 0x15, 0x94, 0x49, 0xec, 0x7e, 0x87,
	0xee, 0xab, 0x9a, 0x86, 0x8c, 0xb0, 0xf3, 0xe1, 0x42, 0x68, 0xd9, 0xa1, 0x1e, 0x04, 0xab, 0x62,
	0xca, 0x88, 0x5c, 0x0b, 0x8a, 0x9d, 0xb6, 0xd3, 0x0d, 0xfb, 0x7f, 0x6d, 0x97, 0xa7, 0xa5, 0xe2,
	0xf7, 0x1a, 0x74, 0x0e, 0x35, 0xd5, 0x32, 0xae, 0xeb, 0xc4, 0x93, 0x03, 0xd7, 0x8c, 0xaf, 0xa5,
	0x14, 0xb7, 0x4c, 0x8a, 0x6d, 0xa2, 0xe5, 0x28, 0x06, 0xdf, 0xbc, 0x0b, 0xf6, 0x34, 0xfa, 0xcf,
	0x46, 0x47, 0x3c, 0xa3, 0x2a, 0xe8, 0x37, 0x2b, 0x17, 0xa8, 0x0f, 0x9e, 0xa0, 0x12, 0xd7, 0xb4,
	0xb4, 0x7d, 0x28, 0x25, 0xa1, 0xb2, 0x0c, 0x51, 0xe2, 0x68, 0x04, 0xc1, 0x2e, 0x16, 0xb5, 0xc0,
	0x9b, 0xd1, 0xad, 0x6e, 0x29, 0x48, 0xd4, 0x12, 0x9d, 0x41, 0x5d, 0xcf, 0x09, 0xbb, 0xba, 0xcd,
	0x23, 0xdb, 0x54, 0x71, 0xcf, 0xaa, 0x98, 0x94, 0x9a, 0x2b, 0xf7, 0xd2, 0x89, 0x2e, 0xc0, 0x37,
	0x01, 0x3f, 0xd8, 0xfd, 0xb7, 0xed, 0x02, 0x8b, 0x7b, 0xa8, 0xf9, 0x6e, 0xcb, 0xeb, 0xbc, 0x40,
	0x68, 0x8d, 0x00, 0x9d, 0xc2, 0x1f, 0x33, 0xb9, 0x94, 0x91, 0x05, 0xad, 0xac, 0x1a, 0xe6, 0x70,
	0x44, 0x16, 0x54, 0x89, 0xcc, 0xa4, 0x52, 0x55, 0xa8, 0xbc, 0x1b, 0xe6, 0x50, 0xf5, 0x3e, 0xec,
	0x01, 0xe6, 0x62, 0x6a, 0xdf, 0x3e, 0x17, 0x64, 0x41, 0x37, 0x5c, 0xcc, 0x86, 0x4d, 0xf3, 0x3a,
	0x63, 0xf5, 0xdb, 0x56, 0x63, 0xe7, 0xd3, 0x71, 0x26, 0xbf, 0xf4, 0xd7, 0x1b, 0x7c, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0xb3, 0x08, 0x28, 0xc8, 0x02, 0x00, 0x00,
}
