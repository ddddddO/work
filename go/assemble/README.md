# 「もっとCPUの気持ちが知りたいですか？(出村成和)」
- Goは出てこない
- CPUの概要を知るのに良かった

## インデックスを張るためのメモ
### 機械語（CPUが解釈・実行するプログラム。main_bin）/アセンブリ言語（機械語を人間が読みやすい書式で表現したもの。main.s）
- 相互に変換可能

### CPUを構成する回路とその周辺（P30 図も参照）
1. （メインメモリ）
2. （周辺機器）
3. 外部インタフェース（1,2の出入口）
4. インストラクションキャッシュ（メインメモリから命令を取得・キャッシュ）
5. データキャッシュ（データをキャッシュ）
6. レジスタ（演算で参照する値・演算結果が格納される）
    - 汎用レジスタ・浮動小数点レジスタ・専用レジスタ（**P41~**）があり、役割りが分かれている。
7. ALU（四則演算・論理演算をする）
8. バス（各回路を結ぶ信号線）
9. コントローラー（各回路や回路間のバスを流れるデータを制御する）

### 命令セットアーキテクチャ（実行する命令やレジスタなどCPUがソフトウェアによってどのように制御されるかを定義したモデル）/マイクロアーキテクチャ（命令セットアーキテクチャに必要な構造のこと）

### CPUの処理はすべて「命令」として指示する
- 命令の分類（**P52**）

### RISCとCISC（どちらもCPUの設計手法）
- ロード・ストアアーキテクチャ/レジスター・メモリアーキテクチャ

### バイトオーダー（**P86**）
- 2バイト以上で表す数値をメモリに格納する順序
- （ネットワークバイトオーダーはビッグエンディアンであるとされている）

### パイプライン（内部の処理は**P97~**）
- 1つの命令を複数の処理に分割して、複数の命令をずらして実行する仕組み
- A、B、Cという3つの処理が各々1つずつ実行できるとして、a(A)の処理開始->a(B)の処理開始時に、b(A)の処理を開始させて...という感じに効率的に処理させるイメージ
- 多段パイプライン
- 分岐予測と投機的実行

### キャッシュメモリ（**P111~**）
- 無いと速度面でかなり辛い。メインメモリから都度取得するでは遅い。CPU内部に複数ある
- インストラクションキャッシュ・データキャッシュ
- L1キャッシュ・L2キャッシュ（・L3キャッシュ）
- P131で試してみたいコマンドがある（`perf stat`）

### メモリマップドI/O（**P137~**）
- TODO: **再度読む**

### 割り込み

# [解説&翻訳 - A Quick Guide to Go's Assembler](https://zenn.dev/hsaki/articles/godoc-asm-ja)
# [Goアセンブリ入門](https://qiita.com/Akatsuki_py/items/231350711f9ab6eba95e)