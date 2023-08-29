package encoder

import (
	"strconv"
	"testing"
)

func TestOfEncodeBase62Function(t *testing.T) {
	base62Encoder := NewBase62Encoder()

	data, _ := base62Encoder.Encode("aaaa")

	if data != "JDE3Lh4cZcv" {
		t.Errorf("Encode(\"aaaa\") = %s, want JDE3Lh4cZcv", data)
	}
}

func BenchmarkEncodeBase62(b *testing.B) {
	base62Encoder := NewBase62Encoder()

	for i := 0; i < b.N; i++ {
		base62Encoder.Encode("aaaa" + strconv.Itoa(i))
	}
}
