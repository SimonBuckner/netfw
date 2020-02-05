package main

// Device describes the common features of all devices
type Device interface {
	IsFW() bool
	IsSW() bool
	IsIface() bool
	// GetParent() Device
	// SetParent(parent Device) error
}

// Zone represents a network zone
type Zone int

const (
	// AdminZone denotete the local management and monitoring zone
	AdminZone Zone = iota
	// PublicZone denotes public (internet) facing network zone
	PublicZone
	// PrivateZone denotes local private network zone
	PrivateZone
)

var zoneName = map[Zone]string{
	AdminZone:   "admin",
	PublicZone:  "public",
	PrivateZone: "prvate",
}

func main() {

	fw := NewFW("Islington")

	wan0 := NewIface("WAN0", PublicZone)
	fw.AddIface(wan0)
	mad0 := NewIface("MAD0", AdminZone)
	fw.AddIface(mad0)
	lan0 := NewIface("LAN0", PrivateZone)
	fw.AddIface(lan0)

	sw1 := NewSW(14, OddsEvensPorts)
	for i := 1; i <= 12; i++ {
		sw1.SetPortUse(i, EdgePort)
	}
	sw1.SetPortUse(13, LinkPort)
	sw1.SetPortUse(14, LinkPort)

	fw.Dump()

}

func printSWConfigs(devices *Device) {

}
