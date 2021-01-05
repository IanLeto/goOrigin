package backend

import "github.com/jinzhu/gorm"

type Client interface {
	Close() error
	NewClient() interface{}
	// 返回该client 类型
	ShowType() string
}

type BaseClient struct {
	ClientType string
}

func (b BaseClient) Close() error {
	panic("implement me")
}

func (b BaseClient) NewClient() interface{} {
	panic("implement me")
}

func (b BaseClient) ShowType() string {
	return b.ClientType
}

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

