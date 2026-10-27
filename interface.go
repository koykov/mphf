package mphf

type Key interface {
	~uint64 | ~uint32 | ~uint16 | ~uint8 | ~uint | ~uintptr
}

type Interface[T Key] interface {
	Len() int
	Size() int
	Index(key T) uint64
}
