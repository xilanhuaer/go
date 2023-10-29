package entity

type MainCollection struct {
	Name        string `gorm:"type:varchar(255):not null" json:"name"`
	Enabled     string `gorm:"type:varchar(1);default: '1'" json:"enabled"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	BaseModule
}

func (MainCollection) TableName() string {
	return "main_collection"
}
