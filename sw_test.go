package main

import "testing"

func TestArrayToText(t *testing.T) {
	{
		test := intArrayToText([]int{1, 2, 3, 4, 5})
		if test != "1-5" {
			t.Errorf("Expecting '1-5' but got '%v'", test)
		}
	}

	{
		test := intArrayToText([]int{-1, 0, 1, 2, 3})
		if test != "-1-3" {
			t.Errorf("Expecting '-1-3' but got '%v'", test)
		}
	}

	{
		test := intArrayToText([]int{1, 3, 5, 7})
		if test != "1,3,5,7" {
			t.Errorf("Expecting '1,3,5,7' but got '%v'", test)
		}
	}

	{
		test := intArrayToText([]int{1, 2, 3, 4, 5, 8, 10})
		if test != "1-5,8,10" {
			t.Errorf("Expecting '1-5,8,10' but got '%v'", test)
		}
	}

	{
		test := intArrayToText([]int{1, 6, 4, 7, 2, 3, 5})
		if test != "1-7" {
			t.Errorf("Expecting '1-7' but got '%v'", test)
		}
	}

	{
		test := intArrayToText([]int{1, 2, 2, 3})
		if test != "1-3" {
			t.Errorf("Expecting '1-3' but got '%v'", test)
		}
	}

	{
		test := intArrayToText([]int{})
		if test != "" {
			t.Errorf("Expecting '' but got '%v'", test)
		}
	}

}
