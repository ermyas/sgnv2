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
	ChainAssetMap map[uint64]ChainTokenAsset `protobuf:"bytes,1,rep,name=chain_asset_map,json=chainAssetMap,proto3" json:"chain_asset_map" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (m *MultiChainTokenAsset) GetChainAssetMap() map[uint64]ChainTokenAsset {
	if m != nil {
		return m.ChainAssetMap
	}
	return nil
}

type ChainTokenAsset struct {
	TokenAssetMap map[string]Asset `protobuf:"bytes,1,rep,name=token_asset_map,json=tokenAssetMap,proto3" json:"token_asset_map" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ChainTokenAsset) Reset()         { *m = ChainTokenAsset{} }
func (m *ChainTokenAsset) String() string { return proto.CompactTextString(m) }
func (*ChainTokenAsset) ProtoMessage()    {}
func (*ChainTokenAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea88343271b63cf0, []int{2}
}
func (m *ChainTokenAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainTokenAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainTokenAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainTokenAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainTokenAsset.Merge(m, src)
}
func (m *ChainTokenAsset) XXX_Size() int {
	return m.Size()
}
func (m *ChainTokenAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainTokenAsset.DiscardUnknown(m)
}

var xxx_messageInfo_ChainTokenAsset proto.InternalMessageInfo

func (m *ChainTokenAsset) GetTokenAssetMap() map[string]Asset {
	if m != nil {
		return m.TokenAssetMap
	}
	return nil
}

type Asset struct {
	ChainId      uint64 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Token        string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Decimal      uint64 `protobuf:"varint,3,opt,name=decimal,proto3" json:"decimal,omitempty"`
	MaxFeeAmount uint64 `protobuf:"varint,4,opt,name=max_fee_amount,json=maxFeeAmount,proto3" json:"max_fee_amount,omitempty"`
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

func (m *Asset) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

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
	proto.RegisterMapType((map[uint64]ChainTokenAsset)(nil), "sgn.staking.v1.MultiChainTokenAsset.ChainAssetMapEntry")
	proto.RegisterType((*ChainTokenAsset)(nil), "sgn.staking.v1.ChainTokenAsset")
	proto.RegisterMapType((map[string]Asset)(nil), "sgn.staking.v1.ChainTokenAsset.TokenAssetMapEntry")
	proto.RegisterType((*Asset)(nil), "sgn.staking.v1.Asset")
}

func init() { proto.RegisterFile("sgn/cbridge/v1/cbridge.proto", fileDescriptor_ea88343271b63cf0) }

var fileDescriptor_ea88343271b63cf0 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xa3, 0x26, 0x69, 0x57, 0x75, 0x6b, 0x86, 0xc8, 0x98, 0x1b, 0x86, 0x5b, 0x42, 0x0f,
	0x85, 0x51, 0x8b, 0x66, 0x8c, 0x8d, 0xdd, 0xda, 0xb1, 0x41, 0x0f, 0x85, 0x11, 0x06, 0x83, 0x5d,
	0x3c, 0xc5, 0x79, 0x55, 0x8d, 0x2d, 0x29, 0x58, 0x72, 0x16, 0x7f, 0x8b, 0x7d, 0xac, 0x1e, 0x76,
	0xe8, 0x71, 0xa7, 0x51, 0x92, 0x2f, 0x32, 0x2c, 0xa5, 0x9b, 0xeb, 0x04, 0xd6, 0xdb, 0xfb, 0xfb,
	0xff, 0x97, 0xde, 0xfb, 0x3d, 0xdb, 0xf8, 0x85, 0xe6, 0x92, 0x46, 0xa3, 0x2c, 0x1e, 0x73, 0xa0,
	0xd3, 0x93, 0xbb, 0x32, 0x98, 0x64, 0xca, 0x28, 0xb2, 0xab, 0xb9, 0x0c, 0xb4, 0x61, 0x49, 0x2c,
	0x79, 0x30, 0x3d, 0xe9, 0x75, 0xb9, 0xe2, 0xca, 0x5a, 0xb4, 0xac, 0x5c, 0xaa, 0xb7, 0xc7, 0x95,
	0xe2, 0x29, 0x50, 0xab, 0x46, 0xf9, 0x25, 0x65, 0xb2, 0xb8, 0xb3, 0x22, 0xa5, 0x85, 0xd2, 0xa1,
	0x3b, 0xe3, 0x84, 0xb3, 0xfa, 0x09, 0xde, 0xfc, 0xc4, 0x32, 0x26, 0x34, 0x61, 0xf8, 0xb9, 0xc8,
	0x53, 0x13, 0x87, 0xd1, 0x15, 0x8b, 0x65, 0x68, 0x54, 0x02, 0x32, 0x64, 0x5a, 0x83, 0xf1, 0xd0,
	0x01, 0x3a, 0xda, 0x19, 0x1c, 0x06, 0xf7, 0xe7, 0x08, 0x2e, 0xca, 0xf8, 0xfb, 0x32, 0xfd, 0xb9,
	0x0c, 0x9f, 0x96, 0xd9, 0xb3, 0xd6, 0xf5, 0xef, 0xfd, 0xc6, 0xb0, 0x2b, 0xd6, 0x78, 0xfd, 0x5b,
	0x84, 0xbb, 0xeb, 0x0e, 0x11, 0xc0, 0x1d, 0xd7, 0xd5, 0xf6, 0x0b, 0x05, 0x9b, 0x78, 0xe8, 0xa0,
	0x79, 0xb4, 0x33, 0x78, 0xf3, 0x90, 0x9e, 0x81, 0xd5, 0xb6, 0xbc, 0x60, 0x93, 0x0f, 0xd2, 0x64,
	0xc5, 0x72, 0x8c, 0x27, 0x51, 0xd5, 0xe9, 0x31, 0x4c, 0x56, 0xa3, 0xe4, 0x29, 0x6e, 0x26, 0x50,
	0x58, 0xc8, 0xd6, 0xb0, 0x2c, 0xc9, 0x6b, 0xdc, 0x9e, 0xb2, 0x34, 0x07, 0x6f, 0xc3, 0x82, 0xef,
	0xd7, 0x87, 0xa8, 0xf5, 0x1f, 0xba, 0xf4, 0xbb, 0x8d, 0xb7, 0xa8, 0xff, 0x13, 0xe1, 0x4e, 0x9d,
	0xee, 0x1b, 0xee, 0x54, 0xb6, 0x59, 0xa1, 0x1b, 0xfc, 0xe7, 0xe2, 0xe0, 0x5f, 0x59, 0x07, 0x33,
	0x55, 0xa7, 0xf7, 0x05, 0x93, 0xd5, 0x68, 0x15, 0x6c, 0xdb, 0x81, 0xbd, 0xbc, 0x0f, 0xf6, 0xac,
	0xde, 0x7f, 0x05, 0x67, 0x86, 0xdb, 0x8e, 0x61, 0x0f, 0x3f, 0x72, 0x6f, 0x28, 0x1e, 0x2f, 0x37,
	0xb5, 0x65, 0xf5, 0xf9, 0x98, 0x74, 0x71, 0xdb, 0x4e, 0x63, 0x2f, 0xdd, 0x1e, 0x3a, 0x41, 0x3c,
	0xbc, 0x35, 0x86, 0x28, 0x16, 0x2c, 0xf5, 0x9a, 0x2e, 0xbf, 0x94, 0xe4, 0x10, 0xef, 0x0a, 0x36,
	0x0b, 0x2f, 0x01, 0x42, 0x26, 0x54, 0x2e, 0x8d, 0xd7, 0xb2, 0x81, 0xc7, 0x82, 0xcd, 0x3e, 0x02,
	0x9c, 0xda, 0x67, 0x67, 0xe7, 0xd7, 0x73, 0x1f, 0xdd, 0xcc, 0x7d, 0x74, 0x3b, 0xf7, 0xd1, 0x8f,
	0x85, 0xdf, 0xb8, 0x59, 0xf8, 0x8d, 0x5f, 0x0b, 0xbf, 0xf1, 0x95, 0xf2, 0xd8, 0x5c, 0xe5, 0xa3,
	0x20, 0x52, 0x82, 0x46, 0x90, 0x42, 0x76, 0x2c, 0xc1, 0x7c, 0x57, 0x59, 0x42, 0x35, 0x97, 0xc7,
	0xd3, 0x01, 0x9d, 0xfd, 0xfd, 0x95, 0x4c, 0x31, 0x01, 0x3d, 0xda, 0xb4, 0x9f, 0xfa, 0xab, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xe9, 0x5a, 0x4f, 0x66, 0x03, 0x00, 0x00,
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

func (m *ChainTokenAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainTokenAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainTokenAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenAssetMap) > 0 {
		for k := range m.TokenAssetMap {
			v := m.TokenAssetMap[k]
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
		dAtA[i] = 0x20
	}
	if m.Decimal != 0 {
		i = encodeVarintCbridge(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintCbridge(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChainId != 0 {
		i = encodeVarintCbridge(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x8
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

func (m *ChainTokenAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TokenAssetMap) > 0 {
		for k, v := range m.TokenAssetMap {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovCbridge(uint64(len(k))) + 1 + l + sovCbridge(uint64(l))
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
	if m.ChainId != 0 {
		n += 1 + sovCbridge(uint64(m.ChainId))
	}
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
				m.ChainAssetMap = make(map[uint64]ChainTokenAsset)
			}
			var mapkey uint64
			mapvalue := &ChainTokenAsset{}
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
					mapvalue = &ChainTokenAsset{}
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
func (m *ChainTokenAsset) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ChainTokenAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainTokenAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAssetMap", wireType)
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
			if m.TokenAssetMap == nil {
				m.TokenAssetMap = make(map[string]Asset)
			}
			var mapkey string
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
			m.TokenAssetMap[mapkey] = *mapvalue
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCbridge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
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
		case 3:
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
		case 4:
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
