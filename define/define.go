package define

import (
	"goOrigin/config"
	"goOrigin/storage"
)

var InitHandler []func() error

var Conf *config.Config

var MongoSession *storage.MongoConn
