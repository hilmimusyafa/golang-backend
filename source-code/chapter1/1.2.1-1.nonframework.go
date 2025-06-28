package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message": "Halo Dunia dari net/http!"}`)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Println("Server net/http berjalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}