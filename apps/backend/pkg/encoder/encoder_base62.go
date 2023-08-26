package encoder

import (
	"hash/fnv"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type Base62Encoder struct {
}

func NewBase62Encoder() EncoderService {
	return &Base62Encoder{}
}

func (enc *Base62Encoder) Encode(n string) (string, error) {
	h := fnv.New64a()
	length := len(alphabet)
	var encodedBuilder strings.Builder

	h.Write([]byte(n))
	encodedBuilder.Grow(10)

	number := h.Sum64()

	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String(), nil
}
