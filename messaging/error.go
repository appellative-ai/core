package messaging

import (
	"errors"
	"fmt"
)

func EmptyMapError(location string) *Status {
	return NewStatus(StatusInvalidArgument, errors.New("map is nil")).WithLocation(location)
}

func MapContentError(location string, key string) *Status {
	return NewStatus(StatusInvalidArgument, errors.New(fmt.Sprintf("map does not contain key: %v", key))).WithLocation(location)
}

func EmptyReviewError(location string) *Status {
	return NewStatus(StatusInvalidArgument, errors.New("review is nil")).WithLocation(location)
}

/*
func EventError(agentId string, msg *Message) error {
	err := errors.New(fmt.Sprintf("error: message name:%v is invalid for agent:%v", msg.Name(), agentId))
	return err
}

func MessageContentTypeError(agentId string, msg *Message) error {
	err := errors.New(fmt.Sprintf("error: message content:%v is invalid for agent:%v and name:%v", msg.ContentType(), agentId, msg.Name()))
	return err
}



func ConfigContentStatusError(agent Agent, key string) *Status {
	return NewStatus(StatusInvalidArgument, errors.New(fmt.Sprintf("config map does not contain key: %v", key))).WithLocation(agent.Name())
}


*/
