package rand

import (
	"math/rand"
	"sync"
	"time"
)

type MrandN struct {
	lock         sync.Mutex
	bucket       []int
	bucketLength int
	rand         *rand.Rand
}

var globalRand = NewMrandN()

func NewMrandNWithLength(max int) *MrandN {
	foo := MrandN{bucketLength: max, bucket: make([]int, max), rand: rand.New(rand.NewSource(time.Now().UnixNano()))}
	for i := 0; i < max; i++ {
		foo.bucket[i] = i
	}
	return &foo
}

func NewMrandN() *MrandN {
	return NewMrandNWithLength(1024)
}

func (mrn *MrandN) RandSlice(m, n int, randSlice []int) []int {
	mrn.lock.Lock()
	if m > mrn.bucketLength {
		mrn.bucket = make([]int, m)
	}
	for i := 0; i < n; i++ {
		idx := mrn.rand.Intn(m-i) + i
		randSlice[i] = mrn.bucket[idx]
		mrn.bucket[idx], mrn.bucket[i] = mrn.bucket[i], mrn.bucket[idx]
	}
	mrn.reset(n)
	mrn.lock.Unlock()
	return randSlice
}

func RandSlice(m, n int, randSlice []int) []int { return globalRand.RandSlice(m, n, randSlice) }

func (mrn *MrandN) Rand(m, n int) []int {
	randSlice := make([]int, n)
	mrn.RandSlice(m, n, randSlice)
	return randSlice
}
func Rand(m, n int) []int { return globalRand.Rand(m, n) }

func (mrn *MrandN) reset(n int) {
	for i := 0; i < n; i++ {
		idx := mrn.bucket[i]
		mrn.bucket[idx], mrn.bucket[i] = idx, mrn.bucket[idx]
	}
}
