package main

import "testing"

func TestNewFirewall(t *testing.T) {
	{
		test := NewFirewall("TestNewFirewall")
		if test.Name != "TestNewFirewall" {
			t.Errorf("Expected firewall called 'Test 1' but got %v", test.Name)
		}
	}
}

func TestAddInterface(t *testing.T) {
	fw := NewFirewall("TestAddInterface")
	iface := NewInterface("Interface", MADZone)
	err1 := fw.AddInterface(iface)
	if err1 != nil {
		t.Errorf("add interface returned unexpected error; %v", err1)
	}
	err2 := fw.AddInterface(iface)
	if err2 == nil {
		t.Error("add interface returned unexpected nill error")
	}

}
