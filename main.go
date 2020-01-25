package main

import "strconv"

import "sort"

func main() {

	s := NewSwitch(14, OddsEvensPorts)

	s.SetPortUse(7, LinkPort)
	for i := 1; i <= 6; i++ {
		s.SetPortUse(i, EdgePort)
	}
	s.SetPortUse(9, LinkPort)
	s.SetPortUse(14, LinkPort)

	s.DumpSwitch()

	// s.SetPortUse(14, LinkPort)
	// for i := 9; i <= 13; i++ {
	// 	s.SetPortUse(i, EdgePort)
	// }
	// for i := 13; i <= 14; i++ {
	// 	s.SetPortUse(i, LinkPort)
	// }

}

// intArrayToText turns a list of int to a text description where ranges are
// separated by '-', so '1,2,3,4' would become '1-4' and individual numbers are
// separated by ',', so '1,2,3,4,7,9' would become '1-4,7,9'
func intArrayToText(array []int) string {
	sort.Ints(array)
	text := strconv.Itoa(array[0])
	index := 1
	run := 0
	for {
		if index == len(array) {
			if run > 0 {
				text = text + "-" + strconv.Itoa(array[index-1])
			}
			return text
		}
		if array[index]-1 == array[index-1] {
			run++
			index++
			continue
		}
		if array[index] == array[index-1] {
			// Remove duplicates
			index++
			continue
		}
		if run > 0 {
			text = text + "-" + strconv.Itoa(array[index-1])
			run = 0
		}
		text = text + "," + strconv.Itoa(array[index])
		index++
	}
}
