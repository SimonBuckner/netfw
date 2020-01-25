package main

func main() {

	s := NewSwitch(14, OddsEvensPorts)

	s.SetPortUse(7, LinkPort)
	for i := 1; i <= 6; i++ {
		s.SetPortUse(i, EdgePort)
	}
	s.DumpSwitch()

	// s.SetPortUse(14, LinkPort)
	// for i := 9; i <= 13; i++ {
	// 	s.SetPortUse(i, EdgePort)
	// }
	// for i := 13; i <= 14; i++ {
	// 	s.SetPortUse(i, LinkPort)
	// }

}
