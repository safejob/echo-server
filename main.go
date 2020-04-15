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

	router := gin.Default()

	router.Any("/*id", func(c *gin.Context) {
		//b, _ := json.Marshal(c.Request.Header)
		// b := fmt.Sprintf("%#+v\n", c.Request.Header)
		fmt.Println()

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
