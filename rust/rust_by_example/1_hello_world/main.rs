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
  // This struct `struct.i32: 3`. Display!
  println!("This struct `{:?}`. Debug!", Structure(3));
  // This struct `struct details: 3`. Debug!

  let pi = 3.141592;
  println!("Pi is roughly {:.3}", pi);
  // Pi is roughly 3.142

  println!("{:?} months in a year.", 12);
  // 12 months in a year.
  #[derive(Debug)]
  struct Structure2(i32);
  #[derive(Debug)]
  struct Deep(Structure2);
  println!("Now {:?} will print!", Structure2(3));
  // Now Structure2(3) will print!
  println!("Now {:#?} will print!", Structure2(3));
  // Now Structure2(
  //     3,
  // ) will print!
  println!("Now {:?} will print!", Deep(Structure2(7)));
  // Now Deep(Structure2(7)) will print!
  println!("Now {:#?} will print!", Deep(Structure2(7)));
  // Now Deep(
  //   Structure2(
  //       7,
  //   ),
  // ) will print!

  struct Complex {
    real: f64,
    imag: f64,
  }
  impl fmt::Display for Complex {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      write!(f, "{} + {}i", &self.real, &self.imag)
    }
  }
  let complex = Complex{real: 3.3, imag: 7.2};
  println!("Display: {}", complex);
  // Display: 3.3 + 7.2i
  impl fmt::Debug for Complex {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      write!(f, "Complex {{ real: {}, imag: {} }}", &self.real, &self.imag)
    }
  }
  println!("Debug: {:?}", complex);
  // Debug: Complex { real: 3.3, imag: 7.2 }

  struct List(Vec<i32>);
  impl fmt::Display for List {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      let vec = &self.0;
      write!(f, "[")?;

      for (count, v) in vec.iter().enumerate() {
        if count != 0 { write!(f, ", ")?; }
        write!(f, "{}: {}", count, v)?;
      }

      write!(f, "]")
    }
  }
  let v = List(vec![1, 2, 3, 4]);
  println!("{}", v);
  // [0: 1, 1: 2, 2: 3, 3: 4]

  struct Color {
      red: u8,
      green: u8,
      blue: u8,
  }
  impl fmt::Display for Color {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      write!(f, "RGB ({r}, {g}, {b}) {r:#02X}{g:02X}{b:02X}",
        r=&self.red, g=&self.green, b=&self.blue)
    }
  }
  for color in [
    Color { red: 128, green: 255, blue: 90 },
    Color { red: 0, green: 3, blue: 254 },
    Color { red: 0, green: 0, blue: 0 },
  ].iter() {
    println!("{}", *color);
  }
  // RGB (128, 255, 90) 0x80FF5A
  // RGB (0, 3, 254) 0x003FE
  // RGB (0, 0, 0) 0x00000
}