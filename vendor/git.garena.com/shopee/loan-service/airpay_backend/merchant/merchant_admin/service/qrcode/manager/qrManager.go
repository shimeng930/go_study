package manager

import (
	"context"
	"fmt"
	"git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_admin/common/util"
	"git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_admin/service/qrcode/utils"
	"os"
)

type QrCodeUploadModel struct {
	QrID 		uint64
	QrToken		string
	OutletID 	uint64
	OutletName 	string
}

type QrCodeManager struct {
	Ctx context.Context
	Log util.Log
}

func (q *QrCodeManager) GetQrCodePngFiles(qrCodes []QrCodeUploadModel) ([]string, error) {
	if qrCodes == nil {
		return nil, nil
	}
	//qrNumber := len(qrCodes)

	var qrCodesPath []string
	qrUtil := utils.NewQrCodePng(0, 0, "", "")
	//ch := make(chan string)
	for _, item := range qrCodes {
		qrName := fmt.Sprintf("empty_qr_%d.png", item.QrID)
		qrUtil.SetQrID(item.QrID)
		qrUtil.SetQrToken(item.QrToken)
		if item.OutletName != "" {
			qrUtil.SetOutletName(item.OutletName)
		}
		if item.OutletID != 0 {
			qrUtil.SetOutletID(item.OutletID)
			qrName = fmt.Sprintf("outlet_%d.png", item.OutletID)
		}
		qrUtil.SetQrID(item.QrID)

		qrPath, err := qrUtil.FileQrCodeWithBg(qrName)
		if err != nil {
			q.Log.Errorf("GetQrCodePngFiles|FileQrCodeWithBg error|qrID=%d, err=%v", item.QrID, err)
		}
		qrCodesPath = append(qrCodesPath, qrPath)

		//go func(item QrCodeUploadModel) {
		//	qrUtil.SetQrID(item.QrID)
		//	qrUtil.SetQrToken(item.QrToken)
		//	if item.OutletName != "" {
		//		qrUtil.SetOutletName(item.OutletName)
		//	}
		//	qrUtil.SetQrID(item.QrID)
		//	qrPath, err := qrUtil.FileQrCodeWithBg(fmt.Sprintf("cqr_%d.png", item.QrID))
		//	if err != nil {
		//		q.Log.Errorf("GetQrCodePngFiles|FileQrCodeWithBg error|qrID=%d, err=%v", item.QrID, err)
		//	}
		//	ch <- qrPath
		//}(item)
	}

	//var chData int
	//for data := range ch {
	//	chData++
	//	qrCodesPath = append(qrCodesPath, data)
	//	if chData == qrNumber {
	//		break
	//	}
	//}

	return qrCodesPath, nil
}

//func (q *QrCodeManager) GetQrCodeZip(qrCodePath []string) (string, error) {
//	var targetPath string
//	return targetPath, utils.Compress(qrCodePath, targetPath)
//}

func (q *QrCodeManager) UploadQrCodes(qrCodes []QrCodeUploadModel, fileName string, timeNow uint64) (string, error) {
	q.Log.Infof("test_log|qrCodes len=%v", len(qrCodes))
	//fileName := fmt.Sprintf(fileNameFmt, timeNow)
	qrPaths, err := q.GetQrCodePngFiles(qrCodes)
	if err != nil {
		return "", err
	}
	q.Log.Infof("test_log|make qr success|qrPaths=%s,err=%v", qrPaths, err)
	targetPath := getZipFilePath(fileName)
	body, err := utils.Compress(qrPaths)
	if err != nil {
		return "", err
	}
	q.Log.Infof("test_log|compress success|err=%v", err)
	upload := utils.NewUploadFile(q.Ctx, targetPath, int(timeNow))
	err = upload.Upload(body, fileName)
	if err != nil {
		return "", err
	}
	q.Log.Infof("test_log|upload success|err=%v", err)
	return upload.GetFileReturnUrl(), nil
}

func getZipFilePath(fileName string) string {
	path, _ := os.Getwd()
	return fmt.Sprintf("%s/QRTemp/%s", path, fileName)
}