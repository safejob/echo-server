package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	Port := flag.String("p", "8080", "服务监听端口")
	Code := flag.Int("s", 200, "响应状态码")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Any("/*id", func(c *gin.Context) {
		fmt.Println("<================= 请求属性 ====================>")
		fmt.Printf("Host => %s\nMethod => %s\nHTTP => %s\nRemoteAddr => %s\nRequestURI => %s\nURL => %s\n",
			c.Request.Host, c.Request.Method, c.Request.Proto,
			c.Request.RemoteAddr, c.Request.RequestURI, c.Request.URL)

		fmt.Println("<================= 请求头 ====================>")
		for k, v := range c.Request.Header {
			for i, j := range v {
				fmt.Printf("%s -> %d: %s\n", k, i, j)
			}
		}

		fmt.Println("<================= 请求体 ====================>")
		body, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println(string(body))

		//c.String(200, c.Param("id"))
		//c.String(200, string(b))
		c.JSON(*Code, c.Request.Header)
	})

	router.Run(":" + *Port)
}
