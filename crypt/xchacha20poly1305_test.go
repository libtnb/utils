package crypt

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type XChacha20Poly1305TestSuite struct {
	suite.Suite
}

func TestXChacha20Poly1305TestSuite(t *testing.T) {
	suite.Run(t, &XChacha20Poly1305TestSuite{})
}

func (s *XChacha20Poly1305TestSuite) TestEncrypt() {
	encrypter, err := NewXChacha20Poly1305([]byte("12345678901234567890123456789012"))
	s.NoError(err)
	payload, err := encrypter.Encrypt([]byte("test"))
	s.NotEmpty(payload)
	s.NoError(err)
}

func (s *XChacha20Poly1305TestSuite) TestEncryptEmpty() {
	encrypter, err := NewXChacha20Poly1305([]byte("12345678901234567890123456789012"))
	s.NoError(err)
	payload, err := encrypter.Encrypt([]byte(""))
	s.NotEmpty(payload)
	s.NoError(err)
}

func (s *XChacha20Poly1305TestSuite) TestDecrypt() {
	encrypter, err := NewXChacha20Poly1305([]byte("12345678901234567890123456789012"))
	s.NoError(err)
	value, err := encrypter.Encrypt([]byte("test"))
	s.NoError(err)
	plaintext, err := encrypter.Decrypt(value)
	s.NoError(err)
	s.Equal("test", string(plaintext))
}

func (s *XChacha20Poly1305TestSuite) TestDecryptEmpty() {
	encrypter, err := NewXChacha20Poly1305([]byte("12345678901234567890123456789012"))
	s.NoError(err)
	value, err := encrypter.Encrypt([]byte(""))
	s.NoError(err)
	plaintext, err := encrypter.Decrypt(value)
	s.NoError(err)
	s.Empty(plaintext)
}
