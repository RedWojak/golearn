package main

import "fmt"

type car struct {
	gasPedal      uint16
	breakPedal    uint16
	steeromgWheel uint16
	topSpeedKmh   float64
}

func main() {
	aCar := car{gasPedal: 22341,
		breakPedal:    0,
		steeromgWheel: 1256,
		topSpeedKmh:   225.0}

	fmt.Println(aCar.gasPedal)

}
