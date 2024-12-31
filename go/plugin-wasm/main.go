package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	wasi "github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

const wasmDir = ".plugin-wasm"

// refs:
// https://qiita.com/uh-zz/items/3c8941d2059b6171e808
// https://github.com/tetratelabs/wazero/tree/main/examples/basic
// https://github.com/aquasecurity/trivy/blob/e8085bae3e71fc5c9839feb13e34b75deba4ce9d/pkg/module/module.go#L129
func main() {
	fmt.Println("plugin!")

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)
	// wasi_snapshot_preview1.MustInstantiate(ctx, r)

	// Instantiate a Go-defined module named "env" that exports functions.
	envBuilder := r.NewHostModuleBuilder("env")
	if _, err := envBuilder.Instantiate(ctx); err != nil {
		return fmt.Errorf("wasm module build error: %w", err)
	}

	if _, err := wasi.NewBuilder(r).Instantiate(ctx); err != nil {
		return fmt.Errorf("WASI init error: %w", err)
	}

	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	return err
	// }
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	wasmPath := filepath.Join(cwd, wasmDir)
	log.Println("wasm directory path:", wasmPath)

	finfo, err := os.Stat(wasmPath)
	if err != nil {
		return err
	}
	if !finfo.IsDir() {
		return fmt.Errorf("not directory")
	}

	wasmCodes := [][]byte{}
	if err := filepath.Walk(wasmPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(info.Name()) != ".wasm" {
			return nil
		}

		log.Println("wasm code path:", path)

		wasmCode, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		wasmCodes = append(wasmCodes, wasmCode)

		return nil
	}); err != nil {
		return err
	}

	compileds := []*CompiledProtocol{}
	for _, wasmCode := range wasmCodes {
		// Compile the WebAssembly module using the default configuration.
		compiledMod, err := r.CompileModule(ctx, wasmCode)
		if err != nil {
			return fmt.Errorf("module compile error: %w", err)
		}

		// InstantiateModule runs the "_start" function which is what TinyGo compiles "main" to.
		mod, err := r.InstantiateModule(ctx, compiledMod, wazero.NewModuleConfig())
		if err != nil {
			return fmt.Errorf("module init error: %w", err)
		}

		nameFunc := mod.ExportedFunction("name")
		if nameFunc == nil {
			return fmt.Errorf("required Name function")
		}
		portFunc := mod.ExportedFunction("port")
		if portFunc == nil {
			return fmt.Errorf("required Port function")
		}

		compiled := &CompiledProtocol{
			mod:    mod,
			malloc: mod.ExportedFunction("malloc"),
			free:   mod.ExportedFunction("free"),
			getStr: mod.ExportedFunction("get_str"),
			name:   nameFunc,
			port:   portFunc,
		}

		compileds = append(compileds, compiled)
	}

	proto := compileds[0] // dhcp.wasm

	str, err := proto.GetStr(ctx)
	if err != nil {
		return err
	}
	fmt.Println("GetStr result:", str)

	name, err := proto.Name(ctx, "DHCP!!")
	if err != nil {
		return err
	}
	fmt.Println("Name result:", name)

	port, err := proto.Port(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Port result:", port)
	// Output:
	// plugin!
	// 2024/12/31 23:18:58 wasm directory path: /home/ddddddo/github.com/ddddddO/work/go/plugin-wasm/.plugin-wasm
	// 2024/12/31 23:18:58 wasm code path: /home/ddddddo/github.com/ddddddO/work/go/plugin-wasm/.plugin-wasm/dhcp.wasm
	// GetStr result: xxxxxx
	// Name result: protocol name: DHCP!!
	// Port result: 68

	return nil
}

type CompiledProtocol struct {
	mod    api.Module
	malloc api.Function
	free   api.Function

	getStr api.Function
	name   api.Function
	port   api.Function
}

func (c *CompiledProtocol) GetStr(ctx context.Context) (string, error) {
	results, err := c.getStr.Call(ctx)
	if err != nil {
		return "", err
	}
	str, err := ptrSizeToString(c.mod.Memory(), results[0])
	if err != nil {
		return "", err
	}
	return str, nil
}

func (c *CompiledProtocol) Name(ctx context.Context, x string) (string, error) {
	inputPtr, size, err := stringToPtrSize(ctx, x, c.mod, c.malloc)
	if err != nil {
		return "", err
	}
	defer c.free.Call(ctx, inputPtr)

	results, err := c.name.Call(ctx, inputPtr, size)
	if err != nil {
		return "", err
	}

	var ret string
	if err = unmarshal(c.mod.Memory(), results[0], &ret); err != nil {
		return "", fmt.Errorf("invalid return value: %w", err)
	}

	return ret, nil
}

func (c *CompiledProtocol) Port(ctx context.Context) (uint64, error) {
	results, err := c.port.Call(ctx)
	if err != nil {
		return 0, err
	}
	return results[0], nil
}

// copied: https://github.com/aquasecurity/trivy/blob/e8085bae3e71fc5c9839feb13e34b75deba4ce9d/pkg/module/module.go#L209
// stringToPtr returns a pointer and size pair for the given string in a way compatible with WebAssembly numeric types.
func stringToPtrSize(ctx context.Context, s string, mod api.Module, malloc api.Function) (uint64, uint64, error) {
	size := uint64(len(s))
	results, err := malloc.Call(ctx, size)
	if err != nil {
		return 0, 0, fmt.Errorf("malloc error: %w", err)
	}

	// The pointer is a linear memory offset, which is where we write the string.
	ptr := results[0]
	if !mod.Memory().Write(uint32(ptr), []byte(s)) {
		return 0, 0, fmt.Errorf("Memory.Write(%d, %d) out of range of memory size %d",
			ptr, size, mod.Memory().Size())
	}

	return ptr, size, nil
}

func ptrSizeToString(mem api.Memory, ptrSize uint64) (string, error) {
	ptr, size := splitPtrSize(ptrSize)
	buf := readMemory(mem, ptr, size)
	if buf == nil {
		return "", fmt.Errorf("unable to read memory")
	}
	return string(buf), nil
}

func splitPtrSize(u uint64) (uint32, uint32) {
	ptr := uint32(u >> 32)
	size := uint32(u)
	return ptr, size
}

func readMemory(mem api.Memory, offset, size uint32) []byte {
	buf, ok := mem.Read(offset, size)
	if !ok {
		log.Printf("Memory.Read() out of range: offset: %d, size: %d", int(offset), int(size))
		return nil
	}
	return buf
}

func marshal(ctx context.Context, m api.Module, malloc api.Function, v any) (uint64, uint64, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return 0, 0, fmt.Errorf("marshal error: %w", err)
	}

	size := uint64(len(b))
	results, err := malloc.Call(ctx, size)
	if err != nil {
		return 0, 0, fmt.Errorf("malloc error: %w", err)
	}

	// The pointer is a linear memory offset, which is where we write the marshaled value.
	ptr := results[0]
	if !m.Memory().Write(uint32(ptr), b) {
		return 0, 0, fmt.Errorf("Memory.Write(%d, %d) out of range of memory size %d",
			ptr, size, m.Memory().Size())
	}

	return ptr, size, nil
}

func unmarshal(mem api.Memory, ptrSize uint64, v any) error {
	ptr, size := splitPtrSize(ptrSize)
	buf := readMemory(mem, ptr, size)
	if buf == nil {
		return fmt.Errorf("unable to read memory")
	}
	if err := json.Unmarshal(buf, v); err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	return nil
}
