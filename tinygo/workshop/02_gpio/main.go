package main

import (
	"fmt"
	"machine"
	"time"
)

const (
	led = machine.LCD_BACKLIGHT
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttonA, buttonB, buttonC := prepareButton()
	up, down, left, right, press := prepareJuziKey()

	for {
		if !buttonA.Get() {
			led.Low()
			fmt.Printf("machine.WIO_KEY_A pressed\r\n")
		} else if !buttonB.Get() {
			led.Low()
			fmt.Printf("machine.WIO_KEY_B pressed\r\n")
		} else if !buttonC.Get() {
			led.Low()
			fmt.Printf("machine.WIO_KEY_C pressed\r\n")
		} else if !up.Get() {
			led.Low()
			fmt.Printf("machine.WIO_5S_UP pressed\r\n")
		} else if !left.Get() {
			led.Low()
			fmt.Printf("machine.WIO_5S_LEFT pressed\r\n")
		} else if !right.Get() {
			led.Low()
			fmt.Printf("machine.WIO_5S_RIGHT pressed\r\n")
		} else if !down.Get() {
			led.Low()
			fmt.Printf("machine.WIO_5S_DOWN pressed\r\n")
		} else if !press.Get() {
			led.Low()
			fmt.Printf("machine.WIO_5S_PRESS pressed\r\n")
		} else {
			led.High()
		}

		time.Sleep(time.Millisecond * 10)
	}
}

func prepareButton() (machine.Pin, machine.Pin, machine.Pin) {
	buttonA := machine.WIO_KEY_A
	buttonA.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	buttonB := machine.WIO_KEY_B
	buttonB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	buttonC := machine.WIO_KEY_C
	buttonC.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	return buttonA, buttonB, buttonC
}

func prepareJuziKey() (machine.Pin, machine.Pin, machine.Pin, machine.Pin, machine.Pin) {
	up := machine.WIO_5S_UP
	up.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	down := machine.WIO_5S_DOWN
	down.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	left := machine.WIO_5S_LEFT
	left.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	right := machine.WIO_5S_RIGHT
	right.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	press := machine.WIO_5S_PRESS
	press.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	return up, down, left, right, press
}
