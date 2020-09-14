package server

import (
	"context"
	"github.com/Gitforxuyang/walleManage/client/redis"
	"github.com/Gitforxuyang/walleManage/config"
	"github.com/Gitforxuyang/walleManage/util/logger"
	"github.com/Gitforxuyang/walleManage/util/utils"
	"sync"
)

//api控制器
type apiController struct {
	//访问地址对应的方法
	apis map[string]*Method
}

type Method struct {
	Service string   `json:"service"`
	Method  string   `json:"method"`
	Plugins []Plugin `json:"plugins"` //需要通过哪几个插件
}

type Plugin struct {
	Name  string `json:"name"`
	Param string `json:"param"`
}

var (
	apiOnce sync.Once
	api     *apiController
)

const (
	ETCD_API_KEY = "/eva/walleManage/api"
)

func InitApi() {
	apiOnce.Do(func() {
		api = &apiController{}
		api.apis = make(map[string]*Method)
		refreshApi()
		go watch()
	})
}

//刷新api
func refreshApi() {
	client := redis.GetRedisClient()
	res := client.HGetAll(context.TODO(), "walleManage:api")
	if res.Err() != nil {
		logger.GetLogger().Error(context.TODO(), "获取redis报错", logger.Fields{"err": res.Err()})
		return
	}
	maps, err := res.Result()
	if err != nil {
		logger.GetLogger().Error(context.TODO(), "获取redis报错", logger.Fields{"err": err})
		return
	}
	for k, v := range maps {
		m := Method{}
		err := utils.JsonToStruct(v, &m)
		if err != nil {
			logger.GetLogger().Error(context.TODO(), "api转json报错", logger.Fields{"err": err})
			continue
		}
		api.apis[k] = &m
	}
	//logger.GetLogger().Info(context.TODO(), "刷新api成功", logger.Fields{
	//	"api": api.apis,
	//})
}
func watch() {
	etcd := config.GetEtcdClient()
	log := logger.GetLogger()
	for {
		chs := etcd.Watch(context.TODO(), ETCD_API_KEY)
		for v := range chs {
			for _, event := range v.Events {
				version := string(event.Kv.Value)
				refreshApi()
				log.Info(context.TODO(), "etcd更新api", logger.Fields{
					"version": version,
				})
			}
		}
	}
}

