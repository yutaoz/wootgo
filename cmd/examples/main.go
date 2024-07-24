package main

import (
	"C"
	"time"

	gowootrgb "github.com/yutaoz/wootgo/gowootrgb"
	gowootusb "github.com/yutaoz/wootgo/gowootusb"
)
import "fmt"

func main() {
	testCount()
	autoUpdateTest()
}

// if count = 0, metadata related functions WILL CRASH
// this includes direct set, use array set instead
func testCount() {
	cnt := gowootusb.DeviceCount()

	fmt.Println("Count: ", cnt)
}

// crashes with my keyboard due to some meta issues,
// hopefully it works for you ;-;
func directTest() {
	x := uint8(1)
	y := uint8(2)
	r := uint8(50)
	g := uint8(50)
	b := uint8(100)
	_ = gowootrgb.DirectSetKey(x, y, r, g, b)
	time.Sleep(1 * time.Second)

	gowootrgb.Reset()
	gowootrgb.Close()
}

// sets 2 keys red and green, swaps, and resets.
// SHOULD NOT DISPLAY BLUE. IF SO, AUTO UPDATE NO WORK
func autoUpdateTest() {
	gowootrgb.ArrayAutoUpdate(true)
	gowootrgb.ArraySetSingle(2, 3, 200, 0, 0)
	gowootrgb.ArraySetSingle(3, 4, 0, 200, 0)
	time.Sleep(1 * time.Second)
	gowootrgb.ArraySetSingle(2, 3, 0, 200, 0)
	gowootrgb.ArraySetSingle(3, 4, 200, 0, 0)
	time.Sleep(1 * time.Second)
	gowootrgb.ArrayAutoUpdate(false)
	gowootrgb.ArraySetSingle(2, 3, 0, 0, 200)
	gowootrgb.ArraySetSingle(3, 4, 0, 0, 200)
	time.Sleep(1 * time.Second)
	gowootrgb.Reset()
	gowootrgb.Close()
}

// test for array updates
// sets 2 keys diagonnaly red and green, swaps them, then resets
func testArray() {
	gowootrgb.ArraySetSingle(2, 3, 200, 0, 0)
	gowootrgb.ArraySetSingle(3, 4, 0, 200, 0)
	gowootrgb.ArrayUpdateKeyboard()

	time.Sleep(1 * time.Second)

	gowootrgb.ArraySetSingle(2, 3, 0, 200, 0)
	gowootrgb.ArraySetSingle(3, 4, 200, 0, 0)
	gowootrgb.ArrayUpdateKeyboard()

	time.Sleep(1 * time.Second)

	gowootrgb.Reset()
	gowootrgb.Close()
}

// test for setting full array
// flashes entire keyboard red, then green, then resets
func testFullArray() {
	buffsize := gowootusb.WOOTING_RGB_ROWS * gowootusb.WOOTING_RGB_COLS
	colbuff := make([]uint8, buffsize)

	for i := 0; i < buffsize; i++ {
		if i%3 == 0 {
			colbuff[i] = 255
		} else {
			colbuff[i] = 0
		}
	}
	gowootrgb.ArraySetFull(colbuff)
	gowootrgb.ArrayUpdateKeyboard()
	time.Sleep(1 * time.Second)
	for i := 0; i < buffsize; i++ {
		if i%3 == 1 {
			colbuff[i] = 255
		} else {
			colbuff[i] = 0
		}
	}
	gowootrgb.ArraySetFull(colbuff)
	gowootrgb.ArrayUpdateKeyboard()
	time.Sleep(1 * time.Second)

	gowootrgb.Reset()
	gowootrgb.Close()
}
