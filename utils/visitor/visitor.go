package visitor

type Visitor[V any] func(value V) bool

type KvVisitor[K, V any] func(key K, value V) bool
