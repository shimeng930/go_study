package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_admin/global/cfg"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	TimeStamp 	= "timestamp"
	AccessKey	= "access_key"
	Signature	= "signature"
	NonStr		= "nonstr"
	NonStrLen   = 20

	UploadFunction = "/upload?"
)

type UploadFile struct {
	Ctx 			context.Context
	FilePath 		string
	TimeNow			int
	uploadHost 		string
	fileReturnUrl 	string
	accessKey		string
	secretKey		string
}

func NewUploadFile(ctx context.Context, filePath string, timeNow int) *UploadFile {
	uploadFile := &UploadFile{Ctx: ctx, FilePath: filePath, TimeNow: timeNow}
	uploadFile.uploadHost = cfg.GetGlobalConfig().FileServer.Host
	uploadFile.accessKey = cfg.GetGlobalConfig().FileServer.AccessKey
	uploadFile.secretKey = cfg.GetGlobalConfig().FileServer.SecretKey
	return uploadFile
}

func (u *UploadFile) GetFileReturnUrl() string {
	return u.fileReturnUrl
}

func (u *UploadFile) GetUploadUrl() string {
	return u.uploadHost + UploadFunction + u.GetParams()
}

func (u *UploadFile) GetParams() string {
	var param = make(url.Values)
	param.Add(TimeStamp, strconv.Itoa(u.TimeNow))
	param.Add(NonStr, CryptUtil.RandomStringUseDefaultAllowedChars(NonStrLen))
	param.Add(AccessKey, u.accessKey)
	queryStr := sortStr(param.Encode())
	signature := CryptUtil.HmacSha256(queryStr, u.secretKey)
	param.Add(Signature, signature)
	return param.Encode()
}

func (u *UploadFile) GetUploadData(body *bytes.Buffer, fileName string) (*bytes.Buffer, error) {
	var data bytes.Buffer
	mp := multipart.NewWriter(&data)
	w, err := mp.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	_, err = w.Write(body.Bytes())
	if err != nil {
		return nil, err
	}
	mp.Close()
	return &data, nil
}

func (u *UploadFile) GetUploadDataByFile() (*bytes.Buffer, error) {
	var body bytes.Buffer
	mp := multipart.NewWriter(&body)
	w, err := mp.CreateFormFile("file", u.FilePath)
	if err != nil {
		return nil, err
	}

	fileContent, err := ioutil.ReadFile(u.FilePath)
	if err != nil {
		return nil, err
	}
	_, err = w.Write(fileContent)
	if err != nil {
		return nil, err
	}
	mp.Close()
	return &body, nil
}

func (u *UploadFile) Upload(body *bytes.Buffer, fileName string) error {
	var data bytes.Buffer
	mp := multipart.NewWriter(&data)
	w, err := mp.CreateFormFile("file", fileName)
	if err != nil {
		return err
	}

	_, err = w.Write(body.Bytes())
	if err != nil {
		return err
	}
	mp.Close()

	// var host = "https://fs.uat.airpay.in.th"
	// var host = "https://fs.uat.airpay.vn"
	// var host = "https://fs.mitra.shopee.co.id"
	// var host = "https://fs.uat.mitra.shopee.co.id"
	// var host = "https://fs.airpay.in.th"
	// var host = "https://fs.airpay.vn"
	// var host = "https://fs.uat.airpay.co.id"
	// var host = "https://fs.airpay.co.id"

	header := http.Header{}
	header.Set("Content-Type", mp.FormDataContentType())
	endU := u.GetUploadUrl()
	println(endU)
	httpReq := &HttpReqWrapper{
		Url:    	u.GetUploadUrl(),
		Method:    	"POST",
		ReqBody:    &data,
		Headers:    header,
		RetryTimes: cfg.GetGlobalConfig().FileServer.RetryTimes,
		Timeout:    time.Duration(cfg.GetGlobalConfig().FileServer.TimeOut) * time.Millisecond,
	}
	rsp, err := HttpPostRaw(u.Ctx, httpReq)
	if err != nil {
		return err
	}
	var s UploadResult
	err = json.Unmarshal(rsp.RspBody, &s)
	if err != nil {
		return err
	}
	u.fileReturnUrl = s.Data.Url
	return nil
}

func sortStr(s string) string {
	var keys []string
	paramMap := make(map[string]string)

	params := strings.Split(s, "&")
	for _, item := range params {
		param := strings.Split(item, "=")
		if param == nil || len(param) != 2 {
			return ""
		}
		keys = append(keys, param[0])
		paramMap[param[0]] = param[1]
	}
	sort.Strings(keys)
	var sortParams string
	for index, item := range keys {
		sortParams = fmt.Sprintf("%s%s=%s", sortParams, item, paramMap[item])
		if index != len(keys)-1 {
			sortParams = fmt.Sprintf("%s&", sortParams)
		}
	}
	return sortParams
}

type UploadResult struct {
	Code int
	Data UploadResultData
}

type UploadResultData struct {
	UploadID uint64 `json:"upload_id"`
	Url 	 string
	Msg 	 string
}