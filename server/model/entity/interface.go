package entity

type Interface struct {
	BaseModule
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Type        string `gorm:"type:varchar(255);not null" json:"type"`
	Path        string `gorm:"type:varchar(255);not null" json:"path"`
	Header      string `gorm:"type:varchar(255);default:'{\"Content-Type\":\"application/json\",\"Authorization\":\"${token}\"}'" json:"header"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Enabled     string `gorm:"type:char(1);not null;default:'1'" json:"enabled"`
}

func (Interface) TableName() string {
	return "interface"
}
