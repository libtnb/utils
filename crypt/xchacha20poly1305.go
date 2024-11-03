package crypt

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/chacha20poly1305"
)

type XChacha20Poly1305 struct {
	key []byte
}

// NewXChacha20Poly1305 returns a new XChacha20Poly1305 crypt instance.
func NewXChacha20Poly1305(key []byte) (Crypter, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, errors.New("the key must be 32 bytes")
	}

	return &XChacha20Poly1305{
		key: key,
	}, nil
}

// Encrypt encrypts the given plaintext, and returns the base64 encoded payload.
func (r *XChacha20Poly1305) Encrypt(plaintext []byte) (string, error) {
	aead, err := chacha20poly1305.NewX(r.key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(aead.Seal(nonce, nonce, plaintext, nil)), nil
}

// Decrypt decrypts the given base64 encoded payload, and returns the plaintext.
func (r *XChacha20Poly1305) Decrypt(payload string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	aead, err := chacha20poly1305.NewX(r.key)
	if err != nil {
		return nil, err
	}

	if len(decoded) < aead.NonceSize() {
		return nil, errors.New("payload too short")
	}

	nonce, ciphertext := decoded[:aead.NonceSize()], decoded[aead.NonceSize():]
	return aead.Open(nil, nonce, ciphertext, nil)
}
