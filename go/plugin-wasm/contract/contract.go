package contract

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type Protocol interface {
	GetStr() string
	Name(x string) string
	Port() uint64
}

var p Protocol

func Register(proto Protocol) {
	p = proto
}

//export get_str
func _getStr() uint64 {
	str := p.GetStr()
	ptr, size := stringToPtr(str)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

//export name
func _name(ptr, size uint32) uint64 {
	x := ptrToString(ptr, size)
	sentence := p.Name(x)
	return marshal(sentence)
}

//export port
func _port() uint64 {
	return p.Port()
}

// 以降、copied: https://github.com/aquasecurity/trivy/blob/e8085bae3e71fc5c9839feb13e34b75deba4ce9d/pkg/module/wasm/sdk.go
func marshal(v any) uint64 {
	b, err := json.Marshal(v)
	if err != nil {
		Error(fmt.Sprintf("marshal error: %s", err))
		return 0
	}

	p := uintptr(unsafe.Pointer(&b[0]))
	return (uint64(p) << uint64(32)) | uint64(len(b))
}

func unmarshal(ptr, size uint32, v any) error {
	s := ptrToString(ptr, size)
	if err := json.Unmarshal([]byte(s), v); err != nil {
		return fmt.Errorf("unmarshal error: %s", err)
	}
	return nil
}

// ptrToString returns a string from WebAssembly compatible numeric types representing its pointer and length.
func ptrToString(ptr uint32, size uint32) string {
	b := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), size)
	return *(*string)(unsafe.Pointer(&b))
}

// stringToPtr returns a pointer and size pair for the given string in a way compatible with WebAssembly numeric types.
func stringToPtr(s string) (uint32, uint32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return uint32(unsafePtr), uint32(len(buf))
}

func Debug(message string) {
	message = fmt.Sprintf("Protocol %s: %s", p.GetStr(), message)
	ptr, size := stringToPtr(message)
	_debug(ptr, size)
}

func Info(message string) {
	message = fmt.Sprintf("Protocol %s: %s", p.GetStr(), message)
	ptr, size := stringToPtr(message)
	_info(ptr, size)
}

func Warn(message string) {
	message = fmt.Sprintf("Protocol %s: %s", p.GetStr(), message)
	ptr, size := stringToPtr(message)
	_warn(ptr, size)
}

func Error(message string) {
	message = fmt.Sprintf("Protocol %s: %s", p.GetStr(), message)
	ptr, size := stringToPtr(message)
	_error(ptr, size)
}

//go:wasmimport env debug
func _debug(ptr uint32, size uint32)

//go:wasmimport env info
func _info(ptr uint32, size uint32)

//go:wasmimport env warn
func _warn(ptr uint32, size uint32)

//go:wasmimport env error
func _error(ptr uint32, size uint32)
