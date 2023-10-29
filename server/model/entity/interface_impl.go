package entity

import (
	"gorm.io/gorm"
)

type InterfaceImpl struct {
	BaseModule
	InterfaceID        uint    `gorm:"not null" json:"interface_id"`
	InterfaceName      string  `gorm:"not null" json:"interface_name"`
	Name               string  `gorm:"type:varchar(255)" json:"name"`
	Path               string  `gorm:"type:varchar(255);not null" json:"path"`
	Type               string  `gorm:"type:varchar(255)" json:"type"`
	Params             *string `gorm:"type:text" json:"params"`
	Headers            *string `gorm:"type:text" json:"headers"`
	JsonBody           *string `gorm:"type:text" json:"json_body"`
	Enabled            string  `gorm:"type:char(1);not null;default:1" json:"enabled"`
	Description        string  `gorm:"type:varchar(255)" json:"description"`
	MainCollectionID   uint    `json:"main_collection_id"`
	MainCollectionName string  `gorm:"type:varchar(255)" json:"main_collection_name"`
	SubCollectionID    uint    `json:"sub_collection_id"`
	SubCollectionName  string  `gorm:"type:varchar(255)" json:"sub_collection_name"`
}

func (InterfaceImpl) TableName() string {
	return "interface_impl"
}

// 创建接口实现钩子，设置main_collection_name,sub_collection_name,interface_name,path
func (interfaceImpl *InterfaceImpl) BeforeCreate(tx *gorm.DB) (err error) {
	var (
		interface_name, main_collection_name, sub_collection_name, path string
	)
	// 查询对应的关联表中存储的名称
	if err = tx.Raw("select name from interface where id = ?", interfaceImpl.InterfaceID).Scan(&interface_name).Error; err != nil {
		return err
	}
	if err = tx.Raw("select name from main_collection where id = ?", interfaceImpl.MainCollectionID).Scan(&main_collection_name).Error; err != nil {
		return err
	}
	if err = tx.Raw("select name from sub_collection where id = ?", interfaceImpl.SubCollectionID).Scan(&sub_collection_name).Error; err != nil {
		return err
	}
	if err = tx.Raw("select path from interface where id = ?", interfaceImpl.InterfaceID).Scan(&path).Error; err != nil {
		return err
	}
	{
		interfaceImpl.InterfaceName = interface_name
		interfaceImpl.MainCollectionName = main_collection_name
		interfaceImpl.SubCollectionName = sub_collection_name
		interfaceImpl.Path = path
	}
	return nil
}
