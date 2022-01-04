package slice

func Index[X any](xs []X, f func(X) bool) int {
	for i, x := range xs {
		if f(x) {
			return i
		}
	}
	return -1
}

func IndexOf[X comparable](xs []X, x X) int {
	return Index(xs, func(x_ X) bool {
		return x_ == x
	})
}

func FoldL[X, Y any](xs []X, z Y, op func(Y, X) Y) Y {
	result := z
	for i := 0; i < len(xs); i++ {
		result = op(result, xs[i])
	}
	return result
}

func FoldR[X, Y any](xs []X, z Y, op func(X, Y) Y) Y {
	result := z
	for i := len(xs) - 1; i >= 0; i-- {
		result = op(xs[i], result)
	}
	return result
}

func ScanL[X, Y any](xs []X, z Y, op func(Y, X) Y) []Y {
	result := make([]Y, len(xs)+1)
	result[0] = z
	for i := 0; i < len(xs); i++ {
		result[i+1] = op(result[i], xs[i])
	}
	return result
}

func ScanR[X, Y any](xs []X, z Y, op func(X, Y) Y) []Y {
	result := make([]Y, len(xs)+1)
	result[len(xs)] = z
	for i := len(xs) - 1; i >= 0; i-- {
		result[i] = op(xs[i], result[i+1])
	}
	return result
}

func Map[X, Y any](xs []X, f func(X) Y) []Y {
	ys := make([]Y, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func Filter[X any](xs []X, f func(X) bool) []X {
	xs_ := make([]X, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			xs_ = append(xs_, x)
		}
	}
	return xs_
}

func FilterNot[X any](xs []X, f func(X) bool) []X {
	return Filter(xs, func(x X) bool {
		return !f(x)
	})
}

func ForAll[X any](xs []X, pred func(x X) bool) bool {
	notPred := func(x X) bool { return !pred(x) }
	return Index(xs, notPred) == -1
}

func ForAny[X any](xs []X, pred func(x X) bool) bool {
	return Index(xs, pred) != -1
}

func Contains[X comparable](xs []X, x X) bool {
	return IndexOf(xs, x) != -1
}
