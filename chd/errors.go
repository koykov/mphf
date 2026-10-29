package chd

import "errors"

var (
	ErrLFRanges  = errors.New("load factor must be in range [0.0..1.0]")
	ErrEmpty     = errors.New("keyset is empty")
	ErrDuplicate = errors.New("keyset contains duplicates")
)
