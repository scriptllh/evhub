// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.0
// source: transaction/transaction.proto

package transaction

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

type CallbackReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *comm.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Tx    *comm.Tx    `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *CallbackReq) Reset() {
	*x = CallbackReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackReq) ProtoMessage() {}

func (x *CallbackReq) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackReq.ProtoReflect.Descriptor instead.
func (*CallbackReq) Descriptor() ([]byte, []int) {
	return file_transaction_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *CallbackReq) GetEvent() *comm.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *CallbackReq) GetTx() *comm.Tx {
	if x != nil {
		return x.Tx
	}
	return nil
}

type CallbackRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret      *comm.Result       `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	TxStatus comm.TxStatus      `protobuf:"varint,2,opt,name=tx_status,json=txStatus,proto3,enum=evhub_comm.TxStatus" json:"tx_status,omitempty"`
	Trigger  *comm.EventTrigger `protobuf:"bytes,3,opt,name=trigger,proto3" json:"trigger,omitempty"`
}

func (x *CallbackRsp) Reset() {
	*x = CallbackRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackRsp) ProtoMessage() {}

func (x *CallbackRsp) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackRsp.ProtoReflect.Descriptor instead.
func (*CallbackRsp) Descriptor() ([]byte, []int) {
	return file_transaction_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *CallbackRsp) GetRet() *comm.Result {
	if x != nil {
		return x.Ret
	}
	return nil
}

func (x *CallbackRsp) GetTxStatus() comm.TxStatus {
	if x != nil {
		return x.TxStatus
	}
	return comm.TxStatus_TX_STATUS_INVALID
}

func (x *CallbackRsp) GetTrigger() *comm.EventTrigger {
	if x != nil {
		return x.Trigger
	}
	return nil
}

var File_transaction_transaction_proto protoreflect.FileDescriptor

var file_transaction_transaction_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x74,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x54, 0x78, 0x52, 0x02, 0x74, 0x78, 0x22, 0x9a, 0x01, 0x0a, 0x0b,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x03, 0x72,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x03, 0x72, 0x65,
	0x74, 0x12, 0x31, 0x0a, 0x09, 0x74, 0x78, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x74, 0x78, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x32, 0x5e, 0x0a, 0x10, 0x65, 0x76, 0x68, 0x75,
	0x62, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x08,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1e, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x65, 0x76, 0x68, 0x75, 0x62,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x42, 0x1d, 0x5a, 0x1b, 0x65, 0x76, 0x68, 0x75,
	0x62, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_transaction_proto_rawDescOnce sync.Once
	file_transaction_transaction_proto_rawDescData = file_transaction_transaction_proto_rawDesc
)

func file_transaction_transaction_proto_rawDescGZIP() []byte {
	file_transaction_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_transaction_proto_rawDescData)
	})
	return file_transaction_transaction_proto_rawDescData
}

var file_transaction_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transaction_transaction_proto_goTypes = []interface{}{
	(*CallbackReq)(nil),       // 0: evhub_transaction.CallbackReq
	(*CallbackRsp)(nil),       // 1: evhub_transaction.CallbackRsp
	(*comm.Event)(nil),        // 2: evhub_comm.Event
	(*comm.Tx)(nil),           // 3: evhub_comm.Tx
	(*comm.Result)(nil),       // 4: evhub_comm.Result
	(comm.TxStatus)(0),        // 5: evhub_comm.TxStatus
	(*comm.EventTrigger)(nil), // 6: evhub_comm.EventTrigger
}
var file_transaction_transaction_proto_depIdxs = []int32{
	2, // 0: evhub_transaction.CallbackReq.event:type_name -> evhub_comm.Event
	3, // 1: evhub_transaction.CallbackReq.tx:type_name -> evhub_comm.Tx
	4, // 2: evhub_transaction.CallbackRsp.ret:type_name -> evhub_comm.Result
	5, // 3: evhub_transaction.CallbackRsp.tx_status:type_name -> evhub_comm.TxStatus
	6, // 4: evhub_transaction.CallbackRsp.trigger:type_name -> evhub_comm.EventTrigger
	0, // 5: evhub_transaction.evhubTransaction.Callback:input_type -> evhub_transaction.CallbackReq
	1, // 6: evhub_transaction.evhubTransaction.Callback:output_type -> evhub_transaction.CallbackRsp
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_transaction_transaction_proto_init() }
func file_transaction_transaction_proto_init() {
	if File_transaction_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackReq); i {
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
		file_transaction_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackRsp); i {
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
			RawDescriptor: file_transaction_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transaction_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_transaction_proto_msgTypes,
	}.Build()
	File_transaction_transaction_proto = out.File
	file_transaction_transaction_proto_rawDesc = nil
	file_transaction_transaction_proto_goTypes = nil
	file_transaction_transaction_proto_depIdxs = nil
}
