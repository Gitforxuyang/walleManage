package server

import (
	"errors"
	"fmt"
	mysql2 "github.com/Gitforxuyang/walleManage/client/mysql"
	"github.com/Gitforxuyang/walleManage/server/dao"
	"github.com/Gitforxuyang/walleManage/server/dto"
	"github.com/Gitforxuyang/walleManage/util/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func controller() {
	mysql := mysql2.GetMySQLClient()
	r.POST("/v1/walle/manage/login", handler(func(ctx *gin.Context) (i interface{}, e error) {
		body := dto.Login{}
		buf,err:=ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		err = utils.JsonToStruct(string(buf), &body)
		if err != nil {
			return nil, err
		}
		if body.UserName == "" || body.Password == "" {
			return nil, errors.New("参数错误")
		}
		var admin dao.Admin
		ok, err := mysql.SqlMapClient("user_select_by_login",
			map[string]interface{}{"user_name": body.UserName}).Get(&admin)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("账号不存在或密码错误")
		}
		if body.Password == admin.Password {
			return map[string]interface{}{"token": admin.Token, "name": admin.Name, "id": admin.Id}, nil
		}
		return nil, errors.New("登陆失败")
	}))
	r.Group("/v1/walle/manage", func(ctx *gin.Context) {
		//做登陆状态验证
		token:=ctx.GetHeader("token")
		exists,err:=mysql.Exist(&dao.Admin{Token:token})
		if err!=nil{
			ctx.Set("err",err)
			ctx.Abort()
		}
		if !exists{
			ctx.Set("err",errors.New("token验证失败"))
			ctx.Abort()
		}
	}).GET("/service", func(ctx *gin.Context) {
		fmt.Println("service")
	})
}
