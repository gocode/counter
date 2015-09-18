package main

import (
	"net/http"
)

type counter struct {
	count int
	name  string
	desc  string
}

var counters = []counter{}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "static/counter.html")
}

func staticFilesHandler(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/static/", staticFilesHandler)
	http.ListenAndServe(":9090", nil)
}
