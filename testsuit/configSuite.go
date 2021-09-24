package testsuit

import "goOrigin/config"

func InitTestConfig(conf config.Config) {
	config.Conf = &conf
}
