// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: redstoneconsolidator/price_feed.proto

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

type PriceFeed struct {
	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BirthTime string   `protobuf:"bytes,2,opt,name=birthTime,proto3" json:"birthTime,omitempty"`
	Interval  string   `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
	Providers []string `protobuf:"bytes,4,rep,name=providers,proto3" json:"providers,omitempty"`
	Assets    []string `protobuf:"bytes,5,rep,name=assets,proto3" json:"assets,omitempty"`
}

func (m *PriceFeed) Reset()         { *m = PriceFeed{} }
func (m *PriceFeed) String() string { return proto.CompactTextString(m) }
func (*PriceFeed) ProtoMessage()    {}
func (*PriceFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa922ab09faa1645, []int{0}
}
func (m *PriceFeed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceFeed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceFeed.Merge(m, src)
}
func (m *PriceFeed) XXX_Size() int {
	return m.Size()
}
func (m *PriceFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceFeed.DiscardUnknown(m)
}

var xxx_messageInfo_PriceFeed proto.InternalMessageInfo

func (m *PriceFeed) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PriceFeed) GetBirthTime() string {
	if m != nil {
		return m.BirthTime
	}
	return ""
}

func (m *PriceFeed) GetInterval() string {
	if m != nil {
		return m.Interval
	}
	return ""
}

func (m *PriceFeed) GetProviders() []string {
	if m != nil {
		return m.Providers
	}
	return nil
}

func (m *PriceFeed) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func init() {
	proto.RegisterType((*PriceFeed)(nil), "kodek16.redstoneconsolidator.redstoneconsolidator.PriceFeed")
}

func init() {
	proto.RegisterFile("redstoneconsolidator/price_feed.proto", fileDescriptor_fa922ab09faa1645)
}

var fileDescriptor_fa922ab09faa1645 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x4a, 0x4d, 0x29,
	0x2e, 0xc9, 0xcf, 0x4b, 0x4d, 0xce, 0xcf, 0x2b, 0xce, 0xcf, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f,
	0xd2, 0x2f, 0x28, 0xca, 0x4c, 0x4e, 0x8d, 0x4f, 0x4b, 0x4d, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x32, 0xcc, 0xce, 0x4f, 0x49, 0xcd, 0x36, 0x34, 0xd3, 0xc3, 0xa6, 0x1c, 0xab, 0xa0,
	0x52, 0x3f, 0x23, 0x17, 0x67, 0x00, 0xc8, 0x1c, 0xb7, 0xd4, 0xd4, 0x14, 0x21, 0x21, 0x2e, 0x96,
	0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x86, 0x8b,
	0x33, 0x29, 0xb3, 0xa8, 0x24, 0x23, 0x24, 0x33, 0x37, 0x55, 0x82, 0x09, 0x2c, 0x81, 0x10, 0x10,
	0x92, 0xe2, 0xe2, 0xc8, 0xcc, 0x2b, 0x49, 0x2d, 0x2a, 0x4b, 0xcc, 0x91, 0x60, 0x06, 0x4b, 0xc2,
	0xf9, 0x20, 0x9d, 0x05, 0x45, 0xf9, 0x65, 0x99, 0x29, 0xa9, 0x45, 0xc5, 0x12, 0x2c, 0x0a, 0xcc,
	0x20, 0x9d, 0x70, 0x01, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xe2, 0xe2, 0xd4, 0x92, 0x62, 0x09, 0x56,
	0xb0, 0x14, 0x94, 0xe7, 0x14, 0x7f, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51,
	0xae, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x50, 0x9f, 0xea, 0xc3,
	0x3c, 0xa5, 0x8b, 0x12, 0x32, 0x15, 0xfa, 0x58, 0x03, 0xac, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89,
	0x0d, 0x1c, 0x58, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xb1, 0xe2, 0x93, 0x55, 0x01,
	0x00, 0x00,
}

func (m *PriceFeed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceFeed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceFeed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Assets[iNdEx])
			copy(dAtA[i:], m.Assets[iNdEx])
			i = encodeVarintPriceFeed(dAtA, i, uint64(len(m.Assets[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Providers) > 0 {
		for iNdEx := len(m.Providers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Providers[iNdEx])
			copy(dAtA[i:], m.Providers[iNdEx])
			i = encodeVarintPriceFeed(dAtA, i, uint64(len(m.Providers[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Interval) > 0 {
		i -= len(m.Interval)
		copy(dAtA[i:], m.Interval)
		i = encodeVarintPriceFeed(dAtA, i, uint64(len(m.Interval)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BirthTime) > 0 {
		i -= len(m.BirthTime)
		copy(dAtA[i:], m.BirthTime)
		i = encodeVarintPriceFeed(dAtA, i, uint64(len(m.BirthTime)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPriceFeed(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPriceFeed(dAtA []byte, offset int, v uint64) int {
	offset -= sovPriceFeed(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PriceFeed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPriceFeed(uint64(l))
	}
	l = len(m.BirthTime)
	if l > 0 {
		n += 1 + l + sovPriceFeed(uint64(l))
	}
	l = len(m.Interval)
	if l > 0 {
		n += 1 + l + sovPriceFeed(uint64(l))
	}
	if len(m.Providers) > 0 {
		for _, s := range m.Providers {
			l = len(s)
			n += 1 + l + sovPriceFeed(uint64(l))
		}
	}
	if len(m.Assets) > 0 {
		for _, s := range m.Assets {
			l = len(s)
			n += 1 + l + sovPriceFeed(uint64(l))
		}
	}
	return n
}

func sovPriceFeed(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPriceFeed(x uint64) (n int) {
	return sovPriceFeed(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PriceFeed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPriceFeed
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
			return fmt.Errorf("proto: PriceFeed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceFeed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriceFeed
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
				return ErrInvalidLengthPriceFeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPriceFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BirthTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriceFeed
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
				return ErrInvalidLengthPriceFeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPriceFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BirthTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriceFeed
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
				return ErrInvalidLengthPriceFeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPriceFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Interval = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Providers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriceFeed
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
				return ErrInvalidLengthPriceFeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPriceFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Providers = append(m.Providers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPriceFeed
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
				return ErrInvalidLengthPriceFeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPriceFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assets = append(m.Assets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPriceFeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPriceFeed
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
func skipPriceFeed(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPriceFeed
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
					return 0, ErrIntOverflowPriceFeed
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
					return 0, ErrIntOverflowPriceFeed
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
				return 0, ErrInvalidLengthPriceFeed
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPriceFeed
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPriceFeed
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPriceFeed        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPriceFeed          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPriceFeed = fmt.Errorf("proto: unexpected end of group")
)