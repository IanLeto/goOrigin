package httpclient

import (
	"goOrigin/config"
	"goOrigin/event"
)

var CC *CCClient

func init() {
	CC = NewCCClient(config.GlobalConfig.Client.CC)
	// 发布心跳事件
	event.GlobalEventBus.SubPeriodicTask("ccPing", PingCCClient(CC))
}
