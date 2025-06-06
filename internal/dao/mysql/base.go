package mysql

import (
	"gorm.io/gorm"
)

type Table interface {
	GetID() uint
}

func Transaction(db *gorm.DB, txFunc func(*gorm.DB) error) error {
	return db.Transaction(txFunc)
}

func Create(db *gorm.DB, value interface{}) (interface{}, *gorm.DB, error) {
	result := db.Create(value)
	return result, result, result.Error
}

func ExecSQL(db *gorm.DB, sql string) (*gorm.DB, error) {
	result := db.Exec(sql)
	return result, result.Error
}

func BatchCreate(db *gorm.DB, values interface{}) (interface{}, *gorm.DB, error) {
	result := db.Create(values)
	return result, result, result.Error
}

func GetAllValues(db *gorm.DB, output []interface{}) (interface{}, *gorm.DB, error) {
	result := db.Find(output).Limit(-1)
	return result, result, result.Error
}

func GetValues(db *gorm.DB, output interface{}, limit int) (interface{}, *gorm.DB, error) {
	result := db.Find(output).Limit(limit)
	return result, result, result.Error
}

func GetValue(db *gorm.DB, output interface{}, tableName string) (interface{}, *gorm.DB, error) {
	result := db.Table(tableName).First(output)
	return result, result, result.Error
}

func GetValueByRaw(db *gorm.DB, output interface{}, tableName string, where string) (interface{}, *gorm.DB, error) {
	result := db.Table(tableName).Where(where).First(output)
	return result, result, result.Error
}

func GetValuesByRaw(db *gorm.DB, output []interface{}, tableName string, where string, limit int) (interface{}, *gorm.DB, error) {
	result := db.Table(tableName).Where(where).Find(output).Limit(limit)
	return result, result, result.Error
}
func GetValueByID(db *gorm.DB, input interface{}) (interface{}, error) {

	result := db.First(&input)
	return result, result.Error
}

func DeleteValue(db *gorm.DB, value interface{}) error {
	result := db.Delete(value)
	return result.Error
}

func DeleteValues(db *gorm.DB, value interface{}) error {
	result := db.Delete(value)
	return result.Error
}

func UpdateValue(db *gorm.DB, tableName string, where string, value interface{}) (interface{}, *gorm.DB, error) {
	result := db.Table(tableName).Where(where).Updates(value)
	return result, result, result.Error
}

func GetValuesByField(db *gorm.DB, fieldName string, fieldValue interface{}, output interface{}) error {
	result := db.Where(fieldName+" = ?", fieldValue).Find(output)
	return result.Error
}
