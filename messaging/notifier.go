package messaging

type NotifyFunc func(status *Status) *Status

type Notifier interface {
	Notify(status *Status)
}
