package common_util

import (
	PBC "git.garena.com/shopee/loan-service/airpay_backend/public/merchant/merchant_core_proto/protobuf/common/pb2"
)

func GenMerchantHeader(code PBC.Merchant_ErrorCode, msg string) *PBC.Merchant_ResponseHeader {
	return &PBC.Merchant_ResponseHeader{
		ErrorCode: &code,
		Message:   &msg,
	}
}
