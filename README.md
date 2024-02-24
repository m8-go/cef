# CEF

CEF builder generated from CEF implementation standard.

## Usage

```shell
go get go.m8.ru/cef
```

```go
package main

import (
	"fmt"
	"net"

	"go.m8.ru/cef"
)

func main() {
	l1 := new(cef.CEF)
	l1.SetCEFVersion(0)
	l1.SetDeviceVendor("Security")
	l1.SetDeviceProduct("threatmanager")
	l1.SetDeviceVersion("1.0")
	l1.SetDeviceEventClassID("100")
	l1.SetName("worm successfully stopped")
	l1.SetAgentSeverity(cef.AgentSeverityVeryHigh10)
	l1.SetSrc(net.ParseIP("10.0.0.1"))
	l1.SetDst(net.ParseIP("2.1.2.2"))
	l1.SetSpt(1232)

	text, err1 := l1.MarshalText()

	fmt.Printf("%s\n", text)
	fmt.Println(err1)

	l2 := new(cef.CEF)
	err2 := l2.UnmarshalText(text)

	fmt.Println(l2.CEFVersion())
	fmt.Println(l2.DeviceVendor())
	fmt.Println(l2.DeviceProduct())
	fmt.Println(l2.DeviceVersion())
	fmt.Println(l2.DeviceEventClassID())
	fmt.Println(l2.Name())
	fmt.Println(l2.AgentSeverity())
	fmt.Printf("dst=%v spt=%v src=%v\n", l2.Dst(), l2.Spt(), l2.Src())
	fmt.Println(err2)
}
```

See [example](example_test.go) for more info.

## Reference

- https://www.microfocus.com/documentation/arcsight/arcsight-smartconnectors-8.4/pdfdoc/cef-implementation-standard/cef-implementation-standard.pdf
