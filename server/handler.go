package server

import (
    "net"
)

func HandleEncrypted(conn net.Conn, protocol *Protocol) {
    defer conn.Close()

    buf := make([]byte, 65535)

    for {
        n, err := conn.Read(buf)
        if err != nil {
            return
        }

        pkt, err := DecodePacket(buf[:n])
        if err != nil {
            continue
        }

        plaintext, err := protocol.Decrypt(pkt.Nonce, pkt.Payload)
        if err != nil {
            continue
        }

        proto := ProcessIPPacket(plaintext)

        var response []byte

        switch proto {
        case 6:
            response = HandleTCP(plaintext)
        case 17:
            response = HandleUDP(plaintext)
        default:
            response = nil
        }

        if response == nil {
            continue
        }

        nonce := protocol.NewNonce()
        encrypted, _ := protocol.Encrypt(nonce, response)
        encoded := EncodePacket(nonce, encrypted)

        conn.Write(encoded)
    }
}