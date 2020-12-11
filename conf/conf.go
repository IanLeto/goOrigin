package conf

import (
	"github.com/fsnotify/fsnotify"
	logger "github.com/lexkong/log"
	"github.com/spf13/viper"
	"goOrigin/define"
	"log"
	"os"
)

type Config struct {
	Name string
}

func init() {
	define.InitHandler = append(define.InitHandler, InitConfig, InitLog)
	for _, v := range define.InitHandler {
		if err := v(); err != nil {
			panic("init config file %v")
		}
	}
}
func InitConfig() error {

	viper.AddConfigPath("../config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("reading config err %v", err)
		}
		viper.AddConfigPath(dir)
		if err != viper.ReadInConfig() {
			panic(err)
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func InitLog() error {
	logger.InitWithConfig(&logger.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),       // 输出位置，有两个可选项 —— file 和 stdout。选择 file 会将日志记录到 logge
		LoggerLevel:    viper.GetString("log.logger_level"),  // 日志级别，DEBUG、INFO、WARN、ERROR、FATAL
		LoggerFile:     viper.GetString("log.logger_file"),   // 日志文件
		LogFormatText:  viper.GetBool("log.log_format_text"), // 日志的输出格式，JSON 或者 plaintext，true 会输出成非 JSON 格式，false 会输出成 JSON 格式
		RollingPolicy:  viper.GetString("log.rollingPolicy"), // rotate 依据，可选的有 daily 和 size。如果选 daily 则根据天进行转存，如果是 size 则根据大小进行转存
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),  // rotate 转存时间，配 合rollingPolicy: daily 使用
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),  // rotate 转存大小，配合 rollingPolicy: size 使用
		LogBackupCount: viper.GetInt("log.log_backup_count"), // 当日志文件达到转存标准时，log 系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数
	})
	return nil
}
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("config file changed")
	})
}
