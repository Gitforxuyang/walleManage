package main

import (
	"github.com/Gitforxuyang/walleManage/client/mysql"
	"github.com/Gitforxuyang/walleManage/client/redis"
	"github.com/Gitforxuyang/walleManage/config"
	"github.com/Gitforxuyang/walleManage/registory/etcd"
	"github.com/Gitforxuyang/walleManage/server"
	"github.com/Gitforxuyang/walleManage/util/logger"
	"github.com/Gitforxuyang/walleManage/util/sentry"
)

func main() {
	config.Init()
	conf := config.GetConfig()
	logger.Init(conf.GetName())
	sentry.Init()
	redis.Init()
	mysql.Init()
	etcd.Init()
	server.InitServer()
}
