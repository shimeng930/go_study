package utils

import (
	"bytes"
	"context"
	"fmt"
	"git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_admin/common/util"
	"io/ioutil"
	"net/http"
	"time"
)

var httpClient *http.Client

func InitHttpClient() *http.Client {
	return &http.Client{}
}

type HttpReqWrapper struct {
	Url    		string
	Method    	string
	ReqBody     *bytes.Buffer
	Headers     http.Header
	Timeout     time.Duration
	RetryTimes  int
}

func (h HttpReqWrapper) GoString() string {
	return h.String()
}

func (h HttpReqWrapper) String() string {
	return fmt.Sprintf("host=%s, headers=%s, reqBody=%s, timeout=%s, retry_times=%d", h.Url, h.Headers,
		string(h.ReqBody.Bytes()), h.Timeout, h.RetryTimes)
}

func (h HttpReqWrapper) StringWithoutBody() string {
	return fmt.Sprintf("host=%s, headers=%s, timeout=%s, retry_times=%d", h.Url, h.Headers, h.Timeout, h.RetryTimes)
}

type HttpRspWrapper struct {
	RspBody    []byte
	Headers    http.Header
	HttpStatus int
}

func (h HttpRspWrapper) GoString() string {
	return h.String()
}

func (h HttpRspWrapper) String() string {
	return fmt.Sprintf("status_code:%d, headers=%s, rsp_body=%s", h.HttpStatus, h.Headers, string(h.RspBody))
}

func HttpPostRaw(ctx context.Context, req *HttpReqWrapper) (rsp *HttpRspWrapper, err error) {
	ctxLog := util.CtxLog(ctx)
	ctxLog.Debugf("[http_post_raw]begin|req=%+v", req.StringWithoutBody())
	for i := 1; i <= req.RetryTimes; i++ {
		start := time.Now()
		rsp, err = doPost(ctx, req)
		cost := time.Now().Sub(start) / time.Millisecond
		if err == nil {
			ctxLog.Warn(fmt.Sprintf("[%d]http_post_success|cost=%d, url=%s", i, cost, req.Url))
			break
		}
		ctxLog.Warn(fmt.Sprintf("[%d]http_post_error|cost=%d, url=%s, error=%v", i, cost, req.Url, err))
	}
	ctxLog.Debugf("[http_post_raw]end|rsp=%+v", rsp)
	if err != nil {
		return
	}
	return rsp, nil
}

func doPost(ctx context.Context, req *HttpReqWrapper) (*HttpRspWrapper, error) {
	request, err := http.NewRequest(req.Method, req.Url, req.ReqBody)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, req.Timeout)
	defer cancel()
	request = request.WithContext(ctx)
	request.Header = req.Headers
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &HttpRspWrapper{
		RspBody:    body,
		Headers:    response.Header,
		HttpStatus: response.StatusCode,
	}, nil
}