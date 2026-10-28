package chd

import "github.com/koykov/mphf"

type Option[T mphf.Key] func(x *chd[T])

func WithSeed[T mphf.Key](seed T) Option[T] {
	return func(x *chd[T]) {
		x.seed = seed
	}
}

func WithSalt[T mphf.Key](salt T) Option[T] {
	return func(x *chd[T]) {
		x.salt = salt
	}
}
