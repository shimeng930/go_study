package log

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"git.garena.com/shopee/loan-service/airpay_backend/public/common/log/lumberjack"
	"git.garena.com/shopee/loan-service/airpay_backend/public/common/utils"
)

type LogLevel = zapcore.Level

// zap 中没有trace 和 critial, 分别用debug和fatal替代
const (
	TraceLvl    = zapcore.DebugLevel
	DebugLvl    = zapcore.DebugLevel
	InfoLvl     = zapcore.InfoLevel
	WarnLvl     = zapcore.WarnLevel
	ErrorLvl    = zapcore.ErrorLevel
	CriticalLvl = zapcore.FatalLevel
	PanicLvl    = zapcore.PanicLevel
	DPanicLvl   = zapcore.DPanicLevel
	FatalLvl    = zapcore.FatalLevel
	Off         = zapcore.Level(-2)
)

var (
	logger *zap.Logger
)

const (
	customTimeLayout = "2006-01-02 15:04:05.999999+07:00"
)

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(customTimeLayout)
	cfg.Sampling = nil
	logger, _ = cfg.Build()
}

func InitLog(level LogLevel) (string, error) {
	// zap 和 seelog 对level的定义不一样，这里转换下和现在配置文件保持一致
	level = LogLevel(int8(level) - 2)
	if level.Enabled(DebugLvl) {
		level = DebugLvl
	}
	randDir := ""
	podName, ok := os.LookupEnv("POD_NAME")
	if ok {
		randDir = podName
	}
	if randDir == "" && utils.IsInsideDocker() {
		rand.Seed(time.Now().UnixNano() ^ int64(os.Getpid()))
		randDir = fmt.Sprintf("%s-%d", time.Now().Format("20060102150405"), rand.Int()%1000)
	}
	filename := filepath.Join(".", "log", randDir, "rlog.log")

	logger = zap.New(buildCore(filename, level, true))
	rLog = &rLogger{logger, false}

	InitFlowLog(randDir)

	return randDir, nil
}

func buildCore(fileName string, level zapcore.Level, isRlog bool) zapcore.Core {
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		LocalTime: true,
		Compress:  false,
	})
	if _useIODiscard {
		syncWriter = zapcore.AddSync(ioutil.Discard)
	}

	encodeConfig := zapcore.EncoderConfig{
		MessageKey:       "msg",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout(customTimeLayout),
		ConsoleSeparator: " ",
	}
	if isRlog {
		encodeConfig.LevelKey = "level"
		encodeConfig.TimeKey = "ts"
	}

	encoder := zapcore.NewConsoleEncoder(encodeConfig)
	if _useJSONEncoder {
		encoder = zapcore.NewJSONEncoder(encodeConfig)
	}

	core := zapcore.NewCore(encoder, syncWriter, level)
	return core
}

func formatLog(level LogLevel, skip int) string {
	fileName, line, funcName := findFileInfo(skip)
	//return fmt.Sprintf("[%s] [%v] [%v/%v]", level.String(), funcName, fileName, line)
	// 这里level已经由zap框架打印了, 不需要再输出了
	return fmt.Sprintf("[%v] [%v/%v]", funcName, fileName, line)
}

func findFileInfo(skip int) (string, int, string) {
	fileName, line, funcName := "???", 0, "???"
	pc, fileName, line, ok := runtime.Caller(skip)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
		if idx := strings.LastIndexByte(funcName, '.'); idx > 0 {
			funcName = funcName[idx+1:]
		}

		if idx := strings.LastIndexByte(fileName, '/'); idx > 0 {
			fileName = fileName[idx+1:]
			// 下面输出为runtime/asm_amd64.s
			// if idx = strings.LastIndexByte(fileName[:idx], '/'); idx > 0 {
			// 	fileName = fileName[idx+1:]
			// }
		}
	}
	return fileName, line, funcName
}
