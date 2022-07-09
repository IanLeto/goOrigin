package cmd

//var dataCmd = &cobra.Command{
//	Use: "data",
//	Run: func(cmd *cobra.Command, args []string) {
//		PreRun("")
//		logrus.SetLevel(logrus.InfoLevel)
//		if err := storage.Mongo.Ping(); err != nil {
//			fmt.Println("无效的mongo链接")
//		}
//		c := storage.Mongo.C("ian")
//		switch paramsStr(cmd.Flags().GetString("mongo")) {
//		case "index":
//			indexes, err := c.Indexes()
//			if err != nil {
//				fmt.Println(err)
//			}
//			for _, index := range indexes {
//				v, _ := json.Marshal(index)
//				logrus.Infof("index:%s", string(v))
//			}
//		case "initData":
//			err := storage.InitSchema(storage.Mongo, paramsStr(cmd.Flags().GetString("collection")))
//			if err != nil {
//				fmt.Println(err)
//			}
//		}
//
//	},
//}
//
//func init() {
//	RootCmd.AddCommand(dataCmd)
//	dataCmd.Flags().StringP("path", "", "", "run dev version")
//	dataCmd.Flags().StringP("mongo", "", "", "run dev version")
//	dataCmd.Flags().StringP("collection", "c", "", "run dev version")
//
//}
