package chd

import (
	"math"
	"math/bits"
	"slices"
	"sort"

	"github.com/koykov/bitvector"
	"github.com/koykov/mphf"
)

const defaultLoadFactor = .85

type chd[T mphf.Key] struct {
	lf     float64
	seed   T
	seeds  []uint32
	salt   T
	hasher Hasher[T]
}

func New[T mphf.Key](keys []T, options ...Option[T]) (mphf.Interface[T], error) {
	if len(keys) == 0 {
		return nil, ErrEmpty
	}

	x := chd[T]{}
	for _, fn := range options {
		fn(&x)
	}
	if x.lf < 0 || x.lf > 1 {
		return nil, ErrLFRanges
	}
	if x.lf == 0 {
		x.lf = defaultLoadFactor
	}

	m := uint64(float64(len(keys)) / x.lf)
	m = 1 << bits.Len64(m) // get nearest pow2

	bb := make(buckets[T], m)
	for i := uint64(0); i < m; i++ {
		bb[i].slot = T(i)
	}
	for i := 0; i < len(keys); i++ {
		pos := x.hasher.Hash(keys[i])
		bb[pos].keys = append(bb[pos].keys, keys[i])
	}

	slices.Sort(keys)
	var pkey T
	for i, key := range keys {
		if i > 0 && key == pkey {
			return nil, ErrDuplicate
		}
		pkey = key
		h := x.hasher.Hash(key)
		b := &bb[h]
		b.keys = append(b.keys, key)
	}

	sort.Sort(bb)

	vec0, _ := bitvector.NewVector(m)
	vec1, _ := bitvector.NewVector(m)
	seeds := make([]uint32, m)
	var mxs uint32
	for i := uint64(0); i < m; i++ {
		b := &bb[i]
		for j := 0; j < math.MaxUint16; j++ {
			vec1.Reset()
			for _, k := range b.keys {
				h := x.hasher.Hash(k)
				if vec0.Get(h) != 0 || vec1.Get(h) != 0 {
					// use next seed
				}
				vec1.Set(h)
			}
			vec0.Merge(vec1)
			seeds[b.slot] = uint32(j)
			mxs = max(mxs, uint32(j))
			// use next bucket
		}
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
