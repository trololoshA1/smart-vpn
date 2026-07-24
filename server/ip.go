package server

import (
    "net"
)

func ProcessIPPacket(packet []byte) []byte {
    if len(packet) < 20 {
        return nil
    }

    proto := packet[9]

    switch proto {
    case 6: // TCP
        return HandleTCP(packet)
    case 17: // UDP
        return HandleUDP(packet)
    default:
        return nil
    }
}