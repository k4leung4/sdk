// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: role_binding.platform.proto

package v1

import (
	_ "chainguard.dev/sdk/proto/annotations"
	v1 "chainguard.dev/sdk/proto/platform/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RoleBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id, the UID of this role binding.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// identity, UID of the Identity to bind.
	Identity string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	// group, UIDP of the group to bind. This field is ignored and will be removed
	// in the future. The group is always the parent of the UIDP.
	//
	// Deprecated: Do not use.
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	// role, UIDP of the Role to bind
	Role string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RoleBinding) Reset() {
	*x = RoleBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_binding_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBinding) ProtoMessage() {}

func (x *RoleBinding) ProtoReflect() protoreflect.Message {
	mi := &file_role_binding_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBinding.ProtoReflect.Descriptor instead.
func (*RoleBinding) Descriptor() ([]byte, []int) {
	return file_role_binding_platform_proto_rawDescGZIP(), []int{0}
}

func (x *RoleBinding) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleBinding) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

// Deprecated: Do not use.
func (x *RoleBinding) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *RoleBinding) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type RoleBindingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*RoleBindingList_Binding `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *RoleBindingList) Reset() {
	*x = RoleBindingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_binding_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBindingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBindingList) ProtoMessage() {}

func (x *RoleBindingList) ProtoReflect() protoreflect.Message {
	mi := &file_role_binding_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBindingList.ProtoReflect.Descriptor instead.
func (*RoleBindingList) Descriptor() ([]byte, []int) {
	return file_role_binding_platform_proto_rawDescGZIP(), []int{1}
}

func (x *RoleBindingList) GetItems() []*RoleBindingList_Binding {
	if x != nil {
		return x.Items
	}
	return nil
}

type RoleBindingFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the exact UID of the record.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// uidp filters records based on their position in the group hierarchy.
	Uidp *v1.UIDPFilter `protobuf:"bytes,2,opt,name=uidp,proto3" json:"uidp,omitempty"`
}

func (x *RoleBindingFilter) Reset() {
	*x = RoleBindingFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_binding_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBindingFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBindingFilter) ProtoMessage() {}

func (x *RoleBindingFilter) ProtoReflect() protoreflect.Message {
	mi := &file_role_binding_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBindingFilter.ProtoReflect.Descriptor instead.
func (*RoleBindingFilter) Descriptor() ([]byte, []int) {
	return file_role_binding_platform_proto_rawDescGZIP(), []int{2}
}

func (x *RoleBindingFilter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleBindingFilter) GetUidp() *v1.UIDPFilter {
	if x != nil {
		return x.Uidp
	}
	return nil
}

type CreateRoleBindingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// parent, The Group UIDP path under which the new RoleBinding resides.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// RoleBinding to create.
	RoleBinding *RoleBinding `protobuf:"bytes,2,opt,name=role_binding,json=roleBinding,proto3" json:"role_binding,omitempty"`
}

func (x *CreateRoleBindingRequest) Reset() {
	*x = CreateRoleBindingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_binding_platform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleBindingRequest) ProtoMessage() {}

func (x *CreateRoleBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_binding_platform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleBindingRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleBindingRequest) Descriptor() ([]byte, []int) {
	return file_role_binding_platform_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRoleBindingRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateRoleBindingRequest) GetRoleBinding() *RoleBinding {
	if x != nil {
		return x.RoleBinding
	}
	return nil
}

type DeleteRoleBindingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the exact UID of the record.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRoleBindingRequest) Reset() {
	*x = DeleteRoleBindingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_binding_platform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleBindingRequest) ProtoMessage() {}

func (x *DeleteRoleBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_binding_platform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleBindingRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleBindingRequest) Descriptor() ([]byte, []int) {
	return file_role_binding_platform_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRoleBindingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoleBindingList_Binding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id, the UID of this role binding.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// group of the bound role.
	Group *Group `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	// identity, UID of the Identity bound.
	Identity string `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	// role of the bound identity.
	Role *Role `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	// email of the bound identity.
	Email string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	// ClaimMatch issuer of the bound identity.
	ClaimMatchIssuer string `protobuf:"bytes,6,opt,name=claim_match_issuer,json=claimMatchIssuer,proto3" json:"claim_match_issuer,omitempty"`
	// ClaimMatch subject of the bound identity.
	ClaimMatchSubject string `protobuf:"bytes,7,opt,name=claim_match_subject,json=claimMatchSubject,proto3" json:"claim_match_subject,omitempty"`
}

func (x *RoleBindingList_Binding) Reset() {
	*x = RoleBindingList_Binding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_binding_platform_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBindingList_Binding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBindingList_Binding) ProtoMessage() {}

func (x *RoleBindingList_Binding) ProtoReflect() protoreflect.Message {
	mi := &file_role_binding_platform_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBindingList_Binding.ProtoReflect.Descriptor instead.
func (*RoleBindingList_Binding) Descriptor() ([]byte, []int) {
	return file_role_binding_platform_proto_rawDescGZIP(), []int{1, 0}
}

func (x *RoleBindingList_Binding) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleBindingList_Binding) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *RoleBindingList_Binding) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *RoleBindingList_Binding) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *RoleBindingList_Binding) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RoleBindingList_Binding) GetClaimMatchIssuer() string {
	if x != nil {
		return x.ClaimMatchIssuer
	}
	return ""
}

func (x *RoleBindingList_Binding) GetClaimMatchSubject() string {
	if x != nil {
		return x.ClaimMatchSubject
	}
	return ""
}

var File_role_binding_platform_proto protoreflect.FileDescriptor

var file_role_binding_platform_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0x90, 0xaf, 0xa8, 0xd2, 0x05, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0xee, 0x02, 0x0a, 0x0f, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x1a, 0x92, 0x02, 0x0a, 0x07, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x34, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x31, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5f, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a,
	0x04, 0x75, 0x69, 0x64, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x49, 0x44, 0x50, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x04, 0x75, 0x69, 0x64, 0x70, 0x22, 0x83, 0x01, 0x0a, 0x18, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x90, 0xaf, 0xa8, 0xd2, 0x05, 0x01, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22,
	0x32, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x90, 0xaf, 0xa8, 0xd2, 0x05, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x32, 0x93, 0x06, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0xe5, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x31, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x81, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x30, 0x22, 0x20, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x2a, 0x2a, 0x7d, 0x3a, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x8a, 0xaf, 0xa8, 0xd2, 0x05, 0x06, 0x12, 0x04, 0x0a, 0x02, 0x91, 0x03, 0xc2, 0xf0, 0x8e,
	0xfc, 0x0b, 0x39, 0x0a, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x12, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x12, 0xc8, 0x01, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x24, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x22, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x1a, 0x1c, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x2a, 0x7d, 0x3a, 0x01, 0x2a, 0x8a, 0xaf, 0xa8, 0xd2,
	0x05, 0x06, 0x12, 0x04, 0x0a, 0x02, 0x92, 0x03, 0xc2, 0xf0, 0x8e, 0xfc, 0x0b, 0x39, 0x0a, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x12, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x12, 0x88, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x28, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x8a, 0xaf, 0xa8, 0xd2, 0x05, 0x08, 0x12, 0x06, 0x0a, 0x02, 0x93, 0x03,
	0x10, 0x01, 0x12, 0xc4, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x31, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x2a, 0x1c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x2a, 0x7d, 0x8a, 0xaf,
	0xa8, 0xd2, 0x05, 0x06, 0x12, 0x04, 0x0a, 0x02, 0x94, 0x03, 0xc2, 0xf0, 0x8e, 0xfc, 0x0b, 0x39,
	0x0a, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x12, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x42, 0x6d, 0x0a, 0x22, 0x64, 0x65, 0x76,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42,
	0x1b, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x41, 0x4d, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_binding_platform_proto_rawDescOnce sync.Once
	file_role_binding_platform_proto_rawDescData = file_role_binding_platform_proto_rawDesc
)

func file_role_binding_platform_proto_rawDescGZIP() []byte {
	file_role_binding_platform_proto_rawDescOnce.Do(func() {
		file_role_binding_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_binding_platform_proto_rawDescData)
	})
	return file_role_binding_platform_proto_rawDescData
}

var file_role_binding_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_role_binding_platform_proto_goTypes = []interface{}{
	(*RoleBinding)(nil),              // 0: chainguard.platform.iam.RoleBinding
	(*RoleBindingList)(nil),          // 1: chainguard.platform.iam.RoleBindingList
	(*RoleBindingFilter)(nil),        // 2: chainguard.platform.iam.RoleBindingFilter
	(*CreateRoleBindingRequest)(nil), // 3: chainguard.platform.iam.CreateRoleBindingRequest
	(*DeleteRoleBindingRequest)(nil), // 4: chainguard.platform.iam.DeleteRoleBindingRequest
	(*RoleBindingList_Binding)(nil),  // 5: chainguard.platform.iam.RoleBindingList.Binding
	(*v1.UIDPFilter)(nil),            // 6: chainguard.platform.common.UIDPFilter
	(*Group)(nil),                    // 7: chainguard.platform.iam.Group
	(*Role)(nil),                     // 8: chainguard.platform.iam.Role
	(*emptypb.Empty)(nil),            // 9: google.protobuf.Empty
}
var file_role_binding_platform_proto_depIdxs = []int32{
	5, // 0: chainguard.platform.iam.RoleBindingList.items:type_name -> chainguard.platform.iam.RoleBindingList.Binding
	6, // 1: chainguard.platform.iam.RoleBindingFilter.uidp:type_name -> chainguard.platform.common.UIDPFilter
	0, // 2: chainguard.platform.iam.CreateRoleBindingRequest.role_binding:type_name -> chainguard.platform.iam.RoleBinding
	7, // 3: chainguard.platform.iam.RoleBindingList.Binding.group:type_name -> chainguard.platform.iam.Group
	8, // 4: chainguard.platform.iam.RoleBindingList.Binding.role:type_name -> chainguard.platform.iam.Role
	3, // 5: chainguard.platform.iam.RoleBindings.Create:input_type -> chainguard.platform.iam.CreateRoleBindingRequest
	0, // 6: chainguard.platform.iam.RoleBindings.Update:input_type -> chainguard.platform.iam.RoleBinding
	2, // 7: chainguard.platform.iam.RoleBindings.List:input_type -> chainguard.platform.iam.RoleBindingFilter
	4, // 8: chainguard.platform.iam.RoleBindings.Delete:input_type -> chainguard.platform.iam.DeleteRoleBindingRequest
	0, // 9: chainguard.platform.iam.RoleBindings.Create:output_type -> chainguard.platform.iam.RoleBinding
	0, // 10: chainguard.platform.iam.RoleBindings.Update:output_type -> chainguard.platform.iam.RoleBinding
	1, // 11: chainguard.platform.iam.RoleBindings.List:output_type -> chainguard.platform.iam.RoleBindingList
	9, // 12: chainguard.platform.iam.RoleBindings.Delete:output_type -> google.protobuf.Empty
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_role_binding_platform_proto_init() }
func file_role_binding_platform_proto_init() {
	if File_role_binding_platform_proto != nil {
		return
	}
	file_group_platform_proto_init()
	file_role_platform_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_role_binding_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBinding); i {
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
		file_role_binding_platform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBindingList); i {
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
		file_role_binding_platform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBindingFilter); i {
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
		file_role_binding_platform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleBindingRequest); i {
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
		file_role_binding_platform_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleBindingRequest); i {
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
		file_role_binding_platform_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBindingList_Binding); i {
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
			RawDescriptor: file_role_binding_platform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_binding_platform_proto_goTypes,
		DependencyIndexes: file_role_binding_platform_proto_depIdxs,
		MessageInfos:      file_role_binding_platform_proto_msgTypes,
	}.Build()
	File_role_binding_platform_proto = out.File
	file_role_binding_platform_proto_rawDesc = nil
	file_role_binding_platform_proto_goTypes = nil
	file_role_binding_platform_proto_depIdxs = nil
}
