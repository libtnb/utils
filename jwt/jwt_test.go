package jwt

import (
	"testing"
	"time"

	"github.com/cristalhq/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateValidToken(t *testing.T) {
	j := NewJWT("secret", time.Hour)
	claims := &Claims{ID: "123", Subject: "test"}
	token, err := j.Generate(claims)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGenerateWithInvalidKey(t *testing.T) {
	j := NewJWT("", time.Hour)
	claims := &Claims{ID: "123", Subject: "test"}
	_, err := j.Generate(claims)
	assert.Error(t, err)
}

func TestParseValidToken(t *testing.T) {
	j := NewJWT("secret", time.Hour)
	claims := &Claims{ID: "123", Subject: "test"}
	token, err := j.Generate(claims)
	assert.NoError(t, err)

	parsedClaims, err := j.Parse(token)
	assert.NoError(t, err)
	assert.Equal(t, claims.ID, parsedClaims.ID)
	assert.Equal(t, claims.Subject, parsedClaims.Subject)
}

func TestParseInvalidToken(t *testing.T) {
	j := NewJWT("secret", time.Hour)
	_, err := j.Parse("invalid.token")
	assert.Error(t, err)
}

func TestParseWithInvalidKey(t *testing.T) {
	j := NewJWT("secret", time.Hour)
	claims := &Claims{ID: "123", Subject: "test"}
	token, err := j.Generate(claims)
	assert.NoError(t, err)

	jInvalid := NewJWT("wrongsecret", time.Hour)
	_, err = jInvalid.Parse(token)
	assert.Error(t, err)
}

func TestNotBeforeSetWhenNil(t *testing.T) {
	j := NewJWT("secret", time.Hour)
	claims := &Claims{ID: "123", Subject: "test"}
	token, err := j.Generate(claims)
	assert.NoError(t, err)

	parsedClaims, err := j.Parse(token)
	assert.NoError(t, err)
	assert.NotNil(t, parsedClaims.NotBefore)
}

func TestNotBeforePreservedWhenSet(t *testing.T) {
	j := NewJWT("secret", time.Hour)
	notBefore := jwt.NewNumericDate(time.Now().Add(-time.Hour))
	claims := &Claims{ID: "123", Subject: "test", NotBefore: notBefore}
	token, err := j.Generate(claims)
	assert.NoError(t, err)

	parsedClaims, err := j.Parse(token)
	assert.NoError(t, err)
	assert.WithinDuration(t, notBefore.Time, parsedClaims.NotBefore.Time, time.Second)
}
