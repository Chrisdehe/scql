// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: pkg/audit/common.proto

package audit

import (
	scql "github.com/secretflow/scql/pkg/proto-gen/scql"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventName int32

const (
	EventName_UNCATEGORIZED    EventName = 0
	EventName_RUN_SYNC_QUERY   EventName = 1
	EventName_RUN_ASYNC_QUERY  EventName = 2
	EventName_ASYNC_COMPLETE   EventName = 3
	EventName_SESSION_PARAMS   EventName = 4
	EventName_FETCH_RESULT     EventName = 5
	EventName_PLAN_DETAIL      EventName = 6
	EventName_DAG_DETAIL       EventName = 7
	EventName_CCL_DETAIL       EventName = 8
	EventName_RUN_SUB_DAG      EventName = 9
	EventName_RUN_EXEC_PLAN    EventName = 10
	EventName_CREATE_SESSION   EventName = 11
	EventName_STOP_SESSION     EventName = 12
	EventName_JOIN_DETAIL      EventName = 13
	EventName_IN_DETAIL        EventName = 14
	EventName_SQL_DETAIL       EventName = 15
	EventName_DUMP_FILE_DETAIL EventName = 16
	EventName_PUBLISH_DETAIL   EventName = 17
)

// Enum value maps for EventName.
var (
	EventName_name = map[int32]string{
		0:  "UNCATEGORIZED",
		1:  "RUN_SYNC_QUERY",
		2:  "RUN_ASYNC_QUERY",
		3:  "ASYNC_COMPLETE",
		4:  "SESSION_PARAMS",
		5:  "FETCH_RESULT",
		6:  "PLAN_DETAIL",
		7:  "DAG_DETAIL",
		8:  "CCL_DETAIL",
		9:  "RUN_SUB_DAG",
		10: "RUN_EXEC_PLAN",
		11: "CREATE_SESSION",
		12: "STOP_SESSION",
		13: "JOIN_DETAIL",
		14: "IN_DETAIL",
		15: "SQL_DETAIL",
		16: "DUMP_FILE_DETAIL",
		17: "PUBLISH_DETAIL",
	}
	EventName_value = map[string]int32{
		"UNCATEGORIZED":    0,
		"RUN_SYNC_QUERY":   1,
		"RUN_ASYNC_QUERY":  2,
		"ASYNC_COMPLETE":   3,
		"SESSION_PARAMS":   4,
		"FETCH_RESULT":     5,
		"PLAN_DETAIL":      6,
		"DAG_DETAIL":       7,
		"CCL_DETAIL":       8,
		"RUN_SUB_DAG":      9,
		"RUN_EXEC_PLAN":    10,
		"CREATE_SESSION":   11,
		"STOP_SESSION":     12,
		"JOIN_DETAIL":      13,
		"IN_DETAIL":        14,
		"SQL_DETAIL":       15,
		"DUMP_FILE_DETAIL": 16,
		"PUBLISH_DETAIL":   17,
	}
)

func (x EventName) Enum() *EventName {
	p := new(EventName)
	*p = x
	return p
}

func (x EventName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventName) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_audit_common_proto_enumTypes[0].Descriptor()
}

func (EventName) Type() protoreflect.EnumType {
	return &file_pkg_audit_common_proto_enumTypes[0]
}

func (x EventName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventName.Descriptor instead.
func (EventName) EnumDescriptor() ([]byte, []int) {
	return file_pkg_audit_common_proto_rawDescGZIP(), []int{0}
}

type AuditHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Status    *scql.Status           `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	EventName EventName              `protobuf:"varint,3,opt,name=event_name,json=eventName,proto3,enum=audit.pb.EventName" json:"event_name,omitempty"`
	SessionId string                 `protobuf:"bytes,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *AuditHeader) Reset() {
	*x = AuditHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_audit_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditHeader) ProtoMessage() {}

func (x *AuditHeader) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_audit_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditHeader.ProtoReflect.Descriptor instead.
func (*AuditHeader) Descriptor() ([]byte, []int) {
	return file_pkg_audit_common_proto_rawDescGZIP(), []int{0}
}

func (x *AuditHeader) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *AuditHeader) GetStatus() *scql.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AuditHeader) GetEventName() EventName {
	if x != nil {
		return x.EventName
	}
	return EventName_UNCATEGORIZED
}

func (x *AuditHeader) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type UncategorizedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlPath  string `protobuf:"bytes,1,opt,name=url_path,json=urlPath,proto3" json:"url_path,omitempty"`
	SourceIp string `protobuf:"bytes,2,opt,name=source_ip,json=sourceIp,proto3" json:"source_ip,omitempty"`
}

func (x *UncategorizedEvent) Reset() {
	*x = UncategorizedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_audit_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UncategorizedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UncategorizedEvent) ProtoMessage() {}

func (x *UncategorizedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_audit_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UncategorizedEvent.ProtoReflect.Descriptor instead.
func (*UncategorizedEvent) Descriptor() ([]byte, []int) {
	return file_pkg_audit_common_proto_rawDescGZIP(), []int{1}
}

func (x *UncategorizedEvent) GetUrlPath() string {
	if x != nil {
		return x.UrlPath
	}
	return ""
}

func (x *UncategorizedEvent) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Inputs  map[string]*Strings `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Outputs map[string]*Strings `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_audit_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_audit_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_pkg_audit_common_proto_rawDescGZIP(), []int{2}
}

func (x *NodeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeInfo) GetInputs() map[string]*Strings {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *NodeInfo) GetOutputs() map[string]*Strings {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type Strings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ss []string `protobuf:"bytes,1,rep,name=ss,proto3" json:"ss,omitempty"`
}

func (x *Strings) Reset() {
	*x = Strings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_audit_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strings) ProtoMessage() {}

func (x *Strings) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_audit_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strings.ProtoReflect.Descriptor instead.
func (*Strings) Descriptor() ([]byte, []int) {
	return file_pkg_audit_common_proto_rawDescGZIP(), []int{3}
}

func (x *Strings) GetSs() []string {
	if x != nil {
		return x.Ss
	}
	return nil
}

var File_pkg_audit_common_proto protoreflect.FileDescriptor

var file_pkg_audit_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x70, 0x62, 0x1a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x64, 0x69, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32,
	0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x4c, 0x0a, 0x12, 0x55, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x22,
	0xae, 0x02, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x36, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x1a, 0x4c, 0x0a, 0x0b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x4d, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x19, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x73, 0x73, 0x2a, 0xd6, 0x02, 0x0a, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x52, 0x55, 0x4e, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x52, 0x55, 0x4e, 0x5f, 0x41, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x51, 0x55,
	0x45, 0x52, 0x59, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x53, 0x10, 0x04, 0x12, 0x10, 0x0a,
	0x0c, 0x46, 0x45, 0x54, 0x43, 0x48, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x05, 0x12,
	0x0f, 0x0a, 0x0b, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x06,
	0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x41, 0x47, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x07,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x43, 0x4c, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x08,
	0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x44, 0x41, 0x47, 0x10,
	0x09, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x55, 0x4e, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x5f, 0x50, 0x4c,
	0x41, 0x4e, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x53,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x4f, 0x50,
	0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x0c, 0x12, 0x0f, 0x0a, 0x0b, 0x4a, 0x4f,
	0x49, 0x4e, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x0d, 0x12, 0x0d, 0x0a, 0x09, 0x49,
	0x4e, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x0e, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x51,
	0x4c, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x55,
	0x4d, 0x50, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x10,
	0x12, 0x12, 0x0a, 0x0e, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x5f, 0x44, 0x45, 0x54, 0x41,
	0x49, 0x4c, 0x10, 0x11, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_audit_common_proto_rawDescOnce sync.Once
	file_pkg_audit_common_proto_rawDescData = file_pkg_audit_common_proto_rawDesc
)

func file_pkg_audit_common_proto_rawDescGZIP() []byte {
	file_pkg_audit_common_proto_rawDescOnce.Do(func() {
		file_pkg_audit_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_audit_common_proto_rawDescData)
	})
	return file_pkg_audit_common_proto_rawDescData
}

var file_pkg_audit_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_audit_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_audit_common_proto_goTypes = []interface{}{
	(EventName)(0),                // 0: audit.pb.EventName
	(*AuditHeader)(nil),           // 1: audit.pb.AuditHeader
	(*UncategorizedEvent)(nil),    // 2: audit.pb.UncategorizedEvent
	(*NodeInfo)(nil),              // 3: audit.pb.NodeInfo
	(*Strings)(nil),               // 4: audit.pb.strings
	nil,                           // 5: audit.pb.NodeInfo.InputsEntry
	nil,                           // 6: audit.pb.NodeInfo.OutputsEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*scql.Status)(nil),           // 8: scql.pb.Status
}
var file_pkg_audit_common_proto_depIdxs = []int32{
	7, // 0: audit.pb.AuditHeader.time:type_name -> google.protobuf.Timestamp
	8, // 1: audit.pb.AuditHeader.status:type_name -> scql.pb.Status
	0, // 2: audit.pb.AuditHeader.event_name:type_name -> audit.pb.EventName
	5, // 3: audit.pb.NodeInfo.inputs:type_name -> audit.pb.NodeInfo.InputsEntry
	6, // 4: audit.pb.NodeInfo.outputs:type_name -> audit.pb.NodeInfo.OutputsEntry
	4, // 5: audit.pb.NodeInfo.InputsEntry.value:type_name -> audit.pb.strings
	4, // 6: audit.pb.NodeInfo.OutputsEntry.value:type_name -> audit.pb.strings
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_audit_common_proto_init() }
func file_pkg_audit_common_proto_init() {
	if File_pkg_audit_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_audit_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditHeader); i {
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
		file_pkg_audit_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UncategorizedEvent); i {
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
		file_pkg_audit_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_pkg_audit_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strings); i {
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
			RawDescriptor: file_pkg_audit_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_audit_common_proto_goTypes,
		DependencyIndexes: file_pkg_audit_common_proto_depIdxs,
		EnumInfos:         file_pkg_audit_common_proto_enumTypes,
		MessageInfos:      file_pkg_audit_common_proto_msgTypes,
	}.Build()
	File_pkg_audit_common_proto = out.File
	file_pkg_audit_common_proto_rawDesc = nil
	file_pkg_audit_common_proto_goTypes = nil
	file_pkg_audit_common_proto_depIdxs = nil
}
