// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crypto/types/types.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// PublicKey specifies a public key
type PublicKey struct {
	// sum specifies which type of public key is wrapped
	//
	// Types that are valid to be assigned to Sum:
	//	*PublicKey_Secp256K1
	//	*PublicKey_Ed25519
	//	*PublicKey_Sr25519
	//	*PublicKey_Multisig
	//	*PublicKey_Secp256R1
	//	*PublicKey_AnyPubkey
	Sum isPublicKey_Sum `protobuf_oneof:"sum"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_2165b2a1badb1b0c, []int{0}
}
func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(m, src)
}
func (m *PublicKey) XXX_Size() int {
	return m.Size()
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

type isPublicKey_Sum interface {
	isPublicKey_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type PublicKey_Secp256K1 struct {
	Secp256K1 []byte `protobuf:"bytes,1,opt,name=secp256k1,proto3,oneof" json:"secp256k1,omitempty"`
}
type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,2,opt,name=ed25519,proto3,oneof" json:"ed25519,omitempty"`
}
type PublicKey_Sr25519 struct {
	Sr25519 []byte `protobuf:"bytes,3,opt,name=sr25519,proto3,oneof" json:"sr25519,omitempty"`
}
type PublicKey_Multisig struct {
	Multisig *PubKeyMultisigThreshold `protobuf:"bytes,4,opt,name=multisig,proto3,oneof" json:"multisig,omitempty"`
}
type PublicKey_Secp256R1 struct {
	Secp256R1 []byte `protobuf:"bytes,5,opt,name=secp256r1,proto3,oneof" json:"secp256r1,omitempty"`
}
type PublicKey_AnyPubkey struct {
	AnyPubkey *types.Any `protobuf:"bytes,15,opt,name=any_pubkey,json=anyPubkey,proto3,oneof" json:"any_pubkey,omitempty"`
}

func (*PublicKey_Secp256K1) isPublicKey_Sum() {}
func (*PublicKey_Ed25519) isPublicKey_Sum()   {}
func (*PublicKey_Sr25519) isPublicKey_Sum()   {}
func (*PublicKey_Multisig) isPublicKey_Sum()  {}
func (*PublicKey_Secp256R1) isPublicKey_Sum() {}
func (*PublicKey_AnyPubkey) isPublicKey_Sum() {}

func (m *PublicKey) GetSum() isPublicKey_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *PublicKey) GetSecp256K1() []byte {
	if x, ok := m.GetSum().(*PublicKey_Secp256K1); ok {
		return x.Secp256K1
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetSum().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetSr25519() []byte {
	if x, ok := m.GetSum().(*PublicKey_Sr25519); ok {
		return x.Sr25519
	}
	return nil
}

func (m *PublicKey) GetMultisig() *PubKeyMultisigThreshold {
	if x, ok := m.GetSum().(*PublicKey_Multisig); ok {
		return x.Multisig
	}
	return nil
}

func (m *PublicKey) GetSecp256R1() []byte {
	if x, ok := m.GetSum().(*PublicKey_Secp256R1); ok {
		return x.Secp256R1
	}
	return nil
}

func (m *PublicKey) GetAnyPubkey() *types.Any {
	if x, ok := m.GetSum().(*PublicKey_AnyPubkey); ok {
		return x.AnyPubkey
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PublicKey) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PublicKey_Secp256K1)(nil),
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_Sr25519)(nil),
		(*PublicKey_Multisig)(nil),
		(*PublicKey_Secp256R1)(nil),
		(*PublicKey_AnyPubkey)(nil),
	}
}

// PubKeyMultisigThreshold specifies a public key type which nests multiple public
// keys and a threshold
type PubKeyMultisigThreshold struct {
	K       uint32       `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty" yaml:"threshold"`
	PubKeys []*PublicKey `protobuf:"bytes,2,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty" yaml:"pubkeys"`
}

func (m *PubKeyMultisigThreshold) Reset()         { *m = PubKeyMultisigThreshold{} }
func (m *PubKeyMultisigThreshold) String() string { return proto.CompactTextString(m) }
func (*PubKeyMultisigThreshold) ProtoMessage()    {}
func (*PubKeyMultisigThreshold) Descriptor() ([]byte, []int) {
	return fileDescriptor_2165b2a1badb1b0c, []int{1}
}
func (m *PubKeyMultisigThreshold) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKeyMultisigThreshold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKeyMultisigThreshold.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKeyMultisigThreshold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKeyMultisigThreshold.Merge(m, src)
}
func (m *PubKeyMultisigThreshold) XXX_Size() int {
	return m.Size()
}
func (m *PubKeyMultisigThreshold) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKeyMultisigThreshold.DiscardUnknown(m)
}

var xxx_messageInfo_PubKeyMultisigThreshold proto.InternalMessageInfo

func (m *PubKeyMultisigThreshold) GetK() uint32 {
	if m != nil {
		return m.K
	}
	return 0
}

func (m *PubKeyMultisigThreshold) GetPubKeys() []*PublicKey {
	if m != nil {
		return m.PubKeys
	}
	return nil
}

// MultiSignature wraps the signatures from a PubKeyMultisigThreshold.
// See cosmos_sdk.tx.v1.ModeInfo.Multi for how to specify which signers signed
// and with which modes
type MultiSignature struct {
	Signatures       [][]byte `protobuf:"bytes,1,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MultiSignature) Reset()         { *m = MultiSignature{} }
func (m *MultiSignature) String() string { return proto.CompactTextString(m) }
func (*MultiSignature) ProtoMessage()    {}
func (*MultiSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_2165b2a1badb1b0c, []int{2}
}
func (m *MultiSignature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiSignature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiSignature.Merge(m, src)
}
func (m *MultiSignature) XXX_Size() int {
	return m.Size()
}
func (m *MultiSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiSignature.DiscardUnknown(m)
}

var xxx_messageInfo_MultiSignature proto.InternalMessageInfo

func (m *MultiSignature) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// CompactBitArray is an implementation of a space efficient bit array.
// This is used to ensure that the encoded data takes up a minimal amount of
// space after proto encoding.
// This is not thread safe, and is not intended for concurrent usage.
type CompactBitArray struct {
	ExtraBitsStored uint32 `protobuf:"varint,1,opt,name=extra_bits_stored,json=extraBitsStored,proto3" json:"extra_bits_stored,omitempty"`
	Elems           []byte `protobuf:"bytes,2,opt,name=elems,proto3" json:"elems,omitempty"`
}

func (m *CompactBitArray) Reset()      { *m = CompactBitArray{} }
func (*CompactBitArray) ProtoMessage() {}
func (*CompactBitArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_2165b2a1badb1b0c, []int{3}
}
func (m *CompactBitArray) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompactBitArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompactBitArray.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompactBitArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactBitArray.Merge(m, src)
}
func (m *CompactBitArray) XXX_Size() int {
	return xxx_messageInfo_CompactBitArray.Size(m)
}
func (m *CompactBitArray) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactBitArray.DiscardUnknown(m)
}

var xxx_messageInfo_CompactBitArray proto.InternalMessageInfo

func (m *CompactBitArray) GetExtraBitsStored() uint32 {
	if m != nil {
		return m.ExtraBitsStored
	}
	return 0
}

func (m *CompactBitArray) GetElems() []byte {
	if m != nil {
		return m.Elems
	}
	return nil
}

func init() {
	proto.RegisterType((*PublicKey)(nil), "cosmos_sdk.crypto.v1.PublicKey")
	proto.RegisterType((*PubKeyMultisigThreshold)(nil), "cosmos_sdk.crypto.v1.PubKeyMultisigThreshold")
	proto.RegisterType((*MultiSignature)(nil), "cosmos_sdk.crypto.v1.MultiSignature")
	proto.RegisterType((*CompactBitArray)(nil), "cosmos_sdk.crypto.v1.CompactBitArray")
}

func init() { proto.RegisterFile("crypto/types/types.proto", fileDescriptor_2165b2a1badb1b0c) }

var fileDescriptor_2165b2a1badb1b0c = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xc1, 0x6a, 0xdb, 0x4a,
	0x14, 0x95, 0xec, 0xf8, 0x25, 0x1e, 0xe7, 0xc5, 0xed, 0x60, 0xa8, 0x62, 0xa8, 0x64, 0xb4, 0x28,
	0x6e, 0x21, 0x12, 0x76, 0x71, 0x4a, 0xbd, 0x8b, 0xb2, 0x09, 0x98, 0x82, 0x51, 0xba, 0x2a, 0x14,
	0x21, 0xc9, 0x53, 0x59, 0x58, 0xf2, 0x88, 0xb9, 0xa3, 0xd2, 0xf9, 0x89, 0xd2, 0x65, 0x97, 0xc9,
	0x37, 0xf4, 0x27, 0xba, 0xf4, 0xb2, 0x2b, 0x53, 0xec, 0x3f, 0xc8, 0x17, 0x94, 0x68, 0xa4, 0xd8,
	0x94, 0x66, 0x23, 0xe9, 0x9e, 0x73, 0x46, 0xf7, 0xdc, 0x73, 0x25, 0xa4, 0x85, 0x4c, 0x64, 0x9c,
	0xda, 0x5c, 0x64, 0x04, 0xe4, 0xd5, 0xca, 0x18, 0xe5, 0x14, 0x77, 0x42, 0x0a, 0x29, 0x05, 0x0f,
	0x66, 0x0b, 0x4b, 0x8a, 0xac, 0xcf, 0x83, 0xee, 0x0b, 0x3e, 0x8f, 0xd9, 0xcc, 0xcb, 0x7c, 0xc6,
	0x85, 0x5d, 0x08, 0xed, 0x88, 0x46, 0x74, 0xf7, 0x24, 0x4f, 0x77, 0x4f, 0x23, 0x4a, 0xa3, 0x84,
	0x48, 0x49, 0x90, 0x7f, 0xb2, 0xfd, 0xa5, 0x90, 0x94, 0xf9, 0xb5, 0x86, 0x9a, 0xd3, 0x3c, 0x48,
	0xe2, 0x70, 0x42, 0x04, 0xd6, 0x51, 0x13, 0x48, 0x98, 0x0d, 0x47, 0xe7, 0x8b, 0x81, 0xa6, 0xf6,
	0xd4, 0xfe, 0xf1, 0x95, 0xe2, 0xee, 0x20, 0xdc, 0x45, 0x87, 0x64, 0x36, 0x1c, 0x8d, 0x06, 0x6f,
	0xb5, 0x5a, 0xc9, 0x56, 0xc0, 0x3d, 0x07, 0x4c, 0x72, 0xf5, 0x8a, 0x2b, 0x01, 0x3c, 0x41, 0x47,
	0x69, 0x9e, 0xf0, 0x18, 0xe2, 0x48, 0x3b, 0xe8, 0xa9, 0xfd, 0xd6, 0xf0, 0xcc, 0xfa, 0xd7, 0x44,
	0xd6, 0x34, 0x0f, 0x26, 0x44, 0xbc, 0x2b, 0xb5, 0xef, 0xe7, 0x8c, 0xc0, 0x9c, 0x26, 0xb3, 0x2b,
	0xc5, 0x7d, 0x78, 0xc1, 0x9e, 0x49, 0x36, 0xd0, 0x1a, 0x7f, 0x99, 0x64, 0x03, 0x3c, 0x42, 0xc8,
	0x5f, 0x0a, 0x2f, 0xcb, 0x83, 0x05, 0x11, 0x5a, 0xbb, 0x68, 0xd7, 0xb1, 0x64, 0x04, 0x56, 0x15,
	0x81, 0x75, 0xb1, 0x14, 0xf7, 0xc7, 0xfc, 0xa5, 0x98, 0x16, 0x42, 0xa7, 0x81, 0xea, 0x90, 0xa7,
	0xe6, 0x0f, 0x15, 0x3d, 0x7b, 0xc4, 0x05, 0x7e, 0x83, 0x9a, 0xbc, 0x2a, 0x8a, 0x78, 0xfe, 0x77,
	0x4e, 0x37, 0x6b, 0x43, 0x9d, 0xdc, 0xad, 0x8d, 0x27, 0xc2, 0x4f, 0x93, 0xb1, 0xf9, 0xc0, 0x9b,
	0xee, 0x4e, 0x8b, 0x3d, 0xd4, 0xca, 0x8a, 0x90, 0xbd, 0x05, 0x11, 0xa0, 0xd5, 0x7a, 0xf5, 0x7e,
	0x6b, 0x68, 0x3c, 0x1a, 0x81, 0xdc, 0x86, 0xf3, 0x7c, 0xb3, 0x36, 0x0e, 0xa5, 0x17, 0xb8, 0x5b,
	0x1b, 0x27, 0xb2, 0x83, 0x9c, 0x0b, 0x4c, 0x17, 0x65, 0x95, 0x12, 0xcc, 0x73, 0x74, 0x52, 0xd8,
	0xbd, 0x8e, 0xa3, 0xa5, 0xcf, 0x73, 0x46, 0xb0, 0x8e, 0x10, 0x54, 0x05, 0x68, 0x6a, 0xaf, 0xde,
	0x3f, 0x76, 0xf7, 0x90, 0xf1, 0xc1, 0xea, 0xd6, 0x50, 0xcd, 0x8f, 0xa8, 0x7d, 0x49, 0xd3, 0xcc,
	0x0f, 0xb9, 0x13, 0xf3, 0x0b, 0xc6, 0x7c, 0x81, 0x5f, 0xa1, 0xa7, 0xe4, 0x0b, 0x67, 0xbe, 0x17,
	0xc4, 0x1c, 0x3c, 0xe0, 0x94, 0x91, 0x72, 0x58, 0xb7, 0x5d, 0x10, 0x4e, 0xcc, 0xe1, 0xba, 0x80,
	0x71, 0x07, 0x35, 0x48, 0x42, 0x52, 0x90, 0x5f, 0x83, 0x2b, 0x8b, 0xf1, 0xd1, 0xf7, 0x1b, 0x43,
	0xb9, 0xb9, 0x35, 0x14, 0xe7, 0xf2, 0xe7, 0x46, 0x57, 0x57, 0x1b, 0x5d, 0xfd, 0xbd, 0xd1, 0xd5,
	0x6f, 0x5b, 0x5d, 0x59, 0x6d, 0x75, 0xe5, 0xd7, 0x56, 0x57, 0x3e, 0xbc, 0x8c, 0x62, 0x3e, 0xcf,
	0x03, 0x2b, 0xa4, 0xa9, 0x2d, 0x63, 0x28, 0x6f, 0x67, 0x30, 0x5b, 0xd8, 0xfb, 0xff, 0x41, 0xf0,
	0x5f, 0xb1, 0xb3, 0xd7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x46, 0x56, 0x25, 0x1e, 0x03,
	0x00, 0x00,
}

func (m *PublicKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *PublicKey_Secp256K1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Secp256K1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Secp256K1 != nil {
		i -= len(m.Secp256K1)
		copy(dAtA[i:], m.Secp256K1)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Secp256K1)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *PublicKey_Ed25519) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Ed25519) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Ed25519 != nil {
		i -= len(m.Ed25519)
		copy(dAtA[i:], m.Ed25519)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Ed25519)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *PublicKey_Sr25519) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Sr25519) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Sr25519 != nil {
		i -= len(m.Sr25519)
		copy(dAtA[i:], m.Sr25519)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Sr25519)))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *PublicKey_Multisig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Multisig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Multisig != nil {
		{
			size, err := m.Multisig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *PublicKey_Secp256R1) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_Secp256R1) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Secp256R1 != nil {
		i -= len(m.Secp256R1)
		copy(dAtA[i:], m.Secp256R1)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Secp256R1)))
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *PublicKey_AnyPubkey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicKey_AnyPubkey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AnyPubkey != nil {
		{
			size, err := m.AnyPubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x7a
	}
	return len(dAtA) - i, nil
}
func (m *PubKeyMultisigThreshold) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKeyMultisigThreshold) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKeyMultisigThreshold) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKeys) > 0 {
		for iNdEx := len(m.PubKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PubKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.K != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.K))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MultiSignature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiSignature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiSignature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signatures[iNdEx])
			copy(dAtA[i:], m.Signatures[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Signatures[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CompactBitArray) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompactBitArray) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompactBitArray) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Elems) > 0 {
		i -= len(m.Elems)
		copy(dAtA[i:], m.Elems)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Elems)))
		i--
		dAtA[i] = 0x12
	}
	if m.ExtraBitsStored != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ExtraBitsStored))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PublicKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *PublicKey_Secp256K1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Secp256K1 != nil {
		l = len(m.Secp256K1)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *PublicKey_Ed25519) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ed25519 != nil {
		l = len(m.Ed25519)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *PublicKey_Sr25519) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sr25519 != nil {
		l = len(m.Sr25519)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *PublicKey_Multisig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Multisig != nil {
		l = m.Multisig.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *PublicKey_Secp256R1) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Secp256R1 != nil {
		l = len(m.Secp256R1)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *PublicKey_AnyPubkey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AnyPubkey != nil {
		l = m.AnyPubkey.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *PubKeyMultisigThreshold) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.K != 0 {
		n += 1 + sovTypes(uint64(m.K))
	}
	if len(m.PubKeys) > 0 {
		for _, e := range m.PubKeys {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *MultiSignature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for _, b := range m.Signatures {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PublicKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PublicKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secp256K1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Sum = &PublicKey_Secp256K1{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ed25519", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Sum = &PublicKey_Ed25519{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sr25519", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Sum = &PublicKey_Sr25519{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multisig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PubKeyMultisigThreshold{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &PublicKey_Multisig{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secp256R1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Sum = &PublicKey_Secp256R1{v}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnyPubkey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.Any{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &PublicKey_AnyPubkey{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *PubKeyMultisigThreshold) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PubKeyMultisigThreshold: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubKeyMultisigThreshold: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field K", wireType)
			}
			m.K = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.K |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKeys = append(m.PubKeys, &PublicKey{})
			if err := m.PubKeys[len(m.PubKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *MultiSignature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: MultiSignature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiSignature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, make([]byte, postIndex-iNdEx))
			copy(m.Signatures[len(m.Signatures)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CompactBitArray) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: CompactBitArray: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompactBitArray: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraBitsStored", wireType)
			}
			m.ExtraBitsStored = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExtraBitsStored |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elems", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elems = append(m.Elems[:0], dAtA[iNdEx:postIndex]...)
			if m.Elems == nil {
				m.Elems = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
