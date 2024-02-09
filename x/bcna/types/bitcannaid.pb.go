// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bcna/bcna/bitcannaid.proto

package types

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

type Bitcannaid struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bcnaid  string `protobuf:"bytes,2,opt,name=bcnaid,proto3" json:"bcnaid,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Bitcannaid) Reset()         { *m = Bitcannaid{} }
func (m *Bitcannaid) String() string { return proto.CompactTextString(m) }
func (*Bitcannaid) ProtoMessage()    {}
func (*Bitcannaid) Descriptor() ([]byte, []int) {
	return fileDescriptor_230c1d9be0912c82, []int{0}
}
func (m *Bitcannaid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bitcannaid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bitcannaid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bitcannaid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bitcannaid.Merge(m, src)
}
func (m *Bitcannaid) XXX_Size() int {
	return m.Size()
}
func (m *Bitcannaid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bitcannaid.DiscardUnknown(m)
}

var xxx_messageInfo_Bitcannaid proto.InternalMessageInfo

func (m *Bitcannaid) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Bitcannaid) GetBcnaid() string {
	if m != nil {
		return m.Bcnaid
	}
	return ""
}

func (m *Bitcannaid) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Bitcannaid) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Bitcannaid)(nil), "bcna.bcna.Bitcannaid")
}

func init() { proto.RegisterFile("bcna/bcna/bitcannaid.proto", fileDescriptor_230c1d9be0912c82) }

var fileDescriptor_230c1d9be0912c82 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x4a, 0xce, 0x4b,
	0xd4, 0x87, 0x10, 0x99, 0x25, 0xc9, 0x89, 0x79, 0x79, 0x89, 0x99, 0x29, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x9c, 0x20, 0x61, 0x3d, 0x10, 0xa1, 0x94, 0xc1, 0xc5, 0xe5, 0x04, 0x97, 0x16,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xca, 0x4c, 0x11,
	0x12, 0xe3, 0x62, 0x03, 0xa9, 0xca, 0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2,
	0x84, 0x24, 0xb8, 0xd8, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x98, 0xc1, 0x12, 0x30,
	0x2e, 0x48, 0x26, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0xbf, 0x48, 0x82, 0x05, 0x22, 0x03, 0xe5, 0x3a,
	0xb9, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x76, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x53, 0x66, 0x89, 0x33, 0xc8, 0x31, 0xee, 0x39,
	0xf9, 0x49, 0x89, 0x39, 0x10, 0xf7, 0x57, 0x40, 0xa8, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36,
	0xb0, 0x17, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39, 0xeb, 0xa5, 0xd0, 0xe0, 0x00, 0x00,
	0x00,
}

func (m *Bitcannaid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bitcannaid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bitcannaid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBitcannaid(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintBitcannaid(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Bcnaid) > 0 {
		i -= len(m.Bcnaid)
		copy(dAtA[i:], m.Bcnaid)
		i = encodeVarintBitcannaid(dAtA, i, uint64(len(m.Bcnaid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintBitcannaid(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBitcannaid(dAtA []byte, offset int, v uint64) int {
	offset -= sovBitcannaid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bitcannaid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBitcannaid(uint64(m.Id))
	}
	l = len(m.Bcnaid)
	if l > 0 {
		n += 1 + l + sovBitcannaid(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovBitcannaid(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBitcannaid(uint64(l))
	}
	return n
}

func sovBitcannaid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBitcannaid(x uint64) (n int) {
	return sovBitcannaid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bitcannaid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBitcannaid
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
			return fmt.Errorf("proto: Bitcannaid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bitcannaid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcannaid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bcnaid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcannaid
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
				return ErrInvalidLengthBitcannaid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcannaid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bcnaid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcannaid
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
				return ErrInvalidLengthBitcannaid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcannaid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcannaid
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
				return ErrInvalidLengthBitcannaid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcannaid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBitcannaid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBitcannaid
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
func skipBitcannaid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBitcannaid
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
					return 0, ErrIntOverflowBitcannaid
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
					return 0, ErrIntOverflowBitcannaid
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
				return 0, ErrInvalidLengthBitcannaid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBitcannaid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBitcannaid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBitcannaid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBitcannaid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBitcannaid = fmt.Errorf("proto: unexpected end of group")
)
