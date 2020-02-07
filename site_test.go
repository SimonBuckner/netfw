package netfw

import "testing"

func TestSiteSetPrimary(t *testing.T) {
	site := NewSite("work")
	primary := NewFirewall("one")
	backup := NewFirewall("two")

	{
		err := site.SetPrimaryFirewall(nil)
		if err == nil {
			t.Errorf("error setting primary nil firewall, expected error, got nil")
		}
	}
	{
		err := site.SetBackupFirewall(nil)
		if err == nil {
			t.Errorf("error setting backup nil firewall, expected error, got nil")
		}
	}

	{
		err := site.SetBackupFirewall(backup)
		if err == nil {
			t.Errorf("error setting backup firewall first, expected error, got nil")
		}
	}
	{
		err := site.SetPrimaryFirewall(primary)
		if err != nil {
			t.Errorf("error setting primary firewall, unexpected error %v", err)
		}
	}
	{
		err := site.SetPrimaryFirewall(primary)
		if err == nil {
			t.Errorf("error setting primary firewall twice, expected error, got nil")
		}
	}
	{
		err := site.SetBackupFirewall(primary)
		if err == nil {
			t.Errorf("error setting primary firewall as backup, expected error, got nil")
		}
	}
	{
		err := site.SetBackupFirewall(backup)
		if err != nil {
			t.Errorf("error setting backup firewall, unexpected error %v", err)
		}
	}
	{
		err := site.SetBackupFirewall(backup)
		if err == nil {
			t.Errorf("error setting backup firewall twice, expected error, got nil")
		}
	}
}
