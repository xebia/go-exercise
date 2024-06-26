// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.8
// source: registration.proto

package protobuf

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegistrationStatus int32

const (
	RegistrationStatus_REGISTRATION_UNKNOWN   RegistrationStatus = 0
	RegistrationStatus_REGISTRATION_PENDING   RegistrationStatus = 1
	RegistrationStatus_REGISTRATION_CONFIRMED RegistrationStatus = 2
	RegistrationStatus_REGISTRATION_BLOCKED   RegistrationStatus = 3
)

// Enum value maps for RegistrationStatus.
var (
	RegistrationStatus_name = map[int32]string{
		0: "REGISTRATION_UNKNOWN",
		1: "REGISTRATION_PENDING",
		2: "REGISTRATION_CONFIRMED",
		3: "REGISTRATION_BLOCKED",
	}
	RegistrationStatus_value = map[string]int32{
		"REGISTRATION_UNKNOWN":   0,
		"REGISTRATION_PENDING":   1,
		"REGISTRATION_CONFIRMED": 2,
		"REGISTRATION_BLOCKED":   3,
	}
)

func (x RegistrationStatus) Enum() *RegistrationStatus {
	p := new(RegistrationStatus)
	*p = x
	return p
}

func (x RegistrationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegistrationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_registration_proto_enumTypes[0].Descriptor()
}

func (RegistrationStatus) Type() protoreflect.EnumType {
	return &file_registration_proto_enumTypes[0]
}

func (x RegistrationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegistrationStatus.Descriptor instead.
func (RegistrationStatus) EnumDescriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{0}
}

type RegisterPatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient *Patient `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *RegisterPatientRequest) Reset() {
	*x = RegisterPatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPatientRequest) ProtoMessage() {}

func (x *RegisterPatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPatientRequest.ProtoReflect.Descriptor instead.
func (*RegisterPatientRequest) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterPatientRequest) GetPatient() *Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

type RegisterPatientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientUid string `protobuf:"bytes,1,opt,name=patientUid,proto3" json:"patientUid,omitempty"`
}

func (x *RegisterPatientResponse) Reset() {
	*x = RegisterPatientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPatientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPatientResponse) ProtoMessage() {}

func (x *RegisterPatientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPatientResponse.ProtoReflect.Descriptor instead.
func (*RegisterPatientResponse) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterPatientResponse) GetPatientUid() string {
	if x != nil {
		return x.PatientUid
	}
	return ""
}

type Patient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BSN      string             `protobuf:"bytes,1,opt,name=BSN,proto3" json:"BSN,omitempty"`
	FullName string             `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Address  *Address           `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Contact  *Contact           `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
	Status   RegistrationStatus `protobuf:"varint,5,opt,name=status,proto3,enum=protobuf.RegistrationStatus" json:"status,omitempty"`
}

func (x *Patient) Reset() {
	*x = Patient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patient) ProtoMessage() {}

func (x *Patient) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patient.ProtoReflect.Descriptor instead.
func (*Patient) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{2}
}

func (x *Patient) GetBSN() string {
	if x != nil {
		return x.BSN
	}
	return ""
}

func (x *Patient) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Patient) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Patient) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *Patient) GetStatus() RegistrationStatus {
	if x != nil {
		return x.Status
	}
	return RegistrationStatus_REGISTRATION_UNKNOWN
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostalCode  string `protobuf:"bytes,1,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	HouseNumber int64  `protobuf:"varint,2,opt,name=houseNumber,proto3" json:"houseNumber,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{3}
}

func (x *Address) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Address) GetHouseNumber() int64 {
	if x != nil {
		return x.HouseNumber
	}
	return 0
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailAddress string `protobuf:"bytes,1,opt,name=emailAddress,proto3" json:"emailAddress,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{4}
}

func (x *Contact) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

var File_registration_proto protoreflect.FileDescriptor

var file_registration_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x22, 0x48, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72,
	0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x39, 0x0a, 0x17, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x22, 0xd0, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x53, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x42, 0x53, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1f, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4b, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x2a, 0x7e, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x52,
	0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x1a, 0x0a, 0x16, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x52,
	0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x45, 0x44, 0x10, 0x03, 0x32, 0x75, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0f,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x23, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x63, 0x47,
	0x72, 0x6f, 0x6c, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registration_proto_rawDescOnce sync.Once
	file_registration_proto_rawDescData = file_registration_proto_rawDesc
)

func file_registration_proto_rawDescGZIP() []byte {
	file_registration_proto_rawDescOnce.Do(func() {
		file_registration_proto_rawDescData = protoimpl.X.CompressGZIP(file_registration_proto_rawDescData)
	})
	return file_registration_proto_rawDescData
}

var file_registration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_registration_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_registration_proto_goTypes = []interface{}{
	(RegistrationStatus)(0),         // 0: protobuf.RegistrationStatus
	(*RegisterPatientRequest)(nil),  // 1: protobuf.RegisterPatientRequest
	(*RegisterPatientResponse)(nil), // 2: protobuf.RegisterPatientResponse
	(*Patient)(nil),                 // 3: protobuf.Patient
	(*Address)(nil),                 // 4: protobuf.Address
	(*Contact)(nil),                 // 5: protobuf.Contact
}
var file_registration_proto_depIdxs = []int32{
	3, // 0: protobuf.RegisterPatientRequest.patient:type_name -> protobuf.Patient
	4, // 1: protobuf.Patient.address:type_name -> protobuf.Address
	5, // 2: protobuf.Patient.contact:type_name -> protobuf.Contact
	0, // 3: protobuf.Patient.status:type_name -> protobuf.RegistrationStatus
	1, // 4: protobuf.RegistrationService.RegisterPatient:input_type -> protobuf.RegisterPatientRequest
	2, // 5: protobuf.RegistrationService.RegisterPatient:output_type -> protobuf.RegisterPatientResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_registration_proto_init() }
func file_registration_proto_init() {
	if File_registration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPatientRequest); i {
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
		file_registration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPatientResponse); i {
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
		file_registration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patient); i {
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
		file_registration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_registration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
			RawDescriptor: file_registration_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_registration_proto_goTypes,
		DependencyIndexes: file_registration_proto_depIdxs,
		EnumInfos:         file_registration_proto_enumTypes,
		MessageInfos:      file_registration_proto_msgTypes,
	}.Build()
	File_registration_proto = out.File
	file_registration_proto_rawDesc = nil
	file_registration_proto_goTypes = nil
	file_registration_proto_depIdxs = nil
}
