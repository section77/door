// +build arm

package main

// #cgo LDFLAGS: -lbcm2835
// #include <bcm2835.h>
import "C"
import "errors"

func pwm(value int) error {
	//
	// init bcm2835 / configure for pwm
	//
	if C.bcm2835_init() == 0 {
		return errors.New("unable to initialize bcm2835")
	}
	defer C.bcm2835_close()

	C.bcm2835_gpio_fsel(C.RPI_V2_GPIO_P1_12, C.BCM2835_GPIO_FSEL_ALT5)
	C.bcm2835_pwm_set_clock(C.BCM2835_PWM_CLOCK_DIVIDER_512)
	// C.bcm2835_pwm_set_mode(channel, markspace, enabled)
	C.bcm2835_pwm_set_mode(0, 1, 1)
	// C.bcm2835_pwm_set_range(channel, range)
	C.bcm2835_pwm_set_range(0, 1024)

	//
	// send value
	//

	// bcm2835_pwm_set_data(channel, value)
	C.bcm2835_pwm_set_data(0, C.uint32_t(value))

	return nil
}
