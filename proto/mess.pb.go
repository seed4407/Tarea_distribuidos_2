// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: mess.proto

package Tarea_distribuidos_2

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

type Datos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
	Estado   string `protobuf:"bytes,3,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *Datos) Reset() {
	*x = Datos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datos) ProtoMessage() {}

func (x *Datos) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datos.ProtoReflect.Descriptor instead.
func (*Datos) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{0}
}

func (x *Datos) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Datos) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

func (x *Datos) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type Recepcion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok string `protobuf:"bytes,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Recepcion) Reset() {
	*x = Recepcion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recepcion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recepcion) ProtoMessage() {}

func (x *Recepcion) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recepcion.ProtoReflect.Descriptor instead.
func (*Recepcion) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{1}
}

func (x *Recepcion) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

type Datos_DataNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *Datos_DataNode) Reset() {
	*x = Datos_DataNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datos_DataNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datos_DataNode) ProtoMessage() {}

func (x *Datos_DataNode) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datos_DataNode.ProtoReflect.Descriptor instead.
func (*Datos_DataNode) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{2}
}

func (x *Datos_DataNode) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Datos_DataNode) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

type Lista_Datos_DataNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListaDatos_DataNode []*Datos_DataNode `protobuf:"bytes,1,rep,name=lista_datos_DataNode,json=listaDatosDataNode,proto3" json:"lista_datos_DataNode,omitempty"`
}

func (x *Lista_Datos_DataNode) Reset() {
	*x = Lista_Datos_DataNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lista_Datos_DataNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lista_Datos_DataNode) ProtoMessage() {}

func (x *Lista_Datos_DataNode) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lista_Datos_DataNode.ProtoReflect.Descriptor instead.
func (*Lista_Datos_DataNode) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{3}
}

func (x *Lista_Datos_DataNode) GetListaDatos_DataNode() []*Datos_DataNode {
	if x != nil {
		return x.ListaDatos_DataNode
	}
	return nil
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListaId []string `protobuf:"bytes,1,rep,name=lista_id,json=listaId,proto3" json:"lista_id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{4}
}

func (x *Id) GetListaId() []string {
	if x != nil {
		return x.ListaId
	}
	return nil
}

type Estado_Persona struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado string `protobuf:"bytes,1,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *Estado_Persona) Reset() {
	*x = Estado_Persona{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Estado_Persona) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Estado_Persona) ProtoMessage() {}

func (x *Estado_Persona) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Estado_Persona.ProtoReflect.Descriptor instead.
func (*Estado_Persona) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{5}
}

func (x *Estado_Persona) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

var File_mess_proto protoreflect.FileDescriptor

var file_mess_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x22, 0x53, 0x0a, 0x05, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x65, 0x70,
	0x63, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x6f, 0x6b, 0x22, 0x44, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x22, 0x5e, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x61, 0x5f, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x46, 0x0a, 0x14, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x6f,
	0x73, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x6f, 0x73, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x1f, 0x0a, 0x02, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x45,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x5f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x32, 0x3c, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x30, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x5f, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x6f, 0x73,
	0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x70, 0x63, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x32, 0xc7, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x43, 0x0a, 0x19, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x5f, 0x49, 0x6e,
	0x66, 0x6f, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x5f, 0x31, 0x12, 0x08, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x61, 0x5f, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x19, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74,
	0x75, 0x64, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65,
	0x5f, 0x32, 0x12, 0x08, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x1a, 0x1a, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x5f, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x5f,
	0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0f, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x0b, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x32, 0x42, 0x0a,
	0x03, 0x4f, 0x4e, 0x55, 0x12, 0x3b, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61,
	0x72, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x5f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x1a, 0x0f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x42, 0x5a, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x42, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x65, 0x65, 0x64, 0x34, 0x34, 0x30, 0x37, 0x2f, 0x54, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x69, 0x64, 0x6f, 0x73, 0x5f, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mess_proto_rawDescOnce sync.Once
	file_mess_proto_rawDescData = file_mess_proto_rawDesc
)

func file_mess_proto_rawDescGZIP() []byte {
	file_mess_proto_rawDescOnce.Do(func() {
		file_mess_proto_rawDescData = protoimpl.X.CompressGZIP(file_mess_proto_rawDescData)
	})
	return file_mess_proto_rawDescData
}

var file_mess_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mess_proto_goTypes = []interface{}{
	(*Datos)(nil),                // 0: grpc.Datos
	(*Recepcion)(nil),            // 1: grpc.Recepcion
	(*Datos_DataNode)(nil),       // 2: grpc.Datos_DataNode
	(*Lista_Datos_DataNode)(nil), // 3: grpc.Lista_Datos_DataNode
	(*Id)(nil),                   // 4: grpc.Id
	(*Estado_Persona)(nil),       // 5: grpc.Estado_Persona
}
var file_mess_proto_depIdxs = []int32{
	2, // 0: grpc.Lista_Datos_DataNode.lista_datos_DataNode:type_name -> grpc.Datos_DataNode
	0, // 1: grpc.NameNode.Recepcion_Info:input_type -> grpc.Datos
	4, // 2: grpc.DataNode.Solicitud_Info_DataNode_1:input_type -> grpc.Id
	4, // 3: grpc.DataNode.Solicitud_Info_DataNode_2:input_type -> grpc.Id
	0, // 4: grpc.DataNode.RegistrarNombre:input_type -> grpc.Datos
	5, // 5: grpc.ONU.ConsultarNombres:input_type -> grpc.Estado_Persona
	1, // 6: grpc.NameNode.Recepcion_Info:output_type -> grpc.Recepcion
	3, // 7: grpc.DataNode.Solicitud_Info_DataNode_1:output_type -> grpc.Lista_Datos_DataNode
	3, // 8: grpc.DataNode.Solicitud_Info_DataNode_2:output_type -> grpc.Lista_Datos_DataNode
	1, // 9: grpc.DataNode.RegistrarNombre:output_type -> grpc.Recepcion
	1, // 10: grpc.ONU.ConsultarNombres:output_type -> grpc.Recepcion
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mess_proto_init() }
func file_mess_proto_init() {
	if File_mess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mess_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datos); i {
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
		file_mess_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recepcion); i {
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
		file_mess_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datos_DataNode); i {
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
		file_mess_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lista_Datos_DataNode); i {
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
		file_mess_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_mess_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Estado_Persona); i {
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
			RawDescriptor: file_mess_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_mess_proto_goTypes,
		DependencyIndexes: file_mess_proto_depIdxs,
		MessageInfos:      file_mess_proto_msgTypes,
	}.Build()
	File_mess_proto = out.File
	file_mess_proto_rawDesc = nil
	file_mess_proto_goTypes = nil
	file_mess_proto_depIdxs = nil
}
