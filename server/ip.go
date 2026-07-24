package core

import (
    "net"
)

func ProcessIPPacket(packet []byte) (dstIP net.IP, dstPort int, proto int) {
    if len(packet) < 20 {
        return nil, 0, 0
    }

    proto = int(packet[9])
    dstIP = net.IP(packet[16:20])
    dstPort = int(packet[22])<<8 | int(packet[23])

    return
}