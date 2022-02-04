package cat

type Traversable[V any] interface {
	ForEach(do func(V))
	ForEachIndex(do func(int))
	ForEachIndexed(do func(int, V))
}

type Appendable[V any] interface {
	Append(v V)
	AppendAll(values ...V)
	AppendFrom(src Traversable[V])
}
