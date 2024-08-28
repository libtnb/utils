package hash

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BcryptTestSuite struct {
	suite.Suite
}

func TestBcryptTestSuite(t *testing.T) {
	suite.Run(t, &BcryptTestSuite{})
}

func (s *BcryptTestSuite) TestMake() {
	hasher := NewBcrypt()
	hash, err := hasher.Make("password")
	s.NotEmpty(hash)
	s.NoError(err)
}

func (s *BcryptTestSuite) TestCheck() {
	hasher := NewBcrypt()
	value, err := hasher.Make("password")
	s.NoError(err)
	s.True(hasher.Check("password", value))
	s.False(hasher.Check("password1", value))
	s.False(hasher.Check("password", "hash"))
	s.False(hasher.Check("password", "hashhash"))
	s.True(hasher.Check("password", "$2a$12$4V6LZSkQpOP..TQD2v6WP.XVbT946ImH9lIsczrAqp3rD0ft5Q/Zu"))
	s.False(hasher.Check("password", "$2a$12$4V6LZSkQpOP..AQD2v6WP.XVbT946ImH9lIsczrAqp3rD0ft5Q/Zu"))
	s.False(hasher.Check("password", "$2a$12$CQxHxgAZhtMWvLB.oZQqye./9QTXGgFq0jd5sHQhlm/5b4pWdJFhK"))
}

func (s *BcryptTestSuite) TestConfigurationOverride() {
	hasher := NewBcrypt()
	value := "$2a$10$dl7Gf.WnkTe3.a8F6gps0u5OLv5yoSMwhQ/NEf.WBE2FwPvVOO62q"
	s.True(hasher.Check("rat", value))
	s.True(hasher.NeedsRehash(value))
}

func (s *BcryptTestSuite) TestNeedsRehash() {
	hasher := NewBcrypt()
	value, err := hasher.Make("password")
	s.NoError(err)
	s.False(hasher.NeedsRehash(value))
	s.True(hasher.NeedsRehash("hash"))
	s.True(hasher.NeedsRehash("hashhash"))
	s.True(hasher.NeedsRehash("$2a$10$jZc6dp3gRgotMAyyZrHYYunokxCVerHkokFs9NkGsI9gJ4VtRkt1m"))
	s.True(hasher.NeedsRehash("$2a$04$T1AxuaXbZiIWy3r1eLlp8OPguGw8XQwUBKwU6w88oTkXBp6RH5tnK"))
}
