package main

// Iface represents a physical Ethernet Interface on the firewall
type Iface struct {
	Parent *Device
	Child  *Device
	Label  string
	Zone   Zone
}

// NewIface returns a new interface
func NewIface(label string, zone Zone) *Iface {
	iface := Iface{
		Label: label,
		Zone:  zone,
	}
	return &iface
}

// ConnectSwitch connects a switch to the specified interface
// func (iface *Interface) ConnectSwitch(sw *Switch) error {
// 	if iface.
// }
