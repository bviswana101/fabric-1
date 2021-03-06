// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/chaincode_shim.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_UNDEFINED           ChaincodeMessage_Type = 0
	ChaincodeMessage_REGISTER            ChaincodeMessage_Type = 1
	ChaincodeMessage_REGISTERED          ChaincodeMessage_Type = 2
	ChaincodeMessage_INIT                ChaincodeMessage_Type = 3
	ChaincodeMessage_READY               ChaincodeMessage_Type = 4
	ChaincodeMessage_TRANSACTION         ChaincodeMessage_Type = 5
	ChaincodeMessage_COMPLETED           ChaincodeMessage_Type = 6
	ChaincodeMessage_ERROR               ChaincodeMessage_Type = 7
	ChaincodeMessage_GET_STATE           ChaincodeMessage_Type = 8
	ChaincodeMessage_PUT_STATE           ChaincodeMessage_Type = 9
	ChaincodeMessage_DEL_STATE           ChaincodeMessage_Type = 10
	ChaincodeMessage_INVOKE_CHAINCODE    ChaincodeMessage_Type = 11
	ChaincodeMessage_RESPONSE            ChaincodeMessage_Type = 13
	ChaincodeMessage_GET_STATE_BY_RANGE  ChaincodeMessage_Type = 14
	ChaincodeMessage_GET_QUERY_RESULT    ChaincodeMessage_Type = 15
	ChaincodeMessage_QUERY_STATE_NEXT    ChaincodeMessage_Type = 16
	ChaincodeMessage_QUERY_STATE_CLOSE   ChaincodeMessage_Type = 17
	ChaincodeMessage_KEEPALIVE           ChaincodeMessage_Type = 18
	ChaincodeMessage_GET_HISTORY_FOR_KEY ChaincodeMessage_Type = 19
	ChaincodeMessage_EXECUTE_UPDATE      ChaincodeMessage_Type = 20
)

var ChaincodeMessage_Type_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "REGISTER",
	2:  "REGISTERED",
	3:  "INIT",
	4:  "READY",
	5:  "TRANSACTION",
	6:  "COMPLETED",
	7:  "ERROR",
	8:  "GET_STATE",
	9:  "PUT_STATE",
	10: "DEL_STATE",
	11: "INVOKE_CHAINCODE",
	13: "RESPONSE",
	14: "GET_STATE_BY_RANGE",
	15: "GET_QUERY_RESULT",
	16: "QUERY_STATE_NEXT",
	17: "QUERY_STATE_CLOSE",
	18: "KEEPALIVE",
	19: "GET_HISTORY_FOR_KEY",
	20: "EXECUTE_UPDATE",
}
var ChaincodeMessage_Type_value = map[string]int32{
	"UNDEFINED":           0,
	"REGISTER":            1,
	"REGISTERED":          2,
	"INIT":                3,
	"READY":               4,
	"TRANSACTION":         5,
	"COMPLETED":           6,
	"ERROR":               7,
	"GET_STATE":           8,
	"PUT_STATE":           9,
	"DEL_STATE":           10,
	"INVOKE_CHAINCODE":    11,
	"RESPONSE":            13,
	"GET_STATE_BY_RANGE":  14,
	"GET_QUERY_RESULT":    15,
	"QUERY_STATE_NEXT":    16,
	"QUERY_STATE_CLOSE":   17,
	"KEEPALIVE":           18,
	"GET_HISTORY_FOR_KEY": 19,
	"EXECUTE_UPDATE":      20,
}

func (x ChaincodeMessage_Type) String() string {
	return proto.EnumName(ChaincodeMessage_Type_name, int32(x))
}
func (ChaincodeMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type ChaincodeMessage struct {
	Type      ChaincodeMessage_Type       `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeMessage_Type" json:"type,omitempty"`
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                      `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Txid      string                      `protobuf:"bytes,4,opt,name=txid" json:"txid,omitempty"`
	Proposal  *SignedProposal             `protobuf:"bytes,5,opt,name=proposal" json:"proposal,omitempty"`
	// event emmited by chaincode. Used only with Init or Invoke.
	// This event is then stored (currently)
	// with Block.NonHashData.TransactionResult
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,6,opt,name=chaincode_event,json=chaincodeEvent" json:"chaincode_event,omitempty"`
}

func (m *ChaincodeMessage) Reset()                    { *m = ChaincodeMessage{} }
func (m *ChaincodeMessage) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeMessage) ProtoMessage()               {}
func (*ChaincodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ChaincodeMessage) GetType() ChaincodeMessage_Type {
	if m != nil {
		return m.Type
	}
	return ChaincodeMessage_UNDEFINED
}

func (m *ChaincodeMessage) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChaincodeMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ChaincodeMessage) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *ChaincodeMessage) GetProposal() *SignedProposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func (m *ChaincodeMessage) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

type PutStateInfo struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PutStateInfo) Reset()                    { *m = PutStateInfo{} }
func (m *PutStateInfo) String() string            { return proto.CompactTextString(m) }
func (*PutStateInfo) ProtoMessage()               {}
func (*PutStateInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *PutStateInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutStateInfo) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetStateByRange struct {
	StartKey string `protobuf:"bytes,1,opt,name=startKey" json:"startKey,omitempty"`
	EndKey   string `protobuf:"bytes,2,opt,name=endKey" json:"endKey,omitempty"`
}

func (m *GetStateByRange) Reset()                    { *m = GetStateByRange{} }
func (m *GetStateByRange) String() string            { return proto.CompactTextString(m) }
func (*GetStateByRange) ProtoMessage()               {}
func (*GetStateByRange) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *GetStateByRange) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *GetStateByRange) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

type GetQueryResult struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *GetQueryResult) Reset()                    { *m = GetQueryResult{} }
func (m *GetQueryResult) String() string            { return proto.CompactTextString(m) }
func (*GetQueryResult) ProtoMessage()               {}
func (*GetQueryResult) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *GetQueryResult) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type GetHistoryForKey struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetHistoryForKey) Reset()                    { *m = GetHistoryForKey{} }
func (m *GetHistoryForKey) String() string            { return proto.CompactTextString(m) }
func (*GetHistoryForKey) ProtoMessage()               {}
func (*GetHistoryForKey) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *GetHistoryForKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type QueryStateNext struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateNext) Reset()                    { *m = QueryStateNext{} }
func (m *QueryStateNext) String() string            { return proto.CompactTextString(m) }
func (*QueryStateNext) ProtoMessage()               {}
func (*QueryStateNext) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *QueryStateNext) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryStateClose struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateClose) Reset()                    { *m = QueryStateClose{} }
func (m *QueryStateClose) String() string            { return proto.CompactTextString(m) }
func (*QueryStateClose) ProtoMessage()               {}
func (*QueryStateClose) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *QueryStateClose) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryResultBytes struct {
	ResultBytes []byte `protobuf:"bytes,1,opt,name=resultBytes,proto3" json:"resultBytes,omitempty"`
}

func (m *QueryResultBytes) Reset()                    { *m = QueryResultBytes{} }
func (m *QueryResultBytes) String() string            { return proto.CompactTextString(m) }
func (*QueryResultBytes) ProtoMessage()               {}
func (*QueryResultBytes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *QueryResultBytes) GetResultBytes() []byte {
	if m != nil {
		return m.ResultBytes
	}
	return nil
}

type QueryResponse struct {
	Results []*QueryResultBytes `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	HasMore bool                `protobuf:"varint,2,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
	Id      string              `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *QueryResponse) GetResults() []*QueryResultBytes {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *QueryResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *QueryResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateResponse struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *UpdateResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*ChaincodeMessage)(nil), "protos.ChaincodeMessage")
	proto.RegisterType((*PutStateInfo)(nil), "protos.PutStateInfo")
	proto.RegisterType((*GetStateByRange)(nil), "protos.GetStateByRange")
	proto.RegisterType((*GetQueryResult)(nil), "protos.GetQueryResult")
	proto.RegisterType((*GetHistoryForKey)(nil), "protos.GetHistoryForKey")
	proto.RegisterType((*QueryStateNext)(nil), "protos.QueryStateNext")
	proto.RegisterType((*QueryStateClose)(nil), "protos.QueryStateClose")
	proto.RegisterType((*QueryResultBytes)(nil), "protos.QueryResultBytes")
	proto.RegisterType((*QueryResponse)(nil), "protos.QueryResponse")
	proto.RegisterType((*UpdateResponse)(nil), "protos.UpdateResponse")
	proto.RegisterEnum("protos.ChaincodeMessage_Type", ChaincodeMessage_Type_name, ChaincodeMessage_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChaincodeSupport service

type ChaincodeSupportClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
}

type chaincodeSupportClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeSupportClient(cc *grpc.ClientConn) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
}

func (c *chaincodeSupportClient) Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChaincodeSupport_serviceDesc.Streams[0], c.cc, "/protos.ChaincodeSupport/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeSupportRegisterClient{stream}
	return x, nil
}

type ChaincodeSupport_RegisterClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeSupportRegisterClient struct {
	grpc.ClientStream
}

func (x *chaincodeSupportRegisterClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChaincodeSupport service

type ChaincodeSupportServer interface {
	Register(ChaincodeSupport_RegisterServer) error
}

func RegisterChaincodeSupportServer(s *grpc.Server, srv ChaincodeSupportServer) {
	s.RegisterService(&_ChaincodeSupport_serviceDesc, srv)
}

func _ChaincodeSupport_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeSupportServer).Register(&chaincodeSupportRegisterServer{stream})
}

type ChaincodeSupport_RegisterServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeSupportRegisterServer struct {
	grpc.ServerStream
}

func (x *chaincodeSupportRegisterServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChaincodeSupport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChaincodeSupport",
	HandlerType: (*ChaincodeSupportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer/chaincode_shim.proto",
}

func init() { proto.RegisterFile("peer/chaincode_shim.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x94, 0x51, 0x6f, 0xe2, 0x46,
	0x10, 0xc7, 0x0f, 0x02, 0x09, 0x0c, 0xc4, 0xec, 0x6d, 0xd2, 0xd4, 0x87, 0x54, 0x95, 0x5a, 0x55,
	0x45, 0x5f, 0x4c, 0x4b, 0xab, 0xaa, 0x6f, 0x15, 0xc1, 0x1b, 0x62, 0x85, 0xd8, 0xdc, 0xda, 0x9c,
	0x42, 0x5f, 0x2c, 0x07, 0x6f, 0x8c, 0x15, 0x60, 0x5d, 0x7b, 0x39, 0x9d, 0xdf, 0xfa, 0x3d, 0xfb,
	0x65, 0xaa, 0xb5, 0x31, 0xe1, 0x72, 0xba, 0x27, 0xfb, 0x3f, 0xf3, 0x9b, 0x99, 0xff, 0xae, 0x56,
	0x03, 0xef, 0x62, 0xc6, 0x92, 0xc1, 0x72, 0xe5, 0x47, 0xdb, 0x25, 0x0f, 0x98, 0x97, 0xae, 0xa2,
	0x8d, 0x1e, 0x27, 0x5c, 0x70, 0x7c, 0x9a, 0x7f, 0xd2, 0x6e, 0xf7, 0x15, 0xc2, 0x3e, 0xb2, 0xad,
	0x28, 0x98, 0xee, 0x45, 0x9e, 0x8b, 0x13, 0x1e, 0xf3, 0xd4, 0x5f, 0xef, 0x83, 0xdf, 0x87, 0x9c,
	0x87, 0x6b, 0x36, 0xc8, 0xd5, 0xe3, 0xee, 0x69, 0x20, 0xa2, 0x0d, 0x4b, 0x85, 0xbf, 0x89, 0x0b,
	0x40, 0xfb, 0xb7, 0x0e, 0x68, 0x5c, 0xf6, 0xbb, 0x67, 0x69, 0xea, 0x87, 0x0c, 0xff, 0x0a, 0x35,
	0x91, 0xc5, 0x4c, 0xad, 0xf4, 0x2a, 0x7d, 0x65, 0xf8, 0x5d, 0x81, 0xa6, 0xfa, 0x6b, 0x4e, 0x77,
	0xb3, 0x98, 0xd1, 0x1c, 0xc5, 0x7f, 0x42, 0xf3, 0xd0, 0x5a, 0xad, 0xf6, 0x2a, 0xfd, 0xd6, 0xb0,
	0xab, 0x17, 0xc3, 0xf5, 0x72, 0xb8, 0xee, 0x96, 0x04, 0x7d, 0x81, 0xb1, 0x0a, 0x67, 0xb1, 0x9f,
	0xad, 0xb9, 0x1f, 0xa8, 0x27, 0xbd, 0x4a, 0xbf, 0x4d, 0x4b, 0x89, 0x31, 0xd4, 0xc4, 0xa7, 0x28,
	0x50, 0x6b, 0xbd, 0x4a, 0xbf, 0x49, 0xf3, 0x7f, 0x3c, 0x84, 0x46, 0x79, 0x44, 0xb5, 0x9e, 0x8f,
	0xb9, 0x2a, 0xed, 0x39, 0x51, 0xb8, 0x65, 0xc1, 0x6c, 0x9f, 0xa5, 0x07, 0x0e, 0xff, 0x05, 0x9d,
	0x57, 0x57, 0xa6, 0x9e, 0x7e, 0x5e, 0x7a, 0x38, 0x19, 0x91, 0x59, 0xaa, 0x2c, 0x3f, 0xd3, 0xda,
	0x7f, 0x55, 0xa8, 0xc9, 0xb3, 0xe2, 0x73, 0x68, 0xce, 0x2d, 0x83, 0xdc, 0x98, 0x16, 0x31, 0xd0,
	0x1b, 0xdc, 0x86, 0x06, 0x25, 0x13, 0xd3, 0x71, 0x09, 0x45, 0x15, 0xac, 0x00, 0x94, 0x8a, 0x18,
	0xa8, 0x8a, 0x1b, 0x50, 0x33, 0x2d, 0xd3, 0x45, 0x27, 0xb8, 0x09, 0x75, 0x4a, 0x46, 0xc6, 0x02,
	0xd5, 0x70, 0x07, 0x5a, 0x2e, 0x1d, 0x59, 0xce, 0x68, 0xec, 0x9a, 0xb6, 0x85, 0xea, 0xb2, 0xe5,
	0xd8, 0xbe, 0x9f, 0x4d, 0x89, 0x4b, 0x0c, 0x74, 0x2a, 0x51, 0x42, 0xa9, 0x4d, 0xd1, 0x99, 0xcc,
	0x4c, 0x88, 0xeb, 0x39, 0xee, 0xc8, 0x25, 0xa8, 0x21, 0xe5, 0x6c, 0x5e, 0xca, 0xa6, 0x94, 0x06,
	0x99, 0xee, 0x25, 0xe0, 0x4b, 0x40, 0xa6, 0xf5, 0xc1, 0xbe, 0x23, 0xde, 0xf8, 0x76, 0x64, 0x5a,
	0x63, 0xdb, 0x20, 0xa8, 0x55, 0x18, 0x74, 0x66, 0xb6, 0xe5, 0x10, 0x74, 0x8e, 0xaf, 0x00, 0x1f,
	0x1a, 0x7a, 0xd7, 0x0b, 0x8f, 0x8e, 0xac, 0x09, 0x41, 0x8a, 0xac, 0x95, 0xf1, 0xf7, 0x73, 0x42,
	0x17, 0x1e, 0x25, 0xce, 0x7c, 0xea, 0xa2, 0x8e, 0x8c, 0x16, 0x91, 0x82, 0xb7, 0xc8, 0x83, 0x8b,
	0x10, 0xfe, 0x06, 0xde, 0x1e, 0x47, 0xc7, 0x53, 0xdb, 0x21, 0xe8, 0xad, 0x74, 0x73, 0x47, 0xc8,
	0x6c, 0x34, 0x35, 0x3f, 0x10, 0x84, 0xf1, 0xb7, 0x70, 0x21, 0x3b, 0xde, 0x9a, 0x8e, 0x6b, 0xd3,
	0x85, 0x77, 0x63, 0x53, 0xef, 0x8e, 0x2c, 0xd0, 0x05, 0xc6, 0xa0, 0x90, 0x07, 0x32, 0x9e, 0xbb,
	0xc4, 0x9b, 0xcf, 0x0c, 0x69, 0xfd, 0x52, 0xfb, 0x03, 0xda, 0xb3, 0x9d, 0x70, 0x84, 0x2f, 0x98,
	0xb9, 0x7d, 0xe2, 0x18, 0xc1, 0xc9, 0x33, 0xcb, 0xf2, 0xc7, 0xd7, 0xa4, 0xf2, 0x17, 0x5f, 0x42,
	0xfd, 0xa3, 0xbf, 0xde, 0xb1, 0xfc, 0x61, 0xb5, 0x69, 0x21, 0x34, 0x02, 0x9d, 0x09, 0x2b, 0xea,
	0xae, 0x33, 0xea, 0x6f, 0x43, 0x86, 0xbb, 0xd0, 0x48, 0x85, 0x9f, 0x88, 0xbb, 0x43, 0xfd, 0x41,
	0xe3, 0x2b, 0x38, 0x65, 0xdb, 0x40, 0x66, 0xaa, 0x79, 0x66, 0xaf, 0xb4, 0x9f, 0x40, 0x99, 0x30,
	0xf1, 0x7e, 0xc7, 0x92, 0x8c, 0xb2, 0x74, 0xb7, 0x16, 0x72, 0xdc, 0x3f, 0x52, 0xee, 0x5b, 0x14,
	0x42, 0xfb, 0x11, 0xd0, 0x84, 0x89, 0xdb, 0x28, 0x15, 0x3c, 0xc9, 0x6e, 0x78, 0x22, 0x7b, 0x7e,
	0x61, 0x55, 0xeb, 0x81, 0x92, 0xb7, 0xca, 0x6d, 0x59, 0xec, 0x93, 0xc0, 0x0a, 0x54, 0xa3, 0x60,
	0x8f, 0x54, 0xa3, 0x40, 0xfb, 0x01, 0x3a, 0x2f, 0xc4, 0x78, 0xcd, 0x53, 0xf6, 0x05, 0xf2, 0x3b,
	0xa0, 0x23, 0x3f, 0xd7, 0x99, 0x60, 0x29, 0xee, 0x41, 0x2b, 0x79, 0x91, 0x39, 0xdc, 0xa6, 0xc7,
	0x21, 0x6d, 0x0b, 0xe7, 0x65, 0x55, 0xcc, 0xb7, 0x29, 0xc3, 0x43, 0x38, 0x2b, 0xf2, 0x12, 0x3f,
	0xe9, 0xb7, 0x86, 0x6a, 0xf9, 0xde, 0x5f, 0x77, 0xa7, 0x25, 0x88, 0xdf, 0x41, 0x63, 0xe5, 0xa7,
	0xde, 0x86, 0x27, 0xc5, 0x6d, 0x37, 0xe8, 0xd9, 0xca, 0x4f, 0xef, 0x79, 0x52, 0xba, 0x3c, 0x39,
	0xb8, 0xec, 0x81, 0x32, 0x8f, 0x03, 0x5f, 0xb0, 0xc3, 0x40, 0x05, 0xaa, 0xfc, 0x39, 0xb7, 0xd6,
	0xa0, 0x55, 0xfe, 0x3c, 0x7c, 0x38, 0xda, 0x2d, 0xce, 0x2e, 0x8e, 0x79, 0x22, 0xb0, 0x01, 0x0d,
	0xca, 0xc2, 0x28, 0x15, 0x2c, 0xc1, 0xea, 0xd7, 0x36, 0x4b, 0xf7, 0xab, 0x19, 0xed, 0x4d, 0xbf,
	0xf2, 0x4b, 0xe5, 0xda, 0x06, 0x8d, 0x27, 0xa1, 0xbe, 0xca, 0x62, 0x96, 0xac, 0x59, 0x10, 0xb2,
	0x44, 0x7f, 0xf2, 0x1f, 0x93, 0x68, 0x59, 0xd6, 0xc9, 0x65, 0xf8, 0xf7, 0xcf, 0x61, 0x24, 0x56,
	0xbb, 0x47, 0x7d, 0xc9, 0x37, 0x83, 0x23, 0x74, 0x50, 0xa0, 0xc5, 0x52, 0x4c, 0x07, 0x12, 0x7d,
	0x2c, 0x36, 0xec, 0x6f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x09, 0x8b, 0x5f, 0x85, 0x05,
	0x00, 0x00,
}
