package chd

import "github.com/koykov/mphf"

type Hasher[T mphf.Key] interface {
	Hash(key T) uint64
}
