package main

import "fmt"

type IPAddr [4]byte

type IPAddrInterface interface{
	String()
}

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string{
	return "hello"
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
