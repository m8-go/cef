package cef_test

import (
	"fmt"
	"net"

	"go.m8.ru/cef"
)

func ExampleCEF() {
	const veryHigh cef.AgentSeverity = 10

	log := cef.New()
	log.SetCEFVersion(0)
	log.SetDeviceVendor("Security")
	log.SetDeviceProduct("threatmanager")
	log.SetDeviceVersion("1.0")
	log.SetDeviceEventClassID("100")
	log.SetName("worm successfully stopped")
	log.SetAgentSeverity(veryHigh)
	log.SetSrc(net.ParseIP("10.0.0.1"))
	log.SetDst(net.ParseIP("2.1.2.2"))
	log.SetSpt(1232)

	text, err := log.MarshalText()

	fmt.Printf("%s\n", text)
	fmt.Println(err)
	// Output:
	// CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|10|dst=2.1.2.2 spt=1232 src=10.0.0.1
	// <nil>
}

func ExampleCEF_UnmarshalText() {
	log := cef.New()

	text := "CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|10|dst=2.1.2.2 spt=1232 src=10.0.0.1"

	err := log.UnmarshalText([]byte(text))

	fmt.Println(log.CEFVersion())
	fmt.Println(log.DeviceVendor())
	fmt.Println(log.DeviceProduct())
	fmt.Println(log.DeviceVersion())
	fmt.Println(log.DeviceEventClassID())
	fmt.Println(log.Name())
	fmt.Println(log.AgentSeverity())
	fmt.Printf("dst=%v spt=%v src=%v\n", log.Dst(), log.Spt(), log.Src())
	fmt.Println(err)
	// Output:
	// 0
	// Security
	// threatmanager
	// 1.0
	// 100
	// worm successfully stopped
	// Very-High
	// dst=2.1.2.2 spt=1232 src=10.0.0.1
	// <nil>
}
