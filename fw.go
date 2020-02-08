package netfw

import "fmt"

// Firewall represents the network border device
type Firewall struct {
	device
	interfaces map[string]*Iface
}

// NewFirewall builds a new firewall object
func NewFirewall(name string) *Firewall {
	return &Firewall{
		device: device{
			name:  name,
			class: FirewallClass,
		},
		interfaces: make(map[string]*Iface),
	}
}

// SetSite connects a firewall to a specific site
func (fw *Firewall) SetSite(site *Site) error {
	if site == nil {
		return fmt.Errorf("error setting site, cannot be nil")
	}
	if fw.parent != nil {
		return fmt.Errorf("error setting site, site already set")
	}
	fw.setParent(site)
	return nil
}

// AddIface adds a network interace to the firewall
func (fw *Firewall) AddIface(iface *Iface) error {
	if iface == nil {
		return fmt.Errorf("error adding interface, cannot be nil")
	}
	if _, ok := fw.interfaces[iface.name]; ok {
		return fmt.Errorf("unable to add interface as interface already exists")
	}
	if iface.getParent() != nil {
		return fmt.Errorf("unable to add interface as interface already has a parent")
	}
	fw.interfaces[iface.name] = iface
	iface.setParent(fw)
	return nil
}
