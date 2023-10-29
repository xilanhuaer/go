package entity

type MainSubEntity struct {
	MainCollectionId   int                 `json:"main_collection_id"`
	MainCollectionName string              `json:"main_collection_name"`
	SubCollectionList  []SubCollectionData `json:"sub_collection_list"`
}
type SubCollectionData struct {
	SubCollectionId   int    `json:"sub_collection_id"`
	SubCollectionName string `json:"sub_collection_name"`
}
