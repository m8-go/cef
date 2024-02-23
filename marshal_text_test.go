package cef_test

import (
	"bytes"
	"testing"

	"go.m8.ru/cef"
)

func TestCEF_MarshalText(t *testing.T) {
	t.Parallel()

	cefVersion := 1
	deviceVendor := "Security"
	deviceProduct := "threatmanager"
	deviceVersion := "1.0"
	deviceEventClassID := "100"
	agentSeverity := cef.AgentSeverity(10)

	for _, tc := range []struct {
		name               string
		cefVersion         int
		deviceVendor       string
		deviceProduct      string
		deviceVersion      string
		deviceEventClassID string
		agentSeverity      cef.AgentSeverity
		wantText           string
	}{
		{
			name:               "negative cef version",
			cefVersion:         -1,
			deviceVendor:       deviceVendor,
			deviceProduct:      deviceProduct,
			deviceVersion:      deviceVersion,
			deviceEventClassID: deviceEventClassID,
			agentSeverity:      agentSeverity,
			wantText:           "CEF:0|Security|threatmanager|1.0|100|worm successfully stopped|10|",
		},
		{
			name:               "cef version greater than 1",
			cefVersion:         2,
			deviceVendor:       deviceVendor,
			deviceProduct:      deviceProduct,
			deviceVersion:      deviceVersion,
			deviceEventClassID: deviceEventClassID,
			agentSeverity:      agentSeverity,
			wantText:           "CEF:1|Security|threatmanager|1.0|100|worm successfully stopped|10|",
		},
		{
			name:               "negative agent severity",
			cefVersion:         cefVersion,
			deviceVendor:       deviceVendor,
			deviceProduct:      deviceProduct,
			deviceVersion:      deviceVersion,
			deviceEventClassID: deviceEventClassID,
			agentSeverity:      -1,
			wantText:           "CEF:1|Security|threatmanager|1.0|100|worm successfully stopped|0|",
		},
		{
			name:               "agent severity greater than 10",
			cefVersion:         cefVersion,
			deviceVendor:       deviceVendor,
			deviceProduct:      deviceProduct,
			deviceVersion:      deviceVersion,
			deviceEventClassID: deviceEventClassID,
			agentSeverity:      11,
			wantText:           "CEF:1|Security|threatmanager|1.0|100|worm successfully stopped|10|",
		},
		{
			name:               "device vendor with backslashes",
			cefVersion:         cefVersion,
			deviceVendor:       "Security\\",
			deviceProduct:      deviceProduct,
			deviceVersion:      deviceVersion,
			deviceEventClassID: deviceEventClassID,
			agentSeverity:      agentSeverity,
			wantText:           "CEF:1|Security\\\\|threatmanager|1.0|100|worm successfully stopped|10|",
		},
		{
			name:               "device vendor with vertical bar",
			cefVersion:         cefVersion,
			deviceVendor:       "Security|",
			deviceProduct:      deviceProduct,
			deviceVersion:      deviceVersion,
			deviceEventClassID: deviceEventClassID,
			agentSeverity:      agentSeverity,
			wantText:           "CEF:1|Security\\||threatmanager|1.0|100|worm successfully stopped|10|",
		},
		{
			name:               "ok",
			cefVersion:         cefVersion,
			deviceVendor:       deviceVendor,
			deviceProduct:      deviceProduct,
			deviceVersion:      deviceVersion,
			deviceEventClassID: deviceEventClassID,
			agentSeverity:      agentSeverity,
			wantText:           "CEF:1|Security|threatmanager|1.0|100|worm successfully stopped|10|",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			log := cef.New()
			log.SetCEFVersion(tc.cefVersion)
			log.SetDeviceVendor(tc.deviceVendor)
			log.SetDeviceProduct(tc.deviceProduct)
			log.SetDeviceVersion(tc.deviceVersion)
			log.SetDeviceEventClassID(tc.deviceEventClassID)
			log.SetName("worm successfully stopped")
			log.SetAgentSeverity(tc.agentSeverity)

			t.Log(log.DeviceProduct())

			gotText, _ := log.MarshalText()
			if !bytes.Equal(gotText, []byte(tc.wantText)) {
				t.Fatalf("got text %q, want %q", gotText, tc.wantText)
			}
		})
	}
}
