package cef

type CEF struct {
	headerFields
	extensionFields

	err error
}
