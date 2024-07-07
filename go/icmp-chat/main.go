package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
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
	var dstIPAddr string
	description := "specify destination ip address"
	// TODO: rename "peer"
	flag.StringVar(&dstIPAddr, "dst", "", description)
	flag.Parse()

	if len(dstIPAddr) == 0 {
		log.Fatal(description)
	}

	log.Println("start icmp chat!")

	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go func() {
		for {
			rb := make([]byte, 1500)
			n, peer, err := c.ReadFrom(rb)
			if err != nil {
				log.Fatal(err)
			}

			_ = peer
			// 経路の途中でSNATされるとダメ
			// if peer.String() != dstIPAddr {
			// 	log.Printf("received from unknown destination: %s\n", peer)
			// 	continue
			// }

			const PROTOCOL_NUM_ICMPv4 = 1
			rm, err := icmp.ParseMessage(PROTOCOL_NUM_ICMPv4, rb[:n])
			if err != nil {
				log.Fatal(err)
			}
			switch rm.Type {
			case ipv4.ICMPTypeEchoReply:
				// log.Printf("passive icmp echo reply from %v", peer)
				// noop
			case ipv4.ICMPTypeEcho:
				// log.Printf("passive icmp echo from %v", peer)

				b, err := rm.Body.Marshal(PROTOCOL_NUM_ICMPv4)
				if err != nil {
					log.Printf("marshal err: %v\n", err)
					continue
				}

				// log.Printf("data: %x\n", b)
				fmt.Printf("> %s\n", string(b))
			default:
				log.Printf("got %+v; want echo reply", rm)
			}
		}
	}()

	go func() {
		fmt.Print("< ")
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			fmt.Print("< ")
			// https://www.infraexpert.com/study/tcpip4.html
			wm := icmp.Message{
				Type: ipv4.ICMPTypeEcho,
				Code: 0,
				Body: &icmp.Echo{
					ID:  os.Getpid() & 0xffff,
					Seq: 1,
					// Data: []byte("HELLO-R-U-THERE!!!!!"),
					Data: []byte(sc.Text()),
				},
			}
			wb, err := wm.Marshal(nil)
			if err != nil {
				log.Fatal(err)
			}

			if _, err := c.WriteTo(wb, &net.IPAddr{
				IP: net.ParseIP(dstIPAddr),
			}); err != nil {
				log.Fatal(err)
			}
		}

		if sc.Err() != nil {
			log.Fatal(sc.Err())
			return
		}
	}()

	time.Sleep(time.Second * 500)
}
