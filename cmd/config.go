package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/config"
	"goOrigin/pkg/utils"
	"os"
)

var (
	brokers          string
	topic            string
	consumerGroup    string
	consumerMember   string
	outFile          string
	consumeMsgsLimit int
	username         string
	password         string
	auth             string
)

var configCmd = &cobra.Command{
	Use: "config",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := paramsStr(cmd.Flags().GetString("path"))
		utils.NoError(os.Setenv("configPath", configPath))
		logger.Sugar().Info("%s", utils.ToJson(config.NewV2Config()))
	},
}

func init() {
	configCmd.Flags().StringP("path", "p", "", "config")
}
