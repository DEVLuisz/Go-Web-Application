package main

import (
	"Luís/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":9000", nil)
}
