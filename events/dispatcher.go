package events

import (
	"fmt"
	"sort"

	"github.com/crowdeco/skeleton/configs"
)

type Dispatcher struct {
	Events map[string]configs.Listener
}

func (d *Dispatcher) Register(listeners []configs.Listener) {
	sort.Slice(listeners, func(i, j int) bool {
		return listeners[i].Priority() > listeners[j].Priority()
	})

	for _, listener := range listeners {
		if _, ok := d.Events[listener.Listen()]; ok {
			panic(fmt.Sprintf("the '%s' event is already registered", listener.Listen()))
		}

		d.Events[listener.Listen()] = listener
	}
}

func (d *Dispatcher) Dispatch(name string, event interface{}) error {
	if _, ok := d.Events[name]; !ok {
		return fmt.Errorf("the '%s' event is already registered", name)
	}

	d.Events[name].Handle(event)

	return nil
}
