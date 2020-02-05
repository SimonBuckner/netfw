package main

import "testing"

func TestNewFW(t *testing.T) {
	{
		test := NewFW("TestNewFirewall")
		if test.Name != "TestNewFirewall" {
			t.Errorf("Expected firewall called 'Test 1' but got %v", test.Name)
		}
	}
}

func TestAddInterface(t *testing.T) {
	fw := NewFW("TestAddInterface")
	iface := NewIface("Interface", AdminZone)
	err1 := fw.AddIface(iface)
	if err1 != nil {
		t.Errorf("add interface returned unexpected error; %v", err1)
	}
	err2 := fw.AddIface(iface)
	if err2 == nil {
		t.Error("add interface returned unexpected nill error")
	}

}
