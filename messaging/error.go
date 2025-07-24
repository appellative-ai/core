package messaging

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/core/std"
)

func EmptyMapError(location string) *std.Status {
	return std.NewStatus(std.StatusInvalidArgument, location, errors.New("map is nil"))
}

func MapContentError(location string, key string) *std.Status {
	return std.NewStatus(std.StatusInvalidArgument, location, errors.New(fmt.Sprintf("map does not contain key: %v", key)))
}

func EmptyReviewError(location string) *std.Status {
	return std.NewStatus(std.StatusInvalidArgument, location, errors.New("review is nil"))
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
