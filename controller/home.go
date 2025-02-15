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
	appointment := map[string]string{
		"img":          "/public/pics/appointment.png",
		"title":        "Appointment",
		"description":  "Időpontok foglalása, kezelése",
		"technologies": "Laravel, Tailwind CSS, Blade, MySQL",
		"link":         "https://appointment.davidfabian.hu/",
		"github":       "https://github.com/Dazannah/appointment",
	}
	// myblog := map[string]string{
	// 	"img":          "/public/pics/myblog.png",
	// 	"title":        "Blog",
	// 	"description":  "Bloggerek követése, blogokra kommentelés",
	// 	"technologies": "Laravel, Tailwind CSS, Blade, MySQL",
	// 	"link":         "https://myblog.davidfabian.hu/",
	// 	"github":       "https://github.com/Dazannah/blog-blade",
	// }
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
		"description":   "",
		"worklocation":  "Szeged, Makó, Hódmezővásárhely",
	}

	infolan := map[string]string{
		"workPlaceName": "Infolan Kft.",
		"date":          "2020 február - jelenleg is",
		"position":      "Rendszergazda",
		"description": `Rendszergazdai feladatok elvégzése melett, CSMEKHM intrawebes jogosultságigénylő alkalmazás fejlesztése, karbantartása.
		NodeJS, MongoDB, és React-ot használva. 
		`,
		"worklocation": "Hódmezővásárhely, Makó",
	}

	iec := map[string]string{
		"schoolName":  "International Education Center",
		"date":        "2023 március - 2024 december",
		"course":      "Junior Fullstack API fejlesztő",
		"description": "Tanultak: PHP, JavaScript, Yii2, MySQL, REST API.",
	}

	jgypk := map[string]string{
		"schoolName":  "Szegedi Tudományegyetem Juhász Gyula Pedagógusképző Kar",
		"date":        "2016 szeptember - 2019 január",
		"course":      "Felsőfokú rendszergazda mérnökinformatikus-asszisztens",
		"description": "",
	}

	gd := map[string]string{
		"schoolName":  "Szegedi SzC Gábor Dénes Szakgimnáziuma és Szakközépiskolája",
		"date":        "2013 szeptember - 2016 június",
		"course":      "Rendszergazda",
		"description": "",
	}

	signum := map[string]string{
		"schoolName":  "Szent István Egyházi Általános Iskola, Gimnázium és Kollégium",
		"date":        "2009 szeptember - 2013 június",
		"course":      "Érettségi",
		"description": "",
	}

	c.HTML(200, "index.html", gin.H{
		"title":      "Fábián Dávid | Portfólió",
		"buttons":    []map[string]string{introductionButton, workplacesButton, studiesButton, projectsButton, contactButton},
		"projects":   []map[string]string{appointment, todoApp, portfolio},
		"workplaces": []map[string]string{infolan, csmpsz},
		"studies":    []map[string]string{iec, jgypk, gd, signum},
	})
}
