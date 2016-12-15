package codec

import (
	"testing"
	"math/rand"
)

const seed = 0x12345

func TestCodec(t *testing.T) {
	rnd := rand.New(rand.NewSource(seed))
	encoder, _ := encode("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	decoder, _ := decode("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := 1; i <= 10000; i++ {
		n := uint64(rnd.Uint32())<<32 + uint64(rnd.Uint32())
		e := encoder(n)
		d, _ := decoder(e)
		if d != n {
			t.Errorf("decode(%q) == %q, want %q", e, d, n)
		}
	}
}
