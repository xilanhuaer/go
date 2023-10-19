package impl

import (
	"interface/global"
	"interface/model/entity"
)

type MainSubService struct {
}

func (mss *MainSubService) MainSubList() ([]entity.MainSubEntity, error) {
	var mainList []entity.MainCollection
	if err := global.DB.Where("enabled=?", "1").Find(&mainList).Error; err != nil {
		return []entity.MainSubEntity{}, err
	}
	var subList []entity.SubCollection
	if err := global.DB.Where("enabled=?", "1").Find(&subList).Error; err != nil {
		return []entity.MainSubEntity{}, err
	}
	var data []entity.MainSubEntity
	for _, main := range mainList {
		var subCollectionList []entity.SubCollectionData
		for _, sub := range subList {
			if main.ID == uint(sub.MainCollectionId) {
				subCollectionList = append(subCollectionList, entity.SubCollectionData{
					SubCollectionId:   int(sub.Id),
					SubCollectionName: sub.Name,
				})
			}
		}
		data = append(data, entity.MainSubEntity{
			MainCollectionId:   int(main.ID),
			MainCollectionName: main.Name,
			SubCollectionList:  subCollectionList,
		})
	}
	return data, nil
}
