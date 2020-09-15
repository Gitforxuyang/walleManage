package mysql

import (
	"fmt"
	"github.com/Gitforxuyang/walleManage/config"
	"github.com/Gitforxuyang/walleManage/util/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"sync"
)

var (
	mysqlOnce sync.Once
	client    *xorm.Engine
)

func Init() {
	mysqlOnce.Do(func() {
		conf := config.GetConfig()
		fmt.Println(conf.GetMySQL("walle").Addr)
		engine, err := xorm.NewMySQL(xorm.MYSQL_DRIVER, conf.GetMySQL("walle").Addr)
		utils.Must(err)
		client = engine
	})
}

func GetMySQLClient() *xorm.Engine {
	if client == nil {
		panic("mysql客户端不存在")
	}
	return client
}
