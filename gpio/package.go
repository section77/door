package gpio

import "C"

// Pin represents a RPi v2 GPIO Pin (hw numbering schema)
type Pin uint8

type LowOrHigh int

const (
	Pin03 Pin = 2
	Pin05 Pin = 3
	Pin07 Pin = 4
	Pin08 Pin = 14
	Pin10 Pin = 15
	Pin11 Pin = 17
	Pin12 Pin = 18
	Pin13 Pin = 27
	Pin15 Pin = 22
	Pin16 Pin = 23
	Pin18 Pin = 24
	Pin19 Pin = 10
	Pin21 Pin = 9
	Pin22 Pin = 25
	Pin23 Pin = 11
	Pin24 Pin = 8
	Pin26 Pin = 7
	Pin29 Pin = 5
	Pin31 Pin = 6
	Pin32 Pin = 12
	Pin33 Pin = 13
	Pin35 Pin = 19
	Pin36 Pin = 16
	Pin37 Pin = 26
	Pin38 Pin = 20
	Pin40 Pin = 21

	Low  LowOrHigh = 0
	High LowOrHigh = 1
)

func (x LowOrHigh) String() string {
	if x == 0 {
		return "LOW"
	}
	return "HIGH"
}
