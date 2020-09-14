package catch

import (
	"github.com/Gitforxuyang/walleManage/server/vo"
	"github.com/Gitforxuyang/walleManage/util/logger"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

func ServerCatch() gin.HandlerFunc {
	log := logger.GetLogger()
	return func(ctx *gin.Context) {
		defer func() {
		if e := recover(); e != nil {
			log.Error(ctx, "发生panic", logger.Fields{"e": e})
			sentry.CaptureException(errors.New(e))
		}
		}()
		ctx.Next()
		res:=ctx.Value("res")
		err:=ctx.Value("err")
		if err!=nil{
			msg:=err.(error).Error()
			ctx.JSON(500,vo.Resp{Code:500,Message:msg})
		}else{
			ctx.JSON(200,vo.Resp{Data:res})
		}
	}
}

