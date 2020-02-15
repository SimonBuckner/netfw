package netfw

// EdgePatch represents a physical patch point
type EdgePatch struct {
	*device
}

// NewEdgePort factory
func NewEdgePort(name string) *EdgePatch {
	return &EdgePatch{
		device: newDevice(name, EdgeClass),
	}
}
