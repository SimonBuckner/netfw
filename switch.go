package netfw

import (
	"fmt"
	"strconv"
)

// Switch represents a network switch
type Switch struct {
	*device
}

// NewSwitch factory
func NewSwitch(name string) *Switch {
	return &Switch{
		device: newDevice(name, SwitchClass),
	}
}

// AddPort add a single port to a switch
func (sw *Switch) AddPort(path string, index int, speed Speed) error {
	name := path + strconv.Itoa(index)
	iface := NewIface(name)
	iface.speed = speed
	if sw.hasChild(iface) == true {
		return fmt.Errorf("error: a port with that deisgnation already exists")
	}
	sw.addChild(iface)
	return nil
}

// AddPortRange adss a range of ports to a switch
func (sw *Switch) AddPortRange(path string, first, last int, speed Speed) error {
	if last <= first {
		return fmt.Errorf("Last port must be higher than first port")
	}
	if first < 0 || last < 0 {
		return fmt.Errorf("first and last port must be 0 and higher")
	}
	for i := first; i <= last; i++ {
		if err := sw.AddPort(path, i, speed); err != nil {
			return err
		}
	}
	return nil
}
