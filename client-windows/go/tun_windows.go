package go

import (
    "github.com/WireGuard/wintun"
)

type WindowsTun struct {
    adapter *wintun.Adapter
    session *wintun.Session
}

func NewWindowsTun() (*WindowsTun, error) {
    adapter, err := wintun.CreateAdapter("SmartVPN", "SmartVPN", nil)
    if err != nil {
        return nil, err
    }

    session, err := adapter.StartSession(0x400000)
    if err != nil {
        return nil, err
    }

    return &WindowsTun{
        adapter: adapter,
        session: session,
    }, nil
}

func (t *WindowsTun) ReadPacket() ([]byte, error) {
    return t.session.ReceivePacket()
}

func (t *WindowsTun) WritePacket(packet []byte) error {
    t.session.SendPacket(packet)
    return nil
}