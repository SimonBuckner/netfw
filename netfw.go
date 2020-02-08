package netfw

// Device are objects that can return their lacation
type Device interface {
	GetPath() string
	setParent(parent Device)
	getParent() Device
}

var _ Device = &device{}

// device represents the core features of a device
type device struct {
	name   string
	class  Class
	parent Device
}

// Class denotes the type of device
type Class int

const (
	// SiteClass ..
	SiteClass Class = iota
	// FirewallClass ..
	FirewallClass
	// IfaceClass ..
	IfaceClass
	// SwitchClass ..
	SwitchClass
	// VlanClass ..
	VlanClass
)

var classText = map[Class]string{
	SiteClass:     "site",
	FirewallClass: "firewall",
	IfaceClass:    "iface",
	SwitchClass:   "switch",
	VlanClass:     "vlan",
}

// ToString returns a string representation of a class
func (c Class) ToString() string {
	return classText[c]
}

// GetPath returns this devices location
func (d *device) GetPath() string {
	if d.parent == nil {
		return d.class.ToString() + "=" + d.name

	}
	return d.parent.GetPath() + "," + d.class.ToString() + "=" + d.name
}

// getParent returns the parent device
func (d *device) getParent() Device {
	return d.parent
}

// setParent set the parent device for this device
func (d *device) setParent(parent Device) {
	d.parent = parent
}
