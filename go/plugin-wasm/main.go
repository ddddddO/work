package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
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

type Protocol interface {
	Name(x string) string
	Port() uint64
}

func run() error {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

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
		// dhcp.wasm
		mod, err := r.Instantiate(ctx, wasmCode)
		if err != nil {
			return err
		}

		nameFunc := mod.ExportedFunction("Name")
		if nameFunc == nil {
			return fmt.Errorf("required Name function")
		}
		portFunc := mod.ExportedFunction("Port")
		if portFunc == nil {
			return fmt.Errorf("required Port function")
		}

		compiled := &CompiledProtocol{
			mod:    mod,
			malloc: mod.ExportedFunction("malloc"),
			free:   mod.ExportedFunction("free"),
			name:   nameFunc,
			port:   portFunc,
		}

		compileds = append(compileds, compiled)
	}

	proto := compileds[0] // dhcp.wasm

	results, err := proto.Name(ctx, "9")
	if err != nil {
		return err
	}

	fmt.Println("Result:", results[0])

	return nil
}

type CompiledProtocol struct {
	mod    api.Module
	malloc api.Function
	free   api.Function

	name api.Function
	port api.Function
}

func (c *CompiledProtocol) Name(ctx context.Context, x string) ([]uint64, error) {
	inputPtr, size, err := stringToPtrSize(ctx, x, c.mod, c.malloc)
	if err != nil {
		return nil, err
	}
	defer c.free.Call(ctx, inputPtr)

	return c.name.Call(ctx, inputPtr, size)
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
