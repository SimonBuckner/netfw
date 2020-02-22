package netfw

import (
	"testing"
)

var _ Device = &device{}

func TestDeviceSetParent(t *testing.T) {
	p := newDevice("parent", SiteClass)
	c := newDevice("child", FirewallClass)
	{
		c.setParent(p)
		if c.parent != p {
			t.Errorf("error getting parent")
		}
	}
}

func TestDeviceGetParent(t *testing.T) {
	p := newDevice("parent", SiteClass)
	c := newDevice("child", FirewallClass)
	c.setParent(p)
	if c.getParent() != p {
		t.Errorf("error getting parent")
	}
}

func TestDeviceGetPath(t *testing.T) {

	parent := newDevice("parent", SiteClass)
	child := newDevice("child", FirewallClass)

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

func TestDeviceSetChild(t *testing.T) {
	parent := newDevice("parent", SiteClass)
	child := newDevice("child", FirewallClass)
	childKey := child.getKey()
	{
		parent.addChild(child)
		if parent.children[childKey] != child {
			t.Errorf("error adding child to device; child is not correct")
		}
		if child.parent != parent {
			t.Errorf("error adding child to device; parent is not correct")
		}
	}
}

func TestDeviceChildExists(t *testing.T) {
	parent := newDevice("parent", SiteClass)
	child := newDevice("child", FirewallClass)
	parent.addChild(child)
	{
		exists := parent.hasChild(child)
		if !exists {
			t.Errorf("error checking child exists; expecting true, got false")
		}
	}
}

func TestDeviceHasChildren(t *testing.T) {
	parent := newDevice("parent", SiteClass)
	child := newDevice("child", FirewallClass)

	{
		if parent.hasChildren() == true {
			t.Errorf("error checking children exist; expecting false, got true")
		}
	}

	parent.addChild(child)
	{
		if parent.hasChildren() == false {
			t.Errorf("error checking children exist; expecting true, got false")
		}
	}
}

func TestDeviceGetChildren(t *testing.T) {
	parent := newDevice("parent", SiteClass)
	child1 := newDevice("child1", FirewallClass)
	child2 := newDevice("child2", FirewallClass)
	{
		parent.addChild(child1)
		if len(parent.getChildren()) != 1 {
			t.Errorf("error getting children; expecting 1, got %v", len(parent.getChildren()))
		}
		c1 := parent.getChildren()
		if c1[0] != child1 {
			t.Errorf("error getting children; expected %v, got %v", child1, c1)
		}
	}
	{
		parent.addChild(child2)
		if len(parent.getChildren()) != 2 {
			t.Errorf("error getting children; expecting 2, got %v", len(parent.getChildren()))
		}
	}
}

func TestPatchable(t *testing.T) {
	eth0 := NewIface("eth0")
	eth1 := NewIface("eth1")
	eth2 := NewIface("eth2")

	{
		if eth0.IsPatched() == true {
			t.Errorf("error checking eth0 patching; expecting false, got true")
		}
	}

	{
		err := Patch(eth0, eth1)
		if err != nil {
			t.Errorf("error patching eth0 to eth1; expecting nil, got %v", err)
		}
	}

	{
		err := Patch(eth1, eth2)
		if err == nil {
			t.Errorf("error patching eth1 to eth2; expecting an error, got nil")
		}
	}

	{
		err := Patch(eth2, eth1)
		if err == nil {
			t.Errorf("error patching eth2 to eth1; expecting an error, got nil")
		}
	}

	{
		if eth0.IsPatched() == false {
			t.Errorf("error checking eth0 patching; expecting true, got false")
		}
	}

	{
		if eth1.IsPatched() == false {
			t.Errorf("error checking eth1 patching; expecting true, got false")
		}
	}

	{
		if eth2.IsPatched() == true {
			t.Errorf("error checking eth2 patching; expecting false, got true")
		}
	}

}

func TestGetConfig(t *testing.T) {
	parent := newDevice("parent", SiteClass)
	child1 := newDevice("child1", FirewallClass)
	child2 := newDevice("child2", FirewallClass)

	{
		cfg := parent.GetConfig()
		if len(cfg) != 1 {
			t.Errorf("error getting config; expecting 1 line, got %v", len(cfg))
		}
		if cfg[0] != "site=parent\n" {
			t.Errorf("error getting config; unexpected output, got %v", cfg[0])
		}
	}

	parent.addChild(child1)
	{
		cfg := parent.GetConfig()
		if len(cfg) != 2 {
			t.Errorf("error getting config; expecting 2 line, got %v", len(cfg))
		}
		if cfg[0] != "site=parent\n" {
			t.Errorf("error getting config; unexpected output, got %v", cfg[0])
		}
		if cfg[1] != "site=parent,firewall=child1\n" {
			t.Errorf("error getting config; unexpected output, got %v", cfg[1])
		}
	}

	parent.addChild(child2)
	{
		cfg := parent.GetConfig()
		if len(cfg) != 3 {
			t.Errorf("error getting config; expecting 2 line, got %v", len(cfg))
		}
		if cfg[0] != "site=parent\n" {
			t.Errorf("error getting config; unexpected output, got %v", cfg[0])
		}
		if cfg[1] != "site=parent,firewall=child1\n" {
			t.Errorf("error getting config; unexpected output, got %v", cfg[1])
		}
		if cfg[2] != "site=parent,firewall=child2\n" {
			t.Errorf("error getting config; unexpected output, got %v", cfg[2])
		}
	}

}
