package event

import (
	"github.com/asaskevich/EventBus"
)

var Bus *Event

type Event struct {
	EventBus.Bus
}

// 订阅周期任务
// to del 注意语法 fn 必须要这样写
func (receiver *Event) SubPeriodicTask(topic string, fn func(args ...interface{}) error) {
	//utils.ErrorLog(receiver.Subscribe(topic, fn))
	panic("")
}

// 发布周期任务
func (receiver *Event) PubPeriodicTask(topic string, args ...interface{}) {
	receiver.Publish(topic, args...)
}

// 发布无回调事件
func (receiver *Event) Pub(topic string) {
	receiver.Publish(topic)
}

func InitEvent() error {
	Bus = &Event{EventBus.New()}
	return nil
}

func init() {

}
