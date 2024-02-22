// Code generated by go.m8.ru/cef, DO NOT EDIT.

package cef

import (
	"bytes"
	"errors"
	"net"
	"strconv"
	"strings"
)

var ErrNumHeaderFields = errors.New("number of header fields less than 7")

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
	cef.SetDeviceEventClassId(deviceEventClassID)
	name := string(ss[5])
	cef.SetName(name)
	agentSeverity, err := strconv.Atoi(string(ss[6]))
	if err != nil {
		return err
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

func (cef *CEF) set(key string, val string) error {
	switch key {
	case "act":
		cef.SetAct(val)

	case "app":
		cef.SetApp(val)

	case "c6a1":
		cef.SetC6a1(net.ParseIP(val))

	case "c6a1Label":
		cef.SetC6a1Label(val)

	case "c6a3":
		cef.SetC6a3(net.ParseIP(val))

	case "c6a3Label":
		cef.SetC6a3Label(val)

	case "c6a4":
		cef.SetC6a4(net.ParseIP(val))

	case "c6a4Label":
		cef.SetC6a4Label(val)

	case "cat":
		cef.SetCat(val)

	case "cfp1":
		cfp1, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		cef.SetCfp1(cfp1)

	case "cfp1Label":
		cef.SetCfp1Label(val)

	case "cfp2":
		cfp2, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		cef.SetCfp2(cfp2)

	case "cfp2Label":
		cef.SetCfp2Label(val)

	case "cfp3":
		cfp3, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		cef.SetCfp3(cfp3)

	case "cfp3Label":
		cef.SetCfp3Label(val)

	case "cfp4":
		cfp4, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		cef.SetCfp4(cfp4)

	case "cfp4Label":
		cef.SetCfp4Label(val)

	case "cn1":
		cn1, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		cef.SetCn1(cn1)

	case "cn1Label":
		cef.SetCn1Label(val)

	case "cn2":
		cn2, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		cef.SetCn2(cn2)

	case "cn2Label":
		cef.SetCn2Label(val)

	case "cn3":
		cn3, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		cef.SetCn3(cn3)

	case "cn3Label":
		cef.SetCn3Label(val)

	case "cnt":
		cnt, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		cef.SetCnt(cnt)

	case "cs1":
		cef.SetCs1(val)

	case "cs1Label":
		cef.SetCs1Label(val)

	case "cs2":
		cef.SetCs2(val)

	case "cs2Label":
		cef.SetCs2Label(val)

	case "cs3":
		cef.SetCs3(val)

	case "cs3Label":
		cef.SetCs3Label(val)

	case "cs4":
		cef.SetCs4(val)

	case "cs4Label":
		cef.SetCs4Label(val)

	case "cs5":
		cef.SetCs5(val)

	case "cs5Label":
		cef.SetCs5Label(val)

	case "cs6":
		cef.SetCs6(val)

	case "cs6Label":
		cef.SetCs6Label(val)

	case "destinationDnsDomain":
		cef.SetDestinationDnsDomain(val)

	case "destinationServiceName":
		cef.SetDestinationServiceName(val)

	case "destinationTranslatedAddress":
		cef.SetDestinationTranslatedAddress(net.ParseIP(val))

	case "destinationTranslatedPort":
		destinationTranslatedPort, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		cef.SetDestinationTranslatedPort(destinationTranslatedPort)

	case "deviceCustomDate1":
		cef.SetDeviceCustomDate1(val)

	case "deviceCustomDate1Label":
		cef.SetDeviceCustomDate1Label(val)

	case "deviceCustomDate2":
		cef.SetDeviceCustomDate2(val)

	default:
	}
	return nil
}