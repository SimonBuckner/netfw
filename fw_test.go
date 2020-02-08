package netfw

import "testing"

func TestNewFirewall(t *testing.T) {
	fw := NewFirewall("work")
	if fw.class != FirewallClass {
		t.Errorf("expecting SiteClass, got %v", fw.class.ToString())
	}
}

func TestFirewallSetSite(t *testing.T) {
	s1 := NewSite("work")
	s2 := NewSite("home")
	fw := NewFirewall("fw")

	{
		err := fw.SetSite(nil)

		if err == nil {
			t.Errorf("expecting error setting site, got nil")
		}
	}
	{
		err := fw.SetSite(s1)
		if err != nil {
			t.Errorf("unexpected error setting site; expected nil, got %v", err)
		}
	}
	{
		err := fw.SetSite(s2)
		if err == nil {
			t.Errorf("error overwritting site; expected error, got nil")
		}
	}

}
