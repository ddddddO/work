#[derive(Debug)]
struct Point {
  x: f32,
  y: f32,
}

struct Pair(i32, f32);

fn main() {
  let point: Point = Point { x: 10.3, y: 0.4 };

  // 構造体の更新記法を用いて、別の構造体のフィールドの値を基に
  // 新たなpointを生成
  let bottom_right = Point { x: 5.2, ..point };
  println!("{:?}", bottom_right);
  // Point { x: 5.2, y: 0.4 }

  let pair = Pair(1, 0.1);
  // タプルをデストラクト
  let Pair(integer, decimal) = pair;
  println!("pair contains {:?} and {:?}", integer, decimal);
  // pair contains 1 and 0.1

  exercises_struct();
}

fn exercises_struct() {
  let rect = Rectangle{
    top_left: Point{x: 1.5, y: 12.3},
    bottom_right: Point{x: 20.4, y: 1.0},
  };
  println!("answer 1: {}", rect_area(rect));
  // answer 1: 213.56999

  let rect2 = Rectangle{
    top_left: Point{x: 1.5, y: 12.3},
    bottom_right: Point{x: 20.4, y: 1.0},
  };
  println!("answer 2: {}", square(rect2, 10.0));
  // answer 2: 76.5
}

struct Rectangle {
  top_left: Point,
  bottom_right: Point,
}

// 構造体のデストラクト
// ref: https://doc.rust-jp.rs/rust-by-example-ja/flow_control/match/destructuring/destructure_structures.html
fn rect_area(rect: Rectangle) -> f32 {
  // ネストした構造体のデストラクト。多分
  let Rectangle{
    top_left: Point{x: x_a, y: y_a},
    bottom_right: Point{x: x_b, y: y_b},
  } = rect;

  (x_b - x_a) * (y_a - y_b)
}

fn square(rect: Rectangle, len: f32) -> f32 {
  let Rectangle{
    top_left: Point{x: _x, ..},
    bottom_right: Point{y: _y, ..},
  } = rect;

  (len - _x) * (len - _y)
}