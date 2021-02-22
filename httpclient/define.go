package httpclient

import (
	"goOrigin/config"
	"goOrigin/event"
)

var CC *CCClient

func init() {
	var ch = make(chan struct{})
	CC = NewCCClient(config.GlobalConfig.Client.CC)
	// 发布心跳事件
	event.GlobalEventBus.SubPeriodicTask("ccPing", PingCCClient(CC, ch))

}
