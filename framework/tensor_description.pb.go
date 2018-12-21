// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensor_description.proto

package framework

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TensorDescription struct {
	// Data type of tensor elements
	Dtype DataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=framework.DataType" json:"dtype,omitempty"`
	// Shape of the tensor.
	Shape *TensorShapeProto `protobuf:"bytes,2,opt,name=shape" json:"shape,omitempty"`
	// Information about the size and allocator used for the data
	AllocationDescription *AllocationDescription `protobuf:"bytes,4,opt,name=allocation_description,json=allocationDescription" json:"allocation_description,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *TensorDescription) Reset()      { *m = TensorDescription{} }
func (*TensorDescription) ProtoMessage() {}
func (*TensorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_tensor_description_01eb5e1625769f57, []int{0}
}
func (m *TensorDescription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TensorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TensorDescription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TensorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorDescription.Merge(dst, src)
}
func (m *TensorDescription) XXX_Size() int {
	return m.Size()
}
func (m *TensorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_TensorDescription proto.InternalMessageInfo

func (m *TensorDescription) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *TensorDescription) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *TensorDescription) GetAllocationDescription() *AllocationDescription {
	if m != nil {
		return m.AllocationDescription
	}
	return nil
}

func init() {
	proto.RegisterType((*TensorDescription)(nil), "framework.TensorDescription")
}
func (this *TensorDescription) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*TensorDescription)
	if !ok {
		that2, ok := that.(TensorDescription)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *TensorDescription")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *TensorDescription but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *TensorDescription but is not nil && this == nil")
	}
	if this.Dtype != that1.Dtype {
		return fmt.Errorf("Dtype this(%v) Not Equal that(%v)", this.Dtype, that1.Dtype)
	}
	if !this.Shape.Equal(that1.Shape) {
		return fmt.Errorf("Shape this(%v) Not Equal that(%v)", this.Shape, that1.Shape)
	}
	if !this.AllocationDescription.Equal(that1.AllocationDescription) {
		return fmt.Errorf("AllocationDescription this(%v) Not Equal that(%v)", this.AllocationDescription, that1.AllocationDescription)
	}
	return nil
}
func (this *TensorDescription) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TensorDescription)
	if !ok {
		that2, ok := that.(TensorDescription)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Dtype != that1.Dtype {
		return false
	}
	if !this.Shape.Equal(that1.Shape) {
		return false
	}
	if !this.AllocationDescription.Equal(that1.AllocationDescription) {
		return false
	}
	return true
}
func (this *TensorDescription) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&framework.TensorDescription{")
	s = append(s, "Dtype: "+fmt.Sprintf("%#v", this.Dtype)+",\n")
	if this.Shape != nil {
		s = append(s, "Shape: "+fmt.Sprintf("%#v", this.Shape)+",\n")
	}
	if this.AllocationDescription != nil {
		s = append(s, "AllocationDescription: "+fmt.Sprintf("%#v", this.AllocationDescription)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTensorDescription(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TensorDescription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TensorDescription) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Dtype != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTensorDescription(dAtA, i, uint64(m.Dtype))
	}
	if m.Shape != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTensorDescription(dAtA, i, uint64(m.Shape.Size()))
		n1, err := m.Shape.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.AllocationDescription != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTensorDescription(dAtA, i, uint64(m.AllocationDescription.Size()))
		n2, err := m.AllocationDescription.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintTensorDescription(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedTensorDescription(r randyTensorDescription, easy bool) *TensorDescription {
	this := &TensorDescription{}
	this.Dtype = DataType([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123}[r.Intn(47)])
	if r.Intn(10) != 0 {
		this.Shape = NewPopulatedTensorShapeProto(r, easy)
	}
	if r.Intn(10) != 0 {
		this.AllocationDescription = NewPopulatedAllocationDescription(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTensorDescription interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTensorDescription(r randyTensorDescription) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTensorDescription(r randyTensorDescription) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTensorDescription(r)
	}
	return string(tmps)
}
func randUnrecognizedTensorDescription(r randyTensorDescription, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTensorDescription(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTensorDescription(dAtA []byte, r randyTensorDescription, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTensorDescription(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTensorDescription(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTensorDescription(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTensorDescription(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTensorDescription(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTensorDescription(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTensorDescription(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *TensorDescription) Size() (n int) {
	var l int
	_ = l
	if m.Dtype != 0 {
		n += 1 + sovTensorDescription(uint64(m.Dtype))
	}
	if m.Shape != nil {
		l = m.Shape.Size()
		n += 1 + l + sovTensorDescription(uint64(l))
	}
	if m.AllocationDescription != nil {
		l = m.AllocationDescription.Size()
		n += 1 + l + sovTensorDescription(uint64(l))
	}
	return n
}

func sovTensorDescription(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTensorDescription(x uint64) (n int) {
	return sovTensorDescription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TensorDescription) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TensorDescription{`,
		`Dtype:` + fmt.Sprintf("%v", this.Dtype) + `,`,
		`Shape:` + strings.Replace(fmt.Sprintf("%v", this.Shape), "TensorShapeProto", "TensorShapeProto", 1) + `,`,
		`AllocationDescription:` + strings.Replace(fmt.Sprintf("%v", this.AllocationDescription), "AllocationDescription", "AllocationDescription", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTensorDescription(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TensorDescription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensorDescription
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TensorDescription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TensorDescription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtype", wireType)
			}
			m.Dtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dtype |= (DataType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shape", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTensorDescription
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Shape == nil {
				m.Shape = &TensorShapeProto{}
			}
			if err := m.Shape.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationDescription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensorDescription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTensorDescription
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AllocationDescription == nil {
				m.AllocationDescription = &AllocationDescription{}
			}
			if err := m.AllocationDescription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTensorDescription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTensorDescription
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
func skipTensorDescription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTensorDescription
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
					return 0, ErrIntOverflowTensorDescription
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTensorDescription
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTensorDescription
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTensorDescription
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTensorDescription(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTensorDescription = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTensorDescription   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("tensor_description.proto", fileDescriptor_tensor_description_01eb5e1625769f57)
}

var fileDescriptor_tensor_description_01eb5e1625769f57 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x8a, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x2b, 0x4a, 0xcc, 0x4d, 0x2d, 0xcf, 0x2f, 0xca, 0x96,
	0x52, 0x4a, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x0b, 0x27, 0x95, 0xa6, 0xe9, 0x83, 0x78, 0x60, 0x0e,
	0x98, 0x05, 0x51, 0x2e, 0xc5, 0x5d, 0x52, 0x59, 0x90, 0x5a, 0x0c, 0xe5, 0x08, 0x41, 0x4d, 0x2d,
	0xce, 0x48, 0x2c, 0x48, 0x85, 0x8a, 0xc9, 0x24, 0xe6, 0xe4, 0xe4, 0x27, 0x27, 0x82, 0x6c, 0xc0,
	0xb4, 0x4d, 0xe9, 0x14, 0x23, 0x97, 0x60, 0x08, 0x58, 0x93, 0x0b, 0x42, 0x4e, 0x48, 0x93, 0x8b,
	0x35, 0x05, 0x64, 0xae, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0xb0, 0x1e, 0xdc, 0x4d, 0x7a,
	0x2e, 0x89, 0x25, 0x89, 0x21, 0x95, 0x05, 0xa9, 0x41, 0x10, 0x15, 0x42, 0x86, 0x5c, 0xac, 0x60,
	0xdb, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xa4, 0x91, 0x94, 0x42, 0xcc, 0x0d, 0x06, 0xc9,
	0x06, 0x80, 0x2c, 0x0b, 0x82, 0xa8, 0x14, 0x0a, 0xe7, 0x12, 0xc3, 0xee, 0x26, 0x09, 0x16, 0xb0,
	0x19, 0x0a, 0x48, 0x66, 0x38, 0xc2, 0x15, 0x22, 0xb9, 0x2f, 0x48, 0x34, 0x11, 0x9b, 0xb0, 0x53,
	0xca, 0x8d, 0x87, 0x72, 0x0c, 0x0f, 0x1e, 0xca, 0x31, 0x7e, 0x78, 0x28, 0xc7, 0xf8, 0xe3, 0xa1,
	0x1c, 0x63, 0xc3, 0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0x77, 0x3c, 0x92, 0x63, 0x3c, 0xf1,
	0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xe0,
	0x92, 0xc8, 0x2f, 0x4a, 0xd7, 0x83, 0x04, 0x58, 0x5a, 0x4e, 0x7e, 0x39, 0xc2, 0x52, 0x27, 0x71,
	0x8c, 0x10, 0x01, 0xbb, 0xbf, 0x38, 0x80, 0xf1, 0x07, 0x23, 0x63, 0x12, 0x1b, 0x38, 0xe4, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x47, 0xa3, 0xc6, 0xc3, 0x01, 0x00, 0x00,
}
