package server

import (
	"GoWork/models"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server starting...")
	controllerFactory := models.NewControllerFactory("config/ControllerFactoryConfig.json")
	if controllerFactory != nil {
		http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
			var req models.Request
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			controller, exists := controllerFactory.GetController(req.Domain, req.Command)
			if !exists {
				http.Error(w, "Controller not found", http.StatusNotFound)
				return
			}

			resp := controller.HandleRequest(req)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
