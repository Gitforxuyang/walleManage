package server

import (
	"github.com/gin-gonic/gin"
)

func controller(){
	r.POST("/v1/walle/manage/login", handler(func(ctx *gin.Context) (i interface{}, e error) {
		return map[string]string{"d":"123"},nil
	}))
	r.Group("/v1/walle/manage", func(ctx *gin.Context) {
		//做登陆状态验证
	}).GET("/service", func(ctx *gin.Context) {

	})
}

