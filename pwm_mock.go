// +build !arm
//
// pwm mock
//  - zu entwicklung unter nicht arm architekturen (keine bcm2835 lib vorhanden)
//  - gibt die pwm werte auf stdout aus

package main

import (
	"fmt"
)

func pwm(value int) error {
	fmt.Printf("dummy pwm - value: %d\n", value)
	return nil
}
