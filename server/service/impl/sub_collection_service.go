package impl

import (
	"fmt"
	"interface/global"
	"interface/model/entity"
	"interface/utils"
)

type SubCollectionService struct {
}

func (s *SubCollectionService) CreateSubCollection(e entity.SubCollection) error {
	var m entity.MainCollection
	if err := global.DB.Where("id = ?", e.MainCollectionId).First(&m).Error; err != nil {
		return err
	}
	e.MainCollectionName = m.Name
	return global.DB.Create(&e).Error
}

func (s *SubCollectionService) FindSubCollections(page, page_size string, params map[string]string) ([]entity.SubCollection, int64, error) {
	var e []entity.SubCollection
	limit, offset, err := utils.PageUtil(page, page_size)
	if err != nil {
		return e, 0, err
	}
	query := global.DB.Model(&entity.SubCollection{})
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
		return e, 0, err
	}
	err = query.Limit(limit).Offset(offset).Order("enabled desc,id desc").Find(&e).Error
	return e, count, err

}
func (s *SubCollectionService) FindSubCollection(id string) (entity.SubCollection, error) {
	var e entity.SubCollection
	err := global.DB.Where("id=?", id).First(&e).Error
	return e, err
}
func (s *SubCollectionService) UpdateSubCollection(id string, e entity.SubCollection) error {
	var m entity.MainCollection
	if err := global.DB.Where("id = ?", e.MainCollectionId).First(&m).Error; err != nil {
		return err
	}
	e.MainCollectionName = m.Name
	return global.DB.Where("id = ?", id).Model(&entity.SubCollection{}).Updates(&e).Error
}
func (s *SubCollectionService) CheckSubCollectionEnable(id string, e entity.SubCollection) error {
	return global.DB.Where("id = ?", id).Model(&entity.SubCollection{}).Update("enabled", e.Enabled).Error
}
func (s *SubCollectionService) DeleteSubCollection(id string) error {
	return global.DB.Where("id = ?", id).Model(&entity.SubCollection{}).Error
}
