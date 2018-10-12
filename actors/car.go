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
	currentNode graph.Node
	path        []graph.Node
}

func (c *Car) Act(tick int) {
	switch c.state {
	case ready:
		df := path.DijkstraFrom(simple.Node(c.Origin), c.Network)
		shortestPath, _ := df.To(c.Destination)

		c.currentNode = simple.Node(c.Origin)
		c.path = shortestPath[1:]

		c.state = traveling
		c.moveToNextNode(tick)

	case traveling:
		c.moveToNextNode(tick)
	}
}

func (c *Car) moveToNextNode(tick int) {

	nextNode := c.path[0]

	if len(c.path) == 1 {
		c.state = finished
	} else {
		c.path = c.path[1:]
	}

	speed := float64(14)
	edge := c.Network.WeightedEdge(c.currentNode.ID(), nextNode.ID())
	time := int(edge.Weight() / speed)

	log.Printf("t=%d, name=%v orig=%v dest=%v len=%v time=%v\n",
		tick,
		c.Name,
		c.currentNode,
		nextNode,
		edge.Weight(),
		time,
	)

	c.EventQueue.Push(&goscsim.Event{Time: tick + time, Actor: c})

	c.currentNode = nextNode
}
