package persistence

import (
	//"database/sql"
	"github.com/erpe/linkpong_api/model"
	"github.com/jmoiron/sqlx"
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

func CreateStore(store *model.Store, db *sqlx.DB) model.Store {
	log.Println("About to create/persist a Store")

	result, err := db.Exec(`INSERT INTO stores (title, uuid) VALUES( $1, $2)`, store.Title, "1234567890")

	if err != nil {
		log.Println("ERROR: creating store")
		panic(err)
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		log.Println("Result caught error...")
		panic(err)
	}

	return model.Store{uint64(lastId), store.Title, "1234567890"}
}

func CreateLink(link *model.Link, db *sqlx.DB) model.Link {
	log.Println("about to create/persist a Link")

	result, err := db.Exec(`INSERT INTO links (title, url, store_id) VALUES($1, $2, $3)`, link.Title, link.Url, link.StoreId)

	if err != nil {
		log.Println("ERROR creating link...")
		panic(err)
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		log.Println("Result caught error...")
		panic(err)
	}

	return model.Link{uint64(lastId), link.Title, link.Url, link.StoreId}
}

func AllLinks(db *sqlx.DB) []model.Link {
	log.Println("about to find all links")

	mappedLinks := []LinkMapper{}
	links := []model.Link{}
	err := db.Select(&mappedLinks, "SELECT * FROM links")

	if err != nil {
		log.Println("ERROR getting all links..." + err.Error())
		panic(err)
	}

	for _, value := range mappedLinks {
		links = append(links, value.ToLink())
	}

	return links
}

func NewDB() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "linkpong.sqlite")
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
