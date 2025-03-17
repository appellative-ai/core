package messaging

import (
	"fmt"
	"time"
)

type ActivityFunc func(agent Agent, event, source string, content any)

func Activity(agent Agent, event, source string, content any) {
	fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), agent.Uri(), event, source, content)
}
