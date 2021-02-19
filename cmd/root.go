package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/httpclient"
	"goOrigin/utils"
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
	utils.CheckNoError(err)
	configVersion, err = flags.GetString("configVersion")
	utils.CheckNoError(err)
	serverVersion, err = flags.GetString("serverVersion")
	utils.CheckNoError(err)
	serverName, err = flags.GetString("serverName")
	utils.CheckNoError(err)
	configTemplate, err = flags.GetString("configTemplate")
	utils.CheckNoError(err)
	configTemplatePath, err = flags.GetString("configTemplatePath")
	utils.CheckNoError(err)
	businessTemplate, err = flags.GetString("businessTemplate")
	utils.CheckNoError(err)
	businessTemplatePath, err = flags.GetString("businessTemplatePath")
	utils.CheckNoError(err)
	commonFile, err = flags.GetString("commonFile")
	utils.CheckNoError(err)
	commonFilePath, err = flags.GetString("commonFilePath")
	utils.CheckNoError(err)
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
		utils.CheckNoError(err)
		if path == "" {
			output = os.Stdout
		} else {
			output, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0777)
			utils.CheckNoError(err)
			defer func() {
				utils.CheckNoError(err)
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
	RootCmd.Flags().StringP("config_version", "cv", "", "run dev version")
	RootCmd.Flags().StringP("server_version", "sv", "", "run dev version")
	RootCmd.Flags().StringP("target_config_version", "tcv", "", "run dev version")
	RootCmd.Flags().StringP("target_server_version", "tsv", "", "run dev version")
	RootCmd.Flags().BoolP("version", "v", false, "current version")
	RootCmd.Flags().StringP("output", "", "", "run dev version")

}
