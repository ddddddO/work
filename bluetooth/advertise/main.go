package main

import (
	"fmt"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	sender := newSender()

	// NOTE: 以降、アドバタイズの際に、書き込みをしていたバイト列をべた書きで生成（アドバタイズに関わりそうなところのみ）
	// https://www.dododo.site/posts/bluetooth%E3%82%92%E8%A9%A6%E3%81%97%E3%81%A6%E3%81%BF%E3%82%8B/
	LEReadAdvertisingChannelTxPower, len := newLEReadAdvertisingChannelTxPower()
	fmt.Println(LEReadAdvertisingChannelTxPower)
	if err := sender.send(LEReadAdvertisingChannelTxPower, len); err != nil {
		panic(err)
	}

	LESetAdvertisingParameters, len := newLESetAdvertisingParameters()
	fmt.Println(LESetAdvertisingParameters)
	if err := sender.send(LESetAdvertisingParameters, len); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(LESetAdvertisingParameters)
	if err := sender.send(LESetAdvertisingParameters, len); err != nil {
		panic(err)
	}

	LESetAdvertisingData, len := newLESetAdvertisingData()
	fmt.Println(LESetAdvertisingData)
	if err := sender.send(LESetAdvertisingData, len); err != nil {
		panic(err)
	}

	LESetAdvertiseEnable, len := newLESetAdvertiseEnable()
	fmt.Println(LESetAdvertiseEnable)
	if err := sender.send(LESetAdvertiseEnable, len); err != nil {
		panic(err)
	}

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println(LESetAdvertiseEnable)
		if err := sender.send(LESetAdvertiseEnable, len); err != nil {
			panic(err)
		}
	}
}

type sender struct {
	fd int
}

func newSender() *sender {
	fd, err := newFDBluetooth()
	if err != nil {
		panic(err)
	}
	return &sender{
		fd: fd,
	}
}

func (s *sender) send(b []byte, length int) error {
	if _, err := unix.Write(s.fd, b[:4+length]); err != nil {
		return err
	}
	return nil
}

// NOTE: 以下リンク先からコピペ
// https://github.com/currantlabs/ble/blob/master/linux/hci/socket/socket.go

const (
	ioctlSize     = 4
	hciMaxDevices = 16
	typHCI        = 72 // 'H'
)

// NOTE: 以下リンク先コードのGoバージョンのよう
// https://code.woboq.org/qt5/include/bluetooth/hci.h.html#88
var (
	hciUpDevice      = ioW(typHCI, 201, ioctlSize) // HCIDEVUP
	hciDownDevice    = ioW(typHCI, 202, ioctlSize) // HCIDEVDOWN
	hciResetDevice   = ioW(typHCI, 203, ioctlSize) // HCIDEVRESET
	hciGetDeviceList = ioR(typHCI, 210, ioctlSize) // HCIGETDEVLIST
)

func ioR(t, nr, size uintptr) uintptr {
	return (2 << 30) | (t << 8) | nr | (size << 16)
}

func ioW(t, nr, size uintptr) uintptr {
	return (1 << 30) | (t << 8) | nr | (size << 16)
}

// NOTE: ioctlとは
// https://www.wdic.org/w/TECH/ioctl
func ioctl(fd, op, arg uintptr) error {
	if _, _, ep := unix.Syscall(unix.SYS_IOCTL, fd, op, arg); ep != 0 {
		return ep
	}
	return nil
}

type devListRequest struct {
	devNum     uint16
	devRequest [hciMaxDevices]struct {
		id  uint16
		opt uint32
	}
}

func newFDBluetooth() (int, error) {
	fd, err := unix.Socket(unix.AF_BLUETOOTH, unix.SOCK_RAW, unix.BTPROTO_HCI)
	if err != nil {
		return -1, err
	}

	req := devListRequest{devNum: 16}
	if err = ioctl(uintptr(fd), hciGetDeviceList, uintptr(unsafe.Pointer(&req))); err != nil {
		panic(err)
	}

	var e error
	for id := 0; id < int(req.devNum); id++ {
		// Reset the device in case previous session didn't cleanup properly.
		if err := ioctl(uintptr(fd), hciDownDevice, uintptr(id)); err != nil {
			e = err
			continue
		}
		if err := ioctl(uintptr(fd), hciUpDevice, uintptr(id)); err != nil {
			e = err
			continue
		}

		// HCI User Channel requires exclusive access to the device.
		// The device has to be down at the time of binding.
		if err := ioctl(uintptr(fd), hciDownDevice, uintptr(id)); err != nil {
			e = err
			continue
		}

		// Bind the RAW socket to HCI User Channel
		sa := unix.SockaddrHCI{Dev: uint16(id), Channel: unix.HCI_CHANNEL_USER}
		if err := unix.Bind(fd, &sa); err != nil {
			e = err
			continue
		}
		e = nil
		break
	}

	fmt.Printf("fd: %d\n", fd)
	return fd, e
}

func newLEReadAdvertisingChannelTxPower() ([]byte, int) {
	LEReadAdvertisingChannelTxPower := make([]byte, 64)
	LEReadAdvertisingChannelTxPower[0] = 1
	LEReadAdvertisingChannelTxPower[1] = 7
	LEReadAdvertisingChannelTxPower[2] = 32
	return LEReadAdvertisingChannelTxPower, 0
}

func newLESetAdvertisingParameters() ([]byte, int) {
	LESetAdvertisingParameters := make([]byte, 64)
	LESetAdvertisingParameters[0] = 1
	LESetAdvertisingParameters[1] = 6
	LESetAdvertisingParameters[2] = 32
	LESetAdvertisingParameters[3] = 15
	LESetAdvertisingParameters[4] = 32
	LESetAdvertisingParameters[6] = 32
	LESetAdvertisingParameters[17] = 7
	return LESetAdvertisingParameters, 15
}

func newLESetAdvertisingData() ([]byte, int) {
	LESetAdvertisingData := make([]byte, 64)
	LESetAdvertisingData[0] = 1
	LESetAdvertisingData[1] = 8
	LESetAdvertisingData[2] = 32
	LESetAdvertisingData[3] = 32
	LESetAdvertisingData[4] = 29
	LESetAdvertisingData[5] = 2
	LESetAdvertisingData[6] = 1
	LESetAdvertisingData[7] = 6
	LESetAdvertisingData[8] = 17
	LESetAdvertisingData[9] = 7

	LESetAdvertisingData[10] = 251
	LESetAdvertisingData[11] = 52
	LESetAdvertisingData[12] = 155
	LESetAdvertisingData[13] = 95
	LESetAdvertisingData[14] = 128
	LESetAdvertisingData[17] = 128
	LESetAdvertisingData[19] = 16

	LESetAdvertisingData[20] = 1
	LESetAdvertisingData[24] = 1
	LESetAdvertisingData[26] = 7
	LESetAdvertisingData[27] = 9
	LESetAdvertisingData[28] = 71
	LESetAdvertisingData[29] = 111

	LESetAdvertisingData[30] = 112
	LESetAdvertisingData[31] = 104
	LESetAdvertisingData[32] = 101
	LESetAdvertisingData[33] = 114
	return LESetAdvertisingData, 32
}

func newLESetAdvertiseEnable() ([]byte, int) {
	LESetAdvertiseEnable := make([]byte, 64)
	LESetAdvertiseEnable[0] = 1
	LESetAdvertiseEnable[1] = 10
	LESetAdvertiseEnable[2] = 32
	LESetAdvertiseEnable[3] = 1
	LESetAdvertiseEnable[4] = 1
	return LESetAdvertiseEnable, 1
}
