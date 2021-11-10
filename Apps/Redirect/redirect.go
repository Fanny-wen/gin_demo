package Redirect

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RedirectHandler 重定向
func RedirectHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/404")
}
