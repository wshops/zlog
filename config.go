package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetLevel
// @param level 日志等级 -1 to 5
// @param config logger配置
// @Description:
func SetLevel(level int8, cfg *zap.Config) {
	if cfg == nil {
		panic("cfg is nil")
	}
	cfg.Level = zap.NewAtomicLevelAt(zapcore.Level(level))
}
