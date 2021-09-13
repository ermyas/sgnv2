// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sgn/cbridge/v1/cbridge.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type Params struct {
	MultiChainTokenAsset MultiChainTokenAsset `protobuf:"bytes,1,opt,name=multi_chain_token_asset,json=multiChainTokenAsset,proto3" json:"multi_chain_token_asset"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea88343271b63cf0, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMultiChainTokenAsset() MultiChainTokenAsset {
	if m != nil {
		return m.MultiChainTokenAsset
	}
	return MultiChainTokenAsset{}
}

type MultiChainTokenAsset struct {
	TokenSymbolAssetMap map[string]ChainAsset `protobuf:"bytes,1,rep,name=token_symbol_asset_map,json=tokenSymbolAssetMap,proto3" json:"token_symbol_asset_map" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MultiChainTokenAsset) Reset()         { *m = MultiChainTokenAsset{} }
func (m *MultiChainTokenAsset) String() string { return proto.CompactTextString(m) }
func (*MultiChainTokenAsset) ProtoMessage()    {}
func (*MultiChainTokenAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea88343271b63cf0, []int{1}
}
func (m *MultiChainTokenAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiChainTokenAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiChainTokenAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiChainTokenAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiChainTokenAsset.Merge(m, src)
}
func (m *MultiChainTokenAsset) XXX_Size() int {
	return m.Size()
}
func (m *MultiChainTokenAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiChainTokenAsset.DiscardUnknown(m)
}

var xxx_messageInfo_MultiChainTokenAsset proto.InternalMessageInfo

func (m *MultiChainTokenAsset) GetTokenSymbolAssetMap() map[string]ChainAsset {
	if m != nil {
		return m.TokenSymbolAssetMap
	}
	return nil
}

type ChainAsset struct {
	ChainAssetMap map[uint64]Asset `protobuf:"bytes,1,rep,name=chain_asset_map,json=chainAssetMap,proto3" json:"chain_asset_map" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ChainAsset) Reset()         { *m = ChainAsset{} }
func (m *ChainAsset) String() string { return proto.CompactTextString(m) }
func (*ChainAsset) ProtoMessage()    {}
func (*ChainAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea88343271b63cf0, []int{2}
}
func (m *ChainAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainAsset.Merge(m, src)
}
func (m *ChainAsset) XXX_Size() int {
	return m.Size()
}
func (m *ChainAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainAsset.DiscardUnknown(m)
}

var xxx_messageInfo_ChainAsset proto.InternalMessageInfo

func (m *ChainAsset) GetChainAssetMap() map[uint64]Asset {
	if m != nil {
		return m.ChainAssetMap
	}
	return nil
}

type Asset struct {
	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Decimal      uint64 `protobuf:"varint,2,opt,name=decimal,proto3" json:"decimal,omitempty"`
	MaxFeeAmount uint64 `protobuf:"varint,3,opt,name=max_fee_amount,json=maxFeeAmount,proto3" json:"max_fee_amount,omitempty"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea88343271b63cf0, []int{3}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return m.Size()
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Asset) GetDecimal() uint64 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *Asset) GetMaxFeeAmount() uint64 {
	if m != nil {
		return m.MaxFeeAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "sgn.staking.v1.Params")
	proto.RegisterType((*MultiChainTokenAsset)(nil), "sgn.staking.v1.MultiChainTokenAsset")
	proto.RegisterMapType((map[string]ChainAsset)(nil), "sgn.staking.v1.MultiChainTokenAsset.TokenSymbolAssetMapEntry")
	proto.RegisterType((*ChainAsset)(nil), "sgn.staking.v1.ChainAsset")
	proto.RegisterMapType((map[uint64]Asset)(nil), "sgn.staking.v1.ChainAsset.ChainAssetMapEntry")
	proto.RegisterType((*Asset)(nil), "sgn.staking.v1.Asset")
}

func init() { proto.RegisterFile("sgn/cbridge/v1/cbridge.proto", fileDescriptor_ea88343271b63cf0) }

var fileDescriptor_ea88343271b63cf0 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0x34, 0x49, 0xc5, 0x57, 0xad, 0x32, 0x46, 0x5d, 0x17, 0x59, 0x4b, 0xe8, 0xa1, 0x20,
	0xd9, 0xb1, 0xf1, 0x22, 0x82, 0x87, 0x56, 0x14, 0x3c, 0x14, 0x24, 0x0a, 0x82, 0x22, 0xcb, 0xec,
	0x76, 0x3a, 0x5d, 0x76, 0x67, 0x26, 0xec, 0xcc, 0xc6, 0xec, 0xbf, 0xf0, 0x37, 0x79, 0xea, 0xb1,
	0x47, 0x4f, 0x22, 0xc9, 0x5f, 0xf0, 0x07, 0xc8, 0xce, 0x34, 0x9a, 0x6e, 0x1a, 0xe8, 0xed, 0x7b,
	0xf3, 0x7d, 0xef, 0xbd, 0xef, 0x7b, 0xec, 0xc2, 0x63, 0xcd, 0x25, 0x49, 0xe2, 0x22, 0x3d, 0xe6,
	0x8c, 0x4c, 0xf6, 0x17, 0x30, 0x1c, 0x17, 0xca, 0x28, 0xbc, 0xad, 0xb9, 0x0c, 0xb5, 0xa1, 0x59,
	0x2a, 0x79, 0x38, 0xd9, 0xf7, 0x7b, 0x5c, 0x71, 0x65, 0x29, 0x52, 0x23, 0xa7, 0xf2, 0x1f, 0x71,
	0xa5, 0x78, 0xce, 0x88, 0xad, 0xe2, 0xf2, 0x84, 0x50, 0x59, 0x2d, 0xa8, 0x44, 0x69, 0xa1, 0x74,
	0xe4, 0x7a, 0x5c, 0xe1, 0xa8, 0x7e, 0x06, 0x9b, 0xef, 0x69, 0x41, 0x85, 0xc6, 0x14, 0x1e, 0x8a,
	0x32, 0x37, 0x69, 0x94, 0x9c, 0xd2, 0x54, 0x46, 0x46, 0x65, 0x4c, 0x46, 0x54, 0x6b, 0x66, 0x3c,
	0xb4, 0x83, 0xf6, 0xb6, 0x86, 0xbb, 0xe1, 0x65, 0x1f, 0xe1, 0x51, 0x2d, 0x7f, 0x5d, 0xab, 0x3f,
	0xd6, 0xe2, 0x83, 0x5a, 0x7b, 0xd8, 0x39, 0xfb, 0xf5, 0xa4, 0x35, 0xea, 0x89, 0x2b, 0xb8, 0xfe,
	0x1f, 0x04, 0xbd, 0xab, 0x9a, 0xf0, 0x14, 0x1e, 0xb8, 0x7d, 0xba, 0x12, 0xb1, 0xca, 0xdd, 0xda,
	0x48, 0xd0, 0xb1, 0x87, 0x76, 0xda, 0x7b, 0x5b, 0xc3, 0x57, 0xd7, 0x59, 0x1d, 0x5a, 0xf8, 0xc1,
	0x4e, 0xb0, 0x0f, 0x47, 0x74, 0xfc, 0x46, 0x9a, 0xa2, 0xba, 0xf0, 0x74, 0xcf, 0xac, 0xf2, 0x7e,
	0x0c, 0xde, 0xba, 0x36, 0x7c, 0x17, 0xda, 0x19, 0xab, 0x6c, 0xfa, 0x9b, 0xa3, 0x1a, 0xe2, 0x67,
	0xd0, 0x9d, 0xd0, 0xbc, 0x64, 0xde, 0x86, 0xbd, 0x88, 0xdf, 0xb4, 0x65, 0x1d, 0xd9, 0x21, 0x23,
	0x27, 0x7c, 0xb9, 0xf1, 0x02, 0xf5, 0x7f, 0x20, 0x80, 0xff, 0x0c, 0xfe, 0x02, 0x77, 0xdc, 0x89,
	0x9b, 0x29, 0x07, 0xeb, 0xc7, 0x2d, 0xc1, 0x46, 0xaa, 0xdb, 0xc9, 0x32, 0xe3, 0x7f, 0x02, 0xbc,
	0x2a, 0x5d, 0x4e, 0xd2, 0x71, 0x49, 0x9e, 0x5e, 0x4e, 0x72, 0xbf, 0xb9, 0x7a, 0x25, 0xc4, 0x57,
	0xe8, 0x3a, 0xfb, 0x3d, 0xe8, 0xda, 0x43, 0x5e, 0xdc, 0xc5, 0x15, 0xd8, 0x83, 0x1b, 0xc7, 0x2c,
	0x49, 0x05, 0xcd, 0xed, 0xc4, 0xce, 0x68, 0x51, 0xe2, 0x5d, 0xd8, 0x16, 0x74, 0x1a, 0x9d, 0x30,
	0x16, 0x51, 0xa1, 0x4a, 0x69, 0xbc, 0xb6, 0x15, 0xdc, 0x12, 0x74, 0xfa, 0x96, 0xb1, 0x03, 0xfb,
	0x76, 0xf8, 0xee, 0x6c, 0x16, 0xa0, 0xf3, 0x59, 0x80, 0x7e, 0xcf, 0x02, 0xf4, 0x7d, 0x1e, 0xb4,
	0xce, 0xe7, 0x41, 0xeb, 0xe7, 0x3c, 0x68, 0x7d, 0x26, 0x3c, 0x35, 0xa7, 0x65, 0x1c, 0x26, 0x4a,
	0x90, 0x84, 0xe5, 0xac, 0x18, 0x48, 0x66, 0xbe, 0xa9, 0x22, 0x23, 0x9a, 0xcb, 0xc1, 0x64, 0x48,
	0xa6, 0xff, 0xfe, 0x1c, 0x53, 0x8d, 0x99, 0x8e, 0x37, 0xed, 0x97, 0xfd, 0xfc, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7a, 0x20, 0x26, 0x66, 0x55, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MultiChainTokenAsset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCbridge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MultiChainTokenAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiChainTokenAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiChainTokenAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenSymbolAssetMap) > 0 {
		for k := range m.TokenSymbolAssetMap {
			v := m.TokenSymbolAssetMap[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCbridge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintCbridge(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintCbridge(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChainAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainAssetMap) > 0 {
		for k := range m.ChainAssetMap {
			v := m.ChainAssetMap[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCbridge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i = encodeVarintCbridge(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintCbridge(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Asset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxFeeAmount != 0 {
		i = encodeVarintCbridge(dAtA, i, uint64(m.MaxFeeAmount))
		i--
		dAtA[i] = 0x18
	}
	if m.Decimal != 0 {
		i = encodeVarintCbridge(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintCbridge(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCbridge(dAtA []byte, offset int, v uint64) int {
	offset -= sovCbridge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MultiChainTokenAsset.Size()
	n += 1 + l + sovCbridge(uint64(l))
	return n
}

func (m *MultiChainTokenAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TokenSymbolAssetMap) > 0 {
		for k, v := range m.TokenSymbolAssetMap {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovCbridge(uint64(len(k))) + 1 + l + sovCbridge(uint64(l))
			n += mapEntrySize + 1 + sovCbridge(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ChainAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChainAssetMap) > 0 {
		for k, v := range m.ChainAssetMap {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + sovCbridge(uint64(k)) + 1 + l + sovCbridge(uint64(l))
			n += mapEntrySize + 1 + sovCbridge(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Asset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovCbridge(uint64(l))
	}
	if m.Decimal != 0 {
		n += 1 + sovCbridge(uint64(m.Decimal))
	}
	if m.MaxFeeAmount != 0 {
		n += 1 + sovCbridge(uint64(m.MaxFeeAmount))
	}
	return n
}

func sovCbridge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCbridge(x uint64) (n int) {
	return sovCbridge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCbridge
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultiChainTokenAsset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCbridge
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
				return ErrInvalidLengthCbridge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCbridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MultiChainTokenAsset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCbridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCbridge
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
func (m *MultiChainTokenAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCbridge
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
			return fmt.Errorf("proto: MultiChainTokenAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiChainTokenAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenSymbolAssetMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCbridge
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
				return ErrInvalidLengthCbridge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCbridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TokenSymbolAssetMap == nil {
				m.TokenSymbolAssetMap = make(map[string]ChainAsset)
			}
			var mapkey string
			mapvalue := &ChainAsset{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCbridge
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCbridge
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthCbridge
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthCbridge
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCbridge
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthCbridge
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthCbridge
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ChainAsset{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCbridge(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthCbridge
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.TokenSymbolAssetMap[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCbridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCbridge
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
func (m *ChainAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCbridge
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
			return fmt.Errorf("proto: ChainAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainAssetMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCbridge
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
				return ErrInvalidLengthCbridge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCbridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChainAssetMap == nil {
				m.ChainAssetMap = make(map[uint64]Asset)
			}
			var mapkey uint64
			mapvalue := &Asset{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCbridge
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCbridge
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCbridge
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthCbridge
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthCbridge
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Asset{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCbridge(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthCbridge
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ChainAssetMap[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCbridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCbridge
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
func (m *Asset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCbridge
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
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCbridge
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
				return ErrInvalidLengthCbridge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCbridge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCbridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimal |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxFeeAmount", wireType)
			}
			m.MaxFeeAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCbridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxFeeAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCbridge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCbridge
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
func skipCbridge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCbridge
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
					return 0, ErrIntOverflowCbridge
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
					return 0, ErrIntOverflowCbridge
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
				return 0, ErrInvalidLengthCbridge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCbridge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCbridge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCbridge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCbridge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCbridge = fmt.Errorf("proto: unexpected end of group")
)
