package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
type Todos []Todo

func IndexHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(rw, "index")
}

func TodoIndexHandler(rw http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "buy milk", Completed: false},
		Todo{Name: "pay rent", Completed: false},
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(todos)
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/todo/", TodoIndexHandler)

	n := negroni.New(
		negroni.NewRecovery(),
	)
	n.UseHandler(router)
	n.Run(":8080")
}
