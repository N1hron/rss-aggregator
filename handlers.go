package main

import "net/http"

func handlerReadiness(writer http.ResponseWriter, request *http.Request) {
	respondWithJson(writer, 200, struct{}{})
}

func handlerError(writer http.ResponseWriter, request *http.Request) {
	respondWithError(writer, 400, "Something went wrong")
}
