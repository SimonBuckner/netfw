package netfw

// Mode represents the operating mode of the interface
type Mode int

const (
	// EdgeMode indicates that edge devices are connected to this interface
	// Edge devices inlude phones and computers and printers
	EdgeMode Mode = iota
	// TrunkMode indicates the attached devices are other switches or firewalls
	TrunkMode
)

var modeText = map[Mode]string{
	EdgeMode:  "edge",
	TrunkMode: "trunk",
}

// ToString returns a string representation of a Mode
func (m Mode) ToString() string {
	return modeText[m]
}

// Iface represents a network interfaces in a switch or firewall
type Iface struct {
	device
	Trunk bool
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
