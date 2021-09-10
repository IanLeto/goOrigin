package httpclient

var CC *CCClient

func init() {
	// 心跳
	//var ch = make(chan struct{})
	//CC = NewCCClient(run.Conf.Client.CC)
	//// 发布心跳事件
	//event.GlobalEventBus.SubPeriodicTask("ccPing", PingCCClient(CC, ch))
	//go func() {
	//	<-ch
	//}()
}
