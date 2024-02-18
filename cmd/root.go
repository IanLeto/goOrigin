package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/config"
	"goOrigin/internal/dao/elastic"
	"goOrigin/internal/dao/mysql"
	"goOrigin/pkg/k8s"
	"goOrigin/pkg/storage"
	"goOrigin/pkg/utils"
	"os"
)

func paramsStr(v string, err error) string {
	utils.NoError(err)
	return v
}

var compInit = map[string]func() error{
	"mongo": storage.InitMongo,
	"zk":    storage.InitZk,
	"k8s":   k8s.InitK8s,
	"redis": storage.InitRedis,
	"mysql": mysql.NewMySQLConns,
	"es":    elastic.InitEs,
}

func migrate() error {
	for _, i := range config.Conf.Components {
		if v, ok := compInit[i]; ok {
			utils.NoError(v())
		}

	}
	return nil
}

var RootCmd = &cobra.Command{
	Use:   "tool", // 这个是命令的名字,跟使用没啥关系
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		if v, err := cmd.Flags().GetBool("pass"); err != nil && v {
			utils.NoError(os.Setenv("PASS", "true"))
		}
		configPath := paramsStr(cmd.Flags().GetString("config"))
		PreRun(configPath)
		if v, err := cmd.Flags().GetString("init"); err != nil {
			utils.NoError(mysql.DBMigrate(v))
			return
		}
		DebugServer()
	},
}

func init() {
	RootCmd.Flags().StringP("config", "c", "", "config")
	RootCmd.Flags().BoolP("pass", "p", false, "pass")
	RootCmd.Flags().Bool("debug", false, "debug")
	RootCmd.Flags().String("init", "", "init db 啥的，要现保证各个依赖项，安装部署成功")

}
