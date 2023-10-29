package entity

import (
	"time"

	"github.com/gin-gonic/gin"
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
func (base *BaseModule) BeforeCreate(tx *gorm.DB, c *gin.Context) (err error) {
	base.CreatedAt = time.Now().Add(8 * time.Hour)
	base.UpdatedAt = time.Now().Add(8 * time.Hour)
	base.Creator = c.MustGet("username").(string)
	base.Updator = c.MustGet("username").(string)
	return
}

// update hooks
func (base *BaseModule) BeforeUpdate(tx *gorm.DB, c *gin.Context) (err error) {
	base.UpdatedAt = time.Now().Add(8 * time.Hour)
	base.Updator = c.MustGet("username").(string)
	return
}

// delete hooks
func (base *BaseModule) BeforeDelete(tx *gorm.DB, c *gin.Context) (err error) {
	base.DeletedAt = gorm.DeletedAt{
		Time:  time.Now().Add(8 * time.Hour),
		Valid: true,
	}
	base.Updator = c.MustGet("username").(string)
	return
}
