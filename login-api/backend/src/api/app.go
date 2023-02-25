package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type AppHandler struct {
	http.Handler
}

func (a *AppHandler) getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "say hello")
}

func MakeHandler() *AppHandler {
	mux := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
		negroni.NewRecovery(),
	)
	n.UseHandler(mux)

	a := &AppHandler{
		Handler: n,
	}

	mux.HandleFunc("/hello", a.getHello).Methods("GET")

	return a
}
