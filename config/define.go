package config

import (
	"encoding/json"
	"fmt"
	"goOrigin/pkg/utils"
)

var Conf *Config
var ConfV2 *V2Config

func InitConf() {
	Conf = NewConfig()
	ConfV2 = NewV2Config()
	v, err := json.MarshalIndent(Conf, "", "  ")
	utils.NoError(err)
	fmt.Println(string(v))
}
