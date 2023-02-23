package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLog struct {
	SugaredLogger *zap.SugaredLogger
	RegularLogger *zap.Logger
	Cfg           *zap.Config
}

var zlogInstance *ZapLog

type LogLevel int

const (
	LevelProd LogLevel = iota
	LevelDev
)

func New(level ...LogLevel) *ZapLog {
	if len(level) == 0 {
		level = append(level, LevelProd)
	}
	zaplog := &ZapLog{}
	var logCfg zap.Config
	if level[0] == LevelProd {
		logCfg = zap.NewProductionConfig()
	} else {
		logCfg = zap.NewDevelopmentConfig()
		logCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	zaplog.Cfg = &logCfg
	logger, err := zaplog.Cfg.Build()
	if err != nil {
		panic(err)
	}
	zaplog.SugaredLogger = logger.Sugar()
	zaplog.RegularLogger = logger
	zlogInstance = zaplog
	return zlogInstance
}

func Log() *zap.SugaredLogger {
	return zlogInstance.SugaredLogger
}

func GetLogger() *zap.Logger {
	return zlogInstance.RegularLogger
}
