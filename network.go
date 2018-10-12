package goscsim

import "encoding/xml"

// TODO: Maybe this should be xml-specific types

type Link struct {
	XMLName   xml.Name `xml:"link"`
	ID        int64    `xml:"id,attr"`
	From      int64    `xml:"from,attr"`
	To        int64    `xml:"to,attr"`
	Length    float64  `xml:"length,attr"`
	Freespeed float64  `xml:"freespeed,attr"`
	Lanes     float64  `xml:"permlanes,attr"`
}

type Network struct {
	XMLName      xml.Name     `xml:"network"`
	LinksElement LinksElement `xml:"links"`
}

type LinksElement struct {
	XMLName xml.Name `xml:"links"`
	Links   []Link   `xml:"link"`
}
