package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"

	"gonum.org/v1/gonum/graph"

	"github.com/gustavocovas/goscsim"
	"github.com/gustavocovas/goscsim/actors"
	"github.com/gustavocovas/goscsim/events"
	"gonum.org/v1/gonum/graph/simple"
)

var cellSize = 7.5

func loadTrips(filename string) ([]goscsim.Trip, error) {
	log.Printf("Loading trips from %s\n", filename)
	tripsFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to open trips file: %v", err)
	}

	var matrix goscsim.InterSCSimulatorMatrix
	xml.Unmarshal(tripsFile, &matrix)

	log.Println("Done loading trips")
	return matrix.Trips, nil
}

func loadNetwork(filename string) (graph.WeightedDirected, error) {
	log.Printf("Loading network from %s\n", filename)
	networkFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to open network file: %v", err)
	}

	var network goscsim.XMLNetwork
	xml.Unmarshal(networkFile, &network)

	networkGraph := simple.NewWeightedDirectedGraph(0, math.Inf(1))

	for _, link := range network.LinksElement.Links {
		capacity := int64(link.Lanes * (link.Length / cellSize))

		networkGraph.SetWeightedEdge(goscsim.Link{
			F:         link.From,
			T:         link.To,
			W:         link.Length,
			Freespeed: link.Freespeed,
			Capacity:  capacity,
		})
	}

	log.Println("Done loading network")
	return networkGraph, nil
}

func main() {
	var (
		tripsFile   string
		networkFile string
	)

	flag.StringVar(&tripsFile, "trips", "", "File containing trips definitions")
	flag.StringVar(&networkFile, "network", "", "File containing network")
	flag.Parse()

	trips, err := loadTrips(tripsFile)
	if err != nil {
		log.Fatalf("Error loading trips: %v", err)
	}

	network, err := loadNetwork(networkFile)
	if err != nil {
		log.Fatalf("Error loading network: %v", err)
	}

	eventQueue := events.New()

	for _, trip := range trips {
		car := &actors.Car{
			EventQueue:  eventQueue,
			Network:     network,
			Name:        trip.Name,
			Origin:      trip.Origin,
			Destination: trip.Destination,
		}
		eventQueue.Push(&goscsim.Event{Time: trip.StartTime, Actor: car})
	}

	log.Println("Starting simulation")
	for eventQueue.Len() > 0 {
		event := eventQueue.Pop()

		if event.Time > 86400 {
			break
		}

		event.Actor.Act(event.Time)
	}
	log.Println("Finished simulation")
}
