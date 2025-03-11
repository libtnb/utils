package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientOriginalExtension(t *testing.T) {
	assert.Equal(t, ClientOriginalExtension("file.go"), "go")
}

func TestContain(t *testing.T) {
	assert.True(t, Contain("file.go", "package file"))
}

func TestWriteAndGetLineNum(t *testing.T) {
	pwd, _ := os.Getwd()
	path := pwd + "/goravel/goravel.txt"
	assert.Nil(t, Write(path, []byte(`goravel`)))
	assert.Nil(t, WriteString(path, `goravel`))

	assert.Equal(t, 1, GetLineNum(path))
	assert.True(t, Exists(path))
	assert.Nil(t, Remove(path))
	assert.Nil(t, Remove(pwd+"/goravel"))
}

func TestExists(t *testing.T) {
	assert.True(t, Exists("file.go"))
}

func TestExtension(t *testing.T) {
	extension, err := Extension("file.go")
	assert.Nil(t, err)
	assert.Equal(t, "txt", extension)
}

func TestLastModified(t *testing.T) {
	ti, err := LastModified("file.go", "UTC")
	assert.Nil(t, err)
	assert.NotNil(t, ti)
}

func TestMimeType(t *testing.T) {
	mimeType, err := MimeType("file.go")
	assert.Nil(t, err)
	assert.Equal(t, "text/plain; charset=utf-8", mimeType)
}

func TestRemove(t *testing.T) {
	pwd, _ := os.Getwd()
	path := pwd + "/goravel/goravel.txt"
	assert.Nil(t, WriteString(path, `goravel`))

	assert.Nil(t, Remove(path))
	assert.Nil(t, Remove(pwd+"/goravel"))
}

func TestSize(t *testing.T) {
	size, err := Size("file.go")
	assert.Nil(t, err)
	assert.True(t, size > 100)
}
