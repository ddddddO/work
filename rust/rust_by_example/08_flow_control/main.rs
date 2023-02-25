#![allow(unreachable_code)]

fn main() {
  ifelse_fn();
  loop_fn();
  label_fn();
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