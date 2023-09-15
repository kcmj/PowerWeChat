package contract

import "net/http"

type EventInterface interface {
	GetToUserName() string
	GetFromUserName() string
	GetCreateTime() string
	GetMsgType() string
	GetEvent() string
	GetChangeType() string
	ReadMessage(msg interface{}) error
	GetContent() []byte
	GetSessionFrom() string
	GetMiniGameEventPayload() string
	GetMiniGameEventIsMock() bool
}

type EventHandlerInterface interface {
	Handle(request *http.Request, header EventInterface, content interface{}) interface{}
}
