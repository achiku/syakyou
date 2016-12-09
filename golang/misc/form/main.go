package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func receivePostHandler(t *testing.T) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		buf, _ := ioutil.ReadAll(r.Body)
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		b, _ := ioutil.ReadAll(rdr1)
		r.Body = rdr2
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		t.Logf("%s", b)
		t.Logf("%+v", r.Form)
		t.Logf("%+v", r.PostForm)
		t.Logf("%s", r.FormValue("order[custom]"))
		t.Logf("%s", r.FormValue("orderAamount"))
		return
	}
}

func newMux(t *testing.T) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/post", receivePostHandler(t))
	return mux
}
