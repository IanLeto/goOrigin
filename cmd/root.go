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
		config := paramsStr(cmd.Flags().GetString("config"))
		PreRun(config)
		DebugServer()
	},
}

//func Execute() {
//	utils.NoError(RootCmd.Execute())
//}

func init() {
	RootCmd.Flags().StringP("config", "c", "", "config")
	RootCmd.Flags().BoolP("pass", "p", false, "pass")
	RootCmd.Flags().Bool("debug", false, "debug")

}
