package go

import (
    "smart-vpn/core"
    "smart-vpn/geoip"
    "smart-vpn/adblock"
    "smart-vpn/subscriptions"
)

type App struct {
    vpn *core.VPN
    tun *WindowsTun
}

func NewApp() *App {
    return &App{}
}

// Инициализация VPN ядра
func (a *App) Init(geoPath, adPath, subsPath string) error {
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

    a.vpn = core.NewVPN(geo, ad, subs)
    return nil
}

// Подключение VPN
func (a *App) Connect() error {
    err := a.vpn.Connect()
    if err != nil {
        return err
    }

    tun, err := NewWindowsTun()
    if err != nil {
        return err
    }
    a.tun = tun

    handler := core.NewTunHandler(tun, a.vpn)
    go handler.Start()

    return nil
}

// Получить список подписок
func (a *App) GetSubs() []subscriptions.Subscription {
    return a.vpn.SubsManager.Subs
}

// Проверить узлы (ping)
func (a *App) CheckNodes() []subscriptions.Subscription {
    a.vpn.SubsManager.CheckAll()
    return a.vpn.SubsManager.Subs
}

// Лучший узел
func (a *App) BestNode() *subscriptions.Subscription {
    return a.vpn.SubsManager.Best()
}

// Статус VPN
func (a *App) Status() string {
    if a.vpn == nil {
        return "Не инициализировано"
    }
    return "Работает"
}