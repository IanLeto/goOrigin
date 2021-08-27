package backend

import (
	"github.com/jinzhu/gorm"
	"goOrigin/config"
)

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

type Connection interface {
	NewConn(config config.Config) Connection
	Create() (interface{}, error)
	Update() (interface{}, error)
	Delete() (interface{}, error)
	Close() error
}
