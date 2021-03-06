package log

import "github.com/blessmylovexy/pkg/logger"

var log = logger.NewLogger()

func init() {
	log.SetCallerSkip(2)
}

/*Debug Debug log*/
func Debug(args ...interface{}) {
	log.Debug(args...)
}

/*Debugf Debug format log*/
func Debugf(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

/*Info Info log*/
func Info(args ...interface{}) {
	log.Info(args...)
}

/*Infof Info format log*/
func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

/*Warn Warn log*/
func Warn(args ...interface{}) {
	log.Warn(args...)
}

/*Warnf Warn format log*/
func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

/*Error Error log*/
func Error(args ...interface{}) {
	log.Error(args...)
}

/*Errorf Error format log*/
func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

/*Panic Panic log*/
func Panic(args ...interface{}) {
	log.Panic(args...)
}

/*Panicf Panic format log*/
func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}

/*Fatal Fatal log*/
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

/*Fatalf Fatal format log*/
func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}

/*SetLogLevel 设置日志级别*/
func SetLogLevel(level string) {
	log.SetLogLevel(level)
}

/*SetProjectName 设置日志名字*/
func SetProjectName(projectName string) {
	log.SetProjectName(projectName)
}

/*SetStacktraceLevel 设置堆栈跟踪的日志级别*/
func SetStacktraceLevel(level string) {
	log.SetStacktraceLevel(level)
}

/*SetLogFile 设置日志文件*/
func SetLogFile(fileName string, rotationTime, rotationCount uint) {
	log.SetLogFile(fileName, rotationTime, rotationCount)
}
