package netfw

import (
	"fmt"
	"sort"
)

// Device are objects that can return their lacation
type Device interface {
	GetName() string
	GetPath() string
	GetConfig() []string
	GetClass() Class
	getParent() Device
	setParent(parent Device)

	hasChild(child Device) bool
	addChild(child Device)

	hasChildren() bool
	getChildren() []Device

	getKey() string
}

// Patchable objects can be directly patched to each other
type Patchable interface {
	Device
	IsPatched() bool
	patch(dev Patchable)
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

// GetClass returns the class of the device
func (d *device) GetClass() Class {
	return d.class
}

// getKey returns the key that identifies the device
func (d *device) getKey() string {
	return d.GetClass().ToString() + "," + d.GetName()
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
	key := child.getKey()
	d.children[key] = child
	d.childKeys = append(d.childKeys, key)
	sort.Strings(d.childKeys)
	child.setParent(d)
}

// hasChildren checks to see if a child already exists
func (d *device) hasChildren() bool {
	return (len(d.children) > 0)
}

// hasChild checks to see if a child already exists
func (d *device) hasChild(child Device) bool {
	key := child.getKey()
	if _, ok := d.children[key]; ok {
		return true
	}
	return false
}

// getChildren returns all child devices
func (d *device) getChildren() []Device {
	children := make([]Device, 0, len(d.children))

	for _, k := range d.childKeys {
		children = append(children, d.children[k])
	}
	return children
}

// Read ..
func (d *device) GetConfig() []string {
	conf := []string{d.GetPath() + "\n"}

	for _, v := range d.children {
		conf = append(conf, v.GetConfig()...)
	}
	return conf
}

// Patch connects to devices together
func Patch(dev1, dev2 Patchable) error {
	if dev1.IsPatched() {
		return fmt.Errorf("unable to patch %v as it is laready patched", dev1.GetName())
	}

	if dev2.IsPatched() {
		return fmt.Errorf("unable to patch %v as it is laready patched", dev2.GetName())
	}
	dev1.patch(dev2)
	dev2.patch(dev1)
	return nil
}
