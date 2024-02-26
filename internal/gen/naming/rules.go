package naming

import "strings"

var (
	rules = [...]string{
		"CEF", "ID", "DNS",
		"C6A1",
		"C6A1Label",
		"C6A3",
		"C6A3Label",
		"C6A4",
		"C6A4Label",
		"CFP1",
		"CFP1Label",
		"CFP2",
		"CFP2Label",
		"CFP3",
		"CFP3Label",
		"CFP4",
		"CFP4Label",
		"CN1",
		"CN1Label",
		"CN2",
		"CN2Label",
		"CN3",
		"CN3Label",
		"CS1",
		"CS1Label",
		"CS2",
		"CS2Label",
		"CS3",
		"CS3Label",
		"CS4",
		"CS4Label",
		"CS5",
		"CS5Label",
		"CS6",
		"CS6Label",
		// Destination HostName.
		"DHost",
		// Destination Process ID.
		"DPID",
		"DPriv",
		"DProc",
		// Destination Port
		"DPt",
		// Device Time Zone.
		"DTZ",
		// Destination User ID
		"DUID",
		"DUser",
		"DvcHost",
		"DvcMAC",
		"DvcPID",
		"FName",
		"FSize",
		"SHost",
		"SMAC",
		"SNtDom",
		"SPID",
		"SPriv",
		"SProc",
		"SPt",
		"SUID",
		"SUser",
		"AHost",
		// Agent ID.
		"AID",
		"AMAC",
		// Agent Type.
		"AT",
		"ATZ",
		// Agent Version
		"AV",
		"DLat",
		"DLong",
		"SLat",
		"SLong",
	}

	rulesMap = func() (r map[string]string) {
		r = make(map[string]string)
		for _, v := range rules {
			r[strings.ToLower(v)] = v
		}
		return r
	}()
)

func Rule(part string) (string, bool) {
	v, ok := rulesMap[strings.ToLower(part)]
	return v, ok
}
