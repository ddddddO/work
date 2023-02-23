fn main() {
  let long_lived_binding = 1;
  println!{"long: {}", long_lived_binding};
  // long: 1

  {
    // シャドーイング
    let long_lived_binding = 5_f32;
    println!("inner long: {}", long_lived_binding);
    // inner long: 5
  }

  println!("outer block long: {}", long_lived_binding);
  // outer block long: 1

  // シャドーイング
  let long_lived_binding = 'a';
  println!("end long: {}", long_lived_binding);
  // end long: a

  let a_binding;
  a_binding = "aaaa";
  println!("{}", a_binding);
  // aaaa

  freeze();
}

// データを同じ名前のイミュータブルな変数に束縛しなおすと、データは凍結されます。
// 凍結したデータは、イミュータブルな束縛がスコープ外になるまで変更できません。
fn freeze() {
  let mut _mutable_integer = 7i32;
  {
    // イミュータブルな`_mutable_integer`でシャドーイングする
    let _mutable_integer = _mutable_integer;

    // エラー! `_mutable_integer`はこのスコープでは凍結している。
    // _mutable_integer = 50;
  }

  // OK! `_mutable_integer`はこのスコープでは凍結していない。
  _mutable_integer = 3;
  println!("{}", _mutable_integer);
  // 3
}