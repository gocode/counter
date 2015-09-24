package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type counter struct {
	Count int
	Name  string
	Desc  string
}

var counters = []counter{}

func getHandler(rw http.ResponseWriter, req *http.Request) {
	e := json.NewEncoder(rw)
	e.Encode(counters)
}

func addHandler(rw http.ResponseWriter, req *http.Request) {
	c := counter{}
	c.Name = req.FormValue("Name")
	c.Desc = req.FormValue("Desc")
	c.Count, _ = strconv.Atoi(req.FormValue("Count"))
	counters = append(counters, c)
}

func incHandler(rw http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	for _, c := range counters {
		if c.Name == name {
			c.Count++
			break
		}
	}
}

func decHandler(rw http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	for _, c := range counters {
		if c.Name == name {
			c.Count--
			break
		}
	}
}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "static/index.html")
}

func staticFilesHandler(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, req.URL.Path[1:])
}

func main() {

	//test data
	counters = append(counters, counter{Name: "mohan", Desc: "desc"})
	counters = append(counters, counter{Name: "ram", Desc: "csed"})

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/static/", staticFilesHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/inc", incHandler)
	http.HandleFunc("/dec", decHandler)
	http.ListenAndServe(":9090", nil)
}
