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
    // Connect VPN core
    err := a.vpn.Connect()
    if err != nil {
        return err
    }

    // Create Windows TUN
    tun, err := NewWindowsTun()
    if err != nil {
        return err
    }
    a.tun = tun

    // Start TUN handler
    handler := core.NewTunHandler(tun, a.vpn)
    go handler.Start()

    return nil
}

func (a *App) Status() string {
    if a.vpn == nil {
        return "Not initialized"
    }
    return "Connected"
}