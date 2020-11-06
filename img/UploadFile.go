package img

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Upload(body *bytes.Buffer, filename string) {
	if body.Len() > uploadItemLen {
		UploadLarge(body, filename)
	} else {
		UploadSmall(body, filename)
	}
}

func getParams() string {
	var timestamps = time.Now().Unix()
	var nonstr = CryptUtil.RandomStringUseDefaultAllowedChars(20)

	var accessKey = "89J1jxm3MWdyxF0uIHu43wcieqt582ypaZm6"
	var secretKey = "iCrmC3tN5IbUmYKNSRnwxVNiJVgLXQ2BLlEK"

	var param = make(url.Values)
	param.Add("timestamp", strconv.Itoa(int(timestamps)))
	param.Add("nonstr", nonstr)
	param.Add("access_key", accessKey)
	queryStr := sortStr(param.Encode())
	signature := CryptUtil.HmacSha256(queryStr, secretKey)
	param.Add("signature", signature)
	return param.Encode()
}

const uploadItemLen = 10 * 1024 * 1024

func doUpload(data *bytes.Buffer, contentType, requestUrl string) ([]byte, error) {
	var host = "https://file-server.test.airpay.in.th"
	var param = getParams()
	var client = &http.Client{}

	//req, err := http.NewRequest("POST", host+"/upload?skip_signature=1", &body)
	req, err := http.NewRequest("POST", host+requestUrl+param, data)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	res, err := client.Do(req)
	fmt.Println(res.StatusCode)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	return content, err
}

func UploadLarge(body *bytes.Buffer, filename string) {
	var partNum = body.Len() / uploadItemLen
	//if body.Len() % uploadItemLen != 0 {
	//	partNum = partNum + 1
	//}
	//var partIndex = 1
	uploadLargeLogic(body, filename, partNum, 0, 0)
}

func uploadLargeLogic(body *bytes.Buffer, filename string, partNum, partIndex, uploadId int) error {
	//var partIndex = 1
	//var lastIndex = partIndex * uploadItemLen
	//if partIndex == partNum {
	//	lastIndex = body.Len()
	//}
	//sliceData := body.Bytes()[(partIndex-1)*uploadItemLen:lastIndex]
	
	if partIndex == partNum {
		return errors.New("params error")
	}

	mp, data, _ := _getMP(body, filename, partIndex, partNum, uploadId)
	mp.Close()

	content, err := doUpload(data, mp.FormDataContentType(), "/multipart_upload_large?")
	if err != nil {
		return err
	}
	var s UploadBigResult
	fmt.Println(string(content))
	err = json.Unmarshal(content, &s)
	if err != nil {
		return err
	}

	partIndex = partIndex + 1
	if s.Data.IsFinished {
		fmt.Println(s.Data.Url)
	} else {
		uploadId := s.Data.UploadID
		uploadLargeLogic(body, filename, partNum, partIndex, int(uploadId))
	}
	return nil

}

func _getMP(body *bytes.Buffer, filename string, partIndex, partNum, uploadId int) (*multipart.Writer, *bytes.Buffer, error) {
	var data bytes.Buffer
	mp := multipart.NewWriter(&data)
	w, err := mp.CreateFormFile("file", filename)
	if err != nil {
		return nil, nil, err
	}

	var lastIndex = (partIndex+1) * uploadItemLen
	if partIndex == (partNum-1) {
		lastIndex = body.Len()
	}
	sliceData := body.Bytes()[partIndex*uploadItemLen:lastIndex]
	_, err = w.Write(sliceData)
	if err != nil {
		//return err
	}
	mp.WriteField("part_index", strconv.Itoa(partIndex))
	mp.WriteField("part_number", strconv.Itoa(partNum))
	if partIndex != 0 {
		mp.WriteField("upload_id", strconv.Itoa(uploadId))
	}
	return mp, &data, err
}

func UploadSmall(body *bytes.Buffer, filename string) {
	var data bytes.Buffer
	mp := multipart.NewWriter(&data)
	w, err := mp.CreateFormFile("file", filename)
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
		return
	}
	req.Header.Set("Content-Type", mp.FormDataContentType())
	res, err := client.Do(req)
	fmt.Println(res.StatusCode)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

type HttpRspWrapper struct {
	RspBody    []byte
	Headers    http.Header
	HttpStatus int
}

func (h HttpRspWrapper) String() string {
	return fmt.Sprintf("status_code:%d, headers=%s, rsp_body=%s", h.HttpStatus, h.Headers, string(h.RspBody))
}