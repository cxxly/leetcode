// bitmap_test.go
package main

import (
	"testing"
)

func TestBitmapSort(t *testing.T) {
	bmap := NewBitmap(14)
	original := [5]uint64{4, 6, 3, 1, 14}
	expected := [5]uint64{1, 3, 4, 6, 14}
	actual := [5]uint64{}

	for _, offset := range original {
		bmap.SetBit(offset, 1)
	}

	var i, offset uint64 = 0, 0
	for ; offset < bmap.Maxpos()+1; offset++ {
		if bmap.GetBit(offset) == 1 {
			actual[i] = offset
			i++
		}
	}

	if expected != actual {
		t.Errorf("Expected:%v, actual:%v", expected, actual)
	}
}
