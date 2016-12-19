package rand

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandByteSlice(l int, buff []byte) []byte {
	if len(buff) < l {
		buff = make([]byte, l)
	}
	for i := 0; i < l; i++ {
		buff[i] = r.Intn(1 << 8)
	}
	return buff
}
