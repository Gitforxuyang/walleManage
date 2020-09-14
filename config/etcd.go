package config

const (
	//服务注册前缀
	ETCD_CONFIG_PREFIX = "/eva/config/"
)
//
//func watch(client *clientv3.Client) {
//	go func() {
//		for true {
//		LOOP:
//			w := client.Watch(context.TODO(), fmt.Sprintf("%s%s", ETCD_CONFIG_PREFIX, config.name))
//			for wresp := range w {
//				for _, ev := range wresp.Events {
//					err := config.v.MergeConfig(bytes.NewBuffer(ev.Kv.Value))
//					if err != nil {
//						logger.GetLogger().Error(context.TODO(), "动态更新配置出错", logger.Fields{"err": err})
//						goto LOOP
//					}
//				}
//			}
//
//		}
//	}()
//}
