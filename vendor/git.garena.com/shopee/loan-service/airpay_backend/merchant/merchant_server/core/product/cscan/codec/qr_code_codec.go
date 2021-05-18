package codec

import (
	"fmt"
	"strings"

	comutil "git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_server/common/util"
	"git.garena.com/shopee/loan-service/airpay_backend/public/common/utils"
	"github.com/howeyc/crc16"
)

const (
	CodePayloadFormatIndicator = "000201"
	CodeInitiationMethod       = "010211"
	CodeChecksumPrefix         = "6304"

	CodeCustomBusinessLength = 20
	CodeCustomRandomLength   = 20
)

const (
	CodeMerchantCategoryVN = "52045812"
	CodeMerchantNameVN     = "5910RESTAURANT"
	CodeMerchantCityVN     = "6009HOCHIMINH"
	CodeCountryVN          = "5802VN"
	CodeBusinessDataVN     = "01%s%s"
	CodeMerchantDomainVN   = "0013vn.airpay.www"
	CustomContentIndexVN   = "26"
)

const (
	CodeCountryTH            = "5802TH"
	CodeMerchantDataPrefixTH = "0016A90000000000000001031900203CsB03%s%s"
	CustomContentIndexTH     = "39"
)

type QrCodec interface {
	IsValidteCQrCode(qrCode string) bool
	GetQrCodeBody() string
	GenerateBusinessContent() string
}

type BaseQrCodec struct {
	RandomContent string
	BusinessData  string
}

type QrCodecForVN struct {
	BaseQrCodec
}

type QrCodecForTH struct {
	BaseQrCodec
}

func NewCodec(region string, randomContent string, businessData string) QrCodec {
	if randomContent == "" {
		randomContent = comutil.CryptUtil.RandomStringUseDefaultAllowedChars(CodeCustomRandomLength)
	}

	baseQrCodec := BaseQrCodec{
		RandomContent: randomContent,
		BusinessData:  businessData,
	}

	if region == utils.VN {
		return &QrCodecForVN{
			BaseQrCodec: baseQrCodec,
		}
	} else if region == utils.TH {
		return &QrCodecForTH{
			BaseQrCodec: baseQrCodec,
		}
	}
	return &baseQrCodec
}

func (c *BaseQrCodec) GenerateCRC16Sum(qrBody string) string {
	checkSum := crc16.ChecksumCCITTFalse([]byte(qrBody))
	return fmt.Sprintf("%04X", checkSum)
}

func (c *BaseQrCodec) IsValidteCQrCode(qrCode string) bool {
	prefix := CodePayloadFormatIndicator + CodeInitiationMethod
	if !strings.HasPrefix(qrCode, prefix) {
		return false
	}
	qrLen := len(qrCode)
	if c.GenerateCRC16Sum(qrCode[:qrLen-4]) == qrCode[qrLen-4:] {
		return true
	}
	return false
}

func (c *BaseQrCodec) GetQrCodeBody() string {
	panic("no need to implement")
}

func (c *BaseQrCodec) GenerateBusinessContent() string {
	panic("no need to implement")
}

func (c *BaseQrCodec) GenerateCustomContent(fieldIndex string, codec QrCodec) string {
	customContent := codec.GenerateBusinessContent()
	lenInStr := fmt.Sprintf("%02d", len(customContent))
	return fieldIndex + lenInStr + customContent
}

func (c *QrCodecForVN) GenerateBusinessContent() string {
	var businessData string
	if c.BusinessData != "" {
		businessData = c.BusinessData
	} else {
		businessData = "01000101010200000000"
	}
	businessDataLen := CodeCustomBusinessLength + CodeCustomRandomLength
	businessDataLenStr := fmt.Sprintf("%02d", businessDataLen)
	businessData = businessData + c.RandomContent
	formatBusinessData := fmt.Sprintf(CodeBusinessDataVN, businessDataLenStr, businessData)
	return CodeMerchantDomainVN + formatBusinessData
}

func (c *QrCodecForVN) GetMerchantAccountInformation() string {
	return c.GenerateCustomContent(CustomContentIndexVN, c)
}

func (c *QrCodecForVN) GetQrCodeBody() string {
	merchantInfo := c.GetMerchantAccountInformation()
	txnCurrency := c.GetTransactionCurrencyCode()
	return CodePayloadFormatIndicator + CodeInitiationMethod + merchantInfo + CodeMerchantCategoryVN +
		txnCurrency + CodeMerchantNameVN + CodeMerchantCityVN + CodeCountryVN + CodeChecksumPrefix

}

func (c *QrCodecForVN) GetTransactionCurrencyCode() string {
	return "5303704" //As defined by ISO 4217. For Vietnam Baht is "704"
}

func (c *QrCodecForTH) GetQrCodeBody() string {
	return CodePayloadFormatIndicator + CodeInitiationMethod + c.GenerateCustomContent(CustomContentIndexTH, c) +
		c.GetTransactionCurrencyCode() + CodeCountryTH + CodeChecksumPrefix
}

func (c *QrCodecForTH) GenerateBusinessContent() string {
	var businessData string
	if c.BusinessData != "" {
		businessData = c.BusinessData
	} else {
		businessData = "01000101010100000000"
	}
	businessDataLen := CodeCustomBusinessLength + CodeCustomRandomLength
	businessDataLenStr := fmt.Sprintf("%02d", businessDataLen)
	businessData = businessData + c.RandomContent
	return fmt.Sprintf(CodeMerchantDataPrefixTH, businessDataLenStr, businessData)
}

func (c *QrCodecForTH) GetTransactionCurrencyCode() string {
	return "5303764" //As defined by ISO 4217. For Thai Baht is "764"
}
