package global

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"log/slog"
)

var (
	c = cron.New(cron.WithSeconds())
)

func init() {
	c.Start()
}

func AddCronJob(fn func() (err error), spec string) {
	if _, err := c.AddFunc(spec, func() {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("cron job panic", slog.Any("err", err))
			}
		}()
		err := fn()
		if err != nil {
			Logger.Error("cron job error", zap.String("id", spec), zap.Error(err))
		}
	}); err != nil {
		Logger.Error("failed to add cron job", zap.Error(err))
		panic(err)
	}
}
