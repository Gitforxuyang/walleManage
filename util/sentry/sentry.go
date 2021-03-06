package sentry

import (
	"context"
	"github.com/Gitforxuyang/walleManage/config"
	"github.com/Gitforxuyang/walleManage/util/logger"
	"github.com/getsentry/sentry-go"
	"sync"
)

var (
	sentryOnce sync.Once
)

func Init() {
	sentryOnce.Do(func() {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: config.GetSentry(),
		})
		if err != nil {
			logger.GetLogger().Error(context.TODO(), "sentry error", logger.Fields{
				"err": err,
			})
		}
	})
}
