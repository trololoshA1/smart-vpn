package mobile

import (
    "smart-vpn/core"
    "smart-vpn/geoip"
    "smart-vpn/adblock"
    "smart-vpn/subscriptions"
)

var vpnCore *core.VPN

func Init(geoPath string, adPath string, subsPath string, updateURL string) error {
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
    vpnCore.SetUpdateURL(updateURL)

    return nil
}

func AddSubscription(name, url string, auto bool, interval int) {
    vpnCore.SubsManager.AddSubscription(name, url, auto, interval)
}

func Connect() error {
    return vpnCore.Connect()
}

func HandlePacket(packet []byte) ([]byte, error) {
    return vpnCore.HandlePacket(packet)
}

func GetSubscriptions() []subscriptions.Subscription {
    return vpnCore.SubsManager.Subs
}

func CheckNodes() []subscriptions.Subscription {
    vpnCore.SubsManager.CheckAll()
    return vpnCore.SubsManager.Subs
}

func GetBestNode() *subscriptions.Subscription {
    return vpnCore.SubsManager.Best()
}

func Status() string {
    if vpnCore == nil {
        return "Не инициализировано"
    }
    return "Работает"
}