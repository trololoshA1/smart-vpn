package go

import (
    "smart-vpn/core"
    "smart-vpn/geoip"
    "smart-vpn/adblock"
    "smart-vpn/subscriptions"
)

type App struct {
    vpn *core.VPN
}

func NewApp() *App {
    return &App{}
}

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

func (a *App) Connect() error {
    return a.vpn.Connect()
}

func (a *App) Status() string {
    if a.vpn == nil {
        return "Not initialized"
    }
    return "Connected"
}