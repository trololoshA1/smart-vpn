package core

import (
    "net"
)

func HandleTCPLocal(packet []byte) ([]byte, error) {
    info := ParseIP(packet)
    if info == nil {
        return nil, nil
    }

    conn, err := net.Dial("tcp", net.JoinHostPort(info.DstIP.String(), itoa(info.DstPort)))
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    // Отправляем только TCP payload, не весь IP пакет
    tcpPayload := packet[20:]
    _, err = conn.Write(tcpPayload)
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

func itoa(p uint16) string {
    return net.Itoa(int(p))
}