package netfw

import "fmt"

// Site represents a physical location
type Site struct {
	device
	primary *Firewall
	backup  *Firewall
}

// NewSite builsd a new Site object
func NewSite(name string) *Site {
	return &Site{
		device: device{
			name:  name,
			class: SiteClass,
		},
	}
}

// SetPrimaryFirewall sets the primary firewall for the site
func (s *Site) SetPrimaryFirewall(fw *Firewall) error {
	if fw == nil {
		return fmt.Errorf("error setting primary firewall, cannot be nil")
	}
	if s.primary != nil {
		return fmt.Errorf("error setting primary firewall, firewall already set")
	}
	s.primary = fw
	if err := fw.SetSite(s); err != nil {
		s.primary = nil
		return err
	}
	return nil
}

// SetBackupFirewall sets the backup firewall for the site
func (s *Site) SetBackupFirewall(fw *Firewall) error {
	if fw == nil {
		return fmt.Errorf("error setting backup firewall, cannot be nil")
	}
	if s.primary == nil {
		return fmt.Errorf("error setting backup firewall, no primary firewall set")
	}
	if s.backup != nil {
		return fmt.Errorf("error setting backup firewall, firewall already set")
	}
	if fw == s.primary {
		return fmt.Errorf("error setting backup firewall, primary & backup cannot be the same")
	}
	s.backup = fw
	if err := fw.SetSite(s); err != nil {
		s.backup = nil
		return err
	}
	return nil
}
