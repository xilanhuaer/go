package impl

import (
	"fmt"
	"interface/global"
	"interface/model/entity"
	"interface/utils"
	"log"
)

type MainCollectionService struct{}

func (m *MainCollectionService) CreateMainCollection(e entity.MainCollection) error {
	return global.DB.Create(&e).Error
}
func (m *MainCollectionService) List(page, page_size string, params map[string]string) ([]entity.MainCollection, int64, error) {
	var mainCollections []entity.MainCollection
	limit, offset, err := utils.PageUtil(page, page_size)
	if err != nil {
		return mainCollections, 0, err
	}
	query := global.DB.Model(&entity.MainCollection{})
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
		return mainCollections, 0, err
	}
	err = query.Limit(limit).Offset(offset).Order("enabled desc,id desc").Find(&mainCollections).Error
	return mainCollections, count, err

}
func (m *MainCollectionService) Find(id string) (entity.MainCollection, error) {
	var mainCollection entity.MainCollection
	err := global.DB.Where("id=?", id).First(&mainCollection).Error
	return mainCollection, err
}
func (m *MainCollectionService) UpdateMainCollection(id string, mainCollection entity.MainCollection) error {
	log.Println(id, mainCollection)
	return global.DB.Where("id=?", id).Model(&entity.MainCollection{}).Updates(&mainCollection).Error
}
func (m *MainCollectionService) CheckMainCollectionEnable(id string, mainCollection entity.MainCollection) error {
	return global.DB.Where("id=?", id).Model(&entity.MainCollection{}).Update("enabled", mainCollection.Enabled).Error
}
func (m *MainCollectionService) DeleteMainCollection(id string) error {
	return global.DB.Where("id=?", id).Model(&entity.MainCollection{}).Error
}
