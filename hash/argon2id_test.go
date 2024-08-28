package hash

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Argon2idTestSuite struct {
	suite.Suite
}

func TestArgon2idTestSuite(t *testing.T) {
	suite.Run(t, &Argon2idTestSuite{})
}

func (s *Argon2idTestSuite) TestMake() {
	hasher := NewArgon2id()
	hash, err := hasher.Make("password")
	s.NotEmpty(hash)
	s.NoError(err)
}

func (s *Argon2idTestSuite) TestCheck() {
	hasher := NewArgon2id()
	value, err := hasher.Make("password")
	s.NoError(err)
	s.True(hasher.Check("password", value))
	s.False(hasher.Check("password1", value))
	s.False(hasher.Check("password", "hash"))
	s.False(hasher.Check("password", "hashhash"))
	s.False(hasher.Check("password", "$argon2id$v=20$m=16,t=2,p=1$dTltTmtGb0JmNE9Zb0lTeQ$2lHJsAodBnV4u7j39gj7Uw"))
	s.False(hasher.Check("password", "$argon2id$v=$m=16,t=2,p=1$dTltTmtGb0JmNE9Zb0lTeQ$2lHJsAodBnV4u7j39gj7Uw"))
	s.False(hasher.Check("password", "$argon2id$v=19$m=16,t=2$dTltTmtGb0JmNE9Zb0lTeQ$2lHJsAodBnV4u7j39gj7Uw"))
	s.False(hasher.Check("password", "$argon2id$v=19$m=16,t=2,p=1$dTltTmtGb0JmNE9Zb0lTeQ$123456"))
	s.False(hasher.Check("password", "$argon2id$v=19$m=16,t=2,p=1$123456$2lHJsAodBnV4u7j39gj7xx"))
}

func (s *Argon2idTestSuite) TestConfigurationOverride() {
	hasher := NewArgon2id()
	value := "$argon2id$v=19$m=65536,t=8,p=1$UEFLYmZsbW9BMVIwcDhnZw$BSlJvzKjHUrAwwTOLauybA"
	s.True(hasher.Check("rat", value))
	s.True(hasher.NeedsRehash(value))
}

func (s *Argon2idTestSuite) TestNeedsRehash() {
	hasher := NewArgon2id()
	value, err := hasher.Make("password")
	s.NoError(err)
	s.False(hasher.NeedsRehash(value))
	s.True(hasher.NeedsRehash("hash"))
	s.True(hasher.NeedsRehash("hashhash"))
	s.True(hasher.NeedsRehash("$argon2id$v=$m=16,t=2,p=1$dTltTmtGb0JmNE9Zb0lTeQ$2lHJsAodBnV4u7j39gj7Uw"))
	s.True(hasher.NeedsRehash("$argon2id$v=19$m=16,t=2$dTltTmtGb0JmNE9Zb0lTeQ$2lHJsAodBnV4u7j39gj7Uw"))
}
