package netfw

import "fmt"

// Site represents a physical location
type Site struct {
	*device
}

// NewSite factory
func NewSite(name string) *Site {
	return &Site{
		device: newDevice(name, SiteClass),
	}
}

// AddFirewall adds a firewall to the site
func (s *Site) AddFirewall(fw *Firewall) error {
	if fw == nil {
		return fmt.Errorf("error setting firewall, cannot be nil")
	}
	name := fw.GetName()
	if s.hasChild(fw) {
		return fmt.Errorf("error setting %s firewall, firewall already set", name)
	}
	if fw.getParent() != nil {
		return fmt.Errorf("error setting %s firewall, firewall already has a site", name)
	}
	s.addChild(fw)
	return nil
}

// AddSwitch adds a switch to the site
func (s *Site) AddSwitch(sw *Switch) error {
	if sw == nil {
		return fmt.Errorf("error add switch, cannot be nil")
	}
	name := sw.GetName()
	if s.hasChild(sw) {
		return fmt.Errorf("error adding %s switch, switch already present", name)
	}
	if sw.getParent() != nil {
		return fmt.Errorf("error setting %s site, switch already has a site", name)
	}
	s.addChild(sw)
	return nil
}
