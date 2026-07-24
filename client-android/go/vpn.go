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

func AddSubscription(name, url string, auto bool, interval int) {
    vpnCore.SubsManager.AddSubscription(name, url, auto, interval)
}

func RemoveSubscription(index int) {
    vpnCore.SubsManager.RemoveSubscription(index)
}

func EditSubscription(index int, name, url string, auto bool, interval int) {
    vpnCore.SubsManager.UpdateLocalSubscription(index, name, url, auto, interval)
}

func GetSubscriptions() []subscriptions.Subscription {
    return vpnCore.SubsManager.Subs
}