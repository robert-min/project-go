package main

import (
	"net/http"
	"project-go/login-api/backend/src/api"
)

func main() {
	router := api.MakeHandler()
	http.ListenAndServe(":8080", router)
}
