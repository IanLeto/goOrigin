package backend

import "github.com/jinzhu/gorm"

var GlobalBackend Backend

type Backend interface {
	NewBackend() Backend
	Close() error
}

type DataBase struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

func (d *DataBase) NewBackend() Backend {
	panic("implement me")
}

func (d *DataBase) Close() error {
	panic("implement me")
}
