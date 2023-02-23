use std::fmt;

fn main() {
  println!("Hello World!");
  println!("I'm a Rustacean!");

  println!("{} days", 31);
  println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");

  println!("{subject} {verb} {object}",
    object="ooo",
    subject="subsub",
    verb="vvvv",
  );

  println!("Base 10:               {}",   69420); // 69420
  println!("Base 2 (binary):       {:b}", 69420); // 10000111100101100
  println!("Base 8 (octal):        {:o}", 69420); // 207454
  println!("Base 16 (hexadecimal): {:x}", 69420); // 10f2c
  println!("Base 16 (hexadecimal): {:X}", 69420); // 10F2C

  println!("{number:>width$}", number=1, width=6);
  //     1
  println!("{number:>5}", number=1);
  //    1
  println!("{number:0<5}", number=1);
  //10000

  // println!("My name is {0}, {1} {0}", "Bond");

  let number: f64 = 1.0;
  let width: usize = 10;
  println!("{number:>width$}");
  //         1

  #[allow(dead_code)]
  struct Structure(i32);
  impl fmt::Display for Structure {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      write!(f, "struct.i32: {}", self.0)
    }
  }
  impl fmt::Debug for Structure {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      write!(f, "struct details: {}", self.0)
    }
  }

  println!("This struct `{}`. Display!", Structure(3));
  println!("This struct `{:?}`. Debug!", Structure(3));

  // TODO:
  // https://doc.rust-jp.rs/rust-by-example-ja/hello/print.html の以下から
  // println!マクロを追加し、表示される小数部の桁数を調整してPi is roughly 3.142という文字列を出力しましょう。 ただし、円周率の値はlet pi = 3.141592を使ってください。（ヒント: 小数部の桁数を調整する方法については、std::fmtをチェックする必要があるかもしれません。）

}