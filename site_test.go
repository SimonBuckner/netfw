package netfw

import "testing"

func TestNewSites(t *testing.T) {
	site := NewSite("work")
	if site.class != SiteClass {
		t.Errorf("expecting SiteClass, got %v", site.class.ToString())
	}
}
func TestSiteSetPrimaryFirewall(t *testing.T) {
	site := NewSite("work")
	primary := NewFirewall("one")
	backup := NewFirewall("backup")
	backup.SetSite(site)

	{
		err := site.SetPrimaryFirewall(backup)
		if err == nil {
			t.Errorf("expecting error setting primary firewall that already has a site")

		}
	}
	{
		err := site.SetPrimaryFirewall(nil)

		if err == nil {
			t.Errorf("expecting error setting nil site")
		}
	}
	{
		err := site.SetPrimaryFirewall(primary)

		if err != nil {
			t.Errorf("error setting primary firewall, expected nil, got %v", err)
		}
	}

	{
		err := site.SetPrimaryFirewall(primary)

		if err == nil {
			t.Errorf("error setting primary firewall, expected error, got nil")
		}
	}

}

func TestSiteSetBackupFirewall(t *testing.T) {
	site := NewSite("work")
	backup := NewFirewall("backup")

	{
		err := site.SetBackupFirewall(backup)
		if err == nil {
			t.Errorf("error setting backup firewall with no primary firewall, expected error, got nil")
		}
	}

	primary := NewFirewall("primary")
	site.SetPrimaryFirewall(primary)
	backup2 := NewFirewall("backup")
	backup2.SetSite(site)

	{
		err := site.SetBackupFirewall(nil)

		if err == nil {
			t.Errorf("expecting error setting backup firewall, got nil")
		}
	}
	{
		err := site.SetBackupFirewall(primary)
		if err == nil {
			t.Errorf("expecting error setting same firewall as primary & backup, got nil")
		}
	}

	{
		err := site.SetBackupFirewall(backup2)
		if err == nil {
			t.Errorf("expecting error setting backup firewall that already has a site")

		}
	}
	{
		err := site.SetBackupFirewall(backup)
		if err != nil {
			t.Errorf("error setting backup firewall; expecting nil, got %v", err)

		}
	}
	{
		err := site.SetBackupFirewall(backup)
		if err == nil {
			t.Errorf("expecting error setting backup firewall tice, got nil")
		}
	}
}
