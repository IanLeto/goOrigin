package cmd

import "github.com/spf13/cobra"

var beatCmd = &cobra.Command{
	Use: "beat",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			_ = cmd.Flags()
		)
	},
}

func init() {
	beatCmd.Flags().Bool("init", true, "初始化true 时会重新将jinja2模板初始化为新的filebeat文件。")
	//beatCmd.Flags().Bool("init", true, "rota")
}
