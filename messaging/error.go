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

func ConfigEmptyStatusError(agent Agent) *Status {
	return NewStatusError(StatusInvalidArgument, errors.New("config map is nil"), agent.Uri())
}

func ConfigContentStatusError(agent Agent, key string) *Status {
	return NewStatusError(StatusInvalidArgument, errors.New(fmt.Sprintf("config map does not contain key: %v", key)), agent.Uri())
}
