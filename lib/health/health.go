package health

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/unravellingtechnologies/unravelling-cloud/cli/lib/server"
	"net/http"
)

type ReadinessResponse struct {
	Ready   bool   `json:"ready"`
	Message string `json:"message"`
}

type HealthinessResponse struct {
	Healthy bool   `json:"healthy"`
	Message string `json:"message"`
}

var healthRouter *mux.Router

// Initialize the sub-router for the health endpoint.
func Initialize(r *mux.Router) {
	log.Infoln("Initializing health endpoint...")
	healthRouter = mux.NewRouter()
	healthRouter.HandleFunc("/readyz", readiness)
	r.HandleFunc("/health", healthiness)
	server.Mount(r, "/health/", healthRouter)
}

// ready Returns the readiness of our application.
func readiness(w http.ResponseWriter, _ *http.Request) {
	log.Debugln("Calculating readiness...")

	r := ReadinessResponse{
		Ready:   true,
		Message: "OK",
	}
	response, _ := json.Marshal(r)
	_, err := fmt.Fprintf(w, string(response))
	if err != nil {
		fmt.Println("Error")
		return
	}
}

// ready Returns the readiness of our application.
func healthiness(w http.ResponseWriter, _ *http.Request) {
	log.Debugln("Calculating health...")

	response, _ := json.Marshal(HealthinessResponse{
		Healthy: true,
		Message: "OK",
	})

	_, err := fmt.Fprintf(w, string(response))
	if err != nil {
		fmt.Println("Error")
		return
	}
}
