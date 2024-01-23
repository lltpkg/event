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

	if evManager.evGroups[eventName] != nil {
		return evManager.evGroups[eventName]
	}
	evManager.Mt.Lock()
	appEventInstance := &eventGroup{EventName: eventName}
	evManager.evGroups[eventName] = appEventInstance
	evManager.Mt.Unlock()
	return appEventInstance
}

func (a *eventGroup) fireEvent(data interface{}) {

	var wg sync.WaitGroup
	evManager.Mt.Lock()
	for _, listener := range a.Listeners {
		wg.Add(1)
		go func(l chan interface{}) {
			defer wg.Done()
			l <- data
		}(listener)
	}
	wg.Wait()
	evManager.Mt.Unlock()

}

func (a *eventGroup) addListenerChan(listener chan interface{}) *eventGroup {
	evManager.Mt.Lock()
	a.Listeners = append(a.Listeners, listener)
	evManager.Mt.Unlock()
	return a
}

func (a *eventGroup) rmListeningChan(listener chan interface{}) {

	for i, l := range a.Listeners {
		if l == listener {
			evManager.Mt.Lock()
			a.Listeners = append(a.Listeners[:i], a.Listeners[i+1:]...)
			evManager.Mt.Unlock()
			break
		}
	}

	if len(a.Listeners) == 0 {
		evManager.Mt.Lock()
		evManager.evGroups[a.EventName] = nil
		evManager.Mt.Unlock()
	}

}
