// +build !arm
//
// pwm mock
//  - gpio mock to develop the library under x86
//  - the real gpio implemenation depend on the bcm2835 library
//  - only log the events

package main

import (
	"fmt"
)

func enableUMTSStick(enable bool) error {
	fmt.Printf("enableUMTSStick: %t\n", enable)
	return nil
}

func pwm(value int) error {
	fmt.Printf("dummy pwm - value: %d\n", value)
	return nil
}

func reactOnBtnEvent() {
}
