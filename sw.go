package main

import (
	"fmt"
	"sort"
	"strconv"
)

// SW represents a network swith
type SW struct {
	PortDirection
	MaxPort int
	Ports   map[int]*Port
}

// Port represents a port on a network switch
type Port struct {
	Index       int
	Description string
	Use         PortUse
}

// PortDirection denotes the direction on the switch ports are laid out
type PortDirection int

const (
	// SequentialPorts denotes a switch with ports numebrs 1-x along the top row
	// with ports numbered x+1 to y along the bottom row
	SequentialPorts PortDirection = iota
	// OddsEvensPorts denotes a switch with odd ports along the top and even port along the bottom row
	OddsEvensPorts
)

var portDirectionText = map[PortDirection]string{
	SequentialPorts: "Sequential",
	OddsEvensPorts:  "Odds then Evens",
}

// ToString returns the text representation of the PortDirection
func (pd PortDirection) ToString() string {
	return portDirectionText[pd]
}

// PortUse denotes the type of use to the port will be used for
// Either Edge for connecting devices or Link for connecting switches
type PortUse int

// ToString returns the text representation of the PortUse
func (p PortUse) ToString() string {
	return portUseText[p]
}

const (
	// EdgePort denotes a port that will be used to connect devices
	EdgePort PortUse = iota
	// LinkPort denotes a port that will be used to connect switches
	LinkPort
	// NotConfigured denotes a port that has not been configured
	NotConfigured
)

var portUseText = map[PortUse]string{
	EdgePort:      "Edge Port",
	LinkPort:      "Link Port",
	NotConfigured: "Not Configured",
}

// NewSW creates a new switch
func NewSW(maxPort int, direction PortDirection) *SW {
	s := SW{
		PortDirection: direction,
		MaxPort:       maxPort,
		Ports:         make(map[int]*Port),
	}
	for i := 1; i <= maxPort; i++ {
		s.Ports[i] = &Port{
			Index:       i,
			Description: fmt.Sprintf("Port %2d", i),
			Use:         NotConfigured,
		}
	}
	return &s
}

// SetPortUse configures the specified port for a specific use
func (s *SW) SetPortUse(index int, use PortUse) {
	if port, ok := s.Ports[index]; ok {
		port.Use = use
	}
}

// Dump produces a textual representation of a zone config on a firewall
func (s *SW) Dump() {

	keys := make([]int, s.MaxPort)
	for k := range s.Ports {
		keys[k-1] = k
	}

	var edgeKeys []int
	var linkKeys []int
	var noConfig []int

	for _, key := range keys {
		switch s.Ports[key].Use {
		case EdgePort:
			edgeKeys = append(edgeKeys, key)
		case LinkPort:
			linkKeys = append(linkKeys, key)
		default:
			noConfig = append(noConfig, key)
		}
	}

	fmt.Printf("\nPort Layout : %v", s.PortDirection.ToString())
	fmt.Printf("\nMax Ports   : %2d", s.MaxPort)
	fmt.Printf("\nNo Config   : %v", intArrayToText(noConfig))
	fmt.Printf("\nLink Ports  : %v", intArrayToText(linkKeys))
	fmt.Printf("\nEdge Ports  : %v", intArrayToText(edgeKeys))

}

// intArrayToText turns a list of int to a text description where ranges are
// separated by '-', so '1,2,3,4' would become '1-4' and individual numbers are
// separated by ',', so '1,2,3,4,7,9' would become '1-4,7,9'
func intArrayToText(array []int) string {
	if len(array) == 0 {
		return ""
	}

	sort.Ints(array)
	text := strconv.Itoa(array[0])
	index := 1
	run := 0
	for {
		if index == len(array) {
			if run > 0 {
				text = text + "-" + strconv.Itoa(array[index-1])
			}
			return text
		}
		if array[index]-1 == array[index-1] {
			run++
			index++
			continue
		}
		if array[index] == array[index-1] {
			// Remove duplicates
			index++
			continue
		}
		if run > 0 {
			text = text + "-" + strconv.Itoa(array[index-1])
			run = 0
		}
		text = text + "," + strconv.Itoa(array[index])
		index++
	}
}
