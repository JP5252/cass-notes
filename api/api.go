package api

import (
	"log"
	"net/http"
)

func StartServer() {
	routes := getRoutes()

	for _, route := range routes {
		http.HandleFunc(route.path, route.function)
	}

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
