package str

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"math/big"
	"net/url"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/go-rat/utils/convert"
)

const (
	Letters = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers = "0123456789"
)

// Escape 转义字符串
func Escape(str string) string {
	return template.HTMLEscapeString(str)
}

// MD5 生成字符串的 MD5 值
func MD5(str string) string {
	sum := md5.Sum(convert.UnsafeBytes(str))
	dst := make([]byte, hex.EncodedLen(len(sum)))
	hex.Encode(dst, sum[:])
	return convert.UnsafeString(dst)
}

// SHA256 生成字符串的 SHA256 值
func SHA256(str string) string {
	sum := sha256.Sum256(convert.UnsafeBytes(str))
	dst := make([]byte, hex.EncodedLen(len(sum)))
	hex.Encode(dst, sum[:])
	return convert.UnsafeString(dst)
}

// SHA3 生成字符串的 SHA3 值
func SHA3(str string) string {
	sum := sha3.Sum256(convert.UnsafeBytes(str))
	dst := make([]byte, hex.EncodedLen(len(sum)))
	hex.Encode(dst, sum[:])
	return convert.UnsafeString(dst)
}

// IsPhone 判断是否为手机号
func IsPhone(phone string) bool {
	phoneRegex := `^1[3-9]\d{9}$`
	matched, err := regexp.MatchString(phoneRegex, phone)
	if err != nil {
		return false
	}
	return matched
}

// IsEmail 判断是否为邮箱
func IsEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false
	}
	return matched
}

// IsURL 判断字符串是否是 URL
func IsURL(str string) bool {
	unescape, err := url.QueryUnescape(str)
	if err == nil {
		str = unescape
	}
	parsed, err := url.Parse(str)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") || len(parsed.Host) == 0 {
		return false
	}

	return true
}

// Cut 返回字符串中 begin 和 end 之间的内容
// 如果未找到 begin 或 end，或者它们的位置无效，则返回空字符串。
// 例如：
//
//	Cut("hello[world]", "[", "]") 返回 "world"
//	Cut("hello[world", "[", "]") 返回 ""
func Cut(str, begin, end string) string {
	if str == "" || begin == "" || end == "" {
		return ""
	}

	bIndex := strings.Index(str, begin)
	if bIndex == -1 {
		return ""
	}

	afterBegin := bIndex + len(begin)
	eIndex := strings.Index(str[afterBegin:], end)
	if eIndex == -1 {
		return ""
	}
	eIndex += afterBegin

	if bIndex >= eIndex || afterBegin > eIndex {
		return ""
	}

	return str[afterBegin:eIndex]
}

// Substr 返回字符串的子串
// start 表示起始位置，可以是负数（表示从末尾开始计数）
// length 可选，表示子串长度。如果省略，则返回到字符串末尾。
// 如果 length 为负数，表示从末尾往前数的位置。
// 例如：
//
//	Substr("hello世界", 0, 5) 返回 "hello"
//	Substr("hello世界", -2) 返回 "世界"
func Substr(str string, start int, length ...int) string {
	if str == "" {
		return ""
	}

	runes := []rune(str)
	strLen := len(runes)

	if start < 0 {
		start = strLen + start
	}
	if start < 0 {
		start = 0
	}
	if start >= strLen {
		return ""
	}

	end := strLen
	if len(length) > 0 {
		if length[0] >= 0 {
			end = start + length[0]
		} else {
			end = strLen + length[0]
		}
	}

	if end < start {
		return ""
	}
	if end > strLen {
		end = strLen
	}

	return string(runes[start:end])
}

// Random 生成长度为 length 的随机字符串
func Random(length int) string {
	sb := new(strings.Builder)
	sb.Grow(length)

	lettersLen := big.NewInt(int64(len(Letters)))
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, lettersLen)
		if err != nil {
			panic(err)
		}
		sb.WriteByte(Letters[idx.Int64()])
	}

	return sb.String()
}

// RandomN 生成长度为 length 随机数字字符串
func RandomN(length int) string {
	sb := new(strings.Builder)
	sb.Grow(length)

	numbersLen := big.NewInt(int64(len(Numbers)))
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, numbersLen)
		if err != nil {
			panic(err)
		}
		sb.WriteByte(Numbers[idx.Int64()])
	}

	return sb.String()
}

// Case2Camel 下划线命名转驼峰命名
func Case2Camel(name string) string {
	names := strings.Split(name, "_")
	sb := new(strings.Builder)

	for _, item := range names {
		buffer := new(bytes.Buffer)
		for i, r := range item {
			if i == 0 {
				buffer.WriteRune(unicode.ToUpper(r))
			} else {
				buffer.WriteRune(r)
			}
		}
		sb.Write(buffer.Bytes())
	}

	return sb.String()
}

// Camel2Case 驼峰命名转下划线命名
func Camel2Case(name string) string {
	sb := new(strings.Builder)

	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
