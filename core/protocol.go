package core

import (
    "crypto/rand"
)

type Protocol struct {
    Key [32]byte
}

func NewProtocol(key [32]byte) *Protocol {
    return &Protocol{Key: key}
}

func (p *Protocol) NewNonce() [12]byte {
    var nonce [12]byte
    rand.Read(nonce[:])
    return nonce
}