package main

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/txthinking/socks5"
)

// 準備
// - ~/github.com/ddddddO/gtree/docs 配下で、以下コマンド実行でターゲットのHTTPサーバー起動
//   - python3 -m http.server 8080
//
// - SOCKS5サーバー(このGoプログラム)起動後、端末で以下を実行で動作確認
//   - curl --proxy socks5://127.0.0.1:1080 http://localhost:8080
func main() {
	fmt.Println("--- Start ---")

	srv, err := socks5.NewClassicServer("127.0.0.1:1080", "127.0.0.1", "", "", 0, 60)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	proxy := &proxyHandler{
		DefaultHandle: &socks5.DefaultHandle{},
	}
	if err := srv.ListenAndServe(proxy); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

type proxyHandler struct {
	*socks5.DefaultHandle
}

func (p *proxyHandler) TCPHandle(srv *socks5.Server, conn *net.TCPConn, req *socks5.Request) error {
	slog.Info("--- In proxyHandler ---")

	p.debug(req)

	// conn を io.Copy しても止まる。EOFでないからっぽい(curl)
	// FireFoxからだと大丈夫そう？
	// go io.Copy(os.Stdout, conn)

	// 多分、送信元リクエストに手を加えるなら conn に write する
	// 今は、外部の socks5 package の DefaultHandle.TCPHandle を呼んでるけど、自前で用意した方が writeするとき都合がいいかも?
	return p.DefaultHandle.TCPHandle(srv, conn, req)
}

func (*proxyHandler) debug(req *socks5.Request) {
	out := os.Stdout
	// fmt.Fprintf(out, "Request: %+v\n", req)
	fmt.Fprintf(out, "Ver     : %d\n", uint8(req.Ver))
	fmt.Fprintf(out, "Cmd     : %d\n", uint8(req.Cmd))
	fmt.Fprintf(out, "Rsv     : %d\n", uint8(req.Rsv))
	fmt.Fprintf(out, "Atyp    : %d\n", uint8(req.Atyp))
	fmt.Fprintf(out, "Dst Addr: %d.%d.%d.%d\n", uint8(req.DstAddr[0]), uint8(req.DstAddr[1]), uint8(req.DstAddr[2]), uint8(req.DstAddr[3]))
	fmt.Fprintf(out, "Dst Port: %d\n", binary.BigEndian.Uint16(req.DstPort))
	fmt.Fprintln(out)
	// Output(curl --proxy socks5://127.0.0.1:1080 http://localhost:8080):
	// 2023/12/24 20:01:13 INFO --- In proxyHandler ---
	// Ver     : 5
	// Cmd     : 1
	// Rsv     : 0
	// Atyp    : 1
	// Dst Addr: 127.0.0.1
	// Dst Port: 8080

	// Output(from FireFox: https://ddddddo.github.io/gtree/)
	// 2023/12/24 20:03:33 INFO --- In proxyHandler ---
	// Ver     : 5
	// Cmd     : 1
	// Rsv     : 0
	// Atyp    : 1
	// Dst Addr: 185.199.108.153
	// Dst Port: 443

	// ref...
	// ~/github.com/ddddddO/work/go/proxy
	// 20:04:15 > nslookup github.io
	// Server:         172.23.240.1
	// Address:        172.23.240.1#53
	//
	// Non-authoritative answer:
	// Name:   github.io
	// Address: 185.199.109.153
	// Name:   github.io
	// Address: 185.199.110.153
	// Name:   github.io
	// Address: 185.199.111.153
	// Name:   github.io
	// Address: 185.199.108.153
	//
	// ~/github.com/ddddddO/work/go/proxy
	// 20:04:29 >
}
