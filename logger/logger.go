package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zlog struct {
	_logger           *zap.SugaredLogger //SugaredLogger
	callerSkip        int                //CallerSkip次数
	logLevel          zap.AtomicLevel    //日志记录级别
	isStdOut          bool               //是否输出到console
	projectName       string             //项目名称
	fileLogName       string             //日志保存路径和名称
	fileLogMaxSize    int                //日志分割的尺寸 MB
	fileLogMaxAge     int                //分割日志保存的时间 day
	fileLogMaxBackups int                // 文件最多备份个数
	fileLogLocalTime  bool               // 是否为备份文件生成时间戳
	fileLogCompress   bool               // 是否启用备份压缩
	stacktrace        string             //记录堆栈的级别
}

func NewLogger() *zlog {
	log := &zlog{
		callerSkip:        1,
		logLevel:          zap.NewAtomicLevel(),
		isStdOut:          true,
		projectName:       "",
		fileLogName:       "",
		fileLogMaxSize:    1024,
		fileLogMaxAge:     7,
		fileLogMaxBackups: 15,
		fileLogLocalTime:  true,
		fileLogCompress:   true,
		stacktrace:        "panic",
	}
	log.logLevel.SetLevel(zap.InfoLevel)
	log.build()
	return log
}

/*build 用当前配置构建logger*/
func (z *zlog) build() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
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
		// 配置log轮滚
		hook := lumberjack.Logger{
			Filename:   z.fileLogName,
			MaxSize:    z.fileLogMaxSize,
			MaxAge:     z.fileLogMaxAge,
			MaxBackups: z.fileLogMaxBackups,
			LocalTime:  z.fileLogLocalTime,
			Compress:   z.fileLogCompress,
		}
		if z.isStdOut {
			core = zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
				z.logLevel,
			)
		} else {
			core = zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),
				z.logLevel,
			)
		}

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
func (z *zlog) Debug(args ...interface{}) {
	z._logger.Debug(args...)
}

/*Debugf Debug format log*/
func (z *zlog) Debugf(template string, args ...interface{}) {
	z._logger.Debugf(template, args...)
}

/*Info Info log*/
func (z *zlog) Info(args ...interface{}) {
	z._logger.Info(args...)
}

/*Infof Info format log*/
func (z *zlog) Infof(template string, args ...interface{}) {
	z._logger.Infof(template, args...)
}

/*Warn Warn log*/
func (z *zlog) Warn(args ...interface{}) {
	z._logger.Warn(args...)
}

/*Warnf Warn format log*/
func (z *zlog) Warnf(template string, args ...interface{}) {
	z._logger.Warnf(template, args...)
}

/*Error Error log*/
func (z *zlog) Error(args ...interface{}) {
	z._logger.Error(args...)
}

/*Errorf Error format log*/
func (z *zlog) Errorf(template string, args ...interface{}) {
	z._logger.Errorf(template, args...)
}

/*Panic Panic log*/
func (z *zlog) Panic(args ...interface{}) {
	z._logger.Panic(args...)
}

/*Panicf Panic format log*/
func (z *zlog) Panicf(template string, args ...interface{}) {
	z._logger.Panicf(template, args...)
}

/*Fatal Fatal log*/
func (z *zlog) Fatal(args ...interface{}) {
	z._logger.Fatal(args...)
}

/*Fatalf Fatal format log*/
func (z *zlog) Fatalf(template string, args ...interface{}) {
	z._logger.Fatalf(template, args...)
}

/*SetLogLevel 设置日志级别*/
func (z *zlog) SetLogLevel(level string) {
	z.logLevel.SetLevel(getLevelByString(level))
}

/*SetProjectName 设置日志名字*/
func (z *zlog) SetProjectName(projectName string) {
	z.projectName = projectName
	z._logger = z._logger.Named(projectName)
}

/*SetStacktraceLevel 设置堆栈跟踪的日志级别*/
func (z *zlog) SetStacktraceLevel(level string) {
	z.stacktrace = level
	z.build()
}

/*SetCallerSkip 设置日志名字*/
func (z *zlog) SetCallerSkip(callerSkip int) {
	z.callerSkip = callerSkip
	z.build()
}
