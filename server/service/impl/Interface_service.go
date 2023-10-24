package impl

import (
	"fmt"
	"interface/global"
	"interface/model/entity"
	"interface/utils"
)

type InterfaceService struct {
}

func (m *InterfaceService) CreateInterface(i entity.Interface) error {
	return global.DB.Create(&i).Error
}
func (m *InterfaceService) FindInterfaces(page, page_size string, params map[string]string) ([]entity.Interface, int64, error) {
	var i []entity.Interface
	limit, offset, err := utils.PageUtil(page, page_size)
	if err != nil {
		return i, 0, err
	}
	query := global.DB.Model(&entity.Interface{})
	for key, field := range params {
		if params[key] != "" {
			if key != "name" {
				query.Where(fmt.Sprintf("%s=?", key), field)
			} else {
				query.Where("name like ?", "%"+field+"%")
			}
		}
	}
	var count int64
	err = query.Count(&count).Error
	if err != nil {
		return i, 0, err
	}
	err = query.Limit(limit).Offset(offset).Order("enabled desc,id desc").Find(&i).Error
	return i, count, err

}
func (m *InterfaceService) FindInterface(id string) (entity.Interface, error) {
	var i entity.Interface
	err := global.DB.Where("id=?", id).First(&i).Error
	return i, err
}
func (m *InterfaceService) UpdateInterface(id string, i entity.Interface) error {
	return global.DB.Where("id=?", id).Model(&entity.Interface{}).Updates(&i).Error
}
func (m *InterfaceService) CheckInterfaceEnable(id, enabled, name string) error {
	return global.DB.Raw("update interface set enabled = ?, updator=? where id = ?", enabled, name, id).Error
}
func (m *InterfaceService) DeleteInterface(id, name string) error {
	return global.DB.Raw("update interface set deleted_at = now(), updator = ? where id = ?", name, id).Error
}
