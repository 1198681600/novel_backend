package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"log/slog"
	"os"
)

var (
	Logger *zap.Logger
)

func initLogger() {
	var config zap.Config
	if os.Getenv("LOG_ENV") == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}
	if viper.GetBool("debug.logger") {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		slog.Info("DEBUG Level Enabled")
	}
	Logger, _ = config.Build()

	zap.ReplaceGlobals(Logger)

	slog.SetDefault(slog.New(zapslog.NewHandler(Logger.Core(), &zapslog.HandlerOptions{
		AddSource: true,
	})))

	slog.Debug("DEBUG SLOG")
	Logger.Debug("DEBUG ZAP LOG")
}
