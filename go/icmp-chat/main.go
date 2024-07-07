package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/sync/errgroup"
)

/*
# icmpを待ち受ける機能
- icmp echo request きたら、replyを返す
	- 返せたら、data部分をパースして表示

# icmp echo requestを送る機能
- 送れたら、data部にメッセージつめて送る

# 多分これらで楽に作れるかな
- https://pkg.go.dev/golang.org/x/net/icmp
- https://github.com/google/gopacket

*/

// https://pkg.go.dev/golang.org/x/net/icmp#example-PacketConn-NonPrivilegedPing
func main() {
	var peerIPAddr string
	description := "specify peer ip address"
	flag.StringVar(&peerIPAddr, "peer", "", description)
	flag.Parse()

	if len(peerIPAddr) == 0 {
		log.Fatal(description)
	}

	log.Println("start icmp chat!")

	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	eg, ctx := errgroup.WithContext(context.Background())
	_ = ctx

	// TODO: この辺り、別にエラー返さないで継続でいいかも
	eg.Go(func() error {
		for {
			rb := make([]byte, 1500)
			n, peer, err := c.ReadFrom(rb)
			if err != nil {
				return err
			}

			_ = peer
			// 経路の途中でSNATされるとダメ
			// if peer.String() != peerIPAddr {
			// 	log.Printf("received from unknown destination: %s\n", peer)
			// 	continue
			// }

			const PROTOCOL_NUM_ICMPv4 = 1
			rm, err := icmp.ParseMessage(PROTOCOL_NUM_ICMPv4, rb[:n])
			if err != nil {
				return err
			}
			switch rm.Type {
			case ipv4.ICMPTypeEchoReply:
				// noop
			case ipv4.ICMPTypeEcho:
				b, err := rm.Body.Marshal(PROTOCOL_NUM_ICMPv4)
				if err != nil {
					log.Printf("marshal err: %v\n", err)
					continue
				}

				fmt.Printf("> %s\n", string(b))
			default:
				// noop
			}
		}
	})

	eg.Go(func() error {
		fmt.Print("< ")
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			fmt.Print("< ")
			// https://www.infraexpert.com/study/tcpip4.html
			wm := icmp.Message{
				Type: ipv4.ICMPTypeEcho,
				Code: 0,
				Body: &icmp.Echo{
					ID:   os.Getpid() & 0xffff,
					Seq:  1,
					Data: []byte(sc.Text()), // TODO: data部にこのアプリ独自のマークを仕込んで、読み取り側もそのマークであればパース、のようにした方がいいかも。普通のpingとかを出力してもノイズかな
				},
			}
			wb, err := wm.Marshal(nil)
			if err != nil {
				return err
			}

			if _, err := c.WriteTo(wb, &net.IPAddr{
				IP: net.ParseIP(peerIPAddr),
			}); err != nil {
				return err
			}
		}

		if sc.Err() != nil {
			return err
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
