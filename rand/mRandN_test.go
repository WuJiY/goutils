package rand

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMRandN(t *testing.T) {
	for i := 0; i < 10; i++ {
		rst := Rand(i, i)
		fmt.Println(rst)
		for j := 0; j < i; j++ {
			if rst[j] > i {
				t.Error(rst)
			}
		}
	}
}

func BenchmarkMRandN(b *testing.B) {
	r := NewMrandN()
	for i := 0; i < b.N; i++ {
		r.Rand(100000, i)
	}
}

func BenchmarkMRandNSlice(b *testing.B) {
	r := NewMrandN()
	rs := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		r.RandSlice(100000, i, rs)
	}
}

func BenchmarkMathRand(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		r.Perm(100000)
	}
}
