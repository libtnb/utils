package crypt

type Crypter interface {
	Encrypt(plaintext []byte) (string, error)
	Decrypt(payload string) ([]byte, error)
}
