# 公式サイト
https://yaml.org/  

# 用語
- スカラー：値
    - 文字列(aaa, "aaa", 'aaa' のどれでも可)
    - 数値
    - 真偽値 (true/false, yes/no, on/off)
    - null (`~`で表わすこともできる)
    - 日付け
- シーケンス：配列
- マッピング：ハッシュ
- アンカー：変数
- エイリアス：変数呼び出し
- --- ：データの始まり
- ...：データの終わり(省略可)

# Try
`ruby parser.rb` で`sample.yml`をパースして出力
`sample.yml`内で`---`と`...`でデータを区切っているため、パースする場合は、対象のデータ以外をコメントアウトして実行する。

# Memo
rubyのto_yamlメソッドで、rubyのオブジェクトからyamlへ変換したものを出力してくれるので便利そう。