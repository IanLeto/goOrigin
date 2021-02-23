package event_test

import (
	"goOrigin/config"
	"goOrigin/event"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// eventBusSuite :
type eventBusSuite struct {
	suite.Suite
	debug bool
	conf  *config.Config
	event *event.Event
	count string
}

func (s *eventBusSuite) SetupTest() {
	s.event = event.GlobalEventBus
	s.conf = config.GlobalConfig
}

// TestMarshal :
func (s *eventBusSuite) TestEventBus() {
	if s.conf.RunMode != "debug" {
		return
	}
	testCases := []struct {
		except, topic string
		fn            func(args ...interface{}) error
	}{
		{except: "ok", topic: "case1", fn: func(args ...interface{}) error {
			s.count = "ok"
			return nil
		}},
	}
	for _, v := range testCases {
		s.event.SubPeriodicTask(v.topic, v.fn)
		s.event.PubPeriodicTask(v.topic, nil)
		time.Sleep(1 * time.Second)
		s.Equal(v.except, s.count)
	}

}

// TestViperConfiguration :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(eventBusSuite))
}
