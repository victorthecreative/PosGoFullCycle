package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {

	return e.Name
}

func (e *TestEvent) GetDataTime() string {

	return ""
}

func (e *TestEvent) GetPayload() interface{} {

	return e.Payload
}

type TestEventHandler struct {
}

func (h *TestEventHandler) Handle(event EventInterface) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = *NewEventDispatcher()

	suite.handler = TestEventHandler{}
	suite.handler2 = TestEventHandler{}
	suite.handler3 = TestEventHandler{}
	suite.event = TestEvent{
		Name:    "test",
		Payload: "test",
	}
	suite.event2 = TestEvent{
		Name:    "test2",
		Payload: "test2",
	}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	assert.True(suite.T(), true)
}
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
