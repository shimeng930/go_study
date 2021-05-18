package log

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"go.uber.org/zap"
)

const (
	defaultRespStr = "<rsp not print>"
	defaultReqStr  = "<request not print>"
)

type MonitorFLogHandlerType func(grpcMethod string)

var (
	flowLog            *zap.Logger
	defaultMaxLogLen   = 200
	monitorFLogHandler MonitorFLogHandlerType
)

func SetFlogLen(l int) {
	if l <= 0 {
		l = defaultMaxLogLen
	}
	defaultMaxLogLen = l
}

func SetMonitorFLogHandler(monitorFLogHandlerTypeInput MonitorFLogHandlerType) {
	monitorFLogHandler = monitorFLogHandlerTypeInput
}

func InitFlowLog(randDir string) {
	filename := filepath.Join(".", "log", randDir, "flog.log")
	flowLog = zap.New(buildCore(filename, DebugLvl, false))
}

func FlowLogInfo(traceId string, retCode int, timeCost time.Duration, uid string, fullMethod string, custStr string, req interface{}, resp interface{}) {
	flowLog.Info(build(traceId, retCode, timeCost, uid, fullMethod, custStr, req, resp, true, false, "", false))
}

func FlowLogInfoWithFlag(traceId string, retCode int, timeCost time.Duration, uid string, fullMethod string, custStr string, req interface{}, resp interface{}, printResp bool) {
	if monitorFLogHandler != nil { //上报flog的调用次数
		monitorFLogHandler(fullMethod)
	}
	flowLog.Info(build(traceId, retCode, timeCost, uid, fullMethod, custStr, req, resp, true, printResp, "", false))
}

func FlowLogInfoWithConfig(traceId string, retCode int, timeCost time.Duration, uid string, fullMethod string, custStr string, req interface{}, resp interface{}, printResp bool, printRequest bool, callFrom string) {
	if monitorFLogHandler != nil {
		monitorFLogHandler(fullMethod)
	}
	flowLog.Info(build(traceId, retCode, timeCost, uid, fullMethod, custStr, req, resp, printResp, printRequest, callFrom, true))
}

// 这里的格式真的很诡异。。。
func build(traceId string, retCode int, timeCost time.Duration, uid string, fullMethod string, custStr string, req interface{}, resp interface{}, printResp bool, printRequest bool, callFrom string, withConfig bool) string {
	respStr := defaultRespStr
	if printResp {
		respStr = fmt.Sprint(resp)
		strLen := len(respStr)
		if strLen > defaultMaxLogLen {
			respStr = respStr[0:defaultMaxLogLen-1] + "...(leave " + strconv.Itoa(strLen-defaultMaxLogLen) + ")"
		}
	}

	reqStr := defaultReqStr
	if printRequest {
		reqStr = fmt.Sprint(req)
	}

	buffer := GetBuffer()
	buffer.AppendString(time.Now().Format(customTimeLayout))
	buffer.AppendByte('|')
	buffer.AppendString(traceId)
	buffer.AppendByte('|')
	buffer.AppendInt(int64(retCode))
	buffer.AppendByte('|')
	buffer.AppendInt(int64(timeCost))
	buffer.AppendByte('|')
	buffer.AppendString(uid)
	buffer.AppendByte('|')
	buffer.AppendString(fullMethod)
	buffer.AppendByte('|')
	if !withConfig {
		buffer.AppendString(custStr)
	} else {
		buffer.AppendString(callFrom)
	}

	buffer.AppendByte('|')
	buffer.AppendString(reqStr)
	buffer.AppendByte('|')
	buffer.AppendString(respStr)
	if withConfig {
		buffer.AppendByte('|')
		buffer.AppendString(custStr)
	}
	buffer.AppendByte('|')

	str := buffer.String()
	buffer.Free()
	return str
}

func FlushFlowLog() {
	if flowLog != nil {
		flowLog.Sync()
	}
}
