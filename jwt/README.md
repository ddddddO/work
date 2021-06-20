# JWT
- https://jwt.io/introduction
    - 何かあればこちら確認するとよさそう。
- https://jwt.io/#debugger-io
    - JWTのデバッガー。JWT生成を試せる。

- JWTの構造
    - ドット区切りで以下な形。
        - `xxxxx.yyyyy.zzzzz`
    - ヘッダ（`xxxxx` の部分）
    - ペイロード（`yyyyy` の部分）
    - 署名（`zzzzz` の部分）

- Use Case
    - Authorization: 認可したいとき。一般的なJWTの使い方らしい。
    - Information Exchange: 本人確認可能な情報交換したい
        - 公開鍵と秘密鍵で署名できるから本人であると確認できる。
        - 署名は、ヘッダとペイロードから計算されるから改ざん検知可能。
    - [認証と認可の違い](https://qiita.com/kaysquare1231/items/c4e4736f2a924b03777b)

- メモ
    - 署名のフローを把握する。
    - 一旦、ドキュメントとデバッガーの値を確認しながら実装してみた（jwt.go）。
        - これをどう使っていくのか？を次に試していく。