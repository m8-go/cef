package cef

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func (cef *CEF) MarshalText() (text []byte, err error) {
	fields := []string{
		fmt.Sprintf("CEF:%d", cef.CEFVersion()),
		cef.DeviceVendor(),
		cef.DeviceProduct(),
		cef.DeviceVersion(),
		cef.DeviceEventClassId(),
		cef.Name(),
		strconv.Itoa(int(cef.AgentSeverity())),
		cef.extension(),
	}

	b := new(bytes.Buffer)
	b.WriteString(strings.Join(fields, "|"))

	return b.Bytes(), nil
}
