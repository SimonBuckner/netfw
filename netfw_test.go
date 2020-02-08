package netfw

import (
	"testing"
)

func TestDeviceSetParent(t *testing.T) {
	p := &device{
		name:  "parent",
		class: SiteClass,
	}
	c := &device{
		name:  "child",
		class: FirewallClass,
	}
	{
		c.setParent(p)
		if c.parent != p {
			t.Errorf("error getting parent")
		}
	}
}

func TestDeviceGetParent(t *testing.T) {
	p := &device{
		name:  "parent",
		class: SiteClass,
	}
	c := &device{
		name:  "child",
		class: FirewallClass,
	}
	c.setParent(p)
	if c.getParent() != p {
		t.Errorf("error getting parent")
	}
}

func TestDeviceGetPath(t *testing.T) {

	parent := &device{
		name:  "parent",
		class: SiteClass,
	}
	child := &device{
		name:  "child",
		class: FirewallClass,
	}

	{
		exp := "site=parent"
		got := parent.GetPath()
		if got != exp {
			t.Errorf("error getting firewall path, expected %s, got %s", exp, got)
		}
	}

	{
		exp := "firewall=child"
		got := child.GetPath()
		if got != exp {
			t.Errorf("error getting primary firewall path, expected %s, got %s", exp, got)
		}
	}

	child.setParent(parent)
	{
		exp := "site=parent,firewall=child"
		got := child.GetPath()
		if got != exp {
			t.Errorf("error getting primary firewall path, expected %s, got %s", exp, got)
		}
	}

}
