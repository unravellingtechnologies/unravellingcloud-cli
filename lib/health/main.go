package health

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Ready Returns the readiness of our application.
func readiness(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calculating readiness")

	response, _ := json.Marshal(Readiness{
		ready:   true,
		message: "Ready to rock & roll...",
	})
	_, err := fmt.Fprintf(w, string(response))
	if err != nil {
		fmt.Println("Error")
		return
	}
}
