fn main() {
   let x: i64 = 56;
   let y = if x == 55 { "hol" } else if x == 56{"a"} else if x == 57{"b"} else if x == 58{"c"} else { "hola" };
   println!("{}", y);
}
