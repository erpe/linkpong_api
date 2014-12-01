package model

type Link struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	StoreId uint64 `json:"store_id"`
}

type Store struct {
	Id    uint64   `json:"id"`
	Title string   `json:"title"`
	Uuid  string   `json:"uuid"`
	Links []uint64 `json:"links"`
}
