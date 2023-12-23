package routes

import (
	"portfolio/controller"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	files := []string{
		"views/index.html", "views/404.html",
		"views/includes/footer.tmpl", "views/includes/header.tmpl",
	}

	router.LoadHTMLFiles(files...)

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
