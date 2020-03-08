package logger

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*Zlog zap logger*/
type Zlog struct {
	_logger          *zap.SugaredLogger //SugaredLogger
	callerSkip       int                //CallerSkip次数
	logLevel         zap.AtomicLevel    //日志记录级别
	projectName      string             //项目名称
	fileName         string             //日志保存路径和名称
	fileRotationTime int                //分割日志的时间间隔（小时）
	fileMaxAge       int                //日志保存的时间（小时）
	stacktrace       string             //记录堆栈的级别
}

/*NewLogger New logger*/
func NewLogger() *Zlog {
	log := &Zlog{
		callerSkip:       1,
		logLevel:         zap.NewAtomicLevel(),
		projectName:      "",
		fileName:         "",
		fileRotationTime: 24,
		fileMaxAge:       7 * 24,
		stacktrace:       "panic",
	}
	log.logLevel.SetLevel(zap.InfoLevel)
	log.build()
	return log
}

/*build 用当前配置构建logger*/
func (z *Zlog) build() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		NameKey:        "N",
		LevelKey:       "L",
		MessageKey:     "M",
		CallerKey:      "C",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 初始化core
	var core zapcore.Core

	if z.fileLogName != "" {
		hook, err := rotatelogs.New(
			z.fileLogName+".%Y%m%d",
			rotatelogs.WithLinkName(z.fileLogName),
			rotatelogs.WithMaxAge(time.Hour*z.fileMaxAge),
			rotatelogs.WithRotationTime(time.Hour*z.fileRotationTime),
		)
		if err != nil {
			panic(err)
		}
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),
			z.logLevel,
		)

	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			z.logLevel,
		)
	}

	log := zap.New(core, zap.AddCaller(), zap.AddStacktrace(getLevelByString(z.stacktrace)), zap.AddCallerSkip(z.callerSkip))
	z._logger = log.Sugar()

	if z.projectName != "" {
		z._logger = z._logger.Named(z.projectName)
	}
}

/*getLevelByString 通过字符串获取zaplog等级*/
func getLevelByString(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.DebugLevel
	}
}

/*Debug Debug log*/
func (z *Zlog) Debug(args ...interface{}) {
	z._logger.Debug(args...)
}

/*Debugf Debug format log*/
func (z *Zlog) Debugf(template string, args ...interface{}) {
	z._logger.Debugf(template, args...)
}

/*Info Info log*/
func (z *Zlog) Info(args ...interface{}) {
	z._logger.Info(args...)
}

/*Infof Info format log*/
func (z *Zlog) Infof(template string, args ...interface{}) {
	z._logger.Infof(template, args...)
}

/*Warn Warn log*/
func (z *Zlog) Warn(args ...interface{}) {
	z._logger.Warn(args...)
}

/*Warnf Warn format log*/
func (z *Zlog) Warnf(template string, args ...interface{}) {
	z._logger.Warnf(template, args...)
}

/*Error Error log*/
func (z *Zlog) Error(args ...interface{}) {
	z._logger.Error(args...)
}

/*Errorf Error format log*/
func (z *Zlog) Errorf(template string, args ...interface{}) {
	z._logger.Errorf(template, args...)
}

/*Panic Panic log*/
func (z *Zlog) Panic(args ...interface{}) {
	z._logger.Panic(args...)
}

/*Panicf Panic format log*/
func (z *Zlog) Panicf(template string, args ...interface{}) {
	z._logger.Panicf(template, args...)
}

/*Fatal Fatal log*/
func (z *Zlog) Fatal(args ...interface{}) {
	z._logger.Fatal(args...)
}

/*Fatalf Fatal format log*/
func (z *Zlog) Fatalf(template string, args ...interface{}) {
	z._logger.Fatalf(template, args...)
}

/*SetLogLevel 设置日志级别*/
func (z *Zlog) SetLogLevel(level string) {
	z.logLevel.SetLevel(getLevelByString(level))
}

/*SetProjectName 设置日志名字*/
func (z *Zlog) SetProjectName(projectName string) {
	z.projectName = projectName
	z._logger = z._logger.Named(projectName)
}

/*SetStacktraceLevel 设置堆栈跟踪的日志级别*/
func (z *Zlog) SetStacktraceLevel(level string) {
	z.stacktrace = level
	z.build()
}

/*SetCallerSkip 设置日志名字*/
func (z *Zlog) SetCallerSkip(callerSkip int) {
	z.callerSkip = callerSkip
	z.build()
}

/*SetLogFile 设置日志文件*/
func (z *Zlog) SetLogFile(fileName string, fileRotationTime, fileMaxAge int) {
	z.fileName = fileName
	z.fileRotationTime = fileRotationTime
	z.fileMaxAge = fileMaxAge
	z.build()
}
