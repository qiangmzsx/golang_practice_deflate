package mydeflate

import (
	"testing"
)

func TestTakeBit(t *testing.T) {
	for i := 0; i < 256; i++ {
	//for i := 0; i < 3; i++ {
		src := []byte{byte(i)}
		res, err := takeBit(src, 0, 8)
		if err != nil {
			t.Error(err)
			t.Fail()
		} else if res != uint(i) {
			t.Errorf("value missing expected %d actual %d, value: %d", i, res, int(src[0]))
			t.Fail()
		} else {
			// OK success!
		}
	}

	// dst must be "aaaaa"
}

func TestSmallRegular(t *testing.T) {
	x := []byte{0xb4, 0xc4, 0x04, 0x02, 0x00}
	dst := []byte{}
	Decode(dst, x, -8)

	// dst must be "aaaaa"
}
