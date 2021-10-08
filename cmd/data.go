package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/pkg/utils"
)

var dataCmd = &cobra.Command{
	Use: "config",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			flags = cmd.Flags()
		)
		_, err := flags.GetBool("version")
		utils.NoError(err)
		print("1.0.1")
	},
}

func init() {
	RootCmd.AddCommand(dataCmd)
	dataCmd.Flags().StringP("path", "", "", "run dev version")

}
