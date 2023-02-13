package zlog

import (
	"go.uber.org/zap"
)

var sugaredLogger *zap.SugaredLogger

// InitZap
// @param cfg 配置选项
// @Description: 初始化logger
func InitZap(cfg *zap.Config) {
	if cfg == nil {
		panic("cfg is nil")
	}
	logger, _ := cfg.Build()
	sugaredLogger = logger.Sugar()
}

func NewProductionConfig() zap.Config {
	return zap.NewProductionConfig()
}

func NewDevelopmentConfig() zap.Config {
	return zap.NewDevelopmentConfig()
}
