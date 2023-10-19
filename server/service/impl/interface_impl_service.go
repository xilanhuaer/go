package impl

import (
	"interface/global"
	"interface/model/entity"
)

type InterfaceImplService struct {
}

func (iis *InterfaceImplService) CreateImpl(ii entity.InterfaceImpl) error {
	var (
		interface_name, main_collection_name, sub_collection_name string
	)
	// 查询对应的关联表中存储的名称
	if err := global.DB.Raw("select name from interface where id = ?", ii.InterfaceID).Scan(&interface_name).Error; err != nil {
		return err
	} else {
		ii.InterfaceName = interface_name
	}
	if err := global.DB.Raw("select name from main_collection where id = ?", ii.MainCollectionID).Scan(&main_collection_name).Error; err != nil {
		return err
	} else {
		ii.MainCollectionName = main_collection_name
	}
	if err := global.DB.Raw("select name from sub_collection where id = ?", ii.SubCollectionID).Scan(&sub_collection_name).Error; err != nil {
		return err
	} else {
		ii.SubCollectionName = sub_collection_name
	}
	err := global.DB.Create(&ii).Error
	if err != nil {
		return err
	}
	return nil
}
func (iis *InterfaceImplService) FindInterfaceImplements(limit, offset int, params map[string]string) ([]entity.InterfaceImpl, error) {
	return nil, nil
}
