package main

import (
	"github.com/gustavocovas/goscsim"
	"github.com/gustavocovas/goscsim/actors"
	"github.com/gustavocovas/goscsim/events"
)

func main() {
	eventQueue := events.New()

	car1 := &actors.Car{EventQueue: eventQueue, Name: "car1"}
	car2 := &actors.Car{EventQueue: eventQueue, Name: "car2"}

	eventQueue.Push(&goscsim.Event{Time: 1, Actor: car1})
	eventQueue.Push(&goscsim.Event{Time: 2, Actor: car2})

	for eventQueue.Len() > 0 {
		event := eventQueue.Pop()

		if event.Time > 86400 {
			break
		}

		event.Actor.Act(event.Time)
	}
}
