// Code generated by go.m8.ru/cef, DO NOT EDIT.

package cef

import (
	"fmt"
	"strings"
)

func (cef *CEF) extension() string {
	var extension []string

	if act := cef.Act(); act != "" {
		extension = append(extension, fmt.Sprintf("act=%v", act))
	}

	if app := cef.App(); app != "" {
		extension = append(extension, fmt.Sprintf("app=%v", app))
	}

	if c6a1 := cef.C6a1(); len(c6a1) > 0 {
		extension = append(extension, fmt.Sprintf("c6a1=%v", c6a1))
	}

	if c6a1Label := cef.C6a1Label(); c6a1Label != "" {
		extension = append(extension, fmt.Sprintf("c6a1Label=%v", c6a1Label))
	}

	if c6a3 := cef.C6a3(); len(c6a3) > 0 {
		extension = append(extension, fmt.Sprintf("c6a3=%v", c6a3))
	}

	if c6a3Label := cef.C6a3Label(); c6a3Label != "" {
		extension = append(extension, fmt.Sprintf("c6a3Label=%v", c6a3Label))
	}

	if c6a4 := cef.C6a4(); len(c6a4) > 0 {
		extension = append(extension, fmt.Sprintf("c6a4=%v", c6a4))
	}

	if c6a4Label := cef.C6a4Label(); c6a4Label != "" {
		extension = append(extension, fmt.Sprintf("c6a4Label=%v", c6a4Label))
	}

	if cat := cef.Cat(); cat != "" {
		extension = append(extension, fmt.Sprintf("cat=%v", cat))
	}

	if cfp1 := cef.Cfp1(); cfp1 != 0 {
		extension = append(extension, fmt.Sprintf("cfp1=%v", cfp1))
	}

	if cfp1Label := cef.Cfp1Label(); cfp1Label != "" {
		extension = append(extension, fmt.Sprintf("cfp1Label=%v", cfp1Label))
	}

	if cfp2 := cef.Cfp2(); cfp2 != 0 {
		extension = append(extension, fmt.Sprintf("cfp2=%v", cfp2))
	}

	if cfp2Label := cef.Cfp2Label(); cfp2Label != "" {
		extension = append(extension, fmt.Sprintf("cfp2Label=%v", cfp2Label))
	}

	if cfp3 := cef.Cfp3(); cfp3 != 0 {
		extension = append(extension, fmt.Sprintf("cfp3=%v", cfp3))
	}

	if cfp3Label := cef.Cfp3Label(); cfp3Label != "" {
		extension = append(extension, fmt.Sprintf("cfp3Label=%v", cfp3Label))
	}

	if cfp4 := cef.Cfp4(); cfp4 != 0 {
		extension = append(extension, fmt.Sprintf("cfp4=%v", cfp4))
	}

	if cfp4Label := cef.Cfp4Label(); cfp4Label != "" {
		extension = append(extension, fmt.Sprintf("cfp4Label=%v", cfp4Label))
	}

	if cn1 := cef.Cn1(); cn1 != 0 {
		extension = append(extension, fmt.Sprintf("cn1=%v", cn1))
	}

	if cn1Label := cef.Cn1Label(); cn1Label != "" {
		extension = append(extension, fmt.Sprintf("cn1Label=%v", cn1Label))
	}

	if cn2 := cef.Cn2(); cn2 != 0 {
		extension = append(extension, fmt.Sprintf("cn2=%v", cn2))
	}

	if cn2Label := cef.Cn2Label(); cn2Label != "" {
		extension = append(extension, fmt.Sprintf("cn2Label=%v", cn2Label))
	}

	if cn3 := cef.Cn3(); cn3 != 0 {
		extension = append(extension, fmt.Sprintf("cn3=%v", cn3))
	}

	if cn3Label := cef.Cn3Label(); cn3Label != "" {
		extension = append(extension, fmt.Sprintf("cn3Label=%v", cn3Label))
	}

	if cnt := cef.Cnt(); cnt != 0 {
		extension = append(extension, fmt.Sprintf("cnt=%v", cnt))
	}

	if cs1 := cef.Cs1(); cs1 != "" {
		extension = append(extension, fmt.Sprintf("cs1=%v", cs1))
	}

	if cs1Label := cef.Cs1Label(); cs1Label != "" {
		extension = append(extension, fmt.Sprintf("cs1Label=%v", cs1Label))
	}

	if cs2 := cef.Cs2(); cs2 != "" {
		extension = append(extension, fmt.Sprintf("cs2=%v", cs2))
	}

	if cs2Label := cef.Cs2Label(); cs2Label != "" {
		extension = append(extension, fmt.Sprintf("cs2Label=%v", cs2Label))
	}

	if cs3 := cef.Cs3(); cs3 != "" {
		extension = append(extension, fmt.Sprintf("cs3=%v", cs3))
	}

	if cs3Label := cef.Cs3Label(); cs3Label != "" {
		extension = append(extension, fmt.Sprintf("cs3Label=%v", cs3Label))
	}

	if cs4 := cef.Cs4(); cs4 != "" {
		extension = append(extension, fmt.Sprintf("cs4=%v", cs4))
	}

	if cs4Label := cef.Cs4Label(); cs4Label != "" {
		extension = append(extension, fmt.Sprintf("cs4Label=%v", cs4Label))
	}

	if cs5 := cef.Cs5(); cs5 != "" {
		extension = append(extension, fmt.Sprintf("cs5=%v", cs5))
	}

	if cs5Label := cef.Cs5Label(); cs5Label != "" {
		extension = append(extension, fmt.Sprintf("cs5Label=%v", cs5Label))
	}

	if cs6 := cef.Cs6(); cs6 != "" {
		extension = append(extension, fmt.Sprintf("cs6=%v", cs6))
	}

	if cs6Label := cef.Cs6Label(); cs6Label != "" {
		extension = append(extension, fmt.Sprintf("cs6Label=%v", cs6Label))
	}

	if destinationDnsDomain := cef.DestinationDnsDomain(); destinationDnsDomain != "" {
		extension = append(extension, fmt.Sprintf("destinationDnsDomain=%v", destinationDnsDomain))
	}

	if destinationServiceName := cef.DestinationServiceName(); destinationServiceName != "" {
		extension = append(extension, fmt.Sprintf("destinationServiceName=%v", destinationServiceName))
	}

	if destinationTranslatedAddress := cef.DestinationTranslatedAddress(); len(destinationTranslatedAddress) > 0 {
		extension = append(extension, fmt.Sprintf("destinationTranslatedAddress=%v", destinationTranslatedAddress))
	}

	if destinationTranslatedPort := cef.DestinationTranslatedPort(); destinationTranslatedPort != 0 {
		extension = append(extension, fmt.Sprintf("destinationTranslatedPort=%v", destinationTranslatedPort))
	}

	if deviceCustomDate1 := cef.DeviceCustomDate1(); deviceCustomDate1 != "" {
		extension = append(extension, fmt.Sprintf("deviceCustomDate1=%v", deviceCustomDate1))
	}

	if deviceCustomDate1Label := cef.DeviceCustomDate1Label(); deviceCustomDate1Label != "" {
		extension = append(extension, fmt.Sprintf("deviceCustomDate1Label=%v", deviceCustomDate1Label))
	}

	if deviceCustomDate2 := cef.DeviceCustomDate2(); deviceCustomDate2 != "" {
		extension = append(extension, fmt.Sprintf("deviceCustomDate2=%v", deviceCustomDate2))
	}

	if deviceCustomDate2Label := cef.DeviceCustomDate2Label(); deviceCustomDate2Label != "" {
		extension = append(extension, fmt.Sprintf("deviceCustomDate2Label=%v", deviceCustomDate2Label))
	}

	if deviceDirection := cef.DeviceDirection(); deviceDirection != 0 {
		extension = append(extension, fmt.Sprintf("deviceDirection=%v", deviceDirection))
	}

	if deviceDnsDomain := cef.DeviceDnsDomain(); deviceDnsDomain != "" {
		extension = append(extension, fmt.Sprintf("deviceDnsDomain=%v", deviceDnsDomain))
	}

	if deviceExternalId := cef.DeviceExternalId(); deviceExternalId != "" {
		extension = append(extension, fmt.Sprintf("deviceExternalId=%v", deviceExternalId))
	}

	if deviceFacility := cef.DeviceFacility(); deviceFacility != "" {
		extension = append(extension, fmt.Sprintf("deviceFacility=%v", deviceFacility))
	}

	if deviceInboundInterface := cef.DeviceInboundInterface(); deviceInboundInterface != "" {
		extension = append(extension, fmt.Sprintf("deviceInboundInterface=%v", deviceInboundInterface))
	}

	if deviceNtDomain := cef.DeviceNtDomain(); deviceNtDomain != "" {
		extension = append(extension, fmt.Sprintf("deviceNtDomain=%v", deviceNtDomain))
	}

	if deviceOutboundInterface := cef.DeviceOutboundInterface(); deviceOutboundInterface != "" {
		extension = append(extension, fmt.Sprintf("deviceOutboundInterface=%v", deviceOutboundInterface))
	}

	if devicePayloadId := cef.DevicePayloadId(); devicePayloadId != "" {
		extension = append(extension, fmt.Sprintf("DevicePayloadId=%v", devicePayloadId))
	}

	if deviceProcessName := cef.DeviceProcessName(); deviceProcessName != "" {
		extension = append(extension, fmt.Sprintf("deviceProcessName=%v", deviceProcessName))
	}

	if deviceTranslatedAddress := cef.DeviceTranslatedAddress(); len(deviceTranslatedAddress) > 0 {
		extension = append(extension, fmt.Sprintf("deviceTranslatedAddress=%v", deviceTranslatedAddress))
	}

	if dhost := cef.Dhost(); dhost != "" {
		extension = append(extension, fmt.Sprintf("dhost=%v", dhost))
	}

	if dntdom := cef.Dntdom(); dntdom != "" {
		extension = append(extension, fmt.Sprintf("dntdom=%v", dntdom))
	}

	if dpid := cef.Dpid(); dpid != 0 {
		extension = append(extension, fmt.Sprintf("dpid=%v", dpid))
	}

	if dpriv := cef.Dpriv(); dpriv != "" {
		extension = append(extension, fmt.Sprintf("dpriv=%v", dpriv))
	}

	if dproc := cef.Dproc(); dproc != "" {
		extension = append(extension, fmt.Sprintf("dproc=%v", dproc))
	}

	if dpt := cef.Dpt(); dpt != 0 {
		extension = append(extension, fmt.Sprintf("dpt=%v", dpt))
	}

	if dst := cef.Dst(); len(dst) > 0 {
		extension = append(extension, fmt.Sprintf("dst=%v", dst))
	}

	if dtz := cef.Dtz(); dtz != "" {
		extension = append(extension, fmt.Sprintf("dtz=%v", dtz))
	}

	if duid := cef.Duid(); duid != "" {
		extension = append(extension, fmt.Sprintf("duid=%v", duid))
	}

	if duser := cef.Duser(); duser != "" {
		extension = append(extension, fmt.Sprintf("duser=%v", duser))
	}

	if dvc := cef.Dvc(); len(dvc) > 0 {
		extension = append(extension, fmt.Sprintf("dvc=%v", dvc))
	}

	if dvchost := cef.Dvchost(); dvchost != "" {
		extension = append(extension, fmt.Sprintf("dvchost=%v", dvchost))
	}

	if dvcmac := cef.Dvcmac(); len(dvcmac) > 0 {
		extension = append(extension, fmt.Sprintf("dvcmac=%v", dvcmac))
	}

	if dvcpid := cef.Dvcpid(); dvcpid != 0 {
		extension = append(extension, fmt.Sprintf("dvcpid=%v", dvcpid))
	}

	if end := cef.End(); end != "" {
		extension = append(extension, fmt.Sprintf("end=%v", end))
	}

	if externalId := cef.ExternalId(); externalId != "" {
		extension = append(extension, fmt.Sprintf("externalId=%v", externalId))
	}

	if fileCreateTime := cef.FileCreateTime(); fileCreateTime != "" {
		extension = append(extension, fmt.Sprintf("fileCreateTime=%v", fileCreateTime))
	}

	if fileHash := cef.FileHash(); fileHash != "" {
		extension = append(extension, fmt.Sprintf("fileHash=%v", fileHash))
	}

	if fileId := cef.FileId(); fileId != "" {
		extension = append(extension, fmt.Sprintf("fileId=%v", fileId))
	}

	if fileModificationTime := cef.FileModificationTime(); fileModificationTime != "" {
		extension = append(extension, fmt.Sprintf("fileModificationTime=%v", fileModificationTime))
	}

	if filePath := cef.FilePath(); filePath != "" {
		extension = append(extension, fmt.Sprintf("filePath=%v", filePath))
	}

	if filePermission := cef.FilePermission(); filePermission != "" {
		extension = append(extension, fmt.Sprintf("filePermission=%v", filePermission))
	}

	if fileType := cef.FileType(); fileType != "" {
		extension = append(extension, fmt.Sprintf("fileType=%v", fileType))
	}

	if flexDate1 := cef.FlexDate1(); flexDate1 != "" {
		extension = append(extension, fmt.Sprintf("flexDate1=%v", flexDate1))
	}

	if flexDate1Label := cef.FlexDate1Label(); flexDate1Label != "" {
		extension = append(extension, fmt.Sprintf("flexDate1Label=%v", flexDate1Label))
	}

	if flexString1 := cef.FlexString1(); flexString1 != "" {
		extension = append(extension, fmt.Sprintf("flexString1=%v", flexString1))
	}

	if flexString1Label := cef.FlexString1Label(); flexString1Label != "" {
		extension = append(extension, fmt.Sprintf("flexString1Label=%v", flexString1Label))
	}

	if flexString2 := cef.FlexString2(); flexString2 != "" {
		extension = append(extension, fmt.Sprintf("flexString2=%v", flexString2))
	}

	if flexString2Label := cef.FlexString2Label(); flexString2Label != "" {
		extension = append(extension, fmt.Sprintf("flexString2Label=%v", flexString2Label))
	}

	if fname := cef.Fname(); fname != "" {
		extension = append(extension, fmt.Sprintf("fname=%v", fname))
	}

	if fsize := cef.Fsize(); fsize != 0 {
		extension = append(extension, fmt.Sprintf("fsize=%v", fsize))
	}

	if in := cef.In(); in != 0 {
		extension = append(extension, fmt.Sprintf("in=%v", in))
	}

	if msg := cef.Msg(); msg != "" {
		extension = append(extension, fmt.Sprintf("msg=%v", msg))
	}

	if oldFileCreateTime := cef.OldFileCreateTime(); oldFileCreateTime != "" {
		extension = append(extension, fmt.Sprintf("oldFileCreateTime=%v", oldFileCreateTime))
	}

	if oldFileHash := cef.OldFileHash(); oldFileHash != "" {
		extension = append(extension, fmt.Sprintf("oldFileHash=%v", oldFileHash))
	}

	if oldFileId := cef.OldFileId(); oldFileId != "" {
		extension = append(extension, fmt.Sprintf("oldFileId=%v", oldFileId))
	}

	if oldFileModificationTime := cef.OldFileModificationTime(); oldFileModificationTime != "" {
		extension = append(extension, fmt.Sprintf("oldFileModificationTime=%v", oldFileModificationTime))
	}

	if oldFileName := cef.OldFileName(); oldFileName != "" {
		extension = append(extension, fmt.Sprintf("oldFileName=%v", oldFileName))
	}

	if oldFilePath := cef.OldFilePath(); oldFilePath != "" {
		extension = append(extension, fmt.Sprintf("oldFilePath=%v", oldFilePath))
	}

	if oldFilePermission := cef.OldFilePermission(); oldFilePermission != "" {
		extension = append(extension, fmt.Sprintf("oldFilePermission=%v", oldFilePermission))
	}

	if oldFileSize := cef.OldFileSize(); oldFileSize != 0 {
		extension = append(extension, fmt.Sprintf("oldFileSize=%v", oldFileSize))
	}

	if oldFileType := cef.OldFileType(); oldFileType != "" {
		extension = append(extension, fmt.Sprintf("oldFileType=%v", oldFileType))
	}

	if out := cef.Out(); out != 0 {
		extension = append(extension, fmt.Sprintf("out=%v", out))
	}

	if outcome := cef.Outcome(); outcome != "" {
		extension = append(extension, fmt.Sprintf("outcome=%v", outcome))
	}

	if proto := cef.Proto(); proto != "" {
		extension = append(extension, fmt.Sprintf("proto=%v", proto))
	}

	if reason := cef.Reason(); reason != "" {
		extension = append(extension, fmt.Sprintf("reason=%v", reason))
	}

	if request := cef.Request(); request != "" {
		extension = append(extension, fmt.Sprintf("request=%v", request))
	}

	if requestClientApplication := cef.RequestClientApplication(); requestClientApplication != "" {
		extension = append(extension, fmt.Sprintf("requestClientApplication=%v", requestClientApplication))
	}

	if requestContext := cef.RequestContext(); requestContext != "" {
		extension = append(extension, fmt.Sprintf("requestContext=%v", requestContext))
	}

	if requestCookies := cef.RequestCookies(); requestCookies != "" {
		extension = append(extension, fmt.Sprintf("requestCookies=%v", requestCookies))
	}

	if requestMethod := cef.RequestMethod(); requestMethod != "" {
		extension = append(extension, fmt.Sprintf("requestMethod=%v", requestMethod))
	}

	if rt := cef.Rt(); rt != "" {
		extension = append(extension, fmt.Sprintf("rt=%v", rt))
	}

	if shost := cef.Shost(); shost != "" {
		extension = append(extension, fmt.Sprintf("shost=%v", shost))
	}

	if smac := cef.Smac(); len(smac) > 0 {
		extension = append(extension, fmt.Sprintf("smac=%v", smac))
	}

	if sntdom := cef.Sntdom(); sntdom != "" {
		extension = append(extension, fmt.Sprintf("sntdom=%v", sntdom))
	}

	if sourceDnsDomain := cef.SourceDnsDomain(); sourceDnsDomain != "" {
		extension = append(extension, fmt.Sprintf("sourceDnsDomain=%v", sourceDnsDomain))
	}

	if sourceServiceName := cef.SourceServiceName(); sourceServiceName != "" {
		extension = append(extension, fmt.Sprintf("sourceServiceName=%v", sourceServiceName))
	}

	if sourceTranslatedAddress := cef.SourceTranslatedAddress(); len(sourceTranslatedAddress) > 0 {
		extension = append(extension, fmt.Sprintf("sourceTranslatedAddress=%v", sourceTranslatedAddress))
	}

	if sourceTranslatedPort := cef.SourceTranslatedPort(); sourceTranslatedPort != 0 {
		extension = append(extension, fmt.Sprintf("sourceTranslatedPort=%v", sourceTranslatedPort))
	}

	if spid := cef.Spid(); spid != 0 {
		extension = append(extension, fmt.Sprintf("spid=%v", spid))
	}

	if spriv := cef.Spriv(); spriv != "" {
		extension = append(extension, fmt.Sprintf("spriv=%v", spriv))
	}

	if sproc := cef.Sproc(); sproc != "" {
		extension = append(extension, fmt.Sprintf("sproc=%v", sproc))
	}

	if spt := cef.Spt(); spt != 0 {
		extension = append(extension, fmt.Sprintf("spt=%v", spt))
	}

	if src := cef.Src(); len(src) > 0 {
		extension = append(extension, fmt.Sprintf("src=%v", src))
	}

	if start := cef.Start(); start != "" {
		extension = append(extension, fmt.Sprintf("start=%v", start))
	}

	if suid := cef.Suid(); suid != "" {
		extension = append(extension, fmt.Sprintf("suid=%v", suid))
	}

	if suser := cef.Suser(); suser != "" {
		extension = append(extension, fmt.Sprintf("suser=%v", suser))
	}

	if typ := cef.Type(); typ != 0 {
		extension = append(extension, fmt.Sprintf("typ=%v", typ))
	}

	if agentDnsDomain := cef.AgentDnsDomain(); agentDnsDomain != "" {
		extension = append(extension, fmt.Sprintf("agentDnsDomain=%v", agentDnsDomain))
	}

	if agentNtDomain := cef.AgentNtDomain(); agentNtDomain != "" {
		extension = append(extension, fmt.Sprintf("agentNtDomain=%v", agentNtDomain))
	}

	if agentTranslatedAddress := cef.AgentTranslatedAddress(); len(agentTranslatedAddress) > 0 {
		extension = append(extension, fmt.Sprintf("agentTranslatedAddress=%v", agentTranslatedAddress))
	}

	if agentTranslatedZoneExternalID := cef.AgentTranslatedZoneExternalID(); agentTranslatedZoneExternalID != "" {
		extension = append(extension, fmt.Sprintf("agentTranslatedZoneExternalID=%v", agentTranslatedZoneExternalID))
	}

	if agentTranslatedZoneURI := cef.AgentTranslatedZoneURI(); agentTranslatedZoneURI != "" {
		extension = append(extension, fmt.Sprintf("agentTranslatedZoneURI=%v", agentTranslatedZoneURI))
	}

	if agentZoneExternalID := cef.AgentZoneExternalID(); agentZoneExternalID != "" {
		extension = append(extension, fmt.Sprintf("agentZoneExternalID=%v", agentZoneExternalID))
	}

	if agentZoneURI := cef.AgentZoneURI(); agentZoneURI != "" {
		extension = append(extension, fmt.Sprintf("agentZoneURI=%v", agentZoneURI))
	}

	if agt := cef.Agt(); len(agt) > 0 {
		extension = append(extension, fmt.Sprintf("agt=%v", agt))
	}

	if ahost := cef.Ahost(); ahost != "" {
		extension = append(extension, fmt.Sprintf("ahost=%v", ahost))
	}

	if aid := cef.Aid(); aid != "" {
		extension = append(extension, fmt.Sprintf("aid=%v", aid))
	}

	if amac := cef.Amac(); len(amac) > 0 {
		extension = append(extension, fmt.Sprintf("amac=%v", amac))
	}

	if art := cef.Art(); art != "" {
		extension = append(extension, fmt.Sprintf("art=%v", art))
	}

	if at := cef.At(); at != "" {
		extension = append(extension, fmt.Sprintf("at=%v", at))
	}

	if atz := cef.Atz(); atz != "" {
		extension = append(extension, fmt.Sprintf("atz=%v", atz))
	}

	if av := cef.Av(); av != "" {
		extension = append(extension, fmt.Sprintf("av=%v", av))
	}

	if customerExternalID := cef.CustomerExternalID(); customerExternalID != "" {
		extension = append(extension, fmt.Sprintf("customerExternalID=%v", customerExternalID))
	}

	if customerURI := cef.CustomerURI(); customerURI != "" {
		extension = append(extension, fmt.Sprintf("customerURI=%v", customerURI))
	}

	if destinatioTranslatedZoneExternalID := cef.DestinatioTranslatedZoneExternalID(); destinatioTranslatedZoneExternalID != "" {
		extension = append(extension, fmt.Sprintf("destinatioTranslatedZoneExternalID=%v", destinatioTranslatedZoneExternalID))
	}

	if destinationTranslatedZoneURI := cef.DestinationTranslatedZoneURI(); destinationTranslatedZoneURI != "" {
		extension = append(extension, fmt.Sprintf("destinationTranslatedZoneURI=%v", destinationTranslatedZoneURI))
	}

	if destinationZoneExternalID := cef.DestinationZoneExternalID(); destinationZoneExternalID != "" {
		extension = append(extension, fmt.Sprintf("destinationZoneExternalID=%v", destinationZoneExternalID))
	}

	if destinationZoneURI := cef.DestinationZoneURI(); destinationZoneURI != "" {
		extension = append(extension, fmt.Sprintf("destinationZoneURI=%v", destinationZoneURI))
	}

	if deviceTranslatedZoneExternalID := cef.DeviceTranslatedZoneExternalID(); deviceTranslatedZoneExternalID != "" {
		extension = append(extension, fmt.Sprintf("deviceTranslatedZoneExternalID=%v", deviceTranslatedZoneExternalID))
	}

	if deviceTranslatedZoneURI := cef.DeviceTranslatedZoneURI(); deviceTranslatedZoneURI != "" {
		extension = append(extension, fmt.Sprintf("deviceTranslatedZoneURI=%v", deviceTranslatedZoneURI))
	}

	if deviceZoneExternalID := cef.DeviceZoneExternalID(); deviceZoneExternalID != "" {
		extension = append(extension, fmt.Sprintf("deviceZoneExternalID=%v", deviceZoneExternalID))
	}

	if deviceZoneURI := cef.DeviceZoneURI(); deviceZoneURI != "" {
		extension = append(extension, fmt.Sprintf("deviceZoneURI=%v", deviceZoneURI))
	}

	if dlat := cef.Dlat(); dlat != 0 {
		extension = append(extension, fmt.Sprintf("dlat=%v", dlat))
	}

	if dlong := cef.Dlong(); dlong != 0 {
		extension = append(extension, fmt.Sprintf("dlong=%v", dlong))
	}

	if eventId := cef.EventId(); eventId != 0 {
		extension = append(extension, fmt.Sprintf("eventId=%v", eventId))
	}

	if rawEvent := cef.RawEvent(); rawEvent != "" {
		extension = append(extension, fmt.Sprintf("rawEvent=%v", rawEvent))
	}

	if slat := cef.Slat(); slat != 0 {
		extension = append(extension, fmt.Sprintf("slat=%v", slat))
	}

	if slong := cef.Slong(); slong != 0 {
		extension = append(extension, fmt.Sprintf("slong=%v", slong))
	}

	if sourceTranslatedZoneExternalID := cef.SourceTranslatedZoneExternalID(); sourceTranslatedZoneExternalID != "" {
		extension = append(extension, fmt.Sprintf("sourceTranslatedZoneExternalID=%v", sourceTranslatedZoneExternalID))
	}

	if sourceTranslatedZoneURI := cef.SourceTranslatedZoneURI(); sourceTranslatedZoneURI != "" {
		extension = append(extension, fmt.Sprintf("sourceTranslatedZoneURI=%v", sourceTranslatedZoneURI))
	}

	if sourceZoneExternalID := cef.SourceZoneExternalID(); sourceZoneExternalID != "" {
		extension = append(extension, fmt.Sprintf("sourceZoneExternalID=%v", sourceZoneExternalID))
	}

	if sourceZoneURI := cef.SourceZoneURI(); sourceZoneURI != "" {
		extension = append(extension, fmt.Sprintf("sourceZoneURI=%v", sourceZoneURI))
	}

	if agentTranslatedZoneKey := cef.AgentTranslatedZoneKey(); agentTranslatedZoneKey != 0 {
		extension = append(extension, fmt.Sprintf("agentTranslatedZoneKey=%v", agentTranslatedZoneKey))
	}

	if agentZoneKey := cef.AgentZoneKey(); agentZoneKey != 0 {
		extension = append(extension, fmt.Sprintf("agentZoneKey=%v", agentZoneKey))
	}

	if customerKey := cef.CustomerKey(); customerKey != 0 {
		extension = append(extension, fmt.Sprintf("customerKey=%v", customerKey))
	}

	if destinationTranslatedZoneKey := cef.DestinationTranslatedZoneKey(); destinationTranslatedZoneKey != 0 {
		extension = append(extension, fmt.Sprintf("destinationTranslatedZoneKey=%v", destinationTranslatedZoneKey))
	}

	if dZoneKey := cef.DZoneKey(); dZoneKey != 0 {
		extension = append(extension, fmt.Sprintf("dZoneKey=%v", dZoneKey))
	}

	if deviceTranslatedZoneKey := cef.DeviceTranslatedZoneKey(); deviceTranslatedZoneKey != 0 {
		extension = append(extension, fmt.Sprintf("deviceTranslatedZoneKey=%v", deviceTranslatedZoneKey))
	}

	if deviceZoneKey := cef.DeviceZoneKey(); deviceZoneKey != 0 {
		extension = append(extension, fmt.Sprintf("deviceZoneKey=%v", deviceZoneKey))
	}

	if sTranslatedZoneKey := cef.STranslatedZoneKey(); sTranslatedZoneKey != 0 {
		extension = append(extension, fmt.Sprintf("sTranslatedZoneKey=%v", sTranslatedZoneKey))
	}

	if sZoneKey := cef.SZoneKey(); sZoneKey != 0 {
		extension = append(extension, fmt.Sprintf("sZoneKey=%v", sZoneKey))
	}

	return strings.Join(extension, " ")
}
