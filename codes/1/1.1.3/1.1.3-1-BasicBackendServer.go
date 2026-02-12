package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseData is a template (struct) for the form of data to be sent.
type ResponseData struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {
	// 1. Create an endpoint (Menu List at the Restaurant)
	http.HandleFunc("/api/greetings", func(w http.ResponseWriter, r *http.Request) {
		// Set header so the client knows the incoming package contains JSON
		w.Header().Set("Content-Type", "application/json")

		// 2. Prepare the data to be sent (Cooking the order)
		data := ResponseData{
			Message: "Hello! Order successfully processed by Backend.",
			Status:  200, // 200 is the standard internet code for "OK/Success"
		}

		// 3. Convert Go data structure to JSON and send it (Serving the food)
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// 4. Turn on the restaurant machine
	log.Println("Backend Server starts running on port 8080...")
	
	// Server will keep running and listening for requests on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}