package netfw

import "testing"

var _ Device = &Site{}

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

func TestSiteAddSwitch(t *testing.T) {
	site1 := NewSite("site1")
	sw1 := NewSwitch("sw1")
	site2 := NewSite("site2")
	sw2 := NewSwitch("sw2")
	site2.AddSwitch(sw2)

	{
		err := site1.AddSwitch(nil)
		if err == nil {
			t.Errorf("error adding nil switch, expected error, got nil")
		}
	}

	{
		err := site1.AddSwitch(sw1)
		if err != nil {
			t.Errorf("error adding switch, expected nil, got %v", err)
		}
	}

	{
		err := site1.AddSwitch(sw1)
		if err == nil {
			t.Errorf("error adding switch, expected error, got nil")
		}
	}

	{
		err := site1.AddSwitch(sw2)
		if err == nil {
			t.Errorf("error adding switch, expected error, got nil")
		}
	}

}
