package rand

import (
	"math/rand"
	"time"

	"github.com/niubaoshu/goutils"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandBytesA(l int, buff []byte) []byte {
	if cap(buff) < l {
		b := make([]byte, l-cap(buff))
		buff = append(buff[:], b...)
	} else {
		buff = buff[:len(buff)+l]
	}

	for i := len(buff); i < l; i++ {
		buff[i] = byte(r.Intn(1 << 8))
	}
	return buff
}

func RandBytes(l int) []byte {
	return RandBytesA(l, nil)
}

const (
	NUM = 1 << iota
	LOWER
	UPPER

	ALL = NUM | LOWER | UPPER

	num   = "1234567890"
	lower = "abcdefghijklmnopqrstuvwxyz"
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	chars = [ALL + 1][]byte{
		0:                   []byte(),
		NUM:                 []byte(num),
		LOWER:               []byte(lower),
		NUM | LOWER:         []byte(num + lower),
		UPPER:               []byte(upper),
		NUM | UPPER:         []byte(num + upper),
		LOWER | UPPER:       []byte(lower + upper),
		NUM | LOWER + UPPER: []byte(num + upper + lower),
	}
	lens = [ALL + 1]int{
		0:                   0,
		NUM:                 len(num),
		LOWER:               len(lower),
		NUM | LOWER:         len(num + lower),
		UPPER:               len(upper),
		NUM | UPPER:         len(num + upper),
		LOWER | UPPER:       len(lower + upper),
		NUM | LOWER + UPPER: len(num + upper + lower),
	}
)

func RandStringAWithChars(l int, chars []byte, data []byte) []byte {
	buf, n := goutils.EnlargeByte(data, l)
	ret := buf[n:]
	for i := 0; i < l; i++ {
		ret[i] = chars[r.Intn(len(chars))]
	}
	return buf
}

func RandStringWithChars(l int, chars []byte) []byte {
	return RandStringAWithChars(l, chars, nil)
}

func RandStringAWithType(l int, typ int, data []byte) []byte {
	return RandStringAWithChars(l, chars[typ], data)
}
func RandStringWithType(l int, typ int) []byte {
	return RandStringAWithType(l, typ, nil)
}

func RandString(l int) []byte {
	return RandStringWithType(l, ALL)
}
func RandStringA(l int, data []byte) []byte {
	return RandStringAWithType(l, ALL, data)
}
