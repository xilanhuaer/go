package entity

import (
	"time"

	"gorm.io/gorm"
)

type SubCollection struct {
	Id                 uint           `gorm:"primary_key;auto_increment" json:"id"`
	MainCollectionId   int            `gorm:"type:int;not null" json:"main_collection_id"`
	MainCollectionName string         `gorm:"type:varchar(255):not null" json:"main_collection_name"`
	Name               string         `gorm:"type:varchar(255):not null" json:"name"`
	Enabled            string         `gorm:"type:varchar(1);default: '1'" json:"enabled"`
	Description        string         `gorm:"type:varchar(255)" json:"description"`
	Creator            string         `gorm:"type:varchar(255)" json:"creator"`
	Updator            string         `gorm:"type:varchar(255)" json:"updator"`
	CreatedAt          time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"type:datetime" json:"deleted_at"`
}

func (SubCollection) TableName() string {
	return "sub_collection"
}
