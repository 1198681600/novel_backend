package utils

import (
	"gitea.peekaboo.tech/peekaboo/crushon-backend/internal/global"
	"go.uber.org/zap"
	"time"
)

func FunctionTimeProfile(name string, f func() error) error {
	global.Logger.Info("FunctionTimeProfile Start", zap.String("name", name))
	timestamp := time.Now()
	defer func() {
		global.Logger.Info("FunctionTimeProfile Finish", zap.String("name", name), zap.Duration("duration", time.Since(timestamp)))
	}()
	return f()
}
