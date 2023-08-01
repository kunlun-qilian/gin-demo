package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/kunlun-qilian/gin-demo/cmd/demo-docker/apis/hello"
)

func NewRooter(r *gin.Engine) {
	v1 := r.Group("/demo-docker/api/v1")
	hello.HelloRouter(v1)
}
