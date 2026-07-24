package core

import (
    "net"
)

func HandleUDPLocal(packet []byte) ([]byte, error) {
    dstIP, dstPort, _ := ProcessIPPacket(packet)
    if dstIP == nil {
        return nil, nil
    }

    conn, err := net.Dial("udp", net.JoinHostPort(dstIP.String(), string(dstPort)))
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    _, err = conn.Write(packet)
    if err != nil {
        return nil, err
    }

    buf := make([]byte, 65535)
    n, err := conn.Read(buf)
    if err != nil {
        return nil, err
    }

    return buf[:n], nil
}