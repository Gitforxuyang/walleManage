package config

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Gitforxuyang/walleManage/util/utils"
	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/viper"
	"strings"
	"sync"
	"time"
)

//配置发送变更时的通知
type ChangeNotify func(config map[string]interface{})

type HttpClientConfig struct {
	Endpoint string
	Timeout  int64
	MaxConn  int //最大连接数
}
type RedisConfig struct {
	Addr         string
	Password     string
	DB           int
	PoolSize     int
	MinIdleConns int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
type EvaConfig struct {
	name              string
	port              int32
	env               string
	v                 *viper.Viper
	changeNotifyFuncs []ChangeNotify
	http              map[string]*HttpClientConfig
	etcd              []string
	sentryDSN         string
	redis             map[string]*RedisConfig
}

var (
	config     *EvaConfig
	configOnce sync.Once
	etcdClient *clientv3.Client
)

func Init() {
	configOnce.Do(func() {
		config = &EvaConfig{}
		config.changeNotifyFuncs = make([]ChangeNotify, 0)
		config.redis = make(map[string]*RedisConfig)
		config.etcd = make([]string, 0, 3)
		v := viper.New()
		v.SetConfigName("config.default")
		v.AddConfigPath("./conf")
		v.SetConfigType("json")
		err := v.ReadInConfig()
		utils.Must(err)
		v.BindEnv("ENV")
		env := v.GetString("ENV")
		if env == "" {
			env = "local"
		}
		config.env = env
		v.SetConfigName(fmt.Sprintf("config.%s", env))
		err = v.MergeInConfig()
		utils.Must(err)
		config.name = v.GetString("name")
		if config.name == "" {
			panic("配置文件中name不能为空")
		}
		err = v.UnmarshalKey("etcd", &config.etcd)
		utils.Must(err)
		if len(config.etcd) == 0 {
			panic("配置文件中etcd不能为空")
		}

		client, err := clientv3.New(clientv3.Config{
			Endpoints:   config.etcd,
			DialTimeout: time.Second * 3,
		})
		utils.Must(err)
		etcdClient = client
		resp, err := client.Get(context.TODO(), fmt.Sprintf("%s%s", ETCD_CONFIG_PREFIX, "global"))
		utils.Must(err)
		if len(resp.Kvs) == 0 {
			panic("配置中心global未找到")
		}
		v.MergeConfig(bytes.NewBuffer(resp.Kvs[0].Value))
		resp, err = client.Get(context.TODO(), fmt.Sprintf("%s%s", ETCD_CONFIG_PREFIX, config.name))
		utils.Must(err)
		if len(resp.Kvs) == 0 {
			panic(fmt.Sprintf("配置中心%s未找到", config.name))
		}
		v.MergeConfig(bytes.NewBuffer(resp.Kvs[0].Value))

		config.port = v.GetInt32("port")
		if config.port == 0 {
			panic("配置文件中port不能为空")
		}
		config.v = v
		err = v.UnmarshalKey("http", &config.http)
		utils.Must(err)
		err = v.UnmarshalKey("redis", &config.redis)
		utils.Must(err)
		if utils.IsNil(v.Get("trace")) {
			panic("trace设置不能为空")
		}
		config.sentryDSN = v.GetString("sentry")
		//watch(client)
	})
}

func GetConfig() *EvaConfig {
	return config
}

func GetEtcdClient() *clientv3.Client {
	return etcdClient
}
func RegisterNotify(f ChangeNotify) {
	config.changeNotifyFuncs = append(config.changeNotifyFuncs, f)
}

func (m *EvaConfig) changeNotify(config map[string]interface{}) {
	for _, v := range m.changeNotifyFuncs {
		v(config)
	}
}
func (m *EvaConfig) GetName() string {
	return m.name
}
func (m *EvaConfig) GetPort() int32 {
	return m.port
}
func (m *EvaConfig) GetENV() string {
	return m.env
}

func (m *EvaConfig) GetHttp(http string) *HttpClientConfig {
	c := m.http[strings.ToLower(http)]
	if c == nil {
		panic(fmt.Sprintf("http：%s配置未找到", http))
	}
	return c
}

func (m *EvaConfig) GetEtcd() []string {
	if m.etcd == nil {
		panic(fmt.Sprintf("etcd配置未找到"))
	}
	return m.etcd
}
func GetSentry() string {
	return config.sentryDSN
}

func (m *EvaConfig) GetRedis(name string) *RedisConfig {
	c := m.redis[strings.ToLower(name)]
	if c == nil {
		panic(fmt.Sprintf("redis：%s配置未找到", name))
	}
	return c
}
