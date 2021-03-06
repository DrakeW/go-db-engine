// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.13.0
// source: sstable.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SSTableBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*SSTableKeyValue `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SSTableBlock) Reset() {
	*x = SSTableBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sstable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSTableBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSTableBlock) ProtoMessage() {}

func (x *SSTableBlock) ProtoReflect() protoreflect.Message {
	mi := &file_sstable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSTableBlock.ProtoReflect.Descriptor instead.
func (*SSTableBlock) Descriptor() ([]byte, []int) {
	return file_sstable_proto_rawDescGZIP(), []int{0}
}

func (x *SSTableBlock) GetData() []*SSTableKeyValue {
	if x != nil {
		return x.Data
	}
	return nil
}

type SSTableKeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SSTableKeyValue) Reset() {
	*x = SSTableKeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sstable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSTableKeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSTableKeyValue) ProtoMessage() {}

func (x *SSTableKeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_sstable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSTableKeyValue.ProtoReflect.Descriptor instead.
func (*SSTableKeyValue) Descriptor() ([]byte, []int) {
	return file_sstable_proto_rawDescGZIP(), []int{1}
}

func (x *SSTableKeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SSTableKeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SSTableIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*SSTableIndexEntry `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SSTableIndex) Reset() {
	*x = SSTableIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sstable_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSTableIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSTableIndex) ProtoMessage() {}

func (x *SSTableIndex) ProtoReflect() protoreflect.Message {
	mi := &file_sstable_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSTableIndex.ProtoReflect.Descriptor instead.
func (*SSTableIndex) Descriptor() ([]byte, []int) {
	return file_sstable_proto_rawDescGZIP(), []int{2}
}

func (x *SSTableIndex) GetData() []*SSTableIndexEntry {
	if x != nil {
		return x.Data
	}
	return nil
}

type SSTableIndexEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartKey string `protobuf:"bytes,1,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	EndKey   string `protobuf:"bytes,2,opt,name=end_key,json=endKey,proto3" json:"end_key,omitempty"`
	Offset   uint64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Size     uint64 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SSTableIndexEntry) Reset() {
	*x = SSTableIndexEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sstable_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSTableIndexEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSTableIndexEntry) ProtoMessage() {}

func (x *SSTableIndexEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sstable_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSTableIndexEntry.ProtoReflect.Descriptor instead.
func (*SSTableIndexEntry) Descriptor() ([]byte, []int) {
	return file_sstable_proto_rawDescGZIP(), []int{3}
}

func (x *SSTableIndexEntry) GetStartKey() string {
	if x != nil {
		return x.StartKey
	}
	return ""
}

func (x *SSTableIndexEntry) GetEndKey() string {
	if x != nil {
		return x.EndKey
	}
	return ""
}

func (x *SSTableIndexEntry) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SSTableIndexEntry) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_sstable_proto protoreflect.FileDescriptor

var file_sstable_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x34, 0x0a, 0x0c, 0x53, 0x53, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x53, 0x53, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x0f, 0x53, 0x53, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x36, 0x0a, 0x0c, 0x53, 0x53, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x53, 0x53, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x75, 0x0a, 0x11, 0x53, 0x53, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e,
	0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x64,
	0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x42,
	0x04, 0x5a, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sstable_proto_rawDescOnce sync.Once
	file_sstable_proto_rawDescData = file_sstable_proto_rawDesc
)

func file_sstable_proto_rawDescGZIP() []byte {
	file_sstable_proto_rawDescOnce.Do(func() {
		file_sstable_proto_rawDescData = protoimpl.X.CompressGZIP(file_sstable_proto_rawDescData)
	})
	return file_sstable_proto_rawDescData
}

var file_sstable_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sstable_proto_goTypes = []interface{}{
	(*SSTableBlock)(nil),      // 0: SSTableBlock
	(*SSTableKeyValue)(nil),   // 1: SSTableKeyValue
	(*SSTableIndex)(nil),      // 2: SSTableIndex
	(*SSTableIndexEntry)(nil), // 3: SSTableIndexEntry
}
var file_sstable_proto_depIdxs = []int32{
	1, // 0: SSTableBlock.data:type_name -> SSTableKeyValue
	3, // 1: SSTableIndex.data:type_name -> SSTableIndexEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sstable_proto_init() }
func file_sstable_proto_init() {
	if File_sstable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sstable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSTableBlock); i {
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
		file_sstable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSTableKeyValue); i {
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
		file_sstable_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSTableIndex); i {
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
		file_sstable_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSTableIndexEntry); i {
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
			RawDescriptor: file_sstable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sstable_proto_goTypes,
		DependencyIndexes: file_sstable_proto_depIdxs,
		MessageInfos:      file_sstable_proto_msgTypes,
	}.Build()
	File_sstable_proto = out.File
	file_sstable_proto_rawDesc = nil
	file_sstable_proto_goTypes = nil
	file_sstable_proto_depIdxs = nil
}
