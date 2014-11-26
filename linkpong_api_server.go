package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
	Links []Link `json:"links"`
}

type StoreJSON struct {
	Store Store `json:"store"`
}

type StoresJSON struct {
	Stores []Store `json:"stores"`
}

type LinkJSON struct {
	Link Link `json:"link"`
}
type LinksJSON struct {
	Links []Link `json:"links"`
}

var stores []Store
var links []Link

func main() {
	//db := NewDB()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	// stores collection
	stores := r.Path("/stores").Subrouter()
	stores.Methods("GET").HandlerFunc(StoresIndexHandler)
	stores.Methods("POST").HandlerFunc(StoresCreateHandler)

	// Store singular
	store := r.Path("/stores/{id}").Subrouter()
	store.Methods("GET", "OPTIONS").HandlerFunc(StoreShowHandler)
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
	//http.Handle("/", &CorsServer{r})
	//http.ListenAndServe(":8080", db)

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

	if origin := r.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	store1 := Store{1, "Golang", "lakjei38fasjifasifhjasdfaqcnv", links}
	store2 := Store{2, "Javascript", "asdkfjalsdj3r3r3ljlm3i3r3", links}

	stores = append(stores, store1, store2)

	js, err := json.Marshal(StoresJSON{Stores: stores})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if origin := r.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
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

	link1 := Link{42, "The linkpong api", "http://github.com/erpe/linkpong_api", storeId}
	link2 := Link{43, "The linkpong app", "https://github.com/pixelkritzel/linkpong-ember-client", storeId}

	links = append(links, link1, link2)
	store := Store{storeId, "Golang", "lakjei38fasjifasifhjasdfaqcnv", links}

	js, err := json.Marshal(StoreJSON{Store: store})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if origin := r.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
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

	if origin := r.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	storeId, err := strconv.ParseUint(vars["store_id"], 0, 64)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	link1 := Link{42, "The linkpong api", "http://github.com/erpe/linkpong_api", storeId}
	link2 := Link{43, "The linkpong app", "https://github.com/pixelkritzel/linkpong-ember-client", storeId}

	links = append(links, link1, link2)

	js, err := json.Marshal(LinksJSON{Links: links})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if origin := r.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)

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

	js, err := json.Marshal(LinkJSON{Link: link})
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

// database

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "example.sqlite")
	if err != nil {
		panic(err)
	}

	log.Println("preparing database")
	_, err = db.Exec("create table if not exists links(id integer, title string, url text, store_id integer)")
	_, err = db.Exec("create table if not exists stores(id integer, title string, uuid string)")
	if err != nil {
		panic(err)
	}
	return db
}

// TODO: get rid of this...

type CorsServer struct {
	r *mux.Router
}

func (s *CorsServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}
