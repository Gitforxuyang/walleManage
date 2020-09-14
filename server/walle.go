package server

import (
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
	ETCD_API_KEY = "/eva/walle/api"
	REDIS_API_KEY ="walle:api"
	ETCD_WALLE_SERVICE_PREFIX = "/eva/walle/service/"
)
