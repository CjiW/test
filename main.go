package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(context *gin.Context) {
		context.JSON(200,gin.H{"msg":"ok","time":time.Now()})
	})
	_=r.Run(":8089")
}
