package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var data1 uint16 = 65
	WriteData(data1)
	// writeLittleEndian
	// [0 0 0 0]
	// [0 0 0 0 65 0]
	// A

	// writeBigEndian
	// [0 0 0 0]
	// [0 0 0 0 0 65]
	// A

	// writeLittleEndianBufferOnly
	// [65 0]
	// A

	// writeBigEndianBufferOnly
	// [0 65]
	// A

	var data2 = struct {
		Param1 uint8
		Param2 uint16
		Param3 [4]byte
		Param4 uint16
	}{
		Param1: 0x0020, // 32
		Param2: 0x00ff,
		Param3: [4]byte{},
		Param4: 0x0001,
	}
	WriteData(data2)
	// writeLittleEndian
	// [0 0 0 0]
	// [0 0 0 0 32 255 0 0 0 0 0 1 0]
	//

	// writeBigEndian
	// [0 0 0 0]
	// [0 0 0 0 32 0 255 0 0 0 0 0 1]
	//

	// writeLittleEndianBufferOnly
	// [32 255 0 0 0 0 0 1 0]
	//

	// writeBigEndianBufferOnly
	// [32 0 255 0 0 0 0 0 1]
	//

	var data3 = struct {
		Param1 uint8
		Param2 uint8
		Param3 uint8
	}{
		Param1: 0x0020, // 32
		Param2: 0x00ff,
		Param3: 0x0001,
	}
	WriteData(data3)
	// writeLittleEndian
	// [0 0 0 0]
	// [0 0 0 0 32 255 1]
	//

	// writeBigEndian
	// [0 0 0 0]
	// [0 0 0 0 32 255 1]
	//

	// writeLittleEndianBufferOnly
	// [32 255 1]
	//

	// writeBigEndianBufferOnly
	// [32 255 1]
	//

	// NOTE: 一番わかりやすい
	var data4 = struct {
		Param1 uint32
		Param2 uint32
		Param3 uint32
	}{
		Param1: 0x0020, // 32
		Param2: 0x00ff,
		Param3: 0x0001,
	}
	WriteData(data4)
	// writeLittleEndian
	// [0 0 0 0]
	// [0 0 0 0 32 0 0 0 255 0 0 0 1 0 0 0]
	//

	// writeBigEndian
	// [0 0 0 0]
	// [0 0 0 0 0 0 0 32 0 0 0 255 0 0 0 1]
	//

	// writeLittleEndianBufferOnly
	// [32 0 0 0 255 0 0 0 1 0 0 0]
	//

	// writeBigEndianBufferOnly
	// [0 0 0 32 0 0 0 255 0 0 0 1]
	//
}

func WriteData(data interface{}) {
	if err := writeLittleEndian(data); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := writeBigEndian(data); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := writeLittleEndianBufferOnly(data); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := writeBigEndianBufferOnly(data); err != nil {
		panic(err)
	}
}

func writeLittleEndian(data interface{}) error {
	fmt.Println("writeLittleEndian")
	b := make([]byte, 4)
	buf := bytes.NewBuffer(b)

	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		return err
	}

	fmt.Println(b)
	fmt.Println(buf.Bytes())
	fmt.Println(buf.String())
	return nil
}

func writeLittleEndianBufferOnly(data interface{}) error {
	fmt.Println("writeLittleEndianBufferOnly")
	buf := &bytes.Buffer{}

	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		return err
	}

	fmt.Println(buf.Bytes())
	fmt.Println(buf.String())
	return nil
}

func writeBigEndian(data interface{}) error {
	fmt.Println("writeBigEndian")
	b := make([]byte, 4)
	buf := bytes.NewBuffer(b)

	if err := binary.Write(buf, binary.BigEndian, data); err != nil {
		return err
	}

	fmt.Println(b)
	fmt.Println(buf.Bytes())
	fmt.Println(buf.String())
	return nil
}

func writeBigEndianBufferOnly(data interface{}) error {
	fmt.Println("writeBigEndianBufferOnly")
	buf := &bytes.Buffer{}

	if err := binary.Write(buf, binary.BigEndian, data); err != nil {
		return err
	}

	fmt.Println(buf.Bytes())
	fmt.Println(buf.String())
	return nil
}
