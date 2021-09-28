// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eth/staking.proto

package eth

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

type StakingReward struct {
	Recipient              []byte `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	CumulativeRewardAmount []byte `protobuf:"bytes,2,opt,name=cumulative_reward_amount,json=cumulativeRewardAmount,proto3" json:"cumulative_reward_amount,omitempty"`
}

func (m *StakingReward) Reset()         { *m = StakingReward{} }
func (m *StakingReward) String() string { return proto.CompactTextString(m) }
func (*StakingReward) ProtoMessage()    {}
func (*StakingReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_a119100e37990352, []int{0}
}
func (m *StakingReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakingReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakingReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakingReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingReward.Merge(m, src)
}
func (m *StakingReward) XXX_Size() int {
	return m.Size()
}
func (m *StakingReward) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingReward.DiscardUnknown(m)
}

var xxx_messageInfo_StakingReward proto.InternalMessageInfo

func (m *StakingReward) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *StakingReward) GetCumulativeRewardAmount() []byte {
	if m != nil {
		return m.CumulativeRewardAmount
	}
	return nil
}

type Slash struct {
	Validator   []byte         `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	Nonce       uint64         `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	SlashFactor uint64         `protobuf:"varint,3,opt,name=slash_factor,json=slashFactor,proto3" json:"slash_factor,omitempty"`
	ExpireTime  uint64         `protobuf:"varint,4,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	JailPeriod  uint64         `protobuf:"varint,5,opt,name=jail_period,json=jailPeriod,proto3" json:"jail_period,omitempty"`
	Collectors  []*AcctAmtPair `protobuf:"bytes,6,rep,name=collectors,proto3" json:"collectors,omitempty"`
}

func (m *Slash) Reset()         { *m = Slash{} }
func (m *Slash) String() string { return proto.CompactTextString(m) }
func (*Slash) ProtoMessage()    {}
func (*Slash) Descriptor() ([]byte, []int) {
	return fileDescriptor_a119100e37990352, []int{1}
}
func (m *Slash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Slash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Slash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Slash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slash.Merge(m, src)
}
func (m *Slash) XXX_Size() int {
	return m.Size()
}
func (m *Slash) XXX_DiscardUnknown() {
	xxx_messageInfo_Slash.DiscardUnknown(m)
}

var xxx_messageInfo_Slash proto.InternalMessageInfo

func (m *Slash) GetValidator() []byte {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *Slash) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Slash) GetSlashFactor() uint64 {
	if m != nil {
		return m.SlashFactor
	}
	return 0
}

func (m *Slash) GetExpireTime() uint64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *Slash) GetJailPeriod() uint64 {
	if m != nil {
		return m.JailPeriod
	}
	return 0
}

func (m *Slash) GetCollectors() []*AcctAmtPair {
	if m != nil {
		return m.Collectors
	}
	return nil
}

type AcctAmtPair struct {
	Account []byte `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Amount  []byte `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *AcctAmtPair) Reset()         { *m = AcctAmtPair{} }
func (m *AcctAmtPair) String() string { return proto.CompactTextString(m) }
func (*AcctAmtPair) ProtoMessage()    {}
func (*AcctAmtPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_a119100e37990352, []int{2}
}
func (m *AcctAmtPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AcctAmtPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AcctAmtPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AcctAmtPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcctAmtPair.Merge(m, src)
}
func (m *AcctAmtPair) XXX_Size() int {
	return m.Size()
}
func (m *AcctAmtPair) XXX_DiscardUnknown() {
	xxx_messageInfo_AcctAmtPair.DiscardUnknown(m)
}

var xxx_messageInfo_AcctAmtPair proto.InternalMessageInfo

func (m *AcctAmtPair) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *AcctAmtPair) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterType((*StakingReward)(nil), "eth.StakingReward")
	proto.RegisterType((*Slash)(nil), "eth.Slash")
	proto.RegisterType((*AcctAmtPair)(nil), "eth.AcctAmtPair")
}

func init() { proto.RegisterFile("eth/staking.proto", fileDescriptor_a119100e37990352) }

var fileDescriptor_a119100e37990352 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xc7, 0xa9, 0x7c, 0x18, 0xb7, 0x98, 0xe8, 0xc6, 0x90, 0x1e, 0x4c, 0x45, 0x4e, 0x78, 0xa0,
	0x35, 0x78, 0xf1, 0x66, 0xd0, 0xc4, 0x33, 0x29, 0x9e, 0xbc, 0x34, 0xcb, 0x32, 0xb6, 0x2b, 0xed,
	0x6e, 0xb3, 0x9d, 0x82, 0x8f, 0xe1, 0x63, 0x79, 0x24, 0x9e, 0x3c, 0x1a, 0x78, 0x11, 0xd3, 0xad,
	0x08, 0xde, 0x3a, 0xbf, 0xff, 0xc7, 0xa4, 0xb3, 0xe4, 0x14, 0x30, 0xf6, 0x73, 0x64, 0x73, 0x21,
	0x23, 0x2f, 0xd3, 0x0a, 0x15, 0xad, 0x03, 0xc6, 0xbd, 0x88, 0x1c, 0x4f, 0x2a, 0x1a, 0xc0, 0x92,
	0xe9, 0x19, 0x3d, 0x27, 0x47, 0x1a, 0xb8, 0xc8, 0x04, 0x48, 0x74, 0xac, 0xae, 0xd5, 0x6f, 0x07,
	0x3b, 0x40, 0x6f, 0x89, 0xc3, 0x8b, 0xb4, 0x48, 0x18, 0x8a, 0x05, 0x84, 0xda, 0x44, 0x42, 0x96,
	0xaa, 0x42, 0xa2, 0x73, 0x60, 0xcc, 0x9d, 0x9d, 0x5e, 0x35, 0x8e, 0x8c, 0xda, 0xfb, 0xb4, 0x48,
	0x73, 0x92, 0xb0, 0x3c, 0x2e, 0x37, 0x2c, 0x58, 0x22, 0x66, 0x0c, 0x95, 0xde, 0x6e, 0xf8, 0x03,
	0xf4, 0x8c, 0x34, 0xa5, 0x92, 0x1c, 0x4c, 0x5d, 0x23, 0xa8, 0x06, 0x7a, 0x49, 0xda, 0x79, 0x19,
	0x0e, 0x5f, 0x18, 0x2f, 0x63, 0x75, 0x23, 0xda, 0x86, 0x3d, 0x1a, 0x44, 0x2f, 0x88, 0x0d, 0x6f,
	0x99, 0xd0, 0x10, 0xa2, 0x48, 0xc1, 0x69, 0x18, 0x07, 0xa9, 0xd0, 0x93, 0x48, 0xa1, 0x34, 0xbc,
	0x32, 0x91, 0x84, 0x19, 0x68, 0xa1, 0x66, 0x4e, 0xb3, 0x32, 0x94, 0x68, 0x6c, 0x08, 0xbd, 0x26,
	0x84, 0xab, 0x24, 0x81, 0xb2, 0x2e, 0x77, 0x5a, 0xdd, 0x7a, 0xdf, 0x1e, 0x9e, 0x78, 0x80, 0xb1,
	0x37, 0xe2, 0x1c, 0x47, 0x29, 0x8e, 0x99, 0xd0, 0xc1, 0x9e, 0xa7, 0x77, 0x47, 0xec, 0x3d, 0x89,
	0x3a, 0xe4, 0x90, 0x71, 0x6e, 0x8e, 0x51, 0xfd, 0xd7, 0x76, 0xa4, 0x1d, 0xd2, 0xfa, 0x77, 0xa5,
	0xdf, 0xe9, 0xfe, 0xe1, 0x63, 0xed, 0x5a, 0xab, 0xb5, 0x6b, 0x7d, 0xaf, 0x5d, 0xeb, 0x7d, 0xe3,
	0xd6, 0x56, 0x1b, 0xb7, 0xf6, 0xb5, 0x71, 0x6b, 0xcf, 0x57, 0x91, 0xc0, 0xb8, 0x98, 0x7a, 0x5c,
	0xa5, 0x3e, 0x87, 0x04, 0xf4, 0x40, 0x02, 0x2e, 0x95, 0x9e, 0xfb, 0x79, 0x24, 0x07, 0x8b, 0xa1,
	0x6f, 0x1e, 0xd1, 0x07, 0x8c, 0xa7, 0x2d, 0xf3, 0x79, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x09,
	0xf4, 0x4d, 0xcc, 0xe4, 0x01, 0x00, 0x00,
}

func (m *StakingReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakingReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakingReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CumulativeRewardAmount) > 0 {
		i -= len(m.CumulativeRewardAmount)
		copy(dAtA[i:], m.CumulativeRewardAmount)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.CumulativeRewardAmount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Slash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Slash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Slash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Collectors) > 0 {
		for iNdEx := len(m.Collectors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Collectors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.JailPeriod != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.JailPeriod))
		i--
		dAtA[i] = 0x28
	}
	if m.ExpireTime != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.ExpireTime))
		i--
		dAtA[i] = 0x20
	}
	if m.SlashFactor != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.SlashFactor))
		i--
		dAtA[i] = 0x18
	}
	if m.Nonce != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AcctAmtPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AcctAmtPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AcctAmtPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovStaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakingReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.CumulativeRewardAmount)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	return n
}

func (m *Slash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovStaking(uint64(m.Nonce))
	}
	if m.SlashFactor != 0 {
		n += 1 + sovStaking(uint64(m.SlashFactor))
	}
	if m.ExpireTime != 0 {
		n += 1 + sovStaking(uint64(m.ExpireTime))
	}
	if m.JailPeriod != 0 {
		n += 1 + sovStaking(uint64(m.JailPeriod))
	}
	if len(m.Collectors) > 0 {
		for _, e := range m.Collectors {
			l = e.Size()
			n += 1 + l + sovStaking(uint64(l))
		}
	}
	return n
}

func (m *AcctAmtPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	return n
}

func sovStaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStaking(x uint64) (n int) {
	return sovStaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakingReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
			return fmt.Errorf("proto: StakingReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakingReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = append(m.Recipient[:0], dAtA[iNdEx:postIndex]...)
			if m.Recipient == nil {
				m.Recipient = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeRewardAmount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CumulativeRewardAmount = append(m.CumulativeRewardAmount[:0], dAtA[iNdEx:postIndex]...)
			if m.CumulativeRewardAmount == nil {
				m.CumulativeRewardAmount = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
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
func (m *Slash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
			return fmt.Errorf("proto: Slash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Slash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = append(m.Validator[:0], dAtA[iNdEx:postIndex]...)
			if m.Validator == nil {
				m.Validator = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFactor", wireType)
			}
			m.SlashFactor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SlashFactor |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireTime", wireType)
			}
			m.ExpireTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailPeriod", wireType)
			}
			m.JailPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JailPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collectors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collectors = append(m.Collectors, &AcctAmtPair{})
			if err := m.Collectors[len(m.Collectors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
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
func (m *AcctAmtPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
			return fmt.Errorf("proto: AcctAmtPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AcctAmtPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = append(m.Account[:0], dAtA[iNdEx:postIndex]...)
			if m.Account == nil {
				m.Account = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount[:0], dAtA[iNdEx:postIndex]...)
			if m.Amount == nil {
				m.Amount = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
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
func skipStaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStaking
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
					return 0, ErrIntOverflowStaking
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
					return 0, ErrIntOverflowStaking
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
				return 0, ErrInvalidLengthStaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStaking = fmt.Errorf("proto: unexpected end of group")
)
