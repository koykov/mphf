package chd

import "errors"

var (
	ErrEmpty     = errors.New("keyset is empty")
	ErrDuplicate = errors.New("keyset contains duplicates")
)
