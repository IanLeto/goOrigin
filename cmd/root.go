package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/config"
	"goOrigin/internal/model/dao"
	"goOrigin/pkg/utils"
	"os"
)

func paramsStr(v string, err error) string {
	utils.NoError(err)
	return v
}

var compInit = map[string]func() error{
	//"mongo": storage.InitMongo,
	//"zk":    storage.InitZk,
	//"k8s":   k8s.InitK8s,
	//"redis": storage.InitRedis,
	"mysql": dao.DBMigrate,
	//"es":    elastic.InitEs,
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
		init, err := cmd.Flags().GetBool("init")
		utils.NoError(err)
		if !init {
			DebugServer()
		}

	},
}

func init() {
	RootCmd.Flags().StringP("config", "c", "", "config")
	RootCmd.Flags().BoolP("pass", "p", false, "pass")
	RootCmd.Flags().Bool("debug", false, "debug")
	RootCmd.Flags().Bool("init", false, "init db 啥的，要现保证各个依赖项，安装部署成功")

}
