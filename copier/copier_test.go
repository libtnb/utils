package copier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SampleStruct struct {
	Name string
	Age  int
}

func TestCopyStruct(t *testing.T) {
	from := SampleStruct{Name: "John", Age: 30}
	to, err := Copy[SampleStruct](from)
	assert.NoError(t, err)
	assert.Equal(t, from, *to)
}

func TestCopyMap(t *testing.T) {
	from := map[string]interface{}{"Name": "John", "Age": float64(30)}
	to, err := Copy[map[string]any](from)
	assert.NoError(t, err)
	assert.Equal(t, from, *to)
}

func TestCopyInvalidJSON(t *testing.T) {
	from := make(chan int)
	_, err := Copy[SampleStruct](from)
	assert.Error(t, err)
}
