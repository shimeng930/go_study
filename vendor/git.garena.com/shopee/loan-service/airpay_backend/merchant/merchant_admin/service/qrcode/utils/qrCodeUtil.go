package utils

import (
	"bytes"
	"fmt"
	"git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_admin/service/qrcode/static"
	"git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_server/core/product/cscan/codec"
	"git.garena.com/shopee/loan-service/airpay_backend/public/common/utils"
	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"image/png"
	"os"
)

const (
	thBgPath	= "service/qrcode/static/bg_th.png"
	vnBgPath	= "service/qrcode/static/bg_vn.png"
	defaultFont	= "service/qrcode/static/sf_ui_text_medium.ttf"
	thFont		= "service/qrcode/static/thai.ttf"
	vnFont		= "service/qrcode/static/vi.ttf"
	CodeCustomRandomLength = 20
)

type QrCodePng struct {
	BgPath  string
	QrToken string
	QrName 	string
	*Extra
}

type Extra struct {
	OutletName string
	OutletID   uint64
	QrID 	   uint64
}

func NewQrCodePng(qrID, outletID uint64, qrToken, outletName string) *QrCodePng {
	qrCode := &QrCodePng{
		QrToken: qrToken,
	}
	qrCode.QrName = fmt.Sprintf("cqr_%d.png", qrID)
	qrCode.Extra = &Extra{
		QrID: 		qrID,
		OutletID: 	outletID,
		OutletName: outletName,
	}

	if utils.GetRegion() == "vn" {
		qrCode.BgPath = vnBgPath
	} else {
		qrCode.BgPath = thBgPath
	}
	return qrCode
}

func (q *QrCodePng) SetQrID(qrId uint64) {
	q.QrID = qrId
}

func (q *QrCodePng) SetOutletID(outletID uint64) {
	q.OutletID = outletID
}

func (q *QrCodePng) SetOutletName(outletName string) {
	q.OutletName = outletName
}

func (q *QrCodePng) SetQrToken(qrToken string) {
	q.QrToken = qrToken
}

func (q *QrCodePng) GetQrCodeName() string {
	if q.OutletID > 0 {
		return fmt.Sprintf("cqr_store_%d.png", q.OutletID)
	} else if q.QrID > 0 {
		return q.QrName
	}
	return "cqr.png"
}

func (q *QrCodePng) makeQrCode() (barcode.Barcode, error) {
	code, err := qr.Encode(q.QrToken, qr.L, qr.Unicode)
	if err != nil {
		return nil, err
	}

	return code, nil
}

func (q *QrCodePng) adjustShopName() (string, string) {
	if len(q.OutletName) <= 34 {
		return q.OutletName, ""
	} else if len(q.OutletName) <= 64 {
		return q.OutletName[0:32], q.OutletName[32:]
	} else {
		return fmt.Sprintf("%s...", q.OutletName[0:62]), ""
	}
}

func (q *QrCodePng) addText(img draw.Image, maxWidth, maxHeight, heightTimes int) error {
	font := getFontPath()
	fontSourceBytes, err := static.Asset(font)
	if err != nil {
		return err
	}

	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)
	if err != nil {
		return err
	}

	fc := freetype.NewContext()
	fc.SetDPI(72)
	fc.SetFont(trueTypeFont)
	fc.SetSrc(image.Black)
	fc.SetFontSize(80)

	if q.OutletName != "" {
		_showName, _secondName := q.adjustShopName()
		tempImg := image.NewRGBA(image.Rect(0, 0, maxWidth, maxHeight))
		nameLength := q.calOutletNameLen(_showName, tempImg, fc)
		fc.SetClip(img.Bounds())
		fc.SetDst(img)
		pt := freetype.Pt((maxWidth-nameLength)/2, maxHeight-28*heightTimes)
		_, err := fc.DrawString(_showName, pt)
		if err != nil {
			return err
		}

		if _secondName != "" {
			nameLength := q.calOutletNameLen(_secondName, tempImg, fc)
			fc.SetClip(img.Bounds())
			fc.SetDst(img)
			pt := freetype.Pt((maxWidth-nameLength)/2, maxHeight-18*heightTimes)
			_, err := fc.DrawString(_secondName, pt)
			if err != nil {
				return err
			}
		}
	}

	showVal := ""
	if q.OutletID != 0 {
		showVal = fmt.Sprintf("Shop-%d", q.OutletID)
	} else if q.QrID != 0 {
		showVal = fmt.Sprintf("QR-%d", q.QrID)
	}
	if showVal != "" {
		fc.SetClip(img.Bounds())
		fc.SetDst(img)
		fc.SetFontSize(64)
		pt := freetype.Pt(maxWidth-75*heightTimes, maxHeight-20*heightTimes)
		_, err = fc.DrawString(showVal, pt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (q *QrCodePng) calOutletNameLen(showName string, img draw.Image, fc *freetype.Context) int {
	fc.SetClip(img.Bounds())
	fc.SetDst(img)
	pt := freetype.Pt(100, 100)
	eP, err := fc.DrawString(showName, pt)
	if err != nil {
		return 0
	}
	return int(eP.X - pt.X)  / 64
}

func (q *QrCodePng) MakeQrCodeWithBg() (image.Image, error) {
	// 获取背景图内容
	byteData, err := static.Asset(q.BgPath)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(byteData)
	// 解码背景图获取Image
	bgImage, err := png.Decode(reader)
	if err != nil {
		return nil, err
	}

	qrImage, err := q.makeQrCode()
	if err != nil {
		return nil, err
	}

	max := bgImage.Bounds().Max
	maxWidth := max.X
	maxHeight := max.Y
	widthTimes := maxWidth/320
	heightTimes := maxHeight/460
	embedLeft := 74*widthTimes
	embedUp := 74*heightTimes

	qrImage, err = barcode.Scale(qrImage, 170*widthTimes, 195*heightTimes)
	if err != nil {
		return nil, err
	}

	jpg := image.NewRGBA(image.Rect(0, 0, maxWidth, maxHeight))
	draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
	draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(embedLeft, embedUp)), draw.Over)
	if q.OutletName != "" || q.OutletID != 0 || q.QrID != 0 {
		if err = q.addText(jpg, maxWidth,maxHeight, heightTimes); err != nil {
			return nil, err
		}
	}
	return jpg, nil
}

func (q *QrCodePng) FileQrCodeWithBg(fileName string) (string, error) {
	mergedF, err := q.OpenMergedImage(fileName, "QRTemp/")
	if err != nil {
		return "", err
	}
	defer mergedF.Close()

	img, err := q.MakeQrCodeWithBg()
	if err != nil {
		return "", err
	}
	png.Encode(mergedF, img)
	return mergedF.Name(), nil
}

func (q *QrCodePng) OpenMergedImage(fileName, filePath string) (*os.File, error) {
	f, err := file.MustOpen(fileName, filePath)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func getFontPath() string {
	if utils.GetRegion() == "vn" {
		return vnFont
	} else if utils.GetRegion() == "th" {
		return thFont
	} else {
		return defaultFont
	}
}

//**** QrCode model

type StoreQrCode struct {
	BusinessType string
	OutletID 	 uint64
}

func NewStoreQrCode(outletID uint64) *StoreQrCode {
	return &StoreQrCode{BusinessType: "02", OutletID: outletID}
}

func (s *StoreQrCode) GenerateBusinessData() string {
	// static_prefix_data + business_content + random_content
	// static_prefix_data: 0016A900000000000000 | 0103190 | 0203CsB {protocol defined}
	// self define: 03|length|data
	// business_content: version|purpose|business_type|level|type|category
	// version=01
	// purpose=00(payment code) ...
	// business_type=01 {merchant collection} 02 {store collection}
	// level=01 {store level, reserved filed}
	// type=01 {reserved filed}
	// category=01 {reserved filed}
	// id = 00000000:  8 characters,merchant_id or store_id base on business_type
	outletVal := fmt.Sprintf("%08d", s.OutletID)
	version := "00"
	purpose := "00"
	level := "01"
	_type := "01"
	category := "01"

	businessContent := version + purpose + s.BusinessType + level
	businessContent = businessContent + _type + category + outletVal

	return businessContent
}

func (s *StoreQrCode) GetQrToken() string {
	businessContent := s.GenerateBusinessData()
	cc := s.NewCodec(utils.GetRegion(), "", businessContent)
	if cc == nil {
		return ""
	}

	body := cc.GetQrCodeBody()
	c := &codec.BaseQrCodec{}
	return body + c.GenerateCRC16Sum(body)
}

func (s *StoreQrCode) NewCodec(region string, randomContent string, businessData string) codec.QrCodec {
	if randomContent == "" {
		randomContent = CryptUtil.RandomStringUseDefaultAllowedChars(CodeCustomRandomLength)
	}

	baseQrCodec := codec.BaseQrCodec{
		RandomContent: randomContent,
		BusinessData: businessData,
	}

	if region == utils.VN {
		return &codec.QrCodecForVN{
			BaseQrCodec: baseQrCodec,
		}
	} else if region == utils.TH {
		return &codec.QrCodecForTH{
			BaseQrCodec: baseQrCodec,
		}
	}
	return &baseQrCodec
}