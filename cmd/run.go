package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/pkg/utils"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run",
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
	RootCmd.AddCommand(runCmd)
	runCmd.Flags().Bool("version", true, "run dev version")
	runCmd.Flags().StringP("compare", "", "", "run dev version")

}
