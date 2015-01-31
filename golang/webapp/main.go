package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
	"log"
	"net/http"
)

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("public"))
}

func PostIndexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post index")
}

func PostCreateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post create")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(rw, "post edit:", id)
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(rw, "post show:", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(rw, "post update:", id)
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(rw, "post delete:", id)
}

func MyMiddleWare(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there...")
	if r.URL.Query().Get("password") == "pass" {
		log.Println("authorized...")
		next(rw, r)
	} else {
		log.Println("failed to authorize request...")
		http.Error(rw, "Not Authorized", 401)
	}
	log.Println("Logging on the way back...")
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/markdown", GenerateMarkdown)

	// Posts collections
	posts := r.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(PostIndexHandler)
	posts.Methods("POST").HandlerFunc(PostCreateHandler)

	// Posts singular
	post := r.PathPrefix("/posts/{id}").Subrouter()
	post.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler)
	post.Methods("GET").HandlerFunc(PostShowHandler)
	post.Methods("PUT", "POST").HandlerFunc(PostUpdateHandler)
	post.Methods("DELETE").HandlerFunc(PostUpdateHandler)

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleWare),
		negroni.NewStatic(http.Dir("public")),
	)
	n.Run(":8080")
}
