- packemonでARPリクエストし、loadしたebpfプログラムのarpCountがupされてることを確認した

- 394c007 のコミットで、ARPリクエストをdrop出来たことを確認した
  - egress_packet.c の `return TC_ACT_SHOT;` のコメントアウトを外して実行で
  - 修正前は、packemon の Monitor で表示されてたが、修正後は表示されなくなった。また、修正前後でARPリクエストでarpCountはupされていることも確認

- 実行手順

```console
# 別ターミナルで tmux で画面分割して、packemon の Monitor を起動
$ cd ~/github.com/ddddddO/packemon
$ sudo go run cmd/packemon/main.go

# egress_packet.c 修正したら以下
$ go generate

# 実行
$ sudo go run .

$ arping -c 1 8.8.8.8
```

