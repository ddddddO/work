#![allow(unreachable_code)]

// TODO: 以下以降の8章もう一度目を通していいかも
// https://doc.rust-jp.rs/rust-by-example-ja/flow_control/match.html
fn main() {
  ifelse_fn();
  loop_fn();
  label_fn();
  ret_loop_fn();
  for_fn();
}

fn ifelse_fn() {
  let n = 5;

  let big_n =
    if -10 < n && n < 10 {
      println!("a");
      10 * n
    } else {
      println!("b");
      n / 2
    };
  
  println!("{} -> {}", n, big_n);
  // a
  // 5 -> 50
}

fn loop_fn() {
  let mut count = 0u32;

  loop {
    count += 1;
    if count == 3 {
      println!("three");
      continue;
    }

    println!("count: {}", count);

    if count == 5 {
      println!("end!");
      break;
    }
  }
  // count: 1
  // count: 2
  // three
  // count: 4
  // count: 5
  // end!
}

fn label_fn() {
  'outer: loop {
    println!("outer loop");

    'inner: loop {
      println!("inner loop");
      // break: // 内側のループを中断
      break 'outer;
    }
    println!("This point will never be reached");
  }
  println!("end");
  // outer loop
  // inner loop
  // end
}

fn ret_loop_fn() {
  let mut count = 0;
  let ret = loop {
    count += 1;

    if count == 10 {
      break count * 2;
    }
  };
  println!("count: {}", ret);
  // count: 20
}

fn for_fn() {
  for n in 0..3 {
    println!("a: {}", n);
  }

  for n in 0..=3 {
    println!("b: {}", n);
  }

  // a: 0
  // a: 1
  // a: 2
  // b: 0
  // b: 1
  // b: 2
  // b: 3
}