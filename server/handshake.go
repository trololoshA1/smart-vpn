package server

import (
    "crypto/rand"
    "errors"

    "golang.org/x/crypto/curve25519"
)

type ServerHandshake struct {
    ServerPriv [32]byte
    ServerPub  [32]byte
    ClientPub  [32]byte
    SharedKey  [32]byte
}

func NewServerHandshake() (*ServerHandshake, error) {
    h := &ServerHandshake{}

    _, err := rand.Read(h.ServerPriv[:])
    if err != nil {
        return nil, err
    }

    curve25519.ScalarBaseMult(&h.ServerPub, &h.ServerPriv)

    return h, nil
}

func (h *ServerHandshake) SetClientPub(pub [32]byte) {
    h.ClientPub = pub
}

func (h *ServerHandshake) ComputeSharedKey() error {
    if h.ClientPub == ([32]byte{}) {
        return errors.New("client public key not set")
    }

    curve25519.ScalarMult(&h.SharedKey, &h.ServerPriv, &h.ClientPub)
    return nil
}

func (h *ServerHandshake) GetServerPub() [32]byte {
    return h.ServerPub
}

func (h *ServerHandshake) GetSharedKey() [32]byte {
    return h.SharedKey
}