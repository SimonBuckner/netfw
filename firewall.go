package main

import "fmt"

/*
	Firewall has x ports
	 A port can be WAN, LAN or MAD
		 WAN - Wide Area Network
		 LAN - Local Area Network
		 MAD - Management And Diagnostics

	A LAN can have a Switch
*/

// Zone represents a network zone
type Zone int

const (
	// MADZone denotes Management And Diagnostics
	MADZone Zone = iota
	// WANZone denotes Wide Area Network Zone
	WANZone
	// LANZone denotes Local Area Network
	LANZone
)

var zoneName = map[Zone]string{
	MADZone: "mad",
	WANZone: "wan",
	LANZone: "lan",
}

// Firewall represents a firewall at a site
type Firewall struct {
	Name       string
	Interfaces map[string]*Interface
}

// NewFirewall creates a new Firewall
func NewFirewall(name string) *Firewall {
	return &Firewall{
		Name:       name,
		Interfaces: make(map[string]*Interface),
	}
}

// AddInterface adds a new interface to the firewall
func (fw *Firewall) AddInterface(iface *Interface) error {
	if _, found := fw.Interfaces[iface.Label]; found {
		return fmt.Errorf("firewall '%v' already has interface '%v''", fw.Name, iface.Label)
	}

	fw.Interfaces[iface.Label] = iface
	return nil
}

// Dump produces a textual representation of the firewall config
func (fw *Firewall) Dump() {
	fmt.Printf("hostname %v\n", fw.Name)
	fw.DumpZone(MADZone)
	fw.DumpZone(WANZone)
	fw.DumpZone(LANZone)
}

// DumpZone produces a textual representation of a zone config on a firewall
func (fw *Firewall) DumpZone(zone Zone) {
	ifaces := fw.GetZoneInterfaces(MADZone)
	for _, iface := range ifaces {
		fmt.Printf("iface %v\n", iface.Label)
		fmt.Printf("  zone %v\n", zoneName[iface.Zone])

		fmt.Printf("  end\n")
	}
}

// GetZoneInterfaces returns all interfaces matching a zone
func (fw *Firewall) GetZoneInterfaces(zone Zone) []*Interface {
	ifaces := []*Interface{}
	for _, iface := range fw.Interfaces {
		if iface.Zone == zone {
			ifaces = append(ifaces, iface)
		}
	}
	return ifaces
}

// Interface represents a physical Ethernet Interface on the firewall
type Interface struct {
	ParentFirewall *Firewall
	Label          string
	Zone           Zone
	Switch         *Switch
}

// NewInterface returns a new interface
func NewInterface(label string, zone Zone) *Interface {
	iface := Interface{
		Label: label,
		Zone:  zone,
	}
	return &iface
}
