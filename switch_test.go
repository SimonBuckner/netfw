package netfw

import "testing"

func TestSwitchAddPort(t *testing.T) {

	sw := NewSwitch("sw1")

	{
		err := sw.AddPort("eth0/", 0, Speed1000)
		if err != nil {
			t.Errorf("error adding port; expecting nil, got %v", err)
		}
	}

	{
		err := sw.AddPort("eth0/", 0, Speed1000)
		if err == nil {
			t.Errorf("error adding port; expecting and error, got nil")
		}
	}
	{
		ports := sw.getChildren()

		if len(ports) != 1 {
			t.Errorf("error adding port; expecting 1 port, got %v", len(ports))
		}

		if ports[0].GetName() != "eth0/0" {
			t.Errorf("error adding port; unexpected name, got %v", ports[0].GetName())
		}
	}

}

func TestSwitchAddPortRange(t *testing.T) {

	sw := NewSwitch("switch")

	{
		err := sw.AddPortRange("eth0/", 10, 5, Speed100)
		if err == nil {
			t.Errorf("error adding ports 10 thru 5; expecting err, got nil")
		}
	}

	{
		err := sw.AddPortRange("eth0/", -1, 5, Speed100)
		if err == nil {
			t.Errorf("error adding ports -1 thru 5; expecting err, got nil")
		}
	}

	{
		err := sw.AddPortRange("eth0/", 5, -1, Speed100)
		if err == nil {
			t.Errorf("error adding ports 5 thru -1; expecting err, got nil")
		}
	}

	{
		err := sw.AddPortRange("eth0/", 10, 19, Speed100)
		if err != nil {
			t.Errorf("error adding ports 10 thru 19; expecting nil, got %v", err)
		}
		ports := sw.getChildren()
		if len(ports) != 10 {
			t.Errorf("error adding ports 0 thru 9; expeting 10 ports, got %v", len(ports))
		}
		name := ports[0].GetName()
		if name != "eth0/10" {
			t.Errorf("error adding port; unexpected name, got %v", name)
		}
	}

	{
		err := sw.AddPortRange("eth0/", 9, 24, Speed100)
		if err == nil {
			t.Errorf("error adding additional ports 9 thru 24; expecting err, got nil")
		}
	}

}
