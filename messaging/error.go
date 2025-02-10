package messaging

import (
	"errors"
	"fmt"
)

func EventErrorStatus(agentId string, msg *Message) *aspect.Status {
	err := errors.New(fmt.Sprintf("error: message event:%v is invalid for agent:%v", msg.Event(), agentId))
	return aspect.NewStatusError(aspect.StatusInvalidArgument, err)
}

func MessageContentTypeErrorStatus(agentId string, msg *Message) *aspect.Status {
	err := errors.New(fmt.Sprintf("error: message content:%v is invalid for agent:%v and event:%v", msg.ContentType(), agentId, msg.Event()))
	return aspect.NewStatusError(aspect.StatusInvalidArgument, err)
}
