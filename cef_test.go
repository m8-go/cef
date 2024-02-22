package cef_test

import (
	"fmt"

	"go.m8.ru/cef"
)

func ExampleCEF() {
	c := new(cef.CEF)

	fmt.Println(c.AgentSeverity().String())
	// Output:
	// Low
}
