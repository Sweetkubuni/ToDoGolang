package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type App struct {
	*mux.Router
	Todos []string
}

func Initialize() *App {
	app := &App{mux.NewRouter(), make([]string, 0)}
	app.HandleFunc("/todo", app.TodoPost).Methods("POST")
	app.HandleFunc("/todo", app.TodoGet).Methods("GET")
	return app
}

func (a *App) TodoPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	a.Todos = append(a.Todos, string(body))
	fmt.Fprintf(w, "OK")
}

func (a *App) TodoGet(w http.ResponseWriter, r *http.Request) {
	b := new(strings.Builder)
	json.NewEncoder(b).Encode(a.Todos)
	fmt.Fprintf(w, b.String())
}
