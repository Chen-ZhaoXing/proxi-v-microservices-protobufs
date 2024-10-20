// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: checkout.proto

package proto

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

type CheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*RecipeOrderItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // Using a message defined in this proto
}

func (x *CheckoutRequest) Reset() {
	*x = CheckoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequest) ProtoMessage() {}

func (x *CheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequest.ProtoReflect.Descriptor instead.
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{0}
}

func (x *CheckoutRequest) GetItems() []*RecipeOrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type RecipeOrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipeId string `protobuf:"bytes,1,opt,name=recipe_id,json=recipeId,proto3" json:"recipe_id,omitempty"`
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"` // Quantity of the recipe being checked out
}

func (x *RecipeOrderItem) Reset() {
	*x = RecipeOrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeOrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeOrderItem) ProtoMessage() {}

func (x *RecipeOrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeOrderItem.ProtoReflect.Descriptor instead.
func (*RecipeOrderItem) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{1}
}

func (x *RecipeOrderItem) GetRecipeId() string {
	if x != nil {
		return x.RecipeId
	}
	return ""
}

func (x *RecipeOrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CheckoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CheckoutResponse) Reset() {
	*x = CheckoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutResponse) ProtoMessage() {}

func (x *CheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutResponse.ProtoReflect.Descriptor instead.
func (*CheckoutResponse) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{2}
}

func (x *CheckoutResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CheckoutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_checkout_proto protoreflect.FileDescriptor

var file_checkout_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x53, 0x47, 0x43, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x22, 0x4a, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x53, 0x47, 0x43, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x4a,
	0x0a, 0x0f, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x46, 0x0a, 0x10, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x64, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x12, 0x21, 0x2e, 0x53, 0x47, 0x43, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x53, 0x47, 0x43, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x65, 0x6e, 0x2d, 0x5a, 0x68, 0x61, 0x6f,
	0x58, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x2d, 0x76, 0x2d, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_checkout_proto_rawDescOnce sync.Once
	file_checkout_proto_rawDescData = file_checkout_proto_rawDesc
)

func file_checkout_proto_rawDescGZIP() []byte {
	file_checkout_proto_rawDescOnce.Do(func() {
		file_checkout_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkout_proto_rawDescData)
	})
	return file_checkout_proto_rawDescData
}

var file_checkout_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_checkout_proto_goTypes = []any{
	(*CheckoutRequest)(nil),  // 0: SGCooks.Checkout.CheckoutRequest
	(*RecipeOrderItem)(nil),  // 1: SGCooks.Checkout.RecipeOrderItem
	(*CheckoutResponse)(nil), // 2: SGCooks.Checkout.CheckoutResponse
}
var file_checkout_proto_depIdxs = []int32{
	1, // 0: SGCooks.Checkout.CheckoutRequest.items:type_name -> SGCooks.Checkout.RecipeOrderItem
	0, // 1: SGCooks.Checkout.CheckoutService.Checkout:input_type -> SGCooks.Checkout.CheckoutRequest
	2, // 2: SGCooks.Checkout.CheckoutService.Checkout:output_type -> SGCooks.Checkout.CheckoutResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_checkout_proto_init() }
func file_checkout_proto_init() {
	if File_checkout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkout_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CheckoutRequest); i {
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
		file_checkout_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RecipeOrderItem); i {
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
		file_checkout_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CheckoutResponse); i {
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
			RawDescriptor: file_checkout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checkout_proto_goTypes,
		DependencyIndexes: file_checkout_proto_depIdxs,
		MessageInfos:      file_checkout_proto_msgTypes,
	}.Build()
	File_checkout_proto = out.File
	file_checkout_proto_rawDesc = nil
	file_checkout_proto_goTypes = nil
	file_checkout_proto_depIdxs = nil
}
