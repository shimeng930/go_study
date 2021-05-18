// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/pb2/merchant_common_pb2.proto

package merchant_pb_common_pb2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Merchant_ErrorCode int32

const (
	Merchant_ErrorCode_Success              Merchant_ErrorCode = 0
	Merchant_ErrorCode_ServerError          Merchant_ErrorCode = 30001
	Merchant_ErrorCode_InvalidParam         Merchant_ErrorCode = 30002
	Merchant_ErrorCode_MerchantNotExists    Merchant_ErrorCode = 30003
	Merchant_ErrorCode_MerchantTxnNotExists Merchant_ErrorCode = 30004
	Merchant_ErrorCode_MerchantTxnConfirmed Merchant_ErrorCode = 30005
	Merchant_ErrorCode_UserNotExists        Merchant_ErrorCode = 30006
	Merchant_ErrorCode_Timeout              Merchant_ErrorCode = 30007
	Merchant_ErrorCode_InvalidRequest       Merchant_ErrorCode = 32000
	Merchant_ErrorCode_InvalidClient        Merchant_ErrorCode = 32001
	Merchant_ErrorCode_InvalidGrant         Merchant_ErrorCode = 32002
	Merchant_ErrorCode_UnauthorizedClient   Merchant_ErrorCode = 32003
	Merchant_ErrorCode_UnsupportedGrantType Merchant_ErrorCode = 32004
	Merchant_ErrorCode_InvalidScope         Merchant_ErrorCode = 32005
	Merchant_ErrorCode_RequestExpired       Merchant_ErrorCode = 32006
)

var Merchant_ErrorCode_name = map[int32]string{
	0:     "ErrorCode_Success",
	30001: "ErrorCode_ServerError",
	30002: "ErrorCode_InvalidParam",
	30003: "ErrorCode_MerchantNotExists",
	30004: "ErrorCode_MerchantTxnNotExists",
	30005: "ErrorCode_MerchantTxnConfirmed",
	30006: "ErrorCode_UserNotExists",
	30007: "ErrorCode_Timeout",
	32000: "ErrorCode_InvalidRequest",
	32001: "ErrorCode_InvalidClient",
	32002: "ErrorCode_InvalidGrant",
	32003: "ErrorCode_UnauthorizedClient",
	32004: "ErrorCode_UnsupportedGrantType",
	32005: "ErrorCode_InvalidScope",
	32006: "ErrorCode_RequestExpired",
}

var Merchant_ErrorCode_value = map[string]int32{
	"ErrorCode_Success":              0,
	"ErrorCode_ServerError":          30001,
	"ErrorCode_InvalidParam":         30002,
	"ErrorCode_MerchantNotExists":    30003,
	"ErrorCode_MerchantTxnNotExists": 30004,
	"ErrorCode_MerchantTxnConfirmed": 30005,
	"ErrorCode_UserNotExists":        30006,
	"ErrorCode_Timeout":              30007,
	"ErrorCode_InvalidRequest":       32000,
	"ErrorCode_InvalidClient":        32001,
	"ErrorCode_InvalidGrant":         32002,
	"ErrorCode_UnauthorizedClient":   32003,
	"ErrorCode_UnsupportedGrantType": 32004,
	"ErrorCode_InvalidScope":         32005,
	"ErrorCode_RequestExpired":       32006,
}

func (x Merchant_ErrorCode) Enum() *Merchant_ErrorCode {
	p := new(Merchant_ErrorCode)
	*p = x
	return p
}

func (x Merchant_ErrorCode) String() string {
	return proto.EnumName(Merchant_ErrorCode_name, int32(x))
}

func (x *Merchant_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Merchant_ErrorCode_value, data, "Merchant_ErrorCode")
	if err != nil {
		return err
	}
	*x = Merchant_ErrorCode(value)
	return nil
}

func (Merchant_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0, 0}
}

type Merchant_ProductType int32

const (
	Merchant_ProductType_None            Merchant_ProductType = 0
	Merchant_ProductType_Jump_App_Pay    Merchant_ProductType = 1
	Merchant_ProductType_Link_And_Pay    Merchant_ProductType = 2
	Merchant_ProductType_Web_Pay         Merchant_ProductType = 3
	Merchant_ProductType_Auth_Direct_Pay Merchant_ProductType = 4
	Merchant_ProductType_Native_App_Pay  Merchant_ProductType = 5
	Merchant_ProductType_H5_App_Pay      Merchant_ProductType = 6
	Merchant_ProductType_Scan_QR_Pay     Merchant_ProductType = 7
	Merchant_ProductType_Static_QR_Pay   Merchant_ProductType = 8
	Merchant_ProductType_Account_Link    Merchant_ProductType = 9
	Merchant_ProductType_OAuth           Merchant_ProductType = 10
	Merchant_ProductType_Dynamic_QR_Pay  Merchant_ProductType = 11
	Merchant_ProductType_P2p_Transfer    Merchant_ProductType = 12
	Merchant_ProductType_Shopee_API      Merchant_ProductType = 13
	Merchant_ProductType_New_Auth_Pay    Merchant_ProductType = 14
	Merchant_ProductType_Batch_Transfer  Merchant_ProductType = 15
	Merchant_ProductType_Offline_TopUp   Merchant_ProductType = 16
)

var Merchant_ProductType_name = map[int32]string{
	0:  "ProductType_None",
	1:  "ProductType_Jump_App_Pay",
	2:  "ProductType_Link_And_Pay",
	3:  "ProductType_Web_Pay",
	4:  "ProductType_Auth_Direct_Pay",
	5:  "ProductType_Native_App_Pay",
	6:  "ProductType_H5_App_Pay",
	7:  "ProductType_Scan_QR_Pay",
	8:  "ProductType_Static_QR_Pay",
	9:  "ProductType_Account_Link",
	10: "ProductType_OAuth",
	11: "ProductType_Dynamic_QR_Pay",
	12: "ProductType_P2p_Transfer",
	13: "ProductType_Shopee_API",
	14: "ProductType_New_Auth_Pay",
	15: "ProductType_Batch_Transfer",
	16: "ProductType_Offline_TopUp",
}

var Merchant_ProductType_value = map[string]int32{
	"ProductType_None":            0,
	"ProductType_Jump_App_Pay":    1,
	"ProductType_Link_And_Pay":    2,
	"ProductType_Web_Pay":         3,
	"ProductType_Auth_Direct_Pay": 4,
	"ProductType_Native_App_Pay":  5,
	"ProductType_H5_App_Pay":      6,
	"ProductType_Scan_QR_Pay":     7,
	"ProductType_Static_QR_Pay":   8,
	"ProductType_Account_Link":    9,
	"ProductType_OAuth":           10,
	"ProductType_Dynamic_QR_Pay":  11,
	"ProductType_P2p_Transfer":    12,
	"ProductType_Shopee_API":      13,
	"ProductType_New_Auth_Pay":    14,
	"ProductType_Batch_Transfer":  15,
	"ProductType_Offline_TopUp":   16,
}

func (x Merchant_ProductType) Enum() *Merchant_ProductType {
	p := new(Merchant_ProductType)
	*p = x
	return p
}

func (x Merchant_ProductType) String() string {
	return proto.EnumName(Merchant_ProductType_name, int32(x))
}

func (x *Merchant_ProductType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Merchant_ProductType_value, data, "Merchant_ProductType")
	if err != nil {
		return err
	}
	*x = Merchant_ProductType(value)
	return nil
}

func (Merchant_ProductType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0, 1}
}

type Merchant_PaymentMethod int32

const (
	Merchant_PaymentMethod_None           Merchant_PaymentMethod = 0
	Merchant_PaymentMethod_Credit_Card    Merchant_PaymentMethod = 1
	Merchant_PaymentMethod_Debit_Card     Merchant_PaymentMethod = 2
	Merchant_PaymentMethod_Wallet_Balance Merchant_PaymentMethod = 3
)

var Merchant_PaymentMethod_name = map[int32]string{
	0: "PaymentMethod_None",
	1: "PaymentMethod_Credit_Card",
	2: "PaymentMethod_Debit_Card",
	3: "PaymentMethod_Wallet_Balance",
}

var Merchant_PaymentMethod_value = map[string]int32{
	"PaymentMethod_None":           0,
	"PaymentMethod_Credit_Card":    1,
	"PaymentMethod_Debit_Card":     2,
	"PaymentMethod_Wallet_Balance": 3,
}

func (x Merchant_PaymentMethod) Enum() *Merchant_PaymentMethod {
	p := new(Merchant_PaymentMethod)
	*p = x
	return p
}

func (x Merchant_PaymentMethod) String() string {
	return proto.EnumName(Merchant_PaymentMethod_name, int32(x))
}

func (x *Merchant_PaymentMethod) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Merchant_PaymentMethod_value, data, "Merchant_PaymentMethod")
	if err != nil {
		return err
	}
	*x = Merchant_PaymentMethod(value)
	return nil
}

func (Merchant_PaymentMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0, 2}
}

type Merchant_OrderScene int32

const (
	Merchant_OrderScene_None           Merchant_OrderScene = 0
	Merchant_OrderScene_BScanC         Merchant_OrderScene = 10
	Merchant_OrderScene_CScanB         Merchant_OrderScene = 11
	Merchant_OrderScene_BScanC_NEW     Merchant_OrderScene = 12
	Merchant_OrderScene_Standard_API   Merchant_OrderScene = 30
	Merchant_OrderScene_Foody_API      Merchant_OrderScene = 50
	Merchant_OrderScene_P2P_TRANSFER   Merchant_OrderScene = 60
	Merchant_OrderScene_Google_Play    Merchant_OrderScene = 70
	Merchant_OrderScene_Shopee_API     Merchant_OrderScene = 80
	Merchant_OrderScene_Batch_TRANSFER Merchant_OrderScene = 90
	Merchant_OrderScene_TopUp          Merchant_OrderScene = 100
)

var Merchant_OrderScene_name = map[int32]string{
	0:   "OrderScene_None",
	10:  "OrderScene_BScanC",
	11:  "OrderScene_CScanB",
	12:  "OrderScene_BScanC_NEW",
	30:  "OrderScene_Standard_API",
	50:  "OrderScene_Foody_API",
	60:  "OrderScene_P2P_TRANSFER",
	70:  "OrderScene_Google_Play",
	80:  "OrderScene_Shopee_API",
	90:  "OrderScene_Batch_TRANSFER",
	100: "OrderScene_TopUp",
}

var Merchant_OrderScene_value = map[string]int32{
	"OrderScene_None":           0,
	"OrderScene_BScanC":         10,
	"OrderScene_CScanB":         11,
	"OrderScene_BScanC_NEW":     12,
	"OrderScene_Standard_API":   30,
	"OrderScene_Foody_API":      50,
	"OrderScene_P2P_TRANSFER":   60,
	"OrderScene_Google_Play":    70,
	"OrderScene_Shopee_API":     80,
	"OrderScene_Batch_TRANSFER": 90,
	"OrderScene_TopUp":          100,
}

func (x Merchant_OrderScene) Enum() *Merchant_OrderScene {
	p := new(Merchant_OrderScene)
	*p = x
	return p
}

func (x Merchant_OrderScene) String() string {
	return proto.EnumName(Merchant_OrderScene_name, int32(x))
}

func (x *Merchant_OrderScene) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Merchant_OrderScene_value, data, "Merchant_OrderScene")
	if err != nil {
		return err
	}
	*x = Merchant_OrderScene(value)
	return nil
}

func (Merchant_OrderScene) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0, 3}
}

type Merchant_UnifyOrderingType int32

const (
	Merchant_UnifyOrderingType_None            Merchant_UnifyOrderingType = 0
	Merchant_UnifyOrderingType_GetLandingUrlV1 Merchant_UnifyOrderingType = 1
	Merchant_UnifyOrderingType_GetLandingUrlV2 Merchant_UnifyOrderingType = 2
	Merchant_UnifyOrderingType_DynamicQRCode   Merchant_UnifyOrderingType = 3
)

var Merchant_UnifyOrderingType_name = map[int32]string{
	0: "UnifyOrderingType_None",
	1: "UnifyOrderingType_GetLandingUrlV1",
	2: "UnifyOrderingType_GetLandingUrlV2",
	3: "UnifyOrderingType_DynamicQRCode",
}

var Merchant_UnifyOrderingType_value = map[string]int32{
	"UnifyOrderingType_None":            0,
	"UnifyOrderingType_GetLandingUrlV1": 1,
	"UnifyOrderingType_GetLandingUrlV2": 2,
	"UnifyOrderingType_DynamicQRCode":   3,
}

func (x Merchant_UnifyOrderingType) Enum() *Merchant_UnifyOrderingType {
	p := new(Merchant_UnifyOrderingType)
	*p = x
	return p
}

func (x Merchant_UnifyOrderingType) String() string {
	return proto.EnumName(Merchant_UnifyOrderingType_name, int32(x))
}

func (x *Merchant_UnifyOrderingType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Merchant_UnifyOrderingType_value, data, "Merchant_UnifyOrderingType")
	if err != nil {
		return err
	}
	*x = Merchant_UnifyOrderingType(value)
	return nil
}

func (Merchant_UnifyOrderingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0, 4}
}

type Merchant_PaymentSource int32

const (
	Merchant_PaymentSource_None       Merchant_PaymentSource = 0
	Merchant_PaymentSource_AirPay     Merchant_PaymentSource = 1
	Merchant_PaymentSource_Shopee     Merchant_PaymentSource = 2
	Merchant_PaymentSource_SDK        Merchant_PaymentSource = 3
	Merchant_PaymentSource_S2S        Merchant_PaymentSource = 4
	Merchant_PaymentSource_DP         Merchant_PaymentSource = 5
	Merchant_PaymentSource_Foody      Merchant_PaymentSource = 6
	Merchant_PaymentSource_Shopee_API Merchant_PaymentSource = 7
	Merchant_PaymentSource_IS         Merchant_PaymentSource = 8
)

var Merchant_PaymentSource_name = map[int32]string{
	0: "PaymentSource_None",
	1: "PaymentSource_AirPay",
	2: "PaymentSource_Shopee",
	3: "PaymentSource_SDK",
	4: "PaymentSource_S2S",
	5: "PaymentSource_DP",
	6: "PaymentSource_Foody",
	7: "PaymentSource_Shopee_API",
	8: "PaymentSource_IS",
}

var Merchant_PaymentSource_value = map[string]int32{
	"PaymentSource_None":       0,
	"PaymentSource_AirPay":     1,
	"PaymentSource_Shopee":     2,
	"PaymentSource_SDK":        3,
	"PaymentSource_S2S":        4,
	"PaymentSource_DP":         5,
	"PaymentSource_Foody":      6,
	"PaymentSource_Shopee_API": 7,
	"PaymentSource_IS":         8,
}

func (x Merchant_PaymentSource) Enum() *Merchant_PaymentSource {
	p := new(Merchant_PaymentSource)
	*p = x
	return p
}

func (x Merchant_PaymentSource) String() string {
	return proto.EnumName(Merchant_PaymentSource_name, int32(x))
}

func (x *Merchant_PaymentSource) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Merchant_PaymentSource_value, data, "Merchant_PaymentSource")
	if err != nil {
		return err
	}
	*x = Merchant_PaymentSource(value)
	return nil
}

func (Merchant_PaymentSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0, 5}
}

type Merchant struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Merchant) Reset()         { *m = Merchant{} }
func (m *Merchant) String() string { return proto.CompactTextString(m) }
func (*Merchant) ProtoMessage()    {}
func (*Merchant) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0}
}

func (m *Merchant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Merchant.Unmarshal(m, b)
}
func (m *Merchant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Merchant.Marshal(b, m, deterministic)
}
func (m *Merchant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Merchant.Merge(m, src)
}
func (m *Merchant) XXX_Size() int {
	return xxx_messageInfo_Merchant.Size(m)
}
func (m *Merchant) XXX_DiscardUnknown() {
	xxx_messageInfo_Merchant.DiscardUnknown(m)
}

var xxx_messageInfo_Merchant proto.InternalMessageInfo

type Merchant_ResponseHeader struct {
	ErrorCode            *Merchant_ErrorCode `protobuf:"varint,1,req,name=error_code,json=errorCode,enum=merchant.pb.common.pb2.Merchant_ErrorCode" json:"error_code,omitempty"`
	Message              *string             `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Merchant_ResponseHeader) Reset()         { *m = Merchant_ResponseHeader{} }
func (m *Merchant_ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*Merchant_ResponseHeader) ProtoMessage()    {}
func (*Merchant_ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0, 0}
}

func (m *Merchant_ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Merchant_ResponseHeader.Unmarshal(m, b)
}
func (m *Merchant_ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Merchant_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *Merchant_ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Merchant_ResponseHeader.Merge(m, src)
}
func (m *Merchant_ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_Merchant_ResponseHeader.Size(m)
}
func (m *Merchant_ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_Merchant_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_Merchant_ResponseHeader proto.InternalMessageInfo

func (m *Merchant_ResponseHeader) GetErrorCode() Merchant_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return Merchant_ErrorCode_Success
}

func (m *Merchant_ResponseHeader) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type Merchant_VoidResponse struct {
	Header               *Merchant_ResponseHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Merchant_VoidResponse) Reset()         { *m = Merchant_VoidResponse{} }
func (m *Merchant_VoidResponse) String() string { return proto.CompactTextString(m) }
func (*Merchant_VoidResponse) ProtoMessage()    {}
func (*Merchant_VoidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0, 1}
}

func (m *Merchant_VoidResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Merchant_VoidResponse.Unmarshal(m, b)
}
func (m *Merchant_VoidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Merchant_VoidResponse.Marshal(b, m, deterministic)
}
func (m *Merchant_VoidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Merchant_VoidResponse.Merge(m, src)
}
func (m *Merchant_VoidResponse) XXX_Size() int {
	return xxx_messageInfo_Merchant_VoidResponse.Size(m)
}
func (m *Merchant_VoidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Merchant_VoidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Merchant_VoidResponse proto.InternalMessageInfo

func (m *Merchant_VoidResponse) GetHeader() *Merchant_ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type Merchant_UInt64IdsRequest struct {
	Ids                  []uint64 `protobuf:"varint,1,rep,name=ids" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Merchant_UInt64IdsRequest) Reset()         { *m = Merchant_UInt64IdsRequest{} }
func (m *Merchant_UInt64IdsRequest) String() string { return proto.CompactTextString(m) }
func (*Merchant_UInt64IdsRequest) ProtoMessage()    {}
func (*Merchant_UInt64IdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0, 2}
}

func (m *Merchant_UInt64IdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Merchant_UInt64IdsRequest.Unmarshal(m, b)
}
func (m *Merchant_UInt64IdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Merchant_UInt64IdsRequest.Marshal(b, m, deterministic)
}
func (m *Merchant_UInt64IdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Merchant_UInt64IdsRequest.Merge(m, src)
}
func (m *Merchant_UInt64IdsRequest) XXX_Size() int {
	return xxx_messageInfo_Merchant_UInt64IdsRequest.Size(m)
}
func (m *Merchant_UInt64IdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Merchant_UInt64IdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Merchant_UInt64IdsRequest proto.InternalMessageInfo

func (m *Merchant_UInt64IdsRequest) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Merchant_BoolResponse struct {
	Header               *Merchant_ResponseHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Value                *bool                    `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Merchant_BoolResponse) Reset()         { *m = Merchant_BoolResponse{} }
func (m *Merchant_BoolResponse) String() string { return proto.CompactTextString(m) }
func (*Merchant_BoolResponse) ProtoMessage()    {}
func (*Merchant_BoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_10cfd20c64a3cbe7, []int{0, 3}
}

func (m *Merchant_BoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Merchant_BoolResponse.Unmarshal(m, b)
}
func (m *Merchant_BoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Merchant_BoolResponse.Marshal(b, m, deterministic)
}
func (m *Merchant_BoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Merchant_BoolResponse.Merge(m, src)
}
func (m *Merchant_BoolResponse) XXX_Size() int {
	return xxx_messageInfo_Merchant_BoolResponse.Size(m)
}
func (m *Merchant_BoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Merchant_BoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Merchant_BoolResponse proto.InternalMessageInfo

func (m *Merchant_BoolResponse) GetHeader() *Merchant_ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Merchant_BoolResponse) GetValue() bool {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return false
}

func init() {
	proto.RegisterEnum("merchant.pb.common.pb2.Merchant_ErrorCode", Merchant_ErrorCode_name, Merchant_ErrorCode_value)
	proto.RegisterEnum("merchant.pb.common.pb2.Merchant_ProductType", Merchant_ProductType_name, Merchant_ProductType_value)
	proto.RegisterEnum("merchant.pb.common.pb2.Merchant_PaymentMethod", Merchant_PaymentMethod_name, Merchant_PaymentMethod_value)
	proto.RegisterEnum("merchant.pb.common.pb2.Merchant_OrderScene", Merchant_OrderScene_name, Merchant_OrderScene_value)
	proto.RegisterEnum("merchant.pb.common.pb2.Merchant_UnifyOrderingType", Merchant_UnifyOrderingType_name, Merchant_UnifyOrderingType_value)
	proto.RegisterEnum("merchant.pb.common.pb2.Merchant_PaymentSource", Merchant_PaymentSource_name, Merchant_PaymentSource_value)
	proto.RegisterType((*Merchant)(nil), "merchant.pb.common.pb2.Merchant")
	proto.RegisterType((*Merchant_ResponseHeader)(nil), "merchant.pb.common.pb2.Merchant.ResponseHeader")
	proto.RegisterType((*Merchant_VoidResponse)(nil), "merchant.pb.common.pb2.Merchant.VoidResponse")
	proto.RegisterType((*Merchant_UInt64IdsRequest)(nil), "merchant.pb.common.pb2.Merchant.UInt64IdsRequest")
	proto.RegisterType((*Merchant_BoolResponse)(nil), "merchant.pb.common.pb2.Merchant.BoolResponse")
}

func init() {
	proto.RegisterFile("common/pb2/merchant_common_pb2.proto", fileDescriptor_10cfd20c64a3cbe7)
}

var fileDescriptor_10cfd20c64a3cbe7 = []byte{
	// 990 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4b, 0x6f, 0xdb, 0x46,
	0x10, 0xc7, 0x43, 0xc9, 0xcf, 0xb1, 0xe3, 0xac, 0x37, 0x7e, 0x28, 0xb2, 0xe3, 0x38, 0xaa, 0x0b,
	0x18, 0x3d, 0xc8, 0xa8, 0xd0, 0xf6, 0xd4, 0x8b, 0x2c, 0x3f, 0xa2, 0x36, 0xb1, 0x15, 0x4a, 0x8a,
	0x81, 0x5e, 0x88, 0x15, 0x39, 0xb6, 0x88, 0x92, 0xbb, 0xec, 0x72, 0xe9, 0x5a, 0x3d, 0xf5, 0x99,
	0x0f, 0x52, 0xa0, 0x97, 0xde, 0xfa, 0xfe, 0x4c, 0x15, 0xfa, 0x09, 0x04, 0x14, 0x28, 0x48, 0x8a,
	0x16, 0x49, 0xb9, 0x8f, 0x43, 0x6f, 0xdc, 0xff, 0x6f, 0xf8, 0x9f, 0x99, 0x9d, 0xdd, 0x85, 0x3d,
	0x53, 0xb8, 0xae, 0xe0, 0x07, 0x5e, 0xaf, 0x76, 0xe0, 0xa2, 0x34, 0xfb, 0x8c, 0x2b, 0x23, 0xd6,
	0x0c, 0xaf, 0x57, 0xab, 0x7a, 0x52, 0x28, 0x41, 0x37, 0x12, 0x54, 0xf5, 0x7a, 0xd5, 0x98, 0x56,
	0xbd, 0x5e, 0xad, 0xf2, 0xdd, 0x2a, 0x2c, 0xbc, 0x18, 0xa3, 0x72, 0x00, 0x2b, 0x3a, 0xfa, 0x9e,
	0xe0, 0x3e, 0x3e, 0x43, 0x66, 0xa1, 0xa4, 0x4d, 0x00, 0x94, 0x52, 0x48, 0xc3, 0x14, 0x16, 0x96,
	0xb4, 0xdd, 0xc2, 0xfe, 0x4a, 0xed, 0xad, 0xea, 0xdd, 0x5e, 0xd5, 0xc4, 0xa7, 0x7a, 0x1c, 0xfe,
	0xd2, 0x10, 0x16, 0xea, 0x8b, 0x98, 0x7c, 0xd2, 0x12, 0xcc, 0xbb, 0xe8, 0xfb, 0xec, 0x0a, 0x4b,
	0x85, 0x5d, 0x6d, 0x7f, 0x51, 0x4f, 0x96, 0xe5, 0x0b, 0x58, 0x7e, 0x25, 0x6c, 0x2b, 0x49, 0x4d,
	0x4f, 0x61, 0xae, 0x1f, 0xa5, 0x8f, 0x12, 0x2e, 0xd5, 0x0e, 0xfe, 0x35, 0x61, 0xb6, 0x6a, 0x7d,
	0xfc, 0x7b, 0x79, 0x0f, 0x48, 0xb7, 0xc9, 0xd5, 0x7b, 0xef, 0x34, 0x2d, 0x5f, 0xc7, 0x4f, 0x02,
	0xf4, 0x15, 0x25, 0x50, 0xb4, 0x2d, 0xbf, 0xa4, 0xed, 0x16, 0xf7, 0x67, 0xf4, 0xf0, 0xb3, 0xec,
	0xc2, 0xf2, 0xa1, 0x10, 0xce, 0xff, 0x9e, 0x9e, 0xae, 0xc1, 0xec, 0x35, 0x73, 0x82, 0xb8, 0xdf,
	0x05, 0x3d, 0x5e, 0x54, 0x7e, 0x2f, 0xc2, 0xe2, 0xed, 0x06, 0xd1, 0x75, 0x58, 0xbd, 0x5d, 0x18,
	0xed, 0xc0, 0x34, 0xd1, 0xf7, 0xc9, 0x3d, 0xba, 0x05, 0xeb, 0x29, 0x19, 0xe5, 0x35, 0xca, 0x68,
	0x4d, 0x7e, 0x18, 0x6a, 0x74, 0x1b, 0x36, 0x26, 0xb0, 0xc9, 0xaf, 0x99, 0x63, 0x5b, 0x2d, 0x26,
	0x99, 0x4b, 0x7e, 0x1c, 0x6a, 0xf4, 0x29, 0x6c, 0x4d, 0x68, 0x52, 0xe2, 0x99, 0x50, 0xc7, 0x37,
	0xb6, 0xaf, 0x7c, 0xf2, 0xd3, 0x50, 0xa3, 0x7b, 0xb0, 0x33, 0x1d, 0xd2, 0xb9, 0xe1, 0x93, 0xa8,
	0x9f, 0xff, 0x21, 0xaa, 0x21, 0xf8, 0xa5, 0x2d, 0x5d, 0xb4, 0xc8, 0x2f, 0x43, 0x8d, 0x3e, 0x86,
	0xcd, 0x49, 0x54, 0xd7, 0x47, 0x39, 0x31, 0xf9, 0x75, 0xa8, 0xd1, 0xcd, 0x74, 0x7f, 0x1d, 0xdb,
	0x45, 0x11, 0x28, 0xf2, 0xdb, 0x50, 0xa3, 0x3b, 0x50, 0x9a, 0x6a, 0x62, 0x3c, 0x23, 0xf2, 0xf9,
	0x28, 0xe7, 0x3b, 0xe6, 0x0d, 0xc7, 0x46, 0xae, 0xc8, 0x17, 0xa3, 0xbb, 0xf7, 0xe0, 0x54, 0x32,
	0xae, 0xc8, 0x97, 0x23, 0x8d, 0x56, 0x60, 0x3b, 0x55, 0x14, 0x67, 0x81, 0xea, 0x0b, 0x69, 0x7f,
	0x86, 0x89, 0xc3, 0x57, 0xa3, 0x5c, 0x7b, 0x5d, 0xee, 0x07, 0x9e, 0x27, 0xa4, 0xc2, 0xd8, 0xa5,
	0x33, 0xf0, 0x90, 0x7c, 0xfd, 0x37, 0x79, 0xda, 0xa6, 0xf0, 0x90, 0x7c, 0x33, 0xca, 0x35, 0x31,
	0xae, 0xfe, 0xf8, 0xc6, 0xb3, 0x25, 0x5a, 0xe4, 0xf5, 0x48, 0xab, 0xfc, 0x59, 0x84, 0xa5, 0x96,
	0x14, 0x56, 0x60, 0x46, 0x8e, 0x74, 0x0d, 0x48, 0x6a, 0x69, 0x9c, 0x09, 0x8e, 0xe4, 0x1e, 0xdd,
	0x86, 0x52, 0x5a, 0xfd, 0x20, 0x70, 0x3d, 0xa3, 0xee, 0x79, 0x46, 0x8b, 0x0d, 0x88, 0x96, 0xa7,
	0xcf, 0x6d, 0xfe, 0xb1, 0x51, 0xe7, 0x56, 0x44, 0x0b, 0x74, 0x13, 0x1e, 0xa6, 0xe9, 0x05, 0xf6,
	0x22, 0x50, 0xa4, 0x4f, 0x60, 0x2b, 0x0d, 0xea, 0x81, 0xea, 0x1b, 0x47, 0xb6, 0x44, 0x53, 0x45,
	0x01, 0x33, 0x74, 0x07, 0xca, 0x99, 0x5a, 0x98, 0xb2, 0xaf, 0xf1, 0x36, 0xef, 0x2c, 0x2d, 0xc3,
	0x46, 0x9a, 0x3f, 0x7b, 0xf7, 0x96, 0xcd, 0xd1, 0x2d, 0xd8, 0x4c, 0xb3, 0xb6, 0xc9, 0xb8, 0xf1,
	0x52, 0x8f, 0xe0, 0x3c, 0x7d, 0x0c, 0x8f, 0x32, 0x50, 0x31, 0x65, 0x9b, 0x09, 0x5e, 0xc8, 0xf7,
	0x53, 0x37, 0x4d, 0x11, 0x70, 0x15, 0xf5, 0x45, 0x16, 0xc3, 0xfb, 0x90, 0xa6, 0xe7, 0x61, 0xdd,
	0x04, 0xf2, 0xc5, 0x1e, 0x0d, 0x38, 0x73, 0x27, 0xa6, 0x4b, 0x79, 0xd3, 0x56, 0xcd, 0x33, 0x3a,
	0x92, 0x71, 0xff, 0x12, 0x25, 0x59, 0xce, 0xb7, 0xd2, 0xee, 0x0b, 0x0f, 0xd1, 0xa8, 0xb7, 0x9a,
	0xe4, 0x7e, 0xfe, 0xcf, 0x33, 0xfc, 0x34, 0xde, 0xab, 0xd0, 0x77, 0x25, 0x9f, 0xf7, 0x90, 0x29,
	0xb3, 0x3f, 0x71, 0x7e, 0x90, 0xef, 0xf5, 0xfc, 0xf2, 0xd2, 0xb1, 0x39, 0x1a, 0x1d, 0xe1, 0x75,
	0x3d, 0x42, 0x2a, 0xaf, 0x35, 0xb8, 0xdf, 0x62, 0x03, 0x17, 0xb9, 0x7a, 0x81, 0xaa, 0x2f, 0x2c,
	0xba, 0x01, 0x34, 0x23, 0x24, 0x67, 0x20, 0x34, 0xca, 0xe8, 0x0d, 0x89, 0x96, 0xad, 0x8c, 0x06,
	0x93, 0xd6, 0xf8, 0x10, 0x64, 0xf0, 0x11, 0xf6, 0x12, 0x5a, 0xa0, 0xbb, 0xb0, 0x9d, 0xa5, 0x17,
	0xcc, 0x71, 0x50, 0x19, 0x87, 0xcc, 0x61, 0xdc, 0x44, 0x52, 0xac, 0x7c, 0x5f, 0x00, 0x38, 0x97,
	0x16, 0xca, 0xb6, 0x89, 0x1c, 0xe9, 0x43, 0x78, 0x30, 0x59, 0x25, 0x25, 0xac, 0xc3, 0x6a, 0x4a,
	0x3c, 0x0c, 0x87, 0xda, 0x20, 0x90, 0x93, 0x1b, 0xa1, 0x7c, 0x48, 0x96, 0xe8, 0x23, 0x58, 0x9f,
	0x8a, 0x36, 0xce, 0x8e, 0x2f, 0xc8, 0x72, 0x78, 0x3a, 0x52, 0xa8, 0xad, 0x18, 0xb7, 0x98, 0xb4,
	0xa2, 0xfd, 0xde, 0xa1, 0x25, 0x58, 0x4b, 0xc1, 0x13, 0x21, 0xac, 0x41, 0x44, 0x6a, 0xb9, 0xdf,
	0x5a, 0xb5, 0x96, 0xd1, 0xd1, 0xeb, 0x67, 0xed, 0x93, 0x63, 0x9d, 0xbc, 0x1f, 0x8e, 0x30, 0x05,
	0x4f, 0x85, 0xb8, 0x72, 0xd0, 0x68, 0x39, 0x6c, 0x40, 0x4e, 0x72, 0xa5, 0xa4, 0xa6, 0xdb, 0x0a,
	0xb7, 0x35, 0x5d, 0x65, 0x3c, 0xbe, 0xc4, 0xf5, 0xa3, 0xf0, 0x3e, 0xa6, 0x70, 0x3c, 0x35, 0xab,
	0xf2, 0xad, 0x06, 0xab, 0x5d, 0x6e, 0x5f, 0x0e, 0x22, 0x66, 0xf3, 0xab, 0xe8, 0xee, 0x96, 0x61,
	0x63, 0x4a, 0x4c, 0xb6, 0xee, 0x4d, 0x78, 0x3a, 0xcd, 0x4e, 0x51, 0x3d, 0x67, 0xdc, 0xb2, 0xf9,
	0x55, 0x57, 0x3a, 0xaf, 0xde, 0x26, 0xda, 0x7f, 0x09, 0xab, 0x91, 0x02, 0x7d, 0x03, 0x9e, 0x4c,
	0x87, 0x8d, 0x8f, 0xfc, 0x4b, 0x3d, 0x7c, 0x6b, 0x48, 0xb1, 0xf2, 0xc7, 0xe4, 0x68, 0xb5, 0x45,
	0x20, 0x4d, 0x4c, 0x1d, 0xad, 0x58, 0x48, 0x8a, 0x2b, 0xc1, 0x5a, 0x56, 0xaf, 0xdb, 0x32, 0x7e,
	0x5a, 0xa6, 0x48, 0xbc, 0x77, 0xa4, 0x10, 0x5d, 0xc3, 0x2c, 0x39, 0xfa, 0x90, 0x14, 0xef, 0x90,
	0x6b, 0x6d, 0x32, 0x13, 0x3d, 0x6b, 0x19, 0xf9, 0xa8, 0x45, 0x66, 0xa3, 0xa7, 0x29, 0xa3, 0x46,
	0xc3, 0x26, 0x73, 0xa9, 0xc3, 0x9c, 0x49, 0x1b, 0x8d, 0x6c, 0x7e, 0xda, 0xac, 0xd9, 0x26, 0x0b,
	0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x4a, 0x51, 0x97, 0xe6, 0x08, 0x00, 0x00,
}
