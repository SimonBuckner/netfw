package main

import "testing"

func TestNewFirewall(t *testing.T) {
	{
		test := NewFirewall("Test 1")
		if test.Name != "Test 1" {
			t.Errorf("Expected firewall called 'Test 1' but got %v", test.Name)
		}
	}
}
