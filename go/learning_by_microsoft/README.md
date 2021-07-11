- https://docs.microsoft.com/ja-jp/learn/paths/go-first-steps/

- 2021/07/11実施


# メモ
- モジュール化する（他のGoコードからimportしてもらう）
    - [ref](https://docs.microsoft.com/ja-jp/learn/modules/go-variables-functions-packages/4-packages)の「モジュールを作成する」
    - `go mod init` する必要がある。

- 「Go では、if ブロック内で変数を宣言するのが特徴です。つまり、これは、Go で頻繁に実行される規約を使用して効率的にプログラミングする方法であることを意味します。」

- switch, case , `fallthrough`
    - fallthrough キーワードを使用する場合は注意が必要です。
- 「組み込みの panic() 関数を使用すると、通常の制御フローが停止されます。 すべての遅延関数呼び出しは、正常に実行されます。」
    - [ref](https://docs.microsoft.com/ja-jp/learn/modules/go-control-flow/4-use-defer-statement)
- 「Go の組み込み関数 recover() を使用すると、パニックの後で制御を取り戻すことができます。 この関数は、遅延された関数の中でのみ使用できます。」
    - [ref](https://docs.microsoft.com/ja-jp/learn/modules/go-control-flow/4-use-defer-statement)
    - 「panic と recover の組み合わせは、Go での特徴的な例外処理方法です。 他のプログラミング言語では、try/catch ブロックが使用されます。 Go の場合、ここで調べたようなアプローチが好まれます。」
- [スライスについて（長さ、容量）](https://docs.microsoft.com/ja-jp/learn/modules/go-data-types/2-slices)
    - うーん
    - 「 "それより多くの要素を保持するだけの十分な容量がスライスにない場合、Go によってその容量が 2 倍にされます"。 新しい基になる配列が、新しい容量で作成されます。 ユーザーは容量を増やすために何もする必要はありません。 Go によって自動的に実行されます。 注意する必要があります。 ある時点で、スライスの容量が必要以上に多くなり、メモリが無駄に消費されるおそれがあります。」
    - nil mapに要素追加しようとしないように。パニックになる。以下をやらないように。
        - ```go
          var studentsAge map[string]int
          studentsAge["john"] = 32
          ```
- JSON, 構造体, デコード・エンコードは以下を見ればよさそう。
    - [ref](https://docs.microsoft.com/ja-jp/learn/modules/go-data-types/4-structs) の「JSON を使用して構造体をエンコードおよびデコードする」

- 「また、規則では、エラー変数に Err プレフィックスを含めることになっているため、注意してください。」
    - [ref](https://docs.microsoft.com/ja-jp/learn/modules/go-errors-logs/1-errors) の「再利用可能なエラーの作成」

- [Go でエラーを処理する場合は、次のような推奨事項を考慮してください。](https://docs.microsoft.com/ja-jp/learn/modules/go-errors-logs/1-errors)

    - エラーが予想されていなくても、エラーがないかどうかを常に確認します。 そのうえで、それらを適切に処理して、不要な情報がエンド ユーザーに公開されないようにします。
    - エラー メッセージにプレフィックスを含めて、エラーの発生元がわかるようにします。 たとえば、パッケージや関数の名前を含めることができます。
    - 可能な限り、再利用可能なエラー変数を作成します。
    - エラーを返すことと、パニックの違いを理解します。 他に対処方法がない場合は、パニックが発生します。 たとえば、依存関係の準備ができていない場合、プログラムは動作しません (既定の動作を実行する場合を除きます)。
    - 可能な限り多くの詳細情報でエラーをログに記録し (次のセクションでその方法を説明します)、エンド ユーザーが理解可能なエラーを出力します。
- 「log.SetPrefix() があります。 これを使用すると、プログラムのログ メッセージにプレフィックスを追加できます。」

- 「出力が JSON 形式に変更されることに注目してください。 JSON は、集中管理された場所で検索を実行する場合に便利なログの形式です。」
    - [ref](https://docs.microsoft.com/ja-jp/learn/modules/go-errors-logs/2-logs) の「ログ記録フレームワーク」
    - zerologとかzapのドキュメントを確認した方がよさそう。

- 「オブジェクト指向プログラミング (OOP) は、ほとんどの (少なくともある程度の) プログラミング言語でサポートされている一般的なプログラミング パラダイムです。 Go はこれらの言語の 1 つですが、OOP のすべての原則が完全にサポートされているわけではありません。」
    - [ref](https://docs.microsoft.com/ja-jp/learn/modules/go-methods-interfaces/0-introduction)

- 「メソッドで変数を更新する必要がある場合があります。また、引数が大きすぎる場合に、そのコピーを回避したいことがあります。 このような場合、ポインターを使用して、変数のアドレスを渡す必要があります。」
    - [ref](https://docs.microsoft.com/ja-jp/learn/modules/go-methods-interfaces/1-methods) の「メソッド内のポインター」
    - 「メソッドがレシーバーの情報にアクセスするだけの場合、レシーバーの変数にポインターは必要ありません。 ただし、Go の規則では、構造体のいずれかのメソッドにポインター レシーバーがある場合、メソッドには必要がなくても、その構造体のすべてのメソッドにポインター レシーバーが必要であると定められています。」
- 「Go ではメソッドを "オーバーライド" しても、必要に応じて "元の" ものにアクセスすることができます。」
    - [ref](https://docs.microsoft.com/ja-jp/learn/modules/go-methods-interfaces/1-methods) の「メソッドをオーバーロードする」
        - 「オーバーライド」が正しいと思う。

- [Go のコンカレンシーのしくみについて学ぶ](https://docs.microsoft.com/ja-jp/learn/modules/go-concurrency/) のすべて
    - [「多くのプログラムでは、次のように匿名関数を使用して goroutine を作成することが好まれています。」](https://docs.microsoft.com/ja-jp/learn/modules/go-concurrency/1-goroutines)
    - [「Go のチャネルは、goroutine 間の通信メカニズムです。 これが、Go のコンカレンシーに対するアプローチが "メモリを共有して通信しないでください。そうではなく、通信してメモリを共有してください" であると説明した理由です。 ある goroutine から別の goroutine に値を送信する必要がある場合は、チャネルを使用します。」](https://docs.microsoft.com/ja-jp/learn/modules/go-concurrency/2-channels)
    - [「チャネルを閉じると、そのチャネルでデータは再び送信されなくなります。 閉じたチャネルにデータを送信しようとすると、プログラムはパニックになります。 また、閉じたチャネルからデータを受信しようとすると、送信されたすべてのデータを読み取ることができます。 それ以降、"読み取り" を行うたびにゼロ値が返されます。」](https://docs.microsoft.com/ja-jp/learn/modules/go-concurrency/2-channels)
    - [「チャネルに何かを送信するたびに、要素がキューに追加されます。 次に、受信操作によって要素がキューから削除されます。 チャネルがいっぱいになると、データを保持する領域ができるまで、送信操作は待機されます。 逆に、チャネルが空で読み取り操作がある場合、読み取る対象が発生するまでチャネルはブロックされます。」 (バッファ有りチャネル)](https://docs.microsoft.com/ja-jp/learn/modules/go-concurrency/3-buffered-channels)
    - [「チャネルを使用するときは、常に goroutine を使用することをお勧めします。」](https://docs.microsoft.com/ja-jp/learn/modules/go-concurrency/3-buffered-channels)
    - [バッファーなしとバッファーありのチャネル](https://docs.microsoft.com/ja-jp/learn/modules/go-concurrency/3-buffered-channels)
        - 「バッファーなしのチャネルの場合、同期的に通信されます。 データを送信するたびに、誰かがチャネルから読み取るまでプログラムがブロックされることが保証されます。」
        - 「バッファーありのチャネルの場合、送信操作と受信操作は分離されています。 プログラムはブロックされませんが、(前述のように) デッドロックが発生する可能性があるので注意する必要があります。 バッファーなしのチャネルを使用する場合、同時に実行できる goroutine の数を制御できます。」

- [こちら](https://docs.microsoft.com/ja-jp/learn/modules/go-write-test-program/0-introduction)の総集編的な章いいと思う。アツい。
    - これまでの学習を踏まえたアプリ開発
    - DB使わない。基本コピペで進められる。
        - **最後、お題の機能追加を自分で考えて実装する**
    - なにをつくるか？の要件が提示される
    - TDDで進む(red -> greenになるように)