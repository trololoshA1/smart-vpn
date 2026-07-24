package mobile

import (
    "smart-vpn/subscriptions"
    "smart-vpn/core"
    "smart-vpn/geoip"
    "smart-vpn/adblock"
)

var vpnCore *core.VPN

func Init(geoPath string, adPath string, subsPath string, updateURL string) error {
    geo, _ := geoip.NewGeoIP(geoPath)
    ad, _ := adblock.NewAdBlock(adPath)

    subs := subscriptions.NewManager()
    subs.Load(subsPath)

    vpnCore = core.NewVPN(geo, ad, subs)
    vpnCore.SetUpdateURL(updateURL)

    return nil
}

func Connect() error {
    return vpnCore.Connect()
}

func BestNode() (*subscriptions.Node, error) {
    return vpnCore.BestNode()
}

func GetSubscriptions() []subscriptions.Subscription {
    return vpnCore.SubsManager.Subs
}