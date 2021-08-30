package config

var Conf *Config

func InitConf(path string) {
	Conf = NewConfig(path)
}
func init() {

}
