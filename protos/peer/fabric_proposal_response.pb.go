// Code generated by protoc-gen-go.
// source: peer/fabric_proposal_response.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A ProposalResponse is returned from an endorser to the proposal submitter.
// The idea is that this message contains the endorser's response to the
// request of a client to perform an action over a chaincode (or more
// generically on the ledger); the response might be success/error (conveyed in
// the Response field) together with a description of the action and a
// signature over it by that endorser.  If a sufficient number of distinct
// endorsers agree on the same action and produce signature to that effect, a
// transaction can be generated and sent for ordering.
type ProposalResponse struct {
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Timestamp is the time that the message
	// was created as  defined by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// A response message indicating whether the
	// endorsement of the action was successful
	Response *Response `protobuf:"bytes,4,opt,name=response" json:"response,omitempty"`
	// The payload of response. It is the bytes of ProposalResponsePayload
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	// The endorsement of the proposal, basically
	// the endorser's signature over the payload
	Endorsement *Endorsement `protobuf:"bytes,6,opt,name=endorsement" json:"endorsement,omitempty"`
}

func (m *ProposalResponse) Reset()                    { *m = ProposalResponse{} }
func (m *ProposalResponse) String() string            { return proto.CompactTextString(m) }
func (*ProposalResponse) ProtoMessage()               {}
func (*ProposalResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *ProposalResponse) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ProposalResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ProposalResponse) GetEndorsement() *Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

// A response with a representation similar to an HTTP response that can
// be used within another message.
type Response struct {
	// A status code that should follow the HTTP status codes.
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	// A message associated with the response code.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

// ProposalResponsePayload is the payload of a proposal response.  This message
// is the "bridge" between the client's request and the endorser's action in
// response to that request. Concretely, for chaincodes, it contains a hashed
// representation of the proposal (proposalHash) and a representation of the
// chaincode state changes and events inside the extension field.
type ProposalResponsePayload struct {
	// Hash of the proposal that triggered this response. The hash is used to
	// link a response with its proposal, both for bookeeping purposes on an
	// asynchronous system and for security reasons (accountability,
	// non-repudiation). The hash usually covers the entire Proposal message
	// (byte-by-byte). However this implies that the hash can only be verified
	// if the entire proposal message is available when ProposalResponsePayload is
	// included in a transaction or stored in the ledger. For confidentiality
	// reasons, with chaincodes it might be undesirable to store the proposal
	// payload in the ledger.  If the type is CHAINCODE, this is handled by
	// separating the proposal's header and
	// the payload: the header is always hashed in its entirety whereas the
	// payload can either be hashed fully, or only its hash may be hashed, or
	// nothing from the payload can be hashed. The PayloadVisibility field in the
	// Header's extension controls to which extent the proposal payload is
	// "visible" in the sense that was just explained.
	ProposalHash []byte `protobuf:"bytes,1,opt,name=proposalHash,proto3" json:"proposalHash,omitempty"`
	// Extension should be unmarshaled to a type-specific message. The type of
	// the extension in any proposal response depends on the type of the proposal
	// that the client selected when the proposal was initially sent out.  In
	// particular, this information is stored in the type field of a Header.  For
	// chaincode, it's a ChaincodeAction message
	Extension []byte `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *ProposalResponsePayload) Reset()                    { *m = ProposalResponsePayload{} }
func (m *ProposalResponsePayload) String() string            { return proto.CompactTextString(m) }
func (*ProposalResponsePayload) ProtoMessage()               {}
func (*ProposalResponsePayload) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

// An endorsement is a signature of an endorser over a proposal response.  By
// producing an endorsement message, an endorser implicitly "approves" that
// proposal response and the actions contained therein. When enough
// endorsements have been collected, a transaction can be generated out of a
// set of proposal responses.  Note that this message only contains an identity
// and a signature but no signed payload. This is intentional because
// endorsements are supposed to be collected in a transaction, and they are all
// expected to endorse a single proposal response/action (many endorsements
// over a single proposal response)
type Endorsement struct {
	// Identity of the endorser (e.g. its certificate)
	Endorser []byte `protobuf:"bytes,1,opt,name=endorser,proto3" json:"endorser,omitempty"`
	// Signature of the payload included in ProposalResponse concatenated with
	// the endorser's certificate; ie, sign(ProposalResponse.payload + endorser)
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Endorsement) Reset()                    { *m = Endorsement{} }
func (m *Endorsement) String() string            { return proto.CompactTextString(m) }
func (*Endorsement) ProtoMessage()               {}
func (*Endorsement) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func init() {
	proto.RegisterType((*ProposalResponse)(nil), "protos.ProposalResponse")
	proto.RegisterType((*Response)(nil), "protos.Response")
	proto.RegisterType((*ProposalResponsePayload)(nil), "protos.ProposalResponsePayload")
	proto.RegisterType((*Endorsement)(nil), "protos.Endorsement")
}

func init() { proto.RegisterFile("peer/fabric_proposal_response.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x52, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0xc5, 0xd9, 0x4d, 0x36, 0x99, 0xe4, 0x10, 0xb4, 0xb0, 0x6b, 0x42, 0xa1, 0xc1, 0xbd, 0xa4,
	0xb4, 0xd8, 0xd0, 0x52, 0xe8, 0xb9, 0x50, 0xda, 0x63, 0x10, 0xa5, 0x87, 0xf6, 0x10, 0xe4, 0x64,
	0xe2, 0x18, 0x6c, 0x4b, 0x68, 0xe4, 0xd2, 0xfc, 0xe0, 0xfe, 0x8f, 0x62, 0x59, 0x72, 0x92, 0x9e,
	0xcc, 0x1b, 0x3f, 0xbd, 0x0f, 0x69, 0xe0, 0x42, 0x21, 0xea, 0x64, 0x2b, 0x52, 0x9d, 0xaf, 0x57,
	0x4a, 0x4b, 0x25, 0x49, 0x14, 0x2b, 0x8d, 0xa4, 0x64, 0x45, 0x18, 0x2b, 0x2d, 0x8d, 0x64, 0x03,
	0xfb, 0xa1, 0xd9, 0x79, 0x26, 0x65, 0x56, 0x60, 0x62, 0x61, 0x5a, 0x6f, 0x13, 0x93, 0x97, 0x48,
	0x46, 0x94, 0xaa, 0x25, 0x46, 0x5f, 0x01, 0x4c, 0x97, 0x4e, 0x84, 0x3b, 0x0d, 0x16, 0xc2, 0x9f,
	0x0f, 0xd4, 0x94, 0xcb, 0x2a, 0x0c, 0xe6, 0xc1, 0xa2, 0xcf, 0x3d, 0x64, 0xf7, 0x30, 0xea, 0x14,
	0xc2, 0xde, 0x3c, 0x58, 0x8c, 0x6f, 0x66, 0x71, 0xeb, 0x11, 0x7b, 0x8f, 0xf8, 0xc5, 0x33, 0xf8,
	0x81, 0xcc, 0xae, 0x61, 0xe8, 0x33, 0x86, 0xbf, 0xed, 0xc1, 0x69, 0x7b, 0x82, 0x62, 0xef, 0xcb,
	0x3b, 0x46, 0x93, 0x40, 0x89, 0x7d, 0x21, 0xc5, 0x26, 0xec, 0xcf, 0x83, 0xc5, 0x84, 0x7b, 0xc8,
	0xee, 0x60, 0x8c, 0xd5, 0x46, 0x6a, 0xc2, 0x12, 0x2b, 0x13, 0x0e, 0xac, 0xd4, 0x5f, 0x2f, 0xf5,
	0x78, 0xf8, 0xc5, 0x8f, 0x79, 0xd1, 0x2b, 0x0c, 0xbb, 0x7a, 0xff, 0x60, 0x40, 0x46, 0x98, 0x9a,
	0x5c, 0x3b, 0x87, 0x1a, 0xd3, 0x12, 0x89, 0x44, 0x86, 0xb6, 0xda, 0x88, 0x7b, 0x78, 0x1c, 0xe7,
	0xd7, 0x49, 0x9c, 0xe8, 0x1d, 0xfe, 0xff, 0xbc, 0xbe, 0xa5, 0x4b, 0x1a, 0xc1, 0xc4, 0x3f, 0xcf,
	0xb3, 0xa0, 0x9d, 0x35, 0x9b, 0xf0, 0x93, 0x19, 0x3b, 0x83, 0x11, 0x7e, 0x1a, 0xac, 0xec, 0x5d,
	0xf7, 0x2c, 0xe1, 0x30, 0x88, 0x9e, 0x60, 0x7c, 0x54, 0x88, 0xcd, 0x60, 0xe8, 0x2a, 0x69, 0x27,
	0xd6, 0xe1, 0x46, 0x88, 0xf2, 0xac, 0x12, 0xa6, 0xd6, 0xe8, 0x85, 0xba, 0xc1, 0xc3, 0xd5, 0xdb,
	0x65, 0x96, 0x9b, 0x5d, 0x9d, 0xc6, 0x6b, 0x59, 0x26, 0xbb, 0xbd, 0x42, 0x5d, 0xe0, 0x26, 0xeb,
	0xf6, 0xa8, 0xdd, 0x0f, 0x4a, 0x9a, 0xd5, 0x4a, 0xdb, 0xdd, 0xb9, 0xfd, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x43, 0xe1, 0xd2, 0xab, 0x69, 0x02, 0x00, 0x00,
}
