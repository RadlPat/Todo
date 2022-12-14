// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: dbs.protoc

package dbs

import (
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

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Todo   string `protobuf:"bytes,2,opt,name=Todo,proto3" json:"Todo,omitempty"`
	Active int32  `protobuf:"varint,3,opt,name=Active,proto3" json:"Active,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbs_protoc_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_dbs_protoc_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_dbs_protoc_rawDescGZIP(), []int{0}
}

func (x *Todo) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Todo) GetTodo() string {
	if x != nil {
		return x.Todo
	}
	return ""
}

func (x *Todo) GetActive() int32 {
	if x != nil {
		return x.Active
	}
	return 0
}

type TodoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoList []*Todo `protobuf:"bytes,1,rep,name=TodoList,proto3" json:"TodoList,omitempty"`
	Result   string  `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *TodoListResponse) Reset() {
	*x = TodoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbs_protoc_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListResponse) ProtoMessage() {}

func (x *TodoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dbs_protoc_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListResponse.ProtoReflect.Descriptor instead.
func (*TodoListResponse) Descriptor() ([]byte, []int) {
	return file_dbs_protoc_rawDescGZIP(), []int{1}
}

func (x *TodoListResponse) GetTodoList() []*Todo {
	if x != nil {
		return x.TodoList
	}
	return nil
}

func (x *TodoListResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type TodoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TodoListRequest) Reset() {
	*x = TodoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbs_protoc_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListRequest) ProtoMessage() {}

func (x *TodoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dbs_protoc_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListRequest.ProtoReflect.Descriptor instead.
func (*TodoListRequest) Descriptor() ([]byte, []int) {
	return file_dbs_protoc_rawDescGZIP(), []int{2}
}

var File_dbs_protoc protoreflect.FileDescriptor

var file_dbs_protoc_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x12, 0x03, 0x64, 0x62,
	0x73, 0x22, 0x42, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x6f, 0x64,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x51, 0x0a, 0x10, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x54, 0x6f, 0x64,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64, 0x62,
	0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x47, 0x0a, 0x09, 0x44,
	0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x6f, 0x44, 0x6f, 0x73, 0x12, 0x14, 0x2e, 0x64, 0x62, 0x73, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x64, 0x62, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x64, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dbs_protoc_rawDescOnce sync.Once
	file_dbs_protoc_rawDescData = file_dbs_protoc_rawDesc
)

func file_dbs_protoc_rawDescGZIP() []byte {
	file_dbs_protoc_rawDescOnce.Do(func() {
		file_dbs_protoc_rawDescData = protoimpl.X.CompressGZIP(file_dbs_protoc_rawDescData)
	})
	return file_dbs_protoc_rawDescData
}

var file_dbs_protoc_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dbs_protoc_goTypes = []interface{}{
	(*Todo)(nil),             // 0: dbs.Todo
	(*TodoListResponse)(nil), // 1: dbs.TodoListResponse
	(*TodoListRequest)(nil),  // 2: dbs.TodoListRequest
}
var file_dbs_protoc_depIdxs = []int32{
	0, // 0: dbs.TodoListResponse.TodoList:type_name -> dbs.Todo
	2, // 1: dbs.DbService.GetAllToDos:input_type -> dbs.TodoListRequest
	1, // 2: dbs.DbService.GetAllToDos:output_type -> dbs.TodoListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dbs_protoc_init() }
func file_dbs_protoc_init() {
	if File_dbs_protoc != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dbs_protoc_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_dbs_protoc_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoListResponse); i {
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
		file_dbs_protoc_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoListRequest); i {
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
			RawDescriptor: file_dbs_protoc_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dbs_protoc_goTypes,
		DependencyIndexes: file_dbs_protoc_depIdxs,
		MessageInfos:      file_dbs_protoc_msgTypes,
	}.Build()
	File_dbs_protoc = out.File
	file_dbs_protoc_rawDesc = nil
	file_dbs_protoc_goTypes = nil
	file_dbs_protoc_depIdxs = nil
}
