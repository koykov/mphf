package chd

type bucket struct {
	slot uint64
	keys []uint64
}

type buckets []bucket

func (b buckets) Len() int {
	return len(b)
}

func (b buckets) Less(i, j int) bool {
	return len(b[i].keys) > len(b[j].keys)
}

func (b buckets) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
