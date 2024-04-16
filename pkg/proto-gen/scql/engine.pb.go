// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: api/engine.proto

package scql

import (
	context "context"
	spu "github.com/secretflow/scql/pkg/proto-gen/spu"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StopSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Reason    string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *StopSessionRequest) Reset() {
	*x = StopSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopSessionRequest) ProtoMessage() {}

func (x *StopSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopSessionRequest.ProtoReflect.Descriptor instead.
func (*StopSessionRequest) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{0}
}

func (x *StopSessionRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *StopSessionRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type SessionStartParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartyCode     string                      `protobuf:"bytes,1,opt,name=party_code,json=partyCode,proto3" json:"party_code,omitempty"`
	Parties       []*SessionStartParams_Party `protobuf:"bytes,2,rep,name=parties,proto3" json:"parties,omitempty"`
	SessionId     string                      `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	SpuRuntimeCfg *spu.RuntimeConfig          `protobuf:"bytes,4,opt,name=spu_runtime_cfg,json=spuRuntimeCfg,proto3" json:"spu_runtime_cfg,omitempty"`
	TimeZone      string                      `protobuf:"bytes,5,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
}

func (x *SessionStartParams) Reset() {
	*x = SessionStartParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionStartParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionStartParams) ProtoMessage() {}

func (x *SessionStartParams) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionStartParams.ProtoReflect.Descriptor instead.
func (*SessionStartParams) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{1}
}

func (x *SessionStartParams) GetPartyCode() string {
	if x != nil {
		return x.PartyCode
	}
	return ""
}

func (x *SessionStartParams) GetParties() []*SessionStartParams_Party {
	if x != nil {
		return x.Parties
	}
	return nil
}

func (x *SessionStartParams) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *SessionStartParams) GetSpuRuntimeCfg() *spu.RuntimeConfig {
	if x != nil {
		return x.SpuRuntimeCfg
	}
	return nil
}

func (x *SessionStartParams) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

type RunExecutionPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionParams *SessionStartParams `protobuf:"bytes,1,opt,name=session_params,json=sessionParams,proto3" json:"session_params,omitempty"`
	Graph         *SubGraph           `protobuf:"bytes,2,opt,name=graph,proto3" json:"graph,omitempty"`
	Async         bool                `protobuf:"varint,3,opt,name=async,proto3" json:"async,omitempty"`
	CallbackUrl   string              `protobuf:"bytes,4,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
	DebugOpts     *DebugOptions       `protobuf:"bytes,5,opt,name=debug_opts,json=debugOpts,proto3" json:"debug_opts,omitempty"`
}

func (x *RunExecutionPlanRequest) Reset() {
	*x = RunExecutionPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunExecutionPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunExecutionPlanRequest) ProtoMessage() {}

func (x *RunExecutionPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunExecutionPlanRequest.ProtoReflect.Descriptor instead.
func (*RunExecutionPlanRequest) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{2}
}

func (x *RunExecutionPlanRequest) GetSessionParams() *SessionStartParams {
	if x != nil {
		return x.SessionParams
	}
	return nil
}

func (x *RunExecutionPlanRequest) GetGraph() *SubGraph {
	if x != nil {
		return x.Graph
	}
	return nil
}

func (x *RunExecutionPlanRequest) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *RunExecutionPlanRequest) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

func (x *RunExecutionPlanRequest) GetDebugOpts() *DebugOptions {
	if x != nil {
		return x.DebugOpts
	}
	return nil
}

type RunExecutionPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          *Status   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	OutColumns      []*Tensor `protobuf:"bytes,2,rep,name=out_columns,json=outColumns,proto3" json:"out_columns,omitempty"`
	SessionId       string    `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	PartyCode       string    `protobuf:"bytes,4,opt,name=party_code,json=partyCode,proto3" json:"party_code,omitempty"`
	NumRowsAffected int64     `protobuf:"varint,5,opt,name=num_rows_affected,json=numRowsAffected,proto3" json:"num_rows_affected,omitempty"`
}

func (x *RunExecutionPlanResponse) Reset() {
	*x = RunExecutionPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunExecutionPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunExecutionPlanResponse) ProtoMessage() {}

func (x *RunExecutionPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunExecutionPlanResponse.ProtoReflect.Descriptor instead.
func (*RunExecutionPlanResponse) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{3}
}

func (x *RunExecutionPlanResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *RunExecutionPlanResponse) GetOutColumns() []*Tensor {
	if x != nil {
		return x.OutColumns
	}
	return nil
}

func (x *RunExecutionPlanResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *RunExecutionPlanResponse) GetPartyCode() string {
	if x != nil {
		return x.PartyCode
	}
	return ""
}

func (x *RunExecutionPlanResponse) GetNumRowsAffected() int64 {
	if x != nil {
		return x.NumRowsAffected
	}
	return 0
}

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          *Status   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	OutColumns      []*Tensor `protobuf:"bytes,2,rep,name=out_columns,json=outColumns,proto3" json:"out_columns,omitempty"`
	SessionId       string    `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	PartyCode       string    `protobuf:"bytes,4,opt,name=party_code,json=partyCode,proto3" json:"party_code,omitempty"`
	NumRowsAffected int64     `protobuf:"varint,5,opt,name=num_rows_affected,json=numRowsAffected,proto3" json:"num_rows_affected,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{4}
}

func (x *ReportRequest) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ReportRequest) GetOutColumns() []*Tensor {
	if x != nil {
		return x.OutColumns
	}
	return nil
}

func (x *ReportRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ReportRequest) GetPartyCode() string {
	if x != nil {
		return x.PartyCode
	}
	return ""
}

func (x *ReportRequest) GetNumRowsAffected() int64 {
	if x != nil {
		return x.NumRowsAffected
	}
	return 0
}

type SessionStartParams_Party struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Host      string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Rank      int32  `protobuf:"varint,4,opt,name=rank,proto3" json:"rank,omitempty"`
	PublicKey string `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *SessionStartParams_Party) Reset() {
	*x = SessionStartParams_Party{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionStartParams_Party) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionStartParams_Party) ProtoMessage() {}

func (x *SessionStartParams_Party) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionStartParams_Party.ProtoReflect.Descriptor instead.
func (*SessionStartParams_Party) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SessionStartParams_Party) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SessionStartParams_Party) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SessionStartParams_Party) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SessionStartParams_Party) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *SessionStartParams_Party) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

var File_api_engine_proto protoreflect.FileDescriptor

var file_api_engine_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6c, 0x69, 0x62, 0x73, 0x70, 0x75,
	0x2f, 0x73, 0x70, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0xe0, 0x02, 0x0a, 0x12, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x3a, 0x0a, 0x0f, 0x73, 0x70, 0x75, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x63, 0x66, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x70, 0x75,
	0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d,
	0x73, 0x70, 0x75, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x66, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x1a, 0x76, 0x0a, 0x05, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x22, 0xf5, 0x01, 0x0a, 0x17, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42,
	0x0a, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x52, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x52, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x73, 0x79, 0x6e, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e,
	0x63, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x55, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x6f, 0x70,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x09, 0x64, 0x65, 0x62, 0x75, 0x67, 0x4f, 0x70, 0x74, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x18, 0x52,
	0x75, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x30, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x61, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x75, 0x6d,
	0x52, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0xd4, 0x01, 0x0a,
	0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x5f, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x6f,
	0x75, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x5f, 0x72,
	0x6f, 0x77, 0x73, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x32, 0xa9, 0x01, 0x0a, 0x11, 0x53, 0x43, 0x51, 0x4c, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x52, 0x75, 0x6e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x20, 0x2e,
	0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x70,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0x50, 0x0a, 0x14, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x38, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x13, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x73,
	0x63, 0x71, 0x6c, 0x80, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_engine_proto_rawDescOnce sync.Once
	file_api_engine_proto_rawDescData = file_api_engine_proto_rawDesc
)

func file_api_engine_proto_rawDescGZIP() []byte {
	file_api_engine_proto_rawDescOnce.Do(func() {
		file_api_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_engine_proto_rawDescData)
	})
	return file_api_engine_proto_rawDescData
}

var file_api_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_engine_proto_goTypes = []interface{}{
	(*StopSessionRequest)(nil),       // 0: scql.pb.StopSessionRequest
	(*SessionStartParams)(nil),       // 1: scql.pb.SessionStartParams
	(*RunExecutionPlanRequest)(nil),  // 2: scql.pb.RunExecutionPlanRequest
	(*RunExecutionPlanResponse)(nil), // 3: scql.pb.RunExecutionPlanResponse
	(*ReportRequest)(nil),            // 4: scql.pb.ReportRequest
	(*SessionStartParams_Party)(nil), // 5: scql.pb.SessionStartParams.Party
	(*spu.RuntimeConfig)(nil),        // 6: spu.RuntimeConfig
	(*SubGraph)(nil),                 // 7: scql.pb.SubGraph
	(*DebugOptions)(nil),             // 8: scql.pb.DebugOptions
	(*Status)(nil),                   // 9: scql.pb.Status
	(*Tensor)(nil),                   // 10: scql.pb.Tensor
	(*emptypb.Empty)(nil),            // 11: google.protobuf.Empty
}
var file_api_engine_proto_depIdxs = []int32{
	5,  // 0: scql.pb.SessionStartParams.parties:type_name -> scql.pb.SessionStartParams.Party
	6,  // 1: scql.pb.SessionStartParams.spu_runtime_cfg:type_name -> spu.RuntimeConfig
	1,  // 2: scql.pb.RunExecutionPlanRequest.session_params:type_name -> scql.pb.SessionStartParams
	7,  // 3: scql.pb.RunExecutionPlanRequest.graph:type_name -> scql.pb.SubGraph
	8,  // 4: scql.pb.RunExecutionPlanRequest.debug_opts:type_name -> scql.pb.DebugOptions
	9,  // 5: scql.pb.RunExecutionPlanResponse.status:type_name -> scql.pb.Status
	10, // 6: scql.pb.RunExecutionPlanResponse.out_columns:type_name -> scql.pb.Tensor
	9,  // 7: scql.pb.ReportRequest.status:type_name -> scql.pb.Status
	10, // 8: scql.pb.ReportRequest.out_columns:type_name -> scql.pb.Tensor
	2,  // 9: scql.pb.SCQLEngineService.RunExecutionPlan:input_type -> scql.pb.RunExecutionPlanRequest
	0,  // 10: scql.pb.SCQLEngineService.StopSession:input_type -> scql.pb.StopSessionRequest
	4,  // 11: scql.pb.EngineResultCallback.Report:input_type -> scql.pb.ReportRequest
	3,  // 12: scql.pb.SCQLEngineService.RunExecutionPlan:output_type -> scql.pb.RunExecutionPlanResponse
	9,  // 13: scql.pb.SCQLEngineService.StopSession:output_type -> scql.pb.Status
	11, // 14: scql.pb.EngineResultCallback.Report:output_type -> google.protobuf.Empty
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_engine_proto_init() }
func file_api_engine_proto_init() {
	if File_api_engine_proto != nil {
		return
	}
	file_api_common_proto_init()
	file_api_core_proto_init()
	file_api_status_proto_init()
	file_api_subgraph_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopSessionRequest); i {
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
		file_api_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionStartParams); i {
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
		file_api_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunExecutionPlanRequest); i {
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
		file_api_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunExecutionPlanResponse); i {
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
		file_api_engine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
		file_api_engine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionStartParams_Party); i {
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
			RawDescriptor: file_api_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_engine_proto_goTypes,
		DependencyIndexes: file_api_engine_proto_depIdxs,
		MessageInfos:      file_api_engine_proto_msgTypes,
	}.Build()
	File_api_engine_proto = out.File
	file_api_engine_proto_rawDesc = nil
	file_api_engine_proto_goTypes = nil
	file_api_engine_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SCQLEngineServiceClient is the client API for SCQLEngineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SCQLEngineServiceClient interface {
	RunExecutionPlan(ctx context.Context, in *RunExecutionPlanRequest, opts ...grpc.CallOption) (*RunExecutionPlanResponse, error)
	StopSession(ctx context.Context, in *StopSessionRequest, opts ...grpc.CallOption) (*Status, error)
}

type sCQLEngineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSCQLEngineServiceClient(cc grpc.ClientConnInterface) SCQLEngineServiceClient {
	return &sCQLEngineServiceClient{cc}
}

func (c *sCQLEngineServiceClient) RunExecutionPlan(ctx context.Context, in *RunExecutionPlanRequest, opts ...grpc.CallOption) (*RunExecutionPlanResponse, error) {
	out := new(RunExecutionPlanResponse)
	err := c.cc.Invoke(ctx, "/scql.pb.SCQLEngineService/RunExecutionPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCQLEngineServiceClient) StopSession(ctx context.Context, in *StopSessionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/scql.pb.SCQLEngineService/StopSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SCQLEngineServiceServer is the server API for SCQLEngineService service.
type SCQLEngineServiceServer interface {
	RunExecutionPlan(context.Context, *RunExecutionPlanRequest) (*RunExecutionPlanResponse, error)
	StopSession(context.Context, *StopSessionRequest) (*Status, error)
}

// UnimplementedSCQLEngineServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSCQLEngineServiceServer struct {
}

func (*UnimplementedSCQLEngineServiceServer) RunExecutionPlan(context.Context, *RunExecutionPlanRequest) (*RunExecutionPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunExecutionPlan not implemented")
}
func (*UnimplementedSCQLEngineServiceServer) StopSession(context.Context, *StopSessionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopSession not implemented")
}

func RegisterSCQLEngineServiceServer(s *grpc.Server, srv SCQLEngineServiceServer) {
	s.RegisterService(&_SCQLEngineService_serviceDesc, srv)
}

func _SCQLEngineService_RunExecutionPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunExecutionPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCQLEngineServiceServer).RunExecutionPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scql.pb.SCQLEngineService/RunExecutionPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCQLEngineServiceServer).RunExecutionPlan(ctx, req.(*RunExecutionPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCQLEngineService_StopSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCQLEngineServiceServer).StopSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scql.pb.SCQLEngineService/StopSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCQLEngineServiceServer).StopSession(ctx, req.(*StopSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SCQLEngineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scql.pb.SCQLEngineService",
	HandlerType: (*SCQLEngineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunExecutionPlan",
			Handler:    _SCQLEngineService_RunExecutionPlan_Handler,
		},
		{
			MethodName: "StopSession",
			Handler:    _SCQLEngineService_StopSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/engine.proto",
}

// EngineResultCallbackClient is the client API for EngineResultCallback service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EngineResultCallbackClient interface {
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type engineResultCallbackClient struct {
	cc grpc.ClientConnInterface
}

func NewEngineResultCallbackClient(cc grpc.ClientConnInterface) EngineResultCallbackClient {
	return &engineResultCallbackClient{cc}
}

func (c *engineResultCallbackClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scql.pb.EngineResultCallback/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineResultCallbackServer is the server API for EngineResultCallback service.
type EngineResultCallbackServer interface {
	Report(context.Context, *ReportRequest) (*emptypb.Empty, error)
}

// UnimplementedEngineResultCallbackServer can be embedded to have forward compatible implementations.
type UnimplementedEngineResultCallbackServer struct {
}

func (*UnimplementedEngineResultCallbackServer) Report(context.Context, *ReportRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}

func RegisterEngineResultCallbackServer(s *grpc.Server, srv EngineResultCallbackServer) {
	s.RegisterService(&_EngineResultCallback_serviceDesc, srv)
}

func _EngineResultCallback_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineResultCallbackServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scql.pb.EngineResultCallback/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineResultCallbackServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EngineResultCallback_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scql.pb.EngineResultCallback",
	HandlerType: (*EngineResultCallbackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _EngineResultCallback_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/engine.proto",
}
