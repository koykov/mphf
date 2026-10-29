package chd

import "github.com/koykov/mphf"

type bucket[T mphf.Key] struct {
	slot T
	keys []T
}

type buckets[T mphf.Key] []bucket[T]

func (b buckets[T]) Len() int {
	return len(b)
}

func (b buckets[T]) Less(i, j int) bool {
	return len(b[i].keys) > len(b[j].keys)
}

func (b buckets[T]) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
