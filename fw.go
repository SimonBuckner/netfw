package netfw

import "fmt"

// Firewall represents the network border device
type Firewall struct {
	name string
	site *Site
}

// NewFirewall builds a new firewall object
func NewFirewall(name string) *Firewall {
	return &Firewall{
		name: name,
	}
}

// GetPath returns this divices location
func (fw *Firewall) GetPath() string {
	path := ""
	if fw.site != nil {
		path = fw.site.GetPath() + ","
	}
	return path + "firewall=" + fw.name
}

// SetSite connects a firewall to a specific site
func (fw *Firewall) SetSite(site *Site) error {
	if site == nil {
		return fmt.Errorf("error setting site, cannot be nil")
	}
	if fw.site != nil {
		return fmt.Errorf("error setting site, site already set")
	}
	fw.site = site
	return nil
}
