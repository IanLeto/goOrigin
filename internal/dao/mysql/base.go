package mysql

import (
	"github.com/jinzhu/gorm"
	"goOrigin/config"
)

type Meta struct {
	ID         uint  `swaggerignore:"true" gorm:"primary_key" json:"id" binding:"-" `
	CreateTime int64 `swaggerignore:"true" gorm:"autoCreateTime;" json:"created_time" binding:"-"`
	ModifyTime int64 `swaggerignore:"true" gorm:"autoUpdateTime;" json:"modify_time" binding:"-"`
}

type DBOpt interface {
	Create() (DBOpt, error)
	Update() (DBOpt, error)
	Delete() (DBOpt, error)
	List() (DBOpt, error)
	BatchCreate([]DBOpt) error
	BatchUpdate([]DBOpt)
}

func Transaction(db *gorm.DB, txFunc func(*gorm.DB) error) error {
	return db.Transaction(txFunc)
}

func Create(db *gorm.DB, value interface{}) (interface{}, *gorm.DB, error) {
	result := db.Create(value)
	return result.Value, result, result.Error
}

func BatchCreate(db *gorm.DB, values interface{}) (interface{}, *gorm.DB, error) {
	result := db.Create(values)
	return result.Value, result, result.Error
}

func CreateBySql(db *gorm.DB, value interface{}) {

}

func GetAllValues(db *gorm.DB, output []interface{}) (interface{}, *gorm.DB, error) {
	result := db.Find(output).Limit(-1)
	return result.Value, result, result.Error
}

func GetValues(db *gorm.DB, output interface{}, limit int) (interface{}, *gorm.DB, error) {
	result := db.Find(output).Limit(limit)
	return result.Value, result, result.Error
}

func GetValue(db *gorm.DB, output interface{}, tableName string) (interface{}, *gorm.DB, error) {
	result := db.Table(tableName).First(output)
	return result.Value, result, result.Error
}

func GetValueByRaw(db *gorm.DB, output interface{}, tableName string, where string) (interface{}, *gorm.DB, error) {
	result := db.Table(tableName).Where(where).First(output)
	return result.Value, result, result.Error
}

func GetValuesByRaw(db *gorm.DB, output []interface{}, tableName string, where string, limit int) (interface{}, *gorm.DB, error) {
	result := db.Table(tableName).Where(where).Find(output).Limit(limit)
	return result.Value, result, result.Error
}
func GetValueByID(db *gorm.DB, input interface{}) (interface{}, error) {

	result := db.First(&input)
	return result.Value, result.Error
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
	return result.Value, result, result.Error
}

func GetValuesByField(db *gorm.DB, fieldName string, fieldValue interface{}, output interface{}) error {
	result := db.Where(fieldName+" = ?", fieldValue).Find(output)
	return result.Error
}

var migrate = map[string]interface{}{
	"t_record": &TRecord{},
}

func DBMigrate(region string) error {
	for _, table := range migrate {
		err := NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client.AutoMigrate(table).Error
		if err != nil {
			return err
		}
	}
	return nil
}
