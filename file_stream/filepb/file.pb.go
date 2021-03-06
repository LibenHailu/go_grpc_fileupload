// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: filepb/file.proto

package filepb

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

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UploadFileRequest_Info
	//	*UploadFileRequest_ChunkData
	Data isUploadFileRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filepb_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filepb_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_filepb_file_proto_rawDescGZIP(), []int{0}
}

func (m *UploadFileRequest) GetData() isUploadFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadFileRequest) GetInfo() *FileInfo {
	if x, ok := x.GetData().(*UploadFileRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadFileRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*UploadFileRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isUploadFileRequest_Data interface {
	isUploadFileRequest_Data()
}

type UploadFileRequest_Info struct {
	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadFileRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadFileRequest_Info) isUploadFileRequest_Data() {}

func (*UploadFileRequest_ChunkData) isUploadFileRequest_Data() {}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId   string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileType string `protobuf:"bytes,3,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filepb_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_filepb_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_filepb_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *FileInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileInfo) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filepb_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_filepb_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_filepb_file_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadFileResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ServeFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkData []byte `protobuf:"bytes,1,opt,name=chunk_data,json=chunkData,proto3" json:"chunk_data,omitempty"`
}

func (x *ServeFileResponse) Reset() {
	*x = ServeFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filepb_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServeFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServeFileResponse) ProtoMessage() {}

func (x *ServeFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_filepb_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServeFileResponse.ProtoReflect.Descriptor instead.
func (*ServeFileResponse) Descriptor() ([]byte, []int) {
	return file_filepb_file_proto_rawDescGZIP(), []int{3}
}

func (x *ServeFileResponse) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

type ServeFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *ServeFileRequest) Reset() {
	*x = ServeFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filepb_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServeFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServeFileRequest) ProtoMessage() {}

func (x *ServeFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filepb_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServeFileRequest.ProtoReflect.Descriptor instead.
func (*ServeFileRequest) Descriptor() ([]byte, []int) {
	return file_filepb_file_proto_rawDescGZIP(), []int{4}
}

func (x *ServeFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type RegisterPeersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip        string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port      int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	FileNames []string `protobuf:"bytes,3,rep,name=fileNames,proto3" json:"fileNames,omitempty"`
}

func (x *RegisterPeersRequest) Reset() {
	*x = RegisterPeersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filepb_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPeersRequest) ProtoMessage() {}

func (x *RegisterPeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filepb_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPeersRequest.ProtoReflect.Descriptor instead.
func (*RegisterPeersRequest) Descriptor() ([]byte, []int) {
	return file_filepb_file_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterPeersRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RegisterPeersRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RegisterPeersRequest) GetFileNames() []string {
	if x != nil {
		return x.FileNames
	}
	return nil
}

type RegisterPeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerAddress string `protobuf:"bytes,1,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
}

func (x *RegisterPeersResponse) Reset() {
	*x = RegisterPeersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filepb_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPeersResponse) ProtoMessage() {}

func (x *RegisterPeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_filepb_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPeersResponse.ProtoReflect.Descriptor instead.
func (*RegisterPeersResponse) Descriptor() ([]byte, []int) {
	return file_filepb_file_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterPeersResponse) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

var File_filepb_file_proto protoreflect.FileDescriptor

var file_filepb_file_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x22, 0x69, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5d, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x32, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x14, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x32, 0x8d, 0x02, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x1e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x51, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x58, 0x0a, 0x0d, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x21, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_filepb_file_proto_rawDescOnce sync.Once
	file_filepb_file_proto_rawDescData = file_filepb_file_proto_rawDesc
)

func file_filepb_file_proto_rawDescGZIP() []byte {
	file_filepb_file_proto_rawDescOnce.Do(func() {
		file_filepb_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_filepb_file_proto_rawDescData)
	})
	return file_filepb_file_proto_rawDescData
}

var file_filepb_file_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_filepb_file_proto_goTypes = []interface{}{
	(*UploadFileRequest)(nil),     // 0: file_stream.UploadFileRequest
	(*FileInfo)(nil),              // 1: file_stream.FileInfo
	(*UploadFileResponse)(nil),    // 2: file_stream.UploadFileResponse
	(*ServeFileResponse)(nil),     // 3: file_stream.ServeFileResponse
	(*ServeFileRequest)(nil),      // 4: file_stream.ServeFileRequest
	(*RegisterPeersRequest)(nil),  // 5: file_stream.RegisterPeersRequest
	(*RegisterPeersResponse)(nil), // 6: file_stream.RegisterPeersResponse
}
var file_filepb_file_proto_depIdxs = []int32{
	1, // 0: file_stream.UploadFileRequest.info:type_name -> file_stream.FileInfo
	0, // 1: file_stream.FileService.UploadFile:input_type -> file_stream.UploadFileRequest
	4, // 2: file_stream.FileService.DownloadFile:input_type -> file_stream.ServeFileRequest
	5, // 3: file_stream.FileService.RegisterPeers:input_type -> file_stream.RegisterPeersRequest
	2, // 4: file_stream.FileService.UploadFile:output_type -> file_stream.UploadFileResponse
	3, // 5: file_stream.FileService.DownloadFile:output_type -> file_stream.ServeFileResponse
	6, // 6: file_stream.FileService.RegisterPeers:output_type -> file_stream.RegisterPeersResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_filepb_file_proto_init() }
func file_filepb_file_proto_init() {
	if File_filepb_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filepb_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_filepb_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_filepb_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
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
		file_filepb_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServeFileResponse); i {
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
		file_filepb_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServeFileRequest); i {
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
		file_filepb_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPeersRequest); i {
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
		file_filepb_file_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPeersResponse); i {
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
	file_filepb_file_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UploadFileRequest_Info)(nil),
		(*UploadFileRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filepb_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filepb_file_proto_goTypes,
		DependencyIndexes: file_filepb_file_proto_depIdxs,
		MessageInfos:      file_filepb_file_proto_msgTypes,
	}.Build()
	File_filepb_file_proto = out.File
	file_filepb_file_proto_rawDesc = nil
	file_filepb_file_proto_goTypes = nil
	file_filepb_file_proto_depIdxs = nil
}
