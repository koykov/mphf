package chd

import (
	"slices"

	"github.com/koykov/mphf"
)

type chd[T mphf.Key] struct {
	seed T
	salt T
}

func New[T mphf.Key](keys []T, options ...Option[T]) (mphf.Interface[T], error) {
	if len(keys) == 0 {
		return nil, ErrEmpty
	}

	x := chd[T]{
		// ...
	}
	for _, fn := range options {
		fn(&x)
	}

	slices.Sort(keys)
	var pkey T
	for i, key := range keys {
		if i > 0 && key == pkey {
			return nil, ErrDuplicate
		}
		// todo add key to mphf
	}

	return &x, nil
}

func (x *chd[T]) Len() int {
	// todo implement me
	return 0
}

func (x *chd[T]) Size() int {
	// todo implement me
	return 0
}

func (x *chd[T]) Index(key T) uint64 {
	// todo implement me
	return 0
}
