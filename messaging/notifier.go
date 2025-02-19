package messaging

type Notifier interface {
	Notify(status *Status)
}
