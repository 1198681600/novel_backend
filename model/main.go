package model

import (
	"go.uber.org/zap"
	"novel_backend/global"
	"os"
)

func RegisterTables() {
	err := global.DB.AutoMigrate(
		Novel{},
		Book{},
	)
	if err != nil {
		zap.L().Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	zap.L().Info("register table success")
}
