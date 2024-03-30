package controller

import "github.com/gin-gonic/gin"

func RenderHome(c *gin.Context) {

	// menu buttons
	introductionButton := map[string]string{
		"href": "#introduction",
		"name": "Bemutatkozás",
	}
	workplacesButton := map[string]string{
		"href": "#workplaces",
		"name": "Munkahelyek",
	}
	studiesButton := map[string]string{
		"href": "#studies",
		"name": "Tanulmányok",
	}
	projectsButton := map[string]string{
		"href": "#projects",
		"name": "Projektek",
	}
	contactButton := map[string]string{
		"href": "#contact",
		"name": "Elérhetőségek",
	}

	//apps
	portfolio := map[string]string{
		"img":          "/public/pics/portfolio.png",
		"title":        "Portfolió oldal",
		"description":  "Statikus weboldal.",
		"technologies": "Golang, Tailwind CSS",
		"link":         "https://www.davidfabian.hu",
		"github":       "https://github.com/Dazannah/portfolio",
	}
	todoApp := map[string]string{
		"img":          "/public/pics/ToDo.png",
		"title":        "ToDo app",
		"description":  "Egy egyszerű CRUD aplikáció.",
		"technologies": "Node.js, .ejs, htmx, Tailwind CSS, MongoDB",
		"link":         "https://www.todoapp.davidfabian.hu",
		"github":       "https://github.com/Dazannah/todo-nodejs",
	}

	//workplaces
	csmpsz := map[string]string{
		"workPlaceName": "Csongrád-Csanád Vármegyei Pedagógiai Szakszolgálat",
		"date":          "2019 január - 2020 január",
		"position":      "Rendszergazda",
		"description":   "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		"worklocation":  "Szeged, Makó, Hódmezővásárhely",
	}

	infolan := map[string]string{
		"workPlaceName": "Infolan Kft.",
		"date":          "2020 február -",
		"position":      "Rendszergazda",
		"description":   "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		"worklocation":  "Hódmezővásárhely, Makó",
	}

	c.HTML(200, "index.html", gin.H{
		"title":      "Fábián Dávid | Portfólió",
		"buttons":    []map[string]string{introductionButton, workplacesButton, studiesButton, projectsButton, contactButton},
		"projects":   []map[string]string{todoApp, portfolio},
		"workplaces": []map[string]string{infolan, csmpsz},
		"studies":    []map[string]string{},
	})
}
