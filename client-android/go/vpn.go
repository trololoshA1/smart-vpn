package mobile

import (
    "smart-vpn/core"
    "smart-vpn/subscriptions"
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
    return nil
}

func Connect() error {
    return vpnCore.Connect()
}

func HandlePacket(packet []byte) ([]byte, error) {
    return vpnCore.HandlePacket(packet)
}

func BestNode() (*subscriptions.Node, error) {
    node := vpnCore.SubsManager.BestNode()
    if node == nil {
        return nil, errors.New("нет узлов")
    }
    return node, nil
}

func GetSubscriptions() []subscriptions.Subscription {
    return vpnCore.SubsManager.Subs
}