package event

import (
	"sync"
	"testing"
)

func TestEventChannelAndFireEvent(t *testing.T) {
	// Arrange
	evName := "testEvent"
	data := "testData"

	// Act
	evChan, cleanup := EventChannel(evName)

	// Assert
	// Check that the channel is not nil
	if evChan == nil {
		t.Error("Event channel is nil")
	}

	// Check that the channel is registered in the event manager
	evManager.Mt.Lock()
	group := evManager.evGroups[evName]
	evManager.Mt.Unlock()

	if group == nil {
		t.Error("Event group not registered for EventChannel")
	}

	// Act
	FireEvent(evName, data)

	// Assert
	// Check that the data is received on the channel
	receivedData := <-evChan
	if receivedData != data {
		t.Errorf("Expected data %v, but received %v", data, receivedData)
	}

	// Cleanup
	cleanup()

	// Assert
	// Check that the channel is closed after cleanup
	_, ok := <-evChan
	if ok {
		t.Error("Event channel is not closed after cleanup")
	}

	// Check that the listener is removed from the event group
	evManager.Mt.Lock()
	listeners := group.Listeners
	evManager.Mt.Unlock()

	if len(listeners) != 0 {
		t.Error("Listener not removed from the event group")
	}

	// Check that the event group is removed from the event manager
	evManager.Mt.Lock()
	removedGroup := evManager.evGroups[evName]
	evManager.Mt.Unlock()

	if removedGroup != nil {
		t.Error("Event group not removed from the event manager")
	}
}

func TestConcurrentFireEvent(t *testing.T) {
	// Arrange
	evName := "concurrentTestEvent"
	data := "concurrentTestData"
	numListeners := 30
	var wg sync.WaitGroup

	// Act
	evChan, cleanup := EventChannel(evName)

	for i := 0; i < numListeners; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			FireEvent(evName, data)
		}()
	}

	// Assert
	// Wait for all goroutines to finish
	wg.Wait()

	// Check that data is received on the channel for each listener
	for i := 0; i < numListeners; i++ {
		receivedData := <-evChan
		if receivedData != data {
			t.Errorf("Expected data %v, but received %v", data, receivedData)
		}
	}

	// Cleanup
	cleanup()

	// Assert
	// Check that the channel is closed after cleanup
	_, ok := <-evChan
	if ok {
		t.Error("Event channel is not closed after cleanup")
	}
}

func TestMultiEventChannelForSingleEvent(t *testing.T) {
	// Arrange
	evName := "testEvent1"
	data := "testData1"
	numListeners := 30
	var wg sync.WaitGroup

	// Act
	evChan1, cleanup1 := EventChannel(evName)
	evChan2, cleanup2 := EventChannel(evName)

	for i := 0; i < numListeners; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			FireEvent(evName, data)
		}()
	}

	// Assert
	// Wait for all goroutines to finish
	wg.Wait()

	// Check that data is received on the channel for each listener
	for i := 0; i < numListeners; i++ {
		receivedData1 := <-evChan1
		if receivedData1 != data {
			t.Errorf("Expected data %v, but received %v", data, receivedData1)
		}

		receivedData2 := <-evChan2
		if receivedData2 != data {
			t.Errorf("Expected data %v, but received %v", data, receivedData2)
		}
	}

	// Cleanup
	cleanup1()
	cleanup2()

	// Assert
	// Check that the channels are closed after cleanup
	_, ok1 := <-evChan1
	if ok1 {
		t.Error("Event channel 1 is not closed after cleanup")
	}

	_, ok2 := <-evChan2
	if ok2 {
		t.Error("Event channel 2 is not closed after cleanup")
	}
}
