package event

func EventChannel(evName string) (eventChannel chan interface{}, unsubscribe func()) {
	eventChannel = make(chan interface{})
	evMan := newEvent(evName).addListenerChan(eventChannel)
	return eventChannel, func() {
		// cleanup
		evMan.rmListeningChan(eventChannel)
		close(eventChannel)
	}
}

func FireEvent(evName string, data interface{}) {
	newEvent(evName).fireEvent(data)
}
