package cef

type CEF struct {
	headerFields
	extensionFields

	err error
}

const (
	AgentSeverityLow0 AgentSeverity = iota
	AgentSeverityLow1
	AgentSeverityLow2
	AgentSeverityLow3
	AgentSeverityMedium4
	AgentSeverityMedium5
	AgentSeverityMedium6
	AgentSeverityHigh7
	AgentSeverityHigh8
	AgentSeverityVeryHigh9
	AgentSeverityVeryHigh10
)

// AgentSeverity is a string or integer and it reflects the importance of the event.
//
//	The valid string values are: Unknown, Low, Medium, High, and Very-High.
//	The valid integer values are: 0-3=Low, 4-6=Medium, 7- 8=High, and 9- 10=Very-High
type AgentSeverity int

func (s AgentSeverity) String() string {
	switch {
	case s >= 0 && s <= 3:
		return "Low"

	case s >= 4 && s <= 6:
		return "Medium"

	case s >= 7 && s <= 8:
		return "High"

	case s >= 9 && s <= 10:
		return "Very-High"

	default:
		return "Unknown"
	}
}
