package mysql

import (
	"github.com/jinzhu/gorm"
)

func Transaction(db *gorm.DB, txFunc func(*gorm.DB) error) error {
	return db.Transaction(txFunc)
}

func Create(db *gorm.DB, value interface{}) (interface{}, *gorm.DB, error) {
	result := db.Create(value)
	return result.Value, result, result.Error
}

func GetValueByID(db *gorm.DB, id uint, input interface{}) (interface{}, error) {

	result := db.First(&input, id)
	return result.Value, result.Error
}

func DeleteValue(db *gorm.DB, value interface{}) error {
	result := db.Delete(value)
	return result.Error
}

func UpdateValue(db *gorm.DB, id uint, value interface{}) error {
	result := db.Model(value).Where("id = ?", id).Updates(value)
	return result.Error
}

func GetValuesByField(db *gorm.DB, fieldName string, fieldValue interface{}, output interface{}) error {
	result := db.Where(fieldName+" = ?", fieldValue).Find(output)
	return result.Error
}
