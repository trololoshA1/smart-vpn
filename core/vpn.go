package core

import (
    "errors"
    "smart-vpn/subscriptions"
)

type VPN struct {
    Geo        *GeoIP
    Ad         *AdBlock
    SubsManager *subscriptions.Manager
    UpdateURL  string
}

func NewVPN(geo *GeoIP, ad *AdBlock, subs *subscriptions.Manager) *VPN {
    return &VPN{
        Geo:        geo,
        Ad:         ad,
        SubsManager: subs,
    }
}

func (v *VPN) SetUpdateURL(url string) {
    v.UpdateURL = url
}

func (v *VPN) BestNode() (*subscriptions.Node, error) {
    node := v.SubsManager.BestNode()
    if node == nil {
        return nil, errors.New("нет доступных узлов")
    }
    return node, nil
}

func (v *VPN) Connect() error {
    node, err := v.BestNode()
    if err != nil {
        return err
    }

    // Здесь будет логика подключения к узлу
    // node.Address — IP:порт
    // node.Region — регион
    // node.Name — имя узла

    return nil
}