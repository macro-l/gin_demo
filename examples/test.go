package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func mw() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 洋葱模型，层层中间件向核心处理，Next层层由核心向外处理
		//c.Next()
		//c.Abort()

		c.GetHeader("")

		c.AbortWithStatusJSON(200, gin.H{
			"a": "b",
		})
		log.Print("abort after")
		//c.GetInt("key")
	}
}

func main() {
	r := gin.Default()
	r.Use(mw())

	r.GET("ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run()
}
