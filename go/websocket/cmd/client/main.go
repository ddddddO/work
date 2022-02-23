package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// normal()

	const target = "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(target, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 2)

		messageType := websocket.TextMessage
		cnt := 0
		for range ticker.C {
			p := []byte(fmt.Sprintf("wow!!, %d", cnt))
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
			cnt++
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, os.Interrupt)
	<-sig

	log.Println("end")

}

// serverへ書き込み出来なさそう
func normal() {
	const target = "http://localhost:8080/ws"
	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Upgrade", "websocket")
	req.Header.Add("Connection", "upgrade")
	req.Header.Add("Sec-Websocket-Version", "13")
	req.Header.Add("Sec-WebSocket-Key", "aaa")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// defer resp.Body.Close()

	go func() {
		ticker := time.NewTicker(time.Second * 2)
		for range ticker.C {
			buf := make([]byte, 1024)
			if _, err := resp.Body.Read(buf); err != nil {
				log.Fatal(err)
			}
			log.Print(string(buf))
		}
	}()
}
