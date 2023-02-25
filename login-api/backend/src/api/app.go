package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "say hello")
}

func MakeHandler() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/hello", getHello).Methods("GET")

	return mux
}
