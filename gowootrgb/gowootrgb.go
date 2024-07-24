package gowootrgb

// DISCLAIMER ----------------------------------------------------------------
// FUNCTIONS LIKE CONNECTED WORK, BUT DEVICE COUNT AND META DO NOT WORK FOR ME
// EVEN IN C
// RGB AND ARRAY RELATED FUNCTIONS STILL WORK, HOPE THE OTHER STUFF WORKS FOR
// YOU

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -L../lib -lwooting-rgb-sdk64
#include "wooting-rgb-sdk.h"
#include "wooting-usb.h"

extern void goCallback();
*/
import "C"
import (
	"fmt"
	"unsafe"

	gowootusb "github.com/yutaoz/wootgo/gowootusb"
)

type RGBMatrix [gowootusb.WOOTING_RGB_ROWS][gowootusb.WOOTING_RGB_COLS]uint16

var rgbBufferMatrix RGBMatrix

var rgbAutoUpdate = false

// checks if keyboard is connected
// true does NOT indicate that meta-related functions will work
// see device count in wootusb
func KbdConnected() bool {
	return bool(C.wooting_rgb_kbd_connected())
}

var goCallbackFunc gowootusb.VoidCB

func goCallback() {
	if goCallbackFunc != nil {
		goCallbackFunc()
	}
}

// sets a callback function for when keyboard is disconnected
// if devicecount returns 0, will do nothing on disconnect.
func SetDisconnectedCb(cb gowootusb.VoidCB) {
	fmt.Println("Loaded")
	goCallbackFunc = cb
	C.wooting_rgb_set_disconnected_cb((C.void_cb)(unsafe.Pointer(&goCallbackFunc)))
}

// resets rgbs
func ResetRgb() bool {
	return bool(C.wooting_rgb_reset_rgb())
}

// closes keyboard handle
func Close() bool {
	return bool(C.wooting_rgb_close())
}

// resets all back to original state
func Reset() bool {
	return bool(C.wooting_rgb_reset())
}

// sets a key to a color
// IF THIS CRASHES AND DUMPS ASM, USE ARRAY SET SINGLE AND ARRAY UPDATE
// DIRECT SET USES DEVICE META WHICH IS BUGGYI
func DirectSetKey(row, column, red, green, blue uint8) bool {
	return bool(C.wooting_rgb_direct_set_key(
		C.uint8_t(row),
		C.uint8_t(column),
		C.uint8_t(red),
		C.uint8_t(green),
		C.uint8_t(blue),
	))
}

// resets a key
// IF THIS CRASHES AND DUMPS ASM, USE ARRAY SET SINGLE AND ARRAY UPDATE
// DIRECT SET USES DEVICE META WHICH IS BUGGYI
func DirectResetKey(row, column uint8) bool {
	return bool(C.wooting_rgb_direct_reset_key(
		C.uint8_t(row),
		C.uint8_t(column),
	))
}

// updates keyboard based on current rgb matrix
func ArrayUpdateKeyboard() bool {
	return bool(C.wooting_usb_send_buffer_v2(
		(*[gowootusb.WOOTING_RGB_COLS]C.uint16_t)(unsafe.Pointer(&rgbBufferMatrix)),
	))
	//return bool(C.wooting_rgb_array_update_keyboard())
}

// sets auto-update
func ArrayAutoUpdate(autoUpdate bool) {
	rgbAutoUpdate = autoUpdate
}

func encodeColor(red, green, blue uint8) uint16 {
	var encode uint16

	encode |= (uint16(red) & 0xf8) << 8
	encode |= (uint16(green) & 0xfc) << 3
	encode |= (uint16(blue) & 0xf8) >> 3

	return encode
}

func decodeColor(color uint16) (red, green, blue uint8) {
	red = uint8((color >> 8) & 0xf8)
	green = uint8((color >> 3) & 0xfc)
	blue = uint8((color << 3) & 0xf8)
	return
}

// sets a single value in the rgb matrix
func ArraySetSingle(row, column, red, green, blue uint8) bool {
	prevVal := (rgbBufferMatrix)[row][column]
	newVal := encodeColor(red, green, blue)
	if prevVal != newVal {
		(rgbBufferMatrix)[row][column] = newVal
	}

	if rgbAutoUpdate {
		return ArrayUpdateKeyboard()
	}

	return true
}

// sets full rgb matrix
// since device metadata does not work, defaults to 6x21 rgbs
func ArraySetFull(colorBuffer []uint8) bool {
	for row := 0; row < gowootusb.WOOTING_RGB_ROWS; row++ {
		idx := row * 3
		colorSlice := colorBuffer[idx : idx+3]
		for col := 0; col < gowootusb.WOOTING_RGB_COLS; col++ {
			red := colorSlice[0]
			green := colorSlice[1]
			blue := colorSlice[2]
			color := encodeColor(red, green, blue)
			rgbBufferMatrix[row][col] = color
		}
	}
	if rgbAutoUpdate {
		return ArrayUpdateKeyboard()
	}
	return true
}
