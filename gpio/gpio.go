// +build arm
//

package gpio

// #cgo LDFLAGS: -lbcm2835
// #include <bcm2835.h>
import "C"
import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// lock to synchronize the gpio access
// if it's no synchronized, the underlying bcm2835
// library triggers a panic
var mutex = sync.Mutex{}

func Read(pin Pin) (LowOrHigh, error) {
	// acquire the lock
	mutex.Lock()
	defer mutex.Unlock()

	// open the library
	if C.bcm2835_init() == 0 {
		return Low, errors.New("unable to initialize bcm2835")
	}
	defer C.bcm2835_close()

	// configure the port for input
	C.bcm2835_gpio_fsel(C.uint8_t(pin), C.BCM2835_GPIO_FSEL_INPT)
	// with a pullup
	C.bcm2835_gpio_fsel(C.uint8_t(pin), C.BCM2835_GPIO_PUD_UP)

	// read
	switch C.bcm2835_gpio_lev(C.uint8_t(pin)) {
	case C.LOW:
		return Low, nil
	case C.HIGH:
		return High, nil
	}
	return Low, errors.New("unexpeced response")
}

func Write(pin Pin, value LowOrHigh) error {
	// acquire the lock
	mutex.Lock()
	defer mutex.Unlock()

	// open the library
	if C.bcm2835_init() == 0 {
		return errors.New("unable to initialize bcm2835")
	}
	defer C.bcm2835_close()

	// configure the port for output
	C.bcm2835_gpio_fsel(C.uint8_t(pin), C.BCM2835_GPIO_FSEL_OUTP)

	switch value {
	case Low:
		C.bcm2835_gpio_write(C.uint8_t(pin), C.LOW)
	case High:
		C.bcm2835_gpio_write(C.uint8_t(pin), C.HIGH)
	default:
		return fmt.Errorf("Unexpected value for LowOrHigh: %s", value)
	}

	return nil
}

func Pwm(pin Pin, value int) error {
	// acquire the lock
	mutex.Lock()
	defer mutex.Unlock()

	// open the library
	if C.bcm2835_init() == 0 {
		return errors.New("unable to initialize bcm2835")
	}
	defer C.bcm2835_close()

	// configure the port for pwm
	C.bcm2835_gpio_fsel(C.uint8_t(pin), C.BCM2835_GPIO_FSEL_ALT5)
	C.bcm2835_pwm_set_clock(C.BCM2835_PWM_CLOCK_DIVIDER_512)
	C.bcm2835_pwm_set_mode(0, 1, 1)  // (channel, markspace, enabled)
	C.bcm2835_pwm_set_range(0, 1024) // (channel, range)

	// write the value
	C.bcm2835_pwm_set_data(0, C.uint32_t(value)) // (channel, value)

	//
	// give the servo some time to reach the target position
	//
	time.Sleep(2 * time.Second)
	C.bcm2835_gpio_fsel(C.uint8_t(pin), C.BCM2835_GPIO_FSEL_OUTP)

	return nil
}

func BlockWhileIsLow(pin Pin) (bool, time.Duration, error) {
	// acquire the lock
	mutex.Lock()
	defer mutex.Unlock()

	// open the library
	if C.bcm2835_init() == 0 {
		return false, 0, errors.New("unable to initialize bcm2835")
	}
	defer C.bcm2835_close()

	// configure the port for input
	C.bcm2835_gpio_fsel(C.uint8_t(pin), C.BCM2835_GPIO_FSEL_INPT)
	// with a pullup
	C.bcm2835_gpio_fsel(C.uint8_t(pin), C.BCM2835_GPIO_PUD_UP)

	var blockWhileIsLow = func() (bool, time.Duration) {
		isLow := false
		start := time.Now()
		for C.bcm2835_gpio_lev(C.uint8_t(pin)) == C.LOW {
			isLow = true
			time.Sleep(100 * time.Millisecond)
		}
		return isLow, time.Since(start)
	}

	isLow, duration := blockWhileIsLow()
	return isLow, duration, nil
}
