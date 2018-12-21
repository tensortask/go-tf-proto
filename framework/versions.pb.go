// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: versions.proto

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

// Version information for a piece of serialized data
//
// There are different types of versions for each type of data
// (GraphDef, etc.), but they all have the same common shape
// described here.
//
// Each consumer has "consumer" and "min_producer" versions (specified
// elsewhere).  A consumer is allowed to consume this data if
//
//   producer >= min_producer
//   consumer >= min_consumer
//   consumer not in bad_consumers
//
type VersionDef struct {
	// The version of the code that produced this data.
	Producer int32 `protobuf:"varint,1,opt,name=producer,proto3" json:"producer,omitempty"`
	// Any consumer below this version is not allowed to consume this data.
	MinConsumer int32 `protobuf:"varint,2,opt,name=min_consumer,json=minConsumer,proto3" json:"min_consumer,omitempty"`
	// Specific consumer versions which are disallowed (e.g. due to bugs).
	BadConsumers         []int32  `protobuf:"varint,3,rep,packed,name=bad_consumers,json=badConsumers" json:"bad_consumers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionDef) Reset()      { *m = VersionDef{} }
func (*VersionDef) ProtoMessage() {}
func (*VersionDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_versions_a97c0ac4de396b7b, []int{0}
}
func (m *VersionDef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersionDef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *VersionDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionDef.Merge(dst, src)
}
func (m *VersionDef) XXX_Size() int {
	return m.Size()
}
func (m *VersionDef) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionDef.DiscardUnknown(m)
}

var xxx_messageInfo_VersionDef proto.InternalMessageInfo

func (m *VersionDef) GetProducer() int32 {
	if m != nil {
		return m.Producer
	}
	return 0
}

func (m *VersionDef) GetMinConsumer() int32 {
	if m != nil {
		return m.MinConsumer
	}
	return 0
}

func (m *VersionDef) GetBadConsumers() []int32 {
	if m != nil {
		return m.BadConsumers
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionDef)(nil), "framework.VersionDef")
}
func (this *VersionDef) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*VersionDef)
	if !ok {
		that2, ok := that.(VersionDef)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *VersionDef")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *VersionDef but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *VersionDef but is not nil && this == nil")
	}
	if this.Producer != that1.Producer {
		return fmt.Errorf("Producer this(%v) Not Equal that(%v)", this.Producer, that1.Producer)
	}
	if this.MinConsumer != that1.MinConsumer {
		return fmt.Errorf("MinConsumer this(%v) Not Equal that(%v)", this.MinConsumer, that1.MinConsumer)
	}
	if len(this.BadConsumers) != len(that1.BadConsumers) {
		return fmt.Errorf("BadConsumers this(%v) Not Equal that(%v)", len(this.BadConsumers), len(that1.BadConsumers))
	}
	for i := range this.BadConsumers {
		if this.BadConsumers[i] != that1.BadConsumers[i] {
			return fmt.Errorf("BadConsumers this[%v](%v) Not Equal that[%v](%v)", i, this.BadConsumers[i], i, that1.BadConsumers[i])
		}
	}
	return nil
}
func (this *VersionDef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VersionDef)
	if !ok {
		that2, ok := that.(VersionDef)
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
	if this.Producer != that1.Producer {
		return false
	}
	if this.MinConsumer != that1.MinConsumer {
		return false
	}
	if len(this.BadConsumers) != len(that1.BadConsumers) {
		return false
	}
	for i := range this.BadConsumers {
		if this.BadConsumers[i] != that1.BadConsumers[i] {
			return false
		}
	}
	return true
}
func (this *VersionDef) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&framework.VersionDef{")
	s = append(s, "Producer: "+fmt.Sprintf("%#v", this.Producer)+",\n")
	s = append(s, "MinConsumer: "+fmt.Sprintf("%#v", this.MinConsumer)+",\n")
	s = append(s, "BadConsumers: "+fmt.Sprintf("%#v", this.BadConsumers)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringVersions(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *VersionDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Producer != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintVersions(dAtA, i, uint64(m.Producer))
	}
	if m.MinConsumer != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintVersions(dAtA, i, uint64(m.MinConsumer))
	}
	if len(m.BadConsumers) > 0 {
		dAtA2 := make([]byte, len(m.BadConsumers)*10)
		var j1 int
		for _, num1 := range m.BadConsumers {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintVersions(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func encodeVarintVersions(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedVersionDef(r randyVersions, easy bool) *VersionDef {
	this := &VersionDef{}
	this.Producer = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Producer *= -1
	}
	this.MinConsumer = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.MinConsumer *= -1
	}
	v1 := r.Intn(10)
	this.BadConsumers = make([]int32, v1)
	for i := 0; i < v1; i++ {
		this.BadConsumers[i] = int32(r.Int31())
		if r.Intn(2) == 0 {
			this.BadConsumers[i] *= -1
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyVersions interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneVersions(r randyVersions) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringVersions(r randyVersions) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneVersions(r)
	}
	return string(tmps)
}
func randUnrecognizedVersions(r randyVersions, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldVersions(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldVersions(dAtA []byte, r randyVersions, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateVersions(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateVersions(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateVersions(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateVersions(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateVersions(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateVersions(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateVersions(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *VersionDef) Size() (n int) {
	var l int
	_ = l
	if m.Producer != 0 {
		n += 1 + sovVersions(uint64(m.Producer))
	}
	if m.MinConsumer != 0 {
		n += 1 + sovVersions(uint64(m.MinConsumer))
	}
	if len(m.BadConsumers) > 0 {
		l = 0
		for _, e := range m.BadConsumers {
			l += sovVersions(uint64(e))
		}
		n += 1 + sovVersions(uint64(l)) + l
	}
	return n
}

func sovVersions(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozVersions(x uint64) (n int) {
	return sovVersions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *VersionDef) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VersionDef{`,
		`Producer:` + fmt.Sprintf("%v", this.Producer) + `,`,
		`MinConsumer:` + fmt.Sprintf("%v", this.MinConsumer) + `,`,
		`BadConsumers:` + fmt.Sprintf("%v", this.BadConsumers) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringVersions(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *VersionDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVersions
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
			return fmt.Errorf("proto: VersionDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Producer", wireType)
			}
			m.Producer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Producer |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinConsumer", wireType)
			}
			m.MinConsumer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinConsumer |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVersions
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BadConsumers = append(m.BadConsumers, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVersions
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthVersions
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVersions
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BadConsumers = append(m.BadConsumers, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BadConsumers", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVersions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVersions
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
func skipVersions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVersions
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
					return 0, ErrIntOverflowVersions
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
					return 0, ErrIntOverflowVersions
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
				return 0, ErrInvalidLengthVersions
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVersions
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
				next, err := skipVersions(dAtA[start:])
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
	ErrInvalidLengthVersions = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVersions   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("versions.proto", fileDescriptor_versions_a97c0ac4de396b7b) }

var fileDescriptor_versions_a97c0ac4de396b7b = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4b, 0x2d, 0x2a,
	0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x2b, 0x4a, 0xcc,
	0x4d, 0x2d, 0xcf, 0x2f, 0xca, 0x96, 0x52, 0x4a, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x0b, 0x27, 0x95,
	0xa6, 0xe9, 0x83, 0x78, 0x60, 0x0e, 0x98, 0x05, 0x51, 0xae, 0x54, 0xc0, 0xc5, 0x15, 0x06, 0x31,
	0xc0, 0x25, 0x35, 0x4d, 0x48, 0x8a, 0x8b, 0xa3, 0xa0, 0x28, 0x3f, 0xa5, 0x34, 0x39, 0xb5, 0x48,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x08, 0xce, 0x17, 0x52, 0xe4, 0xe2, 0xc9, 0xcd, 0xcc, 0x8b,
	0x4f, 0xce, 0xcf, 0x2b, 0x2e, 0xcd, 0x4d, 0x2d, 0x92, 0x60, 0x02, 0xcb, 0x73, 0xe7, 0x66, 0xe6,
	0x39, 0x43, 0x85, 0x84, 0x94, 0xb9, 0x78, 0x93, 0x12, 0x53, 0xe0, 0x4a, 0x8a, 0x25, 0x98, 0x15,
	0x98, 0x35, 0x58, 0x83, 0x78, 0x92, 0x12, 0x53, 0x60, 0x6a, 0x8a, 0x9d, 0xa2, 0x6f, 0x3c, 0x94,
	0x63, 0x78, 0xf0, 0x50, 0x8e, 0xf1, 0xc3, 0x43, 0x39, 0xc6, 0x1f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e,
	0xc9, 0x31, 0xae, 0x78, 0x24, 0xc7, 0xb8, 0xe3, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x03, 0x97, 0x44, 0x7e, 0x51,
	0xba, 0x5e, 0x49, 0x6a, 0x5e, 0x71, 0x7e, 0x51, 0x5a, 0x4e, 0x7e, 0xb9, 0x1e, 0xdc, 0x77, 0x4e,
	0x7c, 0x50, 0x77, 0x17, 0x07, 0x80, 0x3c, 0x52, 0x1c, 0xc0, 0xf8, 0x83, 0x91, 0x31, 0x89, 0x0d,
	0xec, 0x2b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0xa2, 0x50, 0xab, 0x16, 0x01, 0x00,
	0x00,
}
