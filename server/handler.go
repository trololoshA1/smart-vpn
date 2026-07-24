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

        // Здесь обработка IP-пакета
        response := ProcessIPPacket(plaintext)

        nonce := protocol.NewNonce()
        encrypted, _ := protocol.Encrypt(nonce, response)
        encoded := EncodePacket(nonce, encrypted)

        conn.Write(encoded)
    }
}