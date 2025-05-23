// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: project.proto

package project_service_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IndexMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IndexMessage) Reset() {
	*x = IndexMessage{}
	mi := &file_project_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndexMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexMessage) ProtoMessage() {}

func (x *IndexMessage) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexMessage.ProtoReflect.Descriptor instead.
func (*IndexMessage) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{0}
}

func (x *IndexMessage) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type MenuMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pid           int64                  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Icon          string                 `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	Url           string                 `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	FilePath      string                 `protobuf:"bytes,6,opt,name=filePath,proto3" json:"filePath,omitempty"`
	Params        string                 `protobuf:"bytes,7,opt,name=params,proto3" json:"params,omitempty"`
	Node          string                 `protobuf:"bytes,8,opt,name=node,proto3" json:"node,omitempty"`
	Sort          int32                  `protobuf:"varint,9,opt,name=sort,proto3" json:"sort,omitempty"`
	Status        int32                  `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	CreateBy      int64                  `protobuf:"varint,11,opt,name=createBy,proto3" json:"createBy,omitempty"`
	IsInner       int32                  `protobuf:"varint,12,opt,name=isInner,proto3" json:"isInner,omitempty"`
	Values        string                 `protobuf:"bytes,13,opt,name=values,proto3" json:"values,omitempty"`
	ShowSlider    int32                  `protobuf:"varint,14,opt,name=showSlider,proto3" json:"showSlider,omitempty"`
	Children      []*MenuMessage         `protobuf:"bytes,15,rep,name=children,proto3" json:"children,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MenuMessage) Reset() {
	*x = MenuMessage{}
	mi := &file_project_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MenuMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuMessage) ProtoMessage() {}

func (x *MenuMessage) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuMessage.ProtoReflect.Descriptor instead.
func (*MenuMessage) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{1}
}

func (x *MenuMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MenuMessage) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *MenuMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MenuMessage) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *MenuMessage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MenuMessage) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *MenuMessage) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *MenuMessage) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *MenuMessage) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *MenuMessage) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MenuMessage) GetCreateBy() int64 {
	if x != nil {
		return x.CreateBy
	}
	return 0
}

func (x *MenuMessage) GetIsInner() int32 {
	if x != nil {
		return x.IsInner
	}
	return 0
}

func (x *MenuMessage) GetValues() string {
	if x != nil {
		return x.Values
	}
	return ""
}

func (x *MenuMessage) GetShowSlider() int32 {
	if x != nil {
		return x.ShowSlider
	}
	return 0
}

func (x *MenuMessage) GetChildren() []*MenuMessage {
	if x != nil {
		return x.Children
	}
	return nil
}

type IndexResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Menus         []*MenuMessage         `protobuf:"bytes,1,rep,name=menus,proto3" json:"menus,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IndexResponse) Reset() {
	*x = IndexResponse{}
	mi := &file_project_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexResponse) ProtoMessage() {}

func (x *IndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexResponse.ProtoReflect.Descriptor instead.
func (*IndexResponse) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{2}
}

func (x *IndexResponse) GetMenus() []*MenuMessage {
	if x != nil {
		return x.Menus
	}
	return nil
}

type ProjectMessage struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Id                 int64                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Cover              string                 `protobuf:"bytes,2,opt,name=Cover,proto3" json:"Cover,omitempty"`
	Name               string                 `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Description        string                 `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	AccessControlType  int32                  `protobuf:"varint,5,opt,name=AccessControlType,proto3" json:"AccessControlType,omitempty"`
	WhiteList          string                 `protobuf:"bytes,6,opt,name=WhiteList,proto3" json:"WhiteList,omitempty"`
	Order              int32                  `protobuf:"varint,7,opt,name=Order,proto3" json:"Order,omitempty"`
	Deleted            int32                  `protobuf:"varint,8,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	TemplateCode       string                 `protobuf:"bytes,9,opt,name=TemplateCode,proto3" json:"TemplateCode,omitempty"`
	Schedule           float64                `protobuf:"fixed64,10,opt,name=Schedule,proto3" json:"Schedule,omitempty"`
	CreateTime         string                 `protobuf:"bytes,11,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	OrganizationCode   int64                  `protobuf:"varint,12,opt,name=OrganizationCode,proto3" json:"OrganizationCode,omitempty"`
	DeletedTime        string                 `protobuf:"bytes,13,opt,name=DeletedTime,proto3" json:"DeletedTime,omitempty"`
	Private            int32                  `protobuf:"varint,14,opt,name=Private,proto3" json:"Private,omitempty"`
	Prefix             string                 `protobuf:"bytes,15,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	OpenPrefix         int32                  `protobuf:"varint,16,opt,name=OpenPrefix,proto3" json:"OpenPrefix,omitempty"`
	Archive            int32                  `protobuf:"varint,17,opt,name=Archive,proto3" json:"Archive,omitempty"`
	ArchiveTime        int64                  `protobuf:"varint,18,opt,name=ArchiveTime,proto3" json:"ArchiveTime,omitempty"`
	OpenBeginTime      int32                  `protobuf:"varint,19,opt,name=OpenBeginTime,proto3" json:"OpenBeginTime,omitempty"`
	OpenTaskPrivate    int32                  `protobuf:"varint,20,opt,name=OpenTaskPrivate,proto3" json:"OpenTaskPrivate,omitempty"`
	TaskBoardTheme     string                 `protobuf:"bytes,21,opt,name=TaskBoardTheme,proto3" json:"TaskBoardTheme,omitempty"`
	BeginTime          int64                  `protobuf:"varint,22,opt,name=BeginTime,proto3" json:"BeginTime,omitempty"`
	EndTime            int64                  `protobuf:"varint,23,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	AutoUpdateSchedule int32                  `protobuf:"varint,24,opt,name=AutoUpdateSchedule,proto3" json:"AutoUpdateSchedule,omitempty"`
	ProjectCode        int64                  `protobuf:"varint,25,opt,name=ProjectCode,proto3" json:"ProjectCode,omitempty"`
	MemberCode         int64                  `protobuf:"varint,26,opt,name=MemberCode,proto3" json:"MemberCode,omitempty"`
	JoinTime           int64                  `protobuf:"varint,27,opt,name=JoinTime,proto3" json:"JoinTime,omitempty"`
	IsOwner            int64                  `protobuf:"varint,28,opt,name=IsOwner,proto3" json:"IsOwner,omitempty"`
	Authorize          string                 `protobuf:"bytes,29,opt,name=Authorize,proto3" json:"Authorize,omitempty"`
	Code               string                 `protobuf:"bytes,30,opt,name=Code,proto3" json:"Code,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ProjectMessage) Reset() {
	*x = ProjectMessage{}
	mi := &file_project_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectMessage) ProtoMessage() {}

func (x *ProjectMessage) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectMessage.ProtoReflect.Descriptor instead.
func (*ProjectMessage) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectMessage) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *ProjectMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProjectMessage) GetAccessControlType() int32 {
	if x != nil {
		return x.AccessControlType
	}
	return 0
}

func (x *ProjectMessage) GetWhiteList() string {
	if x != nil {
		return x.WhiteList
	}
	return ""
}

func (x *ProjectMessage) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *ProjectMessage) GetDeleted() int32 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

func (x *ProjectMessage) GetTemplateCode() string {
	if x != nil {
		return x.TemplateCode
	}
	return ""
}

func (x *ProjectMessage) GetSchedule() float64 {
	if x != nil {
		return x.Schedule
	}
	return 0
}

func (x *ProjectMessage) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ProjectMessage) GetOrganizationCode() int64 {
	if x != nil {
		return x.OrganizationCode
	}
	return 0
}

func (x *ProjectMessage) GetDeletedTime() string {
	if x != nil {
		return x.DeletedTime
	}
	return ""
}

func (x *ProjectMessage) GetPrivate() int32 {
	if x != nil {
		return x.Private
	}
	return 0
}

func (x *ProjectMessage) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ProjectMessage) GetOpenPrefix() int32 {
	if x != nil {
		return x.OpenPrefix
	}
	return 0
}

func (x *ProjectMessage) GetArchive() int32 {
	if x != nil {
		return x.Archive
	}
	return 0
}

func (x *ProjectMessage) GetArchiveTime() int64 {
	if x != nil {
		return x.ArchiveTime
	}
	return 0
}

func (x *ProjectMessage) GetOpenBeginTime() int32 {
	if x != nil {
		return x.OpenBeginTime
	}
	return 0
}

func (x *ProjectMessage) GetOpenTaskPrivate() int32 {
	if x != nil {
		return x.OpenTaskPrivate
	}
	return 0
}

func (x *ProjectMessage) GetTaskBoardTheme() string {
	if x != nil {
		return x.TaskBoardTheme
	}
	return ""
}

func (x *ProjectMessage) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *ProjectMessage) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ProjectMessage) GetAutoUpdateSchedule() int32 {
	if x != nil {
		return x.AutoUpdateSchedule
	}
	return 0
}

func (x *ProjectMessage) GetProjectCode() int64 {
	if x != nil {
		return x.ProjectCode
	}
	return 0
}

func (x *ProjectMessage) GetMemberCode() int64 {
	if x != nil {
		return x.MemberCode
	}
	return 0
}

func (x *ProjectMessage) GetJoinTime() int64 {
	if x != nil {
		return x.JoinTime
	}
	return 0
}

func (x *ProjectMessage) GetIsOwner() int64 {
	if x != nil {
		return x.IsOwner
	}
	return 0
}

func (x *ProjectMessage) GetAuthorize() string {
	if x != nil {
		return x.Authorize
	}
	return ""
}

func (x *ProjectMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ProjectRpcMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MemberId      int64                  `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	Page          int64                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int64                  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectRpcMessage) Reset() {
	*x = ProjectRpcMessage{}
	mi := &file_project_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectRpcMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectRpcMessage) ProtoMessage() {}

func (x *ProjectRpcMessage) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectRpcMessage.ProtoReflect.Descriptor instead.
func (*ProjectRpcMessage) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectRpcMessage) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *ProjectRpcMessage) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ProjectRpcMessage) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type MyProjectResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pm            []*ProjectMessage      `protobuf:"bytes,1,rep,name=pm,proto3" json:"pm,omitempty"`
	Total         int64                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MyProjectResponse) Reset() {
	*x = MyProjectResponse{}
	mi := &file_project_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MyProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyProjectResponse) ProtoMessage() {}

func (x *MyProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyProjectResponse.ProtoReflect.Descriptor instead.
func (*MyProjectResponse) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{5}
}

func (x *MyProjectResponse) GetPm() []*ProjectMessage {
	if x != nil {
		return x.Pm
	}
	return nil
}

func (x *MyProjectResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_project_proto protoreflect.FileDescriptor

var file_project_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0x24, 0x0a, 0x0c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8a, 0x03, 0x0a, 0x0b, 0x4d, 0x65,
	0x6e, 0x75, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x73, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68, 0x6f,
	0x77, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x68, 0x6f, 0x77, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x0d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x22, 0xae,
	0x07, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a,
	0x11, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x57,
	0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x70, 0x65,
	0x6e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4f,
	0x70, 0x65, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x6e, 0x42, 0x65, 0x67,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x4f, 0x70,
	0x65, 0x6e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x4f,
	0x70, 0x65, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x54,
	0x61, 0x73, 0x6b, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x1c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x49, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x5f, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x70, 0x63, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x5d, 0x0a, 0x11, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x02, 0x70, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x70, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32,
	0xc6, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x64, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x70, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x25, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_project_proto_rawDescOnce sync.Once
	file_project_proto_rawDescData []byte
)

func file_project_proto_rawDescGZIP() []byte {
	file_project_proto_rawDescOnce.Do(func() {
		file_project_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_proto_rawDesc), len(file_project_proto_rawDesc)))
	})
	return file_project_proto_rawDescData
}

var file_project_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_project_proto_goTypes = []any{
	(*IndexMessage)(nil),      // 0: project.service.v1.IndexMessage
	(*MenuMessage)(nil),       // 1: project.service.v1.MenuMessage
	(*IndexResponse)(nil),     // 2: project.service.v1.IndexResponse
	(*ProjectMessage)(nil),    // 3: project.service.v1.ProjectMessage
	(*ProjectRpcMessage)(nil), // 4: project.service.v1.ProjectRpcMessage
	(*MyProjectResponse)(nil), // 5: project.service.v1.MyProjectResponse
}
var file_project_proto_depIdxs = []int32{
	1, // 0: project.service.v1.MenuMessage.children:type_name -> project.service.v1.MenuMessage
	1, // 1: project.service.v1.IndexResponse.menus:type_name -> project.service.v1.MenuMessage
	3, // 2: project.service.v1.MyProjectResponse.pm:type_name -> project.service.v1.ProjectMessage
	0, // 3: project.service.v1.ProjectService.Index:input_type -> project.service.v1.IndexMessage
	4, // 4: project.service.v1.ProjectService.FindProjectByMemId:input_type -> project.service.v1.ProjectRpcMessage
	2, // 5: project.service.v1.ProjectService.Index:output_type -> project.service.v1.IndexResponse
	5, // 6: project.service.v1.ProjectService.FindProjectByMemId:output_type -> project.service.v1.MyProjectResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_project_proto_init() }
func file_project_proto_init() {
	if File_project_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_proto_rawDesc), len(file_project_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_proto_goTypes,
		DependencyIndexes: file_project_proto_depIdxs,
		MessageInfos:      file_project_proto_msgTypes,
	}.Build()
	File_project_proto = out.File
	file_project_proto_goTypes = nil
	file_project_proto_depIdxs = nil
}
