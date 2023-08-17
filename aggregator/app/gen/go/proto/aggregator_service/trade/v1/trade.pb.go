// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/aggregator_service/trade/v1/trade.proto

package pb_trade

import (
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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price     float64                `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Size      float64                `protobuf:"fixed64,3,opt,name=size,proto3" json:"size,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Buy       bool                   `protobuf:"varint,5,opt,name=buy,proto3" json:"buy,omitempty"`
	Sell      bool                   `protobuf:"varint,6,opt,name=sell,proto3" json:"sell,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aggregator_service_trade_v1_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aggregator_service_trade_v1_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_aggregator_service_trade_v1_trade_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateRequest) GetSize() float64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CreateRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *CreateRequest) GetBuy() bool {
	if x != nil {
		return x.Buy
	}
	return false
}

func (x *CreateRequest) GetSell() bool {
	if x != nil {
		return x.Sell
	}
	return false
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_aggregator_service_trade_v1_trade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_aggregator_service_trade_v1_trade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_aggregator_service_trade_v1_trade_proto_rawDescGZIP(), []int{1}
}

var File_proto_aggregator_service_trade_v1_trade_proto protoreflect.FileDescriptor

var file_proto_aggregator_service_trade_v1_trade_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x75, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x62, 0x75, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x65, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73,
	0x65, 0x6c, 0x6c, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x81, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x6b, 0x5a, 0x69, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x72, 0x4d, 0x6f, 0x6c, 0x6f, 0x64, 0x79,
	0x61, 0x6b, 0x6f, 0x76, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62,
	0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_aggregator_service_trade_v1_trade_proto_rawDescOnce sync.Once
	file_proto_aggregator_service_trade_v1_trade_proto_rawDescData = file_proto_aggregator_service_trade_v1_trade_proto_rawDesc
)

func file_proto_aggregator_service_trade_v1_trade_proto_rawDescGZIP() []byte {
	file_proto_aggregator_service_trade_v1_trade_proto_rawDescOnce.Do(func() {
		file_proto_aggregator_service_trade_v1_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_aggregator_service_trade_v1_trade_proto_rawDescData)
	})
	return file_proto_aggregator_service_trade_v1_trade_proto_rawDescData
}

var file_proto_aggregator_service_trade_v1_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_aggregator_service_trade_v1_trade_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),         // 0: proto.aggregator_service.trade.v1.CreateRequest
	(*CreateResponse)(nil),        // 1: proto.aggregator_service.trade.v1.CreateResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_aggregator_service_trade_v1_trade_proto_depIdxs = []int32{
	2, // 0: proto.aggregator_service.trade.v1.CreateRequest.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: proto.aggregator_service.trade.v1.TradeService.Create:input_type -> proto.aggregator_service.trade.v1.CreateRequest
	1, // 2: proto.aggregator_service.trade.v1.TradeService.Create:output_type -> proto.aggregator_service.trade.v1.CreateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_aggregator_service_trade_v1_trade_proto_init() }
func file_proto_aggregator_service_trade_v1_trade_proto_init() {
	if File_proto_aggregator_service_trade_v1_trade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_aggregator_service_trade_v1_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_proto_aggregator_service_trade_v1_trade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_proto_aggregator_service_trade_v1_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_aggregator_service_trade_v1_trade_proto_goTypes,
		DependencyIndexes: file_proto_aggregator_service_trade_v1_trade_proto_depIdxs,
		MessageInfos:      file_proto_aggregator_service_trade_v1_trade_proto_msgTypes,
	}.Build()
	File_proto_aggregator_service_trade_v1_trade_proto = out.File
	file_proto_aggregator_service_trade_v1_trade_proto_rawDesc = nil
	file_proto_aggregator_service_trade_v1_trade_proto_goTypes = nil
	file_proto_aggregator_service_trade_v1_trade_proto_depIdxs = nil
}
