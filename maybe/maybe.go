package maybe

import (
	"reflect"
)

type Maybe[A any] interface {
	IsNone() bool
	IsSome() bool
	Value() A
	ValueOr(ifNone A) A
	ToNullable() *A
}

func Of[A any](a *A) Maybe[A] {
	if a == nil {
		return None[A]()
	}
	return Some(*a)
}

type some[A any] struct {
	a A
}

func Some[A any](a A) Maybe[A] {
	return some[A]{a: a}
}

func (s some[A]) IsNone() bool {
	return false
}

func (s some[A]) IsSome() bool {
	return true
}

func (s some[A]) Value() A {
	return s.a
}

func (s some[A]) ValueOr(ifNone A) A {
	return s.Value()
}

func (s some[A]) ToNullable() *A {
	return &s.a
}

type noneType[A any] struct{}

var none Maybe[any] = noneType[any]{}

func None[A any]() Maybe[A] {
	return *(*(Maybe[A]))(reflect.ValueOf(&none).UnsafePointer())
}

func (n noneType[A]) IsNone() bool {
	return true
}

func (n noneType[A]) IsSome() bool {
	return false
}

func (n noneType[A]) Value() A {
	panic("None.Some()")
}

func (n noneType[A]) ValueOr(ifNone A) A {
	return ifNone
}

func (n noneType[A]) ToNullable() *A {
	return nil
}

func FoldL[A, B any](m Maybe[A], ifNone B, f func(A) B) B {
	if m.IsNone() {
		return ifNone
	}
	return f(m.Value())
}

func FoldR[A, B any](m Maybe[A], ifNone func() B, f func(A) B) B {
	if m.IsNone() {
		return ifNone()
	}
	return f(m.Value())
}

func Bind[A, B any](m Maybe[A], f func(A) Maybe[B]) Maybe[B] {
	return FoldR(m, None[B], func(a A) Maybe[B] {
		return f(m.Value())
	})
}

func Map[A, B any](m Maybe[A], f func(A) B) Maybe[B] {
	return Bind(m, func(a A) Maybe[B] {
		return Some(f(m.Value()))
	})
}
