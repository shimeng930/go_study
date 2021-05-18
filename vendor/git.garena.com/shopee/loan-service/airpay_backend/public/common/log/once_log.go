package log

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"git.garena.com/shopee/loan-service/airpay_backend/public/common/metadata"
)

const (
	skipLevel     = 3
	OnceLogStrKey = "OnceLogKey"
)

type Field = zap.Field

type OnceLogKey struct{}

type MonitorOnceLogHandlerType func(level LogLevel, grpcMethod string)

var (
	monitorOnceLogHandler MonitorOnceLogHandlerType
)

func SetMonitorOnceLogHandler(monitorOnceLogHandlerInput MonitorOnceLogHandlerType) {
	monitorOnceLogHandler = monitorOnceLogHandlerInput
}

type OnceLog struct {
	skipLevel       int
	uid             string
	grpcMethod      string
	traceID         string
	customerPrefix  string
	prefixStr       string // "|method:%s|traceId:%s|uid:%s|"  前后有竖线
	blPrintFileLine bool   //是否打印文件名和行号,默认是 false 不打印, 高并发下比较耗性能，谨慎开启.参考: https://cloud.tencent.com/developer/article/1385947
}

// GetTraceID 返回jaeger traceid，ctx是grpc server method的context
func getTraceID(ctx context.Context) string {
	traceId := ""
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		if spanCtx, ok := span.Context().(jaeger.SpanContext); ok {
			traceId = spanCtx.TraceID().String()
		}
	}
	return traceId
}

func GetTraceIDFromContext(ctx context.Context) string { return getTraceID(ctx) }

func (c *OnceLog) SetCustomerPrefix(customerPrefix string) { c.customerPrefix = customerPrefix }

func (c *OnceLog) SetUid(uid string) { c.uid = uid }

func (c *OnceLog) SetGrpcMethod(grpcMethod string) { c.grpcMethod = grpcMethod }

func (c *OnceLog) SetTraceID(traceID string) { c.traceID = traceID }

func (c *OnceLog) SetSkipLevel(skipLevel int) { c.skipLevel = skipLevel }

func (c *OnceLog) SetPrintFileLine(blPrintFileLine bool) { c.blPrintFileLine = blPrintFileLine }

func newOnceLog(method, traceID, uid string) *OnceLog {
	onceLog := &OnceLog{
		traceID:         traceID,
		grpcMethod:      method,
		uid:             uid,
		skipLevel:       skipLevel,
		blPrintFileLine: true,
	}
	return onceLog
}

func NewOnceLogFromContext(ctx context.Context) *OnceLog {
	traceID := getTraceID(ctx)
	uid := metadata.GetUidFromCtx(ctx)
	method, _ := grpc.Method(ctx)
	return newOnceLog(method, traceID, uid)
}

func NewOnceLogFromContextGin(ctx *gin.Context) *OnceLog {
	method := ctx.Request.URL.RequestURI()
	traceID := getTraceID(ctx.Request.Context())
	if traceID == "" {
		traceID = ctx.GetHeader("tracingId")
	}
	uid := ctx.GetHeader("uid")
	return newOnceLog(method, traceID, uid)
}

func ExtractOnceLog(ctx context.Context) *OnceLog {
	onceLog, ok := ctx.Value(OnceLogKey{}).(*OnceLog)
	if !ok {
		return NewOnceLogFromContext(ctx)
	}
	return onceLog
}

func ExtractOnceLogGin(ctx *gin.Context) *OnceLog {
	onceLog, ok := ctx.Get(OnceLogStrKey)
	if !ok {
		return NewOnceLogFromContextGin(ctx)
	}
	return onceLog.(*OnceLog)
}

func (c *OnceLog) Trace(fields ...interface{}) {
	if logger.Core().Enabled(TraceLvl) {
		prefix := c.getFullPrefix(TraceLvl, c.skipLevel)
		logger.Debug(prefix + fmt.Sprint(fields...))
	}
}

func (c *OnceLog) Debug(fields ...interface{}) {
	if logger.Core().Enabled(DebugLvl) {
		prefix := c.getFullPrefix(DebugLvl, c.skipLevel)
		logger.Debug(prefix + fmt.Sprint(fields...))
	}
}

func (c *OnceLog) Info(fields ...interface{}) {
	if logger.Core().Enabled(InfoLvl) {
		prefix := c.getFullPrefix(InfoLvl, c.skipLevel)
		logger.Info(prefix + fmt.Sprint(fields...))
	}
}

func (c *OnceLog) Warn(fields ...interface{}) {
	if logger.Core().Enabled(WarnLvl) {
		prefix := c.getFullPrefix(WarnLvl, c.skipLevel)
		logger.Warn(prefix + fmt.Sprint(fields...))
	}
}

func (c *OnceLog) Error(fields ...interface{}) {
	if logger.Core().Enabled(ErrorLvl) {
		prefix := c.getFullPrefix(ErrorLvl, c.skipLevel)
		logger.Error(prefix + fmt.Sprint(fields...))
	}
}

func (c *OnceLog) Critical(fields ...interface{}) {
	if logger.Core().Enabled(CriticalLvl) {
		prefix := c.getFullPrefix(CriticalLvl, c.skipLevel)
		logger.Fatal(prefix + fmt.Sprint(fields...))
	}
}

// 格式化接口
func (c *OnceLog) Tracef(format string, fields ...interface{}) {
	if logger.Core().Enabled(TraceLvl) {
		prefix := c.getFullPrefix(TraceLvl, c.skipLevel)
		logger.Debug(prefix + fmt.Sprintf(format, fields...))
	}
}

func (c *OnceLog) Debugf(format string, fields ...interface{}) {
	if logger.Core().Enabled(DebugLvl) {
		prefix := c.getFullPrefix(DebugLvl, c.skipLevel)
		logger.Debug(prefix + fmt.Sprintf(format, fields...))
	}
}

func (c *OnceLog) Infof(format string, fields ...interface{}) {
	if logger.Core().Enabled(InfoLvl) {
		prefix := c.getFullPrefix(InfoLvl, c.skipLevel)
		logger.Info(prefix + fmt.Sprintf(format, fields...))
	}
}

func (c *OnceLog) Warnf(format string, fields ...interface{}) {
	if logger.Core().Enabled(WarnLvl) {
		prefix := c.getFullPrefix(WarnLvl, c.skipLevel)
		logger.Warn(prefix + fmt.Sprintf(format, fields...))
	}
}

func (c *OnceLog) Errorf(format string, fields ...interface{}) {
	if logger.Core().Enabled(ErrorLvl) {
		prefix := c.getFullPrefix(ErrorLvl, c.skipLevel)
		logger.Error(prefix + fmt.Sprintf(format, fields...))
	}
}

func (c *OnceLog) Criticalf(format string, fields ...interface{}) {
	if logger.Core().Enabled(CriticalLvl) {
		prefix := c.getFullPrefix(CriticalLvl, c.skipLevel)
		logger.Fatal(prefix + fmt.Sprintf(format, fields...))
	}
}

func (c *OnceLog) getFullPrefix(level LogLevel, skip int, msg ...string) string {
	fileName, line, funcName := "not_print", 0, "not_print"
	if c.blPrintFileLine {
		fileName, line, funcName = findFileInfo(skip)
	}

	buffer := GetBuffer()
	buffer.AppendString("|grpcMethod:")
	buffer.AppendString(c.grpcMethod)
	buffer.AppendString("|traceId:")
	buffer.AppendString(c.traceID)
	buffer.AppendString("|uid:")
	buffer.AppendString(c.uid)
	buffer.AppendString("|file:")
	buffer.AppendString(fileName)
	buffer.AppendByte('/')
	buffer.AppendInt(int64(line))
	buffer.AppendString("|func:")
	buffer.AppendString(funcName)
	buffer.AppendByte('|')
	if c.customerPrefix != "" {
		buffer.AppendString(c.customerPrefix)
		buffer.AppendByte('|')
	}
	if len(msg) > 0 {
		buffer.AppendString(msg[0])
	}
	str := buffer.String()
	buffer.Free()

	if monitorOnceLogHandler != nil {
		monitorOnceLogHandler(level, c.grpcMethod)
	}
	return str
}

func (c *OnceLog) Tracew(fields ...interface{}) {
	if logger.Core().Enabled(TraceLvl) {
		prefix := c.getFullPrefix(TraceLvl, c.skipLevel)
		logger.Debug(prefix + formatw(fields))
	}
}

func (c *OnceLog) Debugw(fields ...interface{}) {
	if logger.Core().Enabled(DebugLvl) {
		prefix := c.getFullPrefix(DebugLvl, c.skipLevel)
		logger.Debug(prefix + formatw(fields))
	}
}

func (c *OnceLog) Infow(fields ...interface{}) {
	if logger.Core().Enabled(InfoLvl) {
		prefix := c.getFullPrefix(InfoLvl, c.skipLevel)
		logger.Info(prefix + formatw(fields))
	}
}

func (c *OnceLog) Warnw(fields ...interface{}) {
	if logger.Core().Enabled(WarnLvl) {
		prefix := c.getFullPrefix(WarnLvl, c.skipLevel)
		logger.Warn(prefix + formatw(fields))
	}
}

func (c *OnceLog) Errorw(fields ...interface{}) {
	if logger.Core().Enabled(ErrorLvl) {
		prefix := c.getFullPrefix(ErrorLvl, c.skipLevel)
		logger.Error(prefix + formatw(fields))
	}
}

func (c *OnceLog) Criticalw(fields ...interface{}) {
	if logger.Core().Enabled(CriticalLvl) {
		prefix := c.getFullPrefix(CriticalLvl, c.skipLevel)
		logger.Fatal(prefix + formatw(fields))
	}
}

func (c *OnceLog) Flush() {
	if err := logger.Sync(); err != nil {
		fmt.Println("Sync error: ", err)
	}
}

func Flush() {
	FlushFlowLog()
	logger.Sync()
}

/****************推荐使用下列函数*******************************/
func (c *OnceLog) DebugF(msg string, fields ...Field) {
	if logger.Core().Enabled(DebugLvl) {
		prefix := c.getFullPrefix(DebugLvl, c.skipLevel, msg)
		logger.Debug(prefix, fields...)
	}
}

func (c *OnceLog) InfoF(msg string, fields ...Field) {
	if logger.Core().Enabled(InfoLvl) {
		prefix := c.getFullPrefix(InfoLvl, c.skipLevel, msg)
		logger.Info(prefix, fields...)
	}
}

func (c *OnceLog) WarnF(msg string, fields ...Field) {
	if logger.Core().Enabled(WarnLvl) {
		prefix := c.getFullPrefix(WarnLvl, c.skipLevel, msg)
		logger.Warn(prefix, fields...)
	}
}

func (c *OnceLog) ErrorF(msg string, fields ...Field) {
	if logger.Core().Enabled(ErrorLvl) {
		prefix := c.getFullPrefix(ErrorLvl, c.skipLevel, msg)
		logger.Error(prefix, fields...)
	}
}

func (c *OnceLog) DPanicF(msg string, fields ...Field) {
	if logger.Core().Enabled(DPanicLvl) {
		prefix := c.getFullPrefix(DPanicLvl, c.skipLevel, msg)
		logger.DPanic(prefix, fields...)
	}
}

func (c *OnceLog) PanicF(msg string, fields ...Field) {
	if logger.Core().Enabled(PanicLvl) {
		prefix := c.getFullPrefix(PanicLvl, c.skipLevel, msg)
		logger.Panic(prefix, fields...)
	}
}

func (c *OnceLog) FatalF(msg string, fields ...Field) {
	if logger.Core().Enabled(FatalLvl) {
		prefix := c.getFullPrefix(FatalLvl, c.skipLevel, msg)
		logger.Fatal(prefix, fields...)
	}
}

/***********************************************************/
