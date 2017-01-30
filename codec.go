package codec

import (
	"bytes"
	"errors"
	"math"
	"unicode/utf8"
)

func reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func Encode(alphabet string) (func(uint64) string, error) {
	if alphabet == "" || utf8.RuneCountInString(alphabet) == 0 {
		return nil, errors.New("encode: alphabet cannot be empty")
	}

	m := make(map[uint64]rune)
	d := uint64(utf8.RuneCountInString(alphabet))

	for index, runeValue := range alphabet {
		m[uint64(index)] = runeValue
    }

	return func(n uint64) string {
		var buffer bytes.Buffer
	    r := uint64(0)

		for n > uint64(0) {
			r = n % d
			n = n / d
			buffer.WriteRune(m[r])
		}

		return reverse(buffer.String())
	}, nil
}

func Decode(alphabet string) (func(string) (uint64, error), error) {
	if alphabet == "" || utf8.RuneCountInString(alphabet) == 0 {
		return nil, errors.New("decode: alphabet cannot be empty")
	}

	m := make(map[rune]uint64)
	d := uint64(utf8.RuneCountInString(alphabet))

	for index, runeValue := range alphabet {
		m[runeValue] = uint64(index)
    }

	return func(s string) (uint64, error) {
		if s == "" {
			return 0, errors.New("decode: cannot decode empty string")
		}

		l := utf8.RuneCountInString(s)
	    r := uint64(0)

	    for index, runeValue := range s {
	    	p := (l - index - 1)
	    	v := uint64(math.Pow(float64(d), float64(p)))
	    	r += m[runeValue] * v
	    }

		return r, nil
	}, nil
}