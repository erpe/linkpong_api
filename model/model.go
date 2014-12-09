package model

type Link struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	//StoreString string `json:"store_id"`
	StoreId uint64 `json:"store_id, string"`
}

type Store struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Uuid  string `json:"uuid"`
}

//func (link *Link) GetStoreId() Link {
//	store_id := uint64(link.StoreString)
//	return Link{link.Id, link.Title, link.Url, link.StoreString, store_id}
//}
