// Code generated by protoc-gen-go.
// source: proto/security_e2ekeys_v1/e2ekeys.proto
// DO NOT EDIT!

/*
Package security_e2ekeys_v1 is a generated protocol buffer package.

It is generated from these files:
	proto/security_e2ekeys_v1/e2ekeys.proto

It has these top-level messages:
	HkpLookupRequest
	HttpResponse
	GetEntryResponse
	Committed
	Profile
	Entry
	PublicKey
	KeyValue
	SignedKV
	GetEntryRequest
	ListEntryHistoryRequest
	ListEntryHistoryResponse
	EntryUpdate
	UpdateEntryRequest
	UpdateEntryResponse
*/
package security_e2ekeys_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import security_ctmap "github.com/google/e2e-key-server/proto/security_ctmap"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HkpLookupRequest contains query parameters for retrieving PGP keys.
type HkpLookupRequest struct {
	// Op specifies the operation to be performed on the keyserver.
	// - "get" returns the pgp key specified in the search parameter.
	// - "index" returns 501 (not implemented).
	// - "vindex" returns 501 (not implemented).
	Op string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	// Search specifies the email address or key id being queried.
	Search string `protobuf:"bytes,2,opt,name=search" json:"search,omitempty"`
	// Options specifies what output format to use.
	// - "mr" machine readable will set the content type to "application/pgp-keys"
	// - other options will be ignored.
	Options string `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	// Exact specifies an exact match on search. Always on. If specified in the
	// URL, its value will be ignored.
	Exact string `protobuf:"bytes,4,opt,name=exact" json:"exact,omitempty"`
	// fingerprint is ignored.
	Fingerprint string `protobuf:"bytes,5,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *HkpLookupRequest) Reset()                    { *m = HkpLookupRequest{} }
func (m *HkpLookupRequest) String() string            { return proto.CompactTextString(m) }
func (*HkpLookupRequest) ProtoMessage()               {}
func (*HkpLookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// HttpBody represents an http body.
type HttpResponse struct {
	// Header content type.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	// The http body itself.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HttpResponse) Reset()                    { *m = HttpResponse{} }
func (m *HttpResponse) String() string            { return proto.CompactTextString(m) }
func (*HttpResponse) ProtoMessage()               {}
func (*HttpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// GetEntryResponse
type GetEntryResponse struct {
	// vrf is the output of VRF on user_id.
	Vrf []byte `protobuf:"bytes,1,opt,name=vrf,proto3" json:"vrf,omitempty"`
	// vrf_proof is the proof for VRF on user_id.
	VrfProof []byte `protobuf:"bytes,2,opt,name=vrf_proof,json=vrfProof,proto3" json:"vrf_proof,omitempty"`
	// committed contains the profile for this account and connects the data
	// in profile to the commitment in leaf_proof.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
	// leaf_proof contains an Entry and an inclusion proof in the sparse merkle tree at end_epoch.
	LeafProof *security_ctmap.GetLeafResponse `protobuf:"bytes,5,opt,name=leaf_proof,json=leafProof" json:"leaf_proof,omitempty"`
	// seh contains the signed epoch head for the sparse merkle tree.
	// seh is also stored in the append only log.
	Seh *security_ctmap.SignedEpochHead `protobuf:"bytes,6,opt,name=seh" json:"seh,omitempty"`
	// seh_sct is the SCT showing that seh was submitted to CT logs.
	// TODO: Support storing seh in multiple logs.
	SehSct []byte `protobuf:"bytes,7,opt,name=seh_sct,json=sehSct,proto3" json:"seh_sct,omitempty"`
}

func (m *GetEntryResponse) Reset()                    { *m = GetEntryResponse{} }
func (m *GetEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryResponse) ProtoMessage()               {}
func (*GetEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetEntryResponse) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

func (m *GetEntryResponse) GetLeafProof() *security_ctmap.GetLeafResponse {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

func (m *GetEntryResponse) GetSeh() *security_ctmap.SignedEpochHead {
	if m != nil {
		return m.Seh
	}
	return nil
}

// Committed represents the data comitted to in a cryptographic commitment.
// commitment = HMAC_SHA512_256(key, data)
type Committed struct {
	// key is the 16 byte random commitment key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// data is the data being comitted to.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Committed) Reset()                    { *m = Committed{} }
func (m *Committed) String() string            { return proto.CompactTextString(m) }
func (*Committed) ProtoMessage()               {}
func (*Committed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Profile contains data hidden behind the crypto comitment.
type Profile struct {
	// Keys is a map of appIds to keys.
	Keys map[string][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Profile) GetKeys() map[string][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

// Entry contains a commitment to profile and a set of authorized update keys.
// Entry is placed in the verifiable map as leaf data.
type Entry struct {
	// commitment is a cryptographic commitment to arbitrary data.
	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	// authorized_keys is the set of keys allowed to sign updates for this entry.
	AuthorizedKeys []*PublicKey `protobuf:"bytes,2,rep,name=authorized_keys,json=authorizedKeys" json:"authorized_keys,omitempty"`
	// update_count prevents replay attacks. Monotonically increasing.
	UpdateCount uint64 `protobuf:"varint,3,opt,name=update_count,json=updateCount" json:"update_count,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Entry) GetAuthorizedKeys() []*PublicKey {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

// PublicKey defines a key this domain uses to sign EpochHeads with.
type PublicKey struct {
	// KeyFormats from Keyczar.
	//
	// Types that are valid to be assigned to KeyType:
	//	*PublicKey_Ed25519
	//	*PublicKey_RsaVerifyingSha256_2048
	//	*PublicKey_EcdsaVerifyingP256
	KeyType isPublicKey_KeyType `protobuf_oneof:"key_type"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isPublicKey_KeyType interface {
	isPublicKey_KeyType()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}
type PublicKey_RsaVerifyingSha256_2048 struct {
	RsaVerifyingSha256_2048 []byte `protobuf:"bytes,2,opt,name=rsa_verifying_sha256_2048,json=rsaVerifyingSha2562048,proto3,oneof"`
}
type PublicKey_EcdsaVerifyingP256 struct {
	EcdsaVerifyingP256 []byte `protobuf:"bytes,3,opt,name=ecdsa_verifying_p256,json=ecdsaVerifyingP256,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_KeyType()                 {}
func (*PublicKey_RsaVerifyingSha256_2048) isPublicKey_KeyType() {}
func (*PublicKey_EcdsaVerifyingP256) isPublicKey_KeyType()      {}

func (m *PublicKey) GetKeyType() isPublicKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetRsaVerifyingSha256_2048() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_RsaVerifyingSha256_2048); ok {
		return x.RsaVerifyingSha256_2048
	}
	return nil
}

func (m *PublicKey) GetEcdsaVerifyingP256() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_EcdsaVerifyingP256); ok {
		return x.EcdsaVerifyingP256
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublicKey) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublicKey_OneofMarshaler, _PublicKey_OneofUnmarshaler, _PublicKey_OneofSizer, []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_RsaVerifyingSha256_2048)(nil),
		(*PublicKey_EcdsaVerifyingP256)(nil),
	}
}

func _PublicKey_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EcdsaVerifyingP256)
	case nil:
	default:
		return fmt.Errorf("PublicKey.KeyType has unexpected type %T", x)
	}
	return nil
}

func _PublicKey_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublicKey)
	switch tag {
	case 1: // key_type.ed25519
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_Ed25519{x}
		return true, err
	case 2: // key_type.rsa_verifying_sha256_2048
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_RsaVerifyingSha256_2048{x}
		return true, err
	case 3: // key_type.ecdsa_verifying_p256
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_EcdsaVerifyingP256{x}
		return true, err
	default:
		return false, nil
	}
}

func _PublicKey_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ed25519)))
		n += len(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsaVerifyingSha256_2048)))
		n += len(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EcdsaVerifyingP256)))
		n += len(x.EcdsaVerifyingP256)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// KeyValue is a map entry.
type KeyValue struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// SignedKV is a signed change to a map entry.
type SignedKV struct {
	// keyvalue is a serialized KeyValue.
	KeyValue []byte `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	// signatures on keyvalue. Must be signed by keys from both previous and
	// current epochs. The first proves ownership of new epoch key, and the
	// second proves the the correct owner is making this change.
	Signatures map[uint64][]byte `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"fixed64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *SignedKV) Reset()                    { *m = SignedKV{} }
func (m *SignedKV) String() string            { return proto.CompactTextString(m) }
func (*SignedKV) ProtoMessage()               {}
func (*SignedKV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SignedKV) GetSignatures() map[uint64][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// Get request for a user object.
type GetEntryRequest struct {
	// Last trusted epoch by the client.
	// int64 epoch_start = 3;
	// Absence of the epoch_end field indicates a request for the current value.
	EpochEnd int64 `protobuf:"varint,1,opt,name=epoch_end,json=epochEnd" json:"epoch_end,omitempty"`
	// User identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetEntryRequest) Reset()                    { *m = GetEntryRequest{} }
func (m *GetEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryRequest) ProtoMessage()               {}
func (*GetEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// Get a list of historical values for a user.
type ListEntryHistoryRequest struct {
	// The user identifier.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// from_epoch is the starting epcoh.
	StartEpoch int64 `protobuf:"varint,2,opt,name=start_epoch,json=startEpoch" json:"start_epoch,omitempty"`
	// The maximum number of entries to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListEntryHistoryRequest) Reset()                    { *m = ListEntryHistoryRequest{} }
func (m *ListEntryHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryRequest) ProtoMessage()               {}
func (*ListEntryHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// A paginated history of values for a user.
type ListEntryHistoryResponse struct {
	// The list of values this user_id has contained over time.
	Values []*GetEntryResponse `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// The next time to query for pagination.
	NextEpoch int64 `protobuf:"varint,2,opt,name=next_epoch,json=nextEpoch" json:"next_epoch,omitempty"`
}

func (m *ListEntryHistoryResponse) Reset()                    { *m = ListEntryHistoryResponse{} }
func (m *ListEntryHistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryResponse) ProtoMessage()               {}
func (*ListEntryHistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListEntryHistoryResponse) GetValues() []*GetEntryResponse {
	if m != nil {
		return m.Values
	}
	return nil
}

type EntryUpdate struct {
	// update authorizes the change to profile.
	Update *SignedKV `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	// commitment contains the serialized Profile protobuf.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
}

func (m *EntryUpdate) Reset()                    { *m = EntryUpdate{} }
func (m *EntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*EntryUpdate) ProtoMessage()               {}
func (*EntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *EntryUpdate) GetUpdate() *SignedKV {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *EntryUpdate) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

// Update a user's profile.
type UpdateEntryRequest struct {
	// user_id specifies the id for the new account to be registered.
	UserId      string       `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	EntryUpdate *EntryUpdate `protobuf:"bytes,2,opt,name=entry_update,json=entryUpdate" json:"entry_update,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UpdateEntryRequest) GetEntryUpdate() *EntryUpdate {
	if m != nil {
		return m.EntryUpdate
	}
	return nil
}

// UpdateEntryResponse contains a proof once the update has been included in
// the Merkel Tree.
type UpdateEntryResponse struct {
	Proof *GetEntryResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateEntryResponse) Reset()                    { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()               {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *UpdateEntryResponse) GetProof() *GetEntryResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*HkpLookupRequest)(nil), "security.e2ekeys.v1.HkpLookupRequest")
	proto.RegisterType((*HttpResponse)(nil), "security.e2ekeys.v1.HttpResponse")
	proto.RegisterType((*GetEntryResponse)(nil), "security.e2ekeys.v1.GetEntryResponse")
	proto.RegisterType((*Committed)(nil), "security.e2ekeys.v1.Committed")
	proto.RegisterType((*Profile)(nil), "security.e2ekeys.v1.Profile")
	proto.RegisterType((*Entry)(nil), "security.e2ekeys.v1.Entry")
	proto.RegisterType((*PublicKey)(nil), "security.e2ekeys.v1.PublicKey")
	proto.RegisterType((*KeyValue)(nil), "security.e2ekeys.v1.KeyValue")
	proto.RegisterType((*SignedKV)(nil), "security.e2ekeys.v1.SignedKV")
	proto.RegisterType((*GetEntryRequest)(nil), "security.e2ekeys.v1.GetEntryRequest")
	proto.RegisterType((*ListEntryHistoryRequest)(nil), "security.e2ekeys.v1.ListEntryHistoryRequest")
	proto.RegisterType((*ListEntryHistoryResponse)(nil), "security.e2ekeys.v1.ListEntryHistoryResponse")
	proto.RegisterType((*EntryUpdate)(nil), "security.e2ekeys.v1.EntryUpdate")
	proto.RegisterType((*UpdateEntryRequest)(nil), "security.e2ekeys.v1.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "security.e2ekeys.v1.UpdateEntryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for E2EKeyService service

type E2EKeyServiceClient interface {
	// GetEntry returns a user's entry in the Merkle Tree. Entries contain
	// signed commitments to a profile, which is also returned.
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error)
	// ListEntryHistory returns a list of GetEntryRespons covering several epochs.
	ListEntryHistory(ctx context.Context, in *ListEntryHistoryRequest, opts ...grpc.CallOption) (*ListEntryHistoryResponse, error)
	// blocking or polling?
	// UpdateEntry submits a SignedEntryUpdate.  Returns empty until this update
	// has been included in an epoch.  Clients must retry until this function
	// returns a proof.
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error)
	//
	// // List the Signed Epoch Heads, from epoch to epoch.
	// rpc ListSEH(ListSEHRequest) returns (ListSEHResponse);
	//
	// // List the EntryUpdates by update number.
	// rpc ListUpdate(ListUpdateRequest) returns (ListUpdateResponse);
	//
	// // ListSteps combines SEH and EntryUpdates into single list.
	// rpc ListSteps(ListStepsRequest) returns (ListStepsResponse);
	HkpLookup(ctx context.Context, in *HkpLookupRequest, opts ...grpc.CallOption) (*HttpResponse, error)
}

type e2EKeyServiceClient struct {
	cc *grpc.ClientConn
}

func NewE2EKeyServiceClient(cc *grpc.ClientConn) E2EKeyServiceClient {
	return &e2EKeyServiceClient{cc}
}

func (c *e2EKeyServiceClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error) {
	out := new(GetEntryResponse)
	err := grpc.Invoke(ctx, "/security.e2ekeys.v1.E2EKeyService/GetEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyServiceClient) ListEntryHistory(ctx context.Context, in *ListEntryHistoryRequest, opts ...grpc.CallOption) (*ListEntryHistoryResponse, error) {
	out := new(ListEntryHistoryResponse)
	err := grpc.Invoke(ctx, "/security.e2ekeys.v1.E2EKeyService/ListEntryHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyServiceClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error) {
	out := new(UpdateEntryResponse)
	err := grpc.Invoke(ctx, "/security.e2ekeys.v1.E2EKeyService/UpdateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyServiceClient) HkpLookup(ctx context.Context, in *HkpLookupRequest, opts ...grpc.CallOption) (*HttpResponse, error) {
	out := new(HttpResponse)
	err := grpc.Invoke(ctx, "/security.e2ekeys.v1.E2EKeyService/HkpLookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for E2EKeyService service

type E2EKeyServiceServer interface {
	// GetEntry returns a user's entry in the Merkle Tree. Entries contain
	// signed commitments to a profile, which is also returned.
	GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error)
	// ListEntryHistory returns a list of GetEntryRespons covering several epochs.
	ListEntryHistory(context.Context, *ListEntryHistoryRequest) (*ListEntryHistoryResponse, error)
	// blocking or polling?
	// UpdateEntry submits a SignedEntryUpdate.  Returns empty until this update
	// has been included in an epoch.  Clients must retry until this function
	// returns a proof.
	UpdateEntry(context.Context, *UpdateEntryRequest) (*UpdateEntryResponse, error)
	//
	// // List the Signed Epoch Heads, from epoch to epoch.
	// rpc ListSEH(ListSEHRequest) returns (ListSEHResponse);
	//
	// // List the EntryUpdates by update number.
	// rpc ListUpdate(ListUpdateRequest) returns (ListUpdateResponse);
	//
	// // ListSteps combines SEH and EntryUpdates into single list.
	// rpc ListSteps(ListStepsRequest) returns (ListStepsResponse);
	HkpLookup(context.Context, *HkpLookupRequest) (*HttpResponse, error)
}

func RegisterE2EKeyServiceServer(s *grpc.Server, srv E2EKeyServiceServer) {
	s.RegisterService(&_E2EKeyService_serviceDesc, srv)
}

func _E2EKeyService_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(E2EKeyServiceServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.e2ekeys.v1.E2EKeyService/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(E2EKeyServiceServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _E2EKeyService_ListEntryHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntryHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(E2EKeyServiceServer).ListEntryHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.e2ekeys.v1.E2EKeyService/ListEntryHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(E2EKeyServiceServer).ListEntryHistory(ctx, req.(*ListEntryHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _E2EKeyService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(E2EKeyServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.e2ekeys.v1.E2EKeyService/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(E2EKeyServiceServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _E2EKeyService_HkpLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HkpLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(E2EKeyServiceServer).HkpLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.e2ekeys.v1.E2EKeyService/HkpLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(E2EKeyServiceServer).HkpLookup(ctx, req.(*HkpLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _E2EKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "security.e2ekeys.v1.E2EKeyService",
	HandlerType: (*E2EKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _E2EKeyService_GetEntry_Handler,
		},
		{
			MethodName: "ListEntryHistory",
			Handler:    _E2EKeyService_ListEntryHistory_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _E2EKeyService_UpdateEntry_Handler,
		},
		{
			MethodName: "HkpLookup",
			Handler:    _E2EKeyService_HkpLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("proto/security_e2ekeys_v1/e2ekeys.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1068 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0x51, 0x73, 0xdb, 0x44,
	0x10, 0x46, 0x76, 0xec, 0x58, 0xeb, 0x90, 0x66, 0x2e, 0x25, 0x35, 0x0a, 0x69, 0x53, 0x15, 0x68,
	0x99, 0x21, 0x16, 0x16, 0x04, 0x4a, 0x4a, 0x79, 0x20, 0xe3, 0xa9, 0x99, 0x86, 0x99, 0x8c, 0x0c,
	0x79, 0xd5, 0xc8, 0xd2, 0xd9, 0xd6, 0xc4, 0x96, 0x84, 0x74, 0xf2, 0xc4, 0x30, 0xf0, 0xd0, 0xd7,
	0xf2, 0x56, 0x5e, 0x78, 0xe7, 0x2f, 0xf0, 0x4b, 0xf8, 0x0b, 0xfc, 0x10, 0xf6, 0xee, 0x24, 0x5b,
	0x76, 0x55, 0x92, 0x19, 0x9e, 0x7c, 0xb7, 0xf7, 0xed, 0xee, 0xb7, 0xdf, 0xee, 0x9d, 0x05, 0x0f,
	0xa3, 0x38, 0x64, 0xa1, 0x91, 0x50, 0x37, 0x8d, 0x7d, 0x36, 0xb7, 0xa9, 0x49, 0x2f, 0xe9, 0x3c,
	0xb1, 0x67, 0x1d, 0x23, 0x5b, 0xb6, 0x05, 0x82, 0xec, 0xe6, 0x90, 0x76, 0x6e, 0x9f, 0x75, 0xb4,
	0xb3, 0x91, 0xcf, 0xc6, 0xe9, 0xa0, 0xed, 0x86, 0x53, 0x63, 0xe4, 0x0d, 0xe8, 0x64, 0xe6, 0x07,
	0xdc, 0xef, 0x08, 0x01, 0x47, 0x09, 0x8d, 0x67, 0x34, 0x36, 0xd6, 0x12, 0xb8, 0x6c, 0xea, 0x44,
	0x6b, 0x5b, 0x99, 0x42, 0x7b, 0x6f, 0x14, 0x86, 0xa3, 0x09, 0x35, 0x9c, 0xc8, 0x37, 0x9c, 0x20,
	0x08, 0x99, 0xc3, 0xfc, 0x30, 0xc8, 0x08, 0xe8, 0xbf, 0x29, 0xb0, 0xd3, 0xbb, 0x8c, 0xce, 0xc2,
	0xf0, 0x32, 0x8d, 0x2c, 0xfa, 0x63, 0x4a, 0x13, 0x46, 0xb6, 0xa1, 0x12, 0x46, 0x2d, 0xe5, 0x50,
	0x79, 0xa4, 0x5a, 0xb8, 0x22, 0x7b, 0x50, 0x4f, 0xa8, 0x13, 0xbb, 0xe3, 0x56, 0x45, 0xd8, 0xb2,
	0x1d, 0x69, 0xc1, 0x66, 0x18, 0x89, 0x68, 0xad, 0xaa, 0x38, 0xc8, 0xb7, 0xe4, 0x36, 0xd4, 0xe8,
	0x95, 0xe3, 0xb2, 0xd6, 0x86, 0xb0, 0xcb, 0x0d, 0x39, 0x84, 0xe6, 0xd0, 0x0f, 0x46, 0x34, 0x8e,
	0x62, 0x3f, 0x60, 0xad, 0x9a, 0x38, 0x2b, 0x9a, 0xf4, 0x2e, 0x6c, 0xf5, 0x18, 0x43, 0x22, 0x49,
	0x84, 0x61, 0x28, 0xb9, 0x0f, 0x5b, 0x6e, 0x18, 0x30, 0x1a, 0x30, 0x9b, 0xcd, 0x23, 0x9a, 0x71,
	0x6a, 0x66, 0xb6, 0xef, 0xd1, 0x44, 0x08, 0x6c, 0x0c, 0x42, 0x6f, 0x2e, 0xa8, 0x6d, 0x59, 0x62,
	0xad, 0xbf, 0xac, 0xc0, 0xce, 0x33, 0xca, 0xba, 0x01, 0x8b, 0xe7, 0x8b, 0x58, 0x3b, 0x50, 0x9d,
	0xc5, 0x43, 0x11, 0x62, 0xcb, 0xe2, 0x4b, 0xb2, 0x0f, 0x2a, 0xfe, 0xd8, 0xa8, 0x44, 0x38, 0xcc,
	0xfc, 0x1b, 0x68, 0x38, 0xe7, 0x7b, 0xf2, 0x15, 0xa8, 0xd8, 0x80, 0xa9, 0xcf, 0x18, 0xf5, 0x44,
	0x79, 0x4d, 0xf3, 0x6e, 0xbb, 0xa4, 0x5d, 0xed, 0xd3, 0x1c, 0x65, 0x2d, 0x1d, 0xc8, 0xd7, 0x00,
	0x13, 0xea, 0xe4, 0xb1, 0x6b, 0xc2, 0xfd, 0x5e, 0x7b, 0xad, 0x41, 0x48, 0xf1, 0x0c, 0x41, 0x39,
	0x43, 0x4b, 0xe5, 0x2e, 0x32, 0x7b, 0x07, 0xaa, 0x09, 0x1d, 0xb7, 0xea, 0xe5, 0x8e, 0x7d, 0x7f,
	0x14, 0x50, 0xaf, 0x1b, 0x85, 0xee, 0xb8, 0x47, 0x1d, 0xcf, 0xe2, 0x58, 0x72, 0x07, 0x36, 0xf1,
	0xc7, 0x4e, 0x50, 0xf5, 0x4d, 0x51, 0x0b, 0xb6, 0x69, 0xdc, 0x77, 0x99, 0xde, 0x01, 0x75, 0xc1,
	0x91, 0xab, 0x80, 0xc4, 0x73, 0x15, 0x70, 0xc9, 0x05, 0xf4, 0x1c, 0xe6, 0xe4, 0x02, 0xf2, 0xb5,
	0xfe, 0x2b, 0x6c, 0x22, 0x8f, 0xa1, 0x3f, 0xa1, 0xe4, 0x04, 0x36, 0x78, 0xa5, 0xe8, 0x51, 0x45,
	0x2a, 0x1f, 0x96, 0x4a, 0x90, 0x61, 0xdb, 0xcf, 0x71, 0x2f, 0x45, 0x17, 0x3e, 0xda, 0x17, 0xa0,
	0x2e, 0x4c, 0xc5, 0xcc, 0xaa, 0xcc, 0x8c, 0x53, 0x32, 0x73, 0x26, 0x29, 0xcd, 0x52, 0xcb, 0xcd,
	0x49, 0xe5, 0xb1, 0xa2, 0xbf, 0x52, 0xa0, 0x26, 0xbd, 0xee, 0x02, 0x48, 0x55, 0xa7, 0xd8, 0xf0,
	0x8c, 0x76, 0xc1, 0x42, 0x9e, 0xc1, 0x2d, 0x27, 0x65, 0xe3, 0x30, 0xf6, 0x7f, 0xa2, 0x9e, 0x2d,
	0x98, 0x56, 0x04, 0xd3, 0xf2, 0x66, 0x9d, 0xa7, 0x83, 0x89, 0xef, 0x22, 0x29, 0x6b, 0x7b, 0xe9,
	0xc6, 0x39, 0xf2, 0x51, 0x4b, 0x23, 0x2c, 0x9e, 0xda, 0x6e, 0x98, 0x62, 0x2a, 0xde, 0xf2, 0x0d,
	0xab, 0x29, 0x6d, 0xa7, 0xdc, 0xa4, 0xff, 0xa9, 0x80, 0xba, 0x08, 0x40, 0x34, 0xd8, 0xa4, 0x9e,
	0x79, 0x7c, 0xdc, 0xf9, 0x52, 0xd2, 0xea, 0xbd, 0x65, 0xe5, 0x06, 0xf2, 0x04, 0xde, 0x8d, 0x13,
	0xc7, 0xc6, 0x9b, 0xea, 0x0f, 0xe7, 0x38, 0xdf, 0x76, 0x32, 0x76, 0xcc, 0xe3, 0xcf, 0x6d, 0xf3,
	0x93, 0xcf, 0x1e, 0xcb, 0x6a, 0x11, 0xbd, 0x87, 0x90, 0x8b, 0x1c, 0xd1, 0x17, 0x00, 0x7e, 0x4e,
	0x4c, 0xb8, 0x4d, 0x5d, 0x6f, 0xc5, 0x3d, 0xc2, 0x33, 0xc1, 0x88, 0xfb, 0x11, 0x71, 0xba, 0xf0,
	0x3c, 0xc7, 0xb3, 0x6f, 0x00, 0x1a, 0x58, 0xa2, 0xb8, 0x24, 0xba, 0x09, 0x0d, 0xe4, 0x77, 0xc1,
	0xc5, 0x2c, 0x69, 0x77, 0xa9, 0xe8, 0xfa, 0x5f, 0x0a, 0x34, 0xe4, 0x54, 0x3d, 0xbf, 0xe0, 0xf7,
	0x82, 0x07, 0x93, 0x30, 0xe9, 0xca, 0xa3, 0xcb, 0x88, 0xdf, 0x01, 0x24, 0x08, 0x74, 0x58, 0x1a,
	0xd3, 0x5c, 0xeb, 0xa3, 0x52, 0xad, 0xf3, 0x78, 0x62, 0x21, 0xf1, 0x72, 0x38, 0x0a, 0x01, 0xb4,
	0xa7, 0x70, 0x6b, 0xed, 0xb8, 0xc8, 0xb9, 0x7e, 0xdd, 0xa0, 0x60, 0xfb, 0x97, 0x17, 0x5d, 0xbe,
	0x5e, 0xc8, 0x9e, 0xf2, 0x9b, 0x61, 0xd3, 0xc0, 0x13, 0x41, 0xaa, 0x56, 0x43, 0x18, 0xba, 0x81,
	0xc7, 0x2f, 0x49, 0x8a, 0x4f, 0xa8, 0xed, 0x7b, 0xf9, 0x5b, 0xc6, 0xb7, 0xdf, 0x7a, 0x7a, 0x04,
	0x77, 0xce, 0xfc, 0x44, 0x46, 0xea, 0xe1, 0x22, 0x5c, 0x06, 0x2c, 0xf8, 0x28, 0x45, 0x1f, 0x72,
	0x0f, 0x9a, 0x09, 0x73, 0x62, 0x66, 0x8b, 0xf0, 0x22, 0x60, 0x15, 0x8b, 0xe3, 0x26, 0x71, 0x37,
	0x39, 0x95, 0xc8, 0x19, 0x51, 0x3b, 0xc1, 0x29, 0x13, 0xed, 0xab, 0x59, 0x0d, 0x6e, 0xe8, 0xe3,
	0x5e, 0xbf, 0x82, 0xd6, 0xeb, 0x19, 0xb3, 0xb7, 0xea, 0x29, 0xd4, 0x45, 0x8d, 0xf9, 0xb5, 0xfb,
	0xa0, 0x54, 0xe0, 0xf5, 0x27, 0xce, 0xca, 0x9c, 0xc8, 0x01, 0x40, 0x40, 0xaf, 0x56, 0x79, 0xa9,
	0xdc, 0x22, 0x68, 0xe9, 0x2f, 0x14, 0x68, 0x0a, 0xc7, 0x1f, 0xc4, 0x70, 0x93, 0x63, 0xa8, 0xcb,
	0x31, 0x17, 0xd0, 0xa6, 0x79, 0xf0, 0x9f, 0xed, 0xb4, 0x32, 0xf0, 0xff, 0x7b, 0x21, 0xf5, 0x18,
	0x88, 0x4c, 0xbf, 0xd2, 0xbc, 0x37, 0x6a, 0x7d, 0x0a, 0x5b, 0x94, 0x03, 0xed, 0x15, 0xa6, 0x87,
	0xa5, 0xf9, 0x0a, 0xb5, 0x59, 0x4d, 0xba, 0xdc, 0xe8, 0x16, 0xec, 0xae, 0xe4, 0xcc, 0xd4, 0x7e,
	0x02, 0x35, 0xf9, 0x4e, 0x2b, 0x22, 0xe8, 0x0d, 0xc5, 0x96, 0x3e, 0xe6, 0x1f, 0x1b, 0xf0, 0x76,
	0xd7, 0xec, 0xe2, 0x8d, 0xeb, 0xe3, 0x5f, 0xb3, 0xef, 0x52, 0xc2, 0xa0, 0x91, 0x83, 0xc9, 0xfb,
	0xd7, 0xc4, 0x12, 0x55, 0x6b, 0x37, 0xcb, 0xa8, 0xef, 0xbf, 0xf8, 0xfb, 0x9f, 0x57, 0x95, 0x77,
	0xc8, 0xae, 0x81, 0x1f, 0x12, 0x5c, 0x97, 0xc4, 0xf8, 0x39, 0x53, 0xeb, 0x17, 0xf2, 0x3b, 0xfe,
	0x93, 0xaf, 0xcf, 0x13, 0xf9, 0xb8, 0x34, 0xf0, 0x1b, 0x06, 0x5d, 0x3b, 0xba, 0x21, 0x3a, 0xa3,
	0xf3, 0x40, 0xd0, 0x39, 0x20, 0xfb, 0x25, 0x74, 0x8c, 0x71, 0xc6, 0xe0, 0x25, 0xce, 0x5a, 0x41,
	0x73, 0xf2, 0xb0, 0x34, 0xc7, 0xeb, 0x93, 0xa0, 0x3d, 0xba, 0x1e, 0x98, 0xf1, 0xf8, 0x48, 0xf0,
	0x78, 0xa0, 0x95, 0xc9, 0x72, 0xb2, 0x32, 0x35, 0x64, 0x0a, 0xea, 0xe2, 0x6b, 0x87, 0x94, 0xab,
	0xbe, 0xfe, 0x35, 0xa4, 0xdd, 0x2f, 0x87, 0x15, 0x3e, 0x53, 0xf4, 0x3d, 0xc1, 0x60, 0x87, 0x6c,
	0x73, 0x06, 0xe3, 0xcb, 0xc8, 0x98, 0x88, 0x08, 0x83, 0xba, 0xf8, 0xc8, 0xfa, 0xf4, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb4, 0xfc, 0x55, 0x06, 0x10, 0x0a, 0x00, 0x00,
}
