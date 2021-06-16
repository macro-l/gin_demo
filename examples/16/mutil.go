package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	// 如果使用LoadHTMLFiles 的话这么做(需要列举所有需要加载的文件，不如上述 LoadHTMLGlob 模式匹配方便):
	// router.LoadHTMLFiles("templates/templates.html","templates/templates2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
			"title": "Home",
		})
	})
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users",
		})
	})
	router.Run(":8080")
}
