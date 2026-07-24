package core

import (
    "encoding/binary"
    "errors"
)

const (
    MagicHigh = 0x53 // 'S'
    MagicLow  = 0x56 // 'V'
    Version   = 1
)

type Packet struct {
    Nonce   [12]byte
    Payload []byte
}

func EncodePacket(nonce [12]byte, payload []byte) []byte {
    totalLen := 2 + 1 + 12 + 2 + len(payload)
    buf := make([]byte, totalLen)

    // Magic
    buf[0] = MagicHigh
    buf[1] = MagicLow

    // Version
    buf[2] = Version

    // Nonce
    copy(buf[3:15], nonce[:])

    // Payload length
    binary.BigEndian.PutUint16(buf[15:17], uint16(len(payload)))

    // Payload
    copy(buf[17:], payload)

    return buf
}

func DecodePacket(data []byte) (*Packet, error) {
    if len(data) < 17 {
        return nil, errors.New("packet too short")
    }

    if data[0] != MagicHigh || data[1] != MagicLow {
        return nil, errors.New("invalid magic")
    }

    if data[2] != Version {
        return nil, errors.New("invalid version")
    }

    var nonce [12]byte
    copy(nonce[:], data[3:15])

    length := binary.BigEndian.Uint16(data[15:17])

    if int(length)+17 > len(data) {
        return nil, errors.New("invalid payload length")
    }

    payload := data[17 : 17+length]

    return &Packet{
        Nonce:   nonce,
        Payload: payload,
    }, nil
}