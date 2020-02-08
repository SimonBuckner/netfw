package netfw

import "testing"

func TestNewIface(t *testing.T) {
	iface := NewIface("eth0")
	if iface.class != IfaceClass {
		t.Errorf("expecting IfaceClass, got %v", iface.class.ToString())
	}
}

func TestIfaceSetTrunkMode(t *testing.T) {
	iface := NewIface("eth0")
	{
		err := iface.SetTrunkMode()
		if err != nil {
			t.Errorf("error setting trunk mode; expecting nil, got %v", err)
		}
	}
	{
		if iface.mode.ToString() != "trunk" {
			t.Errorf("unexepcted interface mode; expecting 'trunk', got '%v'", iface.mode.ToString())
		}
	}
}

func TestIfaceSetEdgeMode(t *testing.T) {
	iface := NewIface("eth0")
	{
		err := iface.SetEdgeMode()
		if err != nil {
			t.Errorf("error setting edge mode; expecting nil, got %v", err)
		}
	}
	{
		if iface.mode.ToString() != "edge" {
			t.Errorf("unexepcted interface mode; expecting 'edge', got '%v'", iface.mode.ToString())
		}
	}
}
