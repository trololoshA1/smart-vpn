package mobile

import (
    "smart-vpn/core"
    "smart-vpn/geoip"
    "smart-vpn/adblock"
    "smart-vpn/subscriptions"
)

var vpnCore *core.VPN

// Инициализация VPN ядра
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

// Подключение к лучшему узлу
func Connect() error {
    return vpnCore.Connect()
}

// Обработка пакета из TUN
func HandlePacket(packet []byte) ([]byte, error) {
    return vpnCore.HandlePacket(packet)
}

// Получить список подписок
func GetSubscriptions() []subscriptions.Subscription {
    return vpnCore.SubsManager.Subs
}

// Проверить все узлы (ping)
func CheckNodes() []subscriptions.Subscription {
    vpnCore.SubsManager.CheckAll()
    return vpnCore.SubsManager.Subs
}

// Получить лучший узел
func GetBestNode() *subscriptions.Subscription {
    return vpnCore.SubsManager.Best()
}

// Получить статус VPN
func Status() string {
    if vpnCore == nil {
        return "Не инициализировано"
    }
    return "Работает"
}