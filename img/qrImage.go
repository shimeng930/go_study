package img

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
)

type QrCode struct {
	BgPath  string
	QrID 	uint64
	QrToken string
	QrName 	string
	*Extra
}

type Extra struct {
	OutletName string
	OutletID   uint64
	QrID 	   uint64
}

func NewQrCode(qrID, outletID uint64, qrToken, outletName string) *QrCode {
	qrCode := &QrCode{
		QrID: 	 qrID,
		QrToken: qrToken,
	}
	qrCode.QrName = fmt.Sprintf("cqr_%d.png", qrID)
	qrCode.Extra = &Extra{
		QrID: 		qrID,
		OutletID: 	outletID,
		OutletName: outletName,
	}
	qrCode.BgPath = "static/bg_th.png"
	return qrCode
}

func (q *QrCode) makeQrCode() (barcode.Barcode, error) {
	code, err := qr.Encode(q.QrToken, qr.L, qr.Unicode)
	if err != nil {
		return nil, err
	}

	return code, nil
}

func (q *QrCode) addText(img draw.Image, maxWidth, maxHeight, heightTimes int) error {
	font := getFontPath()
	fontSourceBytes, err := ioutil.ReadFile(font)
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
	fc.SetClip(img.Bounds())
	fc.SetDst(img)
	fc.SetSrc(image.Black)
	fc.SetFontSize(80)

	if q.OutletName != "" {
		pt := freetype.Pt(10*heightTimes, maxHeight-28*heightTimes)
		_, err = fc.DrawString(q.OutletName, pt)
		if err != nil {
			return err
		}
	}

	showVal := ""
	if q.OutletID != 0 {
		showVal = fmt.Sprintf("Shop-%d", q.OutletID)
	} else if q.QrID != 0 {
		showVal = fmt.Sprintf("QR-%d", q.QrID)
	}
	if showVal != "" {
		fc.SetFontSize(64)
		pt := freetype.Pt(maxWidth-75*heightTimes, maxHeight-20*heightTimes)
		_, err = fc.DrawString(showVal, pt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (q *QrCode) MakeQrCodeWithBg() (image.Image, error) {
	// 打开背景图
	bgF, err := file.MustOpen("", q.BgPath)
	if err != nil {
		return nil, err
	}
	defer bgF.Close()
	// 解码背景图获取Image
	bgImage, err := png.Decode(bgF)
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

	print(qrImage.Bounds().Max.X, "-sas-", qrImage.Bounds().Max.Y)

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

func (q *QrCode) FileQrCodeWithBg(name, path string) (string, error) {
	mergedF, err := q.OpenMergedImage(name, path)
	if err != nil {
		return "", err
	}
	defer mergedF.Close()

	img, err := q.MakeQrCodeWithBg()
	if err != nil {
		return "", nil
	}
	png.Encode(mergedF, img)
	return path+name, nil
}

func (q *QrCode) OpenMergedImage(name, path string) (*os.File, error) {
	f, err := file.MustOpen(name, path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func getFontPath() string {
	path, _ := os.Getwd()
	return fmt.Sprintf("%s/static/sf_ui_text_medium.ttf", path)
}