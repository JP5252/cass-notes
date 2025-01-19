package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Route struct {
	path     string
	function func(w http.ResponseWriter, r *http.Request)
}

func getRoutes() []Route {
	routes := []Route{
		{path: "/health", function: healthcheckHandler},
		{path: "/login", function: loginHandler},
	}

	return routes
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Health check request received from: %s", r.RemoteAddr)
	response := HealthCheckResponse{Status: "OK"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

}
