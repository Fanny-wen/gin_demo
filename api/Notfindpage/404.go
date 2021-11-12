package Notfindpage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PageNotFindHandler  404页面
func PageNotFindHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"title": "404",
	})
}


