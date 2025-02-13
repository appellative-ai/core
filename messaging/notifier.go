package messaging

type Notifier interface {
	Notify(err error)
}
