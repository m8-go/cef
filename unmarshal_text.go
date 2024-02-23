package cef

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

var (
	ErrNumHeaderFields  = errors.New("cef: number of header fields less than 7")
	ErrBadCEFVersion    = errors.New("cef: bad CEF version")
	ErrBadAgentSeverity = errors.New("cef: bad agent severity")
	ErrBadExtension     = errors.New("cef: bad extension")
)

func (cef *CEF) UnmarshalText(text []byte) (err error) {
	defer func() {
		cef.err = err
	}()

	ss := bytes.Split(text, []byte("|"))
	if len(ss) < headerFieldsNum {
		return ErrNumHeaderFields
	}

	hasExtension := bytes.ContainsRune(ss[len(ss)-1], '=')
	if len(ss) == headerFieldsNum && hasExtension {
		return ErrNumHeaderFields
	}

	cefVer, err := mkCEFVersion(ss[0])
	if err != nil {
		return err
	}

	cef.SetCEFVersion(cefVer)

	deviceVendor := string(ss[1])
	cef.SetDeviceVendor(deviceVendor)

	deviceProduct := string(ss[2])
	cef.SetDeviceProduct(deviceProduct)

	deviceVersion := string(ss[3])
	cef.SetDeviceVersion(deviceVersion)

	deviceEventClassID := string(ss[4])
	cef.SetDeviceEventClassID(deviceEventClassID)

	name := string(ss[5])
	cef.SetName(name)

	agentSeverity, err := strconv.Atoi(string(ss[6]))
	if err != nil {
		return errors.Join(err, ErrBadAgentSeverity)
	}

	cef.SetAgentSeverity(AgentSeverity(agentSeverity))

	if len(ss) == headerFieldsNum {
		return nil
	}

	c := mkCollection(ss[7])

	for _, kv := range c {
		ss := strings.Split(kv, "=")

		key := ss[0]
		val := ss[1]

		if serr := cef.set(key, val); serr != nil {
			return serr
		}
	}

	return nil
}

func mkCEFVersion(v []byte) (int, error) {
	ss := bytes.Split(v, []byte(":"))

	if len(ss) != 2 {
		return 0, ErrBadCEFVersion
	}

	ver, err := strconv.Atoi(string(ss[1]))
	if err != nil {
		return 0, errors.Join(err, ErrBadCEFVersion)
	}

	return ver, nil
}

func mkCollection(b []byte) []string {
	var collection []string

	str := string(b)
	ss := strings.Split(str, " ")

	for _, s := range ss {
		if !strings.ContainsRune(s, '=') {
			collection[len(collection)-1] = collection[len(collection)-1] + " " + s

			continue
		}

		collection = append(collection, s)
	}

	return collection
}
