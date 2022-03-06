refs:
- https://github.com/sago35/tinygo-workshop
- [TinyGo, VSCode, gopls, wioterminal](https://qiita.com/sago35/items/c30cbce4a0a3e12d899c)
- 回路図の確認
    >tips: 回路図を確認する
    組込みにおいて各種センサー等がどこにどのように繋がっているかを確認しておく必要があります。 Wio Terminal の場合は以下にまとまっているので確認しておくと良いです。上記の図は、下記 pdf からの抜粋になります。

    - https://wiki.seeedstudio.com/Wio-Terminal-Getting-Started/
    - https://files.seeedstudio.com/wiki/Wio-Terminal/res/Wio-Terminal-SCH-v1.2.pdf


# メモ
- VSCodeから開けるPowershellだとうまくいかん
- Powershellから`code .` でvs codeを開く
- Powershellから書き込む
    - e.g. ` tinygo flash --target wioterminal --size short .\01_blinky\`