package app

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)
// InitLogger 初始化日志
func InitLogger() error {

	conf := zapcore.EncoderConfig{
		TimeKey: "time",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "file",
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	//基础信息
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})
	//达到系统警告、错误、致命级别
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	encoder := zapcore.NewConsoleEncoder(conf)
	infoHook,err := getHook("./logs/info/")
	if err != nil {
		return err
	}
	warnHook,err := getHook("./logs/warn/")
	if err != nil {
		return err
	}
	// 设置日志级别
	core := zapcore.NewTee(
		zapcore.NewCore(encoder,zapcore.AddSync(infoHook),infoLevel),
		zapcore.NewCore(encoder,zapcore.AddSync(warnHook),warnLevel),
		zapcore.NewCore(
			encoder,
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			zap.DebugLevel,
			),
		)

	if Config.Debug {
		caller := zap.AddCaller()
		development := zap.Development()
		Logger = newLogger(zap.New(core,caller,development))
	}else{
		Logger = newLogger(zap.New(core))
	}
	return nil
}

func getHook(filename string) (io.Writer,error) {
	return rotatelogs.New(
		filename+"%Y%m%d/%H.log",
		rotatelogs.WithLinkName(filename+"symlink"),
		rotatelogs.WithMaxAge(time.Hour*24*30),    // 保存30天
		rotatelogs.WithRotationTime(time.Hour*24), //切割频率 24小时
	)
}

type AppLogger struct {
	logger *zap.Logger
}
func newLogger(logger *zap.Logger) *AppLogger {
	l := new(AppLogger)
	l.logger = logger
	return l
}
func (l *AppLogger) Debug(format string,args ...interface{})  {
	l.logger.Debug(fmt.Sprintf(format,args...))
}
func (l *AppLogger) Info(format string,args ...interface{})  {
	l.logger.Info(fmt.Sprintf(format,args...))
}
func (l *AppLogger) Warn(format string,args ...interface{})  {
	l.logger.Warn(fmt.Sprintf(format,args...))
}
func (l *AppLogger) Error(format string,args ...interface{})  {
	l.logger.Error(fmt.Sprintf(format,args...))
}
func (l *AppLogger) Fatal(format string,args ...interface{})  {
	l.logger.Fatal(fmt.Sprintf(format,args...))
}

func (l *AppLogger) PrintDebug(msg string, fields ...zap.Field)  {
	l.logger.Debug(msg,fields...)
}
func (l *AppLogger) PrintInfo(msg string, fields ...zap.Field)  {
	l.logger.Info(msg,fields...)
}
func (l *AppLogger) PrintWarn(msg string, fields ...zap.Field)  {
	l.logger.Warn(msg,fields...)
}
func (l *AppLogger) PrintError(msg string, fields ...zap.Field)  {
	l.logger.Error(msg,fields...)
}
func (l *AppLogger) PrintFatal(msg string, fields ...zap.Field)  {
	l.logger.Fatal(msg,fields...)
}