package messaging

type NotifyFunc func(status *Status)

type Notifier interface {
	Notify(status *Status)
}
