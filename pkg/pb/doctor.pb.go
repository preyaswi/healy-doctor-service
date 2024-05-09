// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: pkg/pb/doctor.proto

package pb

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

type DoctorSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName          string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email             string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber       string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Password          string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword   string `protobuf:"bytes,5,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
	Specialization    string `protobuf:"bytes,6,opt,name=specialization,proto3" json:"specialization,omitempty"`
	YearsOfExperience int32  `protobuf:"varint,7,opt,name=years_of_experience,json=yearsOfExperience,proto3" json:"years_of_experience,omitempty"`
	LicenseNumber     string `protobuf:"bytes,8,opt,name=license_number,json=licenseNumber,proto3" json:"license_number,omitempty"`
}

func (x *DoctorSignUpRequest) Reset() {
	*x = DoctorSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_doctor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorSignUpRequest) ProtoMessage() {}

func (x *DoctorSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_doctor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorSignUpRequest.ProtoReflect.Descriptor instead.
func (*DoctorSignUpRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_doctor_proto_rawDescGZIP(), []int{0}
}

func (x *DoctorSignUpRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *DoctorSignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DoctorSignUpRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *DoctorSignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DoctorSignUpRequest) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

func (x *DoctorSignUpRequest) GetSpecialization() string {
	if x != nil {
		return x.Specialization
	}
	return ""
}

func (x *DoctorSignUpRequest) GetYearsOfExperience() int32 {
	if x != nil {
		return x.YearsOfExperience
	}
	return 0
}

func (x *DoctorSignUpRequest) GetLicenseNumber() string {
	if x != nil {
		return x.LicenseNumber
	}
	return ""
}

type DoctorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName          string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email             string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber       string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Specialization    string `protobuf:"bytes,5,opt,name=specialization,proto3" json:"specialization,omitempty"`
	YearsOfExperience int32  `protobuf:"varint,6,opt,name=years_of_experience,json=yearsOfExperience,proto3" json:"years_of_experience,omitempty"`
	LicenseNumber     string `protobuf:"bytes,7,opt,name=license_number,json=licenseNumber,proto3" json:"license_number,omitempty"`
}

func (x *DoctorDetail) Reset() {
	*x = DoctorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_doctor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorDetail) ProtoMessage() {}

func (x *DoctorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_doctor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorDetail.ProtoReflect.Descriptor instead.
func (*DoctorDetail) Descriptor() ([]byte, []int) {
	return file_pkg_pb_doctor_proto_rawDescGZIP(), []int{1}
}

func (x *DoctorDetail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DoctorDetail) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *DoctorDetail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DoctorDetail) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *DoctorDetail) GetSpecialization() string {
	if x != nil {
		return x.Specialization
	}
	return ""
}

func (x *DoctorDetail) GetYearsOfExperience() int32 {
	if x != nil {
		return x.YearsOfExperience
	}
	return 0
}

func (x *DoctorDetail) GetLicenseNumber() string {
	if x != nil {
		return x.LicenseNumber
	}
	return ""
}

type DoctorSignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorDetail *DoctorDetail `protobuf:"bytes,1,opt,name=DoctorDetail,proto3" json:"DoctorDetail,omitempty"`
	AccessToken  string        `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken string        `protobuf:"bytes,3,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
}

func (x *DoctorSignUpResponse) Reset() {
	*x = DoctorSignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_doctor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorSignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorSignUpResponse) ProtoMessage() {}

func (x *DoctorSignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_doctor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorSignUpResponse.ProtoReflect.Descriptor instead.
func (*DoctorSignUpResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_doctor_proto_rawDescGZIP(), []int{2}
}

func (x *DoctorSignUpResponse) GetDoctorDetail() *DoctorDetail {
	if x != nil {
		return x.DoctorDetail
	}
	return nil
}

func (x *DoctorSignUpResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *DoctorSignUpResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type DoctorLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *DoctorLoginRequest) Reset() {
	*x = DoctorLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_doctor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorLoginRequest) ProtoMessage() {}

func (x *DoctorLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_doctor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorLoginRequest.ProtoReflect.Descriptor instead.
func (*DoctorLoginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_doctor_proto_rawDescGZIP(), []int{3}
}

func (x *DoctorLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DoctorLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DoctorLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorDetail *DoctorDetail `protobuf:"bytes,1,opt,name=DoctorDetail,proto3" json:"DoctorDetail,omitempty"`
	AccessToken  string        `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken string        `protobuf:"bytes,3,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
}

func (x *DoctorLoginResponse) Reset() {
	*x = DoctorLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_doctor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorLoginResponse) ProtoMessage() {}

func (x *DoctorLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_doctor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorLoginResponse.ProtoReflect.Descriptor instead.
func (*DoctorLoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_doctor_proto_rawDescGZIP(), []int{4}
}

func (x *DoctorLoginResponse) GetDoctorDetail() *DoctorDetail {
	if x != nil {
		return x.DoctorDetail
	}
	return nil
}

func (x *DoctorLoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *DoctorLoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type Doreq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Doreq) Reset() {
	*x = Doreq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_doctor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Doreq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Doreq) ProtoMessage() {}

func (x *Doreq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_doctor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Doreq.ProtoReflect.Descriptor instead.
func (*Doreq) Descriptor() ([]byte, []int) {
	return file_pkg_pb_doctor_proto_rawDescGZIP(), []int{5}
}

type DoctorsDetailr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName          string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email             string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber       string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Specialization    string `protobuf:"bytes,5,opt,name=specialization,proto3" json:"specialization,omitempty"`
	YearsOfExperience int32  `protobuf:"varint,6,opt,name=years_of_experience,json=yearsOfExperience,proto3" json:"years_of_experience,omitempty"`
	LicenseNumber     string `protobuf:"bytes,7,opt,name=license_number,json=licenseNumber,proto3" json:"license_number,omitempty"`
	Rating            int32  `protobuf:"varint,8,opt,name=Rating,proto3" json:"Rating,omitempty"`
}

func (x *DoctorsDetailr) Reset() {
	*x = DoctorsDetailr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_doctor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorsDetailr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorsDetailr) ProtoMessage() {}

func (x *DoctorsDetailr) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_doctor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorsDetailr.ProtoReflect.Descriptor instead.
func (*DoctorsDetailr) Descriptor() ([]byte, []int) {
	return file_pkg_pb_doctor_proto_rawDescGZIP(), []int{6}
}

func (x *DoctorsDetailr) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DoctorsDetailr) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *DoctorsDetailr) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DoctorsDetailr) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *DoctorsDetailr) GetSpecialization() string {
	if x != nil {
		return x.Specialization
	}
	return ""
}

func (x *DoctorsDetailr) GetYearsOfExperience() int32 {
	if x != nil {
		return x.YearsOfExperience
	}
	return 0
}

func (x *DoctorsDetailr) GetLicenseNumber() string {
	if x != nil {
		return x.LicenseNumber
	}
	return ""
}

func (x *DoctorsDetailr) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type DoctorsDetailre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorsDetailr []*DoctorsDetailr `protobuf:"bytes,1,rep,name=DoctorsDetailr,proto3" json:"DoctorsDetailr,omitempty"`
}

func (x *DoctorsDetailre) Reset() {
	*x = DoctorsDetailre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_doctor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorsDetailre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorsDetailre) ProtoMessage() {}

func (x *DoctorsDetailre) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_doctor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorsDetailre.ProtoReflect.Descriptor instead.
func (*DoctorsDetailre) Descriptor() ([]byte, []int) {
	return file_pkg_pb_doctor_proto_rawDescGZIP(), []int{7}
}

func (x *DoctorsDetailre) GetDoctorsDetailr() []*DoctorsDetailr {
	if x != nil {
		return x.DoctorsDetailr
	}
	return nil
}

var File_pkg_pb_doctor_proto protoreflect.FileDescriptor

var file_pkg_pb_doctor_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xb1, 0x02,
	0x0a, 0x13, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x79, 0x65,
	0x61, 0x72, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x79, 0x65, 0x61, 0x72, 0x73, 0x4f, 0x66,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0xf3, 0x01, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2e, 0x0a, 0x13, 0x79, 0x65, 0x61, 0x72, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x79,
	0x65, 0x61, 0x72, 0x73, 0x4f, 0x66, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x96, 0x01, 0x0a, 0x14, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0c, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x46, 0x0a, 0x12, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x13, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0c, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x07, 0x0a, 0x05, 0x44, 0x6f, 0x72, 0x65, 0x71, 0x22, 0x8d, 0x02, 0x0a, 0x0e, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x79, 0x65,
	0x61, 0x72, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x79, 0x65, 0x61, 0x72, 0x73, 0x4f, 0x66,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x51, 0x0a, 0x0f, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x72, 0x65, 0x12, 0x3e, 0x0a, 0x0e,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x72, 0x52, 0x0e, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x72, 0x32, 0xda, 0x01, 0x0a,
	0x06, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x4b, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1b, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0d, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x0d, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x72, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x72, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_doctor_proto_rawDescOnce sync.Once
	file_pkg_pb_doctor_proto_rawDescData = file_pkg_pb_doctor_proto_rawDesc
)

func file_pkg_pb_doctor_proto_rawDescGZIP() []byte {
	file_pkg_pb_doctor_proto_rawDescOnce.Do(func() {
		file_pkg_pb_doctor_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_doctor_proto_rawDescData)
	})
	return file_pkg_pb_doctor_proto_rawDescData
}

var file_pkg_pb_doctor_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_pb_doctor_proto_goTypes = []interface{}{
	(*DoctorSignUpRequest)(nil),  // 0: doctor.DoctorSignUpRequest
	(*DoctorDetail)(nil),         // 1: doctor.DoctorDetail
	(*DoctorSignUpResponse)(nil), // 2: doctor.DoctorSignUpResponse
	(*DoctorLoginRequest)(nil),   // 3: doctor.DoctorLoginRequest
	(*DoctorLoginResponse)(nil),  // 4: doctor.DoctorLoginResponse
	(*Doreq)(nil),                // 5: doctor.Doreq
	(*DoctorsDetailr)(nil),       // 6: doctor.DoctorsDetailr
	(*DoctorsDetailre)(nil),      // 7: doctor.DoctorsDetailre
}
var file_pkg_pb_doctor_proto_depIdxs = []int32{
	1, // 0: doctor.DoctorSignUpResponse.DoctorDetail:type_name -> doctor.DoctorDetail
	1, // 1: doctor.DoctorLoginResponse.DoctorDetail:type_name -> doctor.DoctorDetail
	6, // 2: doctor.DoctorsDetailre.DoctorsDetailr:type_name -> doctor.DoctorsDetailr
	0, // 3: doctor.Doctor.DoctorSignUp:input_type -> doctor.DoctorSignUpRequest
	3, // 4: doctor.Doctor.DoctorLogin:input_type -> doctor.DoctorLoginRequest
	5, // 5: doctor.Doctor.DoctorsDetail:input_type -> doctor.Doreq
	2, // 6: doctor.Doctor.DoctorSignUp:output_type -> doctor.DoctorSignUpResponse
	4, // 7: doctor.Doctor.DoctorLogin:output_type -> doctor.DoctorLoginResponse
	7, // 8: doctor.Doctor.DoctorsDetail:output_type -> doctor.DoctorsDetailre
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_pb_doctor_proto_init() }
func file_pkg_pb_doctor_proto_init() {
	if File_pkg_pb_doctor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_doctor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorSignUpRequest); i {
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
		file_pkg_pb_doctor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorDetail); i {
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
		file_pkg_pb_doctor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorSignUpResponse); i {
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
		file_pkg_pb_doctor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorLoginRequest); i {
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
		file_pkg_pb_doctor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorLoginResponse); i {
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
		file_pkg_pb_doctor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Doreq); i {
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
		file_pkg_pb_doctor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorsDetailr); i {
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
		file_pkg_pb_doctor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorsDetailre); i {
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
			RawDescriptor: file_pkg_pb_doctor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_doctor_proto_goTypes,
		DependencyIndexes: file_pkg_pb_doctor_proto_depIdxs,
		MessageInfos:      file_pkg_pb_doctor_proto_msgTypes,
	}.Build()
	File_pkg_pb_doctor_proto = out.File
	file_pkg_pb_doctor_proto_rawDesc = nil
	file_pkg_pb_doctor_proto_goTypes = nil
	file_pkg_pb_doctor_proto_depIdxs = nil
}
