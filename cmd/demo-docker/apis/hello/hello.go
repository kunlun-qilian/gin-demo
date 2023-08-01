package hello

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunlun-qilian/gin-demo/cmd/demo-docker/global"
	"net/http"
)

func HelloRouter(r *gin.RouterGroup) {
	r.GET("/hello", HelloGet)
}

type HelloResp struct {
	Data string `json:"data"`
}

func HelloGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HelloResp{
		Data: fmt.Sprintf("hello testEnv is %s", global.Config.TestEnv),
	})
}
