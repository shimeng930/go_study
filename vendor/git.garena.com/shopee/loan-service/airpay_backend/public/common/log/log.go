package log

import "fmt"

func Trace(fields ...interface{}) {
	head := formatLog(TraceLvl, skipLevel)
	logger.Debug(head + fmt.Sprint(fields...))
}

func Debug(fields ...interface{}) {
	head := formatLog(DebugLvl, skipLevel)
	logger.Debug(head + fmt.Sprint(fields...))
}

func Info(fields ...interface{}) {
	head := formatLog(InfoLvl, skipLevel)
	logger.Info(head + fmt.Sprint(fields...))
}

func Warn(fields ...interface{}) {
	head := formatLog(WarnLvl, skipLevel)
	logger.Warn(head + fmt.Sprint(fields...))
}

func Error(fields ...interface{}) {
	head := formatLog(ErrorLvl, skipLevel)
	logger.Error(head + fmt.Sprint(fields...))
}

func Critical(fields ...interface{}) {
	head := formatLog(CriticalLvl, skipLevel)
	logger.Fatal(head + fmt.Sprint(fields...))
}

func Tracef(template string, fields ...interface{}) {
	head := formatLog(TraceLvl, skipLevel)
	logger.Debug(head + fmt.Sprintf(template, fields...))
}

func Debugf(template string, fields ...interface{}) {
	head := formatLog(DebugLvl, skipLevel)
	logger.Debug(head + fmt.Sprintf(template, fields...))
}

func Infof(template string, fields ...interface{}) {
	head := formatLog(InfoLvl, skipLevel)
	logger.Info(head + fmt.Sprintf(template, fields...))
}

func Warnf(template string, fields ...interface{}) {
	head := formatLog(WarnLvl, skipLevel)
	logger.Warn(head + fmt.Sprintf(template, fields...))
}

func Errorf(template string, fields ...interface{}) {
	head := formatLog(ErrorLvl, skipLevel)
	logger.Error(head + fmt.Sprintf(template, fields...))
}

func Criticalf(template string, fields ...interface{}) {
	head := formatLog(CriticalLvl, skipLevel)
	logger.Fatal(head + fmt.Sprintf(template, fields...))
}
