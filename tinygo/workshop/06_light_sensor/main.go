package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	machine.InitADC()

	led := machine.LCD_BACKLIGHT
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	sensor := machine.ADC{Pin: machine.WIO_LIGHT}
	sensor.Configure(machine.ADCConfig{})

	for {
		val := sensor.Get()
		fmt.Printf("%04X\r\n", val)
		// 面白い。wio terminalの背面に光を当てる
		if val < 0x5000 {
			led.Low()
		} else {
			led.High()
		}
		time.Sleep(time.Millisecond * 100)
	}
}
