// Copyright 2016, 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package codec

import (
	"testing"
	"math/rand"
)

const seed = 0x12345

func TestCodec(t *testing.T) {
	rnd := rand.New(rand.NewSource(seed))
	encoder, _ := Encode("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	decoder, _ := Decode("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := 1; i <= 10000; i++ {
		n := uint64(rnd.Uint32())<<32 + uint64(rnd.Uint32())
		e := encoder(n)
		d, _ := decoder(e)
		if d != n {
			t.Errorf("decode(%q) == %q, want %q", e, d, n)
		}
	}
}
