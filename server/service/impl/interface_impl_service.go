package impl

import (
	"fmt"
	"interface/global"
	"interface/model/entity"
	"log"
)

type InterfaceImplService struct {
}

func (interfaceImplService *InterfaceImplService) Create(ii entity.InterfaceImpl) error {
	var (
		err error
	)
	if err = global.DB.Create(&ii).Error; err != nil {
		return err
	}
	return nil
}

// 查询接口实现列表
// return []entity.InterfaceImpl, int64, error
func (interfaceImplService *InterfaceImplService) List(limit, offset int, params map[string]string) ([]entity.InterfaceImpl, int64, error) {
	var (
		interface_implements []entity.InterfaceImpl
	)
	query := global.DB.Model(&entity.InterfaceImpl{})
	for key, field := range params {
		if params[key] != "" {
			if key == "name" || key == "interface_name" {
				query.Where(fmt.Sprintf("%s like ?", key), "%"+field+"%")
			} else {
				query.Where(fmt.Sprintf("%s = ?", key), field)
			}
		}
	}
	log.Println(query)
	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return interface_implements, 0, err
	}
	err = query.Limit(limit).Offset(offset).Order("enabled desc,id desc").Find(&interface_implements).Error
	return interface_implements, count, err

}

// 根据id查询接口实现
// return entity.InterfaceImpl, error
func (interfaceImplService *InterfaceImplService) Find(id string) (entity.InterfaceImpl, error) {
	var (
		interface_impl entity.InterfaceImpl
	)
	err := global.DB.Where("id = ?", id).First(&interface_impl).Error
	return interface_impl, err
}

// 根据id更新接口实现
// params id, entity.InterfaceImpl
// return error
func (interfaceImplService *InterfaceImplService) Update(id string, ii entity.InterfaceImpl, name string) error {
	var (
		interface_name, main_collection_name, sub_collection_name, path string
		err                                                             error
	)
	// 查询对应的关联表中存储的名称
	if err = global.DB.Raw("select name from interface where id = ?", ii.InterfaceID).Scan(&interface_name).Error; err != nil {
		return err
	}
	if err = global.DB.Raw("select path from interface where id = ?", ii.InterfaceID).Scan(&path).Error; err != nil {
		return err
	}
	if err = global.DB.Raw("select name from main_collection where id = ?", ii.MainCollectionID).Scan(&main_collection_name).Error; err != nil {
		return err
	}
	if err = global.DB.Raw("select name from sub_collection where id = ?", ii.SubCollectionID).Scan(&sub_collection_name).Error; err != nil {
		return err
	}
	{
		ii.InterfaceName = interface_name
		ii.Path = path
		ii.MainCollectionName = main_collection_name
		ii.SubCollectionName = sub_collection_name
	}
	err = global.DB.Where("id = ?", id).Model(&entity.InterfaceImpl{}).Updates(&ii).Update("updator", name).Error
	return err
}

// 根据id删除接口实现
// params id name
// return error
func (interfaceImplService *InterfaceImplService) Delete(id, name string) error {
	err := global.DB.Raw("update interface_impl set deleted_at = now(),updator = ? where id = ?", name, id).Error
	return err
}

// 根据ID切换接口实现的状态
// params id, entity.InterfaceImpl
// return error
func (interfaceImplService *InterfaceImplService) Enable(id string, ii entity.InterfaceImpl) error {
	err := global.DB.Raw("update interface_impl set enabled = ? ,updator = ? where id = ?", ii.Enabled, ii.Updator, id).Error
	return err
}
