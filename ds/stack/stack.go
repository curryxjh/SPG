package stack

import (
	"SPL/ds/container"
	"SPL/utils/sync"
)

var (
	defaultLocker sync.FakeLocker
)

type Options[T any] struct {
	locker    sync.Locker
	container container.Container[T]
}

type Option[T any] func(options *Options[T])
