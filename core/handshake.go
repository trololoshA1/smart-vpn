package core

import (
    "crypto/rand"
    "errors"

    "golang.org/x/crypto/curve25519"
)

type Handshake struct {
    ClientPriv [32]byte
    ClientPub  [32]byte
    ServerPub  [32]byte
    SharedKey  [32]byte
}

func NewHandshake() (*Handshake, error) {
    h := &Handshake{}

    // Генерация приватного ключа клиента
    _, err := rand.Read(h.ClientPriv[:])
    if err != nil {
        return nil, err
    }

    // Генерация публичного ключа клиента
    curve25519.ScalarBaseMult(&h.ClientPub, &h.ClientPriv)

    return h, nil
}

func (h *Handshake) SetServerPub(pub [32]byte) {
    h.ServerPub = pub
}

func (h *Handshake) ComputeSharedKey() error {
    if h.ServerPub == ([32]byte{}) {
        return errors.New("server public key not set")
    }

    curve25519.ScalarMult(&h.SharedKey, &h.ClientPriv, &h.ServerPub)
    return nil
}

func (h *Handshake) GetClientPub() [32]byte {
    return h.ClientPub
}

func (h *Handshake) GetSharedKey() [32]byte {
    return h.SharedKey
}