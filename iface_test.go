package netfw

import "testing"

var _ Device = &Iface{}
var _ Patchable = &Iface{}

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

func TestIfaceEnumStrings(t *testing.T) {

	{
		mode := EdgeMode
		if mode.ToString() != "edge" {
			t.Errorf("unexepcted interface mode; expecting 'edge', got '%v'", mode.ToString())
		}
	}

	{
		mode := TrunkMode
		if mode.ToString() != "trunk" {
			t.Errorf("unexepcted interface mode; expecting 'trunk', got '%v'", mode.ToString())
		}
	}

	{
		duplex := HalfDuplex
		if duplex.ToString() != "half-duplex" {
			t.Errorf("unexepcted interface duplex; expecting 'half-duplex', got '%v'", duplex.ToString())
		}
	}

	{
		duplex := FullDuplex
		if duplex.ToString() != "full-duplex" {
			t.Errorf("unexepcted interface duplex; expecting 'full-duplex', got '%v'", duplex.ToString())
		}
	}

	{
		speed := Speed10
		if speed.ToString() != "10-mbps" {
			t.Errorf("unexepcted interface speed; expecting '10-mbps', got '%v'", speed.ToString())
		}
	}

	{
		speed := Speed100
		if speed.ToString() != "100-mbps" {
			t.Errorf("unexepcted interface speed; expecting '100-mbps', got '%v'", speed.ToString())
		}
	}

	{
		speed := Speed1000
		if speed.ToString() != "1000-mbps" {
			t.Errorf("unexepcted interface speed; expecting '1000-mbps', got '%v'", speed.ToString())
		}
	}
}

func TestIfacepatch(t *testing.T) {

	eth0 := NewIface("eth0")
	eth1 := NewIface("eth1")
	eth2 := NewIface("eth2")
	eth1.patch(eth2)
	eth2.patch(eth1)

	{
		if eth0.IsPatched() == true {
			t.Errorf("error checking eth0 patching; expecting false, got true")
		}
	}

	{
		if eth1.IsPatched() == false {
			t.Errorf("error checking eth1 patching; expecting true, got false")
		}
	}

	{
		if eth1.patched != eth2 {
			t.Errorf("error checking eth1 patching; expecting eth2, got %v", eth1.patched)
		}
	}

	{
		if eth2.patched != eth1 {
			t.Errorf("error checking eth2 patching; expecting eth1, got %v", eth2.patched)
		}
	}

}
