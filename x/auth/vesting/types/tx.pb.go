// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/vesting/v1beta1/tx.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// MsgCreateVestingAccount defines a message that enables creating a vesting
// account.
type MsgCreateVestingAccount struct {
	FromAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"to_address,omitempty" yaml:"to_address"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	EndTime     int64                                         `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" yaml:"end_time"`
	Delayed     bool                                          `protobuf:"varint,5,opt,name=delayed,proto3" json:"delayed,omitempty"`
}

func (m *MsgCreateVestingAccount) Reset()         { *m = MsgCreateVestingAccount{} }
func (m *MsgCreateVestingAccount) String() string { return proto.CompactTextString(m) }
func (*MsgCreateVestingAccount) ProtoMessage()    {}
func (*MsgCreateVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_5338ca97811f9792, []int{0}
}
func (m *MsgCreateVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateVestingAccount.Merge(m, src)
}
func (m *MsgCreateVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateVestingAccount proto.InternalMessageInfo

func (m *MsgCreateVestingAccount) GetFromAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.FromAddress
	}
	return nil
}

func (m *MsgCreateVestingAccount) GetToAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *MsgCreateVestingAccount) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MsgCreateVestingAccount) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *MsgCreateVestingAccount) GetDelayed() bool {
	if m != nil {
		return m.Delayed
	}
	return false
}

func init() {
	proto.RegisterType((*MsgCreateVestingAccount)(nil), "cosmos.vesting.v1beta1.MsgCreateVestingAccount")
}

func init() { proto.RegisterFile("cosmos/vesting/v1beta1/tx.proto", fileDescriptor_5338ca97811f9792) }

var fileDescriptor_5338ca97811f9792 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0xcf, 0xd2, 0x40,
	0x18, 0xc7, 0x7b, 0x16, 0x01, 0x0f, 0x12, 0x63, 0x31, 0x5a, 0x19, 0x7a, 0x4d, 0xa7, 0x2e, 0x5c,
	0x45, 0x37, 0x36, 0x4a, 0x62, 0x4c, 0x8c, 0x4b, 0x63, 0x1c, 0x5c, 0xc8, 0xf5, 0x7a, 0x96, 0x06,
	0xda, 0x23, 0xbd, 0x83, 0xc0, 0xb7, 0xf0, 0x23, 0x18, 0x47, 0x3f, 0x09, 0x23, 0xa3, 0x53, 0x35,
	0xb0, 0x38, 0x33, 0x3a, 0x99, 0xf6, 0x5a, 0xde, 0x77, 0x7a, 0xf3, 0xe6, 0x9d, 0xda, 0x27, 0xff,
	0xff, 0xf3, 0xff, 0x3d, 0xcf, 0xd3, 0x42, 0x44, 0xb9, 0x48, 0xb9, 0xf0, 0xb6, 0x4c, 0xc8, 0x24,
	0x8b, 0xbd, 0xed, 0x38, 0x64, 0x92, 0x8c, 0x3d, 0xb9, 0xc3, 0xeb, 0x9c, 0x4b, 0x6e, 0xbc, 0x50,
	0x06, 0x5c, 0x1b, 0x70, 0x6d, 0x18, 0x3e, 0x8f, 0x79, 0xcc, 0x2b, 0x8b, 0x57, 0xbe, 0x29, 0xf7,
	0xd0, 0xaa, 0xe3, 0x42, 0x22, 0xd8, 0x35, 0x8b, 0xf2, 0x24, 0x53, 0xba, 0xf3, 0x43, 0x87, 0x2f,
	0x3f, 0x8a, 0x78, 0x96, 0x33, 0x22, 0xd9, 0x67, 0x15, 0x39, 0xa5, 0x94, 0x6f, 0x32, 0x69, 0x2c,
	0x61, 0xff, 0x6b, 0xce, 0xd3, 0x39, 0x89, 0xa2, 0x9c, 0x09, 0x61, 0x02, 0x1b, 0xb8, 0x7d, 0xff,
	0xfd, 0xa5, 0x40, 0x83, 0x3d, 0x49, 0x57, 0x13, 0xe7, 0xb6, 0xea, 0xfc, 0x2b, 0xd0, 0x28, 0x4e,
	0xe4, 0x62, 0x13, 0x62, 0xca, 0x53, 0xaf, 0xe6, 0xaa, 0xc7, 0x48, 0x44, 0x4b, 0x4f, 0xee, 0xd7,
	0x4c, 0xe0, 0x29, 0xa5, 0x53, 0xd5, 0x11, 0xf4, 0xca, 0xfe, 0xba, 0x30, 0x18, 0x84, 0x92, 0x5f,
	0x51, 0x8f, 0x2a, 0xd4, 0xbb, 0x4b, 0x81, 0x9e, 0x29, 0xd4, 0x8d, 0xf6, 0x00, 0xd0, 0x13, 0xc9,
	0x1b, 0x0c, 0x85, 0x6d, 0x92, 0x96, 0xdb, 0x99, 0xba, 0xad, 0xbb, 0xbd, 0x37, 0xaf, 0x70, 0x7d,
	0xce, 0xf2, 0x40, 0xcd, 0x2d, 0xf1, 0x8c, 0x27, 0x99, 0xff, 0xfa, 0x50, 0x20, 0xed, 0xe7, 0x6f,
	0xe4, 0xde, 0x03, 0x56, 0x36, 0x88, 0xa0, 0x8e, 0x36, 0x30, 0xec, 0xb2, 0x2c, 0x9a, 0xcb, 0x24,
	0x65, 0x66, 0xcb, 0x06, 0xae, 0xee, 0x0f, 0x2e, 0x05, 0x7a, 0xaa, 0x36, 0x69, 0x14, 0x27, 0xe8,
	0xb0, 0x2c, 0xfa, 0x94, 0xa4, 0xcc, 0x30, 0x61, 0x27, 0x62, 0x2b, 0xb2, 0x67, 0x91, 0xf9, 0xd8,
	0x06, 0x6e, 0x37, 0x68, 0xca, 0x49, 0xeb, 0xef, 0x77, 0x04, 0xfc, 0x0f, 0x87, 0x93, 0x05, 0x8e,
	0x27, 0x0b, 0xfc, 0x39, 0x59, 0xe0, 0xdb, 0xd9, 0xd2, 0x8e, 0x67, 0x4b, 0xfb, 0x75, 0xb6, 0xb4,
	0x2f, 0xe3, 0x3b, 0x67, 0xdb, 0x79, 0x64, 0x23, 0x17, 0xd7, 0x5f, 0xa9, 0x1a, 0x35, 0x6c, 0x57,
	0x1f, 0xfe, 0xed, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x00, 0x09, 0x49, 0x69, 0x02, 0x00,
	0x00,
}

func (this *MsgCreateVestingAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgCreateVestingAccount)
	if !ok {
		that2, ok := that.(MsgCreateVestingAccount)
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
	if !bytes.Equal(this.FromAddress, that1.FromAddress) {
		return false
	}
	if !bytes.Equal(this.ToAddress, that1.ToAddress) {
		return false
	}
	if len(this.Amount) != len(that1.Amount) {
		return false
	}
	for i := range this.Amount {
		if !this.Amount[i].Equal(&that1.Amount[i]) {
			return false
		}
	}
	if this.EndTime != that1.EndTime {
		return false
	}
	if this.Delayed != that1.Delayed {
		return false
	}
	return true
}
func (m *MsgCreateVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Delayed {
		i--
		if m.Delayed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.EndTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.EndTime != 0 {
		n += 1 + sovTx(uint64(m.EndTime))
	}
	if m.Delayed {
		n += 2
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = append(m.FromAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.FromAddress == nil {
				m.FromAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = append(m.ToAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ToAddress == nil {
				m.ToAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delayed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Delayed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)