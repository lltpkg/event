package event

import "sync"

type eventGroup struct {
	EventName string
	Listeners []chan interface{}
}

// singletons
var evManager = &singletonMutex{evGroups: make(map[string]*eventGroup)}

type singletonMutex struct {
	Mt       sync.Mutex
	evGroups map[string]*eventGroup
}

func newEvent(eventName string) *eventGroup {
	evManager.Mt.Lock()
	defer evManager.Mt.Unlock()
	if evManager.evGroups[eventName] != nil {
		return evManager.evGroups[eventName]
	}
	appEventInstance := &eventGroup{EventName: eventName}
	evManager.evGroups[eventName] = appEventInstance
	return appEventInstance
}

func (a *eventGroup) fireEvent(data interface{}) {
	evManager.Mt.Lock()
	defer evManager.Mt.Unlock()

	for _, listener := range a.Listeners {
		go func(ch chan interface{}) {
			ch <- data
		}(listener)
	}

}

func (a *eventGroup) addListenerChan(listener chan interface{}) *eventGroup {
	evManager.Mt.Lock()
	defer evManager.Mt.Unlock()
	a.Listeners = append(a.Listeners, listener)
	return a
}

func (a *eventGroup) rmListeningChan(listener chan interface{}) {
	evManager.Mt.Lock()
	defer evManager.Mt.Unlock()

	for i, l := range a.Listeners {
		if l == listener {
			a.Listeners = append(a.Listeners[:i], a.Listeners[i+1:]...)
			break
		}
	}

	if len(a.Listeners) == 0 {
		evManager.evGroups[a.EventName] = nil
	}
}
