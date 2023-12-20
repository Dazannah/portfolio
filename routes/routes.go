package routes

import (
	"portfolio/controller"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("views/*")

	/*api := router.Group("/api")
	{
		api.GET("/", controller.LoginController)
		api.POST("/", controller.Login)
	}*/

	base := router.Group("/")
	{
		base.GET("/", controller.RenderHome)
		base.GET("/:catchAll", func(ctx *gin.Context) {
			ctx.HTML(200, "404.html", "")
		})
	}

	return router
}
