package actors

import (
	"log"

	"github.com/gustavocovas/goscsim"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

type state int

const (
	ready     = state(0)
	traveling = state(1)
	finished  = state(2)
)

type Car struct {
	EventQueue  goscsim.EventQueue
	Network     graph.WeightedDirected
	Name        string
	Origin      int64
	Destination int64

	state       state
	current     int64
	currentLink *goscsim.Link
	path        []graph.Node

	firstTick        int
	distanceTraveled float64
}

func (c *Car) Act(tick int) {
	switch c.state {
	case ready:
		c.firstTick = tick

		df := path.DijkstraFrom(simple.Node(c.Origin), c.Network)
		shortestPath, _ := df.To(c.Destination)

		c.current = c.Origin
		c.path = shortestPath[1:]

		c.state = traveling
		c.moveToNextNode(tick)

	case traveling:
		c.moveToNextNode(tick)
	}
}

func (c *Car) moveToNextNode(tick int) {
	if c.currentLink != nil {
		c.currentLink.Vehicles--
	}

	if c.path == nil {
		log.Printf("Warning: Nil path for car %v at tick %v", c, tick)
		return
	}

	if len(c.path) == 0 {
		log.Printf("Warning: Empty path for car %v at tick %v", c, tick)
		return
	}

	nextNode := c.path[0]

	if len(c.path) == 1 {
		c.state = finished
		log.Printf(
			"Finished trip for car %v: orig=%v, dest=%v, len=%v, time=%v",
			c.Name,
			c.Origin,
			c.Destination,
			c.distanceTraveled,
			tick-c.firstTick,
		)
	} else {
		c.path = c.path[1:]
	}

	link, ok := c.Network.WeightedEdge(c.current, nextNode.ID()).(goscsim.Link)
	if !ok {
		panic("Failed to cast network WeightedEdge to goscsim.Link")
	}

	c.distanceTraveled += link.Weight()

	speed := link.Freespeed * (1 - (float64(link.Vehicles) / float64(link.Capacity)))
	time := int(link.Weight() / speed)

	c.EventQueue.Push(&goscsim.Event{Time: tick + time, Actor: c})
	link.Vehicles++

	c.current = nextNode.ID()
	c.currentLink = &link
}
