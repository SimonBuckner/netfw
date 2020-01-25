package main

/*
	Firewall has x ports
	 A port can be WAN, LAN or MAD
		 WAN - Wide Area Network
		 LAN - Local Area Network
		 MAD - Management And Diagnostics

	A LAN can have a Switch

*/

// Firewall represents a firewall at a site
type Firewall struct {
	Name       string
	Interfaces map[string]*Interface
}

// Interface represents a physical Ethernet Interface on the firewall
type Interface struct {
	ParentFirewall *Firewall
	Label          string
	Zone           Zone
}

// NewFirewall creates a new Firewall
func NewFirewall(name string) *Firewall {
	return &Firewall{
		Name:       name,
		Interfaces: make(map[string]*Interface),
	}
}

// Zone represents a network zone
type Zone int

const (
	// WANZone denotes Wide Area Network Zone
	WANZone Zone = iota
	// LANZone denotes Local Area Network
	LANZone
	// MADZone denotes Management And Diagnostics
	MADZone
)
