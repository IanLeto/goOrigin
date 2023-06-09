package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/pkg/utils"
)

func paramsStr(v string, err error) string {
	utils.NoError(err)
	return v
}

var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		config := paramsStr(cmd.Flags().GetString("config"))
		PreRun(config)
		DebugServer()
	},
}

func Execute() {
	utils.NoError(RootCmd.Execute())
}

func init() {
	RootCmd.Flags().StringP("config", "c", "", "config")
	RootCmd.Flags().Bool("debug", false, "debug")

}
