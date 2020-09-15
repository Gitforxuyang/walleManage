package etcd
//
//import (
//	"github.com/Gitforxuyang/walleManage/config"
//	"github.com/Gitforxuyang/walleManage/util/utils"
//	"github.com/coreos/etcd/clientv3"
//	"sync"
//	"time"
//)
//
//const (
//	//服务注册前缀
//	ETCD_SERVICE_PREFIX = "/eva/service/"
//)
//
//type ServiceNode struct {
//	Name     string `json:"name"`     //服务名
//	Id       string `json:"id"`       //节点id 服务启动时随机生成的唯一id
//	Endpoint string `json:"endpoint"` //服务的访问地址
//}
//
//var (
//	client   *clientv3.Client
//	etcdOnce sync.Once
//)
//
//func GetClient() *clientv3.Client {
//	if client == nil {
//		panic("client不存在")
//	}
//	return client
//}
//func Init() {
//	etcdOnce.Do(func() {
//		var err error
//		client, err = clientv3.New(clientv3.Config{
//			Endpoints:   config.GetConfig().GetEtcd(),
//			DialTimeout: time.Second * 3,
//		})
//		utils.Must(err)
//	})
//}
