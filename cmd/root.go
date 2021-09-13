package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/pkg/httpclient"
	"goOrigin/pkg/utils"
	"io"
	"os"
)

func paramCheck(cmd *cobra.Command) map[string]string {
	var (
		err                  error
		flags                = cmd.Flags()
		product              string
		configVersion        string
		serverVersion        string
		serverName           string
		configTemplate       string
		configTemplatePath   string
		businessTemplate     string
		businessTemplatePath string
		commonFile           string
		commonFilePath       string
	)
	product, err = flags.GetString("product")
	utils.NoError(err)
	configVersion, err = flags.GetString("configVersion")
	utils.NoError(err)
	serverVersion, err = flags.GetString("serverVersion")
	utils.NoError(err)
	serverName, err = flags.GetString("serverName")
	utils.NoError(err)
	configTemplate, err = flags.GetString("configTemplate")
	utils.NoError(err)
	configTemplatePath, err = flags.GetString("configTemplatePath")
	utils.NoError(err)
	businessTemplate, err = flags.GetString("businessTemplate")
	utils.NoError(err)
	businessTemplatePath, err = flags.GetString("businessTemplatePath")
	utils.NoError(err)
	commonFile, err = flags.GetString("commonFile")
	utils.NoError(err)
	commonFilePath, err = flags.GetString("commonFilePath")
	utils.NoError(err)
	return map[string]string{
		"product":              product,
		"configVersion":        configVersion,
		"serverVersion":        serverVersion,
		"serverName":           serverName,
		"configTemplate":       configTemplate,
		"configTemplatePath":   configTemplatePath,
		"businessTemplate":     businessTemplate,
		"businessTemplatePath": businessTemplatePath,
		"commonFile":           commonFile,
		"commonFilePath":       commonFilePath,
	}

}

func formatResponseInfo(res *httpclient.CCResponseInfo, err error) string {

	return ""
}

var RootCmd = &cobra.Command{
	Use:   "details",
	Short: "print details of cc",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err error
			//input  *os.File
			output *os.File
		)
		client := httpclient.NewCCClient(nil)
		params := paramCheck(cmd)
		path, err := cmd.Flags().GetString("output")
		utils.NoError(err)
		if path == "" {
			output = os.Stdout
		} else {
			output, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0777)
			utils.NoError(err)
			defer func() {
				utils.NoError(err)
			}()
		}
		_, err = io.WriteString(output, formatResponseInfo(client.Mkc(params["product"], params["configVersion"], params["serverVersion"], params["serverName"], params["configTemplate"], params["configTemplatePath"], params["businessTemplate"], params["businessTemplatePath"], params["commonFile"], params["commonFilePath"], false)))
	},
}

func init() {
	RootCmd.Flags().StringP("run", "d", "", "run dev version")
	RootCmd.Flags().StringP("compare", "", "", "run dev version")
	RootCmd.Flags().StringP("test", "", "", "run dev version")
	RootCmd.Flags().StringP("mkc", "", "", "run dev version")
	RootCmd.Flags().StringP("config_version", "c", "", "run dev version")
	RootCmd.Flags().StringP("server_version", "s", "", "run dev version")
	RootCmd.Flags().StringP("target_config_version", "", "", "run dev version")
	RootCmd.Flags().StringP("target_server_version", "", "", "run dev version")
	RootCmd.Flags().BoolP("version", "v", false, "current version")
	RootCmd.Flags().StringP("output", "", "", "run dev version")

}
