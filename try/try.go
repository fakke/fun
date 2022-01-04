package try

type Try[A any] interface {
	IsFailure() bool
	IsSuccess() bool
	Error() error
	Value() A
	ValueOr(ifError A) A
}

func Of[A any](a A, e error) Try[A] {
	if e != nil {
		return Failure[A](e)
	}
	return Success(a)
}

type success[A any] struct {
	a A
}

func Success[A any](a A) Try[A] {
	return success[A]{a: a}
}

func (r success[A]) IsFailure() bool {
	return false
}

func (r success[A]) IsSuccess() bool {
	return true
}

func (r success[A]) Error() error {
	panic("Success.Error()")
}

func (r success[A]) Value() A {
	return r.a
}

func (r success[A]) ValueOr(ifError A) A {
	return r.Value()
}

type failure[A any] struct {
	e error
}

func Failure[A any](e error) Try[A] {
	if e == nil {
		panic("Failure(nil)")
	}
	return failure[A]{e: e}
}

func (e failure[A]) IsFailure() bool {
	return true
}

func (e failure[A]) IsSuccess() bool {
	return false
}

func (e failure[A]) Error() error {
	return e.e
}

func (e failure[A]) Value() A {
	panic("Failure.Value()")
}

func (e failure[A]) ValueOr(ifError A) A {
	return ifError
}

func FoldL[A, B any](t Try[A], ifError B, f func(A) B) B {
	if t.IsFailure() {
		return ifError
	}
	return f(t.Value())
}

func FoldR[A, B any](t Try[A], ifError func() B, f func(A) B) B {
	if t.IsFailure() {
		return ifError()
	}
	return f(t.Value())
}

func Bind[A, B any](t Try[A], f func(A) Try[B]) Try[B] {
	return FoldR(t, func() Try[B] { return Failure[B](t.Error()) }, func(a A) Try[B] {
		return f(t.Value())
	})
}

func Map[A, B any](t Try[A], f func(A) B) Try[B] {
	return Bind(t, func(a A) Try[B] {
		return Success(f(a))
	})
}
