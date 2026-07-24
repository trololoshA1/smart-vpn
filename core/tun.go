package core

import (
    "io"
    "log"
)

// TUN interface abstraction
type TunInterface interface {
    ReadPacket() ([]byte, error)
    WritePacket([]byte) error
}

// Universal TUN handler
type TunHandler struct {
    tun TunInterface
    vpn *VPN
}

func NewTunHandler(tun TunInterface, vpn *VPN) *TunHandler {
    return &TunHandler{
        tun: tun,
        vpn: vpn,
    }
}

// Main loop: read → process → write
func (h *TunHandler) Start() {
    log.Println("TUN handler started")

    for {
        packet, err := h.tun.ReadPacket()
        if err == io.EOF {
            log.Println("TUN closed")
            return
        }
        if err != nil {
            log.Println("Read error:", err)
            continue
        }

        // Process packet through VPN core
        response, err := h.vpn.HandlePacket(packet)
        if err != nil {
            log.Println("VPN error:", err)
            continue
        }

        // If response is nil (adblock or drop), skip
        if response == nil {
            continue
        }

        // Write back to TUN
        err = h.tun.WritePacket(response)
        if err != nil {
            log.Println("Write error:", err)
        }
    }
}