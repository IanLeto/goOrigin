package config

import (
	"fmt"
	"goOrigin/pkg/utils"
)

var ConfV2 *V2Config

func InitConf() {
	ConfV2 = NewV2Config()
	fmt.Println(utils.ToJson(ConfV2))
}
