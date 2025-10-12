package crypt

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AES256GCMTestSuite struct {
	suite.Suite
}

func TestAES256GCMTestSuite(t *testing.T) {
	suite.Run(t, &AES256GCMTestSuite{})
}

func (s *AES256GCMTestSuite) TestEncrypt() {
	encrypter, err := NewAES256GCM([]byte("12345678901234567890123456789012"))
	s.NoError(err)
	payload, err := encrypter.Encrypt([]byte("test"))
	s.NotEmpty(payload)
	s.NoError(err)
}

func (s *AES256GCMTestSuite) TestEncryptEmpty() {
	encrypter, err := NewAES256GCM([]byte("12345678901234567890123456789012"))
	s.NoError(err)
	payload, err := encrypter.Encrypt([]byte(""))
	s.NotEmpty(payload)
	s.NoError(err)
}

func (s *AES256GCMTestSuite) TestDecrypt() {
	encrypter, err := NewAES256GCM([]byte("12345678901234567890123456789012"))
	s.NoError(err)
	value, err := encrypter.Encrypt([]byte("test"))
	s.NoError(err)
	plaintext, err := encrypter.Decrypt(value)
	s.NoError(err)
	s.Equal("test", string(plaintext))
}

func (s *AES256GCMTestSuite) TestDecryptEmpty() {
	encrypter, err := NewAES256GCM([]byte("12345678901234567890123456789012"))
	s.NoError(err)
	value, err := encrypter.Encrypt([]byte(""))
	s.NoError(err)
	plaintext, err := encrypter.Decrypt(value)
	s.NoError(err)
	s.Empty(plaintext)
}
