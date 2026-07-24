package core

import (
    "log"
    "net"
    "strings"

    "smart-vpn/adblock"
    "smart-vpn/geoip"
    "smart-vpn/subscriptions"
)

type VPN struct {
    Tunnel      *Tunnel
    Geo         *geoip.GeoIP
    AdBlock     *adblock.AdBlock
    SubsManager *subscriptions.Manager
}

// Initialize VPN core
func NewVPN(geo *geoip.GeoIP, ad *adblock.AdBlock, subs *subscriptions.Manager) *VPN {
    return &VPN{
        Geo:         geo,
        AdBlock:     ad,
        SubsManager: subs,
    }
}

// Connect to best subscription
func (v *VPN) Connect() error {
    v.SubsManager.CheckAll()
    best := v.SubsManager.Best()
    if best == nil {
        return ErrNoAliveNodes
    }

    tunnel, err := NewTunnel("wss://" + best.Address + "/tunnel")
    if err != nil {
        return err
    }

    v.Tunnel = tunnel
    return nil
}

var ErrNoAliveNodes = &VPNError{"Нет рабочих подписок"}

type VPNError struct {
    msg string
}

func (e *VPNError) Error() string {
    return e.msg
}

// Main packet handler
func (v *VPN) HandlePacket(packet []byte) ([]byte, error) {
    // DNS packet?
    if isDNS(packet) {
        domain := extractDomain(packet)

        // AdBlock check
        if v.AdBlock.IsBlocked(domain) {
            log.Println("Blocked ad domain:", domain)
            return nil, nil // drop packet
        }
    }

    // Extract destination IP
    dstIP := extractDstIP(packet)
    if dstIP == "" {
        return nil, nil
    }

    // GeoIP routing
    useTunnel := v.Geo.ShouldUseTunnel(dstIP)

    if useTunnel {
        // Send through VPN tunnel
        return v.Tunnel.SendPacket(packet)
    } else {
        // Direct connection
        return directForward(packet, dstIP)
    }
}

// Very simple direct forward (will be improved later)
func directForward(packet []byte, ip string) ([]byte, error) {
    conn, err := net.Dial("tcp", ip+":80")
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
    if err != nil {
        return nil, err
    }

    return buf[:n], nil
}

// Helpers (simplified for now)
func isDNS(packet []byte) bool {
    return len(packet) > 2 && packet[2] == 1
}

func extractDomain(packet []byte) string {
    s := string(packet)
    if strings.Contains(s, "www.") {
        return s[strings.Index(s, "www."):]
    }
    return ""
}

func extractDstIP(packet []byte) string {
    if len(packet) < 20 {
        return ""
    }
    ip := net.IP(packet[16:20])
    return ip.String()
}