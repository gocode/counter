package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type counter struct {
	count int
	name  string
	desc  string
}

var counters = []counter{}

func getHandler(rw http.ResponseWriter, req *http.Request) {
	e := json.NewEncoder(rw)
	e.Encode(counters)
}

func addHandler(rw http.ResponseWriter, req *http.Request) {
	c := counter{}
	c.name = req.FormValue("name")
	c.desc = req.FormValue("desc")
	c.count, _ = strconv.Atoi(req.FormValue("count"))
	counters = append(counters, c)
}

func incHandler(rw http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	for _, c := range counters {
		if c.name == name {
			c.count++
			break
		}
	}
}

func decHandler(rw http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	for _, c := range counters {
		if c.name == name {
			c.count--
			break
		}
	}
}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "static/counter.html")
}

func staticFilesHandler(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/static/", staticFilesHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/inc", incHandler)
	http.HandleFunc("/dec", decHandler)
	http.ListenAndServe(":9090", nil)
}
