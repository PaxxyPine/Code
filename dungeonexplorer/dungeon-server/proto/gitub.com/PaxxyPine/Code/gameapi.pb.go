// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: gameapi.proto

package Code

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

type Elements int32

const (
	Elements_Elements_NOT_SET Elements = 0
	Elements_Elements_FIRE    Elements = 1
	Elements_Elements_WIND    Elements = 2
	Elements_Elements_WATER   Elements = 3
	Elements_Elements_WOOD    Elements = 4
)

// Enum value maps for Elements.
var (
	Elements_name = map[int32]string{
		0: "Elements_NOT_SET",
		1: "Elements_FIRE",
		2: "Elements_WIND",
		3: "Elements_WATER",
		4: "Elements_WOOD",
	}
	Elements_value = map[string]int32{
		"Elements_NOT_SET": 0,
		"Elements_FIRE":    1,
		"Elements_WIND":    2,
		"Elements_WATER":   3,
		"Elements_WOOD":    4,
	}
)

func (x Elements) Enum() *Elements {
	p := new(Elements)
	*p = x
	return p
}

func (x Elements) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Elements) Descriptor() protoreflect.EnumDescriptor {
	return file_gameapi_proto_enumTypes[0].Descriptor()
}

func (Elements) Type() protoreflect.EnumType {
	return &file_gameapi_proto_enumTypes[0]
}

func (x Elements) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Elements.Descriptor instead.
func (Elements) EnumDescriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{0}
}

type ArmorType int32

const (
	ArmorType_ArmorType_NOT_SET   ArmorType = 0
	ArmorType_ArmorType_HELM      ArmorType = 1
	ArmorType_ArmorType_SHOULDERS ArmorType = 2
	ArmorType_ArmorType_CHEST     ArmorType = 3
	ArmorType_ArmorType_GOVES     ArmorType = 4
	ArmorType_ArmorType_BELT      ArmorType = 5
	ArmorType_ArmorType_LEGS      ArmorType = 6
	ArmorType_ArmorType_BOOTS     ArmorType = 7
	ArmorType_ArmorType_HELD      ArmorType = 8
	ArmorType_ArmorType_NECK      ArmorType = 9
	ArmorType_ArmorType_FINGERS   ArmorType = 10
)

// Enum value maps for ArmorType.
var (
	ArmorType_name = map[int32]string{
		0:  "ArmorType_NOT_SET",
		1:  "ArmorType_HELM",
		2:  "ArmorType_SHOULDERS",
		3:  "ArmorType_CHEST",
		4:  "ArmorType_GOVES",
		5:  "ArmorType_BELT",
		6:  "ArmorType_LEGS",
		7:  "ArmorType_BOOTS",
		8:  "ArmorType_HELD",
		9:  "ArmorType_NECK",
		10: "ArmorType_FINGERS",
	}
	ArmorType_value = map[string]int32{
		"ArmorType_NOT_SET":   0,
		"ArmorType_HELM":      1,
		"ArmorType_SHOULDERS": 2,
		"ArmorType_CHEST":     3,
		"ArmorType_GOVES":     4,
		"ArmorType_BELT":      5,
		"ArmorType_LEGS":      6,
		"ArmorType_BOOTS":     7,
		"ArmorType_HELD":      8,
		"ArmorType_NECK":      9,
		"ArmorType_FINGERS":   10,
	}
)

func (x ArmorType) Enum() *ArmorType {
	p := new(ArmorType)
	*p = x
	return p
}

func (x ArmorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArmorType) Descriptor() protoreflect.EnumDescriptor {
	return file_gameapi_proto_enumTypes[1].Descriptor()
}

func (ArmorType) Type() protoreflect.EnumType {
	return &file_gameapi_proto_enumTypes[1]
}

func (x ArmorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArmorType.Descriptor instead.
func (ArmorType) EnumDescriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{1}
}

type ArmorSlot int32

const (
	ArmorSlot_ArmorSlot_NOT_SET       ArmorSlot = 0
	ArmorSlot_ArmorSlot_HEAD          ArmorSlot = 1
	ArmorSlot_ArmorSlot_SHOULDERS     ArmorSlot = 2
	ArmorSlot_ArmorSlot_CHEST         ArmorSlot = 3
	ArmorSlot_ArmorSlot_GLOVES        ArmorSlot = 4
	ArmorSlot_ArmorSlot_WAIST         ArmorSlot = 5
	ArmorSlot_ArmorSlot_LEGS          ArmorSlot = 6
	ArmorSlot_ArmorSlot_FEET          ArmorSlot = 7
	ArmorSlot_ArmorSlot_LEFT_HAND     ArmorSlot = 8
	ArmorSlot_ArmorSlot_NECK          ArmorSlot = 9
	ArmorSlot_ArmorSlot_LEFT_FINGERS  ArmorSlot = 10
	ArmorSlot_ArmorSlot_RIGHT_FINGERS ArmorSlot = 11
)

// Enum value maps for ArmorSlot.
var (
	ArmorSlot_name = map[int32]string{
		0:  "ArmorSlot_NOT_SET",
		1:  "ArmorSlot_HEAD",
		2:  "ArmorSlot_SHOULDERS",
		3:  "ArmorSlot_CHEST",
		4:  "ArmorSlot_GLOVES",
		5:  "ArmorSlot_WAIST",
		6:  "ArmorSlot_LEGS",
		7:  "ArmorSlot_FEET",
		8:  "ArmorSlot_LEFT_HAND",
		9:  "ArmorSlot_NECK",
		10: "ArmorSlot_LEFT_FINGERS",
		11: "ArmorSlot_RIGHT_FINGERS",
	}
	ArmorSlot_value = map[string]int32{
		"ArmorSlot_NOT_SET":       0,
		"ArmorSlot_HEAD":          1,
		"ArmorSlot_SHOULDERS":     2,
		"ArmorSlot_CHEST":         3,
		"ArmorSlot_GLOVES":        4,
		"ArmorSlot_WAIST":         5,
		"ArmorSlot_LEGS":          6,
		"ArmorSlot_FEET":          7,
		"ArmorSlot_LEFT_HAND":     8,
		"ArmorSlot_NECK":          9,
		"ArmorSlot_LEFT_FINGERS":  10,
		"ArmorSlot_RIGHT_FINGERS": 11,
	}
)

func (x ArmorSlot) Enum() *ArmorSlot {
	p := new(ArmorSlot)
	*p = x
	return p
}

func (x ArmorSlot) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArmorSlot) Descriptor() protoreflect.EnumDescriptor {
	return file_gameapi_proto_enumTypes[2].Descriptor()
}

func (ArmorSlot) Type() protoreflect.EnumType {
	return &file_gameapi_proto_enumTypes[2]
}

func (x ArmorSlot) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArmorSlot.Descriptor instead.
func (ArmorSlot) EnumDescriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{2}
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gameapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{0}
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gameapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{1}
}

var File_gameapi_proto protoreflect.FileDescriptor

var file_gameapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x10, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2a, 0x6d, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x46,
	0x49, 0x52, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x57, 0x41, 0x54, 0x45, 0x52, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x57, 0x4f, 0x4f, 0x44, 0x10, 0x04, 0x2a,
	0xf5, 0x01, 0x0a, 0x09, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a,
	0x11, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x45, 0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x5f, 0x48, 0x45, 0x4c, 0x4d, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x72, 0x6d, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x53, 0x48, 0x4f, 0x55, 0x4c, 0x44, 0x45, 0x52, 0x53, 0x10,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x43,
	0x48, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x47, 0x4f, 0x56, 0x45, 0x53, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x41,
	0x72, 0x6d, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x42, 0x45, 0x4c, 0x54, 0x10, 0x05, 0x12,
	0x12, 0x0a, 0x0e, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4c, 0x45, 0x47,
	0x53, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x42, 0x4f, 0x4f, 0x54, 0x53, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x72, 0x6d, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x48, 0x45, 0x4c, 0x44, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e,
	0x41, 0x72, 0x6d, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e, 0x45, 0x43, 0x4b, 0x10, 0x09,
	0x12, 0x15, 0x0a, 0x11, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x46, 0x49,
	0x4e, 0x47, 0x45, 0x52, 0x53, 0x10, 0x0a, 0x2a, 0x9d, 0x02, 0x0a, 0x09, 0x41, 0x72, 0x6d, 0x6f,
	0x72, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x53, 0x6c,
	0x6f, 0x74, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x41, 0x72, 0x6d, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x10, 0x01,
	0x12, 0x17, 0x0a, 0x13, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x53, 0x48,
	0x4f, 0x55, 0x4c, 0x44, 0x45, 0x52, 0x53, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x43, 0x48, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x14,
	0x0a, 0x10, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x47, 0x4c, 0x4f, 0x56,
	0x45, 0x53, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x53, 0x6c, 0x6f,
	0x74, 0x5f, 0x57, 0x41, 0x49, 0x53, 0x54, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x72, 0x6d,
	0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x4c, 0x45, 0x47, 0x53, 0x10, 0x06, 0x12, 0x12, 0x0a,
	0x0e, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x46, 0x45, 0x45, 0x54, 0x10,
	0x07, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x4c,
	0x45, 0x46, 0x54, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x72,
	0x6d, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x4e, 0x45, 0x43, 0x4b, 0x10, 0x09, 0x12, 0x1a,
	0x0a, 0x16, 0x41, 0x72, 0x6d, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x4c, 0x45, 0x46, 0x54,
	0x5f, 0x46, 0x49, 0x4e, 0x47, 0x45, 0x52, 0x53, 0x10, 0x0a, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x72,
	0x6d, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x46, 0x49,
	0x4e, 0x47, 0x45, 0x52, 0x53, 0x10, 0x0b, 0x32, 0x6f, 0x0a, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x41,
	0x50, 0x49, 0x12, 0x64, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x2e,
	0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x64, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x78, 0x78, 0x79, 0x50, 0x69, 0x6e, 0x65, 0x2f,
	0x43, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gameapi_proto_rawDescOnce sync.Once
	file_gameapi_proto_rawDescData = file_gameapi_proto_rawDesc
)

func file_gameapi_proto_rawDescGZIP() []byte {
	file_gameapi_proto_rawDescOnce.Do(func() {
		file_gameapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_gameapi_proto_rawDescData)
	})
	return file_gameapi_proto_rawDescData
}

var file_gameapi_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_gameapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gameapi_proto_goTypes = []interface{}{
	(Elements)(0),           // 0: dungeonexplorer.gameapi.v1.Elements
	(ArmorType)(0),          // 1: dungeonexplorer.gameapi.v1.ArmorType
	(ArmorSlot)(0),          // 2: dungeonexplorer.gameapi.v1.ArmorSlot
	(*GetInfoRequest)(nil),  // 3: dungeonexplorer.gameapi.v1.GetInfoRequest
	(*GetInfoResponse)(nil), // 4: dungeonexplorer.gameapi.v1.GetInfoResponse
}
var file_gameapi_proto_depIdxs = []int32{
	3, // 0: dungeonexplorer.gameapi.v1.GameAPI.GetInfo:input_type -> dungeonexplorer.gameapi.v1.GetInfoRequest
	4, // 1: dungeonexplorer.gameapi.v1.GameAPI.GetInfo:output_type -> dungeonexplorer.gameapi.v1.GetInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gameapi_proto_init() }
func file_gameapi_proto_init() {
	if File_gameapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gameapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoRequest); i {
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
		file_gameapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse); i {
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
			RawDescriptor: file_gameapi_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gameapi_proto_goTypes,
		DependencyIndexes: file_gameapi_proto_depIdxs,
		EnumInfos:         file_gameapi_proto_enumTypes,
		MessageInfos:      file_gameapi_proto_msgTypes,
	}.Build()
	File_gameapi_proto = out.File
	file_gameapi_proto_rawDesc = nil
	file_gameapi_proto_goTypes = nil
	file_gameapi_proto_depIdxs = nil
}
