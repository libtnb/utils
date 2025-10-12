package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

type AES256GCM struct {
	key []byte
}

// NewAES256GCM returns a new AES256GCM crypt instance.
func NewAES256GCM(key []byte) (Crypter, error) {
	if len(key) != 32 {
		return nil, errors.New("the key must be 32 bytes")
	}

	return &AES256GCM{
		key: key,
	}, nil
}

// Encrypt encrypts the given plaintext, and returns the base64 encoded payload.
func (r *AES256GCM) Encrypt(plaintext []byte) (string, error) {
	block, err := aes.NewCipher(r.key)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(block)
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
func (r *AES256GCM) Decrypt(payload string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(r.key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(decoded) < aead.NonceSize() {
		return nil, errors.New("payload too short")
	}

	nonce, ciphertext := decoded[:aead.NonceSize()], decoded[aead.NonceSize():]
	return aead.Open(nil, nonce, ciphertext, nil)
}
