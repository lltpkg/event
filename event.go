package event

func EventChanel(evName string) (chan interface{}, func()) {
	evChan := make(chan interface{})
	evMan := newEvent(evName).addListenerChan(evChan)
	return evChan, func() {
		// cleanup
		evMan.rmListeningChan(evChan)
		close(evChan)
	}
}

func FireEvent(evName string, data interface{}) {
	newEvent(evName).fireEvent(data)
}
