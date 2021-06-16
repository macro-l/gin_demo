package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	// 如果使用LoadHTMLFiles 的话这么做(需要列举所有需要加载的文件，不如上述 LoadHTMLGlob 模式匹配方便):
	// router.LoadHTMLFiles("templates/templates.html","templates/templates2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
