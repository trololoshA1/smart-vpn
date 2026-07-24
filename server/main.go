package main

import (
    "encoding/binary"
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

        go handle(conn)
    }
}

func handle(conn net.Conn) {
    defer conn.Close()

    hs, err := server.NewServerHandshake()
    if err != nil {
        return
    }

    // Получаем публичный ключ клиента
    var clientPub [32]byte
    _, err = conn.Read(clientPub[:])
    if err != nil {
        return
    }

    hs.SetClientPub(clientPub)
    hs.ComputeSharedKey()

    // Отправляем публичный ключ сервера
    conn.Write(hs.ServerPub[:])

    // Теперь соединение готово к шифрованному обмену
    // SharedKey — общий ключ для ChaCha20-Poly1305
}