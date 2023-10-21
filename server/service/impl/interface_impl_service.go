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
		err                                                       error
	)
	// 查询对应的关联表中存储的名称
	if err = global.DB.Raw("select name from interface where id = ?", ii.InterfaceID).Scan(&interface_name).Error; err != nil {
		return err
	}
	if err = global.DB.Raw("select name from main_collection where id = ?", ii.MainCollectionID).Scan(&main_collection_name).Error; err != nil {
		return err
	}
	if err = global.DB.Raw("select name from sub_collection where id = ?", ii.SubCollectionID).Scan(&sub_collection_name).Error; err != nil {
		return err
	}
	ii.InterfaceName = interface_name
	ii.MainCollectionName = main_collection_name
	ii.SubCollectionName = sub_collection_name
	if err = global.DB.Create(&ii).Error; err != nil {
		return err
	}
	return nil
}

// 查询接口实现列表
// return []entity.InterfaceImpl, int64, error
func (iis *InterfaceImplService) FindInterfaceImplements(limit, offset int, params map[string]string) ([]entity.InterfaceImpl, int64, error) {
	var (
		interface_implements []entity.InterfaceImpl
	)
	query := global.DB.Model(&entity.InterfaceImpl{})
	for key, field := range params {
		if field != "" {
			// 模糊查询
			if key == "name" {
				query = query.Where(key+" like ?", "%"+field+"%")
			} else {
				query = query.Where(key+" = ?", field)
			}
		}
	}
	var total int64
	// 查询总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Limit(limit).Offset(offset).Find(&interface_implements).Error
	return interface_implements, total, err
}

// 根据id查询接口实现
// return entity.InterfaceImpl, error
func (iis *InterfaceImplService) FindInterfaceImplById(id string) (entity.InterfaceImpl, error) {
	var (
		interface_impl entity.InterfaceImpl
	)
	err := global.DB.Where("id = ?", id).First(&interface_impl).Error
	return interface_impl, err
}

// 根据id更新接口实现
//
//	error
func (iis *InterfaceImplService) UpdateInterfaceImplById(id string, ii entity.InterfaceImpl) error {
	var (
		interface_impl entity.InterfaceImpl
	)
	err := global.DB.Model(&interface_impl).Where("id = ?", id).Updates(ii).Error
	return err
}

// 根据id删除接口实现
// return error
func (iis *InterfaceImplService) DeleteInterfaceImplById(id string) error {
	var (
		interface_impl entity.InterfaceImpl
	)
	err := global.DB.Where("id = ?", id).Delete(&interface_impl).Error
	return err
}

// 根据ID切换接口实现的状态
// params id, enabled
// return error
func (iis *InterfaceImplService) SwitchInterfaceImplById(id string, enabled string) error {
	var (
		interface_impl entity.InterfaceImpl
	)
	err := global.DB.Model(&interface_impl).Where("id = ?", id).Update("enabled", enabled).Error
	return err
}