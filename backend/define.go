package backend

var Persistence *Persistence

type Persistence struct {
	Mysql *MySQLBackend
}

func InitPersistence()  {

}

