package netfw

import "fmt"

// Site represents a physical location
type Site struct {
	name    string
	primary *Firewall
	backup  *Firewall
}

// NewSite builsd a new Site object
func NewSite(name string) *Site {
	return &Site{
		name: name,
	}
}

// GetPath returns this divices location
func (s *Site) GetPath() string {
	return "site=" + s.name
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
	fw.SetSite(s)
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
	fw.SetSite(s)
	return nil
}
