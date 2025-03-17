package messaging

import (
	"fmt"
	"time"
)

type ActivityItem struct {
	Agent   Agent
	Event   string
	Source  string
	Content any
}

func (a ActivityItem) IsEmpty() bool {
	return a.Agent == nil
}

type ActivityFunc func(a ActivityItem)

func Activity(a ActivityItem) {
	uri := "<nil>"
	if a.Agent != nil {
		uri = a.Agent.Uri()
	}
	fmt.Printf("active-> %v [%v] [%v] [%v] [%v]\n", FmtRFC3339Millis(time.Now().UTC()), uri, a.Event, a.Source, a.Content)
}
