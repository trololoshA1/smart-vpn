package core

import (
    "errors"
    "net"
)

type VPN struct {
    Geo         *GeoIP
    Ad          *AdBlock
    SubsManager *subscriptions.Manager
    Protocol    *Protocol
    Handshake   *Handshake
    Conn        net.Conn
}

func NewVPN(geo *GeoIP, ad *AdBlock, subs *subscriptions.Manager) *VPN {
    return &VPN{
        Geo:         geo,
        Ad:          ad,
        SubsManager: subs,
    }
}

func (v *VPN) Connect() error {
    node := v.SubsManager.BestNode()
    if node == nil {
        return errors.New("нет узлов")
    }

    // TCP соединение с сервером узла
    conn, err := net.Dial("tcp", node.Address)
    if err != nil {
        return err
    }
    v.Conn = conn

    // Handshake
    hs, err := NewHandshake()
    if err != nil {
        return err
    }
    v.Handshake = hs

    // Отправляем публичный ключ клиента
    _, err = conn.Write(hs.ClientPub[:])
    if err != nil {
        return err
    }

    // Получаем публичный ключ сервера
    var serverPub [32]byte
    _, err = conn.Read(serverPub[:])
    if err != nil {
        return err
    }

    hs.SetServerPub(serverPub)
    hs.ComputeSharedKey()

    // Создаём протокол с общим ключом
    v.Protocol = NewProtocol(hs.SharedKey)

    return nil
}

func (v *VPN) HandlePacket(packet []byte) ([]byte, error) {
    if v.Protocol == nil {
        return nil, errors.New("нет протокола")
    }

    nonce := v.Protocol.NewNonce()

    encrypted, err := v.Protocol.Encrypt(nonce, packet)
    if err != nil {
        return nil, err
    }

    encoded := EncodePacket(nonce, encrypted)

    // Отправляем на сервер
    _, err = v.Conn.Write(encoded)
    if err != nil {
        return nil, err
    }

    // Получаем ответ
    buf := make([]byte, 65535)
    n, err := v.Conn.Read(buf)
    if err != nil {
        return nil, err
    }

    decoded, err := DecodePacket(buf[:n])
    if err != nil {
        return nil, err
    }

    plaintext, err := v.Protocol.Decrypt(decoded.Nonce, decoded.Payload)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}