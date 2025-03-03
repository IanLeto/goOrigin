package cmd

import (
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Initialize resources like MySQL and Elasticsearch",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取命令行参数

	},
}

func init() {
	// 添加命令行参数
	syncCmd.Flags().StringArray("type", []string{"record", "github"}, "")

}
