package cmd

import (
	"github.com/spf13/cobra"
	"goOrigin/config"
	"goOrigin/pkg/utils"
	"os"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Initialize resources like MySQL and Elasticsearch",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取命令行参数
		dbType, _ := cmd.Flags().GetString("type") // 资源类型：mysql 或 es
		configPath := paramsStr(cmd.Flags().GetString("path"))
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		user, _ := cmd.Flags().GetString("user")
		//password, _ := cmd.Flags().GetString("password")
		dbName, _ := cmd.Flags().GetString("db") // 对于 MySQL 来说是数据库名

		// 设置配置路径
		utils.NoError(os.Setenv("configPath", configPath))
		logger.Sugar().Info("Starting initialization for type: %s", dbType)

		// 打印初始化参数（可替换为实际的初始化逻辑）
		logger.Sugar().Info("Host: %s, Port: %d, User: %s, DB: %s", host, port, user, dbName)

		// 在这里调用对应的初始化逻辑
		config.ConfV2 = config.NewV2ConfigFromPath(configPath)
		logger.Sugar().Infof("%s", utils.ToJson(config.ConfV2))
		switch dbType {
		case "mysql":
			logger.Sugar().Info("Initializing MySQL...")

		case "es":
			logger.Sugar().Info("Initializing Elasticsearch...")
			// 调用 Elasticsearch 初始化逻辑
		default:
			logger.Sugar().Error("Unsupported type: %s. Use 'mysql' or 'es'", dbType)
		}
	},
}

func init() {
	// 添加命令行参数
	migrateCmd.Flags().StringP("type", "t", "", "Type of resource to initialize (mysql or es)")
	migrateCmd.Flags().StringP("path", "c", "", "Path to configuration file")
	migrateCmd.Flags().StringP("host", "", "", "Host of the resource (e.g., MySQL or Elasticsearch)")
	migrateCmd.Flags().IntP("port", "", 0, "Port of the resource (e.g., MySQL or Elasticsearch)")
	migrateCmd.Flags().StringP("user", "u", "", "Username for the resource")
	migrateCmd.Flags().StringP("password", "p", "", "Password for the resource")
	migrateCmd.Flags().StringP("db", "d", "", "Database name (for MySQL)")

}
