package manager

import "git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_admin/common/cache"

const (
	UploadStatusUnkonwn	= 0
	UploadStatusPending	= 1
	UploadStatusSuccess	= 2
	UploadStatusFailed	= 3

	qrCodesGenerateChannel = "qrcode-generate-channel"
	qrCodesDownloadChannel = "qrcode-download-channel"
)

type UploadMessage struct {
	ID 			uint64 `json:"id"`
	Status 		uint32 `json:"status"`
	Percent 	uint32 `json:"percent"`
	DownloadUrl string `json:"download_url"`
	Email 		string `json:"email"`
	Message 	string `json:"result_message"`
}

func BuildUploadMesage(id uint64, status uint32, url, email, message string) *UploadMessage {
	msg := &UploadMessage{
		ID: id,
		Status: status,
		DownloadUrl: url,
		Email: email,
		Message: message,
	}
	if status == UploadStatusSuccess {
		msg.Percent = 100
	} else {
		msg.Percent = 0
	}
	return msg
}

// 发布上传成功信息
func PublishQrCodesGenerateMessage(value string) error {
	return cache.GetCacheImpl().PublishIntoRedis(qrCodesGenerateChannel, value)
}

func PublishQrCodesDownloadMessage(value string) error {
	return cache.GetCacheImpl().PublishIntoRedis(qrCodesDownloadChannel, value)
}
