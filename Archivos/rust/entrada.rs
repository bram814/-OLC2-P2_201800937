fn main() {
  // Expresion loop
  let mut cont = 0;
  let result = loop {
    cont = cont + 1;
    if cont == 10 {
      break cont * 2;
    }
  };
// ¡No olvidar el punto y coma!
  println!("{}", result);
}