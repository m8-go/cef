package cef_test

import (
	"fmt"
	"net"

	"go.m8.ru/cef"
)

func ExampleCEF() {
	const veryHigh cef.AgentSeverity = 10

	c := new(cef.CEF)
	c.SetCEFVersion(0)
	c.SetDeviceVendor("Security")
	c.SetDeviceProduct("threatmanager")
	c.SetDeviceVersion("1.0")
	c.SetDeviceEventClassID("100")
	c.SetName("worm successfully stopped")
	c.SetAgentSeverity(veryHigh)
	c.SetSrc(net.ParseIP("10.0.0.1"))
	c.SetDst(net.ParseIP("2.1.2.2"))
	c.SetSpt(1232)

	text, err := c.MarshalText()

	fmt.Printf("%s\n", text)
	fmt.Println(err)
	// Output:
	// CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|10|dst=2.1.2.2 spt=1232 src=10.0.0.1
	// <nil>
}
