package persistence

import (
	"github.com/erpe/linkpong_api/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/twinj/uuid"
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
	arr := make([]uint64, 0)
	return model.Store{sm.Id, sm.Title, sm.Uuid, arr}
}

func TestMapper() model.Link {
	lnk := LinkMapper{42, "Ein Map title", "http://www.heise.de", 23}
	return lnk.ToLink()
}

func CreateStore(store *model.Store, db *sqlx.DB) model.Store {
	log.Println("About to create/persist a Store")

	supplyUuid(store)

	result, err := db.Exec(`INSERT INTO stores (title, uuid) VALUES( $1, $2)`, store.Title, store.Uuid)

	if err != nil {
		log.Println("ERROR: creating store: " + err.Error())
		panic(err)
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		log.Println("Result caught error..." + err.Error())
		panic(err)
	}

	return model.Store{uint64(lastId), store.Title, store.Uuid, make([]uint64, 0)}
}

func CreateLink(link *model.Link, db *sqlx.DB) model.Link {
	log.Println("about to create/persist a Link")

	result, err := db.Exec(`INSERT INTO links (title, url, store_id) VALUES($1, $2, $3)`, link.Title, link.Url, link.StoreId)

	if err != nil {
		log.Println("ERROR creating link..." + err.Error())
		panic(err)
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		log.Println("Result caught error..." + err.Error())
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

func AllStores(db *sqlx.DB) []model.Store {
	log.Println("about to find all stores")

	mappedStores := []StoreMapper{}
	stores := []model.Store{}

	err := db.Select(&mappedStores, "SELECT * FROM stores")

	if err != nil {
		log.Println("ERROR getting all stores..." + err.Error())
		panic(err)
	}

	for _, value := range mappedStores {

		store := value.ToStore()

		links := FindLinksForStore(store.Id, db)
		linkIds := make([]uint64, 0)

		for _, link := range links {
			linkIds = append(linkIds, link.Id)
		}

		store.Links = linkIds

		stores = append(stores, store)
	}
	return stores
}

func FindLinksForStore(storeId uint64, db *sqlx.DB) []model.Link {
	mappedLinks := []LinkMapper{}
	links := []model.Link{}

	err := db.Select(&mappedLinks, "SELECT * FROM links WHERE store_id = $1", storeId)

	if err != nil {
		log.Println("ERROR getting links for Store: " + string(storeId) + " " + err.Error())
		panic(err)
	}

	for _, value := range mappedLinks {
		links = append(links, value.ToLink())
	}

	return links
}

func FindStore(storeId uint64, db *sqlx.DB) model.Store {
	mappedStore := StoreMapper{}
	store := model.Store{}
	err := db.Get(&mappedStore, "SELECT * FROM stores WHERE id = $1", storeId)

	if err != nil {
		log.Println("ERROR getting Store: " + string(storeId) + " " + err.Error())
		panic(err)
	}

	store = mappedStore.ToStore()

	links := FindLinksForStore(storeId, db)
	linkIds := make([]uint64, 0)

	for _, val := range links {
		linkIds = append(linkIds, val.Id)
	}
	store.Links = linkIds

	return store
}

func FindLink(linkId uint64, db *sqlx.DB) model.Link {
	mappedLink := LinkMapper{}
	link := model.Link{}

	err := db.Get(&mappedLink, "SELECT * FROM links WHERE id = $1", linkId)

	if err != nil {
		log.Println("ERROR getting Link: " + string(linkId) + " " + err.Error())
	}

	link = mappedLink.ToLink()
	return link
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

func supplyUuid(store *model.Store) {
	uuid.SwitchFormat(uuid.CleanHyphen)
	id := uuid.NewV4()
	store.Uuid = id.String()
}
