package controller

import "github.com/gin-gonic/gin"

func RenderHome(c *gin.Context) {

	project0 := gin.H{
		"img":          "/public/pics/portfolio.png",
		"title":        "Portfolió oldal",
		"description":  "Statikus weboldal.",
		"technologies": "Golang, Tailwind CSS",
		"link":         "https://www.davidfabian.hu",
		"github":       "https://github.com/Dazannah/portfolio",
	}

	project1 := gin.H{
		"img":          "/public/pics/ToDo.png",
		"title":        "ToDo app",
		"description":  "Egy egyszerű CRUD aplikáció.",
		"technologies": "Node.js, .ejs, htmx, Tailwind CSS, MongoDB",
		"link":         "https://www.todoapp.davidfabian.hu",
		"github":       "https://github.com/Dazannah/todo-nodejs",
	}

	c.HTML(200, "index.html", gin.H{
		"title": "This is the title",
		"projects": gin.H{
			"project1": project1,
			"project0": project0,
		},
	})
}
