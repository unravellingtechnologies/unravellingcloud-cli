package agent

import (
	"fmt"
	"github.com/unravellingtechnologies/unravelling-cloud/cli/lib/health"
	"net/http"
)

func start_agent() {
	// Creates the router
	// TODO: expose endpoints for health and readiness

	router.Handle("/healthz/ready", health.Ready)

	// Use the HandleFunc method on the router to register a new route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Use the health.ListenAndServe method to start the HTTP server
	// This will listen on port 8080 on all available interfaces (0.0.0.0)

}
