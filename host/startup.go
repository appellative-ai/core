package host

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"time"
)

const (
	startupLocation = PkgPath + ":Startup"
)

// Exchange - host package controller2
var Exchange = messaging.NewExchange()

// ContentMap - slice of any content to be included in a message
type ContentMap map[string]map[string]string

type ResourceMap map[string]aspect.HttpExchange

func RegisterControlAgent(uri string, handler messaging.Handler) (messaging.Agent, error) {
	a, err := messaging.NewControlAgent(uri, handler)
	if err != nil {
		return a, err
	}
	return a, Exchange.Register(a)
}

// Startup - templated function to start all registered resources.
func Startup(duration time.Duration, resources ResourceMap) bool {
	return startup(Exchange, duration, resources)
}

func startup(ex *messaging.Exchange, duration time.Duration, resources ResourceMap) bool {
	var failures []string
	var count = ex.Count()

	if count == 0 {
		return true
	}
	cache := messaging.NewCache()
	toSend := createToSend(ex, resources, messaging.NewCacheHandler(cache))
	sendMessages(ex, toSend)
	for wait := time.Duration(float64(duration) * 0.25); duration >= 0; duration -= wait {
		time.Sleep(wait)
		// Check for completion
		if cache.Count() < count {
			continue
		}
		// Check for failed resources
		failures = cache.Exclude(messaging.StartupEvent, http.StatusOK)
		if len(failures) == 0 {
			handleStatus(cache)
			return true
		}
		break
	}
	shutdownHost(messaging.NewMessage(messaging.ControlChannelType, "", "", messaging.ShutdownEvent, nil))
	if len(failures) > 0 {
		handleErrors(failures, cache)
		return false
	}
	fmt.Printf("error: startup failure [%v]\n", errors.New(fmt.Sprintf("response counts < directory entries [%v] [%v]", cache.Count(), ex.Count())))
	return false
}

func createToSend(ex *messaging.Exchange, resources ResourceMap, fn messaging.Handler) messaging.Map {
	m := make(messaging.Map)
	for _, k := range ex.List() {
		msg := messaging.NewMessage(messaging.ControlChannelType, k, startupLocation, messaging.StartupEvent, nil)
		msg.ReplyTo = fn
		if resources != nil {
			if ex, ok := resources[k]; ok {
				msg.SetContent(messaging.ContentTypeConfig, ex)
			}
		}
		m[k] = msg
	}
	return m
}

func sendMessages(ex *messaging.Exchange, msgs messaging.Map) {
	for k := range msgs {
		ex.Send(msgs[k])
	}
}

func handleErrors(failures []string, cache *messaging.Cache) {
	for _, uri := range failures {
		msg, ok := cache.Get(uri)
		if !ok {
			continue
		}
		if msg.Status() != nil && msg.Status().Err != nil {
			fmt.Printf("error: startup failure [%v]\n", msg.Status().Err)
		}
	}
}

func handleStatus(cache *messaging.Cache) {
	for _, uri := range cache.Uri() {
		msg, ok := cache.Get(uri)
		if !ok {
			continue
		}
		if msg.Status() != nil {
			fmt.Printf("startup successful: [%v] : %s\n", uri, msg.Status().Duration)
		}
	}
}

func shutdownHost(m *messaging.Message) {

}
