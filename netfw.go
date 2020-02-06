package netfw

import "fmt"

// Device ..
type Device interface {
	GetPath() string
	SetParent(parent Device) error
	SetChild(child Device) error
}

// baseDevice
type baseDevice struct {
	label  string
	parent *baseDevice
}

// newBaseDevice ..
func newBaseDevice(label string) *baseDevice {
	return &baseDevice{
		label:  label,
		parent: nil,
	}
}

// GetPath ..
func (bd *baseDevice) GetPath() string {
	fmt.Println(bd.label)
	if bd.parent == nil {
		return "\\" + bd.label
	}
	return bd.parent.GetPath() + "\\" + bd.label
}

// SetParent ..
func (bd *baseDevice) SetParent(parent *baseDevice) error {
	if bd.parent != nil {
		return fmt.Errorf("cannot set parent when parent already exists")
	}
	bd.parent = parent
	return nil
}

func main() {

}
