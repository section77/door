// +build arm
//
//

package main

// #cgo LDFLAGS: -lbcm2835
// #include <bcm2835.h>
import "C"
import "time"
import "sync"
import "os"
import "log"
import "fmt"

var lock = sync.Mutex{}

func init() {
	go reactOnBtnEvent()
}

func pwm(value int) error {
	lock.Lock()
	defer lock.Unlock()

	//
	// open the library
	//
	if C.bcm2835_init() == 0 {
		// bcm2835_init prints the failure reason on stderr
		fmt.Println("bcm2835 init error")
		os.Exit(1)
	}
	defer C.bcm2835_close()

	//
	// configure port
	//
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

	//
	// give the servo some time to reach the target
	// position and then close the library
	//
	time.Sleep(2 * time.Second)
	C.bcm2835_gpio_fsel(C.RPI_V2_GPIO_P1_12, C.BCM2835_GPIO_FSEL_OUTP)
	C.bcm2835_close()

	return nil
}

func reactOnBtnEvent() {
	lock.Lock()
	defer lock.Unlock()

	//
	// open the library
	//
	if C.bcm2835_init() == 0 {
		// bcm2835_init prints the failure reason on stderr
		fmt.Println("bcm2835 init error")
		os.Exit(1)
	}
	defer C.bcm2835_close()

	//
	// configure port
	//

	// input
	C.bcm2835_gpio_fsel(C.RPI_V2_GPIO_P1_11, C.BCM2835_GPIO_FSEL_INPT)

	// enable low detection
	C.bcm2835_gpio_len(C.RPI_V2_GPIO_P1_11)

	var btnState = func() (bool, time.Duration) {
		isPushed := false
		start := time.Now()
		for C.bcm2835_gpio_eds(C.RPI_V2_GPIO_P1_11) == 1 {
			C.bcm2835_gpio_set_eds(C.RPI_V2_GPIO_P1_11)
			isPushed = true
			time.Sleep(100 * time.Millisecond)
		}
		C.bcm2835_gpio_set_eds(C.RPI_V2_GPIO_P1_11)
		return isPushed, time.Since(start)
	}

	if isPushed, duration := btnState(); isPushed {
		log.Printf("button was pushed for: '%s'", duration)
		switch {
		case duration.Seconds() < 0.5:
			// ignore - to short
			log.Println("ignore")
		default:
			// toggle
			log.Println("delay toggel ....")
			go func() {
				time.Sleep(10 * time.Second)
				log.Println("toogle!")
				lockToggle()
			}()
		}
	}

	go func() {
		time.Sleep(500 * time.Millisecond)
		reactOnBtnEvent()
	}()
}
