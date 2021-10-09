package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"goOrigin/pkg/httpclient"
	"goOrigin/pkg/utils"
)

func paramsStr(v string, err error) string {
	utils.NoError(err)
	return v
}

func paramsBool(v bool, err error) bool {
	utils.NoError(err)
	return v
}
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
	Use:   "",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		config := ""
		if v, err := cmd.Flags().GetBool("debug"); err != nil && v {
			config = paramsStr(cmd.Flags().GetString("config"))
		}
		logrus.Debugf(config)
		PreRun(config)
	},
}

func Execute() {
	utils.NoError(RootCmd.Execute())
}

func init() {
	RootCmd.Flags().StringP("config", "c", "", "config")
	RootCmd.Flags().Bool("debug", false, "debug")

}
