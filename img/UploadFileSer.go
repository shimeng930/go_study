package img

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	TimeStamp = "timestamp"
	AccessKey = "access_key"
	Signature = "signature"
	NonStr    = "nonstr"
	NonStrLen = 20

	UploadFunction = "/upload?"
	UploadLargeFunction = "/multipart_upload_large?"

	FileThreshold = 10 * 1024 * 1024 // 上传使用大文件上传接口的阈值
	UploadFileItem = 10 * 1024 * 1024 // 文件分片上传的分片大小
)

type UploadFile struct {
	TimeNow 	  int
	FilePath      string
	uploadHost    string
	fileReturnUrl string
	accessKey     string
	secretKey     string
	httpRequest   *HttpReqWrapper
}

type HttpReqWrapper struct {
	Url        string
	Method     string
	ReqBody    *bytes.Buffer
	Headers    http.Header
	Timeout    time.Duration
	RetryTimes int
}

func NewUploadFile(filePath string, timeNow int) *UploadFile {
	// var host = "https://fs.uat.airpay.in.th"
	// var host = "https://fs.uat.airpay.vn"
	// var host = "https://fs.mitra.shopee.co.id"
	// var host = "https://fs.uat.mitra.shopee.co.id"
	// var host = "https://fs.airpay.in.th"
	// var host = "https://fs.airpay.vn"
	// var host = "https://fs.uat.airpay.co.id"
	// var host = "https://fs.airpay.co.id"
	var host = "https://file-server.test.airpay.in.th"

	// var accessKey = "UKBfmTf7i0lcZcqCRxgRZhK6mtIQLyMIGjnk"
	// var secretKey = "CylbAO7pY8F9npU6TbZSC0gSixHy23XUnMv1"

	// test
	var accessKey = "89J1jxm3MWdyxF0uIHu43wcieqt582ypaZm6"
	var secretKey = "iCrmC3tN5IbUmYKNSRnwxVNiJVgLXQ2BLlEK"

	uploadFile := &UploadFile{FilePath: filePath, TimeNow: timeNow}
	uploadFile.uploadHost = host
	uploadFile.accessKey = accessKey
	uploadFile.secretKey = secretKey
	uploadFile.httpRequest = &HttpReqWrapper{
		Method:     "POST",
		//RetryTimes: cfg.GetGlobalConfig().FileServer.RetryTimes,
		//Timeout:    time.Duration(cfg.GetGlobalConfig().FileServer.TimeOut) * time.Millisecond,
	}
	return uploadFile
}

func (u *UploadFile) GetFileReturnUrl() string {
	return u.fileReturnUrl
}

func (u *UploadFile) GetUploadUrl(uploadFunction string) string {
	return u.uploadHost + uploadFunction + u.GetParams()
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

func (u *UploadFile) Upload(body *bytes.Buffer, fileName string) error {
	if body.Len() > FileThreshold {
		u.httpRequest.Url = u.GetUploadUrl(UploadLargeFunction)
		var partNum = body.Len() / UploadFileItem
		return u.uploadLarge(body, fileName, partNum, 0, 0)
	} else {
		var data bytes.Buffer
		mp := multipart.NewWriter(&data)
		w, err := mp.CreateFormFile("file", fileName)
		if err != nil {
			panic(err)
		}
		_, err = w.Write(body.Bytes())
		if err != nil {
			panic(err)
		}
		mp.Close()

		//content, err := doUpload(body, mp.FormDataContentType(), "/upload?")
		var host = "https://file-server.test.airpay.in.th"
		var param = getParams()
		var client = &http.Client{}

		req, err := http.NewRequest("POST", host+"/upload?"+param, &data)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", mp.FormDataContentType())
		res, err := client.Do(req)
		fmt.Println(res.StatusCode)
		if err != nil {
			panic(err)
		}
		content, err := ioutil.ReadAll(res.Body)
		fmt.Println(string(content))

		u.httpRequest.Url = u.GetUploadUrl(UploadFunction)
		return u.uploadSmall(body, fileName)
	}
}

func (u *UploadFile) uploadSmall(body *bytes.Buffer, fileName string) error {
	err := u._getMultipart(body, fileName, 0, 1, 0)
	if err != nil {
		return err
	}

	content, err := u._uploadCall()
	var s UploadResult
	err = json.Unmarshal(content, &s)
	if err != nil {
		return err
	}
	u.fileReturnUrl = s.Data.Url
	return nil
}

func (u *UploadFile) uploadLarge(body *bytes.Buffer, filename string, partNum, partIndex, uploadId int) error {
	if partIndex >= partNum {
		return errors.New("partIndex should less than partNum")
	}

	err := u._getMultipart(body, filename, partIndex, partNum, uploadId)
	if err != nil {
		return err
	}

	content, err := u._uploadCall()
	if err != nil {
		return err
	}
	var s UploadBigResult
	err = json.Unmarshal(content, &s)
	if err != nil {
		return err
	}

	partIndex = partIndex + 1
	if s.Data.IsFinished {
		u.fileReturnUrl = s.Data.Url
	} else {
		uploadId := s.Data.UploadID
		u.uploadLarge(body, filename, partNum, partIndex, int(uploadId))
	}
	return nil
}

func (u *UploadFile) _uploadCall() ([]byte, error) {
	var client = &http.Client{}
	req, err := http.NewRequest("POST", u.httpRequest.Url, u.httpRequest.ReqBody)
	req.Header = u.httpRequest.Headers
	res, err := client.Do(req)
	fmt.Println(res.StatusCode)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
	return content, err
}

func (u *UploadFile) _getMultipart(body *bytes.Buffer, filename string, partIndex, partNum, uploadId int) error {
	var data bytes.Buffer
	mp := multipart.NewWriter(&data)
	w, err := mp.CreateFormFile("file", filename)
	if err != nil {
		return err
	}

	if partNum == 1 {
		_, err = w.Write(body.Bytes())
		if err != nil {
			return err
		}
	} else {
		var lastIndex = (partIndex+1) * UploadFileItem
		if partIndex == (partNum-1) {
			lastIndex = body.Len()
		}
		sliceData := body.Bytes()[partIndex*UploadFileItem:lastIndex]
		_, err = w.Write(sliceData)
		if err != nil {
			return err
		}
		mp.WriteField("part_index", strconv.Itoa(partIndex))
		mp.WriteField("part_number", strconv.Itoa(partNum))
		if partIndex != 0 {
			mp.WriteField("upload_id", strconv.Itoa(uploadId))
		}
	}

	u.httpRequest.ReqBody = &data
	header := http.Header{}
	header.Set("Content-Type", mp.FormDataContentType())
	u.httpRequest.Headers = header
	mp.Close()
	return nil
}

//func (u *UploadFile) Upload(body *bytes.Buffer, fileName string) error {
//	var data bytes.Buffer
//	mp := multipart.NewWriter(&data)
//	w, err := mp.CreateFormFile("file", fileName)
//	if err != nil {
//		return err
//	}
//
//	_, err = w.Write(body.Bytes())
//	if err != nil {
//		return err
//	}
//	mp.Close()
//
//	// var host = "https://fs.uat.airpay.in.th"
//	// var host = "https://fs.uat.airpay.vn"
//	// var host = "https://fs.mitra.shopee.co.id"
//	// var host = "https://fs.uat.mitra.shopee.co.id"
//	// var host = "https://fs.airpay.in.th"
//	// var host = "https://fs.airpay.vn"
//	// var host = "https://fs.uat.airpay.co.id"
//	// var host = "https://fs.airpay.co.id"
//
//	header := http.Header{}
//	header.Set("Content-Type", mp.FormDataContentType())
//	endU := u.GetUploadUrl()
//	println(endU)
//	httpReq := &HttpReqWrapper{
//		Url:        u.GetUploadUrl(),
//		Method:     "POST",
//		ReqBody:    &data,
//		Headers:    header,
//		RetryTimes: cfg.GetGlobalConfig().FileServer.RetryTimes,
//		Timeout:    time.Duration(cfg.GetGlobalConfig().FileServer.TimeOut) * time.Millisecond,
//	}
//	rsp, err := HttpPostRaw(u.Ctx, httpReq)
//	if err != nil {
//		return err
//	}
//	var s UploadResult
//	err = json.Unmarshal(rsp.RspBody, &s)
//	if err != nil {
//		return err
//	}
//	u.fileReturnUrl = s.Data.Url
//	return nil
//}

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
	Url      string
	Msg      string
}

type UploadBigResult struct {
	Code int
	Data UploadBigResultData
}

type UploadBigResultData struct {
	UploadID 	uint64 `json:"upload_id"`
	Url 	 	string
	IsFinished 	bool	`json:"is_finished"`
	PartNumber  int		`json:"part_number"`
	PartIndex   int		`json:"part_index"`
}