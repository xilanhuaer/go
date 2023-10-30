package orm

import (
	"interface/global"
	"reflect"
	"strconv"
)

// model is a ptr
func Create(model interface{}, name string) error {
	reflect.ValueOf(model).Elem().FieldByName("Creator").SetString(name)
	reflect.ValueOf(model).Elem().FieldByName("Updator").SetString(name)
	return global.DB.Create(model).Error
}

// model is a ptr
func Update(model interface{}, name string) error {
	reflect.ValueOf(model).Elem().FieldByName("Updator").SetString(name)
	return global.DB.Model(model).Updates(model).Error
}

// model is a ptr
func Enable(model interface{}, id, name string) error {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	reflect.ValueOf(model).Elem().FieldByName("ID").SetUint(uint64(uid))
	reflect.ValueOf(model).Elem().FieldByName("Updator").SetString(name)
	return global.DB.Model(model).Updates(model).Error
}

// model is a ptr
func Delete(model interface{}, id, name string) error {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	reflect.ValueOf(model).Elem().FieldByName("ID").SetUint(uint64(uid))
	reflect.ValueOf(model).Elem().FieldByName("Updator").SetString(name)
	return global.DB.Delete(model).Error
}
