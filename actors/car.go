package actors

import (
	"fmt"

	"github.com/gustavocovas/goscsim"
)

type Car struct {
	EventQueue goscsim.EventQueue
	Name       string
}

func (c *Car) Act(tick int) {
	fmt.Printf("%d %s vroom\n", tick, c.Name)
	c.EventQueue.Push(&goscsim.Event{Time: tick + 10, Actor: c})
}
