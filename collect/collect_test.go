package collect

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    any
		expected any
	}{
		{
			name:     "string slice to int slice",
			input:    []string{"1", "2", "3"},
			expected: []int{1, 2, 3},
		},
		{
			name:     "int slice to string slice",
			input:    []int{1, 2, 3},
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "float64 slice to int slice",
			input:    []float64{1.1, 2.2, 3.3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "string slice to bool slice",
			input:    []string{"true", "false", "1", "0"},
			expected: []bool{true, false, true, false},
		},
		{
			name:     "int slice to float64 slice",
			input:    []int{1, 2, 3},
			expected: []float64{1.0, 2.0, 3.0},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			switch input := tt.input.(type) {
			case []string:
				switch expected := tt.expected.(type) {
				case []int:
					result := To[string, int](input)
					assert.Equal(t, expected, result)
				case []bool:
					result := To[string, bool](input)
					assert.Equal(t, expected, result)
				}
			case []int:
				switch expected := tt.expected.(type) {
				case []string:
					result := To[int, string](input)
					assert.Equal(t, expected, result)
				case []float64:
					result := To[int, float64](input)
					assert.Equal(t, expected, result)
				}
			case []float64:
				switch expected := tt.expected.(type) {
				case []int:
					result := To[float64, int](input)
					assert.Equal(t, expected, result)
				}
			}
		})
	}
}

func TestToE(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		input       any
		expected    any
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid string slice to int slice",
			input:       []string{"1", "2", "3"},
			expected:    []int{1, 2, 3},
			expectError: false,
		},
		{
			name:        "invalid string slice to int slice",
			input:       []string{"1", "invalid", "3"},
			expected:    []int(nil),
			expectError: true,
			errorMsg:    "failed to convert value at index 1:",
		},
		{
			name:        "valid float64 slice to int slice",
			input:       []float64{1.0, 2.0, 3.0},
			expected:    []int{1, 2, 3},
			expectError: false,
		},
		{
			name:        "valid string slice to bool slice",
			input:       []string{"true", "false", "1", "0"},
			expected:    []bool{true, false, true, false},
			expectError: false,
		},
		{
			name:        "invalid string slice to bool slice",
			input:       []string{"true", "invalid", "1"},
			expected:    []bool(nil),
			expectError: true,
			errorMsg:    "failed to convert value at index 1:",
		},
		{
			name:        "overflow test uint8 to int8",
			input:       []uint8{200, 100, 50},
			expected:    []int8(nil),
			expectError: true,
			errorMsg:    "failed to convert value at index 0:",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var result any
			var err error

			switch input := tt.input.(type) {
			case []string:
				switch tt.expected.(type) {
				case []int:
					result, err = ToE[string, int](input)
				case []bool:
					result, err = ToE[string, bool](input)
				}
			case []float64:
				switch tt.expected.(type) {
				case []int:
					result, err = ToE[float64, int](input)
				}
			case []uint8:
				switch tt.expected.(type) {
				case []int8:
					result, err = ToE[uint8, int8](input)
				}
			}

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestCount(t *testing.T) {
	count := Count([]int{1, 5, 1})
	assert.Equal(t, 3, count)
}

func TestCountBy(t *testing.T) {
	count := CountBy([]int{1, 5, 1}, func(i int) bool {
		return i < 4
	})
	assert.Equal(t, 2, count)
}

func TestDiff(t *testing.T) {
	diff := Diff([]int{1, 2, 3}, []int{2, 3, 4})
	assert.Equal(t, []int{1}, diff)

	diffStr := Diff([]string{"a", "b", "c"}, []string{"b", "c", "d"})
	assert.Equal(t, []string{"a"}, diffStr)
}

func TestEach(t *testing.T) {
	Each([]string{"hello", "world"}, func(x string, i int) {
		if i == 0 {
			assert.Equal(t, "hello", x)
		} else {
			assert.Equal(t, "world", x)
		}
	})
	Each([]int{0, 1, 2, 3}, func(x int, i int) {
		assert.Equal(t, i, x)
	})
}

func TestFilter(t *testing.T) {
	even := Filter([]int{1, 2, 3, 4}, func(x int, index int) bool {
		return x%2 == 0
	})
	assert.Equal(t, []int{2, 4}, even)
}

func TestFirst(t *testing.T) {
	first := First([]int{1, 2, 3})
	assert.Equal(t, 1, *first)
	first2 := First([]int{})
	assert.Nil(t, first2)
}

func TestGroupBy(t *testing.T) {
	groups := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	assert.Equal(t, map[int][]int{0: {0, 3}, 1: {1, 4}, 2: {2, 5}}, groups)
}

func TestKeys(t *testing.T) {
	keys1 := Keys[int, string](map[int]string{1: "foo", 2: "bar"})
	keys2 := Keys[string, int](map[string]int{"foo": 1, "bar": 2})
	sort.Ints(keys1)
	sort.Strings(keys2)
	assert.Equal(t, []int{1, 2}, keys1)
	assert.Equal(t, []string{"bar", "foo"}, keys2)
}

func TestMap(t *testing.T) {
	results1 := Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	results2 := Map([]int64{1, 2, 3, 4}, func(x int64, _ int) int64 {
		return x + 1
	})
	assert.Equal(t, []string{"1", "2", "3", "4"}, results1)
	assert.Equal(t, []int64{2, 3, 4, 5}, results2)
}

func TestMax(t *testing.T) {
	max1 := Max([]int{1, 2, 3})
	max2 := Max([]int{})
	assert.Equal(t, 3, max1)
	assert.Equal(t, 0, max2)
}

func TestMerge(t *testing.T) {
	mergedMaps1 := Merge[string, int](
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
	)
	mergedMaps2 := Merge[int, string](
		map[int]string{1: "a", 2: "b"},
		map[int]string{2: "b", 4: "c"},
	)
	assert.Equal(t, map[string]int{"a": 1, "b": 3, "c": 4}, mergedMaps1)
	assert.Equal(t, map[int]string{1: "a", 2: "b", 4: "c"}, mergedMaps2)
}

func TestMin(t *testing.T) {
	min1 := Min([]int{1, 2, 3})
	min2 := Min([]int{})
	assert.Equal(t, 1, min1)
	assert.Equal(t, 0, min2)
}

func TestReverse(t *testing.T) {
	reverseOrder1 := Reverse([]int{0, 1, 2, 3, 4, 5})
	reverseOrder2 := Reverse([]string{"a", "b", "c", "d"})
	assert.Equal(t, []int{5, 4, 3, 2, 1, 0}, reverseOrder1)
	assert.Equal(t, []string{"d", "c", "b", "a"}, reverseOrder2)
}

func TestShuffle(t *testing.T) {
	shuffled1 := Shuffle([]int{0, 1, 2, 3, 4, 5})
	shuffled2 := Shuffle([]string{"a", "b", "c", "d"})
	assert.Len(t, shuffled1, 6)
	assert.Len(t, shuffled2, 4)
}

func TestSplit(t *testing.T) {
	result := Split([]int{0, 1, 2, 3, 4, 5}, 2)
	result1 := Split([]int{0, 1, 2, 3, 4, 5, 6}, 2)
	result2 := Split([]int{}, 2)
	result3 := Split([]int{0}, 2)
	result4 := Split([]string{"a", "b", "c", "d"}, 2)

	assert.Equal(t, [][]int{{0, 1}, {2, 3}, {4, 5}}, result)
	assert.Equal(t, [][]int{{0, 1}, {2, 3}, {4, 5}, {6}}, result1)
	assert.Equal(t, [][]int{}, result2)
	assert.Equal(t, [][]int{{0}}, result3)
	assert.Equal(t, [][]string{{"a", "b"}, {"c", "d"}}, result4)
}

func TestSum(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	sum := Sum(list)
	assert.Equal(t, 15, sum)
}

func TestUnique(t *testing.T) {
	uniqValues1 := Unique([]int{1, 2, 2, 1})
	uniqValues2 := Unique([]string{"a", "b"}, []string{"b", "a"})
	assert.Equal(t, []int{1, 2}, uniqValues1)
	assert.Equal(t, []string{"a", "b"}, uniqValues2)
}

func TestValues(t *testing.T) {
	values1 := Values[string, int](map[string]int{"foo": 1, "bar": 2})
	values2 := Values[int, string](map[int]string{1: "foo", 2: "bar"})
	sort.Ints(values1)
	sort.Strings(values2)
	assert.Equal(t, []int{1, 2}, values1)
	assert.Equal(t, []string{"bar", "foo"}, values2)
}
