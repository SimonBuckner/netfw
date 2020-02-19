package netfw

// EdgePatch represents a physical patch point
type EdgePatch struct {
	*device
}

// NewEdgePatch factory
func NewEdgePatch(name string) *EdgePatch {
	return &EdgePatch{
		device: newDevice(name, EdgeClass),
	}
}
