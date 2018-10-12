package goscsim

import "encoding/xml"

type Trip struct {
	Name                string `xml:"name,attr"`
	Origin              int64  `xml:"origin,attr"`
	Destination         int64  `xml:"destination,attr"`
	Count               int    `xml:"count,attr"`
	StartTime           int    `xml:"start,attr"`
	DigitalRailsCapable bool   `xml:"digital_rails_capable,attr"`
}

type InterSCSimulatorMatrix struct {
	XMLName xml.Name `xml:"scsimulator_matrix"`
	Trips   []Trip   `xml:"trip"`
}
