package jwt

import (
	"errors"
	"time"

	"github.com/cristalhq/jwt/v5"

	"github.com/libtnb/utils/convert"
	"github.com/libtnb/utils/uuid"
)

type Claims = jwt.RegisteredClaims

type JWT struct {
	key []byte
	ttl time.Duration
}

var ErrInvalidClaims = errors.New("invalid claims")

func NewJWT(key string, ttl time.Duration) *JWT {
	return &JWT{
		key: convert.UnsafeBytes(key),
		ttl: ttl,
	}
}

func (r *JWT) Generate(claims *Claims) (string, error) {
	signer, err := jwt.NewSignerHS(jwt.HS256, r.key)
	if err != nil {
		return "", err
	}

	now := time.Now()
	claims.IssuedAt = jwt.NewNumericDate(now)
	claims.ExpiresAt = jwt.NewNumericDate(now.Add(r.ttl))
	if claims.NotBefore == nil {
		claims.NotBefore = jwt.NewNumericDate(now)
	}
	if claims.ID == "" {
		claims.ID = uuid.UUID()
	}

	token, err := jwt.NewBuilder(signer).Build(claims)
	if err != nil {
		return "", err
	}

	return convert.UnsafeString(token.Bytes()), nil
}

func (r *JWT) Parse(token string) (*Claims, error) {
	verifier, err := jwt.NewVerifierHS(jwt.HS256, r.key)
	if err != nil {
		return nil, err
	}

	claims := new(Claims)
	if err = jwt.ParseClaims(convert.UnsafeBytes(token), verifier, claims); err != nil {
		return nil, err
	}
	if !claims.IsValidAt(time.Now()) {
		return nil, ErrInvalidClaims
	}

	return claims, nil
}
