package netfw

// Device are objects that can return their lacation
type Device interface {
	GetName() string
	GetPath() string
	setParent(parent Device)
	getParent() Device
	addChild(child Device)
	childExists(child Device) bool
	getChildren() []Device
}

var _ Device = &device{}

// device represents the core features of a device
type device struct {
	name     string
	class    Class
	parent   Device
	children map[string]Device
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

// GetName returnes the name f the device
func (d *device) GetName() string {
	return d.name
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

// addChild addsa child for this device
func (d *device) addChild(child Device) {
	if d.children == nil {
		d.children = make(map[string]Device)
	}
	d.children[child.GetName()] = child
	child.setParent(d)
}

// childExists checks to see if a child already exists
func (d *device) childExists(child Device) bool {
	if _, ok := d.children[child.GetName()]; ok {
		return true
	}
	return false
}

// getChildren returns all child devices
func (d *device) getChildren() []Device {
	children := make([]Device, 0, len(d.children))
	for _, v := range d.children {
		children = append(children, v)
	}
	return children
}
