package netfw

import (
	"testing"
)

func TestDevice(t *testing.T) {
	parent := &device{
		name: "parent",
	}
	child := &device{
		name: "child",
	}
	child.setParent(parent)
	if child.parent != parent {
		t.Errorf("error setting parent twice; child.parent & parent don't match")
	}

}
