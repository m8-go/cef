package cef

func New() *CEF {
	return &CEF{}
}

type CEF struct {
	headerFields
	extensionFields

	err error
}
