The `github.com/lltpkg/event` package provides a simple and thread-safe mechanism for managing events and listeners in Go applications.

## Installation

To use the `event` package in your Go project, you can use the following `go get` command:

```bash
go get -u github.com/lltpkg/event
```

## Quick start

```go
package main

import (
	"fmt"

	"github.com/lltpkg/event"
)

func main() {
	evName := "Greeting"
	go func() {
		data := "World"
		event.FireEvent(evName, data)
	}()
	evChan, unSub := event.EventChanel(evName)
	defer unSub()
	receivedData := <-evChan
	fmt.Println("Hello,", receivedData)
}

```

## Usage

### Creating Events and Listeners

The package allows you to create named events and associate listeners with them. Use the EventChanel function to create an event channel and associate a cleanup function with it:

```go
// Create an event channel(listener) for "exampleEvent"
eventChan, cleanup := event.EventChanel("exampleEvent")
// Cleanup resources when done
defer cleanup()


```

### Triggering Events

You can trigger events using the FireEvent function. This function allows you to send data to all registered listeners for a specific event:

```go
// Trigger the event anywhere else
event.FireEvent("exampleEvent", "event data")
```

## Contribution

Feel free to open issues and submit pull requests for improvements or bug fixes. We appreciate any contributions that make the event package more robust and versatile.

## Maintainers

<a href="https://github.com/lyluongthien" target="_blank">
    <img src="https://avatars.githubusercontent.com/u/43800313?v=4" alt="Kai - Ly Luong Thien - Maintainer of github.com/lltpkg/event" style="width:100px; display:block"/>
    1. Kai - @lyluongthien
</a>

## License

This package is licensed under the MIT License.

> Note: Update the documentation based on the actual functionality and features provided by the `event` package. Include details about how to use the package, any configuration options, and examples for common use cases.
