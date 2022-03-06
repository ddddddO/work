package main

import (
	"fmt"
	"machine"
	"time"
)

const (
	led = machine.LCD_BACKLIGHT
)

// machine.でグローバルにアクセスできてしまうけど一応構造体で持っておく

type button struct {
	a machine.Pin
	b machine.Pin
	c machine.Pin
}

func newButton() button {
	buttonA := machine.WIO_KEY_A
	buttonA.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	buttonB := machine.WIO_KEY_B
	buttonB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	buttonC := machine.WIO_KEY_C
	buttonC.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	return button{
		a: buttonA,
		b: buttonB,
		c: buttonC,
	}
}

func (b button) isPressedA() bool {
	return !b.a.Get()
}

func (b button) isPressedB() bool {
	return !b.b.Get()
}

func (b button) isPressedC() bool {
	return !b.c.Get()
}

type juziKey struct {
	up    machine.Pin
	down  machine.Pin
	left  machine.Pin
	right machine.Pin
	press machine.Pin
}

func newJuziKey() juziKey {
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

	return juziKey{
		up:    up,
		down:  down,
		left:  left,
		right: right,
		press: press,
	}
}

func (j juziKey) isPressedUp() bool {
	return !j.up.Get()
}

func (j juziKey) isPressedDown() bool {
	return !j.down.Get()
}

func (j juziKey) isPressedLeft() bool {
	return !j.left.Get()
}

func (j juziKey) isPressedRight() bool {
	return !j.right.Get()
}

func (j juziKey) isPressedPress() bool {
	return !j.press.Get()
}

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := newButton()
	juziKey := newJuziKey()

	for {
		if button.isPressedA() {
			led.Low()
			fmt.Printf("machine.WIO_KEY_A pressed\r\n")
		} else if button.isPressedB() {
			led.Low()
			fmt.Printf("machine.WIO_KEY_B pressed\r\n")
		} else if button.isPressedC() {
			led.Low()
			fmt.Printf("machine.WIO_KEY_C pressed\r\n")
		} else if juziKey.isPressedUp() {
			led.Low()
			fmt.Printf("machine.WIO_5S_UP pressed\r\n")
		} else if juziKey.isPressedLeft() {
			led.Low()
			fmt.Printf("machine.WIO_5S_LEFT pressed\r\n")
		} else if juziKey.isPressedRight() {
			led.Low()
			fmt.Printf("machine.WIO_5S_RIGHT pressed\r\n")
		} else if juziKey.isPressedDown() {
			led.Low()
			fmt.Printf("machine.WIO_5S_DOWN pressed\r\n")
		} else if juziKey.isPressedPress() {
			led.Low()
			fmt.Printf("machine.WIO_5S_PRESS pressed\r\n")
		} else {
			led.High()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
