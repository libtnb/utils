package file

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"

	"github.com/libtnb/utils/convert"
)

func ClientOriginalExtension(file string) string {
	return strings.ReplaceAll(filepath.Ext(file), ".", "")
}

func Contain(file string, search string) bool {
	if Exists(file) {
		data, err := os.ReadFile(file)
		if err != nil {
			return false
		}
		return strings.Contains(convert.UnsafeString(data), search)
	}

	return false
}

func Write(file string, content []byte, perm ...os.FileMode) error {
	if len(perm) == 0 {
		perm = append(perm, os.ModePerm)
	}
	if err := os.MkdirAll(filepath.Dir(file), perm[0]); err != nil {
		return err
	}

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm[0])
	if err != nil {
		return err
	}

	if _, err = f.Write(content); err != nil {
		return err
	}

	return f.Close()
}

func WriteString(file string, content string, perm ...os.FileMode) error {
	if len(perm) == 0 {
		perm = append(perm, os.ModePerm)
	}
	if err := os.MkdirAll(filepath.Dir(file), perm[0]); err != nil {
		return err
	}

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm[0])
	if err != nil {
		return err
	}

	if _, err = f.WriteString(content); err != nil {
		return err
	}

	return f.Close()
}

func Exists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// Extension Supported types: https://github.com/gabriel-vasile/mimetype/blob/master/supported_mimes.md
func Extension(file string, originalWhenUnknown ...bool) (string, error) {
	mtype, err := mimetype.DetectFile(file)
	if err != nil {
		return "", err
	}

	if mtype.String() == "" {
		if len(originalWhenUnknown) > 0 {
			if originalWhenUnknown[0] {
				return ClientOriginalExtension(file), nil
			}
		}

		return "", errors.New("unknown file extension")
	}

	return strings.TrimPrefix(mtype.Extension(), "."), nil
}

func LastModified(file, timezone string) (time.Time, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return time.Time{}, err
	}

	l, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return fileInfo.ModTime().In(l), nil
}

func MimeType(file string) (string, error) {
	mtype, err := mimetype.DetectFile(file)
	if err != nil {
		return "", err
	}

	return mtype.String(), nil
}

func Remove(file string) error {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return os.RemoveAll(file)
}

func Size(file string) (int64, error) {
	fileInfo, err := os.Open(file)
	if err != nil {
		return 0, err
	}

	fi, err := fileInfo.Stat()
	if err != nil {
		return 0, err
	}

	return fi.Size(), fileInfo.Close()
}

func GetLineNum(file string) int {
	total := 0
	f, _ := os.OpenFile(file, os.O_RDONLY, 0444)
	buf := bufio.NewReader(f)

	for {
		_, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				total++

				break
			}
		} else {
			total++
		}
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	return total
}
