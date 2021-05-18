package util

import (
	"context"
	"path/filepath"
	"runtime"
	"strings"

	_log "git.garena.com/shopee/loan-service/airpay_backend/public/common/log"
)

type Log interface {
	Trace(fields ...interface{})
	Debug(fields ...interface{})
	Info(fields ...interface{})
	Warn(fields ...interface{})
	Error(fields ...interface{})
	Critical(fields ...interface{})

	Tracef(format string, fields ...interface{})
	Debugf(format string, fields ...interface{})
	Infof(format string, fields ...interface{})
	Warnf(format string, fields ...interface{})
	Errorf(format string, fields ...interface{})
	Criticalf(format string, fields ...interface{})
}

func CtxLog(ctx context.Context) Log {
	return _log.NewOnceLogFromContext(ctx)
}

func FindFuncName(skip int) string {
	funcName := "???"
	pc, _, _, ok := runtime.Caller(skip)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
		funcName = filepath.Ext(funcName)
		funcName = strings.TrimPrefix(funcName, ".")
	}
	return funcName
}

// to use like defer FuncTimeExpenseEnter
func FuncTimeExpenseEnter(log Log, strTag string, nanoTimeNow uint64) {
	if log != nil {
		log.Infof("Time expense: [%s] done, cost:%dms", strTag, (TimeNowNano()-nanoTimeNow)/1000000)
	}
}
