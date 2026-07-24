package core

import (
    "log"
    "net"
    "time"

    "github.com/gorilla/websocket"
)

type Tunnel struct {
    conn *websocket.Conn
}

func NewTunnel(serverURL string) (*Tunnel, error) {
    dialer := websocket.Dialer{
        HandshakeTimeout: 10 * time.Second,
    }

    conn, _, err := dialer.Dial(serverURL, nil)
    if err != nil {
        return nil, err
    }

    return &Tunnel{conn: conn}, nil
}

func (t *Tunnel) SendPacket(packet []byte) ([]byte, error) {
    err := t.conn.WriteMessage(websocket.BinaryMessage, packet)
    if err != nil {
        return nil, err
    }

    _, response, err := t.conn.ReadMessage()
    if err != nil {
        return nil, err
    }

    return response, nil
}

func (t *Tunnel) Close() {
    if t.conn != nil {
        t.conn.Close()
    }
}

// Simple TCP ping for subscription checking
func TcpPing(address string) time.Duration {
    start := time.Now()
    conn, err := net.DialTimeout("tcp", address, 3*time.Second)
    if err != nil {
        return -1
    }
    conn.Close()
    return time.Since(start)
}