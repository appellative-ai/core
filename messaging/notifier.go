package messaging

type NotifyFunc func(status *Status) *Status

type Notifier interface {
	Notify(status *Status)
}

func Notify(notifier NotifyFunc, status *Status) *Status {
	if notifier != nil {
		return notifier(status)
	}
	return status
}
