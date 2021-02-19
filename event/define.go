package event

import (
	"github.com/asaskevich/EventBus"
	"goOrigin/utils"
)

var GlobalEventBus *Event

type Event struct {
	event EventBus.Bus
}

// 订阅周期任务
// to del 注意语法 fn 必须要这样写
func (receiver *Event) SubPeriodicTask(topic string, fn func(args ...interface{}) error) {
	utils.ErrorLog(receiver.event.Subscribe(topic, fn))
}

// 发布周期任务
func (receiver *Event) PubPeriodicTask(topic string, args ...interface{}) {
	receiver.event.Publish(topic, args...)
}

func init() {
	GlobalEventBus = &Event{
		event: EventBus.New(),
	}

}
