package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"os"
)

// k8s 新服务默认容器 根据环境变量设置监听端口 并响应任意请求 状态码 200
func main() {
	Code := flag.Int("s", 200, "响应状态码")
	flag.Parse()

	Port := os.Getenv("APP_PORT")
	if Port == "" {
		Port = "80"
	}

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.Any("/*id", func(c *gin.Context) {
		//c.String(200, c.Param("id"))
		//c.String(200, string(b))
		c.JSON(*Code, c.Request.Header)
	})

	router.Run(":" + Port)
}
