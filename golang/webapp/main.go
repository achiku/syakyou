package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
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
	http.ListenAndServe(":8080", nil)
}
