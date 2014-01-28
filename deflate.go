package mydeflate

import (
	"errors"
)

func takeBit(src []byte, current int, bit int) (uint, error) {
	if (current+bit)/8 > len(src) {
		return 0, errors.New("source is too short")
	}

	value := uint(0)
	base := uint(1)
	for max := current + bit; current < max; current++ {
		bit_pos := uint(current%8)
		if uint(src[current/8]) >> bit_pos & 1 == 1 {
			value += base
		}
		base = base * 2
	}
	return value, nil
}

func Compress(dst, src []byte, opt int) ([]byte, error) {
	src_size := len(src)
	dst = make([]byte, src_size)

	return dst, nil
}

func Decode(dst, src []byte, opt int) ([]byte, error) {
	// Take
	src_size := len(src)
	dst = make([]byte, src_size)

	switch opt {
	case -8:
		// raw mode (wihtout header, crc)

	default:
	}

	finished := false
	for finished {
		// take is last block (1 bit)
		// take block type (2 bits)
	}

	return dst, nil
}
