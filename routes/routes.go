package routes

import (
	"portfolio/controller"

	"github.com/gin-gonic/gin"
)

func Routes(htmlTemplates []string) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.LoadHTMLFiles(htmlTemplates...)
	router.Static("/public", "./public")

	base := router.Group("/")
	{
		base.GET("/", controller.RenderHome)
		base.GET("/:catchAll", func(ctx *gin.Context) {
			ctx.HTML(200, "404.html", "")
		})
	}

	return router
}
