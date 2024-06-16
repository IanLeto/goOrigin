package dao

var Conns []Connection

type Connection interface {
	Migrate() error
}
