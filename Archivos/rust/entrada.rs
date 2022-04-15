fn main() {
  let x = 5;
  let numCadena = match x {
    1 | 2 => {"a"}
    3 => "b",
    4 => "c",
    5 | 6 => "d",
    235 => "e",
  _ => "f",
  };
  
  println!("{}",numCadena);
}