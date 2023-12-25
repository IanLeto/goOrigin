package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/pkg/utils"
	"os"
)

func paramsStr(v string, err error) string {
	utils.NoError(err)
	return v
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
		//v, err := cmd.Flags().GetString("init")
		//utils.NoError(err)
		//if v != "" {
		//	utils.NoError(mysql.DBMigrate(v))
		//	return
		//}
		DebugServer()
	},
}

func init() {
	RootCmd.Flags().StringP("config", "c", "", "config")
	RootCmd.Flags().BoolP("pass", "p", false, "pass")
	RootCmd.Flags().Bool("debug", false, "debug")
	RootCmd.Flags().String("init", "", "init db 啥的，要现保证各个依赖项，安装部署成功")

}
