package main

import (
	"log"
	"os"
	"portfolio/routes"
)

func main() {
	htmlTemplates := []string{}

	readViewsFolderContent("views/", &htmlTemplates)

	r := routes.Routes(htmlTemplates)

	r.Run(":4000")
}

func readViewsFolderContent(path string, htmlTemplates *[]string) {

	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() {
			readViewsFolderContent(path+file.Name()+"/", htmlTemplates)
		} else {
			//fmt.Println(path + file.Name())
			*htmlTemplates = append(*htmlTemplates, path+file.Name())
		}
	}
}
