package messaging

import (
	"errors"
	"fmt"
)

func EventError(agentId string, msg *Message) error {
	err := errors.New(fmt.Sprintf("error: message name:%v is invalid for agent:%v", msg.Name(), agentId))
	return err
}

func MessageContentTypeError(agentId string, msg *Message) error {
	err := errors.New(fmt.Sprintf("error: message content:%v is invalid for agent:%v and name:%v", msg.ContentType(), agentId, msg.Name()))
	return err
}

func ConfigEmptyStatusError(agent Agent) *Status {
	return NewStatus(StatusInvalidArgument, errors.New("config map is nil")).WithLocation(agent.Name())
}

func ConfigContentStatusError(agent Agent, key string) *Status {
	return NewStatus(StatusInvalidArgument, errors.New(fmt.Sprintf("config map does not contain key: %v", key))).WithLocation(agent.Name())
}
