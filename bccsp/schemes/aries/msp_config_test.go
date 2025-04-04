/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package aries_test

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IdemixMSPConfig struct {
	// Name holds the identifier of the MSP
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ipk represents the (serialized) issuer public key
	Ipk []byte `protobuf:"bytes,2,opt,name=ipk,proto3" json:"ipk,omitempty"`
	// signer may contain crypto material to configure a default signer
	Signer *IdemixMSPSignerConfig `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	// revocation_pk is the public key used for revocation of credentials
	RevocationPk []byte `protobuf:"bytes,4,opt,name=revocation_pk,json=revocationPk,proto3" json:"revocation_pk,omitempty"`
	// epoch represents the current epoch (time interval) used for revocation
	Epoch int64 `protobuf:"varint,5,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// curve_id indicates which Elliptic Curve should be used
	CurveId              string   `protobuf:"bytes,6,opt,name=curve_id,json=curveId,proto3" json:"curve_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdemixMSPConfig) Reset()         { *m = IdemixMSPConfig{} }
func (m *IdemixMSPConfig) String() string { return proto.CompactTextString(m) }
func (*IdemixMSPConfig) ProtoMessage()    {}
func (*IdemixMSPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_783ba658a7862c73, []int{0}
}

func (m *IdemixMSPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdemixMSPConfig.Unmarshal(m, b)
}
func (m *IdemixMSPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdemixMSPConfig.Marshal(b, m, deterministic)
}
func (m *IdemixMSPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdemixMSPConfig.Merge(m, src)
}
func (m *IdemixMSPConfig) XXX_Size() int {
	return xxx_messageInfo_IdemixMSPConfig.Size(m)
}
func (m *IdemixMSPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IdemixMSPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IdemixMSPConfig proto.InternalMessageInfo

func (m *IdemixMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IdemixMSPConfig) GetIpk() []byte {
	if m != nil {
		return m.Ipk
	}
	return nil
}

func (m *IdemixMSPConfig) GetSigner() *IdemixMSPSignerConfig {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *IdemixMSPConfig) GetRevocationPk() []byte {
	if m != nil {
		return m.RevocationPk
	}
	return nil
}

func (m *IdemixMSPConfig) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *IdemixMSPConfig) GetCurveId() string {
	if m != nil {
		return m.CurveId
	}
	return ""
}

// IdemixMSPSignerConfig contains the crypto material to set up an idemix signing identity
type IdemixMSPSignerConfig struct {
	// cred represents the serialized idemix credential of the default signer
	Cred []byte `protobuf:"bytes,1,opt,name=cred,proto3" json:"cred,omitempty"`
	// sk is the secret key of the default signer, corresponding to credential Cred
	Sk []byte `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty"`
	// organizational_unit_identifier defines the organizational unit the default signer is in
	OrganizationalUnitIdentifier string `protobuf:"bytes,3,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier,proto3" json:"organizational_unit_identifier,omitempty"`
	// role defines whether the default signer is admin, peer, member or client
	Role int32 `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	// enrollment_id contains the enrollment id of this signer
	EnrollmentId string `protobuf:"bytes,5,opt,name=enrollment_id,json=enrollmentId,proto3" json:"enrollment_id,omitempty"`
	// credential_revocation_information contains a serialized CredentialRevocationInformation
	CredentialRevocationInformation []byte `protobuf:"bytes,6,opt,name=credential_revocation_information,json=credentialRevocationInformation,proto3" json:"credential_revocation_information,omitempty"`
	// RevocationHandle is the handle used to single out this credential and determine its revocation status
	RevocationHandle     string   `protobuf:"bytes,7,opt,name=revocation_handle,json=revocationHandle,proto3" json:"revocation_handle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdemixMSPSignerConfig) Reset()         { *m = IdemixMSPSignerConfig{} }
func (m *IdemixMSPSignerConfig) String() string { return proto.CompactTextString(m) }
func (*IdemixMSPSignerConfig) ProtoMessage()    {}
func (*IdemixMSPSignerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_783ba658a7862c73, []int{1}
}

func (m *IdemixMSPSignerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdemixMSPSignerConfig.Unmarshal(m, b)
}
func (m *IdemixMSPSignerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdemixMSPSignerConfig.Marshal(b, m, deterministic)
}
func (m *IdemixMSPSignerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdemixMSPSignerConfig.Merge(m, src)
}
func (m *IdemixMSPSignerConfig) XXX_Size() int {
	return xxx_messageInfo_IdemixMSPSignerConfig.Size(m)
}
func (m *IdemixMSPSignerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IdemixMSPSignerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IdemixMSPSignerConfig proto.InternalMessageInfo

func (m *IdemixMSPSignerConfig) GetCred() []byte {
	if m != nil {
		return m.Cred
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

func (m *IdemixMSPSignerConfig) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *IdemixMSPSignerConfig) GetEnrollmentId() string {
	if m != nil {
		return m.EnrollmentId
	}
	return ""
}

func (m *IdemixMSPSignerConfig) GetCredentialRevocationInformation() []byte {
	if m != nil {
		return m.CredentialRevocationInformation
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetRevocationHandle() string {
	if m != nil {
		return m.RevocationHandle
	}
	return ""
}

func init() {
	proto.RegisterType((*IdemixMSPConfig)(nil), "idemixmsp.IdemixMSPConfig")
	proto.RegisterType((*IdemixMSPSignerConfig)(nil), "idemixmsp.IdemixMSPSignerConfig")
}

func init() { proto.RegisterFile("idemixmsp/msp_config.proto", fileDescriptor_783ba658a7862c73) }

var fileDescriptor_783ba658a7862c73 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x8b, 0xd4, 0x30,
	0x14, 0xc7, 0x49, 0x67, 0x67, 0xd6, 0xc6, 0xee, 0xba, 0x06, 0x17, 0xaa, 0x88, 0x5b, 0xd7, 0xcb,
	0x80, 0x30, 0x03, 0x7a, 0xf1, 0x3c, 0x0a, 0x1a, 0x61, 0xa0, 0x64, 0x11, 0x16, 0x11, 0x4a, 0xb6,
	0xc9, 0xcc, 0x84, 0x36, 0x49, 0x49, 0x33, 0x8b, 0x78, 0xf6, 0x93, 0x78, 0xf4, 0xea, 0x87, 0x10,
	0xfc, 0x18, 0x1e, 0xfd, 0x14, 0x92, 0xd7, 0xdd, 0xb6, 0x82, 0xb7, 0xff, 0x7b, 0xf9, 0xf5, 0xff,
	0xde, 0x3f, 0x0d, 0x7e, 0xa4, 0x84, 0xd4, 0xea, 0xb3, 0x6e, 0x9b, 0xa5, 0x6e, 0x9b, 0xa2, 0xb4,
	0x66, 0xa3, 0xb6, 0x8b, 0xc6, 0x59, 0x6f, 0x49, 0xdc, 0x9f, 0x9d, 0xff, 0x44, 0xf8, 0x1e, 0x85,
	0x6a, 0x7d, 0x91, 0xbf, 0x06, 0x88, 0x10, 0x7c, 0x60, 0xb8, 0x96, 0x29, 0xca, 0xd0, 0x3c, 0x66,
	0xa0, 0xc9, 0x09, 0x9e, 0xa8, 0xa6, 0x4a, 0xa3, 0x0c, 0xcd, 0x13, 0x16, 0x24, 0x79, 0x85, 0x67,
	0xad, 0xda, 0x1a, 0xe9, 0xd2, 0x49, 0x86, 0xe6, 0x77, 0x5f, 0x64, 0x8b, 0xde, 0x75, 0xd1, 0x3b,
	0x5e, 0x00, 0xd1, 0xf9, 0xb2, 0x1b, 0x9e, 0x3c, 0xc3, 0x47, 0x4e, 0x5e, 0xdb, 0x92, 0x7b, 0x65,
	0x4d, 0xd1, 0x54, 0xe9, 0x01, 0xb8, 0x26, 0x43, 0x33, 0xaf, 0xc8, 0x03, 0x3c, 0x95, 0x8d, 0x2d,
	0x77, 0xe9, 0x34, 0x43, 0xf3, 0x09, 0xeb, 0x0a, 0xf2, 0x10, 0xdf, 0x29, 0xf7, 0xee, 0x5a, 0x16,
	0x4a, 0xa4, 0x33, 0x58, 0xef, 0x10, 0x6a, 0x2a, 0xce, 0x7f, 0x44, 0xf8, 0xf4, 0xbf, 0x73, 0x43,
	0x9e, 0xd2, 0x49, 0x01, 0x79, 0x12, 0x06, 0x9a, 0x1c, 0xe3, 0xa8, 0xbd, 0x8d, 0x13, 0xb5, 0x15,
	0x79, 0x83, 0x9f, 0x58, 0xb7, 0xe5, 0x46, 0x7d, 0x81, 0x05, 0x78, 0x5d, 0xec, 0x8d, 0xf2, 0x85,
	0x12, 0xd2, 0x78, 0xb5, 0x51, 0x37, 0x29, 0x63, 0xf6, 0xf8, 0x5f, 0xea, 0x83, 0x51, 0x9e, 0xf6,
	0x4c, 0x98, 0xe4, 0x6c, 0x2d, 0x21, 0xd0, 0x94, 0x81, 0x0e, 0x69, 0xa5, 0x71, 0xb6, 0xae, 0xb5,
	0x34, 0xc1, 0x10, 0x02, 0xc5, 0x2c, 0x19, 0x9a, 0x54, 0x90, 0xf7, 0xf8, 0x69, 0x58, 0x2b, 0x18,
	0xf1, 0xba, 0x18, 0xdd, 0x8e, 0x32, 0x1b, 0xeb, 0x34, 0x68, 0x08, 0x9c, 0xb0, 0xb3, 0x01, 0x64,
	0x3d, 0x47, 0x07, 0x8c, 0x3c, 0xc7, 0xf7, 0x47, 0x06, 0x3b, 0x6e, 0x44, 0x2d, 0xd3, 0x43, 0x18,
	0x7a, 0x32, 0x1c, 0xbc, 0x83, 0xfe, 0xea, 0x2b, 0xc2, 0x47, 0xa5, 0xd5, 0xc3, 0xbf, 0x5b, 0x1d,
	0xaf, 0xdb, 0xa6, 0xbb, 0xb8, 0x3c, 0x3c, 0x96, 0x1c, 0x7d, 0x3c, 0xdb, 0x2a, 0xbf, 0xdb, 0x5f,
	0x2d, 0x4a, 0xab, 0x97, 0x74, 0xb5, 0x5e, 0x76, 0xec, 0xb2, 0xff, 0xe4, 0x5b, 0x34, 0xa1, 0x97,
	0x97, 0xdf, 0xa3, 0x98, 0xde, 0x76, 0x7e, 0x8d, 0xf4, 0xef, 0xe8, 0xb4, 0xd7, 0x9f, 0xde, 0xe6,
	0xab, 0xb5, 0xf4, 0x5c, 0x70, 0xcf, 0xff, 0x8c, 0x98, 0xab, 0x19, 0x3c, 0xcc, 0x97, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0x8d, 0x2a, 0xd8, 0xb6, 0x02, 0x00, 0x00,
}
