package persistence

import (
	"database/sql"
	"github.com/erpe/linkpong_api/model"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type LinkMapper struct {
	Id      uint64 `db:"id"`
	Title   string `db:"title"`
	Url     string `db:"url"`
	StoreId uint64 `db:"store_id"`
}

type StoreMapper struct {
	Id    uint64 `db:"id"`
	Title string `db:"title"`
	Uuid  string `db:"uuid"`
}

func (lm *LinkMapper) ToLink() model.Link {
	return model.Link{lm.Id, lm.Title, lm.Url, lm.StoreId}
}

func (sm *StoreMapper) ToStore() model.Store {
	return model.Store{sm.Id, sm.Title, sm.Uuid}
}

func TestMapper() model.Link {
	lnk := LinkMapper{42, "Ein Map title", "http://www.heise.de", 23}
	return lnk.ToLink()
}

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
