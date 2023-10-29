package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseModule struct {
	ID        uint           `gorm:"primary_key;auto_increment" json:"id"`
	Creator   string         `gorm:"type:varchar(255)" json:"creator"`
	Updator   string         `gorm:"type:varchar(255)" json:"updator"`
	CreatedAt time.Time      `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime" json:"deleted_at"`
}

// create hooks
func (base *BaseModule) BeforeCreate(tx *gorm.DB) (err error) {
	base.CreatedAt = time.Now().Add(8 * time.Hour)
	base.UpdatedAt = time.Now().Add(8 * time.Hour)
	return
}

// update hooks
func (base *BaseModule) BeforeUpdate(tx *gorm.DB) (err error) {
	base.UpdatedAt = time.Now().Add(8 * time.Hour)
	return
}

// delete hooks
func (base *BaseModule) BeforeDelete(tx *gorm.DB) (err error) {
	base.DeletedAt = gorm.DeletedAt{
		Time:  time.Now().Add(8 * time.Hour),
		Valid: true,
	}
	return
}
