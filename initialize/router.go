package initialize

import (
	"go_fusion/relay/upload"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.POST("/upload", upload.UploadFile)
}
