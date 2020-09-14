package server

import (
	"context"
	"fmt"
	"github.com/Gitforxuyang/walleManage/config"
	"github.com/Gitforxuyang/walleManage/middleware/catch"
	"github.com/Gitforxuyang/walleManage/util/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//注册关闭服务时的回调
type RegisterShutdown func()

var (
	shutdownFunc []RegisterShutdown = make([]RegisterShutdown, 0)
)

func InitServer() {
	conf := config.GetConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	//内部转发
	r.Group("/rpc",
		func(ctx *gin.Context) {
		}).
		Use(catch.ServerCatch()).
		POST("/:Service/:Method", func(ctx *gin.Context) {

		})


	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.GetPort()),
		Handler: r,
	}
	go func() {
		srv.ListenAndServe()
		time.Sleep(time.Millisecond * 500)
	}()
	time.Sleep(time.Millisecond * 200)
	logger.GetLogger().Info(context.TODO(), "server started", logger.Fields{
		"port":   config.GetConfig().GetPort(),
		"server": config.GetConfig().GetName(),
		"env":    config.GetConfig().GetENV(),
	})
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	logger.GetLogger().Info(context.TODO(), "signal", logger.Fields{
		"signal": s.String(),
	})
	srv.Shutdown(context.TODO())
	logger.GetLogger().Info(context.TODO(), "server stop", logger.Fields{})
}
