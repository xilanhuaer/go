package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"primary_key;auto_increment" json:"id"`
	Account   string         `gorm:"type:varchar(24);not null;unique" json:"account"`
	Password  string         `gorm:"type:varchar(256);not null" json:"password"`
	Name      string         `gorm:"type:varchar(24);not null" json:"name"`
	Email     string         `gorm:"type:varchar(255);not null" json:"email"`
	Phone     string         `gorm:"type:varchar(11);not null" json:"phone"`
	Enabled   string         `gorm:"type:varchar(1);not null;default:1" json:"enabled"`
	Creator   string         `gorm:"type:varchar(24)" json:"creator"`
	Updator   string         `gorm:"type:varchar(24)" json:"updator"`
	CreatedAt time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime" json:"deleted_at"`
}

func (User) TableName() string {
	return "user"
}
