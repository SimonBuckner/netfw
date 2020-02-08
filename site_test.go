package netfw

import "testing"

func TestNewSites(t *testing.T) {
	site := NewSite("work")
	if site.class != SiteClass {
		t.Errorf("expecting SiteClass, got %v", site.class.ToString())
	}
}
func TestSiteAddFirewall(t *testing.T) {
	site := NewSite("work")
	primary := NewFirewall("one")
	backup := NewFirewall("backup")
	fail := NewFirewall("fail")

	{
		err := site.AddFirewall(nil)

		if err == nil {
			t.Errorf("expecting error setting nil site")
		}
	}
	{
		err := site.AddFirewall(primary)

		if err != nil {
			t.Errorf("error setting primary firewall, expected nil, got %v", err)
		}
	}
	{
		err := site.AddFirewall(primary)

		if err == nil {
			t.Errorf("error setting primary firewall, expected error, got nil")
		}
	}
	{
		err := site.AddFirewall(backup)
		if err != nil {
			t.Errorf("error setting backup firewall; expecting nil, got %v", err)

		}
	}
	{
		err := site.AddFirewall(backup)
		if err == nil {
			t.Errorf("expecting error setting backup firewall tice, got nil")
		}
	}
	{
		fail.setParent(site)
		err := site.AddFirewall(fail)
		if err == nil {
			t.Errorf("error adding firewall already connected to a site; expected error, got nil")
		}
	}
}
