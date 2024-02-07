package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Index .
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi bardia gopher!")
}

// Health .
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Readiness .
func Readiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	http.HandleFunc("/test", Index)

	http.HandleFunc("/health", Health)
	http.HandleFunc("/readiness", Readiness)

	http.ListenAndServe(os.Getenv("PORT_API"), nil)
}
