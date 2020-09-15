package server

import (
	"context"
	"github.com/Gitforxuyang/walleManage/config"
	"github.com/Gitforxuyang/walleManage/util/utils"
	"github.com/coreos/etcd/clientv3"
	"sync"
)

//api控制器
type apiController struct {
	//访问地址对应的方法
	apis map[string]*Method
}

type HttpMethod struct {
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
	ETCD_API_KEY              = "/eva/walle/api"
	REDIS_API_KEY             = "walle:api"
	ETCD_WALLE_SERVICE_PREFIX = "/eva/walle/service/"
)

type Service struct {
	Package string            `json:"package"`
	Name    string            `json:"name"`
	AppId   string            `json:"appId"`
	Methods map[string]Method `json:"methods"`
}
type Method struct {
	Req  map[string]string `json:"req"`
	Resp map[string]string `json:"resp"`
}

func GetService() ([]*Service, error) {
	client := config.GetEtcdClient()
	res, err := client.Get(context.TODO(), ETCD_WALLE_SERVICE_PREFIX,clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	svcs := make([]*Service, 0)
	for _, v := range res.Kvs {
		service := Service{}
		err = utils.JsonToStruct(string(v.Value), &service)
		if err != nil {
			return nil, err
		}
		svcs = append(svcs, &service)
	}
	return svcs, nil
}
