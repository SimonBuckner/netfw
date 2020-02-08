package netfw

import "fmt"

// ShowSite lists out all the devices in the site
func ShowSite(site *Site) {
	fmt.Println(site.GetPath())
	// for _, fw := range site.children {
	// ShowFirewall(fw)
	// }
}

// // ShowFirewall ..
// func ShowFirewall(fw Device) {
// 	fmt.Println(fw.GetPath())
// 	for _, iface := range fw.G .interfaces {
// 		ShowIface(iface)
// 	}
// }

// ShowIface ..
// func ShowIface(iface *Iface) {
// 	fmt.Println(iface.GetPath())
// }
