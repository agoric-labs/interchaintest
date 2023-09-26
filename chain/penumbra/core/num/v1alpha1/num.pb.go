// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: penumbra/core/num/v1alpha1/num.proto

package numv1alpha1

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// The quantity of a particular Asset. Represented as a 128-bit unsigned integer,
// split over two fields, `lo` and `hi`, representing the low- and high-order bytes
// of the 128-bit value, respectively. Clients must assemble these bits in their
// implementation into a `uint128` or comparable data structure, in order to model
// the Amount accurately.
type Amount struct {
	Lo uint64 `protobuf:"varint,1,opt,name=lo,proto3" json:"lo,omitempty"`
	Hi uint64 `protobuf:"varint,2,opt,name=hi,proto3" json:"hi,omitempty"`
}

func (m *Amount) Reset()         { *m = Amount{} }
func (m *Amount) String() string { return proto.CompactTextString(m) }
func (*Amount) ProtoMessage()    {}
func (*Amount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a515e4f6cefdd105, []int{0}
}
func (m *Amount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Amount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Amount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Amount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Amount.Merge(m, src)
}
func (m *Amount) XXX_Size() int {
	return m.Size()
}
func (m *Amount) XXX_DiscardUnknown() {
	xxx_messageInfo_Amount.DiscardUnknown(m)
}

var xxx_messageInfo_Amount proto.InternalMessageInfo

func (m *Amount) GetLo() uint64 {
	if m != nil {
		return m.Lo
	}
	return 0
}

func (m *Amount) GetHi() uint64 {
	if m != nil {
		return m.Hi
	}
	return 0
}

func init() {
	proto.RegisterType((*Amount)(nil), "penumbra.core.num.v1alpha1.Amount")
}

func init() {
	proto.RegisterFile("penumbra/core/num/v1alpha1/num.proto", fileDescriptor_a515e4f6cefdd105)
}

var fileDescriptor_a515e4f6cefdd105 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4a, 0xc4, 0x30,
	0x18, 0xc7, 0x2f, 0x51, 0x0e, 0xe9, 0xe0, 0x70, 0xd3, 0x71, 0x60, 0x10, 0x11, 0xb9, 0xc5, 0x84,
	0xc3, 0x45, 0xe2, 0x64, 0x6f, 0x70, 0xb2, 0x14, 0x07, 0x07, 0x29, 0x42, 0x5a, 0xc3, 0xb5, 0xd0,
	0x24, 0x25, 0xcd, 0xd7, 0xe7, 0xf0, 0x19, 0x1c, 0x05, 0xdf, 0x43, 0x9c, 0x6e, 0x74, 0x94, 0x76,
	0xf3, 0x29, 0x24, 0xa7, 0xc1, 0xa9, 0x53, 0xf2, 0xe7, 0xf7, 0x83, 0xef, 0xff, 0x7d, 0xd1, 0x69,
	0x23, 0x35, 0xa8, 0xdc, 0x0a, 0x56, 0x18, 0x2b, 0x99, 0x06, 0xc5, 0xba, 0x95, 0xa8, 0x9b, 0x52,
	0xac, 0x7c, 0xa0, 0x8d, 0x35, 0xce, 0xcc, 0x16, 0xc1, 0xa2, 0xde, 0xa2, 0x1e, 0x04, 0xeb, 0x64,
	0x19, 0x4d, 0xaf, 0x95, 0x01, 0xed, 0x66, 0x87, 0x11, 0xae, 0xcd, 0x1c, 0x1d, 0xa3, 0xe5, 0xfe,
	0x1d, 0xae, 0x8d, 0xcf, 0x65, 0x35, 0xc7, 0xbf, 0xb9, 0xac, 0xe2, 0x37, 0xfc, 0xde, 0x13, 0xb4,
	0xed, 0x09, 0xfa, 0xea, 0x09, 0x7a, 0x1e, 0xc8, 0x64, 0x3b, 0x90, 0xc9, 0xe7, 0x40, 0x26, 0x11,
	0x29, 0x8c, 0xa2, 0xe3, 0x43, 0xe2, 0x83, 0x04, 0x54, 0xea, 0xab, 0xa4, 0xe8, 0xe1, 0x71, 0x53,
	0xb9, 0x12, 0x72, 0x5a, 0x18, 0xc5, 0x5a, 0x67, 0x85, 0xde, 0xc8, 0xda, 0x74, 0xf2, 0xbc, 0x93,
	0xda, 0x81, 0x95, 0x2d, 0xab, 0xb4, 0x93, 0xb6, 0x28, 0x85, 0x7f, 0x5b, 0xc7, 0xba, 0x4b, 0xb6,
	0x0b, 0x6c, 0x7c, 0xd5, 0x2b, 0x0d, 0x2a, 0xfc, 0x5f, 0xf0, 0x5e, 0xba, 0x4e, 0x5e, 0xf1, 0x22,
	0x0d, 0x75, 0xd6, 0xbe, 0x4e, 0x02, 0x8a, 0xde, 0xff, 0x29, 0x1f, 0xff, 0x30, 0xf3, 0x30, 0x4b,
	0x40, 0x65, 0x01, 0xf6, 0xf8, 0x6c, 0x1c, 0x66, 0x37, 0x69, 0x7c, 0x2b, 0x9d, 0x78, 0x12, 0x4e,
	0x7c, 0xe3, 0xa3, 0x20, 0x72, 0xee, 0x4d, 0xce, 0x13, 0x50, 0x9c, 0x07, 0x37, 0x9f, 0xee, 0x8e,
	0x7f, 0xf1, 0x13, 0x00, 0x00, 0xff, 0xff, 0x67, 0xb5, 0x67, 0x57, 0xa4, 0x01, 0x00, 0x00,
}

func (m *Amount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Amount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Amount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Hi != 0 {
		i = encodeVarintNum(dAtA, i, uint64(m.Hi))
		i--
		dAtA[i] = 0x10
	}
	if m.Lo != 0 {
		i = encodeVarintNum(dAtA, i, uint64(m.Lo))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNum(dAtA []byte, offset int, v uint64) int {
	offset -= sovNum(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Amount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Lo != 0 {
		n += 1 + sovNum(uint64(m.Lo))
	}
	if m.Hi != 0 {
		n += 1 + sovNum(uint64(m.Hi))
	}
	return n
}

func sovNum(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNum(x uint64) (n int) {
	return sovNum(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Amount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNum
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
			return fmt.Errorf("proto: Amount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Amount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lo", wireType)
			}
			m.Lo = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lo |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hi", wireType)
			}
			m.Hi = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hi |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNum
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
func skipNum(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNum
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
					return 0, ErrIntOverflowNum
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
					return 0, ErrIntOverflowNum
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
				return 0, ErrInvalidLengthNum
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNum
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNum
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNum        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNum          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNum = fmt.Errorf("proto: unexpected end of group")
)