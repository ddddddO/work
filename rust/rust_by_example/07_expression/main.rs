fn main() {
  let x = 5u32;

  let y = {
      let x_squared = x * x;
      let x_cube = x_squared * x;

      // This expression will be assigned to `y`
      // この式は`y`に代入されます。
      x_cube + x_squared + x
  };

  let z = {
      // The semicolon suppresses this expression and `()` is assigned to `z`
      // セミコロンがあるので`z`には`()`が入ります。
      2 * x;
  };

  println!("x is {:?}", x);
  // x is 5
  println!("y is {:?}", y);
  // y is 155
  println!("z is {:?}", z);
  // z is ()
}