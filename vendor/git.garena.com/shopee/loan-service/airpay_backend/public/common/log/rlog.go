package log

import (
	"fmt"

	"github.com/go-xorm/core"
	"go.uber.org/zap"
)

type rLogger struct {
	logger  *zap.Logger
	showSQL bool
}

var rLog *rLogger

func (log *rLogger) Debug(v ...interface{}) {
	log.logger.Debug(fmt.Sprint(v...))
}

func (log *rLogger) Debugf(format string, v ...interface{}) {
	log.logger.Debug(fmt.Sprintf(format, v...))
}

func (log *rLogger) Info(v ...interface{}) {
	log.logger.Info(fmt.Sprint(v...))
}

func (log *rLogger) Infof(format string, v ...interface{}) {
	log.logger.Info(fmt.Sprintf(format, v...))
}

func (log *rLogger) Warn(v ...interface{}) {
	log.logger.Warn(fmt.Sprint(v...))
}

func (log *rLogger) Warnf(format string, v ...interface{}) {
	log.logger.Warn(fmt.Sprintf(format, v...))
}

func (log *rLogger) Error(v ...interface{}) {
	log.logger.Error(fmt.Sprint(v...))
}

func (log *rLogger) Errorf(format string, v ...interface{}) {
	log.logger.Error(fmt.Sprintf(format, v...))
}

func (log *rLogger) Level() core.LogLevel {
	return 0
}

func (log *rLogger) SetLevel(core.LogLevel) {}

func (log *rLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		log.showSQL = true
	} else {
		log.showSQL = show[0]
	}
}

func (log *rLogger) IsShowSQL() bool {
	return log.showSQL
}

func GetRLog() *rLogger {
	return rLog
}
