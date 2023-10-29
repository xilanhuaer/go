package entity

type SubCollection struct {
	BaseModule
	MainCollectionId   int    `gorm:"type:int;not null" json:"main_collection_id"`
	MainCollectionName string `gorm:"type:varchar(255):not null" json:"main_collection_name"`
	Name               string `gorm:"type:varchar(255):not null" json:"name"`
	Enabled            string `gorm:"type:varchar(1);default: '1'" json:"enabled"`
	Description        string `gorm:"type:varchar(255)" json:"description"`
}

func (SubCollection) TableName() string {
	return "sub_collection"
}
