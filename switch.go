package main

import (
	"fmt"
)

// Switch represents a network switch
type Switch struct {
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

// NewSwitch creates a new switch
func NewSwitch(maxPort int, direction PortDirection) *Switch {
	s := Switch{
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
func (s *Switch) SetPortUse(index int, use PortUse) {
	if port, ok := s.Ports[index]; ok {
		port.Use = use
	}
}

// DumpSwitch shows the config of a switch
func (s *Switch) DumpSwitch() {

	keys := make([]int, s.MaxPort)
	for k := range s.Ports {
		keys[k-1] = k
	}

	edgeKeys := []int{}
	linkKeys := []int{}
	noConfig := []int{}
	for _, key := range keys {
		switch s.Ports[key].Use {
		case EdgePort:
			edgeKeys = append(edgeKeys, key)
			fmt.Printf("\nEdge : %2d", key)
		case LinkPort:
			linkKeys = append(linkKeys, key)
			fmt.Printf("\nLink :  %2d", key)
		default:
			noConfig = append(noConfig, key)
			fmt.Printf("\nNot  :  %2d", key)
		}
	}
	fmt.Printf("\nPort Layout : %v", s.PortDirection.ToString())
	fmt.Printf("\nMax Ports   : %2d", s.MaxPort)
	fmt.Printf("\nNo Config   : %2d", noConfig)
	fmt.Printf("\nLink Ports  : %2d", linkKeys)
	fmt.Printf("\nEdge Ports  : %2d", edgeKeys)

	// fmt.Printf("Edge Ports  :  ")
	// sepVal := ""
	// for _, k := range edgeKeys {
	// 	fmt.Printf("%s%d", sepVal, k)
	// 	sepVal = ","
	// }
	// fmt.Println()

	// fmt.Printf("Link Ports  :  ")
	// sepVal = ""
	// for _, k := range linkKeys {
	// 	fmt.Printf("%s%d", sepVal, k)
	// 	sepVal = ","
	// }
	// fmt.Println()

	// edgeText := string(edgeKeys[0])
	// firstKey := edgeKeys[0]
	// lastKey := edgeKeys[0]
	// inRange := false

	// for i:=1; i< len(edgeKeys); i++ {
	// 	if edgeKeys[i] == lastKey+1 {
	// 		inRange = true
	// 		lastKey = edgeKeys[i]
	// 		continue
	// 	}
	// 	if !inRange {
	// 		edgeText = edgeText + "," + string(i)
	// 	} else {

	// 	}
	// }

	// }

}
