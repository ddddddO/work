build_wasm:
	tinygo build -o main_wasi.wasm -target wasi main.go

run_wasi:
	wasmer main_wasi.wasm

# NOTE: ダメだった
# cd /home/ddd/
#     ↓
# https://github.com/wasmerio/wasmer-php#install <- once execute
up_php:
	docker run -v /home/ochi/github.com/ddddddO/work/php/wasm:/home/ddd -it php:8.0-fpm bash
