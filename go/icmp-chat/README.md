- TailscaleのWeb管理画面でRemoveした後は、各端末で `sudo systemctl restart tailscaled` -> `sudo tailscale up`

- WSL側でキャプチャしてみる
  `sudo tcpdump -U -i tailscale0 -w - | /mnt/c/Program\ Files/Wireshark/Wireshark.exe -k -i -`