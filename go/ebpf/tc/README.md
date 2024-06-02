- packemonでARPリクエストし、loadしたebpfプログラムのarpCountがupされてることを確認した

- 実行手順

```console
# 別ターミナルで tmux で画面分割して、packemon の Monitor/Generatorを起動
$ cd ~/github.com/ddddddO/packemon
$ sudo go run cmd/packemon/main.go
$ sudo go run cmd/packemon/main.go  --send

# egress_packet.c 修正したら以下
$ go generate

# 実行
$ sudo go run .
```

