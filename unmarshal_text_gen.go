// Code generated by go.m8.ru/cef, DO NOT EDIT.

package cef

import (
	"errors"
	"net"
	"strconv"
)

func (cef *CEF) set(key string, val string) error {
	switch key {
	case "act":
		cef.SetAct(val)

	case "app":
		cef.SetApp(val)

	case "c6a1":
		cef.SetC6A1(net.ParseIP(val))

	case "c6a1Label":
		cef.SetC6A1Label(val)

	case "c6a3":
		cef.SetC6A3(net.ParseIP(val))

	case "c6a3Label":
		cef.SetC6A3Label(val)

	case "c6a4":
		cef.SetC6A4(net.ParseIP(val))

	case "c6a4Label":
		cef.SetC6A4Label(val)

	case "cat":
		cef.SetCat(val)

	case "cfp1":
		cfp1, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetCFP1(float32(cfp1))

	case "cfp1Label":
		cef.SetCFP1Label(val)

	case "cfp2":
		cfp2, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetCFP2(float32(cfp2))

	case "cfp2Label":
		cef.SetCFP2Label(val)

	case "cfp3":
		cfp3, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetCFP3(float32(cfp3))

	case "cfp3Label":
		cef.SetCFP3Label(val)

	case "cfp4":
		cfp4, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetCFP4(float32(cfp4))

	case "cfp4Label":
		cef.SetCFP4Label(val)

	case "cn1":
		CN1, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetCN1(CN1)

	case "cn1Label":
		cef.SetCN1Label(val)

	case "cn2":
		CN2, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetCN2(CN2)

	case "cn2Label":
		cef.SetCN2Label(val)

	case "cn3":
		CN3, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetCN3(CN3)

	case "cn3Label":
		cef.SetCN3Label(val)

	case "cnt":
		cnt, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetCnt(cnt)

	case "cs1":
		cef.SetCS1(val)

	case "cs1Label":
		cef.SetCS1Label(val)

	case "cs2":
		cef.SetCS2(val)

	case "cs2Label":
		cef.SetCS2Label(val)

	case "cs3":
		cef.SetCS3(val)

	case "cs3Label":
		cef.SetCS3Label(val)

	case "cs4":
		cef.SetCS4(val)

	case "cs4Label":
		cef.SetCS4Label(val)

	case "cs5":
		cef.SetCS5(val)

	case "cs5Label":
		cef.SetCS5Label(val)

	case "cs6":
		cef.SetCS6(val)

	case "cs6Label":
		cef.SetCS6Label(val)

	case "destinationDnsDomain":
		cef.SetDestinationDNSDomain(val)

	case "destinationServiceName":
		cef.SetDestinationServiceName(val)

	case "destinationTranslatedAddress":
		cef.SetDestinationTranslatedAddress(net.ParseIP(val))

	case "destinationTranslatedPort":
		destinationTranslatedPort, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDestinationTranslatedPort(destinationTranslatedPort)

	case "deviceCustomDate1":
		cef.SetDeviceCustomDate1(val)

	case "deviceCustomDate1Label":
		cef.SetDeviceCustomDate1Label(val)

	case "deviceCustomDate2":
		cef.SetDeviceCustomDate2(val)

	case "deviceCustomDate2Label":
		cef.SetDeviceCustomDate2Label(val)

	case "deviceDirection":
		deviceDirection, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDeviceDirection(deviceDirection)

	case "deviceDnsDomain":
		cef.SetDeviceDNSDomain(val)

	case "deviceExternalId":
		cef.SetDeviceExternalID(val)

	case "deviceFacility":
		cef.SetDeviceFacility(val)

	case "deviceInboundInterface":
		cef.SetDeviceInboundInterface(val)

	case "deviceNtDomain":
		cef.SetDeviceNtDomain(val)

	case "deviceOutboundInterface":
		cef.SetDeviceOutboundInterface(val)

	case "DevicePayloadId":
		cef.SetDevicePayloadID(val)

	case "deviceProcessName":
		cef.SetDeviceProcessName(val)

	case "deviceTranslatedAddress":
		cef.SetDeviceTranslatedAddress(net.ParseIP(val))

	case "dhost":
		cef.SetDHost(val)

	case "dntdom":
		cef.SetDntdom(val)

	case "dpid":
		dpid, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDPID(dpid)

	case "dpriv":
		cef.SetDPriv(val)

	case "dproc":
		cef.SetDProc(val)

	case "dpt":
		dpt, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDPt(dpt)

	case "dst":
		cef.SetDst(net.ParseIP(val))

	case "dtz":
		cef.SetDTZ(val)

	case "duid":
		cef.SetDUID(val)

	case "duser":
		cef.SetDUser(val)

	case "dvc":
		cef.SetDvc(net.ParseIP(val))

	case "dvchost":
		cef.SetDvcHost(val)

	case "dvcmac":
		dvcmac, err := net.ParseMAC(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDvcMAC(dvcmac)

	case "dvcpid":
		dvcpid, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDvcPID(dvcpid)

	case "end":
		cef.SetEnd(val)

	case "externalId":
		cef.SetExternalID(val)

	case "fileCreateTime":
		cef.SetFileCreateTime(val)

	case "fileHash":
		cef.SetFileHash(val)

	case "fileId":
		cef.SetFileID(val)

	case "fileModificationTime":
		cef.SetFileModificationTime(val)

	case "filePath":
		cef.SetFilePath(val)

	case "filePermission":
		cef.SetFilePermission(val)

	case "fileType":
		cef.SetFileType(val)

	case "flexDate1":
		cef.SetFlexDate1(val)

	case "flexDate1Label":
		cef.SetFlexDate1Label(val)

	case "flexString1":
		cef.SetFlexString1(val)

	case "flexString1Label":
		cef.SetFlexString1Label(val)

	case "flexString2":
		cef.SetFlexString2(val)

	case "flexString2Label":
		cef.SetFlexString2Label(val)

	case "fname":
		cef.SetFName(val)

	case "fsize":
		FSize, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetFSize(FSize)

	case "in":
		in, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetIn(in)

	case "msg":
		cef.SetMsg(val)

	case "oldFileCreateTime":
		cef.SetOldFileCreateTime(val)

	case "oldFileHash":
		cef.SetOldFileHash(val)

	case "oldFileId":
		cef.SetOldFileID(val)

	case "oldFileModificationTime":
		cef.SetOldFileModificationTime(val)

	case "oldFileName":
		cef.SetOldFileName(val)

	case "oldFilePath":
		cef.SetOldFilePath(val)

	case "oldFilePermission":
		cef.SetOldFilePermission(val)

	case "oldFileSize":
		oldFileSize, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetOldFileSize(oldFileSize)

	case "oldFileType":
		cef.SetOldFileType(val)

	case "out":
		out, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetOut(out)

	case "outcome":
		cef.SetOutcome(val)

	case "proto":
		cef.SetProto(val)

	case "reason":
		cef.SetReason(val)

	case "request":
		cef.SetRequest(val)

	case "requestClientApplication":
		cef.SetRequestClientApplication(val)

	case "requestContext":
		cef.SetRequestContext(val)

	case "requestCookies":
		cef.SetRequestCookies(val)

	case "requestMethod":
		cef.SetRequestMethod(val)

	case "rt":
		cef.SetRt(val)

	case "shost":
		cef.SetSHost(val)

	case "smac":
		smac, err := net.ParseMAC(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetSMAC(smac)

	case "sntdom":
		cef.SetSNtDom(val)

	case "sourceDnsDomain":
		cef.SetSourceDNSDomain(val)

	case "sourceServiceName":
		cef.SetSourceServiceName(val)

	case "sourceTranslatedAddress":
		cef.SetSourceTranslatedAddress(net.ParseIP(val))

	case "sourceTranslatedPort":
		sourceTranslatedPort, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetSourceTranslatedPort(sourceTranslatedPort)

	case "spid":
		spid, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetSPID(spid)

	case "spriv":
		cef.SetSPriv(val)

	case "sproc":
		cef.SetSProc(val)

	case "spt":
		spt, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetSPt(spt)

	case "src":
		cef.SetSrc(net.ParseIP(val))

	case "start":
		cef.SetStart(val)

	case "suid":
		cef.SetSUID(val)

	case "suser":
		cef.SetSUser(val)

	case "type":
		typ, err := strconv.Atoi(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetType(typ)

	case "agentDnsDomain":
		cef.SetAgentDNSDomain(val)

	case "agentNtDomain":
		cef.SetAgentNtDomain(val)

	case "agentTranslatedAddress":
		cef.SetAgentTranslatedAddress(net.ParseIP(val))

	case "agentTranslatedZoneExternalID":
		cef.SetAgentTranslatedZoneExternalID(val)

	case "agentTranslatedZoneURI":
		cef.SetAgentTranslatedZoneURI(val)

	case "agentZoneExternalID":
		cef.SetAgentZoneExternalID(val)

	case "agentZoneURI":
		cef.SetAgentZoneURI(val)

	case "agt":
		cef.SetAgt(net.ParseIP(val))

	case "ahost":
		cef.SetAHost(val)

	case "aid":
		cef.SetAID(val)

	case "amac":
		amac, err := net.ParseMAC(val)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetAMAC(amac)

	case "art":
		cef.SetArt(val)

	case "at":
		cef.SetAT(val)

	case "atz":
		cef.SetATZ(val)

	case "av":
		cef.SetAV(val)

	case "customerExternalID":
		cef.SetCustomerExternalID(val)

	case "customerURI":
		cef.SetCustomerURI(val)

	case "destinatioTranslatedZoneExternalID":
		cef.SetDestinatioTranslatedZoneExternalID(val)

	case "destinationTranslatedZoneURI":
		cef.SetDestinationTranslatedZoneURI(val)

	case "destinationZoneExternalID":
		cef.SetDestinationZoneExternalID(val)

	case "destinationZoneURI":
		cef.SetDestinationZoneURI(val)

	case "deviceTranslatedZoneExternalID":
		cef.SetDeviceTranslatedZoneExternalID(val)

	case "deviceTranslatedZoneURI":
		cef.SetDeviceTranslatedZoneURI(val)

	case "deviceZoneExternalID":
		cef.SetDeviceZoneExternalID(val)

	case "deviceZoneURI":
		cef.SetDeviceZoneURI(val)

	case "dlat":
		dlat, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDLat(dlat)

	case "dlong":
		dlong, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDLong(dlong)

	case "eventId":
		eventID, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetEventID(eventID)

	case "rawEvent":
		cef.SetRawEvent(val)

	case "slat":
		slat, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetSLat(slat)

	case "slong":
		slong, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetSLong(slong)

	case "sourceTranslatedZoneExternalID":
		cef.SetSourceTranslatedZoneExternalID(val)

	case "sourceTranslatedZoneURI":
		cef.SetSourceTranslatedZoneURI(val)

	case "sourceZoneExternalID":
		cef.SetSourceZoneExternalID(val)

	case "sourceZoneURI":
		cef.SetSourceZoneURI(val)

	case "agentTranslatedZoneKey":
		agentTranslatedZoneKey, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetAgentTranslatedZoneKey(agentTranslatedZoneKey)

	case "agentZoneKey":
		agentZoneKey, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetAgentZoneKey(agentZoneKey)

	case "customerKey":
		customerKey, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetCustomerKey(customerKey)

	case "destinationTranslatedZoneKey":
		destinationTranslatedZoneKey, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDestinationTranslatedZoneKey(destinationTranslatedZoneKey)

	case "dZoneKey":
		dZoneKey, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDZoneKey(dZoneKey)

	case "deviceTranslatedZoneKey":
		deviceTranslatedZoneKey, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDeviceTranslatedZoneKey(deviceTranslatedZoneKey)

	case "deviceZoneKey":
		deviceZoneKey, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetDeviceZoneKey(deviceZoneKey)

	case "sTranslatedZoneKey":
		sTranslatedZoneKey, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetSTranslatedZoneKey(sTranslatedZoneKey)

	case "sZoneKey":
		sZoneKey, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.Join(err, ErrBadExtension)
		}

		cef.SetSZoneKey(sZoneKey)

	default:
	}

	return nil
}
