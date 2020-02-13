package netfw

// Iface represents a network interfaces in a switch or firewall
type Iface struct {
	device
	mode   Mode
	duplex Duplex
	speed  Speed
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

// SetTrunkMode set the interface mode to Trunk
func (iface *Iface) SetTrunkMode() error {
	iface.mode = TrunkMode
	return nil
}

// SetEdgeMode set the interface mode to Edge
func (iface *Iface) SetEdgeMode() error {
	iface.mode = EdgeMode
	return nil
}

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

// Speed representents the speed of an interface
type Speed int

const (
	// Speed10 indicates a 10 Mbps interface
	Speed10 Speed = 10
	// Speed100 indicates a 100 Mbps interface
	Speed100 Speed = 100
	// Speed1000 indicates a 1000 Mbps interface
	Speed1000 Speed = 1000
)

var speedText = map[Speed]string{
	Speed10:   "10-mbps",
	Speed100:  "100-mbps",
	Speed1000: "1000-mbps",
}

// Duplex represent the duplex mode of an interface
type Duplex int

const (
	// HalfDuplex indicates a half duplex interface
	HalfDuplex Duplex = iota
	// FullDuplex indicated a full duplex interface
	FullDuplex
)

var duplexText = map[Duplex]string{
	HalfDuplex: "half-duplex",
	FullDuplex: "full-duplex",
}

// ToString returns a string representation of a Duplex
func (duplex Duplex) ToString() string {
	return duplexText[duplex]
}

// ToString returns a string representation of a Speed
func (speed Speed) ToString() string {
	return speedText[speed]
}
