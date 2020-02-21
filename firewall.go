package netfw

import "fmt"

// Firewall represents the network border device
type Firewall struct {
	*device
}

// NewFirewall factory
func NewFirewall(name string) *Firewall {
	return &Firewall{
		device: newDevice(name, FirewallClass),
	}
}

// AddIface adds a network interace to the firewall
func (fw *Firewall) AddIface(iface *Iface) error {
	if iface == nil {
		return fmt.Errorf("error adding interface, cannot be nil")
	}
	if fw.hasChild(iface) {
		return fmt.Errorf("unable to add interface as interface already exists")
	}
	if iface.getParent() != nil {
		return fmt.Errorf("unable to add interface as interface already has a parent")
	}
	fw.addChild(iface)
	return nil
}
