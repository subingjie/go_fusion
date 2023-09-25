package main

/*
安装gin

*/
import (
	"fmt"
	"go_fusion/config"
	"go_fusion/initialize"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("欢迎使用golang融合系统！")
	ginDemo()
}

func ginDemo() {
	r := gin.Default()

	config, err := config.LoadConfiguation()
	if err != nil {
		panic(err)
	}

	initialize.InitRouter(r)

	port := config.Server.Port

	r.Run(":" + port) // 监听并在
}
