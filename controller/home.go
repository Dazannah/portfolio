package controller

import "github.com/gin-gonic/gin"

func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", "")
}
