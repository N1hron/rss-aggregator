package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(writer http.ResponseWriter, status int, message string) {
	if status > 499 {
		log.Println("Responding with 5XX status:", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(writer, status, errorResponse{Error: message})
}

func respondWithJson(writer http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v\n", payload)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(data)
}
