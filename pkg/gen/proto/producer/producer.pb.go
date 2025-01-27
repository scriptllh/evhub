// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.0
// source: producer/producer.proto

package producer

import (
	reflect "reflect"
	sync "sync"

	comm "github.com/tencentmusic/evhub/pkg/gen/proto/comm"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event   *comm.Event        `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Trigger *comm.EventTrigger `protobuf:"bytes,2,opt,name=trigger,proto3" json:"trigger,omitempty"`
	Option  *comm.EventOption  `protobuf:"bytes,3,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *ReportReq) Reset() {
	*x = ReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producer_producer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportReq) ProtoMessage() {}

func (x *ReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_producer_producer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportReq.ProtoReflect.Descriptor instead.
func (*ReportReq) Descriptor() ([]byte, []int) {
	return file_producer_producer_proto_rawDescGZIP(), []int{0}
}

func (x *ReportReq) GetEvent() *comm.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *ReportReq) GetTrigger() *comm.EventTrigger {
	if x != nil {
		return x.Trigger
	}
	return nil
}

func (x *ReportReq) GetOption() *comm.EventOption {
	if x != nil {
		return x.Option
	}
	return nil
}

type ReportRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret     *comm.Result `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	EventId string       `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *ReportRsp) Reset() {
	*x = ReportRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producer_producer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRsp) ProtoMessage() {}

func (x *ReportRsp) ProtoReflect() protoreflect.Message {
	mi := &file_producer_producer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRsp.ProtoReflect.Descriptor instead.
func (*ReportRsp) Descriptor() ([]byte, []int) {
	return file_producer_producer_proto_rawDescGZIP(), []int{1}
}

func (x *ReportRsp) GetRet() *comm.Result {
	if x != nil {
		return x.Ret
	}
	return nil
}

func (x *ReportRsp) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type PrepareReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event      *comm.Event      `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	TxCallback *comm.TxCallback `protobuf:"bytes,2,opt,name=tx_callback,json=txCallback,proto3" json:"tx_callback,omitempty"`
}

func (x *PrepareReq) Reset() {
	*x = PrepareReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producer_producer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareReq) ProtoMessage() {}

func (x *PrepareReq) ProtoReflect() protoreflect.Message {
	mi := &file_producer_producer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareReq.ProtoReflect.Descriptor instead.
func (*PrepareReq) Descriptor() ([]byte, []int) {
	return file_producer_producer_proto_rawDescGZIP(), []int{2}
}

func (x *PrepareReq) GetEvent() *comm.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *PrepareReq) GetTxCallback() *comm.TxCallback {
	if x != nil {
		return x.TxCallback
	}
	return nil
}

type PrepareRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *comm.Result `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Tx  *comm.Tx     `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *PrepareRsp) Reset() {
	*x = PrepareRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producer_producer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareRsp) ProtoMessage() {}

func (x *PrepareRsp) ProtoReflect() protoreflect.Message {
	mi := &file_producer_producer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareRsp.ProtoReflect.Descriptor instead.
func (*PrepareRsp) Descriptor() ([]byte, []int) {
	return file_producer_producer_proto_rawDescGZIP(), []int{3}
}

func (x *PrepareRsp) GetRet() *comm.Result {
	if x != nil {
		return x.Ret
	}
	return nil
}

func (x *PrepareRsp) GetTx() *comm.Tx {
	if x != nil {
		return x.Tx
	}
	return nil
}

type CommitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string             `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Trigger *comm.EventTrigger `protobuf:"bytes,2,opt,name=trigger,proto3" json:"trigger,omitempty"`
}

func (x *CommitReq) Reset() {
	*x = CommitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producer_producer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitReq) ProtoMessage() {}

func (x *CommitReq) ProtoReflect() protoreflect.Message {
	mi := &file_producer_producer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitReq.ProtoReflect.Descriptor instead.
func (*CommitReq) Descriptor() ([]byte, []int) {
	return file_producer_producer_proto_rawDescGZIP(), []int{4}
}

func (x *CommitReq) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *CommitReq) GetTrigger() *comm.EventTrigger {
	if x != nil {
		return x.Trigger
	}
	return nil
}

type CommitRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *comm.Result `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Tx  *comm.Tx     `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *CommitRsp) Reset() {
	*x = CommitRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producer_producer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitRsp) ProtoMessage() {}

func (x *CommitRsp) ProtoReflect() protoreflect.Message {
	mi := &file_producer_producer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitRsp.ProtoReflect.Descriptor instead.
func (*CommitRsp) Descriptor() ([]byte, []int) {
	return file_producer_producer_proto_rawDescGZIP(), []int{5}
}

func (x *CommitRsp) GetRet() *comm.Result {
	if x != nil {
		return x.Ret
	}
	return nil
}

func (x *CommitRsp) GetTx() *comm.Tx {
	if x != nil {
		return x.Tx
	}
	return nil
}

type RollbackReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *RollbackReq) Reset() {
	*x = RollbackReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producer_producer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackReq) ProtoMessage() {}

func (x *RollbackReq) ProtoReflect() protoreflect.Message {
	mi := &file_producer_producer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackReq.ProtoReflect.Descriptor instead.
func (*RollbackReq) Descriptor() ([]byte, []int) {
	return file_producer_producer_proto_rawDescGZIP(), []int{6}
}

func (x *RollbackReq) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type RollbackRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *comm.Result `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Tx  *comm.Tx     `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *RollbackRsp) Reset() {
	*x = RollbackRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producer_producer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackRsp) ProtoMessage() {}

func (x *RollbackRsp) ProtoReflect() protoreflect.Message {
	mi := &file_producer_producer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackRsp.ProtoReflect.Descriptor instead.
func (*RollbackRsp) Descriptor() ([]byte, []int) {
	return file_producer_producer_proto_rawDescGZIP(), []int{7}
}

func (x *RollbackRsp) GetRet() *comm.Result {
	if x != nil {
		return x.Ret
	}
	return nil
}

func (x *RollbackRsp) GetTx() *comm.Tx {
	if x != nil {
		return x.Tx
	}
	return nil
}

var File_producer_producer_proto protoreflect.FileDescriptor

var file_producer_producer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x76, 0x68, 0x75, 0x62,
	0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x1a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x09, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x32, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x07, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x6e, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x74,
	0x78, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x54, 0x78,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x0a, 0x74, 0x78, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x22, 0x52, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52,
	0x73, 0x70, 0x12, 0x24, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x54, 0x78, 0x52, 0x02, 0x74, 0x78, 0x22, 0x5a, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x07, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x73,
	0x70, 0x12, 0x24, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x2e, 0x54, 0x78, 0x52, 0x02, 0x74, 0x78, 0x22, 0x28, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x53, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x73, 0x70,
	0x12, 0x24, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e,
	0x54, 0x78, 0x52, 0x02, 0x74, 0x78, 0x32, 0x98, 0x02, 0x0a, 0x0d, 0x65, 0x76, 0x68, 0x75, 0x62,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x12, 0x1a, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x1a, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x73, 0x70, 0x12, 0x3e, 0x0a, 0x06, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x73, 0x70, 0x12, 0x44, 0x0a, 0x08, 0x52,
	0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x73,
	0x70, 0x42, 0x1a, 0x5a, 0x18, 0x65, 0x76, 0x68, 0x75, 0x62, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_producer_producer_proto_rawDescOnce sync.Once
	file_producer_producer_proto_rawDescData = file_producer_producer_proto_rawDesc
)

func file_producer_producer_proto_rawDescGZIP() []byte {
	file_producer_producer_proto_rawDescOnce.Do(func() {
		file_producer_producer_proto_rawDescData = protoimpl.X.CompressGZIP(file_producer_producer_proto_rawDescData)
	})
	return file_producer_producer_proto_rawDescData
}

var file_producer_producer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_producer_producer_proto_goTypes = []interface{}{
	(*ReportReq)(nil),         // 0: evhub_producer.ReportReq
	(*ReportRsp)(nil),         // 1: evhub_producer.ReportRsp
	(*PrepareReq)(nil),        // 2: evhub_producer.PrepareReq
	(*PrepareRsp)(nil),        // 3: evhub_producer.PrepareRsp
	(*CommitReq)(nil),         // 4: evhub_producer.CommitReq
	(*CommitRsp)(nil),         // 5: evhub_producer.CommitRsp
	(*RollbackReq)(nil),       // 6: evhub_producer.RollbackReq
	(*RollbackRsp)(nil),       // 7: evhub_producer.RollbackRsp
	(*comm.Event)(nil),        // 8: evhub_comm.Event
	(*comm.EventTrigger)(nil), // 9: evhub_comm.EventTrigger
	(*comm.EventOption)(nil),  // 10: evhub_comm.EventOption
	(*comm.Result)(nil),       // 11: evhub_comm.Result
	(*comm.TxCallback)(nil),   // 12: evhub_comm.TxCallback
	(*comm.Tx)(nil),           // 13: evhub_comm.Tx
}
var file_producer_producer_proto_depIdxs = []int32{
	8,  // 0: evhub_producer.ReportReq.event:type_name -> evhub_comm.Event
	9,  // 1: evhub_producer.ReportReq.trigger:type_name -> evhub_comm.EventTrigger
	10, // 2: evhub_producer.ReportReq.option:type_name -> evhub_comm.EventOption
	11, // 3: evhub_producer.ReportRsp.ret:type_name -> evhub_comm.Result
	8,  // 4: evhub_producer.PrepareReq.event:type_name -> evhub_comm.Event
	12, // 5: evhub_producer.PrepareReq.tx_callback:type_name -> evhub_comm.TxCallback
	11, // 6: evhub_producer.PrepareRsp.ret:type_name -> evhub_comm.Result
	13, // 7: evhub_producer.PrepareRsp.tx:type_name -> evhub_comm.Tx
	9,  // 8: evhub_producer.CommitReq.trigger:type_name -> evhub_comm.EventTrigger
	11, // 9: evhub_producer.CommitRsp.ret:type_name -> evhub_comm.Result
	13, // 10: evhub_producer.CommitRsp.tx:type_name -> evhub_comm.Tx
	11, // 11: evhub_producer.RollbackRsp.ret:type_name -> evhub_comm.Result
	13, // 12: evhub_producer.RollbackRsp.tx:type_name -> evhub_comm.Tx
	0,  // 13: evhub_producer.evhubProducer.Report:input_type -> evhub_producer.ReportReq
	2,  // 14: evhub_producer.evhubProducer.Prepare:input_type -> evhub_producer.PrepareReq
	4,  // 15: evhub_producer.evhubProducer.Commit:input_type -> evhub_producer.CommitReq
	6,  // 16: evhub_producer.evhubProducer.Rollback:input_type -> evhub_producer.RollbackReq
	1,  // 17: evhub_producer.evhubProducer.Report:output_type -> evhub_producer.ReportRsp
	3,  // 18: evhub_producer.evhubProducer.Prepare:output_type -> evhub_producer.PrepareRsp
	5,  // 19: evhub_producer.evhubProducer.Commit:output_type -> evhub_producer.CommitRsp
	7,  // 20: evhub_producer.evhubProducer.Rollback:output_type -> evhub_producer.RollbackRsp
	17, // [17:21] is the sub-list for method output_type
	13, // [13:17] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_producer_producer_proto_init() }
func file_producer_producer_proto_init() {
	if File_producer_producer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_producer_producer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportReq); i {
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
		file_producer_producer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRsp); i {
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
		file_producer_producer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareReq); i {
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
		file_producer_producer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareRsp); i {
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
		file_producer_producer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitReq); i {
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
		file_producer_producer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitRsp); i {
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
		file_producer_producer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackReq); i {
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
		file_producer_producer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackRsp); i {
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
			RawDescriptor: file_producer_producer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_producer_producer_proto_goTypes,
		DependencyIndexes: file_producer_producer_proto_depIdxs,
		MessageInfos:      file_producer_producer_proto_msgTypes,
	}.Build()
	File_producer_producer_proto = out.File
	file_producer_producer_proto_rawDesc = nil
	file_producer_producer_proto_goTypes = nil
	file_producer_producer_proto_depIdxs = nil
}
