package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeFrontend(router *gin.Engine) {
	router.LoadHTMLGlob("internal/web/templates/*.html")
	router.Static("/static", "./web/static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
