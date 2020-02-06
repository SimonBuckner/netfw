package netfw

import "testing"

func TestGetParent(t *testing.T) {

	parentBD := newBaseDevice("firewall")
	if parentBD.GetPath() != "\\firewall" {
		t.Errorf("GetPath returned %s, expected %s", parentBD.GetPath(), "\\firewall")
	}

	childBD1 := newBaseDevice("child 1")
	childBD1.SetParent(parentBD)
	path := childBD1.GetPath()
	if path != "\\firewall\\child 1" {
		t.Errorf("GetPath returned %s, expected %s", path, "\\firewall\\child 1")
	}

}
