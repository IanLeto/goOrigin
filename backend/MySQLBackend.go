package backend

import (
	"goOrigin/conf"
)

type MySQLBackend struct {
	Client
}

func NewMySQLBackend(conf conf.Configuration) error {
	//config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
	//	username,
	//	password,
	//	addr,
	//	name,
	//	true,
	//	//"Asia/Shanghai"),
	//	"Local")
	return nil
}
