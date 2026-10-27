package mphf

type Option func(b *base)

func WithSeed(seed uint64) Option {
	return func(b *base) {
		b.seed = seed
	}
}
