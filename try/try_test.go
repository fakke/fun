package try

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

var errTest = errors.New("test")
var errDomain = errors.New("domain")
var errDivZero = errors.New("division by zero")

var ignore = func(a any) {}

var izero = func() int { return 0 }
var imul10 = func(a int) int { return a * 10 }
var idiv = func(a, b int) (int, error) {
	if b == 0 {
		return 0, errDivZero
	}
	return a / b, nil
}
var idivTry = func(a, b int) Try[int] {
	if b == 0 {
		return Failure[int](errDivZero)
	}
	return Success(a / b)
}
var strlen = func(s string) int { return len(s) }
var sempty = func() string { return "" }
var itoa = strconv.Itoa
var atoi = strconv.Atoi
var atoiTry = func(s string) Try[int] {
	return Of(atoi(s))
}
var isqrt = func(a int) (int, error) {
	if a < 0 {
		return 0, errDomain
	}
	return int(math.Sqrt(float64(a))), nil
}
var isqrtTry = func(a int) Try[int] {
	if a < 0 {
		return Failure[int](errDomain)
	}
	return Success(int(math.Sqrt(float64(a))))
}

func TestOf(t *testing.T) {
	assert.Equal(t, Of(idiv(10, 2)), Success(5))
	assert.Equal(t, Of(idiv(10, 0)), Failure[int](errDivZero))
}

func TestResult(t *testing.T) {
	assert.Equal(t, Success(izero()), Success(0))
	assert.Equal(t, Success(imul10(3)), Success(30))
}

func TestError(t *testing.T) {
	assert.Equal(t, Failure[any](errTest), Failure[any](errTest))
	assert.Panics(t, func() {
		Failure[any](nil)
	})
}

func TestResult_Error(t *testing.T) {
	assert.Panics(t, func() {
		ignore(Success(42).Error())
	})
	assert.Panics(t, func() {
		ignore(Success("test").Error())
	})
}

func TestResult_Result(t *testing.T) {
	assert.Equal(t, idivTry(10, 5), Success(2))
	assert.Equal(t, atoiTry("42"), Success(42))
}

func TestResult_IsError(t *testing.T) {
	assert.False(t, idivTry(10, 5).IsFailure())
	assert.False(t, atoiTry("42").IsFailure())
}

func TestResult_IsResult(t *testing.T) {
	assert.True(t, idivTry(10, 5).IsSuccess())
	assert.True(t, atoiTry("42").IsSuccess())
}

func TestResult_ResultOr(t *testing.T) {
}

func TestErr_Error(t *testing.T) {
	assert.Equal(t, idivTry(0, 0).Error(), errors.New("division by zero"))
}

func TestErr_IsError(t *testing.T) {
	assert.True(t, idivTry(0, 0).IsFailure())
	assert.True(t, atoiTry("").IsFailure())
}

func TestErr_IsResult(t *testing.T) {
	assert.False(t, idivTry(0, 0).IsSuccess())
	assert.False(t, atoiTry("").IsSuccess())
}

func TestErr_Result(t *testing.T) {
	assert.Panics(t, func() {
		idivTry(0, 0).Value()
	})
	assert.Panics(t, func() {
		atoiTry("").Value()
	})
}

func TestErr_ResultOr(t *testing.T) {
	assert.Equal(t, idivTry(0, 0).ValueOr(42), 42)
	assert.Equal(t, atoiTry("").ValueOr(42), 42)
}

func TestFoldL(t *testing.T) {
	assert.Equal(t, FoldL(Success(42), "", itoa), "42")
	assert.Equal(t, FoldL(Failure[int](errTest), "", itoa), "")
	assert.Equal(t, FoldL(Success("hello"), 0, strlen), 5)
	assert.Equal(t, FoldL(Failure[string](errTest), 0, strlen), 0)
}

func TestFoldR(t *testing.T) {
	assert.Equal(t, FoldR(Success(42), sempty, itoa), "42")
	assert.Equal(t, FoldR(Failure[int](errTest), sempty, itoa), "")
	assert.Equal(t, FoldR(Success("hello"), izero, strlen), len("hello"))
	assert.Equal(t, FoldR(Failure[string](errTest), izero, strlen), 0)
}

func TestBind(t *testing.T) {
	assert.Equal(t, Bind(Success(4), isqrtTry), Success(2))
	assert.Equal(t, Bind(Success(-4), isqrtTry), Failure[int](errDomain))
}

func TestMap(t *testing.T) {
	assert.Equal(t, Map(Success(42), imul10), Success(420))
	assert.Equal(t, Map(Success(42), itoa), Success("42"))
}
