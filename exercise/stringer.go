package exercise

import "fmt"

type IPAddr [4]byte

func (ipAddr *IPAddr) String() string {
	a := uint8(ipAddr[0])
	b := uint8(ipAddr[1])
	c := uint8(ipAddr[2])
	d := uint8(ipAddr[3])

	return fmt.Sprintf("%v.%v.%v.%v", a, b, c, d)
}

func StringerMain() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
