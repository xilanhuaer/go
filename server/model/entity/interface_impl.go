package entity

import (
	"time"

	"gorm.io/gorm"
)

type InterfaceImpl struct {
	ID                 uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	InterfaceID        uint           `gorm:"not null" json:"interface_id"`
	InterfaceName      string         `gorm:"not null" json:"interface_name"`
	Name               string         `gorm:"type:varchar(255)" json:"name"`
	Path               string         `gorm:"type:varchar(255);not null" json:"path"`
	Type               string         `gorm:"type:varchar(255)" json:"type"`
	Params             *string        `gorm:"type:text" json:"params"`
	Headers            *string        `gorm:"type:text" json:"headers"`
	JsonBody           *string        `gorm:"type:text" json:"json_body"`
	Enabled            string         `gorm:"type:char(1);not null;default:1" json:"enabled"`
	Description        string         `gorm:"type:varchar(255)" json:"description"`
	MainCollectionID   uint           `json:"main_collection_id"`
	MainCollectionName string         `gorm:"type:varchar(255)" json:"main_collection_name"`
	SubCollectionID    uint           `json:"sub_collection_id"`
	SubCollectionName  string         `gorm:"type:varchar(255)" json:"sub_collection_name"`
	Creator            string         `gorm:"type:varchar(255)" json:"creator"`
	Updator            string         `gorm:"type:varchar(255)" json:"updator"`
	CreatedAt          time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"type:datetime" json:"deleted_at"`
}

func (InterfaceImpl) TableName() string {
	return "interface_impl"
}
