package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"goOrigin/config"
	"goOrigin/internal/router"
	"goOrigin/pkg/utils"
	"net/http"
)

func DebugServer() {
	g := gin.New()
	router.Load(g, nil)
	utils.NoError(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.ConfV2.Base.Port), g))
}

var runCmd = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			flags = cmd.Flags()
		)
		if v, err := flags.GetBool("version"); err != nil && v {
			fmt.Println("1.1.1")
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	runCmd.Flags().Bool("version", true, "run dev version")

}
