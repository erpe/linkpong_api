package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/erpe/linkpong_api/cors"
	"github.com/erpe/linkpong_api/model"
	"github.com/erpe/linkpong_api/persistence"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
)

//
//	"github.com/jmoiron/sqlx"
type StoreJSON struct {
	Store model.Store `json:"store"`
}

type StoresJSON struct {
	Stores []model.Store `json:"stores"`
}

type LinkJSON struct {
	Link model.Link `json:"link"`
}
type LinksJSON struct {
	Links []model.Link `json:"links"`
}

var stores []model.Store
var links []model.Link
var link_ids []uint64
var db *sql.DB

func main() {
	// dummy data
	link_ids = append(link_ids, 42, 43)
	db = NewDB()

	log.Println("preparing Router...")

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

	// base_links collection
	base_links := r.Path("/links").Subrouter()
	base_links.Methods("GET").HandlerFunc(LinksIndexHandler)
	base_links.Methods("POST").HandlerFunc(LinksCreateHandler)

	// base links singular
	base_link := r.Path("/links/{id}").Subrouter()
	base_link.Methods("GET").HandlerFunc(LinkShowHandler)
	base_link.Methods("PUT", "POST").HandlerFunc(LinkUpdateHandler)
	base_link.Methods("DELETE").HandlerFunc(LinkDeleteHandler)

	log.Println("server starts listening on 8080...")

	n := negroni.New(negroni.NewLogger(), negroni.HandlerFunc(cors.Intercept))
	n.UseHandler(r)
	n.Run(":8080")
	//http.ListenAndServe(":8080", r)
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

	store1 := model.Store{1, "Golang", "lakjei38fasjifasifhjasdfaqcnv", link_ids}
	store2 := model.Store{2, "Javascript", "asdkfjalsdj3r3r3ljlm3i3r3", link_ids}

	stores = append(stores, store1, store2)

	js, err := json.Marshal(StoresJSON{Stores: stores})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
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

	store := model.Store{storeId, "Golang", "lakjei38fasjifasifhjasdfaqcnv", link_ids}

	js, err := json.Marshal(StoreJSON{Store: store})
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

func LinksIndexHandler(rw http.ResponseWriter, r *http.Request) {

	link1 := model.Link{42, "The linkpong api",
		"http://github.com/erpe/linkpong_api", 1}
	link2 := model.Link{43, "The linkpong app",
		"https://github.com/pixelkritzel/linkpong-ember-client", 2}

	links = append(links, link1, link2)

	js, err := json.Marshal(LinksJSON{Links: links})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)

}

func LinksCreateHandler(rw http.ResponseWriter, r *http.Request) {
	text := "POST /stores/{store_id}/links - create link"
	rw.Write([]byte(text))
}

func LinkShowHandler(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	linkId, err := strconv.ParseUint(vars["id"], 0, 64)

	//if err != nil {
	//		http.Error(rw, err.Error(), http.StatusInternalServerError)
	//		return
	//	}

	//link := Link{linkId, "The linkpong api", "http://github.com/erpe/linkpong_api", 2}
	link := persistence.TestMapper()
	link.Id = linkId
	//Link{linkId, "The linkpong api", "http://github.com/erpe/linkpong_api", 2}

	js, err := json.Marshal(LinkJSON{Link: link})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func LinkUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	text := "POST | PUT /stores/{store_id}/links/{id} - update link"
	rw.Write([]byte(text))
}

func LinkDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	text := "DELETE /stores/{store_id}/links/{id} - delete link"
	rw.Write([]byte(text))
}

func StoreLinksIndexHandler(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	storeId, err := strconv.ParseUint(vars["store_id"], 0, 64)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	link1 := model.Link{42, "The linkpong api",
		"http://github.com/erpe/linkpong_api", storeId}
	link2 := model.Link{43, "The linkpong app",
		"https://github.com/pixelkritzel/linkpong-ember-client", storeId}

	links = append(links, link1, link2)

	js, err := json.Marshal(LinksJSON{Links: links})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
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

	link := model.Link{linkId, "The linkpong api",
		"http://github.com/erpe/linkpong_api", storeId}

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
	db, err := sql.Open("sqlite3", "linkpong.sqlite")
	if err != nil {
		panic(err)
	}

	log.Println("preparing database")
	_, err = db.Exec("create table if not exists links(id INTEGER PRIMARY KEY AUTOINCREMENT, title string, url text, store_id integer)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists stores(id INTEGER PRIMARY KEY AUTOINCREMENT, title string, uuid string)")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`INSERT INTO stores (title,uuid) values ('Hier mein Super Store','xxxxx111111sss')`)

	if err != nil {
		panic(err)
	}
	log.Println("database prepared")
	return db
}
