package main

import (
	"fmt"
)

//Bitmapsize represents default size of bitmap
const Bitmapsize = 0x01 << 32

type Bitmap struct {
	//data store bit date of bitmap
	data []byte
	// bitsize represents size of bitmap
	bitsize uint64
	//maxpos represents max position of 1 in bitmap
	maxpos uint64
}

func NewBitmap(size int) *Bitmap {
	if size == 0 || size > Bitmapsize {
		size = Bitmapsize
	}
	if remainder := size % 8; remainder != 0 {
		size += 8 - remainder
	}
	return &Bitmap{
		data:    make([]byte, size>>3),
		bitsize: uint64(size),
	}
}

func (b *Bitmap) SetBit(offset uint64, value uint8) bool {
	index, pos := offset/8, offset%8
	if offset > b.bitsize-1 {
		return false
	}
	if value == 0 {
		b.data[index] &= ^(0x01 << pos)
	} else {
		b.data[index] |= 0x01 << pos
		if b.maxpos < offset {
			b.maxpos = offset
		}
	}
	return true
}

func (b *Bitmap) GetBit(offset uint64) uint8 {
	index, pos := offset/8, offset%8
	if b.bitsize-1 < offset {
		return 0
	}
	return (b.data[index] >> pos) & 0x01
}

func (b *Bitmap) Maxpos() uint64 {
	return b.maxpos
}

// String output top 100 elements
func (b *Bitmap) String() string {
	var maxtotal, bittotal uint64 = 100, b.maxpos + 1
	if b.maxpos > maxtotal {
		bittotal = maxtotal
	}
	numSlice := make([]uint64, 0, bittotal)
	var offset uint64
	for offset = 0; offset < bittotal; offset++ {
		if b.GetBit(offset) == 1 {
			numSlice = append(numSlice, offset)
		}
	}
	return fmt.Sprintf("%v", numSlice)
}
