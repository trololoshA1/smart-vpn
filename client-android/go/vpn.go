package mobile

import (
    "smart-vpn/core"
    "smart-vpn/geoip"
    "smart-vpn/adblock"
    "smart-vpn/subscriptions"
)

var vpnCore *core.VPN

// Initialize VPN core
func Init(geoPath string, adPath string, subsPath string) error {
    geo, err := geoip.NewGeoIP(geoPath)
    if err != nil {
        return err
    }

    ad, err := adblock.NewAdBlock(adPath)
    if err != nil {
        return err
    }

    subs := subscriptions.NewManager()
    subs.Load(subsPath)

    vpnCore = core.NewVPN(geo, ad, subs)
    return nil
}

// Connect to best node
func Connect() error {
    return vpnCore.Connect()
}

// Handle packet from TUN
func HandlePacket(packet []byte) ([]byte, error) {
    return vpnCore.HandlePacket(packet)
}