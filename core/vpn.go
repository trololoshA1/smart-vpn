package core

import (
    "errors"
)

func (v *VPN) HandlePacket(packet []byte) ([]byte, error) {
    if v.Protocol == nil {
        return nil, errors.New("нет протокола")
    }

    info := ParseIP(packet)
    if info == nil {
        return nil, nil
    }

    // Шифруем и отправляем на сервер
    nonce := v.Protocol.NewNonce()
    encrypted, err := v.Protocol.Encrypt(nonce, packet)
    if err != nil {
        return nil, err
    }

    encoded := EncodePacket(nonce, encrypted)

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