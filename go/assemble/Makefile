# アセンブル
asm:
	go tool compile -S main.go > main.s

bin:
	go build -o main_bin main.go

# 逆アセンブル
revasm: bin
	objdump -d main_bin > main_rev.s
