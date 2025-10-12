package collect

import (
	"fmt"
	"slices"

	"github.com/samber/lo"
	"github.com/samber/lo/mutable"
	"github.com/spf13/cast"
	"golang.org/x/exp/constraints"
)

// To converts a slice of any type to a slice of another type.
func To[V any, T any](s []V) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i], _ = convert[V, T](v)
	}
	return result
}

// ToE converts a slice of any type to a slice of another type with error handling.
func ToE[V any, T any](s []V) ([]T, error) {
	result := make([]T, len(s))
	for i, v := range s {
		val, err := convert[V, T](v)
		if err != nil {
			return nil, fmt.Errorf("failed to convert value at index %d: %w", i, err)
		}
		result[i] = val
	}
	return result, nil
}

// Count counts the number of elements in the collection.
func Count[T comparable](collection []T) (count int) {
	return len(collection)
}

// CountBy counts the number of elements in the collection for which predicate is true.
func CountBy[T any](collection []T, predicate func(item T) bool) (count int) {
	return lo.CountBy(collection, predicate)
}

// Diff creates a slice of slice values not included in the other given slice.
func Diff[T comparable](list1, list2 []T) []T {
	diff, _ := lo.Difference(list1, list2)
	return diff
}

// Each iterates over elements of collection and invokes iteratee for each element.
func Each[T any](collection []T, iteratee func(item T, index int)) {
	lo.ForEach(collection, iteratee)
}

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func Filter[V any](collection []V, predicate func(item V, index int) bool) []V {
	return lo.Filter(collection, predicate)
}

// First gets the first element of collection.
func First[V any](args []V) *V {
	if len(args) > 0 {
		return &args[0]
	}
	return nil
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	return lo.GroupBy(collection, iteratee)
}

// Keys creates an array of the map keys.
func Keys[K comparable, V any](in map[K]V) []K {
	return lo.Keys(in)
}

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	return lo.Map(collection, iteratee)
}

// Max searches the maximum value of a collection.
func Max[T constraints.Ordered](collection []T) T {
	return lo.Max(collection)
}

// Merge merges multiple maps from left to right.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	return lo.Assign(maps...)
}

// Min search the minimum value of a collection.
func Min[T constraints.Ordered](collection []T) T {
	return lo.Min(collection)
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
func Reverse[T any](collection []T) []T {
	mutable.Reverse(collection)
	return collection
}

// Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
func Shuffle[T any](collection []T) []T {
	mutable.Shuffle(collection)
	return collection
}

// Split returns an array of elements split into groups the length of size. If array can't be split evenly,
func Split[T any](collection []T, size int) [][]T {
	return lo.Chunk(collection, size)
}

// Sum sums the values in a collection. If collection is empty 0 is returned.
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	return lo.Sum(collection)
}

// Unique returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
func Unique[T comparable](collections ...[]T) []T {
	return lo.Uniq(slices.Concat(collections...))
}

// Values creates an array of the map values.
func Values[K comparable, V any](in map[K]V) []V {
	return lo.Values(in)
}

// convert converts a value to another type.
func convert[V any, T any](v V) (T, error) {
	var result T
	var err error
	var converted any

	switch any(result).(type) {
	case int:
		converted, err = cast.ToIntE(v)
	case int64:
		converted, err = cast.ToInt64E(v)
	case uint:
		converted, err = cast.ToUintE(v)
	case uint64:
		converted, err = cast.ToUint64E(v)
	case float32:
		converted, err = cast.ToFloat32E(v)
	case float64:
		converted, err = cast.ToFloat64E(v)
	case string:
		converted, err = cast.ToStringE(v)
	case bool:
		converted, err = cast.ToBoolE(v)
	default:
		return result, fmt.Errorf("unsupported type conversion to %T", result)
	}

	if err != nil {
		return result, err
	}
	return converted.(T), nil
}
