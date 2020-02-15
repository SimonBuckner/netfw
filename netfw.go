package netfw

import "sort"

// Device are objects that can return their lacation
type Device interface {
	GetName() string
	GetPath() string
	setParent(parent Device)
	getParent() Device
	addChild(child Device)
	hasChildren(child Device) bool
	getChildren() []Device
}

// Patchable objects can be directly patched to each other
type Patchable interface {
	Device
	Patch(to Patchable) error
	IsPatched() bool
	GetPatchedDevice() Device
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
	// EdgeClass ..
	EdgeClass
)

var classText = map[Class]string{
	SiteClass:     "site",
	FirewallClass: "firewall",
	IfaceClass:    "iface",
	SwitchClass:   "switch",
	VlanClass:     "vlan",
	EdgeClass:     "edge",
}

// ToString returns a string representation of a class
func (c Class) ToString() string {
	return classText[c]
}

// device represents the core features of a device
type device struct {
	name      string
	class     Class
	parent    Device
	children  map[string]Device
	childKeys []string
}

// newDevice factory
func newDevice(name string, class Class) *device {
	return &device{
		name:      name,
		class:     class,
		parent:    nil,
		children:  make(map[string]Device),
		childKeys: make([]string, 0),
	}
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
	name := child.GetName()
	d.children[name] = child
	d.childKeys = append(d.childKeys, name)
	sort.Strings(d.childKeys)
	child.setParent(d)
}

// hasChildren checks to see if a child already exists
func (d *device) hasChildren(child Device) bool {
	if _, ok := d.children[child.GetName()]; ok {
		return true
	}
	return false
}

// getChildren returns all child devices
func (d *device) getChildren() []Device {
	children := make([]Device, 0, len(d.children))

	for _, name := range d.childKeys {
		children = append(children, d.children[name])
	}
	return children
}
