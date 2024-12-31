build_dhcp:
	tinygo build -o ../.plugin-wasm/dhcp.wasm -scheduler=none --no-debug -target=wasi dhcp.go

build_imap:
	tinygo build -o ../.plugin-wasm/imap.wasm -scheduler=none --no-debug -target=wasi imap.go

build: build_dhcp build_imap
