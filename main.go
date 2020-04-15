package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	Port := flag.String("p", "8080", "服务监听端口")
	Code := flag.Int("s", 200, "响应状态码")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Any("/*id", func(c *gin.Context) {
		//b, _ := json.Marshal(c.Request.Header)
		// b := fmt.Sprintf("%#+v\n", c.Request.Header)
		fmt.Println()
		fmt.Printf("Host => %s\nMethod => %s\nHTTP => %s\nRemoteAddr => %s\nRequestURI => %s\nURL => %s\n",
			c.Request.Host, c.Request.Method, c.Request.Proto,
			c.Request.RemoteAddr, c.Request.RequestURI, c.Request.URL)
		fmt.Println("-----")

		for k, v := range c.Request.Header {
			for i, j := range v {
				fmt.Printf("%s -> %d: %s\n", k, i, j)
			}
		}

		//c.String(200, c.Param("id"))
		//c.String(200, string(b))
		c.JSON(*Code, c.Request.Header)
	})

	router.Run(":" + *Port)
}
