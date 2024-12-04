package messaging

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/core"
)

func EventErrorStatus(agentId string, msg *Message) *core.Status {
	err := errors.New(fmt.Sprintf("error: message event:%v is invalid for agent:%v", msg.Event(), agentId))
	return core.NewStatusError(core.StatusInvalidArgument, err)
}

func MessageContentTypeErrorStatus(agentId string, msg *Message) *core.Status {
	err := errors.New(fmt.Sprintf("error: message content:%v is invalid for agent:%v and event:%v", msg.ContentType(), agentId, msg.Event()))
	return core.NewStatusError(core.StatusInvalidArgument, err)
}
