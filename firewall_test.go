package netfw

import "testing"

func TestNewFirewall(t *testing.T) {
	fw := NewFirewall("work")
	if fw.class != FirewallClass {
		t.Errorf("expecting SiteClass, got %v", fw.class.ToString())
	}
}

func TestFWAddIface(t *testing.T) {
	fw := NewFirewall("fw")
	eth0 := NewIface("eth0")
	eth1 := NewIface("eth1")
	eth2 := NewIface("eth2")

	{
		err := fw.AddIface(nil)
		if err == nil {
			t.Errorf("error adding nil interface; expecting error, got nil")
		}
	}
	{
		err := fw.AddIface(eth0)
		if err != nil {
			t.Errorf("error adding interface eth0; expecting 'nil', got '%v'", err)
		}
	}
	{
		err := fw.AddIface(eth0)
		if err == nil {
			t.Errorf("error adding interface eth0 twice; expecting error, got nil")
		}
	}
	{
		err := fw.AddIface(eth1)
		if err != nil {
			t.Errorf("error adding interface eth1; expecting 'nil', got '%v'", err)
		}
	}
	{
		err := fw.AddIface(eth1)
		if err == nil {
			t.Errorf("error adding interface eth1 twice; expecting error, got nil")
		}
	}
	{
		eth2.setParent(fw)
		err := fw.AddIface(eth2)
		if err == nil {
			t.Errorf("error adding interface eth2 that already has a parent; expecting nil, got %v", err)
		}
	}
}
