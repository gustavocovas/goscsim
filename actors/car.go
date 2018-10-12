package actors

import (
	"github.com/gustavocovas/goscsim"
)

type Car struct {
	EventQueue goscsim.EventQueue
	Name       string
	count      int
}

func (c *Car) Act(tick int) {
	// fmt.Printf("%d %s vroom\n", tick, c.Name)
	if c.count < 10 {
		c.EventQueue.Push(&goscsim.Event{Time: tick + 10, Actor: c})
		c.count++
	}
}
