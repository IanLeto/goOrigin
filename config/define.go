package config

import (
	"encoding/json"
	"fmt"
	"goOrigin/pkg/utils"
)

var Conf *Config

func InitConf() {
	Conf = NewConfig()
	v, err := json.MarshalIndent(Conf, "", "  ")
	utils.NoError(err)
	fmt.Println(string(v))
}
