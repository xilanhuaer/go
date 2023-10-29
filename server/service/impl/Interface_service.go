package impl

import (
	"fmt"
	"interface/global"
	"interface/model/entity"
	"interface/utils"
	"time"

	"gorm.io/gorm"
)

type InterfaceService struct {
}

func (m *InterfaceService) Create(i entity.Interface) error {
	return global.DB.Create(&i).Error
}
func (m *InterfaceService) List(page, page_size string, params map[string]string) ([]entity.Interface, int64, error) {
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
func (m *InterfaceService) Find(id string) (entity.Interface, error) {
	var i entity.Interface
	err := global.DB.Where("id=?", id).First(&i).Error
	return i, err
}
func (m *InterfaceService) Update(id string, i entity.Interface) error {
	return global.DB.Model(&entity.Interface{}).Where("id=?", id).Updates(i).Error
}
func (m *InterfaceService) Enable(id, enabled, name string) error {
	return global.DB.Model(&entity.Interface{}).Where("id=?", id).Updates(map[string]interface{}{"enabled": enabled, "updator": name}).Error
}
func (m *InterfaceService) Delete(id, name string) error {
	return global.DB.Model(&entity.Interface{}).Where("id=?", id).Updates(map[string]interface{}{"deleted": gorm.DeletedAt{
		Valid: true,
		Time:  time.Now().Add(8 * time.Hour),
	}, "updator": name}).Error
}
