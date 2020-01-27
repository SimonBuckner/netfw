package main

import "strconv"

import "sort"

func main() {

	fw := NewFirewall("Islington")
	wan0 := NewInterface("WAN0", WANZone)
	fw.AddInterface(wan0)
	mad0 := NewInterface("MAD0", MADZone)
	fw.AddInterface(mad0)
	lan0 := NewInterface("LAN0", LANZone)
	fw.AddInterface(lan0)

	sw1 := NewSwitch(14, OddsEvensPorts)
	for i := 1; i <= 12; i++ {
		sw1.SetPortUse(i, EdgePort)
	}
	sw1.SetPortUse(13, LinkPort)
	sw1.SetPortUse(14, LinkPort)

	fw.Dump()

}

// intArrayToText turns a list of int to a text description where ranges are
// separated by '-', so '1,2,3,4' would become '1-4' and individual numbers are
// separated by ',', so '1,2,3,4,7,9' would become '1-4,7,9'
func intArrayToText(array []int) string {
	if len(array) == 0 {
		return ""
	}

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
