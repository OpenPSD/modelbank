// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modelbank.proto

package providers

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Consent struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tpp                  string               `protobuf:"bytes,2,opt,name=tpp,proto3" json:"tpp,omitempty"`
	TppCertificate       string               `protobuf:"bytes,3,opt,name=tppCertificate,proto3" json:"tppCertificate,omitempty"`
	Access               *Consent_Access      `protobuf:"bytes,4,opt,name=access,proto3" json:"access,omitempty"`
	RecurringIndicator   bool                 `protobuf:"varint,5,opt,name=recurringIndicator,proto3" json:"recurringIndicator,omitempty"`
	ValidUntil           *timestamp.Timestamp `protobuf:"bytes,6,opt,name=validUntil,proto3" json:"validUntil,omitempty"`
	FrequencyPerDay      int32                `protobuf:"varint,7,opt,name=frequencyPerDay,proto3" json:"frequencyPerDay,omitempty"`
	LastActionDate       *timestamp.Timestamp `protobuf:"bytes,8,opt,name=lastActionDate,proto3" json:"lastActionDate,omitempty"`
	ConsentStatus        string               `protobuf:"bytes,9,opt,name=consentStatus,proto3" json:"consentStatus,omitempty"`
	ScaStatus            string               `protobuf:"bytes,10,opt,name=scaStatus,proto3" json:"scaStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Consent) Reset()         { *m = Consent{} }
func (m *Consent) String() string { return proto.CompactTextString(m) }
func (*Consent) ProtoMessage()    {}
func (*Consent) Descriptor() ([]byte, []int) {
	return fileDescriptor_modelbank_a3f87e35c54f6826, []int{0}
}
func (m *Consent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consent.Unmarshal(m, b)
}
func (m *Consent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consent.Marshal(b, m, deterministic)
}
func (dst *Consent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consent.Merge(dst, src)
}
func (m *Consent) XXX_Size() int {
	return xxx_messageInfo_Consent.Size(m)
}
func (m *Consent) XXX_DiscardUnknown() {
	xxx_messageInfo_Consent.DiscardUnknown(m)
}

var xxx_messageInfo_Consent proto.InternalMessageInfo

func (m *Consent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consent) GetTpp() string {
	if m != nil {
		return m.Tpp
	}
	return ""
}

func (m *Consent) GetTppCertificate() string {
	if m != nil {
		return m.TppCertificate
	}
	return ""
}

func (m *Consent) GetAccess() *Consent_Access {
	if m != nil {
		return m.Access
	}
	return nil
}

func (m *Consent) GetRecurringIndicator() bool {
	if m != nil {
		return m.RecurringIndicator
	}
	return false
}

func (m *Consent) GetValidUntil() *timestamp.Timestamp {
	if m != nil {
		return m.ValidUntil
	}
	return nil
}

func (m *Consent) GetFrequencyPerDay() int32 {
	if m != nil {
		return m.FrequencyPerDay
	}
	return 0
}

func (m *Consent) GetLastActionDate() *timestamp.Timestamp {
	if m != nil {
		return m.LastActionDate
	}
	return nil
}

func (m *Consent) GetConsentStatus() string {
	if m != nil {
		return m.ConsentStatus
	}
	return ""
}

func (m *Consent) GetScaStatus() string {
	if m != nil {
		return m.ScaStatus
	}
	return ""
}

type Consent_Account struct {
	Iban                 string   `protobuf:"bytes,1,opt,name=iban,proto3" json:"iban,omitempty"`
	Bban                 string   `protobuf:"bytes,2,opt,name=bban,proto3" json:"bban,omitempty"`
	Pan                  string   `protobuf:"bytes,3,opt,name=pan,proto3" json:"pan,omitempty"`
	MaskedPan            string   `protobuf:"bytes,4,opt,name=maskedPan,proto3" json:"maskedPan,omitempty"`
	Msisdn               string   `protobuf:"bytes,5,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	Currency             string   `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consent_Account) Reset()         { *m = Consent_Account{} }
func (m *Consent_Account) String() string { return proto.CompactTextString(m) }
func (*Consent_Account) ProtoMessage()    {}
func (*Consent_Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_modelbank_a3f87e35c54f6826, []int{0, 0}
}
func (m *Consent_Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consent_Account.Unmarshal(m, b)
}
func (m *Consent_Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consent_Account.Marshal(b, m, deterministic)
}
func (dst *Consent_Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consent_Account.Merge(dst, src)
}
func (m *Consent_Account) XXX_Size() int {
	return xxx_messageInfo_Consent_Account.Size(m)
}
func (m *Consent_Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Consent_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Consent_Account proto.InternalMessageInfo

func (m *Consent_Account) GetIban() string {
	if m != nil {
		return m.Iban
	}
	return ""
}

func (m *Consent_Account) GetBban() string {
	if m != nil {
		return m.Bban
	}
	return ""
}

func (m *Consent_Account) GetPan() string {
	if m != nil {
		return m.Pan
	}
	return ""
}

func (m *Consent_Account) GetMaskedPan() string {
	if m != nil {
		return m.MaskedPan
	}
	return ""
}

func (m *Consent_Account) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *Consent_Account) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type Consent_Access struct {
	Accounts             []*Consent_Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Balances             []*Consent_Account `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances,omitempty"`
	Transactions         []*Consent_Account `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	AvailableAccounts    string             `protobuf:"bytes,4,opt,name=availableAccounts,proto3" json:"availableAccounts,omitempty"`
	AllPsd2              string             `protobuf:"bytes,5,opt,name=allPsd2,proto3" json:"allPsd2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Consent_Access) Reset()         { *m = Consent_Access{} }
func (m *Consent_Access) String() string { return proto.CompactTextString(m) }
func (*Consent_Access) ProtoMessage()    {}
func (*Consent_Access) Descriptor() ([]byte, []int) {
	return fileDescriptor_modelbank_a3f87e35c54f6826, []int{0, 1}
}
func (m *Consent_Access) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consent_Access.Unmarshal(m, b)
}
func (m *Consent_Access) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consent_Access.Marshal(b, m, deterministic)
}
func (dst *Consent_Access) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consent_Access.Merge(dst, src)
}
func (m *Consent_Access) XXX_Size() int {
	return xxx_messageInfo_Consent_Access.Size(m)
}
func (m *Consent_Access) XXX_DiscardUnknown() {
	xxx_messageInfo_Consent_Access.DiscardUnknown(m)
}

var xxx_messageInfo_Consent_Access proto.InternalMessageInfo

func (m *Consent_Access) GetAccounts() []*Consent_Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *Consent_Access) GetBalances() []*Consent_Account {
	if m != nil {
		return m.Balances
	}
	return nil
}

func (m *Consent_Access) GetTransactions() []*Consent_Account {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Consent_Access) GetAvailableAccounts() string {
	if m != nil {
		return m.AvailableAccounts
	}
	return ""
}

func (m *Consent_Access) GetAllPsd2() string {
	if m != nil {
		return m.AllPsd2
	}
	return ""
}

func init() {
	proto.RegisterType((*Consent)(nil), "providers.Consent")
	proto.RegisterType((*Consent_Account)(nil), "providers.Consent.Account")
	proto.RegisterType((*Consent_Access)(nil), "providers.Consent.Access")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ModelBankClient is the client API for ModelBank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModelBankClient interface {
	CreateConsent(ctx context.Context, in *Consent, opts ...grpc.CallOption) (*Consent, error)
}

type modelBankClient struct {
	cc *grpc.ClientConn
}

func NewModelBankClient(cc *grpc.ClientConn) ModelBankClient {
	return &modelBankClient{cc}
}

func (c *modelBankClient) CreateConsent(ctx context.Context, in *Consent, opts ...grpc.CallOption) (*Consent, error) {
	out := new(Consent)
	err := c.cc.Invoke(ctx, "/providers.ModelBank/CreateConsent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelBankServer is the server API for ModelBank service.
type ModelBankServer interface {
	CreateConsent(context.Context, *Consent) (*Consent, error)
}

func RegisterModelBankServer(s *grpc.Server, srv ModelBankServer) {
	s.RegisterService(&_ModelBank_serviceDesc, srv)
}

func _ModelBank_CreateConsent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelBankServer).CreateConsent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/providers.ModelBank/CreateConsent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelBankServer).CreateConsent(ctx, req.(*Consent))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModelBank_serviceDesc = grpc.ServiceDesc{
	ServiceName: "providers.ModelBank",
	HandlerType: (*ModelBankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsent",
			Handler:    _ModelBank_CreateConsent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modelbank.proto",
}

func init() { proto.RegisterFile("modelbank.proto", fileDescriptor_modelbank_a3f87e35c54f6826) }

var fileDescriptor_modelbank_a3f87e35c54f6826 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0x36, 0xed, 0x5d, 0xda, 0x8c, 0xde, 0x9d, 0xce, 0x83, 0xac, 0x41, 0xb0, 0x1c, 0x22, 0x7d,
	0x90, 0x1c, 0x56, 0x10, 0xf4, 0x41, 0xe8, 0xf5, 0x10, 0x7c, 0x10, 0x4a, 0xd4, 0x1f, 0x30, 0xd9,
	0x6c, 0xcb, 0xd2, 0x74, 0x37, 0xee, 0x6e, 0x0b, 0xf7, 0xe6, 0x9f, 0xf0, 0xf7, 0xf8, 0xd7, 0x64,
	0x37, 0x49, 0xcf, 0xeb, 0x15, 0xfa, 0x36, 0xf3, 0xcd, 0x37, 0x99, 0x99, 0xef, 0xcb, 0xc2, 0xc5,
	0x5a, 0x97, 0xa2, 0x2a, 0x48, 0xad, 0xb2, 0xda, 0x68, 0xa7, 0x31, 0xa9, 0x8d, 0xde, 0xca, 0x52,
	0x18, 0x9b, 0xbe, 0x5a, 0x6a, 0xbd, 0xac, 0xc4, 0x55, 0x28, 0x14, 0x9b, 0xc5, 0x95, 0x93, 0x6b,
	0x61, 0x1d, 0xad, 0xeb, 0x86, 0x7b, 0xf9, 0x37, 0x86, 0xc1, 0x4c, 0x2b, 0x2b, 0x94, 0xc3, 0x73,
	0xe8, 0xc9, 0x92, 0x45, 0xa3, 0x68, 0x9c, 0xe4, 0x3d, 0x59, 0xe2, 0x53, 0xe8, 0xbb, 0xba, 0x66,
	0xbd, 0x00, 0xf8, 0x10, 0xdf, 0xc0, 0xb9, 0xab, 0xeb, 0x99, 0x30, 0x4e, 0x2e, 0x24, 0x27, 0x27,
	0x58, 0x3f, 0x14, 0xf7, 0x50, 0x7c, 0x07, 0x31, 0x71, 0x2e, 0xac, 0x65, 0x27, 0xa3, 0x68, 0xfc,
	0x78, 0xf2, 0x22, 0xdb, 0xad, 0x94, 0xb5, 0xd3, 0xb2, 0x69, 0x20, 0xe4, 0x2d, 0x11, 0x33, 0x40,
	0x23, 0xf8, 0xc6, 0x18, 0xa9, 0x96, 0x5f, 0x55, 0xe9, 0xbf, 0xa3, 0x0d, 0x3b, 0x1d, 0x45, 0xe3,
	0x61, 0x7e, 0xa0, 0x82, 0x9f, 0x00, 0xb6, 0x54, 0xc9, 0xf2, 0xa7, 0x72, 0xb2, 0x62, 0x71, 0x18,
	0x93, 0x66, 0xcd, 0xb9, 0x59, 0x77, 0x6e, 0xf6, 0xa3, 0x3b, 0x37, 0xff, 0x8f, 0x8d, 0x63, 0xb8,
	0x58, 0x18, 0xf1, 0x6b, 0x23, 0x14, 0xbf, 0x9d, 0x0b, 0x73, 0x43, 0xb7, 0x6c, 0x30, 0x8a, 0xc6,
	0xa7, 0xf9, 0x3e, 0x8c, 0xd7, 0x70, 0x5e, 0x91, 0x75, 0x53, 0xee, 0xa4, 0x56, 0x37, 0xfe, 0xe0,
	0xe1, 0xd1, 0x49, 0x7b, 0x1d, 0xf8, 0x1a, 0xce, 0x78, 0x73, 0xf3, 0x77, 0x47, 0x6e, 0x63, 0x59,
	0x12, 0x34, 0xbb, 0x0f, 0xe2, 0x4b, 0x48, 0x2c, 0xa7, 0x96, 0x01, 0x81, 0x71, 0x07, 0xa4, 0x7f,
	0x22, 0x18, 0x4c, 0x39, 0xd7, 0x1b, 0xe5, 0x10, 0xe1, 0x44, 0x16, 0xa4, 0x5a, 0xa3, 0x42, 0xec,
	0xb1, 0xc2, 0x63, 0x8d, 0x57, 0x21, 0xf6, 0xf6, 0xd5, 0xa4, 0x5a, 0x87, 0x7c, 0xe8, 0x67, 0xac,
	0xc9, 0xae, 0x44, 0x39, 0x27, 0x15, 0x9c, 0x49, 0xf2, 0x3b, 0x00, 0x9f, 0x43, 0xbc, 0xb6, 0xd2,
	0x96, 0x2a, 0xa8, 0x9e, 0xe4, 0x6d, 0x86, 0x29, 0x0c, 0xbd, 0xfa, 0x5e, 0x95, 0xa0, 0x73, 0x92,
	0xef, 0xf2, 0xf4, 0x77, 0x0f, 0xe2, 0xc6, 0x48, 0xfc, 0x00, 0x43, 0x6a, 0x36, 0xb4, 0x2c, 0x1a,
	0xf5, 0x83, 0x48, 0x07, 0x5d, 0xf7, 0x94, 0x7c, 0xc7, 0xf5, 0x7d, 0x05, 0x55, 0xa4, 0xb8, 0xb0,
	0xac, 0x77, 0xbc, 0xaf, 0xe3, 0xe2, 0x67, 0x78, 0xe2, 0x0c, 0x29, 0x4b, 0x41, 0x69, 0xcb, 0xfa,
	0x47, 0x7b, 0xef, 0xf1, 0xf1, 0x2d, 0x3c, 0xa3, 0x2d, 0xc9, 0x8a, 0x8a, 0x4a, 0x4c, 0xbb, 0xc5,
	0x1b, 0x51, 0x1e, 0x16, 0x90, 0xc1, 0x80, 0xaa, 0x6a, 0x6e, 0xcb, 0x49, 0xab, 0x4e, 0x97, 0x4e,
	0xbe, 0x40, 0xf2, 0xcd, 0x3f, 0xc0, 0x6b, 0x52, 0x2b, 0xfc, 0x08, 0x67, 0x33, 0x23, 0xc8, 0x89,
	0xee, 0x4d, 0xe1, 0xc3, 0x7d, 0xd2, 0x03, 0xd8, 0xe5, 0xa3, 0x22, 0x0e, 0xbf, 0xd2, 0xfb, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x18, 0x1b, 0xd8, 0xcf, 0x03, 0x00, 0x00,
}
