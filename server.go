package main

import (
	"portfolio/routes"
)

func main() {
	r := routes.Routes()

	r.Run(":4000")
}
