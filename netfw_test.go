package netfw

import (
	"testing"
)

func TestSites(t *testing.T) {
	site := NewSite("work")
	primary := NewFirewall("one")
	backup := NewFirewall("two")

	{
		exp := "site=work"
		got := site.GetPath()
		if got != exp {
			t.Errorf("error getting firewall path, expected %s, got %s", exp, got)
		}
	}

	{
		exp := "firewall=one"
		got := primary.GetPath()
		if got != exp {
			t.Errorf("error getting primary firewall path, expected %s, got %s", exp, got)
		}
	}

	{
		exp := "firewall=two"
		got := backup.GetPath()
		if got != exp {
			t.Errorf("error getting primary firewall path, expected %s, got %s", exp, got)
		}
	}

	pErr := site.SetPrimaryFirewall(primary)
	bErr := site.SetBackupFirewall(backup)

	if pErr != nil {
		t.Errorf("error setting primary firewall, expected nil, got %v", pErr)
	}
	if bErr != nil {
		t.Errorf("error setting backup firewall, expected nil, got %v", bErr)
	}

	{
		exp := "site=work,firewall=one"
		got := primary.GetPath()
		if got != exp {
			t.Errorf("error getting primary firewall path, expected %s, got %s", exp, got)
		}
	}
	{
		exp := "site=work,firewall=two"
		got := backup.GetPath()
		if got != exp {
			t.Errorf("error getting backup firewall path, expected %s, got %s", exp, got)
		}
	}

}
