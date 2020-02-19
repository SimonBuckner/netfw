package netfw

import "testing"

var _ Device = EdgePatch{}

func TestNewEdgePatch(t *testing.T) {
	NewEdgePatch("21")
}
