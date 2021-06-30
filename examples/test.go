package main

import (
	"github.com/gin-gonic/gin"
)

func mw() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 洋葱模型，层层中间件向核心处理，Next层层由核心向外处理
		c.Next()
		c.GetInt("key")
	}
}

func main() {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run()
}
