package netfw

// Switch represents a network switch
type Switch struct {
	device
}

// NewSwitch factory
func NewSwitch(name string) *Switch {
	return &Switch{
		device: newDevice(name, SwitchClass),
	}
}
