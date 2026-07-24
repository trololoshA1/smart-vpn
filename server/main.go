package main

import (
    "io"
    "log"
    "net"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func main() {
    http.HandleFunc("/tunnel", handleTunnel)

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTunnel(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Upgrade error:", err)
        return
    }
    defer conn.Close()

    for {
        _, data, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            return
        }

        // data = IP packet from client
        response, err := forwardPacket(data)
        if err != nil {
            log.Println("Forward error:", err)
            continue
        }

        conn.WriteMessage(websocket.BinaryMessage, response)
    }
}

func forwardPacket(packet []byte) ([]byte, error) {
    // VERY SIMPLE FOR NOW:
    // We treat packet as raw TCP payload to google.com:80
    // Later we will replace this with real IP routing.

    conn, err := net.Dial("tcp", "google.com:80")
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    _, err = conn.Write(packet)
    if err != nil {
        return nil, err
    }

    buf := make([]byte, 4096)
    n, err := conn.Read(buf)
    if err != nil && err != io.EOF {
        return nil, err
    }

    return buf[:n], nil
}