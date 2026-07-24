package server

import (
    "crypto/rand"

    "golang.org/x/crypto/chacha20poly1305"
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

func (p *Protocol) Encrypt(nonce [12]byte, plaintext []byte) ([]byte, error) {
    aead, err := chacha20poly1305.New(p.Key[:])
    if err != nil {
        return nil, err
    }
    return aead.Seal(nil, nonce[:], plaintext, nil), nil
}

func (p *Protocol) Decrypt(nonce [12]byte, ciphertext []byte) ([]byte, error) {
    aead, err := chacha20poly1305.New(p.Key[:])
    if err != nil {
        return nil, err
    }
    return aead.Open(nil, nonce[:], ciphertext, nil)
}