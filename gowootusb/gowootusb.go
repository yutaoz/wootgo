package gowootusb

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -L../lib -lwooting-rgb-sdk64
#include "wooting-usb.h"
#include "wooting-rgb-sdk.h"


*/
import "C"

const (
	WOOTING_RGB_ROWS = 6
	WOOTING_RGB_COLS = 21
)

type VoidCB func()

type WootingDeviceLayout int32
type WootingDeviceType int32

const (
	LAYOUT_UNKNOWN WootingDeviceLayout = -1
	LAYOUT_ANSI    WootingDeviceLayout = 0
	LAYOUT_ISO     WootingDeviceLayout = 1

	DEVICE_KEYBOARD_TKL WootingDeviceType = 1
	DEVICE_KEYBOARD     WootingDeviceType = 2
	DEVICE_KEYPAD_60    WootingDeviceType = 3
	DEVICE_KEYPAD_3KEY  WootingDeviceType = 4
)

type WootingUSBMeta struct {
	Connected        bool
	Model            string
	MaxRows          uint8
	MaxColumns       uint8
	LedIndexMax      uint8
	DeviceType       WootingDeviceType
	V2Interface      bool
	Layout           WootingDeviceLayout
	UsesSmallPackets bool
}

// for some reason returns 0 for me even in c,
// maybe it works on your machine ;-;
func DeviceCount() uint8 {
	return uint8(C.wooting_usb_device_count())
}

// selects which device to send commands to
// will not work if devicecount is 0
func SelectDevice(dev uint8) bool {
	return bool(C.wooting_usb_select_device(C.uint8_t(dev)))
}

// hacky implementation of deviceinfo to work around cgo shenanigans and keep it in wooting rgb
// I CANT GET DEVICE INFO TO WORK EVEN IN C

// func DevInfoRaw() {
// 	cMeta := C.wooting_rgb_device_info()
// 	//devType := WootingDeviceType(cMeta.device_type)

// 	// t := ""
// 	// if devType == 1 {
// 	// 	t = "TKL"
// 	// } else if devType == 2 {
// 	// 	t = "Normal"
// 	// } else if devType == 3 {
// 	// 	t = "60%"
// 	// } else if devType == 4 {
// 	// 	t = "UWU"
// 	// }

// 	maxr := uint8((*cMeta).max_rows)
// 	fmt.Println("Type: %d", maxr)
// 	//fmt.Println("devtype: " + strconv.Itoa(deviceType))
// }

// func DeviceInfo() *WootingUSBMeta {
// 	cMeta := C.wooting_rgb_device_info_wrapper()
// 	if cMeta == nil {
// 		return nil
// 	}

// 	return &WootingUSBMeta{
// 		Connected:        bool(cMeta.connected),
// 		Model:            cMeta.model,
// 		MaxRows:          uint8(cMeta.max_rows),
// 		MaxColumns:       uint8(cMeta.max_columns),
// 		LedIndexMax:      uint8(cMeta.led_index_max),
// 		DeviceType:       cMeta.device_type,
// 		V2Interface:      bool(cMeta.v2_interface),
// 		Layout:           cMeta.layout,
// 		UsesSmallPackets: bool(cMeta.uses_small_packets),
// 	}
// }
