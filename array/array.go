package array

import (
	. "github.com/fakke/fun/cat"
)

// Array interface
type Array[V any] interface {
	Traversable[V]
	Appendable[V]
	Cap() int
	Len() int
	At(i int) V
	SetAt(i int) func(v V)
	InsertAt(i int) func(v V)
	RemoveAt(i int)
}

// array struct is the default implementation of the Array interface
type array[V any] struct {
	values []V
}

// Of wraps the vararg-specified array / slice
func Of[V any](values ...V) Array[V] {
	return &array[V]{values: values}
}

// OfSize creates an array of specified length and capacity
func OfSize[V any](len, cap int) Array[V] {
	return &array[V]{values: make([]V, len, cap)}
}

// OfLen creates an array of specified length
func OfLen[V any](len int) Array[V] {
	return &array[V]{values: make([]V, len)}
}

// OfCap creates an array of specified capacity
func OfCap[V any](cap int) Array[V] {
	return OfSize[V](0, cap)
}

// ForEach to implement Traversable
func (a *array[V]) ForEach(do func(V)) {
	for _, v := range a.values {
		do(v)
	}
}

// ForEachIndex to implement Traversable
func (a *array[V]) ForEachIndex(do func(int)) {
	for i, _ := range a.values {
		do(i)
	}
}

// ForEachIndexed to implement Traversable
func (a *array[V]) ForEachIndexed(do func(int, V)) {
	for i, v := range a.values {
		do(i, v)
	}
}

// Append to implement Appendable
func (a *array[V]) Append(v V) {
	a.values = append(a.values, v)
}

// AppendAll to implement Appendable
func (a *array[V]) AppendAll(values ...V) {
	a.values = append(a.values, values...)
}

// AppendFrom to implement Appendable
func (a *array[V]) AppendFrom(src Traversable[V]) {
	src.ForEach(a.Append)
}

// Cap returns the capacity of the underlying array/slice
func (a *array[V]) Cap() int {
	return cap(a.values)
}

// Len return the length of the underlying array/slice
func (a *array[V]) Len() int {
	return len(a.values)
}

// At return the value at i
func (a *array[V]) At(i int) V {
	return a.values[i]
}

// SetAt updates the value at index i
func (a *array[V]) SetAt(i int) func(v V) {
	return func(v V) {
		a.values[i] = v
	}
}

// InsertAt inserts a value before the i-th value (in-place)
func (a *array[V]) InsertAt(i int) func(V) {
	return func(v V) {
		if i == a.Len() {
			a.Append(v)
		} else {
			a.values = append(a.values[:i+1], a.values[i:]...)
			a.values[i] = v
		}
	}
}

// RemoveAt removes the i-th value (in-place)
func (a *array[V]) RemoveAt(i int) {
	if i == a.Len() {
		a.values = a.values[:a.Len()-1]
	} else {
		a.values = append(a.values[:i], a.values[i+1:]...)
	}
}

// Map creates an array of B with values mapped via f
func Map[A, B any](a Traversable[A], f func(A) B) Array[B] {
	if arr, isArray := a.(Array[A]); isArray {
		return mapArray(arr, f)
	} else {
		return mapTraversable(a, f)
	}
}

// mapArray implements Map optimized for Array
func mapArray[A, B any](a Array[A], f func(A) B) Array[B] {
	b := OfSize[B](a.Len(), a.Cap())
	a.ForEachIndexed(func(i int, v A) {
		b.SetAt(i)(f(v))
	})
	return b
}

// mapTraversable implements Map optimized for Traversable
func mapTraversable[A, B any](a Traversable[A], f func(A) B) Array[B] {
	b := OfLen[B](0)
	a.ForEach(func(v A) {
		b.Append(f(v))
	})
	return b
}
