package str

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StringTestSuite struct {
	suite.Suite
}

func TestStringTestSuite(t *testing.T) {
	suite.Run(t, &StringTestSuite{})
}

func (s *StringTestSuite) TestEscape() {
	s.Equal("Hello, world!", Escape("Hello, world!"))
	s.Equal("&lt;&gt;&#34;&#39;&amp;=?%# \t\n\r", Escape("<>\"'&=?%# \t\n\r"))
}

func (s *StringTestSuite) TestMD5() {
	s.Equal("5eb63bbbe01eeed093cb22bb8f5acdc3", MD5("hello world"))
	s.Equal("d41d8cd98f00b204e9800998ecf8427e", MD5(""))
}

func (s *StringTestSuite) TestSHA256() {
	s.Equal("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9", SHA256("hello world"))
	s.Equal("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", SHA256(""))
}

func (s *StringTestSuite) TestSHA3() {
	s.Equal("644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938", SHA3("hello world"))
	s.Equal("a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a", SHA3(""))
}

func (s *StringTestSuite) TestIsPhone() {
	s.True(IsPhone("13800138000"))
	s.False(IsPhone("1234567890"))
	s.False(IsPhone("1380013800a"))
	s.False(IsPhone("23800138000"))
	s.False(IsPhone(""))
}

func (s *StringTestSuite) TestIsEmail() {
	s.True(IsEmail("test@example.com"))
	s.False(IsEmail("test@.com"))
	s.False(IsEmail("test@com"))
	s.False(IsEmail("test.com"))
	s.False(IsEmail(""))
}

func (s *StringTestSuite) TestIsURL() {
	s.True(IsURL("http://example.com"))
	s.True(IsURL("https://example.com"))
	s.False(IsURL("ftp://example.com"))
	s.False(IsURL("http://"))
	s.False(IsURL("example.com"))
	s.False(IsURL(""))
}

func (s *StringTestSuite) TestCut() {
	tests := []struct {
		str      string
		begin    string
		end      string
		expected string
	}{
		{"hello[world]", "[", "]", "world"},
		{"hello[world]hello[earth]", "[", "]", "world"},
		{"hello[world", "[", "]", ""},
		{"", "[", "]", ""},
		{"hello[world]", "", "]", ""},
		{"hello[world]", "[", "", ""},
		{"hello[[world]]", "[", "]", "[world"},
		{"Hello, world!", "Hello, ", "!", "world"},
		{"Hello, world!", "world", "Hello", ""},
		{"Hello, world!", "foo", "bar", ""},
		{"Hello, world!", "Hello", "Hello", ""},
		{"", "Hello", "world", ""},
	}

	for _, test := range tests {
		result := Cut(test.str, test.begin, test.end)
		s.Equal(test.expected, result)
	}
}

func (s *StringTestSuite) TestSubstr() {
	tests := []struct {
		str      string
		start    int
		length   []int
		expected string
	}{
		{"hello世界", 0, []int{5}, "hello"},
		{"hello世界", 5, []int{2}, "世界"},
		{"hello世界", -2, nil, "世界"},
		{"hello世界", -2, []int{1}, "世"},
		{"", 0, []int{5}, ""},
		{"hello", 10, []int{5}, ""},
		{"hello", -10, []int{5}, "hello"},
		{"Hello, world!", 7, []int{5}, "world"},
		{"Golang", 10, nil, ""},
		{"Goroutines", -5, []int{4}, "tine"},
		{"Unicode", 2, []int{-3}, "ic"},
		{"Testing", 1, []int{10}, "esting"},
		{"", 0, []int{5}, ""},
		{"你好，世界！", 3, []int{3}, "世界！"},
	}

	for _, test := range tests {
		result := Substr(test.str, test.start, test.length...)
		s.Equal(test.expected, result)
	}
}

func (s *StringTestSuite) TestRandom() {
	s.Len(Random(10), 10)
	s.Empty(Random(0))
	s.Panics(func() {
		Random(-1)
	})
}

func (s *StringTestSuite) TestRandomN() {
	s.Len(RandomN(10), 10)
	s.Empty(RandomN(0))
	s.Panics(func() {
		RandomN(-1)
	})
}

func (s *StringTestSuite) TestCase2Camel() {
	s.Equal("GoravelFramework", Case2Camel("goravel_framework"))
	s.Equal("GoravelFramework1", Case2Camel("goravel_framework1"))
	s.Equal("GoravelFramework", Case2Camel("GoravelFramework"))
}

func (s *StringTestSuite) TestCamel2Case() {
	s.Equal("goravel_framework", Camel2Case("GoravelFramework"))
	s.Equal("goravel_framework1", Camel2Case("GoravelFramework1"))
	s.Equal("goravel_framework", Camel2Case("goravel_framework"))
}
