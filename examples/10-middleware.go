package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求之前。。。
		log.Println("之前。。")
		c.Next()
		log.Println("之后。。")

		// 请求之后。。。

		latency := time.Since(t)
		log.Print(latency)

		// 访问即将发送的响应码
		status := c.Writer.Status()
		log.Println(status)

	}
}

func main() {
	r := gin.New()
	// 使用自定义的Logger 中间件
	r.Use(Logger())

	// 定义路由
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		// 打印 example 值12345
		log.Println(example)
	})

	r.Run(":8080")
}
