// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: src/idl/auth_service/staff_auth.proto

package auth_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_auth_service_staff_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_auth_service_staff_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_src_idl_auth_service_staff_auth_proto_rawDescGZIP(), []int{0}
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role     string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	BranchId string `protobuf:"bytes,3,opt,name=branchId,proto3" json:"branchId,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_auth_service_staff_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_auth_service_staff_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_src_idl_auth_service_staff_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *AuthResponse) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type StaffLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
}

func (x *StaffLoginRequest) Reset() {
	*x = StaffLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_auth_service_staff_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffLoginRequest) ProtoMessage() {}

func (x *StaffLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_auth_service_staff_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffLoginRequest.ProtoReflect.Descriptor instead.
func (*StaffLoginRequest) Descriptor() ([]byte, []int) {
	return file_src_idl_auth_service_staff_auth_proto_rawDescGZIP(), []int{2}
}

func (x *StaffLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *StaffLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *StaffLoginRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type StaffSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Role        string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	BranchId    string `protobuf:"bytes,6,opt,name=branchId,proto3" json:"branchId,omitempty"`
}

func (x *StaffSignUpRequest) Reset() {
	*x = StaffSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_auth_service_staff_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffSignUpRequest) ProtoMessage() {}

func (x *StaffSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_auth_service_staff_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffSignUpRequest.ProtoReflect.Descriptor instead.
func (*StaffSignUpRequest) Descriptor() ([]byte, []int) {
	return file_src_idl_auth_service_staff_auth_proto_rawDescGZIP(), []int{3}
}

func (x *StaffSignUpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StaffSignUpRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *StaffSignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *StaffSignUpRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *StaffSignUpRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *StaffSignUpRequest) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

type StaffResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID     string `protobuf:"bytes,1,opt,name=staffID,proto3" json:"staffID,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
}

func (x *StaffResetPasswordRequest) Reset() {
	*x = StaffResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_idl_auth_service_staff_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffResetPasswordRequest) ProtoMessage() {}

func (x *StaffResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_idl_auth_service_staff_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*StaffResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_src_idl_auth_service_staff_auth_proto_rawDescGZIP(), []int{4}
}

func (x *StaffResetPasswordRequest) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

func (x *StaffResetPasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

var File_src_idl_auth_service_staff_auth_proto protoreflect.FileDescriptor

var file_src_idl_auth_service_staff_auth_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x72, 0x63, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4e,
	0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0x6d,
	0x0a, 0x11, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xb2, 0x01,
	0x0a, 0x12, 0x53, 0x74, 0x61, 0x66, 0x66, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x49, 0x64, 0x22, 0x57, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0xff, 0x01, 0x0a, 0x0b,
	0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_idl_auth_service_staff_auth_proto_rawDescOnce sync.Once
	file_src_idl_auth_service_staff_auth_proto_rawDescData = file_src_idl_auth_service_staff_auth_proto_rawDesc
)

func file_src_idl_auth_service_staff_auth_proto_rawDescGZIP() []byte {
	file_src_idl_auth_service_staff_auth_proto_rawDescOnce.Do(func() {
		file_src_idl_auth_service_staff_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_idl_auth_service_staff_auth_proto_rawDescData)
	})
	return file_src_idl_auth_service_staff_auth_proto_rawDescData
}

var file_src_idl_auth_service_staff_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_src_idl_auth_service_staff_auth_proto_goTypes = []interface{}{
	(*Empty)(nil),                     // 0: auth_service.Empty
	(*AuthResponse)(nil),              // 1: auth_service.AuthResponse
	(*StaffLoginRequest)(nil),         // 2: auth_service.StaffLoginRequest
	(*StaffSignUpRequest)(nil),        // 3: auth_service.StaffSignUpRequest
	(*StaffResetPasswordRequest)(nil), // 4: auth_service.StaffResetPasswordRequest
}
var file_src_idl_auth_service_staff_auth_proto_depIdxs = []int32{
	2, // 0: auth_service.AuthService.StaffLogin:input_type -> auth_service.StaffLoginRequest
	3, // 1: auth_service.AuthService.StaffSignUp:input_type -> auth_service.StaffSignUpRequest
	4, // 2: auth_service.AuthService.StaffResetPassword:input_type -> auth_service.StaffResetPasswordRequest
	1, // 3: auth_service.AuthService.StaffLogin:output_type -> auth_service.AuthResponse
	1, // 4: auth_service.AuthService.StaffSignUp:output_type -> auth_service.AuthResponse
	0, // 5: auth_service.AuthService.StaffResetPassword:output_type -> auth_service.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_idl_auth_service_staff_auth_proto_init() }
func file_src_idl_auth_service_staff_auth_proto_init() {
	if File_src_idl_auth_service_staff_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_idl_auth_service_staff_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_src_idl_auth_service_staff_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_src_idl_auth_service_staff_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffLoginRequest); i {
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
		file_src_idl_auth_service_staff_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffSignUpRequest); i {
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
		file_src_idl_auth_service_staff_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffResetPasswordRequest); i {
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
			RawDescriptor: file_src_idl_auth_service_staff_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_idl_auth_service_staff_auth_proto_goTypes,
		DependencyIndexes: file_src_idl_auth_service_staff_auth_proto_depIdxs,
		MessageInfos:      file_src_idl_auth_service_staff_auth_proto_msgTypes,
	}.Build()
	File_src_idl_auth_service_staff_auth_proto = out.File
	file_src_idl_auth_service_staff_auth_proto_rawDesc = nil
	file_src_idl_auth_service_staff_auth_proto_goTypes = nil
	file_src_idl_auth_service_staff_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	StaffLogin(ctx context.Context, in *StaffLoginRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	StaffSignUp(ctx context.Context, in *StaffSignUpRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	StaffResetPassword(ctx context.Context, in *StaffResetPasswordRequest, opts ...grpc.CallOption) (*Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) StaffLogin(ctx context.Context, in *StaffLoginRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/StaffLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) StaffSignUp(ctx context.Context, in *StaffSignUpRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/StaffSignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) StaffResetPassword(ctx context.Context, in *StaffResetPasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/auth_service.AuthService/StaffResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	StaffLogin(context.Context, *StaffLoginRequest) (*AuthResponse, error)
	StaffSignUp(context.Context, *StaffSignUpRequest) (*AuthResponse, error)
	StaffResetPassword(context.Context, *StaffResetPasswordRequest) (*Empty, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) StaffLogin(context.Context, *StaffLoginRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffLogin not implemented")
}
func (*UnimplementedAuthServiceServer) StaffSignUp(context.Context, *StaffSignUpRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffSignUp not implemented")
}
func (*UnimplementedAuthServiceServer) StaffResetPassword(context.Context, *StaffResetPasswordRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffResetPassword not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_StaffLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).StaffLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/StaffLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).StaffLogin(ctx, req.(*StaffLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_StaffSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffSignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).StaffSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/StaffSignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).StaffSignUp(ctx, req.(*StaffSignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_StaffResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).StaffResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.AuthService/StaffResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).StaffResetPassword(ctx, req.(*StaffResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StaffLogin",
			Handler:    _AuthService_StaffLogin_Handler,
		},
		{
			MethodName: "StaffSignUp",
			Handler:    _AuthService_StaffSignUp_Handler,
		},
		{
			MethodName: "StaffResetPassword",
			Handler:    _AuthService_StaffResetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/idl/auth_service/staff_auth.proto",
}
