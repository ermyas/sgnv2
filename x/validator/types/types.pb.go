// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: types.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthAddress     string       `protobuf:"bytes,1,opt,name=eth_address,json=ethAddress,proto3" json:"eth_address,omitempty"`
	EthSigner      string       `protobuf:"bytes,2,opt,name=eth_signer,json=ethSigner,proto3" json:"eth_signer,omitempty"`
	SgnAddress     string       `protobuf:"bytes,3,opt,name=sgn_address,json=sgnAddress,proto3" json:"sgn_address,omitempty"`
	Transactors    []string     `protobuf:"bytes,4,rep,name=transactors,proto3" json:"transactors,omitempty"`
	Tokens         string       `protobuf:"bytes,5,opt,name=tokens,proto3" json:"tokens,omitempty"`
	Shares         string       `protobuf:"bytes,6,opt,name=shares,proto3" json:"shares,omitempty"`
	CommissionRate uint64       `protobuf:"varint,7,opt,name=commission_rate,json=commissionRate,proto3" json:"commission_rate,omitempty"`
	Description    *Description `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *Validator) GetEthAddress() string {
	if x != nil {
		return x.EthAddress
	}
	return ""
}

func (x *Validator) GetEthSigner() string {
	if x != nil {
		return x.EthSigner
	}
	return ""
}

func (x *Validator) GetSgnAddress() string {
	if x != nil {
		return x.SgnAddress
	}
	return ""
}

func (x *Validator) GetTransactors() []string {
	if x != nil {
		return x.Transactors
	}
	return nil
}

func (x *Validator) GetTokens() string {
	if x != nil {
		return x.Tokens
	}
	return ""
}

func (x *Validator) GetShares() string {
	if x != nil {
		return x.Shares
	}
	return ""
}

func (x *Validator) GetCommissionRate() uint64 {
	if x != nil {
		return x.CommissionRate
	}
	return 0
}

func (x *Validator) GetDescription() *Description {
	if x != nil {
		return x.Description
	}
	return nil
}

type Description struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moniker         string `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Identity        string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	Website         string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	SecurityContact string `protobuf:"bytes,4,opt,name=security_contact,json=securityContact,proto3" json:"security_contact,omitempty"`
	Details         string `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Description) Reset() {
	*x = Description{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Description) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Description) ProtoMessage() {}

func (x *Description) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Description.ProtoReflect.Descriptor instead.
func (*Description) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *Description) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

func (x *Description) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Description) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Description) GetSecurityContact() string {
	if x != nil {
		return x.SecurityContact
	}
	return ""
}

func (x *Description) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type Delegator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthAddress string `protobuf:"bytes,1,opt,name=eth_address,json=ethAddress,proto3" json:"eth_address,omitempty"`
	ValAddress string `protobuf:"bytes,2,opt,name=val_address,json=valAddress,proto3" json:"val_address,omitempty"`
}

func (x *Delegator) Reset() {
	*x = Delegator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delegator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delegator) ProtoMessage() {}

func (x *Delegator) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delegator.ProtoReflect.Descriptor instead.
func (*Delegator) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *Delegator) GetEthAddress() string {
	if x != nil {
		return x.EthAddress
	}
	return ""
}

func (x *Delegator) GetValAddress() string {
	if x != nil {
		return x.ValAddress
	}
	return ""
}

type Syncer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValIndex   uint64 `protobuf:"varint,1,opt,name=val_index,json=valIndex,proto3" json:"val_index,omitempty"`
	SgnAddress string `protobuf:"bytes,2,opt,name=sgn_address,json=sgnAddress,proto3" json:"sgn_address,omitempty"`
}

func (x *Syncer) Reset() {
	*x = Syncer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Syncer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Syncer) ProtoMessage() {}

func (x *Syncer) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Syncer.ProtoReflect.Descriptor instead.
func (*Syncer) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{3}
}

func (x *Syncer) GetValIndex() uint64 {
	if x != nil {
		return x.ValIndex
	}
	return 0
}

func (x *Syncer) GetSgnAddress() string {
	if x != nil {
		return x.SgnAddress
	}
	return ""
}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SyncerDuration   uint64 `protobuf:"varint,1,opt,name=syncer_duration,json=syncerDuration,proto3" json:"syncer_duration,omitempty"`
	EpochLength      uint64 `protobuf:"varint,2,opt,name=epoch_length,json=epochLength,proto3" json:"epoch_length,omitempty"`
	MaxValidatorDiff uint64 `protobuf:"varint,3,opt,name=max_validator_diff,json=maxValidatorDiff,proto3" json:"max_validator_diff,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{4}
}

func (x *Params) GetSyncerDuration() uint64 {
	if x != nil {
		return x.SyncerDuration
	}
	return 0
}

func (x *Params) GetEpochLength() uint64 {
	if x != nil {
		return x.EpochLength
	}
	return 0
}

func (x *Params) GetMaxValidatorDiff() uint64 {
	if x != nil {
		return x.MaxValidatorDiff
	}
	return 0
}

type MsgSetTransactors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactors []string `protobuf:"bytes,1,rep,name=transactors,proto3" json:"transactors,omitempty"`
	Sender      string   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *MsgSetTransactors) Reset() {
	*x = MsgSetTransactors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetTransactors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetTransactors) ProtoMessage() {}

func (x *MsgSetTransactors) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSetTransactors.ProtoReflect.Descriptor instead.
func (*MsgSetTransactors) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{5}
}

func (x *MsgSetTransactors) GetTransactors() []string {
	if x != nil {
		return x.Transactors
	}
	return nil
}

func (x *MsgSetTransactors) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

type MsgEditDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthAddress  string       `protobuf:"bytes,1,opt,name=eth_address,json=ethAddress,proto3" json:"eth_address,omitempty"`
	Description *Description `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Sender      string       `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *MsgEditDescription) Reset() {
	*x = MsgEditDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgEditDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgEditDescription) ProtoMessage() {}

func (x *MsgEditDescription) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgEditDescription.ProtoReflect.Descriptor instead.
func (*MsgEditDescription) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{6}
}

func (x *MsgEditDescription) GetEthAddress() string {
	if x != nil {
		return x.EthAddress
	}
	return ""
}

func (x *MsgEditDescription) GetDescription() *Description {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *MsgEditDescription) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params     *Params      `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Syncer     *Syncer      `protobuf:"bytes,2,opt,name=syncer,proto3" json:"syncer,omitempty"`
	Validators []*Validator `protobuf:"bytes,3,rep,name=validators,proto3" json:"validators,omitempty"`
	Delegators []*Delegator `protobuf:"bytes,4,rep,name=delegators,proto3" json:"delegators,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{7}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetSyncer() *Syncer {
	if x != nil {
		return x.Syncer
	}
	return nil
}

func (x *GenesisState) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *GenesisState) GetDelegators() []*Delegator {
	if x != nil {
		return x.Delegators
	}
	return nil
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02,
	0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x74, 0x68, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x74, 0x68, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x74, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x67, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x67, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa2, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x4d, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x74, 0x68,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61,
	0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x76, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x53,
	0x79, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x67, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x67, 0x6e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x65,
	0x70, 0x6f, 0x63, 0x68, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61,
	0x78, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x69, 0x66, 0x66,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x44, 0x69, 0x66, 0x66, 0x22, 0x4d, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x53,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x7d, 0x0a, 0x12, 0x4d, 0x73, 0x67, 0x45, 0x64,
	0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x74, 0x68, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0xa8, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73,
	0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x63,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x65,
	0x72, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x0a, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x65, 0x6c, 0x65, 0x72, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x67,
	0x6e, 0x2d, 0x76, 0x32, 0x2f, 0x78, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_types_proto_goTypes = []interface{}{
	(*Validator)(nil),          // 0: Validator
	(*Description)(nil),        // 1: Description
	(*Delegator)(nil),          // 2: Delegator
	(*Syncer)(nil),             // 3: Syncer
	(*Params)(nil),             // 4: Params
	(*MsgSetTransactors)(nil),  // 5: MsgSetTransactors
	(*MsgEditDescription)(nil), // 6: MsgEditDescription
	(*GenesisState)(nil),       // 7: GenesisState
}
var file_types_proto_depIdxs = []int32{
	1, // 0: Validator.description:type_name -> Description
	1, // 1: MsgEditDescription.description:type_name -> Description
	4, // 2: GenesisState.params:type_name -> Params
	3, // 3: GenesisState.syncer:type_name -> Syncer
	0, // 4: GenesisState.validators:type_name -> Validator
	2, // 5: GenesisState.delegators:type_name -> Delegator
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Description); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delegator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Syncer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetTransactors); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgEditDescription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
