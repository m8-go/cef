package cef_test

import (
	"testing"

	"go.m8.ru/cef"
)

func TestCEF_UnmarshalText(t *testing.T) {
	c := new(cef.CEF)
	c.UnmarshalText([]byte("CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|10|c6a1=10.0.0.1 dst=2.1.2.2 spt=1232"))
	t.Log(c.C6a1().To16())
}
