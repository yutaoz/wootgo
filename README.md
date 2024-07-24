# wootgo
A wrapper for the Wooting RGB SDK in Go

## Usage
`go get github.com/yutaoz/wootgo`

Make sure you have `wooting-rgb-sdk64.dll` in your project root

## Examples
```
package main

import (
	"fmt"
	"time"

	gwrgb "github.com/yutaoz/wootgo/gowootrgb"
)

// flashes a key red for 1 second, then resets all rgbs back to original
func main() {
	f := gwrgb.KbdConnected()
	if f {
		fmt.Println("connected!")
	} else {
		fmt.Println("no conn")
	}
	gwrgb.ArraySetSingle(3, 5, 250, 0, 0)
	gwrgb.ArrayUpdateKeyboard()
	time.Sleep(1 * time.Second)
	gwrgb.Reset()
	gwrgb.Close()
}
```

There are also example tests in cmd/examples

## IMPORTANT - BETA
I had some issues with the sdk on its own, such as not being able to get a device count or find my keyboard. As a result, there are no implementations for functions that require meta as I cannot reliably test these - they just crash (even in C)

If you can get a device count or meta struct working, I'd love a pr!
