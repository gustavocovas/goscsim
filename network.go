package goscsim

import (
	"encoding/xml"

	"gonum.org/v1/gonum/graph/simple"

	"gonum.org/v1/gonum/graph"
)

type XMLLink struct {
	XMLName   xml.Name `xml:"link"`
	ID        int64    `xml:"id,attr"`
	From      int64    `xml:"from,attr"`
	To        int64    `xml:"to,attr"`
	Length    float64  `xml:"length,attr"`
	Freespeed float64  `xml:"freespeed,attr"`
	Lanes     float64  `xml:"permlanes,attr"`
}

type XMLNetwork struct {
	XMLName      xml.Name        `xml:"network"`
	LinksElement XMLLinksElement `xml:"links"`
}

type XMLLinksElement struct {
	XMLName xml.Name  `xml:"links"`
	Links   []XMLLink `xml:"link"`
}

type Link struct {
	F         int64
	T         int64
	W         float64
	Vehicles  int64
	Capacity  int64
	Freespeed float64
}

func (l Link) From() graph.Node {
	return simple.Node(l.F)
}

func (l Link) To() graph.Node {
	return simple.Node(l.T)
}

func (l Link) Weight() float64 {
	return l.W
}
