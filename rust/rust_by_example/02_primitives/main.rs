use std::fmt;
use std::mem;

struct Matrix(f32, f32, f32, f32);
impl fmt::Display for Matrix {
  fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
    write!(f, "( {} {} )\n( {} {} )", self.0, self.1, self.2, self.3)
  }
}

fn transpose(origin: Matrix) -> Matrix {
  Matrix(origin.0, origin.2, origin.1, origin.3)
}

// この関数はスライスを借用する
fn analyze_slice(slice: &[i32]) {
  println!("First element of the slice: {}", slice[0]);
  println!("The slice has {} elements", slice.len());
}

fn main() {
  let matrix = Matrix(1.1, 1.2, 2.1, 2.2);
  println!("Matrix:\n{}", matrix);
  // Matrix:
  // ( 1.1 1.2 )
  // ( 2.1 2.2 )
  println!("Transpose:\n{}", transpose(matrix));
  // Transpose:
  // ( 1.1 2.1 )
  // ( 1.2 2.2 )

  /*
  * 配列: 長さはコンパイル時には決定されていて、[T; length]という形で指定できます。
  * スライス: コンパイル時に長さが決定されていません。スライスは配列の一部を借用するのに使用され、&[T]という型シグネチャを持ちます。
  */
  let xs: [i32; 5] = [1, 2, 3, 4, 5];
  // 配列はスタック上に置かれる
  println!("Array occupies {} bytes", mem::size_of_val(&xs));
  // Array occupies 20 bytes

  // 配列は自動的にスライスとして借用される。
  println!("Borrow the whole array as a slice.");
  analyze_slice(&xs);
  // Borrow the whole array as a slice.
  // First element of the slice: 1
  // The slice has 5 elements

  let ys: [i32; 500] = [0; 500];
  // スライスは配列の一部を指すことができる。
  // [starting_index..ending_index] の形をとり、
  // `starting_index` はスライスの先頭の位置を表し、
  // `ending_index` はスライスの末尾の1つ先の位置を表す。
  println!("Borrow a section of the array as a slice.");
  analyze_slice(&ys[1 .. 4]);
  // Borrow a section of the array as a slice.
  // First element of the slice: 0
  // The slice has 3 elements
}