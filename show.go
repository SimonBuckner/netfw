package netfw

import "fmt"

// ShowDevice lists out all the devices
func ShowDevice(dev Device) {
	fmt.Println(dev.GetPath())
	for _, child := range dev.getChildren() {
		ShowDevice(child)
	}
}
