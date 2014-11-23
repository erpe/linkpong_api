package main

import (
		"bytes"
		"net/http"
	)

func main() {
	http.HandleFunc("/stores/new", StoreNew)
	http.HandleFunc("/stores/:id", StoreShow)
	http.HandleFunc("/stores/:id/links", PostStoreLink)
	http.HandleFunc("/stores/:id/links/latest", LatestStoreLink)
	http.HandleFunc("/", Index)
	http.Handle("/ember", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}


func Index(rw http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	buffer.WriteString("GET /\n")
	buffer.WriteString("GET /stores/new\n")
	buffer.WriteString("GET /stores/:id\n")
	buffer.WriteString("GET /stores/:id/links\n")
	buffer.WriteString("POST /stores/:id/links\n")
	rw.Write([]byte(buffer.String()))
}

func StoreNew(rw http.ResponseWriter, r *http.Request) {
	text := "GET to /stores/new"
	rw.Write([]byte(text))
}

func StoreShow(rw http.ResponseWriter, r *http.Request) {
	text := "GET to /stores/:id/show"
	rw.Write([]byte(text))
}

func PostStoreLink(rw http.ResponseWriter, r *http.Request) {
	text := "POST to /stores/:id/links"
	rw.Write([]byte(text))
}

func LatestStoreLink(rw http.ResponseWriter, r *http.Request) {
	text := "GET to /stores/:id/links"
	rw.Write([]byte(text))
}
