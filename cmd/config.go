package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/config"
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
		logger.Sugar().Info("%s", config.NewV2Config())
	},
}

func init() {

}
