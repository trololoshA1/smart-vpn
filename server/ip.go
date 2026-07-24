package core

import (
    "encoding/binary"
    "net"
)

type IPInfo struct {
    Proto   int
    SrcIP   net.IP
    DstIP   net.IP
    SrcPort uint16
    DstPort uint16
}

func ParseIP(packet []byte) *IPInfo {
    if len(packet) < 20 {
        return nil
    }

    proto := int(packet[9])
    srcIP := net.IP(packet[12:16])
    dstIP := net.IP(packet[16:20])

    var srcPort, dstPort uint16

    if proto == 6 || proto == 17 {
        srcPort = binary.BigEndian.Uint16(packet[20:22])
        dstPort = binary.BigEndian.Uint16(packet[22:24])
    }

    return &IPInfo{
        Proto:   proto,
        SrcIP:   srcIP,
        DstIP:   dstIP,
        SrcPort: srcPort,
        DstPort: dstPort,
    }
}