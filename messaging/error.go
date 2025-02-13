package messaging

import (
	"errors"
	"fmt"
)

func EventError(agentId string, msg *Message) error {
	err := errors.New(fmt.Sprintf("error: message event:%v is invalid for agent:%v", msg.Event(), agentId))
	return err
}

func MessageContentTypeError(agentId string, msg *Message) error {
	err := errors.New(fmt.Sprintf("error: message content:%v is invalid for agent:%v and event:%v", msg.ContentType(), agentId, msg.Event()))
	return err
}
