package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gustavocovas/goscsim"
	"github.com/gustavocovas/goscsim/actors"
	"github.com/gustavocovas/goscsim/events"
)

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

func main() {
	var tripsFile string

	flag.StringVar(&tripsFile, "trips", "", "File containing trips definitions")
	flag.Parse()

	trips, err := loadTrips(tripsFile)
	if err != nil {
		log.Fatalf("Error loading trips: %v", err)
	}

	eventQueue := events.New()

	for _, trip := range trips {
		car := &actors.Car{EventQueue: eventQueue, Name: trip.Name}
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
