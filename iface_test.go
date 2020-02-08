package netfw

import "testing"

func TestNewIface(t *testing.T) {
	iface := NewIface("eth0")
	if iface.class != IfaceClass {
		t.Errorf("expecting IfaceClass, got %v", iface.class.ToString())
	}
}
