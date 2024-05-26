### https://ebpf-go.dev/guides/getting-started/#the-go-application

に沿ってやってみた。別ターミナルでcurlするとカウントアップされてるのが確認できた

```console
~/github.com/ddddddO/work/go/ebpf/hello
00:09:48 > sudo ./ebpf-test 
2024/05/27 00:09:56 Counting incoming packets on eth0..
2024/05/27 00:09:57 Received 0 packets
2024/05/27 00:09:58 Received 0 packets
2024/05/27 00:09:59 Received 0 packets
2024/05/27 00:10:00 Received 0 packets
2024/05/27 00:10:01 Received 0 packets
2024/05/27 00:10:02 Received 0 packets
2024/05/27 00:10:03 Received 0 packets
2024/05/27 00:10:04 Received 0 packets
2024/05/27 00:10:05 Received 0 packets
2024/05/27 00:10:06 Received 0 packets
2024/05/27 00:10:07 Received 0 packets
2024/05/27 00:10:08 Received 0 packets
2024/05/27 00:10:09 Received 0 packets
2024/05/27 00:10:10 Received 0 packets
2024/05/27 00:10:11 Received 1 packets
2024/05/27 00:10:12 Received 1 packets
2024/05/27 00:10:13 Received 4 packets
2024/05/27 00:10:14 Received 5 packets
2024/05/27 00:10:15 Received 7 packets
2024/05/27 00:10:16 Received 8 packets
2024/05/27 00:10:17 Received 9 packets
2024/05/27 00:10:18 Received 9 packets
2024/05/27 00:10:19 Received 10 packets
2024/05/27 00:10:20 Received 10 packets
2024/05/27 00:10:21 Received 16 packets
2024/05/27 00:10:22 Received 16 packets
2024/05/27 00:10:23 Received 16 packets
2024/05/27 00:10:24 Received 16 packets
2024/05/27 00:10:25 Received 16 packets
2024/05/27 00:10:26 Received 16 packets
2024/05/27 00:10:27 Received 16 packets
2024/05/27 00:10:28 Received 16 packets
2024/05/27 00:10:29 Received 16 packets
2024/05/27 00:10:30 Received 17 packets
2024/05/27 00:10:31 Received 19 packets
2024/05/27 00:10:32 Received 21 packets
2024/05/27 00:10:33 Received 23 packets
2024/05/27 00:10:34 Received 24 packets
2024/05/27 00:10:35 Received 24 packets
2024/05/27 00:10:36 Received 24 packets
2024/05/27 00:10:37 Received 24 packets
2024/05/27 00:10:38 Received 24 packets
2024/05/27 00:10:39 Received 26 packets
2024/05/27 00:10:40 Received 27 packets
2024/05/27 00:10:41 Received 38 packets
2024/05/27 00:10:42 Received 208 packets
2024/05/27 00:10:43 Received 208 packets
2024/05/27 00:10:44 Received 208 packets
2024/05/27 00:10:45 Received 208 packets
2024/05/27 00:10:46 Received 208 packets
2024/05/27 00:10:47 Received 208 packets
^C2024/05/27 00:10:48 Received signal, exiting..
~/github.com/ddddddO/work/go/ebpf/hello
00:10:48 > 
```
