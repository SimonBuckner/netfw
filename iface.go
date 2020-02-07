package netfw

// Iface represents a network interfaces in a switch or firewall
type Iface struct {
	device
}

// NewIface builds a new interface
func NewIface(name string) *Iface {
	return &Iface{
		device: device{
			name:  name,
			class: IfaceClass,
		},
	}
}
