package entity

import (
	"time"

	"gorm.io/gorm"
)

type Interface struct {
	ID          uint           `gorm:"primary_key;auto_increment" json:"id"`
	Creator     string         `gorm:"type:varchar(255)" json:"creator"`
	Updator     string         `gorm:"type:varchar(255)" json:"updator"`
	CreatedAt   time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"type:datetime" json:"deleted_at"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Type        string         `gorm:"type:varchar(255);not null" json:"type"`
	Path        string         `gorm:"type:varchar(255);not null" json:"path"`
	Header      string         `gorm:"type:varchar(255);default:'{\"Content-Type\":\"application/json\",\"Authorization\":\"${token}\"}'" json:"header"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	Enabled     string         `gorm:"type:char(1);not null;default:'1'" json:"enabled"`
}

func (Interface) TableName() string {
	return "interface"
}
