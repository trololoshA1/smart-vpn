package server

import (
    "net"
)

func HandleTCP(packet []byte) []byte {
    dstIP := net.IP(packet[16:20])
    dstPort := int(packet[22])<<8 | int(packet[23])

    conn, err := net.Dial("tcp", net.JoinHostPort(dstIP.String(), string(dstPort)))
    if err != nil {
        return nil
    }
    defer conn.Close()

    conn.Write(packet)

    buf := make([]byte, 65535)
    n, _ := conn.Read(buf)

    return buf[:n]
}