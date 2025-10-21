package main

import (
	"net/http"

	"github.com/alexsilvabello/loja-golang/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8089", nil)
}
