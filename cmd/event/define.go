package event

import (
	"github.com/asaskevich/EventBus"
	"goOrigin/pkg/utils"
)

var Bus *Event

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

// 发布无回调事件
func (receiver *Event) Pub(topic string) {
	receiver.event.Publish(topic)
}

// 订阅无回调事件
func (receiver *Event) Sub(topic string) {
	//var fn func() error
	//receiver.event.SubscribeAsync(topic, func() {}func, false)
}

func InitEvent() error {
	Bus = &Event{event: EventBus.New()}
	return nil
}

func init() {

}
