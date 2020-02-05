package main

import (
	"fmt"
)

/*
	Firewall has x ports
	 A port can be WAN, LAN or MAD
		 WAN - Wide Area Network
		 LAN - Local Area Network
		 MAD - Management And Diagnostics

	A LAN can have a Switch
*/

// FW represents a firewall at a site
type FW struct {
	Name   string
	Ifaces map[string]*Iface
}

// NewFW creates a new Firewall
func NewFW(name string) *FW {
	return &FW{
		Name:   name,
		Ifaces: make(map[string]*Iface),
	}
}

// IsFW returns true if the device is a FW
func (fw *FW) IsFW() bool {
	return true
}

// IsSW returns true as this is a firewall
func (fw *FW) IsSW() bool {
	return false
}

// IsIface returns false as this is not an Iface
func (fw *FW) IsIface() bool {
	return false
}

// GetParent returns the parent device
func (fw *FW) GetParent() Device {
	return nil
}

// SetParent is invalid as a firewall has no parent
func (fw *FW) SetParent(parent Device) error {
	return fmt.Errorf("unable to set parent on a firewall")
}

// AddIface adds an interface to the firewall
func (fw *FW) AddIface(iface *Iface) error {
	if _, found := fw.Ifaces[iface.Label]; found {
		return fmt.Errorf("firewall '%v' already has interface '%v''", fw.Name, iface.Label)
	}

	fw.Ifaces[iface.Label] = iface
	return nil
}

// Dump produces a textual representation of the firewall config
func (fw *FW) Dump() {
	fmt.Printf("hostname %v\n", fw.Name)
	fw.DumpZone(AdminZone)
	fw.DumpZone(PublicZone)
	fw.DumpZone(PrivateZone)
}

// DumpZone produces a textual representation of a zone config on a firewall
func (fw *FW) DumpZone(zone Zone) {
	ifaces := fw.GetZoneInterfaces(AdminZone)
	for _, iface := range ifaces {
		fmt.Printf("iface %v\n", iface.Label)
		fmt.Printf("  zone %v\n", zoneName[iface.Zone])

		fmt.Printf("  end\n")
	}
}

// GetZoneInterfaces returns all interfaces matching a zone
func (fw *FW) GetZoneInterfaces(zone Zone) []*Iface {
	ifaces := []*Iface{}
	for _, iface := range fw.Ifaces {
		if iface.Zone == zone {
			ifaces = append(ifaces, iface)
		}
	}
	return ifaces
}
