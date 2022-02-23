# websocket試し

## server <-> clientで双方向通信
- server

```console
cd cmd/server
go run main.go
```

- client1

```console
cd cmd/client
go run main.go
```

- client2

```console
cd cmd/client
go run main.go
```

- output
```console
# server
19:12:59 > go run main.go
2022/02/23 19:13:31 recv: wow!!, 0
2022/02/23 19:13:33 recv: wow!!, 1
2022/02/23 19:13:35 recv: wow!!, 2
2022/02/23 19:13:37 recv: wow!!, 3
2022/02/23 19:13:38 recv: wow!!, 0
2022/02/23 19:13:39 recv: wow!!, 4
2022/02/23 19:13:40 recv: wow!!, 1
2022/02/23 19:13:41 recv: wow!!, 5
...

# client1
19:13:13 > go run main.go
2022/02/23 19:13:31 recv: hello!!, 0
2022/02/23 19:13:33 recv: hello!!, 1
2022/02/23 19:13:35 recv: hello!!, 2
2022/02/23 19:13:37 recv: hello!!, 3
2022/02/23 19:13:39 recv: hello!!, 4
2022/02/23 19:13:41 recv: hello!!, 5
...

# client2
19:13:22 > go run main.go
2022/02/23 19:13:38 recv: hello!!, 0
2022/02/23 19:13:40 recv: hello!!, 1
2022/02/23 19:13:42 recv: hello!!, 2
2022/02/23 19:13:44 recv: hello!!, 3
2022/02/23 19:13:46 recv: hello!!, 4
2022/02/23 19:13:48 recv: hello!!, 5
...

```



## gorillaのexample

### websocketでchatを試す

- server起動
```console
cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/chat`
go run *.go
```

- chat画面オープン
```console
open http://localhost:8080
```
