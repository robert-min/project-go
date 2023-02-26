package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project-go/login-api/backend/src/lib"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render = render.New()

type AppHandler struct {
	http.Handler
	db lib.DBHandler
}

type User struct {
	Id       string `json: "id"`
	Password string `json: "password"`
}

func (a *AppHandler) addNewUserHandler(w http.ResponseWriter, r *http.Request) {
	vars, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var user User
	err = json.Unmarshal(vars, &user)
	if err != nil {
		panic(err)
	}
	ok := a.db.AddNewUser(user.Id, user.Password)
	rd.JSON(w, http.StatusOK, ok)
}

func (a *AppHandler) getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "say hello")
}

func MakeHandler(filepath string) *AppHandler {
	mux := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
		negroni.NewRecovery(),
	)
	n.UseHandler(mux)

	a := &AppHandler{
		Handler: n,
		db:      lib.NewDBHandler(filepath),
	}

	mux.HandleFunc("/hello", a.getHello).Methods("GET")
	mux.HandleFunc("/user", a.addNewUserHandler).Methods("POST")

	return a
}
