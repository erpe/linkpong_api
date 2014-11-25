package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Link struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	StoreId uint64 `json:"store_id"`
}

type Store struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Uuid  string `json:"uuid"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	// stores collection
	stores := r.Path("/stores").Subrouter()
	stores.Methods("GET").HandlerFunc(StoresIndexHandler)
	stores.Methods("POST").HandlerFunc(StoresCreateHandler)

	// Store singular
	store := r.Path("/stores/{id}").Subrouter()
	store.Methods("GET").HandlerFunc(StoreShowHandler)
	store.Methods("PUT", "POST").HandlerFunc(StoreUpdateHandler)
	store.Methods("DELETE").HandlerFunc(StoreDeleteHandler)

	// links collection
	links := r.Path("/stores/{store_id}/links").Subrouter()
	links.Methods("GET").HandlerFunc(StoreLinksIndexHandler)
	links.Methods("POST").HandlerFunc(StoreLinksCreateHandler)

	// links singular
	link := r.Path("/stores/{store_id}/links/{id}").Subrouter()
	link.Methods("GET").HandlerFunc(StoreLinkShowHandler)
	link.Methods("PUT", "POST").HandlerFunc(StoreLinkUpdateHandler)
	link.Methods("DELETE").HandlerFunc(StoreLinkDeleteHandler)

	http.ListenAndServe(":8080", r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	buffer.WriteString("GET /\n")
	buffer.WriteString("GET|POST /stores/ - index|creates store\n")
	buffer.WriteString("GET|POST|PUT|DELETE /stores/{id} - reads|updates|deletes store\n")
	buffer.WriteString("GET|POST /stores/{id}/links - index|creates link\n")
	buffer.WriteString("GET|POST|PUT|DELETE /stores/{store_id}/links/{id} - reads|updates|deletes link\n")
	rw.Write([]byte(buffer.String()))
}

func StoresIndexHandler(rw http.ResponseWriter, r *http.Request) {
	text := "GET /stores - stores index"
	rw.Write([]byte(text))
}

func StoresCreateHandler(rw http.ResponseWriter, r *http.Request) {
	text := "POST /stores - create store"
	rw.Write([]byte(text))
}

func StoreShowHandler(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	storeId, err := strconv.ParseUint(vars["id"], 0, 64)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	store := Store{storeId, "Golang", "lakjei38fasjifasifhjasdfaqcnv"}

	js, err := json.Marshal(store)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func StoreUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	text := "POST | PUT /stores/{id} - updates store"
	rw.Write([]byte(text))
}

func StoreDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	text := "DELETE /stores/{id} - deletes store"
	rw.Write([]byte(text))
}

func StoreLinksIndexHandler(rw http.ResponseWriter, r *http.Request) {
	text := "GET /stores/{store_id}/links - links index"
	rw.Write([]byte(text))
}

func StoreLinksCreateHandler(rw http.ResponseWriter, r *http.Request) {
	text := "POST /stores/{store_id}/links - create link"
	rw.Write([]byte(text))
}

func StoreLinkShowHandler(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	storeId, err := strconv.ParseUint(vars["store_id"], 0, 64)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	linkId, err := strconv.ParseUint(vars["id"], 0, 64)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	link := Link{linkId, "The linkpong api", "http://github.com/erpe/linkpong_api", storeId}

	js, err := json.Marshal(link)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func StoreLinkUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	text := "POST | PUT /stores/{store_id}/links/{id} - update link"
	rw.Write([]byte(text))
}

func StoreLinkDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	text := "DELETE /stores/{store_id}/links/{id} - delete link"
	rw.Write([]byte(text))
}
