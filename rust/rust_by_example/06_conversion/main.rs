// Rustはカスタム型（例えばstructやenum）間の変換をトレイトを用いて行います。
// ジェネリックな型変換にはFromおよびIntoトレイトを使用します。
// しかし、よくあるケースにおいて、特にStringとの相互の型変換では、特殊なトレイトが使用されます。

fn main() {
  // Fromトレイトは、ある型に対し、別の型からその型を作る方法を定義できるようにするものです。
  // そのため、複数の型の間で型変換を行うための非常にシンプルな仕組みを提供しています。
  // 標準ライブラリでは、基本データ型やよく使われる型に対して、このトレイトが多数実装されています。
  let my_str = "hello";
  let my_string = String::from(my_str); // strからStringへの型変換
  println!("{}", my_string);
  // hello

  // 自作の型にFromトレイト実装
  #[derive(Debug)]
  struct Number {
    value: i32,
  }
  impl std::convert::From<i32> for Number {
    fn from(item: i32) -> Self {
      Number{ value: item }
    }
  }
  let num = Number::from(100);
  println!("{:?}", num);
  // Number { value: 100 }

  // Intoトレイトは、単にFromトレイトの逆の働きをします。
  let n = 29;
  let num2: Number = n.into(); // Intoトレイトを使用すると、ほとんどの場合、コンパイラが型を決定することができないため、変換する型を指定する必要があります。
  println!("{:?}", num2);
  // Number { value: 29 }

  // TryFromおよびTryIntoも型変換を行うジェネリックなトレイトです。
  // From/Intoと異なり、TryFrom/TryIntoトレイトは失敗する可能性のある型変換に用いられるので、Resultを返します。

  // 任意の型をStringに変換するのは簡単で、その型にToStringトレイトを実装するだけです。
  // これを直接実装するよりも、fmt::Displayトレイトを実装するのがよいでしょう。
  // そうすることで自動的にToStringが提供されるだけでなく、print!の章で説明したように、その型を表示できるようにもなります。

  analyze_string();
}

fn analyze_string() {
  // 文字列からの型変換において、数値への型変換はよく行われるものの一つです。
  // これを行うイディオムはparse関数を使用することですが、このときに型を推論できるようにするか、
  // もしくは turbofish構文を使用して型を指定するかのいずれかを行います。以下の例では、どちらの方法も紹介しています。

  // parse関数は、指定された型にFromStrトレイトが実装されていれば、文字列をその型に変換します。
  // このトレイトは標準ライブラリの多くの型に対して実装されています。
  // ユーザー定義の型でこの機能を利用するには、その型に対してFromStrトレイトを実装するだけです。
  let parsed: i32 = "5".parse().unwrap(); // 型推論
  let turbo_parsed = "10".parse::<i32>().unwrap(); // turbofish構文を使用して型を指定

  let sum = parsed + turbo_parsed;
  println!("Sum: {:?}", sum);
  // Sum: 15
}