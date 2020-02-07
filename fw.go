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
func (fw *Firewall) AddIface(name string, iface *Iface) error {
	if _, ok := fw.interfaces[name]; ok {
		return fmt.Errorf("unable to add interface as interface already exists")
	}

	return nil
}
