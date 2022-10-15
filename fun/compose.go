package fun

import "github.com/fakke/fun/slice"

func Chain[X any](fs ...func(X) X) func(X) X {
	return slice.FoldL(fs[1:], fs[0], Compose[X, X, X])
}

func Compose[X, Y, Z any](fxy func(X) Y, fyz func(Y) Z) func(X) Z {
	return func(x X) Z {
		return fyz(fxy(x))
	}
}

func Compose2[X1, X2, X3 any](f1 func(X1) X2, f2 func(X2) X3) func(X1) X3 {
	return Compose(f1, f2)
}

func Compose3[X1, X2, X3, X4 any](f1 func(X1) X2, f2 func(X2) X3, f3 func(X3) X4) func(X1) X4 {
	return Compose(Compose(f1, f2), f3)
}

func Compose4[X1, X2, X3, X4, X5 any](f1 func(X1) X2, f2 func(X2) X3, f3 func(X3) X4, f4 func(X4) X5) func(X1) X5 {
	return Compose(Compose(Compose(f1, f2), f3), f4)
}
