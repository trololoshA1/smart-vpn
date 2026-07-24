package core

import (
    "crypto/cipher"
    "golang.org/x/crypto/chacha20poly1305"
)

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