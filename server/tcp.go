package server

import (
    "encoding/binary"
    "net"
)

func HandleTCP(packet []byte) []byte {
    if len(packet) < 40 {
        return nil
    }

    dstIP := net.IP(packet[16:20])
    dstPort := binary.BigEndian.Uint16(packet[22:24])

    conn, err := net.Dial("tcp", net.JoinHostPort(dstIP.String(), stringPort(dstPort)))
    if err != nil {
        return nil
    }
    defer conn.Close()

    // Отправляем TCP payload (весь IP-пакет)
    _, err = conn.Write(packet)
    if err != nil {
        return nil
    }

    buf := make([]byte, 65535)
    n, err := conn.Read(buf)
    if err != nil {
        return nil
    }

    return buf[:n]
}

func stringPort(p uint16) string {
    return net.Itoa(int(p))
}