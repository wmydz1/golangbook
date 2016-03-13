package main

import (
	"fmt"
)

type IPAddr [4]byte

func main() {
	adders := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range adders {
		fmt.Printf("%v :%v\n", n, a)
	}
}
