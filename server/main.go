package main

import (
    "log"
    "net"

    "smart-vpn/server"
)

func main() {
    ln, err := net.Listen("tcp", ":9000")
    if err != nil {
        log.Fatal(err)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }

        go func(c net.Conn) {
            hs, _ := server.NewServerHandshake()

            var clientPub [32]byte
            c.Read(clientPub[:])

            hs.SetClientPub(clientPub)
            hs.ComputeSharedKey()

            c.Write(hs.ServerPub[:])

            protocol := server.NewProtocol(hs.SharedKey)

            server.HandleEncrypted(c, protocol)
        }(conn)
    }
}